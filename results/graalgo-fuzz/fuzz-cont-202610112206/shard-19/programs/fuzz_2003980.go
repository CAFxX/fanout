package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 bool = false
    var v2 uint32 = 2
    var v3 float64 = 1.5
    v4 := make(map[uint32]int)
    var v5 int64 = -9223372036854775808
    var v6 int64 = (-(v5))
    v5 = v6
    if false {
        v7 := GenS1{F0: 1}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 16; fi8++ {
            v4[(^(v2))] = 9223372036854775807
            fmt.Println(v6)
            var v9 uint8 = 0
            _ = fi8
            _ = v9
        }
    }
    if v1 {
        v2 += 2147483648
        var v10 int64 = v6
        _ = v10
    } else {
        var v11 string = "foo bar"
        switch "a" {
        case "\n\t":
            var v12 uint32 = v2
            _ = v12
        case "a":
            v3 -= 0.0
        }
        v13 := GenS1{F0: 255}
        _ = v11
        _ = v13
    }
    v6 = v5
    for fi14 := 0; fi14 < 20; fi14++ {
        switch (fi14 - -148) {
        case 128:
            v6 = v6
        case -9223372036854775808:
            v15 := []int{fi14, (fi14 & (fi14 * (fi14 << ((2) & 63)))), 2147483647}
            _ = v15
        default:
            for fi16 := 0; fi16 < 2; fi16++ {
                switch (fi16 >> ((255) & 63)) {
                case 309:
                    if (v1 && v1) {
                        var v17 bool = (v1 || (!v1))
                        _ = v17
                    }
                case 32768:
                    if (!(!v1)) {
                        v2 -= 1
                    } else {
                        v18 := GenS1{F0: 42}
                        fmt.Println(v2, fi14)
                        var v19 int64 = 148
                        switch "foo bar" {
                        case "hello":
                            v4[v2] = 1
                        }
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi16
            }
        }
        var v20 int32 = 0
        _ = len(v4)
        _ = fi14
        _ = v20
    }
    v21 := GenS1{F0: 272}
    var v22 uint64 = 599
    var v23 uint16 = 32767
    var v24 int8 = 42
    v22 += v22
    switch 747 {
    case 117:
        v1 = v1
    case 1:
        var v25 uint16 = (v23 - (v23 - (v23 << ((946) & 63))))
        _ = v25
    }
    var v26 int32 = 42
    for fi27 := 0; fi27 < 5; fi27++ {
        fmt.Println(v24)
        v22 = (v22 - v22)
        _ = fi27
    }
    if (-2147483648 <= 128) {
        fmt.Println(v1)
    } else {
        var v28 int64 = (v5 * 586)
        _ = v28
    }
    v29 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
