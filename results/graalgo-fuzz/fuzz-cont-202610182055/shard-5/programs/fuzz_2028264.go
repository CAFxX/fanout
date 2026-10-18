package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16; F1 bool; F2 int16 }

func main() {
    var v1 uint32 = 4294967295
    v2 := []uint32{(v1 % ((v1) | 1)), 65535, (v1 & v1)}
    v3 := make(map[string]bool)
    v4 := []uint32{(v1 & v1), v1, (^(v1))}
    var v5 uint64 = 2
    switch "\n\t" {
    case "foo bar":
        var v6 uint16 = 65535
        _ = v6
    case "":
        var v7 int8 = 35
        _ = v7
    case "hello":
        switch "a" {
        case "foo bar":
            v5 += (v5 % ((4294967295) | 1))
        case "":
            for fi8 := 0; fi8 < 9; fi8++ {
                v5 = (v5 - (v5 + v5))
                _ = fi8
            }
        case "\n\t":
            v5 = 18446744073709551615
        default:
            if (-63 > 1) {
                v9 := GenS1{F0: 326, F1: ((0 != v1) && ((127 < int8(v5)) && true)), F2: 255}
                switch "a" {
                case "":
                    var v10 uint64 = (v5 + (^(v5)))
                    _ = v10
                case "\n\t":
                    for fi11 := 0; fi11 < 6; fi11++ {
                        v12 := GenS1{F0: 328, F1: (!true), F2: 65535}
                        var v13 int = (^(fi11))
                        if (false || true) {
                            var v14 int32 = 42
                            switch -768 {
                            case -147:
                                for fi15 := 0; fi15 < 16; fi15++ {
                                    _ = len(v2)
                                    delete(v3, (((("foo bar" + "\n\t") + "\n\t") + "hello") + (("hello" + "hello") + ("hello" + "a"))))
                                    _ = fi15
                                }
                            case 127:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                for fi16 := 0; fi16 < 13; fi16++ {
                                    var v17 uint32 = (v1 % ((v1) | 1))
                                    _ = fi16
                                    _ = v17
                                }
                            }
                            _ = v14
                        } else {
                            v18 := GenS1{F0: 1, F1: (v1 < v1), F2: 40}
                            var v19 uint64 = (v5 << ((18446744073709551615) & 63))
                            fmt.Println(fi11)
                            _ = v18
                            _ = v19
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
                delete(v3, "foo bar")
                _ = v9
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v20 int64 = int64(v1)
                _ = v20
            }
        }
    default:
        var v21 uint8 = 1
        _ = v21
    }
    v22 := GenS2{F0: 157, F1: false, F2: -402}
    v1 = v1
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 bool = (!false)
        _ = fi23
        _ = v24
    }
    fmt.Println(v1, v1)
    v1 = v1
    var v25 uint8 = 127
    var v26 uint32 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
