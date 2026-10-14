package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 string; F1 string; F2 uint8 }

func main() {
    var v1 int8 = -1
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    v2[(("foo bar" + "a") + ("hello" + "a"))] = 200
    v4 := GenS1{F0: -408, F1: 747, F2: 181}
    fmt.Println(v1)
    if false {
        _ = len(v3)
    } else {
        if (294 <= 566) {
            var v5 bool = ((v1 / (((v1 / (((v1 - (v1 | -128))) | 1))) | 1)) == v1)
            _ = v5
        } else {
            var v6 int16 = 746
            _ = v6
        }
        if false {
            fmt.Println(v1)
            var v7 bool = (960 <= 0)
            fmt.Println(v7)
            _ = v7
        }
        v8 := GenS1{F0: 0, F1: 475, F2: 1}
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint = 0
            v1 = v1
            if (fi9 == 32768) {
                var v12 uint16 = 2
                v1 = v1
                _ = v12
            } else {
                if (-1 != 705) {
                    v13 := GenS1{F0: 1, F1: 4294967295, F2: 127}
                    _ = v13
                }
                switch "a" {
                case "foo bar":
                    if true {
                        var v14 int = fi9
                        var v15 uint8 = 186
                        var v16 uint16 = 465
                        var v17 int8 = v1
                        var v18 int8 = (^(v1))
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        v1 -= (v1 | v1)
                    }
                case "":
                    v2[(("foo bar" + "") + "\n\t")] = 452
                case "hello":
                    var v19 uint64 = 1
                    _ = v19
                default:
                    fmt.Println(v11, fi9, v11)
                }
            }
            _ = fi10
            _ = v11
        }
        fmt.Println(v1)
        _ = fi9
    }
    switch ((("\n\t" + "foo bar") + ("\n\t" + "")) + "a") {
    case "foo bar":
        v20 := GenS1{F0: 9223372036854775807, F1: 830, F2: 10}
        _ = v20
    case "\n\t":
        var v21 int32 = 38
        _ = v21
    case "hello":
        v22 := []int{-990}
        _ = v22
    }
    if false {
        fmt.Println(v1)
        v1 = (v1 | (v1 | v1))
    }
    switch 149 {
    case 32768:
        v23 := GenS1{F0: 32768, F1: 2147483647, F2: 2}
        _ = v23
    case 0:
        for fi24 := 0; fi24 < 9; fi24++ {
            fmt.Println(v1)
            _ = fi24
        }
    case -305:
        v25 := GenS1{F0: 1, F1: 0, F2: 42}
        _ = v25
    default:
        if (!(65535 <= 329)) {
            switch ((("" + "") + ("" + "")) + "\n\t") {
            case "\n\t":
                fmt.Println(v1)
            }
            v26 := GenS2{F0: 9223372036854775807, F1: 2}
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        v28 := v2["\n\t"]
        _ = fi27
        _ = v28
    }
    v29 := GenS3{F0: "foo bar", F1: "\n\t", F2: 60}
    var v30 float64 = 0.1
    v30 = 0.0
    v31 := GenS3{F0: ("" + ("" + ("foo bar" + "foo bar"))), F1: "hello", F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
