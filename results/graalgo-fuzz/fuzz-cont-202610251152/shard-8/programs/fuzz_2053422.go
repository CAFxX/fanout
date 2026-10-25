package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int }
func genFn0() string {
    var v7 uint32 = 1
    v8 := []int{32767}
    _ = v7
    _ = v8
    return ""
}

func main() {
    var v1 uint64 = 905
    var v2 int32 = 822
    var v3 int16 = 42
    v4 := make(map[int]string)
    var v5 uint32 = 2147483648
    var v6 bool = (true && (543 >= 2))
    v4[9223372036854775807] = ((("hello" + "foo bar") + "hello") + "")
    v9 := []int{128, 939, -9223372036854775808, 4294967295}
    var v10 uint16 = 65535
    if v6 {
        for fi11 := 0; fi11 < 6; fi11++ {
            v2 += (v2 + v2)
            _ = fi11
        }
        v5 += v5
    }
    v12 := GenS1{F0: 40, F1: (v10 >> ((255) & 63)), F2: 1}
    v1 -= (^(v1))
    v3 -= 42
    var v13 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
