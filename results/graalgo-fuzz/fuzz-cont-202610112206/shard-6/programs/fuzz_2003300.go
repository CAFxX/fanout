package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 int16) int16 {
    for fi7 := 0; fi7 < 1; fi7++ {
        p0 += int16(fi7)
        v8 := GenS1{F0: ((1 != 0) || (86 == 78)), F1: 4294967295}
        _ = fi7
        _ = v8
    }
    v9 := []int{-2147483648}
    v10 := GenS1{F0: (!(-114 >= 127)), F1: 2}
    var v11 int = 1
    if ((!true) && (("foo bar" < "foo bar") || false)) {
        switch "hello" {
        case "hello":
            v11 -= -9223372036854775808
        case "":
            var v12 int64 = 9223372036854775807
            _ = v12
        default:
            switch v11 {
            case -3:
                for fi13 := 0; fi13 < 20; fi13++ {
                    switch (("hello" + "hello") + ("a" + ((((("" + (("a" + ("foo bar" + ((("a" + "") + "\n\t") + "a"))) + "")) + "hello") + "foo bar") + ("\n\t" + "\n\t")) + "\n\t"))) {
                    case "":
                        fmt.Println(fi13, fi13, v11)
                    case "a":
                        switch fi13 {
                        case -935:
                            var v14 uint16 = 255
                            _ = v14
                        case 156:
                            if true {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                var v15 float32 = (0.1 * 1.5)
                                if (!true) {
                                    var v16 float32 = v15
                                    var v17 float32 = v15
                                    _ = v16
                                    _ = v17
                                }
                                _ = v15
                            }
                        }
                    }
                    _ = fi13
                }
            case 187:
                if (-108 <= -1) {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        v9 = append(v9, -9223372036854775808)
                        _ = fi18
                    }
                    var v19 bool = false
                    if false {
                        var v20 int = (v11 / (((v11 + v11)) | 1))
                        switch v20 {
                        case 0:
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        }
                        switch ("foo bar" + "foo bar") {
                        case "a":
                            switch v11 {
                            case 2147483647:
                                for fi21 := 0; fi21 < 3; fi21++ {
                                    v19 = (100 > 127)
                                    v20 = 32767
                                    if v19 {
                                        var v22 float64 = float64(-2.25)
                                        v23 := GenS1{F0: ((1e10 + 1.5) == v22), F1: 65535}
                                        _ = v22
                                        _ = v23
                                    } else {
                                        for fi24 := 0; fi24 < 18; fi24++ {
                                            var v25 uint32 = 2
                                            _ = fi24
                                            _ = v25
                                        }
                                        if (!v19) {
                                            if v19 {
                                                if (!(v11 > v11)) {
                                                    v19 = (1e10 > 0.1)
                                                }
                                            }
                                            v20 -= 127
                                        }
                                        for fi26 := 0; fi26 < 11; fi26++ {
                                            fmt.Println(fi21, v11)
                                            v20 -= 915
                                            _ = fi26
                                        }
                                    }
                                    _ = fi21
                                }
                            case -2147483648:
                                fmt.Println(v20, v19, v11)
                            case 875:
                                if len(v9) > 0 {
                                    v9[0] = (-(v20))
                                }
                            default:
                                var v27 uint32 = 42
                                _ = v27
                            }
                        case "":
                            v9 = append(v9, v20)
                        case "foo bar":
                            p0 = 42
                        }
                        _ = v20
                    } else {
                        if true {
                            if v19 {
                                for fi28 := 0; fi28 < 6; fi28++ {
                                    var v29 float64 = 3.14159
                                    v30 := GenS1{F0: true, F1: 255}
                                    fmt.Println(v19)
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                                v31 := GenS1{F0: ((42 == 2147483647) && (!(true && (v19 && false)))), F1: 1}
                                _ = v31
                            }
                        }
                        if ((-1 > -885) && v19) {
                            var v32 uint64 = 2
                            v19 = v19
                            _ = v32
                        }
                    }
                    _ = v19
                }
            default:
                var v33 uint64 = 1
                _ = v33
            }
        }
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return (p0 & (p0 >> ((4294967295) & 63)))
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 2
    var v3 int = -2147483648
    var v4 int = v3
    v5 := make(map[int]int)
    var v6 int64 = 272
    fmt.Println(v2, v3)
    v5[-15] = (v4 * (v4 >> ((255) & 63)))
    v3 = v3
    _ = len(v5)
    if false {
        v34 := GenS2{F0: (v2 / (((v2 * (v2 | (v2 | (v2 << ((519) & 63)))))) | 1)), F1: (1000 < genFn0(-32768)), F2: 68}
        v35 := GenS3{F0: 139, F1: 1, F2: -32768}
        _ = v34
        _ = v35
    } else {
        var v36 int16 = -227
        var v37 string = "a"
        _ = v36
        _ = v37
    }
    var v38 int = v3
    v39 := GenS4{F0: "foo bar", F1: (float32(((3.14159 - ((1.5 * float64((3.14159 * (0.1 - (float32((((1.5 * 1e10) - ((1e10 + (3.14159 - 1e10)) * 1.5)) + 1e10)) - 3.14159))))) - (((-2.25 - 1.5) + 0.0) / (3.14159 * 1.5)))) - ((-2.25 + ((1e10 + float64((3.14159 * 1e10))) / ((float64(-2.25) + -2.25) / 1e10))) * (0.1 * 0.0)))) == 1e10)}
    for fi40 := 0; fi40 < 7; fi40++ {
        var v41 string = "hello"
        _ = fi40
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
