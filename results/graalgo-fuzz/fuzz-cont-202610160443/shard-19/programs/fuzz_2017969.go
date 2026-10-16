package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int16 }
type GenS2 struct { F0 int16; F1 uint; F2 int }
func genFn0(p0 string) int {
    var v5 float64 = 1e10
    var v6 int64 = -9223372036854775808
    p0 = p0
    _ = p0
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint, p1 bool, p2 uint16) string {
    if false {
        fmt.Println(p1, p2, p1)
        p0 = (p0 ^ p0)
        fmt.Println(p0)
    }
    fmt.Println(p0, p2)
    if (!(!(("" != ("foo bar" + ("foo bar" + "\n\t"))) || (-9223372036854775808 >= 696)))) {
        var v7 bool = (!(!false))
        _ = v7
    } else {
        if p1 {
            fmt.Println(p2, p2)
        } else {
            p0 -= p0
            p2 = p2
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return ((("\n\t" + "\n\t") + "") + "a")
}

func main() {
    var v1 int32 = 6
    var v2 float32 = float32(float64(-2.25))
    v3 := []int{4294967295}
    v4 := make(map[int]uint32)
    var v8 string = ""
    v8 = v8
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            _ = len(v4)
            _ = fi10
        }
        switch "" {
        case "":
            v4[(-(fi9))] = 2
        case "a":
            fmt.Println(v1)
        case "foo bar":
            fmt.Println(v1)
        }
        _ = fi9
    }
    switch v8 {
    case "":
        v11 := GenS1{F0: 32767, F1: 159, F2: 1}
        _ = v11
    }
    if ((-2.25 - ((3.14159 + v2) - 1.5)) >= ((1e10 + v2) + v2)) {
        var v12 int32 = (v1 + 982)
        _ = v12
    }
    fmt.Println(v8, v8)
    var v13 float32 = (1.5 / (3.14159 / v2))
    _ = len(v3)
    var v14 int64 = 1
    var v15 int16 = 0
    for fi16 := 0; fi16 < 3; fi16++ {
        if (321 != 65535) {
            v1 = v1
            v2 = 3.14159
        } else {
            switch ((v8 + "") + "") {
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "":
                if true {
                    if true {
                        v2 -= v13
                    }
                    for fi17 := 0; fi17 < 18; fi17++ {
                        v1 = v1
                        var v18 int16 = (v15 & (v15 >> ((64) & 63)))
                        switch -591 {
                        case 0:
                            v14 = 42
                        case 32768:
                            v19 := GenS2{F0: (v15 | v18), F1: 1, F2: fi17}
                            _ = v19
                        }
                        _ = fi17
                        _ = v18
                    }
                    var v20 uint16 = 0
                    _ = v20
                } else {
                    if false {
                        var v21 int16 = v15
                        v22 := GenS2{F0: (-(v21)), F1: 65535, F2: genFn0((v8 + ("hello" + "\n\t")))}
                        _ = v21
                        _ = v22
                    }
                }
            case "hello":
                if ((^(fi16)) == 127) {
                    v23 := GenS2{F0: -400, F1: 4294967295, F2: fi16}
                    var v24 uint32 = 0
                    var v25 int16 = v15
                    _ = v23
                    _ = v24
                    _ = v25
                }
            default:
                var v26 uint16 = 752
                _ = v26
            }
            for fi27 := 0; fi27 < 18; fi27++ {
                switch "foo bar" {
                case "hello":
                    v28 := GenS2{F0: (v15 ^ v15), F1: 1, F2: fi27}
                    _ = v28
                case "":
                    v3 = append(v3, (fi27 % ((fi27) | 1)))
                }
                _ = fi27
            }
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 int32 = 2147483647
                switch genFn0(v8) {
                case 4294967295:
                    v2 = v2
                }
                _ = fi29
                _ = v30
            }
        }
        var v31 int16 = (-(v15))
        if len(v3) > 0 {
            v3[0] = fi16
        }
        _ = fi16
        _ = v31
    }
    v4[2147483647] = 2
    v4[-141] = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
