package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint16 }
func genFn0() uint {
    if false {
        var v6 int32 = 1
        _ = v6
    } else {
        switch -604 {
        case -9223372036854775808:
            v7 := []int{2147483647, -2147483648, -1, 9223372036854775807}
            _ = v7
        case 127:
            v8 := GenS1{F0: 0}
            _ = v8
        case 710:
            switch ("a" + ("foo bar" + "")) {
            case "foo bar":
                v9 := []int{32768, 9223372036854775807, -1, 3}
                _ = v9
            case "a":
                var v10 uint16 = 65535
                _ = v10
            }
        }
        var v11 uint64 = 853
        fmt.Println(v11)
        _ = v11
    }
    var v12 int = 4294967295
    _ = v12
    return 255
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[int]int)
    var v3 int32 = 42
    var v4 int8 = 127
    var v5 int32 = (v3 - (^(v3)))
    var v13 uint64 = 0
    v14 := v2[131]
    if false {
        v15 := GenS2{F0: 5, F1: 4294967295, F2: 2}
        _ = v15
    } else {
        var v16 uint64 = v13
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        v4 = (v4 * v4)
        _ = fi17
    }
    fmt.Println(v14, v3)
    if false {
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS3{F0: 2}
            var v20 bool = true
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    var v21 uint8 = 1
    var v22 bool = true
    v23 := []int{(v14 | v14)}
    v1 = (v1 & (v1 - (v1 * -1)))
    v13 = v13
    var v24 uint32 = 255
    _ = len(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
