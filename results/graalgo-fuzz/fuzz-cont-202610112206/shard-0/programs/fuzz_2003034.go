package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 string }
func genFn0(p0 string, p1 uint32, p2 uint64) uint16 {
    for fi5 := 0; fi5 < 18; fi5++ {
        if ((^(p2)) <= p2) {
            p1 -= 4294967295
        } else {
            if (false && true) {
                var v6 string = p0
                p2 += p2
                _ = v6
            } else {
                var v7 int8 = -96
                _ = v7
            }
            for fi8 := 0; fi8 < 16; fi8++ {
                fmt.Println(fi5, p1)
                var v9 int32 = -2147483648
                var v10 string = (p0 + p0)
                _ = fi8
                _ = v10
                _ = v9
            }
            p2 -= p2
        }
        v11 := GenS1{F0: 0, F1: 2, F2: 37}
        _ = fi5
        _ = v11
    }
    var v12 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 0
}
func genFn1(p0 uint64, p1 int, p2 bool) uint16 {
    p1 = (p1 << ((2) & 63))
    fmt.Println(p0, p0, p0)
    v13 := []int{-1}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 435
}

func main() {
    v1 := []uint32{}
    var v2 int64 = 0
    var v3 uint16 = 785
    var v4 int16 = 1
    if len(v1) > 0 {
        v1[0] = 42
    }
    switch "\n\t" {
    case "":
        var v14 uint32 = 65535
        _ = v14
    case "hello":
        var v15 uint = 877
        _ = v15
    }
    switch 32767 {
    case -2147483648:
        v16 := GenS2{F0: ""}
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 uint64 = 0
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 bool = (1 == 188)
            _ = fi20
            _ = v21
        }
        switch "a" {
        case "\n\t":
            if false {
                v22 := GenS2{F0: ""}
                if ((v3 & v3) > 255) {
                    _ = len(v1)
                    var v23 uint32 = 4294967295
                    _ = v23
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                var v24 int = 2147483647
                _ = v22
                _ = v24
            } else {
                v25 := GenS3{F0: -2147483648, F1: "hello"}
                var v26 string = ""
                _ = v25
                _ = v26
            }
        case "hello":
            var v27 int8 = 42
            _ = v27
        case "":
            var v28 int64 = v2
            _ = v28
        }
        _ = fi19
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 uint64 = 1
        v4 -= v4
        v31 := GenS2{F0: ("\n\t" + "foo bar")}
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 int = -159
    for fi33 := 0; fi33 < 16; fi33++ {
        var v34 string = "hello"
        _ = fi33
        _ = v34
    }
    var v35 string = "\n\t"
    v4 += v4
    v36 := GenS1{F0: (v4 & v4), F1: 1, F2: v2}
    switch v35 {
    case "hello":
        switch "a" {
        case "a":
            _ = len(v1)
        default:
            for fi37 := 0; fi37 < 13; fi37++ {
                if true {
                    for fi38 := 0; fi38 < 8; fi38++ {
                        v39 := GenS1{F0: v4, F1: 18446744073709551615, F2: (v2 * v2)}
                        v40 := GenS2{F0: ""}
                        if (1 != 4294967295) {
                            var v41 string = ((v35 + "") + (("foo bar" + (v35 + v35)) + (v35 + v35)))
                            _ = v41
                        } else {
                            var v42 uint32 = 42
                            var v43 uint = 65535
                            var v44 string = "hello"
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    fmt.Println(v35)
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v45 int16 = (v4 & (v4 / ((-1) | 1)))
                    var v46 int16 = (v4 << ((2) & 63))
                    _ = v45
                    _ = v46
                }
                _ = fi37
            }
        }
    case "a":
        v1 = append(v1, 829)
    }
    v2 = (-(v2))
    var v47 int64 = v2
    v2 += v47
    var v48 uint8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v47
    _ = v48
}
