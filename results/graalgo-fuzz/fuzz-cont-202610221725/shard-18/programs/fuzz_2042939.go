package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint, p1 string, p2 uint16) bool {
    if true {
        var v6 uint16 = (p2 ^ 2)
        p0 += 42
        _ = v6
    } else {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint64 = 2
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS1{F0: 0}
    fmt.Println(p2, p0)
    v10 := GenS2{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return true
}

func main() {
    var v1 uint = 980
    var v2 float32 = float32(1.5)
    v3 := make(map[int]bool)
    v4 := []int{32768, -592, -255}
    v5 := []uint32{2147483647, 104}
    var v11 uint8 = 128
    v11 = v11
    if (false || (!(!false))) {
        for fi12 := 0; fi12 < 4; fi12++ {
            v11 -= v11
            var v13 string = ""
            v2 = ((v2 - (0.1 * (-2.25 / (v2 + v2)))) - 0.0)
            _ = fi12
            _ = v13
        }
    }
    switch 803 {
    case 4294967295:
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 uint8 = 128
            var v16 uint = (^(v1))
            _ = fi14
            _ = v15
            _ = v16
        }
    case 0:
        switch "a" {
        case "foo bar":
            v11 = v11
        case "\n\t":
            v17 := GenS1{F0: 872}
            _ = v17
        case "a":
            var v18 float64 = 3.14159
            _ = v18
        }
    case 2147483647:
        v19 := GenS3{F0: 109}
        _ = v19
    }
    var v20 int64 = 42
    if false {
        var v21 float32 = (0.1 * float32((1e10 * 0.1)))
        v21 += (0.1 + (((((v2 * 1e10) - (v21 / 1.5)) + v2) + v21) / 1.5))
        _ = v21
    }
    if (0 > -631) {
        v1 += (v1 % ((42) | 1))
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v2 = (0.1 / v2)
        _ = len(v4)
    }
    v22 := GenS3{F0: -128}
    var v23 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
