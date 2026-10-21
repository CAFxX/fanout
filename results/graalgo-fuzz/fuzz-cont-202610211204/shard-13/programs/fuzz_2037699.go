package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint16, p1 bool, p2 uint32) int16 {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint32 = 42
        var v9 uint32 = p2
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(p2, p1)
    if (!p1) {
        if p1 {
            switch ("" + "a") {
            case "":
                if p1 {
                    p0 = p0
                    p1 = p1
                    v10 := GenS1{F0: 64, F1: uint16(p2)}
                    _ = v10
                }
            case "hello":
                switch "a" {
                case "":
                    var v11 float64 = ((-2.25 / ((1e10 - 1e10) + 0.1)) + (3.14159 + float64((float32(1e10) * 0.1))))
                    _ = v11
                case "hello":
                    p2 = 42
                default:
                    var v12 string = "foo bar"
                    _ = v12
                }
            case "foo bar":
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := []int{fi13}
                    for fi15 := 0; fi15 < 2; fi15++ {
                        switch ("foo bar" + "a") {
                        case "\n\t":
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        }
                        p1 = p1
                        var v16 bool = (!true)
                        _ = fi15
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                }
            default:
                if false {
                    v17 := GenS1{F0: 39, F1: 65535}
                    if p1 {
                        var v18 bool = (255 == 42)
                        if true {
                            v19 := GenS2{F0: 32768, F1: v18, F2: (p2 & 4294967295)}
                            _ = v19
                        } else {
                            var v20 float32 = float32(((0.1 + 3.14159) + 1.5))
                            var v21 uint64 = 18446744073709551615
                            fmt.Println(p1)
                            _ = v20
                            _ = v21
                        }
                        v18 = ((p0 | p0) < (p0 >> ((uint(p0)) & 63)))
                        _ = v18
                    }
                    fmt.Println(p1, p0, p2)
                    _ = v17
                } else {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        var v23 uint16 = p0
                        _ = fi22
                        _ = v23
                    }
                }
            }
            p1 = (42 != -1)
        } else {
            var v24 bool = true
            p1 = true
            v25 := GenS3{F0: 0}
            _ = v24
            _ = v25
        }
        if p1 {
            v26 := []int{127}
            v27 := GenS2{F0: 42, F1: p1, F2: p2}
            _ = v26
            _ = v27
        }
    }
    v28 := GenS4{F0: p1, F1: 138}
    var v29 uint8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v29
    return -32768
}
func genFn1(p0 uint32, p1 string, p2 uint) uint16 {
    v30 := []int{9223372036854775807, 32768, -148, 400}
    if true {
        var v31 float32 = 3.14159
        if len(v30) > 0 {
            v30[0] = 127
        }
        _ = v31
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 255
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    v5 := make(map[string]uint32)
    var v6 bool = false
    for fi32 := 0; fi32 < 19; fi32++ {
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 int = fi32
            v6 = true
            _ = fi33
            _ = v34
        }
        v2 -= genFn1(2147483647, ("a" + ("\n\t" + (("a" + ("" + "")) + ""))), uint(v2))
        v2 = v2
        _ = fi32
    }
    var v35 int64 = 0
    v36 := []int{0, -1}
    if (true && ((!v6) && v6)) {
        fmt.Println(v2)
    } else {
        if ((^(v35)) > 1) {
            v2 = genFn1(1, (("\n\t" + ("" + "a")) + ("a" + "hello")), 65535)
            for fi37 := 0; fi37 < 16; fi37++ {
                fmt.Println(fi37, fi37, v2)
                _ = fi37
            }
        } else {
            var v38 float64 = 0.1
            _ = v38
        }
    }
    var v39 uint = 999
    v6 = v6
    var v40 uint16 = (v2 / ((65535) | 1))
    delete(v3, -662)
    if len(v36) > 0 {
        v36[0] = int(v35)
    }
    v40 = v2
    v40 = v2
    var v41 uint32 = 4294967295
    _ = len(v4)
    var v42 bool = v6
    v40 = 543
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
