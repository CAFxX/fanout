package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int8, p1 int32) int32 {
    var v7 bool = false
    var v8 uint16 = 2
    for fi9 := 0; fi9 < 7; fi9++ {
        if v7 {
            for fi10 := 0; fi10 < 8; fi10++ {
                fmt.Println(v8)
                p1 -= p1
                p1 += 2147483647
                _ = fi10
            }
        }
        if true {
            switch fi9 {
            case -463:
                switch fi9 {
                case 127:
                    v8 -= v8
                case 128:
                    if v7 {
                        fmt.Println(v8, p1)
                        for fi11 := 0; fi11 < 19; fi11++ {
                            fmt.Println(p0, fi9, p0)
                            if ((p1 << ((1) & 63)) < (-(p1))) {
                                var v12 string = "a"
                                _ = v12
                            } else {
                                var v13 int8 = (p0 | p0)
                                var v14 uint = 2
                                switch -855 {
                                case 127:
                                    fmt.Println(fi9, v7)
                                default:
                                    v14 = (^(v14))
                                }
                                p1 = -1
                                _ = v13
                                _ = v14
                            }
                            _ = fi11
                        }
                    } else {
                        var v15 int16 = 1
                        v16 := GenS1{F0: 647}
                        v15 += v15
                        _ = v15
                        _ = v16
                    }
                case 1:
                    var v17 int64 = 830
                    _ = v17
                }
            case 128:
                if v7 {
                    p0 -= p0
                    switch "a" {
                    case "foo bar":
                        var v18 uint64 = 18446744073709551615
                        _ = v18
                    case "a":
                        switch int(p0) {
                        case -571:
                            p0 -= 100
                        }
                    case "":
                        fmt.Println(p1, p0, v8)
                    default:
                        v7 = (!true)
                    }
                }
            }
            for fi19 := 0; fi19 < 13; fi19++ {
                if v7 {
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v8, v7)
                    case "foo bar":
                        p1 = -608
                    }
                    if v7 {
                        var v20 int8 = p0
                        _ = v20
                    }
                    v21 := GenS1{F0: -894}
                    _ = v21
                }
                v8 -= v8
                _ = fi19
            }
            v7 = (741 != 0)
        } else {
            switch fi9 {
            case 32768:
                v7 = false
            }
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p1
}
func genFn1() int8 {
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 int64 = -382
        v24 := []int{(-(fi22))}
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: 483}
    _ = v25
    return -55
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 0
    var v3 int = -987
    v4 := make(map[string]bool)
    var v5 int = -84
    var v6 uint8 = 255
    var v26 int16 = -1
    if false {
        v3 -= v5
        v5 -= v3
    }
    if true {
        var v27 int8 = int8(v2)
        _ = v27
    } else {
        v28 := GenS1{F0: (v2 / ((v2) | 1))}
        v3 = (v5 / ((v3) | 1))
        v2 += v2
        _ = v28
    }
    _ = len(v1)
    for fi29 := 0; fi29 < 8; fi29++ {
        var v30 string = ("a" + "foo bar")
        v31 := GenS1{F0: (v2 | v2)}
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 int32 = 1
    v33 := []int{v3, (v3 << ((0) & 63)), -591}
    for fi34 := 0; fi34 < 13; fi34++ {
        v35 := GenS2{F0: (-(v32))}
        for fi36 := 0; fi36 < 7; fi36++ {
            var v37 string = "a"
            _ = fi36
            _ = v37
        }
        _ = fi34
        _ = v35
    }
    if (genFn1() > -103) {
        if false {
            v3 = v3
        } else {
            switch "foo bar" {
            case "a":
                for fi38 := 0; fi38 < 8; fi38++ {
                    var v39 bool = (42 <= int8(fi38))
                    var v40 bool = (!(42 <= (-(v26))))
                    var v41 string = "\n\t"
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                }
            case "":
                fmt.Println(v6, v6, v32)
            case "hello":
                v42 := GenS1{F0: int64(v6)}
                _ = v42
            }
        }
        delete(v4, "hello")
        var v43 uint8 = 98
        _ = v43
    } else {
        v44 := GenS1{F0: v2}
        switch v5 {
        case -1:
            var v45 string = ""
            _ = v45
        case 2147483647:
            var v46 float32 = 1e10
            _ = v46
        case 1:
            v26 -= 42
        }
        v3 = int(v2)
        _ = v44
    }
    var v47 float64 = float64(1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
