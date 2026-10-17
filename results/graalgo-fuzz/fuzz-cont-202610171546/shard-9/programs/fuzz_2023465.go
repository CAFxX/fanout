package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS4 struct { F0 bool }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: (-9223372036854775808 < -9223372036854775808), F1: 74, F2: ("a" + "hello")}
        _ = fi6
        _ = v7
    }
    if ("hello" <= ("a" + ("" + "foo bar"))) {
        if false {
            var v8 int = 2147483647
            _ = v8
        }
        v9 := GenS2{F0: 127}
        v10 := GenS2{F0: 127}
        _ = v10
        _ = v9
    }
    return 861
}
func genFn1(p0 int16, p1 int16, p2 uint64) uint8 {
    fmt.Println(p1, p0)
    var v11 uint8 = 0
    var v12 int32 = -65
    if (70 != 948) {
        if (127 > 358) {
            v13 := []int{957}
            v14 := GenS1{F0: (false && (p2 > 4294967295)), F1: v11, F2: "\n\t"}
            _ = v13
            _ = v14
        } else {
            if ((v12 | v12) == -2147483648) {
                var v15 int = 2147483647
                for fi16 := 0; fi16 < 5; fi16++ {
                    p0 -= p1
                    _ = fi16
                }
                if (!(!(0.1 != (((float32(float64((0.1 * 0.0))) + float32(1.5)) + 3.14159) + 0.0)))) {
                    p1 -= (p1 | (p0 | (p0 ^ (p1 * p0))))
                }
                _ = v15
            }
            switch "foo bar" {
            case "hello":
                fmt.Println(p0, p2)
            case "":
                for fi17 := 0; fi17 < 3; fi17++ {
                    v11 = (v11 | (^(v11)))
                    var v18 float64 = 0.1
                    _ = fi17
                    _ = v18
                }
            }
            fmt.Println(p0, v11, p0)
        }
        if false {
            var v19 int64 = -938
            _ = v19
        } else {
            v12 -= v12
        }
    } else {
        var v20 bool = (false && (((p2 & 839) == genFn0()) || (3.14159 < (3.14159 - (0.1 * float64(-2.25))))))
        var v21 uint64 = p2
        if (v20 && (v20 || (1e10 > float64(-2.25)))) {
            for fi22 := 0; fi22 < 19; fi22++ {
                if v20 {
                    switch "a" {
                    case "a":
                        v23 := []int{fi22, fi22}
                        _ = v23
                    default:
                        p2 -= (^(p2))
                    }
                    var v24 int = 32767
                    v11 += v11
                    _ = v24
                } else {
                    v25 := GenS3{F0: genFn0(), F1: (v12 ^ v12), F2: (v21 >> ((2) & 63))}
                    _ = v25
                }
                _ = fi22
            }
            for fi26 := 0; fi26 < 14; fi26++ {
                v27 := GenS3{F0: v21, F1: v12, F2: v21}
                var v28 int8 = 127
                v20 = false
                _ = fi26
                _ = v27
                _ = v28
            }
            if true {
                var v29 uint16 = 255
                _ = v29
            } else {
                p1 = (p1 >> ((2) & 63))
                v20 = (0 != 4294967295)
            }
        }
        _ = v20
        _ = v21
    }
    var v30 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v30
    return (v11 ^ (^(v11)))
}

func main() {
    var v1 float32 = 1e10
    v2 := []string{}
    var v3 uint = 42
    v4 := []uint32{}
    v5 := []int{2147483647, 4294967295, -1, int(v3)}
    var v31 int32 = 120
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v32 float64 = 1e10
    for fi33 := 0; fi33 < 17; fi33++ {
        v34 := GenS3{F0: 332, F1: 2147483647, F2: uint64(v3)}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi33
        _ = v34
    }
    v35 := GenS1{F0: (v3 >= (v3 / ((v3) | 1))), F1: genFn1(1000, -204, 940), F2: "\n\t"}
    for fi36 := 0; fi36 < 2; fi36++ {
        switch (-(fi36)) {
        case 138:
            var v37 int = fi36
            _ = v37
        default:
            var v38 float32 = v1
            _ = v38
        }
        if ((!(true && true)) && true) {
            var v39 float32 = v1
            var v40 uint16 = 0
            v41 := GenS4{F0: true}
            _ = v39
            _ = v40
            _ = v41
        }
        _ = fi36
    }
    fmt.Println(v1, v1)
    v32 = (v32 + 0.1)
    switch "" {
    case "a":
        _ = len(v5)
    case "\n\t":
        for fi42 := 0; fi42 < 10; fi42++ {
            var v43 bool = (330 > 187)
            _ = fi42
            _ = v43
        }
    default:
        if (!false) {
            if len(v4) > 0 {
                v4[0] = 2147483648
            }
            v2 = append(v2, "a")
            if len(v2) > 0 {
                v2[0] = ""
            }
        }
    }
    v1 -= -2.25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v5
}
