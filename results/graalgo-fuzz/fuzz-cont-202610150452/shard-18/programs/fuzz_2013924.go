package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]bool)
    var v3 int32 = 1
    var v4 int64 = 972
    var v5 int32 = v3
    v4 = (v4 >> ((776) & 63))
    delete(v2, (v1 >> ((717) & 63)))
    v6 := GenS1{F0: v1}
    v7 := []int{-2147483648, 771}
    var v8 bool = false
    v4 -= (v4 - (v4 + v4))
    var v9 float64 = -2.25
    if (!(9223372036854775807 != -2147483648)) {
        v3 -= -363
        var v10 uint64 = 375
        _ = v10
    } else {
        for fi11 := 0; fi11 < 17; fi11++ {
            v8 = false
            fmt.Println(v9, v3)
            _ = fi11
        }
    }
    if true {
        var v12 uint8 = 9
        v9 -= v9
        switch 9223372036854775807 {
        case 32768:
            if v8 {
                for fi13 := 0; fi13 < 9; fi13++ {
                    v7 = append(v7, fi13)
                    v14 := GenS1{F0: v1}
                    var v15 float32 = 3.14159
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v16 uint32 = (v1 / ((v1) | 1))
                v3 += (v5 >> ((42) & 63))
                _ = v16
            }
        case -990:
            v17 := v2[2]
            _ = v17
        }
        _ = v12
    } else {
        var v18 uint32 = v1
        v5 -= v5
        _ = v18
    }
    if v8 {
        var v19 uint16 = 65535
        for fi20 := 0; fi20 < 13; fi20++ {
            v21 := GenS1{F0: v1}
            if (v8 && v8) {
                switch ("a" + "a") {
                case "foo bar":
                    var v22 int64 = (v4 << ((547) & 63))
                    _ = v22
                }
                v23 := GenS1{F0: (v1 % ((v1) | 1))}
                _ = v23
            }
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    switch ("" + "hello") {
    case "\n\t":
        fmt.Println(v1, v3, v1)
    case "foo bar":
        fmt.Println(v9)
    default:
        for fi24 := 0; fi24 < 10; fi24++ {
            v1 = 2
            _ = fi24
        }
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
