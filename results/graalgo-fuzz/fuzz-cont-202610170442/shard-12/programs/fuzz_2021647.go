package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint8 = 1
    var v2 uint64 = 4294967295
    v3 := []uint32{2147483648, 65535}
    var v4 string = "a"
    v5 := make(map[uint32]int)
    var v6 float32 = 1.5
    v1 = (v1 / (((v1 ^ 183)) | 1))
    if true {
        var v7 uint64 = v2
        v8 := GenS1{F0: -128, F1: 42, F2: 4294967295}
        if false {
            switch -126 {
            case 629:
                for fi9 := 0; fi9 < 4; fi9++ {
                    for fi10 := 0; fi10 < 8; fi10++ {
                        var v11 int32 = 0
                        var v12 uint = 4294967295
                        var v13 uint16 = 0
                        _ = fi10
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    fmt.Println(fi9)
                    _ = fi9
                }
            case 128:
                v3 = append(v3, 2)
            case 1:
                var v14 uint16 = 65535
                _ = v14
            default:
                for fi15 := 0; fi15 < 12; fi15++ {
                    if true {
                        v6 = (float32(1e10) + v6)
                        if false {
                            if (2 >= 2147483648) {
                                var v16 uint32 = 571
                                if (uint16(v16) > 255) {
                                    switch -846 {
                                    case 2147483647:
                                        switch (v4 + v4) {
                                        case "hello":
                                            v17 := GenS1{F0: -123, F1: -128, F2: (v16 % (((v16 % (((v16 | v16)) | 1))) | 1))}
                                            _ = v17
                                        default:
                                            for fi18 := 0; fi18 < 9; fi18++ {
                                                v19 := GenS1{F0: 42, F1: 100, F2: (^(v16))}
                                                if true {
                                                    var v20 bool = true
                                                    v20 = (9223372036854775807 == 9223372036854775807)
                                                    _ = v20
                                                }
                                                fmt.Println(fi18, v16)
                                                _ = fi18
                                                _ = v19
                                            }
                                        }
                                    case -9223372036854775808:
                                        var v21 uint8 = 2
                                        _ = v21
                                    case 861:
                                        var v22 int = (-(fi15))
                                        _ = v22
                                    }
                                    v16 = (v16 - (v16 ^ (v16 | (v16 - (v16 * v16)))))
                                    switch ("foo bar" + (v4 + "")) {
                                    case "":
                                        switch v4 {
                                        case "\n\t":
                                            var v23 bool = (v6 == v6)
                                            _ = v23
                                        default:
                                            var v24 string = "\n\t"
                                            _ = v24
                                        }
                                    default:
                                        if (-1 == 844) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = 492
                                            }
                                            var v25 int64 = -1
                                            _ = v25
                                        }
                                    }
                                }
                                _ = v16
                            }
                            fmt.Println(v7)
                        }
                        fmt.Println(v6, fi15, v6)
                    } else {
                        v1 = v1
                    }
                    _ = fi15
                }
            }
        } else {
            var v26 int8 = int8(v7)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v6 -= 0.1
            _ = v26
        }
        _ = v7
        _ = v8
    } else {
        for fi27 := 0; fi27 < 3; fi27++ {
            fmt.Println(v6)
            v1 -= v1
            _ = fi27
        }
    }
    var v28 uint8 = v1
    var v29 string = v4
    var v30 uint = 65535
    fmt.Println(v6)
    v31 := GenS2{F0: uint8(v2), F1: uint(v28), F2: (v30 >> ((v30) & 63))}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch v4 {
    case "":
        var v32 string = v4
        _ = v32
    case "a":
        v33 := GenS3{F0: v1}
        _ = v33
    case "\n\t":
        v1 = uint8(v30)
    default:
        var v34 uint16 = 164
        _ = v34
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        var v36 float64 = (0.0 / ((float64(v6) * 0.1) / ((1e10 + 0.0) * float64(0.1))))
        _ = fi35
        _ = v36
    }
    switch int(v28) {
    case 42:
        var v37 int16 = 32767
        _ = v37
    default:
        v6 -= v6
    }
    v38 := GenS4{F0: 65535}
    if (172 <= 850) {
        fmt.Println(v6, v28)
        fmt.Println(v30, v4)
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        v30 = v30
        _ = fi39
    }
    delete(v5, 42)
    fmt.Println(v2, v30)
    for fi40 := 0; fi40 < 7; fi40++ {
        if (-32 < 4294967295) {
            for fi41 := 0; fi41 < 9; fi41++ {
                v28 -= (^(v28))
                v1 = v1
                _ = fi41
            }
            v1 = v1
            if len(v3) > 0 {
                v3[0] = 2
            }
        }
        _ = fi40
    }
    var v42 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
