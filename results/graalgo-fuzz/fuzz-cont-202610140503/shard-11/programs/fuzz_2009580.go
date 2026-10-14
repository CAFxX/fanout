package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 bool = true
    v2 := []string{"foo bar", (("" + ((("" + ("foo bar" + (("hello" + ("hello" + "a")) + "hello"))) + "foo bar") + ("hello" + "a"))) + "")}
    var v3 float32 = (1e10 - 0.1)
    v4 := []string{"a", (("foo bar" + "foo bar") + "a"), "\n\t", "\n\t"}
    var v5 uint64 = 461
    var v6 int16 = 36
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint16 = 0
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        _ = fi9
    }
    var v10 string = "foo bar"
    switch 128 {
    case 128:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        for fi11 := 0; fi11 < 13; fi11++ {
            v1 = ((2 <= uint8(v5)) || v1)
            var v12 int16 = 994
            _ = fi11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v5 -= v5
        _ = fi13
    }
    v14 := GenS1{F0: 1000}
    for fi15 := 0; fi15 < 7; fi15++ {
        fmt.Println(v6, v10, v1)
        _ = fi15
    }
    var v16 uint32 = 4294967295
    var v17 int8 = -1
    _ = len(v4)
    v5 = (v5 & v5)
    v16 -= v16
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
