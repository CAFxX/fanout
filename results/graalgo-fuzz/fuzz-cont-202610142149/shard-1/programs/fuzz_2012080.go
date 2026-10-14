package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []uint32{1, 2147483647, 2147483647}
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 115
    v4 := make(map[uint32]string)
    var v5 int32 = -331
    var v6 uint16 = 1
    if (-870 >= -32768) {
        v3 -= v3
        fmt.Println(v5)
        v7 := GenS1{F0: 421}
        _ = v7
    } else {
        if (!(true || false)) {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 int8 = -1
                _ = fi8
                _ = v9
            }
        } else {
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 float64 = 3.14159
                if true {
                    var v12 int = fi10
                    v11 = (float64(1e10) - v11)
                    _ = v12
                } else {
                    var v13 int8 = -128
                    var v14 int8 = v13
                    _ = v13
                    _ = v14
                }
                if false {
                    v15 := GenS1{F0: uint32(v6)}
                    for fi16 := 0; fi16 < 4; fi16++ {
                        if (229 > 4294967295) {
                            if false {
                                if (!true) {
                                    var v17 uint32 = 900
                                    if len(v1) > 0 {
                                        v1[0] = 985
                                    }
                                    _ = v17
                                }
                                var v18 uint32 = 2147483648
                                v11 = (0.1 + v11)
                                _ = len(v2)
                                if ((fi16 - -1) >= -1) {
                                    v6 = 0
                                    v19 := GenS1{F0: v18}
                                    var v20 float32 = ((0.1 - 0.1) - (1e10 / 1.5))
                                    _ = v19
                                    _ = v20
                                }
                                _ = v18
                            } else {
                                switch fi16 {
                                case 127:
                                    var v21 int64 = 1
                                    _ = v21
                                default:
                                    if (!false) {
                                        v5 -= 0
                                    }
                                }
                                if ((2147483648 != 42) && (-9223372036854775808 == -383)) {
                                    if len(v1) > 0 {
                                        v1[0] = 0
                                    }
                                } else {
                                    if true {
                                        switch "\n\t" {
                                        case "hello":
                                            v6 = (v6 % ((65535) | 1))
                                            v5 = 42
                                        case "foo bar":
                                            var v22 uint32 = 2
                                            _ = v22
                                        default:
                                            v3 += (v3 * (^(v3)))
                                            if true {
                                                v3 = 128
                                            } else {
                                                fmt.Println(v5, fi16)
                                                if (2147483647 != 2) {
                                                    var v23 string = "hello"
                                                    var v24 int64 = 539
                                                    _ = v23
                                                    _ = v24
                                                }
                                                var v25 string = ("" + "\n\t")
                                                _ = v25
                                            }
                                            var v26 string = ("" + (("" + "a") + ""))
                                            v4[uint32(v3)] = v26
                                            _ = v26
                                        }
                                        v27 := GenS1{F0: uint32(v3)}
                                        _ = len(v1)
                                        _ = v27
                                    } else {
                                        if false {
                                            v11 -= (v11 * ((0.1 + float64(0.0)) / v11))
                                            fmt.Println(v5, v3, v11)
                                            switch fi10 {
                                            case 127:
                                                var v28 int32 = v5
                                                if false {
                                                    switch (^(fi10)) {
                                                    case 32768:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        switch 32768 {
                                                        case 32768:
                                                            fmt.Println(v6)
                                                        case 9223372036854775807:
                                                            fmt.Println(v6, fi16, fi10)
                                                        default:
                                                            v6 = v6
                                                        }
                                                        v28 = (v5 + 1)
                                                    }
                                                    fmt.Println(fi16, v5)
                                                }
                                                _ = v28
                                            default:
                                                v1 = append(v1, 42)
                                            }
                                        }
                                    }
                                    var v29 int32 = -307
                                    v6 -= 22
                                    _ = v29
                                }
                            }
                            fmt.Println(fi16, fi10, fi10)
                        }
                        _ = fi16
                    }
                    _ = v15
                }
                _ = fi10
                _ = v11
            }
        }
        fmt.Println(v6, v5, v3)
    }
    v1 = append(v1, 1)
    var v30 int = 0
    for fi31 := 0; fi31 < 13; fi31++ {
        v32 := GenS1{F0: 42}
        v33 := GenS1{F0: 42}
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 int = v30
        v5 = v5
        _ = len(v1)
        _ = fi34
        _ = v35
    }
    v36 := GenS2{F0: 4294967295}
    v5 = v5
    v30 += 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
