package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int16 }

func main() {
    var v1 uint32 = 2147483647
    var v2 int = 127
    v3 := []string{("foo bar" + ((("hello" + ("foo bar" + "")) + "foo bar") + "a")), "\n\t", "\n\t"}
    var v4 float32 = (3.14159 * (float32((0.1 / 1e10)) - (0.1 / (3.14159 + 1e10))))
    v5 := make(map[uint32]string)
    v5[v1] = "\n\t"
    switch v2 {
    case 1:
        switch (v2 / (((^(v2))) | 1)) {
        case 1:
            for fi6 := 0; fi6 < 12; fi6++ {
                if len(v3) > 0 {
                    v3[0] = (((((((("a" + "foo bar") + (("foo bar" + (("foo bar" + ("\n\t" + (("\n\t" + ("foo bar" + "")) + "\n\t"))) + "a")) + (("" + ("\n\t" + "a")) + ""))) + "\n\t") + "hello") + "foo bar") + "hello") + "") + "\n\t")
                }
                var v7 int16 = 0
                _ = fi6
                _ = v7
            }
        }
    }
    var v8 int = 127
    var v9 bool = (0.0 < (float64(v4) + (0.0 + -2.25)))
    fmt.Println(v2, v2, v1)
    v10 := GenS1{F0: -100, F1: 1, F2: -608}
    var v11 int32 = 1
    var v12 float64 = 3.14159
    switch (("" + "") + "") {
    case "\n\t":
        _ = len(v3)
    }
    switch (v8 ^ v2) {
    case 9223372036854775807:
        v13 := v5[(v1 & 65535)]
        _ = v13
    case 1:
        for fi14 := 0; fi14 < 12; fi14++ {
            v8 = 127
            var v15 string = "a"
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 uint16 = 0
                v5[(^(v1))] = v15
                _ = fi16
                _ = v17
            }
            _ = fi14
            _ = v15
        }
    case 174:
        v18 := GenS1{F0: v11, F1: -2147483648, F2: 816}
        _ = v18
    }
    fmt.Println(v9, v8, v1)
    var v19 string = (("foo bar" + ("foo bar" + (("foo bar" + "") + "a"))) + ("foo bar" + ("" + "a")))
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 uint8 = 2
        fmt.Println(v12)
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
