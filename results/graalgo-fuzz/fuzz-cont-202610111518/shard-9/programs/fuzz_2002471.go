package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int; F1 string; F2 int64 }
type GenS5 struct { F0 uint16; F1 int32; F2 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 int64; F1 uint; F2 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 32768
    v3 := []string{"foo bar", "hello", "\n\t", "hello"}
    var v4 int64 = 9223372036854775807
    var v5 uint8 = 1
    var v6 uint8 = v5
    var v7 uint32 = 2147483648
    var v8 uint8 = (v5 >> ((18446744073709551615) & 63))
    v9 := GenS1{F0: true, F1: (^(v5)), F2: (v8 & (v6 % ((195) | 1)))}
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 string = "hello"
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        if (true && (v7 > (v7 >> ((4294967295) & 63)))) {
            if (true && ((!(!false)) || (0 != -38))) {
                v12 := GenS2{F0: (^(v5))}
                _ = v12
            } else {
                v4 = (^(v4))
                v5 = (v5 - 86)
            }
            var v13 uint64 = 18446744073709551615
            switch 0 {
            case 1:
                v14 := GenS1{F0: ((((true || true) || true) && true) || (!(!(v13 != v13)))), F1: (v6 & (v6 % (((v8 << ((18446744073709551615) & 63))) | 1))), F2: (v8 % ((uint8(v7)) | 1))}
                _ = v14
            case 32767:
                if (2 > 1) {
                    var v15 uint16 = 195
                    var v16 uint = 177
                    _ = v15
                    _ = v16
                } else {
                    var v17 int64 = (v4 >> ((42) & 63))
                    v17 = v17
                    v13 += 42
                    _ = v17
                }
            case -946:
                if (!(v11 < (("\n\t" + v11) + (v11 + ((("a" + ((v11 + ("a" + v11)) + "")) + v11) + v11))))) {
                    var v18 uint = 757
                    switch v11 {
                    case "hello":
                        v13 += 1
                    case "\n\t":
                        if ((false && (false && true)) || (!(false && true))) {
                            fmt.Println(v5, v13, fi10)
                            if false {
                                switch (v11 + "hello") {
                                case "\n\t":
                                    switch (v2 + (fi10 - -143)) {
                                    case -1:
                                        switch (v2 * 128) {
                                        case 42:
                                            if (!((^(fi10)) <= int(v18))) {
                                                v19 := GenS3{F0: (v13 << (((v18 ^ v18)) & 63)), F1: 32767}
                                                for fi20 := 0; fi20 < 10; fi20++ {
                                                    if (true && ((42 >= 2147483647) && true)) {
                                                        switch -620 {
                                                        case -76:
                                                            var v21 uint8 = 213
                                                            _ = v21
                                                        case 0:
                                                            var v22 uint8 = (^(v8))
                                                            _ = v22
                                                        }
                                                    }
                                                    _ = fi20
                                                }
                                                _ = v19
                                            }
                                        case 127:
                                            v23 := GenS4{F0: -219, F1: ("\n\t" + v11), F2: v4}
                                            _ = v23
                                        default:
                                            v4 += v4
                                        }
                                    case -631:
                                        var v24 uint = uint(v8)
                                        _ = v24
                                    case 32767:
                                        if false {
                                            fmt.Println(v5, fi10, v7)
                                            var v25 uint8 = 42
                                            _ = v25
                                        } else {
                                            var v26 int = fi10
                                            switch v11 {
                                            case "hello":
                                                fmt.Println(v5, v13, v18)
                                            case "":
                                                var v27 int8 = 42
                                                _ = v27
                                            }
                                            _ = v26
                                        }
                                    }
                                case "a":
                                    var v28 int = 0
                                    _ = v28
                                case "foo bar":
                                    v8 -= 155
                                default:
                                    delete(v1, (v7 << ((v18) & 63)))
                                }
                                v5 = (v8 & 59)
                                switch fi10 {
                                case 127:
                                    fmt.Println(v8)
                                }
                            } else {
                                var v29 int16 = -757
                                switch (^(fi10)) {
                                case 127:
                                    for fi30 := 0; fi30 < 11; fi30++ {
                                        v7 += v7
                                        var v31 int = int(v5)
                                        var v32 string = "a"
                                        _ = fi30
                                        _ = v31
                                        _ = v32
                                    }
                                case -539:
                                    for fi33 := 0; fi33 < 7; fi33++ {
                                        var v34 uint16 = 1
                                        _ = len(v1)
                                        _ = len(v1)
                                        _ = fi33
                                        _ = v34
                                    }
                                default:
                                    switch v11 {
                                    case "hello":
                                        fmt.Println(v4, fi10)
                                    case "foo bar":
                                        switch v2 {
                                        case 32768:
                                            v6 = (v6 + v6)
                                        case -2147483648:
                                            if false {
                                                for fi35 := 0; fi35 < 15; fi35++ {
                                                    var v36 int = (v2 % ((-37) | 1))
                                                    _ = fi35
                                                    _ = v36
                                                }
                                                var v37 uint16 = 65535
                                                switch (^(fi10)) {
                                                case -1:
                                                    v38 := GenS3{F0: v13, F1: (v2 ^ v2)}
                                                    _ = v38
                                                }
                                                _ = v37
                                            }
                                        case 128:
                                            fmt.Println(v6)
                                        default:
                                            fmt.Println(v13, fi10, v11)
                                        }
                                    case "\n\t":
                                        v3 = append(v3, (("\n\t" + (((v11 + v11) + v11) + ((v11 + v11) + v11))) + ("\n\t" + v11)))
                                    default:
                                        v29 += 450
                                    }
                                }
                                _ = v29
                            }
                            v4 = (v4 + int64(v8))
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch v11 {
                            case "hello":
                                v39 := GenS5{F0: 229, F1: -408, F2: (v2 & fi10)}
                                _ = v39
                            default:
                                if true {
                                    v6 -= (v8 ^ (v8 << (((^(v18))) & 63)))
                                    v4 += 548
                                }
                            }
                        }
                    case "":
                        switch fi10 {
                        case 2147483647:
                            switch "" {
                            case "foo bar":
                                v40 := GenS1{F0: (547 > 1000), F1: (v5 * 33), F2: (v6 << ((255) & 63))}
                                _ = v40
                            case "\n\t":
                                v3 = append(v3, (v11 + ("a" + v11)))
                            case "hello":
                                if (false && (v8 >= 255)) {
                                    for fi41 := 0; fi41 < 8; fi41++ {
                                        if (!true) {
                                            delete(v1, uint32(v18))
                                        }
                                        v13 -= v13
                                        _ = len(v3)
                                        _ = fi41
                                    }
                                    v8 = (v6 / ((v5) | 1))
                                }
                            }
                        case -2147483648:
                            if false {
                                if false {
                                    fmt.Println(v4, v8)
                                    v42 := v1[838]
                                    for fi43 := 0; fi43 < 5; fi43++ {
                                        var v44 string = ""
                                        v4 = 0
                                        _ = fi43
                                        _ = v44
                                    }
                                    _ = v42
                                }
                            }
                        default:
                            var v45 int16 = int16(v6)
                            _ = v45
                        }
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    for fi46 := 0; fi46 < 10; fi46++ {
                        v47 := GenS5{F0: 562, F1: -384, F2: (v2 & -160)}
                        var v48 float32 = (float32((3.14159 * 3.14159)) / -2.25)
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    _ = v18
                }
            }
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v5 = v5
    v49 := GenS4{F0: (v2 * int(v7)), F1: ((("foo bar" + ((("hello" + "foo bar") + "\n\t") + "")) + ("a" + (("a" + (("foo bar" + (("a" + ("a" + ("hello" + "a"))) + "foo bar")) + ("hello" + "hello"))) + "a"))) + "a"), F2: (v4 / ((v4) | 1))}
    if ((!((!(!(-2.25 == 3.14159))) && (!(0 > 2147483647)))) || (int64(v7) > v4)) {
        if true {
            v3 = append(v3, "\n\t")
        }
        for fi50 := 0; fi50 < 14; fi50++ {
            v2 -= (v2 | (fi50 & fi50))
            _ = len(v1)
            for fi51 := 0; fi51 < 9; fi51++ {
                switch "foo bar" {
                case "foo bar":
                    if false {
                        v52 := GenS6{F0: 9223372036854775807}
                        v53 := GenS7{F0: "hello", F1: (-2.25 == 0.0)}
                        var v54 int64 = (v4 % (((v4 * int64(v2))) | 1))
                        var v55 int16 = -32768
                        var v56 uint32 = 0
                        v7 = v7
                        _ = v52
                        _ = v53
                        _ = v54
                        _ = v55
                        _ = v56
                    }
                case "hello":
                    var v57 int16 = -48
                    _ = v57
                case "":
                    v58 := GenS1{F0: (0.0 >= (3.14159 + ((1.5 / float64(1.5)) - (1e10 - (float64(1e10) / -2.25))))), F1: v6, F2: v8}
                    _ = v58
                }
                var v59 float32 = 0.0
                switch fi51 {
                case 127:
                    fmt.Println(v7, v2)
                }
                _ = fi51
                _ = v59
            }
            _ = fi50
        }
    }
    fmt.Println(v7)
    var v60 bool = false
    var v61 int16 = 32767
    for fi62 := 0; fi62 < 1; fi62++ {
        if ((!v60) || v60) {
            fmt.Println(v60, v2)
            v63 := GenS3{F0: 42, F1: (fi62 & (^(fi62)))}
            _ = v63
        } else {
            if false {
                v2 = (fi62 % (((fi62 >> ((42) & 63))) | 1))
                for fi64 := 0; fi64 < 2; fi64++ {
                    v5 -= v6
                    var v65 float32 = (0.1 - -2.25)
                    var v66 float64 = 1.5
                    fmt.Println(v61, v60, fi64)
                    _ = fi64
                    _ = v65
                    _ = v66
                }
            } else {
                _ = len(v1)
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v67 := GenS6{F0: int64(v6)}
        _ = fi62
        _ = v67
    }
    var v68 float32 = 1.5
    var v69 string = "foo bar"
    for fi70 := 0; fi70 < 3; fi70++ {
        fmt.Println(fi70, v4)
        v6 = v5
        _ = fi70
    }
    switch 128 {
    case 1:
        for fi71 := 0; fi71 < 13; fi71++ {
            v72 := GenS8{F0: -79, F1: 182, F2: ("\n\t" + "")}
            v8 = (v6 & v6)
            fmt.Println(v60, v68)
            _ = fi71
            _ = v72
        }
    case 377:
        var v73 uint64 = 42
        _ = v73
    default:
        var v74 float32 = float32(0.1)
        _ = v74
    }
    fmt.Println(v7, v7)
    v6 += v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v60
    _ = v61
    _ = v68
    _ = v69
    _ = v7
    _ = v8
    _ = v9
}
