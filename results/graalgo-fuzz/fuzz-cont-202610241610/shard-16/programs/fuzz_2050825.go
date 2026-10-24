package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 int16; F1 bool }
func genFn0(p0 uint, p1 string, p2 int) string {
    p2 = 4294967295
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 int = 0
            var v9 float64 = 0.1
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return ((((p1 + ("a" + (("hello" + ("foo bar" + "\n\t")) + p1))) + "\n\t") + (((p1 + p1) + p1) + ("\n\t" + p1))) + "a")
}
func genFn1(p0 int16, p1 int32) bool {
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            if false {
                switch "a" {
                case "\n\t":
                    var v12 uint32 = 238
                    _ = v12
                case "hello":
                    switch "\n\t" {
                    case "hello":
                        var v13 string = "foo bar"
                        if true {
                            fmt.Println(p0, fi11, v13)
                            var v14 int8 = 1
                            _ = v14
                        } else {
                            var v15 int8 = 1
                            _ = v15
                        }
                        _ = v13
                    case "\n\t":
                        fmt.Println(fi10)
                    case "a":
                        v16 := GenS1{F0: 32767}
                        _ = v16
                    default:
                        if false {
                            if (4294967295 != 1) {
                                fmt.Println(fi11, fi11, p0)
                            } else {
                                switch "foo bar" {
                                case "a":
                                    p0 = (p0 + 0)
                                case "foo bar":
                                    switch (fi11 >> ((0) & 63)) {
                                    case 0:
                                        var v17 uint64 = 190
                                        var v18 int = (fi11 + (fi10 | 366))
                                        _ = v17
                                        _ = v18
                                    case -2147483648:
                                        fmt.Println(fi11, fi11)
                                        var v19 int64 = 9223372036854775807
                                        _ = v19
                                    case -1:
                                        switch 9223372036854775807 {
                                        case 9223372036854775807:
                                            switch fi10 {
                                            case 215:
                                                v20 := []int{}
                                                _ = v20
                                            }
                                        case -1:
                                            var v21 uint32 = 2
                                            _ = v21
                                        default:
                                            var v22 string = ("" + "a")
                                            _ = v22
                                        }
                                    }
                                default:
                                    var v23 uint = 1
                                    _ = v23
                                }
                                fmt.Println(fi11, p0)
                            }
                            p0 -= (p0 - p0)
                        }
                    }
                case "a":
                    var v24 uint8 = 2
                    _ = v24
                }
                switch "" {
                case "a":
                    p0 += p0
                case "foo bar":
                    if (1e10 < 3.14159) {
                        switch fi10 {
                        case 1000:
                            fmt.Println(fi10, p0, p0)
                        case 32768:
                            var v25 float32 = (0.1 / (0.1 + ((-2.25 - 1e10) / float32((0.1 + 0.0)))))
                            _ = v25
                        }
                        v26 := GenS2{F0: 255, F1: fi11, F2: 1}
                        v27 := []int{-572, (fi11 / ((fi11) | 1)), 127}
                        _ = v26
                        _ = v27
                    }
                case "":
                    if false {
                        var v28 int16 = (p0 << ((4294967295) & 63))
                        v29 := GenS3{F0: 226}
                        v30 := []int{fi10, fi10}
                        _ = v28
                        _ = v29
                        _ = v30
                    } else {
                        switch ("a" + ("\n\t" + "a")) {
                        case "hello":
                            var v31 string = "hello"
                            _ = v31
                        default:
                            var v32 int16 = p0
                            _ = v32
                        }
                        if true {
                            var v33 uint8 = 243
                            switch -654 {
                            case -184:
                                v34 := GenS1{F0: (^(p0))}
                                _ = v34
                            default:
                                var v35 string = "foo bar"
                                _ = v35
                            }
                            var v36 int32 = (p1 & (p1 * int32(p0)))
                            v33 = (v33 / (((v33 ^ 127)) | 1))
                            _ = v33
                            _ = v36
                        }
                    }
                    var v37 uint = 215
                    v38 := GenS3{F0: 127}
                    var v39 int16 = p0
                    _ = v37
                    _ = v38
                    _ = v39
                default:
                    if false {
                        v40 := GenS4{F0: 100, F1: "foo bar"}
                        p1 = (p1 * (p1 * p1))
                        _ = v40
                    } else {
                        v41 := []int{0, fi11, (fi11 ^ (fi10 << ((1) & 63))), 42}
                        var v42 int64 = 1
                        _ = v41
                        _ = v42
                    }
                }
            }
            p0 = 32767
            var v43 string = "hello"
            _ = fi11
            _ = v43
        }
        var v44 uint = 1
        _ = fi10
        _ = v44
    }
    p0 += 0
    var v45 bool = true
    _ = p0
    _ = p1
    _ = v45
    return true
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []string{((((("hello" + "\n\t") + ("\n\t" + "a")) + ("a" + "")) + ((("a" + "foo bar") + ("foo bar" + "")) + ("hello" + (("a" + ("a" + "hello")) + ("a" + "foo bar"))))) + (("hello" + ("\n\t" + "")) + (("foo bar" + (((("\n\t" + "\n\t") + "foo bar") + "\n\t") + (("hello" + "\n\t") + "a"))) + "hello"))), "a", ""}
    v3 := make(map[int]bool)
    var v4 uint32 = 65535
    var v5 uint = (v1 & v1)
    v46 := GenS4{F0: -128, F1: "foo bar"}
    fmt.Println(v4, v4)
    fmt.Println(v4, v4)
    v5 = v1
    for fi47 := 0; fi47 < 17; fi47++ {
        for fi48 := 0; fi48 < 4; fi48++ {
            v2 = append(v2, "hello")
            _ = fi48
        }
        for fi49 := 0; fi49 < 20; fi49++ {
            if false {
                v50 := GenS5{F0: 42, F1: 32767}
                _ = v50
            }
            _ = fi49
        }
        _ = fi47
    }
    v5 = v5
    if (1 == 1000) {
        for fi51 := 0; fi51 < 19; fi51++ {
            var v52 bool = ("" != (("" + (("" + "foo bar") + "")) + ("hello" + ("\n\t" + (((("foo bar" + ("" + (("foo bar" + "") + ""))) + "hello") + ("a" + ("a" + ("foo bar" + "")))) + "")))))
            _ = fi51
            _ = v52
        }
        for fi53 := 0; fi53 < 18; fi53++ {
            for fi54 := 0; fi54 < 10; fi54++ {
                if (uint8(v4) != 128) {
                    fmt.Println(fi53)
                }
                var v55 uint8 = 2
                var v56 bool = true
                _ = fi54
                _ = v55
                _ = v56
            }
            fmt.Println(v4, fi53, v5)
            v1 = v5
            _ = fi53
        }
    } else {
        v57 := v3[-360]
        _ = v57
    }
    v1 = v1
    v58 := GenS6{F0: 1, F1: true}
    fmt.Println(v5, v1)
    for fi59 := 0; fi59 < 3; fi59++ {
        var v60 int8 = 0
        _ = fi59
        _ = v60
    }
    v61 := GenS1{F0: -32768}
    v5 -= (v5 ^ (v5 + 0))
    if true {
        v62 := GenS6{F0: 1, F1: (125 >= 107)}
        _ = v62
    } else {
        v63 := GenS4{F0: 94, F1: "foo bar"}
        _ = v63
    }
    v4 = (v4 << (((^(v5))) & 63))
    if (1 == -9223372036854775808) {
        _ = len(v3)
    } else {
        v64 := v3[-786]
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v58
    _ = v61
}
