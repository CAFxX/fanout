package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint32 }

func main() {
    var v1 uint64 = 444
    var v2 float32 = 0.1
    var v3 uint = 439
    v4 := []string{"\n\t"}
    v5 := make(map[string]int)
    v1 -= v1
    var v6 uint16 = 255
    v7 := GenS1{F0: 1, F1: v3}
    v1 = (v1 << (((v3 + 18446744073709551615)) & 63))
    v8 := GenS2{F0: (v3 & v3), F1: 636, F2: 141}
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 uint8 = 255
        v5[("hello" + "hello")] = (fi9 - 870)
        var v11 int64 = -1
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "a" {
    case "hello":
        if true {
            v6 = uint16(v3)
            fmt.Println(v1, v1)
            if (!(false && ("hello" < "foo bar"))) {
                if (true || (!(((-128 >= 103) || ((v1 & (v1 | (v1 << ((uint(v1)) & 63)))) < (v1 ^ v1))) || true))) {
                    for fi12 := 0; fi12 < 10; fi12++ {
                        if len(v4) > 0 {
                            v4[0] = "a"
                        }
                        _ = fi12
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    for fi13 := 0; fi13 < 9; fi13++ {
                        var v14 uint = (^(v3))
                        v15 := GenS1{F0: 1000, F1: 42}
                        v16 := GenS3{F0: 43, F1: 42, F2: v6}
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                } else {
                    for fi17 := 0; fi17 < 19; fi17++ {
                        switch "\n\t" {
                        case "hello":
                            v18 := GenS2{F0: (v3 - uint(v1)), F1: 899, F2: 42}
                            _ = v18
                        case "":
                            switch fi17 {
                            case 32767:
                                if (!((!false) || true)) {
                                    if false {
                                        v3 = 255
                                    } else {
                                        var v19 uint8 = 42
                                        _ = v19
                                    }
                                    switch "a" {
                                    case "a":
                                        var v20 int8 = -128
                                        _ = v20
                                    case "foo bar":
                                        v6 += uint16(v1)
                                    }
                                } else {
                                    v3 += v3
                                }
                            case 127:
                                for fi21 := 0; fi21 < 1; fi21++ {
                                    v1 = v1
                                    v3 -= (v3 + v3)
                                    _ = fi21
                                }
                            default:
                                switch int(v6) {
                                case -9223372036854775808:
                                    if false {
                                        fmt.Println(fi17)
                                    }
                                default:
                                    v22 := GenS1{F0: int16(v6), F1: (v3 * v3)}
                                    _ = v22
                                }
                            }
                        case "\n\t":
                            var v23 uint64 = (^(v1))
                            _ = v23
                        default:
                            switch "a" {
                            case "hello":
                                var v24 float32 = 3.14159
                                _ = v24
                            }
                        }
                        _ = fi17
                    }
                    delete(v5, "a")
                    _ = len(v4)
                }
                var v25 uint32 = 65535
                _ = v25
            }
        } else {
            switch 32768 {
            case 4294967295:
                var v26 int64 = -154
                _ = v26
            }
            for fi27 := 0; fi27 < 19; fi27++ {
                v2 += float32(0.0)
                _ = fi27
            }
        }
    case "a":
        for fi28 := 0; fi28 < 18; fi28++ {
            for fi29 := 0; fi29 < 8; fi29++ {
                fmt.Println(v3, fi29)
                fmt.Println(v6)
                _ = fi29
            }
            _ = fi28
        }
    }
    var v30 float32 = 1e10
    switch "a" {
    case "a":
        fmt.Println(v6)
    case "foo bar":
        v3 += v3
    case "\n\t":
        v1 = (v1 << ((594) & 63))
    }
    var v31 float64 = 1.5
    v31 = ((v31 * v31) - -2.25)
    switch "foo bar" {
    case "":
        v32 := GenS3{F0: -688, F1: 0, F2: 255}
        _ = v32
    case "a":
        var v33 int64 = -9223372036854775808
        _ = v33
    }
    var v34 uint64 = (v1 + v1)
    if (!false) {
        if (!(-921 <= 41)) {
            v1 -= v1
            v35 := GenS4{F0: 1, F1: 1}
            for fi36 := 0; fi36 < 20; fi36++ {
                v2 = (((float32((v31 + v31)) * v2) / v30) * -2.25)
                _ = fi36
            }
            _ = v35
        }
    }
    v3 = 287
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
