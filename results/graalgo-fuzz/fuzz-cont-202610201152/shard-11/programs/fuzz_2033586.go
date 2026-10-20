package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int }
type GenS2 struct { F0 bool; F1 int32; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int8; F2 int8 }
func genFn0(p0 uint32, p1 int16) uint32 {
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 uint32 = p0
        v7 := GenS1{F0: 32767, F1: (p1 * (p1 ^ (p1 | 42))), F2: fi5}
        _ = fi5
        _ = v6
        _ = v7
    }
    switch "\n\t" {
    case "\n\t":
        v8 := []int{}
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: 0, F1: (p1 | p1), F2: -9223372036854775808}
            _ = fi9
            _ = v10
        }
    default:
        p0 = 484
    }
    v11 := GenS1{F0: 514, F1: (p1 | p1), F2: 127}
    _ = p0
    _ = p1
    _ = v11
    return 1
}
func genFn1() int16 {
    var v12 int64 = -1
    switch "foo bar" {
    case "\n\t":
        v13 := []int{4294967295, 4294967295, 128, -9223372036854775808}
        _ = v13
    case "hello":
        for fi14 := 0; fi14 < 10; fi14++ {
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := GenS1{F0: 42, F1: -383, F2: (fi14 % ((fi14) | 1))}
                fmt.Println(v12, fi15)
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    var v17 uint32 = 2147483648
    fmt.Println(v17, v17)
    _ = v12
    _ = v17
    return 976
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    var v4 uint16 = 935
    v18 := GenS1{F0: v4, F1: 42, F2: 2147483647}
    var v19 int32 = 0
    var v20 string = "foo bar"
    var v21 uint16 = v4
    v19 = (v19 >> (((v1 / ((v1) | 1))) & 63))
    v21 += 32767
    v22 := GenS1{F0: (v4 & v21), F1: 0, F2: 140}
    for fi23 := 0; fi23 < 17; fi23++ {
        for fi24 := 0; fi24 < 4; fi24++ {
            if true {
                var v25 int8 = 0
                fmt.Println(fi24, v4, fi23)
                var v26 uint = v1
                _ = v25
                _ = v26
            } else {
                if true {
                    fmt.Println(v4)
                    switch int(v1) {
                    case 349:
                        v27 := GenS2{F0: (2147483648 > 159), F1: (-(v19)), F2: -32768}
                        _ = v27
                    case 128:
                        v28 := GenS2{F0: false, F1: (-(v19)), F2: -795}
                        _ = v28
                    case -956:
                        switch (fi24 * fi23) {
                        case -1:
                            if ((genFn0(0, -1) < 4294967295) && ((^(fi24)) == (fi23 | (fi23 % (((fi24 % (((fi24 + fi24)) | 1))) | 1))))) {
                                var v29 int16 = 42
                                _ = v29
                            }
                        }
                    }
                }
                var v30 float32 = (((0.1 + (1.5 - 1.5)) + (1.5 * float32(1.5))) + -2.25)
                fmt.Println(v20)
                _ = v30
            }
            switch (fi23 / ((-469) | 1)) {
            case 32767:
                if false {
                    v31 := []int{int(v21), fi24, (fi23 << ((0) & 63)), 32768}
                    var v32 float64 = -2.25
                    _ = v31
                    _ = v32
                } else {
                    var v33 int = fi24
                    fmt.Println(v20, v4)
                    _ = v33
                }
            case -1:
                v3[fi23] = (("\n\t" + (v20 + ("foo bar" + v20))) + ("\n\t" + v20))
            case 128:
                v34 := GenS2{F0: (0 < 113), F1: 1, F2: 933}
                _ = v34
            default:
                var v35 uint = (v1 + (v1 << ((42) & 63)))
                _ = v35
            }
            _ = fi24
        }
        _ = fi23
    }
    switch 9223372036854775807 {
    case -1:
        if false {
            v36 := GenS3{F0: -474}
            for fi37 := 0; fi37 < 18; fi37++ {
                for fi38 := 0; fi38 < 5; fi38++ {
                    var v39 int8 = -128
                    var v40 int32 = v19
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                if false {
                    if (!true) {
                        fmt.Println(v19, v4)
                        fmt.Println(v1, v4)
                    }
                    if (2147483648 >= genFn0(1, 1000)) {
                        fmt.Println(v1, fi37, v1)
                    } else {
                        for fi41 := 0; fi41 < 8; fi41++ {
                            var v42 uint64 = 4294967295
                            _ = fi41
                            _ = v42
                        }
                    }
                    fmt.Println(fi37, fi37)
                } else {
                    v1 += v1
                    var v43 uint64 = 4294967295
                    for fi44 := 0; fi44 < 8; fi44++ {
                        switch fi37 {
                        case 127:
                            fmt.Println(v21)
                        case 4294967295:
                            fmt.Println(v1, v21, v4)
                            if ((v1 / (((v1 >> ((v1) & 63))) | 1)) > 65535) {
                                v45 := []int{fi44}
                                v46 := GenS3{F0: 426}
                                v47 := GenS1{F0: (v21 ^ (v21 << ((v1) & 63))), F1: 42, F2: fi37}
                                var v48 uint64 = v43
                                _ = v45
                                _ = v46
                                _ = v47
                                _ = v48
                            } else {
                                v1 += (v1 | (v1 | (v1 % ((0) | 1))))
                            }
                            fmt.Println(v20, v19)
                        default:
                            var v49 int16 = genFn1()
                            _ = v49
                        }
                        _ = fi44
                    }
                    _ = v43
                }
                _ = fi37
            }
            var v50 int = 1
            _ = v36
            _ = v50
        }
    case 2147483647:
        var v51 float64 = 1e10
        _ = v51
    default:
        v4 += 757
    }
    v4 = v21
    fmt.Println(v4, v4, v4)
    fmt.Println(v1, v20, v21)
    switch v20 {
    case "":
        v52 := GenS4{F0: -623, F1: 42, F2: 40}
        _ = v52
    case "\n\t":
        var v53 int8 = -47
        _ = v53
    default:
        for fi54 := 0; fi54 < 3; fi54++ {
            var v55 int = 32767
            var v56 float32 = float32(float64(0.1))
            v19 += v19
            _ = fi54
            _ = v55
            _ = v56
        }
    }
    if true {
        for fi57 := 0; fi57 < 10; fi57++ {
            v2["\n\t"] = (-1 <= 1)
            _ = fi57
        }
        var v58 uint = (v1 ^ (v1 & (v1 << ((v1) & 63))))
        _ = v58
    } else {
        v59 := GenS2{F0: false, F1: v19, F2: 42}
        _ = v59
    }
    fmt.Println(v4, v20)
    var v60 uint16 = 250
    var v61 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v60
    _ = v61
}
