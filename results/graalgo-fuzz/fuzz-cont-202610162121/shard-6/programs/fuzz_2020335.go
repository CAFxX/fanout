package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 string }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int32 }
func genFn0(p0 bool, p1 int64, p2 uint64) int8 {
    v7 := []int{9223372036854775807}
    if p0 {
        p0 = (!p0)
        var v8 float64 = -2.25
        _ = v8
    }
    p2 = (p2 % (((p2 & (p2 + (^(p2))))) | 1))
    if (42 == (p1 * (p1 | 42))) {
        if (!(true || true)) {
            for fi9 := 0; fi9 < 1; fi9++ {
                p1 += p1
                _ = fi9
            }
            fmt.Println(p1)
            switch (("a" + (("foo bar" + "a") + (("" + "a") + "foo bar"))) + (("foo bar" + "hello") + "foo bar")) {
            case "a":
                var v10 uint32 = 42
                _ = v10
            case "\n\t":
                var v11 uint32 = 480
                _ = v11
            case "foo bar":
                var v12 bool = (336 == 153)
                _ = v12
            }
        } else {
            _ = len(v7)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 int8, p1 int16) string {
    for fi13 := 0; fi13 < 8; fi13++ {
        fmt.Println(fi13, p0)
        _ = fi13
    }
    p0 += p0
    if false {
        var v14 float64 = 0.1
        var v15 uint = 164
        _ = v14
        _ = v15
    } else {
        fmt.Println(p0)
        if (4294967295 > 42) {
            p0 -= genFn0((false && (p1 < p1)), -371, 18446744073709551615)
            var v16 int16 = -1
            if (false || true) {
                if false {
                    v17 := GenS1{F0: -198, F1: 128, F2: "foo bar"}
                    _ = v17
                }
            }
            _ = v16
        } else {
            var v18 uint8 = 128
            v19 := []int{1}
            _ = v18
            _ = v19
        }
    }
    _ = p0
    _ = p1
    return ("hello" + "a")
}

func main() {
    var v1 string = "a"
    var v2 uint = 0
    var v3 uint64 = 4294967295
    var v4 uint = v2
    var v5 bool = true
    var v6 float64 = 0.1
    for fi20 := 0; fi20 < 11; fi20++ {
        switch (fi20 ^ fi20) {
        case 128:
            v21 := GenS1{F0: 1, F1: 42, F2: v1}
            _ = v21
        case 0:
            for fi22 := 0; fi22 < 15; fi22++ {
                switch v1 {
                case "hello":
                    v23 := []int{fi22, fi22}
                    _ = v23
                case "\n\t":
                    switch (fi20 ^ (fi22 + fi22)) {
                    case 802:
                        v24 := []int{(fi20 | -9223372036854775808), -9223372036854775808, 128, fi22}
                        _ = v24
                    case -9223372036854775808:
                        v1 = v1
                    case 477:
                        v6 -= v6
                    }
                case "":
                    if false {
                        v25 := GenS2{F0: (v3 - 0), F1: 42}
                        v3 = (^(v3))
                        switch fi20 {
                        case 128:
                            v26 := GenS1{F0: 0, F1: 0, F2: v1}
                            _ = v26
                        case 32768:
                            switch 42 {
                            case -73:
                                if ((!true) || false) {
                                    fmt.Println(v4, v3)
                                } else {
                                    v27 := []int{fi22, fi22}
                                    v1 = ((("hello" + v1) + "\n\t") + "\n\t")
                                    _ = v27
                                }
                                v28 := []int{}
                                _ = v28
                            default:
                                v29 := []int{}
                                _ = v29
                            }
                        default:
                            v2 -= 4294967295
                        }
                        _ = v25
                    }
                }
                v3 -= (v3 >> ((691) & 63))
                var v30 bool = (!v5)
                var v31 uint32 = 2
                v32 := []int{(fi22 & (fi22 << ((uint(v3)) & 63))), fi20, -968, (fi20 & (fi22 ^ (fi22 + fi22)))}
                _ = fi22
                _ = v30
                _ = v31
                _ = v32
            }
        }
        v33 := []int{}
        _ = fi20
        _ = v33
    }
    var v34 uint32 = 2147483647
    if v5 {
        var v35 string = "foo bar"
        var v36 int8 = 1
        _ = v35
        _ = v36
    } else {
        var v37 uint = (v4 ^ (^(v4)))
        v38 := GenS2{F0: 18446744073709551615, F1: 74}
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 7; fi39++ {
        for fi40 := 0; fi40 < 12; fi40++ {
            var v41 string = (((v1 + "\n\t") + "hello") + "a")
            _ = fi40
            _ = v41
        }
        v6 = float64(1.5)
        v42 := []int{(fi39 / ((128) | 1))}
        _ = fi39
        _ = v42
    }
    v34 -= 568
    fmt.Println(v1, v6, v2)
    v43 := []int{int(v3), 32767, 42}
    for fi44 := 0; fi44 < 7; fi44++ {
        if len(v43) >= 2 {
            v43 = v43[1:len(v43)]
        }
        fmt.Println(v4, v1)
        v4 = (v4 / ((v2) | 1))
        _ = fi44
    }
    if (((2147483648 >= 2147483648) && ((v1 != "\n\t") || true)) && true) {
        v5 = false
        v5 = v5
    }
    for fi45 := 0; fi45 < 19; fi45++ {
        v1 = (v1 + (("" + "hello") + (("hello" + "a") + v1)))
        var v46 string = v1
        _ = fi45
        _ = v46
    }
    if (!v5) {
        if v5 {
            var v47 uint16 = 0
            fmt.Println(v6)
            v4 += (v2 + (v2 % (((v4 & 42)) | 1)))
            _ = v47
        } else {
            var v48 int = 32767
            fmt.Println(v6, v48)
            _ = v48
        }
        v49 := GenS1{F0: 42, F1: 206, F2: (v1 + "\n\t")}
        _ = v49
    }
    fmt.Println(v5, v3)
    var v50 uint16 = 1
    for fi51 := 0; fi51 < 17; fi51++ {
        v52 := GenS3{F0: 1000, F1: (v3 & v3)}
        _ = fi51
        _ = v52
    }
    v1 = "\n\t"
    v53 := GenS4{F0: 65535, F1: 151, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v43
    _ = v5
    _ = v50
    _ = v53
    _ = v6
}
