package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint16; F1 int; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int; F2 int }
func genFn0(p0 int16, p1 uint32) uint16 {
    v6 := GenS1{F0: (p1 & p1), F1: "", F2: 128}
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 = p0
        var v8 bool = (42 > 100)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 string = ""
    var v2 int = 32767
    var v3 string = ""
    var v4 int32 = -2147483648
    var v5 uint = 65535
    switch (v2 << ((v5) & 63)) {
    case 127:
        v4 = (v4 >> ((2) & 63))
    case 2147483647:
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS2{F0: -105, F1: -659}
            var v11 uint64 = 0
            _ = fi9
            _ = v10
            _ = v11
        }
    case -198:
        v12 := GenS3{F0: 76, F1: (v2 % (((v2 >> (((v5 * v5)) & 63))) | 1)), F2: 32767}
        _ = v12
    }
    v13 := []int{(v2 % ((-162) | 1)), v2}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    _ = len(v13)
    v5 -= (v5 & (^(v5)))
    fmt.Println(v4)
    var v14 int = (^(v2))
    if (!(42 >= 462)) {
        for fi15 := 0; fi15 < 6; fi15++ {
            if len(v13) > 0 {
                v13[0] = (v14 + v2)
            }
            _ = fi15
        }
        v16 := GenS3{F0: 964, F1: v2, F2: genFn0(-552, 4294967295)}
        var v17 bool = (v14 < 128)
        _ = v16
        _ = v17
    }
    switch 32767 {
    case -2147483648:
        v3 = v1
    }
    if false {
        for fi18 := 0; fi18 < 11; fi18++ {
            v1 = v1
            v19 := GenS1{F0: 216, F1: "", F2: (fi18 * fi18)}
            _ = fi18
            _ = v19
        }
        v3 = "foo bar"
    }
    var v20 int64 = 1
    v21 := GenS3{F0: uint16(v5), F1: -1, F2: 65535}
    var v22 uint8 = 1
    v23 := GenS4{F0: (v22 % (((^(v22))) | 1)), F1: (^(v14)), F2: v14}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
