package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint, p1 int16, p2 uint) uint64 {
    switch "hello" {
    case "hello":
        for fi6 := 0; fi6 < 4; fi6++ {
            if false {
                if true {
                    fmt.Println(p2)
                }
                if true {
                    v7 := []int{}
                    fmt.Println(p2, fi6)
                    var v8 uint16 = 626
                    _ = v7
                    _ = v8
                }
                switch "a" {
                case "a":
                    v9 := GenS1{F0: 42}
                    _ = v9
                case "":
                    v10 := GenS1{F0: 0}
                    _ = v10
                default:
                    v11 := []int{(fi6 - fi6)}
                    _ = v11
                }
            } else {
                v12 := []int{fi6, (fi6 / (((^(fi6))) | 1)), fi6, fi6}
                switch fi6 {
                case 9223372036854775807:
                    if len(v12) > 0 {
                        v12[0] = -9223372036854775808
                    }
                case -883:
                    if false {
                        v13 := GenS2{F0: 78, F1: 1, F2: 71}
                        _ = v13
                    } else {
                        var v14 float32 = (((-2.25 + 1e10) * 0.1) + ((-2.25 + (3.14159 - (0.0 / -2.25))) + ((-2.25 / 1.5) + 1.5)))
                        fmt.Println(p0)
                        _ = v14
                    }
                default:
                    for fi15 := 0; fi15 < 9; fi15++ {
                        var v16 int8 = 42
                        var v17 int32 = 0
                        var v18 int16 = 32767
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                }
                _ = v12
            }
            _ = fi6
        }
    case "foo bar":
        p0 = (^(p0))
    }
    v19 := GenS2{F0: -1, F1: 0, F2: -81}
    v20 := []int{}
    for fi21 := 0; fi21 < 4; fi21++ {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        v22 := GenS1{F0: 42}
        _ = fi21
        _ = v22
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 846
}
func genFn1(p0 int16) int32 {
    var v23 uint32 = 568
    for fi24 := 0; fi24 < 12; fi24++ {
        v25 := []int{(fi24 ^ -9223372036854775808)}
        fmt.Println(p0)
        var v26 int32 = 2147483647
        _ = fi24
        _ = v25
        _ = v26
    }
    v23 -= v23
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := []int{fi27}
        var v29 uint32 = 379
        for fi30 := 0; fi30 < 6; fi30++ {
            v31 := GenS1{F0: 572}
            v32 := GenS1{F0: 0}
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    _ = p0
    _ = v23
    return -2147483648
}

func main() {
    var v1 float64 = 1.5
    var v2 bool = (27 == 0)
    var v3 uint16 = 1
    v4 := []int{}
    v5 := []uint32{2147483648, 0, 777}
    for fi33 := 0; fi33 < 19; fi33++ {
        fmt.Println(v2, v2)
        v2 = v2
        _ = fi33
    }
    switch 4294967295 {
    case -2147483648:
        var v34 float32 = 1e10
        _ = v34
    case 0:
        if ((!true) && true) {
            switch "hello" {
            case "a":
                v35 := GenS3{F0: 42}
                _ = v35
            case "hello":
                fmt.Println(v2, v3, v2)
            }
        }
    case 32768:
        switch "a" {
        case "\n\t":
            var v36 bool = v2
            _ = v36
        case "":
            v2 = (("\n\t" <= "a") && v2)
        }
    default:
        v3 += v3
    }
    if len(v5) > 0 {
        v5[0] = 2
    }
    var v37 string = "a"
    if (v2 && (-307 == -1)) {
        fmt.Println(v3)
        var v38 string = "hello"
        _ = v38
    }
    v39 := GenS1{F0: (v3 % ((234) | 1))}
    for fi40 := 0; fi40 < 4; fi40++ {
        for fi41 := 0; fi41 < 2; fi41++ {
            var v42 int = fi40
            var v43 int8 = 42
            v3 = (v3 & (v3 + 39))
            _ = fi41
            _ = v42
            _ = v43
        }
        v3 -= (v3 ^ v3)
        v2 = ((v1 / 3.14159) >= 0.0)
        _ = fi40
    }
    var v44 string = v37
    v1 = float64(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v44
    _ = v5
}
