package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 bool; F1 int16; F2 uint8 }
func genFn0(p0 string, p1 uint32, p2 uint32) uint32 {
    if (true || true) {
        v6 := []int{127, 2147483647, 127}
        var v7 int = 42
        _ = v6
        _ = v7
    }
    p2 -= 42
    _ = p0
    _ = p1
    _ = p2
    return p2
}
func genFn1() int {
    v8 := []int{-869, -9223372036854775808, 127}
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 bool = (81 == 4294967295)
            if false {
                v11 = v11
                v11 = v11
            } else {
                if len(v8) > 0 {
                    v8[0] = (fi10 - fi9)
                }
                var v12 int = 61
                v13 := GenS1{F0: 42, F1: 120, F2: 32767}
                _ = v12
                _ = v13
            }
            v11 = v11
            v8 = append(v8, (fi10 | 0))
            fmt.Println(fi10, v11, fi10)
            fmt.Println(fi9)
            _ = len(v8)
            v11 = v11
            fmt.Println(v11, v11, fi9)
            var v14 uint8 = 161
            _ = fi10
            _ = v11
            _ = v14
        }
        v15 := GenS1{F0: 4294967295, F1: -128, F2: 800}
        _ = fi9
        _ = v15
    }
    v8 = append(v8, -870)
    _ = v8
    return 127
}

func main() {
    var v1 uint8 = 2
    var v2 string = (("foo bar" + "foo bar") + ("hello" + "a"))
    var v3 int64 = -542
    var v4 uint32 = 1
    var v5 uint32 = (v4 + (v4 % ((v4) | 1)))
    v16 := []int{-214, 127, genFn1()}
    var v17 float32 = ((0.1 / 1e10) / (3.14159 / 0.1))
    switch -262 {
    case -2147483648:
        if false {
            if true {
                switch ((("\n\t" + v2) + (v2 + "hello")) + v2) {
                case "foo bar":
                    v17 = 0.0
                default:
                    if len(v16) > 0 {
                        v16[0] = 1
                    }
                }
            } else {
                v4 = v5
                fmt.Println(v5)
            }
        }
    case -742:
        v18 := GenS2{F0: 65535}
        _ = v18
    default:
        switch 420 {
        case 42:
            if ((-2.25 <= ((0.1 + 3.14159) - (0.1 - 0.1))) || (-373 < 156)) {
                var v19 int64 = v3
                if (!true) {
                    var v20 float32 = v17
                    _ = len(v16)
                    fmt.Println(v1)
                    _ = v20
                }
                _ = v19
            } else {
                fmt.Println(v1)
            }
        }
    }
    v17 += float32(-2.25)
    fmt.Println(v1, v3, v5)
    fmt.Println(v17)
    var v21 int32 = 42
    fmt.Println(v21, v3, v21)
    var v22 int = 32768
    var v23 int = (v22 / ((649) | 1))
    v21 += v21
    for fi24 := 0; fi24 < 9; fi24++ {
        for fi25 := 0; fi25 < 14; fi25++ {
            v21 = v21
            _ = fi25
        }
        switch v23 {
        case 127:
            var v26 uint16 = 255
            _ = v26
        case 776:
            v17 -= float32(1e10)
        }
        _ = fi24
    }
    v27 := GenS3{F0: "foo bar", F1: 666}
    var v28 int = genFn1()
    for fi29 := 0; fi29 < 18; fi29++ {
        v30 := GenS4{F0: v21, F1: ""}
        _ = fi29
        _ = v30
    }
    var v31 uint = 139
    for fi32 := 0; fi32 < 15; fi32++ {
        v17 = v17
        v33 := GenS5{F0: false, F1: 42, F2: (v1 & (^(v1)))}
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
