package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint16 }
func genFn0() int32 {
    v6 := GenS1{F0: -876, F1: 100, F2: 829}
    var v7 int8 = 127
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 1
    var v3 uint8 = 29
    var v4 uint16 = (v2 * (^(v2)))
    var v5 uint16 = v2
    switch ("foo bar" + ("\n\t" + "")) {
    case "hello":
        if (689 == -9223372036854775808) {
            switch "" {
            case "\n\t":
                v8 := []int{}
                _ = v8
            case "foo bar":
                if false {
                    var v9 int64 = 1
                    _ = v9
                } else {
                    var v10 uint8 = 128
                    if true {
                        v11 := []int{-916, -9223372036854775808, 305, -1}
                        _ = len(v1)
                        for fi12 := 0; fi12 < 6; fi12++ {
                            var v13 uint16 = 32767
                            var v14 bool = ((!false) && true)
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        _ = v11
                    }
                    _ = v10
                }
            }
            var v15 float64 = (0.1 / 3.14159)
            _ = v15
        } else {
            var v16 int8 = 30
            _ = v16
        }
    case "":
        v3 -= v3
    default:
        var v17 uint32 = 4294967295
        _ = v17
    }
    switch (("hello" + "foo bar") + ("hello" + "a")) {
    case "foo bar":
        v18 := GenS1{F0: 0, F1: 127, F2: 2147483648}
        _ = v18
    case "":
        for fi19 := 0; fi19 < 5; fi19++ {
            fmt.Println(v4, v5)
            _ = fi19
        }
    default:
        var v20 float32 = float32(1e10)
        _ = v20
    }
    if ("hello" != "hello") {
        var v21 uint16 = (v2 | (^(v5)))
        v22 := []int{}
        v2 = (v2 + (v4 ^ 841))
        _ = v21
        _ = v22
    } else {
        if false {
            for fi23 := 0; fi23 < 2; fi23++ {
                var v24 bool = true
                v25 := v1[(fi23 >> ((379) & 63))]
                if (v3 > 2) {
                    if false {
                        fmt.Println(v4, fi23)
                    }
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 15; fi26++ {
                v3 -= (v3 ^ (v3 % ((v3) | 1)))
                _ = fi26
            }
        } else {
            for fi27 := 0; fi27 < 15; fi27++ {
                var v28 int = -2147483648
                v29 := GenS1{F0: -656, F1: 100, F2: 4294967295}
                var v30 int64 = -9223372036854775808
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        }
        v5 = v2
    }
    fmt.Println(v2, v5, v3)
    for fi31 := 0; fi31 < 14; fi31++ {
        for fi32 := 0; fi32 < 2; fi32++ {
            v3 += v3
            _ = fi32
        }
        var v33 uint16 = (^(v5))
        _ = fi31
        _ = v33
    }
    if false {
        v2 += (^(v2))
        switch "foo bar" {
        case "hello":
            var v34 int64 = 1
            _ = v34
        case "":
            var v35 int64 = 927
            _ = v35
        default:
            for fi36 := 0; fi36 < 6; fi36++ {
                var v37 float32 = (0.1 * 1e10)
                _ = fi36
                _ = v37
            }
        }
    } else {
        var v38 int16 = 42
        var v39 int8 = 1
        _ = v38
        _ = v39
    }
    var v40 int16 = 733
    if ((663 == 0) || false) {
        fmt.Println(v40, v40, v40)
    } else {
        var v41 int32 = 1
        var v42 float32 = -2.25
        _ = v41
        _ = v42
    }
    v43 := GenS2{F0: v4}
    v2 = 970
    v2 -= v5
    v2 += 419
    if false {
        var v44 uint8 = v3
        var v45 uint32 = 313
        v1[2147483647] = (v45 - v45)
        _ = v44
        _ = v45
    } else {
        v1[1] = 42
        var v46 float64 = float64(0.0)
        _ = v46
    }
    v47 := []int{518}
    if true {
        var v48 int64 = 0
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v47
    _ = v5
}
