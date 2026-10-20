package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 bool; F1 uint64; F2 string }

func main() {
    var v1 uint16 = 42
    var v2 float64 = (1.5 - (0.0 / 1.5))
    v3 := make(map[int]bool)
    var v4 uint64 = 0
    v5 := []string{"a", "a", (("foo bar" + "\n\t") + "a"), "hello"}
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 bool = (!false)
        var v8 int8 = -61
        var v9 int = fi6
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v2 -= 0.0
    var v10 int32 = -1
    v2 = 3.14159
    v1 -= 32767
    v11 := GenS1{F0: (v4 % ((1) | 1)), F1: 545, F2: 2}
    if (1 > 2) {
        switch 128 {
        case 32767:
            var v12 uint = 510
            _ = v12
        case 4294967295:
            v4 = v4
        case -9223372036854775808:
            v10 += v10
        default:
            for fi13 := 0; fi13 < 17; fi13++ {
                if len(v5) > 0 {
                    v5[0] = ("hello" + ("a" + "hello"))
                }
                var v14 uint32 = uint32(v4)
                v2 = v2
                _ = fi13
                _ = v14
            }
        }
    } else {
        switch 1 {
        case -575:
            v1 += v1
        }
        v15 := GenS2{F0: 77, F1: "hello"}
        _ = v15
    }
    if false {
        delete(v3, 0)
        v16 := GenS2{F0: -32768, F1: "\n\t"}
        if false {
            v2 -= 3.14159
        }
        _ = v16
    } else {
        v17 := GenS2{F0: -910, F1: ""}
        var v18 int64 = 42
        for fi19 := 0; fi19 < 18; fi19++ {
            v1 += (^(v1))
            for fi20 := 0; fi20 < 4; fi20++ {
                v21 := GenS3{F0: false, F1: v4, F2: "a"}
                if ((!true) || false) {
                    if (!(0.0 < (v2 - v2))) {
                        var v22 uint = 0
                        _ = v22
                    }
                    var v23 bool = true
                    switch 277 {
                    case -456:
                        if ((v10 - (v10 / ((42) | 1))) == v10) {
                            v18 = (v18 + (v18 * (-(v18))))
                        } else {
                            var v24 int8 = -128
                            v2 += v2
                            var v25 float64 = v2
                            _ = v24
                            _ = v25
                        }
                    }
                    _ = v23
                } else {
                    v26 := GenS2{F0: 1, F1: "hello"}
                    fmt.Println(v1)
                    _ = v26
                }
                switch ("foo bar" + ("foo bar" + "")) {
                case "":
                    var v27 uint32 = 1
                    _ = v27
                case "a":
                    var v28 uint8 = 1
                    _ = v28
                case "hello":
                    fmt.Println(fi19, v4, fi20)
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        _ = v17
        _ = v18
    }
    fmt.Println(v2, v10, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
