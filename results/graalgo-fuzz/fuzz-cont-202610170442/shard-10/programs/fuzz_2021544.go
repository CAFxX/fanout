package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint32; F2 bool }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    v3 := []string{("foo bar" + "a"), ("" + "foo bar"), ("\n\t" + (("foo bar" + ("foo bar" + "foo bar")) + "a")), "foo bar"}
    var v4 uint64 = 2
    var v5 int = 42
    v6 := GenS1{F0: 65535, F1: 77, F2: v5}
    switch (v5 & (v5 * -9223372036854775808)) {
    case 2147483647:
        var v7 int8 = -99
        _ = v7
    default:
        delete(v1, "hello")
    }
    switch ("foo bar" + "a") {
    case "hello":
        if false {
            v5 = v5
            v3 = append(v3, "\n\t")
        } else {
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := GenS1{F0: 0, F1: 2147483647, F2: 9223372036854775807}
                _ = fi8
                _ = v9
            }
        }
    case "":
        var v10 string = (("" + "foo bar") + "")
        _ = v10
    default:
        for fi11 := 0; fi11 < 15; fi11++ {
            fmt.Println(v4)
            _ = fi11
        }
    }
    if false {
        _ = len(v3)
    } else {
        if false {
            if false {
                if true {
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v4 = v4
                        _ = fi12
                    }
                }
                fmt.Println(v4, v4)
            } else {
                for fi13 := 0; fi13 < 6; fi13++ {
                    v5 -= v5
                    var v14 uint = 0
                    for fi15 := 0; fi15 < 1; fi15++ {
                        v16 := GenS2{F0: 1}
                        _ = fi15
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                }
                for fi17 := 0; fi17 < 5; fi17++ {
                    switch "hello" {
                    case "":
                        var v18 int32 = 1
                        _ = v18
                    case "a":
                        v19 := GenS1{F0: 65535, F1: 0, F2: 2147483647}
                        _ = v19
                    case "\n\t":
                        v20 := GenS3{F0: ((float32(float64(1.5)) - (0.1 + float32(float64((1.5 * ((float32(0.1) - (-2.25 * (0.0 * 0.0))) + 0.1)))))) < 1e10), F1: 483, F2: (true || true)}
                        _ = v20
                    }
                    _ = fi17
                }
            }
            var v21 string = "foo bar"
            var v22 int = (v5 - v5)
            _ = v21
            _ = v22
        }
    }
    var v23 uint32 = 0
    if (true && (1 > 0)) {
        v24 := GenS3{F0: true, F1: v23, F2: (("" + "foo bar") < "foo bar")}
        fmt.Println(v4, v5, v23)
        _ = v24
    }
    var v25 bool = false
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
