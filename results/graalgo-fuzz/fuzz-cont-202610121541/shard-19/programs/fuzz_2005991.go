package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint16; F2 bool }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint8, p1 string) string {
    v4 := GenS1{F0: 745}
    for fi5 := 0; fi5 < 15; fi5++ {
        p0 += p0
        if false {
            v6 := []int{(-(fi5)), fi5}
            for fi7 := 0; fi7 < 17; fi7++ {
                switch (p1 + p1) {
                case "a":
                    var v8 int8 = 18
                    _ = v8
                default:
                    p1 = "a"
                }
                _ = fi7
            }
            _ = v6
        }
        _ = fi5
    }
    v9 := []int{}
    if true {
        if false {
            v10 := GenS1{F0: 2}
            _ = v10
        }
    } else {
        p0 -= p0
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v9
    return (("\n\t" + p1) + ("\n\t" + ("hello" + ("foo bar" + p1))))
}
func genFn1(p0 uint, p1 int16, p2 uint8) int {
    v11 := GenS2{F0: (p1 & 0)}
    v12 := GenS2{F0: (-(p1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -817
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{-2147483648, 32768}
    v3 := []string{"foo bar", ""}
    v13 := GenS3{F0: -297, F1: 255, F2: (2147483648 >= 2)}
    if (false && (((1 != 0) || true) && (929 <= 0))) {
        v14 := v1["hello"]
        _ = len(v1)
        _ = v14
    } else {
        var v15 bool = false
        if len(v2) > 0 {
            v2[0] = genFn1(0, -651, 0)
        }
        _ = v15
    }
    v16 := GenS4{F0: 42, F1: 127}
    var v17 bool = ((true && (9223372036854775807 >= 32767)) || false)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v18 float64 = 1.5
    if (126 < 100) {
        var v19 uint8 = 0
        fmt.Println(v18, v18, v17)
        _ = v19
    }
    v18 = (float64(float32(v18)) / 1e10)
    v17 = v17
    var v20 uint = 65535
    if false {
        v21 := GenS1{F0: 0}
        if v17 {
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 uint64 = 18446744073709551615
                if len(v2) > 0 {
                    v2[0] = -406
                }
                var v24 string = "foo bar"
                _ = fi22
                _ = v23
                _ = v24
            }
            var v25 uint16 = 2
            var v26 int = 0
            _ = v25
            _ = v26
        }
        if false {
            var v27 int8 = 42
            v28 := GenS4{F0: -9223372036854775808, F1: -1}
            _ = v27
            _ = v28
        } else {
            for fi29 := 0; fi29 < 14; fi29++ {
                v20 -= v20
                for fi30 := 0; fi30 < 19; fi30++ {
                    var v31 uint32 = 0
                    _ = fi30
                    _ = v31
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi29
            }
            var v32 int32 = 169
            _ = v32
        }
        _ = v21
    } else {
        fmt.Println(v20, v20, v17)
        v33 := GenS1{F0: 65535}
        fmt.Println(v18)
        _ = v33
    }
    v18 -= (v18 + v18)
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 float32 = 0.0
        v17 = (!(v17 && false))
        _ = fi34
        _ = v35
    }
    switch 0 {
    case 9223372036854775807:
        v17 = (18446744073709551615 <= 42)
    case 699:
        for fi36 := 0; fi36 < 3; fi36++ {
            switch fi36 {
            case -567:
                var v37 int8 = 1
                _ = v37
            }
            fmt.Println(v17)
            var v38 int = 3
            _ = fi36
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        delete(v1, "")
        _ = fi39
    }
    switch ("\n\t" + "foo bar") {
    case "":
        var v40 bool = v17
        _ = v40
    default:
        var v41 float64 = ((0.1 * v18) - 3.14159)
        _ = v41
    }
    v18 += v18
    if true {
        fmt.Println(v17)
        if v17 {
            for fi42 := 0; fi42 < 15; fi42++ {
                v43 := GenS2{F0: -1}
                var v44 int64 = -1
                _ = fi42
                _ = v43
                _ = v44
            }
        } else {
            for fi45 := 0; fi45 < 19; fi45++ {
                var v46 string = "hello"
                switch fi45 {
                case -52:
                    if v17 {
                        v46 = "a"
                    }
                case 32768:
                    v47 := GenS5{F0: 65535, F1: 1}
                    _ = v47
                default:
                    for fi48 := 0; fi48 < 14; fi48++ {
                        v17 = (1 != -686)
                        v46 = ("" + v46)
                        var v49 uint16 = 666
                        _ = fi48
                        _ = v49
                    }
                }
                fmt.Println(v17)
                _ = fi45
                _ = v46
            }
            var v50 uint8 = 109
            v20 = 667
            _ = v50
        }
        v51 := GenS4{F0: -1, F1: 360}
        _ = v51
    } else {
        v20 = (v20 | (v20 * (v20 - 0)))
        var v52 uint = (v20 / ((1) | 1))
        for fi53 := 0; fi53 < 19; fi53++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi53
        }
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
}
