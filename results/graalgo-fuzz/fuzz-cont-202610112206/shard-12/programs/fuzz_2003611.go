package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 string }

func main() {
    var v1 int32 = -514
    var v2 int64 = -188
    v3 := make(map[string]uint32)
    var v4 uint64 = 319
    var v5 bool = (true || (42 >= 235))
    var v6 bool = v5
    if (20 < 1) {
        fmt.Println(v5, v6)
    }
    if (!(!(509 < 255))) {
        var v7 uint16 = 255
        _ = v7
    } else {
        v4 += 2
        fmt.Println(v6)
        for fi8 := 0; fi8 < 12; fi8++ {
            if v6 {
                v5 = ("a" < (("" + "hello") + ("foo bar" + ("" + ""))))
                v9 := []int{}
                var v10 int64 = v2
                _ = v10
                _ = v9
            }
            _ = fi8
        }
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := []int{(fi12 % ((fi11) | 1))}
            _ = len(v13)
            _ = fi12
            _ = v13
        }
        var v14 uint64 = (v4 % (((v4 * (v4 - (^(v4))))) | 1))
        if (!((fi11 + 1) < fi11)) {
            if false {
                for fi15 := 0; fi15 < 15; fi15++ {
                    var v16 int32 = v1
                    _ = fi15
                    _ = v16
                }
                v17 := []int{}
                _ = v17
            }
            v18 := GenS1{F0: int16(v1), F1: -9223372036854775808, F2: ("foo bar" + ("hello" + "\n\t"))}
            _ = v18
        }
        _ = fi11
        _ = v14
    }
    var v19 int32 = (v1 << ((255) & 63))
    v2 = -1
    var v20 bool = ((v19 + (v1 | v1)) != v19)
    var v21 int = 32768
    var v22 bool = (!v20)
    switch int(v2) {
    case -1:
        _ = len(v3)
    case -9223372036854775808:
        if v5 {
            v19 = v19
        }
    }
    v4 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
