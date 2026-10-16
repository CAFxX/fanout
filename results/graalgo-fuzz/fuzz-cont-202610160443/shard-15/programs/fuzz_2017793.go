package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 string }
type GenS2 struct { F0 uint64; F1 string; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 int; F2 int8 }
func genFn0(p0 int, p1 int, p2 int32) uint8 {
    v5 := GenS1{F0: (true || (p2 > p2)), F1: "foo bar", F2: "hello"}
    var v6 float32 = 1.5
    switch p0 {
    case 276:
        p0 -= (p0 | p1)
    default:
        var v7 int8 = 42
        _ = v7
    }
    v6 += 1e10
    if (259 > 42) {
        switch (p1 << ((2) & 63)) {
        case 128:
            v8 := []int{786, (p1 & (^(p0))), (p1 | 1), 32767}
            _ = v8
        case 32767:
            v9 := GenS1{F0: (v6 > v6), F1: (("a" + "") + ("hello" + "")), F2: "a"}
            _ = v9
        default:
            p1 = -633
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 uint16 = 778
    var v2 uint8 = 134
    v3 := []uint32{}
    v4 := make(map[int]uint32)
    v10 := GenS1{F0: true, F1: ("\n\t" + "hello"), F2: ""}
    switch "\n\t" {
    case "a":
        v1 = v1
    case "\n\t":
        v1 += v1
    default:
        _ = len(v4)
    }
    _ = len(v4)
    var v11 string = ("hello" + "\n\t")
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS2{F0: 2, F1: "hello", F2: -583}
        for fi14 := 0; fi14 < 13; fi14++ {
            v1 -= v1
            var v15 int = (^(fi12))
            if (!(false && true)) {
                v16 := GenS1{F0: (-848 < 0), F1: "", F2: v11}
                if true {
                    switch (-(v15)) {
                    case -2147483648:
                        var v17 uint = 573
                        _ = v17
                    case 1:
                        fmt.Println(v11, v2, v2)
                    }
                } else {
                    fmt.Println(v15, fi14, fi12)
                    _ = len(v3)
                    v2 += v2
                }
                if (1 < 4294967295) {
                    var v18 bool = (!(!(!false)))
                    _ = v18
                }
                switch v15 {
                case 432:
                    var v19 int32 = -2147483648
                    _ = v19
                }
                v1 -= (v1 - v1)
                var v20 int16 = -1
                if (0.0 == -2.25) {
                    v1 = (^(v1))
                    if (2147483647 > fi12) {
                        if false {
                            fmt.Println(v1, v2, fi12)
                        } else {
                            var v21 uint64 = 905
                            fmt.Println(fi14, v1)
                            v22 := GenS2{F0: v21, F1: "", F2: (v15 % ((45) | 1))}
                            _ = v21
                            _ = v22
                        }
                        var v23 int64 = 42
                        fmt.Println(fi12)
                        _ = v23
                    } else {
                        var v24 int8 = -1
                        _ = v24
                    }
                    var v25 int32 = 1
                    _ = v25
                }
                _ = v16
                _ = v20
            }
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v2 = v2
    if (1 != 0) {
        switch -2147483648 {
        case 0:
            var v26 float64 = 1.5
            _ = v26
        case 223:
            for fi27 := 0; fi27 < 1; fi27++ {
                for fi28 := 0; fi28 < 16; fi28++ {
                    fmt.Println(v2)
                    v4[(fi28 + fi27)] = 1
                    v29 := GenS3{F0: v1}
                    _ = fi28
                    _ = v29
                }
                var v30 uint8 = genFn0(-9223372036854775808, (-(fi27)), -1)
                _ = len(v3)
                _ = fi27
                _ = v30
            }
        }
        var v31 string = "\n\t"
        _ = v31
    } else {
        var v32 int64 = 9223372036854775807
        if (!((v1 / ((v1) | 1)) > 218)) {
            v33 := GenS1{F0: true, F1: v11, F2: ""}
            _ = v33
        } else {
            v11 = (v11 + v11)
            if (("hello" + ("a" + (v11 + "\n\t"))) > ("" + (v11 + v11))) {
                switch 0 {
                case 4294967295:
                    fmt.Println(v2, v11, v11)
                default:
                    var v34 int8 = 100
                    _ = v34
                }
            }
            v2 -= genFn0(-682, 349, 42)
        }
        for fi35 := 0; fi35 < 6; fi35++ {
            fmt.Println(v1, v11, v2)
            _ = fi35
        }
        _ = v32
    }
    fmt.Println(v2)
    var v36 bool = true
    fmt.Println(v11, v1, v1)
    for fi37 := 0; fi37 < 14; fi37++ {
        _ = len(v4)
        if (100 <= 127) {
            var v38 float32 = 0.0
            _ = v38
        } else {
            v39 := GenS4{F0: -44, F1: -9223372036854775808, F2: 1}
            switch "" {
            case "":
                fmt.Println(v11, fi37, v2)
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "hello":
                for fi40 := 0; fi40 < 7; fi40++ {
                    v36 = ((("\n\t" + (v11 + ("\n\t" + "\n\t"))) + "\n\t") < "\n\t")
                    _ = fi40
                }
            default:
                for fi41 := 0; fi41 < 18; fi41++ {
                    var v42 uint = 1
                    if v36 {
                        var v43 float64 = 3.14159
                        v4[fi37] = 1
                        _ = v43
                    }
                    v44 := GenS4{F0: -128, F1: -2147483648, F2: 100}
                    fmt.Println(v42, fi37, fi41)
                    _ = fi41
                    _ = v42
                    _ = v44
                }
            }
            _ = v39
        }
        if ((2 < 0) && v36) {
            if (0 <= (^(v2))) {
                switch "hello" {
                case "a":
                    for fi45 := 0; fi45 < 2; fi45++ {
                        var v46 int = 32768
                        _ = fi45
                        _ = v46
                    }
                }
                if false {
                    var v47 uint = 4294967295
                    v48 := GenS3{F0: (v1 - (v1 - (v1 | uint16(v2))))}
                    _ = v47
                    _ = v48
                } else {
                    fmt.Println(v36, v1)
                    v2 += genFn0(fi37, 0, 42)
                }
            }
            for fi49 := 0; fi49 < 7; fi49++ {
                var v50 bool = false
                _ = fi49
                _ = v50
            }
        } else {
            if v36 {
                var v51 int16 = -667
                _ = v51
            }
            for fi52 := 0; fi52 < 7; fi52++ {
                v4[751] = 2
                _ = fi52
            }
            v1 += v1
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
