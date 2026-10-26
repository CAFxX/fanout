package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }

func main() {
    v1 := []string{"", (("foo bar" + "") + "hello")}
    var v2 uint16 = 588
    var v3 string = "hello"
    var v4 string = ("a" + "a")
    v5 := []int{4294967295, 2147483647}
    var v6 string = v4
    for fi7 := 0; fi7 < 16; fi7++ {
        switch (v6 + v4) {
        case "foo bar":
            if true {
                switch fi7 {
                case 1:
                    v3 = v4
                default:
                    v6 = "foo bar"
                }
            }
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v8 := GenS1{F0: -698, F1: v2}
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 51
    fmt.Println(v3, v9, v6)
    v6 = "hello"
    var v10 float64 = 3.14159
    if (-9223372036854775808 < -197) {
        for fi11 := 0; fi11 < 10; fi11++ {
            _ = len(v1)
            _ = fi11
        }
        var v12 bool = (4294967295 == 853)
        _ = v12
    } else {
        var v13 uint = 18446744073709551615
        var v14 int = -2147483648
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: -571, F1: v2}
    v16 := GenS1{F0: -32768, F1: (v2 ^ 687)}
    var v17 int64 = -9223372036854775808
    switch v4 {
    case "hello":
        v9 += v9
    default:
        var v18 uint64 = 1
        _ = v18
    }
    switch 9223372036854775807 {
    case 1:
        v17 += 1
    case 293:
        for fi19 := 0; fi19 < 10; fi19++ {
            v17 = v17
            _ = fi19
        }
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
