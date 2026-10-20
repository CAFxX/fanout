package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int64; F1 int64; F2 string }
func genFn0(p0 uint8, p1 uint8, p2 int32) uint8 {
    p0 = 128
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]int)
    var v3 float64 = 1.5
    var v4 bool = true
    var v5 int = -2147483648
    var v6 float64 = float64(3.14159)
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint = 911
        _ = fi7
        _ = v8
    }
    v5 = (v5 ^ v5)
    v4 = true
    v4 = (!(int32(v1) != 399))
    var v9 bool = (((v3 / 3.14159) + (3.14159 + v6)) > v6)
    v10 := GenS1{F0: ("\n\t" + "")}
    v5 -= -27
    if (!true) {
        var v11 uint8 = 127
        v9 = ((v5 - -1) != 128)
        _ = v11
    }
    switch "" {
    case "\n\t":
        v5 = v5
    default:
        for fi12 := 0; fi12 < 20; fi12++ {
            fmt.Println(v5, v1)
            var v13 uint64 = 0
            var v14 uint32 = 671
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    v15 := GenS1{F0: "foo bar"}
    switch -288 {
    case 32768:
        for fi16 := 0; fi16 < 2; fi16++ {
            if v9 {
                for fi17 := 0; fi17 < 7; fi17++ {
                    switch fi16 {
                    case 128:
                        v1 += -128
                    case -1:
                        v18 := []int{fi17, fi16, 0, fi16}
                        _ = v18
                    default:
                        v4 = false
                    }
                    switch -713 {
                    case -1:
                        v19 := GenS1{F0: ("hello" + "a")}
                        _ = v19
                    case 32767:
                        if true {
                            fmt.Println(v4, v4, v6)
                            v3 -= v3
                        } else {
                            v4 = (fi16 <= -9223372036854775808)
                            var v20 uint8 = 42
                            var v21 int8 = v1
                            _ = v20
                            _ = v21
                        }
                    default:
                        v2["foo bar"] = fi16
                    }
                    if (v4 || false) {
                        switch "" {
                        case "":
                            var v22 uint = 228
                            _ = v22
                        case "a":
                            var v23 string = "a"
                            _ = v23
                        }
                    }
                    _ = fi17
                }
                if ((1 == 0) || ((!true) && v9)) {
                    v1 = v1
                } else {
                    var v24 int8 = (v1 & v1)
                    _ = v24
                }
            } else {
                v3 -= v3
            }
            _ = fi16
        }
    case 127:
        var v25 int = (v5 >> ((1) & 63))
        _ = v25
    case -573:
        v1 += (v1 * v1)
    default:
        switch ("hello" + ("a" + ("\n\t" + "a"))) {
        case "a":
            v6 = float64(-2.25)
        case "\n\t":
            switch (v5 ^ v5) {
            case 550:
                if (v1 != (v1 - (v1 / ((v1) | 1)))) {
                    fmt.Println(v1)
                    var v26 string = (("\n\t" + "\n\t") + ((("foo bar" + "\n\t") + ("hello" + ("a" + "foo bar"))) + (("" + "foo bar") + "")))
                    _ = v26
                } else {
                    v27 := GenS1{F0: "foo bar"}
                    var v28 int16 = int16(v1)
                    switch "hello" {
                    case "\n\t":
                        v3 = 0.1
                    case "a":
                        v9 = (18446744073709551615 > 2)
                    }
                    _ = v27
                    _ = v28
                }
            case 534:
                if (!(v9 && (2 > 18446744073709551615))) {
                    var v29 int16 = 42
                    v30 := []int{791, v5, 42, v5}
                    delete(v2, "\n\t")
                    _ = v29
                    _ = v30
                }
            case 4294967295:
                if (!(1 >= 42)) {
                    for fi31 := 0; fi31 < 12; fi31++ {
                        v32 := GenS2{F0: (fi31 | (^(fi31))), F1: v5}
                        v33 := GenS2{F0: 2147483647, F1: (-(fi31))}
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    var v34 int64 = 0
                    for fi35 := 0; fi35 < 13; fi35++ {
                        switch "foo bar" {
                        case "foo bar":
                            v36 := GenS3{F0: 0, F1: 330}
                            _ = v36
                        case "\n\t":
                            v37 := GenS2{F0: v5, F1: -2147483648}
                            _ = v37
                        case "a":
                            v38 := GenS3{F0: 6, F1: 190}
                            _ = v38
                        }
                        v39 := GenS4{F0: (v34 ^ v34), F1: (v34 / ((v34) | 1)), F2: "a"}
                        _ = fi35
                        _ = v39
                    }
                    _ = v34
                }
            }
        default:
            v4 = (4294967295 < 587)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
