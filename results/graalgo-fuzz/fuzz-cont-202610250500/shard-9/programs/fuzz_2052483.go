package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int8, p1 uint8, p2 uint8) int16 {
    if (128 < -9223372036854775808) {
        v5 := GenS1{F0: 2147483648, F1: (p0 >> ((2) & 63))}
        p2 += (p1 << ((4294967295) & 63))
        _ = v5
    }
    fmt.Println(p1)
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 bool) uint16 {
    var v6 uint8 = 2
    switch "foo bar" {
    case "a":
        if p0 {
            var v7 int16 = -32768
            if false {
                if (!p0) {
                    v8 := GenS1{F0: 1, F1: 41}
                    if (p0 && p0) {
                        v9 := GenS2{F0: 211}
                        _ = v9
                    } else {
                        v10 := []int{32768, -9223372036854775808}
                        var v11 int16 = genFn0(0, v6, (v6 | 42))
                        _ = v10
                        _ = v11
                    }
                    if false {
                        for fi12 := 0; fi12 < 13; fi12++ {
                            for fi13 := 0; fi13 < 20; fi13++ {
                                var v14 uint32 = 640
                                switch "a" {
                                case "\n\t":
                                    v15 := []int{(fi12 << ((505) & 63))}
                                    _ = v15
                                case "hello":
                                    v14 = (v14 + 772)
                                case "foo bar":
                                    var v16 int32 = -266
                                    _ = v16
                                default:
                                    v17 := []int{(fi13 - (fi13 - (fi13 >> ((0) & 63)))), fi13, (-(fi12))}
                                    _ = v17
                                }
                                _ = fi13
                                _ = v14
                            }
                            _ = fi12
                        }
                    }
                    _ = v8
                }
                var v18 string = "foo bar"
                _ = v18
            }
            if true {
                v19 := []int{-889, 32768, -468, 127}
                for fi20 := 0; fi20 < 8; fi20++ {
                    if ((float64(-2.25) - 0.1) == float64((float32(0.0) * 1.5))) {
                        fmt.Println(fi20)
                        if ((!(1 == 27)) && p0) {
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            fmt.Println(v6, p0)
                            v21 := GenS1{F0: 4294967295, F1: 127}
                            _ = v21
                        }
                    }
                    v19 = append(v19, (fi20 << ((0) & 63)))
                    _ = fi20
                }
                _ = v19
            }
            _ = v7
        } else {
            v22 := GenS1{F0: 1, F1: -128}
            _ = v22
        }
    case "hello":
        p0 = (1 == -404)
    case "":
        if (!((((3.14159 * (0.1 - 1.5)) - 1e10) + -2.25) >= 0.1)) {
            if p0 {
                var v23 uint64 = 2
                v6 = 202
                v23 = v23
                _ = v23
            }
            if (803 <= 1) {
                for fi24 := 0; fi24 < 15; fi24++ {
                    var v25 uint8 = v6
                    _ = fi24
                    _ = v25
                }
            }
            v26 := []int{-185, 42}
            _ = v26
        }
    }
    var v27 int = 0
    var v28 string = "\n\t"
    _ = p0
    _ = v27
    _ = v28
    _ = v6
    return 42
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int = 32767
    var v3 uint32 = 0
    var v4 float64 = 1.5
    v29 := GenS1{F0: (^(v3)), F1: 1}
    if (!(false || (!true))) {
        var v30 int32 = -101
        switch (v2 + v2) {
        case -405:
            for fi31 := 0; fi31 < 12; fi31++ {
                v32 := []int{fi31, (-(fi31)), fi31}
                v1 -= (v1 << (((v1 << ((487) & 63))) & 63))
                v30 = -766
                _ = fi31
                _ = v32
            }
        case -1:
            var v33 uint = (v1 ^ v1)
            _ = v33
        case 801:
            v34 := GenS3{F0: 4294967295, F1: 42}
            _ = v34
        default:
            v35 := []int{(-(v2)), 127, v2, v2}
            _ = v35
        }
        v1 = v1
        _ = v30
    }
    if false {
        for fi36 := 0; fi36 < 12; fi36++ {
            v3 = 1
            _ = fi36
        }
        switch "a" {
        case "hello":
            var v37 bool = true
            _ = v37
        case "foo bar":
            var v38 uint32 = 664
            _ = v38
        }
        v39 := GenS4{F0: 331, F1: ("a" + ("foo bar" + "\n\t"))}
        _ = v39
    } else {
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            for fi40 := 0; fi40 < 1; fi40++ {
                v41 := []int{}
                _ = fi40
                _ = v41
            }
        default:
            v3 = v3
        }
        if (v4 >= ((v4 - v4) * v4)) {
            var v42 string = ""
            _ = v42
        } else {
            if false {
                if (!((0.1 - v4) >= v4)) {
                    var v43 float64 = v4
                    _ = v43
                } else {
                    v1 = v1
                }
                var v44 int8 = -88
                if false {
                    v45 := []int{2147483647}
                    if true {
                        v46 := GenS5{F0: v3, F1: 128}
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 17; fi47++ {
                        switch (("hello" + "foo bar") + (((("\n\t" + ("\n\t" + ("hello" + ""))) + "foo bar") + "") + ("hello" + "foo bar"))) {
                        case "hello":
                            v48 := GenS5{F0: (^(v3)), F1: 0}
                            _ = v48
                        case "foo bar":
                            v44 += 42
                        }
                        var v49 bool = (3.14159 == float32(v4))
                        _ = fi47
                        _ = v49
                    }
                    _ = v45
                }
                _ = v44
            } else {
                for fi50 := 0; fi50 < 3; fi50++ {
                    var v51 bool = (false && true)
                    fmt.Println(v4, v1)
                    _ = fi50
                    _ = v51
                }
                v1 = 0
            }
        }
    }
    var v52 uint64 = 4294967295
    if (!(((v1 ^ v1) >= (v1 >> (((v1 & v1)) & 63))) && false)) {
        v53 := GenS5{F0: (v3 >> ((v1) & 63)), F1: 127}
        for fi54 := 0; fi54 < 1; fi54++ {
            fmt.Println(v4, fi54)
            v3 = 42
            _ = fi54
        }
        _ = v53
    } else {
        v55 := []int{-1, (v2 % ((v2) | 1)), (-(v2)), v2}
        v1 = 255
        v3 -= 42
        _ = v55
    }
    var v56 int64 = -951
    v3 = (v3 % (((v3 | v3)) | 1))
    v56 = (-(v56))
    var v57 float32 = 3.14159
    v58 := GenS3{F0: 42, F1: 127}
    var v59 uint16 = 65535
    var v60 int = -569
    v56 = v56
    fmt.Println(v2, v60)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v52
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v60
}
