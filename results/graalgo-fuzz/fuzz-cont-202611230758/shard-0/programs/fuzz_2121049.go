package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int; F2 uint }
type GenS4 struct { F0 uint32; F1 int64; F2 string }
type GenS5 struct { F0 bool }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 uint = 0
        v5 = v5
        _ = fi4
        _ = v5
    }
    v6 := []int{42, 1}
    var v7 int = 2147483647
    switch (v7 >> ((463) & 63)) {
    case 127:
        var v8 string = "foo bar"
        _ = v8
    case -783:
        v7 -= (^(v7))
    default:
        var v9 uint32 = 42
        _ = v9
    }
    var v10 string = ("\n\t" + "")
    _ = v10
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 string, p1 int, p2 uint8) uint16 {
    if true {
        v11 := GenS1{F0: 42, F1: -9223372036854775808, F2: 100}
        _ = v11
    }
    v12 := []int{42, p1, (p1 | p1), p1}
    _ = len(v12)
    switch p1 {
    case 613:
        var v13 int64 = 0
        _ = v13
    case 9223372036854775807:
        v14 := GenS1{F0: 107, F1: -927, F2: -109}
        _ = v14
    case 32768:
        if len(v12) > 0 {
            v12[0] = p1
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 333
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 0.1
    var v3 uint = 42
    v2 -= 0.1
    for fi15 := 0; fi15 < 18; fi15++ {
        if false {
            v1[(fi15 + 423)] = fi15
            var v16 uint = (v3 + 800)
            _ = v16
        } else {
            var v17 int8 = -12
            var v18 uint64 = 2
            for fi19 := 0; fi19 < 15; fi19++ {
                v20 := GenS1{F0: 42, F1: -1, F2: v17}
                fmt.Println(v3, fi19, v17)
                _ = fi19
                _ = v20
            }
            _ = v17
            _ = v18
        }
        fmt.Println(v3)
        _ = fi15
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := GenS2{F0: 2147483647}
        v1[127] = (fi21 >> ((v3) & 63))
        switch fi21 {
        case -2147483648:
            v23 := v1[0]
            _ = v23
        case 127:
            fmt.Println(v2)
        case 9223372036854775807:
            var v24 float64 = 3.14159
            _ = v24
        default:
            if (((true && true) || false) || true) {
                v2 -= (float32(3.14159) + v2)
            }
        }
        _ = fi21
        _ = v22
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        if (-1 != 0) {
            v26 := GenS3{F0: v3, F1: 632, F2: 18446744073709551615}
            _ = v26
        }
        v27 := GenS4{F0: 42, F1: int64(v3), F2: ("foo bar" + "hello")}
        if (1 > -1) {
            v28 := GenS1{F0: uint32(v3), F1: 841, F2: 42}
            _ = v28
        }
        _ = fi25
        _ = v27
    }
    var v29 string = ""
    switch v29 {
    case "a":
        fmt.Println(v2, v3)
    case "hello":
        fmt.Println(v29, v29, v29)
    case "foo bar":
        fmt.Println(v29, v3, v3)
    default:
        v29 = (v29 + v29)
    }
    v2 = float32(((-2.25 / 1e10) / ((float64(float32(1.5)) + 3.14159) / 1.0)))
    if (!(98 != 42)) {
        if (!(-2.25 >= 0.1)) {
            var v30 float32 = (v2 / v2)
            _ = v30
        }
        for fi31 := 0; fi31 < 9; fi31++ {
            v32 := GenS1{F0: 2147483647, F1: int64(v3), F2: 61}
            _ = fi31
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        v2 = (v2 + v2)
        v34 := []int{}
        v35 := GenS5{F0: true}
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 uint8 = genFn0()
    fmt.Println(v3, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v36
}
