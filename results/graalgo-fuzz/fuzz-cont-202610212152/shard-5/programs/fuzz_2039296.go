package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int }
func genFn0(p0 int16, p1 uint, p2 string) uint16 {
    if false {
        if ((1.5 / float32(((0.0 + (3.14159 * -2.25)) / 3.14159))) > 0.0) {
            v7 := GenS1{F0: -2147483648, F1: -998, F2: 597}
            _ = v7
        } else {
            var v8 uint8 = 127
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 float32 = (float32(1.5) - 0.1)
                switch "\n\t" {
                case "foo bar":
                    var v11 int16 = (p0 >> ((p1) & 63))
                    _ = v11
                case "hello":
                    v12 := GenS2{F0: 4294967295}
                    _ = v12
                }
                v13 := []int{}
                _ = fi9
                _ = v10
                _ = v13
            }
            _ = v8
        }
    }
    v14 := GenS2{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 255
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = (float64(3.14159) - 0.0)
    var v3 uint16 = 168
    var v4 uint = 942
    var v5 uint32 = 1
    var v6 float64 = (float64(0.0) * 1.5)
    switch "hello" {
    case "foo bar":
        switch 0 {
        case 127:
            if (-9223372036854775808 == -1) {
                for fi15 := 0; fi15 < 15; fi15++ {
                    v16 := []int{fi15, (fi15 >> (((v4 * (^(v4)))) & 63))}
                    fmt.Println(v2, fi15)
                    v17 := v1[((("hello" + ("foo bar" + "")) + "foo bar") + "foo bar")]
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            } else {
                if ((!false) && true) {
                    v18 := []int{-9223372036854775808}
                    fmt.Println(v6, v3)
                    v19 := GenS1{F0: 9223372036854775807, F1: 634, F2: v5}
                    _ = v18
                    _ = v19
                } else {
                    var v20 uint = (v4 % ((v4) | 1))
                    _ = v20
                }
                for fi21 := 0; fi21 < 12; fi21++ {
                    var v22 uint64 = 42
                    var v23 float64 = -2.25
                    v2 = (v6 * 0.0)
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
        case 696:
            delete(v1, "\n\t")
        }
    default:
        _ = len(v1)
    }
    switch "" {
    case "a":
        var v24 int32 = int32(v3)
        _ = v24
    case "foo bar":
        v25 := GenS1{F0: 32767, F1: -32768, F2: (v5 / ((v5) | 1))}
        _ = v25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        v27 := []int{-713, (fi26 - fi26), -9223372036854775808, 2147483647}
        var v28 string = "a"
        v29 := GenS1{F0: (fi26 - 622), F1: -488, F2: v5}
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        v31 := []int{fi30, -874}
        _ = fi30
        _ = v31
    }
    v5 = 1
    v5 += 4294967295
    v3 += v3
    if false {
        v32 := GenS3{F0: 127, F1: "hello"}
        v4 = v4
        var v33 uint16 = v3
        _ = v32
        _ = v33
    }
    fmt.Println(v6, v6)
    var v34 string = "hello"
    if (!(-9223372036854775808 <= 1)) {
        var v35 string = "\n\t"
        v36 := GenS4{F0: 255}
        _ = v35
        _ = v36
    } else {
        delete(v1, v34)
        fmt.Println(v4)
    }
    var v37 int = -293
    v38 := GenS5{F0: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
