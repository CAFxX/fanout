package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 string; F2 uint16 }
type GenS7 struct { F0 string; F1 int8 }
func genFn0(p0 int8, p1 bool, p2 uint8) int {
    p0 = (p0 * p0)
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int16 = 32767
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 uint = 115
            p2 = p2
            _ = fi9
            _ = v10
        }
        fmt.Println(v8)
        _ = fi7
        _ = v8
    }
    switch 253 {
    case 1:
        p0 -= (p0 / (((p0 << ((255) & 63))) | 1))
    case 127:
        if p1 {
            fmt.Println(p2)
            for fi11 := 0; fi11 < 2; fi11++ {
                p0 -= (p0 ^ p0)
                _ = fi11
            }
            switch 49 {
            case 128:
                switch 0 {
                case -206:
                    switch 128 {
                    case -168:
                        for fi12 := 0; fi12 < 15; fi12++ {
                            fmt.Println(fi12, p2, p0)
                            if p1 {
                                fmt.Println(fi12, fi12)
                            }
                            _ = fi12
                        }
                    case -346:
                        v13 := GenS1{F0: 0, F1: -1, F2: 4294967295}
                        _ = v13
                    }
                default:
                    v14 := GenS2{F0: (p0 << ((42) & 63)), F1: "foo bar"}
                    _ = v14
                }
            case -9223372036854775808:
                if p1 {
                    p1 = p1
                    fmt.Println(p1, p1)
                    var v15 uint64 = 4294967295
                    _ = v15
                } else {
                    fmt.Println(p1)
                }
            }
        }
    default:
        if p1 {
            if p1 {
                fmt.Println(p1)
            }
            var v16 int = 42
            _ = v16
        }
    }
    v17 := GenS2{F0: (p0 + (^(p0))), F1: "hello"}
    switch 32768 {
    case -1:
        v18 := []int{}
        _ = v18
    case -334:
        if (p2 <= 128) {
            if (-233 < 982) {
                for fi19 := 0; fi19 < 16; fi19++ {
                    for fi20 := 0; fi20 < 8; fi20++ {
                        switch (fi19 & fi20) {
                        case -2147483648:
                            v21 := []int{-9223372036854775808, fi19, fi19}
                            _ = v21
                        case 42:
                            if (419 > 65535) {
                                if ((!p1) || (p2 != 2)) {
                                    v22 := GenS3{F0: 895, F1: -986, F2: 1}
                                    v23 := []int{fi19}
                                    if false {
                                        p0 = p0
                                    }
                                    _ = v22
                                    _ = v23
                                } else {
                                    v24 := []int{(fi20 % (((fi20 & (fi19 * (-(fi20))))) | 1)), 128, fi19, (fi20 % (((fi19 | fi19)) | 1))}
                                    _ = v24
                                }
                                var v25 bool = (0 <= 813)
                                _ = v25
                            } else {
                                v26 := GenS4{F0: 303}
                                var v27 int = (fi20 % ((int(p2)) | 1))
                                _ = v26
                                _ = v27
                            }
                        case 2147483647:
                            var v28 int64 = 706
                            _ = v28
                        }
                        _ = fi20
                    }
                    p1 = true
                    _ = fi19
                }
                var v29 uint = 65535
                _ = v29
            }
            for fi30 := 0; fi30 < 12; fi30++ {
                var v31 float64 = -2.25
                if (1.5 < float64((3.14159 / 1.0))) {
                    var v32 string = "a"
                    v33 := []int{}
                    var v34 int16 = 0
                    _ = v32
                    _ = v33
                    _ = v34
                }
                _ = fi30
                _ = v31
            }
        }
    case 127:
        p0 = -95
    default:
        switch 2147483647 {
        case 137:
            var v35 uint32 = 42
            _ = v35
        case 42:
            v36 := GenS1{F0: 422, F1: -249, F2: 255}
            _ = v36
        case -2147483648:
            p2 = 103
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 988
}

func main() {
    var v1 bool = (-2.25 <= (1.5 - 3.14159))
    var v2 uint8 = 127
    var v3 int16 = -959
    var v4 string = ""
    v5 := make(map[string]string)
    var v6 uint32 = 65535
    v6 -= 42
    v37 := GenS5{F0: 255}
    delete(v5, v4)
    for fi38 := 0; fi38 < 1; fi38++ {
        v39 := GenS6{F0: ((("foo bar" + "\n\t") + v4) + ((("a" + ("hello" + "a")) + "") + "\n\t")), F1: "\n\t", F2: 0}
        for fi40 := 0; fi40 < 8; fi40++ {
            v6 = v6
            _ = fi40
        }
        _ = fi38
        _ = v39
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        fmt.Println(v2, fi41)
        _ = fi41
    }
    v42 := GenS4{F0: -1}
    var v43 uint = 255
    if ((42 >= 10) || (-32768 == 646)) {
        v44 := v5[("\n\t" + v4)]
        _ = v44
    }
    var v45 float64 = ((((1.5 + (float64((3.14159 / 0.1)) + (0.1 + float64(1.5)))) - (((float64((((0.0 + 0.1) * 0.1) / 1.0)) * 1.5) * (1e10 + 0.1)) / (float64(1e10) / 1.0))) - (-2.25 - -2.25)) - ((0.0 / 1.5) * float64(3.14159)))
    v46 := []int{9223372036854775807}
    v47 := GenS7{F0: "foo bar", F1: 0}
    if true {
        var v48 int64 = -1
        switch genFn0(-83, (v1 && v1), (^(v2))) {
        case 1:
            var v49 int16 = v3
            _ = v49
        case 2147483647:
            v50 := GenS4{F0: 825}
            _ = v50
        }
        _ = v48
    } else {
        fmt.Println(v4, v2, v1)
        v51 := v5[(v4 + "a")]
        var v52 uint = (v43 * 0)
        _ = v51
        _ = v52
    }
    v4 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v42
    _ = v43
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
