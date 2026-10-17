package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int; F1 bool }
func genFn0(p0 bool) string {
    p0 = p0
    for fi4 := 0; fi4 < 10; fi4++ {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(p0, fi4)
        default:
            fmt.Println(p0)
        }
        var v5 uint64 = 18446744073709551615
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 42, F1: "hello"}
    _ = p0
    _ = v6
    return ("a" + ((("\n\t" + ((("foo bar" + "a") + ((("hello" + (((("" + ("" + ("foo bar" + "a"))) + "\n\t") + "a") + "hello")) + "a") + "a")) + "a")) + (("foo bar" + ("foo bar" + "hello")) + (("hello" + "hello") + ""))) + "foo bar"))
}
func genFn1(p0 uint, p1 bool, p2 int) int {
    fmt.Println(p2)
    fmt.Println(p0, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []string{}
    var v3 uint16 = 32767
    var v7 int8 = 42
    fmt.Println(v3, v1, v3)
    var v8 uint64 = 42
    v1 = v1
    v2 = append(v2, ("foo bar" + "a"))
    var v9 string = "foo bar"
    for fi10 := 0; fi10 < 3; fi10++ {
        v1 -= (v1 & 65535)
        var v11 uint64 = (^(v8))
        _ = len(v2)
        _ = fi10
        _ = v11
    }
    var v12 float64 = ((0.0 - 1.5) / float64(1e10))
    if (4294967295 >= 732) {
        fmt.Println(v12, v3)
        fmt.Println(v7)
    }
    v2 = append(v2, v9)
    v8 -= 18446744073709551615
    v8 += (v8 % ((v8) | 1))
    var v13 string = v9
    switch 128 {
    case 42:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -9223372036854775808:
        fmt.Println(v13, v12)
    case 599:
        var v14 int64 = 42
        _ = v14
    }
    switch v9 {
    case "foo bar":
        fmt.Println(v9, v9)
    case "hello":
        v15 := GenS2{F0: 2147483647, F1: (true || (42 != 9223372036854775807))}
        _ = v15
    case "\n\t":
        var v16 int64 = 42
        _ = v16
    }
    v13 = ("\n\t" + v9)
    for fi17 := 0; fi17 < 13; fi17++ {
        v2 = append(v2, v9)
        fmt.Println(fi17, v1, v9)
        _ = fi17
    }
    v18 := GenS2{F0: 4294967295, F1: (false && true)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
