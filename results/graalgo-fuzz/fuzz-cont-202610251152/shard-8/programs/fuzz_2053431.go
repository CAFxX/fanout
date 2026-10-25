package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 string }
type GenS2 struct { F0 string; F1 int8; F2 int32 }
type GenS3 struct { F0 int }
func genFn0(p0 uint8, p1 string, p2 int8) uint32 {
    v7 := GenS1{F0: 65535, F1: 42, F2: ((p1 + p1) + "\n\t")}
    switch -534 {
    case 9223372036854775807:
        var v8 int = 32767
        _ = v8
    case 1:
        p0 += (^(p0))
    default:
        p0 += p0
    }
    v9 := GenS2{F0: "", F1: (p2 - p2), F2: -1}
    p1 = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return uint32(p0)
}
func genFn1(p0 int, p1 int16, p2 int64) uint16 {
    fmt.Println(p2, p0)
    p1 = 0
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]int)
    var v3 bool = (!((42 == 7) && (false || (!false))))
    var v4 uint16 = 860
    v5 := make(map[uint32]string)
    var v6 bool = v3
    v10 := GenS3{F0: 127}
    v11 := GenS3{F0: 128}
    switch 1 {
    case -454:
        var v12 int16 = 0
        _ = v12
    default:
        v13 := v2[("" + ((("a" + v1) + v1) + v1))]
        _ = v13
    }
    var v14 string = v1
    v15 := []int{}
    if false {
        var v16 string = ""
        v3 = (!false)
        _ = v16
    } else {
        fmt.Println(v14)
        v17 := v5[0]
        v18 := GenS1{F0: 617, F1: 42, F2: ((((v17 + (v14 + "foo bar")) + "hello") + v14) + "foo bar")}
        _ = v17
        _ = v18
    }
    v4 += 2
    switch 4294967295 {
    case -270:
        v1 = ("foo bar" + v14)
    case -9223372036854775808:
        var v19 int32 = 376
        _ = v19
    default:
        for fi20 := 0; fi20 < 11; fi20++ {
            for fi21 := 0; fi21 < 1; fi21++ {
                if v3 {
                    v4 -= 185
                } else {
                    v14 = (((((v14 + v14) + "foo bar") + "foo bar") + (v14 + (v14 + (("" + "\n\t") + (v1 + v1))))) + v14)
                }
                _ = fi21
            }
            for fi22 := 0; fi22 < 10; fi22++ {
                _ = len(v15)
                var v23 uint16 = 1
                v23 = v4
                _ = fi22
                _ = v23
            }
            if (128 >= 187) {
                v24 := GenS1{F0: 2147483647, F1: 0, F2: v1}
                for fi25 := 0; fi25 < 7; fi25++ {
                    v1 = v1
                    var v26 string = ("\n\t" + (v14 + v1))
                    _ = fi25
                    _ = v26
                }
                _ = v24
            } else {
                for fi27 := 0; fi27 < 13; fi27++ {
                    v28 := v5[2]
                    v14 = (v14 + (((v28 + ("foo bar" + v14)) + v28) + v14))
                    var v29 uint8 = 212
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
            _ = fi20
        }
    }
    var v30 int8 = 1
    v14 = ((v1 + ((v1 + v14) + v1)) + v1)
    if (42 < 42) {
        switch 561 {
        case 1:
            if (32767 < 359) {
                var v31 uint8 = 1
                v32 := GenS2{F0: "", F1: (v30 >> ((819) & 63)), F2: 1}
                v1 = v1
                _ = v31
                _ = v32
            }
        case 84:
            v33 := GenS2{F0: "foo bar", F1: (^(v30)), F2: 42}
            _ = v33
        case -659:
            for fi34 := 0; fi34 < 9; fi34++ {
                var v35 float32 = 1e10
                fmt.Println(v3, fi34, v4)
                _ = fi34
                _ = v35
            }
        default:
            var v36 int16 = 817
            _ = v36
        }
        for fi37 := 0; fi37 < 1; fi37++ {
            v6 = v3
            _ = fi37
        }
    } else {
        var v38 string = ((("\n\t" + v1) + (("" + "a") + "foo bar")) + v14)
        fmt.Println(v4, v4, v38)
        for fi39 := 0; fi39 < 7; fi39++ {
            fmt.Println(v3, v14)
            v14 = v1
            _ = fi39
        }
        _ = v38
    }
    var v40 bool = (128 != 91)
    for fi41 := 0; fi41 < 11; fi41++ {
        v5[2147483647] = "\n\t"
        var v42 uint = 255
        _ = fi41
        _ = v42
    }
    var v43 int = -9223372036854775808
    var v44 int16 = 1000
    switch (v43 - -973) {
    case -1:
        var v45 uint16 = (v4 & v4)
        _ = v45
    case 814:
        v2[("hello" + v14)] = v43
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
