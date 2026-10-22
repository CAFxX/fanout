package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int; F1 int16; F2 int64 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS6 struct { F0 string; F1 uint16; F2 uint16 }
func genFn0(p0 bool, p1 bool) bool {
    p0 = p1
    v5 := GenS1{F0: -185, F1: ("hello" + (("hello" + "\n\t") + "a"))}
    if (false || false) {
        fmt.Println(p0, p1)
        if false {
            v6 := GenS1{F0: 4294967295, F1: "hello"}
            _ = v6
        }
    }
    var v7 int32 = 0
    var v8 uint32 = 113
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    _ = v8
    return p1
}
func genFn1(p0 int64, p1 int16, p2 int) int {
    if true {
        p0 -= p0
        v9 := []int{127, 42}
        _ = v9
    } else {
        p0 = int64(p2)
        v10 := []int{}
        _ = v10
    }
    var v11 uint32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p2
}

func main() {
    var v1 int32 = 311
    var v2 string = "hello"
    var v3 int64 = 9223372036854775807
    var v4 int = -1
    switch ("" + "") {
    case "hello":
        for fi12 := 0; fi12 < 10; fi12++ {
            v2 = v2
            v13 := GenS1{F0: (fi12 & -9223372036854775808), F1: v2}
            _ = fi12
            _ = v13
        }
    case "\n\t":
        fmt.Println(v2, v4)
    default:
        switch ((v2 + "hello") + ("a" + v2)) {
        case "":
            v14 := GenS2{F0: (v4 * v4), F1: 42, F2: -991}
            _ = v14
        case "foo bar":
            if (true && false) {
                var v15 float32 = (3.14159 / 1.5)
                v16 := []int{v4, 32768, (v4 % (((v4 / ((v4) | 1))) | 1))}
                _ = v15
                _ = v16
            } else {
                if (v3 == (v3 >> ((42) & 63))) {
                    switch v2 {
                    case "\n\t":
                        var v17 uint = 2
                        _ = v17
                    case "":
                        for fi18 := 0; fi18 < 4; fi18++ {
                            var v19 int = (fi18 - genFn1((v3 >> ((460) & 63)), -1, v4))
                            _ = fi18
                            _ = v19
                        }
                    case "a":
                        v20 := GenS3{F0: 1, F1: (^(v1))}
                        _ = v20
                    default:
                        if false {
                            var v21 int16 = 1
                            v22 := GenS4{F0: v2}
                            for fi23 := 0; fi23 < 16; fi23++ {
                                if true {
                                    if false {
                                        v4 += genFn1(439, v21, fi23)
                                        v24 := GenS4{F0: ""}
                                        fmt.Println(fi23, fi23)
                                        _ = v24
                                    }
                                } else {
                                    var v25 int32 = (v1 | (v1 - v1))
                                    _ = v25
                                }
                                if (((true && (608 > 77)) || (255 >= 65535)) || (-2.25 != 3.14159)) {
                                    v26 := []int{(v4 * 42), 0, genFn1(1, v21, v4)}
                                    var v27 int64 = v3
                                    for fi28 := 0; fi28 < 15; fi28++ {
                                        v26 = append(v26, v4)
                                        _ = fi28
                                    }
                                    _ = v26
                                    _ = v27
                                } else {
                                    if (1 > 255) {
                                        var v29 uint64 = 42
                                        _ = v29
                                    }
                                    v30 := GenS5{F0: 255, F1: 747, F2: 0}
                                    var v31 float64 = (0.1 - (-2.25 * (((0.1 / float64(3.14159)) / float64(0.1)) * (((0.1 / (((0.0 / 1.5) / float64(float32(-2.25))) - (0.0 - (((-2.25 * (1.5 / float64(1e10))) / 3.14159) / 1.5)))) + 1e10) + 0.1))))
                                    _ = v30
                                    _ = v31
                                }
                                switch v2 {
                                case "foo bar":
                                    v2 = v2
                                case "":
                                    v32 := []int{(fi23 | (fi23 * (fi23 * fi23))), -871}
                                    _ = v32
                                default:
                                    for fi33 := 0; fi33 < 10; fi33++ {
                                        v34 := []int{genFn1(v3, (v21 + (v21 - 1)), 602), v4}
                                        v35 := GenS3{F0: 65535, F1: v1}
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                    }
                                }
                                _ = fi23
                            }
                            _ = v21
                            _ = v22
                        }
                    }
                    v4 = 2147483647
                    for fi36 := 0; fi36 < 13; fi36++ {
                        fmt.Println(fi36, fi36)
                        _ = fi36
                    }
                } else {
                    v37 := GenS2{F0: v4, F1: 1, F2: (v3 ^ v3)}
                    switch (("hello" + (v2 + v2)) + (v2 + (("\n\t" + v2) + (v2 + v2)))) {
                    case "a":
                        var v38 uint16 = 1
                        _ = v38
                    case "foo bar":
                        fmt.Println(v4, v3)
                    case "hello":
                        for fi39 := 0; fi39 < 8; fi39++ {
                            if (2 > 811) {
                                var v40 int8 = -128
                                fmt.Println(fi39, v4, v40)
                                _ = v40
                            } else {
                                v41 := []int{v4}
                                _ = v41
                            }
                            var v42 uint8 = 189
                            _ = fi39
                            _ = v42
                        }
                    }
                    _ = v37
                }
                if true {
                    v43 := []int{v4}
                    v43 = append(v43, genFn1((v3 << ((255) & 63)), 809, int(v1)))
                    _ = v43
                } else {
                    var v44 uint32 = 796
                    if true {
                        var v45 string = "hello"
                        _ = v45
                    }
                    _ = v44
                }
            }
        default:
            v46 := GenS5{F0: 42, F1: 262, F2: 2147483648}
            _ = v46
        }
    }
    if (false && (false || (!(false || false)))) {
        v47 := GenS3{F0: 2, F1: v1}
        _ = v47
    } else {
        if true {
            v48 := GenS1{F0: v4, F1: v2}
            for fi49 := 0; fi49 < 5; fi49++ {
                if true {
                    v50 := []int{(fi49 | (^(fi49))), fi49, 1, (fi49 >> ((4294967295) & 63))}
                    if len(v50) > 0 {
                        v50[0] = fi49
                    }
                    _ = v50
                } else {
                    fmt.Println(v2, v4, v4)
                    var v51 uint32 = 4294967295
                    _ = v51
                }
                _ = fi49
            }
            v3 -= v3
            _ = v48
        }
        v52 := GenS6{F0: "", F1: 42, F2: 2}
        _ = v52
    }
    var v53 float32 = (1.5 / 1e10)
    var v54 int16 = -32768
    v55 := GenS3{F0: 191, F1: v1}
    switch "a" {
    case "":
        fmt.Println(v4, v4, v1)
    case "\n\t":
        v56 := []int{(v4 % ((728) | 1))}
        _ = v56
    case "foo bar":
        var v57 string = v2
        _ = v57
    default:
        v2 = (((v2 + v2) + v2) + "\n\t")
    }
    var v58 int = (v4 & genFn1((v3 | (v3 - -1)), v54, v4))
    v53 -= v53
    v59 := []int{127, v4, -865}
    v58 = (v58 + (v58 << ((574) & 63)))
    var v60 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v55
    _ = v58
    _ = v59
    _ = v60
}
