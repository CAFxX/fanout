package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint64; F2 uint }
type GenS4 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint, p1 int64) uint8 {
    p0 = p0
    fmt.Println(p1)
    var v6 int8 = 1
    _ = p0
    _ = p1
    _ = v6
    return 119
}
func genFn1(p0 int, p1 int8, p2 int8) uint32 {
    var v7 uint64 = 1
    fmt.Println(v7, p2, p2)
    v8 := []int{p0, 885, p0, (p0 % ((p0) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 870
}

func main() {
    var v1 int16 = 739
    var v2 int64 = 42
    var v3 uint16 = 1
    var v4 uint16 = (v3 | v3)
    v5 := make(map[uint32]int)
    switch ("\n\t" + ("" + "\n\t")) {
    case "a":
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := v5[310]
            v11 := GenS1{F0: v2}
            var v12 int8 = -4
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    v2 = (v2 % (((v2 - v2)) | 1))
    for fi13 := 0; fi13 < 19; fi13++ {
        if (2147483647 == 42) {
            if ((0.1 == (0.1 * (((3.14159 * -2.25) - 0.1) + 1.5))) && false) {
                if false {
                    for fi14 := 0; fi14 < 11; fi14++ {
                        v15 := []int{(fi13 / ((2147483647) | 1))}
                        _ = fi14
                        _ = v15
                    }
                    var v16 int = int(v1)
                    _ = v16
                }
                fmt.Println(v4, fi13)
                fmt.Println(v2, v1, v3)
            } else {
                for fi17 := 0; fi17 < 7; fi17++ {
                    v18 := GenS2{F0: 80}
                    v1 = (v1 | -948)
                    var v19 uint32 = 0
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            v20 := GenS1{F0: 685}
            _ = v20
        }
        for fi21 := 0; fi21 < 12; fi21++ {
            v22 := v5[2]
            _ = fi21
            _ = v22
        }
        v23 := GenS3{F0: "\n\t", F1: 18446744073709551615, F2: 18446744073709551615}
        _ = fi13
        _ = v23
    }
    var v24 uint32 = 275
    switch 2147483647 {
    case 751:
        var v25 int32 = -1
        _ = v25
    case 0:
        var v26 bool = (1 > -1)
        _ = v26
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        v2 -= v2
        _ = fi27
    }
    var v28 int32 = -48
    switch 4294967295 {
    case 4294967295:
        for fi29 := 0; fi29 < 17; fi29++ {
            var v30 int32 = int32(fi29)
            for fi31 := 0; fi31 < 8; fi31++ {
                v2 = v2
                v24 = 2
                delete(v5, (v24 << ((65535) & 63)))
                _ = fi31
            }
            _ = fi29
            _ = v30
        }
    case 2147483647:
        v5[(v24 - (v24 ^ (v24 + v24)))] = 9223372036854775807
    default:
        var v32 string = "\n\t"
        _ = v32
    }
    if false {
        var v33 uint32 = (^(v24))
        var v34 uint8 = 174
        _ = v33
        _ = v34
    }
    fmt.Println(v1)
    fmt.Println(v1, v28)
    fmt.Println(v28, v4)
    if ("" != "") {
        var v35 float64 = 0.1
        _ = v35
    } else {
        switch -9223372036854775808 {
        case -2147483648:
            v36 := GenS3{F0: "a", F1: 42, F2: 0}
            _ = v36
        case 32767:
            v37 := GenS2{F0: (v28 * v28)}
            _ = v37
        default:
            v2 += 360
        }
        v38 := GenS4{F0: v24, F1: (v28 / (((v28 & v28)) | 1))}
        fmt.Println(v3, v1, v3)
        _ = v38
    }
    _ = len(v5)
    v39 := GenS2{F0: v28}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
