package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -128
    var v3 string = "a"
    v4 := []uint32{42}
    var v5 uint32 = 624
    v6 := make(map[string]bool)
    v5 = (v5 ^ (v5 * v5))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = v3
    if (!true) {
        var v7 uint16 = 2
        _ = v7
    }
    if (32767 <= 436) {
        v6[v3] = (!(921 >= 32767))
        _ = len(v6)
    } else {
        if len(v4) > 0 {
            v4[0] = (v5 ^ (v5 + 0))
        }
        v2 = (v2 - v2)
    }
    if false {
        switch -2147483648 {
        case -9223372036854775808:
            v8 := v1[-577]
            _ = v8
        case 128:
            fmt.Println(v2, v2, v5)
        case 9223372036854775807:
            var v9 uint64 = 1
            _ = v9
        }
        if false {
            for fi10 := 0; fi10 < 3; fi10++ {
                fmt.Println(v3, v3)
                if false {
                    fmt.Println(v2)
                    var v11 int8 = -128
                    _ = v11
                } else {
                    var v12 uint16 = 42
                    for fi13 := 0; fi13 < 12; fi13++ {
                        var v14 int64 = 1
                        v15 := GenS1{F0: (v5 % (((v5 & 4294967295)) | 1)), F1: 42, F2: v12}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    _ = v12
                }
                _ = fi10
            }
            var v16 uint8 = 2
            switch (("a" + ("\n\t" + "foo bar")) + v3) {
            case "\n\t":
                delete(v1, 267)
            case "a":
                fmt.Println(v3, v16)
            }
            _ = v16
        }
    }
    if true {
        v5 -= 65535
    } else {
        var v17 uint32 = v5
        v5 = 0
        for fi18 := 0; fi18 < 11; fi18++ {
            for fi19 := 0; fi19 < 12; fi19++ {
                if true {
                    if false {
                        fmt.Println(v2)
                    } else {
                        v20 := GenS1{F0: (v17 * (v17 << ((0) & 63))), F1: 1, F2: 988}
                        _ = v20
                    }
                    v2 = 2
                    var v21 uint32 = v5
                    _ = v21
                }
                _ = fi19
            }
            var v22 int32 = 0
            _ = fi18
            _ = v22
        }
        _ = v17
    }
    var v23 int32 = -528
    var v24 uint64 = 995
    var v25 uint8 = 133
    var v26 float32 = (3.14159 + 1e10)
    fmt.Println(v2, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
