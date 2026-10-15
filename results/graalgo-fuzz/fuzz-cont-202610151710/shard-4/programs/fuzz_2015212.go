package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int32, p1 int64, p2 uint8) int {
    for fi5 := 0; fi5 < 5; fi5++ {
        if true {
            p1 = p1
            v6 := []int{-2147483648}
            p1 = (p1 & (-(p1)))
            _ = v6
        } else {
            var v7 bool = false
            var v8 int64 = p1
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 20; fi9++ {
            fmt.Println(p2, fi9)
            _ = fi9
        }
        _ = fi5
    }
    switch 9223372036854775807 {
    case 2147483647:
        for fi10 := 0; fi10 < 15; fi10++ {
            fmt.Println(p0, p2)
            v11 := []int{-9223372036854775808}
            _ = fi10
            _ = v11
        }
    case -674:
        var v12 int32 = (^(p0))
        _ = v12
    case 940:
        if ((997 >= 65535) && false) {
            v13 := GenS1{F0: 0}
            if (true || (p2 <= p2)) {
                for fi14 := 0; fi14 < 17; fi14++ {
                    v15 := GenS1{F0: 42}
                    v16 := []int{(fi14 + (^(fi14))), fi14, fi14, 72}
                    p0 += (-(p0))
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                p1 = -347
                p0 += 0
            } else {
                var v17 int64 = p1
                fmt.Println(p2)
                v18 := GenS1{F0: -81}
                _ = v17
                _ = v18
            }
            var v19 float64 = -2.25
            _ = v13
            _ = v19
        }
    default:
        v20 := GenS1{F0: 847}
        _ = v20
    }
    var v21 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return -1
}

func main() {
    var v1 float64 = (0.1 * (1.5 + (0.1 * 0.0)))
    var v2 bool = (-982 != 602)
    v3 := make(map[uint32]int)
    var v4 uint16 = 42
    v22 := GenS2{F0: 0}
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 11; fi24++ {
            v25 := []int{(fi23 + fi23), fi23, fi24}
            if v2 {
                v26 := GenS2{F0: 470}
                v27 := GenS2{F0: 954}
                fmt.Println(v4, fi23, v4)
                _ = v26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    var v28 int8 = 42
    var v29 int64 = 0
    if v2 {
        v29 = (^(v29))
    }
    var v30 uint64 = 245
    var v31 int = 128
    var v32 uint32 = 65535
    switch (v31 << ((2) & 63)) {
    case 869:
        var v33 uint8 = 2
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
