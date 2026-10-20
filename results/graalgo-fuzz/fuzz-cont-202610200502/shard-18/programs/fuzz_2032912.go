package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32, p1 int32) int32 {
    p1 = (p1 - (p1 - 1))
    v6 := []int{128, 903, 4294967295}
    var v7 uint64 = 18446744073709551615
    if (("" + "foo bar") == "\n\t") {
        v8 := GenS1{F0: v7, F1: 100, F2: p1}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 * (p1 >> ((0) & 63)))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    var v4 int8 = 3
    var v5 uint32 = 1
    v4 = v4
    var v9 uint64 = 18446744073709551615
    v4 = (v4 ^ v4)
    v4 += 42
    v10 := GenS1{F0: 0, F1: (v4 << ((1) & 63)), F2: 0}
    if (!(v9 != (^(v9)))) {
        v5 -= 65535
    } else {
        if true {
            var v11 uint = 873
            var v12 int8 = -120
            v13 := GenS1{F0: v9, F1: v12, F2: int32(v5)}
            _ = v11
            _ = v12
            _ = v13
        }
    }
    if (3.14159 != 1.5) {
        v14 := []int{32768, -165, 32768}
        _ = v14
    } else {
        v15 := []int{2147483647, -11, -1}
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := []int{2147483647, (^(fi16)), 4294967295}
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 227}
    if ((-2.25 != (1e10 + 0.1)) || true) {
        switch "foo bar" {
        case "foo bar":
            v9 = 4294967295
        case "a":
            if ("a" == ((("" + "a") + "a") + (((("" + "foo bar") + "hello") + "a") + "hello"))) {
                if true {
                    var v19 int32 = 1
                    v3[127] = 32768
                    switch -9223372036854775808 {
                    case -2147483648:
                        v9 -= (v9 >> ((1) & 63))
                    case -1:
                        var v20 uint64 = v9
                        _ = v20
                    case 32767:
                        for fi21 := 0; fi21 < 16; fi21++ {
                            switch "a" {
                            case "":
                                for fi22 := 0; fi22 < 19; fi22++ {
                                    switch fi21 {
                                    case 127:
                                        var v23 int32 = 2147483647
                                        _ = v23
                                    default:
                                        var v24 int16 = -1
                                        _ = v24
                                    }
                                    var v25 int32 = genFn0(1, (^(v19)))
                                    if ((v4 + v4) <= (v4 * v4)) {
                                        v26 := GenS2{F0: 326}
                                        fmt.Println(v5, v5)
                                        v4 = (v4 << ((452) & 63))
                                        _ = v26
                                    }
                                    _ = fi22
                                    _ = v25
                                }
                            case "a":
                                v27 := GenS1{F0: v9, F1: v4, F2: v19}
                                _ = v27
                            }
                            fmt.Println(v4, v5)
                            _ = fi21
                        }
                    default:
                        var v28 int = -2147483648
                        _ = v28
                    }
                    _ = v19
                }
            }
        }
    } else {
        switch 424 {
        case -355:
            var v29 bool = true
            _ = v29
        case 32768:
            v4 += -1
        case 128:
            switch -595 {
            case 608:
                var v30 int = 32767
                _ = v30
            default:
                if true {
                    for fi31 := 0; fi31 < 8; fi31++ {
                        v9 = 229
                        for fi32 := 0; fi32 < 11; fi32++ {
                            var v33 bool = (3.14159 <= -2.25)
                            v34 := []int{fi32, 278}
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        v35 := GenS1{F0: v9, F1: (v4 - (-(v4))), F2: 2147483647}
                        _ = fi31
                        _ = v35
                    }
                    v4 = v4
                }
            }
        }
        if (!false) {
            var v36 uint64 = (v9 >> ((4294967295) & 63))
            for fi37 := 0; fi37 < 7; fi37++ {
                var v38 float32 = (-2.25 / 1.0)
                _ = fi37
                _ = v38
            }
            var v39 int32 = 42
            _ = v36
            _ = v39
        }
    }
    v40 := GenS2{F0: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v9
}
