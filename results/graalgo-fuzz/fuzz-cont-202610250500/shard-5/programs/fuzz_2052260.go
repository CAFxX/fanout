package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int }
type GenS2 struct { F0 uint8; F1 string; F2 string }
type GenS3 struct { F0 uint }

func main() {
    var v1 int16 = 1
    var v2 uint8 = 1
    v3 := make(map[string]int)
    var v4 uint8 = (v2 - (v2 - 42))
    v5 := []uint32{911, 0}
    v6 := make(map[int]uint32)
    var v7 int64 = 9223372036854775807
    var v8 float64 = 1.5
    for fi9 := 0; fi9 < 13; fi9++ {
        switch ("foo bar" + ((("foo bar" + "") + "\n\t") + ((("a" + "a") + "") + "foo bar"))) {
        case "\n\t":
            var v10 uint16 = 1
            _ = v10
        }
        v11 := GenS1{F0: "a", F1: 26, F2: fi9}
        _ = fi9
        _ = v11
    }
    v12 := v3[("hello" + "foo bar")]
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 string = ""
        v15 := GenS2{F0: (v4 >> ((42) & 63)), F1: ("\n\t" + v14), F2: v14}
        var v16 float32 = -2.25
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := GenS3{F0: 368}
    v18 := GenS3{F0: 1}
    var v19 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
