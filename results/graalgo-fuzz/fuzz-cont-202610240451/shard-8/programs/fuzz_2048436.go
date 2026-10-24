package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint; F1 string }
func genFn0(p0 int16, p1 int) int8 {
    fmt.Println(p1)
    p1 -= 32767
    switch (p1 << ((39) & 63)) {
    case -271:
        var v6 int64 = -9223372036854775808
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        p0 += 1
        p1 += (p1 / ((fi7) | 1))
        _ = fi7
    }
    p1 = 490
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 uint8, p1 string, p2 bool) uint32 {
    fmt.Println(p2)
    var v8 uint8 = p0
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(p1, fi9)
        if (0.1 != -2.25) {
            switch p1 {
            case "\n\t":
                for fi10 := 0; fi10 < 18; fi10++ {
                    p2 = (-1 >= 42)
                    var v11 uint = 1
                    _ = fi10
                    _ = v11
                }
            case "foo bar":
                v12 := GenS1{F0: p1}
                _ = v12
            case "":
                if (!(((0 <= 518) || false) && (207 == 739))) {
                    var v13 int16 = -864
                    var v14 int32 = -1
                    v8 += (v8 << ((4294967295) & 63))
                    _ = v13
                    _ = v14
                }
            default:
                v15 := GenS1{F0: (p1 + p1)}
                _ = v15
            }
            var v16 int16 = -1
            _ = v16
        } else {
            switch "hello" {
            case "a":
                p0 = 92
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                switch p1 {
                case "foo bar":
                    var v18 bool = (!p2)
                    _ = v18
                case "a":
                    var v19 int64 = int64(fi9)
                    _ = v19
                default:
                    v20 := GenS1{F0: p1}
                    _ = v20
                }
                var v21 uint = 65535
                p1 = "\n\t"
                _ = fi17
                _ = v21
            }
        }
        _ = fi9
    }
    fmt.Println(p1, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    var v4 uint64 = 0
    v5 := make(map[uint32]uint32)
    fmt.Println(v4)
    for fi22 := 0; fi22 < 1; fi22++ {
        v4 -= (v4 | v4)
        _ = fi22
    }
    v4 = 0
    delete(v1, 872)
    v4 -= (v4 * v4)
    v4 = v4
    if ((!((103 <= -32768) && (v4 <= v4))) && (!true)) {
        if (537 < 42) {
            v23 := []int{-9223372036854775808, -2147483648, -1}
            var v24 int16 = 1000
            if true {
                if false {
                    for fi25 := 0; fi25 < 9; fi25++ {
                        v26 := v1[genFn1(0, "foo bar", (v24 != (^(v24))))]
                        _ = fi25
                        _ = v26
                    }
                    switch "\n\t" {
                    case "":
                        v4 = 2
                    case "\n\t":
                        v27 := GenS1{F0: ""}
                        _ = v27
                    case "foo bar":
                        var v28 bool = false
                        _ = v28
                    default:
                        if false {
                            _ = len(v3)
                            var v29 int16 = (v24 ^ 1)
                            switch ("a" + "a") {
                            case "a":
                                v4 += (v4 ^ 4294967295)
                            case "\n\t":
                                for fi30 := 0; fi30 < 4; fi30++ {
                                    if false {
                                        var v31 int64 = 9223372036854775807
                                        _ = v31
                                    } else {
                                        if (v29 < 1) {
                                            var v32 int16 = v24
                                            if len(v23) > 0 {
                                                v23[0] = (^(fi30))
                                            }
                                            _ = v32
                                        }
                                        switch ("" + "") {
                                        case "":
                                            switch "a" {
                                            case "\n\t":
                                                switch "" {
                                                case "foo bar":
                                                    v33 := GenS1{F0: "a"}
                                                    _ = v33
                                                case "a":
                                                    v34 := GenS2{F0: 0, F1: 19, F2: 1}
                                                    _ = v34
                                                }
                                            case "":
                                                delete(v1, 2147483648)
                                            case "foo bar":
                                                var v35 uint = 1
                                                _ = v35
                                            default:
                                                var v36 uint64 = (v4 >> ((80) & 63))
                                                _ = v36
                                            }
                                        case "a":
                                            v37 := v5[1]
                                            _ = v37
                                        default:
                                            v1[310] = ("a" < "foo bar")
                                        }
                                        for fi38 := 0; fi38 < 4; fi38++ {
                                            switch fi30 {
                                            case 127:
                                                v39 := GenS1{F0: ("foo bar" + "\n\t")}
                                                _ = v39
                                            case -9223372036854775808:
                                                fmt.Println(v24, fi30)
                                            default:
                                                v24 = (v29 + 1000)
                                            }
                                            var v40 float64 = (-2.25 - 3.14159)
                                            _ = fi38
                                            _ = v40
                                        }
                                    }
                                    var v41 bool = ((v24 ^ (v24 | 1)) < v24)
                                    fmt.Println(v24, v4)
                                    _ = fi30
                                    _ = v41
                                }
                            }
                            _ = v29
                        } else {
                            var v42 uint = 892
                            v43 := GenS3{F0: 4294967295, F1: "foo bar"}
                            _ = v42
                            _ = v43
                        }
                    }
                } else {
                    if (!false) {
                        switch ("" + ("" + "hello")) {
                        case "a":
                            switch (("" + ("a" + ("hello" + "\n\t"))) + ("foo bar" + ((((((((("a" + "foo bar") + "a") + ("foo bar" + (("\n\t" + "hello") + "foo bar"))) + "foo bar") + "a") + "\n\t") + "foo bar") + "foo bar") + (("foo bar" + (("\n\t" + (("hello" + "") + "\n\t")) + ("foo bar" + ("hello" + "\n\t")))) + "hello")))) {
                            case "a":
                                var v44 uint16 = 1
                                _ = v44
                            case "":
                                v4 = v4
                            }
                        case "\n\t":
                            switch 32768 {
                            case -9223372036854775808:
                                fmt.Println(v24, v24)
                            case 32767:
                                _ = len(v1)
                            }
                        default:
                            fmt.Println(v4, v24)
                        }
                        fmt.Println(v24, v4)
                    }
                }
                var v45 int32 = -759
                v45 = (v45 / (((v45 | v45)) | 1))
                _ = v45
            }
            _ = v23
            _ = v24
        }
        _ = len(v3)
    } else {
        var v46 uint16 = 193
        v47 := GenS3{F0: 1, F1: ""}
        if (!(42 <= -968)) {
            v48 := v3[42]
            fmt.Println(v4, v46, v46)
            _ = v48
        } else {
            var v49 int16 = 32767
            for fi50 := 0; fi50 < 17; fi50++ {
                var v51 string = ""
                v52 := []int{}
                var v53 float64 = 0.1
                _ = fi50
                _ = v51
                _ = v52
                _ = v53
            }
            v46 -= (v46 >> ((uint(v4)) & 63))
            _ = v49
        }
        _ = v46
        _ = v47
    }
    var v54 uint64 = 65
    v4 = (v4 | (v54 % ((0) | 1)))
    var v55 int16 = 466
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
}
