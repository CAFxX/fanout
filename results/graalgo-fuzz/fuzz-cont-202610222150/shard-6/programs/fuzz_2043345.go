package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 float64 = (1e10 * -2.25)
    fmt.Println(v3)
    if false {
        v3 -= (v3 / float64(float32((1.5 + v3))))
    } else {
        var v4 uint8 = 1
        _ = v4
    }
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := GenS1{F0: 2, F1: 127}
        _ = fi5
        _ = v6
    }
    v7 := v1[457]
    var v8 uint = 2
    var v9 int64 = 1
    switch ("a" + "\n\t") {
    case "hello":
        for fi10 := 0; fi10 < 7; fi10++ {
            v9 -= v9
            _ = fi10
        }
    case "a":
        switch "" {
        case "foo bar":
            delete(v2, "foo bar")
        case "\n\t":
            v8 -= 2
        default:
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := GenS2{F0: 255, F1: 100}
                v7 += 2147483648
                v13 := GenS3{F0: 42, F1: 104, F2: 0}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    case "":
        var v14 string = "\n\t"
        _ = v14
    }
    v9 -= (v9 << ((v8) & 63))
    var v15 uint64 = 217
    for fi16 := 0; fi16 < 7; fi16++ {
        v9 -= 42
        var v17 int64 = int64(v15)
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        delete(v2, "")
        var v19 uint8 = 83
        _ = fi18
        _ = v19
    }
    v7 = v7
    v20 := GenS1{F0: (v7 / ((183) | 1)), F1: 0}
    fmt.Println(v3, v15, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
