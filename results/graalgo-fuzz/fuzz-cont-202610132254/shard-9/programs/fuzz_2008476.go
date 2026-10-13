package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []string{"foo bar", "foo bar", ("hello" + "a"), "\n\t"}
    var v2 uint8 = 121
    var v3 float32 = 1e10
    var v4 int64 = 1
    var v5 bool = true
    var v6 int64 = v4
    v3 = 0.1
    v1 = append(v1, "")
    v6 = v6
    fmt.Println(v2, v3, v5)
    fmt.Println(v2)
    if v5 {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint16 = 116
            for fi9 := 0; fi9 < 2; fi9++ {
                v6 = int64(v8)
                _ = fi9
            }
            for fi10 := 0; fi10 < 7; fi10++ {
                fmt.Println(v5)
                var v11 int = fi7
                v5 = (!true)
                _ = fi10
                _ = v11
            }
            _ = fi7
            _ = v8
        }
    } else {
        if v5 {
            v1 = append(v1, ("foo bar" + "hello"))
        }
        var v12 uint8 = v2
        v12 = v2
        _ = v12
    }
    v2 -= (v2 * v2)
    if len(v1) > 0 {
        v1[0] = ""
    }
    v4 += 747
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    switch 214 {
    case -1:
        switch -2147483648 {
        case 32767:
            _ = len(v1)
        case 526:
            if len(v1) > 0 {
                v1[0] = "hello"
            }
        default:
            v13 := GenS1{F0: 255}
            _ = v13
        }
    case 128:
        var v14 int8 = 36
        _ = v14
    case 9223372036854775807:
        v2 -= (^(v2))
    }
    if (444 != 42) {
        var v15 bool = true
        _ = v15
    }
    v3 = (v3 * v3)
    v16 := GenS1{F0: 603}
    if v5 {
        _ = len(v1)
        var v17 string = "\n\t"
        for fi18 := 0; fi18 < 20; fi18++ {
            v1 = append(v1, v17)
            _ = fi18
        }
        _ = v17
    } else {
        var v19 uint64 = 18446744073709551615
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
