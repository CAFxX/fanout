package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int8 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint; F1 uint; F2 int32 }
type GenS7 struct { F0 int16; F1 int16 }

func main() {
    var v1 uint16 = 2
    var v2 bool = false
    var v3 uint32 = 1
    v4 := GenS1{F0: 42, F1: (("" + "foo bar") + "\n\t"), F2: 100}
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS2{F0: 1, F1: 42}
        if true {
            switch "" {
            case "hello":
                v2 = v2
            case "":
                var v7 int16 = -917
                _ = v7
            default:
                v1 += (v1 | v1)
            }
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 uint = 459
                _ = fi8
                _ = v9
            }
            var v10 uint32 = v3
            _ = v10
        } else {
            var v11 int8 = 127
            v11 -= -12
            fmt.Println(v1, v3, fi5)
            _ = v11
        }
        _ = fi5
        _ = v6
    }
    var v12 int8 = 42
    v2 = ((v1 | v1) >= (v1 | 65535))
    var v13 bool = false
    v12 += (^(v12))
    if (v2 && (2 == 2)) {
        if (v2 || (-1 < 42)) {
            v14 := GenS3{F0: 255}
            _ = v14
        }
        fmt.Println(v13)
    }
    v15 := GenS4{F0: v3}
    v3 = 0
    var v16 float64 = ((0.1 - 3.14159) * ((((-2.25 - 0.1) + 1.5) - 0.0) * 3.14159))
    switch "\n\t" {
    case "hello":
        for fi17 := 0; fi17 < 16; fi17++ {
            switch (("foo bar" + "a") + "") {
            case "hello":
                var v18 uint32 = 1
                _ = v18
            case "foo bar":
                for fi19 := 0; fi19 < 15; fi19++ {
                    v13 = v13
                    v16 = (v16 * v16)
                    _ = fi19
                }
            case "":
                var v20 string = ("a" + "hello")
                _ = v20
            }
            switch (("\n\t" + ("" + "hello")) + "a") {
            case "a":
                var v21 int8 = (-(v12))
                _ = v21
            default:
                if (-288 > -1) {
                    fmt.Println(fi17)
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v23 := []int{}
                        _ = fi22
                        _ = v23
                    }
                }
            }
            v24 := GenS3{F0: 42}
            _ = fi17
            _ = v24
        }
    case "":
        v25 := GenS5{F0: 42, F1: (v3 / ((1) | 1)), F2: -1}
        _ = v25
    }
    v26 := GenS6{F0: 809, F1: 65535, F2: 42}
    v27 := GenS6{F0: 255, F1: 1, F2: -1}
    v16 -= (0.1 + (v16 / (v16 + v16)))
    v28 := GenS7{F0: 1, F1: -823}
    var v29 uint64 = 273
    var v30 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
