package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int32 }

func main() {
    v1 := []string{("a" + ""), "a", "", "a"}
    v2 := []string{}
    v3 := []int{-493}
    v4 := make(map[int]string)
    var v5 uint16 = 32767
    var v6 uint = 921
    v5 -= (v5 / ((65535) | 1))
    var v7 string = ("\n\t" + "hello")
    v5 += 2
    fmt.Println(v6)
    v6 = v6
    var v8 int16 = 1
    for fi9 := 0; fi9 < 4; fi9++ {
        delete(v4, 658)
        _ = fi9
    }
    v7 = (v7 + (("" + "a") + "a"))
    if (!false) {
        if ((v5 >> ((v6) & 63)) == (v5 / ((2) | 1))) {
            v5 -= v5
            if false {
                var v10 string = v7
                switch 32767 {
                case -9223372036854775808:
                    for fi11 := 0; fi11 < 13; fi11++ {
                        v7 = ("foo bar" + "foo bar")
                        _ = fi11
                    }
                case 523:
                    for fi12 := 0; fi12 < 11; fi12++ {
                        var v13 uint = (v6 >> ((v6) & 63))
                        _ = fi12
                        _ = v13
                    }
                case -235:
                    var v14 uint32 = 807
                    _ = v14
                }
                v10 = v7
                _ = v10
            } else {
                v8 += v8
                if (!(((-128 > 100) && (127 > -955)) || (uint(v5) <= uint(v5)))) {
                    var v15 string = "hello"
                    for fi16 := 0; fi16 < 18; fi16++ {
                        var v17 uint8 = 127
                        v18 := GenS1{F0: v5, F1: true, F2: (fi16 | int(v17))}
                        var v19 float32 = (1e10 - 0.1)
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = v15
                } else {
                    v20 := GenS1{F0: (v5 - 413), F1: false, F2: 128}
                    switch v7 {
                    case "\n\t":
                        if (0.1 > ((3.14159 - (float64(3.14159) + (1.5 * (float64(0.0) / 3.14159)))) * 0.1)) {
                            var v21 uint = v6
                            switch "\n\t" {
                            case "\n\t":
                                v22 := GenS2{F0: 2, F1: -9223372036854775808, F2: 32767}
                                _ = v22
                            default:
                                switch "" {
                                case "a":
                                    var v23 bool = (42 >= v5)
                                    _ = v23
                                }
                            }
                            if len(v3) > 0 {
                                v3[0] = 128
                            }
                            _ = v21
                        } else {
                            v24 := GenS1{F0: v5, F1: false, F2: 897}
                            _ = v24
                        }
                    default:
                        var v25 int = -687
                        _ = v25
                    }
                    _ = v20
                }
                var v26 int = -2147483648
                _ = v26
            }
        } else {
            v27 := GenS3{F0: (^(v5)), F1: 360, F2: 2147483647}
            v6 += v6
            _ = v27
        }
    } else {
        fmt.Println(v6)
    }
    v6 = v6
    switch 0 {
    case -2147483648:
        v7 = v7
    case 4294967295:
        v28 := v4[-9223372036854775808]
        _ = v28
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 854 {
    case 128:
        v8 -= 1
    case -9223372036854775808:
        v29 := GenS2{F0: (^(v5)), F1: 42, F2: v8}
        _ = v29
    default:
        fmt.Println(v6)
    }
    v30 := GenS3{F0: v5, F1: 65535, F2: -341}
    var v31 int64 = -916
    var v32 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
