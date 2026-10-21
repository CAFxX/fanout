package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int; F2 bool }
type GenS4 struct { F0 uint8; F1 int; F2 string }
func genFn0() uint16 {
    var v7 float32 = 3.14159
    v7 -= (float32(1e10) - v7)
    for fi8 := 0; fi8 < 7; fi8++ {
        fmt.Println(v7, fi8)
        switch fi8 {
        case 4294967295:
            v9 := []int{(fi8 & fi8), 4294967295, fi8}
            _ = v9
        default:
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS1{F0: 42}
                v7 = 1.5
                var v12 uint64 = 0
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        _ = fi8
    }
    fmt.Println(v7)
    v7 += 1.5
    _ = v7
    return 32767
}
func genFn1(p0 int64, p1 string, p2 uint64) string {
    fmt.Println(p0, p1, p0)
    for fi13 := 0; fi13 < 12; fi13++ {
        p0 = p0
        for fi14 := 0; fi14 < 9; fi14++ {
            fmt.Println(p2, p1, fi14)
            var v15 int8 = 100
            v16 := []int{(fi13 * 0)}
            _ = fi14
            _ = v15
            _ = v16
        }
        fmt.Println(fi13, p2)
        _ = fi13
    }
    p2 = (p2 + p2)
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    v1 := []uint32{}
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    var v4 int64 = 42
    v5 := []string{"foo bar", ""}
    v6 := []int{}
    fmt.Println(v4)
    v4 -= (v4 | v4)
    switch -9223372036854775808 {
    case -9223372036854775808:
        v4 = v4
    case 4294967295:
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 float32 = 1.5
            var v19 float32 = v18
            _ = fi17
            _ = v18
            _ = v19
        }
    case 32767:
        v4 = (v4 >> ((288) & 63))
    default:
        v20 := GenS2{F0: 1}
        _ = v20
    }
    var v21 bool = (2 == 1)
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    var v22 uint = 4294967295
    for fi23 := 0; fi23 < 17; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            v25 := GenS3{F0: "\n\t", F1: 856, F2: (944 < 4294967295)}
            v21 = v21
            if v21 {
                var v26 uint16 = genFn0()
                var v27 int8 = 0
                _ = v26
                _ = v27
            } else {
                if v21 {
                    if v21 {
                        v22 = 42
                    }
                    v22 -= (^(v22))
                    _ = len(v6)
                }
                if (!(!(("foo bar" + ("a" + "foo bar")) <= ""))) {
                    var v28 float32 = float32(float64(0.0))
                    if v21 {
                        v22 = (v22 * v22)
                        var v29 bool = v21
                        var v30 string = "hello"
                        _ = v29
                        _ = v30
                    }
                    if (v4 > v4) {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v31 int = (^(fi24))
                        var v32 float32 = v28
                        _ = v31
                        _ = v32
                    }
                    var v33 uint32 = 42
                    _ = v28
                    _ = v33
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi24
            _ = v25
        }
        for fi34 := 0; fi34 < 6; fi34++ {
            v4 = (v4 << (((^(v22))) & 63))
            var v35 int64 = 0
            var v36 uint8 = 0
            var v37 bool = true
            var v38 uint64 = 18446744073709551615
            var v39 int8 = 42
            var v40 uint32 = uint32(v36)
            var v41 int = 42
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
            _ = v38
            _ = v39
            _ = v40
            _ = v41
        }
        _ = len(v3)
        _ = fi23
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        for fi43 := 0; fi43 < 20; fi43++ {
            v44 := GenS4{F0: 109, F1: (fi42 << ((255) & 63)), F2: "foo bar"}
            _ = fi43
            _ = v44
        }
        var v45 uint8 = 128
        v2[(("\n\t" + "") + "\n\t")] = ((v4 + 1) == 1)
        _ = fi42
        _ = v45
    }
    fmt.Println(v4, v4)
    delete(v2, "foo bar")
    var v46 int = -259
    var v47 int32 = 2147483647
    var v48 uint32 = 42
    v49 := v2["hello"]
    v47 = (v47 ^ v47)
    var v50 float32 = 0.0
    var v51 uint32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
