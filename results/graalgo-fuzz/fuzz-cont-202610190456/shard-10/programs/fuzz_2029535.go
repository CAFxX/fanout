package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 int64 }
type GenS3 struct { F0 uint16; F1 int16; F2 int16 }

func main() {
    var v1 int32 = 42
    var v2 int64 = -926
    var v3 float32 = 0.0
    var v4 uint16 = 119
    v5 := []int{32767}
    v6 := []int{42, -9223372036854775808, 127, 32767}
    v1 = 42
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        v1 = v1
        var v8 uint8 = 0
        var v9 uint = 215
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v4 -= v4
        fmt.Println(fi10, fi10, v4)
        _ = fi10
    }
    fmt.Println(v2, v2, v4)
    v1 = (v1 & v1)
    if true {
        v1 += (v1 | 0)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v1 = 1
    } else {
        fmt.Println(v1, v4)
        var v11 int = 770
        if true {
            v12 := GenS1{F0: 0, F1: (!false)}
            fmt.Println(v4, v1)
            v2 = v2
            _ = v12
        }
        _ = v11
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := GenS2{F0: 128, F1: v2, F2: -202}
        _ = fi13
        _ = v14
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v15 int8 = 100
    switch 42 {
    case 1:
        if false {
            for fi16 := 0; fi16 < 3; fi16++ {
                if true {
                    v4 -= (v4 ^ v4)
                    if len(v5) > 0 {
                        v5[0] = (fi16 / ((fi16) | 1))
                    }
                    var v17 int8 = v15
                    _ = v17
                } else {
                    fmt.Println(v15, v4)
                }
                for fi18 := 0; fi18 < 20; fi18++ {
                    var v19 int64 = -9223372036854775808
                    var v20 uint64 = 1
                    switch "foo bar" {
                    case "foo bar":
                        var v21 uint64 = (v20 << ((2) & 63))
                        _ = v21
                    case "a":
                        if false {
                            var v22 string = "\n\t"
                            switch ("\n\t" + v22) {
                            case "foo bar":
                                var v23 uint64 = 1
                                _ = v23
                            default:
                                fmt.Println(v22, fi18, v4)
                            }
                            _ = v22
                        }
                    }
                    var v24 int = (fi16 ^ (^(fi18)))
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v24
                }
                fmt.Println(v1, fi16)
                _ = fi16
            }
        }
    default:
        v1 -= v1
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if true {
            v1 += 0
            v5 = append(v5, (fi25 ^ fi25))
        }
        for fi26 := 0; fi26 < 20; fi26++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi26
        }
        _ = fi25
    }
    v27 := GenS3{F0: v4, F1: 1000, F2: 13}
    v2 += 1
    v2 = v2
    var v28 uint64 = 73
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
