package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint32 = 127
    var v2 float64 = -2.25
    v3 := []string{}
    v4 := make(map[int]string)
    var v5 int32 = 90
    var v6 int = 683
    var v7 int = (v6 >> ((2) & 63))
    switch -297 {
    case 32768:
        v7 += (v7 - v6)
    case -2147483648:
        v8 := GenS1{F0: (v5 << ((18446744073709551615) & 63))}
        _ = v8
    case -947:
        v1 = 65535
    }
    v4[(v6 - v6)] = "foo bar"
    fmt.Println(v7)
    v5 = (^(v5))
    fmt.Println(v6, v2)
    var v9 bool = true
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v9 = ((524 >= 462) && (!(false || (!(true && (v5 < (v5 * (v5 + v5))))))))
    var v10 float64 = 1e10
    v1 += 418
    if (591 < 665) {
        v11 := GenS1{F0: v5}
        _ = v11
    } else {
        v12 := GenS1{F0: 2147483647}
        _ = v12
    }
    delete(v4, 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
