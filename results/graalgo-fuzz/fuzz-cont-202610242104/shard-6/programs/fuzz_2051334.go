package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 int16; F2 uint32 }

func main() {
    var v1 int8 = 100
    var v2 int64 = -1
    v3 := []string{}
    v4 := []int{2147483647, 125}
    if true {
        if (!(!false)) {
            switch ("hello" + "foo bar") {
            case "":
                v5 := GenS1{F0: -416, F1: 255, F2: 255}
                _ = v5
            }
            for fi6 := 0; fi6 < 1; fi6++ {
                switch fi6 {
                case -74:
                    if false {
                        v2 = (v2 | -1)
                    }
                case -9223372036854775808:
                    v1 = (-(v1))
                }
                _ = fi6
            }
            v1 += (v1 + (^(v1)))
        } else {
            v2 -= v2
            v7 := GenS1{F0: -791, F1: 1, F2: 621}
            if ((v2 << ((869) & 63)) == (v2 * v2)) {
                fmt.Println(v2, v2)
            }
            _ = v7
        }
        v2 = (v2 / (((v2 ^ 0)) | 1))
    } else {
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        v8 := GenS1{F0: 127, F1: 500, F2: 255}
        var v9 uint8 = 124
        _ = v8
        _ = v9
    }
    v2 += v2
    for fi10 := 0; fi10 < 2; fi10++ {
        v1 = (v1 | -113)
        _ = fi10
    }
    if true {
        var v11 int = 2147483647
        var v12 int16 = -207
        v13 := GenS2{F0: v11, F1: 0}
        _ = v11
        _ = v12
        _ = v13
    } else {
        v1 = (-(v1))
        if true {
            for fi14 := 0; fi14 < 1; fi14++ {
                v15 := GenS3{F0: (fi14 % (((fi14 << ((880) & 63))) | 1))}
                var v16 int32 = 2147483647
                if true {
                    fmt.Println(fi14, v1, fi14)
                    v16 = v16
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    v1 = v1
                    if false {
                        var v17 int32 = -1
                        if false {
                            if (fi14 != int(v1)) {
                                v16 -= 740
                            } else {
                                v18 := GenS3{F0: (^(fi14))}
                                var v19 uint64 = 18446744073709551615
                                _ = v18
                                _ = v19
                            }
                            v2 = (v2 << ((42) & 63))
                        } else {
                            fmt.Println(v1, fi14)
                        }
                        var v20 string = "a"
                        _ = v17
                        _ = v20
                    } else {
                        fmt.Println(v2, v1)
                    }
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            fmt.Println(v1, v2)
        }
        if ("\n\t" >= "foo bar") {
            v1 = v1
            var v21 uint = 2
            fmt.Println(v2, v1, v1)
            _ = v21
        } else {
            if (2147483647 >= 42) {
                fmt.Println(v1, v2)
                switch 32768 {
                case 411:
                    var v22 int32 = -1
                    _ = v22
                }
            }
            var v23 bool = true
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 uint64 = 2
        _ = fi24
        _ = v25
    }
    if true {
        var v26 uint16 = 255
        fmt.Println(v2, v1)
        _ = len(v4)
        _ = v26
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 += v2
    if true {
        _ = len(v3)
        switch 32767 {
        case 447:
            switch "" {
            case "\n\t":
                if false {
                    if (-2.25 <= 0.1) {
                        var v27 bool = false
                        fmt.Println(v1, v27)
                        _ = v27
                    } else {
                        v28 := GenS4{F0: 255}
                        var v29 float32 = 3.14159
                        _ = v28
                        _ = v29
                    }
                    if (((1 == 255) && (false && (false || (805 < 4294967295)))) && (!false)) {
                        if (0.1 <= (float32(1.5) * (1e10 + float32(-2.25)))) {
                            if (!false) {
                                v30 := GenS5{F0: -742}
                                if (!false) {
                                    v31 := GenS1{F0: 127, F1: 638, F2: 1}
                                    v32 := GenS3{F0: 2147483647}
                                    _ = v31
                                    _ = v32
                                }
                                _ = v30
                            }
                        } else {
                            fmt.Println(v2, v1)
                        }
                        v2 -= (v2 + v2)
                        fmt.Println(v2, v2)
                    }
                    for fi33 := 0; fi33 < 19; fi33++ {
                        switch ("hello" + ("foo bar" + "\n\t")) {
                        case "hello":
                            switch ("" + (((("foo bar" + "hello") + "hello") + (("" + (("\n\t" + "\n\t") + "a")) + "a")) + ("\n\t" + ("hello" + "hello")))) {
                            case "foo bar":
                                var v34 int8 = -1
                                _ = v34
                            case "\n\t":
                                for fi35 := 0; fi35 < 14; fi35++ {
                                    v2 = (^(v2))
                                    v36 := GenS5{F0: 723}
                                    _ = fi35
                                    _ = v36
                                }
                            }
                        case "foo bar":
                            fmt.Println(v1, fi33)
                        case "a":
                            for fi37 := 0; fi37 < 18; fi37++ {
                                var v38 uint32 = 65535
                                v39 := GenS4{F0: 65535}
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                        }
                        if len(v4) > 0 {
                            v4[0] = 1
                        }
                        switch "" {
                        case "hello":
                            v2 = (v2 | v2)
                        case "\n\t":
                            var v40 float32 = 1.5
                            _ = v40
                        default:
                            for fi41 := 0; fi41 < 9; fi41++ {
                                v42 := GenS1{F0: fi33, F1: 32767, F2: 0}
                                _ = fi41
                                _ = v42
                            }
                        }
                        _ = fi33
                    }
                }
            default:
                for fi43 := 0; fi43 < 6; fi43++ {
                    for fi44 := 0; fi44 < 13; fi44++ {
                        if (int8(fi44) != (v1 & (v1 | v1))) {
                            if (true && true) {
                                var v45 float32 = 0.1
                                _ = v45
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            var v46 int32 = -947
                            fmt.Println(v46, v46)
                            _ = v46
                        }
                        fmt.Println(fi44)
                        _ = fi44
                    }
                    v47 := GenS6{F0: "foo bar", F1: 0, F2: 714}
                    _ = fi43
                    _ = v47
                }
            }
        case 254:
            var v48 uint64 = 601
            _ = v48
        case 32768:
            fmt.Println(v2, v2)
        default:
            switch -2147483648 {
            case -713:
                v2 = v2
            case 1:
                if false {
                    var v49 uint8 = 42
                    _ = v49
                }
            case 2147483647:
                var v50 string = ("hello" + (("hello" + "\n\t") + "\n\t"))
                _ = v50
            }
        }
    }
    v1 = v1
    for fi51 := 0; fi51 < 17; fi51++ {
        for fi52 := 0; fi52 < 16; fi52++ {
            fmt.Println(v2, v2)
            _ = fi52
        }
        var v53 float64 = 0.1
        fmt.Println(v1)
        _ = fi51
        _ = v53
    }
    for fi54 := 0; fi54 < 5; fi54++ {
        if (false && (!true)) {
            if len(v4) > 0 {
                v4[0] = (fi54 * (fi54 >> ((42) & 63)))
            }
        }
        var v55 int = 32767
        var v56 string = "a"
        _ = fi54
        _ = v55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
