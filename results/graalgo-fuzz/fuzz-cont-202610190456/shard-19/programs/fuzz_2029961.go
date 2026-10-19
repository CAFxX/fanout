package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int }
func genFn0(p0 bool, p1 int, p2 uint32) uint64 {
    if p0 {
        v7 := []int{(p1 - 32767), p1, p1, p1}
        if false {
            var v8 uint = 0
            fmt.Println(v8)
            var v9 bool = false
            _ = v8
            _ = v9
        }
        if p0 {
            v10 := GenS1{F0: ("" + "a")}
            _ = v10
        }
        _ = v7
    } else {
        v11 := []int{(p1 << ((255) & 63)), 42, p1}
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := []int{fi12, (p1 * (-(p1)))}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 32767
    var v3 float64 = 0.1
    v4 := make(map[uint32]int)
    var v5 uint32 = 2
    var v6 string = ""
    v3 = v3
    var v14 float64 = float64(0.1)
    v6 = v6
    for fi15 := 0; fi15 < 1; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            if false {
                if (true && false) {
                    v2 -= (v2 | (v2 * (-(v2))))
                    var v17 int32 = -2147483648
                    v17 -= v17
                    _ = v17
                }
                if false {
                    v2 -= (-(v2))
                }
            } else {
                v18 := []int{(^(fi16)), (fi16 / (((fi16 % (((-(fi16))) | 1))) | 1)), fi16}
                fmt.Println(fi15, v6, v3)
                fmt.Println(v2, v3, v5)
                v6 = v6
                v14 = float64(0.1)
                _ = v18
            }
            _ = fi16
        }
        if false {
            var v19 int64 = 9223372036854775807
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 uint = uint(v5)
                switch "" {
                case "\n\t":
                    v6 = v6
                case "":
                    fmt.Println(v14, v14, v19)
                    v3 = v3
                case "a":
                    switch v6 {
                    case "\n\t":
                        v21 = (^(v21))
                    case "":
                        v6 = "hello"
                    default:
                        v22 := GenS1{F0: v6}
                        _ = v22
                    }
                }
                _ = fi20
                _ = v21
            }
            var v23 int64 = (v19 % ((v19) | 1))
            _ = v19
            _ = v23
        } else {
            if (true || (0 != 255)) {
                delete(v1, (v5 & 2147483647))
            } else {
                if ((true || true) && false) {
                    var v24 string = ((v6 + "hello") + "foo bar")
                    _ = v24
                }
                if (65535 >= 4294967295) {
                    if false {
                        switch -9223372036854775808 {
                        case 1:
                            v25 := []int{}
                            _ = v25
                        case -2147483648:
                            var v26 float32 = 0.1
                            _ = v26
                        case -930:
                            var v27 int32 = 2147483647
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 2; fi28++ {
                            var v29 int32 = 2147483647
                            switch v6 {
                            case "":
                                switch ((v6 + "hello") + v6) {
                                case "a":
                                    delete(v4, v5)
                                    v6 = (v6 + v6)
                                default:
                                    var v30 string = v6
                                    _ = v30
                                }
                            case "a":
                                var v31 float32 = ((3.14159 + 0.0) / 1.5)
                                _ = v31
                            default:
                                v14 = v3
                            }
                            _ = fi28
                            _ = v29
                        }
                        var v32 int = (-(fi15))
                        _ = v32
                    }
                } else {
                    v6 = v6
                }
            }
        }
        _ = fi15
    }
    if false {
        var v33 bool = (v6 == (v6 + v6))
        if false {
            v34 := v1[(v5 % ((2147483647) | 1))]
            v35 := GenS1{F0: v6}
            _ = v34
            _ = v35
        }
        _ = v33
    } else {
        v2 += 0
        v2 = v2
        v36 := GenS2{F0: 9223372036854775807, F1: 9223372036854775807, F2: v5}
        _ = v36
    }
    if (0.1 > 0.0) {
        if true {
            var v37 int64 = 9223372036854775807
            var v38 uint64 = 4294967295
            _ = v37
            _ = v38
        } else {
            for fi39 := 0; fi39 < 19; fi39++ {
                if false {
                    if ((fi39 + (fi39 >> ((70) & 63))) != 1) {
                        for fi40 := 0; fi40 < 1; fi40++ {
                            v3 = (v3 * v14)
                            _ = fi40
                        }
                        var v41 float64 = v3
                        _ = v41
                    }
                } else {
                    var v42 uint16 = 0
                    _ = v42
                }
                _ = fi39
            }
            v5 += (v5 / (((v5 & 528)) | 1))
        }
        if false {
            v43 := GenS2{F0: -53, F1: 0, F2: (v5 / ((v5) | 1))}
            if true {
                var v44 int = 418
                _ = v44
            } else {
                v2 += 42
                fmt.Println(v3, v5, v14)
            }
            if false {
                var v45 float64 = -2.25
                v4[(^(v5))] = -607
                _ = v45
            }
            _ = v43
        } else {
            v3 += ((0.1 * v14) * v14)
            if ((v3 < v14) && (2 == v5)) {
                var v46 bool = (false || false)
                _ = v46
            }
        }
    } else {
        var v47 float64 = v3
        var v48 int32 = 42
        switch "" {
        case "\n\t":
            if true {
                if true {
                    v49 := GenS3{F0: 216, F1: 65535}
                    fmt.Println(v2, v2, v6)
                    switch 1 {
                    case 511:
                        for fi50 := 0; fi50 < 8; fi50++ {
                            v51 := GenS1{F0: v6}
                            v52 := GenS3{F0: 461, F1: 538}
                            _ = fi50
                            _ = v51
                            _ = v52
                        }
                    case 0:
                        var v53 uint16 = 32767
                        _ = v53
                    case 127:
                        for fi54 := 0; fi54 < 7; fi54++ {
                            var v55 uint = 18446744073709551615
                            _ = fi54
                            _ = v55
                        }
                    }
                    _ = v49
                }
            } else {
                v2 = v2
            }
        case "a":
            if false {
                for fi56 := 0; fi56 < 9; fi56++ {
                    switch fi56 {
                    case 128:
                        var v57 int16 = (^(v2))
                        _ = v57
                    case 221:
                        if true {
                            v3 -= float64(float32(v47))
                        } else {
                            v58 := GenS4{F0: (fi56 & (fi56 + (fi56 | fi56)))}
                            fmt.Println(fi56, v2)
                            var v59 float64 = v14
                            _ = v58
                            _ = v59
                        }
                    }
                    v48 += v48
                    var v60 uint = 42
                    _ = fi56
                    _ = v60
                }
            }
        case "":
            v14 += v3
        }
        _ = v47
        _ = v48
    }
    v61 := GenS4{F0: 0}
    v6 = (("" + (v6 + (v6 + (v6 + ("\n\t" + ("a" + "\n\t")))))) + (v6 + "a"))
    fmt.Println(v2, v14, v5)
    if ((v2 | (-(v2))) < v2) {
        switch ("foo bar" + v6) {
        case "":
            v62 := []int{32768}
            _ = v62
        }
    } else {
        var v63 bool = true
        var v64 float32 = 1.5
        v3 += float64(v64)
        _ = v63
        _ = v64
    }
    var v65 int8 = 100
    var v66 int = 32768
    v65 = (v65 | -62)
    v2 -= v2
    if true {
        var v67 string = v6
        _ = v67
    } else {
        var v68 float64 = (v3 + 3.14159)
        v69 := GenS4{F0: (v66 ^ -1)}
        _ = v68
        _ = v69
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v65
    _ = v66
}
