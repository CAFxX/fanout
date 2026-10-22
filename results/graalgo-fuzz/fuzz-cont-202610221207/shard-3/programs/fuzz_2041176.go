package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0() uint16 {
    if (float32(3.14159) != 0.0) {
        var v7 bool = false
        var v8 float32 = ((0.0 * (-2.25 - 3.14159)) / 1e10)
        v9 := GenS1{F0: 1, F1: 264}
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 int = 0
        var v11 uint8 = 42
        var v12 int16 = 42
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 int8 = 100
    v13 = v13
    var v14 string = "hello"
    _ = v13
    _ = v14
    return 1
}
func genFn1(p0 int64, p1 string, p2 uint32) uint32 {
    var v15 uint16 = genFn0()
    var v16 uint32 = (p2 / ((p2) | 1))
    if true {
        var v17 bool = false
        if (4294967295 != 255) {
            v15 = (^(v15))
            if v17 {
                var v18 bool = false
                var v19 uint = 768
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 uint = (v19 * (v19 - 1))
                    _ = fi20
                    _ = v21
                }
                _ = v18
                _ = v19
            }
            v22 := GenS1{F0: 0, F1: -1}
            _ = v22
        } else {
            var v23 uint = 231
            switch p1 {
            case "a":
                if false {
                    fmt.Println(v16)
                }
            default:
                for fi24 := 0; fi24 < 1; fi24++ {
                    v25 := []int{}
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                    v25 = append(v25, 1)
                    _ = fi24
                    _ = v25
                }
            }
            _ = v23
        }
        var v26 int16 = 0
        _ = v17
        _ = v26
    }
    if (!(uint64(p2) != 92)) {
        for fi27 := 0; fi27 < 1; fi27++ {
            if (42 != 255) {
                var v28 uint16 = v15
                v15 -= 42
                var v29 int16 = 32767
                _ = v28
                _ = v29
            }
            fmt.Println(p1, p2)
            _ = fi27
        }
        v30 := GenS1{F0: 67, F1: 0}
        v31 := GenS1{F0: -1, F1: -329}
        _ = v30
        _ = v31
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return (v16 | v16)
}

func main() {
    var v1 int16 = 1000
    var v2 uint16 = 32767
    var v3 uint32 = 65535
    var v4 int8 = -69
    var v5 int16 = v1
    var v6 uint = 18446744073709551615
    var v32 string = ("\n\t" + "foo bar")
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 uint = 806
        if (false || (float64(1e10) != 1e10)) {
            switch fi33 {
            case 0:
                switch (fi33 % ((fi33) | 1)) {
                case 424:
                    for fi35 := 0; fi35 < 20; fi35++ {
                        if true {
                            var v36 uint32 = v3
                            v32 = "a"
                            switch v32 {
                            case "foo bar":
                                if true {
                                    fmt.Println(v5)
                                } else {
                                    var v37 string = ("\n\t" + "\n\t")
                                    v1 = v5
                                    v38 := []int{(fi33 & 2147483647), fi35, (fi35 | int(v3)), int(v2)}
                                    _ = v37
                                    _ = v38
                                }
                            case "hello":
                                v32 = v32
                            }
                            _ = v36
                        }
                        fmt.Println(fi35, v34, fi33)
                        v39 := GenS1{F0: (v4 << (((v34 & 843)) & 63)), F1: (v1 / ((v5) | 1))}
                        _ = fi35
                        _ = v39
                    }
                }
            case 127:
                for fi40 := 0; fi40 < 12; fi40++ {
                    v41 := GenS1{F0: v4, F1: (-(v5))}
                    _ = fi40
                    _ = v41
                }
            case -372:
                fmt.Println(v5, v3)
            default:
                v2 = 255
            }
            var v42 uint32 = (v3 << ((18446744073709551615) & 63))
            _ = v42
        }
        if true {
            if false {
                v43 := GenS2{F0: (fi33 | (fi33 | fi33)), F1: v2}
                v3 += genFn1(42, "a", (v3 - v3))
                _ = v43
            }
            v2 += (v2 << (((v34 ^ v6)) & 63))
        } else {
            v2 = v2
            for fi44 := 0; fi44 < 18; fi44++ {
                v45 := GenS1{F0: v4, F1: int16(v4)}
                _ = fi44
                _ = v45
            }
            v46 := GenS1{F0: int8(v6), F1: v1}
            _ = v46
        }
        _ = fi33
        _ = v34
    }
    var v47 uint8 = 2
    if false {
        fmt.Println(v5, v5, v3)
        v48 := GenS1{F0: (v4 | (v4 & int8(v3))), F1: (v5 | v5)}
        _ = v48
    }
    for fi49 := 0; fi49 < 15; fi49++ {
        v50 := GenS2{F0: (fi49 & 32767), F1: v2}
        if true {
            if false {
                v51 := GenS2{F0: -751, F1: (v2 & 910)}
                _ = v51
            }
            for fi52 := 0; fi52 < 13; fi52++ {
                if (1 <= 2147483647) {
                    var v53 uint = (v6 ^ v6)
                    v54 := GenS2{F0: fi52, F1: (v2 * 1)}
                    fmt.Println(fi49, fi52)
                    var v55 int16 = (v5 << ((v6) & 63))
                    var v56 int = fi52
                    _ = v53
                    _ = v54
                    _ = v55
                    _ = v56
                } else {
                    if false {
                        switch (fi52 + (fi52 >> ((v6) & 63))) {
                        case 167:
                            var v57 uint = 772
                            _ = v57
                        case -439:
                            var v58 float32 = (((((1e10 / 1.5) / float32(float64((3.14159 * 1.5)))) + (1.5 / 1.0)) * float32(1e10)) * 1e10)
                            _ = v58
                        }
                        var v59 uint = v6
                        v2 = genFn0()
                        v5 = v5
                        var v60 bool = (((fi52 > (fi52 + (fi52 & (fi49 * fi52)))) && true) && ((3.14159 - (float64(-2.25) + -2.25)) != (0.0 + float64((3.14159 * ((0.0 - -2.25) * (3.14159 / 1.0)))))))
                        _ = v59
                        _ = v60
                    }
                }
                _ = fi52
            }
            v61 := []int{605}
            _ = v61
        }
        _ = fi49
        _ = v50
    }
    for fi62 := 0; fi62 < 1; fi62++ {
        v63 := GenS2{F0: (fi62 - fi62), F1: 255}
        var v64 uint16 = genFn0()
        _ = fi62
        _ = v63
        _ = v64
    }
    switch (v32 + "foo bar") {
    case "hello":
        var v65 int = -121
        _ = v65
    case "a":
        v66 := []int{}
        _ = v66
    }
    v67 := []int{2147483647, 42}
    for fi68 := 0; fi68 < 14; fi68++ {
        fmt.Println(v1, v3, v3)
        v3 -= (v3 & uint32(v2))
        v67 = append(v67, -2)
        _ = fi68
    }
    switch 2147483647 {
    case 4294967295:
        for fi69 := 0; fi69 < 14; fi69++ {
            var v70 bool = (true || (1 == 2))
            _ = fi69
            _ = v70
        }
    default:
        var v71 uint64 = 2
        _ = v71
    }
    v1 = -283
    if false {
        v5 -= 530
    } else {
        if true {
            var v72 float64 = float64((-2.25 * 3.14159))
            _ = v72
        } else {
            var v73 uint = v6
            for fi74 := 0; fi74 < 10; fi74++ {
                fmt.Println(v1, v2, v3)
                var v75 uint = v6
                _ = fi74
                _ = v75
            }
            _ = v73
        }
        for fi76 := 0; fi76 < 19; fi76++ {
            var v77 int8 = 100
            _ = fi76
            _ = v77
        }
    }
    var v78 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v67
    _ = v78
}
