package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int; F1 uint8; F2 int }
type GenS3 struct { F0 int8; F1 string; F2 uint64 }

func main() {
    var v1 uint8 = 2
    var v2 float32 = 0.1
    var v3 uint32 = 4294967295
    var v4 uint64 = 0
    var v5 string = ((("\n\t" + "\n\t") + "") + ("a" + "hello"))
    var v6 uint = 136
    if (431 != 65535) {
        v7 := GenS1{F0: 101, F1: -1}
        switch 4294967295 {
        case 32768:
            v8 := GenS2{F0: -156, F1: v1, F2: -172}
            _ = v8
        }
        v9 := GenS3{F0: 42, F1: ((v5 + v5) + v5), F2: (v4 * (v4 | v4))}
        _ = v7
        _ = v9
    } else {
        var v10 string = v5
        fmt.Println(v6, v4, v5)
        v6 = v6
        _ = v10
    }
    var v11 int32 = 0
    var v12 bool = true
    var v13 uint64 = (v4 + uint64(v3))
    var v14 int = 32767
    if false {
        fmt.Println(v5, v11, v13)
    } else {
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint32 = (v3 ^ 65535)
            if (v12 || (!false)) {
                v12 = ((^(v6)) == 985)
            }
            _ = fi15
            _ = v16
        }
        v17 := []int{(v14 - (v14 % (((v14 + (v14 ^ 998))) | 1))), 4294967295, (v14 - 32768), int(v6)}
        _ = v17
    }
    v14 += (v14 | v14)
    v18 := GenS1{F0: 75, F1: 909}
    fmt.Println(v4, v3, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
