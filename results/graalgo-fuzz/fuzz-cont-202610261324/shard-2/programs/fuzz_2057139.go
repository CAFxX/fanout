package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint64) int8 {
    p1 = p1
    var v5 float64 = ((float64(-2.25) - 1e10) + 3.14159)
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int8 = 0
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return (p0 * (p0 >> ((42) & 63)))
}

func main() {
    var v1 string = "\n\t"
    var v2 bool = (false && (1 >= 0))
    var v3 float32 = -2.25
    var v4 int8 = -99
    var v8 uint8 = 228
    if true {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        switch 127 {
        case 32768:
            if false {
                var v11 float64 = 1e10
                _ = v11
            }
        case -711:
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 int8 = v4
                _ = fi12
                _ = v13
            }
        default:
            var v14 int16 = 0
            _ = v14
        }
        v15 := []int{1, 2147483647, 42, -493}
        _ = v15
    }
    switch -334 {
    case 263:
        if v2 {
            v8 -= v8
        } else {
            var v16 int64 = 967
            v16 += v16
            _ = v16
        }
    }
    v3 = v3
    var v17 float64 = 3.14159
    v18 := GenS1{F0: 32767}
    for fi19 := 0; fi19 < 4; fi19++ {
        v4 = (v4 / (((v4 / (((v4 ^ v4)) | 1))) | 1))
        var v20 uint8 = 1
        _ = fi19
        _ = v20
    }
    if (958 == 409) {
        for fi21 := 0; fi21 < 3; fi21++ {
            v17 += v17
            _ = fi21
        }
    }
    var v22 float64 = (0.1 / float64(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v8
}
