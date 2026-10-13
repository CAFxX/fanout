package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 string; F1 uint32 }

func main() {
    v1 := []string{"", ("" + "a"), ("" + (("" + "foo bar") + "a"))}
    var v2 bool = (32767 >= 212)
    var v3 bool = (42 <= 32767)
    v4 := make(map[string]int)
    v5 := make(map[int]bool)
    var v6 float32 = ((0.1 / 1.0) + 3.14159)
    var v7 string = (((("" + ((((("" + ((((((("hello" + "foo bar") + "a") + "") + (("foo bar" + "\n\t") + "\n\t")) + ("foo bar" + "hello")) + "hello") + ((("\n\t" + "hello") + ("foo bar" + "")) + ("a" + "a")))) + "a") + "\n\t") + ("foo bar" + "\n\t")) + "foo bar")) + "a") + "\n\t") + ("" + "\n\t"))
    v2 = true
    fmt.Println(v3, v2, v7)
    v7 = v7
    var v8 int64 = 0
    v2 = v3
    v6 = (v6 / 1.0)
    var v9 int64 = -966
    v8 = v9
    v3 = (1 <= 976)
    v10 := GenS1{F0: 254, F1: -2147483648}
    for fi11 := 0; fi11 < 15; fi11++ {
        switch fi11 {
        case 352:
            if len(v1) > 0 {
                v1[0] = v7
            }
        case 2147483647:
            _ = len(v4)
        case 1:
            if len(v1) > 0 {
                v1[0] = ((v7 + (v7 + v7)) + "")
            }
        }
        v2 = (1 <= 300)
        v1 = append(v1, (v7 + ((v7 + "\n\t") + (v7 + v7))))
        _ = fi11
    }
    delete(v5, -795)
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int = fi12
        switch (v13 - fi12) {
        case 2147483647:
            var v14 uint64 = 18446744073709551615
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint = 255
        var v17 int16 = 1
        _ = fi15
        _ = v16
        _ = v17
    }
    switch ("hello" + v7) {
    case "":
        switch 42 {
        case -419:
            var v18 uint32 = 395
            _ = v18
        }
    case "foo bar":
        if v2 {
            fmt.Println(v6, v2)
            v8 -= v9
        } else {
            if v2 {
                v19 := GenS2{F0: "foo bar", F1: 65535}
                _ = v19
            } else {
                v1 = append(v1, (v7 + "foo bar"))
            }
            delete(v5, 128)
        }
    case "hello":
        v2 = v3
    default:
        var v20 bool = (v6 == float32(-2.25))
        _ = v20
    }
    var v21 float64 = float64(v6)
    switch v7 {
    case "a":
        var v22 float32 = v6
        _ = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
