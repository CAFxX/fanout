package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 string; F2 uint }
type GenS3 struct { F0 int8; F1 string }
func genFn0(p0 int16, p1 uint64, p2 uint32) int32 {
    if false {
        v7 := GenS1{F0: int(p2)}
        if true {
            p1 -= (^(p1))
        }
        p2 = 2147483648
        _ = v7
    } else {
        fmt.Println(p0, p0)
        fmt.Println(p1)
    }
    var v8 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return v8
}
func genFn1(p0 bool) uint16 {
    p0 = ("a" >= "\n\t")
    p0 = (p0 || (-1 != -1))
    _ = p0
    return 255
}

func main() {
    var v1 int16 = 42
    v2 := make(map[string]uint32)
    v3 := []uint32{2147483648, 42, 4294967295, 2147483648}
    var v4 uint = 255
    var v5 int8 = 0
    v6 := []string{}
    if (v4 <= v4) {
        if len(v3) > 0 {
            v3[0] = 764
        }
        switch "a" {
        case "hello":
            for fi9 := 0; fi9 < 3; fi9++ {
                switch (fi9 & fi9) {
                case -9223372036854775808:
                    var v10 string = (("\n\t" + "\n\t") + "\n\t")
                    _ = v10
                default:
                    v3 = append(v3, 38)
                }
                _ = fi9
            }
        case "foo bar":
            if (18446744073709551615 <= 929) {
                for fi11 := 0; fi11 < 19; fi11++ {
                    for fi12 := 0; fi12 < 10; fi12++ {
                        switch fi12 {
                        case 32767:
                            switch (fi11 >> (((^(v4))) & 63)) {
                            case -2147483648:
                                fmt.Println(v1)
                            default:
                                _ = len(v2)
                            }
                        }
                        v1 += v1
                        _ = fi12
                    }
                    _ = fi11
                }
            } else {
                v1 = v1
                if false {
                    v1 = v1
                }
                v13 := v2[("" + "\n\t")]
                _ = v13
            }
        case "":
            var v14 int = 32767
            _ = v14
        }
        fmt.Println(v4, v4)
    }
    var v15 string = "hello"
    fmt.Println(v5, v4)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v16 float32 = ((1.5 - (0.0 / 1.5)) / 0.1)
    if true {
        fmt.Println(v4, v15, v5)
        v15 = "a"
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 16; fi18++ {
                v16 -= ((0.0 / 0.1) + v16)
                v1 += v1
                _ = fi18
            }
            v15 = v15
            _ = fi17
        }
    } else {
        var v19 uint16 = 65535
        v20 := GenS2{F0: 1, F1: ("" + ("foo bar" + v15)), F2: 782}
        if true {
            v1 = v1
            for fi21 := 0; fi21 < 5; fi21++ {
                switch 42 {
                case 32767:
                    v16 -= v16
                case 32768:
                    if true {
                        switch "" {
                        case "a":
                            v22 := GenS1{F0: (fi21 - (^(fi21)))}
                            _ = v22
                        case "hello":
                            v23 := GenS3{F0: (v5 * (v5 - (-(v5)))), F1: v15}
                            _ = v23
                        case "\n\t":
                            var v24 uint64 = 18446744073709551615
                            _ = v24
                        default:
                            var v25 int8 = (v5 << ((473) & 63))
                            _ = v25
                        }
                        var v26 string = v15
                        v5 = (-(v5))
                        _ = v26
                    }
                }
                v19 += v19
                var v27 int64 = 0
                _ = fi21
                _ = v27
            }
        }
        _ = v19
        _ = v20
    }
    v28 := GenS2{F0: 0, F1: (("" + v15) + v15), F2: 4294967295}
    v29 := GenS3{F0: (v5 | int8(v4)), F1: (((v15 + v15) + v15) + v15)}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
