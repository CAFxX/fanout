package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint }
func genFn0(p0 uint32, p1 uint) int8 {
    v6 := GenS1{F0: 127}
    var v7 int32 = 1
    v8 := []int{-610, 0, 32767, -2147483648}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int16, p1 int8) uint8 {
    switch (("\n\t" + "") + "hello") {
    case "hello":
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int8 = p1
            _ = fi9
            _ = v10
        }
    }
    p1 -= p1
    v11 := []int{-9223372036854775808, 127}
    var v12 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := []string{((("hello" + "a") + "") + "foo bar"), "foo bar"}
    var v2 uint16 = 360
    v3 := make(map[string]uint32)
    v4 := []string{"a"}
    var v5 bool = (270 < 895)
    for fi13 := 0; fi13 < 11; fi13++ {
        if ((fi13 % ((581) | 1)) < 784) {
            var v14 int8 = 127
            _ = v14
        }
        _ = fi13
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS1{F0: int8(v2)}
        if (18446744073709551615 > 1) {
            switch (-(fi15)) {
            case 9223372036854775807:
                var v17 uint64 = 2
                _ = v17
            default:
                var v18 int = fi15
                _ = v18
            }
        } else {
            var v19 int32 = 2147483647
            var v20 string = ("" + ("hello" + "foo bar"))
            _ = v19
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    if v5 {
        v5 = true
        switch "a" {
        case "a":
            v2 += (v2 * 32767)
        case "\n\t":
            if (true || (3.14159 > 1.5)) {
                var v21 int16 = 0
                if ((1 < 630) && v5) {
                    fmt.Println(v2, v21, v21)
                    v22 := GenS1{F0: 127}
                    v21 -= v21
                    _ = v22
                } else {
                    var v23 int8 = 0
                    switch (("" + "") + "hello") {
                    case "\n\t":
                        v23 += v23
                    }
                    _ = v23
                }
                _ = v21
            } else {
                var v24 uint = 18446744073709551615
                v4 = append(v4, "foo bar")
                fmt.Println(v24)
                _ = v24
            }
        case "foo bar":
            for fi25 := 0; fi25 < 20; fi25++ {
                v2 -= (v2 - 976)
                fmt.Println(fi25)
                if (1 != 920) {
                    if v5 {
                        v26 := v3[(("foo bar" + "foo bar") + ("foo bar" + "hello"))]
                        switch "\n\t" {
                        case "":
                            switch ("a" + "\n\t") {
                            case "\n\t":
                                v26 += (v26 + v26)
                            case "":
                                if false {
                                    v2 += 1
                                    v26 = 61
                                    fmt.Println(v26)
                                } else {
                                    v26 -= (v26 ^ v26)
                                    var v27 float32 = 0.0
                                    delete(v3, ("\n\t" + ("" + ("foo bar" + "\n\t"))))
                                    _ = v27
                                }
                            case "foo bar":
                                v28 := v3["\n\t"]
                                _ = v28
                            }
                        }
                        v26 -= 2147483647
                        _ = v26
                    } else {
                        v2 -= v2
                        v2 += v2
                    }
                }
                _ = fi25
            }
        }
    }
    v2 = (v2 - 73)
    var v29 uint8 = 20
    for fi30 := 0; fi30 < 8; fi30++ {
        if false {
            v31 := GenS2{F0: -32768, F1: 2}
            fmt.Println(v29)
            v29 = 2
            _ = v31
        } else {
            v5 = v5
        }
        _ = fi30
    }
    v29 += 0
    if (-367 == 0) {
        if (-128 <= genFn0(2147483648, 306)) {
            v32 := GenS2{F0: 1000, F1: uint(v2)}
            v2 = 42
            for fi33 := 0; fi33 < 17; fi33++ {
                v34 := GenS1{F0: 42}
                var v35 uint = 499
                var v36 uint = v35
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            _ = v32
        }
        fmt.Println(v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
