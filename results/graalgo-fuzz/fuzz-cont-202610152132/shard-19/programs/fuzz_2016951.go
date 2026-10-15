package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 uint32; F2 uint16 }

func main() {
    v1 := []uint32{42, 4294967295}
    var v2 int = 127
    v3 := []uint32{2147483647, 0}
    var v4 uint8 = 127
    v5 := GenS1{F0: 255, F1: -105, F2: 42}
    v2 = v2
    v4 = v4
    for fi6 := 0; fi6 < 20; fi6++ {
        if (true && false) {
            switch (("hello" + "foo bar") + ("a" + "foo bar")) {
            case "foo bar":
                _ = len(v1)
            case "\n\t":
                v7 := GenS1{F0: 42, F1: 1, F2: -32768}
                _ = v7
            case "":
                if false {
                    fmt.Println(v2)
                } else {
                    v8 := GenS1{F0: 1, F1: -73, F2: 436}
                    _ = v8
                }
            }
            v9 := GenS1{F0: 255, F1: -1, F2: 140}
            for fi10 := 0; fi10 < 8; fi10++ {
                v4 -= (v4 >> ((682) & 63))
                if true {
                    var v11 int32 = -94
                    if true {
                        if len(v3) > 0 {
                            v3[0] = 2
                        }
                        v2 = (fi6 / ((v2) | 1))
                        var v12 int8 = 100
                        _ = v12
                    } else {
                        switch "" {
                        case "\n\t":
                            if ((!true) || (float64(((float32(1.5) - (float32(0.0) + (-2.25 * -2.25))) * 1e10)) >= (1.5 / -2.25))) {
                                switch 32767 {
                                case -2147483648:
                                    v4 += v4
                                case 2147483647:
                                    var v13 int = v2
                                    _ = v13
                                }
                                v14 := GenS2{F0: 807, F1: 789}
                                if (v2 < fi10) {
                                    fmt.Println(v2, v4, fi6)
                                }
                                _ = v14
                            } else {
                                v4 = (v4 + v4)
                                var v15 bool = (1e10 == float32(0.1))
                                var v16 bool = false
                                _ = v15
                                _ = v16
                            }
                        case "foo bar":
                            v17 := GenS3{F0: 545, F1: 2}
                            var v18 int32 = (-(v11))
                            _ = v17
                            _ = v18
                        }
                        var v19 uint = 42
                        var v20 uint = (v19 * v19)
                        _ = v19
                        _ = v20
                    }
                    var v21 string = "a"
                    _ = v11
                    _ = v21
                } else {
                    switch ("\n\t" + "hello") {
                    case "foo bar":
                        var v22 uint8 = v4
                        _ = v22
                    default:
                        if false {
                            _ = len(v1)
                            v23 := GenS4{F0: 1, F1: 0, F2: 2}
                            if false {
                                fmt.Println(v2)
                                var v24 int32 = 42
                                _ = v24
                            } else {
                                v4 += v4
                            }
                            _ = v23
                        }
                    }
                    switch "" {
                    case "":
                        v25 := GenS1{F0: 4294967295, F1: 42, F2: -1}
                        _ = v25
                    case "foo bar":
                        var v26 bool = (false && true)
                        _ = v26
                    default:
                        v2 = fi10
                        switch (v2 + 2147483647) {
                        case 127:
                            var v27 int32 = 0
                            _ = v27
                        }
                        fmt.Println(fi10)
                    }
                    if (0.0 > 0.1) {
                        v2 = -578
                        var v28 uint64 = 1
                        if (0.1 > 3.14159) {
                            v4 = v4
                            if ((v4 & (^(v4))) != (v4 & 1)) {
                                fmt.Println(v28)
                            } else {
                                if true {
                                    _ = len(v1)
                                    var v29 int8 = 100
                                    _ = v29
                                }
                                v28 -= (v28 - 1)
                            }
                            if len(v1) > 0 {
                                v1[0] = 243
                            }
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = v28
                    } else {
                        var v30 uint16 = 4
                        if (-653 <= 2147483647) {
                            switch v2 {
                            case 0:
                                v4 += 2
                            case 42:
                                var v31 float64 = 3.14159
                                _ = v31
                            case 128:
                                v4 = v4
                            }
                            v32 := GenS5{F0: 65535}
                            if true {
                                if len(v3) > 0 {
                                    v3[0] = 393
                                }
                                fmt.Println(fi10, v30, fi10)
                                var v33 uint64 = 4294967295
                                _ = v33
                            }
                            _ = v32
                        } else {
                            var v34 float32 = (0.1 - (3.14159 * (((float32(0.1) * float32(0.1)) / float32(float64((-2.25 / 1.0)))) - 1.5)))
                            if (!false) {
                                _ = len(v3)
                                v4 -= v4
                            } else {
                                v34 += 1.5
                            }
                            fmt.Println(fi10, v30, v30)
                            _ = v34
                        }
                        _ = v30
                    }
                }
                _ = len(v3)
                _ = fi10
            }
            _ = v9
        } else {
            v2 -= (fi6 % ((128) | 1))
            if true {
                if false {
                    switch (fi6 & (v2 % ((fi6) | 1))) {
                    case -2147483648:
                        var v35 int64 = 9223372036854775807
                        _ = v35
                    case 2147483647:
                        for fi36 := 0; fi36 < 11; fi36++ {
                            switch 32768 {
                            case 0:
                                if false {
                                    if ("hello" <= "foo bar") {
                                        v2 = (fi6 * (v2 | fi6))
                                        var v37 int = fi6
                                        v38 := GenS2{F0: 0, F1: 255}
                                        _ = v37
                                        _ = v38
                                    } else {
                                        v4 = v4
                                    }
                                    var v39 string = "hello"
                                    fmt.Println(v39)
                                    v1 = append(v1, 65535)
                                    _ = v39
                                }
                            }
                            _ = fi36
                        }
                    default:
                        v2 += (fi6 - v2)
                    }
                    v2 = (v2 - v2)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v4 += (v4 | 127)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            v40 := GenS3{F0: 42, F1: 65535}
            _ = v40
        }
        if true {
            switch 32767 {
            case -1:
                switch ("" + (("\n\t" + "foo bar") + "")) {
                case "foo bar":
                    v4 = 0
                case "hello":
                    v2 = v2
                case "":
                    for fi41 := 0; fi41 < 6; fi41++ {
                        fmt.Println(fi6)
                        v4 += (v4 ^ v4)
                        _ = fi41
                    }
                }
            case 506:
                v42 := GenS1{F0: 715, F1: 100, F2: -446}
                _ = v42
            default:
                var v43 float64 = (float64(float32((1e10 - 3.14159))) / 1.0)
                _ = v43
            }
        }
        _ = fi6
    }
    v2 -= int(v4)
    v2 = (v2 & v2)
    switch (("hello" + "hello") + "") {
    case "":
        for fi44 := 0; fi44 < 17; fi44++ {
            if ((!false) && true) {
                v4 -= v4
            }
            if false {
                if (fi44 != (fi44 % ((27) | 1))) {
                    if (18446744073709551615 >= 4294967295) {
                        v2 = (v2 + (-(fi44)))
                        v4 -= v4
                        v2 -= 0
                    }
                    var v45 int8 = -1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v45
                }
                if true {
                    v1 = append(v1, 4294967295)
                    v4 += v4
                } else {
                    v46 := GenS6{F0: 1}
                    if false {
                        fmt.Println(fi44, fi44, v2)
                        fmt.Println(v4, v2, fi44)
                        switch ("foo bar" + "") {
                        case "foo bar":
                            v4 -= v4
                        default:
                            for fi47 := 0; fi47 < 19; fi47++ {
                                v2 = 503
                                v48 := GenS3{F0: 2, F1: 65535}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi47
                                _ = v48
                            }
                        }
                    } else {
                        var v49 int32 = 2147483647
                        v2 -= 32767
                        _ = v49
                    }
                    var v50 uint16 = 2
                    _ = v46
                    _ = v50
                }
                if false {
                    var v51 uint = 4294967295
                    _ = v51
                } else {
                    switch fi44 {
                    case -1:
                        fmt.Println(v2, v4, fi44)
                    }
                    if false {
                        var v52 uint64 = 4294967295
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v52
                    }
                    v53 := GenS1{F0: 886, F1: 100, F2: 32767}
                    _ = v53
                }
            } else {
                var v54 uint16 = 255
                v4 = (^(v4))
                for fi55 := 0; fi55 < 6; fi55++ {
                    v4 -= v4
                    var v56 uint = 4294967295
                    fmt.Println(fi44, v2)
                    _ = fi55
                    _ = v56
                }
                _ = v54
            }
            _ = fi44
        }
    case "hello":
        if true {
            v2 -= 32768
            v57 := GenS7{F0: 859, F1: 4294967295, F2: 775}
            _ = v57
        } else {
            var v58 int16 = 1000
            _ = v58
        }
    case "foo bar":
        v4 = 0
    }
    v3 = append(v3, 65535)
    var v59 int16 = 750
    var v60 int32 = -286
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v60
}
