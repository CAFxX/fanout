package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 string }
type GenS2 struct { F0 uint; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint16; F1 int32; F2 int }
func genFn0() uint64 {
    if (1 >= -586) {
        v4 := []int{2147483647, 0, -760, -65}
        var v5 int16 = -339
        _ = v4
        _ = v5
    }
    var v6 bool = false
    v7 := GenS1{F0: 32767, F1: 174, F2: "\n\t"}
    for fi8 := 0; fi8 < 13; fi8++ {
        if true {
            if v6 {
                v9 := GenS1{F0: 1, F1: 42, F2: "hello"}
                _ = v9
            } else {
                fmt.Println(v6, fi8)
                switch (fi8 | fi8) {
                case -1:
                    var v10 bool = ((v6 && (301 >= 42)) && v6)
                    _ = v10
                case 32768:
                    v6 = v6
                default:
                    fmt.Println(v6, v6)
                }
                for fi11 := 0; fi11 < 3; fi11++ {
                    v6 = (42 >= 32767)
                    _ = fi11
                }
            }
        }
        v6 = (1 <= 282)
        _ = fi8
    }
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 string) uint8 {
    v12 := GenS1{F0: -32768, F1: 4294967295, F2: p0}
    fmt.Println(p0)
    v13 := GenS2{F0: 255, F1: 127, F2: 906}
    fmt.Println(p0)
    var v14 int8 = 111
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 128
}

func main() {
    v1 := []uint32{42, 2147483648, 42}
    var v2 uint16 = 0
    var v3 int32 = -16
    for fi15 := 0; fi15 < 18; fi15++ {
        fmt.Println(fi15, v2, v3)
        v1 = append(v1, 42)
        _ = fi15
    }
    _ = len(v1)
    var v16 uint32 = 370
    v17 := GenS3{F0: (v2 * (v2 & 0)), F1: -1, F2: -709}
    for fi18 := 0; fi18 < 17; fi18++ {
        v2 -= 76
        var v19 int = fi18
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: 0, F1: 153, F2: -503}
    v21 := GenS1{F0: -32768, F1: 42, F2: ("" + "\n\t")}
    v16 = (v16 >> ((591) & 63))
    fmt.Println(v2)
    var v22 int32 = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
