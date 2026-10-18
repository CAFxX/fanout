package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS3 struct { F0 string; F1 uint16; F2 uint16 }
type GenS4 struct { F0 string; F1 int }

func main() {
    var v1 uint8 = 255
    var v2 uint8 = v1
    v3 := make(map[string]string)
    var v4 uint = 2
    var v5 bool = false
    var v6 uint32 = 944
    switch -2147483648 {
    case 127:
        var v7 uint64 = 2
        _ = v7
    case -2147483648:
        var v8 int16 = 386
        _ = v8
    case 9223372036854775807:
        switch (("" + (("" + "a") + "foo bar")) + "\n\t") {
        case "a":
            v1 = 105
        case "\n\t":
            v9 := GenS1{F0: 473}
            _ = v9
        default:
            _ = len(v3)
        }
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            v1 += v1
            v12 := GenS2{F0: v2, F1: 211, F2: (^(v6))}
            v1 = 127
            var v13 uint32 = (^(v6))
            _ = fi11
            _ = v12
            _ = v13
        }
        v4 = (v4 ^ 18446744073709551615)
        _ = fi10
    }
    v4 = v4
    v14 := []int{}
    fmt.Println(v4)
    v2 = uint8(v6)
    fmt.Println(v2, v4, v5)
    fmt.Println(v4)
    v15 := GenS1{F0: 1}
    if len(v14) > 0 {
        v14[0] = -1
    }
    var v16 float32 = 0.0
    v1 -= 128
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            switch -455 {
            case 70:
                v19 := GenS2{F0: (v1 ^ v2), F1: 42, F2: 796}
                _ = v19
            case 26:
                if (!((v2 & (v1 >> ((457) & 63))) >= (^(v2)))) {
                    switch fi18 {
                    case 327:
                        switch fi17 {
                        case 1:
                            v4 -= (v4 % ((v4) | 1))
                        case 9223372036854775807:
                            v20 := GenS2{F0: v1, F1: -2147483648, F2: 0}
                            _ = v20
                        default:
                            if v5 {
                                if false {
                                    fmt.Println(v16)
                                    v1 -= (v1 >> ((v4) & 63))
                                    var v21 uint64 = 42
                                    delete(v3, "\n\t")
                                    v22 := GenS3{F0: "hello", F1: 628, F2: 32767}
                                    _ = v21
                                    _ = v22
                                }
                            } else {
                                var v23 bool = v5
                                v16 = ((v16 - 3.14159) / (-2.25 + v16))
                                var v24 float32 = 0.1
                                _ = v23
                                _ = v24
                            }
                        }
                    case -998:
                        var v25 uint = 0
                        _ = v25
                    case 495:
                        var v26 bool = (("foo bar" < "hello") && true)
                        _ = v26
                    default:
                        var v27 uint = (^(v4))
                        _ = v27
                    }
                    fmt.Println(v5, v4)
                    var v28 uint64 = 858
                    _ = v28
                } else {
                    fmt.Println(v16)
                    var v29 int8 = 31
                    if (!v5) {
                        fmt.Println(v4, v1)
                        v5 = v5
                        v16 -= (-2.25 * ((-2.25 - (v16 + float32(3.14159))) + 1.5))
                    }
                    _ = v29
                }
            case 397:
                if false {
                    v4 = (v4 & v4)
                } else {
                    fmt.Println(v2, v5, v16)
                }
            }
            _ = fi18
        }
        _ = len(v3)
        if (!((v4 << ((102) & 63)) >= 513)) {
            v30 := GenS2{F0: v2, F1: -2147483648, F2: (v6 ^ 298)}
            for fi31 := 0; fi31 < 11; fi31++ {
                var v32 uint8 = v1
                _ = fi31
                _ = v32
            }
            _ = v30
        } else {
            delete(v3, (("\n\t" + "a") + "\n\t"))
        }
        _ = fi17
    }
    var v33 int32 = 1
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    fmt.Println(v33, v6)
    v34 := GenS4{F0: "foo bar", F1: 72}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
