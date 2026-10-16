package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int16 }
type GenS2 struct { F0 string; F1 string; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint8; F1 string; F2 bool }
func genFn0(p0 int8, p1 int8) int {
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            p0 = p0
            v8 := []int{}
            p0 = (p0 + (p0 + (p1 | (^(p1)))))
            fmt.Println(p1)
            v9 := GenS1{F0: fi6, F1: -750, F2: -263}
            _ = fi7
            _ = v8
            _ = v9
        }
        switch fi6 {
        case -498:
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 uint = 42
                if false {
                    if (true || false) {
                        fmt.Println(p1, v11, fi6)
                        if false {
                            if true {
                                v12 := []int{(fi6 & 9223372036854775807), -722, (fi6 << ((2) & 63))}
                                p1 -= p0
                                v12 = append(v12, (^(fi6)))
                                _ = v12
                            } else {
                                if false {
                                    v13 := []int{2147483647, fi6, fi10, fi6}
                                    _ = v13
                                } else {
                                    v14 := GenS1{F0: (fi6 * 0), F1: -1, F2: -32768}
                                    p0 += (p1 & p1)
                                    if (4294967295 > 22) {
                                        fmt.Println(fi10)
                                        v15 := GenS1{F0: (fi10 + (fi10 + 128)), F1: -206, F2: int16(v11)}
                                        _ = v15
                                    }
                                    _ = v14
                                }
                                fmt.Println(p0, v11)
                            }
                        }
                        p1 -= p1
                    } else {
                        var v16 uint64 = 338
                        _ = v16
                    }
                }
                var v17 int64 = 0
                _ = fi10
                _ = v11
                _ = v17
            }
        case 4294967295:
            p0 = p0
        case 0:
            p1 -= -52
        default:
            fmt.Println(fi6, p1, fi6)
        }
        if false {
            p1 = p0
            var v18 int = (fi6 ^ 845)
            v19 := GenS2{F0: ("" + "a"), F1: ("hello" + ("\n\t" + ("foo bar" + "\n\t"))), F2: 22}
            _ = v18
            _ = v19
        } else {
            v20 := []int{fi6}
            if true {
                v21 := GenS2{F0: (("\n\t" + "hello") + ""), F1: "", F2: 664}
                var v22 uint8 = 1
                _ = v21
                _ = v22
            } else {
                if true {
                    if false {
                        for fi23 := 0; fi23 < 9; fi23++ {
                            if false {
                                _ = len(v20)
                                if true {
                                    fmt.Println(p1)
                                } else {
                                    var v24 int = (^(fi23))
                                    _ = v24
                                }
                                v20 = append(v20, 128)
                            } else {
                                var v25 uint8 = 127
                                var v26 uint64 = 0
                                _ = v25
                                _ = v26
                            }
                            p1 = p1
                            if true {
                                v27 := GenS3{F0: 1, F1: 849}
                                var v28 string = "a"
                                _ = v27
                                _ = v28
                            }
                            _ = fi23
                        }
                        var v29 uint = 0
                        _ = v29
                    } else {
                        v30 := GenS1{F0: 127, F1: -285, F2: -15}
                        _ = v30
                    }
                } else {
                    var v31 uint32 = 2147483648
                    if len(v20) > 0 {
                        v20[0] = (fi6 & (fi6 % (((fi6 | (fi6 >> ((65535) & 63)))) | 1)))
                    }
                    v32 := GenS3{F0: 1, F1: 65535}
                    _ = v31
                    _ = v32
                }
            }
            _ = v20
        }
        _ = fi6
    }
    if false {
        var v33 uint64 = 816
        for fi34 := 0; fi34 < 3; fi34++ {
            fmt.Println(v33)
            switch 127 {
            case 0:
                p0 = p1
            case 128:
                v35 := []int{fi34, int(v33), fi34, fi34}
                _ = v35
            default:
                for fi36 := 0; fi36 < 4; fi36++ {
                    v37 := GenS1{F0: (^(fi36)), F1: 223, F2: -1}
                    switch "foo bar" {
                    case "foo bar":
                        v38 := GenS4{F0: 244, F1: "", F2: true}
                        var v39 uint = 65535
                        _ = v38
                        _ = v39
                    case "\n\t":
                        v40 := []int{fi34, (fi36 >> ((uint(v33)) & 63)), 746, (fi34 % ((9223372036854775807) | 1))}
                        _ = v40
                    case "hello":
                        v41 := []int{(fi34 >> ((147) & 63)), fi36, (fi34 & fi34)}
                        _ = v41
                    default:
                        var v42 int16 = 42
                        _ = v42
                    }
                    _ = fi36
                    _ = v37
                }
            }
            v33 -= (^(v33))
            _ = fi34
        }
        _ = v33
    } else {
        p1 = (p0 << ((804) & 63))
        v43 := GenS4{F0: 124, F1: "a", F2: false}
        for fi44 := 0; fi44 < 6; fi44++ {
            p1 += p0
            _ = fi44
        }
        _ = v43
    }
    fmt.Println(p1, p0)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return -689
}
func genFn1(p0 int8) int16 {
    var v45 int64 = 379
    var v46 uint = 65535
    _ = p0
    _ = v45
    _ = v46
    return 1000
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{32767, -731, -82}
    var v3 string = (((("hello" + ("\n\t" + "a")) + ("hello" + "foo bar")) + "a") + "a")
    v4 := []int{127}
    var v5 uint8 = 42
    fmt.Println(v3, v3)
    v47 := v1[-2147483648]
    var v48 uint64 = 199
    if len(v2) > 0 {
        v2[0] = (v47 ^ v47)
    }
    for fi49 := 0; fi49 < 1; fi49++ {
        switch (v47 + 32768) {
        case -768:
            var v50 bool = (uint(v5) != 255)
            _ = v50
        case 127:
            fmt.Println(v48)
        case -567:
            _ = len(v4)
        default:
            v47 = v47
        }
        var v51 int = v47
        _ = fi49
        _ = v51
    }
    for fi52 := 0; fi52 < 19; fi52++ {
        switch (fi52 / (((fi52 >> ((4294967295) & 63))) | 1)) {
        case 128:
            if len(v2) > 0 {
                v2[0] = int(v48)
            }
        }
        switch v3 {
        case "":
            var v53 float32 = float32((((0.0 - 1e10) - 0.0) - (0.0 * 0.0)))
            _ = v53
        case "\n\t":
            v48 = 0
        case "a":
            v48 += (v48 + (v48 ^ v48))
        default:
            var v54 int16 = int16(v5)
            _ = v54
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi52
    }
    var v55 int64 = 9223372036854775807
    var v56 uint = 961
    v55 += 42
    v48 = v48
    for fi57 := 0; fi57 < 9; fi57++ {
        var v58 float64 = 0.0
        var v59 int64 = v55
        fmt.Println(v48, v3)
        _ = fi57
        _ = v58
        _ = v59
    }
    v60 := GenS2{F0: v3, F1: "", F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v55
    _ = v56
    _ = v60
}
