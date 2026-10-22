package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int32; F2 int64 }
type GenS3 struct { F0 int32; F1 int; F2 bool }

func main() {
    var v1 uint16 = 0
    v2 := []uint32{273, 2}
    var v3 uint8 = 2
    var v4 uint8 = (v3 | v3)
    v4 = 106
    v1 = v1
    v1 = v1
    _ = len(v2)
    for fi5 := 0; fi5 < 2; fi5++ {
        if false {
            for fi6 := 0; fi6 < 13; fi6++ {
                if false {
                    if ((0 >= -89) || true) {
                        switch "a" {
                        case "a":
                            var v7 uint64 = 857
                            v3 = v3
                            _ = v7
                        case "foo bar":
                            fmt.Println(v4, v3)
                        case "":
                            fmt.Println(fi5)
                        }
                        var v8 bool = (!false)
                        fmt.Println(v3)
                        _ = v8
                    }
                }
                fmt.Println(v4, fi5)
                _ = fi6
            }
            switch fi5 {
            case 4294967295:
                v4 = 2
            case -9223372036854775808:
                v2 = append(v2, 65535)
            default:
                switch (fi5 ^ (fi5 / (((fi5 - 32767)) | 1))) {
                case 42:
                    if (42 > 42) {
                        _ = len(v2)
                    } else {
                        v1 = (v1 & v1)
                        var v9 bool = (18446744073709551615 < 42)
                        _ = v9
                    }
                case 2147483647:
                    v4 += (^(v3))
                case 410:
                    for fi10 := 0; fi10 < 12; fi10++ {
                        v11 := GenS1{F0: 538}
                        v12 := GenS2{F0: 65535, F1: -487, F2: 137}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                default:
                    v4 += 245
                }
            }
            for fi13 := 0; fi13 < 16; fi13++ {
                v2 = append(v2, 4294967295)
                if true {
                    if false {
                        if ((!(true && (!(!(v3 < v3))))) && ((-484 >= 1) && true)) {
                            v1 += 607
                            v4 = (v3 >> ((42) & 63))
                            v1 = (v1 - (v1 << ((307) & 63)))
                        }
                        var v14 uint8 = v3
                        switch (-(fi13)) {
                        case 4294967295:
                            v1 -= 0
                        case -2147483648:
                            switch "" {
                            case "":
                                fmt.Println(v1)
                            default:
                                fmt.Println(v3, fi5)
                            }
                            v4 = (v4 - (v4 * (v4 << ((uint(v1)) & 63))))
                            fmt.Println(v3)
                        }
                        _ = v14
                    }
                } else {
                    v15 := GenS1{F0: -1}
                    _ = v15
                }
                v4 += v3
                _ = fi13
            }
        }
        _ = fi5
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        v17 := GenS1{F0: 1}
        if false {
            v1 = 1
        } else {
            var v18 uint32 = 726
            var v19 uint64 = uint64(v1)
            for fi20 := 0; fi20 < 14; fi20++ {
                if false {
                    var v21 uint16 = (^(v1))
                    v1 = 32767
                    fmt.Println(v3)
                    _ = v21
                } else {
                    v3 += 127
                    fmt.Println(v3, v4)
                    var v22 bool = (0 >= 428)
                    _ = v22
                }
                if false {
                    if len(v2) > 0 {
                        v2[0] = (v18 & (v18 % ((v18) | 1)))
                    }
                }
                _ = fi20
            }
            _ = v18
            _ = v19
        }
        v3 += (v4 | v3)
        _ = fi16
        _ = v17
    }
    var v23 int16 = 0
    if true {
        if len(v2) > 0 {
            v2[0] = 65535
        }
    }
    if ("\n\t" >= "foo bar") {
        v24 := GenS3{F0: 346, F1: 4294967295, F2: true}
        _ = v24
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v25 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
