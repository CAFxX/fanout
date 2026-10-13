package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 string; F2 uint32 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int }
func genFn0(p0 uint8, p1 uint64, p2 uint64) uint64 {
    v4 := []int{32767}
    p2 += 42
    for fi5 := 0; fi5 < 2; fi5++ {
        fmt.Println(fi5, p1)
        _ = fi5
    }
    var v6 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return (p1 + p2)
}
func genFn1(p0 uint, p1 int32) int32 {
    v7 := []int{}
    if false {
        v8 := GenS1{F0: 0}
        for fi9 := 0; fi9 < 15; fi9++ {
            for fi10 := 0; fi10 < 14; fi10++ {
                if true {
                    v11 := GenS1{F0: 100}
                    fmt.Println(p1, p0)
                    _ = len(v7)
                    _ = v11
                }
                fmt.Println(p0, fi10, p0)
                var v12 float64 = 0.1
                _ = fi10
                _ = v12
            }
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 string = "\n\t"
                _ = fi13
                _ = v14
            }
            _ = fi9
        }
        _ = v8
    } else {
        var v15 uint = 167
        fmt.Println(p0, p0)
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v7
    return p1
}

func main() {
    v1 := []uint32{65535}
    var v2 uint = 4294967295
    var v3 float64 = (((float64((1e10 + (-2.25 + (-2.25 * (((0.0 - (1e10 * (0.0 - -2.25))) + 0.1) / 3.14159))))) - (1e10 - 0.1)) + 1e10) + 1.5)
    for fi16 := 0; fi16 < 15; fi16++ {
        v17 := GenS2{F0: v2, F1: "", F2: 1}
        if (v2 <= v2) {
            if true {
                v18 := GenS1{F0: -128}
                _ = v18
            }
            fmt.Println(v2, v3, v2)
            for fi19 := 0; fi19 < 7; fi19++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v3, fi19, v2)
                _ = len(v1)
                switch fi19 {
                case 1:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    var v20 uint16 = 683
                    _ = v20
                }
                if true {
                    v3 = float64(0.1)
                    v2 = (v2 + 42)
                }
                _ = fi19
            }
        } else {
            for fi21 := 0; fi21 < 10; fi21++ {
                v2 = (^(v2))
                _ = fi21
            }
            if (-2.25 < (v3 / v3)) {
                switch "" {
                case "a":
                    for fi22 := 0; fi22 < 19; fi22++ {
                        var v23 uint32 = 4294967295
                        _ = fi22
                        _ = v23
                    }
                }
                var v24 int32 = genFn1(604, 0)
                for fi25 := 0; fi25 < 13; fi25++ {
                    v26 := GenS2{F0: 65535, F1: ("\n\t" + ""), F2: 2}
                    _ = fi25
                    _ = v26
                }
                _ = v24
            } else {
                v2 = v2
                var v27 int = (fi16 * fi16)
                for fi28 := 0; fi28 < 9; fi28++ {
                    v3 = v3
                    var v29 uint16 = 81
                    _ = len(v1)
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
        }
        _ = fi16
        _ = v17
    }
    var v30 float32 = (0.0 / 3.14159)
    var v31 float32 = 0.1
    if ((!false) || true) {
        var v32 float32 = 1.5
        _ = v32
    } else {
        for fi33 := 0; fi33 < 4; fi33++ {
            for fi34 := 0; fi34 < 7; fi34++ {
                v2 += v2
                _ = fi34
            }
            _ = fi33
        }
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        var v36 uint64 = 18446744073709551615
        fmt.Println(v31, v2, v3)
        if true {
            switch "hello" {
            case "foo bar":
                if true {
                    var v37 uint8 = 1
                    v3 += float64(1e10)
                    v30 -= 3.14159
                    _ = v37
                } else {
                    if false {
                        var v38 uint16 = 294
                        switch "hello" {
                        case "":
                            if true {
                                v30 = float32(float64(0.1))
                                var v39 uint = (v2 / ((v2) | 1))
                                v31 = v30
                                _ = v39
                            }
                        case "hello":
                            v40 := GenS1{F0: -128}
                            _ = v40
                        case "a":
                            fmt.Println(v31, v31, v31)
                        }
                        fmt.Println(v38)
                        _ = v38
                    }
                    v41 := GenS3{F0: 18446744073709551615, F1: 0}
                    _ = v41
                }
            default:
                var v42 uint8 = 255
                _ = v42
            }
            v36 = (v36 % ((genFn0(63, v36, v36)) | 1))
        } else {
            if false {
                for fi43 := 0; fi43 < 18; fi43++ {
                    v44 := GenS1{F0: 42}
                    _ = fi43
                    _ = v44
                }
                v36 += (v36 | genFn0(127, v36, (^(v36))))
            }
            v2 -= (v2 * 65535)
        }
        _ = fi35
        _ = v36
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        var v45 uint8 = 32
        _ = v45
    }
    if false {
        var v46 float64 = v3
        if (0 > 42) {
            v31 += (v30 / 1.5)
            var v47 int16 = 0
            for fi48 := 0; fi48 < 20; fi48++ {
                v30 = float32((v3 - 1e10))
                _ = fi48
            }
            _ = v47
        } else {
            var v49 int64 = -1
            _ = v49
        }
        _ = v46
    } else {
        fmt.Println(v2, v30, v30)
        for fi50 := 0; fi50 < 1; fi50++ {
            v2 -= 18446744073709551615
            fmt.Println(v3, v31)
            fmt.Println(v31)
            _ = fi50
        }
    }
    var v51 bool = (!(127 <= 0))
    v3 -= float64(0.1)
    if (v2 <= v2) {
        for fi52 := 0; fi52 < 14; fi52++ {
            var v53 int = (-(fi52))
            v54 := GenS1{F0: -90}
            _ = fi52
            _ = v53
            _ = v54
        }
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if (!v51) {
        v30 -= float32(float64(v31))
        var v55 int32 = genFn1((v2 - (^(v2))), genFn1(v2, -1))
        if v51 {
            v1 = append(v1, 195)
        } else {
            v56 := GenS3{F0: (v2 - (v2 | v2)), F1: -854}
            _ = v56
        }
        _ = v55
    }
    for fi57 := 0; fi57 < 11; fi57++ {
        v51 = v51
        v58 := GenS4{F0: fi57}
        var v59 float32 = v30
        _ = fi57
        _ = v58
        _ = v59
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v51
}
