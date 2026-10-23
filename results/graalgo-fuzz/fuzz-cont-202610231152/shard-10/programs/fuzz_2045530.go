package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int16, p1 uint8, p2 bool) uint16 {
    var v6 bool = false
    for fi7 := 0; fi7 < 5; fi7++ {
        p2 = (801 == 525)
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    var v9 int64 = 42
    v10 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v9
    return 184
}

func main() {
    var v1 uint = 4294967295
    var v2 int8 = -1
    v3 := make(map[string]bool)
    var v4 float32 = 1.5
    v5 := make(map[uint32]bool)
    v11 := GenS1{F0: -159}
    v2 += 85
    v4 += v4
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 uint64 = 1
        _ = fi12
        _ = v13
    }
    v4 -= 0.0
    if false {
        for fi14 := 0; fi14 < 6; fi14++ {
            for fi15 := 0; fi15 < 20; fi15++ {
                v2 = -84
                _ = fi15
            }
            _ = fi14
        }
    }
    v4 -= 0.0
    for fi16 := 0; fi16 < 4; fi16++ {
        v4 = ((((((v4 + 1e10) - v4) * v4) * ((v4 * -2.25) * (v4 / float32(float64(-2.25))))) * 1e10) * float32(0.1))
        var v17 int64 = 42
        _ = fi16
        _ = v17
    }
    v2 -= v2
    fmt.Println(v4, v1)
    v4 -= (v4 - v4)
    switch "\n\t" {
    case "foo bar":
        switch 0 {
        case -356:
            var v18 int64 = -1
            _ = v18
        case 42:
            v19 := []int{}
            _ = v19
        }
    }
    var v20 int8 = v2
    v21 := []int{4294967295, -986, 9223372036854775807, 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
