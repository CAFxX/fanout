package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 int }
type GenS3 struct { F0 int32; F1 int64; F2 uint }
type GenS4 struct { F0 int8 }
func genFn0(p0 int, p1 string) uint {
    var v4 int16 = -881
    p0 = -2147483648
    if false {
        for fi5 := 0; fi5 < 15; fi5++ {
            if (32767 < 32767) {
                var v6 uint8 = 89
                _ = v6
            } else {
                p1 = "hello"
                v7 := []int{fi5, (fi5 + (-(fi5))), fi5, (fi5 << ((65535) & 63))}
                switch fi5 {
                case 4294967295:
                    _ = len(v7)
                case 127:
                    for fi8 := 0; fi8 < 17; fi8++ {
                        v4 += v4
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        v4 += -199
                        _ = fi8
                    }
                case 128:
                    v7 = append(v7, -2147483648)
                }
                _ = v7
            }
            for fi9 := 0; fi9 < 8; fi9++ {
                if (p1 < p1) {
                    v4 += v4
                    var v10 int32 = 2147483647
                    fmt.Println(p1)
                    v11 := GenS1{F0: 2, F1: v4, F2: 1}
                    _ = v10
                    _ = v11
                }
                var v12 uint64 = 42
                var v13 int16 = -548
                v14 := []int{9223372036854775807, fi5, 576, fi9}
                v15 := GenS1{F0: 4294967295, F1: (^(v4)), F2: 232}
                _ = fi9
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            }
            v4 = v4
            _ = fi5
        }
    } else {
        v16 := GenS2{F0: 1, F1: (p0 + 4294967295), F2: p0}
        p0 = (p0 ^ p0)
        _ = v16
    }
    var v17 uint64 = 959
    _ = p0
    _ = p1
    _ = v17
    _ = v4
    return 0
}
func genFn1(p0 string) uint32 {
    fmt.Println(p0)
    var v18 int16 = 0
    v19 := []int{42, 615, -314, int(v18)}
    switch 4294967295 {
    case -1:
        p0 = (("\n\t" + (("" + "foo bar") + ((("" + ("a" + p0)) + ("hello" + ("\n\t" + (p0 + p0)))) + p0))) + p0)
    }
    _ = p0
    _ = v18
    _ = v19
    return 2147483647
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{-9223372036854775808, 68, 127, 32768}
    var v3 float32 = ((1.5 - 0.0) * 1e10)
    v20 := v1["a"]
    v21 := GenS1{F0: 2, F1: 1, F2: 0}
    v20 = v20
    var v22 float32 = (float32(3.14159) + -2.25)
    switch v20 {
    case 32767:
        if ((-631 > -490) && false) {
            for fi23 := 0; fi23 < 7; fi23++ {
                v3 = (((1e10 + v3) * v22) + -2.25)
                fmt.Println(v20, v20, v3)
                if (4294967295 < 298) {
                    v24 := GenS2{F0: 65535, F1: (fi23 << ((genFn0(fi23, "hello")) & 63)), F2: (v20 % ((fi23) | 1))}
                    v25 := GenS3{F0: 0, F1: -9223372036854775808, F2: 0}
                    _ = v24
                    _ = v25
                }
                _ = fi23
            }
            v26 := GenS3{F0: 1, F1: -9223372036854775808, F2: 4294967295}
            v27 := v1[("a" + "\n\t")]
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 4; fi28++ {
                fmt.Println(v20, v22, v22)
                _ = fi28
            }
            v29 := v1["hello"]
            v30 := GenS1{F0: 65535, F1: -137, F2: 42}
            _ = v29
            _ = v30
        }
    }
    var v31 uint64 = 590
    var v32 int64 = 1
    for fi33 := 0; fi33 < 7; fi33++ {
        for fi34 := 0; fi34 < 19; fi34++ {
            if false {
                if (-1 < 0) {
                    v35 := GenS4{F0: -128}
                    v36 := GenS2{F0: genFn1(("a" + "a")), F1: v20, F2: (v20 ^ fi33)}
                    _ = v35
                    _ = v36
                }
                var v37 int = fi34
                v32 = (v32 >> ((2) & 63))
                v38 := GenS2{F0: 1, F1: (fi33 << ((0) & 63)), F2: 1}
                v20 += -175
                switch "foo bar" {
                case "a":
                    v31 -= (^(v31))
                }
                _ = v37
                _ = v38
            } else {
                v1["foo bar"] = -596
            }
            _ = fi34
        }
        for fi39 := 0; fi39 < 17; fi39++ {
            var v40 float32 = 0.0
            _ = fi39
            _ = v40
        }
        _ = fi33
    }
    v31 = v31
    delete(v1, ((("a" + "foo bar") + "a") + ""))
    var v41 string = "\n\t"
    delete(v1, "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v41
}
