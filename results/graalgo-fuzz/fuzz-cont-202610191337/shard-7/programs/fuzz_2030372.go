package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint32; F2 bool }
type GenS3 struct { F0 string; F1 int16; F2 int }
type GenS4 struct { F0 int64 }
func genFn0() uint32 {
    v6 := GenS1{F0: 127, F1: 2147483647, F2: -2147483648}
    var v7 int64 = 42
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 int32, p1 uint64, p2 string) uint64 {
    if (!true) {
        var v8 uint8 = 1
        _ = v8
    } else {
        if true {
            var v9 string = p2
            var v10 bool = true
            _ = v10
            _ = v9
        } else {
            v11 := GenS1{F0: 2, F1: (p0 | (^(p0))), F2: p0}
            _ = v11
        }
    }
    v12 := GenS1{F0: 2, F1: p0, F2: (-(p0))}
    switch (((p2 + p2) + (((p2 + ("a" + p2)) + p2) + "")) + (("\n\t" + "a") + (((p2 + (("" + p2) + p2)) + "\n\t") + (p2 + "\n\t")))) {
    case "\n\t":
        v13 := []int{-59, 56, -171, -9223372036854775808}
        _ = v13
    default:
        if (!(true || false)) {
            p1 -= (p1 + (p1 * 42))
            fmt.Println(p0, p1, p0)
            var v14 uint16 = 42
            _ = v14
        } else {
            switch p2 {
            case "hello":
                switch p2 {
                case "\n\t":
                    fmt.Println(p0)
                default:
                    for fi15 := 0; fi15 < 19; fi15++ {
                        var v16 float64 = (1e10 * -2.25)
                        switch p2 {
                        case "hello":
                            v17 := GenS1{F0: 53, F1: 0, F2: p0}
                            _ = v17
                        case "foo bar":
                            switch (fi15 % (((fi15 - (fi15 + (^(fi15))))) | 1)) {
                            case 4294967295:
                                var v18 int16 = 32767
                                _ = v18
                            default:
                                v19 := GenS2{F0: 42, F1: 4294967295, F2: false}
                                _ = v19
                            }
                        case "a":
                            if false {
                                v20 := []int{4294967295}
                                _ = v20
                            }
                        default:
                            var v21 float32 = 3.14159
                            _ = v21
                        }
                        _ = fi15
                        _ = v16
                    }
                }
            case "foo bar":
                if true {
                    p0 += (p0 >> ((42) & 63))
                    v22 := []int{}
                    v22 = append(v22, 32768)
                    _ = v22
                } else {
                    p2 = (p2 + ("foo bar" + (p2 + p2)))
                    fmt.Println(p1, p2)
                    switch (p2 + (p2 + (p2 + (p2 + "hello")))) {
                    case "":
                        v23 := []int{1, 1}
                        _ = v23
                    }
                }
            case "":
                v24 := GenS3{F0: p2, F1: 1, F2: 392}
                _ = v24
            default:
                var v25 uint16 = 326
                _ = v25
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p1
}

func main() {
    var v1 uint32 = 807
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    var v4 uint64 = 0
    var v5 uint64 = (v4 / (((v4 + (v4 | 4294967295))) | 1))
    v5 += 0
    var v26 float64 = -2.25
    switch "hello" {
    case "":
        for fi27 := 0; fi27 < 13; fi27++ {
            v28 := []int{}
            v26 += 1.5
            _ = fi27
            _ = v28
        }
    default:
        if (true && (-681 < 2147483647)) {
            switch 128 {
            case 32768:
                var v29 float64 = 0.0
                _ = v29
            case 231:
                if true {
                    fmt.Println(v26, v1)
                }
            case 884:
                for fi30 := 0; fi30 < 7; fi30++ {
                    if (v5 <= v5) {
                        for fi31 := 0; fi31 < 18; fi31++ {
                            fmt.Println(v1, v26, v4)
                            _ = fi31
                        }
                        if (42 < -1) {
                            if (true && false) {
                                v1 -= 2147483648
                                v5 = v5
                                v32 := GenS2{F0: 65535, F1: (v1 ^ 65535), F2: (333 >= 0)}
                                _ = v32
                            } else {
                                if ((true || (0.0 <= (v26 + 1e10))) || false) {
                                    var v33 string = "\n\t"
                                    var v34 float64 = v26
                                    var v35 float64 = v34
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                } else {
                                    _ = len(v2)
                                    for fi36 := 0; fi36 < 5; fi36++ {
                                        fmt.Println(fi36)
                                        var v37 bool = (-9223372036854775808 < 1)
                                        v5 -= (v5 << ((315) & 63))
                                        _ = fi36
                                        _ = v37
                                    }
                                }
                                var v38 uint8 = 138
                                var v39 uint16 = 42
                                _ = v38
                                _ = v39
                            }
                        } else {
                            v40 := GenS4{F0: -9223372036854775808}
                            v41 := GenS1{F0: 255, F1: 1, F2: 42}
                            fmt.Println(v1, fi30, v26)
                            _ = v40
                            _ = v41
                        }
                    } else {
                        if false {
                            var v42 string = (("a" + "foo bar") + ("\n\t" + ("foo bar" + (("\n\t" + "a") + "hello"))))
                            _ = v42
                        } else {
                            fmt.Println(fi30)
                            if false {
                                var v43 bool = (-1 == -2147483648)
                                v4 = (^(v4))
                                v44 := []int{579, (fi30 % ((int(v1)) | 1)), (fi30 & fi30), fi30}
                                _ = v43
                                _ = v44
                            }
                        }
                    }
                    _ = fi30
                }
            default:
                switch -480 {
                case 0:
                    var v45 int8 = 100
                    _ = v45
                }
            }
            for fi46 := 0; fi46 < 10; fi46++ {
                if (42 == 100) {
                    v47 := GenS4{F0: -674}
                    _ = v47
                }
                v1 -= uint32(v4)
                _ = fi46
            }
            v26 += v26
        }
    }
    var v48 uint = 0
    v5 = 4294967295
    for fi49 := 0; fi49 < 11; fi49++ {
        v50 := []int{fi49}
        fmt.Println(v26, v1, v1)
        _ = fi49
        _ = v50
    }
    var v51 float32 = -2.25
    var v52 float64 = (v26 / 1.0)
    v1 += v1
    v53 := GenS1{F0: 127, F1: 42, F2: -1}
    for fi54 := 0; fi54 < 10; fi54++ {
        switch "" {
        case "a":
            if false {
                var v55 int8 = -1
                _ = v55
            } else {
                if true {
                    for fi56 := 0; fi56 < 8; fi56++ {
                        fmt.Println(fi56, v4)
                        var v57 float32 = 3.14159
                        _ = fi56
                        _ = v57
                    }
                    var v58 float64 = ((1e10 - 1e10) + (3.14159 - v52))
                    _ = v58
                }
                fmt.Println(v1)
            }
        case "":
            _ = len(v3)
        }
        _ = fi54
    }
    v51 = v51
    v52 += v52
    v1 += 4294967295
    fmt.Println(v51, v26, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v51
    _ = v52
    _ = v53
}
