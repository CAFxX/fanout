package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS2 struct { F0 int; F1 uint8; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint16; F2 int16 }
type GenS5 struct { F0 string }
func genFn0(p0 int64) uint32 {
    for fi4 := 0; fi4 < 4; fi4++ {
        if ((343 == 2147483647) && false) {
            if (255 > 4294967295) {
                v5 := GenS1{F0: 42, F1: p0, F2: 42}
                for fi6 := 0; fi6 < 15; fi6++ {
                    switch "hello" {
                    case "\n\t":
                        v7 := []int{(^(fi4)), 127}
                        _ = v7
                    case "":
                        p0 += p0
                    }
                    v8 := GenS2{F0: -398, F1: 2, F2: (("a" + "hello") + ("\n\t" + "foo bar"))}
                    p0 += (^(p0))
                    _ = fi6
                    _ = v8
                }
                for fi9 := 0; fi9 < 1; fi9++ {
                    if (true || (false || (true && false))) {
                        var v10 uint = 405
                        _ = v10
                    }
                    var v11 float32 = (-2.25 / (((0.1 - (1e10 / (3.14159 + ((3.14159 - 1e10) + float32((0.0 - (0.0 - 1e10))))))) + ((0.1 - 3.14159) * 1.5)) + 1.5))
                    v11 -= v11
                    _ = fi9
                    _ = v11
                }
                _ = v5
            }
            v12 := GenS2{F0: 128, F1: 2, F2: ("a" + ("hello" + "\n\t"))}
            _ = v12
        } else {
            if false {
                v13 := GenS1{F0: 2, F1: 42, F2: 1}
                if (128 == fi4) {
                    var v14 bool = false
                    _ = v14
                } else {
                    switch "a" {
                    case "":
                        if (42 >= 618) {
                            p0 = 1
                            if true {
                                for fi15 := 0; fi15 < 13; fi15++ {
                                    v16 := GenS1{F0: 42, F1: 1, F2: 42}
                                    var v17 uint32 = 248
                                    fmt.Println(fi15, fi4, p0)
                                    v17 += v17
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                }
                            }
                        } else {
                            v18 := GenS2{F0: fi4, F1: 128, F2: ("a" + "hello")}
                            if (42 > 42) {
                                v19 := []int{fi4, fi4, fi4, 1}
                                v20 := GenS3{F0: 42}
                                _ = v19
                                _ = v20
                            }
                            v21 := GenS1{F0: 17, F1: p0, F2: 2}
                            _ = v18
                            _ = v21
                        }
                    case "\n\t":
                        switch fi4 {
                        case 833:
                            v22 := []int{(fi4 / ((893) | 1)), -54}
                            _ = v22
                        case -641:
                            p0 -= p0
                        default:
                            for fi23 := 0; fi23 < 14; fi23++ {
                                v24 := GenS1{F0: 219, F1: p0, F2: 1}
                                v25 := []int{fi4, (^(fi4)), fi4}
                                var v26 uint16 = 255
                                _ = fi23
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                        }
                    case "foo bar":
                        if true {
                            v27 := []int{(fi4 & fi4), fi4, (fi4 % ((-2147483648) | 1)), (fi4 >> ((305) & 63))}
                            fmt.Println(p0)
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                            _ = v27
                        } else {
                            p0 -= (p0 + -1)
                        }
                    default:
                        v28 := []int{fi4, 134}
                        _ = v28
                    }
                    var v29 float64 = 0.1
                    _ = v29
                }
                _ = v13
            }
            for fi30 := 0; fi30 < 15; fi30++ {
                v31 := []int{(fi30 / ((fi30) | 1))}
                _ = fi30
                _ = v31
            }
        }
        _ = fi4
    }
    v32 := []int{}
    switch (("hello" + "foo bar") + ((((("foo bar" + ("a" + "hello")) + (("" + "") + "\n\t")) + "hello") + "\n\t") + ("foo bar" + "\n\t"))) {
    case "":
        for fi33 := 0; fi33 < 18; fi33++ {
            for fi34 := 0; fi34 < 16; fi34++ {
                p0 += p0
                _ = fi34
            }
            v35 := GenS4{F0: 18446744073709551615, F1: 65535, F2: 1000}
            _ = fi33
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        fmt.Println(p0, fi36)
        var v37 int8 = 127
        if len(v32) > 0 {
            v32[0] = int(v37)
        }
        _ = fi36
        _ = v37
    }
    var v38 int = -653
    _ = p0
    _ = v32
    _ = v38
    return 1
}
func genFn1(p0 uint16) int64 {
    p0 = p0
    fmt.Println(p0)
    for fi39 := 0; fi39 < 17; fi39++ {
        switch "\n\t" {
        case "a":
            p0 = (^(p0))
        default:
            v40 := []int{}
            _ = v40
        }
        _ = fi39
    }
    _ = p0
    return 590
}

func main() {
    v1 := []string{"a", ("\n\t" + "a")}
    var v2 int = 1
    var v3 int64 = -1
    for fi41 := 0; fi41 < 11; fi41++ {
        fmt.Println(fi41)
        var v42 uint16 = 42
        var v43 uint = 255
        _ = fi41
        _ = v42
        _ = v43
    }
    if false {
        for fi44 := 0; fi44 < 7; fi44++ {
            var v45 uint = 1
            v46 := GenS5{F0: "hello"}
            _ = fi44
            _ = v45
            _ = v46
        }
        var v47 uint = 257
        _ = v47
    }
    var v48 int16 = -335
    switch ("" + "\n\t") {
    case "hello":
        v48 = v48
    }
    v49 := GenS2{F0: (v2 >> ((65535) & 63)), F1: 2, F2: "hello"}
    fmt.Println(v48)
    switch 708 {
    case 128:
        if ((v48 / ((int16(v2)) | 1)) != v48) {
            v3 = v3
            var v50 uint16 = 32767
            _ = v50
        }
    }
    if false {
        if len(v1) > 0 {
            v1[0] = ("" + "a")
        }
        v2 = 2147483647
        var v51 uint64 = 18446744073709551615
        _ = v51
    } else {
        var v52 int16 = 1
        _ = v52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v49
}
