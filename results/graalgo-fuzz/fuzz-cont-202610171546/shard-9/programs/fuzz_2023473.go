package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 bool; F1 string }
func genFn0(p0 int32) uint {
    var v4 uint = 867
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := []int{}
        fmt.Println(p0, v4, p0)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        v4 += v4
        if false {
            v8 := []int{(-(fi7)), -760, fi7, 4294967295}
            _ = v8
        }
        var v9 float64 = (1.5 * 0.0)
        _ = fi7
        _ = v9
    }
    _ = p0
    _ = v4
    return 4294967295
}
func genFn1() bool {
    if false {
        v10 := GenS1{F0: 745, F1: 42, F2: 18446744073709551615}
        var v11 uint8 = 76
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        if (1e10 == 0.0) {
            var v13 string = "\n\t"
            _ = v13
        }
        if false {
            if ("hello" < "\n\t") {
                var v14 float64 = 0.1
                v15 := GenS2{F0: "", F1: -1}
                v16 := GenS3{F0: false}
                _ = v14
                _ = v15
                _ = v16
            }
            v17 := GenS4{F0: 21}
            _ = v17
        }
        if false {
            v18 := []int{fi12, (fi12 << ((705) & 63))}
            v19 := GenS1{F0: 0, F1: 2, F2: 42}
            var v20 int = (-(fi12))
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi12
    }
    if (249 == 1) {
        var v21 uint8 = 2
        v21 -= (v21 | 233)
        _ = v21
    } else {
        v22 := []int{-857, 2147483647}
        _ = v22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        v24 := []int{1, (fi23 / ((fi23) | 1)), fi23, (-(fi23))}
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 float32 = 3.14159
            v24 = append(v24, fi25)
            var v27 uint = 2
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    var v28 uint16 = 0
    _ = v28
    return (2147483647 < int32(v28))
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 534
    var v3 uint32 = 42
    var v29 uint8 = 2
    var v30 int16 = -32768
    if true {
        var v31 uint16 = 720
        v32 := GenS5{F0: true, F1: (("" + "\n\t") + ("hello" + "foo bar"))}
        _ = v31
        _ = v32
    }
    var v33 float64 = float64(0.0)
    var v34 int64 = -9223372036854775808
    var v35 uint64 = 4294967295
    var v36 uint64 = v35
    switch 0 {
    case 500:
        switch "foo bar" {
        case "foo bar":
            if (v29 == v29) {
                if ((v30 | (v30 % ((-1) | 1))) == (^(v30))) {
                    switch "hello" {
                    case "":
                        var v37 uint32 = v3
                        _ = v37
                    case "hello":
                        if (uint64(v29) > (v36 << ((0) & 63))) {
                            v38 := GenS5{F0: ((!false) && false), F1: ""}
                            fmt.Println(v36, v34, v2)
                            var v39 uint32 = (v3 % ((4294967295) | 1))
                            _ = v38
                            _ = v39
                        }
                    }
                    if (!false) {
                        v36 += v36
                        var v40 string = ""
                        v2 -= 32767
                        _ = v40
                    } else {
                        var v41 uint32 = v3
                        if true {
                            var v42 int32 = 511
                            _ = v42
                        } else {
                            for fi43 := 0; fi43 < 16; fi43++ {
                                var v44 int16 = v30
                                v45 := GenS6{F0: true, F1: "a"}
                                v46 := GenS6{F0: (-1 != fi43), F1: ""}
                                _ = fi43
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                            for fi47 := 0; fi47 < 19; fi47++ {
                                for fi48 := 0; fi48 < 5; fi48++ {
                                    var v49 uint64 = (^(v36))
                                    _ = fi48
                                    _ = v49
                                }
                                for fi50 := 0; fi50 < 8; fi50++ {
                                    var v51 bool = (2 == (^(v2)))
                                    switch (fi50 << ((10) & 63)) {
                                    case 127:
                                        v52 := GenS6{F0: (((1.5 / 3.14159) / 3.14159) == 0.0), F1: "foo bar"}
                                        _ = v52
                                    }
                                    v33 -= v33
                                    v53 := []int{127}
                                    _ = fi50
                                    _ = v51
                                    _ = v53
                                }
                                _ = fi47
                            }
                        }
                        var v54 uint32 = (v41 ^ 650)
                        _ = v41
                        _ = v54
                    }
                }
                var v55 bool = true
                var v56 float64 = v33
                _ = v55
                _ = v56
            } else {
                delete(v1, -9223372036854775808)
            }
        case "":
            v57 := GenS1{F0: -196, F1: 2, F2: uint64(v29)}
            _ = v57
        }
    }
    var v58 int32 = int32(v3)
    if false {
        var v59 int8 = -97
        for fi60 := 0; fi60 < 6; fi60++ {
            v35 = (v36 / ((v35) | 1))
            var v61 uint8 = (v29 ^ 42)
            _ = fi60
            _ = v61
        }
        _ = v59
    } else {
        var v62 float32 = 0.0
        var v63 uint16 = 1000
        _ = v62
        _ = v63
    }
    if (true && false) {
        var v64 float32 = -2.25
        _ = v64
    }
    fmt.Println(v30, v33, v36)
    fmt.Println(v35, v58, v29)
    if ((v30 ^ (^(v30))) >= v30) {
        for fi65 := 0; fi65 < 8; fi65++ {
            v66 := []int{fi65, -591}
            _ = fi65
            _ = v66
        }
    } else {
        if (v58 == (v58 | 692)) {
            v33 = float64((float32(v33) + 0.0))
        }
        if (-257 < -956) {
            var v67 int64 = (v34 >> ((genFn0(int32(v34))) & 63))
            v35 -= v36
            _ = v67
        } else {
            v1[9223372036854775807] = "hello"
        }
        v58 -= v58
    }
    if true {
        for fi68 := 0; fi68 < 7; fi68++ {
            if false {
                v69 := []int{}
                var v70 uint = 42
                _ = v69
                _ = v70
            }
            var v71 uint = 42
            _ = fi68
            _ = v71
        }
    } else {
        if false {
            var v72 float32 = -2.25
            var v73 uint8 = (v29 + v29)
            for fi74 := 0; fi74 < 1; fi74++ {
                if (!((-32768 == 739) && (((^(v58)) > v58) || (v72 == (v72 * 3.14159))))) {
                    v36 -= 1
                }
                _ = fi74
            }
            _ = v72
            _ = v73
        }
        switch "\n\t" {
        case "foo bar":
            v2 += (^(v2))
        case "hello":
            switch "a" {
            case "a":
                var v75 int64 = v34
                _ = v75
            }
        default:
            var v76 uint32 = (^(v3))
            _ = v76
        }
        var v77 uint16 = 2
        _ = v77
    }
    var v78 int8 = -77
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v58:", v58)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v58
    _ = v78
}
