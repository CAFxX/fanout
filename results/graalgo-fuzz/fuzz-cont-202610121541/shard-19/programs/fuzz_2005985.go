package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint; F2 int32 }

func main() {
    var v1 uint8 = 1
    v2 := []int{0, int(v1), 521}
    v3 := make(map[uint32]string)
    v1 = 255
    switch "hello" {
    case "\n\t":
        for fi4 := 0; fi4 < 12; fi4++ {
            v5 := GenS1{F0: 0}
            _ = fi4
            _ = v5
        }
    case "":
        fmt.Println(v1)
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        switch (^(fi6)) {
        case 156:
            var v7 int32 = 2147483647
            _ = v7
        case 407:
            if false {
                var v8 uint8 = v1
                var v9 int8 = -10
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if false {
                        fmt.Println(v8, v1, v1)
                        switch "foo bar" {
                        case "\n\t":
                            v1 = v8
                        case "a":
                            v10 := GenS1{F0: 115}
                            _ = v10
                        case "":
                            var v11 uint64 = 141
                            _ = v11
                        }
                        if false {
                            if true {
                                if ("hello" >= ("a" + (("foo bar" + ("hello" + "a")) + "hello"))) {
                                    fmt.Println(v8, v8)
                                }
                                fmt.Println(fi6)
                                if len(v2) > 0 {
                                    v2[0] = fi6
                                }
                            } else {
                                fmt.Println(v1)
                            }
                            v2 = append(v2, (^(fi6)))
                        }
                    } else {
                        v8 = (v1 * v8)
                        var v12 uint8 = v1
                        v12 += 128
                        _ = v12
                    }
                    v13 := GenS1{F0: 0}
                    _ = v13
                } else {
                    v8 = v8
                    var v14 uint8 = v1
                    if (270 > 255) {
                        if true {
                            v8 -= (v14 >> ((18446744073709551615) & 63))
                            for fi15 := 0; fi15 < 1; fi15++ {
                                v16 := GenS1{F0: 68}
                                _ = len(v3)
                                _ = fi15
                                _ = v16
                            }
                        } else {
                            switch "" {
                            case "a":
                                v17 := GenS2{F0: 2, F1: 908, F2: 1}
                                _ = v17
                            }
                        }
                        var v18 float64 = 0.1
                        _ = v18
                    }
                    _ = v14
                }
                _ = v8
                _ = v9
            }
        default:
            var v19 uint32 = 65535
            _ = v19
        }
        var v20 string = (("hello" + "\n\t") + (("" + "foo bar") + "foo bar"))
        _ = fi6
        _ = v20
    }
    if (-2.25 < float64(0.1)) {
        v1 = 255
        v1 -= (v1 + 255)
    } else {
        v3[4294967295] = "\n\t"
        if (!true) {
            v2 = append(v2, 1000)
        }
    }
    v2 = append(v2, -375)
    v1 = (v1 + v1)
    switch -9223372036854775808 {
    case 444:
        for fi21 := 0; fi21 < 2; fi21++ {
            if false {
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 uint32 = 1
                    v23 = (v23 + 0)
                    _ = fi22
                    _ = v23
                }
            } else {
                for fi24 := 0; fi24 < 11; fi24++ {
                    v1 -= (^(v1))
                    switch fi21 {
                    case 47:
                        delete(v3, 0)
                    case 42:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v25 uint8 = (v1 << ((65535) & 63))
                        _ = v25
                    case -1:
                        v26 := v3[42]
                        _ = v26
                    default:
                        var v27 bool = true
                        _ = v27
                    }
                    _ = fi24
                }
            }
            var v28 uint8 = v1
            _ = fi21
            _ = v28
        }
    }
    var v29 int8 = 0
    v29 += (v29 - 88)
    var v30 uint8 = (v1 & (v1 >> ((2) & 63)))
    switch 383 {
    case -321:
        v30 = 2
    case -1:
        var v31 int16 = 0
        _ = v31
    case 127:
        v32 := GenS3{F0: -835, F1: 65535, F2: -1}
        _ = v32
    }
    if false {
        var v33 int64 = -879
        _ = v33
    }
    var v34 uint16 = 65535
    v35 := GenS3{F0: 2147483647, F1: 18446744073709551615, F2: -1}
    v2 = append(v2, 9223372036854775807)
    for fi36 := 0; fi36 < 7; fi36++ {
        if len(v2) > 0 {
            v2[0] = (-(fi36))
        }
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
}
