package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 int32 }
func genFn0(p0 int16, p1 int) int16 {
    var v5 uint8 = 0
    fmt.Println(p1, p1)
    var v6 float64 = float64(1e10)
    p0 -= (p0 | 1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p0 / ((p0) | 1))
}
func genFn1(p0 int8, p1 string) uint {
    p0 = (p0 & 121)
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: 2, F1: 42}
        fmt.Println(p0)
        p0 = p0
        _ = fi7
        _ = v8
    }
    var v9 float64 = ((-2.25 - 1e10) / 1.5)
    var v10 float32 = 1.5
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 65535
}

func main() {
    var v1 bool = (!true)
    v2 := []int{}
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    v1 = v1
    switch ("foo bar" + "a") {
    case "foo bar":
        if v1 {
            for fi11 := 0; fi11 < 9; fi11++ {
                v1 = ((float32(3.14159) + 1e10) != (0.0 + -2.25))
                for fi12 := 0; fi12 < 15; fi12++ {
                    if len(v2) > 0 {
                        v2[0] = fi11
                    }
                    v4[29] = fi11
                    v1 = v1
                    _ = fi12
                }
                _ = fi11
            }
        } else {
            if v1 {
                v1 = (v1 || (!(42 >= 648)))
                for fi13 := 0; fi13 < 2; fi13++ {
                    if true {
                        v1 = false
                        v1 = false
                    }
                    var v14 float64 = 1e10
                    _ = fi13
                    _ = v14
                }
                switch "a" {
                case "":
                    v1 = (false || false)
                case "foo bar":
                    fmt.Println(v1)
                }
            } else {
                _ = len(v2)
                if v1 {
                    if len(v2) > 0 {
                        v2[0] = -419
                    }
                    var v15 float64 = 3.14159
                    _ = v15
                }
                switch "hello" {
                case "\n\t":
                    fmt.Println(v1)
                case "foo bar":
                    if ("" == ("foo bar" + "foo bar")) {
                        v16 := GenS2{F0: 65535}
                        for fi17 := 0; fi17 < 17; fi17++ {
                            if false {
                                if (4294967295 >= 686) {
                                    v18 := GenS3{F0: 32767, F1: 255, F2: 2147483647}
                                    if len(v2) > 0 {
                                        v2[0] = (fi17 / ((964) | 1))
                                    }
                                    v1 = v1
                                    _ = v18
                                } else {
                                    var v19 float32 = (3.14159 - ((0.1 - 3.14159) - (1.5 - (1e10 / float32(1.5)))))
                                    for fi20 := 0; fi20 < 11; fi20++ {
                                        v1 = (439 <= 1)
                                        var v21 uint16 = 2
                                        v21 -= v21
                                        _ = fi20
                                        _ = v21
                                    }
                                    switch (^(fi17)) {
                                    case 32768:
                                        fmt.Println(fi17, v19, fi17)
                                    case 294:
                                        if v1 {
                                            v22 := GenS4{F0: 2, F1: fi17}
                                            if v1 {
                                                for fi23 := 0; fi23 < 17; fi23++ {
                                                    v1 = (42 < 42)
                                                    if (!v1) {
                                                        v4[297] = fi23
                                                        v24 := GenS1{F0: 2, F1: 0}
                                                        var v25 float32 = float32(1e10)
                                                        _ = v24
                                                        _ = v25
                                                    } else {
                                                        var v26 uint32 = 0
                                                        if v1 {
                                                            if (462 > 458) {
                                                                v27 := GenS4{F0: 4294967295, F1: fi17}
                                                                _ = v27
                                                            }
                                                            var v28 uint32 = v26
                                                            delete(v3, (fi17 & (^(fi17))))
                                                            _ = v28
                                                        }
                                                        _ = v26
                                                    }
                                                    v29 := GenS2{F0: 636}
                                                    _ = fi23
                                                    _ = v29
                                                }
                                            }
                                            for fi30 := 0; fi30 < 1; fi30++ {
                                                var v31 string = ""
                                                v19 = v19
                                                v19 = float32(-2.25)
                                                _ = fi30
                                                _ = v31
                                            }
                                            _ = v22
                                        }
                                    case -9223372036854775808:
                                        v19 = v19
                                    }
                                    _ = v19
                                }
                            } else {
                                var v32 float64 = 1.5
                                switch (fi17 / (((fi17 % (((fi17 >> ((458) & 63))) | 1))) | 1)) {
                                case -1:
                                    var v33 uint = 477
                                    _ = v33
                                case -2147483648:
                                    v34 := v3[237]
                                    _ = v34
                                case 799:
                                    delete(v4, 0)
                                }
                                fmt.Println(fi17, v1, fi17)
                                _ = v32
                            }
                            _ = fi17
                        }
                        switch 0 {
                        case -9223372036854775808:
                            v1 = (("" < ("foo bar" + "foo bar")) && v1)
                        default:
                            v1 = ("\n\t" <= "hello")
                        }
                        _ = v16
                    }
                default:
                    v35 := GenS5{F0: -1, F1: 2, F2: 0}
                    _ = v35
                }
            }
            for fi36 := 0; fi36 < 2; fi36++ {
                _ = len(v3)
                _ = fi36
            }
        }
    default:
        var v37 float32 = 0.0
        _ = v37
    }
    for fi38 := 0; fi38 < 16; fi38++ {
        v1 = v1
        v39 := GenS5{F0: -1, F1: 660, F2: 1}
        var v40 int16 = 1
        _ = fi38
        _ = v39
        _ = v40
    }
    v41 := GenS6{F0: 42, F1: 255}
    if (("a" + "foo bar") == ("" + ("a" + "foo bar"))) {
        var v42 uint32 = 1
        _ = v42
    }
    var v43 float64 = 1e10
    var v44 int8 = -1
    for fi45 := 0; fi45 < 13; fi45++ {
        var v46 int = (fi45 & -2147483648)
        v44 = 1
        _ = fi45
        _ = v46
    }
    for fi47 := 0; fi47 < 5; fi47++ {
        v48 := GenS7{F0: 42}
        switch "foo bar" {
        case "":
            switch ("\n\t" + "a") {
            case "hello":
                for fi49 := 0; fi49 < 16; fi49++ {
                    var v50 uint64 = 18446744073709551615
                    var v51 float32 = (((1.5 + 0.0) + (1.5 / (-2.25 / 1.0))) + (0.0 / (((1e10 + -2.25) * (-2.25 - 1.5)) + float32(-2.25))))
                    _ = fi49
                    _ = v50
                    _ = v51
                }
            default:
                var v52 int16 = -549
                _ = v52
            }
        case "a":
            _ = len(v3)
        }
        _ = fi47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v44
}
