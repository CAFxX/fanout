package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }

func main() {
    var v1 uint16 = 571
    var v2 bool = (0 < 42)
    v3 := make(map[uint32]string)
    var v4 int16 = 0
    v5 := make(map[string]uint32)
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 16; fi7++ {
        if false {
            v8 := []int{(^(fi7)), -946, fi7, (fi7 & fi7)}
            var v9 int16 = v4
            _ = v8
            _ = v9
        }
        _ = len(v5)
        if (!true) {
            var v10 uint64 = 1
            var v11 uint32 = 1
            _ = v10
            _ = v11
        }
        _ = fi7
    }
    var v12 int64 = 112
    if (0 >= 42) {
        v1 += (v1 << ((496) & 63))
    } else {
        if false {
            v2 = false
        } else {
            var v13 float32 = 3.14159
            _ = v13
        }
        v14 := []int{32768, int(v12), -2147483648, 264}
        _ = v14
    }
    var v15 uint16 = v1
    var v16 uint64 = 0
    v17 := []int{}
    for fi18 := 0; fi18 < 11; fi18++ {
        switch "\n\t" {
        case "a":
            switch -9223372036854775808 {
            case 128:
                var v19 uint8 = 14
                _ = v19
            case -979:
                var v20 int8 = -49
                _ = v20
            default:
                if true {
                    var v21 int8 = -1
                    v1 = v15
                    _ = v21
                }
            }
        case "":
            v22 := GenS1{F0: "a", F1: v2}
            _ = v22
        }
        v17 = append(v17, fi18)
        for fi23 := 0; fi23 < 8; fi23++ {
            delete(v6, "a")
            fmt.Println(fi18, v4)
            var v24 uint16 = (v15 / (((v1 - 1)) | 1))
            v25 := GenS1{F0: ("" + ("foo bar" + "a")), F1: v2}
            v24 = (^(v24))
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi18
    }
    v1 += (v15 ^ (v15 >> ((255) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
