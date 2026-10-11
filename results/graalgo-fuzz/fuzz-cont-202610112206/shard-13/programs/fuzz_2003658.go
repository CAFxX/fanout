package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int64 }
func genFn0(p0 bool, p1 int8) uint {
    v4 := []int{128}
    _ = len(v4)
    _ = p0
    _ = p1
    _ = v4
    return 2
}

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint = (v1 >> (((^(v1))) & 63))
        var v7 int8 = 100
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v1)
    if (4294967295 <= 2147483648) {
        switch "a" {
        case "foo bar":
            v1 += 4294967295
        case "":
            for fi8 := 0; fi8 < 13; fi8++ {
                v1 = 65535
                _ = fi8
            }
        default:
            if (false && true) {
                if true {
                    var v9 uint8 = 128
                    _ = v9
                }
            }
        }
        if false {
            fmt.Println(v1)
        }
        var v10 uint16 = 52
        _ = v10
    }
    v1 = v1
    v1 = 316
    var v11 uint8 = 196
    v11 = (^(v11))
    if (true && true) {
        if false {
            fmt.Println(v11, v11)
            switch 4294967295 {
            case -1:
                v1 += (v1 << ((1) & 63))
            case -760:
                var v12 uint8 = v11
                _ = v12
            }
        }
        v1 += uint(v11)
    }
    v13 := GenS1{F0: 373, F1: -9223372036854775808, F2: 1}
    v14 := v3[-9223372036854775808]
    v15 := GenS1{F0: 19, F1: 1, F2: 846}
    var v16 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
