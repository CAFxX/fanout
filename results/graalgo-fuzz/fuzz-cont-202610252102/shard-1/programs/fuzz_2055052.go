package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }

func main() {
    var v1 float64 = ((0.1 + float64(1e10)) - 1e10)
    v2 := make(map[string]string)
    var v3 uint = 4294967295
    var v4 string = (("" + "a") + ("" + ""))
    var v5 bool = ((!(true || true)) || true)
    v4 = ("a" + "a")
    if true {
        switch "a" {
        case "foo bar":
            if (2 <= 1) {
                fmt.Println(v3, v4)
                fmt.Println(v3, v5, v4)
            }
        case "hello":
            v4 = v4
        case "":
            var v6 int16 = 42
            _ = v6
        }
    } else {
        for fi7 := 0; fi7 < 17; fi7++ {
            v1 = (float64(0.0) * -2.25)
            _ = fi7
        }
        v5 = v5
    }
    switch 73 {
    case 2147483647:
        if v5 {
            switch -835 {
            case -9223372036854775808:
                var v8 float64 = v1
                _ = v8
            }
        }
    case 9223372036854775807:
        v5 = (v4 <= "a")
    }
    var v9 bool = (v5 || true)
    v10 := GenS1{F0: 2, F1: -1}
    v9 = true
    if (1 == 189) {
        var v11 bool = (3.14159 != v1)
        _ = v11
    }
    v9 = v9
    var v12 float64 = (3.14159 + 0.1)
    for fi13 := 0; fi13 < 19; fi13++ {
        v12 += (0.1 / ((v1 * 0.0) + float64(3.14159)))
        v5 = (false || ((-2.25 * 1.5) == 1e10))
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 float64 = v12
            var v16 uint8 = 255
            var v17 string = "a"
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    fmt.Println(v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
