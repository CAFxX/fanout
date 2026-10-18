package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }

func main() {
    var v1 bool = (-1 != 0)
    var v2 bool = v1
    v3 := make(map[string]int)
    var v4 int16 = -901
    v5 := []int{6, -780, 1, 4294967295}
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint32 = 2147483647
        for fi8 := 0; fi8 < 10; fi8++ {
            delete(v3, "")
            fmt.Println(fi8, v1)
            _ = fi8
        }
        var v9 uint64 = 492
        _ = fi6
        _ = v7
        _ = v9
    }
    if (!false) {
        switch "foo bar" {
        case "":
            var v10 uint64 = 1
            _ = v10
        case "\n\t":
            if v1 {
                switch "foo bar" {
                case "hello":
                    for fi11 := 0; fi11 < 8; fi11++ {
                        var v12 uint64 = 834
                        if (((("\n\t" + "") + "a") + (("a" + "a") + "hello")) != "foo bar") {
                            if true {
                                if (127 > 204) {
                                    var v13 int32 = 1
                                    _ = len(v3)
                                    _ = v13
                                } else {
                                    v4 += v4
                                }
                                for fi14 := 0; fi14 < 8; fi14++ {
                                    v1 = (!(255 == 2))
                                    fmt.Println(v2, v12, v1)
                                    _ = fi14
                                }
                                v1 = v2
                            }
                        } else {
                            v4 -= (v4 % ((int16(fi11)) | 1))
                            var v15 uint32 = 42
                            var v16 int16 = v4
                            _ = v15
                            _ = v16
                        }
                        _ = fi11
                        _ = v12
                    }
                }
            }
        }
        if v2 {
            switch 4294967295 {
            case 127:
                v17 := GenS1{F0: v4, F1: (((v2 || false) || false) && (715 == 0))}
                _ = v17
            case 1000:
                var v18 uint32 = 4294967295
                _ = v18
            case 898:
                v1 = v2
            default:
                if v1 {
                    switch 721 {
                    case -2147483648:
                        if true {
                            v19 := GenS1{F0: v4, F1: true}
                            _ = v19
                        }
                    }
                    var v20 uint = 0
                    _ = v20
                }
            }
            var v21 uint16 = 2
            v2 = true
            _ = v21
        }
    } else {
        v22 := GenS1{F0: (v4 % (((v4 - 42)) | 1)), F1: v1}
        fmt.Println(v2, v4, v1)
        _ = v22
    }
    fmt.Println(v2)
    v2 = v2
    var v23 int8 = -84
    switch 2147483647 {
    case 42:
        v24 := GenS1{F0: 32767, F1: v1}
        _ = v24
    default:
        _ = len(v5)
    }
    var v25 uint32 = 666
    var v26 int8 = -1
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS1{F0: (v4 % ((138) | 1)), F1: true}
        _ = fi27
        _ = v28
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
