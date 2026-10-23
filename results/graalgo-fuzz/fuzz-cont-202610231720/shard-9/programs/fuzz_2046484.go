package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int16; F2 uint32 }
func genFn0(p0 int32) uint64 {
    var v7 int64 = -1
    fmt.Println(v7, p0)
    switch ("" + (("a" + "\n\t") + "hello")) {
    case "":
        v8 := []int{-9223372036854775808, 0, -413}
        _ = v8
    case "foo bar":
        v7 += 0
    default:
        if false {
            p0 = p0
            p0 = p0
        } else {
            v7 = (^(v7))
            v9 := GenS1{F0: p0, F1: 127}
            v7 += (v7 >> ((821) & 63))
            _ = v9
        }
    }
    _ = p0
    _ = v7
    return 4294967295
}
func genFn1() string {
    v10 := GenS1{F0: 344, F1: 106}
    if (4294967295 <= 581) {
        v11 := GenS1{F0: 846, F1: 1}
        if (849 != 42) {
            var v12 string = ("" + "hello")
            _ = v12
        } else {
            if false {
                v13 := GenS1{F0: 1, F1: 128}
                v14 := GenS2{F0: true}
                _ = v13
                _ = v14
            }
        }
        var v15 int16 = -32768
        _ = v11
        _ = v15
    } else {
        var v16 float32 = -2.25
        _ = v16
    }
    _ = v10
    return "a"
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 32767
    v3 := []int{v2, v2, 42}
    var v4 uint32 = 65535
    var v5 int8 = 82
    var v6 uint8 = 12
    v6 -= (^(v6))
    var v17 bool = false
    if (!("\n\t" != ("foo bar" + ("\n\t" + ("" + ("" + "")))))) {
        _ = len(v1)
        var v18 int32 = 42
        _ = v18
    } else {
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 float32 = 3.14159
            _ = fi19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 uint = 255
        var v23 int16 = 1000
        v17 = v17
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v25 := GenS1{F0: 2147483647, F1: (v2 << ((62) & 63))}
        _ = fi24
        _ = v25
    }
    _ = len(v1)
    v26 := GenS3{F0: -1, F1: -1, F2: v4}
    v27 := GenS2{F0: v17}
    var v28 int = v2
    if len(v3) > 0 {
        v3[0] = 128
    }
    if false {
        if v17 {
            if ((true || true) && true) {
                switch "foo bar" {
                case "foo bar":
                    switch ("" + "hello") {
                    case "hello":
                        var v29 uint16 = 854
                        _ = v29
                    }
                }
            }
            if true {
                var v30 uint32 = 2147483648
                if (1e10 != 1e10) {
                    var v31 int16 = 1000
                    _ = v31
                } else {
                    v30 = 727
                }
                v30 = (v30 - uint32(v6))
                _ = v30
            } else {
                v32 := GenS2{F0: (genFn0(549) < genFn0(42))}
                _ = v32
            }
            fmt.Println(v6, v17, v28)
        } else {
            var v33 uint16 = 0
            _ = v33
        }
        v17 = v17
        v3 = append(v3, (v28 - 872))
    } else {
        v34 := GenS2{F0: true}
        _ = len(v3)
        _ = v34
    }
    var v35 uint8 = (v6 * v6)
    var v36 uint = 18446744073709551615
    v37 := GenS2{F0: (1 > -1)}
    var v38 uint8 = v35
    if ((int32(v38) == int32(v5)) || v17) {
        for fi39 := 0; fi39 < 18; fi39++ {
            for fi40 := 0; fi40 < 3; fi40++ {
                fmt.Println(v17, v6, v6)
                _ = fi40
            }
            var v41 float32 = (0.0 / ((0.1 + 0.0) - 1.5))
            v42 := GenS2{F0: v17}
            _ = fi39
            _ = v41
            _ = v42
        }
        if len(v3) > 0 {
            v3[0] = v2
        }
        if true {
            _ = len(v1)
        } else {
            if len(v3) > 0 {
                v3[0] = v28
            }
        }
    }
    var v43 uint64 = 1
    switch ((("foo bar" + ("foo bar" + "foo bar")) + "hello") + (("a" + "\n\t") + "\n\t")) {
    case "hello":
        fmt.Println(v38)
    case "a":
        var v44 string = "a"
        _ = v44
    default:
        v35 -= (v35 % ((v6) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
