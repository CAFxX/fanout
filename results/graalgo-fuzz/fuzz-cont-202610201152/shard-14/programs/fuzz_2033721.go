package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 bool; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 uint64; F1 uint32; F2 uint16 }
func genFn0(p0 uint32, p1 uint16, p2 int) uint {
    v4 := GenS1{F0: 715}
    if true {
        v5 := GenS1{F0: uint(p0)}
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 int32 = int32(p0)
            v7 = v7
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    switch p2 {
    case 128:
        var v8 bool = false
        _ = v8
    }
    p1 = 42
    for fi9 := 0; fi9 < 19; fi9++ {
        if (!false) {
            v10 := []int{0}
            _ = v10
        }
        v11 := []int{}
        _ = fi9
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 18446744073709551615
}

func main() {
    var v1 string = "a"
    var v2 string = v1
    var v3 uint16 = 32767
    v12 := []int{4294967295, 1, 9223372036854775807}
    if false {
        v12 = append(v12, 128)
    } else {
        var v13 int16 = 1000
        var v14 int8 = 94
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = v13
        _ = v14
    }
    v12 = append(v12, 128)
    var v15 int32 = 0
    v16 := GenS2{F0: v2, F1: false, F2: ((1.5 <= 0.0) || true)}
    v3 += (v3 ^ (v3 - (v3 & v3)))
    var v17 uint16 = (v3 | 2)
    if (true && (v2 >= "a")) {
        switch 32767 {
        case 0:
            v18 := GenS2{F0: v1, F1: true, F2: (v15 < (^(v15)))}
            _ = v18
        case -9223372036854775808:
            var v19 int8 = -128
            _ = v19
        }
        if len(v12) > 0 {
            v12[0] = -180
        }
    }
    v3 = (^(v3))
    switch -972 {
    case 32768:
        if true {
            v20 := GenS3{F0: 65535}
            v1 = v2
            var v21 int8 = 1
            _ = v20
            _ = v21
        } else {
            var v22 int32 = 2147483647
            _ = v22
        }
    case 42:
        v3 = (v3 << ((4294967295) & 63))
    }
    switch v2 {
    case "hello":
        if ((((-1 == 42) && (4294967295 > 2)) && true) || true) {
            switch (v2 + "a") {
            case "":
                if false {
                    _ = len(v12)
                } else {
                    v2 = (v1 + v1)
                }
            default:
                v2 = (((v1 + (v1 + v1)) + "a") + v1)
            }
            switch v2 {
            case "foo bar":
                v23 := GenS2{F0: v1, F1: false, F2: (2147483647 < 2)}
                _ = v23
            case "\n\t":
                var v24 int8 = 42
                _ = v24
            default:
                for fi25 := 0; fi25 < 10; fi25++ {
                    var v26 uint16 = 934
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    v27 := GenS4{F0: v15, F1: fi25}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
            if true {
                switch ((v2 + (v1 + ((v1 + "a") + (((v2 + (v1 + (v1 + ""))) + v2) + "\n\t")))) + "a") {
                case "foo bar":
                    v15 -= v15
                case "":
                    switch -1 {
                    case 518:
                        var v28 int16 = 1
                        _ = v28
                    case 32768:
                        if true {
                            fmt.Println(v17)
                            if true {
                                var v29 int16 = -772
                                switch v1 {
                                case "foo bar":
                                    if (v29 >= (-(v29))) {
                                        fmt.Println(v15, v2, v29)
                                        v29 = (-(v29))
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                    } else {
                                        var v30 int64 = 501
                                        v31 := GenS2{F0: ((("foo bar" + v2) + v1) + v1), F1: ((v29 ^ 1) > (v29 | v29)), F2: (!false)}
                                        if ((-2.25 / ((-2.25 * 1.5) + 3.14159)) >= (0.0 * (0.1 - -2.25))) {
                                            for fi32 := 0; fi32 < 20; fi32++ {
                                                v33 := GenS1{F0: 892}
                                                _ = fi32
                                                _ = v33
                                            }
                                        } else {
                                            var v34 string = (v1 + (v1 + "a"))
                                            _ = v34
                                        }
                                        _ = v30
                                        _ = v31
                                    }
                                default:
                                    switch "foo bar" {
                                    case "":
                                        _ = len(v12)
                                    case "a":
                                        v35 := GenS5{F0: 18446744073709551615, F1: 65535, F2: (v17 & (v3 >> ((1) & 63)))}
                                        _ = v35
                                    case "foo bar":
                                        v15 = 42
                                    }
                                }
                                var v36 uint = 4294967295
                                _ = v29
                                _ = v36
                            }
                        } else {
                            v37 := GenS5{F0: 18446744073709551615, F1: 65535, F2: v3}
                            _ = v37
                        }
                    }
                case "a":
                    var v38 float32 = -2.25
                    _ = v38
                default:
                    switch 128 {
                    case -787:
                        for fi39 := 0; fi39 < 6; fi39++ {
                            v15 = v15
                            _ = fi39
                        }
                    default:
                        var v40 int32 = v15
                        _ = v40
                    }
                }
                v3 = 366
                v1 = v2
            }
        } else {
            fmt.Println(v1)
            var v41 int32 = (v15 ^ -245)
            v42 := GenS4{F0: v41, F1: 706}
            _ = v41
            _ = v42
        }
    default:
        for fi43 := 0; fi43 < 13; fi43++ {
            v17 = (v17 + v17)
            switch (-(fi43)) {
            case 32767:
                v44 := GenS1{F0: 235}
                _ = v44
            case 127:
                if true {
                    var v45 uint = 4294967295
                    _ = v45
                } else {
                    if (-9223372036854775808 == 0) {
                        var v46 float64 = 3.14159
                        _ = v46
                    }
                    fmt.Println(v1, v15, v2)
                    v12 = append(v12, (^(fi43)))
                }
            case -678:
                _ = len(v12)
            }
            _ = fi43
        }
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    switch "\n\t" {
    case "":
        _ = len(v12)
    case "a":
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    var v47 int8 = -124
    var v48 float64 = (-2.25 + 0.1)
    for fi49 := 0; fi49 < 5; fi49++ {
        var v50 int = 2147483647
        var v51 uint16 = 65535
        for fi52 := 0; fi52 < 9; fi52++ {
            var v53 bool = true
            v48 = 0.0
            var v54 uint64 = 42
            _ = fi52
            _ = v53
            _ = v54
        }
        _ = fi49
        _ = v50
        _ = v51
    }
    if true {
        for fi55 := 0; fi55 < 1; fi55++ {
            _ = len(v12)
            var v56 uint = 77
            for fi57 := 0; fi57 < 16; fi57++ {
                var v58 bool = (0.1 == 1.5)
                _ = fi57
                _ = v58
            }
            _ = fi55
            _ = v56
        }
    } else {
        for fi59 := 0; fi59 < 11; fi59++ {
            var v60 uint64 = 321
            _ = fi59
            _ = v60
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v47
    _ = v48
}
