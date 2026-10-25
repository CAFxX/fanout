package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
func genFn0(p0 int64, p1 string, p2 string) int32 {
    fmt.Println(p2, p1, p2)
    p1 = "\n\t"
    p2 = p2
    if false {
        var v5 bool = (1 <= 449)
        v6 := GenS1{F0: 42, F1: 2147483647}
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: 42, F1: fi7}
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := []int{128, fi9}
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        _ = v5
        _ = v6
    }
    switch ((p2 + p2) + p1) {
    case "hello":
        switch ((("\n\t" + p1) + (p2 + (p1 + "foo bar"))) + p1) {
        case "":
            if true {
                var v11 float32 = (0.1 + 0.0)
                _ = v11
            } else {
                switch 0 {
                case 127:
                    switch p2 {
                    case "a":
                        if true {
                            fmt.Println(p2, p0, p1)
                        }
                    }
                }
                v12 := []int{}
                _ = v12
            }
        case "foo bar":
            if (!false) {
                if (!true) {
                    if false {
                        if false {
                            v13 := []int{32768, 32768}
                            switch 42 {
                            case -2147483648:
                                v14 := GenS1{F0: -1, F1: -837}
                                _ = v14
                            case -9223372036854775808:
                                p0 += p0
                            }
                            _ = v13
                        } else {
                            var v15 float32 = 1e10
                            v16 := GenS1{F0: 2147483647, F1: 349}
                            _ = v15
                            _ = v16
                        }
                    }
                }
            } else {
                if (!false) {
                    switch (p1 + p2) {
                    case "":
                        var v17 int = 575
                        _ = v17
                    default:
                        v18 := GenS2{F0: (p0 | (p0 / (((p0 - (p0 & p0))) | 1)))}
                        _ = v18
                    }
                }
            }
        default:
            if true {
                for fi19 := 0; fi19 < 7; fi19++ {
                    p0 += (p0 << ((65535) & 63))
                    for fi20 := 0; fi20 < 17; fi20++ {
                        p2 = (p2 + "")
                        _ = fi20
                    }
                    p1 = p2
                    _ = fi19
                }
                v21 := []int{136, 2147483647}
                _ = v21
            } else {
                var v22 string = p1
                v23 := GenS2{F0: p0}
                for fi24 := 0; fi24 < 7; fi24++ {
                    if true {
                        var v25 int32 = 1
                        _ = v25
                    } else {
                        for fi26 := 0; fi26 < 12; fi26++ {
                            var v27 int32 = -1
                            v22 = v22
                            _ = fi26
                            _ = v27
                        }
                    }
                    for fi28 := 0; fi28 < 4; fi28++ {
                        p0 = (p0 ^ (p0 & (p0 & int64(fi24))))
                        _ = fi28
                    }
                    _ = fi24
                }
                _ = v22
                _ = v23
            }
        }
    case "foo bar":
        if true {
            v29 := GenS1{F0: 2147483647, F1: -810}
            var v30 float32 = float32(1.5)
            var v31 int8 = 42
            _ = v29
            _ = v30
            _ = v31
        }
    case "a":
        fmt.Println(p0)
    default:
        if (!true) {
            var v32 string = p2
            var v33 int8 = 1
            p1 = "a"
            _ = v32
            _ = v33
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -225
}
func genFn1(p0 uint, p1 string, p2 uint) int {
    v34 := GenS2{F0: 622}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    return 127
}

func main() {
    var v1 int8 = -108
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    var v4 int32 = 2147483647
    var v35 float64 = (1.5 * 1.5)
    fmt.Println(v1)
    delete(v2, "hello")
    fmt.Println(v1, v35, v4)
    var v36 uint = 18446744073709551615
    v37 := v3[""]
    v38 := v3["a"]
    var v39 bool = (false && true)
    v40 := GenS3{F0: "foo bar"}
    if (("" + "foo bar") < "foo bar") {
        if ((1.5 * (0.1 * 0.1)) > 1e10) {
            if v39 {
                var v41 float64 = v35
                _ = v41
            }
            var v42 string = "hello"
            fmt.Println(v4, v39, v38)
            _ = v42
        }
        if false {
            fmt.Println(v38)
            v43 := []int{42, (v37 + (v38 % (((v38 ^ (-(v38)))) | 1)))}
            fmt.Println(v1)
            _ = v43
        } else {
            if v39 {
                v2[("foo bar" + (("\n\t" + ("a" + (("foo bar" + ((("a" + ("" + "a")) + ("hello" + ("\n\t" + "\n\t"))) + "hello")) + ("hello" + ("" + "a"))))) + "\n\t"))] = 749
                fmt.Println(v36, v38)
                switch v37 {
                case 9223372036854775807:
                    v36 = v36
                }
            } else {
                v36 = 2
                for fi44 := 0; fi44 < 9; fi44++ {
                    for fi45 := 0; fi45 < 20; fi45++ {
                        var v46 uint16 = 1
                        var v47 uint8 = 127
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    v48 := GenS3{F0: "\n\t"}
                    var v49 int = genFn1((v36 + v36), "hello", v36)
                    _ = fi44
                    _ = v48
                    _ = v49
                }
                fmt.Println(v38, v35, v39)
            }
            for fi50 := 0; fi50 < 12; fi50++ {
                fmt.Println(v39, v36)
                if (v39 || true) {
                    v51 := GenS4{F0: v39}
                    v36 += 4294967295
                    _ = v51
                } else {
                    var v52 uint32 = 65535
                    v38 += int(v36)
                    v39 = ((v36 << ((65535) & 63)) < v36)
                    _ = v52
                }
                for fi53 := 0; fi53 < 8; fi53++ {
                    var v54 uint = v36
                    if (2147483648 != 65535) {
                        var v55 int = genFn1(v54, "\n\t", 0)
                        if (!v39) {
                            v38 = (v55 | v55)
                        }
                        _ = v55
                    } else {
                        var v56 string = ""
                        switch fi50 {
                        case 0:
                            v4 = genFn0(0, "", ((("hello" + "foo bar") + ((v56 + v56) + v56)) + (v56 + v56)))
                        case -658:
                            var v57 uint32 = 65535
                            _ = v57
                        default:
                            fmt.Println(v4)
                        }
                        v58 := GenS1{F0: genFn0(42, "foo bar", v56), F1: (v38 << (((v54 << (((^(v54))) & 63))) & 63))}
                        _ = v56
                        _ = v58
                    }
                    if v39 {
                        fmt.Println(v1, v1)
                        var v59 float64 = (1e10 * v35)
                        var v60 int32 = -620
                        var v61 int = -781
                        _ = v59
                        _ = v60
                        _ = v61
                    } else {
                        v35 += v35
                        if (-289 < 1) {
                            fmt.Println(fi53, v37, v54)
                            var v62 int16 = 32767
                            _ = v62
                        }
                        v63 := []int{}
                        v38 += v37
                        v64 := GenS2{F0: 0}
                        _ = v63
                        _ = v64
                    }
                    _ = fi53
                    _ = v54
                }
                _ = fi50
            }
        }
    }
    var v65 uint8 = 86
    delete(v2, "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v65
}
