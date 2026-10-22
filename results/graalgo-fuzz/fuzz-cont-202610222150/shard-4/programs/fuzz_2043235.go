package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 string }
func genFn0(p0 uint8) string {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 65535, F1: 2147483647, F2: (fi4 >> ((4294967295) & 63))}
        p0 += p0
        _ = fi4
        _ = v5
    }
    p0 = 42
    p0 = (p0 & 127)
    v6 := GenS1{F0: 2, F1: 2, F2: 0}
    p0 -= p0
    _ = p0
    _ = v6
    return ("" + "a")
}
func genFn1(p0 uint64, p1 int, p2 uint64) string {
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS2{F0: 308}
        fmt.Println(p2, p1, p1)
        v9 := []int{(-(p1)), (p1 / ((fi7) | 1)), 263}
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        fmt.Println(p0, p2)
    } else {
        var v10 float64 = 0.1
        var v11 uint = 255
        _ = v10
        _ = v11
    }
    fmt.Println(p2, p2)
    if true {
        if false {
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 int = p1
                p1 -= (fi12 ^ (fi12 * (p1 | v13)))
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 uint16 = 32767
                    var v16 bool = (v15 >= v15)
                    p1 += fi12
                    fmt.Println(v16, v13, p0)
                    var v17 uint = 421
                    v18 := GenS3{F0: -124, F1: p0}
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = fi12
                _ = v13
            }
        } else {
            p0 = (p2 >> ((18446744073709551615) & 63))
            p2 -= p0
        }
    } else {
        for fi19 := 0; fi19 < 15; fi19++ {
            v20 := GenS1{F0: 65535, F1: 0, F2: fi19}
            p2 = p2
            switch (p1 | -1) {
            case -1:
                p1 = (-(p1))
            case 932:
                fmt.Println(p0, p1, p1)
            case 917:
                if false {
                    if (42 < 18446744073709551615) {
                        for fi21 := 0; fi21 < 20; fi21++ {
                            v22 := []int{p1, 0, (-(fi19)), p1}
                            switch (fi21 * fi21) {
                            case 32768:
                                if false {
                                    fmt.Println(p1, p2)
                                    p1 -= fi21
                                } else {
                                    var v23 uint16 = 255
                                    if len(v22) > 0 {
                                        v22[0] = fi19
                                    }
                                    var v24 float64 = (1e10 + (-2.25 + (float64((0.0 - (1e10 * 3.14159))) + float64((float32(1.5) - ((float32(1.5) + float32(((-2.25 / 0.1) * 1e10))) - 0.0))))))
                                    _ = v23
                                    _ = v24
                                }
                            }
                            _ = fi21
                            _ = v22
                        }
                        var v25 string = ""
                        var v26 string = v25
                        _ = v25
                        _ = v26
                    }
                    var v27 bool = false
                    var v28 int16 = 977
                    _ = v27
                    _ = v28
                }
            }
            _ = fi19
            _ = v20
        }
        p2 = 2
    }
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    v1 := []string{(("" + "") + ""), ("a" + "foo bar"), ("hello" + "foo bar")}
    var v2 uint32 = 753
    v3 := []uint32{(v2 * 4294967295), (^(v2))}
    var v29 bool = false
    switch 733 {
    case 4294967295:
        v30 := GenS4{F0: 100}
        _ = v30
    case -1:
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
    case 128:
        fmt.Println(v2, v29)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v29 = v29
    v31 := GenS1{F0: 0, F1: v2, F2: 1}
    if false {
        if (42 <= -2147483648) {
            v3 = append(v3, (v2 / ((v2) | 1)))
        } else {
            var v32 float64 = 1.5
            if v29 {
                v33 := GenS2{F0: 768}
                _ = v33
            } else {
                var v34 int32 = 349
                _ = v34
            }
            _ = v32
        }
        v3 = append(v3, 2)
        v29 = false
    } else {
        var v35 uint = 0
        v35 -= (v35 % ((4294967295) | 1))
        _ = v35
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v36 float64 = 0.0
    v37 := GenS5{F0: "a", F1: ("a" + "\n\t")}
    fmt.Println(v29, v29, v36)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v36
    _ = v37
}
