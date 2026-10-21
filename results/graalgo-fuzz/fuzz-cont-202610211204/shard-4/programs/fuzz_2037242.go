package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 string }
func genFn0() int8 {
    if true {
        var v4 uint64 = 42
        _ = v4
    } else {
        var v5 uint = 511
        fmt.Println(v5)
        _ = v5
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 bool = ((!false) || true)
        _ = fi6
        _ = v7
    }
    var v8 uint64 = 2
    if (("a" + ("foo bar" + "hello")) > ("a" + (("a" + "\n\t") + "foo bar"))) {
        var v9 int64 = 1
        v10 := GenS1{F0: ("" + "hello")}
        fmt.Println(v8, v8)
        _ = v10
        _ = v9
    } else {
        var v11 uint32 = 72
        switch 9223372036854775807 {
        case 379:
            v12 := []int{32767}
            _ = v12
        case -9223372036854775808:
            v11 = 42
        }
        fmt.Println(v8, v11)
        _ = v11
    }
    switch -1 {
    case -1:
        v13 := GenS2{F0: 0, F1: 153, F2: 127}
        _ = v13
    default:
        v14 := GenS2{F0: 0, F1: 42, F2: 4}
        _ = v14
    }
    _ = v8
    return -10
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    v15 := v3[65535]
    var v16 int64 = -747
    fmt.Println(v16, v15)
    v16 -= (v16 | 410)
    v16 -= (v16 << ((0) & 63))
    var v17 float32 = 0.1
    v18 := []int{}
    var v19 int = 1
    switch (("\n\t" + (((("a" + "") + "hello") + "hello") + "foo bar")) + ("a" + "hello")) {
    case "foo bar":
        switch v19 {
        case -2147483648:
            if len(v18) > 0 {
                v18[0] = v19
            }
        case 128:
            var v20 float32 = v17
            _ = v20
        case -1:
            v16 += (v16 << ((42) & 63))
        }
    case "a":
        if len(v18) > 0 {
            v18[0] = v19
        }
    case "\n\t":
        switch (v19 % ((v19) | 1)) {
        case 2147483647:
            v17 = ((v17 - v17) - (0.0 / ((float32(0.0) - 0.0) + (float32(float64(float32((3.14159 * float64(float32(-2.25)))))) + (v17 * (((v17 / 1.5) * float32(3.14159)) + (((((float32((1e10 - -2.25)) * (0.1 / (v17 / 1e10))) / 0.1) / (v17 + -2.25)) / 1.0) * (float32((0.1 / (((3.14159 * 1e10) - (((1.5 / 1.0) * float64((v17 - (v17 / (v17 - 0.0))))) * (1.5 - 0.0))) * 0.0))) / v17))))))))
        default:
            v3[(^(v15))] = 65535
        }
    }
    if false {
        fmt.Println(v19)
    } else {
        delete(v2, (v15 / (((^(v15))) | 1)))
        delete(v2, 386)
    }
    v1[1] = v19
    var v21 int8 = -93
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        if (v16 <= v16) {
            switch (v19 ^ (-(fi22))) {
            case -672:
                v17 = v17
            case -443:
                v17 = v17
            case 32768:
                v23 := GenS3{F0: 0, F1: (true || true)}
                _ = v23
            default:
                for fi24 := 0; fi24 < 10; fi24++ {
                    _ = len(v1)
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    _ = fi24
                }
            }
            v15 = (v15 ^ (v15 | v15))
            fmt.Println(fi22, v16, v16)
        } else {
            v19 = v19
        }
        v19 = fi22
        _ = fi22
    }
    v25 := GenS4{F0: "a"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v25
    _ = v3
}
