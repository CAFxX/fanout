package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint16 }
type GenS2 struct { F0 int; F1 int8; F2 bool }
type GenS3 struct { F0 int32; F1 string; F2 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool }
func genFn0() uint {
    var v7 int64 = 1
    v7 = 1
    var v8 uint8 = 1
    var v9 int64 = v7
    _ = v7
    _ = v8
    _ = v9
    return 207
}
func genFn1() bool {
    if (((0.0 / -2.25) / ((0.0 / 3.14159) + float32((((3.14159 - (1.5 * (3.14159 + 1.5))) / (1.5 + 0.0)) - -2.25)))) == 1e10) {
        if (!false) {
            var v10 float32 = 1.5
            _ = v10
        }
        if ((!true) && true) {
            var v11 string = ("hello" + "foo bar")
            v12 := GenS1{F0: 1, F1: v11, F2: 1}
            _ = v11
            _ = v12
        } else {
            v13 := GenS1{F0: 112, F1: "a", F2: 32767}
            var v14 uint64 = 18446744073709551615
            _ = v13
            _ = v14
        }
    } else {
        var v15 bool = (!(42 != 100))
        switch 4294967295 {
        case 32768:
            var v16 uint16 = 32767
            _ = v16
        }
        _ = v15
    }
    var v17 uint64 = 666
    var v18 bool = false
    fmt.Println(v18)
    _ = v17
    _ = v18
    return (((9223372036854775807 == -1) || (2 >= 65535)) || false)
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[int]int)
    var v3 int16 = 32767
    var v4 uint16 = 1
    v5 := []uint32{575, 2147483648, 603, 745}
    v6 := []string{"foo bar", "", "foo bar", "hello"}
    var v19 bool = (-83 == 127)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v4, v4, v19)
    case "":
        v20 := GenS2{F0: (^(v1)), F1: -1, F2: (128 == 33)}
        _ = v20
    }
    fmt.Println(v1, v19, v3)
    if v19 {
        delete(v2, (v1 ^ (v1 - (v1 >> ((18446744073709551615) & 63)))))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if (-1 >= 9223372036854775807) {
            v4 = v4
            switch ("a" + "hello") {
            case "\n\t":
                if v19 {
                    delete(v2, 9223372036854775807)
                }
            case "hello":
                for fi21 := 0; fi21 < 14; fi21++ {
                    v22 := v2[v1]
                    if ("" != "\n\t") {
                        for fi23 := 0; fi23 < 5; fi23++ {
                            v19 = (0.0 != 0.0)
                            _ = fi23
                        }
                        v22 = (v1 - -342)
                    }
                    for fi24 := 0; fi24 < 17; fi24++ {
                        v4 = 32767
                        var v25 string = (("foo bar" + (((("\n\t" + "foo bar") + "") + ("hello" + ("\n\t" + "hello"))) + ("" + "a"))) + "foo bar")
                        if (983 >= 797) {
                            v1 -= 9223372036854775807
                        } else {
                            if false {
                                var v26 int64 = -1
                                var v27 uint16 = v4
                                if false {
                                    v28 := GenS2{F0: fi21, F1: 127, F2: false}
                                    var v29 int16 = 1000
                                    v29 = 42
                                    _ = v28
                                    _ = v29
                                } else {
                                    v25 = v25
                                    if true {
                                        var v30 int16 = (v3 | (v3 ^ v3))
                                        var v31 bool = false
                                        _ = v30
                                        _ = v31
                                    } else {
                                        v3 = v3
                                        var v32 int64 = int64(v3)
                                        var v33 int8 = 42
                                        v2[1] = (^(fi21))
                                        var v34 int = fi21
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    }
                                }
                                _ = v26
                                _ = v27
                            } else {
                                v1 -= fi24
                            }
                        }
                        v22 += (v22 << ((2) & 63))
                        fmt.Println(v1, fi24)
                        _ = fi24
                        _ = v25
                    }
                    _ = fi21
                    _ = v22
                }
            case "foo bar":
                var v35 uint = 0
                _ = v35
            }
            for fi36 := 0; fi36 < 13; fi36++ {
                fmt.Println(v4, fi36, v3)
                var v37 int16 = v3
                fmt.Println(v1)
                _ = fi36
                _ = v37
            }
        }
    } else {
        var v38 uint64 = 42
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = len(v2)
        _ = v38
    }
    v39 := GenS1{F0: -22, F1: "hello", F2: 0}
    var v40 float64 = -2.25
    v41 := GenS2{F0: v1, F1: 43, F2: v19}
    switch -9223372036854775808 {
    case -2147483648:
        switch 127 {
        case 127:
            for fi42 := 0; fi42 < 20; fi42++ {
                if len(v6) > 0 {
                    v6[0] = (("\n\t" + ("\n\t" + "\n\t")) + ("\n\t" + "foo bar"))
                }
                fmt.Println(v3)
                v43 := GenS3{F0: 0, F1: "a", F2: 0}
                _ = fi42
                _ = v43
            }
        default:
            var v44 int8 = -1
            _ = v44
        }
    }
    v3 = (^(v3))
    var v45 uint32 = 2
    if v19 {
        if false {
            v4 = 632
            if v19 {
                v1 = (v1 - v1)
                if false {
                    if v19 {
                        v46 := GenS2{F0: -123, F1: 6, F2: v19}
                        _ = v46
                    } else {
                        var v47 uint32 = (v45 >> ((genFn0()) & 63))
                        _ = v47
                    }
                    if v19 {
                        var v48 uint = 0
                        v45 += (v45 - (^(v45)))
                        var v49 int32 = 2147483647
                        _ = v48
                        _ = v49
                    } else {
                        switch 2147483647 {
                        case 0:
                            var v50 int8 = 0
                            _ = v50
                        }
                        for fi51 := 0; fi51 < 8; fi51++ {
                            var v52 float64 = ((v40 + (1e10 / v40)) + 0.1)
                            _ = fi51
                            _ = v52
                        }
                        if v19 {
                            v1 = 32768
                            v53 := GenS2{F0: v1, F1: 100, F2: false}
                            _ = v53
                        }
                    }
                    for fi54 := 0; fi54 < 1; fi54++ {
                        v19 = (0 < 1)
                        _ = fi54
                    }
                } else {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v55 := GenS2{F0: int(v45), F1: 127, F2: (v45 >= (v45 + (v45 << ((0) & 63))))}
                    _ = v55
                }
                var v56 uint64 = uint64(v45)
                _ = v56
            } else {
                delete(v2, 4294967295)
            }
            var v57 uint64 = 690
            _ = v57
        } else {
            v6 = append(v6, "")
        }
        if v19 {
            v2[(v1 / (((v1 - (^(v1)))) | 1))] = -641
        }
    }
    fmt.Println(v40, v3, v40)
    for fi58 := 0; fi58 < 14; fi58++ {
        v1 = v1
        _ = fi58
    }
    for fi59 := 0; fi59 < 8; fi59++ {
        if (1 == 69) {
            var v60 int16 = -274
            if ((438 <= v4) || v19) {
                for fi61 := 0; fi61 < 13; fi61++ {
                    fmt.Println(v3)
                    _ = fi61
                }
                v62 := GenS4{F0: (v1 % ((-1) | 1))}
                v63 := GenS5{F0: v19}
                _ = v62
                _ = v63
            }
            v45 = v45
            _ = v60
        }
        _ = fi59
    }
    if len(v6) > 0 {
        v6[0] = ("foo bar" + "")
    }
    var v64 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v5
    _ = v6
    _ = v64
}
