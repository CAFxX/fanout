package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 uint32 }
func genFn0(p0 int, p1 int16) bool {
    var v5 uint64 = 2
    var v6 int16 = 543
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return ("hello" < ("foo bar" + ("foo bar" + "\n\t")))
}
func genFn1(p0 uint, p1 uint32) uint32 {
    p1 -= 4294967295
    if false {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: 117, F1: 1000}
            var v9 int32 = 926
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v10 string = "hello"
    if false {
        v10 = v10
        if (false && ((p0 << ((1) & 63)) > p0)) {
            v11 := GenS2{F0: 42, F1: 0}
            if (true && true) {
                var v12 int8 = 42
                _ = v12
            } else {
                v10 = "hello"
                fmt.Println(p0, v10, v10)
            }
            var v13 float32 = 0.1
            _ = v11
            _ = v13
        }
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 bool = false
            _ = fi14
            _ = v15
        }
    } else {
        p1 = 633
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := GenS2{F0: -1, F1: (-(fi16))}
        if (false || false) {
            var v18 uint32 = 729
            for fi19 := 0; fi19 < 10; fi19++ {
                if true {
                    var v20 uint = (^(p0))
                    _ = v20
                } else {
                    v21 := GenS3{F0: 32767, F1: -1}
                    if ((42 == 127) && true) {
                        var v22 int = fi19
                        var v23 float64 = (float64(1e10) * (3.14159 * -2.25))
                        p1 = (v18 * v18)
                        _ = v22
                        _ = v23
                    }
                    _ = v21
                }
                if false {
                    p0 = p0
                    v24 := GenS4{F0: (!((-338 < 0) || (-2.25 != -2.25)))}
                    _ = v24
                } else {
                    var v25 bool = false
                    var v26 string = "foo bar"
                    _ = v25
                    _ = v26
                }
                v27 := []int{32768, (-(fi16)), fi16}
                v28 := GenS5{F0: p1, F1: v18}
                _ = fi19
                _ = v27
                _ = v28
            }
            _ = v18
        } else {
            var v29 string = "a"
            fmt.Println(p1, fi16)
            fmt.Println(fi16, p0)
            _ = v29
        }
        for fi30 := 0; fi30 < 13; fi30++ {
            if false {
                fmt.Println(fi30, fi30)
                var v31 int32 = 2147483647
                switch 2147483647 {
                case 950:
                    v31 = (v31 | v31)
                    if (fi30 == (fi30 << (((p0 ^ (^(p0)))) & 63))) {
                        var v32 float64 = 1e10
                        v32 = (v32 * (v32 + v32))
                        v33 := []int{(fi16 - (fi16 | (fi30 * (fi16 - fi30)))), 1, (fi16 >> (((^(p0))) & 63)), (fi30 >> ((p0) & 63))}
                        _ = v32
                        _ = v33
                    }
                }
                if false {
                    fmt.Println(v10)
                    var v34 uint64 = 699
                    v35 := GenS1{F0: 0, F1: 36}
                    _ = v34
                    _ = v35
                } else {
                    switch "a" {
                    case "a":
                        var v36 int = (fi16 << ((p0) & 63))
                        _ = v36
                    }
                }
                _ = v31
            } else {
                v10 = "a"
            }
            v10 = ("foo bar" + ("foo bar" + v10))
            _ = fi30
        }
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 255
    v3 := make(map[string]uint32)
    v4 := []int{-206}
    if (!(-1 > 0)) {
        switch ("\n\t" + ("" + "hello")) {
        case "foo bar":
            v2 -= (v2 - (v2 >> ((364) & 63)))
        case "":
            for fi37 := 0; fi37 < 12; fi37++ {
                for fi38 := 0; fi38 < 16; fi38++ {
                    var v39 int64 = 0
                    _ = fi38
                    _ = v39
                }
                fmt.Println(fi37)
                switch ("a" + ("" + "")) {
                case "foo bar":
                    fmt.Println(v2, fi37)
                case "":
                    switch ((("hello" + "foo bar") + "foo bar") + (((("hello" + (("hello" + "\n\t") + "foo bar")) + "") + ((("foo bar" + "foo bar") + "a") + "\n\t")) + ((("\n\t" + ("\n\t" + "")) + ("a" + "hello")) + "hello"))) {
                    case "\n\t":
                        switch "" {
                        case "":
                            var v40 uint32 = genFn1(0, 1)
                            _ = v40
                        default:
                            var v41 uint = 42
                            _ = v41
                        }
                    case "hello":
                        var v42 uint8 = (^(v2))
                        _ = v42
                    }
                case "\n\t":
                    _ = len(v3)
                }
                _ = fi37
            }
        }
    }
    for fi43 := 0; fi43 < 8; fi43++ {
        v2 -= v2
        _ = fi43
    }
    v44 := GenS6{F0: 715, F1: 65535}
    v2 -= 146
    switch 32767 {
    case 32768:
        v2 = 42
    case 2147483647:
        switch 2147483647 {
        case -2147483648:
            var v45 uint8 = 89
            _ = v45
        case 742:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 32767:
            v2 -= v2
        default:
            var v46 float32 = 0.0
            _ = v46
        }
    case 42:
        var v47 int16 = 42
        _ = v47
    default:
        for fi48 := 0; fi48 < 8; fi48++ {
            var v49 string = ("" + "\n\t")
            _ = fi48
            _ = v49
        }
    }
    var v50 uint64 = 2
    if len(v4) > 0 {
        v4[0] = -261
    }
    if (!(42 < 2)) {
        if false {
            for fi51 := 0; fi51 < 18; fi51++ {
                var v52 int64 = -233
                _ = fi51
                _ = v52
            }
        }
    }
    v53 := v1[-378]
    for fi54 := 0; fi54 < 17; fi54++ {
        v55 := GenS6{F0: (v50 ^ (v50 | 4294967295)), F1: 42}
        v56 := v1[(fi54 | fi54)]
        _ = fi54
        _ = v55
        _ = v56
    }
    for fi57 := 0; fi57 < 5; fi57++ {
        for fi58 := 0; fi58 < 19; fi58++ {
            var v59 uint8 = v2
            var v60 uint64 = v50
            _ = fi58
            _ = v59
            _ = v60
        }
        v53 = "\n\t"
        _ = fi57
    }
    v53 = ""
    v53 = v53
    v61 := v3[v53]
    var v62 uint64 = (v50 ^ 2)
    if false {
        v63 := GenS7{F0: 42}
        v62 = 2
        _ = v63
    }
    for fi64 := 0; fi64 < 20; fi64++ {
        for fi65 := 0; fi65 < 14; fi65++ {
            var v66 int64 = -695
            if len(v4) > 0 {
                v4[0] = 9223372036854775807
            }
            v67 := v1[(fi64 >> ((908) & 63))]
            _ = fi65
            _ = v66
            _ = v67
        }
        _ = fi64
    }
    v4 = append(v4, 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v50
    _ = v53
    _ = v61
    _ = v62
}
