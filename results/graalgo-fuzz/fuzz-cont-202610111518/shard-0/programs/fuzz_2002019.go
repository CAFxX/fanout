package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS4 struct { F0 int; F1 uint64; F2 int }
type GenS5 struct { F0 int16; F1 int; F2 int16 }
type GenS6 struct { F0 bool; F1 uint; F2 uint8 }
func genFn0(p0 int32) uint16 {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int32 = p0
        if true {
            v9 := []int{(fi7 / ((fi7) | 1)), fi7, 94, fi7}
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := GenS1{F0: (1 < 1)}
                switch fi10 {
                case 42:
                    v8 += (p0 * (p0 >> ((676) & 63)))
                case 32767:
                    if (!(-14 > -1)) {
                        if ((v8 ^ p0) < v8) {
                            if true {
                                v12 := GenS1{F0: (-128 >= -50)}
                                v8 = v8
                                var v13 float32 = float32(1.5)
                                _ = v12
                                _ = v13
                            } else {
                                v14 := GenS2{F0: (!(float32(float64(((0.1 * 1.5) + float32(3.14159)))) != -2.25))}
                                if true {
                                    fmt.Println(fi10, fi7, p0)
                                    v8 = v8
                                }
                                _ = v14
                            }
                            v8 += (^(v8))
                            _ = len(v9)
                            v8 -= -486
                        }
                    }
                case -889:
                    v15 := GenS3{F0: 2, F1: 27, F2: -128}
                    _ = v15
                default:
                    if true {
                        var v16 uint = 1
                        _ = v16
                    }
                }
                var v17 int8 = -128
                _ = fi10
                _ = v11
                _ = v17
            }
            if (42 != 108) {
                fmt.Println(fi7, v8, v8)
            }
            _ = v9
        } else {
            var v18 uint = 0
            v19 := GenS2{F0: ((1 < -1) && (371 != 4294967295))}
            for fi20 := 0; fi20 < 12; fi20++ {
                fmt.Println(v18)
                _ = fi20
            }
            _ = v18
            _ = v19
        }
        if (!(-1 <= 178)) {
            p0 += (v8 ^ p0)
            v21 := GenS3{F0: 2, F1: 255, F2: -128}
            _ = v21
        }
        _ = fi7
        _ = v8
    }
    if false {
        if (false && true) {
            v22 := GenS4{F0: 32768, F1: 1, F2: 4294967295}
            v23 := []int{-89, -2147483648, 0, 870}
            var v24 bool = false
            _ = v22
            _ = v23
            _ = v24
        } else {
            if (p0 != (p0 / ((p0) | 1))) {
                if false {
                    for fi25 := 0; fi25 < 3; fi25++ {
                        v26 := []int{(fi25 - (fi25 - int(p0))), fi25}
                        fmt.Println(fi25)
                        for fi27 := 0; fi27 < 8; fi27++ {
                            p0 += (p0 + (p0 ^ -1))
                            var v28 float64 = (-2.25 / 0.1)
                            if false {
                                fmt.Println(v28)
                            } else {
                                if (false && true) {
                                    v26 = append(v26, fi25)
                                }
                            }
                            _ = fi27
                            _ = v28
                        }
                        _ = fi25
                        _ = v26
                    }
                } else {
                    fmt.Println(p0)
                    p0 -= p0
                }
            } else {
                p0 = (-(p0))
            }
            v29 := GenS5{F0: -797, F1: 32767, F2: 0}
            _ = v29
        }
        var v30 int8 = 127
        var v31 string = ("\n\t" + ("foo bar" + ("a" + "foo bar")))
        _ = v30
        _ = v31
    }
    _ = p0
    return 113
}
func genFn1() int64 {
    var v32 int64 = 42
    var v33 string = ""
    _ = v32
    _ = v33
    return (v32 >> ((716) & 63))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"", "a", "a"}
    var v3 uint64 = 42
    var v4 int = 2147483647
    var v5 bool = false
    v6 := make(map[uint32]bool)
    switch "foo bar" {
    case "a":
        v2 = append(v2, "hello")
    case "hello":
        var v34 float64 = 1e10
        _ = v34
    }
    var v35 uint16 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        v36 := GenS4{F0: -329, F1: 1, F2: v4}
        _ = v36
    }
    if v5 {
        for fi37 := 0; fi37 < 4; fi37++ {
            v3 = (^(v3))
            _ = fi37
        }
    } else {
        v38 := GenS6{F0: false, F1: 255, F2: 128}
        var v39 int = int(v3)
        _ = len(v2)
        _ = v38
        _ = v39
    }
    if v5 {
        var v40 uint8 = 2
        v41 := GenS5{F0: 32767, F1: 370, F2: -882}
        _ = v40
        _ = v41
    } else {
        _ = len(v6)
        v5 = (!v5)
    }
    for fi42 := 0; fi42 < 4; fi42++ {
        var v43 uint16 = 42
        v44 := GenS5{F0: -150, F1: (v4 << ((852) & 63)), F2: 1000}
        v1[""] = 2147483647
        _ = fi42
        _ = v43
        _ = v44
    }
    var v45 int16 = 0
    if v5 {
        v2 = append(v2, "a")
    }
    v5 = v5
    v35 = (v35 / ((v35) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
