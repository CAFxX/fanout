package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int8; F1 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = (((0.0 / 1e10) / 3.14159) * 0.1)
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    var v5 int32 = -809
    var v6 float64 = v2
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 string = (("\n\t" + ("hello" + "foo bar")) + (("\n\t" + ((((("\n\t" + "") + ("\n\t" + (("hello" + "foo bar") + ((("foo bar" + "hello") + (("\n\t" + "hello") + "foo bar")) + ("hello" + ""))))) + ("foo bar" + ("\n\t" + ("" + (("" + "hello") + "foo bar"))))) + "") + (("hello" + "foo bar") + "foo bar"))) + ("\n\t" + "a")))
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: 42, F1: false}
            v11 := GenS1{F0: 18446744073709551615, F1: (false && (((false || false) && (!(3.14159 != 1e10))) && (42 <= 18446744073709551615)))}
            v12 := GenS1{F0: 4294967295, F1: true}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v5, v2)
    v2 += float64(1.5)
    var v13 uint64 = 4294967295
    if (!((22 >= 32) || true)) {
        v14 := GenS2{F0: 1, F1: -923}
        var v15 int8 = 100
        _ = v14
        _ = v15
    } else {
        var v16 string = "a"
        v17 := []int{1}
        _ = v16
        _ = v17
    }
    v18 := []int{}
    if (v2 != 1e10) {
        if false {
            fmt.Println(v5)
        }
        v6 = v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
