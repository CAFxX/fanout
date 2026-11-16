package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 string; F2 int32 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint16; F2 int }
type GenS6 struct { F0 uint; F1 bool; F2 bool }
func genFn0(p0 string, p1 uint) int {
    switch p0 {
    case "a":
        for fi4 := 0; fi4 < 1; fi4++ {
            for fi5 := 0; fi5 < 10; fi5++ {
                v6 := GenS1{F0: 127}
                _ = fi5
                _ = v6
            }
            _ = fi4
        }
    case "foo bar":
        v7 := GenS2{F0: 2, F1: ((p0 + (((p0 + (p0 + p0)) + p0) + p0)) + "hello"), F2: -1}
        _ = v7
    default:
        v8 := []int{9223372036854775807, 32767}
        _ = v8
    }
    v9 := GenS2{F0: 1, F1: ((p0 + p0) + "\n\t"), F2: 0}
    _ = p0
    _ = p1
    _ = v9
    return -39
}
func genFn1(p0 bool, p1 int16) string {
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            fmt.Println(fi10, p0)
            var v12 uint32 = 2147483648
            _ = fi11
            _ = v12
        }
        fmt.Println(p0)
        p1 += p1
        _ = fi10
    }
    if p0 {
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := GenS3{F0: "\n\t", F1: 889}
            if (fi13 > genFn0("", 1)) {
                v15 := GenS3{F0: "\n\t", F1: 285}
                _ = v15
            }
            _ = fi13
            _ = v14
        }
        var v16 string = "\n\t"
        _ = v16
    }
    _ = p0
    _ = p1
    return ""
}

func main() {
    v1 := []string{"hello", "", (("foo bar" + "\n\t") + ""), (((("foo bar" + ("" + ("" + ("\n\t" + (("hello" + ("" + ("foo bar" + "\n\t"))) + "\n\t"))))) + ("foo bar" + "hello")) + "") + (("a" + "foo bar") + (("hello" + "hello") + "hello")))}
    var v2 int = -2147483648
    v3 := make(map[string]uint32)
    var v17 float32 = 0.0
    delete(v3, ("" + "\n\t"))
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := GenS4{F0: "", F1: 277}
        switch "\n\t" {
        case "":
            var v20 uint64 = 2
            _ = v20
        case "\n\t":
            for fi21 := 0; fi21 < 16; fi21++ {
                v2 -= -1
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v22 int8 = 100
                _ = fi21
                _ = v22
            }
        case "a":
            for fi23 := 0; fi23 < 7; fi23++ {
                if ((("" + "a") < "") || (-579 == 42)) {
                    fmt.Println(fi23)
                }
                _ = fi23
            }
        default:
            if (!false) {
                if (!((((3.14159 * float64((v17 + 1.5))) * 1e10) - (1e10 - 1.5)) <= float64((v17 / 1.0)))) {
                    delete(v3, ("foo bar" + ("hello" + "")))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    for fi24 := 0; fi24 < 19; fi24++ {
                        fmt.Println(fi24, v17, fi18)
                        switch "hello" {
                        case "\n\t":
                            if (float32(1.5) > (1.5 / ((v17 * 0.1) * 0.0))) {
                                var v25 int32 = 42
                                var v26 int64 = -9223372036854775808
                                fmt.Println(v17)
                                _ = v25
                                _ = v26
                            }
                        case "":
                            fmt.Println(v2)
                        default:
                            v17 -= (v17 + v17)
                        }
                        v2 += (fi18 >> ((255) & 63))
                        _ = fi24
                    }
                }
                var v27 uint = 65535
                _ = v27
            } else {
                if (287 == 42) {
                    v3["\n\t"] = 2
                    v1 = append(v1, "a")
                    v28 := GenS3{F0: "", F1: 233}
                    _ = v28
                }
            }
        }
        _ = fi18
        _ = v19
    }
    var v29 uint = 704
    switch "\n\t" {
    case "foo bar":
        var v30 uint64 = uint64(v29)
        _ = v30
    }
    if false {
        v2 += v2
        var v31 uint16 = 255
        fmt.Println(v29, v29, v31)
        _ = v31
    } else {
        for fi32 := 0; fi32 < 1; fi32++ {
            v2 = (^(v2))
            v2 = v2
            _ = fi32
        }
        switch (v2 - (v2 ^ v2)) {
        case -286:
            v33 := GenS5{F0: 2, F1: uint16(v29), F2: -818}
            _ = v33
        case -965:
            if true {
                if ((!false) && (2147483647 != -2147483648)) {
                    for fi34 := 0; fi34 < 18; fi34++ {
                        var v35 int = genFn0((("\n\t" + ("a" + "\n\t")) + ""), (^(v29)))
                        _ = fi34
                        _ = v35
                    }
                    switch v2 {
                    case 2147483647:
                        var v36 uint8 = 47
                        _ = v36
                    }
                } else {
                    for fi37 := 0; fi37 < 3; fi37++ {
                        v17 += 3.14159
                        _ = fi37
                    }
                    var v38 uint32 = 2147483648
                    _ = v38
                }
            }
        }
        v17 = v17
    }
    for fi39 := 0; fi39 < 19; fi39++ {
        v1 = append(v1, "hello")
        var v40 string = "foo bar"
        v40 = v40
        _ = fi39
        _ = v40
    }
    delete(v3, "\n\t")
    for fi41 := 0; fi41 < 6; fi41++ {
        for fi42 := 0; fi42 < 19; fi42++ {
            var v43 int32 = 0
            var v44 bool = (!true)
            var v45 uint64 = 42
            _ = fi42
            _ = v43
            _ = v44
            _ = v45
        }
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        _ = fi41
    }
    v46 := GenS6{F0: 42, F1: true, F2: (true && (210 != 856))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v46
}
