package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint; F2 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    v3 := []string{}
    v4 := make(map[uint32]uint32)
    var v5 bool = true
    v6 := make(map[int]int)
    if v5 {
        v5 = (!(true && true))
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(v5, v5)
            for fi8 := 0; fi8 < 7; fi8++ {
                if (100 < 88) {
                    v5 = v5
                    v5 = true
                } else {
                    var v9 uint16 = 255
                    var v10 int64 = 951
                    _ = v10
                    _ = v9
                }
                v11 := GenS1{F0: 65535, F1: ("foo bar" + "a"), F2: 42}
                if (3.14159 != 0.1) {
                    fmt.Println(fi8, v5)
                    fmt.Println(v5, fi7, fi8)
                }
                _ = fi8
                _ = v11
            }
            _ = fi7
        }
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 float64 = (float64(-2.25) * (((((1.5 / 1.5) + float64(0.1)) / 1.5) * ((((float64((0.0 / 0.1)) + 3.14159) / ((-2.25 / 1.0) + 0.1)) * 3.14159) - 0.1)) + 1e10))
            v5 = ((!v5) && v5)
            _ = fi12
            _ = v13
        }
    }
    var v14 uint64 = 2
    v15 := GenS2{F0: 2}
    v14 = (^(v14))
    switch "a" {
    case "":
        for fi16 := 0; fi16 < 17; fi16++ {
            if false {
                v17 := GenS3{F0: "\n\t", F1: 901, F2: v5}
                fmt.Println(fi16, v14, fi16)
                _ = v17
            } else {
                v5 = (!(0 >= 1))
                var v18 float64 = 1.5
                _ = v18
            }
            _ = fi16
        }
    case "a":
        v19 := GenS3{F0: ("" + "\n\t"), F1: 4294967295, F2: v5}
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v14 = 42
        _ = fi20
    }
    var v21 bool = (424 >= 2147483647)
    switch "\n\t" {
    case "foo bar":
        v5 = ((-32768 < -1) || false)
    }
    v5 = v5
    var v22 int64 = 9223372036854775807
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 int64 = (v22 ^ (-(v22)))
        v22 = v24
        _ = fi23
        _ = v24
    }
    if true {
        v22 = -1
    }
    v21 = (v22 < v22)
    v22 += v22
    var v25 uint64 = (v14 | 4294967295)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
