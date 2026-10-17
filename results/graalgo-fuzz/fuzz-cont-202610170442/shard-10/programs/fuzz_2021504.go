package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 string; F1 int16 }

func main() {
    v1 := []string{("\n\t" + "a"), "foo bar", "a", "\n\t"}
    var v2 uint = 328
    v3 := []int{306, 9223372036854775807}
    var v4 uint = 0
    var v5 uint64 = 4294967295
    var v6 uint64 = v5
    var v7 int16 = -32768
    if ((^(v4)) != (v2 & v2)) {
        v8 := GenS1{F0: v7}
        _ = v8
    }
    if ((931 > 2) && true) {
        var v9 uint64 = 1
        for fi10 := 0; fi10 < 7; fi10++ {
            switch (fi10 * 32768) {
            case 0:
                var v11 string = "foo bar"
                _ = v11
            }
            for fi12 := 0; fi12 < 15; fi12++ {
                v6 -= 1
                _ = len(v3)
                _ = fi12
            }
            if (0 == -2147483648) {
                fmt.Println(v2)
                var v13 float64 = 1.5
                v5 += (^(v9))
                _ = v13
            } else {
                v7 = (v7 + v7)
            }
            _ = fi10
        }
        if (!(255 == 32767)) {
            v2 = 1
        }
        _ = v9
    } else {
        v14 := GenS2{F0: false, F1: false}
        _ = v14
    }
    var v15 int = 303
    switch "" {
    case "":
        var v16 string = ("a" + "\n\t")
        _ = v16
    case "a":
        var v17 int32 = 42
        _ = v17
    }
    v15 -= -1
    v6 -= v5
    switch "" {
    case "a":
        v6 += v6
    case "\n\t":
        switch "\n\t" {
        case "":
            switch "" {
            case "foo bar":
                for fi18 := 0; fi18 < 1; fi18++ {
                    v19 := GenS3{F0: 1}
                    _ = fi18
                    _ = v19
                }
            }
        }
    case "":
        v20 := GenS2{F0: ((false || (false && true)) && ((-327 < -1) && ((!((1e10 + -2.25) > 0.1)) && (!(42 > 725))))), F1: true}
        _ = v20
    }
    var v21 int16 = (^(v7))
    if false {
        if true {
            var v22 int64 = 1
            _ = v22
        }
        var v23 uint64 = (v6 >> (((v2 << (((v4 >> ((v2) & 63))) & 63))) & 63))
        switch (-(v15)) {
        case 1:
            var v24 uint = (v4 & 42)
            _ = v24
        default:
            _ = len(v1)
        }
        _ = v23
    }
    if true {
        switch "\n\t" {
        case "a":
            var v25 bool = (-718 <= 9223372036854775807)
            _ = v25
        case "hello":
            for fi26 := 0; fi26 < 1; fi26++ {
                if false {
                    fmt.Println(v6, v7, v4)
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v7 = -115
                        _ = fi27
                    }
                    for fi28 := 0; fi28 < 5; fi28++ {
                        var v29 int8 = 31
                        var v30 uint32 = 42
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch fi26 {
                        case -2147483648:
                            fmt.Println(v5, v21, v15)
                        case -995:
                            fmt.Println(v2, v6)
                            var v31 int16 = 42
                            fmt.Println(v7, v31, v21)
                            _ = v31
                        case 32768:
                            var v32 uint = (v2 * 289)
                            _ = v32
                        }
                        if true {
                            switch "foo bar" {
                            case "a":
                                var v33 string = "\n\t"
                                _ = v33
                            default:
                                var v34 string = ""
                                _ = v34
                            }
                            v5 -= v5
                            switch fi26 {
                            case 32768:
                                fmt.Println(v2, v15, v4)
                            }
                        }
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                }
                _ = fi26
            }
        case "foo bar":
            if (!(!false)) {
                v21 = v21
                for fi35 := 0; fi35 < 5; fi35++ {
                    v21 -= v21
                    v36 := GenS4{F0: 88, F1: (v5 / ((2) | 1))}
                    _ = fi35
                    _ = v36
                }
                fmt.Println(v2, v7, v7)
            }
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if len(v3) > 0 {
        v3[0] = v15
    }
    v6 -= (v6 / ((v6) | 1))
    v15 += (v15 >> ((v2) & 63))
    v21 = (v21 ^ -1)
    v2 -= (v2 / ((v2) | 1))
    if (0.1 <= -2.25) {
        v7 = v21
        var v37 int = v15
        _ = v37
    } else {
        fmt.Println(v5, v2)
        var v38 bool = true
        v39 := GenS5{F0: (("hello" + ("" + "")) + "a"), F1: (v21 * (v21 - v21))}
        _ = v38
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
