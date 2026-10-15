package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int16; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32; F1 uint64 }
func genFn0() bool {
    if true {
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := GenS1{F0: 32767, F1: -1, F2: 764}
            if false {
                v7 := GenS2{F0: 1, F1: 70}
                v8 := []int{fi5, fi5, (fi5 % ((fi5) | 1)), fi5}
                _ = v7
                _ = v8
            }
            _ = fi5
            _ = v6
        }
    }
    var v9 bool = (((42 == 10) || (311 < 0)) && false)
    fmt.Println(v9)
    if (false && v9) {
        if v9 {
            for fi10 := 0; fi10 < 2; fi10++ {
                for fi11 := 0; fi11 < 5; fi11++ {
                    v12 := []int{354, -4, (fi10 & 42)}
                    _ = fi11
                    _ = v12
                }
                v13 := GenS1{F0: 42, F1: 1, F2: 1}
                for fi14 := 0; fi14 < 8; fi14++ {
                    switch "hello" {
                    case "\n\t":
                        var v15 uint32 = 2147483647
                        var v16 string = "\n\t"
                        _ = v15
                        _ = v16
                    case "a":
                        if true {
                            v17 := GenS1{F0: 1, F1: 0, F2: 0}
                            v18 := []int{fi14, (fi14 % (((-(fi10))) | 1)), (fi14 << ((18446744073709551615) & 63))}
                            _ = v17
                            _ = v18
                        }
                    case "foo bar":
                        if (2 != 42) {
                            var v19 int8 = -1
                            if false {
                                var v20 int16 = 42
                                switch "\n\t" {
                                case "":
                                    fmt.Println(fi10, fi10, fi14)
                                case "foo bar":
                                    var v21 uint32 = 1
                                    _ = v21
                                case "a":
                                    v22 := GenS3{F0: 58, F1: -1}
                                    _ = v22
                                default:
                                    var v23 float32 = 3.14159
                                    _ = v23
                                }
                                var v24 int8 = v19
                                _ = v20
                                _ = v24
                            }
                            if false {
                                switch ("\n\t" + (("\n\t" + "") + (("\n\t" + "a") + "\n\t"))) {
                                case "foo bar":
                                    v25 := []int{128}
                                    _ = v25
                                case "a":
                                    fmt.Println(v19, fi10, fi14)
                                }
                                var v26 bool = v9
                                _ = v26
                            } else {
                                fmt.Println(fi14)
                                if v9 {
                                    v9 = true
                                    var v27 int = fi10
                                    _ = v27
                                } else {
                                    var v28 uint = 4294967295
                                    _ = v28
                                }
                                switch fi10 {
                                case 1:
                                    var v29 int16 = 32767
                                    _ = v29
                                case -1:
                                    var v30 uint64 = 4294967295
                                    _ = v30
                                default:
                                    v31 := GenS2{F0: 1, F1: 449}
                                    _ = v31
                                }
                                if true {
                                    v32 := []int{}
                                    if ((!v9) && true) {
                                        var v33 float64 = 0.1
                                        v34 := GenS3{F0: 127, F1: 1}
                                        var v35 int16 = -388
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                    _ = v32
                                } else {
                                    fmt.Println(fi14, v9)
                                    if (true || true) {
                                        v19 = (v19 << ((65535) & 63))
                                        v19 = (v19 >> ((42) & 63))
                                        if (1.5 <= (0.1 + ((((float64((float32(3.14159) * 0.0)) - (float64(3.14159) * 0.1)) - 1e10) - 0.1) / float64(1.5)))) {
                                            fmt.Println(v9, v9, fi14)
                                        } else {
                                            var v36 uint16 = 0
                                            _ = v36
                                        }
                                    } else {
                                        if true {
                                            v37 := []int{fi10}
                                            _ = v37
                                        }
                                        v38 := GenS4{F0: 1000, F1: -32768, F2: 30}
                                        _ = v38
                                    }
                                    var v39 float32 = float32((-2.25 / 1e10))
                                    var v40 uint8 = 255
                                    _ = v39
                                    _ = v40
                                }
                                v41 := GenS3{F0: 0, F1: -452}
                                switch "hello" {
                                case "hello":
                                    v42 := GenS3{F0: 128, F1: -1}
                                    _ = v42
                                default:
                                    var v43 uint32 = 2147483647
                                    _ = v43
                                }
                                _ = v41
                            }
                            _ = v19
                        } else {
                            switch fi14 {
                            case 42:
                                var v44 bool = true
                                _ = v44
                            default:
                                var v45 uint16 = 255
                                _ = v45
                            }
                        }
                    default:
                        v9 = v9
                    }
                    var v46 uint = 2
                    v47 := []int{(^(fi14)), 0}
                    _ = fi14
                    _ = v46
                    _ = v47
                }
                _ = fi10
                _ = v13
            }
            fmt.Println(v9)
            fmt.Println(v9)
        }
    }
    v9 = false
    _ = v9
    return (v9 || (!false))
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    var v3 uint = 435
    v4 := make(map[int]int)
    for fi48 := 0; fi48 < 3; fi48++ {
        var v49 int64 = 9223372036854775807
        var v50 int32 = 42
        _ = fi48
        _ = v49
        _ = v50
    }
    v51 := GenS5{F0: 18446744073709551615, F1: 32768}
    if ((true || (!(1000 >= 1000))) && false) {
        for fi52 := 0; fi52 < 3; fi52++ {
            var v53 int64 = 9223372036854775807
            _ = fi52
            _ = v53
        }
        fmt.Println(v3)
    } else {
        if false {
            var v54 uint16 = 32767
            _ = v54
        }
        if false {
            delete(v4, 862)
        } else {
            v3 = 4294967295
        }
    }
    switch -9223372036854775808 {
    case 32767:
        _ = len(v4)
    case 1:
        if (32767 >= -736) {
            v55 := GenS6{F0: 127}
            var v56 int32 = 1
            var v57 float32 = 1.5
            _ = v55
            _ = v56
            _ = v57
        }
    }
    var v58 uint16 = 42
    if (!((202 > 42) || false)) {
        var v59 string = "\n\t"
        var v60 string = (v59 + (v59 + ("a" + "foo bar")))
        v60 = v60
        _ = v59
        _ = v60
    }
    v3 += (v3 - 789)
    var v61 int = -9223372036854775808
    v3 = (v3 % ((1) | 1))
    fmt.Println(v61, v61, v3)
    v62 := []int{v61, (v61 << ((2) & 63))}
    switch (v61 & v61) {
    case -368:
        v61 -= 42
    case 4294967295:
        for fi63 := 0; fi63 < 9; fi63++ {
            v64 := GenS7{F0: -2147483648, F1: 42}
            _ = fi63
            _ = v64
        }
    case 9223372036854775807:
        v3 += 284
    }
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v58:", v58)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v58
    _ = v61
    _ = v62
}
