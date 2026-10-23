package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int8, p1 int8, p2 int) string {
    switch 42 {
    case 2147483647:
        var v7 int = p2
        _ = v7
    case 32768:
        if false {
            p2 += p2
        }
    default:
        var v8 uint = 218
        _ = v8
    }
    p2 = (p2 ^ p2)
    _ = p0
    _ = p1
    _ = p2
    return "a"
}
func genFn1(p0 uint, p1 bool, p2 uint64) uint {
    p1 = (9223372036854775807 > 9223372036854775807)
    var v9 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p0
}

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 608
    var v3 int16 = 42
    v4 := make(map[string]int)
    var v5 uint16 = (v1 - 1)
    var v6 bool = true
    v4[(((("" + (("hello" + "foo bar") + (((("a" + "a") + "hello") + "foo bar") + ""))) + "") + ("a" + "foo bar")) + ("foo bar" + "\n\t"))] = 4294967295
    v10 := GenS1{F0: 128}
    fmt.Println(v2, v2)
    var v11 bool = true
    var v12 bool = (1 <= 42)
    var v13 int8 = 100
    if (!(374 >= 1)) {
        fmt.Println(v13)
        v14 := GenS1{F0: -1}
        _ = v14
    } else {
        if (!(2 != 230)) {
            v15 := []int{-1, 1, 0}
            v16 := GenS1{F0: 2147483647}
            _ = v15
            _ = v16
        }
    }
    v1 -= (v5 ^ v1)
    v6 = (-93 < -734)
    fmt.Println(v11, v2)
    fmt.Println(v1, v11, v13)
    v3 = 219
    if v11 {
        if (false && false) {
            v17 := GenS1{F0: 32768}
            v2 = 371
            _ = v17
        }
        fmt.Println(v13, v11)
        delete(v4, "foo bar")
    } else {
        if (int32(v13) > 0) {
            var v18 int64 = -9223372036854775808
            _ = v18
        }
        var v19 uint16 = 255
        switch 1 {
        case -2147483648:
            var v20 uint8 = 253
            _ = v20
        case 42:
            if true {
                switch "hello" {
                case "\n\t":
                    fmt.Println(v6, v11)
                case "foo bar":
                    v12 = ((!(v11 || (((true || true) && (38 < 174)) && (32768 == 967)))) || true)
                }
            } else {
                fmt.Println(v3, v12, v11)
            }
        case 2147483647:
            v21 := GenS1{F0: 0}
            _ = v21
        }
        _ = v19
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := GenS1{F0: (fi22 << ((255) & 63))}
        if true {
            for fi24 := 0; fi24 < 18; fi24++ {
                fmt.Println(v6, v11)
                _ = fi24
            }
            var v25 uint16 = (v5 + (v5 - 32767))
            fmt.Println(fi22, v5)
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    var v26 uint16 = 255
    var v27 int32 = 0
    v28 := GenS2{F0: 2}
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
