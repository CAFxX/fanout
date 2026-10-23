package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint16; F1 string; F2 int64 }
type GenS3 struct { F0 int16 }
func genFn0() int {
    if (0 != -9223372036854775808) {
        var v7 int64 = 211
        var v8 string = "hello"
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: -9223372036854775808, F2: 128}
    if true {
        v10 := GenS2{F0: 866, F1: ("" + "hello"), F2: 1}
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 int64 = 0
            var v13 uint = 4294967295
            v13 += v13
            _ = fi11
            _ = v12
            _ = v13
        }
        if (566 > 0) {
            v14 := GenS2{F0: 42, F1: "hello", F2: 1}
            var v15 bool = false
            v16 := GenS2{F0: 0, F1: ("hello" + ""), F2: -63}
            _ = v14
            _ = v15
            _ = v16
        }
        _ = v10
    }
    _ = v9
    return 2147483647
}
func genFn1() int {
    switch "\n\t" {
    case "a":
        v17 := []int{-1, 446}
        _ = v17
    case "hello":
        var v18 int16 = 243
        _ = v18
    default:
        var v19 uint16 = 65535
        _ = v19
    }
    switch "a" {
    case "a":
        v20 := GenS2{F0: 255, F1: "hello", F2: -591}
        _ = v20
    }
    return 32767
}

func main() {
    var v1 uint16 = 0
    var v2 uint32 = 65535
    var v3 int = 22
    v4 := make(map[string]int)
    var v5 uint16 = (v1 >> ((2) & 63))
    var v6 uint64 = 2
    var v21 bool = true
    v22 := []int{(v3 | 1), genFn0(), (^(v3)), 909}
    fmt.Println(v3)
    if true {
        var v23 uint = 248
        _ = v23
    } else {
        for fi24 := 0; fi24 < 14; fi24++ {
            if v21 {
                var v25 bool = true
                _ = v25
            }
            v26 := v4[""]
            if v21 {
                if v21 {
                    fmt.Println(v6, v5, v5)
                }
            } else {
                v5 = (^(v1))
                if len(v22) > 0 {
                    v22[0] = (-(v26))
                }
                v2 += v2
            }
            _ = fi24
            _ = v26
        }
        v3 = v3
        v22 = append(v22, (v3 >> ((0) & 63)))
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := v4["\n\t"]
        _ = fi27
        _ = v28
    }
    var v29 int16 = int16(v3)
    if (-52 == 118) {
        if (0 == 602) {
            for fi30 := 0; fi30 < 19; fi30++ {
                v4[("hello" + ("foo bar" + ""))] = (v3 >> ((956) & 63))
                v3 += (fi30 | -171)
                v31 := GenS3{F0: 697}
                _ = fi30
                _ = v31
            }
            if (0.1 <= -2.25) {
                var v32 float32 = 3.14159
                _ = v32
            } else {
                _ = len(v4)
            }
            _ = len(v22)
        }
        v3 = v3
        v22 = append(v22, v3)
    }
    if v21 {
        for fi33 := 0; fi33 < 12; fi33++ {
            fmt.Println(v21, v5)
            switch fi33 {
            case -506:
                fmt.Println(v3)
            case -9223372036854775808:
                var v34 float32 = 1e10
                _ = v34
            case -1:
                switch (v3 % ((int(v1)) | 1)) {
                case -101:
                    var v35 int16 = -722
                    _ = v35
                case -1:
                    v3 -= (fi33 & (fi33 * genFn0()))
                case 5:
                    switch v3 {
                    case -849:
                        delete(v4, "hello")
                    case 354:
                        if v21 {
                            if len(v22) > 0 {
                                v22[0] = (fi33 * fi33)
                            }
                            fmt.Println(v21, v5, v5)
                        }
                    default:
                        var v36 uint16 = 170
                        _ = v36
                    }
                }
            default:
                var v37 int8 = int8(v1)
                _ = v37
            }
            v4["foo bar"] = 4294967295
            _ = fi33
        }
    }
    if v21 {
        v38 := GenS2{F0: 255, F1: "", F2: -9223372036854775808}
        for fi39 := 0; fi39 < 18; fi39++ {
            if (128 != 0) {
                fmt.Println(fi39, v21)
            }
            fmt.Println(v21, v29)
            v40 := GenS1{F0: 73, F1: -908, F2: 128}
            _ = fi39
            _ = v40
        }
        v41 := GenS3{F0: -189}
        _ = v38
        _ = v41
    } else {
        v5 = 1
    }
    var v42 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
