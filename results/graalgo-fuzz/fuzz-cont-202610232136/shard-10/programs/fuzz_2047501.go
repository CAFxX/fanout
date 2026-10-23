package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint32 }

func main() {
    var v1 int64 = 985
    v2 := make(map[string]int)
    var v3 float64 = 1e10
    var v4 uint = 54
    v5 := []string{"a", ("a" + "foo bar"), "foo bar", ("foo bar" + ("a" + (("" + (("" + "foo bar") + "a")) + "\n\t")))}
    var v6 float64 = v3
    var v7 uint8 = 2
    if false {
        var v8 int32 = 114
        v9 := GenS1{F0: 1000, F1: (v7 + v7), F2: 65535}
        _ = v8
        _ = v9
    } else {
        v3 -= float64(1.5)
        if (72 > -2147483648) {
            v3 = float64(-2.25)
            switch "hello" {
            case "foo bar":
                fmt.Println(v3, v6)
            case "a":
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
            }
        }
        fmt.Println(v4, v7)
    }
    v4 -= v4
    fmt.Println(v7, v1)
    if false {
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 int16 = 1000
            _ = fi10
            _ = v11
        }
        _ = len(v5)
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            if false {
                v13 := GenS1{F0: 1, F1: 2, F2: 2147483648}
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 float64 = ((v3 * (float64(-2.25) * (v3 - v3))) + v6)
                    switch "" {
                    case "":
                        switch ("\n\t" + ("\n\t" + "")) {
                        case "hello":
                            v16 := GenS1{F0: 280, F1: (v7 >> ((v4) & 63)), F2: 2147483647}
                            _ = v16
                        case "a":
                            v15 = float64(1e10)
                        case "foo bar":
                            v4 -= 18446744073709551615
                        }
                    default:
                        var v17 int8 = int8(v4)
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
            _ = fi12
        }
        v18 := GenS1{F0: 32767, F1: (v7 & 2), F2: 2147483648}
        _ = v18
    }
    var v19 uint32 = 841
    v6 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
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
}
