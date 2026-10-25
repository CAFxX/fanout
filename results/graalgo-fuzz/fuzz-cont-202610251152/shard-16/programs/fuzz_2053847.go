package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 string; F1 int8; F2 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 bool }

func main() {
    var v1 bool = true
    var v2 string = "hello"
    var v3 float32 = 0.1
    v4 := []uint32{42, 2147483648}
    var v5 uint32 = 555
    v6 := GenS1{F0: 42}
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint = 0
        var v9 string = "hello"
        switch fi7 {
        case -1:
            v10 := GenS2{F0: v1, F1: 86}
            _ = v10
        case 2147483647:
            var v11 int = fi7
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v12 float64 = 0.1
    for fi13 := 0; fi13 < 19; fi13++ {
        if false {
            if (682 <= 1) {
                v1 = (1 > 0)
                v1 = (v5 < 2147483647)
                v4 = append(v4, 2)
            }
            _ = len(v4)
        } else {
            v3 -= (v3 + ((v3 / ((float32(float64((v3 / v3))) + float32(0.1)) + (v3 * 3.14159))) - 3.14159))
            for fi14 := 0; fi14 < 7; fi14++ {
                fmt.Println(v12)
                v15 := GenS3{F0: ""}
                _ = fi14
                _ = v15
            }
            _ = len(v4)
        }
        fmt.Println(v2, fi13, fi13)
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := GenS4{F0: 4294967295}
            v18 := GenS4{F0: 963}
            v4 = append(v4, (v5 & v5))
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi13
    }
    v12 = float64(v3)
    var v19 int32 = -1
    v2 = v2
    switch 1 {
    case 662:
        switch v2 {
        case "\n\t":
            v20 := GenS4{F0: 197}
            _ = v20
        }
    case 2147483647:
        if (v2 < v2) {
            v1 = v1
            v19 += 332
            if len(v4) > 0 {
                v4[0] = (v5 / (((v5 >> ((0) & 63))) | 1))
            }
        } else {
            v2 = ("a" + v2)
            if (v3 <= 0.1) {
                v12 = v12
                v5 -= (v5 >> ((255) & 63))
                v21 := GenS3{F0: (v2 + v2)}
                _ = v21
            } else {
                if v1 {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        var v23 uint8 = 0
                        _ = fi22
                        _ = v23
                    }
                    v2 = v2
                }
            }
        }
    }
    v24 := GenS2{F0: v1, F1: 175}
    var v25 uint = uint(v5)
    v4 = append(v4, v5)
    _ = len(v4)
    var v26 int64 = -9223372036854775808
    v27 := GenS2{F0: v1, F1: 4294967295}
    switch 32767 {
    case -2147483648:
        for fi28 := 0; fi28 < 18; fi28++ {
            for fi29 := 0; fi29 < 11; fi29++ {
                if v1 {
                    switch 32767 {
                    case -930:
                        v19 += (v19 / (((v19 / (((v19 / (((v19 & 0)) | 1))) | 1))) | 1))
                    case -239:
                        v4 = append(v4, v5)
                    case 128:
                        fmt.Println(v12, v26, fi29)
                    }
                }
                switch "\n\t" {
                case "hello":
                    if (v1 && v1) {
                        if (v26 < v26) {
                            switch v2 {
                            case "hello":
                                fmt.Println(v3, v5)
                            default:
                                v30 := GenS5{F0: 255, F1: (^(v5))}
                                _ = v30
                            }
                            v4 = append(v4, 0)
                            v4 = append(v4, v5)
                        } else {
                            var v31 int32 = (v19 / (((v19 ^ v19)) | 1))
                            if false {
                                v12 += (v12 / (v12 + v12))
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = uint32(v25)
                                }
                                switch (fi28 >> ((v25) & 63)) {
                                case 9223372036854775807:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case 324:
                                    if (v1 && v1) {
                                        _ = len(v4)
                                        var v32 uint8 = 5
                                        _ = v32
                                    } else {
                                        v19 = v19
                                    }
                                }
                                if (!v1) {
                                    if ((!(true && (1 == 533))) || true) {
                                        v33 := GenS4{F0: 533}
                                        v31 = (v31 * v31)
                                        _ = v33
                                    }
                                    if v1 {
                                        v34 := GenS6{F0: v2, F1: 0, F2: ""}
                                        v5 -= v5
                                        _ = v34
                                    } else {
                                        switch (v2 + v2) {
                                        case "":
                                            if (v1 || true) {
                                                var v35 uint64 = 4294967295
                                                var v36 float64 = (((0.1 / 1.0) + v12) - v12)
                                                switch -740 {
                                                case 205:
                                                    switch int(v26) {
                                                    case 4294967295:
                                                        var v37 uint16 = 731
                                                        _ = v37
                                                    case -724:
                                                        if (!(!v1)) {
                                                            switch ((v2 + "a") + (v2 + (v2 + v2))) {
                                                            case "hello":
                                                                v4 = append(v4, (v5 % (((v5 ^ 268)) | 1)))
                                                            case "":
                                                                v12 = v36
                                                            }
                                                            if false {
                                                                fmt.Println(v3, v3)
                                                                var v38 int = (fi28 & (fi29 % (((-(fi28))) | 1)))
                                                                v5 += 0
                                                                _ = v38
                                                            } else {
                                                                v5 = v5
                                                            }
                                                        }
                                                    case -521:
                                                        if v1 {
                                                            if (1.5 < ((float64(float32((v36 / (v36 / float64(3.14159))))) * (((v12 * v12) - v36) - v36)) * ((float64(v3) + -2.25) * v12))) {
                                                                if v1 {
                                                                    v36 += ((float64((v3 * (1.5 * (1e10 - ((0.1 + v3) * v3))))) - ((v12 + 0.1) / v12)) + float64(float32((v36 + v36))))
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                                v39 := GenS6{F0: (v2 + (v2 + "hello")), F1: -5, F2: "a"}
                                                                v36 += 0.1
                                                                _ = v39
                                                            }
                                                            var v40 int = int(v19)
                                                            if (!(v1 || (!false))) {
                                                                if (771 <= (v25 / ((v25) | 1))) {
                                                                    v40 = (fi29 / (((fi29 * 32767)) | 1))
                                                                    if v1 {
                                                                        v41 := GenS3{F0: (v2 + (("foo bar" + "foo bar") + v2))}
                                                                        _ = v41
                                                                    } else {
                                                                        if len(v4) > 0 {
                                                                            v4[0] = (v5 | (v5 >> ((42) & 63)))
                                                                        }
                                                                        var v42 uint32 = (v5 + 2147483647)
                                                                        v43 := GenS4{F0: 18446744073709551615}
                                                                        _ = v42
                                                                        _ = v43
                                                                    }
                                                                }
                                                            } else {
                                                                if len(v4) > 0 {
                                                                    v4[0] = v5
                                                                }
                                                                var v44 uint16 = 615
                                                                _ = v44
                                                            }
                                                            _ = v40
                                                        }
                                                    }
                                                }
                                                _ = v35
                                                _ = v36
                                            }
                                        case "foo bar":
                                            if len(v4) > 0 {
                                                v4[0] = (^(v5))
                                            }
                                        case "\n\t":
                                            var v45 int16 = 1
                                            _ = v45
                                        }
                                        v46 := GenS7{F0: 391}
                                        _ = v46
                                    }
                                }
                            }
                            v3 += 0.0
                            _ = v31
                        }
                    }
                }
                _ = fi29
            }
            var v47 uint = v25
            _ = fi28
            _ = v47
        }
    case 561:
        v1 = v1
    case 32767:
        fmt.Println(v2, v2, v12)
    default:
        v2 = v2
    }
    if true {
        var v48 uint64 = uint64(v25)
        if (((-2.25 - 0.0) >= v12) && true) {
            v4 = append(v4, 2)
            v49 := GenS8{F0: 2, F1: (true && v1)}
            _ = v49
        } else {
            if (1.5 == 1e10) {
                var v50 int64 = (-(v26))
                _ = v50
            }
            switch v2 {
            case "":
                _ = len(v4)
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v48
    } else {
        for fi51 := 0; fi51 < 19; fi51++ {
            v4 = append(v4, (v5 << (((^(v25))) & 63)))
            _ = fi51
        }
        if (481 == 4294967295) {
            v4 = append(v4, v5)
        }
        for fi52 := 0; fi52 < 11; fi52++ {
            var v53 uint8 = 42
            v26 = (v26 + (v26 >> (((v25 << ((v25) & 63))) & 63)))
            for fi54 := 0; fi54 < 17; fi54++ {
                var v55 uint64 = 1
                var v56 uint16 = 2
                _ = fi54
                _ = v55
                _ = v56
            }
            _ = fi52
            _ = v53
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
