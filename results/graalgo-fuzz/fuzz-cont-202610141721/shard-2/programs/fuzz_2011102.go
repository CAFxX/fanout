package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int16 }
func genFn0(p0 int, p1 int, p2 int8) uint {
    for fi5 := 0; fi5 < 12; fi5++ {
        fmt.Println(fi5, p1)
        p1 = (p0 * (p1 >> ((643) & 63)))
        _ = fi5
    }
    switch (p0 - p0) {
    case 42:
        if (708 >= 110) {
            p2 = (p2 / ((100) | 1))
        } else {
            switch ("" + ("hello" + ("hello" + "hello"))) {
            case "\n\t":
                var v6 uint = 746
                _ = v6
            case "foo bar":
                p1 = p1
            case "a":
                for fi7 := 0; fi7 < 1; fi7++ {
                    var v8 uint8 = 255
                    v9 := []int{p0, (p0 & 2147483647)}
                    fmt.Println(p0)
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
            if false {
                var v10 uint32 = 0
                _ = v10
            } else {
                var v11 int = (p1 + 2147483647)
                fmt.Println(p1)
                _ = v11
            }
        }
    case 0:
        p1 = 516
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS1{F0: -428, F1: (p1 << ((255) & 63)), F2: (^(fi12))}
        p0 -= (^(p0))
        _ = fi12
        _ = v13
    }
    p0 = p0
    if (4294967295 >= 2147483648) {
        if (true && true) {
            if (0 <= 221) {
                var v14 uint = 255
                _ = v14
            } else {
                v15 := []int{(p1 | (p1 + 32767)), p0, p0, p0}
                switch "hello" {
                case "a":
                    v15 = append(v15, (-(p0)))
                case "hello":
                    for fi16 := 0; fi16 < 17; fi16++ {
                        var v17 int8 = 25
                        var v18 float64 = 1e10
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                default:
                    var v19 string = ""
                    _ = v19
                }
                _ = v15
            }
        }
        var v20 int8 = (p2 ^ p2)
        switch p1 {
        case 2147483647:
            p1 = (p0 ^ 127)
        case 42:
            var v21 int32 = -1
            _ = v21
        case -64:
            p1 = p1
        }
        _ = v20
    } else {
        var v22 uint64 = 2
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 112
    var v3 float32 = ((0.0 * 1e10) - (1.5 + 0.0))
    var v4 uint = 2
    v2 = 1
    v4 = (v4 / ((v4) | 1))
    v23 := []int{127, -2147483648, 2147483647}
    switch "" {
    case "":
        v4 = (v4 + (v4 << ((v4) & 63)))
    case "a":
        var v24 int = 42
        _ = v24
    }
    var v25 int16 = 814
    if ((v4 ^ 0) < (v4 / ((428) | 1))) {
        if true {
            v26 := GenS1{F0: -287, F1: 128, F2: 268}
            v23 = append(v23, 4294967295)
            v3 = v3
            _ = v26
        }
        v27 := GenS2{F0: v2, F1: 2}
        if true {
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 uint = 65535
                v25 += (v25 / ((-511) | 1))
                v30 := GenS1{F0: -1, F1: 600, F2: (fi28 + fi28)}
                _ = fi28
                _ = v29
                _ = v30
            }
            v2 = 156
            v31 := GenS3{F0: 42, F1: v2, F2: (v25 * (-(v25)))}
            _ = v31
        } else {
            v4 = genFn0(9223372036854775807, -9223372036854775808, 1)
        }
        _ = v27
    }
    v4 -= v4
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
