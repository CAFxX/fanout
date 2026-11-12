package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 int16; F2 bool }
func genFn0(p0 int16, p1 uint16, p2 uint8) uint {
    v5 := GenS1{F0: 42, F1: 255, F2: "\n\t"}
    var v6 uint64 = 96
    v7 := []int{613, int(v6), 776}
    fmt.Println(p2, p2)
    switch "hello" {
    case "hello":
        if len(v7) > 0 {
            v7[0] = -9223372036854775808
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 278
}
func genFn1(p0 int16, p1 int64, p2 uint16) uint8 {
    var v8 uint8 = 1
    if true {
        var v9 int32 = 42
        _ = v9
    } else {
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (v8 + 2)
}

func main() {
    v1 := []string{"\n\t", "\n\t"}
    v2 := []int{755}
    var v3 int64 = 0
    var v4 uint8 = 249
    if (657 == 18446744073709551615) {
        if false {
            for fi10 := 0; fi10 < 11; fi10++ {
                v3 += -1
                switch "" {
                case "hello":
                    if true {
                        _ = len(v1)
                        v11 := GenS1{F0: fi10, F1: 1, F2: "foo bar"}
                        if len(v2) > 0 {
                            v2[0] = (fi10 - (fi10 << ((65535) & 63)))
                        }
                        _ = v11
                    }
                }
                _ = fi10
            }
        }
        var v12 int8 = 127
        _ = v12
    } else {
        fmt.Println(v4, v3)
        if (0.0 <= 1.5) {
            var v13 uint8 = 2
            var v14 float32 = (3.14159 + (1.5 - (0.0 + 3.14159)))
            v15 := GenS2{F0: -1}
            _ = v13
            _ = v14
            _ = v15
        }
        v16 := GenS2{F0: v3}
        _ = v16
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v17 int32 = 1
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := GenS3{F0: 1, F1: v17}
        switch (fi18 * (fi18 - fi18)) {
        case 128:
            v3 -= (v3 ^ v3)
        case 506:
            var v20 uint64 = 42
            _ = v20
        }
        if (false && false) {
            v4 -= 0
        } else {
            if ((((((("\n\t" + (((("" + ((("" + "") + "\n\t") + "")) + "a") + "\n\t") + "")) + "a") + "foo bar") + ("" + ("foo bar" + "a"))) + (("\n\t" + ("a" + ("hello" + ((("foo bar" + ((("\n\t" + "") + (("\n\t" + "") + ("foo bar" + "\n\t"))) + "hello")) + (("hello" + ("\n\t" + "\n\t")) + "\n\t")) + (("" + ("foo bar" + "foo bar")) + "hello"))))) + "")) + "foo bar") == ((("foo bar" + "a") + "") + ("foo bar" + "\n\t"))) {
                for fi21 := 0; fi21 < 3; fi21++ {
                    if false {
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                    }
                    var v22 uint64 = 42
                    _ = len(v2)
                    _ = fi21
                    _ = v22
                }
            } else {
                v23 := GenS1{F0: fi18, F1: 255, F2: (("\n\t" + "") + "foo bar")}
                _ = v23
            }
            v17 = (v17 & v17)
            switch "hello" {
            case "a":
                v17 = v17
            default:
                v24 := GenS1{F0: fi18, F1: 867, F2: "hello"}
                _ = v24
            }
        }
        _ = fi18
        _ = v19
    }
    v4 += v4
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    var v25 uint32 = 155
    if (((0 == 363) && true) || false) {
        if true {
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            var v26 uint16 = 528
            _ = v26
        } else {
            v27 := GenS4{F0: 353}
            _ = v27
        }
        if (true && (!false)) {
            switch "" {
            case "\n\t":
                if (false || (!(!true))) {
                    var v28 int64 = v3
                    for fi29 := 0; fi29 < 18; fi29++ {
                        if true {
                            fmt.Println(v25, v3)
                            for fi30 := 0; fi30 < 18; fi30++ {
                                var v31 int8 = 1
                                _ = fi30
                                _ = v31
                            }
                            var v32 int32 = v17
                            _ = v32
                        } else {
                            v1 = append(v1, "\n\t")
                            var v33 uint64 = 0
                            fmt.Println(v3, v17, v28)
                            _ = v33
                        }
                        fmt.Println(v3, v4, fi29)
                        if false {
                            v34 := GenS1{F0: fi29, F1: 65535, F2: ((("hello" + "a") + "hello") + "\n\t")}
                            _ = v34
                        }
                        _ = fi29
                    }
                    switch "a" {
                    case "foo bar":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = v28
                } else {
                    v35 := GenS1{F0: int(v17), F1: 18446744073709551615, F2: "a"}
                    for fi36 := 0; fi36 < 14; fi36++ {
                        fmt.Println(fi36, v25, v17)
                        _ = len(v1)
                        v3 -= (v3 & (v3 / (((-(v3))) | 1)))
                        _ = fi36
                    }
                    _ = v35
                }
            default:
                for fi37 := 0; fi37 < 15; fi37++ {
                    v2 = append(v2, int(v17))
                    _ = fi37
                }
            }
            v38 := GenS5{F0: 100, F1: 32767, F2: (44 == 64)}
            var v39 int8 = 69
            _ = v38
            _ = v39
        }
        for fi40 := 0; fi40 < 7; fi40++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi40
        }
    }
    if len(v2) > 0 {
        v2[0] = 32767
    }
    var v41 int16 = -1
    _ = len(v2)
    v42 := GenS5{F0: 1, F1: (v41 ^ v41), F2: true}
    for fi43 := 0; fi43 < 19; fi43++ {
        v25 = v25
        if (false || (-2.25 >= 3.14159)) {
            v1 = append(v1, ("\n\t" + ("\n\t" + ("\n\t" + "a"))))
            var v44 float32 = (1e10 / -2.25)
            _ = v44
        } else {
            switch "\n\t" {
            case "a":
                for fi45 := 0; fi45 < 7; fi45++ {
                    var v46 bool = false
                    _ = fi45
                    _ = v46
                }
            case "":
                if true {
                    fmt.Println(v17)
                    fmt.Println(v25)
                }
            case "\n\t":
                v25 -= v25
            default:
                var v47 int32 = (v17 - v17)
                _ = v47
            }
            v41 += (v41 - v41)
            v41 -= (v41 & v41)
        }
        switch ("" + ("foo bar" + "foo bar")) {
        case "":
            if ((!(v41 <= 1)) || (false && (!true))) {
                if (v4 <= (v4 >> ((uint(v25)) & 63))) {
                    if (false && ((v25 | 113) < (v25 / ((uint32(v4)) | 1)))) {
                        var v48 uint32 = 65535
                        _ = v48
                    }
                } else {
                    v49 := GenS2{F0: v3}
                    v25 += v25
                    _ = v49
                }
            }
        }
        _ = fi43
    }
    if true {
        switch "\n\t" {
        case "hello":
            var v50 float32 = 0.0
            _ = v50
        default:
            for fi51 := 0; fi51 < 7; fi51++ {
                v3 -= v3
                fmt.Println(v3, v4, fi51)
                fmt.Println(v4)
                _ = fi51
            }
        }
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v25, v25)
    case "":
        switch "foo bar" {
        case "hello":
            for fi52 := 0; fi52 < 5; fi52++ {
                v41 = v41
                v53 := GenS5{F0: 42, F1: (v41 & (v41 * v41)), F2: ((((!(690 >= 1)) || true) && (true && false)) && true)}
                var v54 uint8 = v4
                _ = fi52
                _ = v53
                _ = v54
            }
        default:
            var v55 int = 0
            _ = v55
        }
    default:
        if (127 == -128) {
            var v56 uint8 = v4
            fmt.Println(v41)
            _ = v56
        } else {
            for fi57 := 0; fi57 < 3; fi57++ {
                v3 -= -103
                var v58 uint = 18446744073709551615
                _ = fi57
                _ = v58
            }
            var v59 float64 = -2.25
            if len(v1) > 0 {
                v1[0] = ("foo bar" + "hello")
            }
            _ = v59
        }
    }
    var v60 uint16 = 1
    v4 += (v4 << ((65535) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v60
}
