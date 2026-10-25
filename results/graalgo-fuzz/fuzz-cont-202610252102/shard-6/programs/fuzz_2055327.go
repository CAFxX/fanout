package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 bool }
func genFn0(p0 string) uint32 {
    var v6 uint16 = 255
    for fi7 := 0; fi7 < 3; fi7++ {
        switch (fi7 + fi7) {
        case 4294967295:
            if false {
                switch p0 {
                case "hello":
                    for fi8 := 0; fi8 < 10; fi8++ {
                        var v9 int = fi8
                        _ = fi8
                        _ = v9
                    }
                case "\n\t":
                    if (42 > 4294967295) {
                        v10 := GenS1{F0: 0, F1: -195, F2: "a"}
                        _ = v10
                    } else {
                        for fi11 := 0; fi11 < 2; fi11++ {
                            fmt.Println(v6)
                            _ = fi11
                        }
                    }
                case "foo bar":
                    v12 := GenS1{F0: uint(v6), F1: 42, F2: p0}
                    _ = v12
                default:
                    fmt.Println(fi7, p0, p0)
                }
                v13 := []int{42, fi7}
                _ = v13
            }
        case 533:
            switch 981 {
            case 127:
                v14 := GenS1{F0: 341, F1: 32767, F2: "\n\t"}
                _ = v14
            }
        default:
            var v15 bool = (float32(1.5) != 1e10)
            _ = v15
        }
        v16 := []int{fi7, (-(fi7)), (^(fi7)), 32768}
        _ = fi7
        _ = v16
    }
    v6 -= (v6 + v6)
    v17 := GenS1{F0: 605, F1: 0, F2: p0}
    switch 2147483647 {
    case -9223372036854775808:
        for fi18 := 0; fi18 < 9; fi18++ {
            v19 := GenS1{F0: 4294967295, F1: -275, F2: ((("" + "\n\t") + "a") + (p0 + ("\n\t" + (("foo bar" + ((("a" + ((p0 + p0) + "")) + "a") + (p0 + "foo bar"))) + p0))))}
            v20 := GenS2{F0: -1, F1: -97}
            p0 = ((((p0 + ((p0 + "") + ("hello" + p0))) + p0) + p0) + p0)
            _ = fi18
            _ = v19
            _ = v20
        }
    case -1:
        fmt.Println(v6, v6)
    case 63:
        p0 = p0
    default:
        v21 := GenS2{F0: 0, F1: -1}
        _ = v21
    }
    _ = p0
    _ = v17
    _ = v6
    return uint32(v6)
}

func main() {
    var v1 uint = 65535
    v2 := []int{4294967295, 32768, 535, 1}
    var v3 float64 = 0.1
    v4 := make(map[uint32]string)
    var v5 int8 = -93
    if len(v2) > 0 {
        v2[0] = 0
    }
    v1 = (v1 >> ((v1) & 63))
    var v22 float32 = 0.0
    var v23 float32 = 1.5
    if (float32((0.0 * 0.0)) > 0.1) {
        if (v23 <= v22) {
            fmt.Println(v3)
            v24 := GenS2{F0: 9223372036854775807, F1: 0}
            _ = v24
        } else {
            var v25 bool = (genFn0("foo bar") < 42)
            for fi26 := 0; fi26 < 3; fi26++ {
                switch fi26 {
                case -1:
                    v27 := GenS2{F0: 0, F1: 42}
                    _ = v27
                case 128:
                    if v25 {
                        v1 += 673
                        var v28 float64 = 3.14159
                        v25 = v25
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 16; fi29++ {
                            v22 = (0.1 - v22)
                            _ = len(v4)
                            _ = fi29
                        }
                        v30 := GenS3{F0: false}
                        _ = len(v2)
                        _ = v30
                    }
                case 1:
                    switch "a" {
                    case "hello":
                        for fi31 := 0; fi31 < 17; fi31++ {
                            var v32 float32 = float32(v3)
                            _ = fi31
                            _ = v32
                        }
                    }
                }
                var v33 string = "\n\t"
                _ = fi26
                _ = v33
            }
            _ = v25
        }
        v3 = ((v3 + 3.14159) - 0.0)
        v5 += v5
    }
    var v34 int16 = 32767
    var v35 uint = (v1 + v1)
    v2 = append(v2, 9223372036854775807)
    var v36 uint = (^(v1))
    var v37 uint = 65535
    v38 := GenS1{F0: v36, F1: 1000, F2: "foo bar"}
    var v39 bool = (!(v37 < v36))
    v2 = append(v2, 42)
    v34 = v34
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
