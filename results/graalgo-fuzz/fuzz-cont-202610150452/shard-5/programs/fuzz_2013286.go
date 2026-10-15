package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int64 }
func genFn0(p0 uint8, p1 int, p2 int) int {
    fmt.Println(p0, p2, p1)
    var v4 int64 = int64(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p1 | (^(p2)))
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[uint32]uint32)
    var v3 float64 = (float64((-2.25 * (0.0 / 1e10))) * 1e10)
    v3 = (v3 * (0.0 / v3))
    fmt.Println(v3, v1)
    v1 -= 127
    switch ("hello" + "a") {
    case "foo bar":
        var v5 string = ""
        _ = v5
    case "a":
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            var v6 uint16 = 115
            _ = v6
        }
    case "\n\t":
        var v7 int8 = 116
        _ = v7
    }
    var v8 uint16 = 0
    if false {
        switch "foo bar" {
        case "hello":
            fmt.Println(v3, v1, v3)
        case "":
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 string = ""
                switch ((("" + (v10 + (v10 + "hello"))) + "\n\t") + (v10 + v10)) {
                case "hello":
                    for fi11 := 0; fi11 < 5; fi11++ {
                        v8 += v8
                        if true {
                            delete(v2, 2)
                        } else {
                            v10 = "hello"
                            var v12 int64 = -45
                            _ = v12
                        }
                        var v13 uint64 = 314
                        _ = fi11
                        _ = v13
                    }
                case "foo bar":
                    fmt.Println(v8, fi9)
                case "":
                    v10 = (v10 + "foo bar")
                default:
                    v1 -= 83
                }
                v10 = "\n\t"
                _ = fi9
                _ = v10
            }
        case "a":
            for fi14 := 0; fi14 < 11; fi14++ {
                v8 += (v8 * 65535)
                _ = fi14
            }
        default:
            fmt.Println(v1, v1, v3)
        }
        v8 = 773
    }
    var v15 int = -9223372036854775808
    v16 := []int{int(v1), genFn0(255, (v15 + (v15 << ((828) & 63))), genFn0(v1, 128, (v15 | (v15 & v15)))), (v15 - 2147483647)}
    fmt.Println(v8, v8, v15)
    fmt.Println(v8, v3, v1)
    for fi17 := 0; fi17 < 5; fi17++ {
        switch (v15 * 42) {
        case -9223372036854775808:
            var v18 int16 = -32768
            _ = v18
        case 673:
            v19 := GenS1{F0: 788, F1: (v15 | genFn0((v1 / ((0) | 1)), genFn0(v1, 128, 42), fi17)), F2: -864}
            _ = v19
        }
        fmt.Println(fi17)
        fmt.Println(v3, v15)
        _ = fi17
    }
    fmt.Println(v15, v8)
    switch (-(v15)) {
    case 421:
        v20 := GenS1{F0: 4294967295, F1: v15, F2: -9223372036854775808}
        _ = v20
    }
    v21 := GenS1{F0: 1, F1: v15, F2: 47}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v8
}
