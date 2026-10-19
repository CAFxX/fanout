package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int64; F2 string }
type GenS3 struct { F0 int; F1 int64; F2 uint64 }

func main() {
    var v1 int32 = 1
    var v2 float32 = (float32(-2.25) * (3.14159 - (1.5 * ((1.5 * (-2.25 - 1.5)) - (1.5 + ((1e10 - 1e10) * (1.5 * 0.1)))))))
    v3 := make(map[uint32]bool)
    v4 := []string{"\n\t"}
    fmt.Println(v2)
    if false {
        for fi5 := 0; fi5 < 10; fi5++ {
            v1 = (v1 >> ((115) & 63))
            _ = fi5
        }
        if len(v4) > 0 {
            v4[0] = "a"
        }
    }
    v6 := GenS1{F0: (1 == 42)}
    for fi7 := 0; fi7 < 15; fi7++ {
        if ((fi7 * (fi7 - fi7)) != fi7) {
            var v8 int = 9223372036854775807
            if true {
                if (false || ("a" > "a")) {
                    var v9 uint64 = 18446744073709551615
                    _ = v9
                }
                v1 += (v1 ^ v1)
                var v10 bool = true
                _ = v10
            } else {
                switch -798 {
                case 109:
                    if false {
                        switch (("a" + ("\n\t" + "hello")) + "") {
                        case "\n\t":
                            for fi11 := 0; fi11 < 7; fi11++ {
                                var v12 string = ("" + "\n\t")
                                var v13 int32 = v1
                                if true {
                                    if (0.0 != 0.0) {
                                        var v14 uint = 255
                                        v4 = append(v4, (v12 + "\n\t"))
                                        delete(v3, 2147483648)
                                        _ = v14
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = v12
                                    }
                                    v13 = (v13 & v13)
                                    fmt.Println(v12, fi11)
                                    v1 = (v13 & (v13 << ((18446744073709551615) & 63)))
                                } else {
                                    if (127 < 133) {
                                        _ = len(v3)
                                        v1 = (v1 | 42)
                                    }
                                    v2 += ((3.14159 * v2) * float32(0.0))
                                    var v15 int16 = 32767
                                    if (-1 > -9223372036854775808) {
                                        var v16 uint16 = 255
                                        _ = v16
                                    }
                                    if false {
                                        v17 := GenS2{F0: (v15 | v15), F1: 958, F2: v12}
                                        v8 = (fi7 << ((65535) & 63))
                                        _ = v17
                                    }
                                    _ = v15
                                }
                                switch 0 {
                                case -744:
                                    if ((42 <= -360) && (-292 > 9223372036854775807)) {
                                        if true {
                                            var v18 string = (("hello" + "") + "")
                                            var v19 uint64 = 416
                                            v19 += v19
                                            _ = v18
                                            _ = v19
                                        } else {
                                            delete(v3, 42)
                                        }
                                    }
                                case 4294967295:
                                    var v20 int32 = (v13 | (-(v1)))
                                    _ = v20
                                }
                                switch 986 {
                                case 2147483647:
                                    var v21 int = fi7
                                    _ = v21
                                }
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                        case "hello":
                            var v22 int8 = -94
                            _ = v22
                        case "foo bar":
                            v2 = 1e10
                        }
                    } else {
                        v2 = (((v2 + v2) / (((v2 * -2.25) + (v2 - float32(3.14159))) - 0.0)) + (v2 * 1.5))
                        for fi23 := 0; fi23 < 7; fi23++ {
                            v8 = fi23
                            _ = fi23
                        }
                        switch (v8 | v8) {
                        case -2147483648:
                            var v24 int32 = (v1 + (v1 % (((^(v1))) | 1)))
                            _ = v24
                        default:
                            v8 = (fi7 + (v8 % ((fi7) | 1)))
                        }
                    }
                }
                v25 := GenS2{F0: 42, F1: 1, F2: "hello"}
                fmt.Println(v2)
                _ = v25
            }
            for fi26 := 0; fi26 < 6; fi26++ {
                v2 += (-2.25 / 1.0)
                _ = fi26
            }
            _ = v8
        }
        _ = fi7
    }
    var v27 uint64 = 1000
    v2 -= (v2 / v2)
    _ = len(v4)
    for fi28 := 0; fi28 < 15; fi28++ {
        v29 := GenS3{F0: (fi28 >> ((438) & 63)), F1: -9223372036854775808, F2: (v27 >> ((667) & 63))}
        v2 = v2
        v4 = append(v4, ("hello" + ("" + "\n\t")))
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v6
}
