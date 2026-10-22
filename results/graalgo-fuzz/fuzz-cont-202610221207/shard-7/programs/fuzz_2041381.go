package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int32; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int16; F2 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 uint; F2 bool }
func genFn0(p0 string, p1 int, p2 int32) uint8 {
    v4 := GenS1{F0: -22, F1: -2147483648}
    p0 = "a"
    fmt.Println(p2, p1)
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := []int{p1}
        p1 += -2147483648
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 int32, p1 uint16, p2 uint8) bool {
    var v7 float32 = 1e10
    fmt.Println(p2, p2, v7)
    p0 -= (p0 + 42)
    for fi8 := 0; fi8 < 20; fi8++ {
        v7 += float32(1.5)
        _ = fi8
    }
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return (1 > 42)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 4294967295
    v3 := []string{"a", ("a" + ("foo bar" + "hello")), "hello"}
    v2 = (v2 & 750)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int8 = 0
        var v12 uint64 = (v2 | (v2 | 892))
        _ = fi10
        _ = v11
        _ = v12
    }
    v2 = (^(v2))
    v13 := GenS2{F0: 283, F1: 747, F2: 30}
    var v14 uint16 = 980
    switch 42 {
    case 790:
        switch 0 {
        case -2147483648:
            fmt.Println(v2)
        case -9223372036854775808:
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 int64 = 0
                v17 := GenS1{F0: 1, F1: -2147483648}
                for fi18 := 0; fi18 < 18; fi18++ {
                    v1[42] = (fi15 & (fi18 + fi18))
                    v14 = v14
                    _ = fi18
                }
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    case 742:
        if (false && true) {
            v19 := GenS3{F0: 0, F1: 405}
            v14 = v14
            for fi20 := 0; fi20 < 18; fi20++ {
                v2 += v2
                fmt.Println(v2)
                _ = fi20
            }
            _ = v19
        } else {
            var v21 uint = 255
            _ = v21
        }
    case -1:
        v22 := v1[565]
        _ = v22
    default:
        for fi23 := 0; fi23 < 8; fi23++ {
            v14 -= v14
            var v24 int16 = 518
            if true {
                if true {
                    var v25 int64 = 42
                    v25 -= -463
                    var v26 int64 = int64(v24)
                    _ = v25
                    _ = v26
                } else {
                    var v27 uint16 = (v14 * v14)
                    _ = v27
                }
                fmt.Println(fi23, v24)
                for fi28 := 0; fi28 < 4; fi28++ {
                    switch "a" {
                    case "hello":
                        var v29 int64 = 42
                        _ = v29
                    case "\n\t":
                        if true {
                            var v30 int32 = 928
                            _ = v30
                        }
                    default:
                        var v31 int = (fi23 / ((fi28) | 1))
                        _ = v31
                    }
                    _ = fi28
                }
            }
            _ = fi23
            _ = v24
        }
    }
    fmt.Println(v2)
    var v32 int = 32768
    for fi33 := 0; fi33 < 7; fi33++ {
        v34 := GenS1{F0: 0, F1: -2147483648}
        switch ("a" + "") {
        case "a":
            var v35 uint8 = 42
            _ = v35
        case "foo bar":
            v32 += fi33
        default:
            for fi36 := 0; fi36 < 17; fi36++ {
                var v37 uint = 466
                v38 := GenS4{F0: 246, F1: 33, F2: uint64(v37)}
                v39 := v1[42]
                v37 += (v37 & v37)
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
            }
        }
        if (!true) {
            for fi40 := 0; fi40 < 2; fi40++ {
                if (!true) {
                    var v41 int8 = -1
                    _ = v41
                }
                _ = fi40
            }
        }
        _ = fi33
        _ = v34
    }
    v2 += (v2 + 0)
    switch "a" {
    case "a":
        v42 := GenS5{F0: genFn0(("a" + ""), v32, 917), F1: 0, F2: int32(v2)}
        _ = v42
    default:
        switch "\n\t" {
        case "foo bar":
            v32 -= int(v2)
        case "\n\t":
            fmt.Println(v32, v14)
        case "a":
            for fi43 := 0; fi43 < 17; fi43++ {
                var v44 int32 = -905
                _ = fi43
                _ = v44
            }
        }
    }
    v45 := v1[42]
    if false {
        for fi46 := 0; fi46 < 17; fi46++ {
            var v47 int8 = 0
            if ((0.0 <= float64(float32(1.5))) && false) {
                for fi48 := 0; fi48 < 3; fi48++ {
                    var v49 uint32 = 2147483647
                    _ = fi48
                    _ = v49
                }
            } else {
                if ((!true) || true) {
                    if true {
                        fmt.Println(v14)
                        for fi50 := 0; fi50 < 12; fi50++ {
                            switch "" {
                            case "hello":
                                switch "hello" {
                                case "foo bar":
                                    v51 := GenS4{F0: 42, F1: 127, F2: v2}
                                    _ = v51
                                }
                            case "a":
                                v2 = v2
                            }
                            _ = fi50
                        }
                        v32 -= 128
                    }
                    v32 -= 2147483647
                } else {
                    v14 += (^(v14))
                    for fi52 := 0; fi52 < 3; fi52++ {
                        switch "\n\t" {
                        case "":
                            v47 = 1
                        case "\n\t":
                            v53 := GenS3{F0: 166, F1: 227}
                            _ = v53
                        case "foo bar":
                            v54 := v1[6]
                            delete(v1, 65535)
                            _ = v54
                        }
                        _ = fi52
                    }
                }
                if false {
                    var v55 string = ("foo bar" + "")
                    if (v47 == (v47 * 95)) {
                        v1[0] = int(v47)
                    } else {
                        for fi56 := 0; fi56 < 6; fi56++ {
                            v32 = v45
                            if false {
                                var v57 bool = (1e10 <= 0.0)
                                _ = v57
                            } else {
                                var v58 int16 = 1000
                                _ = v58
                            }
                            v14 = 42
                            v2 -= v2
                            _ = fi56
                        }
                        v14 += v14
                        v2 += 42
                    }
                    _ = len(v1)
                    _ = v55
                } else {
                    for fi59 := 0; fi59 < 7; fi59++ {
                        v32 = 2147483647
                        v14 += 32767
                        _ = fi59
                    }
                    switch (("\n\t" + "\n\t") + (("foo bar" + "foo bar") + "")) {
                    case "a":
                        v60 := GenS6{F0: 32767}
                        _ = v60
                    default:
                        v61 := GenS7{F0: 65535, F1: 409, F2: (false && false)}
                        _ = v61
                    }
                }
            }
            _ = fi46
            _ = v47
        }
    }
    var v62 bool = (false && true)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v45:", v45)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v45
    _ = v62
}
