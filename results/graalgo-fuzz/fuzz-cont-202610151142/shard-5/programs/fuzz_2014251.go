package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint; F2 bool }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 int8; F1 string }
func genFn0(p0 uint, p1 int8, p2 bool) int16 {
    switch ("" + ("hello" + (("hello" + "\n\t") + "foo bar"))) {
    case "hello":
        if (2 < 252) {
            var v4 int64 = -9223372036854775808
            v4 += (v4 | (-(v4)))
            switch 283 {
            case 41:
                var v5 int8 = -1
                _ = v5
            case 42:
                var v6 int32 = 2147483647
                _ = v6
            }
            _ = v4
        } else {
            for fi7 := 0; fi7 < 8; fi7++ {
                fmt.Println(p0, fi7, p2)
                switch int(p0) {
                case -2147483648:
                    p1 += p1
                case 32768:
                    switch (fi7 % (((-(fi7))) | 1)) {
                    case 924:
                        var v8 int16 = -370
                        _ = v8
                    default:
                        var v9 bool = p2
                        _ = v9
                    }
                default:
                    fmt.Println(fi7, p0, p2)
                }
                _ = fi7
            }
            v10 := GenS1{F0: p1, F1: (p1 ^ (p1 / ((p1) | 1)))}
            _ = v10
        }
    case "foo bar":
        if (!((1 > 128) || (993 >= 330))) {
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 string = "a"
                _ = fi11
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 5; fi13++ {
                if (-716 != 0) {
                    for fi14 := 0; fi14 < 19; fi14++ {
                        p2 = ((3.14159 + (0.1 - (((((0.1 + 1.5) / ((1.5 / 1e10) - float64(1e10))) / 1.5) + 0.0) - 0.0))) >= 1e10)
                        fmt.Println(fi13)
                        p1 -= (p1 % (((p1 % (((p1 >> (((^(p0))) & 63))) | 1))) | 1))
                        _ = fi14
                    }
                    p1 = p1
                } else {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        if ((fi13 >> (((p0 << (((^(p0))) & 63))) & 63)) <= fi13) {
                            p0 -= 0
                            p2 = p2
                        } else {
                            p0 += 1
                            p2 = (2 > 65535)
                        }
                        v16 := GenS2{F0: 714}
                        _ = fi15
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 9; fi17++ {
                        p1 -= 100
                        _ = fi17
                    }
                }
                var v18 int64 = -203
                _ = fi13
                _ = v18
            }
        }
    default:
        if p2 {
            switch 503 {
            case 42:
                p2 = (!p2)
            }
            v19 := []int{-760, -533}
            _ = v19
        } else {
            var v20 int64 = 470
            v21 := []int{-917, 1}
            _ = v20
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 uint = (p0 | (^(p0)))
        p2 = p2
        _ = fi22
        _ = v23
    }
    var v24 uint = (p0 ^ (p0 ^ (p0 % ((p0) | 1))))
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return -1
}
func genFn1(p0 uint) int8 {
    p0 -= (p0 ^ (^(p0)))
    var v25 int32 = -813
    if false {
        if false {
            v26 := GenS2{F0: 42}
            for fi27 := 0; fi27 < 3; fi27++ {
                v28 := GenS1{F0: 0, F1: -128}
                var v29 uint = p0
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = v26
        } else {
            var v30 int16 = genFn0((p0 ^ 4294967295), 127, (18446744073709551615 == 1))
            fmt.Println(v30, v25, p0)
            _ = v30
        }
        p0 -= (p0 << ((p0) & 63))
    } else {
        if ((-2.25 / 0.1) >= 3.14159) {
            v25 = (v25 * (^(v25)))
            fmt.Println(p0)
        }
    }
    _ = p0
    _ = v25
    return -128
}

func main() {
    var v1 string = "hello"
    var v2 uint16 = 2
    var v3 int64 = 9223372036854775807
    if (!(true && false)) {
        fmt.Println(v1)
        var v31 uint32 = 515
        _ = v31
    }
    v32 := GenS3{F0: 867, F1: 1, F2: false}
    if false {
        fmt.Println(v2, v3)
        v33 := []int{-140, 2147483647, 127}
        var v34 float32 = 0.1
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        for fi36 := 0; fi36 < 9; fi36++ {
            if true {
                fmt.Println(v2, v3)
                v37 := []int{32768, fi35}
                v2 -= v2
                _ = v37
            } else {
                v38 := GenS3{F0: 2147483647, F1: 678, F2: true}
                fmt.Println(fi35)
                _ = v38
            }
            fmt.Println(v3, fi35)
            _ = fi36
        }
        _ = fi35
    }
    v39 := GenS4{F0: 127, F1: 0}
    fmt.Println(v3, v3, v1)
    v2 += (^(v2))
    switch ("foo bar" + v1) {
    case "foo bar":
        if (0 < 4294967295) {
            if (4294967295 == -59) {
                var v40 int32 = 0
                for fi41 := 0; fi41 < 19; fi41++ {
                    v40 -= (-(v40))
                    v42 := GenS2{F0: v40}
                    v1 = "foo bar"
                    _ = fi41
                    _ = v42
                }
                v43 := GenS3{F0: 121, F1: 250, F2: true}
                _ = v40
                _ = v43
            }
            var v44 int8 = 100
            _ = v44
        }
    case "":
        for fi45 := 0; fi45 < 13; fi45++ {
            if true {
                v2 += v2
                for fi46 := 0; fi46 < 9; fi46++ {
                    if true {
                        if ((212 >= 128) || false) {
                            var v47 uint = 517
                            v48 := []int{-9223372036854775808, (fi46 - (fi45 | (fi45 ^ -2147483648))), fi46, fi46}
                            _ = v47
                            _ = v48
                        } else {
                            v49 := GenS3{F0: 42, F1: 0, F2: true}
                            _ = v49
                        }
                    }
                    v50 := []int{232, (fi45 >> ((40) & 63)), fi46}
                    _ = fi46
                    _ = v50
                }
            }
            _ = fi45
        }
    case "hello":
        fmt.Println(v2)
    default:
        v51 := GenS5{F0: 0, F1: "hello"}
        _ = v51
    }
    switch 128 {
    case 553:
        var v52 int64 = int64(v2)
        _ = v52
    case 2147483647:
        fmt.Println(v1)
    case 42:
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v39
}
