package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int }
func genFn0() int32 {
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: true, F1: 1}
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: true, F1: 32767}
        _ = fi7
        _ = v10
    }
    v11 := GenS1{F0: false, F1: 785}
    v12 := []int{}
    var v13 bool = true
    if false {
        var v14 int8 = 127
        v13 = v13
        _ = v14
    }
    _ = v11
    _ = v12
    _ = v13
    return -438
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[int]string)
    v3 := []uint32{}
    var v4 int = -1
    var v5 uint16 = 65535
    var v6 uint = 1
    var v15 int64 = -1
    v16 := GenS2{F0: 42}
    var v17 bool = (false || true)
    _ = len(v3)
    v5 = 849
    switch "foo bar" {
    case "foo bar":
        v18 := GenS1{F0: v17, F1: v5}
        _ = v18
    }
    fmt.Println(v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        if ((v5 | (^(v5))) > v5) {
            v20 := GenS1{F0: true, F1: v5}
            delete(v2, v4)
            var v21 bool = (2147483647 > (v4 << (((v6 | (v6 & v6))) & 63)))
            _ = v20
            _ = v21
        } else {
            _ = len(v2)
        }
        var v22 int64 = v15
        _ = fi19
        _ = v22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 uint64 = 1
        switch "a" {
        case "a":
            var v25 uint32 = v1
            _ = v25
        default:
            delete(v2, (fi23 & (v4 + 42)))
        }
        _ = fi23
        _ = v24
    }
    v4 = int(v1)
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 uint32 = v1
        _ = fi26
        _ = v27
    }
    var v28 uint8 = uint8(v1)
    fmt.Println(v15, v6)
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
