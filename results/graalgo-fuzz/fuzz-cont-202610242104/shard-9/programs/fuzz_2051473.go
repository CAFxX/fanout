package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 bool }
type GenS2 struct { F0 uint16 }
func genFn0() int {
    var v5 int32 = 42
    switch "a" {
    case "foo bar":
        fmt.Println(v5)
    case "hello":
        var v6 uint8 = 2
        _ = v6
    default:
        var v7 int64 = 0
        _ = v7
    }
    _ = v5
    return -168
}

func main() {
    var v1 uint64 = 358
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    var v8 bool = (244 != 42)
    for fi9 := 0; fi9 < 9; fi9++ {
        if v8 {
            v10 := []int{}
            _ = v10
        } else {
            var v11 int = -9223372036854775808
            for fi12 := 0; fi12 < 14; fi12++ {
                v11 = (fi12 + 141)
                v13 := GenS1{F0: 1, F1: v8, F2: (2 >= 1)}
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        _ = fi9
    }
    v8 = (891 == 2)
    if (((((1.5 + -2.25) / 1.5) * 0.0) - (((1e10 + 0.1) + 1.5) / 3.14159)) < 0.0) {
        for fi14 := 0; fi14 < 9; fi14++ {
            if (666 <= 0) {
                fmt.Println(v8)
                if false {
                    if (!(false && (!(-1 < 67)))) {
                        var v15 float64 = ((1e10 - (3.14159 + (-2.25 * 1.5))) - 0.1)
                        switch (fi14 * fi14) {
                        case -9223372036854775808:
                            var v16 string = "hello"
                            _ = v16
                        case 0:
                            fmt.Println(v1, v15)
                        default:
                            v8 = (!v8)
                        }
                        _ = v15
                    }
                    switch fi14 {
                    case -832:
                        switch (^(fi14)) {
                        case 127:
                            var v17 int8 = 112
                            _ = v17
                        case 2147483647:
                            v18 := []int{}
                            _ = v18
                        default:
                            var v19 uint32 = 1
                            _ = v19
                        }
                    case 42:
                        switch (fi14 * -331) {
                        case 128:
                            var v20 int64 = -1
                            _ = v20
                        case -487:
                            var v21 uint64 = 42
                            _ = v21
                        default:
                            var v22 int64 = 42
                            _ = v22
                        }
                    default:
                        var v23 int = (fi14 | fi14)
                        _ = v23
                    }
                    var v24 uint16 = 65535
                    _ = v24
                }
            }
            _ = fi14
        }
        var v25 int8 = -128
        _ = v25
    } else {
        var v26 float64 = 1.5
        var v27 float64 = (float64((0.1 / 1e10)) - v26)
        _ = v26
        _ = v27
    }
    v8 = false
    v1 += (v1 << ((389) & 63))
    if (!(v8 && v8)) {
        switch "\n\t" {
        case "":
            v28 := GenS2{F0: 1}
            _ = v28
        case "\n\t":
            v29 := GenS1{F0: 42, F1: (!true), F2: v8}
            _ = v29
        case "hello":
            var v30 float64 = 0.0
            _ = v30
        default:
            var v31 uint16 = 2
            _ = v31
        }
        switch ((((("\n\t" + "\n\t") + "a") + "\n\t") + "\n\t") + ("foo bar" + "")) {
        case "hello":
            v1 = (^(v1))
        }
        if (!v8) {
            switch "\n\t" {
            case "":
                var v32 uint16 = 106
                _ = v32
            }
        }
    }
    var v33 int = genFn0()
    var v34 int64 = 0
    var v35 uint8 = 113
    var v36 uint64 = v1
    _ = len(v3)
    switch ("hello" + "a") {
    case "hello":
        var v37 float32 = (float32((float64(3.14159) / -2.25)) / (((-2.25 - 0.0) - 1.5) * ((0.1 + float32(0.1)) + float32((0.1 * 3.14159)))))
        _ = v37
    case "a":
        v38 := []int{v33, v33, v33, (v33 & v33)}
        _ = v38
    }
    var v39 uint64 = (v1 / ((2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v8
}
