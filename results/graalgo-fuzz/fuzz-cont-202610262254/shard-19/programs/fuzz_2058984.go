package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 65535
    v4 = 65535
    var v5 string = ("foo bar" + "a")
    var v6 int64 = int64(v4)
    v7 := v2[1]
    v6 = (v6 + v6)
    v6 -= -9223372036854775808
    v4 -= v4
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int8 = 100
        if true {
            fmt.Println(v4, v5)
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 string = "\n\t"
                v12 := GenS1{F0: 4294967295}
                v7 = (((v7 + ((v11 + v5) + v7)) + (v5 + (v5 + ("foo bar" + "hello")))) + "")
                _ = fi10
                _ = v11
                _ = v12
            }
            switch 42 {
            case -1:
                v13 := []int{(fi8 % ((fi8) | 1)), (fi8 & (fi8 | fi8))}
                _ = v13
            }
        }
        v6 -= (v6 % ((v6) | 1))
        _ = fi8
        _ = v9
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
