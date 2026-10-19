package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{2147483647, 720, 4294967295}
    var v3 uint32 = 527
    var v4 int = 1
    v4 = int(v3)
    v5 := GenS1{F0: (193 == -754), F1: (v3 & v3)}
    fmt.Println(v1, v4)
    v4 = v4
    var v6 int16 = 1
    v4 += (^(v4))
    v3 = 2147483647
    if (false && true) {
        for fi7 := 0; fi7 < 3; fi7++ {
            switch (^(fi7)) {
            case 644:
                v2 = append(v2, (^(v3)))
            case 158:
                var v8 float32 = 0.0
                _ = v8
            }
            if len(v2) > 0 {
                v2[0] = v3
            }
            fmt.Println(v3, v1, fi7)
            _ = fi7
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if false {
        var v9 bool = (false && false)
        _ = v9
    } else {
        var v10 int8 = 42
        _ = v10
    }
    switch -97 {
    case -9223372036854775808:
        v11 := GenS2{F0: 2}
        _ = v11
    }
    var v12 int32 = 0
    var v13 uint32 = (v3 | (v3 - 288))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
