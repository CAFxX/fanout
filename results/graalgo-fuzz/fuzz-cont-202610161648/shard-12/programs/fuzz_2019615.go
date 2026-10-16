package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint8; F1 int; F2 uint16 }
type GenS4 struct { F0 string; F1 uint8 }
func genFn0(p0 int64, p1 int16) uint8 {
    var v6 uint16 = 692
    v7 := GenS1{F0: 0}
    var v8 int = 128
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int32, p1 bool) int {
    var v9 uint64 = 891
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 string = ("foo bar" + "")
        if false {
            if p1 {
                switch v11 {
                case "a":
                    v12 := []int{(fi10 >> ((18446744073709551615) & 63)), -9223372036854775808, fi10}
                    _ = v12
                case "hello":
                    v13 := []int{fi10, fi10, 217, (fi10 ^ fi10)}
                    _ = v13
                default:
                    v14 := []int{fi10, (fi10 & fi10)}
                    _ = v14
                }
            } else {
                v15 := GenS1{F0: 0}
                p1 = p1
                switch (fi10 >> ((4294967295) & 63)) {
                case 32767:
                    for fi16 := 0; fi16 < 17; fi16++ {
                        if (!((p1 && (!false)) && p1)) {
                            v17 := []int{(fi16 | (-(fi10))), 42, fi16, 4294967295}
                            if true {
                                v17 = append(v17, (fi10 + (fi16 % ((2147483647) | 1))))
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                            } else {
                                v9 = (v9 ^ v9)
                                var v18 uint8 = 0
                                var v19 int64 = 0
                                _ = v18
                                _ = v19
                            }
                            var v20 int32 = -1
                            _ = v17
                            _ = v20
                        }
                        _ = fi16
                    }
                case 4294967295:
                    v21 := []int{}
                    _ = v21
                case 128:
                    if p1 {
                        if p1 {
                            if (-1 < -325) {
                                v22 := []int{}
                                v23 := GenS2{F0: 255, F1: int(v9)}
                                _ = v22
                                _ = v23
                            } else {
                                v9 -= 4294967295
                                var v24 float32 = ((0.0 / -2.25) - ((1.5 / 1e10) + -2.25))
                                _ = v24
                            }
                            v25 := []int{(fi10 % (((-(fi10))) | 1))}
                            var v26 uint16 = 65535
                            _ = v25
                            _ = v26
                        } else {
                            v27 := []int{954, (-(fi10))}
                            for fi28 := 0; fi28 < 20; fi28++ {
                                var v29 uint16 = 1
                                _ = fi28
                                _ = v29
                            }
                            _ = v27
                        }
                        if true {
                            var v30 int = fi10
                            var v31 float64 = (1e10 / -2.25)
                            _ = v30
                            _ = v31
                        }
                    }
                default:
                    p0 = -1
                }
                _ = v15
            }
        } else {
            v32 := []int{fi10, fi10, fi10}
            v11 = (((v11 + v11) + ("a" + (v11 + v11))) + ("\n\t" + ""))
            _ = v32
        }
        if false {
            v9 = (^(v9))
            p1 = p1
        } else {
            fmt.Println(p0, v11)
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return 502
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 uint8 = 1
    v5 := []int{32767, 42, 32768, 2147483647}
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 int16 = -425
        switch genFn1(1, false) {
        case 2147483647:
            if false {
                var v35 int8 = -128
                _ = v35
            }
        case 32768:
            var v36 uint8 = genFn0(0, 856)
            _ = v36
        }
        _ = fi33
        _ = v34
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        if (genFn1(1, ((537 < fi37) && true)) < 2147483647) {
            if (42 <= -32768) {
                if (!(1.5 != -2.25)) {
                    var v38 float32 = 3.14159
                    _ = v38
                } else {
                    var v39 int = fi37
                    if true {
                        v39 -= (v39 & fi37)
                    } else {
                        var v40 uint8 = (v4 << ((42) & 63))
                        fmt.Println(v4, v40)
                        fmt.Println(fi37, v4, fi37)
                        _ = v40
                    }
                    v4 += 127
                    _ = v39
                }
                var v41 uint = 255
                _ = v41
            }
            if false {
                if (!true) {
                    v42 := GenS2{F0: 65535, F1: (fi37 / (((fi37 / ((fi37) | 1))) | 1))}
                    if false {
                        if ((-128 > 100) && (!((false && true) || (721 != 42)))) {
                            var v43 int16 = 0
                            var v44 int32 = -1
                            _ = v43
                            _ = v44
                        }
                    } else {
                        v45 := GenS2{F0: 4294967295, F1: 4294967295}
                        switch fi37 {
                        case 127:
                            if len(v5) > 0 {
                                v5[0] = fi37
                            }
                        }
                        _ = v45
                    }
                    var v46 int64 = 1
                    _ = v42
                    _ = v46
                } else {
                    if true {
                        switch (fi37 / ((32767) | 1)) {
                        case 128:
                            v4 = v4
                        default:
                            fmt.Println(v4)
                        }
                    }
                    if (!(fi37 != fi37)) {
                        switch (fi37 + fi37) {
                        case 4294967295:
                            var v47 uint16 = 42
                            _ = v47
                        case 0:
                            v48 := GenS2{F0: 400, F1: (-(fi37))}
                            _ = v48
                        case 9223372036854775807:
                            v49 := GenS2{F0: 42, F1: fi37}
                            _ = v49
                        }
                        v50 := GenS1{F0: (v4 << ((748) & 63))}
                        _ = v50
                    } else {
                        delete(v1, "foo bar")
                        var v51 uint64 = 0
                        var v52 bool = false
                        _ = v51
                        _ = v52
                    }
                }
                var v53 string = ("" + "\n\t")
                _ = v53
            }
        } else {
            if (!(-381 == -18)) {
                switch (("\n\t" + "\n\t") + "") {
                case "\n\t":
                    for fi54 := 0; fi54 < 16; fi54++ {
                        v4 = v4
                        fmt.Println(v4, fi54, v4)
                        _ = fi54
                    }
                case "hello":
                    var v55 uint64 = 707
                    _ = v55
                }
            }
            v56 := GenS3{F0: v4, F1: (^(fi37)), F2: 2}
            fmt.Println(fi37, fi37)
            _ = v56
        }
        if len(v5) > 0 {
            v5[0] = fi37
        }
        _ = fi37
    }
    var v57 uint64 = 1
    fmt.Println(v57, v57)
    var v58 float64 = ((1.5 - -2.25) + -2.25)
    v57 = (v57 << ((uint(v57)) & 63))
    for fi59 := 0; fi59 < 11; fi59++ {
        fmt.Println(v58)
        if len(v5) > 0 {
            v5[0] = -826
        }
        _ = fi59
    }
    var v60 float64 = v58
    v57 -= (v57 & v57)
    v60 = 1.5
    if false {
        fmt.Println(v57, v4)
    }
    v58 = v60
    v61 := GenS3{F0: genFn0(-1, 1000), F1: genFn1(-443, false), F2: 65535}
    v4 = genFn0(9223372036854775807, -1)
    v1[""] = 0
    for fi62 := 0; fi62 < 19; fi62++ {
        v63 := GenS4{F0: "a", F1: v4}
        _ = fi62
        _ = v63
    }
    if true {
        v57 += (v57 * v57)
        fmt.Println(v4, v58)
    } else {
        switch 42 {
        case 618:
            var v64 uint16 = 255
            _ = v64
        case 184:
            for fi65 := 0; fi65 < 16; fi65++ {
                if true {
                    v66 := GenS1{F0: v4}
                    var v67 uint = 18446744073709551615
                    _ = v66
                    _ = v67
                }
                var v68 uint16 = 1
                _ = fi65
                _ = v68
            }
        case 4294967295:
            v60 -= v60
        }
        for fi69 := 0; fi69 < 17; fi69++ {
            for fi70 := 0; fi70 < 17; fi70++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi70
            }
            fmt.Println(v58)
            var v71 float64 = ((float64(3.14159) * v58) / -2.25)
            _ = fi69
            _ = v71
        }
        v60 = (((v58 - v60) + (v60 * 0.1)) * v60)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v60
    _ = v61
}
