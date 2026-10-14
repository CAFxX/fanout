package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int8 }

func main() {
    v1 := []uint32{65535, 2147483647, 543}
    v2 := []int{547, 9223372036854775807, 962, 2147483647}
    var v3 string = "hello"
    var v4 int64 = 0
    if true {
        var v5 float32 = 0.1
        var v6 uint16 = 255
        var v7 int64 = v4
        _ = v5
        _ = v6
        _ = v7
    } else {
        var v8 uint = 800
        var v9 int8 = 1
        _ = v8
        _ = v9
    }
    var v10 int32 = 42
    fmt.Println(v10, v3, v10)
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 bool = (432 >= 2147483647)
            _ = fi12
            _ = v13
        }
        var v14 string = "a"
        v15 := GenS1{F0: v10}
        _ = fi11
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        v2 = append(v2, 9223372036854775807)
        _ = len(v2)
        if (v4 < (-(v4))) {
            v10 = (v10 & 1)
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 bool = ((v3 + (v3 + v3)) == v3)
                v10 += v10
                v3 = ((v3 + v3) + "a")
                _ = fi17
                _ = v18
            }
        } else {
            if (!true) {
                if false {
                    if true {
                        v3 = v3
                        var v19 uint8 = 128
                        _ = v19
                    } else {
                        v10 += v10
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch fi16 {
                    case -2147483648:
                        for fi20 := 0; fi20 < 4; fi20++ {
                            v10 -= -235
                            var v21 int32 = v10
                            var v22 uint8 = 127
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    default:
                        for fi23 := 0; fi23 < 13; fi23++ {
                            var v24 int32 = -1
                            if len(v1) > 0 {
                                v1[0] = 0
                            }
                            var v25 float32 = (0.1 / (-2.25 * (3.14159 - float32(float64(1e10)))))
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    }
                }
            }
        }
        _ = fi16
    }
    v4 -= (v4 * (v4 * (v4 | 9223372036854775807)))
    var v26 uint8 = 59
    v26 = v26
    v4 = v4
    _ = len(v1)
    switch ((("foo bar" + v3) + "hello") + (v3 + (v3 + v3))) {
    case "":
        for fi27 := 0; fi27 < 1; fi27++ {
            if ((false || (!true)) || true) {
                v26 = 2
            } else {
                if false {
                    _ = len(v1)
                } else {
                    var v28 string = v3
                    if len(v2) > 0 {
                        v2[0] = int(v10)
                    }
                    v10 += (v10 / ((0) | 1))
                    _ = v28
                }
                _ = len(v2)
            }
            v4 -= (v4 - (v4 * 0))
            for fi29 := 0; fi29 < 10; fi29++ {
                if false {
                    if false {
                        var v30 float32 = 0.1
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v30
                    }
                    v4 -= 0
                } else {
                    if true {
                        var v31 int32 = v10
                        _ = v31
                    } else {
                        v4 += int64(v26)
                        switch 32768 {
                        case -16:
                            switch -2147483648 {
                            case -308:
                                var v32 float32 = (float32(-2.25) + 1.5)
                                _ = v32
                            }
                        case -1:
                            v26 += v26
                            v2 = append(v2, 32768)
                            v3 = "foo bar"
                        }
                    }
                    switch "a" {
                    case "":
                        fmt.Println(v3, v4)
                    default:
                        switch v3 {
                        case "foo bar":
                            v3 = "\n\t"
                        case "a":
                            if (!(!true)) {
                                v33 := GenS1{F0: v10}
                                _ = v33
                            }
                        default:
                            v34 := GenS2{F0: v26}
                            _ = v34
                        }
                    }
                    var v35 int64 = -245
                    if (v10 == v10) {
                        var v36 uint32 = 274
                        if (46 < 4294967295) {
                            if false {
                                var v37 int8 = 1
                                _ = v37
                            } else {
                                v10 += (^(v10))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            v1 = append(v1, (^(v36)))
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            v36 += 2
                            switch "a" {
                            case "hello":
                                var v38 uint64 = 1
                                _ = v38
                            }
                        }
                        if (v36 >= 1) {
                            var v39 int32 = (v10 % (((v10 + v10)) | 1))
                            v40 := GenS1{F0: v39}
                            var v41 float64 = -2.25
                            if (2147483647 == 327) {
                                var v42 uint = 18446744073709551615
                                _ = v42
                            } else {
                                v36 = (v36 & (v36 ^ 2147483648))
                                fmt.Println(v10, v4)
                                var v43 int8 = 1
                                if (!(v10 >= v10)) {
                                    fmt.Println(v35, fi29)
                                } else {
                                    v3 = v3
                                    v41 = -2.25
                                }
                                _ = len(v2)
                                var v44 float64 = v41
                                if false {
                                    fmt.Println(v10, v26, v10)
                                    v43 -= (v43 & (v43 << ((0) & 63)))
                                } else {
                                    v45 := GenS1{F0: v10}
                                    switch fi29 {
                                    case 42:
                                        v46 := GenS3{F0: fi29, F1: (v43 >> ((393) & 63))}
                                        _ = v46
                                    case -9223372036854775808:
                                        fmt.Println(v26)
                                    default:
                                        if (uint16(v26) != 1) {
                                            var v47 int = (fi27 ^ fi29)
                                            _ = v47
                                        } else {
                                            var v48 uint16 = 255
                                            _ = v48
                                        }
                                    }
                                    var v49 int64 = (v4 ^ (v4 * (^(v35))))
                                    _ = v45
                                    _ = v49
                                }
                                var v50 int64 = (v4 ^ v35)
                                _ = v43
                                _ = v44
                                _ = v50
                            }
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        _ = v36
                    }
                    var v51 string = (v3 + (("foo bar" + v3) + ((v3 + "") + (v3 + "hello"))))
                    _ = v35
                    _ = v51
                }
                _ = fi29
            }
            _ = fi27
        }
    case "a":
        v4 = int64(v26)
    }
    if true {
        switch "\n\t" {
        case "\n\t":
            var v52 int32 = v10
            _ = v52
        case "a":
            v2 = append(v2, -421)
        case "hello":
            v53 := GenS3{F0: -1, F1: 1}
            _ = v53
        default:
            var v54 int64 = (v4 << ((0) & 63))
            _ = v54
        }
    } else {
        for fi55 := 0; fi55 < 17; fi55++ {
            v10 += v10
            fmt.Println(fi55)
            _ = fi55
        }
    }
    var v56 uint8 = 139
    v4 -= 457
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v56
}
