package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint32, p1 int64, p2 uint) uint32 {
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := []int{(fi7 % ((-9223372036854775808) | 1))}
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 11
    v10 := GenS1{F0: "foo bar"}
    if false {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint = p2
            var v13 uint = (p2 << ((2) & 63))
            _ = fi11
            _ = v12
            _ = v13
        }
        v14 := []int{4294967295, -696, 9223372036854775807}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return p0
}
func genFn1(p0 int8, p1 uint8) string {
    if true {
        var v15 uint8 = p1
        for fi16 := 0; fi16 < 8; fi16++ {
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := GenS1{F0: ((("a" + (("a" + "a") + "\n\t")) + ("a" + "\n\t")) + ("a" + "hello"))}
                v19 := GenS1{F0: ("hello" + ("hello" + "a"))}
                _ = fi17
                _ = v18
                _ = v19
            }
            if (0 <= 973) {
                for fi20 := 0; fi20 < 2; fi20++ {
                    var v21 string = (((("" + "") + ("" + "")) + "a") + "a")
                    var v22 uint32 = 2147483648
                    v23 := GenS2{F0: p1}
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                fmt.Println(p0, p1)
                if true {
                    if (((!false) && (1 != -954)) || true) {
                        v24 := GenS2{F0: v15}
                        if true {
                            switch ("\n\t" + "hello") {
                            case "a":
                                fmt.Println(v15)
                            case "\n\t":
                                fmt.Println(p1, p1)
                            }
                        }
                        switch "hello" {
                        case "foo bar":
                            v25 := []int{fi16, (fi16 % (((fi16 * fi16)) | 1)), (-(fi16)), (fi16 / ((fi16) | 1))}
                            _ = v25
                        default:
                            if (!(false && true)) {
                                p1 += (v15 | (p1 - (v15 ^ (v15 << ((0) & 63)))))
                                if (!true) {
                                    var v26 float32 = (((float32(0.0) + (1.5 / 3.14159)) * 1e10) + 0.0)
                                    var v27 int8 = (p0 % (((p0 - (p0 | -15))) | 1))
                                    if (!(!false)) {
                                        v28 := GenS1{F0: "a"}
                                        p0 -= 0
                                        v26 = v26
                                        _ = v28
                                    } else {
                                        v27 -= (p0 % (((v27 + int8(fi16))) | 1))
                                        switch (-(fi16)) {
                                        case 2147483647:
                                            fmt.Println(p0, p0)
                                        case -875:
                                            if (2 <= 732) {
                                                p1 = v15
                                            }
                                        }
                                    }
                                    _ = v26
                                    _ = v27
                                }
                            } else {
                                for fi29 := 0; fi29 < 20; fi29++ {
                                    if true {
                                        p0 = 100
                                        v30 := GenS3{F0: -1}
                                        v31 := []int{(fi16 % (((fi29 - fi29)) | 1)), fi16}
                                        _ = v30
                                        _ = v31
                                    } else {
                                        var v32 int8 = (^(p0))
                                        var v33 uint16 = 922
                                        v34 := GenS2{F0: 23}
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    }
                                    _ = fi29
                                }
                                for fi35 := 0; fi35 < 14; fi35++ {
                                    p1 = v15
                                    v36 := []int{0}
                                    _ = fi35
                                    _ = v36
                                }
                                var v37 float32 = float32(0.1)
                                _ = v37
                            }
                        }
                        _ = v24
                    } else {
                        v15 = p1
                    }
                    fmt.Println(v15)
                    for fi38 := 0; fi38 < 8; fi38++ {
                        var v39 uint8 = v15
                        p0 = 42
                        _ = fi38
                        _ = v39
                    }
                }
            } else {
                fmt.Println(p0)
            }
            _ = fi16
        }
        _ = v15
    } else {
        for fi40 := 0; fi40 < 6; fi40++ {
            if true {
                p0 = (p0 * 100)
                v41 := []int{(fi40 & fi40), (fi40 / (((fi40 + (fi40 ^ 127))) | 1)), fi40, fi40}
                _ = v41
            }
            _ = fi40
        }
        var v42 string = "a"
        var v43 float64 = 0.0
        _ = v42
        _ = v43
    }
    switch -349 {
    case -156:
        var v44 uint = 553
        _ = v44
    default:
        var v45 int8 = (p0 >> ((255) & 63))
        _ = v45
    }
    var v46 uint16 = 255
    v46 -= (v46 | 0)
    _ = p0
    _ = p1
    _ = v46
    return ((("foo bar" + "a") + "hello") + "\n\t")
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"", (("\n\t" + "") + "hello"), (("a" + "\n\t") + "")}
    var v3 uint64 = 4294967295
    var v4 int = -997
    var v5 float32 = 0.1
    var v6 int8 = 0
    fmt.Println(v3)
    switch v4 {
    case 0:
        v3 = v3
    }
    switch (v4 ^ (-(v4))) {
    case -1:
        for fi47 := 0; fi47 < 17; fi47++ {
            v4 = (fi47 % ((fi47) | 1))
            _ = fi47
        }
    case 128:
        v5 += 0.1
    case -2147483648:
        v48 := GenS4{F0: 1, F1: false}
        _ = v48
    default:
        _ = len(v2)
    }
    for fi49 := 0; fi49 < 5; fi49++ {
        fmt.Println(v6)
        fmt.Println(v6)
        _ = fi49
    }
    v50 := GenS3{F0: 1}
    if (1e10 >= float64(v5)) {
        v51 := GenS5{F0: 1, F1: 137}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v51
    }
    var v52 uint = 1
    var v53 uint32 = 65535
    v52 -= v52
    var v54 uint = 961
    if (1.5 <= (0.1 * 1.5)) {
        delete(v1, "\n\t")
    } else {
        v5 = 0.1
        var v55 float64 = 0.0
        _ = v55
    }
    v52 -= (v54 / ((v52) | 1))
    v3 = 363
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
