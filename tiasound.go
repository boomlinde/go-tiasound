package tiasound

const SET_TO_1 = 0x00
const POLY4 = 0x01
const DIV31_POLY4 = 0x02
const POLY5_POLY4 = 0x03
const PURE = 0x04
const PURE2 = 0x05
const DIV31_PURE = 0x06
const POLY5_2 = 0x07
const POLY9 = 0x08
const POLY5 = 0x09
const DIV31_POLY5 = 0x0a
const POLY5_POLY5 = 0x0b
const DIV3_PURE = 0x0c
const DIV3_PURE2 = 0x0d
const DIV93_PURE = 0x0e
const DIV3_POLY5 = 0x0f

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

const CHAN1 = 0
const CHAN2 = 1

var bit4 = [POLY4_SIZE]uint8{1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0}
var bit5 = [POLY5_SIZE]uint8{
	0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0,
	1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1,
}
var bit9 = [POLY9_SIZE]uint8{
	0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0,
	0, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1,
	0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1,
	0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0,
	0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0,
	1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0,
	0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1,
	0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0,
	0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0,
	1, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 0,
	0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0,
	0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1,
	0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 0,
	1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0,
	1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0,
	0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0,
}
var div31 = [POLY5_SIZE]uint8{
	0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

type TiaSound struct {
	AUDC       [2]uint8
	AUDF       [2]uint8
	AUDV       [2]uint8
	outvol     [2]uint8
	p4         [2]uint8
	p5         [2]uint8
	p9         [2]uint16
	div_n_cnt  [2]uint8
	div_n_max  [2]uint8
	samp_n_max uint16
	samp_n_cnt uint16
}

// NewTiaSound initializes a TiaSound structure and returns its pointer.
func NewTiaSound(sample_freq, playback_freq int) *TiaSound {
	t := &TiaSound{}

	t.samp_n_max = uint16((sample_freq << 8) / playback_freq)
	t.samp_n_cnt = 0

	for ch := CHAN1; ch <= CHAN2; ch++ {
		t.outvol[ch] = 0
		t.div_n_cnt[ch] = 0
		t.div_n_max[ch] = 0
		t.AUDC[ch] = 0
		t.AUDF[ch] = 0
		t.AUDV[ch] = 0
		t.p4[ch] = 0
		t.p5[ch] = 0
		t.p9[ch] = 0
	}

	return t
}

// Update updates a TIA register with a new value and adjusts the emulation accordingly
func (t *TiaSound) Update(addr uint16, val uint8) {
	var new_val uint16
	var ch uint8

	switch addr {
	case AUDC0:
		t.AUDC[0] = val & 0x0f
		ch = 0
	case AUDC1:
		t.AUDC[1] = val & 0x0f
		ch = 1
	case AUDF0:
		t.AUDF[0] = val & 0x1f
		ch = 0
	case AUDF1:
		t.AUDF[1] = val & 0x1f
		ch = 1
	case AUDV0:
		t.AUDV[0] = (val & 0x0f) << 3
		ch = 0
	case AUDV1:
		t.AUDV[1] = (val & 0x0f) << 3
		ch = 1
	default:
		ch = 255
	}

	if ch != 255 {
		if t.AUDC[ch] == SET_TO_1 {
			new_val = 0
			t.outvol[ch] = t.AUDV[ch]
		} else {
			new_val = uint16(t.AUDF[ch] + 1)
			if (t.AUDC[ch] & DIV3_MASK) == DIV3_MASK {
				new_val *= 3
			}
		}

		if new_val != uint16(t.div_n_max[ch]) {
			t.div_n_max[ch] = uint8(new_val)
			if (t.div_n_cnt[ch] == 0) || (new_val == 0) {
				t.div_n_cnt[ch] = uint8(new_val)
			}
		}
	}
}

// GetSample advances the TIA emulation and returns the next sample.
func (t *TiaSound) GetSample() uint8 {
	for {
		for ch := CHAN1; ch <= CHAN2; ch++ {
			if t.div_n_cnt[ch] > 1 {
				t.div_n_cnt[ch]--
			} else if t.div_n_cnt[ch] == 1 {
				t.div_n_cnt[ch] = t.div_n_max[ch]

				t.p5[ch]++
				if t.p5[ch] == POLY5_SIZE {
					t.p5[ch] = 0
				}

				if ((t.AUDC[ch] & 0x02) == 0) ||
					(((t.AUDC[ch] & 0x01) == 0) && (div31[t.p5[ch]] != 0)) ||
					(((t.AUDC[ch] & 0x01) == 1) && (bit5[t.p5[ch]] != 0)) {
					if t.AUDC[ch]&0x04 != 0 {
						if t.outvol[ch] != 0 {
							t.outvol[ch] = 0
						} else {
							t.outvol[ch] = t.AUDV[ch]
						}
					} else if t.AUDC[ch]&0x08 != 0 {
						if t.AUDC[ch] == POLY9 {
							t.p9[ch]++
							if t.p9[ch] == POLY9_SIZE {
								t.p9[ch] = 0
							}

							if bit9[t.p9[ch]] != 0 {
								t.outvol[ch] = t.AUDV[ch]
							} else {
								t.outvol[ch] = 0
							}
						} else {
							if bit5[t.p5[ch]] != 0 {
								t.outvol[ch] = t.AUDV[ch]
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
							t.outvol[ch] = t.AUDV[ch]
						} else {
							t.outvol[ch] = 0
						}
					}
				}
			}
		}

		t.samp_n_cnt -= 256
		if t.samp_n_cnt < 256 {
			t.samp_n_cnt += t.samp_n_max
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