package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint16 }
func genFn0() uint16 {
    var v5 float32 = 1.5
    v6 := []int{9223372036854775807, -9223372036854775808, 9223372036854775807}
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 bool, p1 int) int {
    p0 = (!false)
    v7 := []int{-2147483648}
    if (((0.1 * 0.0) <= 1e10) && false) {
        v8 := GenS1{F0: 4294967295}
        if true {
            for fi9 := 0; fi9 < 14; fi9++ {
                for fi10 := 0; fi10 < 2; fi10++ {
                    if p0 {
                        var v11 int64 = 104
                        _ = v11
                    } else {
                        fmt.Println(p0, p1)
                        fmt.Println(fi10, fi9, fi9)
                        switch 32767 {
                        case 557:
                            p1 = fi9
                        }
                    }
                    p1 -= p1
                    _ = fi10
                }
                _ = fi9
            }
        }
        _ = v8
    } else {
        var v12 uint32 = 1
        _ = v12
    }
    var v13 uint16 = 744
    p1 += (p1 - p1)
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return int(v13)
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -420
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    if false {
        v14 := []int{9223372036854775807}
        _ = v14
    }
    v2 -= (v2 * 1)
    v15 := GenS2{F0: 45, F1: 412}
    v2 += (v2 - -9223372036854775808)
    if (!false) {
        for fi16 := 0; fi16 < 18; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                v2 = v2
                v2 -= v2
                v2 -= (-(v2))
                var v18 string = "foo bar"
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 uint = 18446744073709551615
                _ = fi19
                _ = v20
            }
            if false {
                fmt.Println(v2, fi16)
                fmt.Println(fi16, fi16)
            }
            _ = fi16
        }
    } else {
        v21 := GenS1{F0: 4294967295}
        v22 := []int{32768, -2147483648}
        v2 = -1
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        switch fi23 {
        case 32767:
            fmt.Println(v2)
        }
        _ = fi23
    }
    if (false && false) {
        var v24 int = 127
        v24 = 32768
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        if false {
            fmt.Println(v2)
            if ((v2 << ((281) & 63)) == -434) {
                if false {
                    v2 -= (v2 / (((v2 + v2)) | 1))
                    var v26 string = "foo bar"
                    for fi27 := 0; fi27 < 12; fi27++ {
                        v28 := GenS3{F0: (!false), F1: 127, F2: genFn0()}
                        _ = fi27
                        _ = v28
                    }
                    _ = v26
                } else {
                    fmt.Println(v2, fi25)
                    for fi29 := 0; fi29 < 3; fi29++ {
                        var v30 int = 9223372036854775807
                        v30 -= v30
                        var v31 uint32 = 4294967295
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                }
                v32 := GenS2{F0: 2, F1: 2147483647}
                _ = v32
            }
            var v33 uint = 18446744073709551615
            _ = v33
        }
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
