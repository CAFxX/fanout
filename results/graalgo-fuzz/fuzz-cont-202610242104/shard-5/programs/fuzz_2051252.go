package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint8; F1 uint }
func genFn0(p0 bool, p1 int16) uint16 {
    var v6 string = "\n\t"
    var v7 int32 = 42
    v8 := []int{283, 32767, -124, 127}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1() uint64 {
    v9 := []int{127}
    if false {
        var v10 int = -9223372036854775808
        if (853 != 2147483647) {
            if false {
                switch "a" {
                case "hello":
                    _ = len(v9)
                case "\n\t":
                    switch v10 {
                    case 2147483647:
                        if (false && false) {
                            v11 := GenS1{F0: "a"}
                            v10 -= (v10 ^ -2147483648)
                            _ = v11
                        }
                    case -2147483648:
                        v12 := GenS1{F0: "foo bar"}
                        _ = v12
                    default:
                        v10 = v10
                    }
                case "foo bar":
                    var v13 uint32 = 2147483648
                    _ = v13
                }
            } else {
                v10 = v10
                v14 := GenS1{F0: "\n\t"}
                v10 = -1
                _ = v14
            }
            switch ("a" + (("a" + ("\n\t" + "hello")) + "foo bar")) {
            case "hello":
                switch v10 {
                case 378:
                    for fi15 := 0; fi15 < 18; fi15++ {
                        fmt.Println(v10, v10)
                        v10 -= (^(v10))
                        v10 = v10
                        _ = fi15
                    }
                }
            case "":
                v16 := GenS2{F0: 127}
                _ = v16
            default:
                for fi17 := 0; fi17 < 19; fi17++ {
                    var v18 string = ("a" + "hello")
                    _ = fi17
                    _ = v18
                }
            }
        } else {
            var v19 int = -9223372036854775808
            var v20 string = "a"
            _ = v19
            _ = v20
        }
        v21 := GenS1{F0: "a"}
        _ = v10
        _ = v21
    }
    var v22 uint = 42
    _ = v22
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{995, 42}
    v2 := make(map[string]uint32)
    var v3 int = 127
    var v4 float32 = 3.14159
    v5 := make(map[int]bool)
    v23 := GenS3{F0: 109, F1: -94, F2: 2}
    _ = len(v1)
    for fi24 := 0; fi24 < 14; fi24++ {
        v25 := v2[("hello" + ((("" + "") + "foo bar") + ("hello" + "hello")))]
        _ = fi24
        _ = v25
    }
    var v26 uint64 = 2
    var v27 int64 = 0
    var v28 float64 = 3.14159
    v27 = v27
    fmt.Println(v26, v28, v4)
    fmt.Println(v26, v3)
    var v29 int8 = 122
    v30 := GenS4{F0: 2, F1: 42}
    _ = len(v5)
    v27 += -523
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
