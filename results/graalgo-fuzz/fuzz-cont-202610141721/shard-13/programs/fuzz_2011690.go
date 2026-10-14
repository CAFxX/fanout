package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 3.14159
    var v3 int8 = -8
    v4 := []string{}
    var v5 uint = 748
    var v6 uint32 = 2147483647
    fmt.Println(v5)
    for fi7 := 0; fi7 < 18; fi7++ {
        if true {
            v8 := GenS1{F0: -9223372036854775808}
            _ = v8
        } else {
            v9 := GenS2{F0: (fi7 / (((fi7 & fi7)) | 1))}
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS3{F0: 32767, F1: (^(v3))}
    var v11 uint16 = 1
    switch int(v11) {
    case -9223372036854775808:
        v12 := GenS1{F0: 420}
        _ = v12
    case 578:
        v1[128] = "foo bar"
    case -213:
        if true {
            var v13 bool = true
            if (("\n\t" < "a") && v13) {
                fmt.Println(v6, v3)
            } else {
                if v13 {
                    v5 += (^(v5))
                    _ = len(v1)
                    v14 := v1[2147483647]
                    _ = v14
                } else {
                    delete(v1, -9223372036854775808)
                    var v15 int = -1
                    var v16 int32 = 2147483647
                    _ = v15
                    _ = v16
                }
                var v17 int = 193
                _ = v17
            }
            v18 := GenS2{F0: -2147483648}
            _ = v13
            _ = v18
        } else {
            v3 -= v3
        }
    }
    v11 = 42
    var v19 uint64 = 216
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
