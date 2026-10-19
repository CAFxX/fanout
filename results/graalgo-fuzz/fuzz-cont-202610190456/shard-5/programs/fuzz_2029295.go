package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{2, 2147483647}
    var v3 bool = (32767 != 32767)
    v4 := make(map[int]string)
    var v5 uint = 986
    v5 -= v5
    _ = len(v2)
    var v6 string = (("" + "") + "foo bar")
    for fi7 := 0; fi7 < 19; fi7++ {
        v3 = (42 != -1)
        v3 = (!false)
        _ = fi7
    }
    fmt.Println(v5)
    var v8 int16 = 1
    var v9 float32 = (1e10 - float32(1.5))
    var v10 bool = true
    if v10 {
        v11 := v1[1]
        for fi12 := 0; fi12 < 20; fi12++ {
            switch "\n\t" {
            case "\n\t":
                var v13 uint8 = 255
                _ = v13
            case "foo bar":
                _ = len(v2)
            case "hello":
                v14 := GenS1{F0: (fi12 + (fi12 / ((9223372036854775807) | 1)))}
                _ = v14
            }
            var v15 uint64 = 16
            if len(v2) > 0 {
                v2[0] = 0
            }
            _ = fi12
            _ = v15
        }
        v5 = v5
        _ = v11
    }
    v6 = "foo bar"
    switch 128 {
    case 127:
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        v17 := GenS1{F0: (fi16 ^ fi16)}
        v3 = (v3 && (42 < 1))
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 9223372036854775807}
    var v19 uint32 = 149
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
