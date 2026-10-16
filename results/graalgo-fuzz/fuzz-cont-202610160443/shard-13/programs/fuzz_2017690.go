package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 255
    var v3 float32 = 1.5
    var v4 int16 = -1
    var v5 int16 = v4
    v6 := []int{39, 9223372036854775807}
    var v7 uint = v2
    v8 := v1["\n\t"]
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v2 = v7
    if ((1e10 < 0.0) && (v2 > 0)) {
        var v9 bool = (true && ("\n\t" < "\n\t"))
        switch "foo bar" {
        case "\n\t":
            v8 = v8
        }
        _ = v9
    } else {
        var v10 uint16 = 1
        _ = v10
    }
    if true {
        fmt.Println(v2)
        v3 = float32(0.1)
        var v11 bool = false
        _ = v11
    } else {
        v12 := GenS1{F0: 0, F1: 243, F2: (v7 >> ((v7) & 63))}
        var v13 int16 = v4
        _ = v12
        _ = v13
    }
    if (!true) {
        var v14 uint32 = 65535
        v15 := GenS1{F0: 2, F1: 0, F2: 42}
        _ = v14
        _ = v15
    }
    var v16 uint8 = 15
    v5 -= 1
    if len(v6) > 0 {
        v6[0] = -591
    }
    v7 = v7
    v16 = 1
    v17 := GenS1{F0: 4294967295, F1: (v16 / (((v16 - v16)) | 1)), F2: (v7 << ((v7) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
