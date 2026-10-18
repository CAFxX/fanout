package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int; F2 uint }

func main() {
    var v1 uint32 = 65535
    var v2 uint16 = 2
    var v3 uint64 = 61
    var v4 uint16 = uint16(v3)
    var v5 int8 = 127
    var v6 string = "foo bar"
    switch v6 {
    case "hello":
        if false {
            v7 := []int{-20, 9223372036854775807, 276}
            v7 = append(v7, 42)
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 float32 = 0.1
                fmt.Println(v3, v5)
                _ = fi8
                _ = v9
            }
            _ = v7
        } else {
            fmt.Println(v2, v2)
            v10 := GenS1{F0: -1, F1: 2147483648, F2: -1}
            if (v3 > v3) {
                v11 := GenS2{F0: 288}
                _ = v11
            } else {
                for fi12 := 0; fi12 < 10; fi12++ {
                    var v13 int = 760
                    v5 = (v5 * (v5 & v5))
                    _ = fi12
                    _ = v13
                }
            }
            _ = v10
        }
    default:
        fmt.Println(v6, v1, v2)
    }
    switch v6 {
    case "":
        var v14 float32 = (0.0 - -2.25)
        _ = v14
    case "hello":
        v2 += v4
    }
    v3 = v3
    var v15 int16 = 1000
    switch "\n\t" {
    case "hello":
        var v16 bool = (!(false && true))
        _ = v16
    case "":
        v15 = (v15 ^ (v15 << ((2) & 63)))
    case "\n\t":
        if (621 <= 2) {
            var v17 uint32 = (v1 ^ v1)
            var v18 int64 = 0
            v19 := []int{-1, -139, -205}
            _ = v17
            _ = v18
            _ = v19
        }
    default:
        var v20 uint32 = v1
        _ = v20
    }
    fmt.Println(v5, v15, v2)
    if (!(9223372036854775807 == -1)) {
        if true {
            v21 := []int{}
            if true {
                switch 2147483647 {
                case 867:
                    fmt.Println(v15, v6, v2)
                }
                fmt.Println(v2, v5, v3)
                for fi22 := 0; fi22 < 15; fi22++ {
                    for fi23 := 0; fi23 < 6; fi23++ {
                        v24 := GenS2{F0: 65535}
                        var v25 float64 = (float64(0.0) * 0.0)
                        var v26 uint8 = 42
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    v1 -= 669
                    _ = fi22
                }
            } else {
                var v27 int16 = (v15 << ((255) & 63))
                var v28 int8 = v5
                switch (("\n\t" + v6) + v6) {
                case "hello":
                    v1 = v1
                case "foo bar":
                    var v29 string = (v6 + v6)
                    _ = v29
                }
                _ = v27
                _ = v28
            }
            _ = v21
        } else {
            fmt.Println(v4, v3, v4)
        }
    } else {
        for fi30 := 0; fi30 < 14; fi30++ {
            v31 := GenS3{F0: "", F1: fi30, F2: 2}
            fmt.Println(v15, v1)
            _ = fi30
            _ = v31
        }
        var v32 float32 = 0.0
        var v33 int16 = v15
        _ = v32
        _ = v33
    }
    var v34 uint = 42
    var v35 float64 = 0.1
    switch 9 {
    case 995:
        for fi36 := 0; fi36 < 14; fi36++ {
            if (42 > 9223372036854775807) {
                fmt.Println(v2, v3)
            }
            v15 = v15
            var v37 float64 = 0.1
            _ = fi36
            _ = v37
        }
    case 42:
        v38 := []int{32768, 170, 754, -9223372036854775808}
        _ = v38
    }
    fmt.Println(v3)
    var v39 string = "foo bar"
    v35 += 3.14159
    var v40 int8 = (-(v5))
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
