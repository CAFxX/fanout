package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int16; F2 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 bool; F2 bool }

func main() {
    var v1 uint = 1
    v2 := make(map[string]bool)
    var v3 uint = (v1 & (v1 << ((v1) & 63)))
    var v4 uint64 = 901
    v5 := GenS1{F0: "a", F1: 709}
    fmt.Println(v1, v4)
    var v6 int16 = -883
    v7 := []int{int(v4), 42, 32767}
    v8 := GenS2{F0: 1, F1: true, F2: 4294967295}
    for fi9 := 0; fi9 < 20; fi9++ {
        v7 = append(v7, 904)
        v1 = (v3 + v3)
        _ = fi9
    }
    switch 2147483647 {
    case 42:
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 float64 = 1e10
            v3 = v1
            _ = fi10
            _ = v11
        }
    case 127:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    default:
        delete(v2, ("hello" + "a"))
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        switch -336 {
        case 9223372036854775807:
            var v14 int = fi13
            _ = v14
        case 255:
            if true {
                if false {
                    v6 -= (v6 ^ (^(v6)))
                }
            }
        case -107:
            if true {
                var v15 bool = true
                if (v15 && ((v6 % (((^(v6))) | 1)) >= (v6 ^ 396))) {
                    var v16 int64 = 1
                    _ = v16
                } else {
                    var v17 uint = (v3 % ((4294967295) | 1))
                    for fi18 := 0; fi18 < 8; fi18++ {
                        if (!v15) {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            var v19 int = (fi13 ^ (fi18 + fi18))
                            _ = v19
                        } else {
                            switch "hello" {
                            case "hello":
                                v20 := GenS1{F0: ("\n\t" + (("\n\t" + "foo bar") + "\n\t")), F1: -712}
                                _ = v20
                            case "\n\t":
                                var v21 uint = v17
                                _ = v21
                            case "a":
                                if v15 {
                                    fmt.Println(fi18, v3, v6)
                                    v4 += uint64(v17)
                                } else {
                                    if (v15 && ((v15 && ((true && (118 > -1)) && false)) && v15)) {
                                        switch -376 {
                                        case 128:
                                            var v22 uint16 = 972
                                            _ = v22
                                        }
                                        v1 += 255
                                        fmt.Println(v15)
                                        v17 += 255
                                        var v23 string = ""
                                        _ = v23
                                    }
                                    if ((v6 % (((v6 ^ v6)) | 1)) < (v6 / (((v6 & v6)) | 1))) {
                                        if v15 {
                                            v15 = v15
                                            if (4294967295 >= 2147483647) {
                                                v24 := GenS3{F0: "a"}
                                                _ = v24
                                            } else {
                                                v1 -= 0
                                            }
                                            v6 = v6
                                        }
                                    } else {
                                        v25 := GenS4{F0: 9223372036854775807, F1: -487, F2: 0}
                                        _ = v25
                                    }
                                    v1 += (v17 - v1)
                                    v17 = (v3 & v3)
                                    fmt.Println(v4, v17, v17)
                                }
                            }
                            var v26 int16 = v6
                            v27 := GenS4{F0: 4294967295, F1: (v6 ^ 32767), F2: -9223372036854775808}
                            switch (fi13 / ((fi13) | 1)) {
                            case -973:
                                switch fi18 {
                                case 127:
                                    v17 += (v3 >> (((v17 >> ((v17) & 63))) & 63))
                                case 128:
                                    var v28 uint32 = 42
                                    _ = v28
                                }
                            case 1:
                                var v29 uint = v1
                                _ = v29
                            default:
                                v30 := GenS5{F0: 1}
                                _ = v30
                            }
                            v2["\n\t"] = v15
                            if (!(814 > (^(v17)))) {
                                v7 = append(v7, fi18)
                            } else {
                                v31 := GenS3{F0: ("a" + "foo bar")}
                                fmt.Println(v6, v1, v3)
                                var v32 int = -758
                                _ = v31
                                _ = v32
                            }
                            _ = v26
                            _ = v27
                        }
                        _ = fi18
                    }
                    _ = v17
                }
                switch fi13 {
                case -2147483648:
                    var v33 uint = v3
                    _ = v33
                default:
                    v4 -= v4
                }
                _ = v15
            }
        default:
            switch (fi13 & (fi13 + 414)) {
            case 4294967295:
                v3 -= v3
            default:
                for fi34 := 0; fi34 < 13; fi34++ {
                    var v35 string = ("hello" + "foo bar")
                    var v36 uint32 = 42
                    v2[v35] = true
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
        }
        v1 += v1
        var v37 uint8 = 128
        _ = fi13
        _ = v37
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 uint64 = (^(v4))
        _ = fi38
        _ = v39
    }
    v4 = 743
    var v40 uint64 = (v4 + (v4 & v4))
    var v41 float64 = -2.25
    if ((^(v3)) <= v3) {
        var v42 int64 = 0
        switch -9223372036854775808 {
        case 68:
            v41 = v41
        default:
            if len(v7) > 0 {
                v7[0] = 32768
            }
        }
        switch -601 {
        case 269:
            v43 := GenS6{F0: 30, F1: (42 <= 4294967295), F2: true}
            _ = v43
        }
        _ = v42
    } else {
        for fi44 := 0; fi44 < 6; fi44++ {
            for fi45 := 0; fi45 < 6; fi45++ {
                v40 += v4
                _ = fi45
            }
            delete(v2, "a")
            _ = fi44
        }
    }
    v2["foo bar"] = true
    delete(v2, "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
