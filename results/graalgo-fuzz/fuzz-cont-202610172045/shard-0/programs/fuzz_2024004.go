package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 int; F1 int16; F2 uint8 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 int16 }
func genFn0() int32 {
    var v4 uint32 = 2147483648
    v4 = v4
    for fi5 := 0; fi5 < 4; fi5++ {
        switch (fi5 + (-(fi5))) {
        case 32767:
            for fi6 := 0; fi6 < 9; fi6++ {
                v7 := []int{(fi6 % ((fi6) | 1)), 32767, 127}
                switch "a" {
                case "":
                    v7 = append(v7, (fi5 | fi6))
                case "a":
                    var v8 uint64 = 0
                    _ = v8
                case "hello":
                    var v9 float64 = ((-2.25 * float64((float32(-2.25) - 3.14159))) - 3.14159)
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
        default:
            switch "a" {
            case "hello":
                v4 = 618
            case "foo bar":
                for fi10 := 0; fi10 < 8; fi10++ {
                    v11 := []int{(fi5 % ((fi5) | 1)), (fi10 ^ (fi5 / (((fi5 >> ((18446744073709551615) & 63))) | 1))), fi5}
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    var v12 int = (fi10 / (((fi5 * (fi5 + fi10))) | 1))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "a":
                if ("" > "foo bar") {
                    v13 := GenS1{F0: false}
                    v4 += v4
                    _ = v13
                }
            default:
                var v14 int32 = 993
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := GenS2{F0: v4, F1: int8(v4)}
            _ = fi15
            _ = v16
        }
        _ = fi5
    }
    _ = v4
    return 42
}
func genFn1(p0 int32, p1 uint64, p2 uint32) string {
    var v17 int8 = -128
    for fi18 := 0; fi18 < 12; fi18++ {
        switch "foo bar" {
        case "\n\t":
            var v19 uint32 = (p2 - p2)
            _ = v19
        case "foo bar":
            switch 127 {
            case 0:
                var v20 uint16 = 1
                _ = v20
            case 312:
                for fi21 := 0; fi21 < 2; fi21++ {
                    switch "" {
                    case "a":
                        v22 := GenS2{F0: 2147483647, F1: v17}
                        v23 := GenS3{F0: 15, F1: 810}
                        _ = v22
                        _ = v23
                    case "foo bar":
                        if ((!true) && (true && false)) {
                            var v24 uint64 = p1
                            var v25 uint8 = 2
                            switch (fi21 * (fi18 * (^(fi18)))) {
                            case 9223372036854775807:
                                if (v24 >= p1) {
                                    v26 := GenS1{F0: (("\n\t" + ((("a" + "\n\t") + (((("a" + ("\n\t" + (("\n\t" + "") + ("hello" + "foo bar")))) + "\n\t") + "foo bar") + "\n\t")) + ("" + ""))) < "")}
                                    if true {
                                        fmt.Println(fi21, v24)
                                    } else {
                                        v25 += uint8(p2)
                                        v25 = (v25 - v25)
                                    }
                                    _ = v26
                                }
                                var v27 int32 = p0
                                _ = v27
                            default:
                                v28 := GenS3{F0: v25, F1: 1}
                                _ = v28
                            }
                            _ = v24
                            _ = v25
                        }
                    case "\n\t":
                        v29 := []int{(fi18 >> ((659) & 63)), 1, (-(fi18)), (fi21 << ((0) & 63))}
                        _ = v29
                    }
                    v30 := []int{fi21, fi18, (-(fi21))}
                    v31 := GenS1{F0: true}
                    _ = fi21
                    _ = v30
                    _ = v31
                }
            }
        case "hello":
            v32 := []int{(fi18 << ((4294967295) & 63)), (fi18 & (-(fi18))), (fi18 - fi18)}
            _ = v32
        }
        _ = fi18
    }
    v33 := GenS1{F0: ((!false) || true)}
    v34 := []int{128, 1}
    switch 4294967295 {
    case 128:
        var v35 float64 = 3.14159
        _ = v35
    case -128:
        var v36 uint8 = 255
        _ = v36
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v33
    _ = v34
    return "hello"
}

func main() {
    var v1 uint64 = 1
    var v2 int16 = 42
    v3 := make(map[int]string)
    for fi37 := 0; fi37 < 1; fi37++ {
        v2 = v2
        for fi38 := 0; fi38 < 7; fi38++ {
            v39 := GenS2{F0: 670, F1: 49}
            _ = fi38
            _ = v39
        }
        _ = fi37
    }
    v2 = -182
    v40 := []int{9223372036854775807, 0}
    if (-9223372036854775808 == -1) {
        fmt.Println(v2)
        v41 := GenS1{F0: (true && true)}
        _ = v41
    }
    if len(v40) >= 2 {
        v40 = v40[1:len(v40)]
    }
    v42 := GenS4{F0: "a", F1: 127}
    switch 42 {
    case 2147483647:
        v43 := GenS5{F0: -2147483648, F1: 34, F2: 128}
        _ = v43
    case 128:
        var v44 uint16 = 2
        _ = v44
    default:
        var v45 int8 = 0
        _ = v45
    }
    v46 := v3[-1]
    fmt.Println(v46, v1, v46)
    var v47 int64 = -9223372036854775808
    var v48 uint16 = 65535
    if true {
        var v49 uint8 = uint8(v48)
        _ = v49
    }
    v48 = v48
    for fi50 := 0; fi50 < 6; fi50++ {
        v51 := GenS2{F0: 0, F1: -103}
        v2 += (v2 % ((v2) | 1))
        v1 -= v1
        _ = fi50
        _ = v51
    }
    if false {
        var v52 uint8 = 2
        v53 := GenS2{F0: 604, F1: 0}
        v54 := GenS6{F0: 4294967295, F1: 404}
        _ = v52
        _ = v53
        _ = v54
    }
    v47 += 42
    v55 := GenS7{F0: (v47 & (-(v47))), F1: (2 >= v48)}
    v56 := GenS8{F0: -590}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v42
    _ = v46
    _ = v47
    _ = v48
    _ = v55
    _ = v56
}
