package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int32) uint32 {
    if true {
        if (128 > 128) {
            v6 := GenS1{F0: 0, F1: 2}
            var v7 int8 = -128
            _ = v6
            _ = v7
        } else {
            v8 := GenS2{F0: 42, F1: 165, F2: 959}
            p0 += 42
            _ = v8
        }
        if false {
            v9 := []int{127, 443, 2147483647, 127}
            _ = v9
        }
        p0 += (p0 / ((p0) | 1))
    } else {
        p0 -= 2147483647
    }
    p0 = (p0 | p0)
    v10 := []int{}
    if false {
        switch 9223372036854775807 {
        case -2147483648:
            fmt.Println(p0)
        case 0:
            switch "foo bar" {
            case "hello":
                var v11 int16 = -301
                _ = v11
            case "a":
                fmt.Println(p0)
            case "foo bar":
                if true {
                    for fi12 := 0; fi12 < 8; fi12++ {
                        fmt.Println(fi12, p0)
                        _ = fi12
                    }
                    var v13 string = ("\n\t" + "hello")
                    p0 -= (-(p0))
                    _ = v13
                }
            }
        case 127:
            p0 = p0
        default:
            if true {
                fmt.Println(p0)
                var v14 bool = false
                _ = v14
            } else {
                fmt.Println(p0)
            }
        }
    } else {
        p0 = p0
    }
    var v15 string = "a"
    _ = p0
    _ = v10
    _ = v15
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    v3 := []int{}
    var v4 uint16 = 255
    v5 := []string{((("foo bar" + ("foo bar" + "\n\t")) + "foo bar") + "hello"), "", "\n\t"}
    v4 -= v4
    v4 -= (v4 | 42)
    v16 := GenS1{F0: 4294967295, F1: 233}
    v3 = append(v3, 128)
    var v17 int32 = -2147483648
    v18 := GenS3{F0: 642, F1: 116, F2: (v4 + (v4 * v4))}
    v19 := GenS4{F0: 274}
    if false {
        if true {
            var v20 uint16 = 42
            var v21 float32 = float32(3.14159)
            _ = v20
            _ = v21
        }
    } else {
        var v22 int32 = (v17 << ((0) & 63))
        v17 = (v22 % (((v17 | (-(v22)))) | 1))
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
