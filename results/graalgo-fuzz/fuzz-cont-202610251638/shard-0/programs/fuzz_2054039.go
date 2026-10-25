package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int8; F2 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16; F1 int64; F2 uint64 }
func genFn0() int8 {
    var v7 uint64 = 284
    v8 := GenS1{F0: 77}
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int64, p1 int) bool {
    v9 := []int{p1, (p1 / ((p1) | 1))}
    v10 := GenS2{F0: -128, F1: genFn0(), F2: -35}
    switch p1 {
    case -1:
        for fi11 := 0; fi11 < 5; fi11++ {
            fmt.Println(fi11, p1)
            var v12 int8 = 1
            _ = fi11
            _ = v12
        }
    case 32768:
        for fi13 := 0; fi13 < 10; fi13++ {
            switch "hello" {
            case "a":
                for fi14 := 0; fi14 < 14; fi14++ {
                    switch fi13 {
                    case -1:
                        switch (fi13 & (fi13 * (fi14 - 128))) {
                        case 42:
                            v15 := GenS1{F0: 2}
                            _ = v15
                        case 1:
                            v16 := GenS3{F0: 128}
                            _ = v16
                        default:
                            var v17 uint64 = 4294967295
                            _ = v17
                        }
                    case 127:
                        fmt.Println(fi13, p0, p0)
                    case -2147483648:
                        var v18 uint16 = 729
                        _ = v18
                    }
                    if true {
                        v19 := GenS1{F0: 1}
                        p1 = fi13
                        if len(v9) > 0 {
                            v9[0] = (-(fi13))
                        }
                        _ = v19
                    }
                    _ = fi14
                }
            case "":
                if len(v9) > 0 {
                    v9[0] = p1
                }
            case "hello":
                fmt.Println(fi13, p1)
            }
            p0 += (p0 | p0)
            _ = fi13
        }
    case -759:
        v9 = append(v9, (p1 << ((379) & 63)))
    default:
        var v20 int8 = genFn0()
        _ = v20
    }
    if true {
        var v21 uint32 = 0
        _ = v21
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return true
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float64 = 0.0
    var v3 uint64 = 42
    v4 := make(map[int]string)
    var v5 uint64 = 809
    var v6 uint64 = (v3 & (v5 * (^(v3))))
    if (!false) {
        var v22 float32 = (float32(float64((((float32(v2) * 0.1) + 0.1) + (0.0 / 1.0)))) - (0.1 * (float32((v2 + (0.1 / v2))) * (0.0 * float32(float64(3.14159))))))
        for fi23 := 0; fi23 < 19; fi23++ {
            for fi24 := 0; fi24 < 4; fi24++ {
                if (uint8(v5) >= 80) {
                    var v25 int16 = 1000
                    if false {
                        _ = len(v4)
                        var v26 int = (^(fi23))
                        var v27 string = ((("" + "hello") + ((("" + "foo bar") + (("\n\t" + "foo bar") + "hello")) + ("" + ("a" + "")))) + ("\n\t" + ("" + "")))
                        _ = v26
                        _ = v27
                    } else {
                        v5 = v5
                        v5 = 492
                        switch 32768 {
                        case 374:
                            fmt.Println(fi23, v22)
                        default:
                            v28 := GenS4{F0: 42, F1: 70}
                            _ = v28
                        }
                        v29 := []int{}
                        _ = v29
                    }
                    var v30 float32 = v22
                    _ = v25
                    _ = v30
                }
                _ = fi24
            }
            _ = fi23
        }
        _ = v22
    }
    for fi31 := 0; fi31 < 19; fi31++ {
        fmt.Println(fi31)
        _ = fi31
    }
    v32 := GenS1{F0: 1}
    v3 = (^(v5))
    for fi33 := 0; fi33 < 13; fi33++ {
        for fi34 := 0; fi34 < 17; fi34++ {
            v3 = v3
            _ = fi34
        }
        for fi35 := 0; fi35 < 16; fi35++ {
            v36 := []int{}
            v37 := v4[(fi35 << ((255) & 63))]
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi33
    }
    v38 := GenS1{F0: 42}
    var v39 bool = true
    for fi40 := 0; fi40 < 6; fi40++ {
        var v41 float32 = 0.0
        v6 = (v6 >> ((255) & 63))
        fmt.Println(v39, v6)
        _ = fi40
        _ = v41
    }
    v42 := GenS5{F0: v1, F1: 0}
    v43 := GenS4{F0: 255, F1: 424}
    v44 := GenS6{F0: 100, F1: -2147483648}
    if (-1 == 42) {
        switch 0 {
        case 127:
            v2 -= (v2 / v2)
        default:
            var v45 int32 = 42
            _ = v45
        }
        switch 9223372036854775807 {
        case -2147483648:
            switch "foo bar" {
            case "foo bar":
                if false {
                    var v46 string = ("a" + ("" + "hello"))
                    v46 = "a"
                    if v39 {
                        v46 = v46
                    }
                    _ = v46
                } else {
                    if ((!(549 != 65535)) || (42 < 4294967295)) {
                        switch "foo bar" {
                        case "\n\t":
                            for fi47 := 0; fi47 < 12; fi47++ {
                                v3 = v5
                                _ = fi47
                            }
                        case "hello":
                            v48 := []int{-2147483648}
                            _ = v48
                        case "":
                            var v49 uint8 = 98
                            _ = v49
                        }
                    }
                }
            case "\n\t":
                for fi50 := 0; fi50 < 16; fi50++ {
                    if true {
                        if v39 {
                            if (!(!(42 < 18446744073709551615))) {
                                v3 += v5
                                for fi51 := 0; fi51 < 7; fi51++ {
                                    var v52 uint64 = (v6 & v3)
                                    _ = fi51
                                    _ = v52
                                }
                                v6 -= v6
                            } else {
                                v53 := GenS7{F0: ("a" + "a")}
                                fmt.Println(v6)
                                var v54 float32 = (0.1 * 1.5)
                                _ = v53
                                _ = v54
                            }
                            switch (("hello" + "") + "a") {
                            case "a":
                                if (v1 != v1) {
                                    switch (("" + "hello") + "a") {
                                    case "a":
                                        fmt.Println(v6, v39, v3)
                                    case "foo bar":
                                        var v55 float32 = (1.5 * -2.25)
                                        _ = v55
                                    default:
                                        for fi56 := 0; fi56 < 6; fi56++ {
                                            if true {
                                                var v57 uint = 255
                                                fmt.Println(v3, v39, v1)
                                                v57 += v57
                                                delete(v4, (^(fi50)))
                                                v58 := []int{int(v1), (fi50 | (fi56 & (-(fi56))))}
                                                _ = v57
                                                _ = v58
                                            }
                                            var v59 float32 = (0.1 - 0.0)
                                            _ = fi56
                                            _ = v59
                                        }
                                    }
                                } else {
                                    v5 = v5
                                    switch fi50 {
                                    case 127:
                                        for fi60 := 0; fi60 < 18; fi60++ {
                                            v1 -= v1
                                            v6 = 0
                                            v5 -= v5
                                            fmt.Println(v1, v2, v6)
                                            var v61 uint64 = (v3 | (^(v3)))
                                            _ = fi60
                                            _ = v61
                                        }
                                    case -1:
                                        fmt.Println(v2, fi50)
                                    }
                                }
                            case "hello":
                                if false {
                                    var v62 uint64 = v5
                                    if (!true) {
                                        v5 = (v5 * v5)
                                    }
                                    _ = v62
                                } else {
                                    var v63 float32 = 3.14159
                                    _ = v63
                                }
                            case "":
                                v5 -= 63
                            }
                            switch (fi50 + 183) {
                            case 2147483647:
                                for fi64 := 0; fi64 < 11; fi64++ {
                                    v2 -= (v2 + 0.0)
                                    _ = fi64
                                }
                            case 127:
                                v6 = (v5 >> ((1) & 63))
                            }
                        }
                        for fi65 := 0; fi65 < 5; fi65++ {
                            fmt.Println(v5)
                            switch fi50 {
                            case 127:
                                v39 = (!v39)
                            case 2147483647:
                                v2 = v2
                            }
                            _ = fi65
                        }
                        v66 := GenS8{F0: -115, F1: -9223372036854775808, F2: (v3 ^ 64)}
                        _ = v66
                    }
                    var v67 uint8 = 255
                    _ = fi50
                    _ = v67
                }
            }
        case 0:
            v39 = (((!(!(v1 > 9223372036854775807))) || (42 < 2)) || (true && (v6 < 42)))
        default:
            for fi68 := 0; fi68 < 1; fi68++ {
                for fi69 := 0; fi69 < 16; fi69++ {
                    v39 = true
                    fmt.Println(v2)
                    v4[fi68] = ("hello" + "hello")
                    _ = fi69
                }
                v5 += (^(v3))
                switch (fi68 % ((fi68) | 1)) {
                case 32768:
                    fmt.Println(v5, v2)
                }
                _ = fi68
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
