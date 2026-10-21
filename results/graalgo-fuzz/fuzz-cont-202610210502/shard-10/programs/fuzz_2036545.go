package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int }

func main() {
    var v1 uint8 = 47
    var v2 uint16 = 65535
    var v3 int16 = -879
    var v4 int64 = 392
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]uint32)
    var v7 int = -2147483648
    var v8 bool = ((^(v4)) < v4)
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 uint16 = v2
        _ = fi9
        _ = v10
    }
    var v11 uint = 18446744073709551615
    var v12 int64 = v4
    if v8 {
        v8 = false
        var v13 int32 = -1
        _ = v13
    } else {
        v14 := GenS1{F0: (v3 % ((v3) | 1)), F1: 986, F2: -2147483648}
        _ = v14
    }
    var v15 int8 = -111
    var v16 int64 = v12
    if false {
        v11 += (^(v11))
        v8 = (false || v8)
        if true {
            v1 = v1
            v1 += 0
            v17 := []int{(v7 - (v7 >> ((v11) & 63))), v7, int(v4)}
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := GenS1{F0: 853, F1: (v2 / ((255) | 1)), F2: fi18}
            v20 := []int{v7, (^(v7))}
            _ = fi18
            _ = v19
            _ = v20
        }
        v21 := v6[2]
        var v22 string = ""
        _ = v21
        _ = v22
    }
    v3 = (^(v3))
    switch 4294967295 {
    case 698:
        v3 -= (v3 >> (((v11 * v11)) & 63))
    case 1:
        delete(v6, 2147483647)
    case 2147483647:
        var v23 int8 = (v15 & int8(v3))
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
