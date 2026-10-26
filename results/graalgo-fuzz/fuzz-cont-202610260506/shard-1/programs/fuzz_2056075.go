package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int32 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 int16; F2 bool }

func main() {
    var v1 string = ("hello" + ("a" + ("" + "foo bar")))
    v2 := []uint32{2, 42, 4294967295}
    var v3 bool = ((255 < 817) || true)
    v4 := []uint32{240, 1}
    v5 := []int{87, -479, 798}
    v6 := []string{(((v1 + (v1 + v1)) + "\n\t") + ((v1 + v1) + v1)), "", "foo bar"}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch v1 {
    case "":
        switch v1 {
        case "foo bar":
            var v7 uint = 18446744073709551615
            _ = v7
        case "hello":
            v1 = ((v1 + "a") + (v1 + v1))
        case "\n\t":
            for fi8 := 0; fi8 < 6; fi8++ {
                v1 = v1
                var v9 string = "a"
                v1 = ("\n\t" + v1)
                _ = fi8
                _ = v9
            }
        default:
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 uint16 = 42
                for fi12 := 0; fi12 < 18; fi12++ {
                    if false {
                        var v13 int8 = 0
                        v1 = v1
                        switch "a" {
                        case "":
                            v14 := GenS1{F0: v1, F1: (fi12 + fi12), F2: 309}
                            _ = v14
                        case "a":
                            var v15 uint32 = 836
                            _ = v15
                        case "hello":
                            v16 := GenS1{F0: v1, F1: (fi12 * (fi10 % ((-1) | 1))), F2: 42}
                            _ = v16
                        }
                        _ = v13
                    }
                    v17 := GenS2{F0: 42, F1: (fi10 - (fi10 | fi12))}
                    _ = len(v2)
                    _ = fi12
                    _ = v17
                }
                _ = fi10
                _ = v11
            }
        }
    case "a":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        v3 = (!true)
    }
    if (((1.5 < 3.14159) && v3) && v3) {
        for fi18 := 0; fi18 < 10; fi18++ {
            switch v1 {
            case "hello":
                for fi19 := 0; fi19 < 11; fi19++ {
                    var v20 uint16 = 1
                    if (1.5 != (1.5 / 3.14159)) {
                        v20 = v20
                        v3 = false
                        if len(v6) > 0 {
                            v6[0] = (v1 + ((v1 + v1) + ""))
                        }
                        fmt.Println(v1, fi18)
                        if (!v3) {
                            var v21 float64 = (1e10 + 0.1)
                            v22 := GenS2{F0: -75, F1: (fi19 << ((2) & 63))}
                            _ = v21
                            _ = v22
                        }
                        _ = len(v6)
                        if v3 {
                            var v23 int64 = 726
                            v3 = ((v20 * 645) > v20)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v23
                        }
                        var v24 int = fi19
                        var v25 float64 = 0.1
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v20 -= 192
                        _ = v24
                        _ = v25
                    }
                    var v26 string = ("foo bar" + "a")
                    _ = fi19
                    _ = v20
                    _ = v26
                }
            default:
                v1 = v1
            }
            if (2 > 65535) {
                _ = len(v2)
                _ = len(v6)
                var v27 int32 = 520
                _ = v27
            }
            _ = fi18
        }
        var v28 uint64 = 0
        if true {
            var v29 uint = 42
            v1 = (v1 + v1)
            _ = v29
        } else {
            for fi30 := 0; fi30 < 16; fi30++ {
                _ = len(v4)
                _ = fi30
            }
            switch 1 {
            case 127:
                var v31 string = v1
                _ = v31
            case 128:
                switch 32768 {
                case 0:
                    v28 = (^(v28))
                case 45:
                    if (-1 <= -32768) {
                        fmt.Println(v3, v1)
                    } else {
                        var v32 float64 = 0.1
                        if true {
                            var v33 int16 = -1
                            v28 = (v28 >> ((0) & 63))
                            _ = v33
                        }
                        _ = v32
                    }
                case 127:
                    for fi34 := 0; fi34 < 3; fi34++ {
                        fmt.Println(v28)
                        switch (("\n\t" + "hello") + "") {
                        case "\n\t":
                            if v3 {
                                for fi35 := 0; fi35 < 11; fi35++ {
                                    fmt.Println(v3, v28)
                                    if v3 {
                                        var v36 bool = (18446744073709551615 >= 42)
                                        if ("foo bar" <= (((v1 + v1) + "\n\t") + v1)) {
                                            v36 = false
                                            v28 = v28
                                        } else {
                                            var v37 string = "\n\t"
                                            v1 = "foo bar"
                                            if len(v5) > 0 {
                                                v5[0] = -1
                                            }
                                            _ = v37
                                        }
                                        _ = v36
                                    } else {
                                        v28 = v28
                                    }
                                    v3 = v3
                                    switch fi34 {
                                    case -1:
                                        v1 = v1
                                    case -939:
                                        if v3 {
                                            var v38 uint32 = 923
                                            _ = v38
                                        } else {
                                            v4 = append(v4, 0)
                                            var v39 uint16 = 255
                                            var v40 int64 = 1
                                            _ = v39
                                            _ = v40
                                        }
                                    default:
                                        var v41 uint32 = 1
                                        _ = v41
                                    }
                                    _ = fi35
                                }
                                v42 := GenS1{F0: (v1 + (v1 + v1)), F1: (fi34 + 0), F2: -606}
                                _ = v42
                            }
                        case "hello":
                            for fi43 := 0; fi43 < 14; fi43++ {
                                fmt.Println(v1, fi43)
                                _ = fi43
                            }
                        }
                        _ = fi34
                    }
                default:
                    var v44 uint = 255
                    _ = v44
                }
            case -500:
                fmt.Println(v28, v1)
            }
        }
        _ = v28
    } else {
        v45 := GenS1{F0: (v1 + "a"), F1: 128, F2: -2147483648}
        _ = v45
    }
    v3 = (32767 >= 1)
    for fi46 := 0; fi46 < 6; fi46++ {
        if v3 {
            fmt.Println(v3, v1)
        }
        _ = fi46
    }
    var v47 int16 = 366
    var v48 string = v1
    fmt.Println(v47, v3, v47)
    v49 := GenS3{F0: 42}
    fmt.Println(v47, v48, v48)
    for fi50 := 0; fi50 < 12; fi50++ {
        fmt.Println(v3, v48, fi50)
        var v51 float32 = 1.5
        _ = fi50
        _ = v51
    }
    v52 := GenS4{F0: 553, F1: (v47 | (v47 % (((v47 % (((v47 & 42)) | 1))) | 1))), F2: false}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v47 += 453
    if (!false) {
        fmt.Println(v48, v47)
    } else {
        if ((((v1 + ((((v48 + v48) + (v48 + ("\n\t" + v48))) + ("hello" + "\n\t")) + v1)) + v48) + "") == "") {
            switch (v48 + v1) {
            case "a":
                for fi53 := 0; fi53 < 6; fi53++ {
                    if ((fi53 ^ fi53) <= fi53) {
                        fmt.Println(v1)
                        v48 = v48
                    }
                    _ = fi53
                }
            case "foo bar":
                v5 = append(v5, 1)
            }
        }
        var v54 uint32 = 2147483647
        var v55 int8 = -128
        _ = v54
        _ = v55
    }
    v47 += v47
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v52
    _ = v6
}
