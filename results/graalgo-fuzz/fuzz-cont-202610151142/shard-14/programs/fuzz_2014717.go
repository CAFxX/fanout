package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint32; F2 int }

func main() {
    var v1 int = 4294967295
    v2 := make(map[string]bool)
    var v3 int = v1
    var v4 string = ""
    var v5 string = v4
    v6 := []int{(v1 / (((v1 + v3)) | 1))}
    if (-120 != -51) {
        var v7 int64 = 9223372036854775807
        _ = v7
    }
    var v8 int8 = 127
    var v9 bool = false
    var v10 int8 = (v8 << ((701) & 63))
    v9 = (float64(-2.25) == 0.0)
    if (0.0 > 0.1) {
        fmt.Println(v10)
        v11 := GenS1{F0: v3, F1: ((v5 + v4) + ("a" + "")), F2: 18446744073709551615}
        switch -74 {
        case 42:
            v12 := v2[""]
            _ = v12
        }
        _ = v11
    } else {
        v13 := GenS2{F0: -1}
        v14 := GenS2{F0: 42}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        delete(v2, (v5 + (((("foo bar" + "\n\t") + ("\n\t" + v4)) + ((v4 + (v4 + v5)) + v5)) + v5)))
        v16 := GenS3{F0: -510, F1: 0, F2: fi15}
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
