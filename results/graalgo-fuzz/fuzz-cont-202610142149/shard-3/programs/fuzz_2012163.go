package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 int32; F2 uint8 }
func genFn0(p0 int16) string {
    p0 += (p0 & p0)
    p0 = p0
    v5 := GenS1{F0: -119, F1: 65}
    if false {
        switch 32767 {
        case -61:
            p0 -= p0
        default:
            var v6 int8 = -1
            _ = v6
        }
    } else {
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 20; fi8++ {
                switch "" {
                case "a":
                    if (0 >= 127) {
                        p0 += 1000
                        var v9 bool = true
                        var v10 uint16 = 255
                        _ = v10
                        _ = v9
                    }
                case "hello":
                    p0 -= (p0 * p0)
                case "":
                    if (true && true) {
                        p0 = p0
                    } else {
                        p0 = p0
                    }
                }
                switch fi8 {
                case -115:
                    if true {
                        v11 := []int{}
                        var v12 uint16 = 262
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        _ = v11
                        _ = v12
                    }
                }
                _ = fi8
            }
            _ = fi7
        }
        if (false && (false || ("hello" < ((("a" + ("\n\t" + "\n\t")) + ("hello" + "a")) + ("" + "a"))))) {
            v13 := []int{}
            if false {
                var v14 int = 200
                _ = v14
            } else {
                var v15 string = "hello"
                _ = v15
            }
            _ = v13
        } else {
            v16 := []int{42, -671, 127, 905}
            switch -9223372036854775808 {
            case 32768:
                var v17 uint = 65535
                _ = v17
            }
            p0 += (p0 + (-(p0)))
            _ = v16
        }
        p0 += 32767
    }
    p0 = (p0 * 32767)
    _ = p0
    _ = v5
    return ("hello" + "hello")
}
func genFn1(p0 uint8, p1 int) string {
    switch p1 {
    case 32767:
        fmt.Println(p0)
    case 32768:
        switch "\n\t" {
        case "a":
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := []int{p1}
                v20 := GenS2{F0: 42}
                _ = fi18
                _ = v19
                _ = v20
            }
        case "":
            v21 := []int{p1, (p1 * (p1 - 32768)), (p1 * p1), (-(p1))}
            _ = v21
        case "\n\t":
            v22 := []int{32767, (p1 - (p1 - p1))}
            _ = v22
        default:
            v23 := []int{p1, p1, p1}
            _ = v23
        }
    default:
        var v24 uint64 = 18446744073709551615
        _ = v24
    }
    if false {
        var v25 float64 = 1e10
        _ = v25
    }
    var v26 int16 = -32768
    var v27 float32 = 1e10
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 32767
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    for fi28 := 0; fi28 < 17; fi28++ {
        switch "hello" {
        case "\n\t":
            v2 = v2
        }
        _ = fi28
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        if true {
            v2 += (v2 ^ v2)
            for fi30 := 0; fi30 < 13; fi30++ {
                switch "a" {
                case "foo bar":
                    v31 := []int{(fi30 - fi29)}
                    _ = v31
                case "\n\t":
                    v32 := GenS2{F0: 1000}
                    _ = v32
                }
                _ = fi30
            }
            v2 -= v2
        } else {
            var v33 int = fi29
            v33 = v33
            _ = v33
        }
        var v34 int16 = 122
        _ = fi29
        _ = v34
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        v2 -= 0
        if (true || true) {
            var v36 uint8 = 2
            var v37 int32 = -2147483648
            var v38 uint16 = 873
            _ = v36
            _ = v37
            _ = v38
        }
        _ = fi35
    }
    if ((v2 ^ (v2 | v2)) == (v2 >> ((224) & 63))) {
        v2 -= v2
    } else {
        v2 = v2
    }
    v39 := GenS3{F0: 9223372036854775807, F1: 255}
    fmt.Println(v2)
    fmt.Println(v2)
    switch -233 {
    case -395:
        v2 = 2
    case -1:
        for fi40 := 0; fi40 < 7; fi40++ {
            fmt.Println(v2)
            fmt.Println(fi40, v2)
            _ = fi40
        }
    }
    v2 += (v2 | 139)
    v41 := GenS2{F0: -1}
    var v42 int16 = 508
    _ = len(v1)
    for fi43 := 0; fi43 < 20; fi43++ {
        var v44 uint8 = 0
        _ = fi43
        _ = v44
    }
    v45 := GenS4{F0: 180}
    v46 := GenS5{F0: 100, F1: 2147483647, F2: 42}
    v47 := GenS2{F0: v42}
    v2 = (v2 & (v2 << ((255) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v46
    _ = v47
}
