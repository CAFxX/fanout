package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int16; F2 int }
type GenS3 struct { F0 uint64; F1 int8; F2 int64 }
func genFn0() uint {
    var v4 int16 = 121
    var v5 float32 = 1e10
    var v6 uint32 = 776
    fmt.Println(v6, v5)
    if ("" == ("a" + "")) {
        var v7 uint8 = 127
        var v8 uint32 = v6
        switch ("foo bar" + "a") {
        case "\n\t":
            v4 = (v4 | -1)
        }
        _ = v7
        _ = v8
    } else {
        if false {
            for fi9 := 0; fi9 < 1; fi9++ {
                v6 += (v6 / (((v6 << ((4294967295) & 63))) | 1))
                _ = fi9
            }
            var v10 int8 = 42
            v6 = (v6 + (v6 / ((v6) | 1)))
            _ = v10
        } else {
            v6 -= (v6 + v6)
            v6 -= v6
        }
        fmt.Println(v5, v6, v6)
        switch "" {
        case "\n\t":
            v4 -= v4
        case "":
            v6 += 65535
        case "hello":
            var v11 int8 = 2
            _ = v11
        }
    }
    _ = v4
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int, p1 uint16, p2 int64) int {
    var v12 uint8 = 47
    fmt.Println(v12)
    if true {
        for fi13 := 0; fi13 < 3; fi13++ {
            v12 = (v12 >> ((genFn0()) & 63))
            var v14 int8 = -47
            _ = fi13
            _ = v14
        }
        v15 := []int{p0, p0}
        switch (p0 >> ((genFn0()) & 63)) {
        case 32767:
            if ((!((-128 == 1) && false)) || false) {
                if (uint(v12) <= 65535) {
                    fmt.Println(p2, p2)
                } else {
                    var v16 bool = (!(0 >= 42))
                    fmt.Println(v12, p2, v12)
                    _ = v16
                }
                fmt.Println(p2)
            }
        case 127:
            for fi17 := 0; fi17 < 14; fi17++ {
                for fi18 := 0; fi18 < 18; fi18++ {
                    fmt.Println(p1, p0)
                    fmt.Println(p0, v12)
                    _ = fi18
                }
                for fi19 := 0; fi19 < 14; fi19++ {
                    if ((true && (25 < genFn0())) || true) {
                        switch ("\n\t" + (("a" + ("foo bar" + ("a" + ""))) + "a")) {
                        case "a":
                            v20 := GenS1{F0: (1 > 406)}
                            _ = v20
                        default:
                            fmt.Println(p2, v12, fi17)
                        }
                        _ = len(v15)
                    } else {
                        v21 := GenS1{F0: (!true)}
                        if (int8(p2) != 35) {
                            var v22 uint64 = 18446744073709551615
                            v23 := GenS2{F0: p1, F1: 1, F2: 30}
                            _ = v22
                            _ = v23
                        } else {
                            var v24 uint64 = 2
                            p0 += p0
                            _ = v24
                        }
                        _ = v21
                    }
                    _ = fi19
                }
                _ = fi17
            }
        default:
            for fi25 := 0; fi25 < 20; fi25++ {
                if false {
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    fmt.Println(fi25, p0, fi25)
                }
                v26 := GenS2{F0: (^(p1)), F1: 1, F2: fi25}
                _ = fi25
                _ = v26
            }
        }
        _ = v15
    } else {
        if false {
            var v27 int8 = 42
            p1 -= (p1 ^ p1)
            _ = v27
        }
    }
    v28 := []int{(p0 % (((p0 % (((p0 | (p0 / ((p0) | 1)))) | 1))) | 1)), -641}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v28
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{("" + ""), "hello"}
    v3 := make(map[string]string)
    for fi29 := 0; fi29 < 2; fi29++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi30 := 0; fi30 < 7; fi30++ {
            var v31 int64 = 42
            v2 = append(v2, "a")
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        if false {
            var v33 int8 = 0
            _ = len(v2)
            _ = v33
        }
        var v34 int16 = 32767
        _ = fi32
        _ = v34
    }
    var v35 bool = ("" == "hello")
    var v36 uint32 = 977
    v36 += 2147483647
    var v37 int64 = 1
    v37 += 1
    var v38 float64 = -2.25
    v39 := GenS3{F0: 174, F1: 100, F2: v37}
    var v40 uint32 = v36
    switch -569 {
    case -324:
        var v41 uint32 = 2147483647
        _ = v41
    case 137:
        var v42 uint16 = 65535
        _ = v42
    case 9223372036854775807:
        _ = len(v1)
    default:
        v38 = 1e10
    }
    var v43 uint32 = v36
    for fi44 := 0; fi44 < 3; fi44++ {
        v45 := GenS3{F0: 4294967295, F1: -1, F2: v37}
        v2 = append(v2, "\n\t")
        for fi46 := 0; fi46 < 11; fi46++ {
            v36 -= (v40 << ((158) & 63))
            _ = fi46
        }
        _ = fi44
        _ = v45
    }
    v47 := GenS1{F0: (!v35)}
    var v48 uint32 = (v43 + (v36 - v40))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v43
    _ = v47
    _ = v48
}
