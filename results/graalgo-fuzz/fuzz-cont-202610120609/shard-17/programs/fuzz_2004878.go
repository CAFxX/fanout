package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int = 32768
    var v2 int8 = 1
    var v3 int = (v1 * v1)
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 13; fi5++ {
        if (!true) {
            var v6 bool = (1.5 < (0.1 - float32(1.5)))
            _ = v6
        }
        var v7 uint16 = 942
        var v8 int = v3
        _ = fi5
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 4294967295}
    for fi10 := 0; fi10 < 2; fi10++ {
        v2 += (v2 | 100)
        for fi11 := 0; fi11 < 16; fi11++ {
            switch "hello" {
            case "foo bar":
                var v12 uint16 = 804
                _ = v12
            }
            _ = fi11
        }
        _ = fi10
    }
    v1 = 1
    v1 = 4294967295
    v1 += 2147483647
    if true {
        var v13 int64 = -390
        v1 = (v1 << ((255) & 63))
        v14 := GenS1{F0: 207}
        _ = v13
        _ = v14
    }
    var v15 uint16 = 65535
    fmt.Println(v3)
    v1 = 0
    v1 = v1
    var v16 uint8 = 127
    var v17 bool = ((v2 * v2) > 0)
    var v18 uint = 2
    if (v2 <= -1) {
        delete(v4, 569)
    }
    if ((v16 % (((v16 & v16)) | 1)) >= (v16 - 119)) {
        if (1.5 > 0.1) {
            var v19 int16 = 42
            var v20 float64 = -2.25
            _ = v19
            _ = v20
        } else {
            v18 = 18446744073709551615
        }
    } else {
        v21 := GenS2{F0: -537, F1: v16}
        var v22 float32 = 1.5
        v3 = v1
        _ = v21
        _ = v22
    }
    fmt.Println(v17, v2)
    v23 := []int{1}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
