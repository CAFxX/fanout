package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int64) int8 {
    p0 = p0
    p0 = -1
    _ = p0
    return -128
}

func main() {
    v1 := []uint32{907, 722, 1}
    var v2 uint64 = 449
    var v3 uint = 65535
    var v4 uint64 = (v2 / (((v2 >> (((v3 >> (((v3 / ((65535) | 1))) & 63))) & 63))) | 1))
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 uint16 = 0
            switch (fi6 >> ((v3) & 63)) {
            case 9223372036854775807:
                v3 += v3
            case 1:
                fmt.Println(v4, v3, fi5)
            case 128:
                switch ("foo bar" + "foo bar") {
                case "":
                    var v8 uint64 = 4294967295
                    _ = v8
                }
            }
            v4 += v2
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    var v9 uint8 = 255
    var v10 uint16 = 0
    fmt.Println(v2, v2)
    if (-986 != int(v10)) {
        v9 = v9
        v3 = 42
        v4 = uint64(v10)
    } else {
        var v11 int = -2147483648
        fmt.Println(v3, v9)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v11
    }
    var v12 int64 = 42
    v9 += uint8(v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((false || (!true)) || ((v10 ^ v10) == v10)) {
        v4 = (v2 % (((v2 - 18446744073709551615)) | 1))
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 uint32 = 42
            _ = fi13
            _ = v14
        }
        var v15 int64 = int64(v2)
        _ = v15
    }
    v16 := GenS1{F0: genFn0(v12)}
    fmt.Println(v10, v10, v2)
    var v17 int16 = -1
    v18 := GenS1{F0: 100}
    v19 := GenS1{F0: 100}
    var v20 int64 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
