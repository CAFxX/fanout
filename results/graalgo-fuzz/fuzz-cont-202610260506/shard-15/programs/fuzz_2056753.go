package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 float32 = 1e10
        v7 -= v7
        var v8 int8 = 127
        _ = fi6
        _ = v7
        _ = v8
    }
    switch ("foo bar" + ("hello" + "a")) {
    case "a":
        var v9 int64 = 0
        _ = v9
    case "foo bar":
        var v10 float32 = (0.0 * 0.1)
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 uint64 = 246
            v13 := GenS1{F0: -109, F1: false}
            var v14 float32 = (0.1 * 1e10)
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    default:
        v15 := GenS1{F0: 0, F1: (707 != 42)}
        _ = v15
    }
    var v16 uint16 = 65535
    v16 = 523
    if (true || false) {
        var v17 uint64 = 755
        v18 := []int{0, 42, 195}
        var v19 int32 = 758
        _ = v17
        _ = v18
        _ = v19
    }
    _ = v16
    return 2
}
func genFn1(p0 uint16) int32 {
    fmt.Println(p0)
    var v20 float32 = 3.14159
    switch 2147483647 {
    case -346:
        v21 := GenS1{F0: -2147483648, F1: (1 == 234)}
        _ = v21
    case 363:
        fmt.Println(v20)
    default:
        switch -908 {
        case 2147483647:
            if (!(!((v20 + float32(-2.25)) <= (v20 / v20)))) {
                var v22 int8 = 1
                var v23 int32 = 1
                v24 := []int{-1, 0, 32767, -549}
                _ = v22
                _ = v23
                _ = v24
            }
        case 128:
            fmt.Println(v20, v20)
        default:
            var v25 string = "\n\t"
            _ = v25
        }
    }
    _ = p0
    _ = v20
    return 832
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[uint32]int)
    var v3 int16 = -1
    v4 := []string{"\n\t"}
    v5 := []string{("" + "hello"), "a", "foo bar"}
    if (!false) {
        v3 += (-(v3))
        v1 = (^(v1))
    } else {
        for fi26 := 0; fi26 < 2; fi26++ {
            v3 += (v3 % ((-897) | 1))
            for fi27 := 0; fi27 < 3; fi27++ {
                v28 := GenS2{F0: 18446744073709551615}
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
        fmt.Println(v3, v3)
    }
    v29 := GenS3{F0: 65535, F1: -1}
    switch ("\n\t" + ((("" + "foo bar") + "\n\t") + "\n\t")) {
    case "foo bar":
        var v30 int16 = 1000
        _ = v30
    default:
        var v31 int = -912
        _ = v31
    }
    switch -2147483648 {
    case 474:
        var v32 int64 = -9223372036854775808
        _ = v32
    default:
        v1 = v1
    }
    if false {
        v1 = (^(v1))
    } else {
        var v33 uint64 = 354
        if len(v4) > 0 {
            v4[0] = ("hello" + ("\n\t" + ("foo bar" + "\n\t")))
        }
        _ = v33
    }
    v1 = 98
    for fi34 := 0; fi34 < 3; fi34++ {
        var v35 float32 = 0.1
        v3 = v3
        v35 = v35
        _ = fi34
        _ = v35
    }
    var v36 int64 = 9223372036854775807
    fmt.Println(v36, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
