package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS5 struct { F0 uint32; F1 string; F2 string }
type GenS6 struct { F0 bool; F1 int32 }
func genFn0() int16 {
    var v6 float32 = (0.0 - 0.1)
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: 128, F1: (255 <= 2), F2: -2147483648}
        var v9 uint64 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int8 = -62
    if (true || (float64((float32((3.14159 + 3.14159)) * (v6 * v6))) < -2.25)) {
        var v11 uint = 65535
        _ = v11
    }
    _ = v10
    _ = v6
    return -32768
}
func genFn1() int8 {
    var v12 string = ("hello" + "a")
    fmt.Println(v12)
    switch v12 {
    case "foo bar":
        for fi13 := 0; fi13 < 19; fi13++ {
            switch 0 {
            case -906:
                for fi14 := 0; fi14 < 9; fi14++ {
                    var v15 bool = (2 != 128)
                    var v16 int16 = 42
                    v17 := []int{}
                    if v15 {
                        var v18 int32 = -267
                        var v19 string = "a"
                        fmt.Println(v18, v12)
                        _ = v18
                        _ = v19
                    } else {
                        var v20 uint64 = 0
                        _ = v20
                    }
                    v17 = append(v17, fi14)
                    v15 = (31 <= 1)
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
            }
            if ((1.5 - -2.25) < ((0.1 * 3.14159) - 3.14159)) {
                v21 := GenS1{F0: 128, F1: (!((1e10 / 0.1) <= 0.1)), F2: 1}
                _ = v21
            }
            _ = fi13
        }
    case "":
        if true {
            for fi22 := 0; fi22 < 13; fi22++ {
                for fi23 := 0; fi23 < 13; fi23++ {
                    if (44 != 0) {
                        switch "foo bar" {
                        case "hello":
                            v12 = v12
                        case "\n\t":
                            v24 := GenS1{F0: 255, F1: true, F2: 1}
                            _ = v24
                        case "":
                            var v25 int = fi23
                            _ = v25
                        }
                    }
                    var v26 int32 = 2147483647
                    _ = fi23
                    _ = v26
                }
                _ = fi22
            }
        }
    case "hello":
        v27 := []int{0, -2147483648, 9223372036854775807, 2147483647}
        _ = v27
    }
    v12 = v12
    if (false || (!((!(255 <= 500)) || true))) {
        v12 = v12
        fmt.Println(v12)
        for fi28 := 0; fi28 < 14; fi28++ {
            v29 := GenS2{F0: 430}
            _ = fi28
            _ = v29
        }
    }
    _ = v12
    return 42
}

func main() {
    var v1 float64 = 0.0
    var v2 string = "\n\t"
    v3 := make(map[string]string)
    v4 := []int{42, -9223372036854775808, 32767, 32768}
    var v5 float32 = 1e10
    var v30 uint8 = 0
    v31 := GenS3{F0: v30}
    v2 = "foo bar"
    v32 := GenS3{F0: (v30 % ((v30) | 1))}
    if (((v5 - (v5 - v5)) > ((1.5 * 3.14159) * v5)) && (false && false)) {
        v33 := GenS4{F0: 1, F1: genFn1(), F2: 2147483648}
        fmt.Println(v1, v1)
        _ = v33
    }
    if false {
        var v34 int16 = -32768
        delete(v3, ("\n\t" + ((v2 + "") + "a")))
        _ = v34
    }
    v35 := GenS1{F0: (v30 / (((^(v30))) | 1)), F1: ((genFn1() == 100) || true), F2: 0}
    v36 := GenS4{F0: 776, F1: 100, F2: 42}
    v4 = append(v4, -1)
    if true {
        v37 := GenS5{F0: 65535, F1: "hello", F2: v2}
        if len(v4) > 0 {
            v4[0] = 894
        }
        v38 := GenS5{F0: 65535, F1: "", F2: v2}
        _ = v37
        _ = v38
    } else {
        for fi39 := 0; fi39 < 19; fi39++ {
            var v40 uint64 = 1
            _ = fi39
            _ = v40
        }
        switch ("\n\t" + "hello") {
        case "foo bar":
            if (42 > v30) {
                v1 += 1.5
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    var v41 bool = (v2 < (v2 + v2))
                    for fi42 := 0; fi42 < 5; fi42++ {
                        switch "a" {
                        case "foo bar":
                            var v43 uint = 255
                            _ = v43
                        default:
                            var v44 bool = (!v41)
                            _ = v44
                        }
                        var v45 float64 = float64(3.14159)
                        _ = fi42
                        _ = v45
                    }
                    _ = v41
                }
            }
        case "a":
            var v46 uint = 1
            _ = v46
        default:
            for fi47 := 0; fi47 < 9; fi47++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if ((fi47 / (((-(fi47))) | 1)) != fi47) {
                    v2 = ""
                    v5 = (v5 * float32(0.1))
                } else {
                    v1 -= v1
                }
                for fi48 := 0; fi48 < 7; fi48++ {
                    v3[v2] = v2
                    _ = fi48
                }
                _ = fi47
            }
        }
        if false {
            v49 := GenS6{F0: false, F1: -2147483648}
            var v50 int32 = -2147483648
            for fi51 := 0; fi51 < 5; fi51++ {
                var v52 float64 = v1
                fmt.Println(fi51, v5)
                var v53 int32 = v50
                _ = fi51
                _ = v52
                _ = v53
            }
            _ = v49
            _ = v50
        }
    }
    switch (v2 + (((v2 + v2) + (v2 + (v2 + v2))) + "hello")) {
    case "foo bar":
        v54 := GenS5{F0: 541, F1: "", F2: (v2 + (v2 + v2))}
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
