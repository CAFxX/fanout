package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint32 }
func genFn0() uint8 {
    var v7 float64 = ((0.0 * (3.14159 + 0.1)) * -2.25)
    v8 := GenS1{F0: 617, F1: 1}
    v7 = v7
    _ = v7
    _ = v8
    return 228
}

func main() {
    var v1 int8 = 127
    var v2 float32 = 3.14159
    var v3 uint16 = 2
    v4 := []int{507, 9223372036854775807, 32768, -248}
    var v5 uint = 2
    var v6 float32 = -2.25
    var v9 uint64 = 42
    fmt.Println(v6, v3)
    var v10 uint = 525
    v11 := GenS1{F0: 0, F1: -859}
    var v12 uint8 = 2
    var v13 int16 = 1
    var v14 uint8 = genFn0()
    if (!(false && false)) {
        fmt.Println(v6, v12)
        var v15 int16 = (v13 << (((v5 * 65535)) & 63))
        switch 128 {
        case 794:
            v16 := GenS1{F0: 2147483648, F1: (^(v15))}
            _ = v16
        case 128:
            if true {
                if len(v4) > 0 {
                    v4[0] = -1
                }
                v13 += (v13 ^ (v15 % (((v13 * (v13 ^ v13))) | 1)))
                v9 -= (v9 % ((42) | 1))
            }
        case 765:
            var v17 bool = (v6 == v6)
            _ = v17
        }
        _ = v15
    } else {
        switch ("hello" + ("foo bar" + "a")) {
        case "foo bar":
            for fi18 := 0; fi18 < 10; fi18++ {
                v14 -= (v12 % ((v12) | 1))
                var v19 int8 = (v1 - (v1 % ((v1) | 1)))
                for fi20 := 0; fi20 < 20; fi20++ {
                    switch (fi20 << ((v10) & 63)) {
                    case 9223372036854775807:
                        if len(v4) > 0 {
                            v4[0] = (-(fi18))
                        }
                    case 128:
                        v21 := GenS1{F0: 449, F1: v13}
                        _ = v21
                    default:
                        var v22 float32 = 0.1
                        _ = v22
                    }
                    _ = fi20
                }
                _ = fi18
                _ = v19
            }
        case "hello":
            if (4294967295 <= 605) {
                switch "foo bar" {
                case "\n\t":
                    var v23 int32 = 42
                    _ = v23
                case "a":
                    fmt.Println(v1, v14, v5)
                case "foo bar":
                    var v24 int = 32768
                    _ = v24
                }
                if true {
                    v9 -= 287
                }
                v25 := GenS1{F0: 84, F1: 32767}
                _ = v25
            }
        case "":
            var v26 int16 = -1
            _ = v26
        default:
            if (!(-9223372036854775808 <= 9223372036854775807)) {
                if ((v2 / ((v6 / (-2.25 - ((3.14159 / ((3.14159 * (((v2 + 0.0) - v6) / 1.5)) / (float32((-2.25 / 1e10)) / float32(-2.25)))) / v2))) * 1e10)) == 1e10) {
                    switch 32767 {
                    case 128:
                        _ = len(v4)
                    case 32768:
                        var v27 float64 = (float64(v2) / 3.14159)
                        _ = v27
                    default:
                        v9 = 18446744073709551615
                    }
                    switch ("" + "a") {
                    case "a":
                        var v28 int16 = (v13 | (v13 >> ((v5) & 63)))
                        _ = v28
                    default:
                        if true {
                            fmt.Println(v5, v3, v14)
                        } else {
                            var v29 string = ""
                            _ = v29
                        }
                    }
                    var v30 uint64 = 18446744073709551615
                    _ = v30
                }
                switch "a" {
                case "a":
                    if len(v4) > 0 {
                        v4[0] = 32768
                    }
                case "":
                    v4 = append(v4, 9223372036854775807)
                default:
                    v31 := GenS1{F0: 236, F1: v13}
                    _ = v31
                }
                v13 -= v13
            }
        }
    }
    var v32 int64 = 838
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v33 bool = true
    v3 += v3
    v4 = append(v4, 0)
    switch -86 {
    case -1:
        if ((-(v32)) <= v32) {
            switch -9223372036854775808 {
            case -676:
                v34 := GenS2{F0: (v12 ^ (v14 + genFn0())), F1: v9, F2: v12}
                _ = v34
            case 191:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 42:
                var v35 float32 = 0.1
                _ = v35
            }
            var v36 uint = (^(v10))
            v37 := GenS3{F0: 2147483647}
            _ = v36
            _ = v37
        }
    }
    var v38 float32 = v2
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
