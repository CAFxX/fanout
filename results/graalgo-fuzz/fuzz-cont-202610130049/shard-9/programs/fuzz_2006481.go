package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := []int{-13, 1}
    var v2 string = "\n\t"
    var v3 uint = 183
    var v4 uint32 = 4294967295
    var v5 bool = false
    var v6 uint32 = v4
    switch v2 {
    case "a":
        for fi7 := 0; fi7 < 19; fi7++ {
            for fi8 := 0; fi8 < 4; fi8++ {
                if len(v1) > 0 {
                    v1[0] = fi8
                }
                var v9 int64 = -1
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    case "hello":
        switch ("a" + ((v2 + ("" + (v2 + v2))) + ((v2 + "\n\t") + v2))) {
        case "foo bar":
            v3 = 65535
        case "":
            v10 := GenS1{F0: 56}
            _ = v10
        }
    default:
        var v11 uint = (^(v3))
        _ = v11
    }
    var v12 float32 = float32(0.1)
    v13 := GenS1{F0: 0}
    var v14 float32 = 3.14159
    _ = len(v1)
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(v4, v6, v12)
        _ = fi15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        fmt.Println(v6)
        _ = fi16
    }
    if v5 {
        var v17 string = v2
        v17 = "a"
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v17
    } else {
        v18 := GenS1{F0: 100}
        v3 -= v3
        var v19 uint16 = 545
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        v3 += 42
        var v21 float64 = -2.25
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        v1 = append(v1, (fi22 >> (((v3 * 4294967295)) & 63)))
        if v5 {
            v3 = 1
            if v5 {
                var v23 int8 = 42
                if true {
                    v6 += v4
                    var v24 uint64 = 4294967295
                    _ = v24
                } else {
                    if v5 {
                        var v25 bool = false
                        for fi26 := 0; fi26 < 6; fi26++ {
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(v6, v5, v5)
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v6 -= 2147483647
                            v4 += (v6 | 4294967295)
                            v27 := GenS1{F0: v23}
                            _ = fi26
                            _ = v27
                        }
                        v28 := GenS2{F0: 808, F1: 0}
                        _ = v25
                        _ = v28
                    }
                }
                fmt.Println(v12)
                _ = v23
            }
            v5 = true
        }
        _ = fi22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
