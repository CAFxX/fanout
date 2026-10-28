package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS5 struct { F0 uint64; F1 int8; F2 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string; F1 int32 }
func genFn0(p0 int16) int64 {
    v7 := []int{128, -719, -9223372036854775808, 32767}
    var v8 int32 = 266
    switch ("a" + "") {
    case "a":
        p0 = p0
    case "hello":
        v9 := GenS1{F0: 42}
        _ = v9
    }
    var v10 int16 = (p0 * 32767)
    switch 42 {
    case 1:
        v7 = append(v7, 2147483647)
    case 9223372036854775807:
        v11 := GenS2{F0: 0, F1: 0, F2: -1}
        _ = v11
    }
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    return 9223372036854775807
}
func genFn1(p0 int32, p1 uint8, p2 uint32) bool {
    p1 -= p1
    if true {
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 uint16 = 748
            _ = fi12
            _ = v13
        }
    } else {
        if (!true) {
            fmt.Println(p1)
        }
        var v14 int = -516
        v15 := []int{42, 1, (v14 % (((v14 >> ((65535) & 63))) | 1)), v14}
        _ = v14
        _ = v15
    }
    switch "a" {
    case "foo bar":
        fmt.Println(p2, p0, p1)
    default:
        switch 32767 {
        case 127:
            p2 = p2
        case -2147483648:
            for fi16 := 0; fi16 < 5; fi16++ {
                p0 = p0
                for fi17 := 0; fi17 < 6; fi17++ {
                    v18 := GenS3{F0: 0}
                    if true {
                        var v19 int8 = 42
                        v20 := []int{(fi17 / ((fi17) | 1)), -806}
                        v20 = append(v20, (fi16 + (fi17 / (((fi17 ^ 1)) | 1))))
                        _ = v19
                        _ = v20
                    }
                    p1 = p1
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        default:
            v21 := GenS4{F0: p2, F1: (!false), F2: 0}
            _ = v21
        }
    }
    switch 4294967295 {
    case -430:
        if true {
            var v22 int = 32767
            _ = v22
        }
    default:
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    var v1 int64 = -1
    v2 := make(map[string]int)
    var v3 int = 0
    v4 := []int{}
    var v5 uint8 = 48
    var v6 bool = true
    if false {
        fmt.Println(v3, v5, v3)
    }
    fmt.Println(v6, v1)
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 int32 = 1
        _ = fi23
        _ = v24
    }
    v1 = genFn0(-179)
    v4 = append(v4, v3)
    if v6 {
        fmt.Println(v3, v1, v6)
        if (-32768 >= -658) {
            v6 = false
            fmt.Println(v1)
            if (!(false || (((v6 || v6) && v6) && v6))) {
                v6 = false
            }
        }
    }
    var v25 uint = 65535
    if false {
        var v26 string = "foo bar"
        if ((v1 - v1) > v1) {
            var v27 int8 = 15
            v27 += (v27 / ((v27) | 1))
            _ = v27
        } else {
            if v6 {
                var v28 uint32 = 65535
                var v29 int = (v3 | v3)
                switch -635 {
                case -170:
                    v30 := GenS5{F0: 18446744073709551615, F1: -128, F2: 127}
                    _ = v30
                default:
                    v5 += 2
                }
                _ = v28
                _ = v29
            }
        }
        if v6 {
            if v6 {
                _ = len(v4)
            } else {
                _ = len(v4)
                fmt.Println(v25)
            }
            v2[(v26 + (v26 + (v26 + "foo bar")))] = (v3 << ((v25) & 63))
        }
        _ = v26
    } else {
        v31 := GenS6{F0: 230}
        fmt.Println(v1)
        _ = v31
    }
    fmt.Println(v25, v25, v1)
    fmt.Println(v1)
    switch v3 {
    case 43:
        v25 -= 955
    case -627:
        v5 = v5
    case -1:
        for fi32 := 0; fi32 < 11; fi32++ {
            v33 := GenS5{F0: 0, F1: 0, F2: 100}
            var v34 int8 = 100
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    v25 += (v25 - (v25 >> ((0) & 63)))
    for fi35 := 0; fi35 < 11; fi35++ {
        if v6 {
            if v6 {
                var v36 uint64 = 1
                switch fi35 {
                case 42:
                    if false {
                        for fi37 := 0; fi37 < 10; fi37++ {
                            fmt.Println(v36, v6, fi35)
                            var v38 int8 = 127
                            _ = fi37
                            _ = v38
                        }
                        var v39 float32 = (1.5 * 1e10)
                        switch ("hello" + "hello") {
                        case "foo bar":
                            v39 = -2.25
                        }
                        _ = v39
                    }
                case 4294967295:
                    var v40 int8 = -128
                    _ = v40
                case -2147483648:
                    v25 += (^(v25))
                default:
                    for fi41 := 0; fi41 < 16; fi41++ {
                        v6 = true
                        v6 = (((-2.25 / (1.5 * 1e10)) * ((0.1 * 0.0) / 0.1)) > 1e10)
                        _ = fi41
                    }
                }
                v2[""] = int(v36)
                _ = v36
            } else {
                v2[(("" + ("a" + "")) + ("a" + "a"))] = (fi35 / (((fi35 & v3)) | 1))
                switch (-(v3)) {
                case -834:
                    var v42 uint32 = 1
                    _ = v42
                case -9223372036854775808:
                    v1 -= genFn0(int16(v1))
                case 0:
                    switch v3 {
                    case -191:
                        for fi43 := 0; fi43 < 12; fi43++ {
                            var v44 uint = v25
                            v4 = append(v4, 647)
                            _ = fi43
                            _ = v44
                        }
                    case -768:
                        v5 = v5
                    }
                default:
                    for fi45 := 0; fi45 < 12; fi45++ {
                        if v6 {
                            var v46 float32 = (((0.0 * float32(0.0)) - 0.1) / ((3.14159 / 1.0) + (1.5 * ((float32(0.0) + 0.1) + 0.1))))
                            v5 = (v5 / ((v5) | 1))
                            _ = v46
                        }
                        v25 += (^(v25))
                        if (42 == 106) {
                            var v47 int16 = -32768
                            v48 := GenS4{F0: 42, F1: v6, F2: 18446744073709551615}
                            fmt.Println(v1, v6)
                            _ = v47
                            _ = v48
                        }
                        _ = fi45
                    }
                }
                var v49 float32 = float32((1e10 - float64(1e10)))
                _ = v49
            }
            var v50 int32 = -2147483648
            _ = v50
        }
        var v51 uint32 = 65535
        v51 += 2147483648
        _ = fi35
        _ = v51
    }
    switch "a" {
    case "":
        switch (v3 - v3) {
        case 42:
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                v3 += v3
            case "a":
                v52 := GenS7{F0: 1}
                _ = v52
            default:
                v53 := GenS8{F0: ((("hello" + (("hello" + "") + "hello")) + "\n\t") + "foo bar"), F1: 1}
                _ = v53
            }
        default:
            for fi54 := 0; fi54 < 8; fi54++ {
                v25 += v25
                var v55 int16 = 1000
                if (false && (v5 == v5)) {
                    v5 = (v5 + v5)
                }
                _ = fi54
                _ = v55
            }
        }
    }
    v56 := GenS7{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
}
