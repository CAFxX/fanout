package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }
type GenS2 struct { F0 string; F1 bool; F2 string }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint64, p1 int8) uint32 {
    v7 := GenS1{F0: "hello", F1: (32767 > 32768), F2: 32767}
    var v8 string = ("foo bar" + (("" + "\n\t") + ("foo bar" + ("a" + "hello"))))
    v9 := []int{-304, 4294967295}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 736
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]uint32)
    var v3 bool = false
    v4 := make(map[int]string)
    v5 := []int{-1, 32767}
    var v6 uint8 = 128
    switch ("a" + "a") {
    case "hello":
        _ = len(v5)
    case "foo bar":
        v10 := GenS2{F0: "\n\t", F1: v3, F2: "\n\t"}
        _ = v10
    case "\n\t":
        if ((3.14159 >= 1e10) || v3) {
            v11 := v2[-555]
            var v12 bool = (v3 || v3)
            v6 = (v6 * 42)
            _ = v11
            _ = v12
        }
    default:
        v13 := GenS2{F0: ("" + "a"), F1: (-9223372036854775808 < 0), F2: (("\n\t" + "\n\t") + (("" + "") + ""))}
        _ = v13
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        fmt.Println(v1, fi14, v1)
        v15 := GenS2{F0: "hello", F1: v3, F2: "foo bar"}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        if (971 >= 767) {
            var v17 uint8 = v6
            switch fi16 {
            case 9223372036854775807:
                if v3 {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        if (650 >= 14) {
                            var v19 int64 = 9223372036854775807
                            fmt.Println(v19, v17, v17)
                            _ = v19
                        } else {
                            v3 = true
                            v20 := GenS3{F0: "a", F1: -1}
                            fmt.Println(v3, v3, v6)
                            _ = v20
                        }
                        _ = fi18
                    }
                    if (-731 > 42) {
                        var v21 int16 = int16(v1)
                        _ = v21
                    }
                } else {
                    v2[(fi16 * 1)] = 930
                }
            case 32768:
                var v22 uint32 = genFn0(2, -1)
                _ = v22
            case 949:
                v23 := GenS4{F0: 0, F1: v6}
                _ = v23
            }
            _ = v17
        }
        _ = fi16
    }
    delete(v4, -1)
    v6 = v6
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 int64 = 42
        if (v3 || v3) {
            for fi26 := 0; fi26 < 11; fi26++ {
                var v27 float64 = 1e10
                var v28 float64 = 0.0
                if (v25 >= v25) {
                    if v3 {
                        v27 -= ((((v27 + float64(3.14159)) - v28) + v27) / 0.1)
                        var v29 bool = true
                        _ = v29
                    }
                    var v30 int = (^(fi26))
                    _ = v30
                }
                _ = fi26
                _ = v27
                _ = v28
            }
            if (!v3) {
                var v31 int64 = int64(fi24)
                switch "\n\t" {
                case "a":
                    for fi32 := 0; fi32 < 5; fi32++ {
                        v6 = (^(v6))
                        _ = fi32
                    }
                }
                var v33 uint8 = 128
                _ = v31
                _ = v33
            } else {
                switch 0 {
                case 0:
                    v34 := v4[(fi24 % (((fi24 % ((-97) | 1))) | 1))]
                    _ = v34
                case 4294967295:
                    v6 += (v6 - v6)
                case 32768:
                    if (true || (false || true)) {
                        var v35 int64 = (^(v25))
                        var v36 float64 = (0.0 - float64((((-2.25 * 1e10) - 1.5) / 0.1)))
                        var v37 string = ((("hello" + ("foo bar" + ("" + "a"))) + "foo bar") + "a")
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                }
                if (-2147483648 <= 0) {
                    var v38 uint64 = v1
                    _ = v38
                } else {
                    v25 = -1
                }
                var v39 int16 = 932
                _ = v39
            }
        } else {
            v25 = (v25 / (((^(v25))) | 1))
            var v40 uint64 = 300
            v40 = v1
            _ = v40
        }
        _ = fi24
        _ = v25
    }
    var v41 bool = false
    switch 115 {
    case -2147483648:
        v42 := GenS5{F0: 104}
        _ = v42
    case 0:
        v6 += v6
    default:
        if ("hello" != "hello") {
            var v43 int8 = -1
            _ = v43
        }
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
