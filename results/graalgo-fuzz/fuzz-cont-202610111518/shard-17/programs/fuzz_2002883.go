package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int32 = -984
    var v2 uint64 = 1
    v3 := []uint32{0, 42, 906}
    var v4 string = "foo bar"
    if false {
        switch 969 {
        case 748:
            v1 = (v1 | 42)
        default:
            var v5 bool = (-32768 <= -646)
            _ = v5
        }
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int16 = 317
            _ = fi6
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 12; fi8++ {
            if false {
                for fi9 := 0; fi9 < 7; fi9++ {
                    v2 = (v2 / (((v2 - (v2 ^ 0))) | 1))
                    if len(v3) > 0 {
                        v3[0] = 603
                    }
                    _ = fi9
                }
            }
            _ = fi8
        }
        if ((42 > 895) && ("hello" < v4)) {
            fmt.Println(v2)
            var v10 uint64 = v2
            v4 = v4
            _ = v10
        } else {
            var v11 int16 = -37
            _ = v11
        }
        v2 -= 0
    }
    v4 = v4
    v1 = (v1 % ((-2147483648) | 1))
    var v12 uint64 = (v2 % ((v2) | 1))
    fmt.Println(v1)
    var v13 uint64 = (v12 & 59)
    if false {
        fmt.Println(v13, v1)
    }
    switch "\n\t" {
    case "a":
        v2 += v13
    case "foo bar":
        for fi14 := 0; fi14 < 13; fi14++ {
            v12 -= (v2 << ((91) & 63))
            _ = fi14
        }
    default:
        var v15 int32 = (v1 >> ((228) & 63))
        _ = v15
    }
    if false {
        if false {
            if (1.5 < 0.1) {
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 int = 0
                    for fi18 := 0; fi18 < 15; fi18++ {
                        v19 := GenS1{F0: true, F1: 1}
                        _ = fi18
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                }
            } else {
                v13 -= v2
                var v20 uint32 = 2147483647
                v21 := GenS1{F0: (343 == 1), F1: 42}
                _ = v20
                _ = v21
            }
        }
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = len(v3)
        }
        fmt.Println(v4, v4)
    }
    var v22 int = int(v12)
    switch (v22 % (((v22 / (((v22 - (v22 % (((v22 & (v22 << ((18446744073709551615) & 63)))) | 1)))) | 1))) | 1)) {
    case 32767:
        var v23 bool = true
        _ = v23
    case -372:
        var v24 uint = 65535
        _ = v24
    }
    fmt.Println(v13, v2, v12)
    if true {
        if ((v1 >> ((0) & 63)) >= v1) {
            if (float32((1.5 - 0.1)) < 1e10) {
                for fi25 := 0; fi25 < 11; fi25++ {
                    switch "\n\t" {
                    case "foo bar":
                        v2 = (v13 - v12)
                    case "a":
                        v2 += v13
                    case "\n\t":
                        switch v22 {
                        case 127:
                            switch (v4 + "foo bar") {
                            case "hello":
                                v4 = "hello"
                            }
                        }
                    default:
                        var v26 uint64 = 4294967295
                        _ = v26
                    }
                    _ = fi25
                }
            }
            v1 = (v1 + v1)
            var v27 uint = uint(v12)
            _ = v27
        }
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 int = -753
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = len(v3)
            _ = fi28
            _ = v29
        }
        var v30 int16 = -1
        _ = v30
    } else {
        fmt.Println(v13, v22, v22)
    }
    switch v4 {
    case "hello":
        switch "foo bar" {
        case "a":
            v4 = (v4 + "hello")
        case "\n\t":
            v22 += 4294967295
        default:
            v31 := GenS2{F0: false}
            _ = v31
        }
    default:
        v32 := GenS2{F0: true}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
