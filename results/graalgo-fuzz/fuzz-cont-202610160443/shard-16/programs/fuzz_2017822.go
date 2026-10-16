package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 42
    v3 := []int{0, 456, 38}
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    for fi4 := 0; fi4 < 13; fi4++ {
        _ = len(v1)
        var v5 bool = true
        _ = fi4
        _ = v5
    }
    var v6 int64 = -1
    var v7 int16 = 32767
    v8 := v1[int(v6)]
    var v9 uint8 = 255
    if (v6 > int64(v7)) {
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 uint16 = 637
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint = 425
            var v14 int32 = 42
            if true {
                var v15 int = (fi12 ^ fi12)
                _ = v15
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        if v8 {
            fmt.Println(v9)
            if v8 {
                if v8 {
                    for fi16 := 0; fi16 < 17; fi16++ {
                        for fi17 := 0; fi17 < 20; fi17++ {
                            var v18 uint16 = 817
                            v1[fi16] = v8
                            var v19 float32 = (0.0 - 0.0)
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        _ = fi16
                    }
                    switch 127 {
                    case -9223372036854775808:
                        fmt.Println(v8, v6)
                    }
                } else {
                    switch ("hello" + "") {
                    case "hello":
                        var v20 uint16 = 0
                        _ = v20
                    default:
                        switch -2147483648 {
                        case -2147483648:
                            v1[379] = (32767 > -2147483648)
                        case 4294967295:
                            _ = len(v3)
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                    var v21 string = ("a" + "a")
                    v7 = v7
                    _ = v21
                }
                var v22 uint16 = 984
                switch 864 {
                case 1:
                    if len(v3) > 0 {
                        v3[0] = 9223372036854775807
                    }
                case 985:
                    if false {
                        var v23 int32 = 0
                        _ = v23
                    }
                default:
                    if (-94 > -68) {
                        switch "a" {
                        case "\n\t":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "a":
                            v7 = v2
                        default:
                            var v24 uint8 = (v9 % ((127) | 1))
                            _ = v24
                        }
                        v6 += (v6 - (v6 * v6))
                    } else {
                        for fi25 := 0; fi25 < 18; fi25++ {
                            for fi26 := 0; fi26 < 16; fi26++ {
                                v8 = (("a" + "") > "hello")
                                _ = fi26
                            }
                            fmt.Println(v9)
                            if ((v22 | (v22 - 923)) < (^(v22))) {
                                fmt.Println(v9)
                            }
                            _ = fi25
                        }
                    }
                }
                _ = v22
            }
            var v27 uint = 42
            _ = v27
        }
    } else {
        v9 = (^(v9))
        v28 := v1[-2147483648]
        var v29 uint = 2
        _ = v28
        _ = v29
    }
    fmt.Println(v9, v2, v9)
    switch "" {
    case "a":
        fmt.Println(v7, v2, v9)
    default:
        fmt.Println(v7, v9, v7)
    }
    _ = len(v1)
    v8 = (720 <= 2147483647)
    fmt.Println(v7)
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 uint32 = 42
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        if (!(!v8)) {
            v6 = (v6 ^ int64(fi32))
        }
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
