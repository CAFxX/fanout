package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int8, p1 uint) bool {
    for fi5 := 0; fi5 < 16; fi5++ {
        p1 = 650
        _ = fi5
    }
    var v6 int16 = 358
    fmt.Println(v6, v6, v6)
    fmt.Println(v6, p0)
    v6 += v6
    _ = p0
    _ = p1
    _ = v6
    return true
}
func genFn1(p0 int32) int32 {
    v7 := []int{1, -2147483648, 32767}
    p0 = (p0 - p0)
    _ = p0
    _ = v7
    return p0
}

func main() {
    var v1 string = ((("\n\t" + "a") + ("" + (("hello" + ("" + "")) + ""))) + "foo bar")
    var v2 string = v1
    var v3 int = 32767
    v4 := make(map[int]bool)
    if (false || true) {
        var v8 uint16 = 344
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 bool = true
            var v11 uint32 = 463
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    }
    if (v3 == (v3 >> ((255) & 63))) {
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(v1, v3, v3)
            var v13 uint16 = 42
            _ = fi12
            _ = v13
        }
        v3 = v3
    }
    v14 := GenS1{F0: 4294967295}
    var v15 uint32 = 2
    var v16 uint8 = 0
    v4[(-(v3))] = ((true || false) || false)
    var v17 float32 = float32(1e10)
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 float64 = float64(1e10)
        _ = fi18
        _ = v19
    }
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
