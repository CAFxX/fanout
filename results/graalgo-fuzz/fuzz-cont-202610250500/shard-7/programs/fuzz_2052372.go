package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint16 }
type GenS2 struct { F0 uint16; F1 bool; F2 int }
type GenS3 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint; F1 bool; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64; F1 uint32; F2 uint8 }
func genFn0(p0 int8) int16 {
    switch "\n\t" {
    case "foo bar":
        var v7 uint = 214
        _ = v7
    }
    p0 -= p0
    if true {
        var v8 int16 = 787
        var v9 float32 = 1e10
        _ = v8
        _ = v9
    }
    _ = p0
    return 32767
}
func genFn1() string {
    if false {
        var v10 uint64 = 462
        switch 127 {
        case 32767:
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 bool = false
                v13 := GenS1{F0: ("" + ""), F1: "hello", F2: 256}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        if (312 < 2147483648) {
            var v14 int64 = 0
            switch "foo bar" {
            case "a":
                if false {
                    fmt.Println(v14)
                } else {
                    var v15 uint = 0
                    switch -857 {
                    case 2147483647:
                        if true {
                            fmt.Println(v10, v14)
                        } else {
                            switch ("hello" + ("a" + ("" + ""))) {
                            case "foo bar":
                                v16 := []int{}
                                _ = v16
                            case "a":
                                for fi17 := 0; fi17 < 14; fi17++ {
                                    if false {
                                        v18 := []int{fi17, (fi17 % ((fi17) | 1)), (fi17 * (fi17 << ((2) & 63)))}
                                        _ = v18
                                    }
                                    var v19 int64 = (v14 - (v14 + v14))
                                    _ = fi17
                                    _ = v19
                                }
                            }
                            for fi20 := 0; fi20 < 3; fi20++ {
                                v21 := GenS2{F0: 925, F1: true, F2: 32767}
                                v22 := []int{fi20, (-(fi20)), 42}
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                            fmt.Println(v15, v15, v10)
                        }
                    }
                    v23 := []int{}
                    _ = v15
                    _ = v23
                }
            case "hello":
                switch "hello" {
                case "hello":
                    v14 += -1
                default:
                    var v24 string = (("hello" + "a") + "\n\t")
                    _ = v24
                }
            default:
                var v25 bool = false
                _ = v25
            }
            v14 = -88
            _ = v14
        } else {
            if (0 >= 0) {
                var v26 int = 32767
                v26 -= -713
                v26 -= v26
                _ = v26
            } else {
                v27 := GenS3{F0: false, F1: 2, F2: 2}
                fmt.Println(v10)
                _ = v27
            }
        }
        _ = v10
    }
    v28 := GenS4{F0: 42, F1: 1}
    var v29 int16 = -1
    switch -862 {
    case 2147483647:
        switch (("hello" + "a") + "a") {
        case "foo bar":
            switch "\n\t" {
            case "":
                v29 += v29
            case "hello":
                fmt.Println(v29)
            case "foo bar":
                v30 := GenS5{F0: 255, F1: (-2.25 < (float64(3.14159) / 0.1)), F2: 1}
                _ = v30
            default:
                v31 := []int{980, 945}
                _ = v31
            }
        case "a":
            for fi32 := 0; fi32 < 20; fi32++ {
                var v33 uint32 = 42
                v29 = genFn0(66)
                _ = fi32
                _ = v33
            }
        }
    case 1:
        var v34 int64 = 1
        _ = v34
    }
    var v35 uint16 = 2
    _ = v28
    _ = v29
    _ = v35
    return ("foo bar" + ("hello" + (("a" + "\n\t") + "foo bar")))
}

func main() {
    v1 := []uint32{1, 972, 988}
    v2 := []string{("hello" + "\n\t"), ""}
    var v3 uint = 4294967295
    var v4 uint64 = 18446744073709551615
    v5 := make(map[int]string)
    var v6 int16 = 294
    var v36 uint64 = (v4 | 4294967295)
    for fi37 := 0; fi37 < 10; fi37++ {
        var v38 uint64 = v36
        _ = len(v5)
        if true {
            if true {
                v39 := GenS6{F0: (v3 % ((2) | 1))}
                _ = v39
            }
            fmt.Println(v4, v4, v38)
        }
        _ = fi37
        _ = v38
    }
    var v40 int8 = 42
    switch 551 {
    case 1:
        v3 = (v3 + v3)
    }
    var v41 bool = (!(!(136 <= 438)))
    v3 -= 0
    for fi42 := 0; fi42 < 8; fi42++ {
        v5[0] = "\n\t"
        v43 := GenS7{F0: 9223372036854775807, F1: 2, F2: 128}
        v44 := GenS2{F0: 255, F1: (v41 || v41), F2: fi42}
        _ = fi42
        _ = v43
        _ = v44
    }
    if v41 {
        var v45 uint16 = 255
        fmt.Println(v41, v40)
        switch "" {
        case "hello":
            v46 := v5[-423]
            _ = v46
        case "a":
            if (255 != 127) {
                var v47 uint8 = 0
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v41 = false
                _ = v47
            }
        }
        _ = v45
    } else {
        fmt.Println(v40, v4, v41)
        fmt.Println(v4, v6)
    }
    fmt.Println(v40, v36)
    for fi48 := 0; fi48 < 12; fi48++ {
        var v49 float64 = 1.5
        var v50 int16 = (-(v6))
        _ = fi48
        _ = v49
        _ = v50
    }
    for fi51 := 0; fi51 < 19; fi51++ {
        fmt.Println(fi51, v4)
        var v52 int8 = -98
        _ = fi51
        _ = v52
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
