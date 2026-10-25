package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 float32 = float32(1e10)
    var v2 int64 = 9223372036854775807
    var v3 string = "a"
    var v4 bool = (true && (((0 <= 65535) || false) && (42 <= -32768)))
    var v5 uint64 = 0
    if (!(0.1 <= float32(((1.5 - 3.14159) + -2.25)))) {
        v6 := []int{42, 4294967295, -224, -1}
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int64 = (-(v2))
            var v9 float32 = v1
            if v4 {
                for fi10 := 0; fi10 < 18; fi10++ {
                    var v11 int = (fi10 + fi10)
                    v5 = (v5 * (^(v5)))
                    var v12 string = (v3 + ("\n\t" + ("\n\t" + "a")))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v13 := GenS1{F0: ("foo bar" + v3), F1: 127, F2: 42}
                if (212 > 220) {
                    var v14 float32 = v9
                    _ = v14
                } else {
                    var v15 uint = 255
                    if v4 {
                        var v16 float64 = (0.1 / 1e10)
                        _ = v16
                    } else {
                        switch fi7 {
                        case 0:
                            var v17 bool = false
                            _ = v17
                        case -343:
                            var v18 uint32 = 65535
                            _ = v18
                        case 128:
                            var v19 uint32 = 65535
                            _ = v19
                        }
                        if v4 {
                            switch -701 {
                            case 32767:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            default:
                                var v20 float32 = (((0.1 - v9) - (0.1 - (1.5 / v1))) * 0.1)
                                _ = v20
                            }
                            for fi21 := 0; fi21 < 3; fi21++ {
                                fmt.Println(fi21)
                                v22 := GenS2{F0: (1 != 18446744073709551615)}
                                v3 = "\n\t"
                                v9 += 0.0
                                v23 := GenS3{F0: 1, F1: (^(v5))}
                                var v24 int64 = (v2 - (v2 - 805))
                                _ = fi21
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                        }
                        var v25 int32 = -1
                        _ = v25
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = v15
                }
                _ = v13
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    } else {
        var v26 int8 = 30
        _ = v26
    }
    v27 := GenS4{F0: 2147483647}
    var v28 int = 32768
    if false {
        if (-798 == -214) {
            v28 += 32767
        } else {
            v29 := GenS3{F0: -139, F1: (v5 & v5)}
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 int16 = 32767
        v1 = 3.14159
        v32 := GenS4{F0: -212}
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println(v2, v4)
    var v33 int = v28
    for fi34 := 0; fi34 < 19; fi34++ {
        var v35 bool = v4
        for fi36 := 0; fi36 < 11; fi36++ {
            v5 -= (v5 * (v5 >> ((4294967295) & 63)))
            var v37 int64 = -248
            _ = fi36
            _ = v37
        }
        _ = fi34
        _ = v35
    }
    var v38 bool = v4
    switch (-(v28)) {
    case 617:
        var v39 string = v3
        _ = v39
    }
    for fi40 := 0; fi40 < 5; fi40++ {
        v28 += (v28 - (fi40 >> ((42) & 63)))
        fmt.Println(v3)
        _ = fi40
    }
    fmt.Println(v28)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v38
    _ = v4
    _ = v5
}
