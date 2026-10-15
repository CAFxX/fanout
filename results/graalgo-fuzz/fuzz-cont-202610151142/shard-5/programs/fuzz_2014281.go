package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint16 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 bool; F1 int64; F2 uint16 }
func genFn0() uint8 {
    switch "" {
    case "hello":
        var v7 uint32 = 68
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint = 44
        _ = fi8
        _ = v9
    }
    return 0
}

func main() {
    v1 := []string{"", "hello", ("a" + "foo bar"), ("hello" + "foo bar")}
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    v5 := []uint32{0, 2}
    var v6 int8 = 100
    v6 += 1
    var v10 float64 = (0.1 - 3.14159)
    var v11 string = "hello"
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: (true && false), F1: (fi12 - (fi12 + fi12)), F2: 77}
        if false {
            switch -251 {
            case 1:
                var v14 int16 = 1
                _ = v14
            }
            if true {
                switch "hello" {
                case "":
                    if false {
                        for fi15 := 0; fi15 < 20; fi15++ {
                            fmt.Println(v11, v11)
                            _ = fi15
                        }
                    } else {
                        var v16 uint16 = 669
                        _ = v16
                    }
                case "a":
                    switch (-(fi12)) {
                    case 32768:
                        if (!false) {
                            var v17 float64 = ((v10 - 0.0) + ((((0.0 + 0.0) + (0.1 + v10)) * v10) / 1.5))
                            _ = v17
                        }
                    case 9223372036854775807:
                        v18 := GenS2{F0: (("foo bar" + v11) + "foo bar"), F1: (128 < 1)}
                        _ = v18
                    case 128:
                        var v19 uint16 = 65535
                        _ = v19
                    default:
                        v10 += v10
                    }
                case "hello":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                default:
                    fmt.Println(v6)
                }
                _ = len(v3)
            } else {
                if false {
                    delete(v3, 922)
                    v10 += v10
                }
                switch (v11 + (v11 + ("" + "hello"))) {
                case "\n\t":
                    var v20 float64 = 3.14159
                    _ = v20
                }
                var v21 bool = (0 > 1)
                _ = v21
            }
            var v22 uint64 = 179
            _ = v22
        }
        v11 = v11
        _ = fi12
        _ = v13
    }
    switch v11 {
    case "hello":
        v23 := GenS1{F0: true, F1: 9223372036854775807, F2: 255}
        _ = v23
    }
    if false {
        v2[("a" + v11)] = true
    } else {
        var v24 int = 128
        _ = v24
    }
    var v25 uint8 = 42
    v26 := GenS3{F0: (42 < 1), F1: 9223372036854775807, F2: 739}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
