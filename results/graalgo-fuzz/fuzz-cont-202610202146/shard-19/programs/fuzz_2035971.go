package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int32 }

func main() {
    v1 := []string{}
    v2 := make(map[string]uint32)
    var v3 bool = false
    var v4 uint32 = 4294967295
    v5 := GenS1{F0: 255, F1: 0, F2: -1}
    if (17 >= 42) {
        fmt.Println(v4, v4)
        if (!v3) {
            v6 := GenS1{F0: 1, F1: 18446744073709551615, F2: 0}
            _ = v6
        }
    }
    var v7 int64 = 0
    v4 -= v4
    var v8 uint = 425
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(v3)
        v4 += (^(v4))
        _ = fi9
    }
    if (!(v8 >= 249)) {
        _ = len(v2)
    } else {
        if (-909 == -2147483648) {
            v7 = v7
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 uint64 = 4294967295
                switch fi10 {
                case 1:
                    if v3 {
                        if true {
                            v1 = append(v1, ("" + ("" + "hello")))
                        }
                        var v12 uint16 = 1
                        var v13 int16 = -1
                        _ = v12
                        _ = v13
                    }
                case 714:
                    var v14 uint = (v8 / ((v8) | 1))
                    _ = v14
                default:
                    var v15 int16 = 0
                    _ = v15
                }
                v1 = append(v1, ("hello" + (("foo bar" + "\n\t") + "foo bar")))
                _ = fi10
                _ = v11
            }
            switch 42 {
            case -904:
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 string = "a"
                    if (!v3) {
                        var v18 int = fi16
                        fmt.Println(v4, v7)
                        _ = v18
                    } else {
                        var v19 string = v17
                        for fi20 := 0; fi20 < 20; fi20++ {
                            fmt.Println(v7)
                            var v21 int = (fi20 / (((fi16 ^ fi20)) | 1))
                            _ = fi20
                            _ = v21
                        }
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                }
            case -906:
                fmt.Println(v8, v7)
            }
        } else {
            if v3 {
                v22 := v2[("foo bar" + ("\n\t" + "hello"))]
                fmt.Println(v8, v3)
                var v23 uint = v8
                _ = v22
                _ = v23
            }
            v3 = true
            switch -9223372036854775808 {
            case 1:
                var v24 int = -9223372036854775808
                _ = v24
            default:
                switch -2147483648 {
                case 935:
                    if true {
                        v25 := GenS1{F0: 127, F1: uint64(v8), F2: -2147483648}
                        _ = v25
                    }
                }
            }
        }
        fmt.Println(v4, v7, v7)
    }
    var v26 uint64 = uint64(v8)
    for fi27 := 0; fi27 < 19; fi27++ {
        v28 := GenS1{F0: 127, F1: uint64(v8), F2: 0}
        _ = fi27
        _ = v28
    }
    v3 = ((255 > v8) || v3)
    var v29 uint8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
