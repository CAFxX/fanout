package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 float32 = (-2.25 * (0.0 * 0.0))
        var v8 string = ""
        var v9 uint = 42
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 float32 = (1e10 - (3.14159 + (3.14159 / 1e10)))
    fmt.Println(v10)
    _ = v10
    return 1
}
func genFn1(p0 int64, p1 uint, p2 uint32) uint16 {
    var v11 int8 = -128
    switch ("hello" + "") {
    case "a":
        v12 := []int{}
        _ = v12
    case "":
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 uint = 42
            fmt.Println(v14, p1, v14)
            _ = fi13
            _ = v14
        }
    case "foo bar":
        var v15 uint8 = 0
        _ = v15
    }
    v16 := []int{127}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v16
    return 255
}

func main() {
    var v1 int = 127
    var v2 uint8 = 0
    var v3 float32 = 0.1
    var v4 string = ("foo bar" + (("hello" + "hello") + ("a" + ("foo bar" + "a"))))
    v5 := []uint32{}
    switch (v1 + (v1 + v1)) {
    case 32768:
        var v17 string = "hello"
        _ = v17
    default:
        v3 = v3
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        switch "" {
        case "":
            v1 = 128
        case "hello":
            v19 := GenS1{F0: (false && false), F1: 461}
            _ = v19
        case "a":
            if (false && (v4 <= (v4 + v4))) {
                if true {
                    v4 = v4
                    switch (fi18 - v1) {
                    case 128:
                        fmt.Println(v2, fi18, fi18)
                    }
                }
            }
        }
        if false {
            fmt.Println(v1, v4, v3)
        } else {
            for fi20 := 0; fi20 < 6; fi20++ {
                v4 = ((v4 + (v4 + v4)) + v4)
                _ = fi20
            }
        }
        v21 := GenS1{F0: (!false), F1: 255}
        _ = fi18
        _ = v21
    }
    v3 -= v3
    v1 -= (v1 * (v1 % ((v1) | 1)))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        var v22 int = 128
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 uint8 = (v2 * (v2 | 42))
            _ = fi23
            _ = v24
        }
        fmt.Println(v4, v4, v22)
        _ = v22
    } else {
        if len(v5) > 0 {
            v5[0] = 209
        }
    }
    var v25 uint64 = 2
    v2 = 127
    var v26 int32 = -1
    var v27 uint = 18446744073709551615
    for fi28 := 0; fi28 < 12; fi28++ {
        switch ("\n\t" + v4) {
        case "":
            var v29 float64 = float64(v3)
            _ = v29
        case "\n\t":
            if false {
                var v30 uint64 = (v25 / ((4294967295) | 1))
                fmt.Println(v30, v26)
                _ = v30
            } else {
                v31 := GenS1{F0: ("foo bar" == v4), F1: 65535}
                if ((4294967295 <= (v25 & (v25 + v25))) || true) {
                    v1 = v1
                } else {
                    var v32 string = v4
                    _ = v32
                }
                _ = v31
            }
        case "foo bar":
            v27 -= v27
        }
        _ = fi28
    }
    fmt.Println(v25, v25, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
