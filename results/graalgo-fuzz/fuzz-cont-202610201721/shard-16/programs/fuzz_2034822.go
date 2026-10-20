package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 string = "hello"
    v3 := make(map[string]uint32)
    for fi4 := 0; fi4 < 15; fi4++ {
        var v5 float64 = -2.25
        v2 = ("" + (((v2 + v2) + "") + (((("hello" + ((v2 + v2) + "hello")) + (v2 + v2)) + "\n\t") + v2)))
        _ = fi4
        _ = v5
    }
    v2 = ((("hello" + v2) + "hello") + v2)
    if true {
        var v6 int8 = -1
        _ = v6
    } else {
        for fi7 := 0; fi7 < 2; fi7++ {
            if true {
                if ((v2 + ((v2 + v2) + v2)) < (v2 + (("a" + ("\n\t" + v2)) + (((v2 + (((v2 + "a") + (("\n\t" + (v2 + "a")) + v2)) + v2)) + v2) + (v2 + ((((((v2 + (v2 + (((("foo bar" + ("foo bar" + "\n\t")) + "a") + (v2 + v2)) + (("foo bar" + "hello") + (v2 + "a"))))) + "") + v2) + ("hello" + v2)) + v2) + v2)))))) {
                    var v8 uint32 = 780
                    fmt.Println(v8)
                    fmt.Println(v2, fi7, v2)
                    _ = v8
                } else {
                    var v9 uint = 0
                    _ = v9
                }
                switch 127 {
                case 9223372036854775807:
                    v10 := GenS1{F0: 1}
                    _ = v10
                }
            }
            _ = fi7
        }
    }
    v11 := GenS2{F0: 4294967295, F1: -9223372036854775808}
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int = fi12
        v2 = "hello"
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        switch (fi14 | fi14) {
        case 4294967295:
            fmt.Println(fi14, v2)
        }
        _ = fi14
    }
    v2 = v2
    for fi15 := 0; fi15 < 16; fi15++ {
        v2 = ((v2 + ((("foo bar" + "foo bar") + v2) + (v2 + ""))) + v2)
        v2 = "hello"
        _ = fi15
    }
    v2 = ("\n\t" + v2)
    v16 := []int{}
    var v17 int = 2147483647
    var v18 string = v2
    v17 = (v17 | v17)
    v18 = (("\n\t" + "foo bar") + ((v2 + ("hello" + v18)) + (v2 + "")))
    if (2 != 255) {
        if (0 <= 0) {
            var v19 string = v18
            _ = v19
        } else {
            var v20 int64 = 42
            v2 = (((v18 + "foo bar") + ("" + "foo bar")) + v18)
            _ = v20
        }
        var v21 uint64 = 18446744073709551615
        _ = v21
    } else {
        var v22 int32 = 0
        _ = v22
    }
    var v23 uint64 = 203
    v24 := GenS2{F0: 79, F1: -9223372036854775808}
    v17 += v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
