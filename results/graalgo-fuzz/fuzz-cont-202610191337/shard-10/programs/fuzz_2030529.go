package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 bool; F1 int8; F2 int64 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint64) uint16 {
    v7 := []int{-2147483648}
    fmt.Println(p0)
    v8 := GenS1{F0: "a", F1: "hello", F2: 1}
    v9 := GenS1{F0: (("" + (("hello" + "hello") + ((("a" + "a") + ("hello" + "foo bar")) + ("hello" + "foo bar")))) + "a"), F1: ("" + "a"), F2: -1}
    if len(v7) > 0 {
        v7[0] = -266
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := []int{9223372036854775807, 9223372036854775807}
    v2 := []string{("" + "a"), (("hello" + "hello") + "")}
    v3 := []uint32{937, 310}
    v4 := make(map[int]string)
    var v5 int8 = 127
    v6 := make(map[string]string)
    v5 -= (v5 ^ (v5 | (^(v5))))
    if (-188 == -410) {
        delete(v4, 0)
        var v10 int64 = 1
        _ = v10
    } else {
        v5 = v5
        var v11 bool = false
        _ = v11
    }
    fmt.Println(v5)
    v5 += (v5 | v5)
    if false {
        delete(v4, 372)
        switch "hello" {
        case "a":
            _ = len(v6)
        case "\n\t":
            v5 -= v5
        case "":
            switch 9223372036854775807 {
            case 1:
                v5 -= (^(v5))
            case -242:
                fmt.Println(v5)
            case -9223372036854775808:
                v12 := GenS2{F0: 42}
                _ = v12
            }
        default:
            fmt.Println(v5)
        }
    }
    v5 += -1
    if false {
        v5 = 42
    }
    var v13 uint8 = 1
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            switch fi15 {
            case -2147483648:
                var v16 float64 = 1e10
                _ = v16
            default:
                if false {
                    v5 = (v5 - 1)
                    var v17 uint8 = (^(v13))
                    _ = v17
                } else {
                    var v18 int8 = v5
                    if true {
                        var v19 uint64 = 4294967295
                        _ = v19
                    } else {
                        v5 = 0
                        var v20 string = ("foo bar" + "hello")
                        v20 = (v20 + "\n\t")
                        _ = v20
                    }
                    var v21 bool = (42 > -32768)
                    _ = v18
                    _ = v21
                }
            }
            var v22 uint16 = 65535
            v22 += v22
            _ = fi15
            _ = v22
        }
        _ = fi14
    }
    fmt.Println(v13, v5)
    var v23 bool = true
    v24 := GenS3{F0: v13, F1: v13}
    for fi25 := 0; fi25 < 7; fi25++ {
        switch ((("hello" + ("\n\t" + "a")) + (((("\n\t" + "") + (("hello" + "a") + "foo bar")) + "foo bar") + ("" + "a"))) + "foo bar") {
        case "a":
            v5 = (v5 % ((v5) | 1))
        case "foo bar":
            for fi26 := 0; fi26 < 10; fi26++ {
                delete(v4, (fi25 - (fi25 ^ (-(fi25)))))
                _ = fi26
            }
        }
        _ = fi25
    }
    var v27 int16 = 0
    if (2 <= 1) {
        if (-2.25 == (0.0 + float32(1.5))) {
            v13 += (v13 / ((0) | 1))
            var v28 float64 = 0.0
            _ = v28
        } else {
            for fi29 := 0; fi29 < 13; fi29++ {
                v13 -= v13
                delete(v6, ((("hello" + (("a" + "hello") + "hello")) + ("hello" + "")) + "\n\t"))
                if true {
                    var v30 int32 = 851
                    if ((^(v27)) <= v27) {
                        if true {
                            v2 = append(v2, "\n\t")
                            if true {
                                var v31 float32 = ((1.5 - -2.25) * (0.1 - -2.25))
                                _ = v31
                            }
                            v30 = 42
                        } else {
                            for fi32 := 0; fi32 < 8; fi32++ {
                                var v33 string = ("\n\t" + "")
                                v4[(fi29 * 32767)] = v33
                                v30 = v30
                                var v34 uint16 = 0
                                if v23 {
                                    v35 := GenS4{F0: false, F1: (v5 >> ((255) & 63)), F2: -958}
                                    if (-1 <= -516) {
                                        if v23 {
                                            v36 := GenS5{F0: v5}
                                            fmt.Println(v5)
                                            _ = v36
                                        } else {
                                            var v37 uint16 = (v34 << ((255) & 63))
                                            _ = v37
                                        }
                                        fmt.Println(fi29, v34, v34)
                                    }
                                    var v38 int8 = v5
                                    if v23 {
                                        v23 = ((v5 + (-(v38))) <= int8(v34))
                                        if (!(false && (true && v23))) {
                                            fmt.Println(fi32)
                                            var v39 int16 = v27
                                            v5 += (v38 % ((0) | 1))
                                            _ = v39
                                        } else {
                                            delete(v4, (fi29 | 2147483647))
                                        }
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = v33
                                    }
                                    if v23 {
                                        if false {
                                            if v23 {
                                                if (false || (-2.25 != (3.14159 - 1e10))) {
                                                    switch int(v34) {
                                                    case 32767:
                                                        v4[(-(fi29))] = ""
                                                    }
                                                }
                                                v23 = true
                                                v40 := GenS1{F0: (v33 + "foo bar"), F1: "\n\t", F2: (-(v30))}
                                                _ = v40
                                            } else {
                                                v41 := GenS2{F0: (v13 << ((65535) & 63))}
                                                switch (fi29 << ((0) & 63)) {
                                                case 9223372036854775807:
                                                    switch (-(fi32)) {
                                                    case 32767:
                                                        switch (fi29 & fi32) {
                                                        case 42:
                                                            v34 = v34
                                                        default:
                                                            var v42 uint64 = 717
                                                            _ = v42
                                                        }
                                                    case 9223372036854775807:
                                                        v23 = true
                                                    default:
                                                        var v43 int = -497
                                                        _ = v43
                                                    }
                                                case -346:
                                                    v6["a"] = (((v33 + ((v33 + v33) + "a")) + (("hello" + v33) + (("\n\t" + ((("\n\t" + (v33 + v33)) + "\n\t") + v33)) + v33))) + (v33 + (v33 + "a")))
                                                }
                                                _ = v41
                                            }
                                            v23 = (0.0 > 1e10)
                                            v5 = (^(v38))
                                        }
                                        var v44 uint8 = 42
                                        _ = v44
                                    } else {
                                        _ = len(v4)
                                        v27 = v27
                                    }
                                    var v45 bool = v23
                                    if (int32(v13) >= 42) {
                                        v46 := GenS2{F0: v13}
                                        v5 = 1
                                        _ = v46
                                    } else {
                                        var v47 bool = ((!true) && false)
                                        v33 = "a"
                                        _ = v47
                                    }
                                    _ = v35
                                    _ = v38
                                    _ = v45
                                } else {
                                    v13 += (v13 >> ((924) & 63))
                                    switch (v33 + v33) {
                                    case "":
                                        if (v23 || ((((3.14159 - 1.5) * (3.14159 * (1.5 + (1e10 / 3.14159)))) + 3.14159) > -2.25)) {
                                            switch v33 {
                                            case "a":
                                                var v48 int8 = v5
                                                _ = v48
                                            case "hello":
                                                var v49 uint32 = 2147483648
                                                _ = v49
                                            }
                                            v50 := v4[128]
                                            _ = v50
                                        }
                                    }
                                }
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                            var v51 int16 = v27
                            _ = v51
                        }
                        _ = len(v3)
                    }
                    _ = v30
                } else {
                    var v52 string = ("" + ("" + "a"))
                    _ = v52
                }
                _ = fi29
            }
            v5 = v5
            fmt.Println(v5, v5)
        }
        fmt.Println(v23, v13)
        if true {
            var v53 uint32 = 65535
            switch ("" + ("" + "foo bar")) {
            case "a":
                v27 = int16(v13)
            case "hello":
                var v54 bool = (!true)
                _ = v54
            }
            _ = v53
        } else {
            for fi55 := 0; fi55 < 18; fi55++ {
                if (!(v23 || true)) {
                    delete(v4, (^(fi55)))
                }
                _ = fi55
            }
        }
    } else {
        var v56 uint8 = (v13 % ((v13) | 1))
        v5 = v5
        _ = v56
    }
    v27 = (v27 / ((-574) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
