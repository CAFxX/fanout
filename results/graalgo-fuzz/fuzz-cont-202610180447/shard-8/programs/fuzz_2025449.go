package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
func genFn0(p0 int, p1 int) uint {
    fmt.Println(p0, p1)
    var v5 int32 = 2147483647
    if (-291 > 9223372036854775807) {
        p0 += p1
        if true {
            v6 := GenS1{F0: (p0 | p1), F1: 65535}
            _ = v6
        } else {
            for fi7 := 0; fi7 < 4; fi7++ {
                for fi8 := 0; fi8 < 20; fi8++ {
                    var v9 string = "a"
                    v10 := GenS1{F0: (fi8 - (p0 | 0)), F1: 7}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                var v11 int16 = -32768
                _ = fi7
                _ = v11
            }
        }
    } else {
        p0 -= p1
        p1 = (p0 + 128)
    }
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1() int {
    v12 := GenS1{F0: 0, F1: 946}
    switch ("foo bar" + "") {
    case "":
        if (54 < 42) {
            if false {
                if (-32768 < 0) {
                    var v13 int16 = -32768
                    _ = v13
                } else {
                    var v14 int = 536
                    var v15 string = "foo bar"
                    _ = v14
                    _ = v15
                }
                var v16 uint16 = 0
                switch 9223372036854775807 {
                case -2147483648:
                    v16 -= (v16 ^ (v16 % ((v16) | 1)))
                case 949:
                    v17 := []int{71}
                    _ = v17
                case 32767:
                    switch "a" {
                    case "hello":
                        fmt.Println(v16)
                    case "foo bar":
                        v18 := []int{979, -724}
                        _ = v18
                    case "":
                        for fi19 := 0; fi19 < 16; fi19++ {
                            fmt.Println(fi19, fi19)
                            v16 = 252
                            _ = fi19
                        }
                    default:
                        if true {
                            v20 := []int{}
                            var v21 uint16 = (v16 | v16)
                            switch -18 {
                            case -1:
                                fmt.Println(v16, v16)
                            }
                            _ = v20
                            _ = v21
                        } else {
                            fmt.Println(v16)
                            if ((true || (false && false)) || false) {
                                v16 = (v16 ^ v16)
                                switch "\n\t" {
                                case "foo bar":
                                    var v22 int64 = 96
                                    _ = v22
                                case "\n\t":
                                    v16 = (v16 >> ((2) & 63))
                                case "a":
                                    for fi23 := 0; fi23 < 7; fi23++ {
                                        var v24 string = "\n\t"
                                        v25 := []int{9223372036854775807, fi23, fi23}
                                        _ = fi23
                                        _ = v24
                                        _ = v25
                                    }
                                }
                                if false {
                                    v26 := GenS1{F0: 0, F1: 4294967295}
                                    _ = v26
                                } else {
                                    for fi27 := 0; fi27 < 20; fi27++ {
                                        for fi28 := 0; fi28 < 13; fi28++ {
                                            fmt.Println(v16)
                                            _ = fi28
                                        }
                                        var v29 uint32 = 65535
                                        _ = fi27
                                        _ = v29
                                    }
                                    v30 := []int{}
                                    _ = v30
                                }
                            } else {
                                v31 := []int{644, 9223372036854775807, 2147483647}
                                var v32 int = 19
                                v32 += 2147483647
                                _ = v31
                                _ = v32
                            }
                        }
                    }
                }
                _ = v16
            } else {
                v33 := GenS1{F0: 32768, F1: 2147483648}
                var v34 bool = true
                var v35 int64 = -731
                _ = v33
                _ = v34
                _ = v35
            }
        } else {
            v36 := []int{1, 127}
            _ = v36
        }
    default:
        for fi37 := 0; fi37 < 5; fi37++ {
            switch "hello" {
            case "hello":
                var v38 int = (fi37 - -1)
                _ = v38
            default:
                var v39 int16 = -32768
                _ = v39
            }
            for fi40 := 0; fi40 < 3; fi40++ {
                var v41 float32 = 3.14159
                var v42 int = fi37
                _ = fi40
                _ = v41
                _ = v42
            }
            fmt.Println(fi37)
            _ = fi37
        }
    }
    _ = v12
    return 42
}

func main() {
    var v1 uint32 = 1
    var v2 int16 = 1
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    v2 += v2
    switch -382 {
    case 127:
        v43 := GenS2{F0: true}
        _ = v43
    case 32767:
        var v44 bool = false
        _ = v44
    default:
        var v45 float64 = 0.1
        _ = v45
    }
    var v46 int32 = 1
    if false {
        switch 411 {
        case -500:
            v47 := []int{-1, 1, -1, 288}
            _ = v47
        case 4294967295:
            for fi48 := 0; fi48 < 17; fi48++ {
                v49 := []int{(fi48 / ((fi48) | 1))}
                for fi50 := 0; fi50 < 4; fi50++ {
                    _ = len(v49)
                    if true {
                        if len(v49) > 0 {
                            v49[0] = fi48
                        }
                    } else {
                        var v51 uint64 = 4294967295
                        if len(v49) > 0 {
                            v49[0] = genFn1()
                        }
                        var v52 uint = genFn0(fi50, (fi50 - (fi48 << ((1) & 63))))
                        _ = v51
                        _ = v52
                    }
                    v53 := GenS1{F0: 9223372036854775807, F1: (v1 | v1)}
                    v49 = append(v49, fi50)
                    _ = fi50
                    _ = v53
                }
                _ = fi48
                _ = v49
            }
        }
        fmt.Println(v2)
    }
    _ = len(v4)
    for fi54 := 0; fi54 < 3; fi54++ {
        v55 := GenS3{F0: 42}
        switch ("\n\t" + ("\n\t" + ("hello" + ""))) {
        case "":
            v1 += v1
        default:
            v2 -= v2
        }
        _ = fi54
        _ = v55
    }
    v56 := GenS4{F0: "hello"}
    v2 = (v2 % (((v2 >> ((601) & 63))) | 1))
    v57 := GenS1{F0: -7, F1: 772}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v56
    _ = v57
}
