package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int16; F2 int32 }
type GenS5 struct { F0 int16 }
func genFn0(p0 int16, p1 int8) uint32 {
    if true {
        p0 = (p0 * (p0 * p0))
    }
    p0 = (^(p0))
    v5 := GenS1{F0: 1, F1: 759, F2: 0}
    var v6 int64 = -876
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 string) bool {
    fmt.Println(p0)
    var v7 uint64 = 119
    if true {
        var v8 float64 = 1.5
        v9 := GenS1{F0: 207, F1: 660, F2: 42}
        for fi10 := 0; fi10 < 5; fi10++ {
            for fi11 := 0; fi11 < 4; fi11++ {
                fmt.Println(v7, p0)
                _ = fi11
            }
            var v12 uint64 = v7
            _ = fi10
            _ = v12
        }
        _ = v8
        _ = v9
    } else {
        switch p0 {
        case "a":
            v13 := GenS1{F0: 1, F1: (v7 * (v7 * v7)), F2: -2147483648}
            _ = v13
        case "hello":
            v7 = v7
        case "":
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 uint32 = genFn0(1, -112)
                _ = fi14
                _ = v15
            }
        default:
            if (false && true) {
                v16 := GenS2{F0: 92}
                v17 := GenS1{F0: 2147483648, F1: 176, F2: 0}
                p0 = (("\n\t" + p0) + "")
                _ = v16
                _ = v17
            } else {
                v18 := []int{}
                if (true || (!(!false))) {
                    v19 := GenS3{F0: (960 > 0), F1: -569, F2: 107}
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    _ = v19
                }
                _ = v18
            }
        }
        fmt.Println(p0)
        var v20 int8 = 127
        _ = v20
    }
    _ = p0
    _ = v7
    return false
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 uint64 = 2
    var v4 float32 = 1.5
    var v21 float64 = 3.14159
    fmt.Println(v21, v21)
    switch 257 {
    case 42:
        v22 := GenS4{F0: 255, F1: -1, F2: -1}
        _ = v22
    default:
        for fi23 := 0; fi23 < 3; fi23++ {
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 uint8 = 1
                _ = fi24
                _ = v25
            }
            for fi26 := 0; fi26 < 1; fi26++ {
                v4 -= ((((3.14159 + 0.0) + v4) + 0.0) / v4)
                var v27 uint32 = 160
                _ = len(v2)
                _ = fi26
                _ = v27
            }
            var v28 int = fi23
            _ = fi23
            _ = v28
        }
    }
    var v29 int16 = -208
    v1[(((("" + "") + "foo bar") + "") + "a")] = "hello"
    fmt.Println(v3, v29)
    v30 := GenS1{F0: 2147483648, F1: v3, F2: -1}
    if ("" < "a") {
        v29 -= v29
    } else {
        if (("a" <= "a") && true) {
            v31 := v1[(("a" + "") + "\n\t")]
            for fi32 := 0; fi32 < 4; fi32++ {
                var v33 int32 = -1
                v21 += 1.5
                _ = fi32
                _ = v33
            }
            for fi34 := 0; fi34 < 18; fi34++ {
                if false {
                    v35 := GenS3{F0: (genFn0((v29 / ((v29) | 1)), -1) < 746), F1: 42, F2: 32767}
                    if true {
                        if false {
                            v31 = v31
                        }
                        for fi36 := 0; fi36 < 16; fi36++ {
                            fmt.Println(fi36, v4, v4)
                            _ = fi36
                        }
                        switch (fi34 - (^(fi34))) {
                        case 0:
                            var v37 uint64 = (v3 + (^(v3)))
                            _ = v37
                        default:
                            v3 = (^(v3))
                        }
                    }
                    _ = v35
                }
                v38 := GenS1{F0: 388, F1: v3, F2: 42}
                switch -1 {
                case -15:
                    if false {
                        if true {
                            var v39 string = (v31 + "foo bar")
                            _ = v39
                        } else {
                            v4 = v4
                        }
                        for fi40 := 0; fi40 < 20; fi40++ {
                            var v41 uint = 930
                            var v42 int32 = -2147483648
                            v3 = v3
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                        v29 = v29
                    }
                }
                _ = fi34
                _ = v38
            }
            _ = v31
        } else {
            for fi43 := 0; fi43 < 4; fi43++ {
                v4 += v4
                var v44 int16 = v29
                _ = fi43
                _ = v44
            }
        }
        switch 2147483647 {
        case -47:
            switch ("foo bar" + "a") {
            case "\n\t":
                v3 -= (v3 + (v3 % ((v3) | 1)))
            case "a":
                v45 := GenS3{F0: true, F1: v29, F2: 65535}
                _ = v45
            case "foo bar":
                fmt.Println(v29, v21)
            default:
                v46 := []int{518, 128, 127, 32768}
                _ = v46
            }
        case 0:
            switch "foo bar" {
            case "foo bar":
                v47 := []int{-2147483648, 127}
                _ = v47
            case "":
                for fi48 := 0; fi48 < 16; fi48++ {
                    if (false || true) {
                        for fi49 := 0; fi49 < 6; fi49++ {
                            fmt.Println(fi49, v3, fi48)
                            var v50 bool = false
                            v4 += v4
                            v51 := GenS4{F0: 1, F1: v29, F2: -2147483648}
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                    }
                    if true {
                        v52 := GenS5{F0: (v29 + -1)}
                        var v53 int16 = 1
                        _ = v52
                        _ = v53
                    }
                    _ = fi48
                }
            }
        case 42:
            for fi54 := 0; fi54 < 1; fi54++ {
                delete(v1, "a")
                _ = fi54
            }
        default:
            delete(v1, "a")
        }
    }
    v55 := []int{4294967295, 873, 4294967295}
    v55 = append(v55, -9223372036854775808)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v55:", len(v55), v55)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v55
}
