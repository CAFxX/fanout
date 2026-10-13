package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 uint64 = 2
    v2 := []string{("" + "hello")}
    var v3 float64 = 0.1
    v4 := []uint32{1, 4294967295}
    v5 := []string{"hello", "foo bar", ("foo bar" + (((("" + "foo bar") + "a") + "foo bar") + ("" + ("" + "a")))), "foo bar"}
    v6 := make(map[string]string)
    fmt.Println(v1, v3)
    var v7 uint64 = v1
    v8 := GenS1{F0: 73, F1: 2147483647}
    if (0 == v7) {
        v9 := GenS2{F0: 255, F1: "\n\t"}
        var v10 float64 = v3
        _ = v10
        _ = v9
    }
    var v11 uint64 = 4294967295
    var v12 float64 = v3
    var v13 uint8 = 227
    for fi14 := 0; fi14 < 19; fi14++ {
        v3 = v12
        if false {
            if true {
                switch "hello" {
                case "foo bar":
                    v12 -= 0.0
                case "":
                    v3 -= (v3 + v12)
                case "hello":
                    v12 = v12
                default:
                    var v15 int = (fi14 ^ fi14)
                    _ = v15
                }
                var v16 int32 = 42
                var v17 uint64 = (v1 ^ 432)
                _ = v16
                _ = v17
            }
            v6["hello"] = "a"
        }
        _ = fi14
    }
    var v18 uint16 = 2
    var v19 uint64 = v11
    v13 -= v13
    var v20 uint = 449
    if true {
        switch -2147483648 {
        case -2147483648:
            var v21 int32 = 970
            _ = v21
        case -1:
            v2 = append(v2, "a")
        case 32768:
            if true {
                for fi22 := 0; fi22 < 12; fi22++ {
                    var v23 uint32 = uint32(v20)
                    var v24 string = "hello"
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                for fi25 := 0; fi25 < 20; fi25++ {
                    switch "\n\t" {
                    case "a":
                        var v26 float64 = v3
                        _ = v26
                    case "hello":
                        v27 := GenS3{F0: 676, F1: 255}
                        _ = v27
                    case "\n\t":
                        if (2147483648 != 65535) {
                            var v28 uint32 = 0
                            _ = v28
                        }
                    }
                    _ = fi25
                }
                switch "\n\t" {
                case "hello":
                    fmt.Println(v1)
                case "foo bar":
                    var v29 int64 = 9223372036854775807
                    _ = v29
                }
            } else {
                var v30 int64 = 671
                _ = v30
            }
        }
        fmt.Println(v11, v18)
    } else {
        v31 := GenS2{F0: v13, F1: "a"}
        _ = v31
    }
    v4 = append(v4, 0)
    var v32 int = -372
    for fi33 := 0; fi33 < 5; fi33++ {
        var v34 uint64 = v7
        if true {
            var v35 int8 = -128
            for fi36 := 0; fi36 < 2; fi36++ {
                v37 := GenS4{F0: 32767}
                if false {
                    var v38 uint = 65535
                    var v39 bool = (!false)
                    _ = v38
                    _ = v39
                }
                _ = len(v6)
                _ = fi36
                _ = v37
            }
            _ = v35
        }
        v3 = float64(0.1)
        _ = fi33
        _ = v34
    }
    for fi40 := 0; fi40 < 9; fi40++ {
        v3 += v3
        v12 = v3
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
