package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int16 }
func genFn0() uint {
    for fi6 := 0; fi6 < 5; fi6++ {
        if (858 <= 0) {
            var v7 uint = 4294967295
            if true {
                if false {
                    var v8 float32 = (0.1 * 0.0)
                    for fi9 := 0; fi9 < 11; fi9++ {
                        v8 = (v8 + (v8 + v8))
                        v8 = (1e10 * (v8 * (0.1 / (3.14159 + (v8 + 3.14159)))))
                        v10 := GenS1{F0: v7}
                        _ = fi9
                        _ = v10
                    }
                    _ = v8
                }
                for fi11 := 0; fi11 < 8; fi11++ {
                    v7 = 352
                    v7 = v7
                    if false {
                        if true {
                            v7 = (v7 + v7)
                            var v12 uint32 = 296
                            _ = v12
                        }
                        v7 += (v7 | (v7 >> ((v7) & 63)))
                        var v13 float64 = 3.14159
                        _ = v13
                    }
                    _ = fi11
                }
                switch fi6 {
                case 2147483647:
                    var v14 int16 = -585
                    _ = v14
                case 1:
                    v15 := []int{}
                    _ = v15
                case -1:
                    switch "hello" {
                    case "\n\t":
                        switch ("foo bar" + "") {
                        case "foo bar":
                            var v16 int8 = 1
                            _ = v16
                        case "":
                            switch "a" {
                            case "\n\t":
                                fmt.Println(fi6, v7)
                            case "foo bar":
                                switch ("" + "\n\t") {
                                case "a":
                                    fmt.Println(v7, fi6)
                                case "\n\t":
                                    switch -255 {
                                    case 9223372036854775807:
                                        v7 += 1
                                    case -93:
                                        var v17 int8 = -128
                                        _ = v17
                                    case -1:
                                        var v18 uint16 = 32767
                                        _ = v18
                                    }
                                default:
                                    var v19 int32 = -1
                                    _ = v19
                                }
                            default:
                                v20 := []int{fi6, (-(fi6)), fi6, -144}
                                _ = v20
                            }
                        case "hello":
                            var v21 bool = (0 == -736)
                            _ = v21
                        default:
                            v7 += v7
                        }
                    default:
                        v22 := GenS1{F0: v7}
                        _ = v22
                    }
                }
            }
            fmt.Println(v7, fi6)
            _ = v7
        } else {
            fmt.Println(fi6)
            var v23 int8 = 100
            var v24 int16 = 19
            _ = v23
            _ = v24
        }
        v25 := GenS2{F0: -931}
        var v26 bool = true
        _ = fi6
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 uint8 = 144
            var v30 float64 = 3.14159
            v29 = 255
            _ = fi28
            _ = v29
            _ = v30
        }
        v31 := []int{(^(fi27))}
        _ = fi27
        _ = v31
    }
    v32 := GenS2{F0: 42}
    if true {
        var v33 uint32 = 972
        _ = v33
    }
    var v34 int32 = 1
    _ = v32
    _ = v34
    return 1
}

func main() {
    var v1 uint8 = 1
    var v2 int16 = -32768
    v3 := []uint32{726, 52, 0}
    var v4 string = ("a" + "hello")
    var v5 int32 = 42
    for fi35 := 0; fi35 < 16; fi35++ {
        v36 := GenS3{F0: v2}
        var v37 int16 = (^(v2))
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 float64 = ((3.14159 + (((0.1 + ((float64((((-2.25 - 0.1) + (0.1 * (1e10 + (0.1 - ((1e10 + -2.25) - 0.0))))) / 0.1)) * 0.0) - (3.14159 / ((1e10 / 3.14159) * float64(-2.25))))) * 1e10) * (((float64(1.5) - (1e10 - 3.14159)) - 1.5) / 1e10))) / (float64(0.1) + 1e10))
    for fi39 := 0; fi39 < 2; fi39++ {
        switch ("\n\t" + v4) {
        case "hello":
            var v40 bool = (!(!(((fi39 - fi39) > fi39) && (false && false))))
            _ = v40
        case "foo bar":
            switch v4 {
            case "\n\t":
                v41 := GenS2{F0: 42}
                _ = v41
            case "foo bar":
                v2 -= (-(v2))
            case "":
                v42 := GenS3{F0: v2}
                _ = v42
            default:
                v43 := GenS2{F0: -1}
                _ = v43
            }
        }
        _ = fi39
    }
    for fi44 := 0; fi44 < 10; fi44++ {
        v45 := GenS1{F0: 42}
        v2 = v2
        _ = fi44
        _ = v45
    }
    switch (v4 + v4) {
    case "hello":
        switch ("" + v4) {
        case "\n\t":
            v2 = (-(v2))
        }
    case "":
        switch "a" {
        case "a":
            if true {
                v46 := GenS4{F0: 9223372036854775807, F1: (v2 - v2)}
                v47 := GenS1{F0: 54}
                _ = v46
                _ = v47
            }
        case "hello":
            v2 = (v2 % ((v2) | 1))
        case "\n\t":
            var v48 uint16 = 32767
            _ = v48
        default:
            var v49 uint16 = 0
            _ = v49
        }
    }
    fmt.Println(v1, v38, v38)
    v5 += v5
    for fi50 := 0; fi50 < 13; fi50++ {
        if false {
            v5 = 0
            var v51 uint32 = 2147483647
            switch (fi50 - fi50) {
            case -33:
                v52 := GenS3{F0: v2}
                _ = v52
            case 2147483647:
                _ = len(v3)
            case -9223372036854775808:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                for fi53 := 0; fi53 < 1; fi53++ {
                    var v54 uint32 = v51
                    switch "foo bar" {
                    case "hello":
                        v55 := GenS2{F0: -9223372036854775808}
                        _ = v55
                    default:
                        v3 = append(v3, 1)
                    }
                    _ = fi53
                    _ = v54
                }
            }
            _ = v51
        } else {
            for fi56 := 0; fi56 < 6; fi56++ {
                if (v5 > (v5 | -102)) {
                    fmt.Println(fi50, v4, v38)
                }
                v57 := GenS3{F0: -905}
                _ = fi56
                _ = v57
            }
            v2 += v2
        }
        _ = fi50
    }
    v58 := GenS4{F0: -260, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v58
}
