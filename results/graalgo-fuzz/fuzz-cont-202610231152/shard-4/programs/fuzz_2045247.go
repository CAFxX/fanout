package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int32; F1 uint; F2 uint16 }
type GenS7 struct { F0 int8 }
func genFn0() uint {
    var v5 int32 = 796
    var v6 bool = ((42 < 65535) || (!false))
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: (v5 << ((0) & 63))}
            v5 = (v5 & (v5 >> ((65535) & 63)))
            _ = fi8
            _ = v9
        }
        v5 += 42
        switch fi7 {
        case -2147483648:
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 string = ("a" + "")
                var v12 uint8 = 42
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        _ = fi7
    }
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int, p1 uint64) int16 {
    p1 -= (^(p1))
    v13 := []int{p0, (p0 * p0)}
    _ = p0
    _ = p1
    _ = v13
    return 1000
}

func main() {
    var v1 uint16 = 2
    var v2 uint16 = 65535
    v3 := make(map[uint32]bool)
    var v4 bool = (float32(0.1) > (0.0 * float32(-2.25)))
    v1 = v2
    v14 := GenS1{F0: 2147483647}
    if v4 {
        if (!true) {
            v2 -= v2
            if ("" > "\n\t") {
                var v15 uint8 = 147
                _ = v15
            } else {
                v16 := GenS2{F0: (!v4)}
                _ = v16
            }
            for fi17 := 0; fi17 < 14; fi17++ {
                v1 -= 0
                if (1e10 > ((-2.25 / 1.0) / float32(-2.25))) {
                    v2 = 2
                    for fi18 := 0; fi18 < 11; fi18++ {
                        v19 := GenS3{F0: "", F1: 42}
                        switch 4294967295 {
                        case 1:
                            v20 := GenS3{F0: ("" + "foo bar"), F1: 980}
                            _ = v20
                        case 4294967295:
                            var v21 uint8 = uint8(v1)
                            _ = v21
                        case 127:
                            v22 := []int{32767}
                            _ = v22
                        }
                        v23 := GenS4{F0: fi18, F1: "a"}
                        if true {
                            var v24 uint32 = 1
                            v25 := []int{fi18}
                            v1 -= 235
                            if v4 {
                                var v26 int32 = 2147483647
                                _ = v26
                            } else {
                                if true {
                                    v2 -= v2
                                    var v27 int16 = -628
                                    _ = v27
                                } else {
                                    if len(v25) > 0 {
                                        v25[0] = (fi17 >> ((42) & 63))
                                    }
                                }
                            }
                            var v28 float64 = ((0.1 + -2.25) / 1.5)
                            _ = v24
                            _ = v25
                            _ = v28
                        } else {
                            v29 := GenS4{F0: fi17, F1: ""}
                            v30 := GenS3{F0: "", F1: 2147483647}
                            switch "a" {
                            case "a":
                                v4 = v4
                            }
                            _ = v29
                            _ = v30
                        }
                        _ = fi18
                        _ = v19
                        _ = v23
                    }
                } else {
                    var v31 uint32 = 443
                    _ = v31
                }
                switch (("" + "a") + "hello") {
                case "foo bar":
                    v2 += 356
                case "\n\t":
                    for fi32 := 0; fi32 < 19; fi32++ {
                        v33 := GenS2{F0: (v4 || (0.0 >= float32((0.0 * 0.1))))}
                        v2 -= 65535
                        _ = fi32
                        _ = v33
                    }
                default:
                    v34 := GenS5{F0: 127, F1: v2}
                    _ = v34
                }
                _ = fi17
            }
        }
        v35 := GenS6{F0: 0, F1: 2, F2: (v1 & (v2 & (v1 - 0)))}
        _ = v35
    } else {
        for fi36 := 0; fi36 < 8; fi36++ {
            for fi37 := 0; fi37 < 11; fi37++ {
                _ = len(v3)
                v4 = (v4 || v4)
                v1 += 255
                var v38 bool = (32767 < v2)
                _ = fi37
                _ = v38
            }
            _ = fi36
        }
        var v39 string = "a"
        if (!(-1 >= -2147483648)) {
            switch v39 {
            case "a":
                switch 32767 {
                case 89:
                    fmt.Println(v1)
                case 32767:
                    v40 := []int{32768, 127}
                    _ = v40
                }
            case "":
                if (!(true && v4)) {
                    fmt.Println(v2, v2, v4)
                } else {
                    v2 += (v1 ^ v2)
                }
            }
            switch 2147483647 {
            case -990:
                for fi41 := 0; fi41 < 7; fi41++ {
                    v42 := []int{0, 0}
                    v1 = v2
                    v2 = 139
                    _ = fi41
                    _ = v42
                }
            default:
                if true {
                    fmt.Println(v4)
                } else {
                    if v4 {
                        var v43 int32 = -1
                        fmt.Println(v39)
                        _ = v43
                    } else {
                        switch 128 {
                        case 59:
                            var v44 int = 32768
                            _ = v44
                        case 4294967295:
                            switch 2147483647 {
                            case -700:
                                var v45 uint64 = 18446744073709551615
                                _ = v45
                            case 128:
                                var v46 int64 = -9223372036854775808
                                _ = v46
                            default:
                                v47 := GenS1{F0: 638}
                                _ = v47
                            }
                        }
                        fmt.Println(v1)
                    }
                    var v48 int16 = 32767
                    _ = v48
                }
            }
        } else {
            v39 = ((v39 + v39) + (v39 + v39))
            if v4 {
                fmt.Println(v1, v4)
            } else {
                if ((-603 < 0) && (1000 > 42)) {
                    for fi49 := 0; fi49 < 3; fi49++ {
                        var v50 string = ("a" + (v39 + "\n\t"))
                        for fi51 := 0; fi51 < 5; fi51++ {
                            v52 := []int{(fi51 << ((genFn0()) & 63)), -827}
                            if (!v4) {
                                v39 = ("a" + v50)
                                var v53 int8 = -128
                                v54 := GenS1{F0: 42}
                                _ = v53
                                _ = v54
                            }
                            if (!true) {
                                v39 = "a"
                            }
                            v2 += (v2 - (v2 | 259))
                            if (1.5 == (((3.14159 + 0.1) + 0.0) * (3.14159 / 1.5))) {
                                var v55 float32 = ((((1.5 + 0.1) - float32((float64((0.0 - 0.1)) * 0.1))) - 0.1) / 1e10)
                                _ = v55
                            } else {
                                if false {
                                    if len(v52) > 0 {
                                        v52[0] = (fi49 >> ((0) & 63))
                                    }
                                    v4 = (485 > 2)
                                    v39 = "a"
                                    _ = len(v52)
                                    var v56 int32 = -235
                                    _ = v56
                                } else {
                                    var v57 bool = (0 > 1)
                                    _ = v57
                                }
                                var v58 int8 = int8(fi51)
                                fmt.Println(v1, v39)
                                _ = v58
                            }
                            _ = fi51
                            _ = v52
                        }
                        v59 := GenS7{F0: 0}
                        _ = fi49
                        _ = v50
                        _ = v59
                    }
                    fmt.Println(v39)
                    var v60 int32 = -2147483648
                    _ = v60
                } else {
                    if v4 {
                        if v4 {
                            var v61 uint = genFn0()
                            _ = v61
                        } else {
                            v1 = v2
                        }
                    } else {
                        var v62 uint64 = 0
                        _ = v62
                    }
                    v1 = v1
                    v2 = 0
                }
                if ((!false) || true) {
                    switch 1 {
                    case 4294967295:
                        v3[646] = (364 >= 0)
                    default:
                        v63 := GenS2{F0: true}
                        _ = v63
                    }
                } else {
                    for fi64 := 0; fi64 < 1; fi64++ {
                        var v65 uint64 = 42
                        v66 := []int{fi64, fi64, fi64, 0}
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                    v67 := v3[831]
                    v68 := GenS2{F0: true}
                    _ = v67
                    _ = v68
                }
            }
            v69 := GenS7{F0: 0}
            _ = v69
        }
        _ = v39
    }
    var v70 uint64 = 0
    v71 := []int{-940}
    if v4 {
        v72 := GenS1{F0: -881}
        if len(v71) > 0 {
            v71[0] = 843
        }
        v73 := GenS2{F0: false}
        _ = v72
        _ = v73
    } else {
        var v74 int32 = -727
        _ = v74
    }
    v75 := GenS6{F0: -241, F1: uint(v70), F2: (v1 | v2)}
    for fi76 := 0; fi76 < 14; fi76++ {
        var v77 uint32 = 205
        switch "foo bar" {
        case "\n\t":
            var v78 int = fi76
            _ = v78
        case "foo bar":
            switch (fi76 | (fi76 * (fi76 * (fi76 ^ 1)))) {
            case -9223372036854775808:
                var v79 float32 = -2.25
                _ = v79
            case 1:
                var v80 uint64 = v70
                _ = v80
            }
        case "a":
            for fi81 := 0; fi81 < 14; fi81++ {
                v1 = v1
                _ = fi81
            }
        default:
            var v82 uint16 = uint16(v77)
            _ = v82
        }
        _ = fi76
        _ = v77
    }
    v2 = (v1 >> ((4294967295) & 63))
    var v83 int32 = 1
    v71 = append(v71, 128)
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v70:", v70)
    fmt.Println("slicelen v71:", len(v71), v71)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v70
    _ = v71
    _ = v75
    _ = v83
}
