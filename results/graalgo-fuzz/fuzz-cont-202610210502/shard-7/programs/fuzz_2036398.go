package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8 }
func genFn0(p0 bool, p1 int16, p2 int8) int32 {
    var v6 uint = 0
    for fi7 := 0; fi7 < 15; fi7++ {
        switch ("hello" + "hello") {
        case "":
            v8 := []int{}
            _ = v8
        case "hello":
            p2 = 0
        default:
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 int64 = -157
                p2 -= -115
                if ((^(v6)) >= (v6 | 550)) {
                    fmt.Println(p1)
                    fmt.Println(p2)
                    v10 = v10
                } else {
                    switch ((("foo bar" + "") + (("foo bar" + "foo bar") + "a")) + "") {
                    case "hello":
                        if (255 <= 181) {
                            v11 := GenS1{F0: 127, F1: v6}
                            switch (fi7 / ((-547) | 1)) {
                            case 42:
                                v12 := []int{fi9, 0, fi9}
                                if true {
                                    v6 = v6
                                } else {
                                    var v13 uint32 = 4294967295
                                    var v14 int64 = v10
                                    p1 = p1
                                    _ = v13
                                    _ = v14
                                }
                                var v15 uint16 = 495
                                var v16 float64 = ((((1e10 - (0.0 * float64(-2.25))) / (-2.25 / float64(-2.25))) + (float64(0.1) - 1.5)) - 0.0)
                                _ = v12
                                _ = v15
                                _ = v16
                            case 32768:
                                if (("a" + "") < "a") {
                                    if true {
                                        v6 = v6
                                        var v17 int = (fi7 % ((fi9) | 1))
                                        _ = v17
                                    }
                                }
                            case 1:
                                var v18 bool = (2 <= 0)
                                _ = v18
                            }
                            v19 := GenS1{F0: (p2 & (p2 << ((v6) & 63))), F1: 0}
                            var v20 int8 = p2
                            v21 := GenS2{F0: 4294967295, F1: ("foo bar" + (("foo bar" + "a") + "hello"))}
                            _ = v11
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    case "":
                        v22 := []int{fi7}
                        _ = v22
                    default:
                        p2 = p2
                    }
                    fmt.Println(fi7, p0)
                    v23 := GenS3{F0: -1}
                    _ = v23
                }
                _ = fi9
                _ = v10
            }
        }
        v24 := GenS1{F0: p2, F1: (^(v6))}
        p1 += p1
        _ = fi7
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -452
}

func main() {
    var v1 int16 = 1
    v2 := make(map[int]uint32)
    var v3 int = -457
    v4 := make(map[int]uint32)
    v5 := []int{(v3 << ((309) & 63))}
    var v25 string = ""
    for fi26 := 0; fi26 < 16; fi26++ {
        v3 = v3
        if (!(1 == 0)) {
            if true {
                var v27 int32 = -815
                v28 := GenS4{F0: 0, F1: 91, F2: 65535}
                _ = v27
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 int64 = 1
            _ = fi29
            _ = v30
        }
        _ = fi26
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        v2[v3] = 4294967295
        _ = fi31
    }
    v32 := GenS3{F0: -1}
    fmt.Println(v3, v1)
    var v33 int32 = 2147483647
    v1 = (v1 % ((495) | 1))
    var v34 float64 = float64((0.0 + 3.14159))
    v35 := GenS5{F0: v3}
    for fi36 := 0; fi36 < 15; fi36++ {
        if true {
            fmt.Println(fi36)
            v34 -= float64((1.5 + (0.0 + 3.14159)))
        }
        _ = fi36
    }
    _ = len(v4)
    v34 = float64(3.14159)
    v37 := GenS5{F0: v3}
    for fi38 := 0; fi38 < 17; fi38++ {
        fmt.Println(v33, v33, v34)
        v39 := GenS6{F0: 57}
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
}
