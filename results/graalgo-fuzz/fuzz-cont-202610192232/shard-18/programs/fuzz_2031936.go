package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    var v4 float32 = (-2.25 - float32(1e10))
    var v5 int16 = -1
    if (!(true && ((("hello" + "") + ((("hello" + "hello") + "\n\t") + "foo bar")) >= "hello"))) {
        var v6 float64 = -2.25
        v7 := GenS1{F0: -1, F1: 65535}
        _ = v6
        _ = v7
    } else {
        v5 += v5
        for fi8 := 0; fi8 < 9; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                _ = len(v3)
                fmt.Println(v5, v4, v4)
                v4 += 1.5
                _ = fi9
            }
            var v10 float32 = float32(1e10)
            _ = fi8
            _ = v10
        }
    }
    v5 = 1000
    switch ("\n\t" + "hello") {
    case "foo bar":
        v5 = (v5 % ((v5) | 1))
    }
    if (!true) {
        var v11 float64 = float64(v4)
        switch 128 {
        case 666:
            v11 = v11
        default:
            switch ("\n\t" + "hello") {
            case "a":
                v12 := GenS1{F0: 42, F1: 623}
                _ = v12
            case "":
                if (!true) {
                    if (1 <= -70) {
                        var v13 int32 = -2147483648
                        _ = v13
                    }
                    v11 = (float64(-2.25) - (v11 * (-2.25 + 3.14159)))
                }
            default:
                v14 := []int{}
                _ = v14
            }
        }
        v4 += (((v4 / 1.0) - ((1.5 - v4) + (v4 - (0.1 - (v4 - v4))))) - (v4 + 1.5))
        _ = v11
    }
    switch "a" {
    case "":
        _ = len(v3)
    default:
        switch int(v5) {
        case 32768:
            v4 -= v4
        }
    }
    var v15 string = ""
    switch 347 {
    case -2147483648:
        switch (v15 + (v15 + "\n\t")) {
        case "\n\t":
            if true {
                var v16 uint32 = 0
                var v17 uint16 = 65535
                _ = v16
                _ = v17
            } else {
                var v18 string = (v15 + v15)
                _ = v18
            }
        }
    case 750:
        fmt.Println(v5)
    default:
        _ = len(v2)
    }
    var v19 uint = 825
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
