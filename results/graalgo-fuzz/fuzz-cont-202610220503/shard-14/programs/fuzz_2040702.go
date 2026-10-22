package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int64; F2 int }
type GenS3 struct { F0 bool; F1 string }
func genFn0(p0 uint16, p1 int16) uint16 {
    fmt.Println(p1)
    switch 9223372036854775807 {
    case -9223372036854775808:
        switch 4294967295 {
        case 24:
            p1 = (^(p1))
        default:
            p0 += p0
        }
    case -2147483648:
        if ((-2147483648 < 0) || (float32(1.5) <= (1.5 + (((0.1 - (0.0 + -2.25)) + ((0.1 / (-2.25 + (1e10 / 1.0))) + -2.25)) - 1e10)))) {
            p1 = p1
            switch 4294967295 {
            case 9223372036854775807:
                v6 := GenS1{F0: 100}
                _ = v6
            case 32768:
                switch "a" {
                case "hello":
                    v7 := GenS1{F0: 42}
                    _ = v7
                case "":
                    var v8 int16 = (p1 - p1)
                    _ = v8
                case "\n\t":
                    p0 -= 913
                }
            case 127:
                var v9 int = 32768
                _ = v9
            }
            var v10 uint = 0
            _ = v10
        } else {
            v11 := GenS2{F0: (^(p0)), F1: 960, F2: 128}
            var v12 float64 = float64(float32((float64(0.0) * 1e10)))
            _ = v11
            _ = v12
        }
    case 385:
        v13 := []int{2147483647, 128, -704, -1}
        _ = v13
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 2147483647
    var v3 uint8 = 220
    v4 := make(map[uint32]string)
    var v5 uint8 = (v3 & 38)
    v2 = (^(v2))
    var v14 float32 = 1.5
    v5 = v5
    v2 = (v2 + v2)
    v3 = (v3 | 128)
    _ = len(v1)
    fmt.Println(v14, v14, v14)
    delete(v4, 23)
    v14 = v14
    switch "\n\t" {
    case "foo bar":
        var v15 uint = 971
        _ = v15
    case "a":
        v2 += 42
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        v3 = (^(v3))
        _ = fi16
    }
    switch (((("" + "") + (("hello" + "") + "foo bar")) + "hello") + "hello") {
    case "":
        switch "a" {
        case "a":
            var v17 int = (v2 - int(v5))
            _ = v17
        }
    case "foo bar":
        if true {
            v18 := GenS3{F0: (true && (!false)), F1: "a"}
            _ = v18
        } else {
            v19 := []int{(v2 >> ((2) & 63)), v2, v2, 181}
            _ = v19
        }
    case "\n\t":
        for fi20 := 0; fi20 < 5; fi20++ {
            switch ("" + "hello") {
            case "\n\t":
                v14 = v14
            case "foo bar":
                v2 = 127
            }
            var v21 uint64 = 2
            _ = fi20
            _ = v21
        }
    default:
        v2 -= v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
