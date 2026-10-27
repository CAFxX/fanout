package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0(p0 uint64, p1 uint16) uint {
    fmt.Println(p1, p0)
    var v5 uint64 = p0
    if false {
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 float64 = float64((3.14159 / 3.14159))
            var v8 uint8 = 83
            if (false && false) {
                fmt.Println(fi6, v5, p1)
                v9 := []int{fi6, fi6, (-(fi6)), (fi6 / ((756) | 1))}
                _ = v9
            } else {
                var v10 uint64 = v5
                v11 := GenS1{F0: v8}
                _ = v10
                _ = v11
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 float64 = -2.25
            _ = fi12
            _ = v13
        }
    } else {
        var v14 string = ((("a" + "a") + "foo bar") + "")
        for fi15 := 0; fi15 < 9; fi15++ {
            if (-818 > -1) {
                p1 = p1
            } else {
                fmt.Println(v5, p1, fi15)
            }
            _ = fi15
        }
        v5 = v5
        _ = v14
    }
    if true {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int8 = 101
            _ = fi16
            _ = v17
        }
        v18 := []int{int(v5)}
        var v19 float32 = float32(0.0)
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v5
    return 277
}

func main() {
    var v1 uint16 = 2
    var v2 uint32 = 112
    v3 := make(map[int]bool)
    var v4 float64 = (float64(0.0) / (-2.25 * 1.5))
    v4 -= (v4 + 1.5)
    switch ((("foo bar" + ("" + (("a" + "a") + ((("a" + "") + "\n\t") + "hello")))) + (("foo bar" + (("" + ("hello" + ("\n\t" + "a"))) + "foo bar")) + "foo bar")) + ("\n\t" + ("\n\t" + "foo bar"))) {
    case "hello":
        v1 = v1
    case "\n\t":
        if false {
            v20 := GenS1{F0: 191}
            _ = v20
        } else {
            var v21 uint32 = v2
            _ = v21
        }
    default:
        v2 = 2
    }
    switch "" {
    case "\n\t":
        var v22 int8 = 100
        _ = v22
    }
    v1 = 0
    v2 += v2
    for fi23 := 0; fi23 < 10; fi23++ {
        fmt.Println(v4, v1)
        v1 += (^(v1))
        var v24 uint16 = 1
        _ = fi23
        _ = v24
    }
    if true {
        var v25 uint16 = (v1 << ((0) & 63))
        v26 := GenS2{F0: -581, F1: 169}
        _ = v25
        _ = v26
    }
    var v27 uint16 = v1
    v28 := v3[42]
    v29 := v3[-9223372036854775808]
    v30 := GenS1{F0: 216}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
