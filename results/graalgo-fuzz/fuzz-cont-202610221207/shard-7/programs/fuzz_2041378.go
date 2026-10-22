package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32, p1 uint32, p2 int32) uint16 {
    v7 := GenS1{F0: 4294967295}
    p2 -= -2147483648
    var v8 int8 = -88
    var v9 uint32 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    v1 := []int{}
    var v2 float64 = float64(0.0)
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    var v5 uint8 = 1
    var v6 uint32 = 148
    _ = len(v4)
    for fi10 := 0; fi10 < 8; fi10++ {
        switch "\n\t" {
        case "\n\t":
            v5 += v5
        case "foo bar":
            var v11 int = (-(fi10))
            _ = v11
        default:
            for fi12 := 0; fi12 < 20; fi12++ {
                fmt.Println(fi10, fi10, v5)
                if ((-1 < 0) || false) {
                    v2 -= (v2 + v2)
                    var v13 int64 = 322
                    _ = v13
                }
                var v14 uint = 2
                _ = fi12
                _ = v14
            }
        }
        v2 -= float64(0.0)
        _ = fi10
    }
    v15 := GenS2{F0: 327}
    v2 = float64(1.5)
    var v16 string = "a"
    for fi17 := 0; fi17 < 4; fi17++ {
        if (!false) {
            var v18 string = "hello"
            v19 := GenS1{F0: 1}
            v2 -= 1.5
            _ = v18
            _ = v19
        }
        var v20 uint64 = 42
        _ = fi17
        _ = v20
    }
    var v21 int64 = -1
    v5 = v5
    if false {
        v5 += v5
    }
    var v22 float64 = float64(-2.25)
    v21 += (v21 | (v21 | (v21 & (v21 % ((v21) | 1)))))
    v23 := v4[((v16 + v16) + v16)]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
