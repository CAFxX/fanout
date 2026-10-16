package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int16; F1 uint; F2 uint16 }
func genFn0(p0 int16) uint {
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint8 = 105
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(fi9, p0)
        p0 = p0
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    p0 += p0
    _ = p0
    return 255
}
func genFn1(p0 uint64, p1 string) string {
    for fi11 := 0; fi11 < 3; fi11++ {
        p1 = p1
        if (0.0 > 0.0) {
            var v12 uint8 = 42
            _ = v12
        }
        switch -9223372036854775808 {
        case 531:
            if true {
                var v13 uint64 = (p0 - p0)
                v13 = (^(v13))
                _ = v13
            } else {
                v14 := GenS1{F0: -128, F1: 669}
                fmt.Println(p1, p0, p1)
                v15 := GenS1{F0: 2, F1: 255}
                _ = v14
                _ = v15
            }
        }
        _ = fi11
    }
    p0 = p0
    _ = p0
    _ = p1
    return (p1 + "hello")
}

func main() {
    var v1 int = -588
    v2 := []string{"foo bar", ("\n\t" + "hello")}
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    var v5 uint32 = 42
    var v6 string = "foo bar"
    var v16 string = "a"
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 string = (v6 + v6)
        _ = fi17
        _ = v18
    }
    if false {
        switch (v1 & v1) {
        case 1:
            v16 = ""
        case -879:
            for fi19 := 0; fi19 < 12; fi19++ {
                for fi20 := 0; fi20 < 2; fi20++ {
                    v21 := v3[(v5 ^ (v5 + 0))]
                    var v22 float32 = float32(0.0)
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = fi19
            }
        default:
            var v23 uint8 = 255
            _ = v23
        }
        v24 := GenS1{F0: 0, F1: 1}
        _ = v24
    }
    var v25 float64 = (3.14159 + 0.0)
    var v26 bool = (false || true)
    v16 = v6
    var v27 float64 = ((-2.25 + (v25 * (v25 + 1e10))) + v25)
    fmt.Println(v1, v25)
    v16 = ("hello" + ((v6 + v6) + "a"))
    if len(v2) > 0 {
        v2[0] = (v6 + ((v6 + "hello") + "a"))
    }
    fmt.Println(v16)
    var v28 uint16 = 255
    var v29 uint8 = 255
    var v30 uint8 = v29
    v31 := GenS2{F0: -1}
    v2 = append(v2, "foo bar")
    if (false || (0 < 1)) {
        for fi32 := 0; fi32 < 10; fi32++ {
            v4[("hello" + v16)] = v5
            if (127 >= v1) {
                _ = len(v2)
            }
            var v33 uint32 = 2147483647
            _ = fi32
            _ = v33
        }
        if v26 {
            v34 := GenS3{F0: 543, F1: v26}
            for fi35 := 0; fi35 < 3; fi35++ {
                if len(v2) > 0 {
                    v2[0] = ""
                }
                _ = fi35
            }
            v36 := GenS4{F0: 1000, F1: 42, F2: v28}
            _ = v34
            _ = v36
        }
        v37 := GenS4{F0: 32767, F1: 4294967295, F2: (v28 / (((v28 >> ((18446744073709551615) & 63))) | 1))}
        _ = v37
    } else {
        v26 = (0 > 913)
        var v38 uint8 = uint8(v5)
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
