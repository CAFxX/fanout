package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int32; F2 string }
func genFn0(p0 bool, p1 int) uint64 {
    fmt.Println(p1)
    p1 += p1
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := []int{p1, p1, (p1 << ((18446744073709551615) & 63)), 4294967295}
        if p0 {
            fmt.Println(p1)
            v6 := GenS1{F0: (fi4 / ((fi4) | 1))}
            _ = v6
        } else {
            switch "hello" {
            case "a":
                v7 := GenS1{F0: 229}
                _ = v7
            case "":
                v5 = append(v5, (fi4 / ((fi4) | 1)))
            }
            var v8 uint32 = 2147483648
            _ = v8
        }
        v9 := GenS2{F0: -97, F1: -2147483648, F2: ((("foo bar" + (("foo bar" + "foo bar") + "a")) + ("hello" + "hello")) + "hello")}
        _ = fi4
        _ = v5
        _ = v9
    }
    if false {
        var v10 float64 = (0.1 / 1.0)
        v11 := GenS1{F0: p1}
        fmt.Println(p0)
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            p1 += (p1 % ((128) | 1))
            _ = fi12
        }
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 int16, p1 int8, p2 bool) int16 {
    for fi13 := 0; fi13 < 19; fi13++ {
        p2 = p2
        _ = fi13
    }
    p0 = (p0 | p0)
    if (-1 <= 1) {
        if p2 {
            var v14 string = "foo bar"
            p1 = p1
            _ = v14
        }
    }
    v15 := GenS2{F0: p1, F1: 2147483647, F2: "\n\t"}
    switch -306 {
    case -695:
        p2 = (p2 || ("a" == ""))
    default:
        p2 = (((^(p1)) == (p1 << ((963) & 63))) && (1 <= -581))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 32767
}

func main() {
    var v1 int32 = -981
    v2 := []string{"\n\t", "a", "hello", ("" + "")}
    v3 := []string{"a", ("foo bar" + "hello"), "foo bar", "hello"}
    var v16 float32 = (-2.25 / -2.25)
    var v17 int8 = 42
    var v18 string = ""
    v3 = append(v3, ("foo bar" + "foo bar"))
    var v19 int = 127
    _ = len(v2)
    var v20 int64 = 1
    v19 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
