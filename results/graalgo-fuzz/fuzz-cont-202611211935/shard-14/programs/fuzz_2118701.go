package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint32 }
func genFn0() string {
    var v7 int = -204
    v8 := []int{}
    v9 := GenS1{F0: true, F1: 120}
    _ = v7
    _ = v8
    _ = v9
    return (("" + (("foo bar" + (("a" + "") + ("" + (("" + "a") + "a")))) + "\n\t")) + "\n\t")
}

func main() {
    var v1 int = -1
    v2 := []uint32{919}
    v3 := make(map[string]int)
    var v4 bool = false
    var v5 uint16 = 0
    var v6 string = "hello"
    if len(v2) > 0 {
        v2[0] = 42
    }
    var v10 float64 = 0.0
    v10 = v10
    var v11 int16 = 849
    v12 := GenS2{F0: 708, F1: -128}
    if (93 == 0) {
        v13 := GenS1{F0: ((!(v11 >= v11)) || v4), F1: -1}
        if (-413 != -640) {
            fmt.Println(v5, v4, v4)
            var v14 uint = 42
            fmt.Println(v4)
            _ = v14
        } else {
            v4 = false
            v15 := GenS3{F0: -104}
            if (190 > 0) {
                v16 := GenS3{F0: 127}
                var v17 uint8 = 251
                _ = len(v2)
                _ = v16
                _ = v17
            }
            _ = v15
        }
        var v18 int64 = 585
        _ = v13
        _ = v18
    }
    switch ((v6 + (v6 + ((("a" + "a") + (v6 + ((v6 + v6) + v6))) + v6))) + ((("" + ("hello" + ("\n\t" + ("\n\t" + v6)))) + v6) + (("\n\t" + (v6 + v6)) + v6))) {
    case "hello":
        var v19 int8 = -1
        _ = v19
    case "a":
        v6 = "foo bar"
    case "foo bar":
        v6 = v6
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        if (0.1 < 1e10) {
            var v21 uint8 = 0
            _ = v21
        }
        v22 := GenS4{F0: 279, F1: 0}
        _ = fi20
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
