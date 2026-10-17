package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int }

func main() {
    v1 := []string{"hello", ("a" + ("hello" + ("foo bar" + "foo bar"))), "a", ("\n\t" + ((("" + (("\n\t" + "hello") + "a")) + ("foo bar" + "")) + "\n\t"))}
    v2 := []string{}
    var v3 int32 = -1
    v4 := make(map[int]string)
    var v5 bool = (((float32((float64(-2.25) * 0.1)) + 0.0) <= 0.0) && (662 > 343))
    v5 = (979 >= 4294967295)
    fmt.Println(v5, v5)
    var v6 uint = 416
    v6 = 4294967295
    var v7 int8 = 104
    switch ("foo bar" + "hello") {
    case "\n\t":
        if v5 {
            v5 = v5
            v6 -= 255
        }
    case "foo bar":
        v4[411] = ""
    case "hello":
        v8 := GenS1{F0: 127, F1: v6}
        _ = v8
    }
    v9 := GenS1{F0: 127, F1: 1}
    _ = len(v4)
    v6 = v6
    if false {
        if (v5 || (v5 || (!v5))) {
            v3 = v3
        }
        var v10 int = 127
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS2{F0: 32767}
        var v13 uint16 = 1
        fmt.Println(v7)
        _ = fi11
        _ = v12
        _ = v13
    }
    switch 496 {
    case 32767:
        fmt.Println(v3, v6)
    }
    var v14 int8 = v7
    v14 = v14
    for fi15 := 0; fi15 < 11; fi15++ {
        fmt.Println(v3, fi15)
        _ = fi15
    }
    switch 50 {
    case -1:
        var v16 uint32 = 42
        _ = v16
    }
    v17 := v4[4294967295]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
