go-tiasound
===========

A Go port of Ron Fries' TIA Chip Sound Simulator (tiasound.c).

Usage
-----

Create a new TiaSound instance. The first parameter is the desired
sample frequency, and the second is the TIA sound clock, twice per
scanline on a real system:

    tia := tiasound.NewTiaSound(44100, 31400)

After that, registers may be set:

    tia.Update(tiasound.AUDV0, 0xf)
    tia.Update(tiasound.AUDF0, 0xf)
    tia.Update(tiasound.AUDC0, 6)

Fetch the next output sample with GetSample:

    sample := float32(tia.GetSample())/128 - 0.5

... or you can fill a buffer with samples:

    buf := make([]uint8, 256)
    tia.FillBuffer(buf)

License
-------

The original GPLv2 license applies (see LICENSE).
