package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int64; F2 int64 }

func main() {
    var v1 int32 = -1
    v2 := []string{("" + ("a" + "a")), "", "", "hello"}
    var v3 uint64 = 787
    v4 := make(map[string]uint32)
    var v5 float32 = 3.14159
    var v6 uint = 586
    v5 = 0.1
    v7 := GenS1{F0: v3, F1: 2}
    v8 := GenS2{F0: 127}
    if false {
        var v9 int64 = 9223372036854775807
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 int64 = (v9 % (((^(v9))) | 1))
            v12 := v4[("\n\t" + "hello")]
            v4["a"] = 2147483648
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    } else {
        fmt.Println(v1)
    }
    v6 += 2
    var v13 string = "a"
    var v14 uint16 = 42
    fmt.Println(v5, v13)
    v15 := GenS3{F0: false, F1: 0, F2: -314}
    for fi16 := 0; fi16 < 10; fi16++ {
        v5 = float32(((float64(v5) + 1.5) - 0.1))
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
