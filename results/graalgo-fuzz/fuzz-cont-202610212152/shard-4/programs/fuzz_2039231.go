package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int8; F2 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 0.1
    var v3 uint = 1
    v4 := make(map[int]int)
    var v5 int8 = 127
    var v6 float64 = v2
    switch -599 {
    case 127:
        if true {
            v3 += v3
        } else {
            fmt.Println(v3)
            switch "hello" {
            case "":
                if false {
                    for fi7 := 0; fi7 < 10; fi7++ {
                        var v8 uint8 = 0
                        _ = fi7
                        _ = v8
                    }
                }
            case "foo bar":
                if false {
                    var v9 int32 = 2147483647
                    v10 := GenS1{F0: -1}
                    _ = v10
                    _ = v9
                } else {
                    for fi11 := 0; fi11 < 15; fi11++ {
                        var v12 uint = v3
                        _ = fi11
                        _ = v12
                    }
                }
            case "hello":
                v13 := GenS1{F0: (v5 | -35)}
                _ = v13
            }
            v14 := []int{9223372036854775807}
            _ = v14
        }
    case 2147483647:
        var v15 int32 = 1
        _ = v15
    }
    if false {
        var v16 int16 = 671
        _ = v16
    } else {
        var v17 float32 = (1.5 / 1.5)
        _ = v17
    }
    switch 9223372036854775807 {
    case 2147483647:
        v18 := GenS1{F0: v5}
        _ = v18
    }
    v19 := GenS2{F0: 128}
    var v20 uint = (v3 - (v3 ^ v3))
    var v21 float32 = 0.1
    var v22 int64 = 0
    var v23 uint = (v3 + v20)
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 uint = 610
        for fi26 := 0; fi26 < 5; fi26++ {
            v27 := GenS2{F0: (fi24 * fi24)}
            v28 := v4[fi24]
            _ = fi26
            _ = v27
            _ = v28
        }
        v29 := GenS1{F0: v5}
        _ = fi24
        _ = v25
        _ = v29
    }
    switch "a" {
    case "a":
        var v30 int64 = -1
        _ = v30
    case "foo bar":
        v31 := GenS3{F0: v22, F1: v5, F2: 237}
        _ = v31
    case "":
        var v32 uint8 = 2
        _ = v32
    }
    var v33 int16 = 0
    var v34 int = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
