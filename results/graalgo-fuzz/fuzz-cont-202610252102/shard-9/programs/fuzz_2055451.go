package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 string }

func main() {
    v1 := []string{"hello", "foo bar", ("a" + "a"), (((("hello" + ("foo bar" + ("\n\t" + "hello"))) + "foo bar") + "hello") + "\n\t")}
    var v2 bool = false
    v3 := make(map[string]bool)
    var v4 int64 = 1
    v5 := []int{0, 266}
    v6 := make(map[string]bool)
    v4 -= (v4 | v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int16 = -32768
        v9 := GenS1{F0: 851, F1: 0, F2: ("a" + "")}
        switch "a" {
        case "\n\t":
            v4 += -917
        case "":
            v4 += (v4 << ((255) & 63))
        case "hello":
            var v10 bool = v2
            _ = v10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if (2 >= 4294967295) {
        _ = len(v5)
        v4 = v4
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 uint8 = 2
            switch "\n\t" {
            case "a":
                var v13 uint8 = (^(v12))
                _ = v13
            default:
                v12 += (v12 << ((255) & 63))
            }
            for fi14 := 0; fi14 < 8; fi14++ {
                v15 := GenS1{F0: uint16(v12), F1: (^(v12)), F2: "a"}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi14
                _ = v15
            }
            _ = fi11
            _ = v12
        }
    } else {
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := v3[("\n\t" + "a")]
            _ = fi16
            _ = v17
        }
        var v18 string = ("" + ("\n\t" + "a"))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v18
    }
    var v19 uint = 1
    var v20 int = 0
    var v21 int = v20
    _ = len(v3)
    var v22 float64 = float64(-2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
