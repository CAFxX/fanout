package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 int; F2 uint64 }

func main() {
    var v1 int64 = 42
    var v2 int64 = -1
    var v3 float64 = ((3.14159 / 3.14159) / ((1.5 + 0.0) / (1e10 / (1e10 - 3.14159))))
    var v4 float64 = 1.5
    v5 := make(map[string]string)
    var v6 uint8 = 42
    v7 := GenS1{F0: (v2 / (((v2 >> ((65535) & 63))) | 1)), F1: (v6 % (((v6 & (v6 >> ((0) & 63)))) | 1)), F2: 656}
    var v8 bool = false
    delete(v5, "")
    for fi9 := 0; fi9 < 1; fi9++ {
        v1 += v2
        var v10 int32 = 550
        _ = fi9
        _ = v10
    }
    v1 = v2
    v3 -= 0.0
    var v11 uint32 = 65535
    fmt.Println(v3, v4)
    if false {
        if true {
            v2 += 373
        } else {
            v12 := GenS2{F0: "\n\t", F1: (v6 / ((uint8(v11)) | 1))}
            _ = v12
        }
        var v13 float32 = 3.14159
        if true {
            v14 := GenS3{F0: 0}
            v11 = 426
            _ = v14
        }
        _ = v13
    } else {
        var v15 int16 = 42
        switch "\n\t" {
        case "":
            switch ("foo bar" + "\n\t") {
            case "a":
                v4 = 0.0
            }
        case "a":
            var v16 int16 = v15
            _ = v16
        default:
            v17 := GenS4{F0: 0, F1: 32767, F2: 369}
            _ = v17
        }
        if v8 {
            delete(v5, "\n\t")
            if (v6 > (^(v6))) {
                v1 = v2
            }
            fmt.Println(v11)
        } else {
            var v18 int8 = 31
            _ = v18
        }
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
