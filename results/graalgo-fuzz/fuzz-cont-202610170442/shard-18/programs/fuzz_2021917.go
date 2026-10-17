package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int; F2 string }
type GenS3 struct { F0 int8; F1 int8 }
func genFn0(p0 uint) uint {
    fmt.Println(p0)
    if true {
        if ("hello" >= "") {
            p0 = (p0 | (p0 / (((p0 | p0)) | 1)))
            switch (((("foo bar" + "foo bar") + ("foo bar" + (((("hello" + "foo bar") + "\n\t") + "") + ("" + "")))) + ("a" + ("" + "a"))) + "") {
            case "":
                p0 = (p0 >> (((p0 % (((p0 << ((p0) & 63))) | 1))) & 63))
            }
            var v6 uint64 = 18446744073709551615
            _ = v6
        } else {
            fmt.Println(p0)
        }
        v7 := GenS1{F0: 2147483647}
        p0 = (^(p0))
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        switch fi8 {
        case -2147483648:
            var v9 bool = (true || (!(false && false)))
            _ = v9
        }
        fmt.Println(fi8)
        _ = fi8
    }
    var v10 uint8 = 0
    _ = p0
    _ = v10
    return 0
}

func main() {
    var v1 int16 = 42
    var v2 int16 = (v1 & v1)
    var v3 string = "\n\t"
    var v4 int32 = 2147483647
    v5 := make(map[string]int)
    v11 := GenS2{F0: v2, F1: -9223372036854775808, F2: v3}
    v12 := v5[("\n\t" + "foo bar")]
    v13 := GenS3{F0: 127, F1: 1}
    v5[("a" + v3)] = v12
    v2 += 72
    if ((v2 != (-(v1))) && ((((1.5 * (float32(1e10) + 1.5)) / ((float32((((0.1 + ((0.1 + (-2.25 + (0.1 / 1.5))) + (float64(1.5) - 1e10))) - 0.1) + (1.5 - (-2.25 / 3.14159)))) / 1.5) * 3.14159)) * 1e10) > float32(((3.14159 / (0.1 + (0.0 - (3.14159 + (1e10 / (1e10 - (-2.25 / ((0.0 + 0.1) - ((0.0 + -2.25) * (3.14159 * 0.0)))))))))) * 1e10)))) {
        v14 := []int{}
        _ = v14
    }
    if (0 != 42) {
        v15 := []int{}
        switch v3 {
        case "a":
            if (-32768 == (^(v2))) {
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                v16 := GenS3{F0: -118, F1: 15}
                _ = v16
            }
        }
        fmt.Println(v4, v1, v3)
        _ = v15
    } else {
        for fi17 := 0; fi17 < 17; fi17++ {
            if true {
                v3 = v3
                if (!true) {
                    v18 := GenS1{F0: 965}
                    _ = v18
                } else {
                    v4 += -1
                }
            } else {
                v19 := []int{v12, 0, -935, (v12 * 0)}
                var v20 int64 = 42
                _ = v19
                _ = v20
            }
            fmt.Println(v3, v12, v3)
            if false {
                switch "\n\t" {
                case "hello":
                    var v21 int8 = -26
                    _ = v21
                case "\n\t":
                    for fi22 := 0; fi22 < 11; fi22++ {
                        var v23 uint = 4294967295
                        _ = fi22
                        _ = v23
                    }
                default:
                    for fi24 := 0; fi24 < 12; fi24++ {
                        var v25 uint32 = 4294967295
                        v4 = v4
                        _ = fi24
                        _ = v25
                    }
                }
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 int16 = (v1 + (v1 * v1))
                    v1 += (v2 ^ (v2 % ((v2) | 1)))
                    var v28 int32 = (v4 | v4)
                    var v29 int8 = 0
                    fmt.Println(v3, v4, fi26)
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
            }
            _ = fi17
        }
    }
    switch "" {
    case "\n\t":
        v12 -= v12
    case "foo bar":
        var v30 float32 = 1.5
        _ = v30
    case "":
        v31 := v5[v3]
        _ = v31
    default:
        for fi32 := 0; fi32 < 20; fi32++ {
            switch v12 {
            case -317:
                v33 := GenS3{F0: 100, F1: 42}
                _ = v33
            default:
                for fi34 := 0; fi34 < 15; fi34++ {
                    v2 = v1
                    _ = fi34
                }
            }
            v35 := []int{(^(fi32)), (fi32 - (^(fi32))), 279}
            _ = fi32
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
