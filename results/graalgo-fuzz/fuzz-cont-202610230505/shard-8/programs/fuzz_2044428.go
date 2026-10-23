package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int16; F1 uint; F2 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 uint }

func main() {
    var v1 bool = (32767 == -32768)
    var v2 float64 = 1e10
    v3 := make(map[string]bool)
    var v4 bool = (!v1)
    v5 := make(map[string]string)
    var v6 int8 = -86
    v1 = (true && v1)
    for fi7 := 0; fi7 < 1; fi7++ {
        v1 = ((false || (!(0 >= 42))) || ((-9223372036854775808 > 0) || v1))
        v8 := GenS1{F0: -1, F1: (-2.25 > (0.0 - ((1.5 / ((-2.25 * -2.25) - 1.5)) - -2.25)))}
        switch fi7 {
        case 1:
            var v9 int32 = -143
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        if ("hello" <= "foo bar") {
            v1 = (1 >= 65535)
            if (!(v1 || v4)) {
                for fi11 := 0; fi11 < 1; fi11++ {
                    var v12 uint = 1
                    switch (fi11 ^ (-(fi11))) {
                    case 1:
                        var v13 uint64 = 359
                        _ = v13
                    case 128:
                        var v14 int32 = 1
                        _ = v14
                    case 42:
                        if false {
                            v15 := []int{}
                            if v4 {
                                v16 := GenS2{F0: 42, F1: 1, F2: 32767}
                                switch 1 {
                                case 1:
                                    var v17 string = ("hello" + (("hello" + "\n\t") + "\n\t"))
                                    _ = v17
                                default:
                                    if len(v15) > 0 {
                                        v15[0] = -2147483648
                                    }
                                }
                                _ = v16
                            } else {
                                if v4 {
                                    if v1 {
                                        v2 += float64((-2.25 + (0.0 + ((1e10 + -2.25) / float32(v2)))))
                                        var v18 int8 = (v6 - 42)
                                        var v19 uint32 = 4294967295
                                        _ = v18
                                        _ = v19
                                    } else {
                                        switch -1 {
                                        case -965:
                                            v12 = v12
                                            v20 := GenS1{F0: -286, F1: (!(42 < 42))}
                                            var v21 int64 = -1
                                            _ = v20
                                            _ = v21
                                        case 543:
                                            var v22 uint8 = 128
                                            _ = v22
                                        case 32768:
                                            v2 = (((v2 / (3.14159 * v2)) / float64((1e10 * 1e10))) / v2)
                                        }
                                        fmt.Println(fi11, v12, fi10)
                                        var v23 int32 = -1
                                        _ = v23
                                    }
                                    v6 -= (^(v6))
                                    v4 = true
                                    var v24 float32 = -2.25
                                    var v25 string = "a"
                                    if false {
                                        v24 += (v24 * v24)
                                        if (((v1 || (3.14159 < (v2 / v2))) && (v4 && v4)) && (!(42 == 869))) {
                                            v3[(v25 + "\n\t")] = v1
                                            v26 := GenS3{F0: 32767, F1: (fi11 & fi11)}
                                            _ = v26
                                        }
                                        var v27 uint = v12
                                        _ = v27
                                    }
                                    if (291 == 2) {
                                        if (!(true || (!true))) {
                                            fmt.Println(v1, v24, v4)
                                        } else {
                                            v28 := v5[v25]
                                            v4 = (3.14159 <= v2)
                                            _ = v28
                                        }
                                        if ((v6 & (v6 / ((0) | 1))) < -1) {
                                            var v29 int16 = 941
                                            _ = v29
                                        }
                                        fmt.Println(v4)
                                    }
                                    _ = v24
                                    _ = v25
                                } else {
                                    v4 = v4
                                    v30 := v3[("hello" + ("hello" + "foo bar"))]
                                    _ = v30
                                }
                            }
                            if (v1 || (!v1)) {
                                switch fi10 {
                                case 127:
                                    v31 := GenS4{F0: -32768, F1: 65535, F2: (fi11 != (fi10 * (fi11 & fi10)))}
                                    _ = v31
                                case -684:
                                    fmt.Println(v12)
                                case -1:
                                    var v32 uint64 = 1
                                    _ = v32
                                }
                                v12 += v12
                            }
                            _ = v15
                        }
                    default:
                        var v33 uint = (^(v12))
                        _ = v33
                    }
                    switch (fi11 << (((v12 % ((v12) | 1))) & 63)) {
                    case 0:
                        v12 = 65535
                    default:
                        if (!(true || true)) {
                            fmt.Println(v1)
                            v2 = 3.14159
                        }
                    }
                    _ = fi11
                    _ = v12
                }
                var v34 int16 = 42
                _ = v34
            } else {
                v35 := GenS2{F0: 2, F1: 1, F2: int16(v6)}
                if (!v4) {
                    v36 := v3[("hello" + "")]
                    v1 = (127 == (v6 << ((4294967295) & 63)))
                    _ = v36
                } else {
                    v6 -= v6
                    fmt.Println(v2, v1)
                }
                _ = v35
            }
        } else {
            v4 = v4
        }
        fmt.Println(v6)
        _ = fi10
    }
    var v37 uint8 = 173
    if (true || v1) {
        v38 := GenS2{F0: 42, F1: 463, F2: 1000}
        v39 := []int{4294967295}
        if v4 {
            var v40 uint = 2
            for fi41 := 0; fi41 < 7; fi41++ {
                var v42 bool = (2147483647 == 785)
                switch (^(fi41)) {
                case -253:
                    var v43 float32 = 0.0
                    _ = v43
                default:
                    if (!true) {
                        _ = len(v3)
                        var v44 float64 = (3.14159 * 1e10)
                        var v45 float64 = (v2 / (1.5 * v2))
                        _ = v44
                        _ = v45
                    } else {
                        var v46 float32 = 3.14159
                        var v47 uint8 = v37
                        _ = v46
                        _ = v47
                    }
                }
                _ = fi41
                _ = v42
            }
            _ = v40
        } else {
            v48 := GenS2{F0: 950, F1: 18446744073709551615, F2: 143}
            _ = v48
        }
        _ = v38
        _ = v39
    }
    var v49 float32 = -2.25
    fmt.Println(v37, v2, v6)
    _ = len(v5)
    switch -2147483648 {
    case -817:
        v37 = v37
    case -9223372036854775808:
        if true {
            v1 = (236 >= 988)
            v50 := []int{0}
            _ = v50
        } else {
            if v4 {
                var v51 int32 = -2147483648
                if v1 {
                    var v52 uint64 = 18446744073709551615
                    var v53 bool = v1
                    _ = v52
                    _ = v53
                } else {
                    switch -300 {
                    case -9223372036854775808:
                        var v54 int = -1
                        _ = v54
                    case -854:
                        var v55 int = 128
                        _ = v55
                    default:
                        var v56 int8 = v6
                        _ = v56
                    }
                    delete(v5, ("" + ("\n\t" + (("\n\t" + ("foo bar" + "a")) + ""))))
                }
                var v57 int64 = 182
                _ = v51
                _ = v57
            } else {
                v58 := GenS5{F0: ("a" > "a")}
                if (42 >= 1) {
                    v1 = (733 != 249)
                }
                _ = v58
            }
        }
    default:
        v5[""] = "\n\t"
    }
    fmt.Println(v49, v4, v4)
    fmt.Println(v2)
    for fi59 := 0; fi59 < 7; fi59++ {
        v60 := GenS6{F0: 0, F1: 528}
        v6 += 1
        _ = fi59
        _ = v60
    }
    v6 -= v6
    v6 += v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
