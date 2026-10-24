package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int16 }
type GenS2 struct { F0 int8; F1 int64; F2 bool }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint) uint64 {
    switch (("" + "a") + "\n\t") {
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            v7 := GenS1{F0: 324, F1: 426, F2: -1}
            _ = v7
        case "":
            fmt.Println(p0)
        }
    case "\n\t":
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int8 = 127
            v10 := []int{0, fi8, (^(fi8))}
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    if (false && true) {
        var v11 int16 = -32768
        _ = v11
    } else {
        switch "hello" {
        case "hello":
            p0 = p0
        }
        if (false || ((p0 ^ (p0 | (p0 + (^(p0))))) == (p0 / ((42) | 1)))) {
            p0 -= 2
        } else {
            var v12 uint16 = 498
            _ = v12
        }
    }
    fmt.Println(p0)
    if ((((2 >= 37) || (!((false && false) && (!(float64(0.0) < 1e10))))) || true) || false) {
        var v13 int64 = 1
        _ = v13
    }
    _ = p0
    return 18446744073709551615
}
func genFn1() int32 {
    var v14 float64 = 0.0
    v15 := GenS2{F0: -128, F1: -1, F2: false}
    v14 = (0.0 - float64((float32((1.5 / 0.1)) + 3.14159)))
    _ = v14
    _ = v15
    return -1
}

func main() {
    var v1 float64 = 1e10
    var v2 string = "\n\t"
    var v3 uint = 18446744073709551615
    v4 := []string{}
    var v5 uint64 = 4294967295
    var v6 int32 = -329
    var v16 bool = (3.14159 < (3.14159 / 0.1))
    var v17 string = (v2 + v2)
    fmt.Println(v17, v1, v3)
    v18 := GenS2{F0: -128, F1: -15, F2: v16}
    var v19 int64 = 9223372036854775807
    var v20 uint64 = (v5 * v5)
    v21 := GenS3{F0: 486}
    switch "" {
    case "\n\t":
        var v22 int64 = (v19 >> ((v3) & 63))
        _ = v22
    }
    v17 = ""
    switch 2147483647 {
    case -9223372036854775808:
        fmt.Println(v1, v6)
    case 2147483647:
        var v23 string = "a"
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v20 += genFn0((^(v3)))
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
