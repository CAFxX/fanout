package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int64 }

func main() {
    var v1 bool = ((!true) || false)
    var v2 int32 = 2147483647
    v3 := []int{127}
    var v4 uint8 = 10
    var v5 int64 = 42
    v6 := make(map[uint32]bool)
    v7 := v6[42]
    if false {
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS1{F0: (fi8 | (fi8 * fi8))}
                var v11 int16 = 42
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
        if len(v3) > 0 {
            v3[0] = -2147483648
        }
    } else {
        if len(v3) > 0 {
            v3[0] = 0
        }
        var v12 int64 = (-(v5))
        for fi13 := 0; fi13 < 7; fi13++ {
            switch "a" {
            case "hello":
                switch "foo bar" {
                case "hello":
                    v14 := GenS2{F0: 156, F1: 0}
                    _ = v14
                case "\n\t":
                    v7 = false
                default:
                    v2 -= v2
                }
            case "a":
                if (42 < 127) {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v16 := GenS3{F0: 0, F1: (("\n\t" + "") + (("hello" + "foo bar") + "\n\t"))}
                        var v17 uint16 = 2
                        var v18 string = "\n\t"
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                } else {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        v12 = v5
                        var v20 uint64 = 0
                        _ = len(v3)
                        _ = fi19
                        _ = v20
                    }
                    if v1 {
                        var v21 int = int(v5)
                        _ = v21
                    }
                }
            case "foo bar":
                v22 := GenS2{F0: fi13, F1: 2147483647}
                _ = v22
            }
            v23 := GenS2{F0: -2147483648, F1: 0}
            v12 += v5
            _ = fi13
            _ = v23
        }
        _ = v12
    }
    var v24 float64 = (((1e10 / 1.0) + (0.1 - (1.5 / 3.14159))) * (1.5 - ((3.14159 - 3.14159) + 3.14159)))
    for fi25 := 0; fi25 < 4; fi25++ {
        var v26 string = "\n\t"
        _ = fi25
        _ = v26
    }
    v27 := GenS2{F0: -2147483648, F1: 260}
    v28 := GenS3{F0: -261, F1: "foo bar"}
    for fi29 := 0; fi29 < 18; fi29++ {
        if ((4294967295 > uint64(v4)) && (fi29 > (-(fi29)))) {
            v30 := GenS2{F0: (fi29 | fi29), F1: 4294967295}
            _ = v30
        }
        _ = fi29
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        var v32 uint8 = (v4 + 127)
        _ = fi31
        _ = v32
    }
    v24 = v24
    var v33 uint16 = 2
    v5 = -1
    switch ("" + "") {
    case "":
        v3 = append(v3, -816)
    case "foo bar":
        if ("a" < "\n\t") {
            var v34 float32 = 1.5
            _ = v34
        }
    case "hello":
        var v35 uint = 4294967295
        _ = v35
    }
    switch "a" {
    case "hello":
        delete(v6, 1)
    case "\n\t":
        v36 := GenS4{F0: v5}
        _ = v36
    case "foo bar":
        for fi37 := 0; fi37 < 3; fi37++ {
            v33 += v33
            _ = fi37
        }
    default:
        v38 := GenS2{F0: -827, F1: 574}
        _ = v38
    }
    v39 := v6[2147483647]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
