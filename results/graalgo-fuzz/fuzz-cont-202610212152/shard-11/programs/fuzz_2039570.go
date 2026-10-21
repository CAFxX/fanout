package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int; F2 bool }
type GenS4 struct { F0 int16; F1 uint32; F2 string }
func genFn0(p0 int32, p1 int, p2 uint8) int {
    p1 += p1
    var v6 float64 = ((0.0 - 1.5) / 1.0)
    v7 := []int{(-(p1))}
    var v8 float32 = (1e10 * 0.0)
    var v9 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (^(p1))
}
func genFn1(p0 int16, p1 uint, p2 uint8) int {
    v10 := []int{-174, 127, 127, genFn0(-2147483648, -1, (p2 >> (((p1 << (((p1 | (p1 ^ p1))) & 63))) & 63)))}
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int32 = 2147483647
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -2147483648
}

func main() {
    var v1 float32 = (((0.0 * (float32(0.1) * float32(1e10))) + 3.14159) + 1e10)
    var v2 string = "\n\t"
    var v3 float32 = 1e10
    var v4 uint64 = 541
    v5 := []string{(("hello" + v2) + v2), v2, "foo bar"}
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int8 = 30
            v4 = (v4 | 0)
            v3 += v3
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    if (float64((1.5 + (v1 * ((1.5 - ((0.0 / 1.0) + ((((0.0 * (v3 + (v1 * (v3 + (float32((1e10 * 0.1)) + v1))))) + 1e10) - float32(float64(v3))) - v1))) / v3)))) != -2.25) {
        var v16 uint32 = 65535
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v17 uint = 4294967295
        _ = v16
        _ = v17
    }
    var v18 uint8 = 42
    switch v2 {
    case "\n\t":
        switch genFn1(-788, 0, v18) {
        case 340:
            for fi19 := 0; fi19 < 5; fi19++ {
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
                var v20 uint32 = 42
                if (0 >= (^(v20))) {
                    v21 := GenS1{F0: (!true)}
                    var v22 float32 = v1
                    switch (fi19 - -738) {
                    case 919:
                        for fi23 := 0; fi23 < 9; fi23++ {
                            var v24 int16 = 32767
                            _ = fi23
                            _ = v24
                        }
                    case 42:
                        v1 -= v3
                    case 128:
                        v4 = 145
                    }
                    _ = v21
                    _ = v22
                } else {
                    if true {
                        var v25 float32 = -2.25
                        switch (fi19 | fi19) {
                        case -2147483648:
                            for fi26 := 0; fi26 < 10; fi26++ {
                                if false {
                                    v27 := GenS1{F0: true}
                                    _ = v27
                                } else {
                                    v5 = append(v5, (v2 + v2))
                                    if (4294967295 > v20) {
                                        switch (fi26 % ((2147483647) | 1)) {
                                        case -2147483648:
                                            fmt.Println(v18, fi26, v20)
                                        case -264:
                                            v4 = 0
                                        }
                                        v5 = append(v5, v2)
                                        v18 -= v18
                                        v1 -= ((1e10 * v1) - (v1 * (v25 * (v25 / v3))))
                                    }
                                }
                                v3 = v3
                                _ = fi26
                            }
                        case -1:
                            for fi28 := 0; fi28 < 17; fi28++ {
                                switch (fi28 << ((255) & 63)) {
                                case 600:
                                    var v29 uint = uint(v4)
                                    _ = v29
                                case 2147483647:
                                    v30 := GenS2{F0: 0}
                                    _ = v30
                                }
                                if false {
                                    var v31 uint8 = (v18 >> ((484) & 63))
                                    if true {
                                        if false {
                                            var v32 uint32 = uint32(v4)
                                            fmt.Println(v31, v18, v18)
                                            _ = v32
                                        } else {
                                            fmt.Println(v20, v31, v1)
                                        }
                                        v18 = (v31 >> ((0) & 63))
                                        var v33 float64 = (((0.1 - 3.14159) * ((1e10 * -2.25) + -2.25)) - 0.1)
                                        _ = v33
                                    } else {
                                        if (42 == 447) {
                                            fmt.Println(fi28, v31, v31)
                                        } else {
                                            if (-1 <= 42) {
                                                var v34 uint32 = v20
                                                _ = v34
                                            }
                                            fmt.Println(fi28, v31, fi19)
                                        }
                                        if false {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            fmt.Println(fi28)
                                        } else {
                                            v35 := GenS3{F0: (v20 >> ((0) & 63)), F1: (fi19 ^ 762), F2: (((^(v18)) >= v31) || ((v31 >= 127) && (!true)))}
                                            _ = len(v5)
                                            var v36 int16 = 42
                                            _ = v35
                                            _ = v36
                                        }
                                        var v37 uint = 528
                                        _ = v37
                                    }
                                    var v38 int64 = 9223372036854775807
                                    _ = v31
                                    _ = v38
                                }
                                var v39 float64 = 1e10
                                _ = fi28
                                _ = v39
                            }
                        case 9223372036854775807:
                            v2 = v2
                        default:
                            v25 = float32((0.1 * 0.1))
                        }
                        var v40 uint8 = v18
                        _ = v25
                        _ = v40
                    } else {
                        v2 = ((v2 + (("\n\t" + v2) + (v2 + v2))) + "foo bar")
                        switch fi19 {
                        case 2147483647:
                            v41 := GenS1{F0: (((v2 + v2) <= v2) || (!((float32(-2.25) > v3) && ((fi19 >> ((65535) & 63)) <= fi19))))}
                            _ = v41
                        }
                    }
                    v42 := GenS3{F0: 2147483648, F1: fi19, F2: false}
                    v1 = ((0.0 + (((v3 + (v3 / ((0.0 / v3) / (-2.25 - v1)))) + 0.0) + (v3 - v1))) * float32(1e10))
                    _ = v42
                }
                _ = fi19
                _ = v20
            }
        case 9223372036854775807:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    case "a":
        v3 = ((v1 - (float32((0.1 / float64(1e10))) + v1)) - v3)
    case "":
        v4 = 18446744073709551615
    }
    var v43 uint64 = (v4 % ((966) | 1))
    switch 32767 {
    case -2147483648:
        for fi44 := 0; fi44 < 10; fi44++ {
            v5 = append(v5, v2)
            v45 := GenS3{F0: 65535, F1: (fi44 + -1), F2: (!(4294967295 <= (^(v4))))}
            _ = fi44
            _ = v45
        }
    case 365:
        if false {
            for fi46 := 0; fi46 < 14; fi46++ {
                for fi47 := 0; fi47 < 7; fi47++ {
                    v48 := GenS3{F0: 1, F1: 128, F2: (!true)}
                    var v49 int = -2147483648
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                _ = fi46
            }
        }
    }
    v4 = v4
    if (false || true) {
        v50 := GenS2{F0: 32767}
        var v51 int = int(v18)
        if ((v3 < v1) || true) {
            v51 += (v51 - v51)
        }
        _ = v50
        _ = v51
    } else {
        for fi52 := 0; fi52 < 18; fi52++ {
            if false {
                if false {
                    v1 -= 0.1
                    v3 = float32((float64(v3) * 0.1))
                }
            }
            _ = fi52
        }
        v18 = (v18 + v18)
        if true {
            var v53 uint16 = 255
            v3 += v1
            v53 = v53
            _ = v53
        } else {
            v5 = append(v5, (v2 + v2))
            v54 := GenS2{F0: 2}
            _ = v54
        }
    }
    var v55 float64 = (0.1 / 3.14159)
    if len(v5) > 0 {
        v5[0] = "a"
    }
    v5 = append(v5, v2)
    fmt.Println(v2, v18, v55)
    var v56 int8 = 1
    v57 := GenS4{F0: 1000, F1: 2, F2: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v55
    _ = v56
    _ = v57
}
