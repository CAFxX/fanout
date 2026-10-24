package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8; F1 bool; F2 uint }
type GenS9 struct { F0 uint; F1 string; F2 int8 }
func genFn0(p0 uint8, p1 int32) uint64 {
    fmt.Println(p1, p0)
    v6 := GenS1{F0: 2}
    if (false || (472 <= 65535)) {
        var v7 bool = true
        p0 -= 0
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return 508
}
func genFn1(p0 uint16, p1 uint) int16 {
    if true {
        var v8 string = "\n\t"
        if true {
            v9 := GenS2{F0: v8}
            _ = v9
        }
        _ = v8
    }
    v10 := []int{127, 32767, -2147483648}
    _ = p0
    _ = p1
    _ = v10
    return -32768
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]bool)
    var v3 int64 = 779
    var v4 uint32 = uint32(v1)
    v5 := []string{((("foo bar" + (("foo bar" + "hello") + "a")) + "foo bar") + ""), "\n\t", "a", ((("a" + "hello") + "") + (((("foo bar" + "a") + ("foo bar" + "")) + "\n\t") + ("foo bar" + "hello")))}
    fmt.Println(v4, v1)
    v3 = v3
    fmt.Println(v1, v3, v4)
    switch "\n\t" {
    case "hello":
        for fi11 := 0; fi11 < 2; fi11++ {
            if false {
                v3 = v3
                var v12 float32 = 1.5
                switch "\n\t" {
                case "a":
                    v13 := GenS2{F0: "hello"}
                    _ = v13
                case "foo bar":
                    v4 -= v4
                default:
                    fmt.Println(v1)
                }
                _ = v12
            } else {
                v4 += 604
                var v14 int32 = -1
                _ = v14
            }
            _ = fi11
        }
    }
    v15 := GenS3{F0: 306, F1: (v4 ^ v4)}
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(v4, v3)
        var v17 uint8 = 15
        var v18 int64 = -465
        _ = fi16
        _ = v17
        _ = v18
    }
    if true {
        switch 859 {
        case 764:
            v1 -= v1
        }
        var v19 float64 = 1.5
        v20 := GenS3{F0: 2, F1: 2147483647}
        _ = v19
        _ = v20
    }
    v3 += v3
    for fi21 := 0; fi21 < 11; fi21++ {
        if (2 > (v4 ^ v4)) {
            for fi22 := 0; fi22 < 12; fi22++ {
                v1 = (^(v1))
                if (376 != -2147483648) {
                    _ = len(v2)
                    v23 := GenS4{F0: 42, F1: (v4 & v4)}
                    v24 := GenS1{F0: 796}
                    _ = v23
                    _ = v24
                } else {
                    v25 := GenS1{F0: 42}
                    v4 += (^(v4))
                    var v26 float32 = -2.25
                    _ = v25
                    _ = v26
                }
                v2[4294967295] = false
                _ = fi22
            }
            for fi27 := 0; fi27 < 11; fi27++ {
                _ = len(v2)
                fmt.Println(fi27)
                fmt.Println(fi27, fi27)
                v28 := GenS2{F0: "foo bar"}
                v4 -= (v4 / (((v4 / (((v4 - uint32(v1))) | 1))) | 1))
                _ = fi27
                _ = v28
            }
            var v29 int32 = -1
            _ = v29
        } else {
            fmt.Println(fi21, v3, fi21)
            v3 = (v3 ^ (v3 >> ((255) & 63)))
        }
        _ = fi21
    }
    var v30 int16 = genFn1((^(v1)), 4294967295)
    v30 = (v30 & (v30 * (v30 / (((v30 + v30)) | 1))))
    v1 = (v1 % ((219) | 1))
    var v31 string = "hello"
    v32 := GenS1{F0: 922}
    if ((920 <= (v4 / ((v4) | 1))) && (!((!true) && false))) {
        if false {
            switch v31 {
            case "hello":
                var v33 int16 = v30
                _ = v33
            case "\n\t":
                var v34 int = 0
                _ = v34
            }
        } else {
            if true {
                fmt.Println(v4)
                fmt.Println(v31, v1)
                var v35 uint16 = v1
                _ = v35
            }
            v4 += (v4 - v4)
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if false {
            v36 := GenS1{F0: 65535}
            fmt.Println(v30)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v36
        } else {
            for fi37 := 0; fi37 < 20; fi37++ {
                if true {
                    for fi38 := 0; fi38 < 5; fi38++ {
                        delete(v2, (v4 | uint32(v1)))
                        var v39 int8 = -1
                        if false {
                            if false {
                                v39 += 1
                                var v40 uint32 = v4
                                v4 = v4
                                v2[v4] = (!false)
                                _ = v40
                            }
                            _ = len(v2)
                            v1 += 0
                        }
                        _ = fi38
                        _ = v39
                    }
                    v4 = 0
                    v5 = append(v5, v31)
                } else {
                    if false {
                        for fi41 := 0; fi41 < 2; fi41++ {
                            switch v31 {
                            case "a":
                                switch (fi41 & (fi37 ^ (fi37 ^ fi41))) {
                                case 128:
                                    v42 := GenS5{F0: (^(v4))}
                                    _ = v42
                                case 523:
                                    switch (fi41 / ((fi41) | 1)) {
                                    case 128:
                                        var v43 string = (v31 + v31)
                                        _ = v43
                                    }
                                    delete(v2, v4)
                                    if false {
                                        v44 := GenS6{F0: "a", F1: v3}
                                        switch (fi37 + 32768) {
                                        case -771:
                                            v45 := GenS7{F0: 92}
                                            _ = v45
                                        case 2147483647:
                                            var v46 int = fi37
                                            _ = v46
                                        default:
                                            v30 += (-(v30))
                                        }
                                        _ = v44
                                    } else {
                                        v47 := GenS1{F0: 301}
                                        v30 -= genFn1(v1, 0)
                                        _ = v47
                                    }
                                }
                            case "\n\t":
                                if ((434 >= 402) && false) {
                                    v48 := GenS5{F0: 680}
                                    if true {
                                        switch -638 {
                                        case 2147483647:
                                            fmt.Println(v4, fi37, v4)
                                        default:
                                            if ((v4 / ((0) | 1)) <= 42) {
                                                v31 = "hello"
                                                v1 = v1
                                                v49 := GenS5{F0: (v4 | (v4 + 633))}
                                                if true {
                                                    var v50 uint16 = 42
                                                    _ = v50
                                                }
                                                _ = v49
                                            } else {
                                                var v51 float64 = 0.1
                                                _ = v51
                                            }
                                        }
                                    }
                                    switch fi37 {
                                    case 127:
                                        fmt.Println(v3, v4)
                                    case 32767:
                                        v52 := GenS8{F0: 0, F1: ((^(v4)) >= (v4 | v4)), F2: 153}
                                        _ = v52
                                    default:
                                        v3 = v3
                                        v53 := GenS9{F0: 1, F1: v31, F2: 0}
                                        _ = v53
                                    }
                                    v1 += v1
                                    v54 := GenS6{F0: v31, F1: v3}
                                    _ = v48
                                    _ = v54
                                }
                            case "foo bar":
                                var v55 int8 = 42
                                _ = v55
                            default:
                                if true {
                                    fmt.Println(fi37)
                                    var v56 int64 = v3
                                    _ = v56
                                }
                            }
                            _ = len(v2)
                            _ = fi41
                        }
                        v30 = 1000
                        var v57 int32 = 619
                        _ = v57
                    }
                }
                var v58 float64 = (0.1 - (1.5 + -2.25))
                _ = fi37
                _ = v58
            }
        }
    } else {
        fmt.Println(v1)
        v3 = 0
    }
    var v59 bool = false
    var v60 string = "\n\t"
    fmt.Println(v31, v59)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v59
    _ = v60
}
