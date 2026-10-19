package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 bool }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int64; F1 int8; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint; F1 int16; F2 bool }
func genFn0() int32 {
    switch 128 {
    case 294:
        var v4 float32 = 0.1
        _ = v4
    case 1:
        v5 := GenS1{F0: 829, F1: (false || true), F2: true}
        _ = v5
    }
    var v6 uint8 = 127
    fmt.Println(v6)
    for fi7 := 0; fi7 < 10; fi7++ {
        if (!false) {
            var v8 string = ("" + "foo bar")
            v9 := GenS1{F0: 41, F1: true, F2: true}
            _ = v8
            _ = v9
        }
        v10 := GenS2{F0: false, F1: "foo bar"}
        _ = fi7
        _ = v10
    }
    v11 := []int{}
    _ = v11
    _ = v6
    return 715
}
func genFn1(p0 uint32, p1 uint64, p2 int16) int64 {
    v12 := GenS3{F0: 1, F1: 42, F2: p2}
    switch "" {
    case "foo bar":
        var v13 int16 = (p2 >> ((4294967295) & 63))
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{388, 301}
    v3 := []int{0, -9223372036854775808, 0, -671}
    switch "hello" {
    case "a":
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 int16 = 1
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 int = fi14
                v18 := GenS4{F0: 349, F1: 0}
                if true {
                    if len(v2) > 0 {
                        v2[0] = 0
                    }
                    fmt.Println(v15)
                    var v19 int16 = 950
                    _ = v19
                } else {
                    switch "foo bar" {
                    case "":
                        if false {
                            v15 = v15
                        }
                    case "a":
                        var v20 int8 = 42
                        _ = v20
                    case "\n\t":
                        switch ("" + ("\n\t" + "hello")) {
                        case "a":
                            var v21 int16 = v15
                            _ = v21
                        }
                    }
                    v22 := GenS5{F0: 796, F1: (("\n\t" + (("a" + "a") + (("\n\t" + "\n\t") + ""))) + ("\n\t" + "\n\t"))}
                    _ = v22
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi14
            _ = v15
        }
    case "foo bar":
        for fi23 := 0; fi23 < 6; fi23++ {
            switch (("foo bar" + "foo bar") + "foo bar") {
            case "a":
                v1[94] = "a"
            case "":
                for fi24 := 0; fi24 < 5; fi24++ {
                    if (3.14159 >= 0.1) {
                        if (32767 > 32767) {
                            v2 = append(v2, 42)
                            var v25 uint16 = 679
                            _ = v25
                        }
                        if (!(true || (653 == 59))) {
                            var v26 string = "a"
                            if ((1.5 - 1e10) > 0.1) {
                                var v27 float32 = 0.1
                                _ = v27
                            }
                            v26 = "a"
                            var v28 uint = 65535
                            switch 128 {
                            case -9223372036854775808:
                                if false {
                                    var v29 int = -149
                                    _ = v29
                                } else {
                                    switch "foo bar" {
                                    case "foo bar":
                                        fmt.Println(v26)
                                    default:
                                        var v30 bool = (true && (genFn0() == 911))
                                        _ = v30
                                    }
                                }
                            case 323:
                                var v31 float64 = 0.1
                                _ = v31
                            case -1:
                                switch (fi23 % (((fi23 & fi23)) | 1)) {
                                case 42:
                                    v32 := GenS1{F0: 553, F1: (2 > 2), F2: (true && (!false))}
                                    _ = v32
                                case 127:
                                    fmt.Println(v26, v26)
                                case 0:
                                    fmt.Println(v28)
                                }
                            }
                            v33 := GenS3{F0: -599, F1: 42, F2: 1000}
                            _ = v26
                            _ = v28
                            _ = v33
                        } else {
                            if false {
                                var v34 int8 = 0
                                _ = v34
                            } else {
                                var v35 int8 = -68
                                if (!(0.0 > (((float32(((1.5 + (((1.5 - float64(3.14159)) - (0.0 / 1.5)) / 1.0)) / -2.25)) / -2.25) * 3.14159) - (3.14159 / 3.14159)))) {
                                    var v36 uint = 0
                                    var v37 int64 = 1
                                    v38 := GenS6{F0: 91}
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                } else {
                                    fmt.Println(fi23, fi23)
                                    v35 = v35
                                    if true {
                                        var v39 uint16 = 0
                                        _ = v39
                                    }
                                }
                                _ = v35
                            }
                            if len(v3) > 0 {
                                v3[0] = 128
                            }
                            if (!(42 == 625)) {
                                fmt.Println(fi23, fi23)
                                _ = len(v2)
                            }
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        if false {
                            var v40 int8 = 1
                            v40 -= v40
                            _ = v40
                        }
                        var v41 uint8 = 171
                        _ = v41
                    }
                    fmt.Println(fi23, fi24)
                    if false {
                        delete(v1, 834)
                        _ = len(v2)
                    }
                    fmt.Println(fi24)
                    _ = fi24
                }
            }
            _ = fi23
        }
    case "\n\t":
        v42 := GenS6{F0: -1}
        _ = v42
    default:
        var v43 int64 = genFn1(2147483648, 2, -32768)
        _ = v43
    }
    var v44 uint8 = 22
    fmt.Println(v44)
    var v45 uint32 = 2147483647
    var v46 int = -9223372036854775808
    if false {
        v47 := GenS7{F0: 65535, F1: 922, F2: (837 >= 65535)}
        v3 = append(v3, (v46 & 2147483647))
        v45 = (^(v45))
        _ = v47
    } else {
        var v48 float64 = 0.1
        v48 += v48
        var v49 int32 = 0
        _ = v48
        _ = v49
    }
    v3 = append(v3, v46)
    v50 := GenS3{F0: 424, F1: 38, F2: 1000}
    _ = len(v1)
    v46 = -632
    v51 := GenS3{F0: 0, F1: 42, F2: -716}
    var v52 float32 = -2.25
    var v53 bool = (127 == 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v50
    _ = v51
    _ = v52
    _ = v53
}
