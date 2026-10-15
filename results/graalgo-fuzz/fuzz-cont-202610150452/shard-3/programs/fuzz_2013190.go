package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 string; F1 int32; F2 int32 }
type GenS6 struct { F0 uint32; F1 string }
func genFn0() uint32 {
    var v7 uint8 = 128
    var v8 uint32 = 4294967295
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 int8 = 73
    v2 := []uint32{2, 0, 2147483647, 2147483648}
    var v3 int64 = -1
    v4 := []int{42, -829}
    v5 := make(map[string]int)
    var v6 string = "\n\t"
    var v9 uint32 = 932
    for fi10 := 0; fi10 < 7; fi10++ {
        switch (fi10 >> ((381) & 63)) {
        case 32768:
            v11 := GenS1{F0: 42}
            _ = v11
        default:
            v6 = v6
        }
        fmt.Println(v9)
        fmt.Println(fi10)
        _ = fi10
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 float32 = -2.25
        var v14 int32 = 2147483647
        v3 = (v3 << ((190) & 63))
        _ = fi12
        _ = v13
        _ = v14
    }
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    v15 := v5[v6]
    v1 = 0
    var v16 uint = 18446744073709551615
    for fi17 := 0; fi17 < 14; fi17++ {
        v1 = 1
        if (2 < 1) {
            switch (v15 - (fi17 % ((fi17) | 1))) {
            case 42:
                v18 := GenS2{F0: false, F1: v16, F2: v1}
                _ = v18
            case 32767:
                var v19 uint16 = uint16(v9)
                _ = v19
            case 543:
                v15 += (fi17 << (((^(v16))) & 63))
            }
            if (285 > (v16 >> ((v16) & 63))) {
                var v20 int = (-(fi17))
                if false {
                    for fi21 := 0; fi21 < 19; fi21++ {
                        fmt.Println(v3, v3)
                        var v22 float64 = -2.25
                        var v23 string = v6
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS1{F0: 255}
                    _ = v24
                }
                v25 := GenS2{F0: (273 > 32767), F1: v16, F2: 49}
                _ = v20
                _ = v25
            } else {
                switch v6 {
                case "\n\t":
                    v26 := v5[""]
                    _ = v26
                default:
                    if true {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v1)
                        case "hello":
                            v27 := GenS3{F0: 102}
                            _ = v27
                        case "":
                            v9 -= v9
                        default:
                            switch ("a" + v6) {
                            case "\n\t":
                                fmt.Println(v9)
                            case "":
                                var v28 int64 = v3
                                _ = v28
                            case "foo bar":
                                v6 = v6
                            }
                        }
                    } else {
                        var v29 int = v15
                        if (int64(v16) <= (^(v3))) {
                            for fi30 := 0; fi30 < 2; fi30++ {
                                var v31 int = (v15 - -9223372036854775808)
                                if ((-2.25 * -2.25) > -2.25) {
                                    if true {
                                        v32 := GenS4{F0: (-(v1)), F1: v15}
                                        _ = v32
                                    } else {
                                        var v33 uint16 = 451
                                        _ = v33
                                    }
                                    if true {
                                        switch (("a" + "hello") + "") {
                                        case "foo bar":
                                            var v34 bool = (!(((((0.0 * 0.1) / 1.5) * (0.0 + ((3.14159 - 1.5) * 1.5))) - 1.5) == -2.25))
                                            _ = v34
                                        case "":
                                            var v35 uint16 = 957
                                            _ = v35
                                        case "hello":
                                            _ = len(v5)
                                        default:
                                            var v36 int64 = (v3 + (v3 << (((^(v16))) & 63)))
                                            _ = v36
                                        }
                                        v3 += (^(v3))
                                        v1 = (v1 & (^(v1)))
                                    }
                                } else {
                                    if false {
                                        var v37 uint32 = 0
                                        if false {
                                            _ = len(v5)
                                            var v38 uint16 = 42
                                            if false {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v39 uint64 = 0
                                                _ = v39
                                            }
                                            if (0 != 2) {
                                                v40 := GenS4{F0: v1, F1: fi17}
                                                if true {
                                                    var v41 uint8 = 128
                                                    var v42 float64 = ((0.1 / 3.14159) + 1e10)
                                                    v38 -= v38
                                                    _ = v41
                                                    _ = v42
                                                } else {
                                                    v5["\n\t"] = v31
                                                    var v43 int16 = -369
                                                    v44 := v5[v6]
                                                    _ = v43
                                                    _ = v44
                                                }
                                                v29 -= (fi30 | (v15 - (fi30 * (fi17 * -188))))
                                                fmt.Println(v3, v1, v31)
                                                _ = v40
                                            }
                                            v31 = (v31 * v15)
                                            _ = v38
                                        }
                                        switch v6 {
                                        case "foo bar":
                                            if (!(v37 == 42)) {
                                                fmt.Println(v15, v29, v3)
                                            } else {
                                                var v45 int8 = v1
                                                _ = v45
                                            }
                                        case "\n\t":
                                            if (v6 >= v6) {
                                                switch "foo bar" {
                                                case "a":
                                                    fmt.Println(v15)
                                                case "\n\t":
                                                    if false {
                                                        if (!true) {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                var v46 uint8 = 1
                                                                var v47 uint = (v16 ^ v16)
                                                                _ = v46
                                                                _ = v47
                                                            case "a":
                                                                var v48 float32 = 1.5
                                                                _ = v48
                                                            default:
                                                                v1 += -66
                                                            }
                                                            switch ("foo bar" + v6) {
                                                            case "":
                                                                v29 = (fi17 >> ((v16) & 63))
                                                            }
                                                        }
                                                    } else {
                                                        v49 := GenS5{F0: (v6 + ("" + "foo bar")), F1: 42, F2: 1}
                                                        _ = v49
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v15, v37, v16)
                                                v50 := GenS6{F0: 4294967295, F1: v6}
                                                v5[((("\n\t" + "") + v6) + (((v6 + (v6 + "")) + v6) + "a"))] = fi17
                                                _ = v50
                                            }
                                        default:
                                            var v51 uint = (^(v16))
                                            _ = v51
                                        }
                                        _ = v37
                                    } else {
                                        fmt.Println(v29, v1, v1)
                                    }
                                }
                                _ = fi30
                                _ = v31
                            }
                            _ = len(v2)
                            if false {
                                if ((!false) && true) {
                                    var v52 uint64 = 267
                                    _ = v52
                                }
                            } else {
                                for fi53 := 0; fi53 < 6; fi53++ {
                                    if true {
                                        v9 -= 4294967295
                                    } else {
                                        var v54 float64 = 3.14159
                                        fmt.Println(fi17, v16)
                                        _ = v54
                                    }
                                    switch v15 {
                                    case 250:
                                        if (true || true) {
                                            v6 = v6
                                            var v55 bool = (!(-5 == v3))
                                            _ = v55
                                        }
                                    case 42:
                                        var v56 string = v6
                                        _ = v56
                                    case -612:
                                        var v57 uint16 = 109
                                        _ = v57
                                    default:
                                        if len(v4) > 0 {
                                            v4[0] = (^(v15))
                                        }
                                    }
                                    var v58 string = v6
                                    var v59 string = (v6 + v58)
                                    _ = fi53
                                    _ = v58
                                    _ = v59
                                }
                                v2 = append(v2, v9)
                            }
                        } else {
                            for fi60 := 0; fi60 < 14; fi60++ {
                                v4 = append(v4, (v29 * -2147483648))
                                var v61 int64 = v3
                                fmt.Println(v3, v61)
                                var v62 uint = (^(v16))
                                switch (^(v29)) {
                                case 4294967295:
                                    var v63 string = v6
                                    _ = v63
                                case 32767:
                                    v9 += 1
                                default:
                                    v62 -= v16
                                }
                                _ = fi60
                                _ = v61
                                _ = v62
                            }
                            v64 := GenS4{F0: v1, F1: 863}
                            _ = v64
                        }
                        var v65 string = (("hello" + v6) + v6)
                        _ = v29
                        _ = v65
                    }
                }
            }
            var v66 bool = false
            _ = v66
        }
        switch v6 {
        case "\n\t":
            _ = len(v5)
        case "a":
            v15 = (fi17 + 42)
        case "":
            var v67 float32 = float32(-2.25)
            _ = v67
        default:
            _ = len(v5)
        }
        _ = fi17
    }
    var v68 uint = 42
    v3 -= (v3 ^ v3)
    fmt.Println(v68, v3)
    for fi69 := 0; fi69 < 20; fi69++ {
        switch fi69 {
        case 32768:
            for fi70 := 0; fi70 < 10; fi70++ {
                fmt.Println(v1)
                switch "a" {
                case "foo bar":
                    v3 += v3
                }
                _ = fi70
            }
        case 1:
            v3 = v3
        default:
            var v71 float64 = float64(3.14159)
            _ = v71
        }
        v5["\n\t"] = (v15 & 9223372036854775807)
        v5[v6] = (v15 + (-(fi69)))
        _ = fi69
    }
    fmt.Println(v6)
    var v72 bool = true
    var v73 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v72
    _ = v73
    _ = v9
}
