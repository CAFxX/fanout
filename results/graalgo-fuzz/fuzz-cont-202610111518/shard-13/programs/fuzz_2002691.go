package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int; F1 bool; F2 int }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }

func main() {
    var v1 uint32 = 500
    var v2 uint = 1
    var v3 uint16 = 2
    v4 := []int{32767}
    var v5 float32 = (0.1 / -2.25)
    switch -25 {
    case -1:
        var v6 uint8 = 17
        _ = v6
    case 2147483647:
        v7 := GenS1{F0: true, F1: 479}
        _ = v7
    case 539:
        v3 += v3
    default:
        switch -927 {
        case -9223372036854775808:
            for fi8 := 0; fi8 < 3; fi8++ {
                fmt.Println(v1)
                fmt.Println(v2)
                if false {
                    v5 = 0.0
                    switch (fi8 % (((fi8 - fi8)) | 1)) {
                    case 9223372036854775807:
                        v1 += (v1 | (v1 ^ (v1 + v1)))
                    case 127:
                        v9 := GenS2{F0: fi8, F1: true, F2: 2147483647}
                        _ = v9
                    case -976:
                        if (!false) {
                            fmt.Println(v2)
                        }
                    }
                    var v10 uint64 = 489
                    _ = v10
                } else {
                    var v11 int32 = -1
                    switch fi8 {
                    case 419:
                        v2 = v2
                    case 1:
                        for fi12 := 0; fi12 < 19; fi12++ {
                            var v13 uint = 4294967295
                            var v14 int32 = (v11 << ((4294967295) & 63))
                            v13 += v2
                            v5 -= float32(1.5)
                            var v15 uint = 391
                            _ = fi12
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                    default:
                        for fi16 := 0; fi16 < 10; fi16++ {
                            v17 := GenS3{F0: (fi8 / ((fi8) | 1)), F1: 18446744073709551615}
                            if (!((1e10 < 0.0) && (!(!(32767 != 42))))) {
                                fmt.Println(v3, fi16, v5)
                            }
                            v1 += (^(v1))
                            _ = fi16
                            _ = v17
                        }
                    }
                    _ = v11
                }
                _ = fi8
            }
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch "foo bar" {
        case "a":
            _ = len(v4)
        }
        v19 := GenS4{F0: "foo bar"}
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        v2 = (v2 | v2)
        fmt.Println(v2)
        if false {
            switch "\n\t" {
            case "a":
                v2 = (v2 * 65535)
            case "hello":
                v2 += v2
            case "foo bar":
                var v21 uint16 = v3
                _ = v21
            }
        } else {
            var v22 uint8 = 2
            v3 = v3
            if len(v4) > 0 {
                v4[0] = fi20
            }
            _ = v22
        }
        _ = fi20
    }
    if false {
        v3 = (v3 * (v3 % (((v3 % (((^(v3))) | 1))) | 1)))
        if len(v4) > 0 {
            v4[0] = -334
        }
        if true {
            switch "\n\t" {
            case "":
                if len(v4) > 0 {
                    v4[0] = 325
                }
            case "a":
                v5 += v5
            default:
                if true {
                    var v23 int64 = -238
                    _ = v23
                }
            }
        } else {
            if len(v4) > 0 {
                v4[0] = 128
            }
            v3 = v3
        }
    } else {
        if len(v4) > 0 {
            v4[0] = 42
        }
        v3 -= v3
    }
    switch ("\n\t" + "foo bar") {
    case "a":
        v24 := GenS4{F0: "a"}
        _ = v24
    case "\n\t":
        v1 += (v1 + 2)
    default:
        if (true || false) {
            var v25 uint16 = (^(v3))
            var v26 uint16 = 32767
            _ = v25
            _ = v26
        }
    }
    if true {
        _ = len(v4)
        var v27 int16 = 1000
        _ = v27
    } else {
        for fi28 := 0; fi28 < 4; fi28++ {
            for fi29 := 0; fi29 < 20; fi29++ {
                _ = len(v4)
                var v30 bool = (-1 > -186)
                if v30 {
                    v31 := GenS5{F0: 0}
                    _ = v31
                }
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
    }
    for fi32 := 0; fi32 < 6; fi32++ {
        if len(v4) > 0 {
            v4[0] = (^(fi32))
        }
        switch "a" {
        case "a":
            v1 = 1
        default:
            v2 = v2
        }
        var v33 int16 = -1
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
