package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
func genFn0(p0 int, p1 uint8) int64 {
    fmt.Println(p0, p0)
    p1 += (^(p1))
    switch "a" {
    case "foo bar":
        if false {
            for fi5 := 0; fi5 < 18; fi5++ {
                var v6 bool = (!(1.5 == (1.5 / (float32(1.5) / (1e10 - 1.5)))))
                _ = fi5
                _ = v6
            }
        } else {
            switch p0 {
            case 9223372036854775807:
                fmt.Println(p1)
            default:
                p1 += (^(p1))
            }
            v7 := []int{(p0 - p0), (p0 >> ((57) & 63)), 746}
            var v8 int32 = 42
            _ = v7
            _ = v8
        }
    case "a":
        var v9 bool = false
        _ = v9
    default:
        if ((0.1 - 1e10) == 0.1) {
            var v10 int64 = 290
            v11 := GenS1{F0: 65535, F1: 2, F2: p0}
            _ = v10
            _ = v11
        }
    }
    if ((true || false) || false) {
        var v12 int64 = 9223372036854775807
        v13 := []int{p0, (^(p0)), (p0 - 0)}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint16 = 0
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int64 = 0
    v2 := make(map[string]bool)
    var v3 uint = 65535
    var v4 int64 = v1
    v16 := GenS2{F0: v3}
    delete(v2, "foo bar")
    var v17 int8 = -1
    v4 += (-(v4))
    var v18 bool = true
    v1 = (v4 << (((v3 * v3)) & 63))
    _ = len(v2)
    v19 := GenS3{F0: v18}
    v17 = v17
    var v20 int32 = 638
    if v18 {
        v20 -= (^(v20))
        for fi21 := 0; fi21 < 3; fi21++ {
            switch -337 {
            case 978:
                v2[(("foo bar" + "foo bar") + "a")] = (436 != 918)
            }
            v4 -= (v4 * v4)
            _ = fi21
        }
    } else {
        fmt.Println(v18, v1)
    }
    if true {
        fmt.Println(v18, v3, v20)
        v17 -= int8(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
