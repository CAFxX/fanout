package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []uint32{}
    v2 := []uint32{891}
    v3 := make(map[int]int)
    var v4 uint32 = 2147483648
    v5 := GenS1{F0: 0}
    var v6 int64 = 42
    v7 := v3[2147483647]
    var v8 uint = 0
    switch int(v4) {
    case -2147483648:
        switch "a" {
        case "a":
            if false {
                for fi9 := 0; fi9 < 16; fi9++ {
                    if true {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        fmt.Println(v7, v4, fi9)
                        for fi10 := 0; fi10 < 9; fi10++ {
                            if true {
                                v6 -= (v6 | 9223372036854775807)
                                v11 := GenS1{F0: 4294967295}
                                _ = v11
                            }
                            _ = fi10
                        }
                    }
                    v6 = 425
                    _ = fi9
                }
            }
        case "\n\t":
            v7 = v7
        default:
            for fi12 := 0; fi12 < 9; fi12++ {
                for fi13 := 0; fi13 < 1; fi13++ {
                    v3[(fi13 * (v7 ^ fi12))] = v7
                    var v14 int = (^(fi12))
                    var v15 string = ""
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v7 -= v7
                if (("a" > "hello") && false) {
                    fmt.Println(v8)
                    for fi16 := 0; fi16 < 20; fi16++ {
                        var v17 uint = v8
                        _ = len(v1)
                        v6 -= v6
                        _ = fi16
                        _ = v17
                    }
                }
                _ = fi12
            }
        }
    default:
        switch "\n\t" {
        case "":
            for fi18 := 0; fi18 < 11; fi18++ {
                for fi19 := 0; fi19 < 1; fi19++ {
                    var v20 int = -215
                    v3[v7] = (-(v7))
                    _ = fi19
                    _ = v20
                }
                fmt.Println(v7, fi18)
                for fi21 := 0; fi21 < 4; fi21++ {
                    v1 = append(v1, (^(v4)))
                    var v22 int = 128
                    switch v22 {
                    case -261:
                        v23 := GenS1{F0: 0}
                        _ = v23
                    default:
                        if (!false) {
                            v24 := GenS1{F0: 1}
                            v3[(fi18 >> (((v8 & v8)) & 63))] = -622
                            v25 := GenS1{F0: uint64(v4)}
                            _ = v24
                            _ = v25
                        }
                    }
                    _ = fi21
                    _ = v22
                }
                _ = fi18
            }
        case "hello":
            v26 := GenS1{F0: 736}
            _ = v26
        default:
            v7 = (-(v7))
        }
    }
    var v27 int16 = 128
    var v28 int16 = (v27 & 1)
    v3[int(v4)] = v7
    fmt.Println(v27, v8)
    for fi29 := 0; fi29 < 3; fi29++ {
        v30 := GenS1{F0: 18446744073709551615}
        v31 := v3[1]
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
