package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint, p1 int, p2 uint) int64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        switch "hello" {
        case "foo bar":
            p2 = p2
        default:
            switch p1 {
            case -1:
                v7 := GenS1{F0: fi6, F1: 18446744073709551615, F2: -1}
                _ = v7
            case 127:
                p2 = p0
            }
        }
        var v8 uint32 = 670
        if true {
            fmt.Println(p0, p1, p2)
            v9 := []int{0, p1}
            _ = v9
        }
        _ = fi6
        _ = v8
    }
    if false {
        for fi10 := 0; fi10 < 2; fi10++ {
            p2 = (p0 - (p0 * p0))
            _ = fi10
        }
    }
    p1 = 421
    if (!true) {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 int = p1
            _ = fi11
            _ = v12
        }
    } else {
        p1 = p1
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 int = p1
        if false {
            v14 = p1
            v14 = fi13
        }
        switch (fi13 % (((v14 << (((p0 / ((p2) | 1))) & 63))) | 1)) {
        case 501:
            var v15 string = "a"
            _ = v15
        case -1:
            v16 := GenS1{F0: int(p0), F1: 1, F2: -1}
            _ = v16
        default:
            v17 := []int{(p1 >> ((p2) & 63)), p1, -9223372036854775808}
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := []string{"hello", ("" + (("foo bar" + "a") + "hello")), ("foo bar" + "foo bar"), ""}
    var v2 uint32 = 0
    var v3 float64 = float64(1.5)
    v4 := []int{32768, 42, 42, 317}
    v5 := []string{"\n\t", "", "", "hello"}
    fmt.Println(v3, v2)
    var v18 int16 = 42
    var v19 uint16 = 255
    if true {
        v2 = v2
        var v20 float64 = (v3 + v3)
        var v21 int32 = -1
        _ = v20
        _ = v21
    }
    switch ("\n\t" + "\n\t") {
    case "hello":
        v22 := GenS2{F0: 0}
        _ = v22
    case "\n\t":
        if (65535 > 1) {
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
            v2 = v2
            v18 = v18
        }
    }
    v2 += (v2 | v2)
    _ = len(v5)
    var v23 int32 = 0
    v19 += (v19 >> ((2) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
