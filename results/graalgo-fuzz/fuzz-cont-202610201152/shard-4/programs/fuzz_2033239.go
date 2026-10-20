package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int8, p1 uint64, p2 uint) int8 {
    v7 := []int{134}
    v8 := GenS1{F0: int8(p2)}
    if len(v7) > 0 {
        v7[0] = 9223372036854775807
    }
    var v9 float32 = float32(1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (p0 ^ -92)
}

func main() {
    var v1 int64 = 340
    v2 := make(map[uint32]uint32)
    v3 := []uint32{2147483648, 937, 42}
    var v4 uint = 2
    var v5 uint32 = 0
    v6 := []uint32{v5, v5, v5, v5}
    fmt.Println(v4, v1, v1)
    v5 = (v5 & (v5 & v5))
    var v10 uint = 2
    var v11 string = "a"
    delete(v2, 42)
    if false {
        v5 -= 2147483648
        for fi12 := 0; fi12 < 1; fi12++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi12
        }
    } else {
        v11 = v11
    }
    v13 := v2[(v5 - v5)]
    _ = len(v2)
    fmt.Println(v11, v1, v10)
    var v14 uint8 = 42
    fmt.Println(v14)
    v15 := GenS2{F0: (v14 * (^(v14))), F1: v14}
    for fi16 := 0; fi16 < 9; fi16++ {
        v1 = v1
        v4 -= v10
        _ = fi16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        switch v11 {
        case "":
            v5 = (v5 / (((v13 / (((v5 / ((2147483648) | 1))) | 1))) | 1))
        case "a":
            var v18 int32 = 42
            _ = v18
        default:
            for fi19 := 0; fi19 < 17; fi19++ {
                if false {
                    v10 = v4
                } else {
                    fmt.Println(v5, v13)
                }
                var v20 uint16 = 42
                _ = fi19
                _ = v20
            }
        }
        v5 -= (v5 + v13)
        var v21 string = (((v11 + v11) + v11) + (((((v11 + v11) + (v11 + v11)) + v11) + "foo bar") + v11))
        _ = fi17
        _ = v21
    }
    switch v11 {
    case "hello":
        switch 1 {
        case 32768:
            switch v11 {
            case "hello":
                for fi22 := 0; fi22 < 16; fi22++ {
                    fmt.Println(v1, v5)
                    var v23 string = ("foo bar" + (v11 + v11))
                    switch fi22 {
                    case 0:
                        for fi24 := 0; fi24 < 7; fi24++ {
                            var v25 uint8 = 127
                            _ = fi24
                            _ = v25
                        }
                    case 200:
                        var v26 uint16 = 743
                        _ = v26
                    default:
                        var v27 int8 = 90
                        _ = v27
                    }
                    _ = fi22
                    _ = v23
                }
            }
        case 9223372036854775807:
            var v28 bool = ((v13 % ((v5) | 1)) == v13)
            _ = v28
        }
    default:
        var v29 uint64 = 1
        _ = v29
    }
    if len(v3) > 0 {
        v3[0] = uint32(v14)
    }
    v30 := GenS2{F0: v14, F1: uint8(v10)}
    v3 = append(v3, 42)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
