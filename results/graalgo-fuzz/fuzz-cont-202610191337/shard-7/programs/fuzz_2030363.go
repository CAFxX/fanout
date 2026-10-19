package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 bool }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint; F1 int16; F2 int32 }

func main() {
    var v1 int = 128
    v2 := make(map[string]bool)
    var v3 int32 = -2147483648
    var v4 uint16 = 1
    var v5 bool = (4294967295 == 0)
    v6 := GenS1{F0: 18446744073709551615, F1: (v3 % (((v3 % ((2147483647) | 1))) | 1)), F2: true}
    if v5 {
        var v7 int8 = 0
        _ = v7
    }
    v3 = v3
    v8 := []int{v1, v1}
    switch 295 {
    case -1:
        var v9 bool = (1 >= 42)
        _ = v9
    case 2147483647:
        switch v1 {
        case -2147483648:
            var v10 uint = 18446744073709551615
            _ = v10
        case 32767:
            v4 += v4
        case 42:
            switch 323 {
            case 0:
                _ = len(v2)
            case 32768:
                v11 := GenS1{F0: 4294967295, F1: (^(v3)), F2: v5}
                _ = v11
            }
        default:
            var v12 int = v1
            _ = v12
        }
    case 127:
        for fi13 := 0; fi13 < 7; fi13++ {
            if true {
                fmt.Println(v4, fi13)
            }
            v14 := GenS2{F0: 0, F1: -32768}
            _ = fi13
            _ = v14
        }
    }
    v4 += 71
    var v15 uint = 42
    var v16 int = 0
    for fi17 := 0; fi17 < 6; fi17++ {
        v2["hello"] = ((!("foo bar" <= "")) && true)
        v4 = v4
        var v18 uint = 18446744073709551615
        _ = fi17
        _ = v18
    }
    v19 := v2["hello"]
    if (!(-319 > 42)) {
        fmt.Println(v3)
    }
    if v5 {
        v4 = v4
        v4 = 546
        v4 = v4
    } else {
        var v20 string = "\n\t"
        v21 := GenS2{F0: 171, F1: 42}
        fmt.Println(v3)
        _ = v20
        _ = v21
    }
    delete(v2, "foo bar")
    var v22 int64 = 1
    delete(v2, "a")
    v23 := GenS3{F0: v15, F1: 595, F2: 680}
    fmt.Println(v22)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
