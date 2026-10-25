package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int32; F1 int8 }

func main() {
    v1 := []string{"foo bar", "hello", ("a" + (("\n\t" + "a") + "a")), "hello"}
    var v2 uint16 = 1
    var v3 uint = 27
    v4 := make(map[string]string)
    var v5 int = -2147483648
    v6 := []string{"\n\t"}
    var v7 float32 = 3.14159
    v5 -= v5
    if true {
        _ = len(v6)
        if false {
            fmt.Println(v2, v5)
            v2 -= (v2 * v2)
            v8 := GenS1{F0: int8(v5), F1: v2, F2: 2}
            _ = v8
        } else {
            v9 := v4[("" + (((((((("hello" + ("" + ("foo bar" + "\n\t"))) + ("hello" + ("hello" + ("\n\t" + "a")))) + "\n\t") + "") + "foo bar") + ("" + ("a" + "a"))) + "\n\t") + ""))]
            v10 := GenS1{F0: 49, F1: v2, F2: 0}
            if len(v6) > 0 {
                v6[0] = v9
            }
            _ = v10
            _ = v9
        }
        if (!(!true)) {
            var v11 int64 = 0
            if (v11 >= v11) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v11)
                fmt.Println(v3, v11)
            }
            _ = v11
        } else {
            var v12 int64 = 1
            _ = len(v4)
            _ = v12
        }
    } else {
        var v13 int16 = -32768
        v4[""] = "a"
        var v14 float64 = 3.14159
        _ = v13
        _ = v14
    }
    v5 -= 37
    if (128 <= 42) {
        v15 := GenS2{F0: -1, F1: -72}
        _ = len(v1)
        _ = v15
    }
    var v16 uint16 = v2
    if false {
        v16 = (v2 - (v16 << ((v3) & 63)))
        fmt.Println(v16)
    }
    v6 = append(v6, ("" + "a"))
    fmt.Println(v7, v5, v5)
    var v17 uint32 = 2147483648
    switch "" {
    case "a":
        var v18 uint64 = 0
        _ = v18
    case "\n\t":
        var v19 uint32 = (v17 | v17)
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
