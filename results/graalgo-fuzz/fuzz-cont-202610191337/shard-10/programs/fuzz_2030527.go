package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 bool; F1 int64; F2 uint16 }

func main() {
    var v1 float32 = 3.14159
    var v2 float32 = v1
    var v3 int16 = -1
    var v4 float64 = 1e10
    var v5 int16 = (v3 / (((-(v3))) | 1))
    v6 := make(map[string]bool)
    v7 := GenS1{F0: -2147483648, F1: 2, F2: 769}
    if (true || (-776 != -1)) {
        var v8 string = ((("" + "a") + "foo bar") + ("\n\t" + "hello"))
        fmt.Println(v5)
        _ = v8
    } else {
        if false {
            v9 := v6["\n\t"]
            if (255 <= 1) {
                fmt.Println(v4, v5)
            }
            _ = v9
        } else {
            var v10 uint16 = 1
            v11 := GenS2{F0: 643, F1: 658}
            for fi12 := 0; fi12 < 8; fi12++ {
                if false {
                    v13 := GenS1{F0: (fi12 * 497), F1: (v10 | v10), F2: -9223372036854775808}
                    var v14 uint8 = 170
                    _ = v13
                    _ = v14
                } else {
                    var v15 uint8 = 117
                    _ = v15
                }
                _ = fi12
            }
            _ = v10
            _ = v11
        }
    }
    switch "foo bar" {
    case "a":
        if (0 == 931) {
            v16 := GenS3{F0: 4294967295, F1: 42}
            _ = v16
        } else {
            v17 := GenS1{F0: 42, F1: 65535, F2: -927}
            _ = v17
        }
    case "hello":
        v1 += v1
    }
    switch -23 {
    case 4294967295:
        for fi18 := 0; fi18 < 13; fi18++ {
            v1 -= v1
            var v19 uint32 = 2
            v20 := GenS4{F0: (v19 + (v19 ^ v19)), F1: -2147483648}
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    var v21 uint = 4294967295
    var v22 float64 = 0.1
    var v23 string = ""
    v3 -= 1
    v23 = ""
    v2 -= 0.1
    v1 = (-2.25 * v2)
    var v24 bool = (127 <= -72)
    switch 314 {
    case -9223372036854775808:
        var v25 uint32 = 374
        _ = v25
    case 42:
        v3 = (v5 | (v5 | -1))
    case 359:
        v26 := GenS5{F0: v21}
        _ = v26
    }
    if v24 {
        fmt.Println(v1)
        v6[v23] = (v3 == v5)
        v27 := GenS6{F0: "a", F1: 42}
        _ = v27
    } else {
        if true {
            fmt.Println(v2)
            v4 = -2.25
            switch v23 {
            case "foo bar":
                v4 += v22
            }
        } else {
            for fi28 := 0; fi28 < 6; fi28++ {
                v29 := []int{fi28, (fi28 - (^(fi28))), (fi28 / (((fi28 >> (((^(v21))) & 63))) | 1))}
                fmt.Println(v3, v2, v5)
                _ = fi28
                _ = v29
            }
            var v30 uint32 = 453
            switch 9223372036854775807 {
            case 9223372036854775807:
                v24 = true
            case 42:
                for fi31 := 0; fi31 < 15; fi31++ {
                    var v32 int64 = -652
                    _ = fi31
                    _ = v32
                }
            case 0:
                var v33 uint16 = 445
                _ = v33
            }
            _ = v30
        }
        if v24 {
            for fi34 := 0; fi34 < 18; fi34++ {
                v5 = (v5 | v3)
                _ = fi34
            }
        }
    }
    v35 := []int{9223372036854775807, 4294967295, 1}
    switch (v23 + (v23 + v23)) {
    case "a":
        v21 = v21
    case "\n\t":
        v36 := GenS7{F0: (((v21 * (v21 - v21)) <= v21) || ((-2.25 * v2) <= float32(3.14159))), F1: -934, F2: 0}
        _ = v36
    default:
        v35 = append(v35, 1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
