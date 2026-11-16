package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int }
func genFn0(p0 bool) uint8 {
    fmt.Println(p0)
    v6 := GenS1{F0: 2}
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            p0 = p0
            if p0 {
                fmt.Println(fi7, p0, fi8)
            }
            _ = fi8
        }
        _ = fi7
    }
    var v9 string = "a"
    _ = p0
    _ = v6
    _ = v9
    return 2
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]bool)
    v3 := []string{}
    v4 := []int{-1}
    v5 := make(map[uint32]int)
    var v10 uint32 = 4294967295
    for fi11 := 0; fi11 < 16; fi11++ {
        if true {
            v12 := GenS2{F0: 65535, F1: fi11}
            v10 -= (v10 | (v10 % ((857) | 1)))
            switch "" {
            case "a":
                v10 = (^(v10))
            case "hello":
                for fi13 := 0; fi13 < 19; fi13++ {
                    v1 += (v1 - (v1 * v1))
                    _ = fi13
                }
            case "\n\t":
                var v14 uint64 = 1
                _ = v14
            }
            _ = v12
        } else {
            var v15 uint16 = 65535
            _ = v15
        }
        var v16 uint32 = (v10 & (^(v10)))
        _ = fi11
        _ = v16
    }
    if true {
        v10 -= (v10 % ((v10) | 1))
        v10 = v10
        var v17 int8 = 1
        _ = v17
    } else {
        for fi18 := 0; fi18 < 4; fi18++ {
            switch (fi18 * -326) {
            case 1:
                v1 = (v1 / ((-2147483648) | 1))
            case 789:
                var v19 int32 = (v1 | v1)
                _ = v19
            }
            v1 = v1
            _ = fi18
        }
        var v20 bool = false
        _ = len(v5)
        _ = v20
    }
    var v21 uint = 42
    var v22 int8 = 73
    var v23 uint32 = (^(v10))
    fmt.Println(v23, v1, v10)
    switch 32767 {
    case 265:
        _ = len(v5)
    }
    _ = len(v3)
    var v24 float64 = ((3.14159 * (1.5 * 0.0)) * 0.0)
    v5[65535] = 1
    if (false || true) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v25 uint16 = 0
    var v26 int64 = int64(v25)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
