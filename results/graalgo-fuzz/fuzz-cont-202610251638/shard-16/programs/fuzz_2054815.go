package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 bool; F1 uint64; F2 uint }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []string{"\n\t", "", "", "\n\t"}
    v2 := make(map[int]string)
    var v3 float32 = (float32(3.14159) + (0.1 - 1e10))
    v4 := make(map[uint32]bool)
    var v5 uint64 = 372
    var v6 bool = (42 >= -32768)
    v7 := GenS1{F0: 1, F1: ("hello" + ("a" + "a"))}
    v6 = ((v3 * ((0.0 / (1e10 + v3)) - 0.1)) < -2.25)
    fmt.Println(v6, v6)
    v8 := GenS2{F0: false, F1: 18446744073709551615, F2: 18446744073709551615}
    if (true || (1 < -2147483648)) {
        if (!(true && false)) {
            if len(v1) > 0 {
                v1[0] = ""
            }
            var v9 int32 = -998
            _ = v9
        }
    }
    var v10 int8 = -123
    fmt.Println(v10)
    v11 := GenS3{F0: 9223372036854775807}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
