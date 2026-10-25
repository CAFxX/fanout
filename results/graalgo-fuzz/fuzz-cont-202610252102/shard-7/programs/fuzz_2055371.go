package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }
type GenS2 struct { F0 uint8; F1 int; F2 int8 }
type GenS3 struct { F0 int8; F1 bool; F2 int }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 string; F1 string; F2 int64 }
func genFn0(p0 uint, p1 int32) uint16 {
    var v4 int32 = (^(p1))
    if (!true) {
        fmt.Println(p1)
        v5 := GenS1{F0: 255, F1: 338, F2: 32767}
        _ = v5
    }
    if false {
        var v6 int = 857
        v6 = v6
        _ = v6
    } else {
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := GenS1{F0: 0, F1: 42, F2: fi7}
            if ((p0 ^ 65535) == p0) {
                v9 := []int{(fi7 << ((832) & 63)), (fi7 | 127)}
                v4 = -2147483648
                _ = v9
            } else {
                var v10 string = "foo bar"
                switch "" {
                case "hello":
                    p0 += (p0 % ((p0) | 1))
                default:
                    v10 = (v10 + v10)
                }
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        v11 := GenS1{F0: 42, F1: 128, F2: 128}
        _ = v11
    }
    fmt.Println(p1)
    v12 := GenS2{F0: 163, F1: -734, F2: 35}
    _ = p0
    _ = p1
    _ = v12
    _ = v4
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 0
    var v3 float32 = 1.5
    switch 0 {
    case 128:
        var v13 uint32 = (v2 / (((v2 | v2)) | 1))
        _ = v13
    default:
        v2 = v2
    }
    switch 1 {
    case 127:
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := GenS2{F0: 42, F1: (^(fi14)), F2: -19}
            v16 := []int{1, 127}
            _ = fi14
            _ = v15
            _ = v16
        }
    case 32768:
        var v17 int = 0
        _ = v17
    }
    if false {
        v2 -= v2
        var v18 int32 = 511
        v19 := []int{32767, 127}
        _ = v18
        _ = v19
    } else {
        fmt.Println(v2, v3)
        var v20 bool = (false && false)
        v21 := GenS1{F0: 4294967295, F1: 32767, F2: 599}
        _ = v20
        _ = v21
    }
    var v22 int32 = -833
    for fi23 := 0; fi23 < 14; fi23++ {
        v2 -= v2
        var v24 float64 = 1e10
        _ = fi23
        _ = v24
    }
    v22 = v22
    var v25 bool = (!(!(false || false)))
    var v26 int32 = v22
    v27 := GenS3{F0: -44, F1: true, F2: 32767}
    v28 := []int{715, 32767, 128, 1}
    switch 440 {
    case 127:
        v29 := GenS1{F0: 4294967295, F1: 32767, F2: 452}
        _ = v29
    case -67:
        var v30 uint64 = 42
        _ = v30
    case 32768:
        if v25 {
            switch -9223372036854775808 {
            case 4294967295:
                if (162 >= 18446744073709551615) {
                    if v25 {
                        for fi31 := 0; fi31 < 2; fi31++ {
                            delete(v1, v2)
                            _ = fi31
                        }
                        v26 = v26
                    } else {
                        v25 = v25
                    }
                    var v32 string = "foo bar"
                    _ = v32
                }
            case 32767:
                if false {
                    v3 = 1e10
                    var v33 int64 = 924
                    _ = v33
                }
            default:
                v2 = 2147483648
            }
            v34 := GenS4{F0: -631, F1: (v25 && true)}
            for fi35 := 0; fi35 < 3; fi35++ {
                v22 = (v26 | -180)
                _ = fi35
            }
            _ = v34
        }
    default:
        v22 -= (v22 + (^(v22)))
    }
    switch -2147483648 {
    case 18:
        v36 := GenS5{F0: "\n\t", F1: "hello", F2: -9223372036854775808}
        _ = v36
    case 128:
        var v37 int8 = 70
        _ = v37
    case 9223372036854775807:
        _ = len(v1)
    default:
        var v38 float32 = v3
        _ = v38
    }
    v2 = 947
    var v39 int16 = -1
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    v26 = v22
    switch 65 {
    case 0:
        switch ("\n\t" + ("hello" + "")) {
        case "hello":
            fmt.Println(v3)
        case "foo bar":
            fmt.Println(v39, v22, v22)
        case "":
            for fi40 := 0; fi40 < 15; fi40++ {
                var v41 uint16 = 2
                _ = fi40
                _ = v41
            }
        }
    default:
        var v42 bool = ((v25 && (v3 < 3.14159)) && (4294967295 == 2))
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v39
}
