package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint16; F1 bool; F2 bool }
type GenS3 struct { F0 string }
func genFn0(p0 string, p1 int32, p2 string) uint16 {
    if (!(255 != 65535)) {
        p0 = p0
        p0 = p0
    } else {
        v7 := GenS1{F0: true, F1: 1}
        v8 := GenS1{F0: (false || true), F1: 0}
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := []int{(fi9 + 32767), fi9, (fi9 % ((fi9) | 1)), fi9}
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: (true && (!false)), F1: 65535}
            p1 -= p1
            _ = fi11
            _ = v12
        }
        if (!false) {
            _ = len(v10)
            for fi13 := 0; fi13 < 5; fi13++ {
                if true {
                    var v14 int = (fi13 / ((fi13) | 1))
                    v14 += -1
                    _ = v14
                } else {
                    fmt.Println(p1, fi9, p1)
                    if len(v10) > 0 {
                        v10[0] = fi13
                    }
                    var v15 int = (fi9 % ((fi13) | 1))
                    _ = v15
                }
                _ = fi13
            }
        } else {
            switch ((p2 + p2) + p2) {
            case "":
                if true {
                    var v16 uint = 4294967295
                    v16 = (^(v16))
                    if true {
                        var v17 int = -49
                        _ = len(v10)
                        _ = len(v10)
                        _ = v17
                    }
                    _ = v16
                }
            default:
                var v18 float32 = (-2.25 + -2.25)
                _ = v18
            }
            p2 = ((((p0 + "foo bar") + (p0 + p2)) + p2) + ("foo bar" + ""))
        }
        _ = fi9
        _ = v10
    }
    if (128 <= 42) {
        v19 := []int{}
        _ = v19
    } else {
        p2 = p0
        p0 = "foo bar"
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 int = 42
            p0 = (("foo bar" + (("" + p2) + p2)) + (p2 + p2))
            _ = fi20
            _ = v21
        }
    }
    p2 = "hello"
    if (0 >= 0) {
        for fi22 := 0; fi22 < 1; fi22++ {
            fmt.Println(p1)
            var v23 uint8 = 127
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := GenS1{F0: (428 != 0), F1: 65535}
            _ = fi24
            _ = v25
        }
        var v26 uint = 105
        _ = v26
    } else {
        var v27 int32 = p1
        var v28 int16 = 703
        v29 := []int{797, 1}
        _ = v27
        _ = v28
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 int = 2147483647
    var v2 int32 = -275
    var v3 int32 = (v2 & (^(v2)))
    var v4 int32 = (v3 * v3)
    v5 := []int{(v1 - -2147483648), (v1 + v1)}
    var v6 uint = 1
    var v30 bool = false
    for fi31 := 0; fi31 < 5; fi31++ {
        fmt.Println(v2, v4, v3)
        v5 = append(v5, fi31)
        _ = len(v5)
        _ = fi31
    }
    var v32 float64 = (3.14159 + -2.25)
    for fi33 := 0; fi33 < 20; fi33++ {
        switch fi33 {
        case 32767:
            v34 := GenS2{F0: 0, F1: v30, F2: (2 <= 648)}
            _ = v34
        }
        _ = fi33
    }
    var v35 int64 = 1
    v36 := GenS1{F0: v30, F1: v6}
    var v37 uint8 = 27
    var v38 uint = 18446744073709551615
    var v39 int32 = 0
    v32 -= (v32 + v32)
    var v40 float64 = v32
    v41 := GenS3{F0: "a"}
    v39 = (v4 / ((v3) | 1))
    switch "a" {
    case "\n\t":
        var v42 int32 = -1
        _ = v42
    case "":
        v35 += 551
    case "hello":
        v6 -= v6
    default:
        v2 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
