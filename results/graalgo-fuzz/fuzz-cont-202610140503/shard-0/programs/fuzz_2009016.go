package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint }
type GenS2 struct { F0 string; F1 int16; F2 int64 }
type GenS3 struct { F0 uint64; F1 int32; F2 bool }
func genFn0(p0 int8, p1 string) uint16 {
    v7 := []int{0, -999, 127}
    if true {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 uint = 622
            for fi10 := 0; fi10 < 17; fi10++ {
                fmt.Println(p1, fi8, v9)
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
        fmt.Println(p1, p0)
        var v11 int64 = -1
        _ = v11
    } else {
        var v12 int16 = 526
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    return 27
}
func genFn1(p0 string, p1 uint16) int64 {
    v13 := []int{}
    p1 = (p1 ^ 65535)
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 string = (p0 + p0)
        var v16 int64 = -13
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint32 = 253
    for fi18 := 0; fi18 < 9; fi18++ {
        fmt.Println(p1, p0, fi18)
        p0 = p0
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v17
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"hello", ""}
    var v3 int32 = 2147483647
    var v4 int8 = 42
    v5 := []uint32{}
    var v6 int64 = -9223372036854775808
    v19 := GenS1{F0: v6, F1: 4294967295, F2: 4294967295}
    v3 = 2147483647
    var v20 int8 = v4
    var v21 uint = 2
    switch "a" {
    case "hello":
        var v22 uint32 = 4294967295
        _ = v22
    case "\n\t":
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 int64 = (v6 | 42)
            _ = len(v2)
            _ = fi23
            _ = v24
        }
    }
    if true {
        var v25 uint32 = 65535
        _ = v25
    } else {
        _ = len(v1)
    }
    switch 802 {
    case -2147483648:
        switch -292 {
        case -9223372036854775808:
            v21 = (v21 << ((v21) & 63))
        case 891:
            var v26 uint64 = 4294967295
            _ = v26
        default:
            switch -123 {
            case 32768:
                switch "foo bar" {
                case "":
                    var v27 uint = 18446744073709551615
                    _ = v27
                }
            case -9223372036854775808:
                if ((!false) || false) {
                    if false {
                        switch 128 {
                        case -1:
                            v3 = v3
                        case 127:
                            for fi28 := 0; fi28 < 4; fi28++ {
                                fmt.Println(v3, v20, v4)
                                _ = fi28
                            }
                        case -372:
                            switch "foo bar" {
                            case "foo bar":
                                v29 := GenS1{F0: 1, F1: 42, F2: v21}
                                _ = v29
                            }
                        }
                    } else {
                        fmt.Println(v21, v4)
                        var v30 int = 1
                        _ = v30
                    }
                    var v31 uint16 = 2
                    fmt.Println(v3, v31, v4)
                    _ = v31
                } else {
                    v32 := GenS2{F0: "hello", F1: 32767, F2: (v6 | genFn1(("hello" + "hello"), genFn0((v20 & 127), ("" + ("\n\t" + "a")))))}
                    switch 4294967295 {
                    case 4294967295:
                        if true {
                            v3 -= -924
                        }
                    case -289:
                        v4 += 127
                    case 32768:
                        if (415 > 4294967295) {
                            fmt.Println(v4, v3, v20)
                            _ = len(v5)
                        }
                    default:
                        var v33 uint16 = 357
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 19; fi34++ {
                        for fi35 := 0; fi35 < 2; fi35++ {
                            fmt.Println(fi35, v4, v21)
                            _ = fi35
                        }
                        v20 = (v4 | (-(v20)))
                        _ = fi34
                    }
                    _ = v32
                }
            }
        }
    case 642:
        for fi36 := 0; fi36 < 12; fi36++ {
            fmt.Println(v4, v4, v3)
            _ = fi36
        }
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        v20 = int8(v6)
        for fi38 := 0; fi38 < 4; fi38++ {
            _ = len(v5)
            v20 += v20
            switch ("foo bar" + ("" + "\n\t")) {
            case "a":
                fmt.Println(fi38)
            default:
                v39 := GenS3{F0: 18446744073709551615, F1: v3, F2: false}
                _ = v39
            }
            fmt.Println(fi37)
            _ = fi38
        }
        _ = fi37
    }
    fmt.Println(v4, v6)
    v3 = (-(v3))
    switch ("" + "hello") {
    case "a":
        v4 -= v20
    case "":
        for fi40 := 0; fi40 < 5; fi40++ {
            var v41 float32 = ((1e10 / float32((0.1 + 1.5))) / 1.0)
            var v42 uint64 = 4294967295
            _ = fi40
            _ = v41
            _ = v42
        }
    }
    v6 = v6
    for fi43 := 0; fi43 < 10; fi43++ {
        v1[(fi43 | (fi43 | fi43))] = true
        _ = fi43
    }
    v1[-419] = false
    var v44 uint8 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
