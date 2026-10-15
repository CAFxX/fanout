package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 bool = (false || ((3.14159 * 0.1) <= 3.14159))
    v2 := make(map[string]string)
    var v3 uint64 = 4294967295
    var v4 string = "hello"
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 uint16 = 2
        _ = fi5
        _ = v6
    }
    var v7 int = 32768
    switch (v4 + v4) {
    case "":
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 bool = v1
            var v10 int8 = 1
            switch fi8 {
            case 0:
                if true {
                    v4 = v4
                } else {
                    v11 := GenS1{F0: -1}
                    v1 = v9
                    if v1 {
                        v1 = false
                        v12 := GenS1{F0: -116}
                        v1 = false
                        _ = v12
                    } else {
                        _ = len(v2)
                    }
                    _ = v11
                }
            case -645:
                v13 := GenS1{F0: 9223372036854775807}
                _ = v13
            default:
                for fi14 := 0; fi14 < 17; fi14++ {
                    switch "a" {
                    case "\n\t":
                        if (255 <= 14) {
                            var v15 bool = true
                            _ = v15
                        }
                    case "a":
                        v7 = fi8
                    case "hello":
                        v1 = true
                    }
                    _ = fi14
                }
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    if (float64(0.1) >= -2.25) {
        v16 := GenS1{F0: -9223372036854775808}
        var v17 int32 = 28
        _ = v16
        _ = v17
    }
    var v18 string = v4
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS2{F0: 2, F1: 1, F2: 127}
        if v1 {
            var v21 uint32 = 999
            _ = v21
        }
        v4 = "\n\t"
        _ = fi19
        _ = v20
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 int16 = -1
        v24 := []int{fi22}
        v1 = (2147483648 >= 1)
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := []int{(v7 + -1), v7}
    switch v18 {
    case "a":
        if false {
            fmt.Println(v4, v7)
            v7 = v7
        } else {
            v7 -= v7
            switch "" {
            case "":
                v26 := GenS2{F0: 2, F1: 4294967295, F2: 100}
                _ = v26
            case "\n\t":
                var v27 float32 = float32(1.5)
                _ = v27
            case "a":
                var v28 int = (v7 * 2147483647)
                _ = v28
            }
        }
    case "\n\t":
        _ = len(v2)
    case "":
        switch "" {
        case "hello":
            for fi29 := 0; fi29 < 15; fi29++ {
                v1 = (854 > 42)
                fmt.Println(v4)
                switch v18 {
                case "\n\t":
                    if (fi29 < v7) {
                        v3 -= v3
                        v30 := GenS3{F0: -280}
                        var v31 uint64 = (^(v3))
                        _ = v30
                        _ = v31
                    } else {
                        v3 -= 930
                    }
                }
                _ = fi29
            }
        }
    default:
        fmt.Println(v18)
    }
    v2[v4] = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v7
}
