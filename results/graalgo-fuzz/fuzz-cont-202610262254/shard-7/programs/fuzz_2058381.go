package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 int16 }
func genFn0(p0 int, p1 uint16, p2 uint16) string {
    fmt.Println(p1, p1, p2)
    var v7 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ""
}
func genFn1(p0 uint16) uint8 {
    for fi8 := 0; fi8 < 2; fi8++ {
        if (32767 < 1) {
            fmt.Println(p0)
        } else {
            switch (("hello" + ("a" + "foo bar")) + "\n\t") {
            case "a":
                v9 := []int{}
                _ = v9
            default:
                if (!(true || (!false))) {
                    fmt.Println(fi8, fi8)
                }
            }
        }
        p0 = p0
        var v10 int16 = -32768
        _ = fi8
        _ = v10
    }
    v11 := []int{-772, 32768, 32768, 2147483647}
    p0 = (p0 | p0)
    var v12 string = "a"
    _ = p0
    _ = v11
    _ = v12
    return 128
}

func main() {
    var v1 float32 = 1e10
    v2 := []int{-209, 424}
    v3 := make(map[int]bool)
    var v4 uint = 0
    var v5 int16 = -418
    v6 := make(map[int]uint32)
    fmt.Println(v5, v5, v5)
    var v13 int64 = 0
    var v14 float32 = v1
    v13 -= (v13 + -500)
    switch ("\n\t" + ("hello" + ("\n\t" + ("a" + "a")))) {
    case "\n\t":
        switch 32767 {
        case -424:
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := GenS1{F0: 358}
                _ = fi15
                _ = v16
            }
        }
    case "":
        var v17 uint8 = 128
        _ = v17
    case "foo bar":
        if (988 == 42) {
            fmt.Println(v1, v4)
        }
    }
    switch "foo bar" {
    case "foo bar":
        var v18 uint = 238
        _ = v18
    case "\n\t":
        switch "a" {
        case "foo bar":
            switch "foo bar" {
            case "a":
                var v19 int = 357
                _ = v19
            }
        }
    case "hello":
        if (9223372036854775807 < -2147483648) {
            fmt.Println(v14, v13, v4)
        } else {
            v20 := GenS1{F0: 1}
            _ = v20
        }
    default:
        fmt.Println(v5)
    }
    v1 = v1
    var v21 float64 = 1.5
    var v22 int16 = (v5 / ((-1) | 1))
    var v23 int32 = 0
    if true {
        if false {
            if len(v2) > 0 {
                v2[0] = 0
            }
        } else {
            v2 = append(v2, -522)
        }
        v13 = v13
    } else {
        if (2 != 105) {
            switch -677 {
            case 32767:
                fmt.Println(v14)
            case 32768:
                var v24 int = 2147483647
                _ = v24
            }
            if true {
                _ = len(v2)
            }
            fmt.Println(v23, v1)
        } else {
            v5 = v5
        }
        var v25 int8 = 1
        if (0 < 0) {
            if (("foo bar" + ("\n\t" + ("a" + ""))) != "") {
                v26 := GenS1{F0: 42}
                if true {
                    var v27 uint8 = 0
                    v1 = v1
                    _ = v27
                }
                var v28 string = "foo bar"
                _ = v26
                _ = v28
            } else {
                switch 9223372036854775807 {
                case 914:
                    v29 := GenS2{F0: 32, F1: 128}
                    _ = v29
                case 384:
                    if (-797 > (v23 % ((v23) | 1))) {
                        v30 := GenS3{F0: (v13 / (((v13 % ((v13) | 1))) | 1))}
                        _ = v30
                    }
                default:
                    var v31 uint64 = 1
                    _ = v31
                }
                for fi32 := 0; fi32 < 2; fi32++ {
                    var v33 int = (fi32 ^ 176)
                    v34 := GenS4{F0: 32767, F1: v5}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
            v35 := GenS3{F0: (^(v13))}
            fmt.Println(v4)
            _ = v35
        }
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
