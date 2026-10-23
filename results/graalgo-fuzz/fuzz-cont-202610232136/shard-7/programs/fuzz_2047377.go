package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 int8; F1 int32; F2 int16 }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 int16; F1 bool }
type GenS10 struct { F0 int8; F1 uint64 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint16; F1 bool }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int8, p1 bool) uint16 {
    fmt.Println(p1, p0)
    if false {
        fmt.Println(p1, p1)
        p0 = (p0 & p0)
        for fi5 := 0; fi5 < 15; fi5++ {
            v6 := GenS1{F0: 875, F1: 9223372036854775807}
            if (4294967295 <= 355) {
                v7 := GenS1{F0: 0, F1: 9223372036854775807}
                v8 := GenS2{F0: 1, F1: 0}
                _ = v7
                _ = v8
            }
            _ = fi5
            _ = v6
        }
    }
    v9 := []int{-822, -763, 4294967295, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 int8 = 20
    v4 := make(map[string]uint32)
    v10 := GenS3{F0: -658}
    fmt.Println(v3)
    switch 128 {
    case 242:
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 int16 = 32767
            var v13 uint16 = 255
            v1[-1] = 850
            _ = fi11
            _ = v12
            _ = v13
        }
    case -9223372036854775808:
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS4{F0: 2}
            switch "hello" {
            case "a":
                if true {
                    v16 := GenS1{F0: -32768, F1: -810}
                    _ = v16
                }
            }
            fmt.Println(v3, v3)
            _ = fi14
            _ = v15
        }
    case 32768:
        var v17 int64 = 42
        _ = v17
    }
    v3 = (v3 % ((v3) | 1))
    v18 := GenS4{F0: 42}
    fmt.Println(v3)
    var v19 int = -558
    if true {
        v19 -= v19
        v20 := v4["\n\t"]
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v3 = v3
        _ = fi21
    }
    v19 += v19
    switch "" {
    case "foo bar":
        if true {
            v19 = 2147483647
            v3 -= 1
        }
    case "\n\t":
        switch (^(v19)) {
        case 2147483647:
            var v22 int64 = 42
            _ = v22
        case 4294967295:
            v23 := GenS5{F0: false, F1: (v3 << ((364) & 63))}
            _ = v23
        case 9223372036854775807:
            if false {
                v24 := GenS3{F0: 42}
                if false {
                    v25 := GenS6{F0: 512, F1: 2147483647}
                    _ = v25
                }
                v26 := []int{(v19 & (v19 + 32767)), (v19 >> ((515) & 63))}
                _ = v24
                _ = v26
            } else {
                for fi27 := 0; fi27 < 10; fi27++ {
                    var v28 int16 = 32767
                    switch ("foo bar" + ("hello" + "a")) {
                    case "a":
                        switch v19 {
                        case -614:
                            v29 := GenS7{F0: v3, F1: 182, F2: (v28 / (((v28 & (v28 >> ((2) & 63)))) | 1))}
                            _ = v29
                        default:
                            if true {
                                var v30 uint32 = 361
                                switch fi27 {
                                case 32767:
                                    v31 := GenS8{F0: 0, F1: (v30 << ((4294967295) & 63))}
                                    _ = v31
                                case 32768:
                                    for fi32 := 0; fi32 < 2; fi32++ {
                                        if ("" < "") {
                                            var v33 string = "foo bar"
                                            v28 -= v28
                                            var v34 uint = 255
                                            v35 := []int{-1, fi27, fi27, fi27}
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        } else {
                                            v19 += 0
                                            v30 = (v30 | (v30 >> ((1) & 63)))
                                            v28 += v28
                                        }
                                        switch fi27 {
                                        case -899:
                                            var v36 float64 = 1.5
                                            _ = v36
                                        default:
                                            if true {
                                                var v37 uint16 = 94
                                                if false {
                                                    delete(v4, "a")
                                                    fmt.Println(v28)
                                                    var v38 uint8 = 1
                                                    _ = v38
                                                } else {
                                                    fmt.Println(v19)
                                                    if true {
                                                        var v39 uint64 = 1
                                                        _ = v39
                                                    }
                                                    if false {
                                                        v19 += (fi32 - fi27)
                                                        if false {
                                                            fmt.Println(v19, v37)
                                                        } else {
                                                            v40 := GenS9{F0: -32768, F1: (true || false)}
                                                            if (-2.25 > 0.0) {
                                                                var v41 int64 = 42
                                                                _ = v41
                                                            } else {
                                                                fmt.Println(v3, v19)
                                                                _ = len(v2)
                                                                fmt.Println(v30)
                                                            }
                                                            _ = v40
                                                        }
                                                    } else {
                                                        v37 = genFn0(v3, (663 > 1))
                                                        v28 -= v28
                                                        if (!(!false)) {
                                                            v42 := GenS3{F0: -303}
                                                            if (!((2147483647 <= 1) && false)) {
                                                                var v43 int = v19
                                                                fmt.Println(fi32, fi27)
                                                                _ = v43
                                                            } else {
                                                                switch (fi32 + (fi27 % ((fi32) | 1))) {
                                                                case 32768:
                                                                    var v44 float32 = (1.5 + -2.25)
                                                                    _ = v44
                                                                case 2147483647:
                                                                    v37 = (v37 << ((4294967295) & 63))
                                                                case 0:
                                                                    if true {
                                                                        v45 := GenS10{F0: v3, F1: 893}
                                                                        if true {
                                                                            v46 := GenS1{F0: (v28 * v28), F1: -143}
                                                                            var v47 string = "hello"
                                                                            v48 := []int{fi27, (fi27 % ((137) | 1))}
                                                                            _ = v46
                                                                            _ = v47
                                                                            _ = v48
                                                                        }
                                                                        _ = v45
                                                                    }
                                                                }
                                                                v49 := GenS1{F0: (v28 - v28), F1: 9223372036854775807}
                                                                _ = v49
                                                            }
                                                            _ = v42
                                                        }
                                                    }
                                                }
                                                var v50 bool = ((!true) || (-2.25 >= 1.5))
                                                _ = v37
                                                _ = v50
                                            } else {
                                                fmt.Println(fi27, v28, v19)
                                            }
                                        }
                                        v3 -= (v3 * 127)
                                        _ = fi32
                                    }
                                case -392:
                                    for fi51 := 0; fi51 < 9; fi51++ {
                                        v52 := GenS4{F0: 0}
                                        if true {
                                            var v53 int32 = -2147483648
                                            v54 := GenS2{F0: 516, F1: 4294967295}
                                            fmt.Println(v19, v28, v19)
                                            v55 := GenS9{F0: (v28 | -707), F1: (!(v30 != v30))}
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        }
                                        _ = fi51
                                        _ = v52
                                    }
                                default:
                                    if (127 != 255) {
                                        switch "" {
                                        case "\n\t":
                                            switch v19 {
                                            case -464:
                                                v56 := GenS7{F0: (v3 << ((42) & 63)), F1: 528, F2: v28}
                                                _ = v56
                                            case -9223372036854775808:
                                                fmt.Println(fi27, v30, v28)
                                            case 32767:
                                                v57 := GenS11{F0: ""}
                                                _ = v57
                                            }
                                        case "a":
                                            switch "hello" {
                                            case "\n\t":
                                                if (float32(1.5) <= 3.14159) {
                                                    var v58 int32 = 0
                                                    _ = v58
                                                } else {
                                                    fmt.Println(v3, v30, v30)
                                                }
                                            case "hello":
                                                for fi59 := 0; fi59 < 11; fi59++ {
                                                    v60 := GenS1{F0: v28, F1: -1}
                                                    v19 += (v19 | (fi27 | 42))
                                                    _ = fi59
                                                    _ = v60
                                                }
                                            default:
                                                for fi61 := 0; fi61 < 10; fi61++ {
                                                    if false {
                                                        if true {
                                                            v30 += (v30 + v30)
                                                            switch fi61 {
                                                            case 32768:
                                                                _ = len(v2)
                                                            default:
                                                                if true {
                                                                    v62 := GenS10{F0: v3, F1: 2}
                                                                    v63 := v1[(fi27 + (^(fi61)))]
                                                                    _ = v62
                                                                    _ = v63
                                                                } else {
                                                                    v64 := GenS12{F0: 892, F1: false}
                                                                    var v65 uint8 = 2
                                                                    delete(v2, "\n\t")
                                                                    _ = v64
                                                                    _ = v65
                                                                }
                                                            }
                                                        }
                                                        var v66 int64 = 1
                                                        if true {
                                                            v67 := GenS13{F0: 255}
                                                            if false {
                                                                v28 = (v28 | v28)
                                                                fmt.Println(v28)
                                                                fmt.Println(v19, fi61)
                                                            }
                                                            v68 := GenS14{F0: 2, F1: 903}
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                        v3 = v3
                                                        _ = v66
                                                    } else {
                                                        fmt.Println(fi61, v30, fi27)
                                                        var v69 uint8 = 0
                                                        _ = v69
                                                    }
                                                    v28 -= (v28 << ((48) & 63))
                                                    _ = fi61
                                                }
                                            }
                                        }
                                        v19 = 783
                                        v3 = (v3 - (v3 | v3))
                                    }
                                }
                                _ = v30
                            } else {
                                v3 = 100
                            }
                        }
                    }
                    _ = fi27
                    _ = v28
                }
                v70 := []int{(v19 - v19), v19, (v19 / ((v19) | 1))}
                v70 = append(v70, v19)
                _ = v70
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
