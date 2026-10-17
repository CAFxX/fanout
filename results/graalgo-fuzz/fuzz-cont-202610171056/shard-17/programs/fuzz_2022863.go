package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint16; F1 int8; F2 string }
type GenS5 struct { F0 int8 }
func genFn0(p0 string, p1 int8, p2 string) int64 {
    if true {
        v4 := GenS1{F0: 1, F1: 871}
        var v5 int8 = p1
        _ = v4
        _ = v5
    } else {
        var v6 uint16 = 1
        var v7 uint64 = 0
        fmt.Println(p1, p2)
        _ = v6
        _ = v7
    }
    fmt.Println(p1)
    switch p2 {
    case "\n\t":
        if (-2147483648 != 1) {
            if (42 == -62) {
                fmt.Println(p1, p1)
                var v8 uint = 2
                v9 := []int{1}
                _ = v8
                _ = v9
            }
        }
    default:
        var v10 float64 = float64(1.5)
        _ = v10
    }
    var v11 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 156
}
func genFn1(p0 int16, p1 uint, p2 int) uint8 {
    fmt.Println(p2)
    if false {
        v12 := GenS1{F0: p0, F1: 65535}
        v13 := []int{127, (^(p2))}
        _ = v12
        _ = v13
    }
    var v14 bool = (p2 == (^(p2)))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 124
}

func main() {
    v1 := []string{(("foo bar" + "hello") + "foo bar")}
    v2 := make(map[int]string)
    var v3 uint32 = 65535
    v3 += (v3 / (((v3 >> ((255) & 63))) | 1))
    switch 4294967295 {
    case 32768:
        var v15 int = -2147483648
        _ = v15
    case 32767:
        var v16 uint16 = 2
        _ = v16
    default:
        v3 -= v3
    }
    fmt.Println(v3)
    if false {
        if ("hello" > "a") {
            if (false && true) {
                delete(v2, 2147483647)
                if false {
                    v3 += 95
                    fmt.Println(v3)
                    v17 := GenS1{F0: int16(v3), F1: 255}
                    _ = v17
                }
            } else {
                fmt.Println(v3)
            }
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 uint8 = 116
                _ = fi18
                _ = v19
            }
            if true {
                fmt.Println(v3)
                v2[32767] = "foo bar"
            } else {
                v3 = v3
                for fi20 := 0; fi20 < 16; fi20++ {
                    var v21 uint = 65535
                    v3 -= 2147483647
                    _ = fi20
                    _ = v21
                }
                _ = len(v1)
            }
        } else {
            if true {
                switch 2147483647 {
                case -2147483648:
                    for fi22 := 0; fi22 < 19; fi22++ {
                        var v23 uint = 4294967295
                        switch (-(fi22)) {
                        case 128:
                            var v24 int = fi22
                            _ = v24
                        case 736:
                            v25 := v2[(fi22 >> ((v23) & 63))]
                            _ = v25
                        }
                        v3 = v3
                        _ = fi22
                        _ = v23
                    }
                case 4294967295:
                    if false {
                        var v26 float64 = -2.25
                        _ = v26
                    } else {
                        var v27 uint32 = v3
                        var v28 float32 = -2.25
                        if (-2.25 < 0.1) {
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                            var v29 uint16 = 1
                            _ = v29
                        } else {
                            var v30 string = "foo bar"
                            var v31 uint16 = 1
                            _ = v30
                            _ = v31
                        }
                        _ = v27
                        _ = v28
                    }
                }
                var v32 int8 = 1
                fmt.Println(v3, v3)
                _ = v32
            } else {
                var v33 uint64 = 823
                switch "foo bar" {
                case "foo bar":
                    v33 = v33
                case "a":
                    if (genFn0("\n\t", 100, ("hello" + "hello")) != 1) {
                        v3 += v3
                    } else {
                        switch (("\n\t" + "foo bar") + "\n\t") {
                        case "":
                            var v34 int32 = 1
                            _ = v34
                        }
                        if false {
                            var v35 uint32 = v3
                            v36 := GenS2{F0: -2147483648}
                            var v37 float64 = -2.25
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                    }
                default:
                    if false {
                        v33 = 948
                        v38 := GenS2{F0: 9223372036854775807}
                        _ = v38
                    }
                }
                _ = v33
            }
            if false {
                v39 := GenS1{F0: -32768, F1: 821}
                var v40 int32 = -974
                switch int(v40) {
                case 127:
                    var v41 uint32 = (v3 >> ((42) & 63))
                    _ = v41
                case -1:
                    v3 = (^(v3))
                }
                _ = v39
                _ = v40
            } else {
                v42 := GenS1{F0: 32767, F1: 188}
                _ = v42
            }
        }
        v43 := GenS3{F0: 65535, F1: 807, F2: 1}
        _ = v43
    }
    v3 -= v3
    v2[32768] = ("\n\t" + "\n\t")
    var v44 uint16 = 812
    v45 := GenS4{F0: 278, F1: 97, F2: "\n\t"}
    v3 = (v3 + (v3 | v3))
    var v46 float64 = 1.5
    var v47 float64 = ((v46 - float64(0.1)) - v46)
    if true {
        v46 = v46
        v48 := GenS5{F0: 114}
        _ = v48
    }
    var v49 int32 = -443
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v49
}
