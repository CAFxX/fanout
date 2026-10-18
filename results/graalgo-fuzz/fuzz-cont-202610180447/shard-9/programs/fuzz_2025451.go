package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 bool }
func genFn0(p0 int8, p1 int8) int16 {
    p1 -= (p0 - (p1 + p1))
    var v6 int = -951
    _ = p0
    _ = p1
    _ = v6
    return -32768
}

func main() {
    var v1 bool = false
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    var v4 float32 = (1e10 / (float32(((3.14159 / 3.14159) + 1e10)) / 3.14159))
    var v5 string = "foo bar"
    v7 := GenS1{F0: 65535}
    if (!(false || ((!true) && false))) {
        switch v5 {
        case "a":
            delete(v3, 268)
        case "foo bar":
            _ = len(v2)
        default:
            var v8 uint = 65535
            _ = v8
        }
        var v9 uint16 = 508
        _ = v9
    } else {
        var v10 float32 = v4
        _ = v10
    }
    fmt.Println(v4)
    v5 = v5
    v5 = v5
    v5 = (v5 + v5)
    var v11 int = 0
    var v12 int16 = -32768
    var v13 int64 = 0
    if (!(v1 && v1)) {
        if v1 {
            var v14 int32 = 673
            _ = v14
        } else {
            if true {
                var v15 string = "foo bar"
                if v1 {
                    v15 = v5
                    switch v11 {
                    case -2147483648:
                        var v16 uint = 2
                        _ = v16
                    case 42:
                        v17 := GenS1{F0: 0}
                        _ = v17
                    }
                    v15 = (v5 + v15)
                } else {
                    v2[v11] = 2147483647
                    if false {
                        var v18 uint8 = 1
                        var v19 uint8 = (v18 << ((953) & 63))
                        _ = v18
                        _ = v19
                    }
                    v4 -= 0.1
                }
                v20 := GenS1{F0: 42}
                _ = v15
                _ = v20
            }
        }
    } else {
        var v21 int = (v11 ^ -992)
        v5 = "hello"
        switch int(v12) {
        case 1:
            v13 += (v13 / (((v13 + 0)) | 1))
        case 2147483647:
            var v22 uint = 4294967295
            _ = v22
        default:
            switch (v11 % ((2147483647) | 1)) {
            case 32768:
                v23 := GenS2{F0: 2, F1: v1}
                _ = v23
            case 2147483647:
                for fi24 := 0; fi24 < 8; fi24++ {
                    v25 := []int{}
                    delete(v2, (^(fi24)))
                    v12 = v12
                    _ = fi24
                    _ = v25
                }
            case 4294967295:
                fmt.Println(v21, v5)
            default:
                var v26 uint32 = 574
                _ = v26
            }
        }
        _ = v21
    }
    v27 := []int{}
    var v28 uint8 = 236
    switch -1 {
    case 782:
        var v29 uint16 = 42
        _ = v29
    case 9223372036854775807:
        fmt.Println(v4)
    case 4294967295:
        for fi30 := 0; fi30 < 16; fi30++ {
            for fi31 := 0; fi31 < 17; fi31++ {
                delete(v3, fi31)
                var v32 uint16 = 428
                _ = fi31
                _ = v32
            }
            var v33 int = 32767
            var v34 uint8 = v28
            _ = fi30
            _ = v33
            _ = v34
        }
    }
    var v35 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v7
}
