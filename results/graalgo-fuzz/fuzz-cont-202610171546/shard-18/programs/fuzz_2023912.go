package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int }
func genFn0(p0 int, p1 string, p2 int) string {
    if false {
        for fi5 := 0; fi5 < 12; fi5++ {
            v6 := []int{42, fi5, (p2 - p0)}
            if (2 == 1) {
                p1 = ("foo bar" + p1)
                var v7 uint64 = 0
                _ = v7
            }
            fmt.Println(p0)
            _ = fi5
            _ = v6
        }
    } else {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 float32 = (1.5 / (1e10 + (3.14159 / 0.1)))
            _ = fi8
            _ = v9
        }
        fmt.Println(p1, p1)
        switch p0 {
        case -9223372036854775808:
            if (98 < 156) {
                for fi10 := 0; fi10 < 9; fi10++ {
                    p0 = (p2 >> ((18446744073709551615) & 63))
                    p1 = (p1 + ("\n\t" + p1))
                    if false {
                        p0 = 32768
                        if true {
                            p0 = p2
                        } else {
                            for fi11 := 0; fi11 < 13; fi11++ {
                                if true {
                                    switch p0 {
                                    case -1:
                                        p0 += (-(p2))
                                    case 127:
                                        p0 -= p2
                                    default:
                                        p2 = p2
                                    }
                                    p0 = (-(p0))
                                }
                                var v12 int16 = 42
                                _ = fi11
                                _ = v12
                            }
                            var v13 string = p1
                            var v14 int = (^(fi10))
                            _ = v13
                            _ = v14
                        }
                        var v15 uint16 = 400
                        _ = v15
                    } else {
                        var v16 uint8 = 0
                        var v17 int32 = -640
                        _ = v16
                        _ = v17
                    }
                    _ = fi10
                }
                v18 := GenS1{F0: 789}
                _ = v18
            } else {
                p1 = p1
                switch p1 {
                case "foo bar":
                    for fi19 := 0; fi19 < 18; fi19++ {
                        var v20 int = (p0 + fi19)
                        v21 := []int{p0, (fi19 * v20), p0, (p2 & 32768)}
                        if (0.0 <= (1.5 - -2.25)) {
                            var v22 bool = false
                            var v23 uint64 = 18446744073709551615
                            switch "hello" {
                            case "foo bar":
                                var v24 string = p1
                                _ = v24
                            case "hello":
                                _ = len(v21)
                            }
                            _ = v22
                            _ = v23
                        } else {
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            p1 = ""
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                case "hello":
                    p2 = (p2 / ((p2) | 1))
                }
                for fi25 := 0; fi25 < 6; fi25++ {
                    fmt.Println(p2, p2)
                    if true {
                        v26 := []int{p2, p2}
                        v27 := GenS2{F0: 1000}
                        var v28 uint = 509
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    _ = fi25
                }
            }
        case 705:
            v29 := []int{}
            _ = v29
        case 128:
            fmt.Println(p2, p1)
        }
    }
    p1 = "\n\t"
    switch (p1 + "\n\t") {
    case "hello":
        p0 += p2
    case "":
        v30 := []int{32768, 1, p2}
        _ = v30
    case "foo bar":
        switch (p0 | 1) {
        case 0:
            if false {
                var v31 bool = true
                v32 := GenS1{F0: 65535}
                _ = v31
                _ = v32
            }
        }
    default:
        for fi33 := 0; fi33 < 18; fi33++ {
            fmt.Println(p2, fi33)
            _ = fi33
        }
    }
    v34 := []int{435, (p0 & p2), (p2 << ((42) & 63))}
    for fi35 := 0; fi35 < 6; fi35++ {
        p1 = p1
        p0 -= -1
        for fi36 := 0; fi36 < 12; fi36++ {
            v34 = append(v34, (p0 >> ((2) & 63)))
            v37 := GenS2{F0: 1}
            p2 += fi35
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    return p1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = -1
    var v3 uint16 = 476
    var v4 uint = 480
    v1 = (v1 << ((v4) & 63))
    var v38 uint = v4
    var v39 uint64 = 994
    v40 := GenS3{F0: 1, F1: 921}
    fmt.Println(v2)
    var v41 string = "a"
    if true {
        v41 = (v41 + ("foo bar" + ("hello" + "foo bar")))
    } else {
        var v42 bool = (((false || false) || (789 > (v3 >> (((v38 / (((v4 % ((v38) | 1))) | 1))) & 63)))) && false)
        v2 -= (v2 + (-(v2)))
        v2 += v2
        _ = v42
    }
    var v43 int8 = -53
    var v44 uint8 = uint8(v3)
    if (!(2 > 4294967295)) {
        v3 = v3
    }
    for fi45 := 0; fi45 < 16; fi45++ {
        if true {
            v43 += v43
        } else {
            fmt.Println(v41, v44, v39)
        }
        v46 := GenS2{F0: v2}
        _ = fi45
        _ = v46
    }
    fmt.Println(v3, v43)
    v47 := []int{0, 32768, 127, 32767}
    if true {
        var v48 int64 = 0
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v47:", len(v47), v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v44
    _ = v47
}
