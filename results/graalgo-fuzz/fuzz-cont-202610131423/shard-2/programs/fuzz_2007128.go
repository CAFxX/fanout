package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 uint32; F2 string }
type GenS7 struct { F0 int32; F1 uint; F2 uint16 }
type GenS8 struct { F0 uint16; F1 bool }
type GenS9 struct { F0 int16; F1 uint8; F2 bool }
type GenS10 struct { F0 bool; F1 string }
func genFn0(p0 uint, p1 uint, p2 bool) int32 {
    if (!p2) {
        v6 := GenS1{F0: p2}
        v7 := GenS2{F0: 42}
        p2 = (float32(1.5) != 0.0)
        _ = v6
        _ = v7
    } else {
        if ((float64((float32(-2.25) - 1e10)) == 0.1) || p2) {
            if (0 != 9223372036854775807) {
                var v8 uint32 = 0
                _ = v8
            }
        } else {
            p0 = p1
            v9 := GenS2{F0: 42}
            v10 := GenS3{F0: -69}
            _ = v10
            _ = v9
        }
        if (34 == 255) {
            switch "\n\t" {
            case "a":
                for fi11 := 0; fi11 < 13; fi11++ {
                    v12 := GenS3{F0: 42}
                    _ = fi11
                    _ = v12
                }
            case "foo bar":
                v13 := []int{-2147483648}
                _ = v13
            case "hello":
                v14 := GenS1{F0: true}
                _ = v14
            }
            var v15 string = "\n\t"
            for fi16 := 0; fi16 < 7; fi16++ {
                if (-160 <= 42) {
                    var v17 string = v15
                    _ = v17
                } else {
                    if (3.14159 > (0.1 - (((0.0 / (0.1 - (0.1 * 1e10))) - 1.5) / 1e10))) {
                        if p2 {
                            p0 = (p0 & (^(p1)))
                            p0 += (p0 - (p0 | (p1 ^ 18446744073709551615)))
                            p2 = ((32767 != 42) || true)
                        } else {
                            var v18 uint8 = 2
                            var v19 uint32 = 65535
                            switch fi16 {
                            case 324:
                                v18 = 128
                            }
                            _ = v18
                            _ = v19
                        }
                    }
                }
                v20 := GenS3{F0: 9}
                fmt.Println(fi16, p1, p0)
                _ = fi16
                _ = v20
            }
            _ = v15
        } else {
            if (-1 <= 964) {
                if false {
                    p1 += (p1 | p0)
                    var v21 int = 32768
                    var v22 int8 = 31
                    _ = v21
                    _ = v22
                }
            } else {
                for fi23 := 0; fi23 < 8; fi23++ {
                    v24 := []int{fi23, fi23}
                    if (0.1 >= 0.0) {
                        fmt.Println(p0, p1, p0)
                        if p2 {
                            fmt.Println(p0, fi23, p0)
                        } else {
                            if (-32768 > 1000) {
                                v24 = append(v24, fi23)
                                p2 = p2
                                if p2 {
                                    for fi25 := 0; fi25 < 11; fi25++ {
                                        p2 = p2
                                        p0 = (p1 | (p0 >> ((p1) & 63)))
                                        v24 = append(v24, fi23)
                                        if p2 {
                                            fmt.Println(fi25, fi23, p2)
                                            p0 += p0
                                            switch "\n\t" {
                                            case "foo bar":
                                                if true {
                                                    v26 := GenS3{F0: 127}
                                                    fmt.Println(p1)
                                                    if (0 < 1) {
                                                        var v27 string = "foo bar"
                                                        p1 += 42
                                                        _ = v27
                                                    }
                                                    _ = v26
                                                }
                                            case "\n\t":
                                                switch "" {
                                                case "\n\t":
                                                    if (("foo bar" <= "\n\t") && (!p2)) {
                                                        var v28 uint32 = 2147483647
                                                        var v29 float64 = 0.1
                                                        switch (fi25 | 426) {
                                                        case -9223372036854775808:
                                                            v30 := GenS4{F0: (v28 - v28)}
                                                            _ = v30
                                                        }
                                                        _ = v28
                                                        _ = v29
                                                    } else {
                                                        if len(v24) >= 2 {
                                                            v24 = v24[1:len(v24)]
                                                        }
                                                        var v31 uint8 = 127
                                                        _ = v31
                                                    }
                                                default:
                                                    v32 := GenS5{F0: ""}
                                                    _ = v32
                                                }
                                            }
                                        } else {
                                            var v33 string = "\n\t"
                                            var v34 string = v33
                                            var v35 float32 = 3.14159
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                        _ = fi25
                                    }
                                    var v36 float64 = -2.25
                                    _ = v36
                                }
                            } else {
                                for fi37 := 0; fi37 < 15; fi37++ {
                                    var v38 int16 = 1
                                    p0 = (^(p1))
                                    var v39 int64 = -9223372036854775808
                                    _ = fi37
                                    _ = v38
                                    _ = v39
                                }
                            }
                            if p2 {
                                p0 = p1
                            }
                            for fi40 := 0; fi40 < 12; fi40++ {
                                var v41 uint = (p0 & (p0 << ((p1) & 63)))
                                var v42 uint = v41
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                                _ = fi40
                                _ = v41
                                _ = v42
                            }
                        }
                        v43 := GenS4{F0: 228}
                        _ = v43
                    } else {
                        p1 += 2
                    }
                    var v44 float64 = (3.14159 + (1.5 + float64(float32(float64(float32(3.14159))))))
                    _ = fi23
                    _ = v24
                    _ = v44
                }
                switch -1 {
                case 942:
                    if p2 {
                        if (42 < 1) {
                            fmt.Println(p2, p2, p1)
                        } else {
                            v45 := []int{712, 4294967295, 32767, 32767}
                            _ = v45
                        }
                    } else {
                        var v46 int16 = -32768
                        _ = v46
                    }
                case 578:
                    var v47 int16 = 389
                    _ = v47
                case 42:
                    v48 := GenS3{F0: 23}
                    _ = v48
                }
            }
        }
        var v49 int8 = 0
        _ = v49
    }
    p1 = (p0 + (p0 | (p0 | (p1 % (((p0 ^ (^(p0)))) | 1)))))
    var v50 uint8 = 174
    var v51 uint = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v50
    _ = v51
    return 0
}

func main() {
    var v1 float64 = (3.14159 + -2.25)
    var v2 uint64 = 1
    v3 := make(map[string]bool)
    var v4 float32 = 0.1
    v5 := []string{"\n\t", (("\n\t" + "hello") + ((("a" + ("a" + "")) + "\n\t") + ("a" + ("hello" + "foo bar"))))}
    for fi52 := 0; fi52 < 1; fi52++ {
        v53 := GenS6{F0: 18446744073709551615, F1: uint32(v2), F2: "foo bar"}
        v54 := GenS6{F0: 4294967295, F1: 2147483647, F2: "hello"}
        if (!(false && (1 > 1))) {
            switch (fi52 - fi52) {
            case -321:
                if true {
                    v1 += v1
                } else {
                    fmt.Println(v4)
                    var v55 float32 = v4
                    var v56 uint = 2
                    _ = v55
                    _ = v56
                }
            case 843:
                fmt.Println(v1, v4)
            case 0:
                for fi57 := 0; fi57 < 17; fi57++ {
                    switch "a" {
                    case "hello":
                        v58 := GenS7{F0: 2147483647, F1: 4294967295, F2: 255}
                        _ = v58
                    case "\n\t":
                        v59 := GenS8{F0: 65535, F1: false}
                        _ = v59
                    default:
                        v60 := GenS9{F0: -517, F1: 255, F2: (42 != -760)}
                        _ = v60
                    }
                    var v61 int64 = -1
                    _ = fi57
                    _ = v61
                }
            }
        } else {
            fmt.Println(v1, v1)
            v1 = (v1 * v1)
            v4 += 0.0
        }
        _ = fi52
        _ = v53
        _ = v54
    }
    var v62 uint32 = 2147483647
    v4 += v4
    for fi63 := 0; fi63 < 10; fi63++ {
        v1 -= v1
        for fi64 := 0; fi64 < 12; fi64++ {
            var v65 int32 = genFn0(18446744073709551615, 18446744073709551615, false)
            _ = fi64
            _ = v65
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi63
    }
    v4 = float32(v1)
    v66 := GenS10{F0: (true || (("hello" + (((("" + "a") + "a") + "a") + "hello")) >= "")), F1: "foo bar"}
    var v67 uint32 = v62
    v68 := GenS6{F0: v2, F1: (v67 >> ((65535) & 63)), F2: "foo bar"}
    var v69 int16 = 1000
    var v70 uint32 = (v62 & 65535)
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    v2 = (v2 ^ v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v62:", v62)
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
