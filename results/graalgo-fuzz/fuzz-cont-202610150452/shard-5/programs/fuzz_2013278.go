package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 bool }
func genFn0(p0 uint32, p1 int16, p2 string) uint16 {
    v6 := []int{688, 2147483647, 32768}
    fmt.Println(p0, p0)
    p0 = (p0 >> ((21) & 63))
    p2 = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 65535
}
func genFn1() string {
    var v7 float64 = 1e10
    fmt.Println(v7)
    if (-678 >= 42) {
        var v8 string = "a"
        v9 := GenS1{F0: -9223372036854775808}
        v8 = v8
        _ = v8
        _ = v9
    }
    _ = v7
    return "hello"
}

func main() {
    var v1 float64 = 1.5
    var v2 uint = 42
    var v3 int = -2147483648
    v4 := []uint32{2147483647}
    v5 := make(map[string]bool)
    if (v3 < (v3 / (((v3 >> ((v2) & 63))) | 1))) {
        var v10 int64 = 1
        _ = v10
    } else {
        switch v3 {
        case 182:
            v11 := GenS2{F0: -94, F1: genFn0(2147483648, -723, (("foo bar" + "hello") + ""))}
            _ = v11
        case 4294967295:
            if (!(!((v3 >= v3) && (!true)))) {
                var v12 bool = ((((141 < 18446744073709551615) && true) || (("" + ((("hello" + ("a" + "\n\t")) + (((("hello" + "") + "\n\t") + "\n\t") + ("foo bar" + ("a" + "\n\t")))) + "foo bar")) != ("" + ("foo bar" + (((((("a" + (("a" + "hello") + "a")) + "a") + "hello") + "\n\t") + "hello") + ""))))) || true)
                var v13 float32 = 3.14159
                v2 += v2
                _ = v12
                _ = v13
            } else {
                for fi14 := 0; fi14 < 18; fi14++ {
                    if (!false) {
                        v1 = float64((float32(v1) / (float32(1e10) - (-2.25 - 3.14159))))
                        if false {
                            v3 += (fi14 + v3)
                        }
                        v3 = (v3 - v3)
                    } else {
                        if false {
                            var v15 int32 = 0
                            for fi16 := 0; fi16 < 10; fi16++ {
                                v17 := v5[("foo bar" + (("" + "\n\t") + "hello"))]
                                if v17 {
                                    var v18 int64 = 1
                                    if len(v4) > 0 {
                                        v4[0] = 42
                                    }
                                    var v19 int8 = 42
                                    var v20 int16 = -32768
                                    v2 -= (v2 % ((v2) | 1))
                                    if ((v1 + v1) == (((1.5 + (v1 + v1)) * v1) + float64(1e10))) {
                                        if ((v19 != v19) || v17) {
                                            var v21 uint32 = 2147483648
                                            var v22 string = "foo bar"
                                            _ = v21
                                            _ = v22
                                        }
                                        fmt.Println(v3, v1, v20)
                                    }
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                } else {
                                    fmt.Println(v1, v3, v17)
                                }
                                v3 += (v3 - fi14)
                                _ = fi16
                                _ = v17
                            }
                            _ = v15
                        }
                        var v23 float64 = -2.25
                        _ = v23
                    }
                    var v24 int8 = 100
                    _ = fi14
                    _ = v24
                }
                v5["\n\t"] = (42 == 973)
                var v25 int8 = -84
                _ = v25
            }
        case -731:
            var v26 uint32 = 264
            _ = v26
        }
        v27 := GenS2{F0: 1, F1: 42}
        v28 := GenS1{F0: 42}
        _ = v27
        _ = v28
    }
    fmt.Println(v3)
    var v29 uint8 = 0
    var v30 int16 = -1
    delete(v5, "foo bar")
    v29 += v29
    v1 = v1
    switch "\n\t" {
    case "foo bar":
        v31 := GenS3{F0: 0, F1: 2147483647}
        _ = v31
    case "":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        switch "" {
        case "a":
            var v32 uint32 = 163
            _ = v32
        case "":
            switch "\n\t" {
            case "a":
                var v33 uint32 = 845
                _ = v33
            case "\n\t":
                var v34 uint = (v2 / (((v2 - v2)) | 1))
                _ = v34
            default:
                if (false || ((v30 / ((v30) | 1)) >= v30)) {
                    fmt.Println(v2)
                    v35 := v5["hello"]
                    v30 += 32767
                    _ = v35
                } else {
                    var v36 uint32 = 639
                    _ = v36
                }
            }
        }
    }
    v1 -= float64((1e10 + (float32((0.1 * ((v1 * v1) + (v1 * v1)))) + (1e10 * 1.5))))
    var v37 int8 = -83
    var v38 int64 = 42
    switch "" {
    case "hello":
        var v39 float32 = 0.0
        _ = v39
    case "foo bar":
        if (32767 <= 1) {
            switch "\n\t" {
            case "\n\t":
                for fi40 := 0; fi40 < 19; fi40++ {
                    var v41 uint32 = 823
                    v38 = 949
                    _ = fi40
                    _ = v41
                }
            case "hello":
                v42 := GenS4{F0: 549}
                _ = v42
            default:
                v43 := GenS3{F0: 1, F1: uint32(v29)}
                _ = v43
            }
            fmt.Println(v1, v3)
            var v44 uint = v2
            _ = v44
        } else {
            switch 128 {
            case 1:
                var v45 bool = false
                _ = v45
            case 32768:
                v46 := GenS3{F0: 2, F1: 824}
                _ = v46
            default:
                if true {
                    if false {
                        delete(v5, ("a" + "\n\t"))
                        for fi47 := 0; fi47 < 20; fi47++ {
                            fmt.Println(v3, v2, v1)
                            _ = fi47
                        }
                        fmt.Println(v2, v29, v3)
                    } else {
                        for fi48 := 0; fi48 < 8; fi48++ {
                            v5[""] = false
                            _ = fi48
                        }
                    }
                }
            }
            v49 := GenS5{F0: 2, F1: (true && (false && false))}
            v37 -= (-(v37))
            _ = v49
        }
    default:
        var v50 uint32 = 2
        _ = v50
    }
    var v51 uint32 = 0
    for fi52 := 0; fi52 < 15; fi52++ {
        v53 := GenS2{F0: (^(v37)), F1: 255}
        for fi54 := 0; fi54 < 13; fi54++ {
            v55 := GenS1{F0: v38}
            v2 = (v2 | 42)
            v56 := GenS2{F0: (v37 & v37), F1: 406}
            _ = fi54
            _ = v55
            _ = v56
        }
        _ = fi52
        _ = v53
    }
    var v57 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v51
    _ = v57
}
