package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int8; F1 uint16; F2 int64 }
type GenS4 struct { F0 int32; F1 bool; F2 int64 }
type GenS5 struct { F0 int64; F1 int64; F2 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool; F1 int8 }
func genFn0() string {
    v5 := GenS1{F0: (false && true)}
    if false {
        var v6 int8 = -1
        _ = v6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        if false {
            var v8 int32 = -2147483648
            _ = v8
        }
        var v9 int = fi7
        _ = fi7
        _ = v9
    }
    _ = v5
    return ((("foo bar" + "") + "hello") + ("foo bar" + ""))
}
func genFn1(p0 int8, p1 int8, p2 bool) int32 {
    if (p1 != (p1 * (p1 / ((p0) | 1)))) {
        p1 = 100
        fmt.Println(p1)
        v10 := GenS2{F0: 2147483647, F1: "hello"}
        _ = v10
    }
    var v11 uint16 = 42
    var v12 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return int32(v11)
}

func main() {
    var v1 uint16 = 32767
    v2 := []int{42, 540, 32767}
    var v3 int64 = -302
    v4 := make(map[int]int)
    v13 := GenS1{F0: true}
    v3 = (v3 - (^(v3)))
    v1 -= 429
    v3 += 0
    var v14 int8 = -1
    v15 := GenS3{F0: (v14 ^ v14), F1: 121, F2: (v3 ^ v3)}
    if false {
        switch "" {
        case "":
            v16 := v4[42]
            _ = v16
        }
        if true {
            if ((^(v14)) > (^(v14))) {
                v3 -= -1
                if (!false) {
                    if len(v2) > 0 {
                        v2[0] = -996
                    }
                    switch -322 {
                    case 127:
                        var v17 uint32 = 65535
                        _ = v17
                    case 1:
                        if false {
                            v18 := GenS4{F0: genFn1(-128, v14, ((1.5 - 0.1) != float32(0.1))), F1: (true || false), F2: (v3 % ((v3) | 1))}
                            _ = v18
                        }
                    }
                } else {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        for fi20 := 0; fi20 < 4; fi20++ {
                            v3 += (-(v3))
                            switch "foo bar" {
                            case "hello":
                                if (!(!(((0.1 - (float32(((1e10 + float64(3.14159)) * float64(1.5))) / (1e10 * (1.5 * (((float32((0.0 - 1.5)) * float32(0.0)) / 1.5) - (0.1 * (3.14159 / -2.25))))))) * (float32((-2.25 * 0.0)) * -2.25)) > -2.25))) {
                                    switch ("" + ("foo bar" + ("a" + "\n\t"))) {
                                    case "hello":
                                        v14 += v14
                                    case "a":
                                        if true {
                                            switch (("\n\t" + ("foo bar" + "hello")) + "") {
                                            case "a":
                                                var v21 uint64 = 2
                                                _ = v21
                                            case "foo bar":
                                                fmt.Println(v14)
                                            }
                                        } else {
                                            v22 := GenS4{F0: 2147483647, F1: true, F2: v3}
                                            _ = v22
                                        }
                                    }
                                }
                            case "\n\t":
                                fmt.Println(v3, fi20)
                            case "":
                                delete(v4, (-(fi19)))
                            }
                            var v23 uint16 = 42
                            _ = fi20
                            _ = v23
                        }
                        v3 = -9223372036854775808
                        _ = fi19
                    }
                }
            }
            fmt.Println(v3, v1)
        } else {
            if false {
                v2 = append(v2, 0)
            }
        }
    } else {
        v14 += (v14 >> ((2) & 63))
        for fi24 := 0; fi24 < 11; fi24++ {
            for fi25 := 0; fi25 < 5; fi25++ {
                if len(v2) > 0 {
                    v2[0] = (fi24 ^ fi25)
                }
                var v26 int8 = 48
                var v27 bool = (false && false)
                switch "foo bar" {
                case "":
                    v27 = (((v1 ^ v1) > v1) || (!(1e10 <= (1e10 - ((float32((1.5 + ((0.0 / (0.0 + 3.14159)) / -2.25))) / 0.1) * (0.0 + (1e10 * 3.14159)))))))
                case "a":
                    if (v27 && (v27 || v27)) {
                        v28 := GenS3{F0: -128, F1: v1, F2: (v3 | int64(v14))}
                        v27 = false
                        _ = v28
                    } else {
                        if (832 > (-(v3))) {
                            fmt.Println(v14)
                        } else {
                            var v29 uint16 = v1
                            _ = v29
                        }
                        v30 := GenS2{F0: 216, F1: ""}
                        v2 = append(v2, fi25)
                        _ = v30
                    }
                }
                v3 = (v3 % ((-1) | 1))
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi24
        }
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        switch "" {
        case "hello":
            if (!true) {
                if false {
                    if false {
                        for fi32 := 0; fi32 < 20; fi32++ {
                            v33 := GenS2{F0: 2, F1: "\n\t"}
                            v4[2147483647] = fi31
                            fmt.Println(fi32, fi31, fi32)
                            _ = fi32
                            _ = v33
                        }
                        v3 = (v3 ^ (v3 / ((-1) | 1)))
                    } else {
                        v3 = v3
                        v34 := GenS5{F0: (^(v3)), F1: 0, F2: 997}
                        _ = v34
                    }
                }
                v35 := GenS6{F0: fi31}
                fmt.Println(v1)
                _ = v35
            } else {
                fmt.Println(v3, v14, fi31)
            }
        case "a":
            v36 := GenS7{F0: false, F1: v14}
            _ = v36
        default:
            for fi37 := 0; fi37 < 17; fi37++ {
                var v38 float32 = 0.1
                _ = fi37
                _ = v38
            }
        }
        _ = len(v4)
        _ = fi31
    }
    var v39 int32 = -401
    if false {
        var v40 float32 = 3.14159
        var v41 int64 = (v3 + int64(v1))
        var v42 uint8 = 1
        _ = v40
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
}
