package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 float64 = float64(3.14159)
    v2 := make(map[string]int)
    var v3 float64 = 0.0
    var v4 bool = (false || (true && (42 <= -2147483648)))
    var v5 bool = (v4 || (0.1 != 1.5))
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(v1, v5)
        _ = fi6
    }
    if v4 {
        delete(v2, (("\n\t" + "\n\t") + "a"))
        for fi7 := 0; fi7 < 9; fi7++ {
            if v4 {
                v4 = v5
                v4 = false
            }
            _ = fi7
        }
        var v8 int64 = 357
        _ = v8
    } else {
        var v9 int32 = -470
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int64 = 9223372036854775807
            _ = fi10
            _ = v11
        }
        v12 := []int{-394, -1}
        _ = v12
        _ = v9
    }
    v13 := []int{-9223372036854775808, 4294967295, 32768, 2147483647}
    fmt.Println(v1, v3)
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS1{F0: 42}
        _ = fi14
        _ = v15
    }
    v5 = false
    switch ("hello" + "foo bar") {
    case "a":
        for fi16 := 0; fi16 < 5; fi16++ {
            v4 = (-1 != 100)
            v4 = v4
            if (-1 <= 2147483647) {
                var v17 uint16 = 2
                v18 := GenS1{F0: 2}
                _ = v17
                _ = v18
            }
            _ = fi16
        }
    case "":
        switch 9223372036854775807 {
        case -1:
            for fi19 := 0; fi19 < 3; fi19++ {
                v4 = (1 >= 32767)
                _ = fi19
            }
        case 4294967295:
            switch "hello" {
            case "hello":
                v4 = (true && (123 >= 0))
            case "\n\t":
                v20 := GenS1{F0: 224}
                _ = v20
            case "":
                v21 := GenS1{F0: 42}
                _ = v21
            }
        case 2147483647:
            var v22 bool = v5
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        fmt.Println(v1, v5, v1)
        _ = fi23
    }
    v5 = (1 < 127)
    _ = len(v13)
    for fi24 := 0; fi24 < 19; fi24++ {
        switch (fi24 ^ (fi24 >> ((2) & 63))) {
        case 2147483647:
            v25 := GenS2{F0: 16, F1: 1}
            _ = v25
        case 1:
            for fi26 := 0; fi26 < 12; fi26++ {
                switch "hello" {
                case "foo bar":
                    v5 = v5
                default:
                    v27 := GenS1{F0: 2}
                    _ = v27
                }
                v28 := GenS3{F0: fi24}
                v29 := GenS4{F0: 4294967295, F1: -953, F2: 2}
                _ = fi26
                _ = v28
                _ = v29
            }
        case -9223372036854775808:
            switch "a" {
            case "foo bar":
                v2["hello"] = fi24
            case "\n\t":
                v30 := GenS5{F0: 896}
                _ = v30
            }
        }
        _ = fi24
    }
    switch "hello" {
    case "hello":
        var v31 int16 = 1000
        _ = v31
    case "a":
        var v32 int8 = 0
        _ = v32
    }
    v13 = append(v13, -981)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
