package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint16; F1 int; F2 int }
type GenS3 struct { F0 uint8; F1 int8; F2 string }
func genFn0() uint32 {
    v6 := []int{32767, 9223372036854775807, 128}
    var v7 int8 = 42
    _ = v6
    _ = v7
    return 139
}
func genFn1(p0 int64, p1 uint64, p2 int16) uint16 {
    var v8 string = (("hello" + "a") + "")
    if (!(!(!(!true)))) {
        var v9 uint64 = (^(p1))
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint = 166
            _ = fi10
            _ = v11
        }
        if (-2147483648 == 4294967295) {
            p1 = (p1 >> ((4294967295) & 63))
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 255
}

func main() {
    var v1 float64 = float64(0.0)
    v2 := make(map[uint32]int)
    var v3 bool = (42 > 255)
    v4 := make(map[uint32]string)
    var v5 int16 = 1000
    fmt.Println(v3, v1)
    v12 := GenS1{F0: 32767, F1: 28}
    if false {
        if (990 <= 2147483647) {
            var v13 int8 = 127
            fmt.Println(v5, v1, v1)
            _ = v13
        } else {
            switch 1 {
            case -1:
                var v14 uint64 = 461
                _ = v14
            case -9223372036854775808:
                v15 := GenS1{F0: -2147483648, F1: 42}
                _ = v15
            }
        }
        v1 = (((float64(3.14159) * v1) - 0.0) * v1)
        v16 := GenS2{F0: 116, F1: 42, F2: -457}
        _ = v16
    }
    fmt.Println(v5)
    v5 = (v5 << ((255) & 63))
    if v3 {
        var v17 string = ""
        var v18 int64 = 1
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 4; fi19++ {
            fmt.Println(fi19, v5, fi19)
            var v20 string = ("foo bar" + ((("" + "\n\t") + ("foo bar" + "a")) + (((("" + "hello") + "hello") + "hello") + "")))
            v1 = ((float64(0.0) - (3.14159 + float64(0.1))) + v1)
            _ = fi19
            _ = v20
        }
        v1 += float64((0.1 - 3.14159))
    }
    var v21 int8 = 42
    var v22 int16 = (v5 % (((^(v5))) | 1))
    var v23 float32 = 0.1
    var v24 uint16 = genFn1(-9223372036854775808, 4294967295, v5)
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 float64 = v1
        _ = fi25
        _ = v26
    }
    v24 += genFn1(-232, 4294967295, 2)
    v22 = 896
    var v27 float32 = v23
    v28 := GenS3{F0: 0, F1: v21, F2: "a"}
    var v29 int8 = v21
    switch ("a" + "foo bar") {
    case "hello":
        fmt.Println(v27, v1, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
