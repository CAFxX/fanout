package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    var v3 int8 = -1
    var v4 int32 = -603
    var v5 uint64 = 0
    v6 := make(map[uint32]int)
    v5 = 42
    var v7 float32 = (0.1 / float32(-2.25))
    for fi8 := 0; fi8 < 17; fi8++ {
        switch -2147483648 {
        case 127:
            if false {
                var v9 string = ((("a" + "") + "\n\t") + "a")
                v7 = 0.0
                var v10 uint8 = 0
                _ = v10
                _ = v9
            }
        default:
            var v11 int32 = (v4 | v4)
            _ = v11
        }
        _ = fi8
    }
    var v12 uint64 = (v5 & 4294967295)
    _ = len(v6)
    var v13 int64 = -9223372036854775808
    for fi14 := 0; fi14 < 11; fi14++ {
        v13 = (v13 << ((996) & 63))
        _ = fi14
    }
    v13 = 9223372036854775807
    if (!(26 < 127)) {
        var v15 int8 = v3
        _ = v15
    }
    v16 := []int{0, 822, 9223372036854775807}
    v3 -= (v3 / (((v3 + v3)) | 1))
    var v17 uint64 = v12
    for fi18 := 0; fi18 < 8; fi18++ {
        if (!((fi18 | 42) <= fi18)) {
            v19 := GenS1{F0: 42, F1: (v13 - -9223372036854775808), F2: 42}
            v1[(("hello" + ((("foo bar" + "foo bar") + "a") + "a")) + "hello")] = ""
            _ = v19
        }
        var v20 int32 = (v4 >> ((738) & 63))
        var v21 int32 = 2147483647
        _ = fi18
        _ = v20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
