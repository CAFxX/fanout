package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 string }
type GenS2 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 uint; F2 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 uint16 }
func genFn0(p0 int32, p1 uint) int16 {
    p0 = (p0 % (((p0 | p0)) | 1))
    v6 := GenS1{F0: 1, F1: 2, F2: ""}
    _ = p0
    _ = p1
    _ = v6
    return 42
}
func genFn1(p0 uint16, p1 uint64, p2 string) int64 {
    for fi7 := 0; fi7 < 15; fi7++ {
        if false {
            var v8 int64 = -1
            switch "" {
            case "foo bar":
                for fi9 := 0; fi9 < 20; fi9++ {
                    var v10 float64 = (((1e10 * (1.5 * 3.14159)) - 0.1) * -2.25)
                    _ = fi9
                    _ = v10
                }
            }
            _ = v8
        }
        _ = fi7
    }
    fmt.Println(p2)
    fmt.Println(p1)
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint = 466
        _ = fi11
        _ = v12
    }
    v13 := []int{32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 407
}

func main() {
    var v1 int64 = 1
    var v2 int = -1
    var v3 string = ""
    var v4 string = "\n\t"
    var v5 int64 = (-(v1))
    v14 := []int{v2, v2}
    v15 := GenS1{F0: 0, F1: 244, F2: "a"}
    fmt.Println(v2)
    switch v3 {
    case "foo bar":
        for fi16 := 0; fi16 < 5; fi16++ {
            v5 += (v5 ^ (v5 + (v1 - v5)))
            _ = fi16
        }
    case "a":
        var v17 uint16 = 1
        _ = v17
    case "\n\t":
        _ = len(v14)
    default:
        var v18 uint32 = 83
        _ = v18
    }
    v19 := GenS2{F0: 0, F1: -1, F2: 0}
    v3 = v3
    var v20 uint32 = 2147483647
    v1 = v5
    var v21 uint8 = 127
    v22 := GenS2{F0: 2, F1: 2147483647, F2: 2}
    v2 -= v2
    switch v3 {
    case "foo bar":
        if true {
            var v23 int = v2
            var v24 string = "hello"
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 20; fi25++ {
                v26 := GenS3{F0: int8(v21)}
                v27 := GenS1{F0: 750, F1: v20, F2: v4}
                _ = fi25
                _ = v26
                _ = v27
            }
        }
    default:
        if false {
            var v28 uint16 = 42
            _ = v28
        }
    }
    v29 := GenS4{F0: 127, F1: 964, F2: -677}
    var v30 bool = false
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    if ((!v30) || v30) {
        v1 += -1
        if (v2 >= (v2 - v2)) {
            v3 = ("hello" + (v4 + ((((v4 + "\n\t") + v4) + "hello") + v3)))
            for fi31 := 0; fi31 < 2; fi31++ {
                var v32 int16 = 1000
                _ = fi31
                _ = v32
            }
            var v33 int32 = 0
            _ = v33
        }
        switch "a" {
        case "foo bar":
            var v34 uint16 = 1
            _ = v34
        }
    }
    if false {
        var v35 string = "foo bar"
        for fi36 := 0; fi36 < 2; fi36++ {
            for fi37 := 0; fi37 < 20; fi37++ {
                switch (fi37 << ((18446744073709551615) & 63)) {
                case 42:
                    fmt.Println(fi36, v5)
                }
                var v38 float64 = 1e10
                v3 = ("foo bar" + v4)
                var v39 int32 = -1
                if (v21 != v21) {
                    if (false && ((((-(fi36)) != (^(fi37))) && (true || true)) || v30)) {
                        v40 := GenS5{F0: v4}
                        v41 := GenS6{F0: 1, F1: 1}
                        _ = v40
                        _ = v41
                    }
                    if false {
                        switch (fi36 * (fi37 << ((4294967295) & 63))) {
                        case 2147483647:
                            _ = len(v14)
                        case 349:
                            if (-554 == 0) {
                                switch 32767 {
                                case 818:
                                    v14 = append(v14, (fi37 << ((65535) & 63)))
                                default:
                                    if len(v14) > 0 {
                                        v14[0] = (v2 ^ (^(fi37)))
                                    }
                                }
                            }
                        case 4294967295:
                            fmt.Println(v21)
                        default:
                            if v30 {
                                v35 = v35
                                _ = len(v14)
                            }
                        }
                    }
                    if (0 < 42) {
                        v1 += (v1 / ((v5) | 1))
                        v39 += -382
                        var v42 uint = 255
                        _ = v42
                    }
                } else {
                    v43 := GenS2{F0: 2, F1: (v39 & v39), F2: (v20 << ((2) & 63))}
                    v2 -= v2
                    fmt.Println(fi36, fi36, v1)
                    _ = v43
                }
                _ = len(v14)
                fmt.Println(v3, v1, v39)
                var v44 int32 = -1
                var v45 int8 = 100
                _ = fi37
                _ = v38
                _ = v39
                _ = v44
                _ = v45
            }
            _ = fi36
        }
        _ = v35
    } else {
        v46 := GenS1{F0: 465, F1: v20, F2: "foo bar"}
        var v47 uint8 = 173
        _ = v46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
