package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 string }
func genFn0(p0 uint32, p1 bool, p2 uint8) int32 {
    if p1 {
        switch 0 {
        case -9223372036854775808:
            p1 = true
        case -539:
            v7 := GenS1{F0: 255, F1: 0, F2: 0}
            _ = v7
        default:
            var v8 uint16 = 210
            _ = v8
        }
    }
    switch "foo bar" {
    case "\n\t":
        switch 32767 {
        case -9223372036854775808:
            fmt.Println(p0, p1, p0)
        case 0:
            v9 := []int{}
            _ = v9
        }
    case "foo bar":
        switch "" {
        case "foo bar":
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := []int{(fi10 | 128), fi10, -1}
                if (p1 || p1) {
                    fmt.Println(p1, p1)
                    v12 := GenS1{F0: 407, F1: 1, F2: fi10}
                    _ = v12
                } else {
                    p1 = p1
                }
                var v13 bool = false
                _ = fi10
                _ = v11
                _ = v13
            }
        case "":
            v14 := []int{32767, -9223372036854775808, 1}
            _ = v14
        }
    }
    var v15 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 37
}
func genFn1(p0 uint) uint16 {
    var v16 uint32 = 278
    if true {
        fmt.Println(v16, v16)
    }
    _ = p0
    _ = v16
    return 65535
}

func main() {
    var v1 int64 = 0
    var v2 uint8 = 61
    v3 := []uint32{1, 4294967295, 803}
    v4 := make(map[uint32]int)
    v5 := make(map[string]string)
    var v6 uint8 = 0
    var v17 float64 = 1e10
    v18 := GenS2{F0: -144}
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 string = "a"
        var v21 float64 = (-2.25 * v17)
        for fi22 := 0; fi22 < 17; fi22++ {
            v23 := GenS3{F0: ("" + "foo bar")}
            v4[2] = fi22
            switch v20 {
            case "hello":
                v24 := v5[(v20 + v20)]
                _ = v24
            case "a":
                var v25 int32 = 25
                _ = v25
            default:
                if true {
                    var v26 uint8 = 42
                    _ = v26
                } else {
                    var v27 uint8 = v2
                    _ = v27
                }
            }
            _ = fi22
            _ = v23
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    v2 = v6
    if (1.5 == (3.14159 * (float64((((((1.5 + 3.14159) / 3.14159) / ((1e10 * 1e10) + -2.25)) + 1.5) + 1e10)) - ((v17 + v17) / (v17 + ((0.0 - v17) * 3.14159)))))) {
        v28 := GenS1{F0: 32767, F1: 172, F2: 665}
        v29 := GenS1{F0: 65535, F1: 973, F2: 127}
        _ = v28
        _ = v29
    } else {
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        if (!(-9223372036854775808 >= -9223372036854775808)) {
            switch "hello" {
            case "a":
                if ("" == "a") {
                    v1 = (v1 * 0)
                } else {
                    var v30 int16 = 1000
                    var v31 uint16 = 298
                    v32 := GenS3{F0: ("a" + (((("hello" + "") + "a") + "hello") + ""))}
                    _ = v30
                    _ = v31
                    _ = v32
                }
            case "\n\t":
                if (v17 == (1.5 - float64(((0.1 / (0.0 * float32(v17))) * 1.5)))) {
                    v6 = v6
                }
            }
            if false {
                v6 -= v6
                v33 := GenS3{F0: ("a" + "")}
                _ = v33
            } else {
                _ = len(v3)
            }
        }
    }
    if true {
        v34 := GenS4{F0: 127, F1: ("a" + ("foo bar" + ("a" + "hello")))}
        _ = v34
    } else {
        v2 += v2
        var v35 float32 = (1.5 + (-2.25 / 1.5))
        if true {
            var v36 int64 = (v1 >> ((0) & 63))
            v2 += (^(v2))
            _ = v36
        }
        _ = v35
    }
    var v37 int64 = 0
    v6 = 127
    fmt.Println(v17, v37, v2)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
