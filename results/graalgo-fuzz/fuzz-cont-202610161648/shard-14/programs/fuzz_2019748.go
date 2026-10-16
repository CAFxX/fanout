package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint; F1 string; F2 string }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 uint; F1 int32 }
func genFn0() uint64 {
    v7 := []int{32767, -264, 637}
    var v8 int = -551
    var v9 uint16 = 270
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(v9, v9)
        var v11 uint = 65535
        _ = fi10
        _ = v11
    }
    if (127 != 0) {
        switch "foo bar" {
        case "foo bar":
            v7 = append(v7, (v8 * (v8 % (((v8 + v8)) | 1))))
        case "":
            v12 := GenS1{F0: 558, F1: 65535}
            _ = v12
        default:
            if false {
                v7 = append(v7, (^(v8)))
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                v13 := GenS1{F0: -9223372036854775808, F1: 4294967295}
                _ = v13
            } else {
                switch v8 {
                case -2147483648:
                    fmt.Println(v9, v9)
                }
            }
        }
        if ((-1 == 32767) && true) {
            _ = len(v7)
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := GenS2{F0: 0}
                var v16 float32 = 1e10
                _ = fi14
                _ = v15
                _ = v16
            }
            if (true || (false || (false || (42 < 42)))) {
                var v17 int16 = -32768
                v9 += (v9 - (v9 * 65535))
                if true {
                    var v18 uint = 42
                    _ = v18
                } else {
                    fmt.Println(v9, v8, v17)
                }
                _ = v17
            }
        } else {
            if true {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 uint8 = 0
                var v21 int32 = -1
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        switch -9223372036854775808 {
        case 32767:
            fmt.Println(v9, v9)
        }
    }
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint = 0
    v2 := []int{32767, 0, -1}
    v3 := make(map[string]uint32)
    var v4 int = -1
    v5 := make(map[uint32]uint32)
    var v6 uint8 = 255
    v1 = 2
    v22 := GenS2{F0: 39}
    _ = len(v2)
    v23 := GenS2{F0: -29}
    v1 = (v1 - (v1 & 376))
    var v24 int8 = -128
    var v25 float32 = 3.14159
    v26 := GenS2{F0: (v24 << (((^(v1))) & 63))}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!(1 == 2)) {
        if (32767 >= uint16(v6)) {
            v27 := v5[2147483648]
            delete(v5, 2)
            _ = v27
        }
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 bool = true
            var v30 int64 = -9223372036854775808
            for fi31 := 0; fi31 < 1; fi31++ {
                var v32 float32 = v25
                _ = fi31
                _ = v32
            }
            _ = fi28
            _ = v29
            _ = v30
        }
        var v33 bool = (v6 > (v6 << (((v1 & v1)) & 63)))
        _ = v33
    }
    v3["foo bar"] = 42
    switch "" {
    case "hello":
        if true {
            if (((v25 * 1.5) * (v25 + -2.25)) <= -2.25) {
                v24 = (v24 + (v24 / ((v24) | 1)))
                v1 += (v1 + (v1 + 18446744073709551615))
                var v34 float32 = v25
                _ = v34
            } else {
                switch "hello" {
                case "foo bar":
                    _ = len(v2)
                case "\n\t":
                    if false {
                        v35 := GenS1{F0: 0, F1: 2}
                        _ = v35
                    } else {
                        _ = len(v5)
                    }
                case "hello":
                    switch 2147483647 {
                    case 9223372036854775807:
                        for fi36 := 0; fi36 < 11; fi36++ {
                            for fi37 := 0; fi37 < 13; fi37++ {
                                v38 := GenS1{F0: 42, F1: 0}
                                v39 := GenS1{F0: 1, F1: 2147483647}
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                            _ = fi36
                        }
                    case -9223372036854775808:
                        v40 := GenS3{F0: 410, F1: (^(v24)), F2: 811}
                        _ = v40
                    case 4294967295:
                        if false {
                            if (v24 > int8(v1)) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v41 := GenS4{F0: (v1 >> (((v1 & (v1 / (((v1 ^ v1)) | 1)))) & 63)), F1: ("" + "a"), F2: "a"}
                                _ = v41
                            } else {
                                v25 = v25
                            }
                            v4 = 4294967295
                        } else {
                            _ = len(v3)
                        }
                    }
                }
                v42 := GenS5{F0: (v4 / ((v4) | 1)), F1: 32767}
                _ = v42
            }
        } else {
            v4 -= v4
            v4 += (v4 * -9223372036854775808)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case "a":
        var v43 uint16 = 65535
        _ = v43
    default:
        v44 := GenS6{F0: (v1 % (((v1 ^ 283)) | 1)), F1: -909}
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
