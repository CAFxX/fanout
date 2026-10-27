package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int64 }
func genFn0() int64 {
    var v7 int16 = 592
    var v8 string = "a"
    if false {
        v9 := GenS1{F0: 128, F1: 42, F2: -1}
        for fi10 := 0; fi10 < 18; fi10++ {
            v7 = (-(v7))
            _ = fi10
        }
        switch v8 {
        case "foo bar":
            if true {
                var v11 uint16 = 42
                if false {
                    fmt.Println(v7, v7, v11)
                    v12 := []int{222, -694, 32768, 4294967295}
                    _ = v12
                } else {
                    var v13 int64 = 156
                    v13 = v13
                    switch 298 {
                    case 0:
                        var v14 bool = ((^(v11)) != 42)
                        _ = v14
                    case -1:
                        if (v13 > v13) {
                            v15 := []int{32768, -356, -2147483648, 32768}
                            v7 -= (v7 / ((32767) | 1))
                            _ = v15
                        }
                    case 364:
                        v16 := []int{-1, -2147483648, 0}
                        _ = v16
                    default:
                        fmt.Println(v11, v11)
                    }
                    _ = v13
                }
                v17 := []int{42, 32768}
                _ = v11
                _ = v17
            } else {
                v8 = v8
                v7 = v7
                if true {
                    if true {
                        var v18 bool = false
                        _ = v18
                    }
                } else {
                    var v19 int8 = 1
                    v20 := []int{-1, 128}
                    _ = v19
                    _ = v20
                }
            }
        case "":
            v8 = ""
        case "a":
            var v21 bool = (-2.25 == 0.1)
            _ = v21
        }
        _ = v9
    } else {
        var v22 uint32 = 2147483647
        if (!false) {
            switch v8 {
            case "foo bar":
                v23 := GenS1{F0: -2147483648, F1: 2147483647, F2: -1}
                _ = v23
            case "a":
                v7 = (v7 % ((280) | 1))
            case "":
                v24 := []int{-2147483648}
                _ = v24
            }
        } else {
            fmt.Println(v7, v7, v22)
            v25 := []int{32767}
            _ = v25
        }
        _ = v22
    }
    _ = v7
    _ = v8
    return -177
}

func main() {
    var v1 string = "\n\t"
    var v2 int8 = 124
    var v3 uint32 = 933
    var v4 int = 2147483647
    var v5 bool = (!((v4 & v4) != (v4 - 127)))
    var v6 uint64 = 0
    v1 = v1
    var v26 int8 = v2
    switch 919 {
    case 32767:
        v27 := []int{v4, (-(v4)), v4, v4}
        _ = v27
    case 32768:
        if v5 {
            var v28 uint8 = 255
            switch 32767 {
            case -2147483648:
                fmt.Println(v3, v4)
            case 128:
                v29 := GenS1{F0: v4, F1: -946, F2: 9223372036854775807}
                _ = v29
            }
            _ = v28
        } else {
            fmt.Println(v4)
            v1 = "\n\t"
        }
    }
    fmt.Println(v6, v3)
    var v30 int8 = (-(v2))
    var v31 uint32 = v3
    v30 = 42
    v5 = false
    var v32 uint32 = 446
    for fi33 := 0; fi33 < 7; fi33++ {
        v34 := GenS1{F0: v4, F1: 1, F2: 42}
        var v35 float32 = ((0.1 / -2.25) * ((3.14159 - 0.0) + 0.1))
        var v36 int = (fi33 ^ 347)
        _ = fi33
        _ = v34
        _ = v35
        _ = v36
    }
    v30 = (-(v26))
    v4 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
