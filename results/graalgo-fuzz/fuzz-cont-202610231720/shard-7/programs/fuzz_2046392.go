package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 string }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 string; F1 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 4294967295
    v3 := make(map[int]int)
    var v4 int32 = -1
    v2 = v2
    v4 = -1
    var v5 float32 = (1.5 / 1.5)
    v6 := GenS1{F0: ("" + "\n\t"), F1: 42, F2: "hello"}
    for fi7 := 0; fi7 < 4; fi7++ {
        switch "hello" {
        case "foo bar":
            fmt.Println(v5, fi7, v4)
        case "hello":
            if true {
                if (!(1 != 1)) {
                    v8 := []int{(fi7 ^ fi7), (fi7 >> ((0) & 63)), fi7, 849}
                    _ = v8
                }
            } else {
                var v9 uint32 = 2147483647
                var v10 uint32 = 4294967295
                var v11 uint16 = 456
                _ = v10
                _ = v11
                _ = v9
            }
        case "":
            v12 := GenS1{F0: "foo bar", F1: (v4 >> ((250) & 63)), F2: "a"}
            _ = v12
        }
        fmt.Println(v5)
        _ = fi7
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int8 = 42
        var v15 int8 = -128
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint64 = (^(v2))
    if false {
        var v17 uint = 42
        var v18 string = ("" + "\n\t")
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 17; fi19++ {
            switch ("hello" + "") {
            case "\n\t":
                v20 := GenS1{F0: ("foo bar" + "\n\t"), F1: v4, F2: "\n\t"}
                _ = v20
            case "a":
                for fi21 := 0; fi21 < 14; fi21++ {
                    switch 1 {
                    case 417:
                        v5 = (float32(0.0) + (v5 * v5))
                    case 1:
                        v22 := GenS2{F0: (!true), F1: int64(v2)}
                        _ = v22
                    case 42:
                        v23 := GenS1{F0: "hello", F1: v4, F2: ("foo bar" + "foo bar")}
                        _ = v23
                    default:
                        v24 := GenS1{F0: ("hello" + ("\n\t" + ("foo bar" + "hello"))), F1: (^(v4)), F2: "hello"}
                        _ = v24
                    }
                    var v25 float64 = (float64(v5) + -2.25)
                    _ = fi21
                    _ = v25
                }
            case "hello":
                if false {
                    fmt.Println(fi19, v2)
                    v3[fi19] = 0
                    var v26 bool = ("\n\t" != ("\n\t" + ("hello" + "foo bar")))
                    _ = v26
                }
            default:
                v27 := []int{}
                _ = v27
            }
            fmt.Println(fi19)
            _ = fi19
        }
        fmt.Println(v16)
        for fi28 := 0; fi28 < 15; fi28++ {
            v29 := GenS3{F0: ("a" + "hello"), F1: 4294967295}
            switch "" {
            case "a":
                v30 := GenS1{F0: "foo bar", F1: (v4 % ((v4) | 1)), F2: ""}
                _ = v30
            }
            var v31 uint64 = v16
            _ = fi28
            _ = v29
            _ = v31
        }
    }
    if true {
        v16 -= (v16 + (v16 % ((v2) | 1)))
    }
    var v32 bool = ((-2.25 + (3.14159 * (float64(1.5) * 0.0))) != 0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
