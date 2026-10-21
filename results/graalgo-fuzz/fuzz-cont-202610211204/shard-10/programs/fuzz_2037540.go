package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint32, p1 int16) bool {
    if (546 >= 705) {
        fmt.Println(p1, p0)
        p1 = -1
        var v7 uint8 = 127
        _ = v7
    }
    if (42 > 42) {
        var v8 uint8 = 0
        _ = v8
    }
    _ = p0
    _ = p1
    return (!(true || (!true)))
}

func main() {
    var v1 float32 = float32(1.5)
    var v2 uint64 = 862
    v3 := make(map[uint32]string)
    var v4 int8 = -28
    var v5 int = 9223372036854775807
    var v6 bool = (255 <= 2)
    switch "\n\t" {
    case "a":
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 uint8 = 1
            var v11 float32 = 1.5
            v11 = v1
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        var v12 int64 = 42
        _ = v12
    }
    var v13 bool = (((v1 + ((v1 + (v1 - float32(-2.25))) / (v1 - float32(0.0)))) >= 1e10) && (2147483647 > 1))
    v14 := GenS1{F0: -85}
    delete(v3, 4294967295)
    var v15 uint16 = 65535
    fmt.Println(v2, v13, v5)
    if false {
        switch "" {
        case "foo bar":
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 int64 = 1
                _ = fi16
                _ = v17
            }
        }
    } else {
        v18 := GenS1{F0: 1}
        _ = v18
    }
    var v19 uint32 = 2
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
