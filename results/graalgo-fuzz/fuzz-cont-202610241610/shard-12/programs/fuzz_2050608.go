package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint16; F2 int16 }
func genFn0() int {
    var v6 int64 = 377
    v7 := []int{994, 655, 128, -598}
    var v8 uint16 = 1
    _ = v6
    _ = v7
    _ = v8
    return 93
}

func main() {
    var v1 float64 = float64(((3.14159 / 0.1) + -2.25))
    v2 := []string{((("foo bar" + ("foo bar" + "\n\t")) + "") + ("\n\t" + ("hello" + "\n\t"))), "", "\n\t", ((("\n\t" + "foo bar") + "foo bar") + ("foo bar" + (("a" + ("hello" + "hello")) + "hello")))}
    var v3 uint8 = 0
    v4 := make(map[int]int)
    var v5 int8 = -1
    switch 4294967295 {
    case 32767:
        var v9 string = "a"
        _ = v9
    case -628:
        var v10 float64 = 1e10
        _ = v10
    case -869:
        for fi11 := 0; fi11 < 14; fi11++ {
            for fi12 := 0; fi12 < 20; fi12++ {
                v1 = (v1 / -2.25)
                if (-2147483648 <= 1) {
                    v5 = v5
                    v13 := GenS1{F0: 562, F1: 42}
                    v3 = 0
                    fmt.Println(fi12, fi12)
                    _ = v13
                } else {
                    v14 := GenS1{F0: 65535, F1: 426}
                    var v15 float32 = (0.0 / 3.14159)
                    _ = v14
                    _ = v15
                }
                fmt.Println(fi12)
                _ = fi12
            }
            v1 += 1.5
            _ = fi11
        }
    default:
        switch ("a" + "hello") {
        case "hello":
            var v16 string = "foo bar"
            _ = v16
        case "foo bar":
            for fi17 := 0; fi17 < 14; fi17++ {
                v3 -= v3
                v18 := GenS1{F0: 690, F1: 42}
                v1 = 0.1
                _ = fi17
                _ = v18
            }
        case "":
            fmt.Println(v1, v3, v1)
        }
    }
    v19 := GenS2{F0: (^(v5))}
    fmt.Println(v5)
    switch "foo bar" {
    case "a":
        v20 := GenS3{F0: true, F1: 801, F2: v3}
        _ = v20
    default:
        var v21 string = ("hello" + (("a" + "\n\t") + "hello"))
        _ = v21
    }
    if false {
        v22 := GenS4{F0: 18446744073709551615, F1: 32767, F2: -210}
        var v23 int8 = 127
        _ = v22
        _ = v23
    } else {
        var v24 uint = 365
        _ = v24
    }
    v25 := GenS4{F0: 4294967295, F1: 1, F2: 32767}
    var v26 int64 = -9223372036854775808
    var v27 int8 = -1
    v1 = 0.0
    v5 = v5
    var v28 uint8 = (v3 % (((v3 * 255)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
