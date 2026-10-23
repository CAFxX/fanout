package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 int64; F2 uint }
func genFn0() uint16 {
    if (!true) {
        v7 := []int{}
        if (((!true) && (!false)) && true) {
            var v8 bool = ((false && false) || (false || (!true)))
            v9 := GenS1{F0: 18446744073709551615}
            for fi10 := 0; fi10 < 2; fi10++ {
                _ = len(v7)
                for fi11 := 0; fi11 < 16; fi11++ {
                    var v12 bool = true
                    v12 = (1.5 > 0.1)
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
            _ = v8
            _ = v9
        }
        if (!(449 < 788)) {
            switch 4294967295 {
            case 816:
                switch -2147483648 {
                case -415:
                    v13 := GenS1{F0: 2}
                    _ = v13
                case -329:
                    switch "\n\t" {
                    case "hello":
                        v14 := GenS1{F0: 808}
                        _ = v14
                    }
                case 9223372036854775807:
                    var v15 float64 = (-2.25 + 1e10)
                    _ = v15
                }
            }
            var v16 uint = 284
            var v17 uint64 = 4294967295
            _ = v16
            _ = v17
        }
        _ = v7
    }
    v18 := GenS2{F0: 100}
    _ = v18
    return 326
}
func genFn1(p0 int64) int32 {
    p0 = p0
    p0 += p0
    var v19 uint16 = 32767
    var v20 uint16 = 65535
    var v21 string = "hello"
    _ = p0
    _ = v19
    _ = v20
    _ = v21
    return 2147483647
}

func main() {
    var v1 int = -601
    var v2 int8 = 0
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    var v5 uint = 0
    v6 := make(map[string]uint32)
    fmt.Println(v5, v2, v1)
    var v22 int32 = 0
    _ = len(v3)
    v2 -= -1
    v23 := v3[(("\n\t" + "foo bar") + "\n\t")]
    if true {
        v24 := GenS2{F0: (-(v2))}
        switch (("foo bar" + "foo bar") + "hello") {
        case "a":
            if false {
                for fi25 := 0; fi25 < 20; fi25++ {
                    fmt.Println(v2, fi25, v5)
                    if false {
                        switch (("foo bar" + "") + "\n\t") {
                        case "\n\t":
                            if (true && false) {
                                if (true && false) {
                                    if false {
                                        switch -9223372036854775808 {
                                        case -9223372036854775808:
                                            fmt.Println(v1, fi25)
                                        }
                                        v26 := []int{4294967295}
                                        switch v1 {
                                        case -2147483648:
                                            v27 := v6["a"]
                                            _ = v27
                                        case -9223372036854775808:
                                            if (genFn0() > 1) {
                                                switch "foo bar" {
                                                case "hello":
                                                    var v28 int64 = 1
                                                    _ = v28
                                                default:
                                                    for fi29 := 0; fi29 < 12; fi29++ {
                                                        var v30 uint16 = 42
                                                        _ = fi29
                                                        _ = v30
                                                    }
                                                }
                                                var v31 uint32 = 4294967295
                                                _ = v31
                                            }
                                        default:
                                            var v32 uint8 = uint8(v23)
                                            _ = v32
                                        }
                                        _ = v26
                                    }
                                    v4["\n\t"] = 42
                                } else {
                                    switch (v1 / (((v1 >> ((v5) & 63))) | 1)) {
                                    case 983:
                                        v2 += (v2 + v2)
                                    case -431:
                                        var v33 int32 = (v22 / (((v22 / (((v22 / ((v22) | 1))) | 1))) | 1))
                                        _ = v33
                                    }
                                    v34 := GenS1{F0: 65535}
                                    for fi35 := 0; fi35 < 8; fi35++ {
                                        if true {
                                            v36 := GenS3{F0: 0, F1: 695}
                                            if true {
                                                if true {
                                                    v23 += v23
                                                    v5 -= (v5 / ((v5) | 1))
                                                    if false {
                                                        var v37 uint = uint(v23)
                                                        delete(v3, "\n\t")
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v38 := GenS3{F0: 0, F1: -1}
                                                            _ = v38
                                                        case "a":
                                                            if true {
                                                                delete(v6, "a")
                                                                fmt.Println(v23)
                                                                v37 -= (^(v5))
                                                            } else {
                                                                if ((0 < (v22 / ((int32(v23)) | 1))) && (!false)) {
                                                                    if false {
                                                                        v37 = (v5 + (v5 % ((2) | 1)))
                                                                    }
                                                                    fmt.Println(v2, v22)
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        var v39 uint = (v5 + v5)
                                                                        _ = v39
                                                                    case "a":
                                                                        v40 := []int{}
                                                                        _ = v40
                                                                    case "hello":
                                                                        fmt.Println(v2)
                                                                    default:
                                                                        switch (^(v1)) {
                                                                        case 792:
                                                                            v41 := GenS4{F0: v22}
                                                                            _ = v41
                                                                        default:
                                                                            v42 := GenS5{F0: 824, F1: -9223372036854775808, F2: v37}
                                                                            _ = v42
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v43 uint = v5
                                                                    v44 := []int{(^(v1)), (v1 & 0), (v1 & 2147483647)}
                                                                    v23 -= v23
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                            }
                                                        case "":
                                                            var v45 uint32 = 4294967295
                                                            _ = v45
                                                        default:
                                                            var v46 string = "a"
                                                            _ = v46
                                                        }
                                                        v47 := GenS5{F0: 18446744073709551615, F1: 0, F2: v37}
                                                        _ = v37
                                                        _ = v47
                                                    } else {
                                                        var v48 bool = (127 != 1)
                                                        _ = v48
                                                    }
                                                } else {
                                                    fmt.Println(v22, v2, v5)
                                                    v49 := GenS5{F0: 1, F1: 0, F2: 2}
                                                    var v50 uint16 = 0
                                                    _ = v49
                                                    _ = v50
                                                }
                                            }
                                            v1 -= fi35
                                            v23 -= v23
                                            _ = v36
                                        }
                                        v22 = (v22 * genFn1(-1))
                                        fmt.Println(v5, fi35)
                                        _ = fi35
                                    }
                                    _ = v34
                                }
                            }
                        }
                    }
                    _ = fi25
                }
            }
        case "":
            v51 := GenS4{F0: genFn1(0)}
            _ = v51
        case "hello":
            if true {
                if (true && false) {
                    v23 -= (v23 ^ 2147483648)
                } else {
                    for fi52 := 0; fi52 < 13; fi52++ {
                        v53 := []int{2147483647, (v1 * v1), (v1 << ((v5) & 63)), (v1 / ((fi52) | 1))}
                        if false {
                            v54 := GenS1{F0: 42}
                            _ = v54
                        }
                        v23 -= 2
                        _ = fi52
                        _ = v53
                    }
                }
                for fi55 := 0; fi55 < 4; fi55++ {
                    var v56 int64 = 0
                    switch "foo bar" {
                    case "foo bar":
                        v57 := []int{(fi55 & (-(v1))), v1, -9223372036854775808, 2147483647}
                        _ = v57
                    }
                    _ = fi55
                    _ = v56
                }
                v58 := []int{v1, (v1 ^ v1), (v1 | v1), (v1 & -217)}
                _ = v58
            }
        default:
            v59 := GenS4{F0: (^(v22))}
            _ = v59
        }
        _ = v24
    }
    v60 := []int{}
    var v61 bool = (false || true)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v60:", len(v60), v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
