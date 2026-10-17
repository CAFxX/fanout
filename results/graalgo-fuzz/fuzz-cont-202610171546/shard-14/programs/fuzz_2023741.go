package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 bool }

func main() {
    v1 := make(map[string]int)
    var v2 float32 = ((-2.25 + float32((1.5 - 1e10))) / 1.0)
    var v3 float64 = 1.5
    v4 := []string{"", "\n\t", "foo bar", "\n\t"}
    v5 := []string{"hello", "foo bar"}
    v6 := make(map[string]int)
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := v1[(("hello" + "hello") + ("" + "hello"))]
        var v9 uint8 = 1
        v3 += v3
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS1{F0: ("hello" + (((("foo bar" + "") + "\n\t") + "\n\t") + "hello")), F1: 133, F2: (v2 == v2)}
        var v12 int32 = -1
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v2, v2)
    var v13 uint16 = 609
    if (!(42 >= -66)) {
        v13 += v13
        if len(v5) > 0 {
            v5[0] = ""
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v2 -= v2
    v13 = 255
    v14 := GenS1{F0: "\n\t", F1: 0, F2: true}
    if len(v5) > 0 {
        v5[0] = ("a" + "a")
    }
    delete(v1, ("" + ("" + "hello")))
    var v15 int = 42
    v3 = v3
    var v16 int64 = -1
    v17 := GenS1{F0: ("" + ("\n\t" + "\n\t")), F1: uint32(v13), F2: (511 <= v16)}
    v15 = v15
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch "hello" {
    case "\n\t":
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 bool = ((!(42 != 42)) || false)
            _ = fi18
            _ = v19
        }
    case "hello":
        fmt.Println(v3)
    case "a":
        var v20 uint8 = 70
        _ = v20
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v15)
        _ = fi21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
