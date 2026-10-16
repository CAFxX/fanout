package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int; F1 uint32; F2 bool }
type GenS3 struct { F0 uint16 }
func genFn0(p0 string, p1 uint32, p2 int8) uint16 {
    v4 := []int{32768}
    var v5 int32 = -350
    switch p0 {
    case "hello":
        if (2 == 247) {
            if false {
                for fi6 := 0; fi6 < 19; fi6++ {
                    fmt.Println(p2, p2, p0)
                    _ = fi6
                }
            }
        }
    case "\n\t":
        fmt.Println(v5, p2, p1)
    case "":
        p0 = p0
    }
    switch p0 {
    case "a":
        var v7 string = (((p0 + p0) + p0) + (p0 + (p0 + (p0 + (p0 + (p0 + p0))))))
        _ = v7
    case "":
        var v8 uint32 = (p1 & p1)
        _ = v8
    default:
        if (p2 <= 42) {
            switch (p0 + "") {
            case "":
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 uint8 = 2
                    _ = fi9
                    _ = v10
                }
            case "foo bar":
                var v11 int8 = p2
                _ = v11
            }
        }
    }
    if (v5 >= v5) {
        _ = len(v4)
        fmt.Println(v5)
        v12 := GenS1{F0: (!true), F1: 0}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 2
    v3 := make(map[string]bool)
    v2 -= (^(v2))
    var v13 uint32 = 1
    v14 := GenS2{F0: -9223372036854775808, F1: (v13 >> ((255) & 63)), F2: true}
    delete(v3, "hello")
    fmt.Println(v13, v2)
    delete(v3, "foo bar")
    v13 = v13
    v15 := []int{-1, -2147483648, 9223372036854775807}
    v16 := GenS1{F0: false, F1: 768}
    if true {
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 int = 593
            var v19 string = "\n\t"
            v2 += (v2 >> ((v2) & 63))
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        v13 += v13
        fmt.Println(v13, v2)
        _ = len(v15)
    }
    if true {
        v2 += v2
        v20 := GenS2{F0: 32767, F1: (v13 << ((0) & 63)), F2: true}
        _ = v20
    } else {
        if len(v15) > 0 {
            v15[0] = -44
        }
    }
    if (!(1e10 == (float64(1e10) / 1e10))) {
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 uint8 = 2
            var v23 int64 = 0
            _ = fi21
            _ = v22
            _ = v23
        }
    } else {
        v2 += v2
        v24 := GenS2{F0: 0, F1: v13, F2: false}
        _ = v24
    }
    if (int(v13) <= -9223372036854775808) {
        var v25 string = ("foo bar" + ("\n\t" + "a"))
        var v26 int8 = -1
        switch v25 {
        case "a":
            _ = len(v1)
        case "foo bar":
            v2 = 65535
        case "hello":
            var v27 uint = 1
            _ = v27
        default:
            var v28 uint32 = 2147483647
            _ = v28
        }
        _ = v25
        _ = v26
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        var v30 int = fi29
        switch "\n\t" {
        case "":
            for fi31 := 0; fi31 < 13; fi31++ {
                switch (fi31 & fi29) {
                case -733:
                    var v32 uint64 = 42
                    _ = v32
                case 629:
                    v33 := GenS2{F0: (v30 >> ((18446744073709551615) & 63)), F1: 4294967295, F2: (false && (!true))}
                    _ = v33
                case 32767:
                    v34 := GenS2{F0: (-(fi31)), F1: v13, F2: (-9223372036854775808 == -1)}
                    _ = v34
                }
                fmt.Println(fi31)
                _ = fi31
            }
        case "a":
            v35 := GenS3{F0: 697}
            _ = v35
        }
        _ = fi29
        _ = v30
    }
    v36 := GenS1{F0: false, F1: 0}
    switch "foo bar" {
    case "":
        delete(v1, v13)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v36
}
