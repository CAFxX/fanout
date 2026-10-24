package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 uint; F1 int32 }
func genFn0(p0 bool) int8 {
    p0 = p0
    var v7 bool = p0
    v8 := GenS1{F0: 0, F1: "foo bar", F2: 594}
    if false {
        var v9 uint64 = 149
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 int16 = 1000
            var v12 uint8 = 255
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 int32 = 2147483647
        _ = v13
        _ = v9
    }
    v7 = v7
    _ = p0
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int32, p1 uint16, p2 int8) uint16 {
    v14 := GenS2{F0: 0}
    v15 := GenS3{F0: "foo bar", F1: 1}
    var v16 string = ""
    v17 := []int{-2147483648, 4294967295, 32767, int(p2)}
    for fi18 := 0; fi18 < 14; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            v17 = append(v17, (fi18 + fi19))
            if false {
                v17 = append(v17, fi18)
                p2 -= -128
                switch ((v16 + "hello") + (((("hello" + v16) + "a") + "a") + "")) {
                case "foo bar":
                    p2 += -128
                case "hello":
                    switch fi18 {
                    case 127:
                        fmt.Println(p1)
                    }
                case "a":
                    if ((((0.1 + (0.0 + (-2.25 + float64((1.5 / 1.0))))) + 0.1) * 1e10) == 1e10) {
                        if false {
                            if (!false) {
                                v20 := GenS3{F0: v16, F1: 47}
                                _ = v20
                            } else {
                                v17 = append(v17, fi18)
                                v21 := GenS3{F0: v16, F1: 398}
                                fmt.Println(v16, p2, fi19)
                                _ = v21
                            }
                            v22 := GenS4{F0: p1}
                            _ = v22
                        }
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                    }
                }
                v23 := GenS2{F0: 521}
                _ = v23
            }
            if (4294967295 > 18446744073709551615) {
                fmt.Println(p1, v16, p1)
                if len(v17) > 0 {
                    v17[0] = (fi18 & int(p0))
                }
                v24 := GenS5{F0: 543}
                _ = v24
            }
            _ = fi19
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return 0
}

