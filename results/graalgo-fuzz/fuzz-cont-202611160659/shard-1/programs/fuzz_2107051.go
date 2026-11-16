package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 string }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 0
    var v3 uint = 4294967295
    var v4 uint64 = 0
    v1[-2147483648] = (796 <= v3)
    var v5 string = "\n\t"
    fmt.Println(v5, v5)
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(fi6, v4)
        _ = fi6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 int16 = -488
            switch (fi7 - (fi8 / (((fi7 + fi7)) | 1))) {
            case 9223372036854775807:
                fmt.Println(fi7, v5, fi8)
            default:
                v10 := GenS1{F0: 0, F1: true, F2: (v5 + (v5 + "hello"))}
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        for fi11 := 0; fi11 < 15; fi11++ {
            if false {
                v2 = v3
                delete(v1, fi7)
                var v12 int64 = 689
                _ = v12
            }
            var v13 string = ((v5 + v5) + v5)
            v3 = 65535
            _ = fi11
            _ = v13
        }
        _ = fi7
    }
    v14 := GenS1{F0: (v4 << (((v3 & (v3 / (((^(v2))) | 1)))) & 63)), F1: (v5 <= ""), F2: (v5 + ("hello" + (v5 + (v5 + (v5 + v5)))))}
    v15 := GenS2{F0: 2}
    for fi16 := 0; fi16 < 11; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := GenS2{F0: 208}
            switch v5 {
            case "foo bar":
                var v19 bool = ((-940 != 166) || (4294967295 <= 2147483647))
                _ = v19
            }
            _ = fi17
            _ = v18
        }
        v4 = (^(v4))
        _ = fi16
    }
    v3 = (v2 * v2)
    v20 := []int{}
    v20 = append(v20, 659)
    if (2147483647 != 61) {
        var v21 int16 = 79
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        var v22 uint32 = 2147483647
        _ = v21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
