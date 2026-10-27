package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int64; F2 int8 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint64, p1 int32, p2 uint64) uint {
    var v6 float64 = 0.1
    p1 += p1
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(p1, p2)
            _ = fi8
        }
        if true {
            var v9 int8 = 125
            _ = v9
        } else {
            if (-1 == 42) {
                switch "" {
                case "\n\t":
                    for fi10 := 0; fi10 < 20; fi10++ {
                        v11 := []int{fi7, (fi7 ^ fi7), fi7}
                        v11 = append(v11, 2147483647)
                        _ = fi10
                        _ = v11
                    }
                case "hello":
                    v12 := GenS1{F0: "hello"}
                    _ = v12
                default:
                    var v13 uint = 472
                    _ = v13
                }
                var v14 float64 = ((((float64(-2.25) / (v6 * (v6 / v6))) * v6) - v6) - 3.14159)
                v15 := []int{}
                _ = v14
                _ = v15
            } else {
                v16 := GenS2{F0: -910}
                _ = v16
            }
            var v17 int32 = (p1 >> ((226) & 63))
            _ = v17
        }
        _ = fi7
    }
    if (!true) {
        v18 := GenS2{F0: -2147483648}
        switch ("a" + "") {
        case "\n\t":
            v19 := GenS3{F0: -122, F1: 9223372036854775807, F2: 82}
            _ = v19
        case "hello":
            v20 := GenS1{F0: "a"}
            _ = v20
        default:
            p2 = (^(p2))
        }
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 128
    var v2 string = (("" + "a") + ("a" + "foo bar"))
    v3 := make(map[uint32]string)
    var v4 int16 = -1
    var v5 int32 = 42
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            if ((0 >= 32767) && false) {
                v2 = v2
                v1 -= v1
            } else {
                v5 -= (v5 - v5)
            }
            _ = fi22
        }
        _ = fi21
    }
    fmt.Println(v4, v5, v4)
    v23 := GenS1{F0: v2}
    var v24 int16 = 1000
    v25 := GenS4{F0: true}
    v2 = (("foo bar" + "a") + ("\n\t" + v2))
    var v26 string = (v2 + ((v2 + v2) + ""))
    for fi27 := 0; fi27 < 14; fi27++ {
        v24 += -1
        fmt.Println(v4, v26)
        var v28 int = -683
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
