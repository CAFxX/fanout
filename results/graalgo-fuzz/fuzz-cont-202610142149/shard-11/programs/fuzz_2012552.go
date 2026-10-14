package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int64; F1 int8; F2 int8 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 string = ("hello" + ("\n\t" + ((("a" + (("" + ("" + ("a" + "hello"))) + ("" + ""))) + ("foo bar" + ("foo bar" + ("a" + "foo bar")))) + (("hello" + "a") + ("" + "")))))
    var v2 int64 = 0
    v3 := make(map[string]uint32)
    v4 := make(map[int]uint32)
    var v5 int = 128
    v2 -= v2
    var v6 bool = (false || false)
    v7 := GenS1{F0: 127, F1: 1, F2: 4294967295}
    v1 = v1
    v8 := GenS1{F0: (^(v5)), F1: 127, F2: 852}
    v5 -= v5
    switch v1 {
    case "a":
        v1 = v1
    case "":
        if v6 {
            if v6 {
                if ((0.0 * (0.1 * (((3.14159 * (3.14159 + (1.5 * -2.25))) + 0.1) * 1.5))) <= 0.1) {
                    for fi9 := 0; fi9 < 10; fi9++ {
                        var v10 int32 = -232
                        v11 := []int{fi9, -568}
                        var v12 uint8 = 128
                        _ = fi9
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    var v13 int32 = 131
                    _ = v13
                }
            }
            switch v1 {
            case "\n\t":
                if v6 {
                    if (67 < 42) {
                        var v14 float64 = 1e10
                        _ = v14
                    }
                    v15 := GenS2{F0: (v2 / ((-828) | 1)), F1: -97, F2: 1}
                    _ = v15
                } else {
                    v5 = (v5 >> ((255) & 63))
                }
            case "a":
                switch ("a" + v1) {
                case "hello":
                    if false {
                        for fi16 := 0; fi16 < 19; fi16++ {
                            var v17 int32 = -1
                            var v18 bool = ((!(v2 >= v2)) || false)
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                default:
                    var v19 int16 = -1
                    _ = v19
                }
            default:
                if (395 > 42) {
                    for fi20 := 0; fi20 < 16; fi20++ {
                        v2 -= (v2 ^ (v2 - 9223372036854775807))
                        _ = fi20
                    }
                    var v21 uint = 1
                    for fi22 := 0; fi22 < 4; fi22++ {
                        v23 := GenS3{F0: (false && (v6 && false)), F1: (v2 & -360)}
                        switch 0 {
                        case 66:
                            v24 := GenS4{F0: 4294967295, F1: 18446744073709551615}
                            _ = v24
                        case -37:
                            fmt.Println(v5, v5)
                        default:
                            var v25 uint = v21
                            _ = v25
                        }
                        _ = fi22
                        _ = v23
                    }
                    _ = v21
                }
            }
            var v26 int8 = 0
            _ = v26
        }
    default:
        v5 += v5
    }
    fmt.Println(v6, v6, v5)
    var v27 float32 = 3.14159
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
