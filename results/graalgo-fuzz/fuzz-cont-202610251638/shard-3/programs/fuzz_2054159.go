package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint64 }
func genFn0(p0 uint8) uint32 {
    var v6 uint = 65535
    v7 := []int{32767, 32767}
    p0 += (p0 - 228)
    fmt.Println(v6, p0)
    fmt.Println(p0, v6)
    _ = p0
    _ = v6
    _ = v7
    return 249
}
func genFn1() uint16 {
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 float32 = 1.5
        _ = fi8
        _ = v9
    }
    if false {
        var v10 float32 = (3.14159 - 0.0)
        _ = v10
    } else {
        var v11 int32 = 42
        var v12 uint64 = 262
        v11 = -568
        _ = v11
        _ = v12
    }
    return 255
}

func main() {
    var v1 float64 = 0.1
    var v2 bool = false
    var v3 int32 = 0
    var v4 int = -9223372036854775808
    var v5 string = ("\n\t" + "foo bar")
    var v13 int8 = 72
    v14 := []int{}
    var v15 uint = 234
    v13 += (v13 >> ((42) & 63))
    v16 := GenS1{F0: 255, F1: v3, F2: 42}
    if v2 {
        for fi17 := 0; fi17 < 1; fi17++ {
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            var v18 uint32 = uint32(v15)
            var v19 uint64 = 4294967295
            _ = fi17
            _ = v18
            _ = v19
        }
        fmt.Println(v3, v13, v5)
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 int16 = 1
        switch v5 {
        case "\n\t":
            switch (fi20 & v4) {
            case -266:
                switch (fi20 | 519) {
                case -1:
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                default:
                    v1 -= (3.14159 * (v1 / -2.25))
                }
            case 1:
                v1 -= 0.1
            default:
                _ = len(v14)
            }
        case "hello":
            var v22 int = (-(fi20))
            _ = v22
        }
        _ = fi20
        _ = v21
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        v5 = (((v5 + v5) + "foo bar") + v5)
        _ = fi23
    }
    if v2 {
        if (!(492 > 0)) {
            for fi24 := 0; fi24 < 2; fi24++ {
                if len(v14) > 0 {
                    v14[0] = (-(v4))
                }
                var v25 uint8 = 2
                var v26 int = 4294967295
                _ = fi24
                _ = v25
                _ = v26
            }
            var v27 uint64 = 18446744073709551615
            switch (v4 << (((v15 | v15)) & 63)) {
            case -2147483648:
                fmt.Println(v2, v1, v27)
            case 32768:
                var v28 uint8 = 182
                _ = v28
            case -1:
                v4 = (v4 >> (((v15 | v15)) & 63))
            default:
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            }
            _ = v27
        } else {
            var v29 int16 = 32767
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = v29
        }
        var v30 int = 32768
        _ = v30
    } else {
        for fi31 := 0; fi31 < 17; fi31++ {
            fmt.Println(v2, v13, v3)
            v32 := GenS1{F0: 96, F1: v3, F2: 0}
            _ = fi31
            _ = v32
        }
        v14 = append(v14, -442)
        for fi33 := 0; fi33 < 20; fi33++ {
            v34 := GenS2{F0: 0}
            for fi35 := 0; fi35 < 14; fi35++ {
                if (157 == 42) {
                    v36 := GenS3{F0: 127}
                    _ = v36
                } else {
                    v1 += ((1e10 * 3.14159) * v1)
                }
                switch "\n\t" {
                case "foo bar":
                    var v37 float64 = ((v1 * (v1 * 3.14159)) - 3.14159)
                    v38 := GenS4{F0: 18446744073709551615, F1: 234, F2: 18446744073709551615}
                    fmt.Println(v3, v3)
                    _ = v37
                    _ = v38
                case "\n\t":
                    var v39 int = fi35
                    _ = v39
                default:
                    if (2 >= 4294967295) {
                        var v40 int8 = (v13 & -27)
                        if (2 >= 0) {
                            v2 = true
                            v41 := GenS4{F0: v15, F1: 1, F2: 18446744073709551615}
                            v2 = v2
                            _ = v41
                        } else {
                            var v42 int64 = -343
                            _ = v42
                        }
                        _ = v40
                    } else {
                        fmt.Println(fi35, fi33, v5)
                    }
                }
                if (!false) {
                    var v43 uint8 = 127
                    var v44 int16 = 0
                    _ = v43
                    _ = v44
                }
                var v45 uint16 = 38
                _ = fi35
                _ = v45
            }
            _ = fi33
            _ = v34
        }
    }
    if (((((v1 - v1) / (float64((((float32(v1) + 1.5) / ((-2.25 - 1e10) / 1.0)) + 1e10)) / 0.1)) / (1.5 + (float64(3.14159) / 1e10))) == v1) && true) {
        v46 := GenS3{F0: 226}
        _ = v46
    }
    v15 = v15
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    for fi47 := 0; fi47 < 14; fi47++ {
        if v2 {
            v4 += (fi47 & 0)
            _ = len(v14)
            for fi48 := 0; fi48 < 13; fi48++ {
                fmt.Println(v5, v2, v5)
                _ = fi48
            }
        }
        if len(v14) > 0 {
            v14[0] = fi47
        }
        v5 = v5
        _ = fi47
    }
    fmt.Println(v13)
    v13 = (v13 | v13)
    _ = len(v14)
    for fi49 := 0; fi49 < 13; fi49++ {
        fmt.Println(v3, v4)
        v1 = v1
        for fi50 := 0; fi50 < 17; fi50++ {
            var v51 string = v5
            v52 := GenS2{F0: -1}
            var v53 int32 = v3
            fmt.Println(fi49, v15)
            _ = fi50
            _ = v51
            _ = v52
            _ = v53
        }
        _ = fi49
    }
    if len(v14) > 0 {
        v14[0] = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
