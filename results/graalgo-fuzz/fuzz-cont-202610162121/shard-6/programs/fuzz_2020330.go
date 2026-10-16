package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint64; F2 string }

func main() {
    var v1 int16 = 1
    var v2 uint8 = 115
    v3 := []uint32{}
    v4 := make(map[uint32]uint32)
    v2 += 1
    if ("hello" == "\n\t") {
        v2 = v2
    }
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(v1)
        _ = fi5
    }
    fmt.Println(v1, v2)
    if true {
        var v6 uint32 = 2147483647
        if len(v3) > 0 {
            v3[0] = 1
        }
        _ = v6
    } else {
        switch 4294967295 {
        case -1:
            var v7 bool = true
            _ = v7
        case 32768:
            v2 += v2
        case 32767:
            if (!(0 < 42)) {
                var v8 uint8 = 1
                v8 = v8
                _ = v8
            } else {
                v9 := GenS1{F0: 127}
                _ = v9
            }
        default:
            if false {
                switch 32767 {
                case 32768:
                    v4[871] = 760
                case 9223372036854775807:
                    for fi10 := 0; fi10 < 13; fi10++ {
                        if ((-9223372036854775808 > -9223372036854775808) || true) {
                            v2 -= 0
                            fmt.Println(fi10, fi10)
                        } else {
                            var v11 string = "hello"
                            var v12 int16 = v1
                            _ = v11
                            _ = v12
                        }
                        v2 = v2
                        if false {
                            var v13 string = "a"
                            fmt.Println(v13)
                            for fi14 := 0; fi14 < 6; fi14++ {
                                if true {
                                    var v15 int8 = 104
                                    var v16 uint32 = 2
                                    var v17 int16 = v1
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                } else {
                                    fmt.Println(v13, fi14)
                                    var v18 int32 = -589
                                    _ = v18
                                }
                                v13 = v13
                                _ = fi14
                            }
                            _ = v13
                        }
                        _ = fi10
                    }
                case -2147483648:
                    v19 := GenS2{F0: "hello"}
                    _ = v19
                }
            }
        }
    }
    var v20 float32 = float32((1.5 / 1.0))
    if (!true) {
        for fi21 := 0; fi21 < 14; fi21++ {
            switch fi21 {
            case 32767:
                fmt.Println(v20)
            default:
                v20 = v20
            }
            _ = fi21
        }
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            if true {
                v24 := v4[926]
                _ = v24
            }
            v2 = (^(v2))
            if (2 >= 18446744073709551615) {
                if (true && (-2147483648 > -210)) {
                    fmt.Println(fi23, v20, fi22)
                    v25 := GenS1{F0: 100}
                    _ = v25
                }
            }
            var v26 uint8 = (^(v2))
            _ = fi23
            _ = v26
        }
        _ = fi22
    }
    if true {
        var v27 int64 = 42
        v1 = -32768
        _ = v27
    } else {
        switch 904 {
        case -1:
            _ = len(v3)
        default:
            fmt.Println(v20)
        }
        v28 := GenS2{F0: "foo bar"}
        switch ("a" + "foo bar") {
        case "":
            if false {
                fmt.Println(v2, v2)
                v1 = v1
            } else {
                fmt.Println(v1)
                v1 -= v1
            }
        default:
            v29 := GenS3{F0: 255, F1: 4294967295, F2: ("hello" + ("foo bar" + "a"))}
            _ = v29
        }
        _ = v28
    }
    var v30 int8 = 0
    var v31 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
