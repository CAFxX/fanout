package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint; F2 int }
func genFn0(p0 int16, p1 bool, p2 int) string {
    fmt.Println(p1)
    v7 := []int{p2, 416, (p2 << ((4294967295) & 63)), (p2 << ((18446744073709551615) & 63))}
    var v8 uint = 0
    if (p1 && (!(false && p1))) {
        for fi9 := 0; fi9 < 4; fi9++ {
            p0 -= p0
            _ = fi9
        }
        p1 = true
    }
    fmt.Println(p2, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "foo bar"
}
func genFn1(p0 int64, p1 bool) int64 {
    var v10 int64 = (p0 / (((p0 / ((p0) | 1))) | 1))
    for fi11 := 0; fi11 < 17; fi11++ {
        fmt.Println(fi11, p0)
        v12 := []int{}
        _ = fi11
        _ = v12
    }
    var v13 int32 = -2147483648
    for fi14 := 0; fi14 < 20; fi14++ {
        if (141 < 127) {
            v15 := []int{0, (fi14 ^ (-(fi14))), fi14}
            _ = v15
        } else {
            var v16 float32 = 3.14159
            _ = v16
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v13
    return (v10 % (((v10 >> ((2) & 63))) | 1))
}

func main() {
    v1 := []uint32{2147483647, 4294967295}
    var v2 int64 = 1
    var v3 int32 = 42
    var v4 uint64 = 1
    var v5 uint64 = (^(v4))
    v6 := []string{"foo bar", ("hello" + "a"), "a", ("a" + "")}
    switch "\n\t" {
    case "a":
        fmt.Println(v3, v3)
    case "foo bar":
        v17 := GenS1{F0: 0, F1: v3}
        _ = v17
    case "hello":
        for fi18 := 0; fi18 < 17; fi18++ {
            if false {
                v3 += (v3 ^ (v3 ^ -2147483648))
                _ = len(v6)
                for fi19 := 0; fi19 < 11; fi19++ {
                    v1 = append(v1, 2147483648)
                    _ = fi19
                }
            } else {
                if false {
                    if len(v1) > 0 {
                        v1[0] = 403
                    }
                    v5 = v4
                } else {
                    for fi20 := 0; fi20 < 20; fi20++ {
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                        v21 := GenS1{F0: 1000, F1: v3}
                        _ = fi20
                        _ = v21
                    }
                    var v22 int8 = 6
                    _ = v22
                }
            }
            if len(v6) > 0 {
                v6[0] = "foo bar"
            }
            switch fi18 {
            case 128:
                if (3.14159 != float32(3.14159)) {
                    v3 += v3
                    for fi23 := 0; fi23 < 15; fi23++ {
                        fmt.Println(v3, fi23)
                        switch (fi23 + fi18) {
                        case -154:
                            v24 := GenS2{F0: 18446744073709551615}
                            if false {
                                var v25 int = 128
                                var v26 uint8 = 0
                                _ = v25
                                _ = v26
                            } else {
                                fmt.Println(fi23)
                            }
                            if len(v1) > 0 {
                                v1[0] = 65535
                            }
                            _ = v24
                        case -2147483648:
                            var v27 bool = (("hello" + "\n\t") < "a")
                            _ = v27
                        }
                        _ = fi23
                    }
                    var v28 uint32 = 533
                    _ = v28
                }
            case 32768:
                var v29 uint64 = 18446744073709551615
                _ = v29
            default:
                for fi30 := 0; fi30 < 18; fi30++ {
                    var v31 uint = 4294967295
                    if ((!false) || false) {
                        switch ("hello" + "hello") {
                        case "a":
                            var v32 float32 = 1.5
                            _ = v32
                        case "\n\t":
                            switch fi18 {
                            case -853:
                                v3 -= v3
                            }
                            var v33 uint32 = 2147483647
                            _ = v33
                        default:
                            if false {
                                v34 := GenS1{F0: 1000, F1: (v3 - (v3 & 643))}
                                v35 := GenS3{F0: -927, F1: 2, F2: (fi18 & (fi18 << ((v31) & 63)))}
                                if true {
                                    if len(v6) > 0 {
                                        v6[0] = "a"
                                    }
                                    v2 -= 645
                                }
                                _ = v34
                                _ = v35
                            }
                        }
                        v6 = append(v6, ("foo bar" + "a"))
                        if true {
                            _ = len(v1)
                        }
                    } else {
                        var v36 string = ""
                        v37 := GenS1{F0: -845, F1: (v3 % (((-(v3))) | 1))}
                        _ = v36
                        _ = v37
                    }
                    _ = fi30
                    _ = v31
                }
            }
            _ = fi18
        }
    }
    v38 := GenS2{F0: 1}
    v5 = v5
    var v39 string = (("foo bar" + "foo bar") + "\n\t")
    _ = len(v6)
    fmt.Println(v2, v2)
    v3 -= v3
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (!(false && (42 > genFn1(genFn1(v2, (true || (!true))), (0.1 != 1.5))))) {
        var v40 uint64 = 18446744073709551615
        _ = v40
    }
    var v41 bool = (!(1 > 1))
    if true {
        switch -916 {
        case 707:
            fmt.Println(v5, v2)
        case 9223372036854775807:
            fmt.Println(v2, v39)
        }
    }
    fmt.Println(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
