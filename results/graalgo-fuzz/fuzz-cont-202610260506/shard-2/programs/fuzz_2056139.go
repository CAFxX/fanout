package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint32; F2 string }
type GenS4 struct { F0 uint32; F1 uint32; F2 int }

func main() {
    v1 := []string{"", "hello", "hello"}
    var v2 uint8 = 255
    var v3 int16 = 1000
    var v4 int32 = -2147483648
    var v5 int = 4294967295
    if false {
        switch 127 {
        case 127:
            _ = len(v1)
        case 9223372036854775807:
            if (!false) {
                var v6 int32 = (v4 * v4)
                if (-2.25 > ((0.1 * ((((0.1 * 3.14159) - (0.1 - -2.25)) + ((1.5 - 0.0) - 3.14159)) * 0.1)) / ((1e10 * 3.14159) / 3.14159))) {
                    switch "hello" {
                    case "":
                        v6 = (v4 & -53)
                    default:
                        var v7 int8 = 100
                        _ = v7
                    }
                    v4 -= 0
                } else {
                    var v8 int64 = -774
                    _ = v8
                }
                _ = v6
            }
        case -104:
            var v9 int = (v5 - v5)
            _ = v9
        default:
            var v10 string = ""
            _ = v10
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            v3 = (-(v3))
            v4 -= (v4 << ((42) & 63))
            _ = fi11
        }
    }
    fmt.Println(v3)
    v12 := GenS1{F0: 2, F1: 2147483647, F2: 2147483648}
    var v13 uint8 = 1
    var v14 string = "hello"
    switch v14 {
    case "hello":
        var v15 int16 = 1000
        _ = v15
    default:
        fmt.Println(v2)
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        if (v14 != v14) {
            fmt.Println(v14, v5, v4)
            for fi17 := 0; fi17 < 15; fi17++ {
                if ((v4 % (((^(v4))) | 1)) > (^(v4))) {
                    v4 -= (-(v4))
                    var v18 uint32 = 65535
                    var v19 uint32 = v18
                    _ = v18
                    _ = v19
                } else {
                    v1 = append(v1, ((v14 + "foo bar") + v14))
                    var v20 int16 = (v3 | (v3 >> ((18446744073709551615) & 63)))
                    v2 -= 0
                    _ = v20
                }
                switch "" {
                case "":
                    v13 = 127
                case "hello":
                    if (true && false) {
                        var v21 float32 = float32(-2.25)
                        v3 = -1
                        _ = v21
                    }
                case "\n\t":
                    v22 := GenS2{F0: 2}
                    _ = v22
                default:
                    var v23 int8 = 0
                    _ = v23
                }
                var v24 bool = ((v5 ^ v5) >= 530)
                var v25 uint64 = 18446744073709551615
                var v26 int = -87
                var v27 int64 = -9223372036854775808
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (32767 > (v3 ^ (v3 * 0))) {
                    _ = len(v1)
                } else {
                    var v28 uint8 = v13
                    if (!v24) {
                        var v29 int = fi16
                        _ = v29
                    }
                    var v30 int16 = v3
                    _ = v28
                    _ = v30
                }
                _ = fi17
                _ = v24
                _ = v25
                _ = v26
                _ = v27
            }
            v31 := GenS2{F0: 1}
            _ = v31
        } else {
            switch ((((v14 + "\n\t") + v14) + ("foo bar" + ("foo bar" + "foo bar"))) + v14) {
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                for fi32 := 0; fi32 < 2; fi32++ {
                    var v33 float64 = 0.1
                    fmt.Println(v4)
                    v2 -= 128
                    switch -997 {
                    case -814:
                        var v34 int16 = v3
                        _ = v34
                    case -484:
                        if (!(-442 == 1)) {
                            v33 -= v33
                            if (!false) {
                                v35 := GenS3{F0: (^(v3)), F1: 2, F2: ("" + (((v14 + "foo bar") + (v14 + ((((v14 + (v14 + (v14 + v14))) + v14) + v14) + v14))) + v14))}
                                _ = v35
                            }
                            fmt.Println(v4, v2)
                        }
                    case -982:
                        v5 += (fi16 ^ v5)
                    default:
                        var v36 uint64 = 42
                        _ = v36
                    }
                    v4 = v4
                    fmt.Println(v2, v33)
                    _ = fi32
                    _ = v33
                }
            }
            if (!(v14 >= "\n\t")) {
                v1 = append(v1, "\n\t")
                fmt.Println(v2, v2, v5)
            } else {
                var v37 int = v5
                _ = v37
            }
            v38 := GenS4{F0: 675, F1: 2147483648, F2: 32768}
            _ = v38
        }
        _ = fi16
    }
    switch ("foo bar" + (("\n\t" + "\n\t") + v14)) {
    case "\n\t":
        var v39 string = "hello"
        _ = v39
    }
    v14 = v14
    v1 = append(v1, (("\n\t" + v14) + v14))
    v13 = v2
    if false {
        var v40 int64 = 1
        v40 += v40
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
