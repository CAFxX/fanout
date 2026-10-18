package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint }
type GenS2 struct { F0 uint16; F1 uint8; F2 int32 }
func genFn0() bool {
    var v5 uint16 = 0
    v6 := []int{4294967295, 331, 4294967295}
    switch 466 {
    case 119:
        var v7 bool = ((v5 / (((^(v5))) | 1)) <= (v5 + (v5 & (v5 << ((65535) & 63)))))
        _ = v7
    }
    v5 = (^(v5))
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 int) uint64 {
    v8 := GenS1{F0: (("a" + "foo bar") + ("\n\t" + "hello")), F1: 9223372036854775807, F2: 65535}
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 string = ("" + "hello")
        for fi11 := 0; fi11 < 5; fi11++ {
            if false {
                p0 -= 32767
            } else {
                if ((0 >= 1) && (100 != 0)) {
                    var v12 int16 = 1000
                    var v13 float32 = (-2.25 * float32(-2.25))
                    v12 = (v12 - (v12 | (v12 << ((56) & 63))))
                    _ = v12
                    _ = v13
                } else {
                    v10 = (v10 + ("" + v10))
                    var v14 uint32 = 2147483647
                    if true {
                        v15 := GenS2{F0: uint16(v14), F1: 42, F2: 42}
                        _ = v15
                    } else {
                        v16 := []int{p0, 32767}
                        v17 := GenS2{F0: 1, F1: 255, F2: -2147483648}
                        var v18 int16 = 32767
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    _ = v14
                }
                var v19 uint64 = 2
                _ = v19
            }
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v8
    return 42
}

func main() {
    var v1 int8 = 127
    var v2 int32 = -1
    v3 := make(map[string]uint32)
    v4 := []uint32{}
    v20 := GenS2{F0: 1, F1: 2, F2: 860}
    switch int(v2) {
    case -631:
        fmt.Println(v2, v1)
    case -2147483648:
        switch ("a" + "") {
        case "\n\t":
            v2 += v2
        case "a":
            if ((("hello" == ("\n\t" + ("" + "hello"))) && ((v1 < v1) && (86 >= genFn1(-9223372036854775808)))) || true) {
                var v21 float64 = float64(0.1)
                switch -243 {
                case 32767:
                    fmt.Println(v21, v21, v21)
                case 88:
                    var v22 string = "a"
                    _ = v22
                default:
                    v2 -= (^(v2))
                }
                _ = v21
            } else {
                if false {
                    v23 := GenS2{F0: 688, F1: 42, F2: -1}
                    _ = v23
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi24 := 0; fi24 < 16; fi24++ {
                    fmt.Println(fi24, v2)
                    _ = fi24
                }
            }
        case "foo bar":
            v2 -= (v2 * v2)
        }
    }
    fmt.Println(v1, v2)
    var v25 int = 128
    var v26 int8 = (v1 & v1)
    v27 := GenS2{F0: 422, F1: 42, F2: -2147483648}
    v28 := GenS1{F0: (("\n\t" + "") + ("hello" + ("\n\t" + ("\n\t" + "a")))), F1: 171, F2: 2}
    if true {
        _ = len(v3)
    } else {
        v25 += (-(v25))
    }
    var v29 int32 = (v2 >> ((65535) & 63))
    fmt.Println(v1, v2)
    var v30 uint64 = 1
    var v31 uint = 566
    v29 -= 2147483647
    switch 9223372036854775807 {
    case 703:
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
    case 32767:
        if false {
            var v32 uint64 = (v30 - v30)
            if false {
                v31 -= (v31 - 2)
                var v33 uint32 = 1
                _ = v33
            } else {
                if ((v25 << (((v31 + v31)) & 63)) <= 1) {
                    var v34 bool = true
                    fmt.Println(v25, v31, v26)
                    _ = v34
                }
            }
            _ = v32
        } else {
            if true {
                v1 -= v26
            }
        }
    case -751:
        var v35 float32 = 1.5
        _ = v35
    default:
        v36 := GenS1{F0: "foo bar", F1: 9223372036854775807, F2: (v31 + (v31 + 0))}
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
