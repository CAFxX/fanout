package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint }

func main() {
    var v1 string = "hello"
    var v2 int = 0
    var v3 int = v2
    v3 -= -30
    v3 += (v2 * (v2 & v2))
    switch 0 {
    case 9223372036854775807:
        for fi4 := 0; fi4 < 2; fi4++ {
            v5 := []int{v2, 128, 4294967295, 4294967295}
            _ = fi4
            _ = v5
        }
    case 32767:
        v2 = v2
    default:
        switch (-(v3)) {
        case -482:
            v1 = v1
        case -9223372036854775808:
            var v6 uint = 356
            _ = v6
        case -2147483648:
            v3 -= 4294967295
        default:
            v1 = "foo bar"
        }
    }
    var v7 float64 = 3.14159
    switch v3 {
    case 250:
        fmt.Println(v2)
    default:
        v8 := []int{1, 330, v2}
        _ = v8
    }
    switch v1 {
    case "\n\t":
        v9 := GenS1{F0: 127, F1: 127, F2: 1}
        _ = v9
    case "foo bar":
        fmt.Println(v3, v7)
    default:
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 uint8 = 255
            v3 = (v3 ^ v2)
            _ = fi10
            _ = v11
        }
    }
    var v12 float32 = 1e10
    var v13 int = v2
    for fi14 := 0; fi14 < 18; fi14++ {
        switch ("hello" + "a") {
        case "foo bar":
            v15 := GenS1{F0: 0, F1: 128, F2: 2}
            _ = v15
        case "a":
            if false {
                for fi16 := 0; fi16 < 17; fi16++ {
                    var v17 uint64 = 4294967295
                    if true {
                        var v18 int64 = 812
                        var v19 int64 = v18
                        _ = v18
                        _ = v19
                    }
                    if false {
                        if (-37 == 42) {
                            v13 += (v2 | int(v17))
                            fmt.Println(fi14, v7, v12)
                            var v20 int64 = 9223372036854775807
                            if false {
                                var v21 bool = (!((v20 == -9223372036854775808) && true))
                                var v22 int16 = 32767
                                if v21 {
                                    v12 = v12
                                    var v23 uint16 = 957
                                    v13 += (^(v13))
                                    _ = v23
                                }
                                _ = v21
                                _ = v22
                            }
                            _ = v20
                        }
                    }
                    _ = fi16
                    _ = v17
                }
                var v24 int8 = 42
                _ = v24
            } else {
                v25 := GenS1{F0: 102, F1: 2, F2: 22}
                _ = v25
            }
        case "":
            switch v1 {
            case "\n\t":
                fmt.Println(v7, v1)
            }
        }
        var v26 int = (^(v2))
        _ = fi14
        _ = v26
    }
    if true {
        var v27 float32 = v12
        v28 := []int{(^(v2)), 32768, v13}
        var v29 int8 = -128
        _ = v27
        _ = v28
        _ = v29
    }
    v30 := []int{}
    if len(v30) > 0 {
        v30[0] = v2
    }
    var v31 bool = (42 >= 13)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v7
}
