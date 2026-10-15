package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int64 }
type GenS2 struct { F0 int; F1 int; F2 uint16 }
type GenS3 struct { F0 uint64 }
func genFn0() int {
    var v6 uint16 = 32767
    v6 += (v6 / (((v6 | v6)) | 1))
    v6 -= v6
    _ = v6
    return 4294967295
}

func main() {
    var v1 int8 = -128
    v2 := []int{-1, -372}
    v3 := make(map[string]bool)
    var v4 float32 = -2.25
    var v5 int8 = v1
    var v7 uint = 42
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint = v7
        var v10 int32 = 2147483647
        switch -220 {
        case -782:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case -271:
            switch (fi8 >> ((65535) & 63)) {
            case 128:
                v1 += -103
            default:
                v11 := GenS1{F0: 1, F1: -321, F2: -1}
                _ = v11
            }
        case 2147483647:
            delete(v3, ("a" + "hello"))
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 -= -59
    v12 := GenS2{F0: -2147483648, F1: 32768, F2: 609}
    v1 += (v5 - 100)
    if true {
        for fi13 := 0; fi13 < 1; fi13++ {
            if (v7 >= v7) {
                if ("\n\t" == "") {
                    var v14 bool = (!(true && (!false)))
                    v4 = v4
                    _ = v14
                }
            } else {
                var v15 string = ""
                v16 := GenS1{F0: 1, F1: 32767, F2: 9223372036854775807}
                _ = v15
                _ = v16
            }
            _ = fi13
        }
    } else {
        if len(v2) > 0 {
            v2[0] = genFn0()
        }
        v17 := GenS3{F0: 808}
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 uint8 = 88
        var v20 int32 = -884
        var v21 uint8 = uint8(v7)
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    v22 := v3["foo bar"]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
