package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint8; F2 int64 }

func main() {
    v1 := []string{((("foo bar" + "foo bar") + ("foo bar" + "")) + ("a" + "hello")), "foo bar", ("foo bar" + "a"), ""}
    var v2 string = "a"
    v3 := []string{"a", v2, v2}
    var v4 uint8 = 1
    switch 4294967295 {
    case 9223372036854775807:
        v2 = ""
    case -53:
        fmt.Println(v2, v2)
    default:
        v4 += v4
    }
    var v5 string = "hello"
    if len(v3) > 0 {
        v3[0] = v2
    }
    var v6 int = -343
    v2 = "hello"
    var v7 uint8 = (v4 / ((42) | 1))
    v4 -= (v7 + v4)
    v5 = v2
    v8 := GenS1{F0: 1}
    switch "\n\t" {
    case "":
        if (1 > 42) {
            if false {
                _ = len(v3)
                _ = len(v1)
            } else {
                if false {
                    if (670 != 4294967295) {
                        v9 := GenS2{F0: true, F1: v4, F2: -9223372036854775808}
                        var v10 uint16 = 194
                        _ = v10
                        _ = v9
                    } else {
                        var v11 uint8 = 2
                        for fi12 := 0; fi12 < 5; fi12++ {
                            v13 := GenS2{F0: false, F1: 127, F2: 1}
                            var v14 uint = 18446744073709551615
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        _ = v11
                    }
                }
                v7 = (v4 * (v7 >> ((65535) & 63)))
            }
            switch (v2 + v5) {
            case "\n\t":
                var v15 uint64 = 18446744073709551615
                _ = v15
            case "":
                var v16 uint16 = 241
                _ = v16
            }
            var v17 string = v2
            _ = v17
        } else {
            var v18 float64 = 1e10
            _ = v18
        }
    }
    if (186 != -112) {
        fmt.Println(v2, v4, v5)
    }
    var v19 string = ("a" + ("hello" + v5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
