package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint32; F2 uint32 }
func genFn0(p0 uint16) uint8 {
    var v4 int64 = 9223372036854775807
    v4 = (v4 % (((v4 >> ((18446744073709551615) & 63))) | 1))
    v5 := GenS1{F0: true}
    v6 := GenS2{F0: 59, F1: 89, F2: 371}
    fmt.Println(v4, p0)
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]int)
    var v3 int = 65
    fmt.Println(v1, v3)
    if true {
        fmt.Println(v3, v1)
        if (false && ((v3 / ((v3) | 1)) >= v3)) {
            var v7 int16 = 695
            _ = v7
        }
    } else {
        delete(v2, v3)
    }
    var v8 string = ("" + "\n\t")
    _ = len(v2)
    var v9 int = -568
    var v10 uint16 = 658
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS1{F0: (-2147483648 != 42)}
        _ = fi11
        _ = v12
    }
    var v13 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
