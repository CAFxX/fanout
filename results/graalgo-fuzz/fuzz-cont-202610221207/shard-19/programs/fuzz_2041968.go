package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int32 }
func genFn0() uint16 {
    var v5 bool = (!false)
    fmt.Println(v5)
    v6 := GenS1{F0: 42}
    _ = v5
    _ = v6
    return 2
}
func genFn1() uint32 {
    var v7 int8 = 42
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: 2}
        _ = fi8
        _ = v9
    }
    var v10 string = ("foo bar" + "\n\t")
    _ = v10
    _ = v7
    return 540
}

func main() {
    var v1 int16 = -744
    v2 := []string{}
    var v3 bool = (false && false)
    var v4 uint64 = 1
    v4 = (v4 >> ((255) & 63))
    if v3 {
        switch -9223372036854775808 {
        case -9223372036854775808:
            v3 = false
        case -732:
            var v11 int8 = 100
            _ = v11
        }
        fmt.Println(v4, v3)
    }
    v4 = v4
    var v12 uint = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v4, v4)
    v1 = -1
    v4 -= (v4 * (v4 + 0))
    if (!true) {
        if (!((v12 <= (v12 << ((2) & 63))) || false)) {
            var v13 uint32 = 2147483647
            var v14 uint8 = 255
            _ = v13
            _ = v14
        } else {
            v1 -= v1
            fmt.Println(v12)
        }
        if (!v3) {
            v2 = append(v2, "")
            v15 := GenS2{F0: 0, F1: -617}
            _ = v15
        } else {
            if true {
                var v16 uint16 = 255
                _ = v16
            }
        }
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 int32 = -1
        _ = fi17
        _ = v18
    }
    var v19 uint64 = (v4 + v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
