package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 bool; F1 int; F2 int64 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int16, p1 int) int16 {
    var v6 uint = 4294967295
    v7 := GenS1{F0: 1, F1: -1}
    for fi8 := 0; fi8 < 1; fi8++ {
        p0 = p0
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -309
}
func genFn1() bool {
    var v9 int64 = 0
    v10 := GenS2{F0: false, F1: int(v9), F2: -373}
    switch 229 {
    case -212:
        for fi11 := 0; fi11 < 9; fi11++ {
            fmt.Println(fi11, fi11)
            v12 := []int{}
            _ = len(v12)
            _ = fi11
            _ = v12
        }
    case -670:
        if false {
            v9 += v9
        } else {
            fmt.Println(v9)
            v9 = v9
            for fi13 := 0; fi13 < 13; fi13++ {
                if (230 >= 918) {
                    v14 := []int{}
                    fmt.Println(fi13)
                    var v15 bool = (-128 < -1)
                    _ = v14
                    _ = v15
                }
                var v16 int16 = -423
                for fi17 := 0; fi17 < 14; fi17++ {
                    v16 = (v16 / (((^(v16))) | 1))
                    switch 748 {
                    case 32767:
                        fmt.Println(fi13)
                    }
                    _ = fi17
                }
                _ = fi13
                _ = v16
            }
        }
    }
    fmt.Println(v9)
    var v18 bool = (65535 == 42)
    _ = v10
    _ = v18
    _ = v9
    return (0.0 == -2.25)
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483648}
    var v3 string = "a"
    var v4 uint32 = 65535
    v5 := []int{}
    var v19 string = (v3 + ("\n\t" + (v3 + "\n\t")))
    if false {
        fmt.Println(v19, v19)
        var v20 float32 = float32((0.0 - 1e10))
        v4 = (v4 - (v4 / ((v4) | 1)))
        _ = v20
    } else {
        var v21 string = ("" + "hello")
        for fi22 := 0; fi22 < 17; fi22++ {
            if true {
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    delete(v1, (fi22 ^ -641))
                    v19 = v19
                }
            } else {
                v5 = append(v5, 9223372036854775807)
                if (42 > -81) {
                    v4 = 2147483647
                    v19 = "hello"
                    v2 = append(v2, (v4 >> ((65535) & 63)))
                } else {
                    fmt.Println(v4, v21, fi22)
                    var v23 int16 = 32767
                    for fi24 := 0; fi24 < 6; fi24++ {
                        var v25 uint16 = 1
                        var v26 float32 = (0.1 + 1e10)
                        v4 = (v4 & v4)
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    _ = v23
                }
                if true {
                    if true {
                        v1[fi22] = fi22
                        fmt.Println(v4, v19, fi22)
                        var v27 uint64 = 18446744073709551615
                        _ = v27
                    }
                    var v28 string = v3
                    var v29 uint16 = 481
                    _ = v28
                    _ = v29
                }
            }
            var v30 int32 = -2147483648
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi22
            _ = v30
        }
        _ = v21
    }
    switch 127 {
    case 9223372036854775807:
        var v31 uint = 935
        _ = v31
    }
    switch (("hello" + v3) + (v3 + "a")) {
    case "hello":
        var v32 float32 = -2.25
        _ = v32
    default:
        for fi33 := 0; fi33 < 20; fi33++ {
            fmt.Println(fi33, fi33, v4)
            _ = fi33
        }
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        var v35 int64 = 0
        _ = len(v5)
        fmt.Println(v19, v3, v35)
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        if (!(v4 > (v4 / ((v4) | 1)))) {
            var v37 int16 = -779
            var v38 uint = 0
            v1[fi36] = 65
            _ = v37
            _ = v38
        } else {
            for fi39 := 0; fi39 < 9; fi39++ {
                v40 := GenS3{F0: 928, F1: 42}
                if len(v5) > 0 {
                    v5[0] = (-(fi36))
                }
                v41 := GenS4{F0: (^(v4)), F1: v4}
                _ = fi39
                _ = v40
                _ = v41
            }
            v42 := GenS4{F0: v4, F1: (v4 >> ((2) & 63))}
            v3 = ((v19 + "\n\t") + v3)
            _ = v42
        }
        var v43 int = fi36
        _ = fi36
        _ = v43
    }
    delete(v1, 198)
    for fi44 := 0; fi44 < 14; fi44++ {
        var v45 uint16 = 65535
        _ = fi44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
