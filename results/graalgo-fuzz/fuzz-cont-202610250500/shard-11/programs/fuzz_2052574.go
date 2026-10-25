package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 string; F2 int16 }
type GenS8 struct { F0 int8; F1 uint; F2 uint }

func main() {
    v1 := make(map[int]string)
    v2 := []int{374, 32768, 2147483647}
    v3 := make(map[string]uint32)
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 uint16 = 32767
        _ = fi4
        _ = v5
    }
    switch "\n\t" {
    case "\n\t":
        var v6 uint64 = 0
        _ = v6
    case "hello":
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 uint16 = 65535
            for fi9 := 0; fi9 < 16; fi9++ {
                v8 = v8
                fmt.Println(fi9, fi9)
                _ = fi9
            }
            if len(v2) > 0 {
                v2[0] = (fi7 & fi7)
            }
            _ = fi7
            _ = v8
        }
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        if false {
            var v11 float32 = (1e10 * 1e10)
            v12 := GenS1{F0: -1000, F1: 255, F2: 2}
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 2; fi13++ {
                var v14 bool = false
                var v15 uint32 = 557
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        var v16 int8 = 127
        if (!false) {
            for fi17 := 0; fi17 < 10; fi17++ {
                fmt.Println(fi17)
                switch (^(fi10)) {
                case 127:
                    switch (("" + "hello") + "a") {
                    case "foo bar":
                        v16 = v16
                    case "a":
                        fmt.Println(fi10)
                    }
                case 4294967295:
                    if true {
                        if (0 > 18446744073709551615) {
                            if false {
                                v16 = -6
                                v18 := GenS2{F0: 545, F1: ""}
                                _ = v18
                            } else {
                                v16 = 100
                                fmt.Println(fi17)
                                v19 := GenS3{F0: 255}
                                v16 += v16
                                var v20 int16 = -1000
                                switch "\n\t" {
                                case "":
                                    if (v16 != v16) {
                                        if false {
                                            v21 := GenS3{F0: 182}
                                            v16 = v16
                                            if false {
                                                _ = len(v2)
                                                _ = len(v2)
                                            }
                                            _ = v21
                                        } else {
                                            v22 := GenS4{F0: 195}
                                            if true {
                                                switch fi10 {
                                                case -1:
                                                    var v23 int32 = -2147483648
                                                    _ = v23
                                                case 2147483647:
                                                    v24 := GenS5{F0: 42}
                                                    _ = v24
                                                default:
                                                    var v25 int8 = int8(v20)
                                                    _ = v25
                                                }
                                                var v26 bool = (224 <= 152)
                                                switch "hello" {
                                                case "":
                                                    v3[""] = 42
                                                case "\n\t":
                                                    delete(v1, fi10)
                                                }
                                                _ = v26
                                            } else {
                                                fmt.Println(v20, v20, v16)
                                            }
                                            if false {
                                                v20 -= (v20 ^ 32767)
                                                _ = len(v3)
                                            }
                                            _ = v22
                                        }
                                        var v27 uint = 65535
                                        _ = v27
                                    } else {
                                        v20 = v20
                                    }
                                }
                                _ = v19
                                _ = v20
                            }
                            v16 = (v16 >> ((255) & 63))
                        }
                        switch fi10 {
                        case 2147483647:
                            fmt.Println(fi10, v16)
                        case 126:
                            v28 := GenS6{F0: (127 != v16)}
                            v29 := GenS2{F0: 2, F1: "hello"}
                            fmt.Println(fi10)
                            _ = v28
                            _ = v29
                        default:
                            if false {
                                if true {
                                    if (100 < (v16 ^ (v16 / ((-128) | 1)))) {
                                        if (((("a" + "a") == ("foo bar" + (("\n\t" + "\n\t") + ""))) || (2 > 255)) || true) {
                                            _ = len(v1)
                                        }
                                        v16 = 106
                                        v2 = append(v2, -2147483648)
                                    }
                                    v16 = 1
                                } else {
                                    if (false || true) {
                                        v16 += (v16 + v16)
                                    } else {
                                        v30 := GenS5{F0: v16}
                                        var v31 uint = 835
                                        fmt.Println(v31, v16, v31)
                                        _ = v30
                                        _ = v31
                                    }
                                    switch (^(fi10)) {
                                    case 32767:
                                        fmt.Println(v16)
                                    default:
                                        _ = len(v2)
                                    }
                                    var v32 float64 = float64(0.1)
                                    _ = v32
                                }
                                if ((v16 << ((4294967295) & 63)) > (v16 / (((v16 >> ((4294967295) & 63))) | 1))) {
                                    _ = len(v2)
                                } else {
                                    var v33 uint16 = 32767
                                    var v34 float64 = 0.0
                                    switch "" {
                                    case "":
                                        v34 -= v34
                                    case "\n\t":
                                        v33 = (v33 | (v33 << ((216) & 63)))
                                    default:
                                        if false {
                                            fmt.Println(v33, v34, v34)
                                            var v35 uint32 = 2147483647
                                            if false {
                                                if true {
                                                    v35 += uint32(v33)
                                                    var v36 uint = 0
                                                    _ = v36
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = 2147483647
                                                }
                                                v37 := GenS6{F0: true}
                                                fmt.Println(v33)
                                                _ = v37
                                            } else {
                                                var v38 uint64 = 550
                                                var v39 uint16 = v33
                                                var v40 float32 = 1e10
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                            }
                                            _ = v35
                                        } else {
                                            var v41 int16 = 42
                                            v42 := GenS3{F0: 128}
                                            _ = v41
                                            _ = v42
                                        }
                                    }
                                    delete(v3, "hello")
                                    _ = v33
                                    _ = v34
                                }
                                v16 += 87
                                if false {
                                    v16 -= v16
                                    var v43 float32 = 1e10
                                    delete(v3, "a")
                                    _ = v43
                                }
                            } else {
                                v16 += (-(v16))
                                var v44 int64 = -9223372036854775808
                                _ = v44
                            }
                        }
                        switch (("" + "a") + "") {
                        case "foo bar":
                            v45 := GenS7{F0: fi17, F1: "foo bar", F2: 1000}
                            _ = v45
                        default:
                            v46 := GenS2{F0: 18446744073709551615, F1: "hello"}
                            _ = v46
                        }
                    }
                default:
                    v16 = -20
                }
                var v47 uint8 = 0
                _ = fi17
                _ = v47
            }
        } else {
            if false {
                var v48 int8 = v16
                v48 = v48
                if (1 >= 2) {
                    v49 := GenS6{F0: ((-2.25 - (0.0 / 1e10)) < (float64(1.5) + -2.25))}
                    _ = v49
                }
                _ = v48
            } else {
                v16 += v16
                for fi50 := 0; fi50 < 20; fi50++ {
                    v51 := GenS8{F0: 1, F1: 719, F2: 0}
                    _ = fi50
                    _ = v51
                }
            }
        }
        _ = fi10
        _ = v16
    }
    if len(v2) > 0 {
        v2[0] = -293
    }
    v52 := GenS8{F0: 1, F1: 42, F2: 65535}
    var v53 uint32 = 322
    var v54 int32 = 0
    v2 = append(v2, 4294967295)
    switch -2147483648 {
    case 1:
        fmt.Println(v53, v54)
    case 32767:
        v2 = append(v2, -722)
    }
    if false {
        switch 127 {
        case -1:
            var v55 float32 = 3.14159
            _ = v55
        default:
            switch "hello" {
            case "foo bar":
                v56 := GenS3{F0: 0}
                _ = v56
            case "":
                v57 := GenS1{F0: 1, F1: 2, F2: 255}
                _ = v57
            case "a":
                v58 := GenS4{F0: 66}
                _ = v58
            default:
                v2 = append(v2, 128)
            }
        }
        for fi59 := 0; fi59 < 13; fi59++ {
            for fi60 := 0; fi60 < 16; fi60++ {
                v61 := GenS6{F0: true}
                v54 = v54
                _ = fi60
                _ = v61
            }
            fmt.Println(v54, v54, v54)
            _ = fi59
        }
    } else {
        switch "hello" {
        case "":
            fmt.Println(v54, v53)
        case "\n\t":
            _ = len(v1)
        }
        switch 32768 {
        case 1:
            fmt.Println(v54, v54)
        default:
            var v62 uint64 = 18446744073709551615
            _ = v62
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v53
    _ = v54
}
