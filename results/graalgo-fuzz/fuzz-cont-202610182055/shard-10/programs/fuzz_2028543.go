package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint16; F1 int32; F2 int64 }

func main() {
    var v1 float64 = -2.25
    v2 := []uint32{680}
    var v3 uint32 = 65535
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 string = "hello"
        _ = fi4
        _ = v5
    }
    var v6 uint8 = 0
    if len(v2) > 0 {
        v2[0] = 170
    }
    v7 := GenS1{F0: 1, F1: 108, F2: 118}
    for fi8 := 0; fi8 < 8; fi8++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch (-(fi8)) {
        case 9223372036854775807:
            v6 = v6
        case -678:
            v6 += v6
        case 995:
            if false {
                v9 := GenS2{F0: 42}
                fmt.Println(fi8)
                _ = v9
            } else {
                v1 = v1
                var v10 uint = 255
                if true {
                    for fi11 := 0; fi11 < 18; fi11++ {
                        v6 = (v6 * 42)
                        var v12 bool = true
                        _ = fi11
                        _ = v12
                    }
                }
                _ = v10
            }
        }
        var v13 int8 = -128
        _ = fi8
        _ = v13
    }
    var v14 uint32 = (v3 % ((v3) | 1))
    v15 := GenS3{F0: 2147483647, F1: 0, F2: 32767}
    var v16 string = "foo bar"
    v17 := GenS4{F0: 2, F1: 1, F2: 42}
    v16 = (v16 + "")
    if len(v2) > 0 {
        v2[0] = v3
    }
    v14 += v3
    if false {
        for fi18 := 0; fi18 < 7; fi18++ {
            if (18446744073709551615 != 42) {
                for fi19 := 0; fi19 < 2; fi19++ {
                    fmt.Println(v16, v16, v14)
                    var v20 uint32 = (v14 ^ v14)
                    _ = fi19
                    _ = v20
                }
            }
            _ = fi18
        }
        v21 := GenS3{F0: -1, F1: 0, F2: 2}
        var v22 uint32 = 262
        _ = v21
        _ = v22
    } else {
        v14 -= 0
    }
    v14 = 488
    fmt.Println(v16, v16)
    v23 := GenS3{F0: -2147483648, F1: 2147483648, F2: uint16(v3)}
    fmt.Println(v14)
    v24 := GenS2{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v6
    _ = v7
}
