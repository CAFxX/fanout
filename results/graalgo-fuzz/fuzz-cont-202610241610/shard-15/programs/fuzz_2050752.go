package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint16 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int64 }
func genFn0() int16 {
    v7 := []int{-2147483648}
    var v8 float64 = 0.1
    _ = v7
    _ = v8
    return 32767
}
func genFn1() int16 {
    var v9 int64 = -246
    if false {
        fmt.Println(v9)
        switch "a" {
        case "foo bar":
            var v10 int = 4294967295
            _ = v10
        case "\n\t":
            v11 := GenS1{F0: 189, F1: 861, F2: 1}
            _ = v11
        default:
            v9 -= v9
        }
        v9 -= 0
    } else {
        var v12 uint8 = 251
        fmt.Println(v12, v12)
        v13 := GenS2{F0: v9, F1: true}
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 2147483648, F1: 128, F2: 679}
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := []int{0, -697}
        switch 0 {
        case 127:
            if (false || (true && true)) {
                fmt.Println(fi15, v9)
                v9 += v9
            }
        case 0:
            var v17 int16 = -1
            _ = v17
        case 1:
            fmt.Println(fi15)
        default:
            _ = len(v16)
        }
        _ = fi15
        _ = v16
    }
    switch 1 {
    case 2147483647:
        switch ("foo bar" + "\n\t") {
        case "":
            switch ("\n\t" + "\n\t") {
            case "":
                for fi18 := 0; fi18 < 5; fi18++ {
                    if true {
                        v19 := GenS1{F0: 0, F1: fi18, F2: 65535}
                        v9 = v9
                        v20 := []int{fi18}
                        _ = v19
                        _ = v20
                    }
                    switch 32767 {
                    case 4294967295:
                        v21 := []int{}
                        _ = v21
                    case 128:
                        v22 := []int{-2147483648, (fi18 / ((4294967295) | 1)), (^(fi18)), 42}
                        _ = v22
                    case -881:
                        var v23 uint = 407
                        _ = v23
                    default:
                        if (("a" + ("hello" + "")) > "foo bar") {
                            v24 := []int{4294967295, fi18}
                            _ = v24
                        }
                    }
                    switch "hello" {
                    case "\n\t":
                        for fi25 := 0; fi25 < 10; fi25++ {
                            if (-461 > v9) {
                                v9 += -1
                                v26 := GenS2{F0: v9, F1: (false || true)}
                                _ = v26
                            }
                            var v27 string = ((((("\n\t" + "") + "") + ("\n\t" + "hello")) + "hello") + "foo bar")
                            _ = fi25
                            _ = v27
                        }
                    case "foo bar":
                        v9 += -1
                    default:
                        fmt.Println(fi18, v9)
                    }
                    _ = fi18
                }
            default:
                if ((0.0 != 3.14159) || true) {
                    v28 := GenS3{F0: (("" + "hello") + "a")}
                    _ = v28
                } else {
                    switch 127 {
                    case 259:
                        var v29 string = "hello"
                        _ = v29
                    }
                }
            }
        case "foo bar":
            var v30 string = ""
            _ = v30
        }
    case 32768:
        fmt.Println(v9)
    default:
        if false {
            var v31 float64 = 3.14159
            v32 := []int{652}
            _ = v31
            _ = v32
        }
    }
    _ = v14
    _ = v9
    return -588
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 201
    var v3 bool = false
    v4 := []uint32{4294967295}
    v5 := make(map[string]bool)
    v6 := []int{-2147483648, -751, 1}
    for fi33 := 0; fi33 < 3; fi33++ {
        if (1 >= 1) {
            fmt.Println(v3, fi33)
            v34 := GenS1{F0: 1, F1: fi33, F2: 32767}
            v35 := GenS4{F0: fi33, F1: "\n\t"}
            _ = v34
            _ = v35
        }
        _ = fi33
    }
    if true {
        var v36 uint8 = 128
        fmt.Println(v2, v36)
        _ = v36
    } else {
        if (v3 || v3) {
            v3 = v3
        }
    }
    var v37 uint8 = 151
    v1[2] = 32767
    var v38 uint32 = 2147483647
    for fi39 := 0; fi39 < 1; fi39++ {
        v38 += (^(v38))
        var v40 int = 4294967295
        v41 := GenS5{F0: 32767, F1: (v37 ^ (v37 - v37))}
        _ = fi39
        _ = v40
        _ = v41
    }
    v42 := GenS3{F0: "foo bar"}
    fmt.Println(v37, v3, v2)
    if (!true) {
        var v43 float32 = (3.14159 + -2.25)
        if v3 {
            v3 = false
            v38 = uint32(v2)
        } else {
            v5[(("" + "foo bar") + ("" + "\n\t"))] = (!false)
            var v44 string = (("hello" + "") + "foo bar")
            _ = v44
        }
        _ = v43
    }
    v2 = (^(v2))
    v38 -= v38
    v45 := GenS6{F0: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v45
    _ = v5
    _ = v6
}
