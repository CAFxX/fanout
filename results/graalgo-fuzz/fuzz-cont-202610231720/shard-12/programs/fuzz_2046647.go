package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int16; F1 int64 }
func genFn0(p0 uint8) int32 {
    fmt.Println(p0)
    if true {
        v5 := []int{0, 591}
        p0 += 163
        _ = v5
    }
    switch 32768 {
    case 128:
        for fi6 := 0; fi6 < 7; fi6++ {
            p0 = (p0 % (((^(p0))) | 1))
            var v7 bool = true
            v7 = v7
            _ = fi6
            _ = v7
        }
    case 2147483647:
        var v8 bool = true
        _ = v8
    case 0:
        if (688 != -9223372036854775808) {
            switch "" {
            case "a":
                if (1.5 == (0.0 / ((-2.25 - 0.0) + 1.5))) {
                    fmt.Println(p0)
                } else {
                    var v9 uint32 = 1
                    v10 := []int{42}
                    v9 = 928
                    _ = v10
                    _ = v9
                }
            }
        }
    default:
        fmt.Println(p0)
    }
    p0 += 42
    var v11 string = "a"
    _ = p0
    _ = v11
    return 1
}
func genFn1(p0 uint32, p1 int8, p2 uint) int {
    var v12 uint64 = 42
    p0 += p0
    switch 127 {
    case -2147483648:
        p0 += (^(p0))
    default:
        fmt.Println(p1, p2)
    }
    p1 = (p1 / ((int8(p2)) | 1))
    v13 := []int{-1, -9223372036854775808, -9223372036854775808, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return -2147483648
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 int16 = -280
    var v4 bool = (float64(-2.25) <= 3.14159)
    v14 := GenS1{F0: v3}
    var v15 int32 = genFn0(0)
    var v16 string = ("a" + ("foo bar" + ""))
    var v17 uint8 = 128
    _ = len(v1)
    for fi18 := 0; fi18 < 8; fi18++ {
        switch v16 {
        case "":
            for fi19 := 0; fi19 < 1; fi19++ {
                fmt.Println(v15, fi19, v4)
                _ = fi19
            }
        }
        v17 += (v17 + 255)
        v15 += v15
        _ = fi18
    }
    var v20 int32 = genFn0((v17 * (^(v17))))
    v21 := GenS2{F0: "foo bar", F1: 128}
    switch "a" {
    case "":
        switch 0 {
        case 2147483647:
            v22 := GenS3{F0: v3, F1: -763}
            _ = v22
        case 4294967295:
            if (v4 || (true && (v4 || (!v4)))) {
                v23 := []int{9223372036854775807, 32767}
                var v24 int8 = -128
                var v25 bool = v4
                _ = v23
                _ = v24
                _ = v25
            }
        case 42:
            var v26 uint = 42
            _ = v26
        }
    case "foo bar":
        v27 := []int{-487, -859, 32767}
        _ = v27
    case "\n\t":
        var v28 int8 = -42
        _ = v28
    }
    var v29 float32 = 3.14159
    if v4 {
        v15 += v20
    } else {
        var v30 uint = 1
        var v31 float32 = ((float32(3.14159) - ((v29 + v29) + ((((-2.25 + float32(0.1)) / v29) / (v29 * v29)) / (((v29 * float32(-2.25)) - 1e10) - v29)))) * 0.0)
        _ = v30
        _ = v31
    }
    if (0 < 255) {
        v32 := GenS2{F0: ("foo bar" + v16), F1: (v17 & (v17 / ((v17) | 1)))}
        var v33 float32 = v29
        _ = v32
        _ = v33
    } else {
        v15 = 0
        v3 = (v3 >> ((368) & 63))
    }
    var v34 uint8 = 127
    var v35 int16 = (v3 << ((0) & 63))
    fmt.Println(v34, v29)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
