package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 uint }
type GenS8 struct { F0 uint16; F1 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16; F1 int16; F2 bool }
type GenS11 struct { F0 int; F1 string }
type GenS12 struct { F0 int16 }

func main() {
    var v1 uint8 = 2
    v2 := []uint32{2147483648, 42, 0, 2147483647}
    var v3 string = ""
    v4 := make(map[uint32]bool)
    var v5 float64 = -2.25
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            switch fi7 {
            case -415:
                if false {
                    switch ("\n\t" + "a") {
                    case "":
                        v8 := GenS1{F0: 100}
                        _ = v8
                    case "hello":
                        if ((-2.25 - (0.0 + (0.0 * 1e10))) != (float32((v5 * v5)) * -2.25)) {
                            switch -1 {
                            case 902:
                                var v9 int8 = 127
                                var v10 bool = true
                                _ = v10
                                _ = v9
                            case 9223372036854775807:
                                var v11 int = 0
                                _ = v11
                            case 4294967295:
                                if (-487 < 42) {
                                    v12 := GenS2{F0: 293, F1: 497}
                                    v13 := GenS2{F0: 18446744073709551615, F1: uint16(v1)}
                                    if false {
                                        v1 -= (v1 - v1)
                                        v5 = v5
                                    } else {
                                        v5 += 0.0
                                        v5 += v5
                                        v14 := GenS3{F0: 127, F1: 863}
                                        _ = v14
                                    }
                                    _ = v12
                                    _ = v13
                                }
                            }
                            v15 := v4[422]
                            _ = v15
                        } else {
                            if true {
                                v16 := GenS4{F0: false, F1: 2}
                                _ = len(v4)
                                _ = v16
                            } else {
                                if true {
                                    v5 = (v5 - (v5 * v5))
                                    _ = len(v2)
                                    var v17 int16 = 769
                                    _ = v17
                                } else {
                                    switch ((v3 + v3) + "foo bar") {
                                    case "a":
                                        var v18 uint32 = 586
                                        _ = v18
                                    case "foo bar":
                                        var v19 float32 = 0.1
                                        _ = v19
                                    case "\n\t":
                                        if ((1e10 + (3.14159 / (-2.25 + 1e10))) != 0.0) {
                                            if (false && true) {
                                                _ = len(v2)
                                                fmt.Println(fi7, v5)
                                                v20 := GenS5{F0: (^(fi6)), F1: -2147483648}
                                                _ = v20
                                            }
                                        }
                                    }
                                    fmt.Println(fi7)
                                    fmt.Println(v5, fi6)
                                }
                                v21 := GenS4{F0: true, F1: 2147483647}
                                _ = v21
                            }
                            v22 := GenS3{F0: 48, F1: 1}
                            v1 = 1
                            _ = v22
                        }
                    case "foo bar":
                        var v23 float64 = v5
                        _ = v23
                    }
                } else {
                    switch fi7 {
                    case 235:
                        var v24 bool = true
                        _ = v24
                    }
                    if (true && true) {
                        if (!("hello" <= ((v3 + (v3 + v3)) + v3))) {
                            if (false || true) {
                                if (false || (-112 >= 42)) {
                                    var v25 int32 = 1
                                    v26 := GenS6{F0: (v25 * v25)}
                                    _ = v25
                                    _ = v26
                                } else {
                                    var v27 uint32 = 4294967295
                                    var v28 int8 = -72
                                    v1 -= v1
                                    _ = v27
                                    _ = v28
                                }
                                switch v3 {
                                case "foo bar":
                                    v1 = v1
                                default:
                                    fmt.Println(v3, v3, v1)
                                }
                                v1 = (^(v1))
                            } else {
                                v5 = v5
                            }
                        } else {
                            var v29 int64 = 0
                            _ = v29
                        }
                        switch fi7 {
                        case -160:
                            if (-46 > 1) {
                                if false {
                                    v1 += v1
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = len(v4)
                                    if true {
                                        _ = len(v4)
                                        if false {
                                            switch 42 {
                                            case -9223372036854775808:
                                                var v30 uint8 = (v1 >> ((255) & 63))
                                                _ = v30
                                            }
                                            v1 += (v1 >> ((42) & 63))
                                        }
                                    }
                                    var v31 uint8 = (v1 + (v1 * v1))
                                    _ = v31
                                } else {
                                    v32 := v4[272]
                                    var v33 bool = v32
                                    if false {
                                        v33 = (1e10 >= 0.0)
                                        v4[42] = false
                                        v32 = true
                                    } else {
                                        fmt.Println(v5)
                                        fmt.Println(v3)
                                        if (!true) {
                                            var v34 int16 = -1
                                            _ = v34
                                        }
                                    }
                                    _ = v32
                                    _ = v33
                                }
                            }
                        case 1:
                            delete(v4, 2)
                        }
                        var v35 uint = 2
                        v35 -= v35
                        var v36 uint8 = (v1 / ((v1) | 1))
                        if (v35 <= (v35 * 519)) {
                            fmt.Println(fi7, v1)
                        }
                        delete(v4, 2147483648)
                        _ = v35
                        _ = v36
                    } else {
                        var v37 float32 = 1e10
                        if true {
                            var v38 uint64 = 2
                            _ = v38
                        }
                        _ = len(v2)
                        _ = v37
                    }
                    v2 = append(v2, 4294967295)
                    var v39 int = (fi6 - fi7)
                    _ = v39
                }
            }
            var v40 uint32 = 4294967295
            switch ("hello" + v3) {
            case "\n\t":
                if (-32768 == 32767) {
                    switch v3 {
                    case "":
                        if (-1 <= -1) {
                            switch v3 {
                            case "a":
                                var v41 float32 = 0.0
                                _ = v41
                            case "hello":
                                v1 += (^(v1))
                            }
                        } else {
                            _ = len(v4)
                            if false {
                                switch fi7 {
                                case 0:
                                    fmt.Println(v5, fi6, v1)
                                case -1:
                                    if false {
                                        fmt.Println(fi6, v1, v5)
                                        if true {
                                            var v42 uint64 = uint64(v40)
                                            var v43 int16 = 796
                                            _ = v42
                                            _ = v43
                                        } else {
                                            switch "" {
                                            case "\n\t":
                                                var v44 uint = 65535
                                                _ = v44
                                            case "a":
                                                if (789 == 2) {
                                                    if false {
                                                        switch ((v3 + ("" + (("a" + (v3 + v3)) + (((v3 + (v3 + v3)) + v3) + "\n\t")))) + (v3 + v3)) {
                                                        case "foo bar":
                                                            if false {
                                                                var v45 bool = false
                                                                _ = v45
                                                            } else {
                                                                fmt.Println(v3, fi7)
                                                            }
                                                        case "a":
                                                            v3 = v3
                                                        default:
                                                            var v46 uint64 = 42
                                                            _ = v46
                                                        }
                                                        v1 = (v1 % (((v1 / ((42) | 1))) | 1))
                                                    } else {
                                                        switch 282 {
                                                        case -275:
                                                            v40 = (v40 << ((65) & 63))
                                                        case -1:
                                                            if true {
                                                                if true {
                                                                    if false {
                                                                        v47 := GenS7{F0: ((v40 << ((4294967295) & 63)) <= 0), F1: 827}
                                                                        var v48 int16 = 1000
                                                                        v3 = ((v3 + (v3 + ("foo bar" + (v3 + (("\n\t" + ((("a" + v3) + v3) + v3)) + "\n\t"))))) + "hello")
                                                                        if (true || false) {
                                                                            v5 += v5
                                                                        } else {
                                                                            fmt.Println(v48, v40)
                                                                            v5 += 1.5
                                                                        }
                                                                        _ = v47
                                                                        _ = v48
                                                                    }
                                                                    fmt.Println(v1)
                                                                    delete(v4, (v40 & v40))
                                                                } else {
                                                                    delete(v4, (v40 >> ((366) & 63)))
                                                                }
                                                                if (((v40 / ((481) | 1)) < v40) || true) {
                                                                    var v49 uint16 = 915
                                                                    _ = v49
                                                                } else {
                                                                    v40 -= 315
                                                                    switch (fi7 + -173) {
                                                                    case 4294967295:
                                                                        if len(v2) > 0 {
                                                                            v2[0] = v40
                                                                        }
                                                                    case 294:
                                                                        switch (fi6 % ((32768) | 1)) {
                                                                        case -2147483648:
                                                                            if ((!(!true)) && (182 < 42)) {
                                                                                var v50 uint = 346
                                                                                v40 = v40
                                                                                v5 = (v5 * 1.5)
                                                                                v1 = (v1 / ((uint8(v40)) | 1))
                                                                                _ = v50
                                                                            }
                                                                            v5 = -2.25
                                                                            switch (fi6 % (((-(fi7))) | 1)) {
                                                                            case -2147483648:
                                                                                _ = len(v4)
                                                                            case 32767:
                                                                                var v51 uint = 25
                                                                                _ = v51
                                                                            case -1:
                                                                                var v52 uint = uint(v1)
                                                                                _ = v52
                                                                            }
                                                                        default:
                                                                            if (-117 < -1) {
                                                                                v53 := GenS2{F0: 42, F1: 134}
                                                                                _ = v53
                                                                            }
                                                                        }
                                                                    }
                                                                    var v54 int32 = -1
                                                                    _ = v54
                                                                }
                                                                var v55 float32 = 0.0
                                                                _ = v55
                                                            }
                                                        case 654:
                                                            v3 = "hello"
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        var v56 int = int(v40)
                                                        v57 := GenS8{F0: 417, F1: 1}
                                                        v40 += v40
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                    var v58 uint = 153
                                                    v1 += v1
                                                    _ = v58
                                                }
                                            case "hello":
                                                if (0 == 42) {
                                                    if len(v2) > 0 {
                                                        v2[0] = (v40 ^ 646)
                                                    }
                                                    if false {
                                                        fmt.Println(v40, fi6)
                                                    }
                                                    v59 := GenS2{F0: 2, F1: 1}
                                                    _ = v59
                                                } else {
                                                    var v60 int = (^(fi6))
                                                    v61 := GenS8{F0: 104, F1: 42}
                                                    v62 := GenS9{F0: (v3 + "\n\t")}
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                }
                                            }
                                            var v63 int16 = -32768
                                            _ = v63
                                        }
                                        var v64 float32 = 0.1
                                        _ = v64
                                    } else {
                                        v65 := GenS10{F0: -530, F1: 42, F2: true}
                                        switch v3 {
                                        case "foo bar":
                                            v66 := GenS2{F0: 4294967295, F1: 672}
                                            _ = v66
                                        case "":
                                            if (916 > 663) {
                                                if false {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                v1 -= v1
                                                v2 = append(v2, (v40 >> ((18446744073709551615) & 63)))
                                            } else {
                                                var v67 uint8 = (v1 ^ v1)
                                                _ = v67
                                            }
                                        default:
                                            v68 := GenS11{F0: fi6, F1: "a"}
                                            _ = v68
                                        }
                                        v69 := GenS4{F0: true, F1: 2147483647}
                                        _ = v65
                                        _ = v69
                                    }
                                case 4294967295:
                                    v3 = ("\n\t" + "foo bar")
                                default:
                                    fmt.Println(fi7)
                                }
                            } else {
                                v1 = (v1 >> ((42) & 63))
                            }
                        }
                        v2 = append(v2, 4294967295)
                        switch 42 {
                        case -2147483648:
                            var v70 uint8 = v1
                            _ = v70
                        case 190:
                            v5 = ((v5 + v5) * (1.5 * v5))
                        }
                        v3 = v3
                    case "hello":
                        v5 = 0.1
                    }
                }
            case "a":
                v40 = (v40 * 2147483647)
            default:
                if len(v2) > 0 {
                    v2[0] = (v40 / ((65535) | 1))
                }
            }
            v71 := GenS12{F0: -1}
            _ = fi7
            _ = v40
            _ = v71
        }
        v5 -= (float64((float32(((v5 / ((((((3.14159 + float64(((1e10 + 0.0) - 0.0))) - 0.1) - 0.0) - float64(1.5)) + v5) - (v5 + v5))) / v5)) - (3.14159 - 3.14159))) * v5)
        fmt.Println(v1, v1)
        _ = fi6
    }
    v1 = (v1 + 127)
    var v72 uint = 18446744073709551615
    if (2 < 65535) {
        var v73 uint = (^(v72))
        _ = v73
    } else {
        v74 := GenS8{F0: 277, F1: 18446744073709551615}
        var v75 float64 = v5
        _ = v74
        _ = v75
    }
    fmt.Println(v5)
    if (!(-32768 <= 42)) {
        if (false && (1e10 <= v5)) {
            v76 := GenS6{F0: 42}
            v1 += v1
            _ = v76
        }
    }
    if false {
        for fi77 := 0; fi77 < 9; fi77++ {
            fmt.Println(v72, v3)
            var v78 int16 = 0
            _ = fi77
            _ = v78
        }
    }
    v5 = float64(0.1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
}
