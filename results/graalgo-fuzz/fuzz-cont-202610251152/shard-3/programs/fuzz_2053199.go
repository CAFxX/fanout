package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint16, p1 int) int16 {
    v7 := GenS1{F0: 730, F1: -963, F2: uint8(p0)}
    var v8 int8 = -128
    p0 -= p0
    fmt.Println(v8, p0)
    v8 -= v8
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1000
}
func genFn1() bool {
    var v9 uint16 = 910
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := []int{}
        for fi12 := 0; fi12 < 12; fi12++ {
            if true {
                if false {
                    var v13 uint32 = 1
                    if false {
                        var v14 uint64 = 1
                        _ = v14
                    }
                    var v15 uint = 0
                    _ = v13
                    _ = v15
                } else {
                    v16 := GenS1{F0: 2147483647, F1: 9223372036854775807, F2: 128}
                    if true {
                        var v17 float64 = ((1e10 / float64(1.5)) / float64(0.1))
                        v18 := GenS2{F0: v9}
                        var v19 int64 = 972
                        if false {
                            var v20 int = -1
                            switch (fi12 * (fi10 << ((255) & 63))) {
                            case -301:
                                fmt.Println(v9, v17)
                            }
                            if (!(!false)) {
                                v9 += 811
                            }
                            v17 -= float64((0.0 + float32(float64(((3.14159 / 1.5) * (1.5 - 1e10))))))
                            _ = v20
                        }
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        if ((false || true) || false) {
                            var v21 uint16 = v9
                            var v22 string = ("\n\t" + "\n\t")
                            switch v22 {
                            case "":
                                if len(v11) > 0 {
                                    v11[0] = fi12
                                }
                            case "\n\t":
                                v22 = ((v22 + "") + "hello")
                            }
                            fmt.Println(fi10, fi12)
                            switch ("" + v22) {
                            case "hello":
                                if true {
                                    if false {
                                        _ = len(v11)
                                    } else {
                                        v22 = v22
                                    }
                                }
                            default:
                                fmt.Println(v22, v21)
                            }
                            v22 = (v22 + ((v22 + v22) + ((v22 + (("a" + v22) + v22)) + ("" + v22))))
                            _ = v21
                            _ = v22
                        } else {
                            var v23 uint8 = 127
                            v23 -= (v23 / ((v23) | 1))
                            if len(v11) > 0 {
                                v11[0] = fi10
                            }
                            var v24 int8 = 0
                            _ = v23
                            _ = v24
                        }
                    }
                    var v25 float64 = (((3.14159 - 0.0) * ((((0.1 / 1.0) / 1.0) * 1.5) + (-2.25 * (((3.14159 / float64((-2.25 / 1.0))) - float64(1e10)) - ((0.1 / (float64(3.14159) / (1e10 / ((1e10 - 1.5) + 1e10)))) / 1.0))))) - (float64(float32((0.0 + ((-2.25 - 1e10) - 1e10)))) * -2.25))
                    _ = v16
                    _ = v25
                }
            } else {
                v11 = append(v11, (fi12 % ((fi10) | 1)))
                var v26 uint64 = 1
                _ = v26
            }
            var v27 uint16 = (v9 % (((v9 << ((18446744073709551615) & 63))) | 1))
            _ = fi12
            _ = v27
        }
        v9 = v9
        _ = fi10
        _ = v11
    }
    v9 += v9
    v28 := []int{2147483647, 42, 0}
    if ((("foo bar" + ("a" + "")) + "foo bar") < (("hello" + "a") + "hello")) {
        for fi29 := 0; fi29 < 10; fi29++ {
            if (47 == 83) {
                v30 := GenS1{F0: 1, F1: -944, F2: 42}
                var v31 int64 = 9223372036854775807
                _ = v30
                _ = v31
            } else {
                fmt.Println(v9)
            }
            var v32 bool = false
            _ = len(v28)
            _ = fi29
            _ = v32
        }
        if true {
            v9 += v9
            v33 := GenS2{F0: v9}
            var v34 uint32 = 2147483647
            _ = v33
            _ = v34
        }
        switch "\n\t" {
        case "\n\t":
            v9 -= 65535
        case "a":
            v9 -= (v9 ^ 0)
        case "hello":
            var v35 int = -845
            _ = v35
        }
    } else {
        v9 += 819
        if false {
            if false {
                if len(v28) > 0 {
                    v28[0] = 32767
                }
                _ = len(v28)
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
            } else {
                for fi36 := 0; fi36 < 10; fi36++ {
                    v9 = (v9 ^ 103)
                    _ = fi36
                }
                v28 = append(v28, -680)
            }
            var v37 int16 = -1
            _ = v37
        }
        fmt.Println(v9)
    }
    _ = v28
    _ = v9
    return (4294967295 == 65535)
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 42
    v3 := []int{1}
    var v4 string = "foo bar"
    var v5 uint = 4294967295
    var v6 int = 32768
    for fi38 := 0; fi38 < 19; fi38++ {
        v39 := GenS3{F0: 1, F1: 9223372036854775807, F2: 288}
        _ = fi38
        _ = v39
    }
    var v40 bool = false
    v41 := GenS4{F0: (v5 & v5)}
    v42 := GenS3{F0: 127, F1: 1, F2: 0}
    switch v4 {
    case "a":
        v43 := GenS4{F0: v5}
        _ = v43
    case "\n\t":
        v44 := GenS4{F0: (v5 + (v5 >> (((v5 & v5)) & 63)))}
        _ = v44
    default:
        v2 -= 127
    }
    v1[-1] = (v6 | (v6 & v6))
    v5 += (v5 & (v5 >> ((v5) & 63)))
    for fi45 := 0; fi45 < 6; fi45++ {
        delete(v1, fi45)
        _ = fi45
    }
    var v46 bool = v40
    if ((!v46) || v46) {
        v47 := GenS2{F0: 2}
        if (!v40) {
            switch v4 {
            case "":
                v2 -= (v2 >> ((v5) & 63))
            case "foo bar":
                v48 := v1[9223372036854775807]
                _ = v48
            }
            var v49 uint8 = v2
            for fi50 := 0; fi50 < 12; fi50++ {
                for fi51 := 0; fi51 < 2; fi51++ {
                    if v40 {
                        v5 = v5
                        v40 = v40
                        if len(v3) > 0 {
                            v3[0] = v6
                        }
                    }
                    _ = fi51
                }
                _ = fi50
            }
            _ = v49
        } else {
            if (-2.25 == -2.25) {
                if v46 {
                    if (4294967295 <= v5) {
                        for fi52 := 0; fi52 < 2; fi52++ {
                            for fi53 := 0; fi53 < 12; fi53++ {
                                var v54 uint16 = 576
                                _ = fi53
                                _ = v54
                            }
                            var v55 int16 = 1
                            var v56 int8 = 127
                            _ = fi52
                            _ = v55
                            _ = v56
                        }
                        switch (((v4 + v4) + "a") + (v4 + "foo bar")) {
                        case "\n\t":
                            v46 = (v40 || (420 != 0))
                        case "foo bar":
                            fmt.Println(v5)
                        case "a":
                            v57 := v1[v6]
                            _ = v57
                        default:
                            var v58 bool = (v46 || (v40 && v40))
                            _ = v58
                        }
                    } else {
                        switch (v4 + v4) {
                        case "\n\t":
                            var v59 int16 = 32767
                            _ = v59
                        }
                        v4 = (v4 + v4)
                        var v60 string = v4
                        _ = v60
                    }
                    v4 = v4
                }
                if ((1e10 + 0.0) >= 0.1) {
                    v40 = v40
                    v61 := GenS4{F0: (v5 + 194)}
                    _ = v61
                }
                if false {
                    if true {
                        v40 = v40
                        v62 := GenS4{F0: v5}
                        _ = v62
                    }
                }
            }
            for fi63 := 0; fi63 < 4; fi63++ {
                var v64 uint8 = v2
                _ = fi63
                _ = v64
            }
            var v65 uint64 = 307
            _ = v65
        }
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v46
    _ = v5
    _ = v6
}
