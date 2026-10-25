package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 string }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 string; F1 uint8; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int8, p1 string, p2 uint64) uint8 {
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := []int{fi7, fi7, fi7, 9223372036854775807}
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: 32767, F1: -712, F2: (p1 + p1)}
            var v11 float64 = 3.14159
            switch p1 {
            case "":
                fmt.Println(fi9, fi9)
            case "foo bar":
                v11 = 1e10
            case "hello":
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            default:
                p2 -= (p2 / (((p2 ^ p2)) | 1))
            }
            fmt.Println(fi7)
            p1 = "foo bar"
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 int8 = p0
            switch fi12 {
            case -661:
                var v14 bool = true
                if ((65535 <= 42) && false) {
                    var v15 string = (p1 + p1)
                    _ = v15
                }
                var v16 bool = (255 >= 380)
                _ = v14
                _ = v16
            case 1:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            case -710:
                fmt.Println(fi7)
            }
            _ = fi12
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v18 := GenS1{F0: fi17, F1: 42, F2: p1}
        v19 := []int{(-(fi17))}
        _ = fi17
        _ = v18
        _ = v19
    }
    switch (("a" + p1) + ("\n\t" + "foo bar")) {
    case "\n\t":
        p1 = "foo bar"
    case "hello":
        v20 := GenS1{F0: -691, F1: -32768, F2: ("foo bar" + p1)}
        _ = v20
    case "":
        var v21 int8 = p0
        _ = v21
    default:
        p2 = (p2 * (p2 + (^(p2))))
    }
    var v22 uint8 = 128
    for fi23 := 0; fi23 < 3; fi23++ {
        v22 += 107
        if (((^(v22)) >= 1) || (-2.25 <= (0.1 * float64(3.14159)))) {
            switch p1 {
            case "foo bar":
                switch p1 {
                case "hello":
                    var v24 uint64 = uint64(v22)
                    _ = v24
                case "\n\t":
                    var v25 int16 = -1
                    _ = v25
                }
            case "hello":
                p1 = p1
            case "a":
                var v26 int32 = -1
                _ = v26
            }
            v27 := []int{(-(fi23)), 32768, (fi23 | (fi23 | (fi23 >> ((4294967295) & 63)))), fi23}
            if (false || true) {
                if len(v27) >= 2 {
                    v27 = v27[1:len(v27)]
                }
            }
            _ = v27
        }
        var v28 bool = false
        _ = fi23
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{}
    v3 := []uint32{}
    v4 := make(map[int]bool)
    var v5 int32 = -1
    v6 := make(map[string]int)
    v5 -= (v5 * (v5 + (v5 | 2147483647)))
    fmt.Println(v5)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 += v5
    if len(v2) > 0 {
        v2[0] = 42
    }
    var v29 uint = 18446744073709551615
    v29 += 42
    v29 -= v29
    var v30 uint64 = 669
    var v31 uint16 = 285
    if false {
        if false {
            v32 := v1[883]
            _ = v32
        } else {
            v29 = (^(v29))
        }
        switch 726 {
        case -2147483648:
            switch "hello" {
            case "hello":
                if (true || (1 < 164)) {
                    switch 422 {
                    case 0:
                        fmt.Println(v5, v5, v30)
                    case 4294967295:
                        var v33 uint64 = 4294967295
                        _ = v33
                    default:
                        if (9223372036854775807 < 128) {
                            if false {
                                switch -1 {
                                case 128:
                                    v31 = (v31 % (((v31 << ((v29) & 63))) | 1))
                                case 0:
                                    var v34 uint32 = 2147483648
                                    _ = v34
                                default:
                                    v31 += v31
                                }
                            } else {
                                v29 = 4294967295
                                if (true || (((!false) && ((v5 >> (((v29 / (((v29 | (v29 * uint(v30)))) | 1))) & 63)) != (v5 ^ 171))) || (!true))) {
                                    var v35 int32 = (v5 / (((v5 * (v5 + (v5 & -764)))) | 1))
                                    _ = v35
                                } else {
                                    var v36 float32 = 1e10
                                    for fi37 := 0; fi37 < 17; fi37++ {
                                        if (true && (-112 <= -128)) {
                                            fmt.Println(v30, v29, v31)
                                        }
                                        v38 := GenS2{F0: -9223372036854775808, F1: "\n\t"}
                                        _ = fi37
                                        _ = v38
                                    }
                                    v39 := v6["hello"]
                                    _ = v36
                                    _ = v39
                                }
                                v40 := GenS3{F0: "a", F1: genFn0(37, "\n\t", (v30 - 60)), F2: v29}
                                _ = v40
                            }
                            if false {
                                v41 := GenS4{F0: 94}
                                _ = v41
                            } else {
                                var v42 uint = 65535
                                v1[440] = 32768
                                var v43 uint = 18446744073709551615
                                _ = v42
                                _ = v43
                            }
                        } else {
                            v31 = (v31 - (^(v31)))
                            v31 = v31
                        }
                    }
                    var v44 int32 = (v5 * 627)
                    _ = v44
                } else {
                    delete(v1, 2147483647)
                    switch -2147483648 {
                    case 127:
                        fmt.Println(v5)
                    default:
                        for fi45 := 0; fi45 < 15; fi45++ {
                            for fi46 := 0; fi46 < 18; fi46++ {
                                fmt.Println(v30, v31, fi46)
                                v5 += (v5 % ((410) | 1))
                                fmt.Println(v31, v29)
                                _ = fi46
                            }
                            _ = fi45
                        }
                    }
                }
            case "a":
                fmt.Println(v5)
            case "":
                var v47 uint = uint(v30)
                _ = v47
            }
        }
    }
    for fi48 := 0; fi48 < 20; fi48++ {
        if true {
            fmt.Println(v29, v29)
            if true {
                for fi49 := 0; fi49 < 20; fi49++ {
                    var v50 int = (fi48 - 1)
                    v29 += v29
                    switch (fi49 ^ -518) {
                    case 0:
                        v51 := GenS1{F0: -37, F1: 42, F2: "\n\t"}
                        _ = v51
                    case 42:
                        fmt.Println(fi49)
                        v52 := GenS4{F0: fi48}
                        _ = v52
                    default:
                        var v53 uint8 = 255
                        _ = v53
                    }
                    _ = fi49
                    _ = v50
                }
                v54 := GenS5{F0: genFn0(1, "foo bar", v30)}
                _ = v54
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if true {
                    v30 += (v30 / ((0) | 1))
                } else {
                    var v55 int = (fi48 ^ (^(fi48)))
                    v29 = 18446744073709551615
                    _ = v55
                }
            }
            fmt.Println(v31, v31, v29)
        }
        fmt.Println(v31, v29)
        var v56 uint32 = 2
        _ = fi48
        _ = v56
    }
    fmt.Println(v5)
    if (!(!(-1 > 173))) {
        var v57 uint8 = 100
        _ = v57
    } else {
        fmt.Println(v5)
        v5 = (^(v5))
    }
    for fi58 := 0; fi58 < 4; fi58++ {
        var v59 float32 = ((1e10 + 0.0) + (3.14159 + 0.0))
        var v60 bool = true
        _ = fi58
        _ = v59
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
