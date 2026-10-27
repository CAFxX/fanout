package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint8, p1 int16, p2 bool) int16 {
    var v4 uint16 = 1
    switch -9223372036854775808 {
    case 438:
        var v5 uint16 = 734
        _ = v5
    default:
        v6 := []int{int(v4), 402, 42}
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(p0, fi7, p0)
        v8 := GenS1{F0: 1}
        for fi9 := 0; fi9 < 13; fi9++ {
            switch (("foo bar" + "") + "\n\t") {
            case "a":
                if false {
                    v10 := []int{(fi9 / ((9223372036854775807) | 1)), fi9}
                    _ = v10
                }
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p1
}

func main() {
    var v1 int8 = 38
    var v2 float32 = 1.5
    v3 := make(map[int]bool)
    var v11 uint8 = 255
    v12 := []int{0, 128, 447}
    if (!(0 >= 18446744073709551615)) {
        fmt.Println(v11)
        _ = len(v12)
    }
    var v13 int64 = 1
    v13 += 0
    fmt.Println(v13, v13)
    v2 = ((3.14159 * (3.14159 * (v2 - v2))) - (float32((-2.25 + -2.25)) - v2))
    v14 := GenS1{F0: v13}
    var v15 uint16 = 2
    v12 = append(v12, 134)
    if (!(!(0.1 <= 0.1))) {
        fmt.Println(v2, v13, v1)
        v15 -= v15
        var v16 int = -9223372036854775808
        _ = v16
    }
    switch int(v13) {
    case -2147483648:
        var v17 int16 = 224
        _ = v17
    }
    var v18 uint16 = (v15 >> ((777) & 63))
    if true {
        v19 := GenS1{F0: v13}
        for fi20 := 0; fi20 < 19; fi20++ {
            if true {
                for fi21 := 0; fi21 < 4; fi21++ {
                    v18 = (v15 * v15)
                    v18 = v15
                    var v22 uint8 = 61
                    _ = fi21
                    _ = v22
                }
            }
            v15 = 42
            v23 := GenS1{F0: -547}
            _ = fi20
            _ = v23
        }
        _ = v19
    } else {
        var v24 int64 = v13
        v2 += v2
        _ = v24
    }
    var v25 uint64 = 1
    switch "\n\t" {
    case "":
        v3[int(v11)] = false
    case "hello":
        v26 := GenS1{F0: (v13 * 9223372036854775807)}
        _ = v26
    case "a":
        var v27 int32 = 42
        _ = v27
    default:
        var v28 int32 = -2147483648
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v25
    _ = v3
}
