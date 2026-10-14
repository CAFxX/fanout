package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int64 {
    if (42 > -1) {
        if (2 == 749) {
            switch "" {
            case "hello":
                var v5 int8 = 127
                _ = v5
            }
            v6 := []int{-345, 787, 128, 1}
            _ = v6
        }
    }
    if false {
        var v7 uint32 = 317
        _ = v7
    }
    return 20
}

func main() {
    v1 := []int{}
    var v2 int = 32767
    v3 := []int{(v2 % ((42) | 1)), (v2 << ((18446744073709551615) & 63)), 1}
    var v4 uint64 = 4294967295
    var v8 uint64 = v4
    fmt.Println(v4, v2, v2)
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int16 = 42
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint16 = 1
            v13 := GenS1{F0: 105}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    if (115 >= -128) {
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 int = (fi14 - (^(fi14)))
            var v16 int16 = 801
            v2 += v15
            _ = fi14
            _ = v15
            _ = v16
        }
        switch "a" {
        case "hello":
            v8 = (v4 * 2)
        default:
            _ = len(v3)
        }
        var v17 string = ""
        _ = v17
    } else {
        var v18 bool = (!false)
        _ = v18
    }
    var v19 bool = true
    for fi20 := 0; fi20 < 2; fi20++ {
        switch ("a" + "foo bar") {
        case "foo bar":
            if v19 {
                switch ("" + "\n\t") {
                case "\n\t":
                    fmt.Println(v2, fi20, v8)
                case "foo bar":
                    var v21 int8 = -128
                    _ = v21
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = fi20
                    }
                }
                if len(v1) > 0 {
                    v1[0] = fi20
                }
            } else {
                var v22 bool = v19
                if (!(v22 || v22)) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if false {
                        for fi23 := 0; fi23 < 3; fi23++ {
                            v24 := GenS1{F0: 42}
                            _ = fi23
                            _ = v24
                        }
                        v22 = v22
                    }
                    var v25 int = fi20
                    _ = v25
                } else {
                    switch "a" {
                    case "\n\t":
                        v2 += v2
                    case "":
                        v2 = fi20
                    }
                    switch v2 {
                    case 612:
                        if true {
                            var v26 uint32 = 7
                            _ = v26
                        }
                    case 32768:
                        fmt.Println(v2, v4, v2)
                    default:
                        switch "" {
                        case "foo bar":
                            var v27 float32 = 1.5
                            _ = v27
                        case "":
                            v8 += v8
                        }
                    }
                    switch "a" {
                    case "\n\t":
                        if true {
                            var v28 int8 = -1
                            _ = v28
                        } else {
                            fmt.Println(v8, v19)
                        }
                    case "foo bar":
                        var v29 uint32 = 843
                        _ = v29
                    default:
                        _ = len(v1)
                    }
                }
                v2 = -1
                _ = v22
            }
        case "hello":
            fmt.Println(v4)
        case "":
            fmt.Println(v2, v19, v19)
        }
        _ = fi20
    }
    fmt.Println(v2, v8, v4)
    var v30 uint32 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v8
}
