package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
func genFn0() int {
    if false {
        var v4 bool = false
        var v5 uint16 = 255
        _ = v4
        _ = v5
    }
    v6 := GenS1{F0: 42}
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    switch 577 {
    case 589:
        var v8 int32 = -2147483648
        _ = v8
    }
    _ = v6
    return 450
}
func genFn1(p0 int64, p1 int64) string {
    fmt.Println(p1)
    if true {
        p1 = (p0 << ((255) & 63))
    }
    var v9 int = genFn0()
    _ = p0
    _ = p1
    _ = v9
    return "foo bar"
}

func main() {
    var v1 float32 = 0.0
    var v2 float32 = v1
    v3 := make(map[int]uint32)
    v10 := []int{2147483647}
    var v11 uint64 = 42
    for fi12 := 0; fi12 < 17; fi12++ {
        v10 = append(v10, fi12)
        _ = fi12
    }
    if false {
        switch 2147483647 {
        case 127:
            v1 = v1
        case 4294967295:
            var v13 uint32 = 65535
            _ = v13
        case 2147483647:
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        }
        v14 := GenS2{F0: 255}
        var v15 uint8 = 1
        _ = v14
        _ = v15
    }
    v11 -= (^(v11))
    for fi16 := 0; fi16 < 15; fi16++ {
        fmt.Println(v11, v11, v11)
        fmt.Println(fi16, v2)
        _ = fi16
    }
    v17 := v3[9223372036854775807]
    v1 -= (v1 + v1)
    v18 := GenS1{F0: 1000}
    var v19 int64 = 1
    var v20 string = ""
    v20 = "foo bar"
    switch 65 {
    case 1:
        v19 += v19
    case -2147483648:
        v11 += v11
    }
    v21 := v3[642]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
