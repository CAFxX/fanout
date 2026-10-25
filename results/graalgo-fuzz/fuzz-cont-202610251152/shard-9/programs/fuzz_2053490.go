package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint; F2 string }

func main() {
    v1 := []uint32{2, 661, 579}
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 144
    var v4 int8 = -128
    var v5 uint16 = 32767
    v6 := make(map[string]string)
    var v7 int64 = -978
    _ = len(v6)
    v7 -= v7
    switch -9223372036854775808 {
    case 125:
        v4 = v4
    case -1:
        v8 := GenS1{F0: 0, F1: 58}
        _ = v8
    case 0:
        v4 += (v4 << ((1) & 63))
    default:
        switch 32767 {
        case -60:
            v3 = 4294967295
        case -710:
            if false {
                if (!("a" > "")) {
                    if ((false || ((!true) && true)) || (42 >= 1)) {
                        switch -386 {
                        case -736:
                            switch 32768 {
                            case 0:
                                fmt.Println(v5, v3)
                            case -2147483648:
                                _ = len(v1)
                            case 1:
                                if (128 > 160) {
                                    switch 32768 {
                                    case 128:
                                        var v9 uint16 = (v5 + (v5 % ((v5) | 1)))
                                        _ = v9
                                    case 32768:
                                        var v10 int8 = (v4 % ((-1) | 1))
                                        _ = v10
                                    }
                                } else {
                                    v11 := GenS2{F0: v4}
                                    v4 = v4
                                    v4 += -128
                                    _ = v11
                                }
                            }
                        case 0:
                            v12 := v6["foo bar"]
                            _ = v12
                        case 9223372036854775807:
                            var v13 string = ""
                            _ = v13
                        }
                        v14 := v2[0]
                        _ = v14
                    }
                } else {
                    v5 = (^(v5))
                    fmt.Println(v3, v3)
                    v15 := GenS3{F0: 2147483647, F1: 0}
                    _ = v15
                }
                var v16 uint8 = 128
                switch "\n\t" {
                case "hello":
                    for fi17 := 0; fi17 < 5; fi17++ {
                        if false {
                            var v18 uint32 = v3
                            for fi19 := 0; fi19 < 18; fi19++ {
                                v5 = v5
                                switch fi17 {
                                case 4294967295:
                                    if len(v1) > 0 {
                                        v1[0] = v3
                                    }
                                case -1:
                                    _ = len(v1)
                                case 9223372036854775807:
                                    v3 -= v3
                                }
                                _ = fi19
                            }
                            v3 = v3
                            _ = v18
                        }
                        _ = fi17
                    }
                }
                _ = v16
            }
        case 0:
            var v20 int = 9223372036854775807
            _ = v20
        default:
            if true {
                v21 := GenS3{F0: 1, F1: 0}
                if len(v1) > 0 {
                    v1[0] = 0
                }
                fmt.Println(v4)
                _ = v21
            } else {
                switch "hello" {
                case "foo bar":
                    v3 = 0
                case "a":
                    v1 = append(v1, 1)
                case "hello":
                    v22 := GenS4{F0: (v3 - 2147483647), F1: 4294967295, F2: "\n\t"}
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 5; fi23++ {
                        fmt.Println(v7)
                        _ = fi23
                    }
                }
                var v24 uint8 = 42
                if false {
                    fmt.Println(v5, v5, v7)
                    for fi25 := 0; fi25 < 4; fi25++ {
                        var v26 uint32 = v3
                        var v27 uint = 124
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
                _ = v24
            }
        }
    }
    if (v7 < -1) {
        v4 = 127
        v7 -= v7
    }
    var v28 int8 = v4
    if true {
        v5 = 1
        var v29 bool = (79 < -32768)
        v4 -= 127
        _ = v29
    }
    v30 := v6["foo bar"]
    var v31 uint = 4294967295
    var v32 int = -115
    if len(v1) > 0 {
        v1[0] = (v3 + (v3 - 65535))
    }
    var v33 int = (^(v32))
    if (v7 > (v7 << (((v31 >> ((v31) & 63))) & 63))) {
        var v34 float32 = 0.0
        _ = v34
    } else {
        if len(v1) > 0 {
            v1[0] = (v3 << ((v31) & 63))
        }
        var v35 int = v33
        var v36 uint = (v31 | 255)
        _ = v35
        _ = v36
    }
    if false {
        for fi37 := 0; fi37 < 13; fi37++ {
            v7 -= 555
            var v38 int8 = 42
            var v39 uint16 = 0
            _ = fi37
            _ = v38
            _ = v39
        }
        var v40 int32 = 1
        v1 = append(v1, (^(v3)))
        _ = v40
    }
    fmt.Println(v3)
    if len(v1) > 0 {
        v1[0] = uint32(v5)
    }
    var v41 int32 = 0
    var v42 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
