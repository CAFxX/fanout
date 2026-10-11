package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int = 121
    var v2 int64 = 1
    v3 := make(map[string]string)
    switch "" {
    case "hello":
        var v4 uint16 = 775
        _ = v4
    }
    fmt.Println(v1)
    var v5 int64 = -598
    v6 := GenS1{F0: 65535}
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint64 = 42
        _ = fi7
        _ = v8
    }
    switch "hello" {
    case "hello":
        fmt.Println(v1, v5, v5)
    case "":
        fmt.Println(v5)
    case "foo bar":
        fmt.Println(v5, v1)
    }
    v9 := []int{128, (v1 | 336), -9223372036854775808}
    if false {
        var v10 int64 = v5
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        var v11 uint64 = 1
        _ = v10
        _ = v11
    }
    if (269 <= v5) {
        v5 = v5
    }
    if ((-113 >= 1) || true) {
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int = (-(v1))
            _ = fi12
            _ = v13
        }
        switch v1 {
        case 788:
            v1 -= v1
        case 821:
            if false {
                switch (^(v1)) {
                case 2147483647:
                    delete(v3, ("" + ("" + "hello")))
                case -2147483648:
                    v14 := GenS1{F0: 458}
                    _ = v14
                case 12:
                    v15 := GenS1{F0: 654}
                    _ = v15
                }
                var v16 float64 = ((-2.25 - 3.14159) / (-2.25 / 0.1))
                var v17 int64 = (^(v2))
                _ = v16
                _ = v17
            }
        case -262:
            var v18 string = (("\n\t" + (((("hello" + "") + "") + ("\n\t" + ("hello" + "\n\t"))) + ("hello" + (((("\n\t" + ("foo bar" + (("\n\t" + "") + ""))) + "hello") + "foo bar") + "hello")))) + "\n\t")
            _ = v18
        default:
            if true {
                fmt.Println(v1, v5, v1)
            }
        }
        if len(v9) > 0 {
            v9[0] = (v1 / ((v1) | 1))
        }
    } else {
        fmt.Println(v2, v5)
    }
    var v19 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v9
}
