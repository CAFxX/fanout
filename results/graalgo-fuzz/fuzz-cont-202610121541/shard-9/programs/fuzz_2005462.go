package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int32; F1 uint32; F2 int8 }
type GenS3 struct { F0 int16; F1 int64; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 uint; F2 int }
type GenS6 struct { F0 int64; F1 int }
func genFn0(p0 int, p1 int64, p2 string) bool {
    for fi4 := 0; fi4 < 13; fi4++ {
        v5 := GenS1{F0: -4, F1: 62}
        var v6 string = (("" + p2) + (p2 + p2))
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        switch "foo bar" {
        case "a":
            var v8 int = p0
            _ = v8
        default:
            var v9 string = (p2 + ((((((((p2 + (((("foo bar" + "hello") + p2) + (p2 + p2)) + (((p2 + (("a" + "") + "hello")) + (("a" + "\n\t") + "")) + p2))) + "foo bar") + (p2 + ((p2 + ("" + "foo bar")) + ("foo bar" + p2)))) + "foo bar") + (p2 + (("" + (p2 + p2)) + (p2 + (p2 + ""))))) + ("a" + (("\n\t" + "a") + p2))) + "a") + p2))
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 bool = false
        switch (((((p2 + "\n\t") + p2) + "a") + "") + p2) {
        case "foo bar":
            v12 := GenS1{F0: 948, F1: 1}
            _ = v12
        case "":
            if false {
                v11 = true
                for fi13 := 0; fi13 < 1; fi13++ {
                    if v11 {
                        if true {
                            v14 := GenS1{F0: 0, F1: -60}
                            var v15 uint16 = 578
                            if ((fi13 == (-(p0))) || ((42 > 568) && true)) {
                                v16 := []int{(fi13 % (((fi13 ^ 42)) | 1)), fi10, 0}
                                p0 = p0
                                _ = v16
                            } else {
                                v17 := GenS2{F0: int32(p1), F1: 4294967295, F2: 1}
                                _ = v17
                            }
                            v18 := []int{(-(fi13)), fi13, fi13, p0}
                            if (151 != 1) {
                                if len(v18) > 0 {
                                    v18[0] = 0
                                }
                                var v19 float32 = ((3.14159 + (1e10 * 0.1)) + 0.0)
                                _ = v19
                            }
                            _ = v14
                            _ = v15
                            _ = v18
                        } else {
                            v20 := GenS1{F0: -595, F1: -124}
                            v21 := GenS3{F0: int16(p1), F1: (p1 >> ((290) & 63)), F2: 42}
                            _ = v20
                            _ = v21
                        }
                        v22 := []int{(fi13 >> ((1) & 63)), (fi13 & fi10), (p0 ^ fi10), p0}
                        _ = v22
                    } else {
                        if v11 {
                            fmt.Println(p2, v11)
                            var v23 uint16 = 32767
                            v24 := GenS4{F0: ("a" + (((("" + p2) + p2) + p2) + p2))}
                            if ((!true) || true) {
                                v25 := GenS1{F0: 1, F1: 126}
                                var v26 float32 = float32(float64(0.1))
                                switch fi13 {
                                case 4294967295:
                                    var v27 int16 = 1
                                    _ = v27
                                }
                                var v28 uint = 385
                                v29 := GenS3{F0: 1, F1: p1, F2: 282}
                                _ = v25
                                _ = v26
                                _ = v28
                                _ = v29
                            }
                            v30 := GenS1{F0: -1, F1: 100}
                            p0 += fi10
                            _ = v23
                            _ = v24
                            _ = v30
                        } else {
                            switch (fi10 / (((fi13 * (fi10 / (((-(fi10))) | 1)))) | 1)) {
                            case 613:
                                v11 = (1e10 == (1.5 + 0.0))
                            case 990:
                                fmt.Println(p0)
                            }
                            p0 = p0
                            v31 := []int{4294967295, fi10, 42, 128}
                            _ = v31
                        }
                        if (0 >= 2) {
                            var v32 uint16 = 42
                            switch (p2 + "") {
                            case "\n\t":
                                var v33 int64 = -9223372036854775808
                                _ = v33
                            case "":
                                v34 := GenS5{F0: (false || (127 < 255)), F1: 679, F2: 518}
                                if false {
                                    var v35 uint8 = 0
                                    v36 := GenS6{F0: (p1 << ((2) & 63)), F1: (fi10 % ((p0) | 1))}
                                    v37 := GenS6{F0: p1, F1: (fi10 | (-(fi13)))}
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                                var v38 float64 = (3.14159 * 0.1)
                                _ = v34
                                _ = v38
                            default:
                                switch ((p2 + "hello") + "hello") {
                                case "a":
                                    if v11 {
                                        fmt.Println(p2)
                                        if (1e10 < 1.5) {
                                            v11 = (v11 || true)
                                        } else {
                                            switch p2 {
                                            case "\n\t":
                                                v11 = ((!true) && (!false))
                                            case "foo bar":
                                                if ((((3.14159 + (float32((0.0 + (0.1 * float64(0.1)))) * 1.5)) * ((3.14159 + (0.0 * (0.0 * (((float32(3.14159) * float32(0.1)) + (1.5 * (3.14159 * 3.14159))) + float32((3.14159 + 0.1)))))) * 0.0)) / 0.1) <= (float32(-2.25) + ((0.1 / -2.25) / 1.0))) {
                                                    v39 := GenS3{F0: 1000, F1: 429, F2: 4294967295}
                                                    _ = v39
                                                } else {
                                                    v40 := GenS4{F0: p2}
                                                    _ = v40
                                                }
                                                switch (fi10 << ((42) & 63)) {
                                                case 128:
                                                    var v41 int32 = 1
                                                    _ = v41
                                                case 238:
                                                    v42 := []int{p0, -675, 127}
                                                    fmt.Println(fi13)
                                                    _ = v42
                                                case 2147483647:
                                                    v43 := []int{0}
                                                    _ = v43
                                                }
                                                switch (fi10 << ((2) & 63)) {
                                                case 1:
                                                    var v44 string = p2
                                                    _ = v44
                                                case -422:
                                                    if v11 {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            p0 += 32768
                                                        }
                                                        p2 = "a"
                                                        fmt.Println(p0)
                                                    } else {
                                                        var v45 uint64 = 0
                                                        _ = v45
                                                    }
                                                }
                                            }
                                            p1 -= 9223372036854775807
                                        }
                                    }
                                default:
                                    fmt.Println(p0, v11, v32)
                                }
                            }
                            v32 -= (v32 << ((65535) & 63))
                            _ = v32
                        } else {
                            p0 += -954
                        }
                    }
                    p2 = p2
                    switch p2 {
                    case "a":
                        p2 = (p2 + p2)
                    }
                    _ = fi13
                }
            } else {
                if v11 {
                    var v46 uint16 = 65535
                    p0 = fi10
                    _ = v46
                }
                v47 := GenS5{F0: false, F1: 14, F2: fi10}
                fmt.Println(v11, fi10)
                _ = v47
            }
        }
        _ = fi10
        _ = v11
    }
    var v48 float64 = float64(0.1)
    v49 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v48
    _ = v49
    return false
}

func main() {
    var v1 bool = false
    var v2 int16 = 32767
    v3 := []int{939, 2147483647, 90}
    for fi50 := 0; fi50 < 7; fi50++ {
        switch "\n\t" {
        case "hello":
            if v1 {
                for fi51 := 0; fi51 < 8; fi51++ {
                    fmt.Println(v1)
                    _ = fi51
                }
            }
        case "a":
            for fi52 := 0; fi52 < 2; fi52++ {
                var v53 int8 = 127
                _ = fi52
                _ = v53
            }
        }
        _ = fi50
    }
    var v54 uint8 = 186
    if v1 {
        var v55 uint64 = 1
        v3 = append(v3, 743)
        _ = v55
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi56 := 0; fi56 < 15; fi56++ {
            v2 += 842
            _ = fi56
        }
        v1 = false
    }
    var v57 int64 = -1
    var v58 int8 = -1
    var v59 uint64 = 18446744073709551615
    switch "" {
    case "foo bar":
        v58 = v58
    case "hello":
        var v60 int64 = v57
        _ = v60
    case "\n\t":
        fmt.Println(v2, v2)
    }
    fmt.Println(v1)
    for fi61 := 0; fi61 < 6; fi61++ {
        var v62 uint8 = (v54 | v54)
        _ = fi61
        _ = v62
    }
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v63 uint16 = 42
    v3 = append(v3, 638)
    for fi64 := 0; fi64 < 8; fi64++ {
        v63 = v63
        switch 4294967295 {
        case -8:
            if len(v3) > 0 {
                v3[0] = fi64
            }
        }
        _ = fi64
    }
    var v65 int64 = (v57 >> ((65535) & 63))
    if len(v3) > 0 {
        v3[0] = -27
    }
    v66 := GenS3{F0: v2, F1: (v57 << ((1) & 63)), F2: uint32(v54)}
    v65 += v57
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v57
    _ = v58
    _ = v59
    _ = v63
    _ = v65
    _ = v66
}
