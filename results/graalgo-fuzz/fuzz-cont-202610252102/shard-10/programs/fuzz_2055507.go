package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS3 struct { F0 string; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 bool = ((!((!(!(!(!(true || false))))) && true)) && true)
    var v2 int64 = 0
    var v3 int32 = 1
    var v4 uint = 0
    v5 := []int{4294967295, 505, -923, 983}
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint64 = 0
        for fi8 := 0; fi8 < 2; fi8++ {
            if (false && (((0.0 / 1.0) > 3.14159) && ((-984 == int32(v2)) && v1))) {
                var v9 int16 = -1
                if v1 {
                    if v1 {
                        var v10 int16 = v9
                        var v11 int32 = (v3 % (((v3 >> ((413) & 63))) | 1))
                        v10 = (v10 ^ (v9 % ((1000) | 1)))
                        _ = v10
                        _ = v11
                    }
                    v3 -= -804
                }
                _ = v9
            } else {
                v2 -= v2
            }
            var v12 uint64 = (^(v7))
            var v13 string = "a"
            v14 := GenS1{F0: ((v13 + v13) + v13), F1: v7}
            _ = fi8
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi6
        _ = v7
    }
    switch ("" + "a") {
    case "a":
        v2 -= v2
    case "":
        switch 0 {
        case -9223372036854775808:
            fmt.Println(v1, v4, v4)
        }
    }
    var v15 int = 42
    if false {
        fmt.Println(v3)
        switch (v15 - v15) {
        case -1:
            v1 = v1
        case 449:
            v16 := GenS1{F0: (("\n\t" + "foo bar") + "a"), F1: 4294967295}
            _ = v16
        default:
            for fi17 := 0; fi17 < 13; fi17++ {
                switch fi17 {
                case -646:
                    v5 = append(v5, (v15 | -394))
                }
                _ = len(v5)
                _ = fi17
            }
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS2{F0: v2, F1: 2, F2: 127}
            _ = fi18
            _ = v19
        }
    } else {
        v20 := GenS2{F0: (v2 + (v2 ^ (v2 | v2))), F1: 1, F2: 128}
        _ = v20
    }
    if (((736 != 0) || true) && (v1 && (v1 || false))) {
        v3 = (v3 + -218)
        var v21 string = ""
        _ = v21
    } else {
        v22 := GenS1{F0: "\n\t", F1: 4294967295}
        v5 = append(v5, v15)
        if len(v5) > 0 {
            v5[0] = 128
        }
        _ = v22
    }
    v4 -= 837
    v23 := GenS1{F0: "a", F1: 4294967295}
    var v24 uint32 = 60
    for fi25 := 0; fi25 < 16; fi25++ {
        fmt.Println(fi25)
        for fi26 := 0; fi26 < 15; fi26++ {
            v2 -= (v2 - (-(v2)))
            var v27 uint16 = 255
            v2 = (v2 % ((v2) | 1))
            v28 := GenS3{F0: ("\n\t" + "\n\t"), F1: 139, F2: v24}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (!v1) {
        v29 := GenS4{F0: 0}
        _ = v29
    } else {
        v4 -= (v4 - v4)
        v5 = append(v5, 61)
        var v30 string = ("a" + "")
        _ = v30
    }
    v2 = v2
    for fi31 := 0; fi31 < 8; fi31++ {
        _ = len(v5)
        _ = fi31
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        var v33 bool = (v1 && (42 <= 182))
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
