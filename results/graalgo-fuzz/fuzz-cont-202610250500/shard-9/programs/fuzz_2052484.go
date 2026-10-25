package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 string; F2 int16 }
type GenS4 struct { F0 int8; F1 int32; F2 int16 }
func genFn0(p0 int64, p1 int8) int8 {
    v7 := GenS1{F0: 255, F1: p0, F2: -32768}
    if true {
        v8 := []int{}
        v9 := GenS2{F0: (p0 | p0)}
        _ = v8
        _ = v9
    } else {
        if (365 > 23) {
            p0 = (p0 & p0)
            if ((("a" + ("hello" + "a")) + "") <= "foo bar") {
                var v10 bool = (0.0 >= (0.1 / -2.25))
                v10 = v10
                _ = v10
            }
            v11 := GenS2{F0: p0}
            _ = v11
        }
        switch ("" + "a") {
        case "hello":
            if ((!true) && false) {
                p1 = p1
                v12 := []int{int(p0), 42}
                _ = v12
            } else {
                for fi13 := 0; fi13 < 9; fi13++ {
                    v14 := GenS2{F0: 1}
                    _ = fi13
                    _ = v14
                }
            }
        case "\n\t":
            switch "hello" {
            case "foo bar":
                for fi15 := 0; fi15 < 1; fi15++ {
                    p0 -= 1
                    _ = fi15
                }
            case "a":
                fmt.Println(p1, p0)
            default:
                var v16 uint = 1
                _ = v16
            }
        default:
            v17 := []int{2147483647, 4294967295, 4294967295, 32767}
            _ = v17
        }
    }
    var v18 float64 = (0.1 / float64((float32(float64(float32(((0.1 + 1e10) * 0.1)))) - float32(float64(3.14159)))))
    fmt.Println(v18, v18)
    _ = p0
    _ = p1
    _ = v18
    _ = v7
    return (p1 & (p1 - p1))
}

func main() {
    var v1 uint64 = 0
    var v2 int32 = 1
    var v3 uint8 = 128
    v4 := make(map[string]uint32)
    v5 := make(map[string]int)
    v6 := make(map[uint32]string)
    v2 += -2147483648
    if (314 < -1) {
        v1 += 956
    } else {
        var v19 int32 = v2
        v20 := v4[("hello" + ((("foo bar" + "foo bar") + (((("foo bar" + "\n\t") + "a") + "a") + "a")) + "a"))]
        v21 := GenS3{F0: 42, F1: "foo bar", F2: 32767}
        _ = v19
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: 1}
    for fi23 := 0; fi23 < 4; fi23++ {
        v24 := []int{-679, (fi23 * fi23)}
        _ = fi23
        _ = v24
    }
    switch -164 {
    case 32767:
        switch 32767 {
        case -459:
            for fi25 := 0; fi25 < 10; fi25++ {
                fmt.Println(fi25)
                switch "\n\t" {
                case "hello":
                    var v26 int8 = genFn0(-1, 42)
                    _ = v26
                default:
                    var v27 float32 = -2.25
                    _ = v27
                }
                _ = fi25
            }
        case 32767:
            for fi28 := 0; fi28 < 20; fi28++ {
                _ = len(v6)
                _ = fi28
            }
        }
    case 732:
        var v29 float32 = (1.5 + 0.0)
        _ = v29
    default:
        var v30 int32 = (v2 - -392)
        _ = v30
    }
    if false {
        v1 -= (v1 >> ((2) & 63))
        var v31 int8 = 100
        _ = v31
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        v1 = v1
        var v33 uint32 = 149
        v34 := GenS4{F0: 100, F1: v2, F2: -32768}
        _ = fi32
        _ = v33
        _ = v34
    }
    if (255 >= 65535) {
        var v35 float64 = float64(1.5)
        _ = v35
    }
    _ = len(v5)
    v36 := GenS1{F0: v3, F1: 9223372036854775807, F2: 32767}
    switch 430 {
    case 175:
        fmt.Println(v2, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
