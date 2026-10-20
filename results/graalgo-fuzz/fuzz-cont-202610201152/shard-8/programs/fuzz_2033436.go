package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 int64; F1 uint32; F2 string }
func genFn0(p0 int8) int64 {
    if false {
        p0 = p0
        p0 = p0
        var v7 string = "foo bar"
        _ = v7
    }
    var v8 string = "foo bar"
    _ = p0
    _ = v8
    return -1
}
func genFn1(p0 bool) uint32 {
    var v9 float32 = 0.0
    switch "\n\t" {
    case "":
        v9 += ((float32((3.14159 / 3.14159)) + 3.14159) - (v9 / (float32((1e10 * (1.5 + float64(v9)))) / float32((((((3.14159 - (((float64(((((v9 - v9) - 3.14159) * v9) - (float32(3.14159) - v9))) * 3.14159) * 1.5) + 1.5)) * 0.0) / 1e10) + 0.0) - -2.25)))))
    default:
        switch "foo bar" {
        case "foo bar":
            p0 = (!p0)
        }
    }
    fmt.Println(v9, v9)
    v10 := GenS1{F0: 42}
    _ = p0
    _ = v10
    _ = v9
    return 586
}

func main() {
    v1 := []uint32{545}
    var v2 int16 = -32768
    v3 := make(map[int]string)
    var v4 float32 = -2.25
    var v5 int16 = (v2 % ((v2) | 1))
    var v6 int8 = 127
    v11 := GenS1{F0: 18446744073709551615}
    v12 := GenS1{F0: 18446744073709551615}
    v2 -= 42
    var v13 float64 = (float64(-2.25) / -2.25)
    var v14 uint16 = 1
    fmt.Println(v2, v6)
    if (true || false) {
        for fi15 := 0; fi15 < 20; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                if (v4 <= ((v4 * v4) / v4)) {
                    v13 += ((float64(((-2.25 + v4) / (v4 - v4))) * (v13 * float64(1e10))) * (3.14159 + float64(v4)))
                } else {
                    if (!true) {
                        var v17 float64 = float64((v4 * v4))
                        v18 := GenS2{F0: (("hello" + "hello") + "")}
                        fmt.Println(v5, v17, v4)
                        _ = len(v1)
                        _ = len(v1)
                        _ = v17
                        _ = v18
                    } else {
                        fmt.Println(v5, fi15, v5)
                        if true {
                            if ((v6 & (v6 % ((76) | 1))) > 1) {
                                v19 := GenS1{F0: 4294967295}
                                _ = v19
                            } else {
                                var v20 uint = 18446744073709551615
                                _ = v20
                            }
                            var v21 string = "a"
                            var v22 float32 = (v4 * ((0.0 - (1e10 / 0.1)) * 0.1))
                            v22 += v22
                            _ = v21
                            _ = v22
                        } else {
                            var v23 float64 = 1.5
                            v24 := GenS3{F0: (v5 & v5)}
                            if ((fi16 ^ fi16) > fi15) {
                                v4 += (3.14159 / (1.5 * v4))
                                if (true && (((("hello" + ("a" + "\n\t")) + "") < (("" + "") + (("\n\t" + ("hello" + "\n\t")) + ("foo bar" + "")))) && true)) {
                                    if (true || true) {
                                        v2 = (-(v2))
                                        var v25 uint8 = 255
                                        v3[9223372036854775807] = "foo bar"
                                        _ = v25
                                    } else {
                                        if (false || true) {
                                            var v26 uint32 = 65535
                                            _ = v26
                                        }
                                    }
                                    delete(v3, (fi16 & 9223372036854775807))
                                }
                                v13 = (v13 + (float64(0.1) * v23))
                            } else {
                                v27 := GenS1{F0: 1}
                                fmt.Println(fi15)
                                _ = v27
                            }
                            _ = v23
                            _ = v24
                        }
                    }
                    var v28 string = ("\n\t" + "\n\t")
                    v29 := GenS1{F0: uint64(v14)}
                    _ = v28
                    _ = v29
                }
                if len(v1) > 0 {
                    v1[0] = 998
                }
                v5 = v2
                _ = fi16
            }
            var v30 float64 = v13
            for fi31 := 0; fi31 < 3; fi31++ {
                if false {
                    var v32 int = 2147483647
                    if (!true) {
                        v33 := GenS3{F0: (v2 & 12)}
                        v34 := GenS4{F0: 255, F1: 285}
                        v2 -= 1000
                        _ = v33
                        _ = v34
                    } else {
                        v35 := GenS4{F0: 0, F1: (^(v14))}
                        _ = v35
                    }
                    if false {
                        var v36 int32 = 42
                        _ = v36
                    }
                    v30 = (0.0 * v30)
                    fmt.Println(v2, v2)
                    _ = v32
                }
                v6 -= v6
                if true {
                    if ("a" >= "\n\t") {
                        v37 := GenS2{F0: "hello"}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v37
                    }
                    fmt.Println(v6, v5, v5)
                    var v38 bool = (-1 < 535)
                    if (42 > -1) {
                        var v39 float32 = 0.0
                        _ = v39
                    } else {
                        v38 = (213 <= 128)
                    }
                    v38 = v38
                    _ = v38
                }
                _ = fi31
            }
            _ = fi15
            _ = v30
        }
    }
    var v40 bool = ("\n\t" != "")
    v5 = 1
    var v41 uint = 42
    switch "foo bar" {
    case "a":
        v42 := GenS3{F0: v2}
        _ = v42
    case "hello":
        for fi43 := 0; fi43 < 18; fi43++ {
            fmt.Println(v5, v41)
            var v44 string = ("hello" + ("a" + "\n\t"))
            var v45 int32 = 42
            _ = fi43
            _ = v44
            _ = v45
        }
    case "foo bar":
        v6 += v6
    default:
        v6 = -20
    }
    if v40 {
        switch 42 {
        case -594:
            delete(v3, 9223372036854775807)
        case 232:
            for fi46 := 0; fi46 < 15; fi46++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi46
            }
        case 0:
            v14 = v14
        }
    } else {
        fmt.Println(v13)
        switch ("" + "") {
        case "\n\t":
            v6 -= 100
        case "foo bar":
            for fi47 := 0; fi47 < 8; fi47++ {
                fmt.Println(v6, v6)
                if false {
                    switch fi47 {
                    case 0:
                        if v40 {
                            var v48 float32 = 3.14159
                            _ = v48
                        } else {
                            var v49 int = (fi47 + fi47)
                            _ = v49
                        }
                    case -2147483648:
                        v50 := GenS3{F0: (^(v2))}
                        _ = v50
                    }
                }
                v13 = v13
                _ = fi47
            }
        default:
            fmt.Println(v41, v14, v14)
        }
    }
    var v51 bool = (v41 > v41)
    switch "foo bar" {
    case "hello":
        for fi52 := 0; fi52 < 11; fi52++ {
            if (105 < 255) {
                v40 = (!(v4 <= 0.1))
                switch (fi52 - fi52) {
                case -2147483648:
                    _ = len(v3)
                case -9223372036854775808:
                    var v53 float64 = v13
                    _ = v53
                case 32768:
                    v14 = v14
                }
                switch "\n\t" {
                case "foo bar":
                    _ = len(v1)
                default:
                    v1 = append(v1, 367)
                }
            }
            for fi54 := 0; fi54 < 2; fi54++ {
                var v55 int64 = -9223372036854775808
                fmt.Println(fi52, v4, v6)
                var v56 uint32 = 346
                var v57 float32 = v4
                _ = fi54
                _ = v55
                _ = v56
                _ = v57
            }
            for fi58 := 0; fi58 < 17; fi58++ {
                v4 -= (1.5 - (float32(v13) - v4))
                if v40 {
                    fmt.Println(fi58, v6)
                    _ = len(v3)
                    _ = len(v1)
                }
                _ = fi58
            }
            _ = fi52
        }
    case "a":
        var v59 uint64 = 2
        _ = v59
    case "foo bar":
        for fi60 := 0; fi60 < 1; fi60++ {
            v61 := GenS5{F0: genFn0(100), F1: 65535, F2: ("\n\t" + "hello")}
            v62 := GenS3{F0: v5}
            _ = fi60
            _ = v61
            _ = v62
        }
    }
    fmt.Println(v5, v51)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v51
    _ = v6
}
