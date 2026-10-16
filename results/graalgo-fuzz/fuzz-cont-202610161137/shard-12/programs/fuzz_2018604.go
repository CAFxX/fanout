package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int8; F1 string; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string }
func genFn0(p0 int32) int16 {
    fmt.Println(p0)
    v4 := []int{}
    var v5 string = ("a" + "hello")
    _ = p0
    _ = v4
    _ = v5
    return -32768
}
func genFn1(p0 uint8, p1 uint8) int32 {
    var v6 int8 = -128
    if (true || true) {
        p0 -= (p1 - (^(p1)))
    }
    v6 -= v6
    switch "foo bar" {
    case "\n\t":
        var v7 float32 = (-2.25 / 1e10)
        _ = v7
    case "":
        v8 := GenS1{F0: 42, F1: ("foo bar" + "\n\t"), F2: 0}
        _ = v8
    case "foo bar":
        switch "" {
        case "a":
            for fi9 := 0; fi9 < 10; fi9++ {
                fmt.Println(p1, p0, fi9)
                _ = fi9
            }
        case "foo bar":
            v6 = 1
        case "\n\t":
            var v10 int = -1
            _ = v10
        default:
            if false {
                if (2147483647 != 32768) {
                    for fi11 := 0; fi11 < 3; fi11++ {
                        switch "a" {
                        case "a":
                            fmt.Println(p0, p0, p0)
                        case "":
                            if (!false) {
                                if false {
                                    if (42 >= -817) {
                                        if false {
                                            var v12 int64 = -1
                                            p0 = p0
                                            _ = v12
                                        }
                                    }
                                    p0 = 143
                                }
                            } else {
                                for fi13 := 0; fi13 < 17; fi13++ {
                                    v6 += (^(v6))
                                    fmt.Println(p1)
                                    fmt.Println(fi13, fi13, v6)
                                    _ = fi13
                                }
                                switch ("foo bar" + (("" + "foo bar") + (("foo bar" + "foo bar") + "\n\t"))) {
                                case "hello":
                                    v6 = (v6 - (v6 / ((v6) | 1)))
                                case "":
                                    v14 := []int{fi11, 32767}
                                    _ = v14
                                case "a":
                                    v15 := GenS1{F0: 4294967295, F1: (("" + "a") + ""), F2: -1}
                                    _ = v15
                                }
                            }
                        default:
                            v16 := []int{(fi11 & -62)}
                            _ = v16
                        }
                        var v17 float32 = (((0.1 / 3.14159) * (0.1 * 1.5)) + (1e10 * 0.0))
                        v18 := []int{(fi11 | (fi11 % ((fi11) | 1)))}
                        _ = fi11
                        _ = v17
                        _ = v18
                    }
                } else {
                    p1 = p1
                }
                fmt.Println(p1, v6, v6)
            } else {
                for fi19 := 0; fi19 < 1; fi19++ {
                    fmt.Println(v6, p1)
                    _ = fi19
                }
                for fi20 := 0; fi20 < 16; fi20++ {
                    v21 := GenS1{F0: 540, F1: "", F2: 2147483647}
                    var v22 int8 = (v6 * (v6 << ((18446744073709551615) & 63)))
                    v6 -= (v22 + v22)
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                for fi23 := 0; fi23 < 8; fi23++ {
                    var v24 string = ""
                    var v25 int = 300
                    v26 := []int{(fi23 - -772), fi23}
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
        }
    default:
        var v27 uint64 = 2
        _ = v27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        v29 := GenS1{F0: 18446744073709551615, F1: "\n\t", F2: -152}
        v30 := GenS2{F0: 2, F1: 2}
        _ = fi28
        _ = v29
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v6
    return -2147483648
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]int)
    var v3 uint32 = 823
    v31 := GenS1{F0: 42, F1: ("hello" + "\n\t"), F2: 1}
    if true {
        if (false || (true || true)) {
            switch 589 {
            case -321:
                v1 = uint(v3)
            }
        }
    } else {
        v32 := GenS3{F0: -9223372036854775808, F1: 32767}
        fmt.Println(v3, v1)
        var v33 string = "a"
        _ = v32
        _ = v33
    }
    v34 := []int{2147483647, 272, -1, -766}
    v1 -= 42
    switch 32768 {
    case 128:
        v1 -= v1
    }
    var v35 int16 = 32767
    v1 -= (^(v1))
    switch "hello" {
    case "\n\t":
        v36 := GenS4{F0: 24, F1: (("a" + ("\n\t" + (((("\n\t" + "a") + "hello") + (("\n\t" + "hello") + (("hello" + "") + "foo bar"))) + "hello"))) + ("a" + (("" + (("foo bar" + ("a" + "a")) + (((("\n\t" + "foo bar") + "") + "foo bar") + "\n\t"))) + ((("\n\t" + "foo bar") + (((("" + "") + (("" + "\n\t") + (("" + "hello") + (("" + ("" + ("hello" + ("a" + "foo bar")))) + "foo bar")))) + "foo bar") + "")) + "a")))), F2: 290}
        _ = v36
    case "":
        v37 := v2["foo bar"]
        _ = v37
    case "hello":
        var v38 uint64 = 2
        _ = v38
    }
    if (((v35 & (^(v35))) > 42) || (true && (!(!(0.0 > 1e10))))) {
        switch 285 {
        case 42:
            v39 := GenS5{F0: (true && true)}
            _ = v39
        case 4294967295:
            var v40 int16 = v35
            _ = v40
        case -701:
            for fi41 := 0; fi41 < 20; fi41++ {
                _ = len(v34)
                var v42 int8 = 127
                v1 -= (v1 << (((v1 & (v1 | v1))) & 63))
                _ = fi41
                _ = v42
            }
        default:
            var v43 uint = v1
            _ = v43
        }
    } else {
        switch 1 {
        case 42:
            if (1 == 0) {
                var v44 int8 = 127
                v44 = -123
                _ = v44
            }
        case 4294967295:
            var v45 uint = v1
            _ = v45
        default:
            var v46 string = ("foo bar" + ("" + "a"))
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 1; fi47++ {
        var v48 uint = (v1 >> (((v1 >> ((42) & 63))) & 63))
        var v49 int16 = 1000
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
        _ = fi47
        _ = v48
        _ = v49
    }
    v3 -= v3
    for fi50 := 0; fi50 < 20; fi50++ {
        for fi51 := 0; fi51 < 7; fi51++ {
            v52 := GenS6{F0: ""}
            var v53 uint8 = 121
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
            _ = fi51
            _ = v52
            _ = v53
        }
        _ = fi50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
}
