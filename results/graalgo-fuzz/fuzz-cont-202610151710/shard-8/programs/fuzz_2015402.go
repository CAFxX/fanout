package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint16, p1 int, p2 int64) string {
    switch (^(p1)) {
    case -9223372036854775808:
        var v5 uint16 = (p0 + 2)
        _ = v5
    }
    v6 := GenS1{F0: 4294967295}
    v7 := []int{1, p1}
    v8 := GenS2{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return ("a" + "hello")
}

func main() {
    var v1 string = ("" + "hello")
    var v2 float64 = -2.25
    v3 := make(map[int]bool)
    var v4 int8 = 1
    fmt.Println(v4)
    v9 := GenS3{F0: 125}
    v1 = ""
    switch v1 {
    case "\n\t":
        fmt.Println(v2, v4, v2)
    }
    fmt.Println(v1, v2, v1)
    v10 := []int{}
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(fi12)
            _ = fi12
        }
        _ = len(v10)
        _ = fi11
    }
    v13 := GenS4{F0: (v1 + ("\n\t" + v1))}
    var v14 uint = 898
    v14 += (v14 % (((v14 << ((18446744073709551615) & 63))) | 1))
    v14 = v14
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 float32 = (1e10 / 1.0)
        if false {
            var v17 int8 = (v4 & 100)
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    v1 = v1
    for fi18 := 0; fi18 < 12; fi18++ {
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := GenS5{F0: 255, F1: 255}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    fmt.Println(v14)
    var v21 uint = 65535
    v22 := GenS1{F0: (v14 * v21)}
    switch "a" {
    case "\n\t":
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 float64 = float64((0.1 * 0.0))
            _ = fi23
            _ = v24
        }
    case "foo bar":
        if (-88 >= (^(v4))) {
            v1 = ((v1 + "\n\t") + v1)
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = len(v10)
        } else {
            var v25 int = 127
            _ = v25
        }
    case "a":
        var v26 bool = false
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
