package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }

func main() {
    var v1 int = 2147483647
    v2 := make(map[string]bool)
    v3 := []string{"", ""}
    for fi4 := 0; fi4 < 17; fi4++ {
        v1 = fi4
        v5 := GenS1{F0: 127, F1: 0}
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: 255, F1: 832}
            fmt.Println(v1, v1, fi4)
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    switch "foo bar" {
    case "foo bar":
        v1 -= (v1 + 9223372036854775807)
    case "\n\t":
        if (2 <= 2) {
            v1 = -446
        }
    }
    fmt.Println(v1)
    var v8 string = ("\n\t" + "foo bar")
    for fi9 := 0; fi9 < 13; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 int = 42
            v8 = ("hello" + v8)
            _ = fi10
            _ = v11
        }
        var v12 uint32 = 2147483648
        fmt.Println(v12)
        _ = fi9
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 int32 = -1
        switch v8 {
        case "hello":
            for fi15 := 0; fi15 < 19; fi15++ {
                if ((true && true) || false) {
                    var v16 float64 = (3.14159 - (((3.14159 - ((1.5 - (1.5 * float64(3.14159))) / 1.0)) * 1e10) + 1e10))
                    var v17 float32 = 0.1
                    _ = v16
                    _ = v17
                }
                v14 = -2147483648
                _ = fi15
            }
        default:
            fmt.Println(v1, v1)
        }
        v8 = ((v8 + ("a" + (v8 + "a"))) + v8)
        _ = fi13
        _ = v14
    }
    v18 := v2[((v8 + ("a" + v8)) + (v8 + (v8 + ((((v8 + v8) + "hello") + "\n\t") + v8))))]
    var v19 uint32 = 2147483647
    _ = len(v2)
    v20 := GenS1{F0: 174, F1: -699}
    fmt.Println(v19, v18)
    var v21 string = v8
    v2[v8] = (127 == 127)
    var v22 string = (v21 + ("a" + v8))
    v23 := GenS1{F0: 192, F1: -2147483648}
    var v24 int32 = 0
    if v18 {
        v21 = "foo bar"
        for fi25 := 0; fi25 < 1; fi25++ {
            if len(v3) > 0 {
                v3[0] = v8
            }
            switch ((v8 + v21) + (v22 + v8)) {
            case "\n\t":
                for fi26 := 0; fi26 < 1; fi26++ {
                    var v27 uint = 18446744073709551615
                    _ = fi26
                    _ = v27
                }
            }
            var v28 uint8 = 42
            _ = fi25
            _ = v28
        }
    }
    var v29 float32 = (1e10 / 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v8
}
