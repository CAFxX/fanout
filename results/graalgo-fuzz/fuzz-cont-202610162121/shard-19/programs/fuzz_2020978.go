package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 string; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS7 struct { F0 int32; F1 int32; F2 uint8 }
type GenS8 struct { F0 int; F1 uint8 }
func genFn0(p0 uint16) int64 {
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := []int{}
        p0 += (p0 % ((p0) | 1))
        _ = fi4
        _ = v5
    }
    if false {
        if (!true) {
            for fi6 := 0; fi6 < 13; fi6++ {
                if true {
                    if (fi6 < (fi6 * (fi6 / ((42) | 1)))) {
                        for fi7 := 0; fi7 < 6; fi7++ {
                            var v8 string = "hello"
                            v9 := GenS1{F0: v8}
                            _ = fi7
                            _ = v8
                            _ = v9
                        }
                    }
                    var v10 int = fi6
                    _ = v10
                } else {
                    switch "hello" {
                    case "":
                        var v11 int16 = 42
                        _ = v11
                    case "a":
                        for fi12 := 0; fi12 < 10; fi12++ {
                            if ((2147483647 <= 0) && (18446744073709551615 != 1)) {
                                var v13 float64 = ((1e10 + (float64(3.14159) * (1e10 * 0.1))) - ((1e10 - 3.14159) + 0.1))
                                _ = v13
                            }
                            var v14 uint16 = p0
                            v15 := GenS2{F0: 42, F1: "foo bar", F2: 42}
                            _ = fi12
                            _ = v14
                            _ = v15
                        }
                    case "hello":
                        for fi16 := 0; fi16 < 7; fi16++ {
                            var v17 int8 = -1
                            switch 9223372036854775807 {
                            case 628:
                                v18 := GenS1{F0: ""}
                                v19 := GenS2{F0: -607, F1: "", F2: v17}
                                _ = v18
                                _ = v19
                            case -711:
                                if false {
                                    var v20 uint8 = 63
                                    _ = v20
                                } else {
                                    var v21 int32 = -1
                                    v21 += 1
                                    v22 := []int{fi6}
                                    var v23 int = (fi6 / (((fi16 + 802)) | 1))
                                    v23 -= (^(fi6))
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                            default:
                                v24 := []int{}
                                _ = v24
                            }
                            v25 := GenS1{F0: "foo bar"}
                            _ = fi16
                            _ = v17
                            _ = v25
                        }
                    default:
                        p0 = p0
                    }
                    fmt.Println(fi6, p0)
                }
                p0 = p0
                _ = fi6
            }
            switch ("a" + "a") {
            case "a":
                if (((p0 / (((^(p0))) | 1)) >= 65535) && false) {
                    var v26 uint8 = 116
                    v27 := []int{32767}
                    _ = v26
                    _ = v27
                } else {
                    p0 -= 1
                }
            }
            fmt.Println(p0)
        }
        switch "" {
        case "a":
            p0 -= (^(p0))
        case "":
            if false {
                for fi28 := 0; fi28 < 10; fi28++ {
                    fmt.Println(fi28)
                    _ = fi28
                }
                var v29 bool = false
                if v29 {
                    for fi30 := 0; fi30 < 3; fi30++ {
                        v31 := GenS3{F0: v29}
                        v32 := GenS3{F0: (!v29)}
                        var v33 float32 = 3.14159
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                }
                _ = v29
            }
        case "foo bar":
            fmt.Println(p0)
        }
        v34 := GenS1{F0: ("foo bar" + "a")}
        _ = v34
    }
    fmt.Println(p0)
    _ = p0
    return -9223372036854775808
}
func genFn1() uint16 {
    var v35 int32 = 97
    var v36 bool = true
    _ = v35
    _ = v36
    return 0
}

func main() {
    var v1 float64 = float64((0.0 * 1.5))
    v2 := []int{42, -1}
    var v3 int = 42
    var v37 float32 = 1.5
    fmt.Println(v3, v1, v1)
    switch (-(v3)) {
    case 32768:
        if (2147483647 <= 0) {
            for fi38 := 0; fi38 < 4; fi38++ {
                v37 -= 1.5
                v39 := GenS4{F0: (242 != 0)}
                var v40 uint16 = 551
                _ = fi38
                _ = v39
                _ = v40
            }
            var v41 float32 = float32(v1)
            _ = v41
        }
    default:
        if len(v2) > 0 {
            v2[0] = (v3 + v3)
        }
    }
    var v42 uint8 = 44
    var v43 uint16 = 90
    if (42 >= 127) {
        fmt.Println(v37)
        v44 := GenS5{F0: 2, F1: (v3 | (^(v3)))}
        _ = v44
    }
    for fi45 := 0; fi45 < 4; fi45++ {
        var v46 int16 = 42
        fmt.Println(v1)
        _ = fi45
        _ = v46
    }
    v43 += 32767
    var v47 float64 = ((((0.0 / (1.5 * (((v1 + 3.14159) + v1) - v1))) + v1) / (v1 - (1e10 * v1))) / (v1 / v1))
    var v48 bool = (true && true)
    v3 = v3
    v42 = (v42 & (v42 ^ 42))
    if (0 <= -387) {
        v48 = true
    } else {
        v49 := GenS6{F0: 4294967295, F1: 0, F2: 1}
        _ = v49
    }
    for fi50 := 0; fi50 < 4; fi50++ {
        if (!(315 <= 9223372036854775807)) {
            for fi51 := 0; fi51 < 10; fi51++ {
                switch ("foo bar" + "") {
                case "foo bar":
                    v42 += (v42 & uint8(v43))
                default:
                    if len(v2) > 0 {
                        v2[0] = -658
                    }
                }
                v52 := GenS7{F0: -2147483648, F1: -1, F2: v42}
                v53 := GenS8{F0: (fi50 | (v3 << ((255) & 63))), F1: (v42 & v42)}
                _ = fi51
                _ = v52
                _ = v53
            }
        } else {
            v43 -= 255
        }
        if v48 {
            _ = len(v2)
        }
        v1 = 0.1
        _ = fi50
    }
    if v48 {
        for fi54 := 0; fi54 < 9; fi54++ {
            v55 := GenS4{F0: v48}
            _ = len(v2)
            _ = fi54
            _ = v55
        }
        var v56 float32 = ((v37 + 3.14159) + float32((v47 + (3.14159 * 0.1))))
        _ = v56
    } else {
        for fi57 := 0; fi57 < 19; fi57++ {
            v42 = 141
            for fi58 := 0; fi58 < 19; fi58++ {
                switch ("\n\t" + "foo bar") {
                case "hello":
                    var v59 int = int(v42)
                    _ = v59
                default:
                    switch ((("foo bar" + ("\n\t" + ("" + "foo bar"))) + ("" + ("foo bar" + ((("\n\t" + "a") + "hello") + "hello")))) + "a") {
                    case "":
                        v60 := GenS8{F0: fi57, F1: (v42 / ((v42) | 1))}
                        fmt.Println(fi57, v3)
                        _ = v60
                    case "foo bar":
                        fmt.Println(v37)
                    case "\n\t":
                        v61 := GenS4{F0: ((v37 - (-2.25 - ((v37 * v37) * (0.0 + (3.14159 - float32(float64(float32((((0.0 * float64(1.5)) * v47) - 1.5))))))))) <= 0.0)}
                        _ = v61
                    default:
                        v43 = 255
                    }
                }
                var v62 int16 = 1000
                _ = fi58
                _ = v62
            }
            fmt.Println(v42, v42)
            _ = fi57
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v42
    _ = v43
    _ = v47
    _ = v48
}
