package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float64 = 3.14159
    var v3 uint = 1
    v4 := []uint32{2147483647, 4294967295, 963}
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 int32 = 2147483647
            var v9 string = "hello"
            v3 = v3
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    if false {
        var v10 uint32 = 2147483647
        fmt.Println(v3, v3)
        var v11 int8 = 127
        _ = v10
        _ = v11
    } else {
        switch ("a" + ("" + "\n\t")) {
        case "foo bar":
            var v12 string = "a"
            _ = v12
        case "\n\t":
            v2 = v2
        }
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 float32 = 1.5
        _ = fi13
        _ = v14
    }
    if (0 < 4294967295) {
        var v15 uint32 = 2
        _ = v15
    }
    fmt.Println(v1, v3, v3)
    v16 := GenS1{F0: 0, F1: -2147483648}
    v3 = 18446744073709551615
    fmt.Println(v2, v1)
    if true {
        fmt.Println(v2)
        v3 = v3
        v2 += (0.0 * (((1.5 * ((v2 + v2) / v2)) - 0.1) - v2))
    }
    if true {
        _ = len(v4)
        var v17 uint8 = 1
        _ = v17
    } else {
        fmt.Println(v3)
        fmt.Println(v3)
        if (-128 < 40) {
            v5[4294967295] = 0
            for fi18 := 0; fi18 < 10; fi18++ {
                _ = len(v5)
                v4 = append(v4, 0)
                _ = fi18
            }
            v19 := GenS1{F0: 32767, F1: 2147483647}
            _ = v19
        } else {
            var v20 int = 127
            fmt.Println(v2, v20)
            if false {
                for fi21 := 0; fi21 < 5; fi21++ {
                    var v22 int64 = (v1 >> (((v3 << (((v3 * v3)) & 63))) & 63))
                    _ = fi21
                    _ = v22
                }
                switch ((((("\n\t" + "\n\t") + (((("" + "") + (("hello" + "") + "hello")) + ("\n\t" + "hello")) + "a")) + "a") + "foo bar") + ("\n\t" + "\n\t")) {
                case "":
                    delete(v5, 65535)
                case "\n\t":
                    v5[2147483648] = (v20 | -1000)
                }
            } else {
                var v23 uint16 = 255
                var v24 bool = (0 <= 42)
                v1 -= (v1 >> (((v3 >> ((65535) & 63))) & 63))
                _ = v23
                _ = v24
            }
            _ = v20
        }
    }
    var v25 int16 = -32768
    var v26 uint = (v3 - (v3 / ((v3) | 1)))
    if ((1 != v1) || (v26 > v26)) {
        switch -1 {
        case 127:
            v3 -= (v26 >> ((v3) & 63))
        case -2147483648:
            var v27 float32 = 3.14159
            _ = v27
        case 330:
            v25 -= (v25 >> ((42) & 63))
        default:
            fmt.Println(v1, v1)
        }
    } else {
        v28 := GenS1{F0: v25, F1: -1}
        fmt.Println(v2)
        for fi29 := 0; fi29 < 15; fi29++ {
            v2 = 0.1
            var v30 uint32 = 2147483648
            for fi31 := 0; fi31 < 5; fi31++ {
                var v32 int = int(v3)
                switch "a" {
                case "a":
                    switch "foo bar" {
                    case "\n\t":
                        var v33 uint = 0
                        _ = v33
                    case "hello":
                        v30 += 2
                    case "":
                        switch 2147483647 {
                        case 2147483647:
                            if false {
                                _ = len(v5)
                                v26 += 0
                            } else {
                                var v34 uint64 = 4294967295
                                v26 -= uint(v34)
                                _ = v34
                            }
                        case 128:
                            var v35 int64 = (v1 % ((-618) | 1))
                            _ = v35
                        default:
                            v30 -= 2147483647
                        }
                    }
                case "":
                    fmt.Println(v1, v1, v26)
                case "foo bar":
                    v36 := GenS1{F0: (v25 * v25), F1: 270}
                    _ = v36
                }
                _ = fi31
                _ = v32
            }
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    if ("" <= ("\n\t" + "foo bar")) {
        var v37 uint32 = 2
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
