package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int }
func genFn0(p0 int16, p1 int64) int64 {
    var v6 float64 = 0.0
    var v7 int16 = p0
    var v8 bool = true
    v9 := GenS1{F0: 0, F1: 0}
    v10 := []int{-1, 983, 128}
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return p1
}
func genFn1(p0 uint32) int8 {
    fmt.Println(p0)
    switch "hello" {
    case "":
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 bool = (2 > 4294967295)
            _ = fi11
            _ = v12
        }
    case "\n\t":
        if false {
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := []int{fi13}
                _ = fi13
                _ = v14
            }
            v15 := GenS1{F0: 0, F1: (p0 % ((p0) | 1))}
            _ = v15
        }
    case "a":
        switch 32767 {
        case -9223372036854775808:
            p0 = (p0 / ((p0) | 1))
        case 127:
            v16 := []int{32767, -746}
            _ = v16
        }
    default:
        p0 = (^(p0))
    }
    p0 = 2147483647
    v17 := GenS2{F0: 9223372036854775807}
    _ = p0
    _ = v17
    return -23
}

func main() {
    v1 := []string{("a" + "hello"), "", ("foo bar" + "foo bar")}
    var v2 float32 = -2.25
    v3 := make(map[uint32]uint32)
    var v4 int8 = 42
    v5 := make(map[string]string)
    switch ("hello" + "hello") {
    case "foo bar":
        if (255 < 243) {
            switch "foo bar" {
            case "\n\t":
                v2 = 0.1
            }
            var v18 int64 = 334
            v18 += 42
            _ = v18
        } else {
            fmt.Println(v2, v2)
            var v19 uint16 = 0
            _ = v19
        }
    case "\n\t":
        fmt.Println(v2, v4)
    case "a":
        var v20 uint16 = 32767
        _ = v20
    }
    var v21 string = ("a" + "a")
    for fi22 := 0; fi22 < 14; fi22++ {
        switch v21 {
        case "foo bar":
            v23 := GenS2{F0: fi22}
            _ = v23
        case "a":
            fmt.Println(fi22, v21, v4)
        case "hello":
            var v24 bool = (true || ("" < "foo bar"))
            _ = v24
        }
        var v25 string = "\n\t"
        _ = fi22
        _ = v25
    }
    v26 := v3[396]
    if ((!false) && (((-549 > -2147483648) || (32767 >= 1)) || true)) {
        v21 = v21
    } else {
        if (210 >= 255) {
            v26 -= 889
            switch "a" {
            case "a":
                var v27 bool = (false && true)
                _ = v27
            default:
                v28 := GenS2{F0: 42}
                _ = v28
            }
        } else {
            for fi29 := 0; fi29 < 3; fi29++ {
                var v30 float64 = ((-2.25 / 1.5) + (-2.25 - 1e10))
                _ = fi29
                _ = v30
            }
            v26 -= v26
            _ = len(v3)
        }
        if false {
            if ((-531 >= 72) && (117 <= 127)) {
                v21 = ("" + v21)
                v31 := GenS2{F0: 853}
                _ = v31
            } else {
                var v32 int64 = 9223372036854775807
                _ = v32
            }
        } else {
            fmt.Println(v4)
        }
    }
    v26 = (^(v26))
    var v33 float64 = 1.5
    if false {
        v4 -= (v4 % ((v4) | 1))
        for fi34 := 0; fi34 < 9; fi34++ {
            v35 := GenS2{F0: (-(fi34))}
            _ = fi34
            _ = v35
        }
    }
    if (!true) {
        v26 = v26
        if true {
            var v36 float64 = v33
            _ = v36
        } else {
            var v37 int16 = -12
            for fi38 := 0; fi38 < 20; fi38++ {
                switch (-(fi38)) {
                case 32767:
                    switch fi38 {
                    case -506:
                        var v39 uint8 = 255
                        _ = v39
                    case 0:
                        fmt.Println(v21)
                    case -9223372036854775808:
                        for fi40 := 0; fi40 < 7; fi40++ {
                            var v41 int64 = 9223372036854775807
                            fmt.Println(v21, v21)
                            _ = fi40
                            _ = v41
                        }
                    }
                case 4294967295:
                    v37 = 962
                case 2147483647:
                    var v42 uint64 = 2
                    _ = v42
                default:
                    delete(v3, (^(v26)))
                }
                v33 -= v33
                _ = fi38
            }
            _ = v37
        }
    }
    if (18446744073709551615 < 2) {
        for fi43 := 0; fi43 < 2; fi43++ {
            v1 = append(v1, "a")
            var v44 uint = 771
            switch "hello" {
            case "":
                for fi45 := 0; fi45 < 13; fi45++ {
                    v2 -= 1.5
                    _ = fi45
                }
            }
            _ = fi43
            _ = v44
        }
        v4 = (-(v4))
    } else {
        v46 := GenS2{F0: -35}
        _ = v46
    }
    var v47 uint16 = 541
    var v48 int8 = (v4 ^ genFn1((^(v26))))
    var v49 uint8 = 127
    var v50 bool = (v2 != (-2.25 / (float32(1.5) - -2.25)))
    var v51 int8 = (v48 % ((127) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
