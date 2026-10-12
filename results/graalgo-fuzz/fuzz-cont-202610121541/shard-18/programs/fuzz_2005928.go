package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 string; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int; F1 int32; F2 int64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint, p1 uint8, p2 int16) int {
    v6 := []int{}
    var v7 string = ""
    if true {
        switch -874 {
        case -68:
            if false {
                switch "hello" {
                case "foo bar":
                    v7 = v7
                case "":
                    p2 += p2
                default:
                    for fi8 := 0; fi8 < 7; fi8++ {
                        var v9 uint8 = (p1 >> ((42) & 63))
                        var v10 uint64 = 18446744073709551615
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                }
                p1 -= p1
                if false {
                    if (!(v7 > "a")) {
                        p2 = p2
                        p2 += p2
                        v11 := GenS1{F0: uint64(p1), F1: 42}
                        _ = v11
                    } else {
                        var v12 int32 = 42
                        v13 := GenS1{F0: 2, F1: -2147483648}
                        _ = v12
                        _ = v13
                    }
                } else {
                    switch 4294967295 {
                    case 127:
                        fmt.Println(p0, p0, v7)
                    case 2147483647:
                        var v14 uint32 = 42
                        _ = v14
                    case -1:
                        for fi15 := 0; fi15 < 7; fi15++ {
                            if false {
                                fmt.Println(p1, fi15, v7)
                            } else {
                                p0 = p0
                                if len(v6) > 0 {
                                    v6[0] = (fi15 | (^(fi15)))
                                }
                            }
                            v16 := GenS1{F0: 304, F1: -300}
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = fi15
                            _ = v16
                        }
                    default:
                        v7 = "a"
                    }
                    fmt.Println(p1, p0)
                }
            }
        case 1:
            v17 := GenS1{F0: 135, F1: 0}
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 int = -799
    var v2 int = v1
    v3 := make(map[int]bool)
    var v4 uint8 = 42
    v5 := []int{(v1 >> ((1) & 63)), v1, -250, 32767}
    for fi18 := 0; fi18 < 1; fi18++ {
        if len(v5) > 0 {
            v5[0] = v1
        }
        if false {
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 uint64 = 43
                _ = fi19
                _ = v20
            }
            var v21 int32 = -11
            _ = v21
        } else {
            for fi22 := 0; fi22 < 13; fi22++ {
                v2 = (^(fi18))
                _ = fi22
            }
            var v23 int64 = 553
            _ = v23
        }
        _ = fi18
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        if (42 >= uint64(v4)) {
            for fi25 := 0; fi25 < 19; fi25++ {
                v1 = (-(fi25))
                v26 := GenS1{F0: 2, F1: -842}
                fmt.Println(v2)
                _ = fi25
                _ = v26
            }
            if true {
                switch (v1 << ((1) & 63)) {
                case -1:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case 560:
                    v3[v2] = false
                case 32768:
                    v2 -= (v2 >> ((0) & 63))
                default:
                    switch "hello" {
                    case "\n\t":
                        v27 := GenS2{F0: "\n\t", F1: 1, F2: 1}
                        _ = v27
                    case "a":
                        for fi28 := 0; fi28 < 19; fi28++ {
                            var v29 string = "hello"
                            var v30 int64 = 54
                            v31 := GenS3{F0: 258, F1: (v30 ^ (v30 / ((v30) | 1)))}
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    default:
                        fmt.Println(v4, v2)
                    }
                }
            }
        }
        switch "\n\t" {
        case "":
            var v32 int = 1
            _ = v32
        case "a":
            var v33 int32 = 0
            _ = v33
        case "foo bar":
            for fi34 := 0; fi34 < 1; fi34++ {
                v2 = v2
                switch genFn0(2, v4, 42) {
                case 1:
                    v35 := GenS4{F0: fi34, F1: 990, F2: -763}
                    _ = v35
                case -9223372036854775808:
                    fmt.Println(v2, v1)
                default:
                    var v36 int = v2
                    _ = v36
                }
                var v37 int8 = 100
                _ = fi34
                _ = v37
            }
        }
        _ = fi24
    }
    delete(v3, v1)
    switch genFn0(4294967295, (^(v4)), -32768) {
    case 127:
        var v38 int64 = 699
        _ = v38
    }
    if false {
        for fi39 := 0; fi39 < 13; fi39++ {
            switch (("\n\t" + ("hello" + "hello")) + "foo bar") {
            case "\n\t":
                _ = len(v3)
            }
            v4 += (v4 * v4)
            switch "foo bar" {
            case "foo bar":
                v4 = (v4 - v4)
            default:
                var v40 float64 = (0.1 - 0.1)
                _ = v40
            }
            _ = fi39
        }
    } else {
        for fi41 := 0; fi41 < 2; fi41++ {
            v42 := GenS5{F0: v4}
            v2 -= (v1 % ((genFn0(42, 42, -326)) | 1))
            _ = fi41
            _ = v42
        }
        v43 := v3[v2]
        _ = v43
    }
    fmt.Println(v2, v2, v4)
    for fi44 := 0; fi44 < 18; fi44++ {
        v45 := v3[fi44]
        _ = fi44
        _ = v45
    }
    fmt.Println(v2, v1)
    var v46 int64 = 0
    v47 := GenS1{F0: 42, F1: 2147483647}
    for fi48 := 0; fi48 < 10; fi48++ {
        var v49 bool = false
        v46 = -423
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
        _ = fi48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
