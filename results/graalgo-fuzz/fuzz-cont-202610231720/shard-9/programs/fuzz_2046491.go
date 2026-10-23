package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int8, p1 uint8, p2 uint16) int32 {
    var v6 int64 = 576
    var v7 uint8 = (p1 / (((p1 & (p1 - 2))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}
func genFn1() int64 {
    v8 := GenS1{F0: "a", F1: (("hello" + ("foo bar" + "")) + ("\n\t" + "hello"))}
    var v9 int16 = 0
    fmt.Println(v9)
    v9 -= v9
    _ = v8
    _ = v9
    return 380
}

func main() {
    var v1 int16 = 1
    var v2 uint32 = 65535
    var v3 int16 = (v1 >> ((1) & 63))
    var v4 int16 = 32767
    v5 := make(map[string]int)
    if true {
        v2 -= v2
    }
    v3 = v1
    var v10 float32 = 0.1
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            if (!(9 == 1)) {
                var v13 int32 = genFn0(100, 128, 2)
                var v14 uint32 = 42
                _ = v13
                _ = v14
            } else {
                switch (-(fi11)) {
                case 128:
                    fmt.Println(v10)
                case 0:
                    v2 = v2
                default:
                    v15 := []int{(fi11 ^ int(v2)), 888, (fi12 / ((fi11) | 1)), (fi11 / ((fi12) | 1))}
                    _ = v15
                }
                if false {
                    v16 := []int{}
                    _ = v16
                } else {
                    v17 := GenS1{F0: "\n\t", F1: "\n\t"}
                    v4 += v4
                    var v18 uint = 0
                    v19 := []int{}
                    _ = v17
                    _ = v18
                    _ = v19
                }
                v20 := GenS1{F0: (("foo bar" + "\n\t") + ((("hello" + ("a" + "")) + "") + "\n\t")), F1: ""}
                fmt.Println(fi11, fi11)
                v4 -= v4
                if false {
                    fmt.Println(v2, v2, v4)
                    switch 304 {
                    case -2147483648:
                        if (true && (true && ((255 <= 74) && false))) {
                            v21 := []int{}
                            _ = v21
                        }
                    default:
                        if false {
                            v10 -= v10
                        }
                    }
                } else {
                    v1 = (v1 / (((v1 << ((426) & 63))) | 1))
                }
                _ = v20
            }
            if false {
                if (110 <= -1) {
                    var v22 float32 = (-2.25 * v10)
                    v23 := []int{(fi12 ^ fi11), 617}
                    var v24 string = "hello"
                    if true {
                        v1 -= (v1 ^ 1000)
                        fmt.Println(v4)
                        if len(v23) > 0 {
                            v23[0] = int(v2)
                        }
                    } else {
                        v25 := GenS1{F0: "hello", F1: "hello"}
                        v3 += (v1 << ((2) & 63))
                        _ = v25
                    }
                    fmt.Println(v2, fi12, v3)
                    _ = v22
                    _ = v23
                    _ = v24
                }
                if ("foo bar" >= "a") {
                    delete(v5, (("" + "foo bar") + ("" + "\n\t")))
                }
                switch 360 {
                case -110:
                    var v26 uint32 = (v2 >> ((4294967295) & 63))
                    _ = v26
                }
            }
            switch "" {
            case "":
                fmt.Println(v3, v4)
            case "a":
                var v27 uint = 1
                v1 += (v4 / (((v3 | (v4 | 0))) | 1))
                _ = v27
            }
            _ = fi12
        }
        switch 32767 {
        case -9223372036854775808:
            for fi28 := 0; fi28 < 17; fi28++ {
                if false {
                    fmt.Println(v3, v10, v4)
                } else {
                    v29 := []int{(fi28 & fi11), fi11, fi11, (fi11 * fi11)}
                    if ("a" != ("foo bar" + "foo bar")) {
                        v30 := GenS2{F0: v2, F1: ("\n\t" + "a")}
                        var v31 uint8 = 21
                        switch ("a" + (("a" + "hello") + "hello")) {
                        case "foo bar":
                            switch fi28 {
                            case -2147483648:
                                v32 := GenS3{F0: (v1 | 1000), F1: (!true)}
                                _ = v32
                            default:
                                if true {
                                    v33 := GenS1{F0: (("\n\t" + "") + ""), F1: "hello"}
                                    var v34 int8 = -73
                                    _ = v33
                                    _ = v34
                                }
                            }
                        }
                        _ = v30
                        _ = v31
                    }
                    var v35 bool = (genFn0(-1, 0, 146) == -2147483648)
                    _ = v29
                    _ = v35
                }
                fmt.Println(v10, v1, v2)
                switch (fi28 + (fi28 * -1)) {
                case -145:
                    if false {
                        switch ("\n\t" + "foo bar") {
                        case "foo bar":
                            if (113 <= 42) {
                                v2 = v2
                            } else {
                                v2 = (v2 & 65535)
                                var v36 int32 = -1
                                if false {
                                    v3 += (v4 & (v4 >> ((384) & 63)))
                                    v37 := GenS4{F0: 1, F1: "hello"}
                                    _ = v37
                                } else {
                                    var v38 int32 = (v36 / (((v36 ^ v36)) | 1))
                                    _ = v38
                                }
                                _ = v36
                            }
                        case "a":
                            fmt.Println(v2)
                        default:
                            v5[("\n\t" + "")] = (fi28 >> ((1) & 63))
                        }
                        var v39 uint8 = 0
                        v10 += float32(1.5)
                        _ = v39
                    }
                }
                _ = fi28
            }
        case 150:
            var v40 int = fi11
            _ = v40
        }
        fmt.Println(v3)
        _ = fi11
    }
    if (-651 <= -374) {
        v2 += (v2 - v2)
        v41 := GenS3{F0: -1, F1: (false && (1.5 <= float64(1.5)))}
        v5[""] = 32767
        _ = v41
    }
    var v42 int16 = v1
    fmt.Println(v2, v10)
    var v43 uint32 = (v2 ^ v2)
    v43 = v2
    if (1 == 0) {
        fmt.Println(v4, v2)
        switch "hello" {
        case "\n\t":
            v42 -= -66
        }
    } else {
        v44 := []int{4294967295}
        fmt.Println(v3)
        var v45 int = -1
        _ = v44
        _ = v45
    }
    v46 := []int{2147483647, 830}
    v47 := GenS5{F0: 1}
    if (((!(127 <= 127)) && (65535 == 69)) && ((!true) && true)) {
        v48 := GenS3{F0: v1, F1: (!true)}
        var v49 uint8 = 0
        _ = v48
        _ = v49
    }
    v1 += (^(v3))
    var v50 float32 = -2.25
    _ = len(v5)
    switch -1 {
    case 128:
        v46 = append(v46, 127)
    }
    var v51 uint16 = 699
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v46
    _ = v47
    _ = v5
    _ = v50
    _ = v51
}
