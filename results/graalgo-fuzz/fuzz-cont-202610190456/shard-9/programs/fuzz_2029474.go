package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint }
type GenS2 struct { F0 uint16; F1 int; F2 uint }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint8, p1 uint64, p2 bool) int16 {
    v7 := GenS1{F0: "hello", F1: (p0 * (p0 | 2)), F2: 178}
    var v8 float32 = 1.5
    if p2 {
        var v9 uint8 = p0
        v10 := []int{-887, 4294967295}
        _ = v10
        _ = v9
    }
    v11 := []int{int(p0), 87, 1}
    v12 := GenS1{F0: "\n\t", F1: p0, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v7
    _ = v8
    return -32768
}
func genFn1(p0 string) int32 {
    var v13 uint = 476
    var v14 bool = false
    v15 := []int{42, 42, 1, 0}
    v14 = (0 != -32768)
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    return 0
}

func main() {
    v1 := []string{}
    var v2 float32 = -2.25
    var v3 uint = 4294967295
    v4 := make(map[uint32]int)
    var v5 float64 = 1e10
    v6 := make(map[uint32]int)
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 int8 = 127
        v2 = (v2 * v2)
        var v18 uint8 = 42
        _ = fi16
        _ = v17
        _ = v18
    }
    v3 += 267
    var v19 uint8 = 2
    _ = len(v6)
    fmt.Println(v2, v5)
    for fi20 := 0; fi20 < 20; fi20++ {
        v21 := GenS2{F0: 255, F1: fi20, F2: v3}
        var v22 string = "a"
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println(v5, v19)
    var v23 uint64 = 42
    v24 := GenS3{F0: (v3 * v3), F1: 1}
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 float32 = v2
        v26 += 3.14159
        v19 = (^(v19))
        _ = fi25
        _ = v26
    }
    if true {
        if true {
            var v27 uint32 = 4294967295
            if false {
                v1 = append(v1, "hello")
                _ = len(v4)
                v28 := GenS1{F0: "", F1: v19, F2: 42}
                _ = v28
            }
            _ = v27
        } else {
            var v29 bool = false
            switch 77 {
            case 4294967295:
                _ = len(v1)
            case 192:
                switch 78 {
                case 978:
                    if (true && (!(!(uint64(v3) >= (v23 % (((v23 / ((v23) | 1))) | 1)))))) {
                        v30 := GenS3{F0: v3, F1: -1}
                        _ = v30
                    } else {
                        var v31 int32 = 0
                        v32 := GenS4{F0: (^(v23)), F1: 9223372036854775807}
                        _ = v31
                        _ = v32
                    }
                case 42:
                    var v33 string = ("foo bar" + "")
                    _ = v33
                default:
                    var v34 int8 = -84
                    _ = v34
                }
            }
            switch "" {
            case "":
                if (!(1 >= -32768)) {
                    for fi35 := 0; fi35 < 18; fi35++ {
                        fmt.Println(v29, v2, v3)
                        for fi36 := 0; fi36 < 17; fi36++ {
                            v19 = v19
                            _ = fi36
                        }
                        fmt.Println(v3)
                        _ = fi35
                    }
                    for fi37 := 0; fi37 < 20; fi37++ {
                        v3 = (v3 | 747)
                        _ = fi37
                    }
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                }
            case "\n\t":
                var v38 uint32 = 0
                _ = v38
            case "foo bar":
                v6[2147483648] = -1
            }
            _ = v29
        }
        if (false && false) {
            if false {
                v19 += (v19 & (v19 + v19))
                for fi39 := 0; fi39 < 5; fi39++ {
                    var v40 uint16 = 42
                    _ = fi39
                    _ = v40
                }
                if true {
                    switch "a" {
                    case "foo bar":
                        for fi41 := 0; fi41 < 13; fi41++ {
                            v42 := GenS4{F0: v23, F1: 0}
                            var v43 uint64 = v23
                            v44 := GenS5{F0: v3, F1: (^(v43))}
                            _ = fi41
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                    case "a":
                        var v45 int = 2147483647
                        _ = v45
                    case "":
                        v19 -= 56
                    }
                    var v46 int8 = 31
                    _ = v46
                }
            } else {
                if true {
                    if (-1 <= 335) {
                        v23 = (v23 % (((^(v23))) | 1))
                        if false {
                            for fi47 := 0; fi47 < 19; fi47++ {
                                v2 = v2
                                _ = fi47
                            }
                            var v48 int8 = 1
                            var v49 int16 = -32768
                            _ = v48
                            _ = v49
                        } else {
                            v50 := GenS3{F0: v3, F1: -775}
                            _ = len(v4)
                            _ = v50
                        }
                    } else {
                        v5 = v5
                    }
                    var v51 string = ("\n\t" + ("hello" + "a"))
                    _ = v51
                } else {
                    if len(v1) > 0 {
                        v1[0] = (("\n\t" + "a") + "\n\t")
                    }
                }
                v5 -= float64((1e10 + v2))
            }
            v52 := GenS4{F0: (^(v23)), F1: 777}
            _ = v52
        }
        var v53 uint16 = 1
        _ = v53
    } else {
        v19 = v19
        if (false || (true && (v2 >= v2))) {
            if false {
                v3 -= v3
                var v54 int32 = 1
                _ = v54
            }
            for fi55 := 0; fi55 < 5; fi55++ {
                if (fi55 < (fi55 ^ fi55)) {
                    var v56 uint32 = 1
                    _ = v56
                }
                if true {
                    var v57 uint32 = 2147483648
                    _ = v57
                }
                _ = fi55
            }
        }
    }
    var v58 int16 = 1000
    var v59 uint = v3
    v5 -= v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
}
