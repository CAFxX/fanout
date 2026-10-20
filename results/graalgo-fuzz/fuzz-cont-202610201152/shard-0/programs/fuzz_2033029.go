package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint }
type GenS2 struct { F0 int16; F1 bool }
func genFn0() int64 {
    var v4 int32 = 2147483647
    var v5 uint32 = 284
    var v6 int16 = -32768
    v7 := GenS1{F0: (v6 << ((42) & 63)), F1: true, F2: 65535}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return int64(v4)
}
func genFn1(p0 bool, p1 string, p2 uint64) int {
    p2 += p2
    p2 = p2
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 string = (p1 + ("foo bar" + "hello"))
        fmt.Println(p0, p1, v9)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    v1 := []string{("\n\t" + ((((("a" + "a") + ("\n\t" + "")) + (("foo bar" + "foo bar") + "")) + ("hello" + "foo bar")) + ("\n\t" + "foo bar")))}
    v2 := make(map[uint32]string)
    v3 := []int{9223372036854775807, -627, -409, 4294967295}
    var v10 int8 = 0
    fmt.Println(v10)
    v10 = v10
    v10 = v10
    fmt.Println(v10)
    var v11 uint64 = 18446744073709551615
    v11 -= v11
    if (124 != v10) {
        for fi12 := 0; fi12 < 4; fi12++ {
            for fi13 := 0; fi13 < 4; fi13++ {
                if (42 < 1) {
                    v3 = append(v3, (fi13 / (((^(fi12))) | 1)))
                }
                _ = fi13
            }
            v10 = (v10 / (((-(v10))) | 1))
            for fi14 := 0; fi14 < 1; fi14++ {
                switch (("" + "a") + "\n\t") {
                case "hello":
                    if ((0 == 255) && false) {
                        switch "" {
                        case "foo bar":
                            var v15 int64 = -312
                            _ = v15
                        case "":
                            v11 -= 847
                        }
                        v16 := GenS2{F0: 980, F1: (true && (!true))}
                        _ = v16
                    } else {
                        v2[4294967295] = ((("foo bar" + (("\n\t" + "hello") + "hello")) + ("hello" + "a")) + "foo bar")
                        v11 = (v11 & v11)
                        switch genFn1(((4294967295 < 18446744073709551615) && (783 != 1)), "foo bar", 42) {
                        case 32768:
                            delete(v2, 42)
                        case -2147483648:
                            var v17 uint8 = 232
                            _ = v17
                        case -9223372036854775808:
                            var v18 uint = 1
                            _ = v18
                        default:
                            v11 -= 0
                        }
                    }
                case "a":
                    var v19 uint = uint(v11)
                    _ = v19
                }
                _ = fi14
            }
            _ = fi12
        }
        var v20 uint = 4294967295
        _ = v20
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi21 := 0; fi21 < 18; fi21++ {
            v10 = (v10 & v10)
            _ = fi21
        }
    }
    v22 := GenS1{F0: -32768, F1: false, F2: 0}
    for fi23 := 0; fi23 < 18; fi23++ {
        v10 = 121
        var v24 bool = (true || ((v10 << ((1) & 63)) <= -128))
        for fi25 := 0; fi25 < 11; fi25++ {
            var v26 float64 = float64(((3.14159 + float32(1e10)) / float32(0.1)))
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    var v27 float64 = float64((0.1 + 1e10))
    v28 := GenS2{F0: -1, F1: false}
    fmt.Println(v27, v11)
    v11 = v11
    if true {
        v29 := GenS2{F0: 32767, F1: false}
        _ = v29
    } else {
        switch 0 {
        case -190:
            v27 = v27
        default:
            if (647 < 65535) {
                v27 = v27
            }
        }
        var v30 uint8 = 1
        _ = v30
    }
    var v31 uint32 = 2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v31
}
