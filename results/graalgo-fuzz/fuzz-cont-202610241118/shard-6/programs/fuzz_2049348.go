package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int8; F2 string }
type GenS4 struct { F0 int; F1 bool; F2 uint }
type GenS5 struct { F0 string; F1 uint8; F2 bool }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 string; F1 int }
func genFn0(p0 string, p1 uint64) uint16 {
    v6 := GenS1{F0: 1000, F1: -1, F2: 407}
    v7 := []int{-2147483648, 42}
    switch 9223372036854775807 {
    case 127:
        v8 := GenS2{F0: 65535}
        _ = v8
    case 32768:
        v9 := GenS1{F0: -1, F1: 510, F2: 65535}
        _ = v9
    case 128:
        for fi10 := 0; fi10 < 9; fi10++ {
            if len(v7) > 0 {
                v7[0] = (fi10 / (((fi10 + (fi10 ^ 128))) | 1))
            }
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 561
}
func genFn1(p0 bool, p1 int16, p2 uint64) uint64 {
    v11 := []int{}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v12 uint64 = (^(p2))
    v13 := GenS3{F0: "a", F1: 1, F2: ""}
    v14 := GenS1{F0: p1, F1: -2147483648, F2: 887}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return v12
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = -1
    var v3 uint64 = 2
    v4 := []uint32{518, 1, 65535, 989}
    v5 := []string{((("foo bar" + "foo bar") + "\n\t") + ("\n\t" + ""))}
    for fi15 := 0; fi15 < 3; fi15++ {
        switch (("foo bar" + (("hello" + ("foo bar" + "a")) + "a")) + (("hello" + "foo bar") + "a")) {
        case "hello":
            v16 := GenS4{F0: fi15, F1: false, F2: 2}
            _ = v16
        default:
            var v17 bool = (!(!(true || ((-455 == fi15) && (1 <= 546)))))
            _ = v17
        }
        for fi18 := 0; fi18 < 2; fi18++ {
            v19 := GenS5{F0: "\n\t", F1: 244, F2: (!false)}
            v1 += (v1 / (((v1 ^ (v1 ^ -561))) | 1))
            _ = fi18
            _ = v19
        }
        _ = fi15
    }
    v2 = (^(v2))
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS6{F0: (false && false), F1: 1}
        for fi22 := 0; fi22 < 3; fi22++ {
            var v23 int64 = 725
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    if (!true) {
        fmt.Println(v2)
    } else {
        v24 := GenS6{F0: false, F1: v1}
        v25 := GenS5{F0: "a", F1: 255, F2: false}
        _ = v24
        _ = v25
    }
    var v26 string = ("" + "foo bar")
    var v27 int32 = 1
    switch ("hello" + (("hello" + "a") + v26)) {
    case "foo bar":
        fmt.Println(v2)
    default:
        if false {
            v28 := GenS2{F0: 255}
            if true {
                var v29 int32 = v27
                _ = v29
            }
            v2 -= v2
            _ = v28
        } else {
            var v30 uint64 = (^(v3))
            _ = v30
        }
    }
    var v31 bool = true
    v3 = (^(v3))
    switch ((v26 + "foo bar") + (v26 + "\n\t")) {
    case "a":
        for fi32 := 0; fi32 < 13; fi32++ {
            if false {
                v3 = (v3 * 658)
                if v31 {
                    v33 := GenS1{F0: v2, F1: (v27 + 798), F2: 136}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v34 := GenS2{F0: 0}
                    _ = v33
                    _ = v34
                }
            } else {
                fmt.Println(v1, v26, v27)
                var v35 bool = (0.1 == 1e10)
                _ = v35
            }
            _ = fi32
        }
    }
    if v31 {
        v26 = v26
    }
    v36 := GenS7{F0: v26, F1: 127}
    v4 = append(v4, 955)
    v1 = v1
    for fi37 := 0; fi37 < 17; fi37++ {
        if v31 {
            v27 += v27
        } else {
            var v38 float32 = float32((3.14159 + 0.0))
            _ = len(v5)
            fmt.Println(v38, v26)
            _ = v38
        }
        switch v26 {
        case "":
            if (44 == 21) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v36
    _ = v4
    _ = v5
}
