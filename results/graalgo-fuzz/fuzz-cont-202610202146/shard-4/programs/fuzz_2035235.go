package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 uint8; F1 uint64; F2 int }
func genFn0(p0 uint8, p1 int32, p2 int64) uint64 {
    if (4294967295 == 18446744073709551615) {
        v6 := GenS1{F0: "foo bar", F1: 2147483647}
        for fi7 := 0; fi7 < 1; fi7++ {
            for fi8 := 0; fi8 < 2; fi8++ {
                p1 = (p1 % ((p1) | 1))
                v9 := GenS1{F0: "\n\t", F1: fi8}
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 13; fi10++ {
                if ((!(!(65535 == 2147483647))) || (p1 != p1)) {
                    fmt.Println(fi7)
                    fmt.Println(p2)
                    v11 := GenS2{F0: "a"}
                    var v12 int64 = p2
                    v13 := GenS1{F0: "foo bar", F1: (fi7 % ((fi7) | 1))}
                    p0 = 127
                    _ = v11
                    _ = v12
                    _ = v13
                }
                var v14 int64 = (p2 & -1)
                _ = fi10
                _ = v14
            }
            _ = fi7
        }
        _ = v6
    } else {
        v15 := GenS3{F0: 927, F1: (p0 << ((0) & 63))}
        _ = v15
    }
    v16 := GenS4{F0: 4294967295}
    switch 0 {
    case 42:
        var v17 float64 = 1e10
        _ = v17
    default:
        var v18 uint = 554
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 1
}
func genFn1(p0 uint16, p1 uint16) bool {
    if true {
        p0 = 32767
        for fi19 := 0; fi19 < 20; fi19++ {
            v20 := []int{fi19, (fi19 | 754), (fi19 >> ((2) & 63))}
            p0 -= (p0 / ((p0) | 1))
            _ = fi19
            _ = v20
        }
        p0 = (p0 - 255)
    } else {
        v21 := []int{42, 1}
        fmt.Println(p1, p0)
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 float64 = -2.25
        var v24 uint32 = 2147483647
        if true {
            v25 := []int{318, (fi22 % ((fi22) | 1)), fi22, -2147483648}
            p0 += 861
            for fi26 := 0; fi26 < 15; fi26++ {
                if len(v25) > 0 {
                    v25[0] = fi26
                }
                _ = fi26
            }
            _ = v25
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    v27 := []int{315, -1, -1}
    v27 = append(v27, -2147483648)
    _ = p0
    _ = p1
    _ = v27
    return false
}

func main() {
    var v1 int32 = 1
    var v2 bool = (!false)
    var v3 uint8 = 12
    var v4 uint32 = 168
    var v5 int64 = 42
    v4 -= (^(v4))
    v4 -= (^(v4))
    v28 := []int{42}
    if len(v28) > 0 {
        v28[0] = 2147483647
    }
    if false {
        v4 -= (^(v4))
        for fi29 := 0; fi29 < 6; fi29++ {
            v1 += v1
            _ = fi29
        }
    }
    var v30 uint32 = v4
    for fi31 := 0; fi31 < 4; fi31++ {
        for fi32 := 0; fi32 < 13; fi32++ {
            switch "\n\t" {
            case "":
                var v33 float64 = (0.1 + 3.14159)
                _ = v33
            }
            v3 += 199
            fmt.Println(v30, v1)
            _ = fi32
        }
        _ = len(v28)
        var v34 uint32 = 287
        _ = fi31
        _ = v34
    }
    v35 := GenS4{F0: 801}
    v36 := GenS1{F0: ((("a" + "\n\t") + "hello") + ""), F1: 9223372036854775807}
    var v37 float64 = (0.1 / 1.0)
    var v38 float64 = ((v37 / v37) + v37)
    for fi39 := 0; fi39 < 16; fi39++ {
        fmt.Println(v1)
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            fmt.Println(v1)
        case "a":
            var v40 int = fi39
            _ = v40
        case "hello":
            if v2 {
                v37 -= float64(0.1)
                var v41 float32 = 3.14159
                _ = v41
            } else {
                for fi42 := 0; fi42 < 6; fi42++ {
                    switch "" {
                    case "hello":
                        if false {
                            fmt.Println(v1, v38)
                            if (65535 == 909) {
                                if v2 {
                                    v28 = append(v28, 28)
                                    var v43 uint8 = (v3 * (^(v3)))
                                    _ = v43
                                }
                            } else {
                                var v44 bool = true
                                _ = v44
                            }
                        } else {
                            _ = len(v28)
                            _ = len(v28)
                        }
                    case "foo bar":
                        var v45 uint64 = 355
                        _ = v45
                    case "\n\t":
                        var v46 float32 = float32(1e10)
                        _ = v46
                    default:
                        var v47 uint32 = 2
                        var v48 uint16 = 65535
                        v49 := GenS5{F0: "foo bar", F1: 1}
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                    v30 = v4
                    v38 = v37
                    _ = fi42
                }
                v50 := GenS6{F0: 2, F1: genFn0((v3 - v3), (v1 % ((int32(v5)) | 1)), v5), F2: (fi39 & int(v4))}
                var v51 int16 = 0
                _ = v50
                _ = v51
            }
        }
        for fi52 := 0; fi52 < 15; fi52++ {
            if (((("hello" + ("hello" + "foo bar")) + "") + "a") != ("\n\t" + "a")) {
                v38 = v37
                if (1000 >= 884) {
                    v3 += v3
                    fmt.Println(v30, fi52)
                } else {
                    switch (("hello" + "a") + "\n\t") {
                    case "foo bar":
                        v37 -= v38
                    case "a":
                        v28 = append(v28, 4294967295)
                        switch "foo bar" {
                        case "foo bar":
                            var v53 int16 = 32767
                            _ = v53
                        }
                        v38 = -2.25
                        if v2 {
                            v2 = false
                            v2 = ((!v2) && v2)
                            v1 += v1
                        }
                        v3 = v3
                    case "\n\t":
                        var v54 string = ("" + "")
                        _ = v54
                    }
                }
                switch (^(fi52)) {
                case -2147483648:
                    fmt.Println(v1, v2)
                case -813:
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                }
            }
            if (1.5 >= 0.0) {
                var v55 int16 = -712
                _ = v55
            }
            v56 := GenS1{F0: "foo bar", F1: fi52}
            _ = fi52
            _ = v56
        }
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
