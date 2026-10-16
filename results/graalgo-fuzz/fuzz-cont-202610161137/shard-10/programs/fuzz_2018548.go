package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint, p1 string) int16 {
    p1 = (p1 + p1)
    var v7 string = p1
    var v8 uint32 = 42
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 uint16, p1 uint32, p2 uint32) uint {
    v9 := []int{9223372036854775807}
    switch -1 {
    case 51:
        v10 := GenS1{F0: 442}
        _ = v10
    case 9223372036854775807:
        p2 += p1
    case 42:
        for fi11 := 0; fi11 < 6; fi11++ {
            p2 += 87
            switch (fi11 | fi11) {
            case 32768:
                fmt.Println(fi11)
            case 127:
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 uint16 = (^(p0))
                    var v14 float32 = 1e10
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case 151:
                p1 += p2
            default:
                for fi15 := 0; fi15 < 7; fi15++ {
                    var v16 uint = 0
                    var v17 int32 = 0
                    if len(v9) > 0 {
                        v9[0] = fi11
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi11
        }
    default:
        p1 = uint32(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 255
}

func main() {
    v1 := []string{}
    var v2 uint64 = 42
    var v3 bool = (42 > 92)
    v4 := []int{1, 9223372036854775807}
    var v5 int16 = 42
    var v6 float32 = (((float32((1e10 / ((-2.25 - float64(1e10)) - 0.0))) + 1e10) - 0.1) - 1.5)
    var v18 uint64 = v2
    switch "a" {
    case "a":
        v18 = v18
    case "\n\t":
        fmt.Println(v18, v3, v3)
    }
    switch ("\n\t" + "a") {
    case "a":
        for fi19 := 0; fi19 < 12; fi19++ {
            _ = len(v4)
            _ = fi19
        }
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    }
    v2 -= v18
    var v20 int32 = 0
    v2 = 487
    v21 := GenS2{F0: (v20 * v20)}
    v22 := GenS3{F0: -1, F1: 2}
    if len(v1) > 0 {
        v1[0] = ""
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
