package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 int; F2 int32 }
type GenS6 struct { F0 uint32; F1 bool; F2 string }

func main() {
    var v1 uint = 199
    var v2 float32 = ((float32(float64(1e10)) + 3.14159) + 0.1)
    var v3 string = (("foo bar" + "hello") + "\n\t")
    v4 := []uint32{650}
    var v5 uint8 = 16
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            v1 = (v1 - v1)
            var v8 int8 = 0
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 uint16 = 65535
    v5 = (^(v5))
    var v10 bool = true
    switch -2147483648 {
    case 859:
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS1{F0: fi11, F1: (-1 == 1), F2: 958}
            var v13 int32 = 0
            _ = fi11
            _ = v12
            _ = v13
        }
    case 128:
        if (v2 > 1.5) {
            switch ("\n\t" + "hello") {
            case "\n\t":
                for fi14 := 0; fi14 < 19; fi14++ {
                    if v10 {
                        var v15 uint8 = 164
                        var v16 float64 = (1.5 + float64(-2.25))
                        _ = v15
                        _ = v16
                    }
                    _ = fi14
                }
            case "hello":
                for fi17 := 0; fi17 < 4; fi17++ {
                    fmt.Println(fi17)
                    var v18 uint64 = 1
                    _ = fi17
                    _ = v18
                }
            case "foo bar":
                v19 := GenS2{F0: 32767, F1: "\n\t"}
                _ = v19
            default:
                v5 -= v5
            }
            fmt.Println(v2, v3)
            switch "" {
            case "\n\t":
                if ("hello" == "") {
                    var v20 int64 = 9223372036854775807
                    _ = len(v4)
                    _ = v20
                }
            case "hello":
                var v21 uint = v1
                _ = v21
            }
        } else {
            var v22 uint8 = 127
            _ = v22
        }
    default:
        switch 175 {
        case 0:
            switch v3 {
            case "a":
                v23 := GenS3{F0: 1, F1: 0}
                _ = v23
            case "":
                switch 32767 {
                case 4294967295:
                    if (795 < 2147483647) {
                        var v24 int64 = 1
                        if (2147483647 > 2147483647) {
                            v25 := GenS4{F0: v9}
                            _ = v25
                        } else {
                            var v26 int32 = 2147483647
                            _ = v26
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v24
                    }
                case 51:
                    switch 4294967295 {
                    case -9223372036854775808:
                        switch -872 {
                        case 2147483647:
                            var v27 string = v3
                            _ = v27
                        }
                    case 137:
                        if (v10 && (uint32(v9) == 2147483648)) {
                            if v10 {
                                var v28 uint32 = 628
                                var v29 uint64 = 1
                                _ = v28
                                _ = v29
                            } else {
                                for fi30 := 0; fi30 < 16; fi30++ {
                                    v10 = v10
                                    if (!false) {
                                        v10 = v10
                                        for fi31 := 0; fi31 < 19; fi31++ {
                                            switch "a" {
                                            case "":
                                                v2 = 1e10
                                            }
                                            _ = fi31
                                        }
                                        var v32 uint16 = 222
                                        _ = v32
                                    }
                                    var v33 int16 = int16(fi30)
                                    _ = fi30
                                    _ = v33
                                }
                            }
                        } else {
                            switch 2147483647 {
                            case 9223372036854775807:
                                v5 -= v5
                            case 371:
                                var v34 uint16 = (v9 & (v9 ^ v9))
                                _ = v34
                            }
                        }
                    case 1:
                        v4 = append(v4, 2147483648)
                    default:
                        for fi35 := 0; fi35 < 20; fi35++ {
                            if len(v4) > 0 {
                                v4[0] = 262
                            }
                            if v10 {
                                switch "" {
                                case "hello":
                                    switch -973 {
                                    case 127:
                                        var v36 int64 = 42
                                        _ = v36
                                    case -521:
                                        switch v3 {
                                        case "\n\t":
                                            var v37 uint32 = 2147483647
                                            _ = v37
                                        }
                                    case 1:
                                        v38 := GenS5{F0: "a", F1: (fi35 << (((v1 >> (((v1 % ((v1) | 1))) & 63))) & 63)), F2: 42}
                                        _ = v38
                                    }
                                case "":
                                    if v10 {
                                        fmt.Println(v10)
                                        v4 = append(v4, 410)
                                    } else {
                                        v10 = (!(!v10))
                                        var v39 int32 = -369
                                        if false {
                                            var v40 uint = (v1 | uint(v9))
                                            _ = len(v4)
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = v40
                                        } else {
                                            var v41 int = (^(fi35))
                                            _ = len(v4)
                                            _ = v41
                                        }
                                        _ = v39
                                    }
                                }
                                if (!v10) {
                                    var v42 float64 = 1.5
                                    var v43 uint32 = 1
                                    _ = v42
                                    _ = v43
                                }
                                var v44 int32 = 1
                                _ = v44
                            }
                            _ = fi35
                        }
                    }
                case 1:
                    var v45 string = (v3 + "")
                    _ = v45
                }
            }
        }
    }
    if v10 {
        for fi46 := 0; fi46 < 5; fi46++ {
            var v47 string = ""
            var v48 int8 = 42
            for fi49 := 0; fi49 < 12; fi49++ {
                fmt.Println(v3, v1, v2)
                var v50 string = v3
                v10 = v10
                if (-67 >= 0) {
                    var v51 uint8 = 250
                    if (((!true) && false) || v10) {
                        _ = len(v4)
                        v52 := GenS6{F0: uint32(v9), F1: v10, F2: "a"}
                        if false {
                            var v53 int16 = int16(v1)
                            _ = v53
                        }
                        _ = v52
                    }
                    v10 = (-1 < -1)
                    _ = v51
                }
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v9 = (v9 >> ((v1) & 63))
                    v54 := GenS1{F0: (fi46 * 4294967295), F1: (("hello" + ((v50 + v47) + (("" + "") + "hello"))) > "foo bar"), F2: 9223372036854775807}
                    _ = v54
                }
                _ = fi49
                _ = v50
            }
            _ = fi46
            _ = v47
            _ = v48
        }
    } else {
        v3 = "a"
        if ((!v10) && ((0 <= 255) || false)) {
            var v55 int8 = 1
            v3 = v3
            var v56 bool = ((v10 && (179 == -2147483648)) && v10)
            _ = v55
            _ = v56
        }
    }
    for fi57 := 0; fi57 < 5; fi57++ {
        v58 := GenS3{F0: 0, F1: 0}
        v4 = append(v4, 754)
        var v59 float64 = ((float64(v2) / 0.1) / 1e10)
        _ = fi57
        _ = v58
        _ = v59
    }
    switch 32768 {
    case 229:
        switch ((("" + ((v3 + (v3 + v3)) + v3)) + v3) + v3) {
        case "\n\t":
            v10 = v10
        default:
            if len(v4) > 0 {
                v4[0] = 2147483648
            }
        }
    }
    v5 = (v5 % ((127) | 1))
    v3 = ((v3 + v3) + (v3 + "a"))
    v2 += 0.0
    var v60 float32 = v2
    var v61 int16 = 32767
    if len(v4) > 0 {
        v4[0] = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v61
    _ = v9
}
