package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    var v1 float32 = float32(3.14159)
    var v2 int64 = -9223372036854775808
    var v3 uint16 = 42
    switch -1 {
    case 1:
        var v4 uint64 = 1
        _ = v4
    }
    var v5 string = "a"
    v6 := []int{}
    if (!false) {
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: 767, F1: -908}
            v5 = v5
            var v9 uint64 = 100
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 14; fi10++ {
            _ = len(v6)
            switch (fi10 ^ -1) {
            case 32768:
                v5 = ((("hello" + "foo bar") + "a") + "a")
            }
            for fi11 := 0; fi11 < 19; fi11++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi11
            }
            _ = fi10
        }
        fmt.Println(v5)
        v2 = (v2 | (-(v2)))
    }
    var v12 uint8 = 0
    v1 += (v1 - v1)
    v1 = v1
    var v13 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
