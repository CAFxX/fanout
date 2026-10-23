package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }
func genFn0(p0 int32, p1 int) int16 {
    if (42 < 0) {
        if true {
            switch p1 {
            case 268:
                v6 := GenS1{F0: 389, F1: 42, F2: ("foo bar" + "hello")}
                _ = v6
            }
            for fi7 := 0; fi7 < 14; fi7++ {
                switch ((("\n\t" + "a") + "hello") + ((("foo bar" + (("hello" + (("hello" + ("a" + "")) + "\n\t")) + ("hello" + ""))) + "a") + "foo bar")) {
                case "foo bar":
                    v8 := []int{p1}
                    _ = v8
                case "\n\t":
                    var v9 float64 = 3.14159
                    _ = v9
                case "hello":
                    var v10 int = (p1 ^ (fi7 >> ((0) & 63)))
                    _ = v10
                default:
                    v11 := []int{p1}
                    _ = v11
                }
                var v12 uint8 = 127
                fmt.Println(v12, p1)
                _ = fi7
                _ = v12
            }
        }
    } else {
        p0 -= (p0 * p0)
    }
    switch (p1 ^ 42) {
    case 32768:
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    return -843
}

func main() {
    var v1 int8 = 75
    var v2 int16 = 31
    var v3 float32 = ((1.5 + ((((0.1 / (0.1 / ((((((1e10 - -2.25) - 0.0) / 1.5) - float32((3.14159 * 1e10))) / 3.14159) / 0.1))) - float32(3.14159)) / 3.14159) / ((0.1 / 1.0) * 1e10))) / 1e10)
    var v4 int = 32767
    var v5 uint8 = 0
    v1 += -64
    switch "" {
    case "foo bar":
        if false {
            v4 = 1
            var v13 uint32 = 4294967295
            _ = v13
        }
    case "\n\t":
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 uint16 = 2
            fmt.Println(v4, fi14, v4)
            switch fi14 {
            case 2147483647:
                var v16 int16 = (-(v2))
                _ = v16
            default:
                v17 := GenS1{F0: 4294967295, F1: 4294967295, F2: "a"}
                _ = v17
            }
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(v3, v3, v1)
    for fi18 := 0; fi18 < 15; fi18++ {
        fmt.Println(v4)
        _ = fi18
    }
    fmt.Println(v3, v2, v5)
    var v19 uint64 = 925
    v2 += v2
    v20 := []int{(v4 & v4), v4, (v4 / (((v4 - 9223372036854775807)) | 1)), v4}
    var v21 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
