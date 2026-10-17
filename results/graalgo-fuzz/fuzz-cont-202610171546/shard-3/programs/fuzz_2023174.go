package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint32 }
func genFn0() bool {
    v7 := []int{224, 32768, 42, 438}
    var v8 uint = 4294967295
    _ = v7
    _ = v8
    return true
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"a", "foo bar", ("hello" + ("" + (("foo bar" + "foo bar") + "a")))}
    var v3 bool = true
    var v4 int64 = -936
    var v5 int32 = -2147483648
    var v6 int32 = 2147483647
    switch -2147483648 {
    case 51:
        var v9 uint64 = 18446744073709551615
        _ = v9
    case 478:
        v4 = (-(v4))
    case -9223372036854775808:
        switch ("" + ("" + "a")) {
        case "":
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := GenS1{F0: (v6 | (v6 - (v6 & -1))), F1: 220, F2: 65535}
                _ = fi10
                _ = v11
            }
        case "foo bar":
            delete(v1, 420)
        case "\n\t":
            for fi12 := 0; fi12 < 3; fi12++ {
                v13 := GenS1{F0: v6, F1: 19, F2: 1}
                fmt.Println(fi12)
                _ = fi12
                _ = v13
            }
        default:
            var v14 uint = 225
            _ = v14
        }
    }
    var v15 uint64 = 1
    var v16 int64 = v4
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    v1[-2147483648] = 707
    for fi17 := 0; fi17 < 9; fi17++ {
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 uint16 = 538
            var v20 uint64 = 2
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 uint16 = 70
        _ = fi17
        _ = v21
    }
    v15 += (v15 << ((42) & 63))
    var v22 uint64 = v15
    var v23 uint16 = 1
    var v24 uint8 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
