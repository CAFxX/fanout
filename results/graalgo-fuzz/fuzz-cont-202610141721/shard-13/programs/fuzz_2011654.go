package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int16; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 int8; F2 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int }
func genFn0(p0 bool, p1 uint32) uint8 {
    fmt.Println(p0, p1)
    var v7 string = "a"
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float32 = 0.1
        fmt.Println(v9, v9)
        _ = fi8
        _ = v9
    }
    switch 9223372036854775807 {
    case 1:
        v10 := []int{-906}
        _ = v10
    case 0:
        v11 := []int{-935, 442}
        _ = v11
    case -2147483648:
        v12 := []int{-2147483648, int(p1)}
        _ = v12
    }
    var v13 float32 = -2.25
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return 8
}

func main() {
    v1 := []int{}
    v2 := make(map[string]int)
    v3 := []uint32{}
    var v4 uint8 = 239
    var v5 float32 = (float32(float64((0.0 / -2.25))) * float32(((3.14159 - 0.1) / float64((0.1 - 0.0)))))
    var v6 float32 = 1e10
    fmt.Println(v6, v5, v6)
    fmt.Println(v4)
    if (((("hello" + "foo bar") + "\n\t") + "") == "hello") {
        v5 -= v6
        v6 = v5
    } else {
        v6 = v5
        v14 := GenS1{F0: "a"}
        _ = v14
    }
    v15 := GenS1{F0: "\n\t"}
    v16 := GenS1{F0: "foo bar"}
    v17 := v2["\n\t"]
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2[("" + ("hello" + "hello"))] = (^(v17))
    if (!false) {
        switch "\n\t" {
        case "foo bar":
            v18 := GenS1{F0: "a"}
            _ = v18
        case "":
            fmt.Println(v17, v17)
        case "hello":
            switch v17 {
            case -788:
                v4 = (v4 << ((65535) & 63))
            }
        }
    } else {
        if (-2147483648 <= 9223372036854775807) {
            if false {
                var v19 int8 = -128
                if (0 > -278) {
                    var v20 float32 = v6
                    _ = v20
                }
                var v21 int64 = 226
                _ = v19
                _ = v21
            } else {
                v22 := v2["hello"]
                v23 := GenS2{F0: (!false), F1: 314, F2: 1}
                switch (-(v17)) {
                case 0:
                    switch (v22 + 32767) {
                    case 9223372036854775807:
                        switch 32767 {
                        case 0:
                            delete(v2, "a")
                        case 4294967295:
                            v6 = (v6 - 1.5)
                        default:
                            if false {
                                if ((!false) || (!(float32((float64((((float32(-2.25) / v5) * (((float32(((1e10 * -2.25) / 1.0)) - (0.0 / (float32((float64((v5 / v6)) + (1.5 * (1e10 + -2.25)))) - float32((float64(3.14159) * (0.0 + 1.5)))))) * v5) - (v5 / v6))) + 1.5)) / (1e10 / 1.5))) < v5))) {
                                    switch (^(v17)) {
                                    case 2147483647:
                                        v24 := GenS3{F0: 1}
                                        _ = v24
                                    }
                                    v25 := GenS1{F0: ""}
                                    _ = v25
                                } else {
                                    _ = len(v2)
                                    var v26 uint = 698
                                    for fi27 := 0; fi27 < 19; fi27++ {
                                        _ = len(v1)
                                        v28 := GenS1{F0: (("a" + "\n\t") + (("foo bar" + "") + "\n\t"))}
                                        delete(v2, ("foo bar" + (("foo bar" + "\n\t") + "hello")))
                                        _ = fi27
                                        _ = v28
                                    }
                                    _ = v26
                                }
                                v29 := v2[""]
                                _ = v29
                            }
                        }
                    case 32767:
                        v17 -= 42
                    default:
                        v17 -= v17
                    }
                case -757:
                    switch (("foo bar" + "") + "a") {
                    case "hello":
                        for fi30 := 0; fi30 < 3; fi30++ {
                            if (1e10 < 1.5) {
                                fmt.Println(v4, v5)
                                v5 = 0.0
                            } else {
                                for fi31 := 0; fi31 < 13; fi31++ {
                                    switch fi30 {
                                    case 4294967295:
                                        v32 := GenS4{F0: false, F1: -82, F2: 1}
                                        _ = v32
                                    }
                                    v1 = append(v1, (fi31 ^ fi31))
                                    _ = fi31
                                }
                                v17 = v22
                                v6 = v5
                            }
                            _ = fi30
                        }
                    case "":
                        fmt.Println(v5)
                    }
                default:
                    v33 := GenS5{F0: 1000}
                    _ = v33
                }
                _ = v22
                _ = v23
            }
            var v34 uint = 336
            switch "" {
            case "foo bar":
                var v35 float32 = 3.14159
                _ = v35
            case "hello":
                v36 := GenS2{F0: false, F1: -32768, F2: -128}
                _ = v36
            case "":
                v37 := v2["\n\t"]
                _ = v37
            default:
                var v38 float32 = float32((1.5 / -2.25))
                _ = v38
            }
            _ = v34
        } else {
            v6 = (0.0 / v6)
            for fi39 := 0; fi39 < 8; fi39++ {
                for fi40 := 0; fi40 < 2; fi40++ {
                    if false {
                        var v41 float64 = ((float64(v6) + float64(((v6 + 1e10) + float32(1.5)))) + (float64(v5) / 1e10))
                        _ = v41
                    }
                    var v42 int64 = 0
                    _ = fi40
                    _ = v42
                }
                v4 += v4
                for fi43 := 0; fi43 < 18; fi43++ {
                    v4 = (v4 / (((v4 / ((v4) | 1))) | 1))
                    _ = fi43
                }
                _ = fi39
            }
            var v44 int64 = -1
            _ = v44
        }
        v45 := GenS6{F0: false, F1: (v4 | (^(v4)))}
        var v46 uint8 = v4
        _ = v45
        _ = v46
    }
    switch -2147483648 {
    case -864:
        var v47 uint16 = 2
        _ = v47
    }
    switch ("a" + ("hello" + "\n\t")) {
    case "foo bar":
        for fi48 := 0; fi48 < 10; fi48++ {
            v4 = (^(v4))
            var v49 uint8 = (v4 | v4)
            for fi50 := 0; fi50 < 14; fi50++ {
                fmt.Println(fi48, v49)
                if ((v6 - -2.25) < ((v6 * (3.14159 + 3.14159)) * 1.5)) {
                    if false {
                        var v51 int32 = 1
                        var v52 uint = 4294967295
                        v49 = v49
                        _ = v51
                        _ = v52
                    }
                }
                _ = fi50
            }
            _ = fi48
            _ = v49
        }
    case "\n\t":
        if true {
            _ = len(v2)
            if true {
                v4 = genFn0((42 < 9223372036854775807), 610)
                v53 := GenS7{F0: v4}
                fmt.Println(v4, v4)
                _ = v53
            }
            for fi54 := 0; fi54 < 5; fi54++ {
                switch fi54 {
                case 42:
                    v55 := GenS8{F0: (fi54 & v17)}
                    _ = v55
                }
                _ = fi54
            }
        }
    case "hello":
        v6 = 1e10
    default:
        v5 += (((-2.25 - 3.14159) / v5) + (v5 - v6))
    }
    v17 = (v17 & v17)
    v6 -= (v6 + (float32((0.0 + float64(1.5))) + float32(1.5)))
    var v56 uint32 = 692
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
}
