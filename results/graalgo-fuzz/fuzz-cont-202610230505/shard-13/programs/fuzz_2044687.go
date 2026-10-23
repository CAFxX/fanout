package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int16; F1 int32; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint }
func genFn0() int64 {
    v7 := GenS1{F0: 255}
    for fi8 := 0; fi8 < 19; fi8++ {
        switch (("foo bar" + "") + "foo bar") {
        case "\n\t":
            var v9 string = "hello"
            _ = v9
        case "foo bar":
            if ((-2.25 * 0.0) != float64(0.1)) {
                v10 := GenS1{F0: 1}
                _ = v10
            } else {
                switch (fi8 | fi8) {
                case 2147483647:
                    switch 0 {
                    case 32767:
                        var v11 string = "hello"
                        _ = v11
                    case 9223372036854775807:
                        var v12 int64 = 85
                        _ = v12
                    }
                case -709:
                    v13 := GenS2{F0: 928, F1: 0}
                    _ = v13
                }
                v14 := GenS1{F0: 183}
                _ = v14
            }
        case "hello":
            var v15 int32 = -1
            _ = v15
        default:
            switch (("\n\t" + "\n\t") + "hello") {
            case "hello":
                var v16 int64 = 42
                _ = v16
            case "\n\t":
                v17 := []int{(fi8 >> ((689) & 63)), fi8, (fi8 | (fi8 % ((-9223372036854775808) | 1)))}
                _ = v17
            case "":
                var v18 uint32 = 4294967295
                _ = v18
            default:
                switch (fi8 ^ fi8) {
                case 550:
                    if (2147483648 < 42) {
                        fmt.Println(fi8)
                        var v19 bool = (42 > 32767)
                        v19 = (("\n\t" + ("foo bar" + ("\n\t" + "\n\t"))) > "foo bar")
                        _ = v19
                    } else {
                        var v20 uint16 = 78
                        v21 := GenS3{F0: 1000, F1: 1, F2: -1}
                        _ = v20
                        _ = v21
                    }
                case 32768:
                    for fi22 := 0; fi22 < 8; fi22++ {
                        var v23 int64 = -9223372036854775808
                        _ = fi22
                        _ = v23
                    }
                }
            }
        }
        switch "hello" {
        case "hello":
            var v24 uint32 = 438
            _ = v24
        case "a":
            v25 := []int{(fi8 >> ((1) & 63)), (fi8 << ((0) & 63))}
            _ = v25
        }
        _ = fi8
    }
    _ = v7
    return 42
}
func genFn1(p0 uint16, p1 uint16) int {
    var v26 int8 = 1
    v27 := GenS4{F0: 1}
    v26 -= 42
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return 413
}

func main() {
    var v1 uint = 255
    var v2 float32 = (0.0 * 3.14159)
    var v3 int8 = 0
    var v4 string = ("a" + "foo bar")
    var v5 int64 = 42
    var v6 uint32 = 1
    if true {
        if true {
            if true {
                var v28 uint = v1
                v2 = (v2 - 3.14159)
                for fi29 := 0; fi29 < 15; fi29++ {
                    for fi30 := 0; fi30 < 10; fi30++ {
                        switch v4 {
                        case "":
                            if (!(1e10 == 0.1)) {
                                var v31 float64 = (3.14159 / 0.1)
                                v31 += v31
                                v32 := GenS5{F0: uint64(v6)}
                                _ = v31
                                _ = v32
                            }
                        case "foo bar":
                            v33 := GenS2{F0: (^(v5)), F1: 90}
                            _ = v33
                        default:
                            var v34 uint16 = 42
                            _ = v34
                        }
                        var v35 uint8 = 127
                        v36 := GenS2{F0: -1, F1: 234}
                        v37 := []int{-1, fi30, (-(fi29)), fi29}
                        var v38 uint = v28
                        _ = fi30
                        _ = v35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    _ = fi29
                }
                _ = v28
            } else {
                var v39 uint32 = (v6 & v6)
                v39 = 463
                _ = v39
            }
        } else {
            if (!(-2.25 >= 1.5)) {
                for fi40 := 0; fi40 < 17; fi40++ {
                    var v41 int = fi40
                    _ = fi40
                    _ = v41
                }
                v6 -= uint32(v1)
                v42 := GenS3{F0: 32767, F1: int32(v3), F2: 0}
                _ = v42
            } else {
                v3 = 25
                for fi43 := 0; fi43 < 6; fi43++ {
                    v44 := []int{(fi43 * fi43), (^(fi43)), fi43, fi43}
                    if true {
                        v2 -= v2
                    } else {
                        if true {
                            v44 = append(v44, 83)
                            if ((v2 <= -2.25) || true) {
                                v5 -= (v5 & 9223372036854775807)
                                switch v4 {
                                case "hello":
                                    if (4294967295 >= uint64(v1)) {
                                        if len(v44) > 0 {
                                            v44[0] = (fi43 << ((490) & 63))
                                        }
                                    } else {
                                        if len(v44) >= 2 {
                                            v44 = v44[1:len(v44)]
                                        }
                                    }
                                }
                            } else {
                                v3 = (v3 ^ 1)
                            }
                        }
                    }
                    _ = fi43
                    _ = v44
                }
            }
            var v45 uint16 = 0
            fmt.Println(v1)
            _ = v45
        }
    } else {
        if true {
            for fi46 := 0; fi46 < 2; fi46++ {
                v1 = (v1 & v1)
                v4 = v4
                v47 := []int{456, genFn1(255, 0)}
                _ = fi46
                _ = v47
            }
        }
    }
    var v48 uint8 = 1
    fmt.Println(v1, v3, v1)
    switch ("foo bar" + v4) {
    case "hello":
        var v49 int64 = genFn0()
        _ = v49
    case "\n\t":
        for fi50 := 0; fi50 < 11; fi50++ {
            if (3.14159 != 0.1) {
                v51 := []int{fi50, fi50, fi50, fi50}
                v48 += 42
                _ = v51
            }
            _ = fi50
        }
    case "a":
        fmt.Println(v3, v2, v48)
    default:
        v2 = 1e10
    }
    var v52 string = v4
    fmt.Println(v1)
    v1 = v1
    v53 := GenS6{F0: 739, F1: v6, F2: v1}
    fmt.Println(v5)
    fmt.Println(v6, v48)
    fmt.Println(v48, v2)
    if false {
        v4 = "hello"
        v54 := GenS3{F0: 0, F1: -2147483648, F2: 1}
        if true {
            var v55 uint = 255
            var v56 int8 = (v3 - (v3 & (v3 - 42)))
            _ = v55
            _ = v56
        } else {
            switch genFn1(42, 825) {
            case -9223372036854775808:
                var v57 uint32 = (v6 / ((0) | 1))
                _ = v57
            }
        }
        _ = v54
    }
    if true {
        var v58 string = "a"
        _ = v58
    } else {
        v59 := []int{-980, 128}
        v6 -= v6
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v52
    _ = v53
    _ = v6
}
