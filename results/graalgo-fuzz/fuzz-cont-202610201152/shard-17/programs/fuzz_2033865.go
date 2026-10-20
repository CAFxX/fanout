package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint32, p1 int64, p2 int8) int {
    var v5 string = "foo bar"
    if false {
        switch "foo bar" {
        case "foo bar":
            var v6 int32 = -2147483648
            _ = v6
        case "":
            v7 := GenS1{F0: 540}
            _ = v7
        default:
            var v8 float64 = (3.14159 * (0.1 + -2.25))
            _ = v8
        }
        var v9 uint64 = 4294967295
        v5 = (("" + (v5 + ((v5 + ("hello" + "foo bar")) + ""))) + v5)
        _ = v9
    } else {
        p0 = p0
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := GenS2{F0: 1, F1: p2}
            _ = fi10
            _ = v11
        }
        v12 := []int{}
        _ = v12
    }
    v13 := []int{851}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v5
    return -899
}
func genFn1(p0 string, p1 uint, p2 int) uint64 {
    p1 += (p1 & p1)
    v14 := []int{(p2 ^ p2), p2, p2}
    if false {
        if (p0 >= p0) {
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 uint64 = 657
                _ = fi15
                _ = v16
            }
            p0 = ""
        }
        var v17 int16 = 235
        _ = v17
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 0
}

func main() {
    var v1 uint = 145
    v2 := make(map[uint32]bool)
    var v3 float64 = 1e10
    v4 := make(map[string]bool)
    fmt.Println(v1)
    for fi18 := 0; fi18 < 9; fi18++ {
        v1 = v1
        switch ("" + (("foo bar" + ("\n\t" + "")) + ("foo bar" + ("foo bar" + "hello")))) {
        case "a":
            fmt.Println(fi18, v1)
        default:
            v3 = v3
        }
        _ = fi18
    }
    delete(v4, "foo bar")
    fmt.Println(v3, v1)
    v19 := GenS1{F0: 32767}
    if true {
        v4["hello"] = false
    }
    switch -791 {
    case 2147483647:
        var v20 int8 = 42
        _ = v20
    }
    fmt.Println(v1, v3)
    var v21 int64 = int64(v1)
    fmt.Println(v3, v1)
    v22 := GenS1{F0: 1}
    v23 := []int{genFn0(427, 210, -87), 9223372036854775807, -836, 0}
    var v24 uint64 = 498
    for fi25 := 0; fi25 < 19; fi25++ {
        if len(v23) > 0 {
            v23[0] = fi25
        }
        var v26 uint8 = 42
        var v27 int8 = -1
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        v3 -= v3
        fmt.Println(v21, fi28, v21)
        var v29 uint64 = v24
        _ = fi28
        _ = v29
    }
    v24 = genFn1("hello", (v1 & (v1 ^ v1)), 62)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
