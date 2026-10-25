package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 string; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int8; F1 bool }
func genFn0(p0 uint, p1 int) int32 {
    var v7 uint32 = 0
    var v8 uint16 = 255
    p0 += (p0 * p0)
    var v9 float64 = 1e10
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 500
}
func genFn1(p0 string, p1 int8) uint {
    p0 = "foo bar"
    switch -181 {
    case -1:
        v10 := GenS1{F0: 4294967295, F1: -2147483648}
        _ = v10
    default:
        var v11 string = ((("\n\t" + p0) + p0) + p0)
        _ = v11
    }
    if false {
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 bool = true
            v14 := GenS1{F0: 4294967295, F1: fi12}
            for fi15 := 0; fi15 < 1; fi15++ {
                v16 := []int{fi12, fi15}
                v13 = (65535 <= 755)
                _ = fi15
                _ = v16
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        p0 = p0
    }
    _ = p0
    _ = p1
    return 106
}

func main() {
    var v1 int64 = -666
    v2 := make(map[int]int)
    var v3 string = ("foo bar" + ("hello" + ""))
    v4 := make(map[string]uint32)
    var v5 int = 127
    var v6 uint8 = 203
    var v17 int = int(v6)
    if false {
        switch (v5 + v17) {
        case -9223372036854775808:
            switch "hello" {
            case "\n\t":
                v1 -= v1
            case "":
                switch v3 {
                case "hello":
                    v17 = (v5 & v17)
                case "":
                    for fi18 := 0; fi18 < 17; fi18++ {
                        var v19 bool = false
                        for fi20 := 0; fi20 < 5; fi20++ {
                            if true {
                                v21 := v2[v17]
                                v22 := GenS2{F0: (v1 << ((60) & 63))}
                                v19 = (-2147483648 < -1)
                                _ = v21
                                _ = v22
                            } else {
                                switch "a" {
                                case "a":
                                    if (4294967295 != 307) {
                                        v23 := GenS1{F0: 4294967295, F1: 127}
                                        _ = v23
                                    }
                                case "hello":
                                    v24 := GenS3{F0: -47, F1: v3, F2: 450}
                                    _ = v24
                                }
                            }
                            var v25 float32 = 0.0
                            if v19 {
                                var v26 bool = v19
                                var v27 uint = genFn1(("a" + v3), 42)
                                var v28 int16 = -880
                                v29 := []int{(fi18 + fi18)}
                                _ = v26
                                _ = v27
                                _ = v28
                                _ = v29
                            } else {
                                switch v3 {
                                case "\n\t":
                                    var v30 string = (((v3 + v3) + (v3 + (v3 + "\n\t"))) + "")
                                    switch 42 {
                                    case 32767:
                                        v3 = v30
                                    case -130:
                                        if v19 {
                                            v31 := []int{(fi20 % ((v17) | 1))}
                                            delete(v4, v30)
                                            _ = v31
                                        }
                                    }
                                    _ = v30
                                default:
                                    fmt.Println(fi20, v3)
                                }
                                var v32 uint16 = 65535
                                _ = v32
                            }
                            _ = fi20
                            _ = v25
                        }
                        var v33 uint = 1
                        _ = fi18
                        _ = v19
                        _ = v33
                    }
                default:
                    var v34 uint8 = v6
                    _ = v34
                }
            case "hello":
                v5 = v17
            default:
                var v35 bool = false
                _ = v35
            }
        case 9223372036854775807:
            switch v3 {
            case "\n\t":
                for fi36 := 0; fi36 < 9; fi36++ {
                    v37 := GenS4{F0: 0, F1: 4294967295}
                    _ = fi36
                    _ = v37
                }
            }
        case 0:
            v17 = (v5 | v17)
        }
        v6 += 42
        var v38 bool = true
        _ = v38
    }
    fmt.Println(v1, v1, v1)
    fmt.Println(v5, v5)
    v39 := GenS1{F0: 18446744073709551615, F1: (v17 | 289)}
    var v40 float32 = 1e10
    v1 -= (v1 | (v1 | (v1 / (((v1 % (((v1 ^ v1)) | 1))) | 1))))
    for fi41 := 0; fi41 < 8; fi41++ {
        v40 += v40
        v4["\n\t"] = 572
        _ = fi41
    }
    v4[v3] = 774
    var v42 uint = 1
    v43 := GenS5{F0: -1, F1: (1 >= 1)}
    var v44 int = (^(v5))
    for fi45 := 0; fi45 < 3; fi45++ {
        v6 -= uint8(v42)
        v17 += (^(v44))
        _ = fi45
    }
    v17 = (v17 / ((v5) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
