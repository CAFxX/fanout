package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int64; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint16; F1 string; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 int16 }
func genFn0(p0 uint64, p1 int64) uint16 {
    switch 127 {
    case 2147483647:
        v4 := GenS1{F0: int(p1), F1: 42}
        _ = v4
    case 32768:
        var v5 int = -9223372036854775808
        _ = v5
    }
    var v6 int8 = 127
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = -9223372036854775808
    var v3 int8 = 1
    v7 := GenS2{F0: -402, F1: 42, F2: 65535}
    if (42 < -1) {
        if true {
            switch "\n\t" {
            case "":
                for fi8 := 0; fi8 < 3; fi8++ {
                    switch fi8 {
                    case 253:
                        switch 9223372036854775807 {
                        case -9223372036854775808:
                            v3 = (v3 & (v3 & 33))
                        case 0:
                            v2 = (fi8 - v2)
                        case -286:
                            v3 = v3
                        default:
                            var v9 int64 = -139
                            _ = v9
                        }
                    }
                    for fi10 := 0; fi10 < 7; fi10++ {
                        var v11 int8 = (v3 / ((127) | 1))
                        _ = fi10
                        _ = v11
                    }
                    v2 += (fi8 + fi8)
                    _ = fi8
                }
            default:
                v12 := GenS3{F0: 221, F1: "\n\t", F2: 530}
                _ = v12
            }
        } else {
            var v13 float64 = (3.14159 / float64(3.14159))
            v14 := GenS4{F0: 23}
            for fi15 := 0; fi15 < 9; fi15++ {
                switch 9223372036854775807 {
                case 1:
                    fmt.Println(v3)
                case 2147483647:
                    v3 = v3
                case -655:
                    var v16 string = ("foo bar" + "foo bar")
                    _ = v16
                }
                v2 += fi15
                _ = fi15
            }
            _ = v13
            _ = v14
        }
        var v17 uint = 65535
        v3 = (v3 + (v3 - 127))
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        if (728 != 255) {
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 int64 = 42
                v21 := []int{(-(fi19)), fi18}
                _ = fi19
                _ = v20
                _ = v21
            }
            v3 += int8(v2)
        }
        var v22 uint8 = 255
        _ = fi18
        _ = v22
    }
    v23 := GenS5{F0: ("a" + "hello"), F1: 42}
    if false {
        var v24 bool = (false || (42 > 65535))
        _ = v24
    }
    v25 := v1[(v2 & 690)]
    v26 := GenS5{F0: ((((("a" + "\n\t") + "\n\t") + "a") + "hello") + ("" + "foo bar")), F1: -1}
    var v27 int8 = v3
    var v28 float32 = (0.1 + 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v7
}
