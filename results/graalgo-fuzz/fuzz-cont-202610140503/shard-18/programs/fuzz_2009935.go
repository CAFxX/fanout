package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 string }
type GenS2 struct { F0 int8; F1 int32; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint8 }
func genFn0(p0 bool, p1 int, p2 uint8) uint16 {
    if p0 {
        var v6 uint32 = 230
        var v7 int8 = 127
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: -425, F1: p1, F2: ""}
        _ = v8
    }
    var v9 uint64 = 81
    var v10 uint16 = 489
    p1 = (p1 % ((-471) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 255
}

func main() {
    var v1 int = -584
    var v2 uint16 = 65535
    var v3 float64 = 1e10
    var v4 int64 = -9223372036854775808
    var v5 int = (^(v1))
    var v11 float64 = v3
    v12 := []int{0}
    v13 := GenS2{F0: 1, F1: 0, F2: 42}
    switch "a" {
    case "":
        var v14 int64 = v4
        _ = v14
    case "foo bar":
        switch 32767 {
        case 4294967295:
            v15 := GenS2{F0: 127, F1: -471, F2: 255}
            _ = v15
        case 9223372036854775807:
            if true {
                v16 := GenS2{F0: 1, F1: 2147483647, F2: 42}
                for fi17 := 0; fi17 < 20; fi17++ {
                    if false {
                        if true {
                            v11 = float64((3.14159 * (float32(float64((1.5 + 0.0))) - (-2.25 + float32(0.0)))))
                        }
                        if (42 >= 127) {
                            var v18 bool = false
                            if false {
                                v12 = append(v12, 42)
                                if ((128 == 66) || v18) {
                                    var v19 float64 = v11
                                    _ = v19
                                }
                                for fi20 := 0; fi20 < 20; fi20++ {
                                    v4 = v4
                                    fmt.Println(v18, fi17, v1)
                                    v21 := GenS3{F0: v5}
                                    _ = fi20
                                    _ = v21
                                }
                            } else {
                                v11 -= 1.5
                            }
                            _ = v18
                        } else {
                            if true {
                                v22 := GenS1{F0: v4, F1: (^(v1)), F2: "foo bar"}
                                var v23 uint8 = 127
                                if (100 == 73) {
                                    v4 -= (v4 - (v4 * 1))
                                    v12 = append(v12, 4294967295)
                                } else {
                                    var v24 int32 = 1
                                    if (!false) {
                                        fmt.Println(fi17, v23, v2)
                                        if false {
                                            v24 += (^(v24))
                                        } else {
                                            v25 := GenS4{F0: -52}
                                            fmt.Println(fi17, v23, v11)
                                            _ = v25
                                        }
                                    } else {
                                        if (v24 <= (v24 & v24)) {
                                            var v26 bool = ((^(v24)) >= (-(v24)))
                                            var v27 float64 = v11
                                            var v28 int64 = (v4 % (((v4 << ((18446744073709551615) & 63))) | 1))
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        }
                                        var v29 int = (^(fi17))
                                        _ = v29
                                    }
                                    for fi30 := 0; fi30 < 16; fi30++ {
                                        v31 := GenS1{F0: (v4 * 0), F1: fi30, F2: "\n\t"}
                                        _ = fi30
                                        _ = v31
                                    }
                                    _ = v24
                                }
                                _ = v22
                                _ = v23
                            }
                            switch "\n\t" {
                            case "a":
                                if (47 == -115) {
                                    if true {
                                        fmt.Println(v11, fi17)
                                    }
                                    v32 := GenS3{F0: (fi17 + (fi17 & (fi17 << ((420) & 63))))}
                                    var v33 uint64 = 18446744073709551615
                                    _ = v32
                                    _ = v33
                                }
                            case "":
                                if len(v12) > 0 {
                                    v12[0] = (^(v5))
                                }
                            }
                            if ((255 <= 71) && false) {
                                switch (v1 >> ((0) & 63)) {
                                case -1:
                                    if true {
                                        fmt.Println(fi17, fi17)
                                        v4 = v4
                                        fmt.Println(v5, v5)
                                    }
                                case 32768:
                                    if (127 < 0) {
                                        var v34 string = ("a" + "\n\t")
                                        if false {
                                            var v35 int16 = -32768
                                            _ = v35
                                        } else {
                                            v12 = append(v12, -216)
                                            var v36 float64 = (-2.25 + v3)
                                            _ = v36
                                        }
                                        _ = v34
                                    }
                                }
                            } else {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            }
                        }
                        switch fi17 {
                        case 1:
                            var v37 bool = (18446744073709551615 <= 18446744073709551615)
                            _ = v37
                        case 32767:
                            v38 := GenS3{F0: (fi17 + v5)}
                            _ = v38
                        }
                    }
                    v39 := GenS2{F0: 127, F1: 0, F2: 255}
                    _ = fi17
                    _ = v39
                }
                if true {
                    var v40 float64 = v3
                    v41 := GenS5{F0: (v5 >> ((1) & 63)), F1: 42}
                    _ = v40
                    _ = v41
                }
                _ = v16
            } else {
                var v42 uint16 = v2
                var v43 int8 = -88
                switch v5 {
                case 550:
                    var v44 uint16 = genFn0(false, v1, uint8(v42))
                    _ = v44
                case 63:
                    fmt.Println(v1, v42)
                }
                _ = v42
                _ = v43
            }
        case 128:
            if (((("hello" + "a") + (("a" + (("a" + "\n\t") + "a")) + "")) + "a") == ((("hello" + (("foo bar" + "hello") + "\n\t")) + "foo bar") + "hello")) {
                switch (v5 | (v1 | 9223372036854775807)) {
                case -2147483648:
                    if true {
                        v11 = v11
                    }
                }
                var v45 uint64 = 2
                _ = v45
            }
        }
    case "a":
        var v46 float32 = 0.1
        _ = v46
    }
    var v47 int16 = -956
    var v48 int32 = -1
    v49 := GenS3{F0: 32767}
    v11 = v3
    v11 = v3
    switch (v1 / (((v5 % (((v5 / (((v1 % (((^(v1))) | 1))) | 1))) | 1))) | 1)) {
    case 127:
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    v1 -= (v5 | (v5 * v5))
    v50 := GenS6{F0: 2147483648, F1: genFn0((!true), 4294967295, 127), F2: (v47 % ((v47) | 1))}
    if ((v4 > (v4 % ((v4) | 1))) && (!(v2 <= genFn0(true, v1, 18)))) {
        v51 := GenS7{F0: 42}
        _ = v51
    } else {
        var v52 uint32 = 523
        switch "hello" {
        case "":
            var v53 int64 = v4
            _ = v53
        case "a":
            v54 := GenS8{F0: 1}
            _ = v54
        }
        var v55 string = "\n\t"
        _ = v52
        _ = v55
    }
    var v56 float32 = 0.0
    v12 = append(v12, (v5 % ((-9223372036854775808) | 1)))
    v11 = 0.0
    if (-2.25 == v56) {
        if (((float64(v56) - float64(v56)) / 0.1) >= -2.25) {
            v57 := GenS1{F0: v4, F1: (^(v5)), F2: (("" + "hello") + "hello")}
            if false {
                var v58 int16 = int16(v1)
                if len(v12) > 0 {
                    v12[0] = v1
                }
                _ = v58
            } else {
                var v59 uint64 = 2
                _ = v59
            }
            _ = v57
        } else {
            var v60 int32 = v48
            var v61 uint64 = 42
            _ = v60
            _ = v61
        }
        v56 = v56
    }
    v56 -= (float32(((v11 / (0.0 + v3)) / (((v3 - (((float64(float32((v3 - 0.1))) / (1.5 / 1.5)) + (float64((-2.25 + (v56 * (v56 * v56)))) * v3)) + (v11 - (v3 - (-2.25 - v11))))) + (v3 + v3)) + -2.25))) + v56)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v56
}
