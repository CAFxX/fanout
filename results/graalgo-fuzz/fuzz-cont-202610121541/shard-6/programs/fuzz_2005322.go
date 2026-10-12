package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint16, p1 int64) uint16 {
    var v5 uint8 = 127
    v6 := []int{-2147483648, 32768}
    if false {
        var v7 int16 = 702
        p0 += p0
        if (float32(-2.25) > 3.14159) {
            fmt.Println(p0, v5)
            v8 := GenS1{F0: 1, F1: -1}
            _ = v8
        }
        _ = v7
    } else {
        v5 = (v5 >> ((0) & 63))
    }
    v9 := GenS1{F0: -36, F1: -9223372036854775808}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v9
    return 35
}

func main() {
    var v1 int8 = 0
    v2 := make(map[int]string)
    var v3 uint8 = 2
    var v4 uint8 = (v3 + (^(v3)))
    switch "" {
    case "a":
        v1 = -128
    case "foo bar":
        var v10 int16 = 0
        _ = v10
    }
    var v11 int32 = -598
    var v12 int64 = 146
    v13 := []int{128, 4294967295}
    v14 := GenS1{F0: 42, F1: v12}
    if false {
        switch "hello" {
        case "foo bar":
            v15 := GenS1{F0: v1, F1: v12}
            _ = v15
        case "\n\t":
            if false {
                v16 := v2[-2147483648]
                v3 = v4
                _ = v16
            } else {
                switch 127 {
                case -969:
                    _ = len(v13)
                case 2147483647:
                    var v17 bool = false
                    _ = v17
                default:
                    fmt.Println(v11)
                }
                if true {
                    delete(v2, 32767)
                    for fi18 := 0; fi18 < 3; fi18++ {
                        var v19 uint64 = 1
                        v4 += (v3 % (((v3 | v4)) | 1))
                        switch (("foo bar" + ("" + ("\n\t" + ""))) + "hello") {
                        case "":
                            v20 := GenS2{F0: (-(v12))}
                            _ = v20
                        case "foo bar":
                            fmt.Println(v1, v3, v12)
                        default:
                            fmt.Println(v19, v19, fi18)
                        }
                        _ = fi18
                        _ = v19
                    }
                } else {
                    if (1000 != 1000) {
                        v12 = (v12 << ((255) & 63))
                    }
                    switch "" {
                    case "\n\t":
                        v12 = v12
                    default:
                        v3 = (v4 >> ((18446744073709551615) & 63))
                    }
                    var v21 int = 128
                    _ = v21
                }
            }
        case "hello":
            switch 1 {
            case 32767:
                for fi22 := 0; fi22 < 9; fi22++ {
                    switch fi22 {
                    case 227:
                        var v23 uint8 = (^(v4))
                        _ = v23
                    case 1:
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    case -9223372036854775808:
                        switch (fi22 % ((fi22) | 1)) {
                        case -1:
                            switch 4294967295 {
                            case -215:
                                v11 += -933
                            case -9223372036854775808:
                                v1 += (v1 + v1)
                            }
                        default:
                            fmt.Println(v11, v3)
                        }
                    default:
                        v24 := GenS2{F0: (v12 % ((v12) | 1))}
                        _ = v24
                    }
                    var v25 int8 = 107
                    _ = fi22
                    _ = v25
                }
            case 42:
                switch 32768 {
                case 127:
                    var v26 uint8 = 255
                    _ = v26
                case 637:
                    if false {
                        for fi27 := 0; fi27 < 18; fi27++ {
                            fmt.Println(fi27, v3)
                            _ = fi27
                        }
                        switch 893 {
                        case 4294967295:
                            v4 = v3
                        default:
                            if false {
                                var v28 uint16 = 255
                                if true {
                                    fmt.Println(v12, v11, v4)
                                }
                                _ = v28
                            }
                        }
                    } else {
                        fmt.Println(v3, v4, v11)
                    }
                }
            }
        default:
            var v29 int64 = v12
            _ = v29
        }
        for fi30 := 0; fi30 < 8; fi30++ {
            var v31 uint8 = 255
            var v32 int16 = 1000
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    fmt.Println(v11)
    v12 = v12
    if len(v13) > 0 {
        v13[0] = -2147483648
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v33 float32 = 0.0
    v11 += v11
    v3 = v3
    v1 -= v1
    v4 = 1
    fmt.Println(v33, v11, v4)
    var v34 uint32 = 42
    v35 := GenS1{F0: v1, F1: v12}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
