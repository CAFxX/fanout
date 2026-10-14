package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int8 }
type GenS2 struct { F0 uint64; F1 string; F2 string }
type GenS3 struct { F0 int; F1 int8 }
func genFn0(p0 int64, p1 uint32, p2 int16) int16 {
    if true {
        p2 = p2
        p1 = 4294967295
        var v7 bool = (int(p1) < 42)
        _ = v7
    } else {
        v8 := []int{42, 4294967295}
        _ = v8
    }
    if ((117 <= p2) || (0 < 1)) {
        v9 := GenS1{F0: 127, F1: uint(p1), F2: 0}
        if false {
            p1 = 959
        } else {
            var v10 float32 = 0.1
            _ = v10
        }
        _ = v9
    }
    v11 := GenS1{F0: -128, F1: 4294967295, F2: int8(p1)}
    v12 := []int{-2147483648, -1, -1, -538}
    if true {
        switch 4294967295 {
        case 1:
            if len(v12) > 0 {
                v12[0] = 128
            }
        case 42:
            v13 := GenS1{F0: -107, F1: 975, F2: 62}
            _ = v13
        default:
            if len(v12) > 0 {
                v12[0] = 2147483647
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return (^(p2))
}
func genFn1(p0 int32, p1 int64, p2 string) int16 {
    p0 += 0
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return -32768
}

func main() {
    v1 := []string{"", "hello", ("a" + ((("hello" + (((("" + "") + "\n\t") + ((("hello" + "") + "hello") + "foo bar")) + "foo bar")) + ("\n\t" + "foo bar")) + ((("" + "") + (("a" + (("\n\t" + ("foo bar" + "foo bar")) + "foo bar")) + (("hello" + ((("foo bar" + "") + (("\n\t" + ("\n\t" + ("a" + ""))) + "")) + "")) + "hello"))) + "hello")))}
    v2 := make(map[string]int)
    var v3 bool = true
    v4 := make(map[string]bool)
    var v5 uint16 = 65535
    v6 := make(map[string]uint32)
    if v3 {
        v14 := GenS2{F0: 1, F1: "", F2: "hello"}
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        var v15 int8 = 0
        _ = v14
        _ = v15
    } else {
        v3 = v3
    }
    var v16 float32 = (3.14159 + 1.5)
    var v17 int16 = -281
    if v3 {
        switch "hello" {
        case "\n\t":
            v18 := GenS3{F0: 32768, F1: 1}
            _ = v18
        }
        if false {
            var v19 int32 = 0
            v20 := GenS2{F0: 18446744073709551615, F1: "", F2: ("\n\t" + ("foo bar" + "\n\t"))}
            _ = v19
            _ = v20
        } else {
            v1 = append(v1, ("\n\t" + "hello"))
            if v3 {
                var v21 int = 42
                _ = v21
            } else {
                var v22 uint16 = (^(v5))
                if v3 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v3 = v3
                    v23 := GenS1{F0: 42, F1: 255, F2: 0}
                    v24 := GenS1{F0: 127, F1: 64, F2: -1}
                    _ = v23
                    _ = v24
                }
                _ = v22
            }
        }
    } else {
        var v25 float32 = float32(0.1)
        fmt.Println(v17, v25, v16)
        var v26 bool = (!true)
        _ = v25
        _ = v26
    }
    v27 := GenS2{F0: 1, F1: "a", F2: ("" + "hello")}
    if len(v1) > 0 {
        v1[0] = ("" + "\n\t")
    }
    var v28 float64 = ((3.14159 * 1e10) * ((((float64(v16) - 1.5) - (((0.1 + 0.1) - 3.14159) - (1e10 - 0.0))) - (-2.25 * 0.0)) + 1.5))
    if (0 >= 0) {
        var v29 uint = 1
        delete(v6, "\n\t")
        _ = v29
    }
    switch -2147483648 {
    case 279:
        for fi30 := 0; fi30 < 15; fi30++ {
            for fi31 := 0; fi31 < 7; fi31++ {
                if v3 {
                    var v32 uint8 = 70
                    _ = v32
                } else {
                    v2["foo bar"] = -9223372036854775808
                }
                _ = fi31
            }
            _ = fi30
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
