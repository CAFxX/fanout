package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint64 }
func genFn0(p0 int16) string {
    fmt.Println(p0)
    fmt.Println(p0)
    p0 += p0
    _ = p0
    return "foo bar"
}

func main() {
    var v1 string = ("\n\t" + "")
    var v2 uint32 = 4294967295
    var v3 int = 127
    v4 := []string{v1, v1, v1, "a"}
    var v5 uint = 65535
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(v1, v1, v5)
        _ = fi6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        v1 = (v1 + v1)
        var v8 uint16 = 42
        _ = fi7
        _ = v8
    }
    var v9 bool = false
    if len(v4) > 0 {
        v4[0] = (((v1 + (v1 + v1)) + ("hello" + "hello")) + "foo bar")
    }
    v9 = false
    fmt.Println(v3, v2, v2)
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS1{F0: fi10, F1: -1}
        switch "hello" {
        case "foo bar":
            if false {
                v1 = v1
            }
        case "":
            switch "a" {
            case "foo bar":
                _ = len(v4)
            case "":
                for fi12 := 0; fi12 < 19; fi12++ {
                    var v13 uint8 = 42
                    _ = fi12
                    _ = v13
                }
            }
        default:
            var v14 uint = v5
            _ = v14
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            fmt.Println(v1)
            var v16 uint32 = v2
            _ = fi15
            _ = v16
        }
        _ = fi10
        _ = v11
    }
    var v17 int8 = 42
    var v18 uint = 65535
    v19 := GenS2{F0: (v2 % ((v2) | 1)), F1: 2, F2: 309}
    switch v3 {
    case 128:
        var v20 int = 915
        _ = v20
    case 734:
        var v21 uint32 = (v2 << ((v5) & 63))
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
