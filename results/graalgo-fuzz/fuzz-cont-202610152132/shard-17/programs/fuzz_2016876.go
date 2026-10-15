package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint64) string {
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 = 2
        _ = fi7
    }
    var v8 int64 = 9223372036854775807
    var v9 bool = (false || (42 != 42))
    _ = p0
    _ = v8
    _ = v9
    return ("foo bar" + (((("a" + "a") + "hello") + "a") + ""))
}

func main() {
    var v1 int16 = 737
    var v2 string = "foo bar"
    var v3 float32 = 3.14159
    v4 := make(map[string]uint32)
    var v5 int = 4294967295
    var v6 int8 = 42
    for fi10 := 0; fi10 < 2; fi10++ {
        v11 := GenS1{F0: "hello", F1: (v5 * v5), F2: (v6 - v6)}
        v12 := GenS2{F0: v1}
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            fmt.Println(fi13)
            if (42 >= 0) {
                fmt.Println(fi14, fi13, v2)
                if false {
                    v15 := GenS3{F0: (v6 >> ((65535) & 63)), F1: v5}
                    v16 := GenS1{F0: v2, F1: fi13, F2: (v6 % ((48) | 1))}
                    v2 = ("foo bar" + ("" + ((v2 + v2) + (((v2 + ("foo bar" + v2)) + v2) + (v2 + (v2 + (v2 + (("foo bar" + v2) + v2))))))))
                    _ = v15
                    _ = v16
                }
                v5 -= (fi13 >> ((65535) & 63))
            } else {
                if (false || (42 <= 42)) {
                    v17 := GenS4{F0: 2, F1: 127, F2: -2147483648}
                    _ = v17
                } else {
                    var v18 string = ""
                    _ = v18
                }
                v3 -= v3
                v19 := []int{(v5 | v5), fi13}
                _ = v19
            }
            _ = fi14
        }
        v20 := GenS5{F0: 349}
        _ = fi13
        _ = v20
    }
    v21 := GenS6{F0: 4294967295, F1: 9223372036854775807}
    if (0.1 != (1.5 * 1e10)) {
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 bool = false
            if (!(302 < 0)) {
                for fi24 := 0; fi24 < 7; fi24++ {
                    fmt.Println(v1, v6)
                    var v25 int32 = 2147483647
                    var v26 uint8 = 255
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                switch 211 {
                case 127:
                    var v27 int = 263
                    _ = v27
                case -2147483648:
                    if v23 {
                        v28 := GenS7{F0: 655, F1: v6}
                        var v29 int32 = 621
                        v30 := []int{}
                        _ = v28
                        _ = v29
                        _ = v30
                    } else {
                        v31 := []int{2147483647, fi22}
                        fmt.Println(v5, v23)
                        _ = v31
                    }
                }
            } else {
                var v32 int = (fi22 * v5)
                var v33 float32 = v3
                _ = v32
                _ = v33
            }
            _ = fi22
            _ = v23
        }
    } else {
        if false {
            switch v2 {
            case "foo bar":
                v2 = (v2 + v2)
            default:
                v34 := []int{2147483647, (v5 & v5), v5}
                _ = v34
            }
        }
        var v35 float32 = float32(float64(v3))
        v36 := []int{v5}
        _ = v35
        _ = v36
    }
    if (!(128 == (v5 % ((v5) | 1)))) {
        v6 = (v6 << ((4294967295) & 63))
        if ((v1 ^ (^(v1))) < 32767) {
            var v37 int8 = (v6 + 1)
            v1 -= 32767
            v1 = v1
            _ = v37
        }
        v1 += v1
    } else {
        for fi38 := 0; fi38 < 2; fi38++ {
            v39 := GenS5{F0: 600}
            switch "a" {
            case "a":
                if true {
                    var v40 int64 = -827
                    v2 = v2
                    var v41 int64 = (v40 << ((1) & 63))
                    _ = v40
                    _ = v41
                } else {
                    v42 := GenS6{F0: (v5 >> ((18446744073709551615) & 63)), F1: -1}
                    var v43 float32 = v3
                    var v44 float32 = (v43 * v43)
                    _ = v42
                    _ = v43
                    _ = v44
                }
            case "":
                var v45 bool = (!((v5 + v5) < (fi38 % ((fi38) | 1))))
                _ = v45
            case "foo bar":
                if (949 <= 42) {
                    v6 = (v6 ^ (v6 & (v6 - v6)))
                } else {
                    v46 := GenS8{F0: 2, F1: 42}
                    fmt.Println(v3, fi38, v1)
                    _ = v46
                }
            }
            switch 4294967295 {
            case 9223372036854775807:
                v47 := []int{(v5 + v5), 2147483647, (fi38 * (fi38 << ((0) & 63)))}
                _ = v47
            }
            _ = fi38
            _ = v39
        }
        v3 -= (0.1 - v3)
        for fi48 := 0; fi48 < 14; fi48++ {
            v1 = (v1 << ((954) & 63))
            v5 = (v5 * (v5 >> ((0) & 63)))
            _ = fi48
        }
    }
    for fi49 := 0; fi49 < 1; fi49++ {
        var v50 int8 = (v6 & (v6 * (v6 ^ (v6 ^ -36))))
        for fi51 := 0; fi51 < 3; fi51++ {
            v52 := GenS2{F0: -32768}
            var v53 int32 = 42
            v6 = v6
            v54 := []int{}
            v55 := GenS8{F0: 523, F1: 42}
            _ = fi51
            _ = v52
            _ = v53
            _ = v54
            _ = v55
        }
        v6 -= (v50 ^ 42)
        _ = fi49
        _ = v50
    }
    v56 := []int{}
    for fi57 := 0; fi57 < 17; fi57++ {
        if true {
            v1 = 906
        } else {
            for fi58 := 0; fi58 < 9; fi58++ {
                v6 -= v6
                var v59 string = "a"
                v59 = v2
                _ = fi58
                _ = v59
            }
            v60 := GenS9{F0: 42, F1: 1}
            fmt.Println(v2, v3)
            _ = v60
        }
        for fi61 := 0; fi61 < 17; fi61++ {
            var v62 int8 = (v6 - (v6 ^ v6))
            var v63 uint32 = 65535
            if ((-(fi57)) < -2147483648) {
                var v64 uint16 = 255
                v56 = append(v56, (v5 ^ fi61))
                _ = v64
            }
            _ = fi61
            _ = v62
            _ = v63
        }
        _ = fi57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
}
