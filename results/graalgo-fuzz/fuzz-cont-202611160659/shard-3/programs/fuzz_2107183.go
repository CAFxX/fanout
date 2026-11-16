package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint64 }
func genFn0(p0 int, p1 int64, p2 int32) uint64 {
    p2 += (p2 * 2147483647)
    var v7 int = 1
    p2 += p2
    p2 += 210
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 string = "a"
    var v2 uint64 = 283
    v3 := []uint32{208, 742}
    var v4 uint = 0
    var v5 string = (v1 + "\n\t")
    var v6 float64 = ((-2.25 + -2.25) + -2.25)
    _ = len(v3)
    v4 += v4
    switch "\n\t" {
    case "hello":
        v5 = ((v1 + v1) + ("\n\t" + "foo bar"))
    case "":
        v1 = v5
    }
    var v8 float32 = (1.5 / 3.14159)
    var v9 string = (v1 + v5)
    fmt.Println(v4, v2, v2)
    if false {
        fmt.Println(v2, v4, v6)
        var v10 int64 = -1
        _ = v10
    } else {
        switch "foo bar" {
        case "":
            v9 = v9
        default:
            switch 32768 {
            case 127:
                v11 := GenS1{F0: 1}
                _ = v11
            case 0:
                for fi12 := 0; fi12 < 11; fi12++ {
                    _ = len(v3)
                    var v13 int8 = -128
                    var v14 int16 = 315
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case 32767:
                for fi15 := 0; fi15 < 17; fi15++ {
                    var v16 string = v5
                    var v17 int32 = 0
                    if (42 >= (^(v4))) {
                        v18 := GenS1{F0: 0}
                        var v19 int64 = 0
                        _ = v18
                        _ = v19
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
        }
    }
    fmt.Println(v8, v4, v2)
    for fi20 := 0; fi20 < 6; fi20++ {
        if ((true && false) && true) {
            var v21 float64 = v6
            v1 = v1
            _ = v21
        } else {
            v6 -= ((((((((v6 * (0.0 + v6)) / (float64(float32(v6)) + ((0.1 / 0.1) * -2.25))) + v6) - v6) + v6) + v6) * v6) - v6)
        }
        if len(v3) > 0 {
            v3[0] = 2
        }
        _ = fi20
    }
    var v22 string = v9
    var v23 uint32 = 144
    if (!false) {
        switch ("foo bar" + v5) {
        case "foo bar":
            var v24 float64 = (float64(v8) / float64(((v8 - v8) + v8)))
            _ = v24
        default:
            if (42 < -32768) {
                if (((v2 % ((2) | 1)) > uint64(v23)) || false) {
                    switch 32767 {
                    case 32768:
                        for fi25 := 0; fi25 < 13; fi25++ {
                            var v26 int16 = -32768
                            var v27 uint = v4
                            var v28 float32 = (v8 / 1.0)
                            _ = fi25
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    }
                    v29 := GenS1{F0: 45}
                    _ = v29
                }
            } else {
                var v30 bool = (0.0 == v8)
                var v31 int32 = -777
                var v32 float32 = v8
                _ = v30
                _ = v31
                _ = v32
            }
        }
        if false {
            for fi33 := 0; fi33 < 18; fi33++ {
                var v34 bool = (!((v8 * v8) < (v8 + (float32(((v6 + ((-2.25 - v6) * 3.14159)) + (0.1 - (-2.25 - ((1e10 - 3.14159) / 0.1))))) * -2.25))))
                switch (fi33 & 0) {
                case 484:
                    var v35 uint32 = (v23 << ((v4) & 63))
                    _ = v35
                case 4294967295:
                    v36 := GenS2{F0: 116, F1: 133, F2: v2}
                    _ = v36
                default:
                    v37 := GenS1{F0: 2}
                    _ = v37
                }
                _ = fi33
                _ = v34
            }
            switch 1 {
            case 42:
                v8 += (0.1 - -2.25)
            default:
                var v38 int32 = 42
                _ = v38
            }
            switch v9 {
            case "a":
                for fi39 := 0; fi39 < 7; fi39++ {
                    v2 = 340
                    for fi40 := 0; fi40 < 14; fi40++ {
                        fmt.Println(v23, fi39, v9)
                        v3 = append(v3, v23)
                        _ = fi40
                    }
                    v41 := GenS2{F0: 0, F1: 42, F2: v2}
                    _ = fi39
                    _ = v41
                }
            default:
                v42 := GenS2{F0: 42, F1: 255, F2: v2}
                _ = v42
            }
        }
        var v43 uint64 = v2
        _ = v43
    }
    fmt.Println(v1, v9)
    v3 = append(v3, v23)
    _ = len(v3)
    switch (("" + "\n\t") + v1) {
    case "":
        v44 := GenS1{F0: 73}
        _ = v44
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = 42
        }
    case "a":
        var v45 int64 = 0
        _ = v45
    default:
        if (42 > 2147483647) {
            if false {
                v3 = append(v3, v23)
                fmt.Println(v2, v4)
            }
            if true {
                v46 := GenS1{F0: 255}
                _ = v46
            }
            v6 -= (((v6 + -2.25) - 0.1) + v6)
        } else {
            var v47 bool = (1 >= 0)
            if v47 {
                v48 := GenS3{F0: 42, F1: 1, F2: genFn0(32768, 0, 0)}
                switch -721 {
                case 4294967295:
                    if v47 {
                        var v49 int8 = int8(v2)
                        if len(v3) > 0 {
                            v3[0] = v23
                        }
                        switch -1 {
                        case -1:
                            switch 32768 {
                            case -1:
                                for fi50 := 0; fi50 < 19; fi50++ {
                                    v8 = v8
                                    _ = fi50
                                }
                            case -9223372036854775808:
                                v47 = ((float32(0.1) * (v8 + v8)) >= (float32(((0.1 - -2.25) + float64((v8 + v8)))) / ((v8 * -2.25) * ((float32(v6) / float32((((v6 + v6) + ((0.0 / ((0.1 * (v6 / float64((v8 - -2.25)))) / v6)) + float64(v8))) * v6))) / v8))))
                            }
                        case -9223372036854775808:
                            var v51 uint = 2
                            _ = v51
                        case 42:
                            v1 = v1
                        }
                        _ = v49
                    } else {
                        v52 := GenS2{F0: 128, F1: 830, F2: v2}
                        _ = v52
                    }
                default:
                    switch (("a" + v22) + v9) {
                    case "foo bar":
                        v1 = (v22 + v22)
                    case "a":
                        var v53 int64 = 472
                        _ = v53
                    }
                }
                _ = v48
            }
            switch (v22 + "hello") {
            case "hello":
                var v54 int8 = -59
                _ = v54
            }
            _ = v47
        }
    }
    v55 := GenS3{F0: 250, F1: 1, F2: (v2 + (v2 >> ((368) & 63)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v8
    _ = v9
}
