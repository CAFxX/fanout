package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int16; F1 int64; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool; F1 uint8; F2 int }
func genFn0(p0 uint8) uint {
    fmt.Println(p0)
    if (1e10 < 3.14159) {
        v6 := GenS1{F0: -436, F1: -1}
        _ = v6
    } else {
        if false {
            switch 1 {
            case 32768:
                switch "\n\t" {
                case "a":
                    v7 := []int{181}
                    _ = v7
                case "":
                    if true {
                        var v8 string = ""
                        if false {
                            v8 = ""
                        } else {
                            if ("\n\t" <= v8) {
                                var v9 float64 = 0.0
                                v10 := []int{}
                                fmt.Println(p0, v8, v8)
                                _ = v10
                                _ = v9
                            } else {
                                v8 = "hello"
                                for fi11 := 0; fi11 < 20; fi11++ {
                                    var v12 string = ((v8 + v8) + v8)
                                    if ((!(42 > 1)) || (fi11 <= (fi11 & fi11))) {
                                        switch 62 {
                                        case -992:
                                            v13 := []int{fi11, fi11, (fi11 ^ 32768)}
                                            _ = v13
                                        }
                                    }
                                    v14 := []int{(fi11 >> ((2) & 63)), 127, (^(fi11)), 907}
                                    _ = fi11
                                    _ = v12
                                    _ = v14
                                }
                                for fi15 := 0; fi15 < 5; fi15++ {
                                    for fi16 := 0; fi16 < 6; fi16++ {
                                        fmt.Println(fi15, fi15, v8)
                                        _ = fi16
                                    }
                                    v17 := GenS1{F0: 691, F1: -815}
                                    _ = fi15
                                    _ = v17
                                }
                            }
                            fmt.Println(v8, p0)
                        }
                        _ = v8
                    }
                }
            }
            var v18 uint32 = 0
            var v19 int8 = 42
            _ = v18
            _ = v19
        } else {
            v20 := []int{-2147483648}
            var v21 float32 = (3.14159 - 1e10)
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 float64 = 3.14159
                var v24 float64 = v23
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = v20
            _ = v21
        }
        if (1 <= 42) {
            if (-1 <= -1) {
                v25 := GenS2{F0: 9223372036854775807, F1: 1000}
                _ = v25
            } else {
                fmt.Println(p0)
                p0 -= p0
                var v26 uint = 255
                _ = v26
            }
            var v27 uint16 = 0
            _ = v27
        }
        v28 := GenS3{F0: 655, F1: "foo bar"}
        _ = v28
    }
    p0 = (p0 * p0)
    fmt.Println(p0)
    _ = p0
    return 0
}
func genFn1() uint {
    for fi29 := 0; fi29 < 17; fi29++ {
        v30 := GenS4{F0: 0, F1: 42, F2: 32767}
        fmt.Println(fi29)
        _ = fi29
        _ = v30
    }
    v31 := []int{}
    v32 := GenS5{F0: 4294967295}
    if (!(2 <= 42)) {
        v33 := GenS6{F0: 2}
        _ = v33
    } else {
        v34 := GenS2{F0: -1, F1: 0}
        for fi35 := 0; fi35 < 11; fi35++ {
            switch "foo bar" {
            case "\n\t":
                if len(v31) > 0 {
                    v31[0] = 127
                }
            case "hello":
                fmt.Println(fi35)
            case "":
                if (-67 > 215) {
                    var v36 uint8 = 2
                    _ = v36
                }
            }
            if ("" == "a") {
                var v37 uint16 = 65535
                v38 := GenS7{F0: -1}
                if (!true) {
                    if (fi35 != fi35) {
                        var v39 int = (fi35 ^ 2147483647)
                        switch 1 {
                        case 470:
                            if (float64((0.0 / 1.0)) != -2.25) {
                                v39 -= fi35
                            }
                        case 128:
                            var v40 int16 = -284
                            _ = v40
                        }
                        _ = v39
                    }
                    var v41 int = (fi35 - 612)
                    var v42 int8 = -1
                    _ = v41
                    _ = v42
                }
                _ = v37
                _ = v38
            } else {
                for fi43 := 0; fi43 < 14; fi43++ {
                    var v44 uint16 = 0
                    v44 = v44
                    _ = fi43
                    _ = v44
                }
                v45 := GenS8{F0: true, F1: 176, F2: -197}
                fmt.Println(fi35)
                _ = v45
            }
            _ = fi35
        }
        _ = v34
    }
    _ = v31
    _ = v32
    return 2
}

func main() {
    var v1 int = 123
    var v2 uint16 = 0
    v3 := []uint32{}
    v4 := []string{"hello", ("a" + "\n\t"), (((((("hello" + ("" + "hello")) + ("hello" + "")) + "foo bar") + ("" + "\n\t")) + "a") + ""), ""}
    var v5 int8 = -1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 = -624
    v46 := GenS8{F0: (!false), F1: 174, F2: v1}
    var v47 float64 = ((0.0 - (((1e10 * 0.1) / (0.1 * 1e10)) + (float64(float32(float64(float32((0.1 / 1e10))))) / 1.0))) - 1e10)
    v48 := GenS8{F0: (0.1 != 0.1), F1: 26, F2: 161}
    var v49 float64 = v47
    v4 = append(v4, "hello")
    switch "foo bar" {
    case "foo bar":
        var v50 int = (v1 ^ 9223372036854775807)
        _ = v50
    case "\n\t":
        v5 = -94
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
}
