package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    v4 := []int{}
    v5 := []uint32{}
    v6 := []string{"\n\t", "a", ("\n\t" + (("a" + ("a" + "a")) + "a")), ("a" + "foo bar")}
    v7 := GenS1{F0: 42, F1: 1, F2: 42}
    switch -2147483648 {
    case 454:
        fmt.Println(v1)
    case -333:
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 float32 = 0.1
            _ = fi8
            _ = v9
        }
    case -117:
        if false {
            switch "\n\t" {
            case "\n\t":
                var v10 float32 = (((1.5 - 3.14159) + 3.14159) + (((3.14159 - ((3.14159 * (1.5 + ((1e10 * -2.25) / 1.0))) * float32(0.0))) * 3.14159) + (-2.25 - 1.5)))
                _ = v10
            case "foo bar":
                fmt.Println(v1)
            case "hello":
                var v11 bool = (32768 >= 1)
                _ = v11
            }
            if false {
                for fi12 := 0; fi12 < 6; fi12++ {
                    for fi13 := 0; fi13 < 18; fi13++ {
                        var v14 float64 = 0.0
                        v1 = 2
                        _ = fi13
                        _ = v14
                    }
                    _ = fi12
                }
                var v15 int32 = -354
                fmt.Println(v1)
                _ = v15
            } else {
                var v16 float32 = 1e10
                _ = v16
            }
        } else {
            v17 := v2[(("a" + ("a" + ("" + "foo bar"))) + "\n\t")]
            _ = v17
        }
    }
    v18 := GenS2{F0: 2}
    v1 += (^(v1))
    if true {
        var v19 int64 = -9223372036854775808
        v20 := GenS1{F0: 462, F1: 4294967295, F2: 0}
        _ = v19
        _ = v20
    }
    var v21 int64 = 498
    var v22 uint16 = 1
    v1 = 480
    _ = len(v6)
    var v23 bool = false
    v24 := GenS3{F0: 42}
    var v25 uint64 = 1
    var v26 uint32 = (v1 / ((2147483648) | 1))
    var v27 float64 = 1e10
    var v28 float64 = v27
    if v23 {
        var v29 float64 = v28
        _ = v29
    }
    v30 := GenS4{F0: v25}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
