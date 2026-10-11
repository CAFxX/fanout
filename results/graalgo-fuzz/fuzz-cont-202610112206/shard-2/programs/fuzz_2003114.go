package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint64; F1 bool; F2 uint }
type GenS7 struct { F0 bool }
func genFn0(p0 bool) uint64 {
    var v6 float32 = ((1.5 + -2.25) * -2.25)
    for fi7 := 0; fi7 < 3; fi7++ {
        if (127 == 29) {
            if p0 {
                v8 := GenS1{F0: 1, F1: 935, F2: 0}
                var v9 int8 = 89
                if (4294967295 == 4294967295) {
                    p0 = p0
                    var v10 float32 = v6
                    _ = v10
                }
                _ = v8
                _ = v9
            } else {
                if (1.5 != 1.5) {
                    fmt.Println(p0)
                }
                v11 := GenS1{F0: 0, F1: 380, F2: 284}
                var v12 float64 = 0.0
                _ = v11
                _ = v12
            }
            v13 := []int{fi7, (fi7 * (fi7 - 1))}
            _ = v13
        }
        p0 = (!(p0 || (464 <= -1)))
        _ = fi7
    }
    v6 += -2.25
    v14 := []int{32768, -2147483648, 32768, -338}
    var v15 uint16 = 32767
    _ = p0
    _ = v14
    _ = v15
    _ = v6
    return uint64(v15)
}
func genFn1(p0 int8, p1 string) int32 {
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(p0, fi16)
        _ = fi16
    }
    var v17 uint = 42
    if (2 < 128) {
        fmt.Println(p0, p1)
        v18 := GenS1{F0: 42, F1: 69, F2: 4294967295}
        _ = v18
    }
    p1 = ("\n\t" + "foo bar")
    _ = p0
    _ = p1
    _ = v17
    return 1
}

func main() {
    var v1 int64 = 0
    var v2 uint8 = 0
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    fmt.Println(v2, v1)
    v19 := GenS2{F0: 42, F1: ("" + ("hello" + ("hello" + ("a" + "\n\t"))))}
    var v20 uint64 = 4294967295
    var v21 float64 = 0.1
    var v22 float32 = ((float32(v21) * -2.25) * 0.1)
    fmt.Println(v20)
    v22 -= v22
    var v23 float64 = (float64(v22) / v21)
    v21 += (v21 - v23)
    if (42 <= 1) {
        if false {
            _ = len(v4)
            var v24 uint32 = 65535
            _ = v24
        }
        switch ("\n\t" + "") {
        case "foo bar":
            v25 := GenS3{F0: -944, F1: 4294967295}
            _ = v25
        case "":
            var v26 int64 = (v1 % ((42) | 1))
            _ = v26
        case "\n\t":
            if false {
                for fi27 := 0; fi27 < 7; fi27++ {
                    for fi28 := 0; fi28 < 6; fi28++ {
                        var v29 uint8 = (v2 ^ v2)
                        v30 := GenS4{F0: (v1 >> ((470) & 63))}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 2; fi31++ {
                        fmt.Println(v21, fi27, v1)
                        v32 := v4[(("hello" + "\n\t") + "hello")]
                        _ = fi31
                        _ = v32
                    }
                    _ = fi27
                }
                var v33 int8 = -114
                v20 -= genFn0(false)
                _ = v33
            }
        }
    } else {
        v20 -= (^(v20))
        v34 := GenS5{F0: 65535, F1: 500}
        for fi35 := 0; fi35 < 11; fi35++ {
            _ = len(v3)
            if false {
                for fi36 := 0; fi36 < 20; fi36++ {
                    v37 := []int{}
                    switch (fi35 ^ fi35) {
                    case 4294967295:
                        if len(v37) > 0 {
                            v37[0] = (fi35 + (^(fi35)))
                        }
                    case -9223372036854775808:
                        var v38 float64 = (v21 - -2.25)
                        _ = v38
                    default:
                        v39 := GenS6{F0: (^(v20)), F1: (v22 != -2.25), F2: 1}
                        _ = v39
                    }
                    _ = fi36
                    _ = v37
                }
                var v40 uint8 = 255
                _ = v40
            } else {
                if false {
                    var v41 bool = (!false)
                    _ = v41
                } else {
                    var v42 int64 = v1
                    var v43 int16 = 1
                    _ = v42
                    _ = v43
                }
            }
            if (false || (false || (true && ((v2 / ((v2) | 1)) >= v2)))) {
                v20 = 42
                var v44 int = 2147483647
                var v45 int64 = -9223372036854775808
                _ = v44
                _ = v45
            }
            _ = fi35
        }
        _ = v34
    }
    v22 -= (v22 / float32(3.14159))
    v23 = 3.14159
    var v46 int32 = 42
    var v47 float32 = (float32(v21) * (v22 - (-2.25 * v22)))
    v47 += v47
    v48 := GenS4{F0: v1}
    v49 := GenS7{F0: false}
    for fi50 := 0; fi50 < 6; fi50++ {
        v1 -= (^(v1))
        fmt.Println(v20, v46, v1)
        v51 := GenS7{F0: true}
        _ = fi50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
}
