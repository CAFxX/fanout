package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 string }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int64 }

func main() {
    var v1 uint64 = 4294967295
    v2 := []int{}
    var v3 uint16 = 690
    var v4 string = ((("" + "") + ("a" + "hello")) + "")
    v5 := make(map[int]uint32)
    var v6 int16 = -642
    var v7 int64 = 9223372036854775807
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            v4 = ((((v4 + v4) + v4) + ("foo bar" + ((v4 + ((("a" + v4) + v4) + v4)) + ""))) + v4)
            switch fi9 {
            case 1:
                v6 = v6
            default:
                var v10 int8 = 0
                v6 -= (v6 << ((255) & 63))
                var v11 float32 = float32((1e10 * -2.25))
                var v12 uint64 = (v1 ^ v1)
                v13 := GenS1{F0: (v10 | (v10 | (v10 << ((4294967295) & 63)))), F1: 4294967295, F2: ("\n\t" + "hello")}
                _ = v10
                _ = v11
                _ = v12
                _ = v13
            }
            switch (^(fi9)) {
            case 680:
                if true {
                    var v14 uint64 = (^(v1))
                    var v15 string = "foo bar"
                    _ = v14
                    _ = v15
                }
            case 32768:
                v6 += (v6 ^ (v6 << ((42) & 63)))
            default:
                var v16 int8 = 1
                _ = v16
            }
            _ = fi9
        }
        _ = fi8
    }
    v17 := GenS1{F0: 1, F1: 2, F2: v4}
    fmt.Println(v1, v3)
    var v18 uint8 = 230
    var v19 uint = 65535
    var v20 int32 = 1
    var v21 uint64 = (^(v1))
    if false {
        fmt.Println(v21)
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 uint64 = (v1 | (v21 * 1))
        for fi24 := 0; fi24 < 14; fi24++ {
            fmt.Println(v23, v3, v1)
            var v25 string = v4
            _ = fi24
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    if (!false) {
        fmt.Println(v20)
    }
    switch v4 {
    case "\n\t":
        for fi26 := 0; fi26 < 4; fi26++ {
            if false {
                switch v4 {
                case "":
                    if (!(v7 > (v7 << ((2) & 63)))) {
                        var v27 float64 = 0.0
                        for fi28 := 0; fi28 < 13; fi28++ {
                            v7 += v7
                            _ = fi28
                        }
                        _ = v27
                    } else {
                        if true {
                            v1 += v21
                            v20 = (v20 & v20)
                        } else {
                            v29 := GenS1{F0: 119, F1: 2, F2: "a"}
                            if ((!true) || true) {
                                v21 = (^(v1))
                                switch fi26 {
                                case 545:
                                    v4 = "a"
                                default:
                                    if ((!true) || (uint8(v19) <= (^(v18)))) {
                                        v20 -= 0
                                        var v30 uint16 = 123
                                        v6 -= (^(v6))
                                        _ = v30
                                    }
                                }
                            }
                            _ = v29
                        }
                        var v31 string = (("\n\t" + v4) + "\n\t")
                        _ = v31
                    }
                case "hello":
                    v21 = 4294967295
                case "a":
                    switch (fi26 + (fi26 + (fi26 | 32767))) {
                    case 127:
                        if true {
                            fmt.Println(v21, v21, fi26)
                            _ = len(v5)
                        } else {
                            if true {
                                switch v4 {
                                case "":
                                    if false {
                                        var v32 float32 = ((float32(0.1) + 3.14159) / -2.25)
                                        switch (v4 + (((v4 + v4) + v4) + "\n\t")) {
                                        case "":
                                            if ((v7 & (v7 & (v7 & v7))) > int64(v6)) {
                                                if false {
                                                    var v33 bool = false
                                                    _ = v33
                                                }
                                                v19 -= v19
                                            } else {
                                                for fi34 := 0; fi34 < 7; fi34++ {
                                                    var v35 uint32 = 1
                                                    var v36 int = fi26
                                                    _ = len(v5)
                                                    _ = fi34
                                                    _ = v35
                                                    _ = v36
                                                }
                                                var v37 uint = (v19 + v19)
                                                _ = v37
                                            }
                                        case "foo bar":
                                            fmt.Println(v32)
                                        case "\n\t":
                                            delete(v5, (fi26 >> ((v19) & 63)))
                                        }
                                        var v38 string = v4
                                        _ = v32
                                        _ = v38
                                    }
                                case "a":
                                    v6 -= v6
                                case "hello":
                                    for fi39 := 0; fi39 < 3; fi39++ {
                                        if true {
                                            v19 -= 438
                                            var v40 uint64 = 1
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v40
                                        }
                                        switch 4294967295 {
                                        case 127:
                                            v41 := GenS1{F0: 127, F1: 4294967295, F2: v4}
                                            _ = v41
                                        case 4294967295:
                                            var v42 uint64 = v21
                                            _ = v42
                                        case -9223372036854775808:
                                            if false {
                                                var v43 float32 = 1.5
                                                _ = v43
                                            } else {
                                                var v44 uint8 = (^(v18))
                                                _ = v44
                                            }
                                        }
                                        _ = fi39
                                    }
                                default:
                                    switch v4 {
                                    case "a":
                                        fmt.Println(v19, v6, fi26)
                                    }
                                }
                            } else {
                                if false {
                                    v4 = v4
                                } else {
                                    v3 += 65535
                                    switch (v4 + "hello") {
                                    case "a":
                                        var v45 bool = true
                                        _ = v45
                                    case "foo bar":
                                        var v46 uint16 = v3
                                        _ = v46
                                    case "\n\t":
                                        switch fi26 {
                                        case 32768:
                                            var v47 bool = true
                                            _ = v47
                                        case -549:
                                            v20 = -983
                                        case -974:
                                            v48 := GenS1{F0: -1, F1: 2147483647, F2: "hello"}
                                            _ = v48
                                        default:
                                            v21 = (v21 - 1)
                                        }
                                    }
                                    v49 := GenS2{F0: 2, F1: (^(v3)), F2: 2}
                                    _ = v49
                                }
                            }
                            if true {
                                var v50 bool = (false && true)
                                _ = v50
                            } else {
                                _ = len(v5)
                            }
                            var v51 uint64 = 18446744073709551615
                            _ = v51
                        }
                    case -158:
                        var v52 int8 = 0
                        _ = v52
                    }
                default:
                    if false {
                        var v53 bool = (1 == v1)
                        _ = v53
                    }
                }
                if (2 <= v18) {
                    v54 := GenS2{F0: 2147483648, F1: (v3 * uint16(v19)), F2: 284}
                    v20 += v20
                    _ = v54
                } else {
                    v55 := GenS2{F0: 4294967295, F1: 65535, F2: 111}
                    v56 := GenS3{F0: uint64(v18), F1: (fi26 & fi26)}
                    fmt.Println(v4, v21, v4)
                    _ = v55
                    _ = v56
                }
                var v57 uint = (^(v19))
                _ = v57
            }
            switch (fi26 * int(v3)) {
            case 9223372036854775807:
                if true {
                    for fi58 := 0; fi58 < 2; fi58++ {
                        var v59 uint16 = 65535
                        var v60 int16 = (v6 * (v6 + (v6 * (v6 + (v6 + v6)))))
                        v20 -= (-(v20))
                        v60 = 1000
                        switch fi26 {
                        case 614:
                            fmt.Println(v21, v7, fi58)
                        }
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                }
            case 0:
                v61 := GenS4{F0: (v7 - (^(v7)))}
                _ = v61
            default:
                var v62 uint64 = (v21 >> (((v19 | uint(v3))) & 63))
                _ = v62
            }
            _ = len(v5)
            _ = fi26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
