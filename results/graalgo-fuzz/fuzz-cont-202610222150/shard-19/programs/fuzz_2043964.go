package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
func genFn0(p0 int, p1 uint16, p2 bool) uint8 {
    var v7 uint = 4294967295
    if false {
        fmt.Println(p2, p2, p1)
        v7 -= 608
        p2 = (p2 && p2)
    } else {
        fmt.Println(p1)
        for fi8 := 0; fi8 < 4; fi8++ {
            p2 = true
            _ = fi8
        }
        v9 := GenS1{F0: "", F1: -463}
        _ = v9
    }
    var v10 int8 = 42
    v11 := []int{p0, (-(p0)), (-(p0)), p0}
    v7 = (^(v7))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    return 112
}
func genFn1(p0 uint16, p1 int64) uint32 {
    fmt.Println(p1, p0)
    v12 := []int{128, -802}
    _ = len(v12)
    _ = len(v12)
    p1 = (p1 & p1)
    _ = p0
    _ = p1
    _ = v12
    return 954
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]string)
    var v3 int16 = 0
    v4 := make(map[string]string)
    var v5 int16 = -725
    var v6 int64 = -59
    var v13 int64 = (-(v6))
    _ = len(v4)
    v14 := GenS1{F0: "", F1: -1}
    fmt.Println(v1, v3)
    var v15 string = "foo bar"
    v16 := GenS1{F0: v15, F1: 0}
    switch (((((v15 + v15) + v15) + "a") + v15) + v15) {
    case "\n\t":
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 float32 = 1.5
            for fi19 := 0; fi19 < 10; fi19++ {
                v1 += v1
                v20 := v2[65535]
                v18 = (v18 + float32(((0.1 * (1.5 * 0.1)) - (0.0 * float64(v18)))))
                _ = fi19
                _ = v20
            }
            var v21 float32 = v18
            _ = fi17
            _ = v18
            _ = v21
        }
    case "foo bar":
        v13 = 0
    }
    var v22 float64 = (1e10 * (1e10 / 0.1))
    fmt.Println(v15)
    v15 = (((v15 + (v15 + (v15 + (("a" + "hello") + (v15 + v15))))) + ((v15 + v15) + ((v15 + "\n\t") + v15))) + (v15 + ""))
    v6 = (v6 >> ((0) & 63))
    delete(v4, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
