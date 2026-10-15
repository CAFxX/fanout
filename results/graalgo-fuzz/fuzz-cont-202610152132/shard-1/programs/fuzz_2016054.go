package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 int, p1 int16) int {
    switch p0 {
    case 9223372036854775807:
        fmt.Println(p1)
    case 128:
        switch (p0 % (((p0 & 42)) | 1)) {
        case 128:
            v6 := GenS1{F0: -716, F1: 2, F2: 42}
            _ = v6
        case 9223372036854775807:
            var v7 int = p0
            _ = v7
        }
    default:
        p1 += (p1 + (p1 | (p1 - 0)))
    }
    p1 -= p1
    p0 = p0
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v8
    return 32768
}
func genFn1(p0 int, p1 uint, p2 int64) uint {
    var v9 int64 = (p2 % ((0) | 1))
    var v10 int32 = 42
    if (0.1 > (0.1 + (0.0 - (float64((1e10 + 1.5)) * 0.1)))) {
        var v11 int = (^(p0))
        for fi12 := 0; fi12 < 19; fi12++ {
            fmt.Println(p2, v10, p1)
            p1 -= (p1 & p1)
            _ = fi12
        }
        _ = v11
    }
    v13 := []int{9223372036854775807, 598, genFn0(p0, 1000)}
    var v14 bool = (!true)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v13
    _ = v14
    _ = v9
    return p1
}

func main() {
    var v1 int = 667
    var v2 string = ("a" + "a")
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    var v5 float32 = -2.25
    if true {
        if (!true) {
            v15 := []int{128, (-(v1))}
            fmt.Println(v2, v1)
            _ = v15
        }
    }
    if false {
        v16 := GenS2{F0: 1}
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := GenS3{F0: genFn0(fi17, 1000), F1: -1}
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    fmt.Println(v2, v5)
    v5 = v5
    if true {
        var v19 int32 = -939
        if (((2 > 2) || (1.5 <= (1e10 - (float32((0.1 / 1.0)) + (v5 - v5))))) || ((-(v1)) <= v1)) {
            v20 := GenS1{F0: 350, F1: 65535, F2: 6}
            v1 = v1
            _ = v20
        } else {
            v5 = v5
        }
        _ = v19
    } else {
        var v21 uint = 255
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                v21 = v21
                _ = fi23
            }
            if false {
                var v24 int32 = 410
                var v25 bool = true
                _ = v24
                _ = v25
            } else {
                if false {
                    if true {
                        v1 = (v1 % ((v1) | 1))
                    } else {
                        v2 = v2
                    }
                    v26 := GenS3{F0: 4294967295, F1: 100}
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 12; fi27++ {
                        if false {
                            if (!(3.14159 > (-2.25 / 1e10))) {
                                var v28 uint16 = 2
                                _ = v28
                            }
                            if (!(-1 >= 42)) {
                                var v29 float32 = 3.14159
                                _ = v29
                            }
                            switch v1 {
                            case 1:
                                fmt.Println(fi22, v21)
                            }
                        } else {
                            if true {
                                v1 -= 127
                                switch v2 {
                                case "foo bar":
                                    if (((!(42 > 255)) && (!(true && false))) && false) {
                                        v2 = v2
                                        var v30 int = v1
                                        var v31 float64 = (1.5 + 1e10)
                                        v30 += genFn0(genFn0((v1 & genFn0(fi27, -32768)), -32768), 750)
                                        _ = v30
                                        _ = v31
                                    }
                                case "hello":
                                    var v32 int32 = 1
                                    _ = v32
                                default:
                                    var v33 int = 127
                                    _ = v33
                                }
                            }
                            v34 := []int{fi22, 32767, v1}
                            fmt.Println(fi22)
                            if true {
                                var v35 uint32 = 2147483647
                                _ = v35
                            } else {
                                switch (v2 + v2) {
                                case "a":
                                    var v36 int16 = 1
                                    _ = v36
                                default:
                                    v37 := GenS3{F0: (fi22 / ((fi27) | 1)), F1: 100}
                                    _ = v37
                                }
                                v5 = v5
                                v38 := GenS4{F0: genFn1(genFn0(genFn0(genFn0((-(fi22)), 1), 42), 1), v21, -373)}
                                _ = v38
                            }
                            _ = v34
                        }
                        v39 := GenS5{F0: 42}
                        _ = fi27
                        _ = v39
                    }
                    v40 := GenS6{F0: 9223372036854775807, F1: -128, F2: 0}
                    _ = v40
                }
                fmt.Println(v1)
            }
            v1 -= 2147483647
            _ = fi22
        }
        v41 := []int{int(v21), v1, v1, (v1 ^ (v1 & v1))}
        _ = v21
        _ = v41
    }
    v2 = "\n\t"
    for fi42 := 0; fi42 < 13; fi42++ {
        var v43 uint8 = 91
        switch 127 {
        case 634:
            v44 := v3[4294967295]
            _ = v44
        case -1:
            v43 = v43
        }
        _ = fi42
        _ = v43
    }
    fmt.Println(v2)
    fmt.Println(v5)
    v5 += v5
    fmt.Println(v1, v2)
    if true {
        switch v1 {
        case 9223372036854775807:
            v45 := GenS6{F0: 9223372036854775807, F1: 42, F2: 2}
            _ = v45
        case 2147483647:
            v4["\n\t"] = 65535
        case 128:
            delete(v3, 0)
        }
        v46 := GenS7{F0: 4294967295}
        v1 -= 0
        _ = v46
    }
    v5 = 0.0
    v5 -= float32(0.1)
    delete(v4, (v2 + ("" + ("hello" + (("foo bar" + "") + v2)))))
    if (true && (false && (1.5 != (-2.25 + 0.1)))) {
        fmt.Println(v1, v1, v2)
        v1 = v1
    }
    var v47 uint = 497
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
