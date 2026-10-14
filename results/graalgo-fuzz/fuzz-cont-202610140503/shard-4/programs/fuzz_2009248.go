package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 string }
func genFn0(p0 uint, p1 int, p2 uint32) bool {
    p1 -= (p1 % (((p1 % ((p1) | 1))) | 1))
    v4 := []int{}
    switch p1 {
    case 32767:
        p0 += (p0 & p0)
    case -2147483648:
        var v5 int16 = -811
        _ = v5
    case 42:
        var v6 string = ((("a" + "hello") + (("" + "foo bar") + ("\n\t" + "\n\t"))) + (("foo bar" + ("foo bar" + ("hello" + "a"))) + ("\n\t" + ("\n\t" + "hello"))))
        _ = v6
    default:
        switch ("" + "a") {
        case "a":
            fmt.Println(p1, p0)
        }
    }
    if ((32767 <= -1) && (0 >= -190)) {
        fmt.Println(p1, p1)
    } else {
        p0 += 4294967295
        _ = len(v4)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (999 <= (p2 % ((p2) | 1)))
}
func genFn1(p0 uint64, p1 int) string {
    var v7 uint64 = p0
    fmt.Println(p0, p1, p1)
    v8 := []int{(p1 * p1), (-(p1))}
    for fi9 := 0; fi9 < 15; fi9++ {
        fmt.Println(fi9, v7)
        var v10 uint64 = 42
        _ = fi9
        _ = v10
    }
    v8 = append(v8, p1)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return "a"
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{867}
    var v3 uint = 255
    fmt.Println(v3)
    v3 -= (v3 % (((v3 * 861)) | 1))
    if (!false) {
        for fi11 := 0; fi11 < 2; fi11++ {
            switch "hello" {
            case "foo bar":
                v12 := GenS1{F0: -2147483648}
                _ = v12
            case "":
                v13 := GenS2{F0: 551, F1: ("hello" + (("\n\t" + "foo bar") + "\n\t"))}
                _ = v13
            }
            v14 := GenS1{F0: 631}
            _ = fi11
            _ = v14
        }
    } else {
        var v15 bool = false
        fmt.Println(v3)
        v2 = append(v2, 0)
        _ = v15
    }
    var v16 int16 = 42
    v3 = 143
    if true {
        fmt.Println(v16)
    }
    v1[("a" + "a")] = true
    if true {
        v17 := GenS1{F0: 531}
        v3 = v3
        _ = v17
    } else {
        var v18 int16 = v16
        v3 -= v3
        _ = v18
    }
    switch -74 {
    case 127:
        fmt.Println(v3, v16)
    case 120:
        if false {
            fmt.Println(v16, v3)
        }
    }
    var v19 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
