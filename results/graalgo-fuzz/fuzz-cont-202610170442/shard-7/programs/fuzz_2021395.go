package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }

func main() {
    var v1 float64 = float64(3.14159)
    var v2 int16 = 0
    v3 := make(map[int]bool)
    fmt.Println(v1)
    var v4 int = -2147483648
    v5 := GenS1{F0: (v2 << ((1) & 63)), F1: 1}
    switch (v4 & v4) {
    case -1:
        v6 := GenS1{F0: (v2 ^ (v2 & 1000)), F1: -1}
        _ = v6
    case -304:
        v7 := GenS1{F0: (v2 & -1), F1: -128}
        _ = v7
    default:
        var v8 uint32 = 2
        _ = v8
    }
    _ = len(v3)
    var v9 uint16 = 843
    v10 := []int{(v4 | v4)}
    var v11 uint64 = 0
    if (!false) {
        if true {
            switch "foo bar" {
            case "":
                if false {
                    for fi12 := 0; fi12 < 9; fi12++ {
                        fmt.Println(fi12, v4, v4)
                        var v13 int64 = 9223372036854775807
                        v2 -= v2
                        _ = fi12
                        _ = v13
                    }
                    var v14 int64 = 0
                    _ = v14
                } else {
                    if (!true) {
                        var v15 uint16 = (v9 - 882)
                        fmt.Println(v1, v1)
                        _ = v15
                    } else {
                        v16 := GenS1{F0: (v2 >> ((0) & 63)), F1: 127}
                        _ = v16
                    }
                }
            case "hello":
                switch "" {
                case "foo bar":
                    fmt.Println(v1, v11)
                case "hello":
                    v17 := v3[(v4 - v4)]
                    _ = v17
                }
            }
        }
        var v18 int = (v4 % (((v4 % ((1) | 1))) | 1))
        switch (("hello" + (((((("\n\t" + "\n\t") + "hello") + "") + "a") + "") + "\n\t")) + "\n\t") {
        case "hello":
            if true {
                v1 -= (v1 / v1)
            } else {
                for fi19 := 0; fi19 < 10; fi19++ {
                    v2 += v2
                    if true {
                        var v20 float64 = v1
                        for fi21 := 0; fi21 < 7; fi21++ {
                            v9 -= (v9 % (((v9 * (v9 - v9))) | 1))
                            v2 -= 1
                            _ = fi21
                        }
                        _ = v20
                    } else {
                        v22 := GenS1{F0: v2, F1: 79}
                        var v23 float32 = 1.5
                        var v24 string = ""
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    var v25 uint32 = 432
                    _ = fi19
                    _ = v25
                }
            }
        case "a":
            var v26 uint16 = (v9 >> ((255) & 63))
            _ = v26
        case "\n\t":
            if len(v10) > 0 {
                v10[0] = v4
            }
        }
        _ = v18
    } else {
        v11 += v11
        if (1 != 42) {
            var v27 uint64 = (v11 + (v11 + (v11 % (((v11 + v11)) | 1))))
            var v28 int8 = -1
            _ = v27
            _ = v28
        }
    }
    v1 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
