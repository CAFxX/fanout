package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int64 = 42
    v2 := make(map[int]int)
    var v3 float32 = float32(0.0)
    v4 := []int{0}
    v5 := make(map[int]int)
    v1 -= v1
    var v6 int = 2147483647
    _ = len(v4)
    switch 128 {
    case 42:
        v3 += ((v3 + v3) * ((v3 + ((v3 + ((-2.25 / (0.0 - v3)) - v3)) / v3)) + v3))
    case -859:
        var v7 int32 = 842
        _ = v7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint64 = 42
        var v10 uint = 18446744073709551615
        switch ("" + "\n\t") {
        case "a":
            if (-2147483648 != (fi8 << (((v10 ^ 831)) & 63))) {
                switch ("hello" + (("a" + "hello") + "foo bar")) {
                case "hello":
                    var v11 string = ("hello" + ("hello" + "\n\t"))
                    _ = v11
                case "a":
                    v12 := GenS1{F0: 255}
                    _ = v12
                case "":
                    var v13 float64 = 3.14159
                    _ = v13
                }
                var v14 float64 = float64(v3)
                for fi15 := 0; fi15 < 8; fi15++ {
                    _ = len(v4)
                    _ = fi15
                }
                _ = v14
            } else {
                switch 546 {
                case 309:
                    v1 -= (^(v1))
                case 9223372036854775807:
                    if len(v4) > 0 {
                        v4[0] = -920
                    }
                case 4294967295:
                    switch "foo bar" {
                    case "a":
                        v4 = append(v4, (v6 / ((v6) | 1)))
                    case "hello":
                        v6 -= fi8
                    case "":
                        var v16 uint = (^(v10))
                        _ = v16
                    default:
                        var v17 int8 = 0
                        _ = v17
                    }
                }
            }
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    delete(v5, (v6 % ((-545) | 1)))
    v3 = float32(((0.1 / 1.0) * 0.1))
    for fi18 := 0; fi18 < 5; fi18++ {
        if false {
            v19 := GenS1{F0: 255}
            _ = v19
        } else {
            v1 -= v1
        }
        _ = fi18
    }
    var v20 string = ""
    fmt.Println(v3, v3)
    if len(v4) > 0 {
        v4[0] = (v6 | -9223372036854775808)
    }
    var v21 uint64 = 0
    switch v6 {
    case 128:
        switch (v20 + (v20 + v20)) {
        case "a":
            v20 = v20
        case "foo bar":
            v22 := GenS1{F0: 2}
            _ = v22
        case "\n\t":
            var v23 int32 = 1
            _ = v23
        }
    case -69:
        v24 := GenS1{F0: 127}
        _ = v24
    case -840:
        v25 := v5[v6]
        _ = v25
    default:
        if true {
            for fi26 := 0; fi26 < 2; fi26++ {
                var v27 float64 = -2.25
                v28 := GenS1{F0: 255}
                _ = fi26
                _ = v27
                _ = v28
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
