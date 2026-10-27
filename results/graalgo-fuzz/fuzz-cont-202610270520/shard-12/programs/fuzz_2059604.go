package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint64) uint32 {
    v5 := []int{-799, 32767, 1}
    switch -1 {
    case -764:
        for fi6 := 0; fi6 < 1; fi6++ {
            if true {
                var v7 uint8 = 128
                _ = v7
            }
            _ = len(v5)
            _ = fi6
        }
    default:
        var v8 string = ""
        _ = v8
    }
    _ = p0
    _ = v5
    return 4294967295
}

func main() {
    var v1 int64 = 377
    var v2 int16 = 1
    var v3 bool = (32768 <= -626)
    var v4 uint64 = 2
    var v9 bool = (2 >= 69)
    var v10 int64 = (v1 - v1)
    v11 := GenS1{F0: v3}
    v12 := []int{-641, 9223372036854775807}
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 uint64 = v4
        var v15 string = "\n\t"
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint64 = v4
    v9 = v9
    v9 = v9
    if ((!false) && true) {
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 6; fi18++ {
                v19 := GenS1{F0: v9}
                _ = fi18
                _ = v19
            }
            switch "foo bar" {
            case "":
                v20 := GenS1{F0: v9}
                _ = v20
            }
            _ = fi17
        }
    }
    fmt.Println(v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
