package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int32 = -1
    var v2 uint8 = 42
    v3 := make(map[int]uint32)
    var v4 uint8 = (v2 ^ (v2 ^ (v2 & 1)))
    var v5 int32 = v1
    v4 = 0
    var v6 int64 = int64(v1)
    if true {
        v2 = (^(v4))
        v7 := []int{}
        _ = v7
    }
    v4 = v4
    v4 = (v4 - 1)
    if false {
        delete(v3, 53)
        for fi8 := 0; fi8 < 20; fi8++ {
            switch "hello" {
            case "foo bar":
                var v9 string = ("a" + "foo bar")
                _ = v9
            default:
                for fi10 := 0; fi10 < 8; fi10++ {
                    v11 := v3[fi8]
                    _ = fi10
                    _ = v11
                }
            }
            v1 = (v1 + (^(v5)))
            _ = fi8
        }
    }
    v12 := GenS1{F0: 1}
    var v13 uint16 = 42
    v13 = 755
    if (("a" < "hello") || (-654 < 42)) {
        var v14 bool = (-2.25 >= 1.5)
        v13 -= (v13 + v13)
        _ = v14
    } else {
        if true {
            var v15 uint64 = 1
            _ = v15
        }
        var v16 uint16 = (^(v13))
        var v17 bool = (v6 <= v6)
        _ = v16
        _ = v17
    }
    delete(v3, 2147483647)
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 uint = 255
        var v20 uint8 = v4
        if false {
            v21 := []int{int(v6), (fi18 ^ (fi18 >> (((v19 & (v19 | v19))) & 63)))}
            _ = v21
        } else {
            v22 := v3[4294967295]
            _ = v22
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    v13 = (v13 * v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
