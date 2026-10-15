package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int16 }
type GenS2 struct { F0 int32; F1 uint16; F2 string }

func main() {
    v1 := []int{128, 32767, 2147483647}
    var v2 uint32 = 42
    v3 := make(map[uint32]int)
    if false {
        if false {
            for fi4 := 0; fi4 < 8; fi4++ {
                v2 = v2
                _ = fi4
            }
        }
    } else {
        v5 := GenS1{F0: 242, F1: 687, F2: -1}
        v2 = (v2 + (v2 << ((2) & 63)))
        _ = v5
    }
    var v6 int8 = -119
    var v7 uint16 = 1
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint32 = 559
        if len(v1) > 0 {
            v1[0] = (fi8 & 2147483647)
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 float64 = 3.14159
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    var v12 uint = 614
    for fi13 := 0; fi13 < 13; fi13++ {
        switch int(v6) {
        case 878:
            v14 := GenS1{F0: 0, F1: 2, F2: 42}
            _ = v14
        }
        fmt.Println(v12, v12)
        _ = fi13
    }
    _ = len(v3)
    v2 += (v2 << ((v12) & 63))
    switch -248 {
    case 2147483647:
        for fi15 := 0; fi15 < 16; fi15++ {
            v12 = 18446744073709551615
            switch fi15 {
            case 1:
                v7 = (v7 ^ 542)
            case 0:
                if false {
                    var v16 uint16 = v7
                    _ = v16
                }
            case -150:
                if (-500 > -1) {
                    v7 = (v7 << ((42) & 63))
                } else {
                    if (false && (128 == 1)) {
                        v17 := GenS1{F0: 1, F1: 795, F2: -1}
                        _ = v17
                    }
                    v7 = (v7 % (((^(v7))) | 1))
                }
            }
            switch 9223372036854775807 {
            case 9223372036854775807:
                v12 += 4294967295
            case 4294967295:
                var v18 uint16 = v7
                _ = v18
            case -2147483648:
                v19 := GenS1{F0: -9223372036854775808, F1: 0, F2: 795}
                _ = v19
            }
            _ = fi15
        }
    default:
        delete(v3, 2147483647)
    }
    v20 := v3[v2]
    switch ("" + ("hello" + "foo bar")) {
    case "":
        v21 := GenS2{F0: -824, F1: 32767, F2: "hello"}
        _ = v21
    case "\n\t":
        v7 = 32767
    default:
        v22 := GenS1{F0: -218, F1: 1, F2: -32768}
        _ = v22
    }
    var v23 bool = (255 > uint8(v12))
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v6
    _ = v7
}
