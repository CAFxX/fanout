package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }

func main() {
    var v1 uint32 = 65535
    var v2 string = ""
    var v3 bool = (!(true && (42 == 230)))
    var v4 bool = (v3 && v3)
    var v5 int16 = 32767
    v6 := make(map[string]string)
    fmt.Println(v5)
    if true {
        switch 0 {
        case 281:
            for fi7 := 0; fi7 < 1; fi7++ {
                v1 = (v1 << ((65535) & 63))
                _ = fi7
            }
        case 267:
            v8 := GenS1{F0: "foo bar", F1: 54}
            _ = v8
        }
    }
    v5 = v5
    var v9 int64 = 173
    v10 := []int{703, 32767, -313}
    if (!v3) {
        for fi11 := 0; fi11 < 6; fi11++ {
            fmt.Println(v9)
            v4 = v3
            var v12 uint64 = 42
            _ = fi11
            _ = v12
        }
        var v13 int32 = -2147483648
        _ = v13
    } else {
        v2 = v2
    }
    v6[v2] = (((v2 + v2) + ("\n\t" + v2)) + v2)
    fmt.Println(v9, v3, v4)
    var v14 uint64 = uint64(v1)
    var v15 int8 = int8(v14)
    v16 := GenS1{F0: v2, F1: v15}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
