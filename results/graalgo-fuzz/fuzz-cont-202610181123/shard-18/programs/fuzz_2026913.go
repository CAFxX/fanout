package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() int32 {
    var v5 float32 = 0.1
    var v6 bool = false
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 int16, p1 int64, p2 uint32) uint16 {
    for fi7 := 0; fi7 < 10; fi7++ {
        p0 = p0
        p2 += p2
        _ = fi7
    }
    var v8 float64 = float64(((1e10 / (float32(1.5) - 1e10)) + (3.14159 * -2.25)))
    p1 = p1
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 255
}

func main() {
    var v1 uint = 543
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v4 float32 = -2.25
    if true {
        v1 = (^(v1))
        var v9 uint64 = 0
        _ = v9
    } else {
        var v10 uint = (v1 | v1)
        _ = v10
    }
    v4 = float32(0.0)
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := []int{127, 220, 237, -1}
        _ = fi11
        _ = v12
    }
    switch 32767 {
    case 357:
        if (42 != 4294967295) {
            v13 := GenS1{F0: 555}
            if false {
                var v14 int32 = genFn0()
                _ = v14
            } else {
                switch -2147483648 {
                case 32767:
                    if false {
                        var v15 int = -721
                        fmt.Println(v15, v15)
                        _ = v15
                    }
                }
                fmt.Println(v4, v4)
                var v16 int64 = 1
                _ = v16
            }
            _ = v13
        } else {
            v17 := GenS1{F0: -1}
            var v18 float64 = 3.14159
            var v19 uint8 = 2
            _ = v17
            _ = v18
            _ = v19
        }
    case -2147483648:
        v20 := GenS1{F0: 2147483647}
        _ = v20
    }
    fmt.Println(v1, v1)
    v4 = 0.0
    fmt.Println(v1, v1)
    var v21 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
