package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint }
type GenS2 struct { F0 int64; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 int; F1 int8; F2 uint32 }
func genFn0() int8 {
    var v7 uint8 = 1
    fmt.Println(v7)
    v8 := GenS1{F0: 128, F1: 2, F2: 2}
    fmt.Println(v7)
    _ = v7
    _ = v8
    return -40
}
func genFn1(p0 uint64, p1 string, p2 uint16) string {
    var v9 int32 = -903
    var v10 uint64 = p0
    v11 := []int{-515, 497, 9223372036854775807}
    v12 := GenS1{F0: -300, F1: 3, F2: 45}
    if (genFn0() >= genFn0()) {
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 uint16 = p2
            _ = fi13
            _ = v14
        }
        v9 += (v9 / ((545) | 1))
    } else {
        v10 = (v10 - 882)
        var v15 int8 = -18
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return ("" + (p1 + ("" + p1)))
}

func main() {
    var v1 float64 = float64(1e10)
    var v2 bool = (!true)
    v3 := []uint32{415, 1, 4294967295}
    v4 := []uint32{0, 809, 511, 0}
    var v5 int8 = 0
    var v6 uint = 219
    for fi16 := 0; fi16 < 2; fi16++ {
        switch (-(fi16)) {
        case 9223372036854775807:
            for fi17 := 0; fi17 < 19; fi17++ {
                if v2 {
                    v18 := GenS1{F0: (fi17 << (((v6 / ((v6) | 1))) & 63)), F1: 676, F2: v6}
                    _ = v18
                }
                var v19 int16 = -32768
                if (!((2 >= 2147483647) || (("a" > "a") && v2))) {
                    v1 = float64(0.1)
                } else {
                    v20 := GenS2{F0: -509, F1: 544, F2: v5}
                    switch 2147483647 {
                    case 928:
                        var v21 int64 = 806
                        _ = v21
                    case 601:
                        fmt.Println(v2, fi17)
                    default:
                        v2 = (true || v2)
                    }
                    _ = v20
                }
                _ = fi17
                _ = v19
            }
        default:
            if (!true) {
                if v2 {
                    for fi22 := 0; fi22 < 11; fi22++ {
                        v2 = true
                        switch "foo bar" {
                        case "\n\t":
                            switch (((("\n\t" + ((("" + "hello") + "foo bar") + ("\n\t" + "a"))) + "hello") + "hello") + "a") {
                            case "\n\t":
                                var v23 string = ("hello" + "")
                                _ = v23
                            case "hello":
                                fmt.Println(fi22)
                            case "":
                                if v2 {
                                    var v24 float64 = float64((3.14159 - float32(v1)))
                                    _ = v24
                                } else {
                                    v5 = genFn0()
                                    var v25 bool = (fi16 != 128)
                                    v2 = true
                                    _ = v25
                                }
                            }
                        case "a":
                            v26 := GenS2{F0: 0, F1: 18446744073709551615, F2: v5}
                            _ = v26
                        case "":
                            fmt.Println(fi16, v2)
                        }
                        v27 := GenS3{F0: (^(v6))}
                        _ = fi22
                        _ = v27
                    }
                } else {
                    v6 += v6
                    if (2 == 0) {
                        v1 -= 0.0
                    } else {
                        var v28 int8 = genFn0()
                        _ = v28
                    }
                }
                if (0 != -9223372036854775808) {
                    v6 = (v6 ^ (v6 >> ((v6) & 63)))
                }
                for fi29 := 0; fi29 < 10; fi29++ {
                    v6 = v6
                    var v30 int8 = 1
                    _ = fi29
                    _ = v30
                }
            } else {
                for fi31 := 0; fi31 < 14; fi31++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v32 string = (("\n\t" + "hello") + "\n\t")
                    _ = fi31
                    _ = v32
                }
                for fi33 := 0; fi33 < 9; fi33++ {
                    v34 := GenS4{F0: 0}
                    if (float32((((-2.25 + 0.1) / (v1 / 3.14159)) - ((-2.25 * float64(0.1)) / 3.14159))) <= float32(float64(0.0))) {
                        v35 := GenS5{F0: (^(v6)), F1: 414}
                        v2 = (!(42 >= 255))
                        _ = v35
                    } else {
                        var v36 uint32 = 0
                        _ = v36
                    }
                    if v2 {
                        v5 += (v5 - (v5 >> ((v6) & 63)))
                    }
                    _ = fi33
                    _ = v34
                }
                for fi37 := 0; fi37 < 8; fi37++ {
                    var v38 uint16 = 758
                    var v39 uint = 255
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi16
    }
    switch "foo bar" {
    case "\n\t":
        if (!false) {
            v3 = append(v3, 65535)
            if v2 {
                if (-9223372036854775808 >= -1) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    var v40 uint = (v6 / (((^(v6))) | 1))
                    var v41 uint8 = 127
                    v6 -= v40
                    _ = v40
                    _ = v41
                }
            } else {
                _ = len(v4)
                var v42 int = 458
                _ = v42
            }
            if ((652 < 154) || v2) {
                v1 -= v1
                var v43 uint32 = 65535
                _ = v43
            } else {
                v44 := GenS6{F0: 2, F1: 2}
                if (false && (("a" + "") != "\n\t")) {
                    var v45 int64 = 1
                    _ = v45
                }
                if (!v2) {
                    v46 := GenS2{F0: -1, F1: 287, F2: (^(v5))}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v46
                } else {
                    v47 := GenS7{F0: 128, F1: -1, F2: 4294967295}
                    var v48 uint16 = 0
                    _ = v47
                    _ = v48
                }
                _ = v44
            }
        }
    case "hello":
        for fi49 := 0; fi49 < 5; fi49++ {
            if (fi49 == 32768) {
                fmt.Println(v5, v1)
                fmt.Println(v6, fi49)
            }
            v2 = true
            _ = fi49
        }
    }
    fmt.Println(v1, v2, v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v50 uint = v6
    if len(v4) > 0 {
        v4[0] = 297
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v51 bool = v2
    var v52 bool = v2
    fmt.Println(v51, v52)
    var v53 int = 42
    switch (v53 * v53) {
    case -9223372036854775808:
        v6 = 2
    case 898:
        var v54 uint8 = 42
        _ = v54
    }
    for fi55 := 0; fi55 < 14; fi55++ {
        v56 := GenS1{F0: fi55, F1: 1, F2: v50}
        for fi57 := 0; fi57 < 18; fi57++ {
            var v58 int8 = (v5 % (((v5 | v5)) | 1))
            var v59 uint64 = 756
            v59 = (v59 | (v59 * v59))
            _ = fi57
            _ = v58
            _ = v59
        }
        _ = fi55
        _ = v56
    }
    var v60 bool = ("a" != "a")
    var v61 string = ""
    var v62 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v6
    _ = v60
    _ = v61
    _ = v62
}
