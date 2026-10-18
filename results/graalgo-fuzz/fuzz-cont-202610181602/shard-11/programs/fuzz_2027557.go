package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []string{"a"}
    var v2 uint16 = 255
    var v3 int = 2147483647
    var v4 int64 = -549
    var v5 int16 = 32767
    var v6 int16 = v5
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    v4 = -331
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(fi7, v6)
        if (!(852 > 888)) {
            v5 -= 496
        }
        if false {
            var v8 int64 = (-(v4))
            fmt.Println(fi7)
            v4 -= (v8 & (v8 / ((570) | 1)))
            _ = v8
        } else {
            v1 = append(v1, (("a" + (("" + "foo bar") + "\n\t")) + "hello"))
        }
        _ = fi7
    }
    v9 := GenS1{F0: v3}
    var v10 float64 = 1e10
    if true {
        v3 = (v3 ^ v3)
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 string = "foo bar"
            _ = fi11
            _ = v12
        }
        switch "foo bar" {
        case "a":
            v13 := GenS1{F0: 716}
            _ = v13
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "\n\t":
            fmt.Println(v3, v5, v6)
        }
    } else {
        var v14 string = ("a" + "a")
        var v15 uint = 265
        _ = v14
        _ = v15
    }
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
