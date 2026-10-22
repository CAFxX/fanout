package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 bool; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int64; F2 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS7 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS8 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS9 struct { F0 uint32; F1 int32; F2 uint32 }
func genFn0(p0 int64, p1 int32) int16 {
    v4 := GenS1{F0: 493}
    v5 := GenS2{F0: "hello", F1: true, F2: ""}
    if false {
        var v6 int64 = (p0 ^ (p0 | p0))
        var v7 int8 = -128
        _ = v6
        _ = v7
    } else {
        v8 := GenS3{F0: 32767}
        var v9 int8 = 42
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 uint = 65535
    var v10 int32 = 905
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := GenS3{F0: 65535}
        v13 := GenS4{F0: v3, F1: 1, F2: 359}
        v14 := []int{32767, (fi11 % (((fi11 ^ fi11)) | 1)), (fi11 << ((255) & 63))}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 float32 = 3.14159
    var v16 bool = true
    if true {
        v16 = true
        for fi17 := 0; fi17 < 19; fi17++ {
            switch "hello" {
            case "":
                v16 = v16
            case "hello":
                var v18 int = 127
                _ = v18
            case "\n\t":
                fmt.Println(v15)
            }
            _ = fi17
        }
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        if v16 {
            _ = len(v2)
        } else {
            v15 -= v15
            fmt.Println(v16)
        }
        switch -390 {
        case 9223372036854775807:
            var v20 uint64 = 4294967295
            _ = v20
        case 1:
            fmt.Println(fi19)
        case -9223372036854775808:
            delete(v1, (fi19 - (fi19 % (((fi19 >> (((v3 + 180)) & 63))) | 1))))
        default:
            if (18446744073709551615 != 2) {
                if (false || (42 <= 471)) {
                    for fi21 := 0; fi21 < 16; fi21++ {
                        switch "foo bar" {
                        case "":
                            if v16 {
                                v10 = (v10 | (v10 * v10))
                                if ("hello" == "") {
                                    v22 := GenS1{F0: -1}
                                    _ = v22
                                }
                                fmt.Println(v10, v15)
                            } else {
                                if v16 {
                                    fmt.Println(v10, v16)
                                    if (v15 == v15) {
                                        if v16 {
                                            fmt.Println(fi19, v3, fi19)
                                            v23 := GenS4{F0: v3, F1: -851, F2: 2147483648}
                                            v10 += v10
                                            _ = v23
                                        } else {
                                            fmt.Println(v10, v16, fi21)
                                            delete(v2, "hello")
                                            v24 := GenS5{F0: (("foo bar" + (("a" + "foo bar") + "foo bar")) + ("\n\t" + "\n\t"))}
                                            _ = v24
                                        }
                                        if (16 >= 246) {
                                            var v25 float64 = (1.5 + 1e10)
                                            if v16 {
                                                v26 := GenS4{F0: 42, F1: 1, F2: 276}
                                                fmt.Println(fi19)
                                                v25 = (v25 * v25)
                                                _ = v26
                                            }
                                            var v27 uint32 = 768
                                            _ = v25
                                            _ = v27
                                        } else {
                                            switch fi19 {
                                            case 32767:
                                                fmt.Println(fi21)
                                                v28 := []int{127}
                                                var v29 uint64 = 904
                                                var v30 uint64 = v29
                                                if len(v28) > 0 {
                                                    v28[0] = fi19
                                                }
                                                v15 = (1.5 / v15)
                                                v30 = v30
                                                if (true && ((fi19 - fi19) < (fi19 | fi19))) {
                                                    v3 -= v3
                                                }
                                                fmt.Println(v10, v30)
                                                v30 = (v30 - (v29 & v30))
                                                var v31 int32 = (v10 | (v10 & v10))
                                                v32 := GenS6{F0: false, F1: 65535, F2: 65535}
                                                v16 = v16
                                                v33 := GenS7{F0: 2, F1: 65535, F2: v10}
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                            case -2147483648:
                                                var v34 int64 = 1
                                                _ = v34
                                            }
                                        }
                                    } else {
                                        v35 := GenS7{F0: 4294967295, F1: 65535, F2: v10}
                                        v3 -= 42
                                        fmt.Println(fi19, v3)
                                        _ = v35
                                    }
                                    v36 := v2[(("foo bar" + (("foo bar" + (("hello" + "\n\t") + "a")) + ("\n\t" + ""))) + (("foo bar" + ("foo bar" + ("a" + ""))) + ("\n\t" + "hello")))]
                                    _ = v36
                                } else {
                                    if ((128 < 0) || ("hello" > "hello")) {
                                        if ((((264 > 2147483648) && (44 > 1)) && (v10 > (v10 / (((v10 * v10)) | 1)))) && (!(!v16))) {
                                            switch "" {
                                            case "\n\t":
                                                fmt.Println(v16)
                                            }
                                            var v37 int = fi21
                                            _ = v37
                                        }
                                        var v38 string = ""
                                        if v16 {
                                            switch v38 {
                                            case "\n\t":
                                                var v39 string = "hello"
                                                v40 := []int{(fi19 & (-(fi19))), (fi21 + -9223372036854775808), (^(fi19))}
                                                v10 = 974
                                                _ = v39
                                                _ = v40
                                            }
                                            var v41 float64 = -2.25
                                            _ = v41
                                        } else {
                                            v38 = (v38 + v38)
                                        }
                                        _ = v38
                                    }
                                }
                                fmt.Println(fi21)
                                var v42 int64 = -138
                                _ = v42
                            }
                        }
                        if (v16 || false) {
                            v3 = (v3 << (((v3 + v3)) & 63))
                            var v43 uint = (v3 | v3)
                            if ((v10 ^ (v10 | (v10 ^ (^(v10))))) == -2147483648) {
                                v43 -= 2
                                var v44 float32 = v15
                                v2["foo bar"] = (fi21 << ((2) & 63))
                                _ = v44
                            } else {
                                var v45 bool = (false || v16)
                                delete(v1, fi19)
                                var v46 uint32 = 42
                                _ = v45
                                _ = v46
                            }
                            _ = v43
                        } else {
                            fmt.Println(v3, v10, v16)
                        }
                        _ = fi21
                    }
                    if (!(1 > 496)) {
                        fmt.Println(v15)
                    } else {
                        for fi47 := 0; fi47 < 15; fi47++ {
                            if (-9223372036854775808 == 1) {
                                v48 := []int{}
                                _ = v48
                            } else {
                                delete(v2, "")
                                var v49 uint = 18446744073709551615
                                _ = v49
                            }
                            _ = fi47
                        }
                    }
                }
                var v50 int16 = -32768
                for fi51 := 0; fi51 < 15; fi51++ {
                    switch int(v10) {
                    case -1:
                        v52 := GenS3{F0: 1}
                        _ = v52
                    }
                    v50 -= (v50 % ((1) | 1))
                    var v53 int32 = v10
                    _ = fi51
                    _ = v53
                }
                _ = v50
            } else {
                switch fi19 {
                case 42:
                    if v16 {
                        if v16 {
                            v15 -= float32(1.5)
                            for fi54 := 0; fi54 < 13; fi54++ {
                                v16 = v16
                                v55 := GenS5{F0: "\n\t"}
                                _ = fi54
                                _ = v55
                            }
                        }
                    } else {
                        fmt.Println(v3, v16, fi19)
                    }
                }
            }
        }
        _ = fi19
    }
    delete(v1, 127)
    v56 := []int{4294967295, -124, -51}
    v57 := GenS8{F0: 4294967295, F1: 4294967295, F2: 35}
    v58 := GenS9{F0: 4294967295, F1: (v10 * v10), F2: 2147483648}
    switch 128 {
    case 32768:
        fmt.Println(v3)
    default:
        var v59 float64 = float64((3.14159 * float32(1e10)))
        _ = v59
    }
    v16 = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v56:", len(v56), v56)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v56
    _ = v57
    _ = v58
}
