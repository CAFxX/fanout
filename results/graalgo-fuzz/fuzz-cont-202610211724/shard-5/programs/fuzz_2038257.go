package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
func genFn0(p0 uint64, p1 uint8, p2 uint16) string {
    v7 := GenS1{F0: (("\n\t" + "hello") + (("hello" + "\n\t") + "hello")), F1: "hello"}
    p2 = (^(p2))
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := []int{fi8, fi8, (-(fi8)), (fi8 | int(p1))}
        _ = fi8
        _ = v9
    }
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("\n\t" + "hello")
}
func genFn1(p0 uint16, p1 int16) uint {
    var v10 string = "\n\t"
    p0 -= p0
    _ = p0
    _ = p1
    _ = v10
    return 0
}

func main() {
    v1 := []uint32{0}
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    var v5 bool = (!(float32((0.1 * 0.1)) < ((3.14159 * (float32((1.5 / 1e10)) / 1.0)) - 1.5)))
    v6 := make(map[int]uint32)
    for fi11 := 0; fi11 < 9; fi11++ {
        delete(v2, fi11)
        _ = fi11
    }
    v5 = true
    if ("" == ("hello" + (("" + "foo bar") + "\n\t"))) {
        var v12 float64 = 0.1
        v12 = 0.0
        _ = v12
    } else {
        v5 = (!(42 <= 65535))
        fmt.Println(v5)
    }
    if (!(2147483648 < 42)) {
        var v13 uint64 = 0
        var v14 uint32 = 2
        var v15 float64 = 1e10
        _ = v13
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 16; fi16++ {
            v5 = v5
            _ = fi16
        }
        var v17 int64 = -9223372036854775808
        _ = v17
    }
    var v18 uint16 = 0
    if ((!v5) || (0 != 42)) {
        var v19 string = ("hello" + "hello")
        var v20 uint32 = 684
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v19
        _ = v20
    } else {
        var v21 int8 = 42
        _ = v21
    }
    v5 = (v5 || (1.5 <= 0.1))
    v22 := GenS1{F0: "a", F1: "hello"}
    var v23 string = "\n\t"
    fmt.Println(v23, v23)
    var v24 uint16 = (^(v18))
    fmt.Println(v18, v24, v23)
    v25 := GenS1{F0: v23, F1: (v23 + ("" + v23))}
    var v26 uint = genFn1((v18 % (((v24 & 1)) | 1)), 0)
    var v27 float64 = 0.1
    v23 = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
