package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint64; F2 int32 }
type GenS5 struct { F0 int64; F1 string }
func genFn0(p0 uint8, p1 int16, p2 int16) int8 {
    p2 = 711
    var v6 string = ""
    if false {
        if true {
            var v7 bool = true
            _ = v7
        }
        var v8 uint8 = 1
        var v9 float64 = float64(3.14159)
        _ = v8
        _ = v9
    } else {
        var v10 uint32 = 0
        v11 := []int{-1}
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 155, F1: (p0 - (p0 % ((p0) | 1))), F2: 0}
    for fi13 := 0; fi13 < 12; fi13++ {
        if ((float64(1e10) > 0.0) || (!((true || false) || (p0 >= (p0 - p0))))) {
            var v14 uint32 = 220
            fmt.Println(p1)
            var v15 uint = 255
            _ = v14
            _ = v15
        }
        var v16 uint64 = 966
        _ = fi13
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    return 0
}
func genFn1() int8 {
    v17 := []int{1, 32767}
    v18 := GenS1{F0: 1, F1: 2, F2: -146}
    var v19 uint64 = 2
    if (42 >= 579) {
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 int8 = 1
            if false {
                var v22 float32 = 1e10
                var v23 int8 = -128
                fmt.Println(fi20, v22, v23)
                _ = v22
                _ = v23
            } else {
                v24 := GenS1{F0: 65535, F1: 127, F2: 0}
                v19 -= (^(v19))
                _ = v24
            }
            _ = fi20
            _ = v21
        }
    }
    _ = v17
    _ = v18
    _ = v19
    return -128
}

func main() {
    var v1 string = "a"
    v2 := []uint32{4294967295, 2}
    v3 := make(map[string]bool)
    var v4 uint32 = 42
    var v5 float32 = 0.0
    for fi25 := 0; fi25 < 19; fi25++ {
        if true {
            var v26 bool = false
            var v27 uint32 = v4
            _ = v26
            _ = v27
        } else {
            v28 := GenS2{F0: 714, F1: -1}
            _ = v28
        }
        if (!(false || (!false))) {
            v29 := GenS3{F0: (-(fi25))}
            var v30 uint = 657
            switch 1 {
            case 32767:
                if len(v2) > 0 {
                    v2[0] = v4
                }
            case 42:
                for fi31 := 0; fi31 < 7; fi31++ {
                    switch (v1 + v1) {
                    case "foo bar":
                        switch v1 {
                        case "foo bar":
                            _ = len(v3)
                        case "a":
                            if (false || true) {
                                var v32 uint16 = 0
                                _ = v32
                            }
                        default:
                            fmt.Println(v5, fi25, v4)
                        }
                    case "\n\t":
                        if false {
                            v5 -= float32(float64(((v5 - v5) / 3.14159)))
                        }
                        switch v1 {
                        case "":
                            _ = len(v2)
                        case "hello":
                            var v33 int = (fi31 ^ 1)
                            _ = v33
                        }
                        v34 := GenS2{F0: 788, F1: 0}
                        _ = v34
                    }
                    v30 = v30
                    delete(v3, (v1 + (((("\n\t" + (v1 + "foo bar")) + v1) + v1) + v1)))
                    _ = fi31
                }
            case -498:
                fmt.Println(v5)
            }
            _ = v29
            _ = v30
        } else {
            var v35 int64 = -1
            _ = v35
        }
        v2 = append(v2, (v4 / ((133) | 1)))
        _ = fi25
    }
    v36 := GenS3{F0: 128}
    var v37 uint32 = (^(v4))
    var v38 float64 = (3.14159 / -2.25)
    v39 := GenS4{F0: false, F1: 1, F2: 910}
    switch v1 {
    case "hello":
        for fi40 := 0; fi40 < 16; fi40++ {
            v41 := GenS5{F0: -1, F1: v1}
            switch (-(fi40)) {
            case 382:
                _ = len(v2)
            }
            _ = fi40
            _ = v41
        }
    case "":
        v38 += v38
    }
    v4 += 65535
    v5 = (1e10 * 0.0)
    var v42 int8 = 73
    if ((100 == 7) || (9223372036854775807 == 1)) {
        fmt.Println(v37, v38, v37)
        v4 = v37
        var v43 int8 = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
}
