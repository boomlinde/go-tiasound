package tiasound

const SET_TO_1 = 0x00
const POLY9 = 0x08
const DIV3_MASK = 0x0c

const AUDC0 = 0x15
const AUDC1 = 0x16
const AUDF0 = 0x17
const AUDF1 = 0x18
const AUDV0 = 0x19
const AUDV1 = 0x1a

const POLY4_SIZE = 0x000f
const POLY5_SIZE = 0x001f
const POLY9_SIZE = 0x01ff

type TiaSound struct {
	audc    [2]uint8
	audf    [2]uint8
	audv    [2]uint8
	outvol  [2]uint8
	p4      [2]uint8
	p5      [2]uint8
	p9      [2]uint16
	divCnt  [2]uint8
	divMax  [2]uint8
	sampMax uint16
	sampCnt uint16
}

// NewTiaSound initializes a TiaSound structure and returns its pointer.
func NewTiaSound(sample_freq, playback_freq int) *TiaSound {
	return &TiaSound{sampMax: uint16((sample_freq << 8) / playback_freq)}
}

// Update updates a TIA register with a new value and adjusts the emulation accordingly
func (t *TiaSound) Update(addr uint16, val uint8) {
	var new_val uint16
	var ch uint8

	switch addr {
	case AUDC0:
		t.audc[0] = val & 0x0f
		ch = 0
	case AUDC1:
		t.audc[1] = val & 0x0f
		ch = 1
	case AUDF0:
		t.audf[0] = val & 0x1f
		ch = 0
	case AUDF1:
		t.audf[1] = val & 0x1f
		ch = 1
	case AUDV0:
		t.audv[0] = (val & 0x0f) << 3
		ch = 0
	case AUDV1:
		t.audv[1] = (val & 0x0f) << 3
		ch = 1
	default:
		ch = 255
	}

	if ch != 255 {
		if t.audc[ch] == SET_TO_1 {
			new_val = 0
			t.outvol[ch] = t.audv[ch]
		} else {
			new_val = uint16(t.audf[ch] + 1)
			if (t.audc[ch] & DIV3_MASK) == DIV3_MASK {
				new_val *= 3
			}
		}

		if new_val != uint16(t.divMax[ch]) {
			t.divMax[ch] = uint8(new_val)
			if (t.divCnt[ch] == 0) || (new_val == 0) {
				t.divCnt[ch] = uint8(new_val)
			}
		}
	}
}

// GetSample advances the TIA emulation and returns the next sample.
func (t *TiaSound) GetSample() uint8 {
	for {
		for ch := 0; ch < 2; ch++ {
			if t.divCnt[ch] > 1 {
				t.divCnt[ch]--
			} else if t.divCnt[ch] == 1 {
				t.divCnt[ch] = t.divMax[ch]

				t.p5[ch]++
				if t.p5[ch] == POLY5_SIZE {
					t.p5[ch] = 0
				}

				if ((t.audc[ch] & 0x02) == 0) ||
					(((t.audc[ch] & 0x01) == 0) && (div31[t.p5[ch]] != 0)) ||
					(((t.audc[ch] & 0x01) == 1) && (bit5[t.p5[ch]] != 0)) {
					if t.audc[ch]&0x04 != 0 {
						if t.outvol[ch] != 0 {
							t.outvol[ch] = 0
						} else {
							t.outvol[ch] = t.audv[ch]
						}
					} else if t.audc[ch]&0x08 != 0 {
						if t.audc[ch] == POLY9 {
							t.p9[ch]++
							if t.p9[ch] == POLY9_SIZE {
								t.p9[ch] = 0
							}

							if bit9[t.p9[ch]] != 0 {
								t.outvol[ch] = t.audv[ch]
							} else {
								t.outvol[ch] = 0
							}
						} else {
							if bit5[t.p5[ch]] != 0 {
								t.outvol[ch] = t.audv[ch]
							} else {
								t.outvol[ch] = 0
							}
						}
					} else {
						t.p4[ch]++
						if t.p4[ch] == POLY4_SIZE {
							t.p4[ch] = 0
						}

						if bit4[t.p4[ch]] != 0 {
							t.outvol[ch] = t.audv[ch]
						} else {
							t.outvol[ch] = 0
						}
					}
				}
			}
		}

		t.sampCnt -= 256
		if t.sampCnt < 256 {
			t.sampCnt += t.sampMax
			return t.outvol[0] + t.outvol[1]
		}
	}
}

// FillBuffer fills a buffer with TIA samples.
func (t *TiaSound) FillBuffer(buffer []uint8) {
	for n := range buffer {
		buffer[n] = t.GetSample()
	}
}
