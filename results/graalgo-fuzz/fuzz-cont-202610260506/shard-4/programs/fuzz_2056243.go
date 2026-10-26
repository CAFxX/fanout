package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int64; F2 bool }
type GenS4 struct { F0 uint8; F1 bool; F2 uint64 }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{}
    v3 := make(map[uint32]uint32)
    var v4 float32 = 0.0
    if len(v2) > 0 {
        v2[0] = 42
    }
    v4 -= (1.5 - float32(1e10))
    v4 = (v4 * v4)
    v4 += 0.1
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 428}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        switch (fi7 ^ fi7) {
        case -656:
            var v8 float32 = v4
            _ = v8
        default:
            v4 += float32(3.14159)
        }
        _ = fi7
    }
    switch ("hello" + "") {
    case "foo bar":
        if true {
            v9 := GenS1{F0: 316}
            _ = v9
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 int = (fi10 & (^(fi10)))
                _ = fi10
                _ = v11
            }
            v4 += -2.25
        }
    }
    if false {
        fmt.Println(v4)
        if true {
            var v12 uint8 = 42
            _ = v12
        } else {
            _ = len(v3)
        }
    } else {
        var v13 uint = 0
        v14 := GenS2{F0: 937, F1: 42}
        v15 := GenS1{F0: -764}
        _ = v13
        _ = v14
        _ = v15
    }
    v16 := GenS3{F0: 4294967295, F1: -9223372036854775808, F2: (!(false && (622 < 0)))}
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 float32 = v4
        if (true && (((v4 > v18) || true) && true)) {
            v18 = v4
        }
        _ = fi17
        _ = v18
    }
    switch -2147483648 {
    case 42:
        if true {
            delete(v3, 811)
            v19 := GenS1{F0: 100}
            fmt.Println(v4)
            _ = v19
        }
    case -636:
        if true {
            for fi20 := 0; fi20 < 1; fi20++ {
                var v21 uint = 4294967295
                _ = fi20
                _ = v21
            }
        }
    case 127:
        if true {
            v4 -= v4
            if false {
                v4 -= (1.5 + 0.0)
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    if (float32(3.14159) == (v4 - float32(((3.14159 - 0.0) - (1e10 / 0.1))))) {
                        _ = len(v1)
                        delete(v1, "hello")
                        if false {
                            v4 += (3.14159 / 1.0)
                        } else {
                            v4 += (0.1 * (v4 / 1e10))
                            if false {
                                switch 128 {
                                case 0:
                                    v4 = 3.14159
                                case 1:
                                    if true {
                                        if true {
                                            switch -280 {
                                            case 0:
                                                if len(v2) > 0 {
                                                    v2[0] = 2
                                                }
                                            default:
                                                v4 = (((0.0 / v4) / v4) - float32((0.0 - (3.14159 - (1e10 * (-2.25 - (-2.25 - (3.14159 - 0.0))))))))
                                            }
                                        } else {
                                            v22 := GenS4{F0: 82, F1: false, F2: 1}
                                            _ = v22
                                        }
                                    } else {
                                        v4 = (v4 * (v4 + v4))
                                        fmt.Println(v4)
                                        if true {
                                            v4 = (v4 / (((((v4 * v4) / v4) * 0.0) - (v4 * v4)) / 1e10))
                                            _ = len(v2)
                                        }
                                    }
                                case 630:
                                    fmt.Println(v4)
                                default:
                                    v4 = v4
                                }
                                for fi23 := 0; fi23 < 6; fi23++ {
                                    var v24 uint64 = 0
                                    switch 1 {
                                    case -549:
                                        for fi25 := 0; fi25 < 12; fi25++ {
                                            v4 += 0.1
                                            v24 = (v24 & v24)
                                            _ = fi25
                                        }
                                    case 128:
                                        var v26 int32 = -2147483648
                                        _ = v26
                                    case 42:
                                        v27 := GenS4{F0: 2, F1: ((255 != 42) && (778 >= 65535)), F2: v24}
                                        _ = v27
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = 1
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                                var v28 uint32 = 321
                                _ = v28
                            }
                        }
                    } else {
                        switch "foo bar" {
                        case "hello":
                            v4 = -2.25
                        case "":
                            fmt.Println(v4)
                        }
                    }
                }
                fmt.Println(v4)
            } else {
                v4 = (((-2.25 - float32(3.14159)) - -2.25) + 0.1)
                v2 = append(v2, 2)
            }
            fmt.Println(v4)
        }
    default:
        if false {
            v4 = 3.14159
            v2 = append(v2, 2147483648)
            v4 = v4
        }
    }
    v29 := GenS3{F0: 2147483648, F1: 1, F2: false}
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
