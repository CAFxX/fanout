package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 string }

func main() {
    var v1 string = "a"
    var v2 uint = 18446744073709551615
    var v3 float64 = 1e10
    fmt.Println(v3)
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 float32 = 0.1
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := []int{-2147483648, (fi6 << ((2) & 63)), fi4}
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    var v8 uint32 = 4294967295
    var v9 uint64 = 599
    v3 = (v3 * ((v3 + v3) + v3))
    if true {
        for fi10 := 0; fi10 < 4; fi10++ {
            for fi11 := 0; fi11 < 8; fi11++ {
                if false {
                    switch (-(fi10)) {
                    case 32767:
                        switch v1 {
                        case "":
                            var v12 int8 = 127
                            _ = v12
                        default:
                            v2 = (v2 % (((v2 >> ((v2) & 63))) | 1))
                        }
                    case -685:
                        v13 := GenS1{F0: -992, F1: -32768, F2: ("a" + v1)}
                        _ = v13
                    }
                } else {
                    fmt.Println(v9)
                    fmt.Println(fi11)
                    switch int(v2) {
                    case -2147483648:
                        if (3.14159 >= 3.14159) {
                            v14 := GenS1{F0: -205, F1: 42, F2: ((v1 + v1) + v1)}
                            var v15 int8 = 42
                            _ = v14
                            _ = v15
                        }
                    case 2147483647:
                        var v16 uint32 = (v8 * v8)
                        _ = v16
                    }
                    switch (fi11 * 128) {
                    case -2147483648:
                        var v17 uint32 = v8
                        _ = v17
                    case 2147483647:
                        if (!false) {
                            if (42 <= uint64(v2)) {
                                fmt.Println(v8)
                            }
                            v18 := []int{32767, fi10, (^(fi11))}
                            var v19 float32 = (-2.25 + 0.0)
                            _ = v18
                            _ = v19
                        }
                    }
                }
                fmt.Println(v1, v1, fi10)
                var v20 int64 = int64(fi10)
                _ = fi11
                _ = v20
            }
            var v21 float32 = 0.1
            switch (fi10 / ((fi10) | 1)) {
            case -766:
                var v22 uint32 = (v8 + v8)
                _ = v22
            case 1:
                v23 := []int{}
                _ = v23
            }
            _ = fi10
            _ = v21
        }
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        if (65535 >= v2) {
            if (float32(1.5) != ((((((1.5 * 1.5) / 1.0) - 1.5) + float32(v3)) / 1e10) * -2.25)) {
                switch (^(fi24)) {
                case 2147483647:
                    v1 = (v1 + (v1 + v1))
                }
            }
            var v25 uint16 = 211
            _ = v25
        } else {
            v26 := GenS1{F0: -594, F1: 1000, F2: "a"}
            var v27 uint32 = 0
            switch "hello" {
            case "foo bar":
                v9 -= v9
            case "hello":
                fmt.Println(v2, v8)
            }
            _ = v26
            _ = v27
        }
        _ = fi24
    }
    switch v1 {
    case "a":
        for fi28 := 0; fi28 < 9; fi28++ {
            fmt.Println(v3, v3)
            v2 += (v2 + (v2 * (v2 - 890)))
            _ = fi28
        }
    case "hello":
        var v29 int32 = 0
        _ = v29
    }
    switch "foo bar" {
    case "hello":
        var v30 int16 = -32768
        _ = v30
    default:
        switch 32768 {
        case 4294967295:
            v2 -= 468
        case 42:
            v31 := GenS1{F0: 9223372036854775807, F1: 1000, F2: v1}
            _ = v31
        default:
            var v32 int64 = 0
            _ = v32
        }
    }
    v33 := GenS1{F0: -9223372036854775808, F1: -32768, F2: ("\n\t" + "\n\t")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v8
    _ = v9
}
