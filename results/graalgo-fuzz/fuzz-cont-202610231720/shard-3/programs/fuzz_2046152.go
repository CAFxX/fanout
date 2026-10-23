package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint16; F2 bool }
type GenS4 struct { F0 uint32 }

func main() {
    v1 := []string{"a"}
    v2 := make(map[int]bool)
    var v3 float32 = 1e10
    var v4 uint8 = 42
    v5 := []int{}
    v6 := []uint32{322, 4294967295, 2147483648, 1}
    if true {
        var v7 int16 = 42
        _ = v7
    }
    v3 += float32(0.0)
    var v8 uint16 = 525
    v3 = v3
    var v9 int32 = 376
    var v10 int16 = -971
    if ("" != ("a" + "a")) {
        if ("a" == "foo bar") {
            v4 -= 128
            v3 = float32(3.14159)
        } else {
            v4 = v4
            var v11 int64 = 172
            _ = v11
        }
    } else {
        var v12 uint32 = 872
        var v13 int64 = 924
        if true {
            var v14 string = ("\n\t" + ("\n\t" + ("a" + ("\n\t" + "foo bar"))))
            switch 32768 {
            case 4294967295:
                v13 = v13
            case 804:
                if (v9 < -894) {
                    if false {
                        v15 := GenS1{F0: uint32(v8)}
                        _ = v15
                    } else {
                        v4 = (v4 % ((v4) | 1))
                        delete(v2, int(v8))
                        v3 = v3
                    }
                    for fi16 := 0; fi16 < 7; fi16++ {
                        if (!false) {
                            if true {
                                v17 := v2[fi16]
                                _ = v17
                            }
                            fmt.Println(v3, v14, v4)
                        } else {
                            v18 := GenS2{F0: fi16, F1: 4294967295}
                            _ = v18
                        }
                        v19 := GenS3{F0: v8, F1: 65535, F2: (v4 != (v4 << ((2) & 63)))}
                        _ = fi16
                        _ = v19
                    }
                    if false {
                        fmt.Println(v8, v14, v3)
                        v20 := GenS2{F0: int(v13), F1: 805}
                        _ = v20
                    }
                }
            }
            v10 += v10
            _ = v14
        }
        _ = v12
        _ = v13
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v9 -= v9
    for fi21 := 0; fi21 < 5; fi21++ {
        switch ("\n\t" + "\n\t") {
        case "":
            switch 127 {
            case -326:
                var v22 int64 = 9223372036854775807
                _ = v22
            case -1:
                var v23 uint8 = (^(v4))
                _ = v23
            }
        case "foo bar":
            v3 -= v3
        case "\n\t":
            fmt.Println(v9)
        }
        _ = fi21
    }
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    v24 := v2[32768]
    var v25 uint64 = 418
    if (!(v8 >= 2)) {
        v5 = append(v5, 4294967295)
        fmt.Println(v25)
        switch -723 {
        case -9223372036854775808:
            if ((v9 * v9) <= (v9 % (((^(v9))) | 1))) {
                for fi26 := 0; fi26 < 8; fi26++ {
                    if true {
                        fmt.Println(v3, v25, v8)
                        v2[(fi26 << ((331) & 63))] = v24
                    }
                    v8 = 42
                    _ = fi26
                }
                for fi27 := 0; fi27 < 15; fi27++ {
                    switch "" {
                    case "\n\t":
                        v25 = v25
                    case "a":
                        v4 = v4
                    }
                    _ = fi27
                }
            } else {
                if v24 {
                    switch 127 {
                    case 32767:
                        if v24 {
                            v8 -= v8
                        } else {
                            v28 := GenS2{F0: int(v10), F1: (^(v25))}
                            switch 1 {
                            case 2147483647:
                                fmt.Println(v24)
                            case 42:
                                if (1 >= 70) {
                                    v9 = v9
                                    v25 = v25
                                }
                            default:
                                v6 = append(v6, 65535)
                            }
                            _ = v28
                        }
                    case 4294967295:
                        switch "hello" {
                        case "a":
                            v9 = v9
                        }
                    case 0:
                        var v29 float32 = 3.14159
                        _ = v29
                    default:
                        var v30 int8 = -128
                        _ = v30
                    }
                    var v31 int16 = v10
                    _ = v31
                } else {
                    var v32 int16 = v10
                    if len(v1) > 0 {
                        v1[0] = ("\n\t" + "")
                    }
                    v33 := GenS2{F0: 1, F1: v25}
                    _ = v32
                    _ = v33
                }
            }
        }
    } else {
        v3 += -2.25
        v34 := GenS4{F0: 4294967295}
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
