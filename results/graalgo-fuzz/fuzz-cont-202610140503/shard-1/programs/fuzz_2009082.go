package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint; F1 int16 }

func main() {
    v1 := []string{"\n\t", "\n\t", ("" + (("\n\t" + ("foo bar" + "a")) + "hello"))}
    var v2 int16 = 42
    v3 := []string{("a" + "foo bar")}
    var v4 bool = false
    switch "\n\t" {
    case "\n\t":
        var v5 float32 = (0.0 / 1.0)
        _ = v5
    case "":
        var v6 uint16 = 32767
        _ = v6
    default:
        switch "\n\t" {
        case "":
            v4 = (!((2 <= 181) || (!true)))
        case "a":
            _ = len(v3)
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    var v7 int8 = 42
    v2 = v2
    fmt.Println(v4, v4, v7)
    if false {
        var v8 float32 = 0.1
        _ = v8
    } else {
        v4 = true
        if v4 {
            if ((v7 + (^(v7))) < v7) {
                if ((v2 * (v2 << ((225) & 63))) <= (v2 - v2)) {
                    v9 := GenS1{F0: 559, F1: v2}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v10 uint = 1
                    _ = v10
                    _ = v9
                } else {
                    for fi11 := 0; fi11 < 6; fi11++ {
                        v12 := GenS1{F0: 1, F1: (v2 >> ((294) & 63))}
                        v13 := GenS1{F0: 0, F1: (-(v2))}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
                switch (("foo bar" + ("a" + ("" + "\n\t"))) + (("hello" + "a") + ("a" + "hello"))) {
                case "a":
                    v1 = append(v1, "\n\t")
                }
                v14 := GenS2{F0: 0, F1: 777}
                _ = v14
            } else {
                switch ("foo bar" + ((("" + ("hello" + (("hello" + "foo bar") + "\n\t"))) + ("hello" + "a")) + "\n\t")) {
                case "foo bar":
                    var v15 uint = 678
                    _ = v15
                case "\n\t":
                    var v16 float64 = (-2.25 / 3.14159)
                    _ = v16
                case "":
                    fmt.Println(v7, v4, v2)
                }
            }
            fmt.Println(v7, v7, v2)
        }
        if len(v1) > 0 {
            v1[0] = ("hello" + "a")
        }
    }
    var v17 string = "a"
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
