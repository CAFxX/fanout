package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint; F2 int }
type GenS3 struct { F0 bool; F1 bool; F2 int }

func main() {
    var v1 uint16 = 42
    var v2 uint64 = 731
    v3 := []string{"a"}
    var v4 string = "foo bar"
    var v5 int64 = -378
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(v4)
        _ = fi6
    }
    if true {
        v7 := GenS1{F0: v4, F1: 42, F2: 128}
        fmt.Println(v1)
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        if false {
            if false {
                if (true && false) {
                    if (float64(0.1) > float64(-2.25)) {
                        v1 += (v1 | uint16(v2))
                    }
                    switch ((v4 + v4) + (v4 + "\n\t")) {
                    case "foo bar":
                        if ((-2.25 - ((0.1 * (float32(3.14159) - 3.14159)) + -2.25)) >= (0.0 / (1.5 + (1.5 + -2.25)))) {
                            if (!false) {
                                var v9 int16 = int16(v5)
                                _ = v9
                            }
                        }
                    default:
                        v4 = ("" + v4)
                    }
                    v10 := GenS1{F0: "a", F1: 0, F2: 42}
                    _ = v10
                }
                var v11 uint = 203
                _ = v11
            }
        } else {
            if ((255 >= 260) || true) {
                fmt.Println(fi8)
                v4 = (v4 + "a")
            }
            var v12 bool = true
            _ = v12
        }
        v2 = uint64(v1)
        switch v4 {
        case "":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "a":
            if (!(v4 < "")) {
                if (v5 > (v5 >> ((457) & 63))) {
                    fmt.Println(fi8, v4)
                    var v13 int64 = 383
                    fmt.Println(v1, v2, v1)
                    _ = v13
                } else {
                    v14 := GenS1{F0: "a", F1: -1, F2: 82}
                    _ = v14
                }
                v4 = ((v4 + "a") + v4)
            } else {
                v4 = v4
                v2 = 42
            }
        default:
            _ = len(v3)
        }
        _ = fi8
    }
    v1 = v1
    fmt.Println(v1, v5)
    v4 = (v4 + "\n\t")
    v15 := GenS2{F0: 2, F1: uint(v2), F2: 1}
    v2 += v2
    if (722 < -1) {
        v16 := GenS3{F0: ((v4 + ("a" + (v4 + v4))) < v4), F1: true, F2: -2147483648}
        _ = v16
    }
    var v17 int8 = 61
    var v18 string = "\n\t"
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
