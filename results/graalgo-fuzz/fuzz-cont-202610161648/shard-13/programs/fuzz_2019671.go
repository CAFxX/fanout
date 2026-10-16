package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int }
type GenS2 struct { F0 uint64 }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := []int{fi5, (^(fi5))}
        _ = fi5
        _ = v6
    }
    if (!(1e10 < 1.5)) {
        if (!(-1 <= 0)) {
            v7 := []int{}
            _ = v7
        }
        v8 := GenS1{F0: -2147483648, F1: ("foo bar" + "a"), F2: 1}
        _ = v8
    }
    return 128
}
func genFn1() string {
    v9 := GenS2{F0: 4294967295}
    switch -1 {
    case 275:
        v10 := GenS2{F0: 4294967295}
        _ = v10
    case -2147483648:
        var v11 uint16 = 309
        _ = v11
    default:
        var v12 uint64 = 18446744073709551615
        _ = v12
    }
    if (100 < -128) {
        if true {
            var v13 uint32 = 1
            _ = v13
        }
    }
    v14 := GenS1{F0: -102, F1: "foo bar", F2: -2147483648}
    _ = v14
    _ = v9
    return "a"
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]string)
    var v3 int16 = 1000
    v4 := []int{-709, 32768}
    fmt.Println(v1, v3)
    var v15 int = -385
    var v16 string = ""
    fmt.Println(v1, v3, v15)
    switch v15 {
    case 32767:
        switch "foo bar" {
        case "foo bar":
            var v17 int32 = 0
            _ = v17
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            v1 -= 18446744073709551615
        }
    case 32768:
        v15 -= (-(v15))
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        if false {
            if false {
                v1 += v1
                if true {
                    var v19 int = 2147483647
                    v16 = v16
                    _ = v19
                } else {
                    var v20 uint = 35
                    for fi21 := 0; fi21 < 12; fi21++ {
                        switch fi21 {
                        case -1:
                            if (false || true) {
                                var v22 string = v16
                                if ((((1.5 * -2.25) * (0.0 / 3.14159)) * float32(((((1e10 + 1e10) + (((3.14159 / (0.1 / 0.1)) / 1.5) * (float64(0.0) - -2.25))) + -2.25) + 1e10))) <= (3.14159 - 3.14159)) {
                                    if (679 < -1) {
                                        v23 := GenS1{F0: 1, F1: (v16 + v22), F2: 244}
                                        v22 = v16
                                        _ = v23
                                    } else {
                                        v16 = "hello"
                                        v16 = (v16 + "foo bar")
                                        var v24 uint32 = 2147483648
                                        _ = v24
                                    }
                                    fmt.Println(v3, fi21)
                                    v16 = v16
                                }
                                _ = v22
                            }
                        case 42:
                            v25 := v2[int(v20)]
                            _ = v25
                        default:
                            switch (fi18 * (fi21 >> ((v20) & 63))) {
                            case 836:
                                switch 9223372036854775807 {
                                case -2147483648:
                                    var v26 int32 = 0
                                    _ = v26
                                default:
                                    v15 += (fi21 % (((fi21 + fi18)) | 1))
                                }
                            case 885:
                                v20 -= uint(v1)
                            case 4294967295:
                                var v27 int64 = 0
                                _ = v27
                            default:
                                _ = len(v4)
                            }
                        }
                        v1 = 42
                        _ = fi21
                    }
                    var v28 uint8 = 42
                    _ = v20
                    _ = v28
                }
            } else {
                v16 = v16
                if false {
                    var v29 string = "hello"
                    for fi30 := 0; fi30 < 6; fi30++ {
                        if false {
                            if true {
                                var v31 uint8 = 1
                                _ = v31
                            }
                        }
                        switch fi18 {
                        case 2147483647:
                            fmt.Println(v16)
                        }
                        var v32 int = (v15 << ((1) & 63))
                        _ = fi30
                        _ = v32
                    }
                    _ = v29
                } else {
                    v1 = (v1 >> ((2) & 63))
                    v15 = fi18
                }
            }
            v1 += (v1 << ((65535) & 63))
        } else {
            switch v16 {
            case "hello":
                if false {
                    v15 += (v15 & (v15 / (((v15 & (^(v15)))) | 1)))
                    var v33 int8 = 0
                    _ = v33
                }
            case "\n\t":
                v15 = (v15 / ((fi18) | 1))
            case "foo bar":
                var v34 bool = (v3 >= -555)
                _ = v34
            }
            if (!false) {
                var v35 uint16 = 256
                v35 -= 669
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v35
            }
            var v36 bool = true
            _ = v36
        }
        v37 := GenS1{F0: 666, F1: v16, F2: 716}
        v4 = append(v4, -497)
        _ = fi18
        _ = v37
    }
    v1 -= v1
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
