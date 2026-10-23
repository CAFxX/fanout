package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint64; F1 uint }

func main() {
    v1 := []int{127, 454, 106, 2147483647}
    var v2 string = ("" + "hello")
    v3 := []uint32{169, 456}
    fmt.Println(v2)
    for fi4 := 0; fi4 < 10; fi4++ {
        v5 := GenS1{F0: 0, F1: -128, F2: -9223372036854775808}
        _ = fi4
        _ = v5
    }
    var v6 uint16 = 781
    var v7 uint16 = v6
    if (true && true) {
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 3; fi9++ {
                if true {
                    v6 += (v7 * (v6 | 2))
                    v2 = v2
                    v10 := GenS2{F0: -32768}
                    if len(v1) > 0 {
                        v1[0] = fi8
                    }
                    var v11 bool = (635 < 0)
                    _ = v10
                    _ = v11
                } else {
                    var v12 int16 = 32767
                    var v13 bool = ("a" != (((("" + v2) + (("hello" + (v2 + (v2 + (((v2 + (v2 + (v2 + v2))) + ("foo bar" + (v2 + ""))) + v2)))) + "foo bar")) + "hello") + ""))
                    _ = v12
                    _ = v13
                }
                _ = fi9
            }
            var v14 uint8 = 1
            _ = fi8
            _ = v14
        }
        if ((!(!true)) || (1 >= 42)) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch v2 {
            case "hello":
                v6 -= v6
            }
            if (1000 >= 793) {
                var v15 int16 = 942
                v7 = (v6 | v7)
                _ = v15
            }
        }
        v1 = append(v1, 127)
    } else {
        var v16 uint8 = 1
        _ = v16
    }
    v17 := GenS3{F0: 128, F1: 42, F2: 42}
    var v18 uint8 = 1
    var v19 int32 = 2147483647
    if true {
        switch 219 {
        case 1:
            var v20 uint16 = v7
            _ = v20
        case 4294967295:
            for fi21 := 0; fi21 < 2; fi21++ {
                v22 := GenS3{F0: (^(v18)), F1: 4294967295, F2: 0}
                for fi23 := 0; fi23 < 20; fi23++ {
                    if ((v19 | 2147483647) < v19) {
                        v24 := GenS4{F0: (v18 + uint8(v6))}
                        v3 = append(v3, 65535)
                        var v25 uint64 = 796
                        v19 -= (v19 + v19)
                        _ = v24
                        _ = v25
                    } else {
                        switch fi23 {
                        case 127:
                            v18 += 42
                        case 32767:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                    v26 := GenS2{F0: 42}
                    _ = fi23
                    _ = v26
                }
                _ = fi21
                _ = v22
            }
        }
    } else {
        v27 := GenS3{F0: (v18 / ((v18) | 1)), F1: 2, F2: 2}
        for fi28 := 0; fi28 < 4; fi28++ {
            if true {
                switch v2 {
                case "foo bar":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                if (!(!(((-2.25 / 0.1) / float64(float32(-2.25))) >= (((1.5 + float64((((1e10 + 0.0) / 1.5) + 1e10))) + 3.14159) + -2.25)))) {
                    v19 += (v19 / (((v19 - v19)) | 1))
                    switch -924 {
                    case -88:
                        v19 -= -1
                    case 1:
                        v19 += (v19 | -948)
                    }
                    for fi29 := 0; fi29 < 16; fi29++ {
                        if true {
                            v2 = v2
                        } else {
                            var v30 uint16 = uint16(v18)
                            fmt.Println(v2, fi29)
                            v19 = v19
                            _ = v30
                        }
                        _ = fi29
                    }
                } else {
                    v18 -= v18
                }
            } else {
                var v31 bool = true
                _ = v31
            }
            fmt.Println(fi28)
            _ = fi28
        }
        _ = v27
    }
    fmt.Println(v18, v7, v19)
    for fi32 := 0; fi32 < 17; fi32++ {
        if false {
            _ = len(v3)
        } else {
            if (42 != 0) {
                v18 = v18
                v7 -= (v7 - v7)
                var v33 float32 = (float32(1e10) + 1e10)
                _ = v33
            } else {
                var v34 int16 = -1
                _ = v34
            }
            if true {
                if (-1 <= 1) {
                    v6 = v7
                } else {
                    if true {
                        for fi35 := 0; fi35 < 7; fi35++ {
                            var v36 bool = true
                            fmt.Println(fi35, fi35)
                            var v37 bool = (true || v36)
                            var v38 uint16 = (v6 + v6)
                            if false {
                                var v39 string = v2
                                v40 := GenS5{F0: 127, F1: -676}
                                if (1.5 < 0.1) {
                                    v41 := GenS2{F0: 1000}
                                    _ = v41
                                }
                                if (!v37) {
                                    v36 = (1.5 > (float64(-2.25) - float64(1e10)))
                                    var v42 float32 = 3.14159
                                    v43 := GenS1{F0: uint(v6), F1: 0, F2: fi32}
                                    _ = v42
                                    _ = v43
                                }
                                _ = v39
                                _ = v40
                            } else {
                                v19 -= 0
                            }
                            if v36 {
                                var v44 int64 = 9223372036854775807
                                _ = v44
                            }
                            _ = fi35
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                    } else {
                        var v45 int32 = (^(v19))
                        var v46 uint64 = 1
                        v47 := GenS6{F0: v2}
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    v18 += v18
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                var v48 uint16 = (v7 >> ((42) & 63))
                v6 = 255
                _ = v48
            } else {
                var v49 int32 = v19
                _ = v49
            }
        }
        if (true || true) {
            switch v2 {
            case "hello":
                switch fi32 {
                case 747:
                    if true {
                        v7 += v6
                    } else {
                        _ = len(v3)
                        v18 -= v18
                    }
                case 1:
                    if (false || (!(float64(1e10) == (1e10 / (1.5 / 1.0))))) {
                        v2 = v2
                        var v50 int = (fi32 / (((fi32 * (fi32 % ((fi32) | 1)))) | 1))
                        switch v2 {
                        case "hello":
                            var v51 float64 = 0.1
                            _ = v51
                        }
                        _ = v50
                    } else {
                        var v52 int = fi32
                        v53 := GenS7{F0: v2, F1: v2}
                        _ = v52
                        _ = v53
                    }
                case 2147483647:
                    v6 -= (v6 >> ((0) & 63))
                }
            case "\n\t":
                for fi54 := 0; fi54 < 3; fi54++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi54
                }
            case "":
                var v55 uint64 = 0
                _ = v55
            }
            if true {
                v3 = append(v3, 4294967295)
                var v56 int32 = (^(v19))
                if (v18 != v18) {
                    _ = len(v3)
                    v57 := GenS8{F0: "hello"}
                    v58 := GenS9{F0: (v2 + "foo bar")}
                    _ = v57
                    _ = v58
                } else {
                    if false {
                        v1 = append(v1, fi32)
                    }
                    v59 := GenS10{F0: 1}
                    _ = v59
                }
                _ = v56
            } else {
                if (1.5 < 0.0) {
                    v60 := GenS11{F0: 193, F1: 170}
                    var v61 string = v2
                    _ = v60
                    _ = v61
                }
                v6 += v6
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            var v62 float64 = 0.1
            _ = v62
        }
        _ = fi32
    }
    var v63 int32 = 42
    _ = len(v1)
    v19 -= (v19 >> ((0) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v63
    _ = v7
}
