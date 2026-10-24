package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 339
    var v3 int32 = 1
    v2 -= (v2 + 621)
    v4 := []int{2147483647}
    var v5 uint64 = 1
    v6 := GenS1{F0: 0}
    var v7 bool = ("a" > ("\n\t" + "\n\t"))
    v3 = (^(v3))
    fmt.Println(v3)
    switch "hello" {
    case "\n\t":
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := GenS1{F0: 1}
            _ = fi8
            _ = v9
        }
    default:
        var v10 int32 = 0
        _ = v10
    }
    switch 32767 {
    case -2147483648:
        for fi11 := 0; fi11 < 16; fi11++ {
            v2 = (v2 << ((4294967295) & 63))
            _ = fi11
        }
    case 9223372036854775807:
        var v12 float32 = 3.14159
        _ = v12
    default:
        v13 := GenS1{F0: 18446744073709551615}
        _ = v13
    }
    v7 = (v7 && true)
    var v14 string = ("" + ((("hello" + "a") + ("a" + "foo bar")) + "hello"))
    v4 = append(v4, 1)
    switch (((v14 + v14) + v14) + v14) {
    case "hello":
        fmt.Println(v3)
    case "":
        delete(v1, int(v3))
    }
    v3 = v3
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
