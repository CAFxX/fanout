package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
func genFn0(p0 bool, p1 uint) uint64 {
    switch "foo bar" {
    case "":
        if (-9223372036854775808 < 42) {
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := GenS1{F0: 1}
                _ = fi7
                _ = v8
            }
        } else {
            var v9 uint = p1
            p1 = (p1 - p1)
            if true {
                var v10 int64 = -9223372036854775808
                _ = v10
            }
            _ = v9
        }
    case "\n\t":
        var v11 uint = p1
        _ = v11
    case "foo bar":
        p1 -= 0
    }
    switch "a" {
    case "\n\t":
        if (!p0) {
            fmt.Println(p1, p0)
            for fi12 := 0; fi12 < 1; fi12++ {
                switch (("" + "foo bar") + "hello") {
                case "foo bar":
                    for fi13 := 0; fi13 < 8; fi13++ {
                        p0 = (fi12 <= (fi13 >> ((339) & 63)))
                        v14 := GenS1{F0: 0}
                        var v15 uint = p1
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                case "hello":
                    var v16 string = ""
                    _ = v16
                case "\n\t":
                    var v17 float64 = 0.1
                    _ = v17
                }
                v18 := []int{2147483647, fi12, fi12}
                fmt.Println(p1)
                _ = fi12
                _ = v18
            }
            p0 = true
        }
    case "a":
        var v19 uint64 = 0
        _ = v19
    }
    v20 := []int{int(p1), -1}
    var v21 uint16 = 226
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return 341
}
func genFn1() int {
    v22 := []int{521, -1}
    var v23 float64 = 1e10
    _ = v22
    _ = v23
    return 128
}

func main() {
    var v1 int8 = 0
    var v2 uint = 2
    var v3 string = "a"
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    var v6 float64 = 1e10
    switch 48 {
    case -1:
        v2 -= (v2 ^ v2)
    }
    var v24 int16 = 0
    fmt.Println(v24)
    switch "\n\t" {
    case "\n\t":
        var v25 uint32 = 935
        _ = v25
    case "":
        v26 := []int{genFn1(), 128}
        _ = v26
    default:
        var v27 uint8 = 217
        _ = v27
    }
    var v28 uint64 = 2
    v6 = float64(0.1)
    if ((false && true) && (!false)) {
        var v29 string = v3
        _ = v29
    }
    var v30 uint64 = genFn0(true, 18446744073709551615)
    var v31 int16 = v24
    fmt.Println(v28, v1, v2)
    v32 := GenS2{F0: 537}
    var v33 int16 = v24
    fmt.Println(v2, v31, v2)
    fmt.Println(v3, v28)
    v28 -= v28
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
