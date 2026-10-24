package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int {
    v7 := []int{}
    switch 2147483647 {
    case -208:
        if false {
            _ = len(v7)
            var v8 uint8 = 2
            fmt.Println(v8)
            _ = v8
        }
    case 2147483647:
        var v9 string = "\n\t"
        _ = v9
    }
    var v10 float32 = (1e10 * 0.1)
    if false {
        for fi11 := 0; fi11 < 10; fi11++ {
            v10 -= v10
            _ = fi11
        }
        var v12 uint16 = 632
        fmt.Println(v10, v12)
        _ = v12
    }
    v10 = (0.0 - (v10 * (v10 + (1e10 + (v10 * (v10 / float32(1.5)))))))
    _ = v10
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]int)
    v3 := []string{""}
    v4 := make(map[uint32]string)
    var v5 float32 = (1.5 / 0.1)
    var v6 int8 = 105
    if true {
        var v13 int16 = -588
        var v14 int64 = -1
        _ = v13
        _ = v14
    }
    v6 += v6
    for fi15 := 0; fi15 < 4; fi15++ {
        v6 = (-(v6))
        var v16 float64 = 1.5
        _ = fi15
        _ = v16
    }
    var v17 int8 = (^(v6))
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 string = "foo bar"
        v6 = 42
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        if (675 <= 309) {
            fmt.Println(v1, v5)
            v21 := GenS1{F0: 2}
            _ = v21
        }
        _ = fi20
    }
    var v22 string = v1
    var v23 int8 = v17
    var v24 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
