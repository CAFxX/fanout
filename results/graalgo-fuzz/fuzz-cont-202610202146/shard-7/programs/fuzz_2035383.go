package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 int64, p1 int32) uint32 {
    if (25 != 100) {
        p0 += p0
        v5 := []int{2147483647}
        _ = v5
    } else {
        if false {
            for fi6 := 0; fi6 < 7; fi6++ {
                var v7 string = "a"
                _ = fi6
                _ = v7
            }
            p1 -= (p1 ^ p1)
        }
    }
    if false {
        if true {
            if false {
                var v8 int8 = 42
                _ = v8
            } else {
                var v9 bool = true
                for fi10 := 0; fi10 < 15; fi10++ {
                    v11 := []int{fi10, (fi10 ^ fi10), 1, (fi10 * fi10)}
                    _ = fi10
                    _ = v11
                }
                var v12 int8 = -128
                _ = v12
                _ = v9
            }
        }
        v13 := GenS1{F0: 559, F1: -1, F2: 127}
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint16 = 65535
            v16 := []int{(fi14 * (fi14 * fi14)), fi14, (fi14 % ((542) | 1)), fi14}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    } else {
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 uint = 65535
            var v19 uint32 = 1
            var v20 uint32 = (v19 << ((v18) & 63))
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        var v21 int = -2147483648
        v22 := []int{v21, 4294967295, v21}
        _ = v21
        _ = v22
    }
    var v23 int64 = p0
    v24 := []int{9223372036854775807, 127}
    v25 := GenS1{F0: 2, F1: 42, F2: 128}
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    _ = v25
    return 553
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (!(42 == 827))
    var v3 bool = (417 != 4294967295)
    v4 := make(map[int]int)
    v26 := GenS1{F0: 512, F1: -667, F2: 255}
    _ = len(v1)
    var v27 uint32 = 2
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 uint64 = 421
        _ = fi28
        _ = v29
    }
    fmt.Println(v3, v3)
    v2 = true
    var v30 float32 = 0.1
    if false {
        for fi31 := 0; fi31 < 7; fi31++ {
            for fi32 := 0; fi32 < 8; fi32++ {
                delete(v1, genFn0(42, -1))
                _ = fi32
            }
            v30 = (1e10 / v30)
            v30 -= (3.14159 + float32(-2.25))
            _ = fi31
        }
        v30 = v30
        var v33 bool = false
        _ = v33
    }
    v34 := GenS2{F0: (927 >= 406)}
    fmt.Println(v27, v27, v30)
    switch "a" {
    case "":
        var v35 uint8 = 2
        _ = v35
    case "a":
        var v36 uint8 = 250
        _ = v36
    default:
        v27 = (v27 & 42)
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        fmt.Println(v30)
        _ = fi37
    }
    var v38 float64 = ((1e10 / (3.14159 / float64(v30))) * 1.5)
    for fi39 := 0; fi39 < 13; fi39++ {
        var v40 int64 = 1
        _ = fi39
        _ = v40
    }
    v41 := GenS2{F0: false}
    v30 -= float32((3.14159 * v38))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v34
    _ = v38
    _ = v4
    _ = v41
}
