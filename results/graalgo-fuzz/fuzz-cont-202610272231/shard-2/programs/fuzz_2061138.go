package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint16 }
type GenS2 struct { F0 string; F1 string; F2 int32 }
type GenS3 struct { F0 int8; F1 int; F2 int }
func genFn0(p0 int16, p1 uint64, p2 uint) string {
    for fi5 := 0; fi5 < 14; fi5++ {
        if true {
            v6 := []int{fi5, 1, fi5}
            var v7 float32 = float32(0.1)
            if false {
                for fi8 := 0; fi8 < 11; fi8++ {
                    var v9 uint = 42
                    _ = fi8
                    _ = v9
                }
            } else {
                v6 = append(v6, 32768)
                v7 = (v7 * v7)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = v6
            _ = v7
        }
        p1 = p1
        v10 := []int{127, 1, fi5, fi5}
        _ = fi5
        _ = v10
    }
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return ("a" + "")
}
func genFn1(p0 uint, p1 int, p2 int64) bool {
    if false {
        if (-1 <= 1) {
            var v11 uint32 = 65535
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 uint = p0
                p2 -= (p2 ^ p2)
                _ = fi12
                _ = v13
            }
            v14 := GenS1{F0: -106, F1: true, F2: 65535}
            _ = v11
            _ = v14
        } else {
            var v15 bool = true
            for fi16 := 0; fi16 < 10; fi16++ {
                fmt.Println(p0, p0)
                v17 := []int{fi16, 128, int(p0)}
                _ = fi16
                _ = v17
            }
            v15 = false
            _ = v15
        }
    } else {
        p2 = -23
        v18 := []int{}
        var v19 uint32 = 40
        _ = v18
        _ = v19
    }
    if true {
        var v20 int32 = 42
        v20 = (^(v20))
        _ = v20
    } else {
        for fi21 := 0; fi21 < 6; fi21++ {
            if false {
                p1 += (p1 << ((p0) & 63))
                p1 -= p1
            } else {
                switch 9223372036854775807 {
                case -2147483648:
                    if false {
                        var v22 float32 = 0.1
                        _ = v22
                    }
                case -9223372036854775808:
                    for fi23 := 0; fi23 < 18; fi23++ {
                        p0 += (p0 + (p0 * (p0 * p0)))
                        switch "a" {
                        case "hello":
                            v24 := []int{32767}
                            _ = v24
                        case "\n\t":
                            if (42 >= 1) {
                                p0 = p0
                                fmt.Println(p1, p1)
                                var v25 int8 = -1
                                var v26 float64 = ((0.0 / 3.14159) * float64(((1e10 + 1.5) * 0.1)))
                                v27 := []int{fi21}
                                _ = v25
                                _ = v26
                                _ = v27
                            } else {
                                p1 = fi23
                                var v28 string = ("hello" + "foo bar")
                                v29 := GenS1{F0: 127, F1: false, F2: 42}
                                _ = v28
                                _ = v29
                            }
                        case "foo bar":
                            if true {
                                p1 -= (fi23 | (p1 / (((p1 - fi23)) | 1)))
                            }
                        }
                        _ = fi23
                    }
                case 1:
                    switch -1 {
                    case 2147483647:
                        for fi30 := 0; fi30 < 19; fi30++ {
                            p1 += 32767
                            var v31 uint64 = 126
                            fmt.Println(p1)
                            fmt.Println(fi30, fi30)
                            v32 := GenS1{F0: 0, F1: (0 <= 127), F2: 2}
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    default:
                        p2 = (p2 ^ p2)
                    }
                }
                p0 = (p0 & p0)
            }
            _ = fi21
        }
    }
    v33 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    return true
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 42
    var v3 int32 = 42
    var v4 int32 = 865
    v3 = 0
    delete(v1, ((("a" + "foo bar") + "\n\t") + "hello"))
    switch -2147483648 {
    case 1:
        v34 := GenS1{F0: 1, F1: (!false), F2: 42}
        _ = v34
    case 128:
        v2 += (v2 >> ((116) & 63))
    case -443:
        for fi35 := 0; fi35 < 12; fi35++ {
            v36 := GenS2{F0: ("hello" + ("a" + ("" + ("hello" + ("foo bar" + "hello"))))), F1: ((("a" + "") + ("a" + ("hello" + "hello"))) + ""), F2: 421}
            _ = fi35
            _ = v36
        }
    default:
        for fi37 := 0; fi37 < 14; fi37++ {
            fmt.Println(fi37, v4, v4)
            fmt.Println(v4, v3, v2)
            v2 += v2
            _ = fi37
        }
    }
    if true {
        v38 := v1["foo bar"]
        _ = v38
    } else {
        for fi39 := 0; fi39 < 9; fi39++ {
            switch -9223372036854775808 {
            case 1:
                v40 := GenS2{F0: "a", F1: "\n\t", F2: (v4 ^ v4)}
                _ = v40
            }
            _ = fi39
        }
    }
    v41 := GenS3{F0: 1, F1: -474, F2: 32767}
    v3 -= -1
    fmt.Println(v2, v2)
    fmt.Println(v2, v4)
    var v42 float32 = 0.0
    var v43 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
