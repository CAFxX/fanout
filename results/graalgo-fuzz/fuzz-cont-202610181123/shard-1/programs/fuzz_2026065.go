package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int8 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint }

func main() {
    v1 := []string{("a" + "foo bar"), "", ((("a" + "hello") + ("" + "a")) + "foo bar")}
    var v2 uint = 0
    var v3 uint64 = 213
    v4 := []int{-634}
    var v5 uint32 = 65535
    v2 = v2
    var v6 int8 = 100
    var v7 uint32 = v5
    switch -1 {
    case 0:
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 bool = (!false)
            if len(v1) > 0 {
                v1[0] = "a"
            }
            _ = fi8
            _ = v9
        }
    case 584:
        for fi10 := 0; fi10 < 18; fi10++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v3, v7)
            var v11 int = fi10
            _ = fi10
            _ = v11
        }
    case -303:
        var v12 float64 = (3.14159 - ((0.0 * 0.1) * 1e10))
        _ = v12
    default:
        switch "a" {
        case "\n\t":
            var v13 uint64 = v3
            _ = v13
        case "a":
            var v14 bool = true
            _ = v14
        }
    }
    v15 := GenS1{F0: 634, F1: "foo bar", F2: (^(v6))}
    v7 = v5
    var v16 int8 = -117
    var v17 int16 = -1
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS2{F0: (v3 - (v3 * uint64(v2))), F1: 2147483647}
        v7 += (^(v7))
        _ = fi18
        _ = v19
    }
    fmt.Println(v3, v2)
    if (true || (((v7 >> (((v2 >> (((v2 & 77)) & 63))) & 63)) != v7) || (v16 < (v16 >> ((65535) & 63))))) {
        var v20 uint16 = 2
        v21 := GenS3{F0: v7, F1: 669, F2: 42}
        v17 += -32768
        _ = v20
        _ = v21
    }
    var v22 string = "foo bar"
    fmt.Println(v7, v7, v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
