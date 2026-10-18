package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int; F1 int32; F2 bool }
func genFn0(p0 int64, p1 bool, p2 uint16) int64 {
    switch -663 {
    case -871:
        if p1 {
            p1 = p1
            v6 := []int{966}
            _ = v6
        }
    }
    fmt.Println(p0, p2, p1)
    if (((1.5 / 1.0) + 1.5) < (0.0 + 3.14159)) {
        if (!(p1 || (false || (!p1)))) {
            p1 = false
            p0 = (-(p0))
        }
    }
    p1 = (115 < int8(p2))
    var v7 uint16 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 | (p0 & p0))
}
func genFn1(p0 int64, p1 uint64, p2 uint64) int {
    fmt.Println(p2, p2, p2)
    v8 := GenS1{F0: false, F1: 8, F2: (^(p1))}
    var v9 uint32 = 617
    var v10 int = 1
    var v11 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return (v10 ^ (v10 + v10))
}

func main() {
    var v1 uint16 = 341
    var v2 float32 = float32((3.14159 + 0.0))
    var v3 uint64 = 746
    v4 := []uint32{518, 42, 0}
    v5 := make(map[int]uint32)
    if (-32768 == 32767) {
        for fi12 := 0; fi12 < 16; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                v2 += v2
                var v14 uint8 = 143
                _ = fi13
                _ = v14
            }
            var v15 uint32 = 312
            _ = fi12
            _ = v15
        }
        v5[127] = 109
        v16 := GenS1{F0: false, F1: 32767, F2: v3}
        _ = v16
    }
    fmt.Println(v2, v1, v2)
    _ = len(v5)
    fmt.Println(v3, v3)
    for fi17 := 0; fi17 < 14; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 uint64 = (^(v3))
            _ = fi18
            _ = v19
        }
        v20 := GenS1{F0: false, F1: v1, F2: v3}
        var v21 int16 = 1
        _ = fi17
        _ = v20
        _ = v21
    }
    var v22 uint16 = (v1 / ((v1) | 1))
    v2 = (3.14159 - (v2 - (((v2 - 0.1) * (-2.25 - 3.14159)) + (v2 / float32(0.1)))))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint32 = 1
        v22 = (v22 / ((686) | 1))
        var v25 uint64 = uint64(v24)
        _ = fi23
        _ = v24
        _ = v25
    }
    switch ("foo bar" + "") {
    case "hello":
        if len(v4) > 0 {
            v4[0] = 337
        }
    }
    v26 := GenS1{F0: true, F1: 0, F2: (v3 / ((v3) | 1))}
    switch ("foo bar" + ("\n\t" + (("foo bar" + "foo bar") + ("foo bar" + (("hello" + ("\n\t" + "\n\t")) + "foo bar"))))) {
    case "\n\t":
        var v27 uint8 = 250
        _ = v27
    case "foo bar":
        if true {
            var v28 uint16 = v22
            _ = v28
        } else {
            if true {
                switch ("hello" + ((("\n\t" + ("a" + "hello")) + "foo bar") + "foo bar")) {
                case "":
                    v22 = v1
                case "hello":
                    if true {
                        for fi29 := 0; fi29 < 12; fi29++ {
                            v30 := GenS2{F0: 42, F1: 2}
                            v31 := GenS1{F0: true, F1: (v1 % (((v1 << ((1) & 63))) | 1)), F2: (v3 + v3)}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        fmt.Println(v22, v3)
                        var v32 uint16 = 42
                        _ = v32
                    }
                case "a":
                    for fi33 := 0; fi33 < 15; fi33++ {
                        var v34 string = ("a" + "hello")
                        _ = fi33
                        _ = v34
                    }
                }
            }
            v35 := GenS1{F0: ((32767 == -847) || true), F1: v1, F2: v3}
            _ = v35
        }
    case "a":
        for fi36 := 0; fi36 < 18; fi36++ {
            v37 := GenS2{F0: 206, F1: 1}
            v1 = (v1 * 1)
            v38 := GenS3{F0: (fi36 - 0), F1: int32(fi36), F2: false}
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    v2 = v2
    if true {
        v22 = v1
        switch ("a" + "\n\t") {
        case "hello":
            for fi39 := 0; fi39 < 4; fi39++ {
                v40 := GenS2{F0: 1, F1: 255}
                fmt.Println(v2, v1, v3)
                var v41 float64 = 1e10
                _ = fi39
                _ = v40
                _ = v41
            }
        default:
            _ = len(v4)
        }
        if true {
            var v42 int = -9223372036854775808
            _ = v42
        }
    }
    v1 -= (v1 >> ((1) & 63))
    var v43 uint16 = v22
    switch 4294967295 {
    case 896:
        var v44 uint64 = (v3 << ((162) & 63))
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
