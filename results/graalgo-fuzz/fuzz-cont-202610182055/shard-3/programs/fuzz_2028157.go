package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int }
type GenS2 struct { F0 int32; F1 string; F2 uint64 }
type GenS3 struct { F0 string; F1 int }

func main() {
    var v1 int16 = 363
    var v2 bool = true
    var v3 int = -1
    var v4 uint = 2
    var v5 uint64 = 18446744073709551615
    if (1 >= 53) {
        if v2 {
            v6 := GenS1{F0: true, F1: 575, F2: 32767}
            v7 := GenS1{F0: true, F1: 42, F2: 2147483647}
            _ = v6
            _ = v7
        } else {
            var v8 uint8 = 0
            for fi9 := 0; fi9 < 15; fi9++ {
                switch v3 {
                case -9223372036854775808:
                    var v10 string = (("hello" + "foo bar") + ("foo bar" + ("foo bar" + "foo bar")))
                    _ = v10
                case 32767:
                    v11 := GenS1{F0: true, F1: 406, F2: 32767}
                    _ = v11
                }
                var v12 uint8 = 42
                var v13 int32 = -1
                _ = fi9
                _ = v12
                _ = v13
            }
            _ = v8
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 int = 32767
            if (2147483647 > 229) {
                v16 := []int{2147483647, -432}
                v16 = append(v16, (-(v3)))
                var v17 uint8 = 127
                _ = v16
                _ = v17
            }
            var v18 int16 = v1
            _ = fi14
            _ = v15
            _ = v18
        }
        if (-50 < 42) {
            v19 := []int{}
            _ = v19
        } else {
            var v20 bool = v2
            if v20 {
                var v21 float64 = (1.5 / (((0.0 / float64(-2.25)) * (((0.1 * 1.5) - -2.25) - (-2.25 / 0.1))) + (3.14159 * (-2.25 + float64(1.5)))))
                v22 := GenS2{F0: 1, F1: "a", F2: v5}
                _ = v21
                _ = v22
            }
            _ = v20
        }
    } else {
        var v23 uint16 = 255
        if (v2 || v2) {
            switch ("hello" + "foo bar") {
            case "hello":
                var v24 int32 = 350
                _ = v24
            }
            switch ("foo bar" + "\n\t") {
            case "a":
                v5 += uint64(v23)
            default:
                var v25 uint32 = 2147483647
                _ = v25
            }
            v1 -= 0
        }
        _ = v23
    }
    v3 += v3
    v26 := GenS3{F0: ("foo bar" + "a"), F1: -9223372036854775808}
    fmt.Println(v1, v4, v1)
    v27 := GenS2{F0: -2147483648, F1: (("hello" + "a") + "foo bar"), F2: v5}
    v4 -= v4
    fmt.Println(v3, v1)
    v2 = ("a" < (("foo bar" + "foo bar") + "\n\t"))
    fmt.Println(v1, v5, v5)
    fmt.Println(v3, v1, v3)
    v28 := GenS3{F0: "a", F1: (^(v3))}
    v29 := []int{v3, v3, v3}
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    var v30 int16 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
