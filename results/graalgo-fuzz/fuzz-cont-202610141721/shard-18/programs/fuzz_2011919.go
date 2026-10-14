package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int16; F2 string }
type GenS3 struct { F0 uint32 }
func genFn0(p0 string, p1 int16) int16 {
    switch (p0 + "a") {
    case "a":
        fmt.Println(p1, p1)
    case "foo bar":
        var v6 string = p0
        _ = v6
    case "hello":
        p0 = p0
    default:
        if false {
            var v7 uint8 = 77
            var v8 string = (p0 + "a")
            _ = v7
            _ = v8
        }
    }
    if (42 == -1) {
        var v9 int64 = 1
        var v10 float32 = -2.25
        _ = v10
        _ = v9
    } else {
        var v11 int32 = 2147483647
        v12 := []int{32768, 9223372036854775807}
        v12 = append(v12, -1)
        _ = v11
        _ = v12
    }
    v13 := []int{-1}
    var v14 int64 = -1
    var v15 uint32 = 1
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return (p1 & (p1 / (((p1 >> ((459) & 63))) | 1)))
}
func genFn1(p0 uint32, p1 uint64) uint8 {
    var v16 uint16 = 985
    var v17 uint = 255
    fmt.Println(v17)
    v16 = (v16 ^ (v16 % (((v16 % ((v16) | 1))) | 1)))
    switch "\n\t" {
    case "":
        fmt.Println(p1, p0)
    case "foo bar":
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 int64 = 0
            _ = fi18
            _ = v19
        }
    case "a":
        fmt.Println(p0)
    default:
        v20 := GenS1{F0: 100}
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 128
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 77
    v4 := make(map[uint32]bool)
    v5 := make(map[string]bool)
    fmt.Println(v3)
    var v21 float64 = ((0.1 / 0.1) * (3.14159 * 3.14159))
    v21 += v21
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 10; fi23++ {
            v24 := GenS2{F0: genFn1(2147483648, 42), F1: 692, F2: "a"}
            v25 := GenS3{F0: 827}
            _ = fi23
            _ = v24
            _ = v25
        }
        v3 = v3
        if ((0.1 / 1.0) < 3.14159) {
            v26 := GenS1{F0: 127}
            for fi27 := 0; fi27 < 7; fi27++ {
                fmt.Println(fi22, fi22)
                v28 := GenS1{F0: 100}
                v21 = v21
                _ = fi27
                _ = v28
            }
            var v29 bool = (1e10 < -2.25)
            _ = v26
            _ = v29
        } else {
            var v30 uint64 = 18446744073709551615
            if (v3 != v3) {
                v31 := []int{fi22, 127, fi22, 128}
                _ = v31
            }
            _ = v30
        }
        _ = fi22
    }
    v1[-9223372036854775808] = ""
    if true {
        var v32 float64 = 0.1
        fmt.Println(v21)
        v3 = (v3 * v3)
        _ = v32
    } else {
        if false {
            v3 -= (v3 + 120)
        }
        switch "" {
        case "\n\t":
            v33 := GenS1{F0: -128}
            _ = v33
        case "":
            if (false && false) {
                switch ("hello" + "a") {
                case "hello":
                    switch 42 {
                    case 126:
                        var v34 string = ("hello" + "hello")
                        _ = v34
                    case 128:
                        if false {
                            fmt.Println(v21, v3)
                        }
                    case 1:
                        var v35 float32 = (1e10 + 1e10)
                        _ = v35
                    default:
                        var v36 bool = (!(false || (42 == 18446744073709551615)))
                        _ = v36
                    }
                case "a":
                    v21 -= (v21 / (-2.25 / 1e10))
                }
            }
        default:
            var v37 int16 = 602
            _ = v37
        }
        if true {
            fmt.Println(v3)
            _ = len(v5)
            v3 = v3
        }
    }
    var v38 int8 = 100
    for fi39 := 0; fi39 < 7; fi39++ {
        switch fi39 {
        case 10:
            var v40 bool = true
            _ = v40
        case 0:
            switch (("" + "hello") + "foo bar") {
            case "hello":
                v41 := GenS1{F0: (^(v38))}
                _ = v41
            case "a":
                var v42 string = "hello"
                _ = v42
            case "foo bar":
                switch "hello" {
                case "":
                    var v43 int16 = -851
                    _ = v43
                }
            }
        }
        _ = fi39
    }
    if (32767 < -984) {
        switch "\n\t" {
        case "":
            if ("foo bar" != ("a" + "hello")) {
                var v44 int32 = -1
                if (true || false) {
                    var v45 uint64 = 42
                    var v46 float64 = (v21 * float64(((0.0 + 0.1) - 1.5)))
                    _ = v45
                    _ = v46
                } else {
                    v38 -= v38
                    delete(v4, 42)
                }
                v47 := GenS3{F0: 736}
                _ = v44
                _ = v47
            }
        case "hello":
            if (!(101 != v38)) {
                v48 := GenS2{F0: 22, F1: 42, F2: ""}
                switch "a" {
                case "hello":
                    v21 -= v21
                case "a":
                    var v49 string = ("\n\t" + "")
                    _ = v49
                }
                for fi50 := 0; fi50 < 12; fi50++ {
                    for fi51 := 0; fi51 < 18; fi51++ {
                        var v52 int16 = -814
                        v21 = ((v21 - (v21 * -2.25)) + -2.25)
                        fmt.Println(v21, fi51)
                        if ((^(v38)) == 1) {
                            _ = len(v2)
                            v52 = v52
                            v52 = 32767
                            v38 = (v38 % ((v38) | 1))
                        } else {
                            switch "\n\t" {
                            case "":
                                v53 := []int{(fi50 / ((fi51) | 1)), (-(fi50)), fi50, fi50}
                                var v54 int = fi50
                                if false {
                                    v3 = 42
                                    v38 += v38
                                    var v55 float32 = 0.0
                                    _ = v55
                                } else {
                                    v21 -= 1e10
                                    switch "a" {
                                    case "a":
                                        if false {
                                            fmt.Println(fi51)
                                            switch (fi51 * (fi51 << ((677) & 63))) {
                                            case 128:
                                                var v56 uint32 = 2147483648
                                                _ = v56
                                            }
                                            v21 = v21
                                        }
                                        v38 = 127
                                        if (!(v3 != v3)) {
                                            var v57 uint32 = 4294967295
                                            if len(v53) >= 2 {
                                                v53 = v53[1:len(v53)]
                                            }
                                            _ = v57
                                        } else {
                                            var v58 float64 = (v21 + v21)
                                            _ = v58
                                        }
                                    case "foo bar":
                                        fmt.Println(v54)
                                    }
                                }
                                v3 += (^(v3))
                                _ = v53
                                _ = v54
                            }
                        }
                        _ = fi51
                        _ = v52
                    }
                    _ = fi50
                }
                _ = v48
            }
        case "foo bar":
            for fi59 := 0; fi59 < 3; fi59++ {
                if false {
                    v2[1] = 340
                    v21 += -2.25
                    v38 += -1
                }
                var v60 uint16 = (v3 ^ (v3 ^ (v3 + v3)))
                for fi61 := 0; fi61 < 15; fi61++ {
                    v3 = (^(v60))
                    var v62 uint = 18446744073709551615
                    _ = fi61
                    _ = v62
                }
                _ = fi59
                _ = v60
            }
        default:
            v63 := GenS2{F0: 0, F1: genFn0("hello", -1), F2: ("foo bar" + "a")}
            _ = v63
        }
        var v64 float64 = v21
        _ = v64
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
