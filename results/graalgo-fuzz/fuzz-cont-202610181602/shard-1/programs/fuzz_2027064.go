package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint16; F1 bool; F2 int }

func main() {
    var v1 uint32 = 322
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]bool)
    var v4 int32 = 0
    var v5 int16 = 0
    var v6 string = "\n\t"
    v4 = (v4 % (((v4 - 1)) | 1))
    if false {
        var v7 int64 = 421
        _ = v7
    } else {
        if (-1 <= 4294967295) {
            for fi8 := 0; fi8 < 14; fi8++ {
                for fi9 := 0; fi9 < 1; fi9++ {
                    v10 := GenS1{F0: int64(v5), F1: true, F2: ("foo bar" + v6)}
                    v5 = v5
                    var v11 uint = 0
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                for fi12 := 0; fi12 < 17; fi12++ {
                    var v13 int8 = 0
                    v2 -= 305
                    _ = fi12
                    _ = v13
                }
                _ = fi8
            }
        }
        if false {
            v14 := GenS1{F0: 0, F1: (!((45 < uint16(v1)) || true)), F2: v6}
            var v15 float32 = 3.14159
            v6 = v6
            _ = v14
            _ = v15
        }
        var v16 int16 = int16(v4)
        _ = v16
    }
    if true {
        v2 = v2
        for fi17 := 0; fi17 < 12; fi17++ {
            if false {
                if true {
                    if ((((fi17 ^ 127) > fi17) && (!false)) || ((v6 != "\n\t") && false)) {
                        v18 := []int{-613, (fi17 + (fi17 / ((fi17) | 1))), (fi17 * (fi17 ^ fi17))}
                        v5 = -32768
                        v19 := GenS2{F0: 154}
                        _ = v18
                        _ = v19
                    }
                } else {
                    if (!(112 == 127)) {
                        var v20 uint32 = v1
                        for fi21 := 0; fi21 < 2; fi21++ {
                            if false {
                                v22 := v3[fi17]
                                v23 := []int{}
                                _ = v22
                                _ = v23
                            }
                            _ = fi21
                        }
                        v1 = v20
                        _ = v20
                    } else {
                        var v24 string = (v6 + "\n\t")
                        for fi25 := 0; fi25 < 6; fi25++ {
                            if true {
                                fmt.Println(fi25, fi25)
                                if (true && true) {
                                    var v26 float64 = 3.14159
                                    var v27 bool = (v26 == v26)
                                    v24 = v24
                                    _ = v26
                                    _ = v27
                                }
                            }
                            fmt.Println(fi25, v2)
                            if true {
                                if true {
                                    v28 := GenS3{F0: v2, F1: -1, F2: v2}
                                    _ = v28
                                } else {
                                    var v29 int = (fi17 * (fi25 * fi25))
                                    switch fi25 {
                                    case -2147483648:
                                        switch "hello" {
                                        case "":
                                            fmt.Println(v6, v5, v4)
                                        case "a":
                                            if false {
                                                if (!(false || false)) {
                                                    delete(v3, (fi25 / ((0) | 1)))
                                                } else {
                                                    if (v1 == (v1 - v1)) {
                                                        switch v24 {
                                                        case "\n\t":
                                                            v2 = 42
                                                        case "hello":
                                                            var v30 int64 = 9223372036854775807
                                                            _ = v30
                                                        }
                                                        v1 = 42
                                                    } else {
                                                        delete(v3, fi17)
                                                        switch 127 {
                                                        case 127:
                                                            v31 := GenS4{F0: 42, F1: false, F2: (fi17 / ((int(v5)) | 1))}
                                                            _ = v31
                                                        case 0:
                                                            if (v5 >= v5) {
                                                                var v32 uint = 755
                                                                v6 = v24
                                                                v33 := GenS2{F0: v32}
                                                                _ = v32
                                                                _ = v33
                                                            }
                                                        }
                                                        var v34 bool = true
                                                        _ = v34
                                                    }
                                                    v35 := GenS2{F0: 18446744073709551615}
                                                    _ = v35
                                                }
                                            } else {
                                                v6 = v24
                                            }
                                        default:
                                            switch (fi17 % ((128) | 1)) {
                                            case 32768:
                                                if (-128 < -28) {
                                                    v5 -= -32768
                                                    if ((!false) && ((-(v5)) < (v5 + (v5 % ((v5) | 1))))) {
                                                        if true {
                                                            v29 -= v29
                                                        } else {
                                                            v4 = v4
                                                            v36 := []int{}
                                                            _ = v36
                                                        }
                                                        v37 := GenS4{F0: 557, F1: (2 == (v2 & 1)), F2: (fi17 << ((0) & 63))}
                                                        v38 := []int{fi17, fi17, 782, (fi17 << ((1) & 63))}
                                                        _ = v37
                                                        _ = v38
                                                    } else {
                                                        var v39 uint8 = 2
                                                        _ = v39
                                                    }
                                                } else {
                                                    v40 := GenS1{F0: 9223372036854775807, F1: true, F2: "foo bar"}
                                                    _ = v40
                                                }
                                            case 344:
                                                delete(v3, v29)
                                            case 42:
                                                if true {
                                                    var v41 int = fi17
                                                    var v42 bool = (true || (1.5 >= 1.5))
                                                    _ = len(v3)
                                                    _ = v41
                                                    _ = v42
                                                }
                                            default:
                                                if true {
                                                    v43 := GenS4{F0: 0, F1: ((-327 < 1) || false), F2: (v29 >> ((255) & 63))}
                                                    var v44 int16 = (v5 % ((v5) | 1))
                                                    _ = v43
                                                    _ = v44
                                                }
                                                v45 := []int{-9223372036854775808}
                                                _ = v45
                                            }
                                        }
                                    case -9223372036854775808:
                                        var v46 float32 = -2.25
                                        _ = v46
                                    }
                                    v1 = 2147483648
                                    _ = v29
                                }
                                if (false && (199 > 480)) {
                                    v6 = ""
                                    var v47 int32 = int32(fi25)
                                    v4 += (v47 << ((2) & 63))
                                    v6 = "hello"
                                    _ = v47
                                }
                                var v48 bool = (0 == (v5 ^ v5))
                                _ = v48
                            } else {
                                v24 = v6
                                switch "" {
                                case "a":
                                    v1 += v1
                                case "hello":
                                    v2 += (v2 + 42)
                                default:
                                    _ = len(v3)
                                }
                                v2 += v2
                            }
                            v4 += (v4 & -1)
                            v4 += v4
                            v1 = (v1 + 4294967295)
                            _ = fi25
                        }
                        _ = v24
                    }
                }
                var v49 uint16 = 255
                _ = v49
            } else {
                var v50 string = v6
                if true {
                    if (!true) {
                        v1 += 42
                        v4 = v4
                        fmt.Println(v2, v2)
                    }
                }
                var v51 uint64 = v2
                _ = v50
                _ = v51
            }
            _ = fi17
        }
    }
    fmt.Println(v4, v2, v4)
    if (672 >= 127) {
        if true {
            v52 := []int{987}
            switch v6 {
            case "":
                var v53 uint64 = v2
                _ = v53
            default:
                var v54 uint16 = 2
                _ = v54
            }
            _ = v52
        } else {
            var v55 uint64 = (v2 % ((545) | 1))
            v56 := []int{}
            _ = v55
            _ = v56
        }
        var v57 uint8 = 63
        _ = v57
    } else {
        var v58 uint = 729
        _ = v58
    }
    if true {
        v6 = ("\n\t" + v6)
    }
    var v59 int16 = (v5 % (((v5 * v5)) | 1))
    var v60 int8 = -40
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
}
