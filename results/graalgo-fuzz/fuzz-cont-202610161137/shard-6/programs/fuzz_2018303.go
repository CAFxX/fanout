package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int32, p1 uint8) bool {
    fmt.Println(p1, p0)
    switch 1 {
    case -9223372036854775808:
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    return false
}
func genFn1() uint64 {
    if false {
        var v6 uint16 = 345
        var v7 uint64 = 18446744073709551615
        _ = v6
        _ = v7
    }
    v8 := []int{-438, 32767, -605}
    var v9 float64 = float64(3.14159)
    _ = len(v8)
    var v10 uint16 = 65535
    _ = v10
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 2
    var v2 int8 = 100
    var v3 uint64 = (v1 >> ((255) & 63))
    v4 := []int{-1, 32768, -866}
    v5 := []string{""}
    var v11 uint32 = 1
    if true {
        fmt.Println(v11, v11)
    } else {
        var v12 int64 = 42
        v13 := GenS1{F0: true}
        v3 += v1
        _ = v12
        _ = v13
    }
    v11 -= 581
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, "\n\t")
    v14 := GenS1{F0: true}
    v2 = -92
    fmt.Println(v11)
    var v15 int8 = v2
    var v16 bool = (((v1 & v3) <= v3) || true)
    switch "" {
    case "\n\t":
        for fi17 := 0; fi17 < 19; fi17++ {
            for fi18 := 0; fi18 < 5; fi18++ {
                v19 := GenS1{F0: (510 <= -9223372036854775808)}
                if v16 {
                    var v20 bool = (!v16)
                    v20 = ((v15 << ((4294967295) & 63)) == (-(v2)))
                    _ = v20
                } else {
                    var v21 int32 = -583
                    var v22 uint8 = 1
                    _ = v21
                    _ = v22
                }
                _ = fi18
                _ = v19
            }
            switch fi17 {
            case 2147483647:
                v1 += genFn1()
            }
            v11 = (v11 ^ 2147483647)
            _ = fi17
        }
    case "hello":
        var v23 int64 = int64(v11)
        _ = v23
    default:
        if len(v5) > 0 {
            v5[0] = (("\n\t" + ("a" + (("" + "hello") + "\n\t"))) + ("hello" + "foo bar"))
        }
    }
    var v24 uint = 1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (v2 >= (v15 << ((v24) & 63))) {
        if v16 {
            for fi25 := 0; fi25 < 5; fi25++ {
                var v26 uint32 = v11
                _ = fi25
                _ = v26
            }
            if false {
                if (890 != 2) {
                    if v16 {
                        v27 := GenS1{F0: ((2147483647 <= 32768) || true)}
                        fmt.Println(v3)
                        if (!false) {
                            if len(v5) > 0 {
                                v5[0] = ""
                            }
                        } else {
                            if (!v16) {
                                _ = len(v4)
                            }
                        }
                        _ = v27
                    } else {
                        v28 := GenS1{F0: false}
                        v3 = genFn1()
                        var v29 uint16 = 0
                        _ = v28
                        _ = v29
                    }
                    fmt.Println(v24)
                } else {
                    for fi30 := 0; fi30 < 9; fi30++ {
                        if (!(!v16)) {
                            for fi31 := 0; fi31 < 8; fi31++ {
                                var v32 float32 = 0.1
                                switch ("a" + ("a" + ("hello" + "\n\t"))) {
                                case "a":
                                    v33 := GenS1{F0: (-904 == 0)}
                                    _ = v33
                                default:
                                    fmt.Println(v15, fi31)
                                }
                                if true {
                                    var v34 int32 = 2147483647
                                    fmt.Println(fi30, v11, v15)
                                    var v35 bool = true
                                    _ = v34
                                    _ = v35
                                }
                                _ = fi31
                                _ = v32
                            }
                        } else {
                            switch "a" {
                            case "\n\t":
                                fmt.Println(v15, v24, v16)
                            case "":
                                var v36 int8 = (v2 - v2)
                                _ = v36
                            }
                        }
                        var v37 uint8 = 2
                        var v38 uint = 1
                        _ = fi30
                        _ = v37
                        _ = v38
                    }
                }
                var v39 uint32 = v11
                _ = v39
            } else {
                if (1.5 >= (1.5 - (3.14159 * 0.1))) {
                    v3 -= (v3 * genFn1())
                    v40 := GenS2{F0: v24, F1: v3}
                    _ = v40
                }
            }
            v41 := GenS3{F0: -32768, F1: 42}
            _ = v41
        } else {
            v1 = (v3 % ((v3) | 1))
            if len(v4) > 0 {
                v4[0] = -260
            }
        }
        if false {
            switch ("foo bar" + ("a" + ("a" + "\n\t"))) {
            case "foo bar":
                v42 := GenS2{F0: (^(v24)), F1: v1}
                _ = v42
            }
        } else {
            v24 += (v24 - (^(v24)))
            var v43 uint64 = (v3 >> ((v24) & 63))
            v3 = (v1 % (((v43 * (v1 % (((v43 >> ((v24) & 63))) | 1)))) | 1))
            _ = v43
        }
        var v44 float64 = (1.5 + 3.14159)
        _ = v44
    } else {
        for fi45 := 0; fi45 < 2; fi45++ {
            for fi46 := 0; fi46 < 8; fi46++ {
                var v47 int = fi46
                var v48 int32 = 2147483647
                v49 := GenS3{F0: 1000, F1: 4294967295}
                _ = fi46
                _ = v47
                _ = v48
                _ = v49
            }
            _ = fi45
        }
    }
    for fi50 := 0; fi50 < 9; fi50++ {
        var v51 int32 = int32(v24)
        for fi52 := 0; fi52 < 2; fi52++ {
            var v53 uint64 = 2
            _ = fi52
            _ = v53
        }
        _ = fi50
        _ = v51
    }
    for fi54 := 0; fi54 < 3; fi54++ {
        fmt.Println(v11, v11)
        if v16 {
            var v55 int64 = 42
            var v56 uint = v24
            _ = v55
            _ = v56
        } else {
            if true {
                switch int(v11) {
                case 4294967295:
                    switch int(v24) {
                    case 128:
                        var v57 int16 = -32768
                        _ = v57
                    }
                case 127:
                    for fi58 := 0; fi58 < 16; fi58++ {
                        v16 = v16
                        var v59 string = ("" + "")
                        v16 = v16
                        _ = fi58
                        _ = v59
                    }
                case -924:
                    var v60 string = ""
                    _ = v60
                }
            }
            var v61 float32 = (float32(((3.14159 * -2.25) * 0.1)) / float32(0.1))
            _ = v61
        }
        var v62 uint = v24
        _ = fi54
        _ = v62
    }
    _ = len(v5)
    if v16 {
        if v16 {
            fmt.Println(v2, v11)
            for fi63 := 0; fi63 < 14; fi63++ {
                if (v11 == v11) {
                    if ("hello" < "\n\t") {
                        v16 = false
                    } else {
                        for fi64 := 0; fi64 < 7; fi64++ {
                            var v65 uint = v24
                            var v66 float32 = float32(float64(1.5))
                            v67 := GenS4{F0: 65535, F1: 946}
                            _ = fi64
                            _ = v65
                            _ = v66
                            _ = v67
                        }
                    }
                    for fi68 := 0; fi68 < 6; fi68++ {
                        switch "a" {
                        case "a":
                            var v69 uint32 = (v11 ^ v11)
                            _ = v69
                        case "":
                            if (("\n\t" < "a") && v16) {
                                v16 = true
                            }
                        }
                        _ = fi68
                    }
                }
                var v70 bool = v16
                for fi71 := 0; fi71 < 3; fi71++ {
                    var v72 bool = false
                    if (true || ((v15 * -54) < 41)) {
                        if len(v4) > 0 {
                            v4[0] = 127
                        }
                        v73 := GenS1{F0: v16}
                        _ = v73
                    }
                    switch "" {
                    case "":
                        if len(v5) > 0 {
                            v5[0] = "\n\t"
                        }
                    }
                    _ = fi71
                    _ = v72
                }
                _ = fi63
                _ = v70
            }
            switch -1 {
            case 32768:
                if (v16 && (v15 == v2)) {
                    switch 4294967295 {
                    case -9223372036854775808:
                        fmt.Println(v16, v3, v15)
                    }
                    _ = len(v4)
                    var v74 int32 = -1
                    _ = v74
                } else {
                    if true {
                        v3 -= (v1 % (((v1 ^ v3)) | 1))
                        for fi75 := 0; fi75 < 3; fi75++ {
                            v76 := GenS1{F0: ((1 >= 0) || v16)}
                            var v77 uint16 = 468
                            if true {
                                v15 = 1
                                if false {
                                    var v78 uint8 = 255
                                    _ = v78
                                }
                            } else {
                                v79 := GenS5{F0: 148, F1: 179}
                                var v80 uint32 = v11
                                _ = v79
                                _ = v80
                            }
                            _ = fi75
                            _ = v76
                            _ = v77
                        }
                    }
                }
            case 4294967295:
                for fi81 := 0; fi81 < 2; fi81++ {
                    switch (fi81 + fi81) {
                    case 32767:
                        v2 += -128
                    }
                    v4 = append(v4, fi81)
                    _ = fi81
                }
            }
        } else {
            v1 += 410
            switch 1 {
            case 4294967295:
                _ = len(v5)
            case 32767:
                v11 = (v11 ^ (v11 + 1))
            }
            var v82 bool = false
            _ = v82
        }
        if len(v5) > 0 {
            v5[0] = "a"
        }
    } else {
        if (65535 == v24) {
            if len(v4) > 0 {
                v4[0] = 9223372036854775807
            }
        } else {
            var v83 int = 128
            fmt.Println(v83)
            _ = v83
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
