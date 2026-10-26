package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 string; F1 string }
func genFn0(p0 uint64, p1 uint64) string {
    var v4 int8 = 1
    v5 := GenS1{F0: -31, F1: 0, F2: -646}
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint = 4294967295
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return ("" + ("" + "hello"))
}

func main() {
    var v1 int = 1
    var v2 bool = (!((-1 < 624) || true))
    var v3 uint8 = 42
    v8 := GenS2{F0: 62}
    var v9 uint64 = 1
    v10 := []int{v1, int(v3), v1}
    fmt.Println(v3, v2, v9)
    v9 += (v9 ^ v9)
    var v11 string = "hello"
    var v12 uint8 = v3
    v1 += 682
    fmt.Println(v1, v9, v12)
    if v2 {
        var v13 int8 = 0
        if (v13 >= int8(v1)) {
            var v14 uint8 = v12
            var v15 int32 = -690
            for fi16 := 0; fi16 < 13; fi16++ {
                v15 = (v15 % (((v15 << ((uint(v9)) & 63))) | 1))
                _ = fi16
            }
            _ = v14
            _ = v15
        } else {
            switch -423 {
            case -1:
                v10 = append(v10, (v1 ^ (v1 + (v1 + (v1 << ((18446744073709551615) & 63))))))
            case -609:
                v17 := GenS3{F0: 0, F1: (v13 << ((334) & 63)), F2: 2147483648}
                _ = v17
            default:
                for fi18 := 0; fi18 < 20; fi18++ {
                    var v19 float64 = 1e10
                    _ = fi18
                    _ = v19
                }
            }
            v20 := GenS4{F0: 4294967295}
            _ = v20
        }
        v21 := GenS5{F0: false, F1: 0}
        _ = v13
        _ = v21
    } else {
        var v22 int64 = -353
        _ = v22
    }
    switch -179 {
    case 109:
        fmt.Println(v2, v12, v1)
    case -877:
        fmt.Println(v3, v12)
    default:
        var v23 int32 = 2147483647
        _ = v23
    }
    v24 := GenS6{F0: v11, F1: ("" + v11)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v8
    _ = v9
}
