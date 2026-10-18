package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 bool }
func genFn0() int {
    v6 := []int{-804, 4294967295, 0, 483}
    var v7 bool = true
    v7 = (255 != 202)
    var v8 int16 = 0
    _ = v6
    _ = v7
    _ = v8
    return int(v8)
}
func genFn1(p0 uint32, p1 uint, p2 int16) int64 {
    var v9 float64 = 3.14159
    if (!(false || (!false))) {
        var v10 int16 = (p2 | (-(p2)))
        var v11 int16 = 0
        _ = v10
        _ = v11
    }
    if true {
        p0 = p0
        p2 = p2
        switch 32768 {
        case 9223372036854775807:
            v12 := []int{}
            _ = v12
        case 42:
            fmt.Println(p1)
        }
    } else {
        var v13 int = -162
        p0 += p0
        _ = v13
    }
    if false {
        switch -401 {
        case 1:
            if (!true) {
                p1 -= p1
                for fi14 := 0; fi14 < 11; fi14++ {
                    var v15 uint8 = 42
                    p0 += uint32(v15)
                    var v16 uint8 = uint8(p0)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            } else {
                var v17 uint8 = 2
                v9 = (v9 - (float64(-2.25) * v9))
                _ = v17
            }
        case 32768:
            var v18 uint64 = 947
            _ = v18
        case 9223372036854775807:
            var v19 int32 = 2147483647
            _ = v19
        default:
            for fi20 := 0; fi20 < 11; fi20++ {
                for fi21 := 0; fi21 < 13; fi21++ {
                    switch "\n\t" {
                    case "foo bar":
                        var v22 float64 = (0.1 - v9)
                        var v23 float32 = float32((((((v22 - v22) - 1e10) / v9) + v9) - v9))
                        _ = v22
                        _ = v23
                    default:
                        fmt.Println(p2, v9)
                    }
                    _ = fi21
                }
                _ = fi20
            }
        }
    } else {
        p0 += p0
    }
    p0 -= 366
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 int8 = 42
    var v2 uint = 42
    var v3 uint16 = 32767
    v4 := make(map[uint32]uint32)
    var v5 string = "foo bar"
    switch 683 {
    case -1:
        if true {
            fmt.Println(v3, v5)
        }
    case -9223372036854775808:
        for fi24 := 0; fi24 < 13; fi24++ {
            for fi25 := 0; fi25 < 17; fi25++ {
                if (1 > 0) {
                    if (v2 > uint(v3)) {
                        v26 := GenS1{F0: (v3 >> ((704) & 63))}
                        var v27 int = (^(fi25))
                        _ = v26
                        _ = v27
                    }
                    var v28 int32 = 0
                    _ = v28
                } else {
                    v29 := GenS2{F0: (v3 << ((676) & 63)), F1: true}
                    var v30 int8 = 42
                    if ((-1 == 0) || ((v2 >> ((v2) & 63)) <= (v2 * (v2 + v2)))) {
                        fmt.Println(v30)
                        switch v5 {
                        case "a":
                            v5 = "\n\t"
                        }
                    }
                    var v31 uint = v2
                    v31 += 103
                    _ = v29
                    _ = v30
                    _ = v31
                }
                var v32 string = "\n\t"
                if (true && true) {
                    var v33 int64 = -1
                    v32 = ("foo bar" + "a")
                    v5 = ("a" + v32)
                    _ = v33
                }
                _ = fi25
                _ = v32
            }
            _ = fi24
        }
    case -992:
        delete(v4, 65535)
    }
    v2 = 255
    for fi34 := 0; fi34 < 5; fi34++ {
        var v35 bool = (987 != 876)
        switch genFn0() {
        case 0:
            v36 := GenS2{F0: 1, F1: (v35 && v35)}
            _ = v36
        case 2147483647:
            var v37 string = (((v5 + (v5 + v5)) + (v5 + (v5 + ((("a" + "\n\t") + v5) + "\n\t")))) + "hello")
            _ = v37
        case 32768:
            var v38 float32 = (((0.1 - (-2.25 - 0.1)) * 3.14159) / float32(3.14159))
            _ = v38
        }
        switch ("a" + "hello") {
        case "a":
            v5 = (v5 + ("foo bar" + (v5 + "hello")))
        case "\n\t":
            v1 = (v1 << ((v2) & 63))
        default:
            v35 = (0 < 42)
        }
        _ = fi34
        _ = v35
    }
    v5 = v5
    if ((v2 % ((2) | 1)) >= v2) {
        if true {
            v5 = v5
        }
        var v39 int = genFn0()
        var v40 int8 = (v1 << ((18446744073709551615) & 63))
        _ = v39
        _ = v40
    } else {
        switch 0 {
        case 128:
            switch 32768 {
            case 42:
                fmt.Println(v3)
            case 148:
                v41 := GenS2{F0: (v3 / ((v3) | 1)), F1: (!false)}
                _ = v41
            default:
                v5 = "hello"
            }
        case -1:
            fmt.Println(v1)
        default:
            v5 = v5
        }
        switch (v5 + (("foo bar" + "") + v5)) {
        case "":
            for fi42 := 0; fi42 < 16; fi42++ {
                v3 += v3
                v43 := []int{fi42, fi42, genFn0(), genFn0()}
                v44 := GenS1{F0: v3}
                _ = fi42
                _ = v43
                _ = v44
            }
        case "\n\t":
            v5 = "\n\t"
        }
        v3 = v3
    }
    var v45 int = 804
    for fi46 := 0; fi46 < 12; fi46++ {
        v2 += (v2 - 727)
        var v47 uint16 = (v3 + (v3 + (^(v3))))
        _ = fi46
        _ = v47
    }
    if true {
        v48 := GenS1{F0: 65535}
        v1 = 100
        v2 += 18446744073709551615
        _ = v48
    } else {
        var v49 uint16 = 255
        var v50 uint16 = v49
        for fi51 := 0; fi51 < 15; fi51++ {
            if (v50 < v49) {
                if (fi51 == fi51) {
                    switch ((v5 + "foo bar") + (v5 + "hello")) {
                    case "foo bar":
                        v52 := []int{}
                        _ = v52
                    }
                    for fi53 := 0; fi53 < 16; fi53++ {
                        var v54 uint8 = 127
                        _ = fi53
                        _ = v54
                    }
                    var v55 uint32 = 2
                    _ = v55
                }
                v56 := []int{fi51}
                switch "a" {
                case "hello":
                    _ = len(v56)
                case "":
                    var v57 float32 = (3.14159 / 1.0)
                    _ = v57
                case "\n\t":
                    if len(v56) >= 2 {
                        v56 = v56[1:len(v56)]
                    }
                default:
                    v3 -= 1
                }
                _ = v56
            }
            _ = fi51
        }
        _ = v49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
