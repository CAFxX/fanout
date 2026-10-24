package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint64 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int32; F1 int8; F2 uint }
func genFn0(p0 int16, p1 int8, p2 string) string {
    v7 := GenS1{F0: (!(false || false)), F1: "\n\t", F2: 63}
    p0 -= p0
    p1 = -128
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := []int{fi8, 32767, fi8, 128}
        p1 = (p1 & (p1 | p1))
        p1 -= (p1 * p1)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 + ("" + ((p2 + ("" + (p2 + "\n\t"))) + p2)))
}

func main() {
    var v1 float64 = 0.1
    v2 := []int{0, -717}
    var v3 bool = ("" >= (((((((("hello" + "hello") + (("hello" + (("foo bar" + "hello") + "")) + ("hello" + ("a" + ("foo bar" + ("a" + ("\n\t" + (("hello" + "a") + "\n\t")))))))) + "") + "") + "foo bar") + "\n\t") + "hello") + "foo bar"))
    v4 := make(map[int]uint32)
    v5 := make(map[int]int)
    v6 := make(map[string]uint32)
    v2 = append(v2, -284)
    for fi10 := 0; fi10 < 9; fi10++ {
        v1 += v1
        if (!true) {
            v1 -= (v1 + 3.14159)
        } else {
            v11 := GenS1{F0: (!false), F1: ((("hello" + "hello") + "\n\t") + ("a" + (("\n\t" + "a") + "\n\t"))), F2: 18446744073709551615}
            v3 = v3
            _ = v11
        }
        v12 := GenS1{F0: (true && false), F1: "", F2: 1}
        _ = fi10
        _ = v12
    }
    var v13 uint16 = 1
    var v14 uint16 = (v13 | (v13 & (v13 >> ((65535) & 63))))
    v15 := GenS1{F0: (false || ((v13 & v13) != v14)), F1: "\n\t", F2: 490}
    var v16 int8 = 123
    if v3 {
        fmt.Println(v14, v13)
        var v17 uint = 18446744073709551615
        var v18 string = "\n\t"
        _ = v17
        _ = v18
    }
    v14 += v13
    if (1 < 42) {
        switch -1 {
        case -1:
            switch ((("\n\t" + "a") + "foo bar") + "a") {
            case "hello":
                var v19 int16 = 1
                _ = v19
            case "\n\t":
                v2 = append(v2, -317)
            }
        }
    } else {
        for fi20 := 0; fi20 < 17; fi20++ {
            if (("\n\t" + "a") >= (("\n\t" + ("foo bar" + "")) + "foo bar")) {
                v3 = false
            }
            if v3 {
                for fi21 := 0; fi21 < 15; fi21++ {
                    var v22 string = "\n\t"
                    v23 := GenS2{F0: v3, F1: 95}
                    v13 += (v14 + 255)
                    v24 := GenS2{F0: (((v16 >= v16) || v3) && (true || v3)), F1: 134}
                    var v25 string = (v22 + "")
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                fmt.Println(v13)
            } else {
                var v26 int = 594
                fmt.Println(v26, v3)
                v27 := GenS3{F0: 0, F1: (v26 & -1)}
                _ = v26
                _ = v27
            }
            _ = fi20
        }
        v2 = append(v2, -262)
    }
    var v28 int8 = (-(v16))
    v16 -= v28
    for fi29 := 0; fi29 < 20; fi29++ {
        v16 -= v16
        v30 := GenS2{F0: v3, F1: 1000}
        _ = fi29
        _ = v30
    }
    if false {
        if true {
            switch int(v16) {
            case 9223372036854775807:
                var v31 float32 = float32(1e10)
                _ = v31
            }
        } else {
            for fi32 := 0; fi32 < 4; fi32++ {
                var v33 string = ("hello" + ("" + ("" + "hello")))
                v16 -= v16
                if v3 {
                    v3 = v3
                } else {
                    v34 := GenS3{F0: 2147483647, F1: 2147483647}
                    switch fi32 {
                    case 0:
                        v35 := GenS4{F0: 1, F1: v28, F2: 546}
                        _ = v35
                    case -751:
                        var v36 int8 = (-(v16))
                        _ = v36
                    case 1:
                        _ = len(v6)
                    default:
                        var v37 uint32 = 2147483648
                        _ = v37
                    }
                    _ = v34
                }
                _ = fi32
                _ = v33
            }
            fmt.Println(v16, v28, v14)
            switch "" {
            case "a":
                var v38 int = -2147483648
                _ = v38
            }
        }
    }
    v2 = append(v2, 42)
    if len(v2) > 0 {
        v2[0] = 32768
    }
    var v39 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
