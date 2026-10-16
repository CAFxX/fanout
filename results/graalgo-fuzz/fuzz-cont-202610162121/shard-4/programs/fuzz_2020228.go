package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []string{}
    var v2 int16 = -1
    v3 := make(map[int]bool)
    var v4 uint = 4294967295
    var v5 int64 = 9223372036854775807
    var v6 uint16 = 42
    var v7 int64 = v5
    v3[-2147483648] = false
    fmt.Println(v6)
    fmt.Println(v5, v2, v2)
    var v8 uint = (v4 << (((v4 * (v4 ^ (v4 << ((v4) & 63))))) & 63))
    for fi9 := 0; fi9 < 16; fi9++ {
        if (false && (!(!(!(0 > (v5 ^ v7)))))) {
            v4 += 436
        } else {
            var v10 int8 = 100
            if (v2 > (v2 - v2)) {
                for fi11 := 0; fi11 < 18; fi11++ {
                    delete(v3, int(v6))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v12 uint32 = 641
                    _ = fi11
                    _ = v12
                }
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                v13 := GenS1{F0: (^(v8))}
                _ = v13
            }
            _ = v10
        }
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            switch "\n\t" {
            case "\n\t":
                var v14 int32 = int32(v4)
                _ = v14
            default:
                v4 -= (v8 * v8)
            }
        case "a":
            v15 := GenS1{F0: v8}
            _ = v15
        }
        _ = fi9
    }
    v6 -= 1
    if false {
        v5 -= (v7 ^ (v5 - (v7 - v5)))
        var v16 bool = false
        _ = v16
    }
    fmt.Println(v6, v7, v7)
    var v17 uint = v4
    v2 += v2
    fmt.Println(v7, v7)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
