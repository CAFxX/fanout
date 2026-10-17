package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
func genFn0(p0 bool) uint32 {
    var v6 string = "hello"
    var v7 uint32 = 225
    var v8 int16 = 0
    if true {
        v9 := GenS1{F0: 1, F1: 18446744073709551615}
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (^(v7))
}

func main() {
    var v1 float64 = 0.1
    var v2 uint8 = 2
    v3 := make(map[uint32]string)
    var v4 uint64 = 518
    var v5 string = "foo bar"
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 bool = true
        _ = fi10
        _ = v11
    }
    if ((322 >= 6) && (false || (!(!true)))) {
        v4 = v4
        v12 := GenS1{F0: 1, F1: (^(v4))}
        var v13 uint = uint(v2)
        _ = v12
        _ = v13
    } else {
        if ((false || false) && ((0.0 >= 1.5) && (false && ((-538 != 271) && (!false))))) {
            v1 += 1.5
        } else {
            if (!((v2 | (v2 ^ v2)) <= (v2 << ((876) & 63)))) {
                v4 += (v4 >> ((18446744073709551615) & 63))
                v5 = ("foo bar" + v5)
            }
            v14 := GenS1{F0: 0, F1: v4}
            _ = v14
        }
    }
    var v15 uint = 858
    if false {
        if false {
            fmt.Println(v4)
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 uint = v15
                v17 -= (v17 ^ v17)
                _ = fi16
                _ = v17
            }
            switch -722 {
            case -376:
                delete(v3, uint32(v4))
            case -2147483648:
                var v18 bool = (!true)
                _ = v18
            default:
                v15 += 255
            }
        }
    } else {
        var v19 uint16 = 1
        _ = v19
    }
    fmt.Println(v4, v4)
    var v20 uint8 = (v2 % (((v2 + 0)) | 1))
    v20 = (v20 + 10)
    v4 += (v4 + v4)
    fmt.Println(v1, v2, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
