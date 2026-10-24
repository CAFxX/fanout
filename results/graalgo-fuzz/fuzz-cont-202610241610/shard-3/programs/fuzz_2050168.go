package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := []uint32{2147483648, 42}
    var v2 int16 = 1
    v3 := make(map[uint32]uint32)
    var v4 int64 = 515
    var v5 uint8 = 128
    var v6 bool = true
    for fi7 := 0; fi7 < 3; fi7++ {
        v2 += 0
        v5 = v5
        _ = fi7
    }
    var v8 string = ("a" + ("\n\t" + ""))
    var v9 uint = 18446744073709551615
    for fi10 := 0; fi10 < 12; fi10++ {
        v9 += (^(v9))
        switch fi10 {
        case 32767:
            if false {
                v5 += (v5 + v5)
            }
        case -171:
            v11 := GenS1{F0: 42}
            _ = v11
        }
        v5 = (v5 | v5)
        _ = fi10
    }
    var v12 uint = (v9 % ((v9) | 1))
    v13 := GenS1{F0: -112}
    switch 42 {
    case 4294967295:
        var v14 uint16 = 255
        _ = v14
    case 32768:
        switch 32767 {
        case -2147483648:
            v15 := v3[0]
            _ = v15
        case 9223372036854775807:
            if (2 >= v9) {
                var v16 int64 = v4
                v17 := GenS1{F0: 47}
                v18 := GenS1{F0: 1}
                _ = v16
                _ = v17
                _ = v18
            }
        }
    case 939:
        v3[618] = 2147483647
    }
    v3[4294967295] = 65535
    switch (v8 + v8) {
    case "\n\t":
        var v19 uint64 = 0
        _ = v19
    case "a":
        fmt.Println(v4, v9)
    }
    v4 += 42
    var v20 int64 = 0
    switch "foo bar" {
    case "a":
        fmt.Println(v6)
    case "":
        if v6 {
            var v21 uint8 = 1
            v22 := GenS2{F0: 374}
            _ = v21
            _ = v22
        }
    case "foo bar":
        v5 += 2
    default:
        fmt.Println(v20, v8, v8)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
