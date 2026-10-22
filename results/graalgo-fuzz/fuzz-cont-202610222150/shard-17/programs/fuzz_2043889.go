package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 string; F1 int8; F2 int }
type GenS3 struct { F0 uint8; F1 string; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint16, p1 string) uint64 {
    for fi7 := 0; fi7 < 9; fi7++ {
        p1 = ((p1 + p1) + p1)
        for fi8 := 0; fi8 < 10; fi8++ {
            if (0 == 0) {
                p0 = (p0 | (p0 ^ p0))
                p1 = (p1 + p1)
            } else {
                v9 := []int{fi7, fi8}
                var v10 int64 = 1
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                if false {
                    var v11 int8 = -128
                    if (0.0 != 1.5) {
                        var v12 uint8 = 42
                        switch p1 {
                        case "\n\t":
                            var v13 string = ("hello" + p1)
                            _ = v13
                        default:
                            p1 = (p1 + p1)
                        }
                        var v14 bool = (4294967295 < 2147483647)
                        _ = v12
                        _ = v14
                    }
                    _ = v11
                }
                v10 += v10
                var v15 uint32 = 65535
                _ = v10
                _ = v15
                _ = v9
            }
            p1 = ("\n\t" + (((p1 + "hello") + "foo bar") + "a"))
            p1 = "a"
            _ = fi8
        }
        _ = fi7
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        p1 = ("\n\t" + (p1 + p1))
        var v17 uint = 51
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 32768, F1: 1}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v18
    return 4294967295
}

func main() {
    var v1 int64 = 42
    var v2 int8 = 42
    var v3 int8 = v2
    var v4 int32 = 896
    var v5 bool = ((v4 / ((v4) | 1)) != v4)
    var v6 uint32 = 675
    var v19 bool = true
    var v20 uint = 255
    v21 := []int{9223372036854775807, 32767}
    if true {
        for fi22 := 0; fi22 < 8; fi22++ {
            v3 = (-(v2))
            for fi23 := 0; fi23 < 3; fi23++ {
                var v24 int16 = 42
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        v25 := GenS1{F0: 339, F1: 42}
        v5 = v19
        _ = v25
    } else {
        switch 0 {
        case -9223372036854775808:
            switch "hello" {
            case "\n\t":
                if v19 {
                    v19 = true
                    var v26 uint32 = (^(v6))
                    for fi27 := 0; fi27 < 18; fi27++ {
                        for fi28 := 0; fi28 < 9; fi28++ {
                            var v29 float32 = float32(((float64((1.5 / (-2.25 / 0.1))) - 0.0) / (float64(1e10) + float64(0.1))))
                            _ = fi28
                            _ = v29
                        }
                        v30 := GenS2{F0: ("foo bar" + (((("foo bar" + ("a" + "foo bar")) + "") + (("\n\t" + "hello") + (("" + "") + "hello"))) + "foo bar")), F1: v2, F2: (fi27 | (fi27 >> ((v20) & 63)))}
                        _ = fi27
                        _ = v30
                    }
                    _ = v26
                }
            case "":
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
            }
        case 32768:
            var v31 int = 2147483647
            _ = v31
        default:
            v32 := GenS3{F0: 2, F1: "hello", F2: -475}
            _ = v32
        }
        if (false || (!v5)) {
            var v33 uint64 = 0
            v19 = v19
            _ = v33
        }
    }
    switch 127 {
    case 853:
        v5 = true
    case 42:
        v34 := GenS2{F0: "a", F1: 0, F2: -790}
        _ = v34
    default:
        if v19 {
            for fi35 := 0; fi35 < 13; fi35++ {
                var v36 int8 = (v3 + 100)
                _ = len(v21)
                _ = fi35
                _ = v36
            }
            var v37 uint32 = 42
            var v38 uint16 = 1
            _ = v37
            _ = v38
        } else {
            fmt.Println(v4, v5, v5)
        }
    }
    v39 := GenS4{F0: 168, F1: 255, F2: 675}
    var v40 int64 = v1
    v41 := GenS5{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
