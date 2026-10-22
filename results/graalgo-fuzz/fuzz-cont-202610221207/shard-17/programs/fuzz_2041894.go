package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 bool; F2 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 bool }
func genFn0() int {
    var v6 int16 = 1000
    v7 := GenS1{F0: 32767, F1: 1000, F2: 127}
    fmt.Println(v6)
    v8 := GenS1{F0: 4294967295, F1: (v6 ^ (v6 >> ((42) & 63))), F2: 0}
    if true {
        if true {
            v6 = v6
        } else {
            for fi9 := 0; fi9 < 20; fi9++ {
                for fi10 := 0; fi10 < 9; fi10++ {
                    switch "foo bar" {
                    case "a":
                        if false {
                            var v11 int = fi10
                            var v12 string = ((("" + "") + (("a" + "a") + "foo bar")) + "foo bar")
                            _ = v11
                            _ = v12
                        }
                    case "foo bar":
                        v13 := GenS1{F0: -972, F1: (-(v6)), F2: 0}
                        v14 := []int{(fi9 ^ fi9), fi9, fi10}
                        v14 = append(v14, (fi9 >> ((65535) & 63)))
                        _ = v13
                        _ = v14
                    case "\n\t":
                        fmt.Println(v6, fi9)
                    default:
                        var v15 float64 = 0.0
                        _ = v15
                    }
                    v16 := []int{fi10, fi9, (-(fi10))}
                    v17 := GenS1{F0: (fi9 ^ (fi10 - (fi10 ^ fi10))), F1: (v6 % ((629) | 1)), F2: 128}
                    var v18 string = ("" + ("hello" + ""))
                    var v19 int8 = -83
                    var v20 uint = 42
                    _ = fi10
                    _ = v16
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                var v21 int8 = -128
                for fi22 := 0; fi22 < 19; fi22++ {
                    if ((162 >= 82) || false) {
                        var v23 int8 = (v21 * v21)
                        v6 = (v6 * 42)
                        if false {
                            var v24 string = ""
                            v25 := []int{(fi22 >> ((2) & 63)), (-(fi9)), fi22}
                            v23 += 127
                            v23 = -1
                            _ = v24
                            _ = v25
                        }
                        fmt.Println(v21, v6, v23)
                        _ = v23
                    }
                    _ = fi22
                }
                _ = fi9
                _ = v21
            }
            var v26 int = -147
            _ = v26
        }
    } else {
        var v27 float64 = 0.1
        _ = v27
    }
    _ = v6
    _ = v7
    _ = v8
    return int(v6)
}

func main() {
    var v1 string = (("" + "foo bar") + "a")
    var v2 float32 = -2.25
    var v3 float64 = float64(float32(-2.25))
    v4 := []string{}
    var v5 uint = 170
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 int8 = 0
        var v30 uint32 = 4294967295
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 uint16 = 65535
    v32 := GenS2{F0: 714, F1: 42}
    if (2 > 0) {
        for fi33 := 0; fi33 < 11; fi33++ {
            v34 := GenS1{F0: (fi33 / (((fi33 >> (((v5 % ((18446744073709551615) | 1))) & 63))) | 1)), F1: 754, F2: 2}
            _ = fi33
            _ = v34
        }
        fmt.Println(v2)
    } else {
        for fi35 := 0; fi35 < 18; fi35++ {
            switch (fi35 + fi35) {
            case 92:
                switch "foo bar" {
                case "":
                    var v36 int32 = 369
                    _ = v36
                case "a":
                    switch ("\n\t" + (((v1 + ("a" + v1)) + "\n\t") + "foo bar")) {
                    case "foo bar":
                        var v37 bool = true
                        _ = v37
                    case "a":
                        var v38 uint8 = 128
                        _ = v38
                    case "\n\t":
                        var v39 bool = false
                        _ = v39
                    }
                }
            }
            for fi40 := 0; fi40 < 11; fi40++ {
                v41 := GenS3{F0: 127}
                if (true && false) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    v31 = 2
                    fmt.Println(v1, v31)
                    switch 990 {
                    case 160:
                        switch genFn0() {
                        case 0:
                            var v42 float32 = v2
                            _ = v42
                        }
                    }
                }
                _ = fi40
                _ = v41
            }
            var v43 uint8 = 128
            _ = fi35
            _ = v43
        }
        if (-1 == 42) {
            if (1 <= 65535) {
                v2 = v2
            } else {
                var v44 uint32 = 128
                for fi45 := 0; fi45 < 15; fi45++ {
                    for fi46 := 0; fi46 < 2; fi46++ {
                        switch genFn0() {
                        case 536:
                            v3 += v3
                        case -2147483648:
                            _ = len(v4)
                        case 128:
                            v47 := GenS4{F0: -1, F1: (-1 <= 702), F2: 291}
                            _ = v47
                        }
                        v1 = ((v1 + (v1 + "\n\t")) + (v1 + (v1 + (v1 + v1))))
                        var v48 string = v1
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v3 = ((0.1 * 1.5) + v3)
                        if (194 != 119) {
                            v2 -= v2
                            var v49 float64 = (v3 - float64(v2))
                            _ = v49
                        }
                        _ = fi46
                        _ = v48
                    }
                    _ = fi45
                }
                v1 = v1
                _ = v44
            }
        } else {
            v50 := GenS5{F0: 0}
            _ = v50
        }
    }
    v51 := GenS6{F0: -9223372036854775808, F1: -9223372036854775808}
    v52 := GenS7{F0: true}
    for fi53 := 0; fi53 < 17; fi53++ {
        v3 = 0.1
        _ = fi53
    }
    var v54 int32 = -2147483648
    var v55 float64 = (1.5 / v3)
    v5 = (^(v5))
    v55 -= (((v3 / v3) - v55) / v55)
    for fi56 := 0; fi56 < 3; fi56++ {
        v3 = (v3 + (((((v3 * (v55 / (v3 - (1.5 + (v55 - 1.5))))) / 1e10) - 3.14159) / 1e10) * float64(v2)))
        _ = fi56
    }
    var v57 bool = true
    var v58 uint32 = 2147483648
    v54 = v54
    v59 := GenS3{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v54
    _ = v55
    _ = v57
    _ = v58
    _ = v59
}
