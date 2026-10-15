package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 1.5
    var v3 int32 = -398
    var v4 int64 = -9223372036854775808
    var v5 float64 = v2
    v6 := make(map[uint32]int)
    for fi7 := 0; fi7 < 1; fi7++ {
        _ = len(v1)
        if true {
            fmt.Println(v2, v5)
            v8 := GenS1{F0: 100, F1: 65535, F2: fi7}
            _ = v8
        } else {
            v9 := GenS2{F0: 42, F1: 1, F2: 255}
            _ = v9
        }
        fmt.Println(v3, v5)
        _ = fi7
    }
    var v10 uint32 = 699
    switch "" {
    case "a":
        delete(v6, (v10 ^ (v10 & v10)))
    case "hello":
        var v11 uint16 = 115
        _ = v11
    case "\n\t":
        v10 = 550
    }
    v5 = (v2 + ((v2 / v2) / v5))
    var v12 string = "\n\t"
    v10 = v10
    v5 = v5
    var v13 uint = 42
    v10 = (v10 & 131)
    v6[(v10 ^ (v10 - 2147483648))] = 127
    v3 = (v3 & v3)
    v5 = v2
    if (!false) {
        switch ("foo bar" + v12) {
        case "\n\t":
            var v14 float32 = 0.1
            _ = v14
        case "a":
            v15 := []int{}
            _ = v15
        case "":
            v16 := GenS2{F0: 1, F1: 2, F2: 65535}
            _ = v16
        }
        v17 := []int{}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := []int{fi18, (-(fi18))}
            var v20 float32 = (0.1 * 0.1)
            if false {
                fmt.Println(v3)
                var v21 int = (fi18 ^ 1)
                _ = v21
            } else {
                var v22 string = ("foo bar" + "\n\t")
                v22 = ""
                for fi23 := 0; fi23 < 2; fi23++ {
                    _ = len(v6)
                    _ = fi23
                }
                _ = v22
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        v13 = 131
        v10 = v10
    }
    v24 := GenS2{F0: 55, F1: uint16(v10), F2: v13}
    if (!false) {
        var v25 uint16 = 2
        if (!false) {
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := []int{(fi26 | fi26), -2147483648}
                var v28 string = (v12 + "hello")
                var v29 int16 = 1000
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            if false {
                if (!false) {
                    var v30 uint8 = 255
                    var v31 int64 = v4
                    _ = v30
                    _ = v31
                }
                if false {
                    for fi32 := 0; fi32 < 17; fi32++ {
                        var v33 uint8 = 16
                        v34 := GenS3{F0: (v13 - 1), F1: (^(v13))}
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    v35 := GenS1{F0: 1, F1: 970, F2: 9223372036854775807}
                    _ = v35
                } else {
                    if ((v3 >> ((v13) & 63)) == v3) {
                        fmt.Println(v12, v2)
                        v3 = v3
                        for fi36 := 0; fi36 < 1; fi36++ {
                            var v37 string = v12
                            switch ("" + v37) {
                            case "":
                                if (!false) {
                                    fmt.Println(v12)
                                    v3 = v3
                                }
                            default:
                                for fi38 := 0; fi38 < 13; fi38++ {
                                    v6[v10] = int(v4)
                                    _ = fi38
                                }
                            }
                            _ = fi36
                            _ = v37
                        }
                    } else {
                        v3 -= v3
                        var v39 string = ("\n\t" + (v12 + ((v12 + v12) + (v12 + ""))))
                        _ = v39
                    }
                }
                v10 -= uint32(v25)
            } else {
                var v40 uint64 = 1
                switch v12 {
                case "foo bar":
                    var v41 float64 = v5
                    _ = v41
                default:
                    if true {
                        v42 := GenS2{F0: v25, F1: (v25 >> ((uint(v10)) & 63)), F2: (v13 & v13)}
                        _ = v42
                    } else {
                        if false {
                            v3 = (v3 - (-(v3)))
                        }
                        v43 := GenS4{F0: v13}
                        _ = v43
                    }
                }
                v44 := []int{}
                _ = v40
                _ = v44
            }
            var v45 int64 = -9223372036854775808
            _ = v45
        }
        var v46 uint = v13
        _ = v25
        _ = v46
    } else {
        v47 := GenS4{F0: v13}
        if ((v13 - 0) == (v13 >> ((4294967295) & 63))) {
            if false {
                switch 1 {
                case 519:
                    v5 += float64((((1e10 * ((0.0 / 1e10) * -2.25)) * -2.25) + float32(0.0)))
                case 4294967295:
                    var v48 uint64 = 18446744073709551615
                    _ = v48
                default:
                    v5 = ((((v5 * ((v2 - v5) - v5)) / v2) + (((v2 - v5) / v2) * ((v5 * v5) * (v2 - v2)))) * (v5 * 1.5))
                }
                v49 := GenS5{F0: 255}
                _ = v49
            }
        }
        if (true && (!(330 >= 4294967295))) {
            switch 737 {
            case 9223372036854775807:
                v1[141] = (0 < 2)
            default:
                fmt.Println(v10, v3)
            }
            switch (v12 + "a") {
            case "a":
                for fi50 := 0; fi50 < 2; fi50++ {
                    switch "\n\t" {
                    case "":
                        fmt.Println(v3)
                    case "a":
                        for fi51 := 0; fi51 < 3; fi51++ {
                            v52 := GenS3{F0: v13, F1: (v13 ^ v13)}
                            _ = fi51
                            _ = v52
                        }
                    default:
                        var v53 int16 = -1
                        _ = v53
                    }
                    _ = fi50
                }
            case "":
                v12 = (v12 + "")
            case "hello":
                _ = len(v6)
            }
        } else {
            v3 -= (-(v3))
        }
        _ = v47
    }
    v54 := GenS6{F0: 0}
    if (-1 < -1) {
        if true {
            v12 = (v12 + ("" + (v12 + ((v12 + ((v12 + v12) + "foo bar")) + "a"))))
            v12 = v12
            var v55 float32 = ((((0.0 - (float32(0.0) - (0.1 * 1e10))) - ((3.14159 - float32((float64((0.0 + (0.0 - (float32((((float64(((-2.25 / 1.5) - 1.5)) / -2.25) + v5) / ((v5 / v2) + v2))) / 0.1)))) - 1e10))) + 0.0)) - 0.0) / float32((1e10 / 1.0)))
            _ = v55
        }
    } else {
        var v56 uint8 = 255
        v10 = v10
        if false {
            var v57 float64 = v2
            _ = v57
        } else {
            if (1 <= 0) {
                if (v13 != 42) {
                    if false {
                        v2 = 3.14159
                        for fi58 := 0; fi58 < 5; fi58++ {
                            v3 = int32(fi58)
                            var v59 uint8 = v56
                            _ = fi58
                            _ = v59
                        }
                        v10 = 2147483648
                    }
                    v60 := []int{}
                    _ = v60
                }
                var v61 int64 = 1
                _ = v61
            }
        }
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
