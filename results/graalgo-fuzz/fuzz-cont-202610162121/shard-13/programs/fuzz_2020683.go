package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    var v4 int16 = 0
    var v5 int8 = 74
    v6 := make(map[string]bool)
    v5 -= 100
    var v7 uint64 = 4294967295
    if ((4294967295 <= v7) || ((1 >= 210) || ((v7 & 0) < 42))) {
        _ = len(v1)
        v8 := GenS1{F0: (v5 >> ((281) & 63)), F1: 1, F2: 1}
        for fi9 := 0; fi9 < 7; fi9++ {
            switch 32767 {
            case 9223372036854775807:
                if ((!("\n\t" <= ("" + ((("" + (("" + "") + ((("hello" + "") + "foo bar") + "foo bar"))) + "foo bar") + "a")))) || false) {
                    var v10 int16 = -919
                    _ = v10
                } else {
                    var v11 uint = 42
                    _ = v11
                }
            case -1:
                var v12 float64 = ((0.1 / (0.1 - (((0.1 * 1.5) / float64(3.14159)) + (1e10 + 3.14159)))) + (3.14159 / 1e10))
                _ = v12
            case 0:
                var v13 bool = (904 == 403)
                _ = v13
            }
            v14 := GenS2{F0: 4294967295, F1: 0}
            v15 := GenS1{F0: v5, F1: 1, F2: (^(fi9))}
            _ = fi9
            _ = v14
            _ = v15
        }
        _ = v8
    }
    var v16 int16 = v4
    switch ("" + ("foo bar" + "foo bar")) {
    case "\n\t":
        v17 := []int{}
        _ = v17
    case "hello":
        v16 = int16(v7)
    case "a":
        var v18 string = (("a" + "a") + "a")
        _ = v18
    }
    fmt.Println(v7, v4, v7)
    v19 := GenS3{F0: 255}
    for fi20 := 0; fi20 < 16; fi20++ {
        _ = len(v3)
        fmt.Println(fi20, v16, fi20)
        v5 = int8(v16)
        _ = fi20
    }
    if (!(false || false)) {
        switch ("\n\t" + "") {
        case "\n\t":
            fmt.Println(v16)
        }
        delete(v6, "")
    } else {
        if true {
            for fi21 := 0; fi21 < 12; fi21++ {
                switch "a" {
                case "\n\t":
                    v22 := []int{(fi21 - (fi21 * fi21)), -615, (fi21 >> ((uint(v7)) & 63))}
                    _ = v22
                case "foo bar":
                    fmt.Println(fi21, v16, v7)
                default:
                    v3["hello"] = 2147483647
                }
                _ = fi21
            }
            var v23 float32 = 1e10
            _ = v23
        }
        fmt.Println(v4, v16)
    }
    var v24 int8 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
