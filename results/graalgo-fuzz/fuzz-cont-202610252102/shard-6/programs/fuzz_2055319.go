package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int32; F1 int32; F2 int8 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint }

func main() {
    var v1 int16 = 0
    v2 := make(map[string]string)
    var v3 int32 = -1
    v4 := GenS1{F0: 0}
    v5 := GenS2{F0: "a", F1: 127, F2: 127}
    v6 := []int{0, int(v3), 9223372036854775807}
    v7 := v2["hello"]
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    v1 = -1
    v8 := GenS3{F0: v3, F1: (v3 | (v3 << ((42) & 63))), F2: -128}
    v7 = ("hello" + (((v7 + "\n\t") + "a") + ""))
    switch v7 {
    case "\n\t":
        v1 += 418
    case "foo bar":
        v6 = append(v6, 128)
    }
    v3 -= v3
    fmt.Println(v1, v1, v3)
    v9 := GenS1{F0: 1}
    if false {
        switch 127 {
        case 701:
            var v10 float32 = (3.14159 / 1.0)
            _ = v10
        }
    } else {
        fmt.Println(v1)
        fmt.Println(v7, v7)
    }
    v11 := GenS2{F0: v7, F1: 0, F2: 42}
    v12 := GenS4{F0: 4294967295, F1: 42, F2: 65535}
    v7 = v7
    v13 := GenS2{F0: "", F1: 103, F2: 9}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
