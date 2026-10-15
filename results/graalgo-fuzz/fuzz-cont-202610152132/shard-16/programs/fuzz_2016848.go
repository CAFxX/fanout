package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }

func main() {
    v1 := []int{}
    v2 := []uint32{2147483647}
    var v3 int32 = 0
    var v4 int8 = 100
    v5 := make(map[uint32]uint32)
    var v6 uint = 109
    switch ("" + "\n\t") {
    case "":
        switch "\n\t" {
        case "a":
            switch "foo bar" {
            case "a":
                for fi7 := 0; fi7 < 7; fi7++ {
                    v8 := GenS1{F0: v4}
                    for fi9 := 0; fi9 < 8; fi9++ {
                        _ = len(v1)
                        switch "" {
                        case "a":
                            v6 += v6
                        default:
                            var v10 bool = true
                            _ = v10
                        }
                        _ = fi9
                    }
                    fmt.Println(v3)
                    _ = fi7
                    _ = v8
                }
            }
        case "":
            var v11 float64 = (0.1 + 1e10)
            _ = v11
        }
    case "\n\t":
        v5[65535] = 34
    case "foo bar":
        v12 := GenS2{F0: "a"}
        _ = v12
    }
    var v13 string = (("hello" + "foo bar") + "foo bar")
    v4 -= v4
    if false {
        if (float64((float32(0.1) * -2.25)) <= 0.1) {
            var v14 uint = 18446744073709551615
            _ = v14
        }
        v13 = ""
        var v15 int64 = -360
        _ = v15
    } else {
        if ((0 != 0) && true) {
            _ = len(v5)
            var v16 int8 = v4
            fmt.Println(v13)
            _ = v16
        }
    }
    fmt.Println(v13, v3)
    if true {
        var v17 int8 = 100
        _ = len(v5)
        _ = v17
    } else {
        v4 = v4
        fmt.Println(v13)
    }
    fmt.Println(v4)
    v13 = v13
    switch 42 {
    case 685:
        if (((float64(1.5) >= -2.25) || true) || (0 != 2)) {
            if ((false && (v13 >= (v13 + v13))) || (1e10 <= -2.25)) {
                v4 += (^(v4))
                v4 += v4
                fmt.Println(v3, v4)
            }
        } else {
            _ = len(v5)
            switch v13 {
            case "hello":
                var v18 int16 = -607
                _ = v18
            case "":
                var v19 int32 = v3
                _ = v19
            case "a":
                for fi20 := 0; fi20 < 17; fi20++ {
                    for fi21 := 0; fi21 < 10; fi21++ {
                        v4 -= 0
                        var v22 int64 = -9223372036854775808
                        _ = fi21
                        _ = v22
                    }
                    if (!true) {
                        var v23 string = ""
                        v6 = (^(v6))
                        _ = v23
                    } else {
                        if true {
                            fmt.Println(fi20, v3)
                            switch -9223372036854775808 {
                            case 1:
                                var v24 int32 = v3
                                _ = v24
                            case 9223372036854775807:
                                v5[42] = 4294967295
                            }
                        }
                    }
                    _ = fi20
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
