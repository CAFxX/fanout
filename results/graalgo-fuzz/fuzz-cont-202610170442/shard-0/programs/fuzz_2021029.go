package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int8 }
type GenS2 struct { F0 uint8; F1 string; F2 int16 }
type GenS3 struct { F0 uint64; F1 int; F2 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int64) uint8 {
    v5 := GenS1{F0: 42, F1: p0, F2: 0}
    p0 -= 42
    p0 += p0
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 bool) uint64 {
    if p0 {
        if (1 >= 255) {
            fmt.Println(p0)
        }
    } else {
        v6 := GenS2{F0: genFn0(616), F1: "a", F2: -335}
        var v7 float64 = 3.14159
        _ = v6
        _ = v7
    }
    v8 := []int{-9223372036854775808, 909, 127, 697}
    v9 := GenS3{F0: 2, F1: 127, F2: 42}
    var v10 float32 = ((0.0 * 0.0) + 0.0)
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (0.1 / 1.5)
    var v3 bool = false
    v4 := make(map[uint32]bool)
    if false {
        var v11 int8 = 115
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS4{F0: ("hello" != "foo bar")}
        for fi14 := 0; fi14 < 19; fi14++ {
            if true {
                v2 += float32(1.5)
                v15 := GenS3{F0: 4294967295, F1: fi12, F2: 1}
                if (v3 || ((1e10 * ((v2 - v2) + v2)) < v2)) {
                    var v16 uint8 = 187
                    var v17 int32 = 1
                    _ = v16
                    _ = v17
                }
                _ = v15
            }
            _ = fi14
        }
        var v18 float32 = ((v2 / v2) - (v2 / 3.14159))
        _ = fi12
        _ = v13
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            switch ("" + ("foo bar" + (("foo bar" + "hello") + ("a" + ("foo bar" + "a"))))) {
            case "foo bar":
                var v21 int = (fi19 % ((-9223372036854775808) | 1))
                _ = v21
            case "a":
                v3 = v3
            case "\n\t":
                v22 := GenS5{F0: 937}
                _ = v22
            }
            v3 = (2147483648 <= 4294967295)
            _ = fi20
        }
        v23 := []int{fi19, fi19, 42}
        var v24 uint16 = 447
        _ = fi19
        _ = v23
        _ = v24
    }
    v3 = v3
    v3 = (0.1 >= (3.14159 / (0.1 + ((0.1 + 1.5) * (-2.25 + 1e10)))))
    for fi25 := 0; fi25 < 12; fi25++ {
        v26 := []int{}
        _ = fi25
        _ = v26
    }
    v3 = (!(127 != 1))
    if (-1 < 42) {
        var v27 int64 = -362
        v2 -= float32((-2.25 / 1.5))
        _ = v27
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        v29 := v4[2147483647]
        v3 = v3
        _ = fi28
        _ = v29
    }
    v3 = v3
    v2 -= 1.5
    for fi30 := 0; fi30 < 19; fi30++ {
        v2 = v2
        _ = fi30
    }
    v3 = true
    v1[429] = (v3 && v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
