package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint }
type GenS4 struct { F0 int16; F1 uint; F2 int32 }
func genFn0() uint {
    switch 32768 {
    case 128:
        v7 := GenS1{F0: 0}
        _ = v7
    }
    var v8 int64 = -1
    _ = v8
    return 294
}
func genFn1(p0 int, p1 int, p2 int32) int32 {
    p2 += (p2 << ((52) & 63))
    if (-1 > 1) {
        switch "" {
        case "foo bar":
            v9 := GenS2{F0: ((((("hello" + "") + "a") + "a") + ("hello" + ("a" + "hello"))) + "hello"), F1: 176}
            _ = v9
        case "\n\t":
            fmt.Println(p2, p2)
        }
        if true {
            p0 -= 42
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 float32 = (1e10 + float32((float64(-2.25) / 1e10)))
                if true {
                    if (-726 > (p2 ^ p2)) {
                        for fi12 := 0; fi12 < 17; fi12++ {
                            var v13 int32 = (p2 << ((2) & 63))
                            switch "" {
                            case "foo bar":
                                var v14 int32 = 847
                                _ = v14
                            case "":
                                v15 := GenS3{F0: 2, F1: 0, F2: 96}
                                _ = v15
                            case "hello":
                                p1 += fi10
                            default:
                                p1 = (-(fi12))
                            }
                            var v16 float32 = 0.0
                            _ = fi12
                            _ = v13
                            _ = v16
                        }
                        for fi17 := 0; fi17 < 8; fi17++ {
                            p1 = (fi17 & 1)
                            _ = fi17
                        }
                    }
                    v18 := GenS4{F0: 800, F1: 0, F2: (p2 % ((p2) | 1))}
                    if (42 == 60) {
                        switch "\n\t" {
                        case "foo bar":
                            for fi19 := 0; fi19 < 8; fi19++ {
                                p2 = p2
                                _ = fi19
                            }
                        case "hello":
                            p1 = fi10
                        default:
                            p2 += (p2 - 1)
                        }
                    }
                    _ = v18
                }
                v20 := GenS4{F0: 42, F1: 65535, F2: p2}
                _ = fi10
                _ = v11
                _ = v20
            }
        } else {
            v21 := []int{9223372036854775807}
            if len(v21) > 0 {
                v21[0] = 9223372036854775807
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 float64 = -2.25
                fmt.Println(p1, p1)
                _ = fi22
                _ = v23
            }
            _ = v21
        }
        v24 := []int{32767}
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    v3 := []string{"hello", ("\n\t" + ("\n\t" + (((("\n\t" + "\n\t") + "\n\t") + "") + (("" + "hello") + "foo bar")))), "hello", "a"}
    v4 := []string{("\n\t" + "\n\t")}
    v5 := []string{}
    var v6 bool = true
    switch 951 {
    case 384:
        v6 = (42 < 42)
    case 0:
        v6 = v6
    case 128:
        v25 := v1[641]
        _ = v25
    default:
        if true {
            _ = len(v1)
            v26 := v1[2147483648]
            fmt.Println(v26)
            _ = v26
        }
    }
    var v27 string = "a"
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        fmt.Println(v27, v27, fi28)
        _ = fi28
    }
    if (-196 == 605) {
        _ = len(v1)
        var v29 string = v27
        var v30 int16 = 146
        _ = v29
        _ = v30
    } else {
        var v31 uint32 = 190
        _ = v31
    }
    if true {
        if (3.14159 >= (-2.25 * -2.25)) {
            fmt.Println(v27, v6)
        }
    }
    var v32 int32 = -2147483648
    var v33 uint16 = 32767
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        if len(v3) > 0 {
            v3[0] = "a"
        }
        _ = fi34
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi35
    }
    v27 = v27
    var v36 uint = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
