package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int32; F2 bool }
type GenS3 struct { F0 uint; F1 bool; F2 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS8 struct { F0 int8; F1 int16 }
type GenS9 struct { F0 int; F1 int8; F2 uint32 }
type GenS10 struct { F0 uint32; F1 uint32 }
type GenS11 struct { F0 int8; F1 uint16 }
type GenS12 struct { F0 int64; F1 int }
type GenS13 struct { F0 int32; F1 uint16 }
func genFn0() uint8 {
    v5 := []int{865, -2147483648, -1, -2147483648}
    var v6 uint = 2
    v6 = (v6 << (((v6 ^ 255)) & 63))
    var v7 uint = v6
    _ = len(v5)
    _ = v5
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 uint8 = 42
    v2 := []int{230}
    var v3 int32 = -1
    var v4 bool = false
    var v8 int32 = v3
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        fmt.Println(v1, v8, v3)
    default:
        var v9 float32 = (-2.25 * 3.14159)
        _ = v9
    }
    fmt.Println(v4, v8)
    v10 := GenS1{F0: v3}
    var v11 bool = v4
    v12 := GenS1{F0: (v3 / (((v8 / (((v8 << ((65535) & 63))) | 1))) | 1))}
    if len(v2) > 0 {
        v2[0] = 449
    }
    v11 = false
    v13 := GenS1{F0: 1}
    switch 344 {
    case 888:
        switch "hello" {
        case "\n\t":
            var v14 float64 = ((((3.14159 / (0.1 - (0.0 - 3.14159))) + 3.14159) * 0.0) + (0.1 + 0.0))
            _ = v14
        case "":
            var v15 uint64 = 1
            _ = v15
        default:
            for fi16 := 0; fi16 < 18; fi16++ {
                if (!(1 > 0)) {
                    _ = len(v2)
                    if false {
                        v4 = v11
                        v1 = (v1 / ((genFn0()) | 1))
                    } else {
                        for fi17 := 0; fi17 < 6; fi17++ {
                            v2 = append(v2, fi17)
                            if (0 == 2147483647) {
                                var v18 int8 = 127
                                var v19 string = "\n\t"
                                var v20 uint8 = 255
                                var v21 int16 = -1
                                if true {
                                    v22 := GenS2{F0: 255, F1: 0, F2: (!false)}
                                    var v23 bool = (!v11)
                                    _ = v22
                                    _ = v23
                                }
                                v1 = (v1 % ((0) | 1))
                                if v4 {
                                    v20 += (v20 / (((v20 / (((v20 - 42)) | 1))) | 1))
                                    var v24 uint8 = v20
                                    var v25 uint = 18446744073709551615
                                    switch (v19 + "\n\t") {
                                    case "hello":
                                        v2 = append(v2, (fi17 & (fi17 * 32768)))
                                    case "\n\t":
                                        v2 = append(v2, int(v3))
                                    }
                                    if v4 {
                                        v2 = append(v2, (-(fi17)))
                                        fmt.Println(v3, v1, v19)
                                        switch fi17 {
                                        case -9223372036854775808:
                                            v19 = (v19 + v19)
                                        }
                                        if v11 {
                                            switch (fi17 * fi17) {
                                            case -9223372036854775808:
                                                if len(v2) > 0 {
                                                    v2[0] = fi16
                                                }
                                            case -959:
                                                var v26 float64 = 1e10
                                                _ = v26
                                            default:
                                                if v11 {
                                                    v27 := GenS3{F0: v25, F1: v11, F2: v19}
                                                    v28 := GenS4{F0: genFn0()}
                                                    fmt.Println(v1)
                                                    if ((fi17 / ((42) | 1)) > (fi16 | (^(fi17)))) {
                                                        v29 := GenS3{F0: 320, F1: false, F2: v19}
                                                        _ = v29
                                                    } else {
                                                        v30 := GenS3{F0: (v25 >> ((v25) & 63)), F1: (!v4), F2: "foo bar"}
                                                        fmt.Println(v20, v24)
                                                        _ = v30
                                                    }
                                                    v31 := GenS5{F0: 42}
                                                    _ = v27
                                                    _ = v28
                                                    _ = v31
                                                }
                                            }
                                        } else {
                                            var v32 float32 = 0.0
                                            v33 := GenS2{F0: 0, F1: (v8 - 2147483647), F2: ((v8 - (v3 % (((v8 & (v3 ^ (^(v8))))) | 1))) != 316)}
                                            if len(v2) > 0 {
                                                v2[0] = (fi16 >> ((v25) & 63))
                                            }
                                            _ = v32
                                            _ = v33
                                        }
                                        var v34 int8 = 51
                                        _ = v34
                                    } else {
                                        v21 = v21
                                        var v35 int8 = -20
                                        var v36 uint = v25
                                        _ = v35
                                        _ = v36
                                    }
                                    var v37 uint8 = (v24 - (^(v1)))
                                    v19 = v19
                                    _ = v24
                                    _ = v25
                                    _ = v37
                                } else {
                                    v21 += (v21 | v21)
                                    switch ("" + v19) {
                                    case "hello":
                                        v8 += (v3 & v8)
                                    case "foo bar":
                                        var v38 uint16 = 65535
                                        _ = v38
                                    default:
                                        if true {
                                            if v11 {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v39 := GenS6{F0: (v11 || true), F1: 0}
                                                _ = v39
                                            } else {
                                                v40 := GenS3{F0: 294, F1: (v11 || v4), F2: (((v19 + "") + "a") + ("hello" + "a"))}
                                                _ = len(v2)
                                                switch ((v19 + v19) + "\n\t") {
                                                case "hello":
                                                    v20 = v1
                                                case "":
                                                    fmt.Println(v3, v11)
                                                }
                                                _ = v40
                                            }
                                            if (-668 <= 9223372036854775807) {
                                                v41 := GenS7{F0: v11, F1: 973, F2: 42}
                                                _ = v41
                                            } else {
                                                var v42 uint8 = (^(v20))
                                                if false {
                                                    v43 := GenS1{F0: v3}
                                                    var v44 uint64 = 564
                                                    v2 = append(v2, 4294967295)
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    var v45 float32 = 3.14159
                                                    fmt.Println(v45, v20)
                                                    _ = v45
                                                }
                                                var v46 bool = v4
                                                if v46 {
                                                    if ("a" <= v19) {
                                                        fmt.Println(v46, v8)
                                                        if v11 {
                                                            v3 -= v3
                                                            if (((^(fi17)) != fi16) && false) {
                                                                v47 := GenS2{F0: 255, F1: -586, F2: (!("a" == (v19 + "\n\t")))}
                                                                _ = v47
                                                            }
                                                            var v48 int8 = 127
                                                            _ = v48
                                                        } else {
                                                            v49 := GenS8{F0: 42, F1: v21}
                                                            fmt.Println(fi16)
                                                            _ = v49
                                                        }
                                                        v2 = append(v2, (fi17 & fi16))
                                                    } else {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                }
                                                switch (fi16 & fi17) {
                                                case 128:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                default:
                                                    switch (-(fi17)) {
                                                    case 1:
                                                        if (613 > 65535) {
                                                            v18 += (v18 >> ((255) & 63))
                                                            fmt.Println(v18)
                                                        }
                                                    case 2147483647:
                                                        if (!(1e10 == (0.1 - (((1.5 - 3.14159) + ((3.14159 / (0.1 / 1e10)) * -2.25)) + 1.5)))) {
                                                            var v50 uint16 = 0
                                                            v11 = v46
                                                            var v51 int16 = 32767
                                                            _ = v50
                                                            _ = v51
                                                        } else {
                                                            if (true && (v1 < (v1 | 0))) {
                                                                fmt.Println(v3)
                                                                v3 = 236
                                                            }
                                                        }
                                                    case 9223372036854775807:
                                                        var v52 int8 = 0
                                                        _ = v52
                                                    default:
                                                        var v53 bool = (v4 && true)
                                                        _ = v53
                                                    }
                                                }
                                                _ = v42
                                                _ = v46
                                            }
                                        } else {
                                            var v54 uint64 = 18446744073709551615
                                            _ = v54
                                        }
                                    }
                                }
                                _ = v18
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            v55 := GenS9{F0: fi17, F1: 127, F2: 42}
                            _ = fi17
                            _ = v55
                        }
                        v8 += v3
                    }
                    v8 = v3
                } else {
                    v56 := GenS10{F0: 2, F1: 2147483648}
                    v57 := GenS11{F0: -128, F1: 462}
                    var v58 int16 = -32768
                    _ = v56
                    _ = v57
                    _ = v58
                }
                for fi59 := 0; fi59 < 12; fi59++ {
                    fmt.Println(v11, v4)
                    v60 := GenS12{F0: -9223372036854775808, F1: fi59}
                    _ = fi59
                    _ = v60
                }
                v61 := GenS4{F0: v1}
                _ = fi16
                _ = v61
            }
        }
    case 128:
        switch 2147483647 {
        case 0:
            v62 := GenS11{F0: -125, F1: 661}
            _ = v62
        case 2147483647:
            fmt.Println(v3, v1, v3)
        case 1:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v63 := GenS13{F0: -2147483648, F1: 405}
            _ = v63
        }
    case 9223372036854775807:
        var v64 uint64 = 2
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
