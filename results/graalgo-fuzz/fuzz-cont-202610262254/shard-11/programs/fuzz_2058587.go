package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    v1 := []string{"a", ((("foo bar" + "hello") + "a") + "foo bar"), "foo bar", ("hello" + ("\n\t" + (("foo bar" + (("a" + "\n\t") + "")) + "\n\t")))}
    var v2 uint = 1
    var v3 float64 = -2.25
    if len(v1) > 0 {
        v1[0] = "a"
    }
    v2 -= (v2 & 339)
    fmt.Println(v3, v3)
    for fi4 := 0; fi4 < 13; fi4++ {
        fmt.Println(v3, fi4)
        v3 -= v3
        _ = fi4
    }
    v5 := GenS1{F0: -64, F1: 2}
    var v6 bool = (v2 != v2)
    var v7 float32 = (3.14159 + 0.0)
    v2 = v2
    v3 -= float64(v7)
    v6 = v6
    v7 = v7
    v8 := GenS2{F0: 240, F1: false}
    v9 := GenS2{F0: 74, F1: ("foo bar" >= "a")}
    var v10 bool = v6
    v10 = v10
    if false {
        var v11 int8 = 0
        for fi12 := 0; fi12 < 19; fi12++ {
            fmt.Println(v2, v7)
            _ = fi12
        }
        if ((float64(3.14159) - (1.5 / ((0.1 + float64(3.14159)) * float64(v7)))) < 1e10) {
            var v13 int64 = -150
            fmt.Println(v3, v6)
            _ = v13
        } else {
            var v14 int64 = 480
            _ = v14
        }
        _ = v11
    } else {
        fmt.Println(v3, v2)
        v1 = append(v1, "\n\t")
        var v15 string = "a"
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 int16 = -32768
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        _ = fi16
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
