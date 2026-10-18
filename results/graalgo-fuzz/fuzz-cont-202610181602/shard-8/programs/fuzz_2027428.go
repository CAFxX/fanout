package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int, p1 uint32) int64 {
    p0 -= -808
    var v4 float64 = 1.5
    _ = p0
    _ = p1
    _ = v4
    return -175
}
func genFn1(p0 int8) int16 {
    var v5 uint64 = 646
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            if false {
                var v8 int16 = -112
                v9 := GenS1{F0: 1}
                fmt.Println(v8)
                _ = v8
                _ = v9
            } else {
                var v10 float64 = 3.14159
                v11 := GenS1{F0: 42}
                _ = v10
                _ = v11
            }
            p0 = p0
            _ = fi7
        }
        _ = fi6
    }
    v12 := []int{42}
    if false {
        var v13 uint32 = 288
        v14 := GenS2{F0: 138}
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v12
    _ = v5
    return 42
}

func main() {
    var v1 uint64 = 783
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    v1 = 803
    for fi15 := 0; fi15 < 11; fi15++ {
        v1 += (v1 | 1)
        _ = fi15
    }
    var v16 string = ""
    var v17 uint16 = 334
    v1 = 1
    v18 := v2[128]
    v16 = (((v16 + v16) + ((((v16 + v16) + "a") + v16) + v16)) + ("\n\t" + ((v16 + v16) + "")))
    var v19 uint16 = v17
    v20 := GenS2{F0: 127}
    var v21 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
