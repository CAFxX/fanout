package main

import (
    "fmt"
)


func main() {
    v1 := []int{32768, 944, 42, 4294967295}
    var v2 int = 128
    var v3 uint32 = 4294967295
    var v4 int = v2
    var v5 int8 = -63
    v4 += (v2 << ((160) & 63))
    var v6 bool = false
    fmt.Println(v3, v5)
    fmt.Println(v3, v5)
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 float64 = ((0.1 * (0.0 + 0.1)) - 0.0)
        var v9 int32 = -411
        v4 -= (v2 / ((v4) | 1))
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v4 = -1
        _ = fi10
    }
    if (1000 > 1000) {
        switch v4 {
        case 32767:
            for fi11 := 0; fi11 < 12; fi11++ {
                switch ((("hello" + ("a" + "hello")) + "hello") + "foo bar") {
                case "hello":
                    v4 += v2
                }
                _ = len(v1)
                _ = fi11
            }
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v2)
            _ = fi12
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
