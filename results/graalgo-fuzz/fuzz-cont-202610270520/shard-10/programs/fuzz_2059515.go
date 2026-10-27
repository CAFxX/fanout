package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int64, p1 bool, p2 int32) int64 {
    v6 := GenS1{F0: p0}
    switch 127 {
    case -2147483648:
        p2 = 42
    case -9223372036854775808:
        if p1 {
            switch 127 {
            case 0:
                p0 = -1
            case 42:
                p2 -= (p2 % ((-387) | 1))
            case 32767:
                var v7 uint32 = 2147483647
                _ = v7
            }
            if true {
                fmt.Println(p1, p2, p0)
                p2 = (p2 >> ((65535) & 63))
                switch 183 {
                case -205:
                    p2 -= p2
                case 9223372036854775807:
                    v8 := GenS2{F0: 769}
                    _ = v8
                case 1:
                    var v9 string = ("" + "")
                    _ = v9
                default:
                    v10 := GenS1{F0: p0}
                    _ = v10
                }
            } else {
                p1 = false
            }
            for fi11 := 0; fi11 < 3; fi11++ {
                p2 = p2
                var v12 int = 32767
                v13 := []int{}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    }
    v14 := []int{1}
    p2 += (p2 + (p2 >> ((2) & 63)))
    switch 128 {
    case 457:
        v15 := GenS3{F0: -70, F1: 286}
        _ = v15
    case 127:
        v14 = append(v14, 2147483647)
    case -9223372036854775808:
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int16 = -32768
            _ = fi16
            _ = v17
        }
    default:
        var v18 uint16 = 32767
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v6
    return p0
}
func genFn1(p0 uint64, p1 int) uint16 {
    for fi19 := 0; fi19 < 19; fi19++ {
        p1 = 2147483647
        v20 := []int{}
        _ = fi19
        _ = v20
    }
    fmt.Println(p1, p1)
    fmt.Println(p0, p1)
    if ((-(p1)) < p1) {
        p1 += (p1 % ((0) | 1))
    } else {
        for fi21 := 0; fi21 < 18; fi21++ {
            p1 -= (p1 ^ (p1 + p1))
            p1 += fi21
            for fi22 := 0; fi22 < 11; fi22++ {
                p1 -= 42
                fmt.Println(fi22, p0)
                var v23 int16 = 32767
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]int)
    var v3 string = ""
    v4 := []int{1}
    v5 := []uint32{314, 577}
    v24 := GenS3{F0: 0, F1: 42}
    if true {
        switch -9223372036854775808 {
        case 9223372036854775807:
            v1 += 3.14159
        }
        var v25 int32 = -240
        if true {
            v4 = append(v4, -2147483648)
        } else {
            if (127 >= 127) {
                v26 := v2[728]
                _ = v26
            }
            v3 = v3
        }
        _ = v25
    } else {
        v1 = (v1 + (3.14159 - (v1 / v1)))
        v3 = v3
        v1 = (1e10 * 1e10)
    }
    v3 = "\n\t"
    _ = len(v2)
    switch "foo bar" {
    case "foo bar":
        for fi27 := 0; fi27 < 13; fi27++ {
            v2[433] = fi27
            _ = fi27
        }
    case "":
        var v28 float32 = 3.14159
        _ = v28
    case "a":
        if ((65535 < 5) && ((!true) || true)) {
            switch (("" + v3) + (v3 + (("a" + "foo bar") + (((v3 + ("hello" + (v3 + v3))) + "a") + v3)))) {
            case "a":
                if false {
                    if false {
                        v1 += v1
                        fmt.Println(v3, v1)
                    }
                    v3 = (v3 + ((((v3 + (v3 + "a")) + v3) + v3) + ((v3 + ((v3 + v3) + "")) + v3)))
                } else {
                    v3 = (v3 + "")
                    v29 := GenS4{F0: 42, F1: 127}
                    v30 := v2[251]
                    _ = v29
                    _ = v30
                }
            case "\n\t":
                v1 = v1
            case "":
                v31 := GenS5{F0: 127, F1: (v3 + v3)}
                _ = v31
            default:
                v3 = "hello"
            }
            for fi32 := 0; fi32 < 13; fi32++ {
                var v33 uint8 = 42
                for fi34 := 0; fi34 < 14; fi34++ {
                    switch (fi32 + (^(fi34))) {
                    case 9223372036854775807:
                        switch "\n\t" {
                        case "hello":
                            v2[812] = (fi32 | 9)
                        case "foo bar":
                            var v35 uint8 = 0
                            _ = v35
                        case "\n\t":
                            switch (fi34 * fi32) {
                            case -1:
                                v36 := GenS4{F0: 4294967295, F1: 42}
                                _ = v36
                            case -479:
                                fmt.Println(v1, fi34, v1)
                            }
                            if true {
                                var v37 uint8 = 128
                                _ = len(v2)
                                if (!false) {
                                    v38 := GenS6{F0: (v33 - (v37 >> ((226) & 63))), F1: 1}
                                    _ = v38
                                }
                                _ = v37
                            }
                        default:
                            switch "hello" {
                            case "":
                                v39 := GenS3{F0: 0, F1: 65535}
                                _ = v39
                            default:
                                v3 = "a"
                            }
                        }
                    case -2147483648:
                        var v40 int8 = 127
                        _ = v40
                    case 127:
                        var v41 int32 = -668
                        _ = v41
                    }
                    _ = fi34
                }
                _ = fi32
                _ = v33
            }
        }
    }
    var v42 uint16 = 244
    var v43 bool = true
    delete(v2, 9)
    var v44 uint16 = v42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