func main() {
    var v1 int8 = 1
    var v2 float64 = ((0.0 * 1.5) + ((1e10 * 1.5) * 3.14159))
    var v3 uint = 321
    v4 := make(map[string]uint32)
    v5 := []uint32{519, 25}
    v6 := []string{"a", "", ((("\n\t" + "") + "hello") + "foo bar")}
    fmt.Println(v2, v1)
    v25 := GenS3{F0: "foo bar", F1: 419}
    v26 := GenS4{F0: 921}
    var v27 int32 = -1
    var v28 uint16 = 1
    v2 = (v2 * v2)
    v1 = 0
    _ = len(v4)
    var v29 int8 = v1
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 float32 = (0.0 - 0.0)
        for fi32 := 0; fi32 < 6; fi32++ {
            v2 = v2
            if (-354 == 32767) {
                v33 := GenS5{F0: v3}
                if (!(!true)) {
                    v31 += v31
                    if ("\n\t" < ("a" + "hello")) {
                        if true {
                            v6 = append(v6, ("" + "hello"))
                        } else {
                            fmt.Println(fi30, v27, fi30)
                            v29 += (-(v1))
                            var v34 string = ""
                            _ = v34
                        }
                        switch (fi30 | (fi32 & fi32)) {
                        case -819:
                            fmt.Println(v27, fi30, v29)
                        case 128:
                            switch "\n\t" {
                            case "\n\t":
                                var v35 string = "\n\t"
                                _ = v35
                            case "foo bar":
                                fmt.Println(v27, v1, v31)
                            case "":
                                v2 = 3.14159
                            }
                        case -2147483648:
                            var v36 int8 = (-(v29))
                            _ = v36
                        }
                        switch ("" + "hello") {
                        case "hello":
                            _ = len(v4)
                        case "\n\t":
                            var v37 uint16 = 42
                            if (0.0 <= v31) {
                                if len(v5) > 0 {
                                    v5[0] = 4294967295
                                }
                                if len(v6) > 0 {
                                    v6[0] = ("" + "\n\t")
                                }
                                v38 := v4[("foo bar" + "hello")]
                                _ = v38
                            } else {
                                if true {
                                    v39 := GenS6{F0: fi32}
                                    v3 = (v3 | (v3 ^ uint(v28)))
                                    var v40 int8 = (v29 % ((v29) | 1))
                                    _ = v39
                                    _ = v40
                                }
                                fmt.Println(v3)
                            }
                            _ = v37
                        default:
                            fmt.Println(fi32, v31)
                        }
                    }
                    var v41 uint8 = 128
                    if ((v28 ^ 593) > genFn1((v27 - (v27 & v27)), v28, (v29 / ((v29) | 1)))) {
                        if ((v1 + genFn0(true)) >= int8(v3)) {
                            var v42 int64 = 42
                            v1 -= (v29 >> ((v3) & 63))
                            _ = v42
                        }
                        if (!true) {
                            v43 := GenS7{F0: 1, F1: 40}
                            _ = v43
                        } else {
                            fmt.Println(v3)
                            if (32767 < 32767) {
                                switch (("foo bar" + ((("" + "foo bar") + "") + ((((("hello" + "\n\t") + "a") + "a") + "foo bar") + ""))) + "a") {
                                case "hello":
                                    switch fi32 {
                                    case 9223372036854775807:
                                        v28 += genFn1(v27, (v28 >> (((v3 & v3)) & 63)), (v29 % ((v1) | 1)))
                                    case 128:
                                        _ = len(v5)
                                    }
                                case "a":
                                    v31 = -2.25
                                case "\n\t":
                                    v31 = v31
                                }
                                if ((v3 % ((v3) | 1)) > (v3 ^ v3)) {
                                    delete(v4, "a")
                                    var v44 uint16 = (^(v28))
                                    v27 = 0
                                    _ = v44
                                }
                                delete(v4, (("\n\t" + "foo bar") + ("a" + ((((("a" + (("hello" + "a") + "hello")) + "hello") + "\n\t") + ("" + ("" + "\n\t"))) + ("foo bar" + ("foo bar" + "a"))))))
                                v28 = (v28 * genFn1((v27 ^ (-(v27))), (^(v28)), v1))
                                if (genFn0((65535 != 65535)) != v29) {
                                    v45 := GenS7{F0: 1, F1: 1}
                                    switch (fi32 / ((fi30) | 1)) {
                                    case 128:
                                        v28 = (v28 / ((v28) | 1))
                                    case -1:
                                        if (genFn0((v3 >= v3)) <= -66) {
                                            fmt.Println(fi30, fi30, v27)
                                        }
                                    case -736:
                                        if len(v6) > 0 {
                                            v6[0] = "a"
                                        }
                                    }
                                    fmt.Println(v27)
                                    _ = v45
                                }
                            }
                            v29 += (v29 ^ (v1 - (v29 << ((497) & 63))))
                        }
                    }
                    _ = v41
                } else {
                    v46 := GenS8{F0: 2, F1: (v27 | 1)}
                    _ = v46
                }
                v28 = genFn1((^(v27)), (v28 % ((v28) | 1)), v1)
                _ = v33
            }
            delete(v4, ("hello" + ""))
            _ = fi32
        }
        if true {
            switch "hello" {
            case "hello":
                switch ((("\n\t" + "foo bar") + "foo bar") + "") {
                case "foo bar":
                    for fi47 := 0; fi47 < 5; fi47++ {
                        fmt.Println(v2, v3, v2)
                        v48 := GenS6{F0: fi47}
                        _ = len(v4)
                        _ = fi47
                        _ = v48
                    }
                }
            }
        } else {
            delete(v4, "hello")
            v31 = 1.5
            v31 -= 1.5
        }
        _ = fi30
        _ = v31
    }
    for fi49 := 0; fi49 < 11; fi49++ {
        v29 = 1
        var v50 bool = false
        if (v50 || (!false)) {
            _ = len(v6)
            _ = len(v4)
        } else {
            fmt.Println(v27, v50, v3)
        }
        _ = fi49
        _ = v50
    }
    fmt.Println(v2)
    var v51 float64 = v2
    var v52 int8 = v29
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
