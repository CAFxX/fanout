package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool }
func genFn0(p0 uint32, p1 uint8, p2 uint8) int8 {
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 float32 = 1e10
        var v7 int64 = -128
        var v8 uint8 = uint8(p0)
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := []int{(fi9 | (fi9 & (^(fi9))))}
        if len(v10) > 0 {
            v10[0] = (fi9 * fi9)
        }
        fmt.Println(p1, fi9, p2)
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: "a"}
    switch 1 {
    case 2147483647:
        var v12 float32 = 0.0
        _ = v12
    }
    v13 := GenS1{F0: "\n\t"}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v13
    return -128
}
func genFn1(p0 int8, p1 uint8, p2 uint8) int32 {
    fmt.Println(p1)
    if false {
        if true {
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := []int{(fi14 * (fi14 - (fi14 + fi14))), fi14, fi14}
                var v16 int16 = 0
                _ = fi14
                _ = v15
                _ = v16
            }
            if (316 <= 0) {
                for fi17 := 0; fi17 < 6; fi17++ {
                    for fi18 := 0; fi18 < 9; fi18++ {
                        p1 += (^(p2))
                        fmt.Println(p2, fi18)
                        _ = fi18
                    }
                    if (fi17 > (fi17 | fi17)) {
                        for fi19 := 0; fi19 < 6; fi19++ {
                            var v20 int32 = -71
                            p0 -= (-(p0))
                            v21 := GenS2{F0: "foo bar", F1: 1, F2: 402}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    }
                    var v22 int32 = 1
                    _ = fi17
                    _ = v22
                }
                v23 := GenS3{F0: 18446744073709551615}
                p2 = p1
                _ = v23
            } else {
                var v24 int16 = 0
                _ = v24
            }
            fmt.Println(p2, p1, p2)
        }
        for fi25 := 0; fi25 < 6; fi25++ {
            switch (fi25 ^ (fi25 >> ((18446744073709551615) & 63))) {
            case 136:
                fmt.Println(p0)
            default:
                if true {
                    for fi26 := 0; fi26 < 20; fi26++ {
                        fmt.Println(p1, fi26)
                        fmt.Println(p1, fi26, p1)
                        _ = fi26
                    }
                    fmt.Println(p0, p1)
                } else {
                    p0 = 42
                }
            }
            v27 := GenS3{F0: 42}
            v28 := GenS4{F0: (((!(true && false)) || (1.5 < float32((3.14159 * 0.0)))) && true)}
            _ = fi25
            _ = v27
            _ = v28
        }
        v29 := []int{-638, 128}
        _ = v29
    }
    var v30 int16 = -815
    switch 42 {
    case 1:
        for fi31 := 0; fi31 < 12; fi31++ {
            if false {
                v32 := []int{32768, (fi31 - (^(fi31))), 32767}
                _ = v32
            } else {
                for fi33 := 0; fi33 < 10; fi33++ {
                    v30 -= (v30 | (v30 >> ((575) & 63)))
                    v34 := GenS3{F0: 878}
                    var v35 float32 = -2.25
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            }
            _ = fi31
        }
    case -761:
        p1 -= p2
    case 127:
        switch "hello" {
        case "foo bar":
            var v36 uint8 = (p1 * (p1 >> ((0) & 63)))
            _ = v36
        case "a":
            v37 := []int{128, -2147483648, 127}
            _ = v37
        }
    }
    v38 := GenS2{F0: "\n\t", F1: 1, F2: 1000}
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    _ = v38
    return -767
}

func main() {
    var v1 bool = (1000 > 0)
    v2 := make(map[int]string)
    var v3 int32 = 528
    v4 := make(map[int]int)
    switch 0 {
    case 1:
        if true {
            for fi39 := 0; fi39 < 6; fi39++ {
                v3 += int32(fi39)
                _ = fi39
            }
            v3 = (v3 | -1)
            v1 = v1
        }
    default:
        for fi40 := 0; fi40 < 7; fi40++ {
            var v41 uint64 = 42
            fmt.Println(fi40)
            _ = fi40
            _ = v41
        }
    }
    fmt.Println(v3)
    switch "\n\t" {
    case "foo bar":
        var v42 int = -742
        _ = v42
    }
    var v43 int = 32767
    var v44 uint64 = 18446744073709551615
    v44 -= (v44 | v44)
    var v45 int8 = 127
    fmt.Println(v1, v44)
    switch "foo bar" {
    case "":
        var v46 uint64 = v44
        _ = v46
    }
    v47 := []int{v43, 4294967295, 32768, (v43 >> ((519) & 63))}
    for fi48 := 0; fi48 < 12; fi48++ {
        var v49 int32 = 42
        v49 = v49
        v49 = 0
        _ = fi48
        _ = v49
    }
    if (("a" + "\n\t") != "a") {
        fmt.Println(v44, v1, v43)
    }
    fmt.Println(v45)
    fmt.Println(v43)
    v50 := GenS5{F0: (true && v1)}
    switch ("\n\t" + (("a" + "foo bar") + "\n\t")) {
    case "\n\t":
        var v51 uint32 = 0
        _ = v51
    case "":
        fmt.Println(v45, v1, v44)
    }
    var v52 int = (v43 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v47
    _ = v50
    _ = v52
}
