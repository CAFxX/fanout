package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 bool; F2 int32 }
type GenS5 struct { F0 int32; F1 string; F2 string }
func genFn0(p0 uint16, p1 int8, p2 int32) int16 {
    fmt.Println(p1, p0)
    fmt.Println(p0, p1)
    var v6 string = ("hello" + "\n\t")
    p1 = (p1 - -128)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 32767
}
func genFn1(p0 uint16, p1 int8) int64 {
    v7 := GenS1{F0: 9223372036854775807}
    v8 := []int{-9223372036854775808, 127}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    v1 := []int{690}
    v2 := make(map[string]string)
    var v3 bool = (103 < 42)
    v4 := make(map[uint32]int)
    var v5 int16 = 42
    fmt.Println(v3)
    _ = len(v1)
    var v9 int16 = (v5 ^ v5)
    switch 4294967295 {
    case -2147483648:
        var v10 bool = v3
        _ = v10
    case -9223372036854775808:
        v3 = (!(v3 && true))
    case 940:
        var v11 string = "a"
        _ = v11
    default:
        if (1 >= 0) {
            var v12 uint64 = 18446744073709551615
            _ = v12
        }
    }
    switch 0 {
    case -971:
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := v4[1]
            v15 := GenS1{F0: 0}
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 11; fi17++ {
            switch fi16 {
            case 127:
                v18 := GenS2{F0: 0, F1: -6}
                _ = v18
            case 42:
                fmt.Println(v3)
            }
            v19 := GenS3{F0: 0}
            _ = fi17
            _ = v19
        }
        _ = fi16
    }
    v20 := GenS3{F0: 751}
    v21 := GenS2{F0: 32767, F1: 73}
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 uint8 = 1
        fmt.Println(v9, v23, v9)
        _ = fi22
        _ = v23
    }
    delete(v2, "\n\t")
    v24 := GenS1{F0: -1}
    var v25 string = "foo bar"
    for fi26 := 0; fi26 < 7; fi26++ {
        v2[v25] = v25
        var v27 uint32 = 0
        v5 += v5
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        for fi29 := 0; fi29 < 14; fi29++ {
            v30 := GenS2{F0: 970, F1: -1}
            v31 := GenS4{F0: 4294967295, F1: false, F2: 211}
            if ((!((!v3) && (12 == 4294967295))) && ((1.5 / 1e10) >= ((((-2.25 + 3.14159) * -2.25) / 0.1) + (1e10 + ((((3.14159 + ((-2.25 / 1e10) * 0.0)) * 0.1) + (0.1 + (((3.14159 * -2.25) + float32((float64((-2.25 * -2.25)) / 1.0))) / 3.14159))) - (1e10 * 1e10)))))) {
                var v32 uint64 = 42
                v32 += (v32 - v32)
                v2[""] = v25
                v33 := GenS3{F0: 2147483648}
                switch (fi29 / (((fi29 & fi28)) | 1)) {
                case 629:
                    fmt.Println(v9)
                case 128:
                    if len(v1) > 0 {
                        v1[0] = fi29
                    }
                case 2147483647:
                    if ("\n\t" >= v25) {
                        switch (v25 + "hello") {
                        case "":
                            if len(v1) > 0 {
                                v1[0] = (^(fi28))
                            }
                        case "hello":
                            switch (fi28 & (fi29 - 336)) {
                            case 9223372036854775807:
                                v3 = v3
                            }
                        case "foo bar":
                            delete(v2, "foo bar")
                        default:
                            v32 = 18446744073709551615
                        }
                        if (255 < 42) {
                            v34 := GenS3{F0: 65535}
                            var v35 int = fi29
                            _ = v34
                            _ = v35
                        }
                        v36 := v4[42]
                        _ = v36
                    } else {
                        v9 -= v9
                        fmt.Println(v3, fi28, v32)
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v32
                _ = v33
            } else {
                v5 -= (v9 ^ (-(v9)))
                v3 = v3
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        v37 := GenS5{F0: 506, F1: v25, F2: (((v25 + ((("foo bar" + ("hello" + v25)) + (v25 + v25)) + "hello")) + v25) + v25)}
        _ = fi28
        _ = v37
    }
    if (-686 == 461) {
        if true {
            v9 = (^(v9))
            if len(v1) > 0 {
                v1[0] = 32767
            }
            if true {
                switch -113 {
                case 127:
                    var v38 uint32 = 42
                    _ = v38
                }
            }
        } else {
            _ = len(v2)
            v5 = (v9 % (((-(v9))) | 1))
        }
        switch int(v5) {
        case 621:
            var v39 uint = 42
            _ = v39
        case -854:
            v40 := GenS2{F0: 32767, F1: -107}
            _ = v40
        case 0:
            v25 = "foo bar"
        default:
            var v41 int32 = -1
            _ = v41
        }
        v1 = append(v1, -756)
    } else {
        _ = len(v1)
        v3 = v3
        for fi42 := 0; fi42 < 16; fi42++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v43 int = -132
            _ = fi42
            _ = v43
        }
    }
    v44 := GenS3{F0: 2}
    v45 := GenS5{F0: 523, F1: v25, F2: v25}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v9
}
