package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int32 = 160
    v2 := make(map[int]int)
    var v3 uint64 = 288
    v4 := make(map[int]int)
    var v5 int32 = -144
    if true {
        if false {
            v3 = v3
        } else {
            var v6 bool = false
            for fi7 := 0; fi7 < 18; fi7++ {
                switch "\n\t" {
                case "foo bar":
                    v1 = v1
                }
                _ = fi7
            }
            _ = v6
        }
        var v8 uint64 = (v3 & 0)
        var v9 uint64 = 4294967295
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: 181}
        _ = v10
    }
    var v11 int16 = 1
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 uint16 = 65535
        v14 := GenS1{F0: 1}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint8 = 128
    v11 = v11
    v1 = 854
    v1 = -1
    fmt.Println(v11, v1, v11)
    v16 := GenS1{F0: (v15 & 1)}
    v17 := []int{0, 32768, 42, -9223372036854775808}
    if (1e10 > 1e10) {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 int8 = 0
            v1 = (v5 & -881)
            for fi20 := 0; fi20 < 3; fi20++ {
                fmt.Println(v1)
                var v21 int8 = (v19 ^ -128)
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        fmt.Println(v1)
    } else {
        var v22 float64 = 1e10
        v11 -= (-(v11))
        _ = v22
    }
    delete(v4, -294)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
