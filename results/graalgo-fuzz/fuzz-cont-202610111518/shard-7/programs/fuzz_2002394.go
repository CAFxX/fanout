package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int32; F1 int; F2 uint16 }
func genFn0() uint16 {
    var v7 float32 = 1e10
    if ((-2.25 <= ((-2.25 + 1e10) / float64(v7))) || false) {
        if false {
            var v8 int64 = -9223372036854775808
            var v9 string = "foo bar"
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            v7 = 1e10
            _ = fi10
        }
        if false {
            v11 := GenS1{F0: 42, F1: 1, F2: 127}
            switch "" {
            case "hello":
                var v12 string = "foo bar"
                _ = v12
            case "foo bar":
                var v13 uint = 42
                _ = v13
            default:
                v14 := GenS1{F0: -32768, F1: 1, F2: 128}
                _ = v14
            }
            v15 := []int{-2147483648, 128, -173}
            _ = v11
            _ = v15
        } else {
            if false {
                v7 = (0.0 + v7)
            } else {
                var v16 int32 = -93
                _ = v16
            }
            if true {
                fmt.Println(v7)
            }
            var v17 int32 = -1
            _ = v17
        }
    } else {
        var v18 bool = ("" == "\n\t")
        if (false || v18) {
            v7 = v7
        } else {
            v19 := GenS1{F0: 32767, F1: 0, F2: 0}
            _ = v19
        }
        _ = v18
    }
    v7 += float32(float64((((v7 - float32(((3.14159 / 1.0) * float64(((v7 / float32(float64(v7))) + (v7 / (v7 / float32(float64(v7))))))))) - (1.5 - 0.0)) * -2.25)))
    switch 9223372036854775807 {
    case -362:
        var v20 uint16 = 2
        _ = v20
    }
    var v21 int8 = -128
    _ = v21
    _ = v7
    return 1
}
func genFn1(p0 int64, p1 int) uint16 {
    if true {
        fmt.Println(p0, p1)
        if (0 != 897) {
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 uint = 828
                var v24 float32 = (-2.25 / -2.25)
                v25 := GenS1{F0: int16(p1), F1: 2, F2: 128}
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
        }
        fmt.Println(p1, p1)
    }
    var v26 uint32 = 2
    var v27 int16 = -1
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return 32767
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -277
    v3 := make(map[int]string)
    v4 := []int{0, 127, 9223372036854775807}
    var v5 float64 = 1.5
    var v6 uint16 = 0
    v5 -= 3.14159
    var v28 string = "a"
    v29 := GenS2{F0: (!false), F1: v6, F2: 539}
    var v30 uint64 = 0
    var v31 uint32 = 2147483648
    switch 1 {
    case 127:
        if len(v4) > 0 {
            v4[0] = -9223372036854775808
        }
    case 2147483647:
        v32 := GenS2{F0: (-209 >= 132), F1: 0, F2: v31}
        _ = v32
    case 128:
        for fi33 := 0; fi33 < 4; fi33++ {
            v6 += (^(v6))
            fmt.Println(v5)
            fmt.Println(v30)
            _ = fi33
        }
    }
    var v34 int8 = 127
    v35 := GenS3{F0: 42, F1: 779, F2: 255}
    if false {
        v2 = 2147483647
    } else {
        fmt.Println(v5, v2, v6)
        v3[4294967295] = v28
        if true {
            if (!true) {
                _ = len(v3)
                v36 := GenS2{F0: true, F1: genFn0(), F2: v31}
                fmt.Println(v5, v6)
                _ = v36
            }
            v5 = v5
            var v37 int32 = v2
            _ = v37
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
