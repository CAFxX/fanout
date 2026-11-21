package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := []string{""}
    v2 := make(map[int]bool)
    var v3 string = "a"
    var v4 int16 = -897
    var v5 uint8 = 154
    fmt.Println(v5, v4, v5)
    if false {
        for fi6 := 0; fi6 < 2; fi6++ {
            for fi7 := 0; fi7 < 20; fi7++ {
                v3 = "\n\t"
                _ = fi7
            }
            _ = fi6
        }
        v5 += (v5 << ((42) & 63))
    }
    v5 = (^(v5))
    delete(v2, 127)
    v3 = "\n\t"
    for fi8 := 0; fi8 < 10; fi8++ {
        delete(v2, 128)
        _ = fi8
    }
    if (-1 < -9223372036854775808) {
        if true {
            var v9 uint32 = 436
            v3 = (v3 + (v3 + ""))
            _ = v9
        } else {
            fmt.Println(v4)
        }
        var v10 int16 = (v4 % ((v4) | 1))
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int = fi11
            if (0 > 932) {
                var v13 uint16 = 254
                if (!false) {
                    fmt.Println(v13)
                    v5 = (v5 << ((4294967295) & 63))
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v13
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        var v14 uint32 = 4294967295
        _ = v14
    }
    v1 = append(v1, v3)
    v5 -= (v5 << ((42) & 63))
    if false {
        switch v3 {
        case "\n\t":
            var v15 uint16 = 0
            _ = v15
        case "foo bar":
            v16 := GenS1{F0: 0, F1: 750}
            _ = v16
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
