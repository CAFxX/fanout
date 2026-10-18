package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int }

func main() {
    var v1 uint16 = 738
    var v2 uint32 = 2147483647
    v3 := []int{42}
    v4 := make(map[uint32]bool)
    if false {
        v2 -= 2147483648
    } else {
        var v5 uint = 325
        v1 += (v1 >> ((v5) & 63))
        v6 := v4[(v2 >> (((^(v5))) & 63))]
        _ = v5
        _ = v6
    }
    fmt.Println(v2)
    v1 -= v1
    if true {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 string = ("\n\t" + "")
            _ = fi7
            _ = v8
        }
    } else {
        v9 := GenS1{F0: -1, F1: 4294967295, F2: 128}
        var v10 string = "a"
        _ = v10
        _ = v9
    }
    if true {
        if true {
            v2 = (v2 * v2)
        }
    } else {
        var v11 uint = 2
        if len(v3) > 0 {
            v3[0] = 32768
        }
        _ = v11
    }
    if len(v3) > 0 {
        v3[0] = 32767
    }
    v2 += v2
    var v12 float32 = ((((((1e10 * 0.1) * 0.1) / float32(3.14159)) - (1.5 * 0.0)) * (1e10 * (3.14159 / -2.25))) + 0.1)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
