package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 bool }
func genFn0(p0 uint, p1 uint64, p2 uint64) int16 {
    var v5 uint16 = 32767
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 262
}
func genFn1(p0 uint64) int64 {
    for fi6 := 0; fi6 < 4; fi6++ {
        p0 = (p0 % ((p0) | 1))
        _ = fi6
    }
    fmt.Println(p0)
    _ = p0
    return 9223372036854775807
}

func main() {
    var v1 int32 = 0
    var v2 float64 = -2.25
    v3 := []string{"foo bar", "", ""}
    var v4 int32 = (^(v1))
    fmt.Println(v2, v2, v2)
    if true {
        var v7 uint64 = 2
        if (!(563 > 2147483647)) {
            var v8 uint16 = 935
            _ = v8
        }
        if false {
            var v9 int = -9223372036854775808
            _ = v9
        }
        _ = v7
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v10 uint8 = 127
        _ = v10
    }
    var v11 int32 = -377
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 bool = (32767 <= (fi12 % (((fi12 | (fi12 << ((42) & 63)))) | 1)))
        var v14 string = (("" + "foo bar") + ("" + "foo bar"))
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint32 = 185
    switch "" {
    case "hello":
        var v16 uint = 0
        _ = v16
    case "\n\t":
        if (1 != 42) {
            var v17 int64 = genFn1(42)
            if false {
                v17 = genFn1(342)
                switch "a" {
                case "foo bar":
                    switch 32768 {
                    case -247:
                        v2 = float64((float32(v2) + (0.0 + (3.14159 - float32(((v2 + ((v2 - 1.5) + float64(1e10))) + 1.5))))))
                    case -686:
                        switch 4294967295 {
                        case 253:
                            fmt.Println(v4, v4, v11)
                        case 128:
                            var v18 float32 = -2.25
                            _ = v18
                        default:
                            for fi19 := 0; fi19 < 6; fi19++ {
                                var v20 uint16 = 311
                                fmt.Println(v2)
                                for fi21 := 0; fi21 < 19; fi21++ {
                                    v4 -= v4
                                    fmt.Println(v15, fi19)
                                    _ = fi21
                                }
                                _ = fi19
                                _ = v20
                            }
                        }
                    }
                }
            } else {
                switch -230 {
                case 1:
                    var v22 float32 = 3.14159
                    _ = v22
                }
                if len(v3) > 0 {
                    v3[0] = ""
                }
            }
            _ = v17
        }
    }
    if false {
        v23 := GenS1{F0: 42, F1: (^(v15)), F2: true}
        if true {
            v3 = append(v3, "a")
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 int16 = 1
                v4 += v4
                fmt.Println(v15, v2, v11)
                _ = fi24
                _ = v25
            }
        } else {
            v1 -= (v4 / (((^(v1))) | 1))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (42 >= 1) {
                v26 := GenS1{F0: 42, F1: v15, F2: false}
                _ = v26
            }
        }
        _ = len(v3)
        _ = v23
    } else {
        v4 = v11
    }
    v27 := GenS1{F0: 65535, F1: (^(v15)), F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
