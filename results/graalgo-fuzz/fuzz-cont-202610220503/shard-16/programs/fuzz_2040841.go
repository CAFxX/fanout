package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint16; F2 int32 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    var v3 uint8 = 2
    fmt.Println(v3)
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 int = fi4
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(fi6, fi6)
        _ = fi6
    }
    v3 = 42
    v7 := GenS1{F0: true}
    for fi8 := 0; fi8 < 8; fi8++ {
        switch fi8 {
        case -2147483648:
            var v9 string = "hello"
            _ = v9
        case 42:
            for fi10 := 0; fi10 < 2; fi10++ {
                switch "hello" {
                case "foo bar":
                    v11 := []int{}
                    _ = v11
                case "hello":
                    v12 := v2[(-(fi8))]
                    _ = v12
                default:
                    var v13 float64 = (1.5 + (1e10 * 0.0))
                    _ = v13
                }
                v3 += v3
                switch ("foo bar" + ("" + "hello")) {
                case "foo bar":
                    var v14 int64 = 619
                    _ = v14
                case "a":
                    v15 := []int{fi8, -1, fi10, fi10}
                    _ = v15
                }
                _ = fi10
            }
        }
        _ = fi8
    }
    var v16 string = "foo bar"
    switch -852 {
    case 32768:
        if (-238 > 42) {
            var v17 uint = 824
            v18 := GenS2{F0: 4294967295, F1: 0, F2: 1}
            _ = v17
            _ = v18
        } else {
            var v19 int8 = -18
            _ = v19
        }
    case 89:
        var v20 float64 = 1.5
        _ = v20
    }
    v3 = 0
    var v21 int64 = 9223372036854775807
    var v22 int32 = -1
    var v23 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v7
}
