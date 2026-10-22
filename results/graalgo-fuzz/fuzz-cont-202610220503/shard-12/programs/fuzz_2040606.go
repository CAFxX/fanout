package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int; F1 uint; F2 bool }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint64; F1 int16 }
func genFn0() uint64 {
    var v7 int8 = 39
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(fi8)
        v9 := GenS1{F0: 1}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 65535, F1: (-2.25 < (1e10 * ((3.14159 + (1e10 * (-2.25 / 1.0))) + -2.25)))}
    v11 := GenS1{F0: 21}
    for fi12 := 0; fi12 < 16; fi12++ {
        switch 32768 {
        case -799:
            v7 -= v7
        case 42:
            if false {
                for fi13 := 0; fi13 < 1; fi13++ {
                    v14 := GenS3{F0: 1, F1: 1, F2: (268 < 907)}
                    v7 -= 100
                    v15 := []int{(fi12 | (fi12 * 32767)), (^(fi12)), 2147483647}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
        case 0:
            if false {
                if false {
                    v7 = v7
                    fmt.Println(v7, v7)
                    var v16 float32 = 0.1
                    _ = v16
                }
                switch "" {
                case "foo bar":
                    var v17 bool = (1 > 255)
                    _ = v17
                case "\n\t":
                    switch 42 {
                    case 720:
                        fmt.Println(v7, fi12)
                    default:
                        switch "" {
                        case "hello":
                            v7 += int8(fi12)
                        }
                    }
                default:
                    var v18 int32 = -2147483648
                    _ = v18
                }
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 bool = true
                    var v21 int64 = 42
                    fmt.Println(fi12)
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        default:
            v22 := []int{}
            _ = v22
        }
        var v23 int = fi12
        v24 := GenS4{F0: 0, F1: false}
        _ = fi12
        _ = v23
        _ = v24
    }
    _ = v10
    _ = v11
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 65535
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]string)
    var v6 string = ""
    fmt.Println(v2, v2)
    for fi25 := 0; fi25 < 12; fi25++ {
        v26 := []int{fi25, (fi25 | fi25), (-(fi25)), (fi25 / ((-9223372036854775808) | 1))}
        var v27 float64 = 1.5
        _ = len(v26)
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 int = 0
    v29 := GenS5{F0: 383, F1: -1}
    _ = len(v3)
    var v30 int64 = -9223372036854775808
    v2 -= (v2 / ((18446744073709551615) | 1))
    fmt.Println(v6, v30)
    v31 := GenS2{F0: 301, F1: true}
    v28 -= -9223372036854775808
    var v32 uint8 = 44
    for fi33 := 0; fi33 < 19; fi33++ {
        v6 = "foo bar"
        _ = fi33
    }
    var v34 string = (((v6 + v6) + ((v6 + ("a" + "foo bar")) + "foo bar")) + v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
