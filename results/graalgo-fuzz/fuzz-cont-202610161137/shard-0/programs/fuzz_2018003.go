package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int8, p1 uint32) uint64 {
    v4 := []int{0, -2147483648, 0, 0}
    var v5 bool = (false || (float32((float64(1e10) * -2.25)) != ((float32((-2.25 / (3.14159 + 3.14159))) - 1e10) * 3.14159)))
    switch ("a" + "foo bar") {
    case "\n\t":
        switch -824 {
        case -384:
            v6 := GenS1{F0: 718, F1: 9223372036854775807}
            _ = v6
        case 0:
            if (false && true) {
                for fi7 := 0; fi7 < 6; fi7++ {
                    v8 := GenS1{F0: 2, F1: -9223372036854775808}
                    _ = fi7
                    _ = v8
                }
                switch "" {
                case "hello":
                    var v9 uint16 = 2
                    _ = v9
                case "foo bar":
                    fmt.Println(p0, p1, p1)
                default:
                    fmt.Println(p0)
                }
                v10 := GenS2{F0: p1, F1: 32767, F2: 65535}
                _ = v10
            } else {
                v5 = v5
                fmt.Println(v5, p0, p1)
            }
        case 9223372036854775807:
            for fi11 := 0; fi11 < 7; fi11++ {
                fmt.Println(fi11)
                _ = fi11
            }
        default:
            v12 := GenS3{F0: -2147483648}
            _ = v12
        }
    case "hello":
        var v13 uint = 2
        _ = v13
    case "foo bar":
        switch 0 {
        case 127:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 32768:
            v14 := GenS2{F0: (p1 / ((p1) | 1)), F1: 1000, F2: p1}
            _ = v14
        }
    default:
        if false {
            var v15 int64 = 1
            var v16 uint32 = p1
            _ = v15
            _ = v16
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch 32768 {
            case -747:
                if (!(!((v5 || v5) || (("a" + "\n\t") < "")))) {
                    p0 -= p0
                    var v17 uint8 = 201
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 15; fi18++ {
                        var v19 bool = ((p1 >> ((65535) & 63)) < (p1 % (((p1 - 1)) | 1)))
                        var v20 uint = 2
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
            case -290:
                p1 -= (^(p1))
            default:
                switch 347 {
                case -534:
                    var v21 uint8 = 0
                    _ = v21
                default:
                    p1 = (p1 / ((695) | 1))
                }
            }
            var v22 int = 792
            _ = v22
        }
    }
    var v23 uint8 = 1
    _ = p0
    _ = p1
    _ = v23
    _ = v4
    _ = v5
    return uint64(p1)
}
func genFn1(p0 int64, p1 uint16, p2 int64) uint32 {
    v24 := []int{622, 32767, 32767}
    p2 = (p2 % (((p2 % ((p2) | 1))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 2147483647
}

func main() {
    var v1 int8 = -1
    var v2 int32 = -2147483648
    v3 := []uint32{1, 42, 0}
    var v25 int16 = 0
    for fi26 := 0; fi26 < 9; fi26++ {
        switch "" {
        case "hello":
            var v27 int = (fi26 | (fi26 | -2147483648))
            _ = v27
        case "a":
            fmt.Println(v25, v1)
        default:
            v2 = v2
        }
        if false {
            v28 := GenS3{F0: (v2 ^ (v2 + (v2 | -2147483648)))}
            v25 += -211
            _ = v28
        }
        _ = fi26
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        if (fi29 == 128) {
            for fi30 := 0; fi30 < 12; fi30++ {
                var v31 int16 = (v25 * v25)
                fmt.Println(v31, v31, v31)
                _ = fi30
                _ = v31
            }
            fmt.Println(v1)
        }
        v32 := GenS3{F0: (v2 - (v2 ^ -2147483648))}
        switch (fi29 * fi29) {
        case 1:
            var v33 uint32 = genFn1(-1, 2, 42)
            _ = v33
        case 193:
            for fi34 := 0; fi34 < 17; fi34++ {
                var v35 uint = 65535
                var v36 string = "\n\t"
                _ = fi34
                _ = v35
                _ = v36
            }
        }
        _ = fi29
        _ = v32
    }
    var v37 uint = 0
    v38 := GenS3{F0: (v2 << ((v37) & 63))}
    var v39 string = "hello"
    for fi40 := 0; fi40 < 16; fi40++ {
        switch v39 {
        case "hello":
            fmt.Println(v37, v2, v2)
        case "foo bar":
            v41 := GenS3{F0: (v2 >> ((v37) & 63))}
            _ = v41
        case "\n\t":
            v25 += int16(v37)
        default:
            var v42 string = v39
            _ = v42
        }
        v3 = append(v3, 1)
        for fi43 := 0; fi43 < 6; fi43++ {
            var v44 bool = true
            _ = fi43
            _ = v44
        }
        _ = fi40
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v37
    _ = v38
    _ = v39
}
