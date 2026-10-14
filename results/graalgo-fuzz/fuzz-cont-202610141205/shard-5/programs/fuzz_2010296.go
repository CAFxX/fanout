package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int64; F1 int16 }

func main() {
    v1 := []string{""}
    v2 := []int{}
    v3 := []string{"a", ("foo bar" + "hello"), "a"}
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := GenS1{F0: 42, F1: 1000}
        _ = fi4
        _ = v5
    }
    var v6 uint32 = 42
    var v7 uint32 = v6
    var v8 int32 = 855
    var v9 bool = (!false)
    fmt.Println(v8, v6)
    v10 := GenS1{F0: -128, F1: 0}
    if v9 {
        for fi11 := 0; fi11 < 14; fi11++ {
            switch fi11 {
            case -1:
                v9 = (v9 && (!v9))
            case -9223372036854775808:
                v12 := GenS1{F0: int8(v7), F1: 0}
                _ = v12
            case 32767:
                if v9 {
                    var v13 uint64 = 2
                    _ = v13
                }
            default:
                var v14 string = "\n\t"
                _ = v14
            }
            _ = fi11
        }
        var v15 string = "\n\t"
        _ = v15
    } else {
        var v16 int8 = 0
        for fi17 := 0; fi17 < 11; fi17++ {
            v9 = false
            _ = fi17
        }
        if false {
            var v18 string = "hello"
            _ = v18
        } else {
            var v19 uint64 = 18446744073709551615
            if (1 < -1) {
                v9 = v9
            }
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            _ = v19
        }
        _ = v16
    }
    if false {
        if v9 {
            var v20 int64 = -9223372036854775808
            var v21 int32 = (v8 << ((18446744073709551615) & 63))
            _ = v20
            _ = v21
        }
        switch 128 {
        case 0:
            if len(v1) > 0 {
                v1[0] = (("" + (("hello" + "foo bar") + "foo bar")) + "foo bar")
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v22 float64 = 3.14159
    if v9 {
        for fi23 := 0; fi23 < 14; fi23++ {
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            _ = fi23
        }
    } else {
        var v24 bool = (4294967295 >= 945)
        v25 := GenS1{F0: -1, F1: -1}
        v1 = append(v1, "\n\t")
        _ = v24
        _ = v25
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v26 uint64 = 2
    v27 := GenS2{F0: 301, F1: 710}
    var v28 int32 = v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
