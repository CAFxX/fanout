package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 bool; F2 string }
type GenS5 struct { F0 int64; F1 string }
func genFn0(p0 int) bool {
    var v5 float64 = (3.14159 - 3.14159)
    var v6 uint8 = 128
    _ = p0
    _ = v5
    _ = v6
    return true
}
func genFn1(p0 int8, p1 int8, p2 uint16) int16 {
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            if false {
                p0 = (p1 + p0)
                v9 := []int{fi8, fi8, fi7}
                if (570 == fi8) {
                    v10 := GenS1{F0: 2, F1: 697, F2: 2}
                    _ = len(v9)
                    fmt.Println(p0)
                    _ = v10
                } else {
                    v11 := GenS1{F0: 2147483647, F1: 65535, F2: 90}
                    if (false || true) {
                        fmt.Println(fi8)
                        if (!(-32768 >= 1)) {
                            v12 := GenS1{F0: 42, F1: 844, F2: 21}
                            p1 = (p1 | 0)
                            if true {
                                if len(v9) > 0 {
                                    v9[0] = 42
                                }
                                p0 += p0
                            }
                            if (0 < 0) {
                                v13 := GenS2{F0: (fi7 | fi8), F1: 32768}
                                var v14 uint32 = 1
                                var v15 uint64 = 0
                                _ = v13
                                _ = v14
                                _ = v15
                            } else {
                                switch fi7 {
                                case 9223372036854775807:
                                    if false {
                                        var v16 float32 = -2.25
                                        if (true && false) {
                                            var v17 int16 = -1
                                            _ = v17
                                        }
                                        p0 = (p0 >> ((4294967295) & 63))
                                        var v18 uint64 = 4294967295
                                        v19 := GenS2{F0: -119, F1: 955}
                                        _ = v16
                                        _ = v18
                                        _ = v19
                                    } else {
                                        v20 := GenS3{F0: -21}
                                        _ = v20
                                    }
                                default:
                                    fmt.Println(fi7, fi8, p0)
                                }
                                fmt.Println(p2)
                            }
                            _ = v12
                        } else {
                            if false {
                                switch (^(fi7)) {
                                case 1:
                                    v21 := GenS1{F0: 747, F1: 96, F2: 0}
                                    _ = v21
                                case 128:
                                    var v22 int32 = -1
                                    _ = v22
                                }
                            }
                            if true {
                                var v23 int32 = -1
                                _ = v23
                            } else {
                                switch ("foo bar" + ("a" + ("hello" + "\n\t"))) {
                                case "foo bar":
                                    var v24 int16 = 0
                                    _ = v24
                                case "a":
                                    if (0 >= 886) {
                                        if len(v9) > 0 {
                                            v9[0] = (^(fi7))
                                        }
                                        p2 -= (p2 >> ((51) & 63))
                                        v25 := GenS4{F0: ((!(18446744073709551615 < 959)) && (!false)), F1: (11 != 128), F2: "hello"}
                                        _ = v25
                                    } else {
                                        fmt.Println(p2)
                                        var v26 uint = 804
                                        _ = v26
                                    }
                                default:
                                    p0 -= (p1 >> ((981) & 63))
                                }
                                fmt.Println(p0, p2, p0)
                            }
                        }
                    }
                    p2 = (p2 | 1)
                    _ = v11
                }
                _ = v9
            } else {
                if true {
                    if (1 <= 128) {
                        switch fi8 {
                        case 332:
                            var v27 uint64 = uint64(p2)
                            _ = v27
                        case -1:
                            var v28 bool = (0.1 < (1e10 / (0.0 - 1e10)))
                            _ = v28
                        default:
                            v29 := GenS4{F0: false, F1: true, F2: "a"}
                            _ = v29
                        }
                    } else {
                        var v30 int16 = 1
                        var v31 float64 = 0.0
                        p0 = int8(p2)
                        _ = v30
                        _ = v31
                    }
                } else {
                    var v32 int8 = p1
                    v33 := []int{}
                    fmt.Println(p1, p1)
                    _ = v32
                    _ = v33
                }
                p1 -= p0
                var v34 string = (("a" + "foo bar") + ((((("\n\t" + "hello") + "hello") + "foo bar") + ("\n\t" + ((("\n\t" + (((("\n\t" + "hello") + ((("" + "foo bar") + (("hello" + "\n\t") + "")) + "foo bar")) + ("hello" + ("hello" + "foo bar"))) + "\n\t")) + "a") + "foo bar"))) + ("foo bar" + "")))
                _ = v34
            }
            fmt.Println(p2)
            _ = fi8
        }
        _ = fi7
    }
    if ((p1 >> ((251) & 63)) > p1) {
        var v35 int64 = 9223372036854775807
        switch "\n\t" {
        case "a":
            v36 := GenS1{F0: 2, F1: 485, F2: 178}
            _ = v36
        }
        for fi37 := 0; fi37 < 2; fi37++ {
            p1 = 37
            v38 := []int{}
            if (255 <= 1) {
                fmt.Println(p0)
                v39 := GenS5{F0: v35, F1: "foo bar"}
                _ = v39
            }
            _ = fi37
            _ = v38
        }
        _ = v35
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 float32 = float32(3.14159)
    v2 := make(map[uint32]bool)
    v3 := []string{"a", "a", "\n\t", "\n\t"}
    var v4 uint64 = 807
    for fi40 := 0; fi40 < 20; fi40++ {
        v1 += float32(float64((v1 + (-2.25 + ((v1 + v1) * v1)))))
        _ = fi40
    }
    for fi41 := 0; fi41 < 14; fi41++ {
        fmt.Println(v4, v1, fi41)
        var v42 float64 = 1.5
        _ = fi41
        _ = v42
    }
    var v43 int64 = 42
    fmt.Println(v43)
    fmt.Println(v4)
    fmt.Println(v1, v4, v1)
    _ = len(v3)
    fmt.Println(v4, v43)
    if (-873 == 75) {
        for fi44 := 0; fi44 < 5; fi44++ {
            if (0 > 1) {
                v1 -= (v1 + (v1 - 0.0))
                var v45 bool = true
                _ = len(v2)
                _ = v45
            } else {
                fmt.Println(v4)
            }
            var v46 int32 = 1
            _ = fi44
            _ = v46
        }
        var v47 uint16 = 65535
        _ = v47
    }
    var v48 float32 = v1
    var v49 int32 = 46
    v4 = 18446744073709551615
    v4 = v4
    for fi50 := 0; fi50 < 14; fi50++ {
        fmt.Println(fi50, v48, v4)
        _ = fi50
    }
    var v51 int16 = -656
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v48
    _ = v49
    _ = v51
}
