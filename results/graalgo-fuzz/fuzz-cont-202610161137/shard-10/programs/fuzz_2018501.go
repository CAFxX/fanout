package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint32; F1 uint; F2 uint }
type GenS4 struct { F0 int16; F1 bool; F2 uint16 }
func genFn0(p0 uint8) bool {
    v4 := GenS1{F0: 42, F1: 42}
    if (214 >= 1) {
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := GenS1{F0: 127, F1: -9223372036854775808}
            if (!(int16(fi5) < -1)) {
                p0 = (p0 << ((802) & 63))
                var v7 int16 = 1000
                _ = v7
            } else {
                switch (^(fi5)) {
                case 692:
                    fmt.Println(fi5, p0)
                default:
                    v8 := []int{}
                    _ = v8
                }
                if (((!false) && (!(!(-2147483648 > 2147483647)))) || true) {
                    var v9 uint16 = 42
                    for fi10 := 0; fi10 < 4; fi10++ {
                        v11 := GenS1{F0: 74, F1: -1}
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                } else {
                    p0 = (p0 + p0)
                }
            }
            var v12 int32 = 0
            _ = fi5
            _ = v12
            _ = v6
        }
        p0 = (p0 ^ (p0 << ((984) & 63)))
    } else {
        var v13 int = 1
        _ = v13
    }
    var v14 uint64 = 1
    fmt.Println(p0)
    _ = p0
    _ = v14
    _ = v4
    return (true || (-2147483648 != -731))
}
func genFn1(p0 uint16) uint64 {
    v15 := GenS1{F0: -95, F1: -635}
    v16 := []int{-634}
    fmt.Println(p0)
    switch -2147483648 {
    case -9223372036854775808:
        fmt.Println(p0)
    case 917:
        v16 = append(v16, 2147483647)
    case 4294967295:
        if ((^(p0)) < p0) {
            fmt.Println(p0)
        } else {
            for fi17 := 0; fi17 < 14; fi17++ {
                p0 = p0
                _ = fi17
            }
        }
    }
    v18 := GenS1{F0: 3, F1: 1}
    _ = p0
    _ = v15
    _ = v16
    _ = v18
    return 2
}

func main() {
    var v1 int64 = -792
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    v1 = v1
    if (((1e10 + (3.14159 + 1.5)) - ((-2.25 * 1.5) + float64(float32(-2.25)))) >= float64((float32((-2.25 * 1e10)) + -2.25))) {
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := GenS1{F0: -1, F1: (^(v1))}
            if (!false) {
                for fi21 := 0; fi21 < 10; fi21++ {
                    var v22 uint16 = 0
                    if (127 >= 76) {
                        v22 += 0
                        v23 := GenS2{F0: 42, F1: -1}
                        _ = v23
                    }
                    _ = fi21
                    _ = v22
                }
            }
            var v24 int16 = 1000
            _ = fi19
            _ = v20
            _ = v24
        }
        v25 := GenS2{F0: 0, F1: 32768}
        _ = v25
    } else {
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS2{F0: genFn1(1), F1: fi26}
            _ = fi26
            _ = v27
        }
        v1 = (v1 & 1)
    }
    var v28 float64 = ((3.14159 + 0.1) - 3.14159)
    var v29 uint8 = 0
    var v30 uint32 = 265
    if true {
        if true {
            v31 := GenS2{F0: 0, F1: 32768}
            _ = v31
        }
    } else {
        var v32 int16 = -652
        for fi33 := 0; fi33 < 14; fi33++ {
            for fi34 := 0; fi34 < 6; fi34++ {
                v28 += v28
                if (-873 >= 42) {
                    if true {
                        var v35 uint = 0
                        var v36 uint32 = v30
                        var v37 uint16 = 42
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                }
                v38 := GenS3{F0: v30, F1: 255, F2: 4294967295}
                if (42 >= 0) {
                    fmt.Println(v32)
                    switch "" {
                    case "hello":
                        v39 := GenS4{F0: v32, F1: false, F2: 1}
                        _ = v39
                    case "\n\t":
                        var v40 float64 = v28
                        _ = v40
                    case "a":
                        delete(v2, (v30 << ((179) & 63)))
                    }
                }
                fmt.Println(v1, fi33, fi33)
                _ = fi34
                _ = v38
            }
            _ = fi33
        }
        _ = v32
    }
    v1 = (v1 >> ((521) & 63))
    var v41 uint8 = 0
    var v42 float64 = 1.5
    v2[(v30 % ((v30) | 1))] = (4294967295 <= 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v41
    _ = v42
}
