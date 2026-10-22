package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32 }
func genFn0() string {
    var v7 float32 = 0.1
    for fi8 := 0; fi8 < 9; fi8++ {
        if (42 == 42) {
            v9 := GenS1{F0: 0, F1: 1}
            switch fi8 {
            case 32768:
                var v10 int = 42
                _ = v10
            case 0:
                var v11 int64 = 9223372036854775807
                _ = v11
            case 42:
                switch "foo bar" {
                case "\n\t":
                    var v12 int16 = -715
                    _ = v12
                case "a":
                    switch ("a" + "\n\t") {
                    case "hello":
                        if (!true) {
                            v13 := GenS1{F0: 42, F1: 396}
                            _ = v13
                        } else {
                            if true {
                                var v14 uint16 = 2
                                v15 := GenS2{F0: 0}
                                _ = v14
                                _ = v15
                            } else {
                                v16 := []int{128}
                                _ = v16
                            }
                            v17 := GenS2{F0: 42}
                            v18 := []int{fi8, (fi8 & 42)}
                            _ = v17
                            _ = v18
                        }
                    case "":
                        var v19 float64 = (1.5 + 3.14159)
                        _ = v19
                    case "a":
                        var v20 int8 = -128
                        _ = v20
                    default:
                        var v21 uint = 4294967295
                        _ = v21
                    }
                }
            }
            for fi22 := 0; fi22 < 11; fi22++ {
                if false {
                    fmt.Println(fi8)
                } else {
                    var v23 uint64 = 0
                    _ = v23
                }
                v24 := []int{(fi8 - fi8), (fi8 & -690), 2147483647}
                _ = fi22
                _ = v24
            }
            _ = v9
        } else {
            fmt.Println(fi8)
            var v25 float64 = float64((v7 + (v7 * (v7 - (-2.25 * v7)))))
            _ = v25
        }
        if (-1 <= 1000) {
            if true {
                for fi26 := 0; fi26 < 14; fi26++ {
                    var v27 uint32 = 65535
                    var v28 int = -539
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            } else {
                v7 = 0.1
            }
            v29 := GenS2{F0: -128}
            var v30 int16 = -32768
            _ = v29
            _ = v30
        }
        switch fi8 {
        case 2147483647:
            v31 := []int{fi8, (fi8 << ((4294967295) & 63))}
            _ = v31
        case -199:
            var v32 float64 = float64(0.1)
            _ = v32
        }
        _ = fi8
    }
    v33 := GenS3{F0: 32767, F1: 42, F2: 651}
    _ = v33
    _ = v7
    return "a"
}
func genFn1(p0 int8, p1 uint32, p2 uint16) uint32 {
    v34 := []int{-1, 0}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    return 764
}

func main() {
    var v1 int64 = 42
    var v2 int16 = 195
    v3 := make(map[string]int)
    v4 := []string{"foo bar"}
    var v5 int = 9223372036854775807
    v6 := []string{"hello", "", "a", "a"}
    var v35 uint = 44
    if (false && (!true)) {
        for fi36 := 0; fi36 < 17; fi36++ {
            for fi37 := 0; fi37 < 15; fi37++ {
                if true {
                    var v38 bool = true
                    _ = v38
                } else {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v1 += v1
                    v3["foo bar"] = 32767
                    v1 += v1
                }
                _ = fi37
            }
            _ = fi36
        }
    } else {
        fmt.Println(v5, v35, v2)
        fmt.Println(v35, v1, v35)
        v39 := GenS4{F0: 1}
        _ = v39
    }
    v40 := GenS4{F0: 1}
    v41 := v3["hello"]
    var v42 float64 = float64(0.0)
    v43 := GenS5{F0: genFn1(-116, 1, 649)}
    v41 -= (-(v41))
    v2 += v2
    var v44 uint = (v35 >> ((65535) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
