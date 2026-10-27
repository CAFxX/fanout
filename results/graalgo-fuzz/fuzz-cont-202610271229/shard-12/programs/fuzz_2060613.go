package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint8 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int8 = 100
    var v2 uint = 4294967295
    var v3 uint64 = 300
    var v4 uint32 = uint32(v2)
    v5 := make(map[int]bool)
    fmt.Println(v2, v3, v1)
    v5[446] = false
    v6 := []int{32, 32768}
    v1 = 112
    if false {
        var v7 string = "\n\t"
        _ = v7
    }
    v6 = append(v6, 593)
    v1 = v1
    var v8 int = int(v4)
    if ((3.14159 == (0.0 / float64(1.5))) || (1000 > 0)) {
        switch ("hello" + "a") {
        case "\n\t":
            v8 = v8
        }
    } else {
        switch v8 {
        case -851:
            fmt.Println(v4)
        default:
            if len(v6) > 0 {
                v6[0] = v8
            }
        }
        v8 = v8
    }
    v1 = v1
    for fi9 := 0; fi9 < 18; fi9++ {
        v2 = (v2 << (((v2 + 62)) & 63))
        v10 := GenS1{F0: (v2 * 722), F1: v2, F2: 117}
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 773}
    var v12 uint16 = 0
    if false {
        var v13 string = "\n\t"
        var v14 uint = v2
        fmt.Println(v13, v8)
        _ = v13
        _ = v14
    } else {
        if true {
            var v15 int64 = 9223372036854775807
            _ = v15
        } else {
            var v16 uint16 = (v12 % (((v12 * (v12 | v12))) | 1))
            _ = v16
        }
        if false {
            v3 += (v3 - uint64(v2))
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            if true {
                var v18 int64 = -1
                if true {
                    var v19 int = 9223372036854775807
                    var v20 uint8 = 0
                    v18 = (v18 % ((v18) | 1))
                    _ = v19
                    _ = v20
                }
                _ = v18
            }
            if len(v6) > 0 {
                v6[0] = v8
            }
            _ = fi17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
