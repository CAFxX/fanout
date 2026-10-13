package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint32; F1 int32 }

func main() {
    var v1 int = 128
    v2 := []uint32{893, 0, 2147483647, 2}
    var v3 int32 = int32(v1)
    var v4 uint16 = 958
    var v5 uint64 = 425
    var v6 int64 = 9223372036854775807
    v1 = v1
    var v7 uint8 = 127
    for fi8 := 0; fi8 < 11; fi8++ {
        switch ((("a" + (("hello" + "hello") + "hello")) + "") + ((("foo bar" + "\n\t") + (("a" + "foo bar") + "")) + "")) {
        case "a":
            switch 127 {
            case 0:
                fmt.Println(fi8, v6, v7)
            case 408:
                v9 := GenS1{F0: false, F1: 1}
                _ = v9
            case -2147483648:
                v7 = v7
            }
        case "foo bar":
            var v10 int32 = -2147483648
            _ = v10
        default:
            switch "hello" {
            case "hello":
                _ = len(v2)
            case "a":
                if (false && false) {
                    v11 := GenS2{F0: v7, F1: 483}
                    if true {
                        v5 = (v5 ^ 121)
                    } else {
                        v5 = v5
                        if true {
                            var v12 string = ("hello" + "hello")
                            switch fi8 {
                            case -9223372036854775808:
                                v13 := GenS1{F0: ((-555 != (v6 >> ((65535) & 63))) && true), F1: 65535}
                                _ = v13
                            case 32767:
                                var v14 int64 = (v6 << ((290) & 63))
                                _ = v14
                            default:
                                v15 := GenS3{F0: v5, F1: (!false)}
                                _ = v15
                            }
                            v12 = ((v12 + v12) + "foo bar")
                            _ = v12
                        } else {
                            fmt.Println(v1, v1, v7)
                        }
                        var v16 uint8 = 118
                        _ = v16
                    }
                    var v17 int8 = -1
                    _ = v11
                    _ = v17
                }
            }
        }
        var v18 float32 = ((3.14159 / 1.0) * (0.1 * float32(0.1)))
        _ = fi8
        _ = v18
    }
    var v19 int32 = 2147483647
    var v20 uint64 = 42
    v21 := GenS4{F0: 0, F1: -782}
    if false {
        switch v1 {
        case 32767:
            if len(v2) > 0 {
                v2[0] = 1
            }
        default:
            switch -185 {
            case 32768:
                if false {
                    var v22 uint16 = (v4 % ((v4) | 1))
                    if true {
                        fmt.Println(v22, v1, v3)
                    }
                    _ = v22
                }
            }
        }
    }
    switch (-(v1)) {
    case 211:
        var v23 int = v1
        _ = v23
    case -510:
        v24 := GenS3{F0: uint64(v4), F1: false}
        _ = v24
    case -9223372036854775808:
        var v25 float64 = 0.0
        _ = v25
    }
    switch (v1 + v1) {
    case 0:
        v2 = append(v2, 4294967295)
    case 9223372036854775807:
        if false {
            var v26 int32 = 2147483647
            _ = v26
        } else {
            var v27 int32 = v3
            _ = v27
        }
    default:
        for fi28 := 0; fi28 < 6; fi28++ {
            fmt.Println(v19, v5)
            for fi29 := 0; fi29 < 1; fi29++ {
                if len(v2) > 0 {
                    v2[0] = 190
                }
                if ("hello" <= "") {
                    if true {
                        var v30 uint8 = v7
                        v5 = (v5 ^ v20)
                        v7 += v7
                        _ = v30
                    }
                }
                var v31 int64 = v6
                switch v1 {
                case 2147483647:
                    v19 -= v3
                default:
                    fmt.Println(v5)
                }
                var v32 uint = 1
                _ = fi29
                _ = v31
                _ = v32
            }
            switch "foo bar" {
            case "":
                for fi33 := 0; fi33 < 4; fi33++ {
                    v7 = 42
                    v5 = (^(v20))
                    v34 := GenS2{F0: (v7 << ((65535) & 63)), F1: 120}
                    _ = fi33
                    _ = v34
                }
            case "foo bar":
                var v35 int32 = (v19 << ((31) & 63))
                _ = v35
            }
            _ = fi28
        }
    }
    fmt.Println(v20, v1)
    if (-63 == 1) {
        var v36 string = ("foo bar" + ("foo bar" + "hello"))
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
