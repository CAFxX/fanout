package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int; F2 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint32; F2 string }
type GenS5 struct { F0 int16; F1 int32; F2 uint }
type GenS6 struct { F0 uint; F1 int8; F2 uint32 }
type GenS7 struct { F0 int64 }
func genFn0(p0 int64) uint8 {
    v7 := GenS1{F0: 0}
    switch 0 {
    case 127:
        var v8 int64 = (p0 - -209)
        _ = v8
    case 9223372036854775807:
        if false {
            var v9 uint32 = 4294967295
            v10 := GenS1{F0: 100}
            _ = v10
            _ = v9
        } else {
            v11 := GenS2{F0: -1, F1: 477, F2: false}
            _ = v11
        }
    default:
        var v12 int64 = -543
        _ = v12
    }
    var v13 string = ""
    v14 := []int{656, 205}
    var v15 int8 = 63
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    _ = v7
    return 95
}
func genFn1(p0 int8) uint16 {
    v16 := []int{32767, 9223372036854775807, 868, -913}
    var v17 int8 = p0
    for fi18 := 0; fi18 < 9; fi18++ {
        v19 := GenS3{F0: "foo bar"}
        v17 -= (v17 | 42)
        v20 := GenS4{F0: (65535 > 2), F1: 1, F2: ""}
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = p0
    _ = v16
    _ = v17
    return 862
}

func main() {
    v1 := []string{"hello", ""}
    var v2 int16 = -1
    v3 := make(map[uint32]bool)
    var v4 float32 = 1.5
    var v5 uint16 = 0
    v6 := []uint32{404, 2, 18, 776}
    var v21 uint8 = 127
    v22 := GenS4{F0: (!true), F1: 0, F2: ("\n\t" + "a")}
    v2 = v2
    if true {
        v21 = 0
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v2, v2, v2)
    v23 := GenS1{F0: 0}
    switch 127 {
    case -583:
        v2 += (v2 * v2)
    case 127:
        for fi24 := 0; fi24 < 17; fi24++ {
            var v25 int32 = 1
            _ = len(v3)
            v26 := GenS3{F0: ""}
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    var v27 string = "hello"
    v21 = genFn0(891)
    if (0 <= 2147483647) {
        var v28 uint = 572
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 float32 = 0.1
            v31 := GenS2{F0: (fi29 * fi29), F1: 127, F2: (!false)}
            if (-1 <= 0) {
                switch "\n\t" {
                case "foo bar":
                    for fi32 := 0; fi32 < 16; fi32++ {
                        switch ((v27 + ((("\n\t" + (("\n\t" + v27) + v27)) + v27) + ("\n\t" + v27))) + v27) {
                        case "foo bar":
                            v21 += v21
                        case "a":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v2 = v2
                        _ = fi32
                    }
                }
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = v28
    } else {
        switch v27 {
        case "":
            v33 := GenS5{F0: 32767, F1: 543, F2: 0}
            _ = v33
        case "\n\t":
            if false {
                fmt.Println(v4)
            }
        default:
            var v34 uint64 = 979
            _ = v34
        }
        var v35 uint32 = 4294967295
        v36 := GenS6{F0: 0, F1: -65, F2: v35}
        _ = v35
        _ = v36
    }
    switch (((v27 + ("foo bar" + (("\n\t" + v27) + (v27 + "foo bar")))) + "\n\t") + v27) {
    case "hello":
        switch (v27 + v27) {
        case "a":
            var v37 float64 = 0.1
            _ = v37
        case "hello":
            if (1.5 >= (v4 * float32(0.0))) {
                for fi38 := 0; fi38 < 2; fi38++ {
                    v4 = 0.0
                    if true {
                        var v39 int64 = -9223372036854775808
                        _ = v39
                    }
                    v40 := GenS7{F0: 230}
                    _ = fi38
                    _ = v40
                }
            } else {
                if true {
                    for fi41 := 0; fi41 < 7; fi41++ {
                        v42 := GenS6{F0: 699, F1: 42, F2: 0}
                        _ = len(v3)
                        _ = fi41
                        _ = v42
                    }
                } else {
                    for fi43 := 0; fi43 < 7; fi43++ {
                        v27 = v27
                        var v44 string = v27
                        _ = fi43
                        _ = v44
                    }
                }
                for fi45 := 0; fi45 < 8; fi45++ {
                    fmt.Println(v21, v21, v21)
                    _ = fi45
                }
                var v46 uint32 = 633
                _ = v46
            }
        }
    case "\n\t":
        v47 := v3[2147483648]
        _ = v47
    case "a":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        fmt.Println(v21, v21, v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
