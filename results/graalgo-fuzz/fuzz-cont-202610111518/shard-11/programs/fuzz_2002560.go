package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
type GenS3 struct { F0 int; F1 uint8; F2 int16 }
func genFn0(p0 int) int32 {
    p0 += p0
    p0 += p0
    _ = p0
    return 1
}
func genFn1(p0 uint) uint64 {
    var v7 float32 = 1e10
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: -9223372036854775808, F1: genFn0((^(fi8)))}
        var v10 uint16 = 1
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 uint16 = 255
    var v2 float64 = 0.1
    v3 := make(map[int]bool)
    v4 := []int{-2147483648, -1, 0}
    var v5 int16 = 920
    var v6 uint8 = 1
    fmt.Println(v1)
    var v11 uint16 = 32767
    v2 += -2.25
    var v12 string = ("a" + (("\n\t" + "foo bar") + "\n\t"))
    fmt.Println(v5)
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 bool = false
        _ = fi13
        _ = v14
    }
    v6 = v6
    var v15 float32 = 0.0
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 15; fi17++ {
            fmt.Println(fi16, v1)
            fmt.Println(fi16, v1)
            var v18 float64 = -2.25
            _ = fi17
            _ = v18
        }
        var v19 uint64 = genFn1(464)
        _ = fi16
        _ = v19
    }
    var v20 float64 = v2
    var v21 uint32 = 138
    _ = len(v4)
    if (v20 < v20) {
        v2 = v20
    }
    fmt.Println(v12)
    if ((false || ((v15 / v15) < float32(((0.1 / v20) / (float64(v15) / v2))))) && (127 > 86)) {
        var v22 bool = (!true)
        _ = v22
    } else {
        for fi23 := 0; fi23 < 4; fi23++ {
            if (18446744073709551615 < 18446744073709551615) {
                fmt.Println(v11, fi23, v21)
            } else {
                var v24 int32 = 979
                var v25 uint16 = (v11 & v1)
                _ = v24
                _ = v25
            }
            v4 = append(v4, 95)
            v26 := GenS1{F0: fi23, F1: genFn0(-112)}
            _ = fi23
            _ = v26
        }
    }
    v27 := GenS1{F0: 622, F1: 1}
    switch v12 {
    case "hello":
        if false {
            if true {
                var v28 int = 0
                v29 := GenS2{F0: (v21 >> ((2) & 63)), F1: v21, F2: 577}
                _ = v28
                _ = v29
            } else {
                v4 = append(v4, 127)
            }
        }
    case "foo bar":
        if (!((v6 * (^(v6))) <= 197)) {
            for fi30 := 0; fi30 < 13; fi30++ {
                switch v12 {
                case "":
                    if false {
                        v6 = (v6 & v6)
                        v15 += ((v15 * v15) * 1e10)
                        v20 = v2
                    } else {
                        var v31 int = int(v11)
                        v32 := GenS3{F0: v31, F1: 255, F2: v5}
                        var v33 uint64 = 42
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                }
                _ = fi30
            }
            for fi34 := 0; fi34 < 1; fi34++ {
                v3[0] = false
                for fi35 := 0; fi35 < 20; fi35++ {
                    var v36 float32 = v15
                    var v37 int32 = -2147483648
                    fmt.Println(v36, v20)
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                var v38 uint16 = 32767
                _ = fi34
                _ = v38
            }
        } else {
            var v39 bool = true
            for fi40 := 0; fi40 < 3; fi40++ {
                var v41 uint8 = (^(v6))
                _ = fi40
                _ = v41
            }
            _ = v39
        }
    case "a":
        var v42 float64 = v2
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
