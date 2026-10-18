package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint = 65535
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    v1 -= (v1 << ((v1) & 63))
    v1 += 65535
    if (0 <= 42) {
        v1 = (v1 >> (((^(v1))) & 63))
        fmt.Println(v1)
        switch 32767 {
        case -1:
            fmt.Println(v1)
        default:
            delete(v3, "")
        }
    } else {
        fmt.Println(v1)
        var v5 uint32 = 65535
        _ = v5
    }
    v1 = v1
    if (-9223372036854775808 >= -9223372036854775808) {
        v6 := GenS1{F0: 1}
        _ = v6
    }
    fmt.Println(v1)
    var v7 uint16 = 314
    if (42 < 4294967295) {
        v7 += 2
        switch "\n\t" {
        case "foo bar":
            var v8 uint16 = (^(v7))
            _ = v8
        case "":
            var v9 uint = v1
            _ = v9
        }
    }
    v1 = v1
    v10 := []int{}
    if true {
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 uint64 = uint64(v1)
            switch 32768 {
            case 2147483647:
                v12 = v12
            case -1:
                switch (fi11 << ((v1) & 63)) {
                case 9223372036854775807:
                    v13 := GenS1{F0: 1}
                    _ = v13
                }
            default:
                var v14 bool = (false || (true && true))
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        v7 -= v7
    }
    v1 += (v1 << ((1) & 63))
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
