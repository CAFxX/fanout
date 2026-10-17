package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
func genFn0(p0 int, p1 uint16, p2 uint16) int8 {
    v4 := []int{p0, (p0 | p0)}
    if (-854 >= 9223372036854775807) {
        v5 := GenS1{F0: -9223372036854775808, F1: 977}
        switch 32768 {
        case 41:
            fmt.Println(p0)
        case 155:
            for fi6 := 0; fi6 < 14; fi6++ {
                var v7 string = "a"
                _ = fi6
                _ = v7
            }
        case 2147483647:
            fmt.Println(p0, p2, p2)
        }
        v8 := GenS1{F0: 9223372036854775807, F1: -976}
        _ = v5
        _ = v8
    } else {
        if false {
            v9 := GenS2{F0: (p1 - 714), F1: 0}
            _ = v9
        } else {
            v10 := GenS2{F0: (^(p2)), F1: 42}
            for fi11 := 0; fi11 < 7; fi11++ {
                p2 += 1000
                var v12 float32 = -2.25
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        v13 := GenS3{F0: -9223372036854775808}
        p1 += (p2 ^ 255)
        _ = v13
    }
    p0 = -674
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -72
}
func genFn1(p0 uint64, p1 int16) int64 {
    switch ("a" + ("" + ("a" + "a"))) {
    case "foo bar":
        for fi14 := 0; fi14 < 12; fi14++ {
            if (true && true) {
                p1 = (-(p1))
                fmt.Println(fi14, fi14, p1)
                fmt.Println(p0, p0)
            }
            _ = fi14
        }
    case "hello":
        var v15 string = "\n\t"
        _ = v15
    case "":
        var v16 int = -877
        _ = v16
    }
    var v17 float64 = 0.0
    var v18 uint8 = 1
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return -28
}

func main() {
    var v1 uint8 = 154
    v2 := make(map[int]bool)
    var v3 int64 = 1
    var v19 string = "hello"
    fmt.Println(v1)
    v20 := GenS1{F0: v3, F1: 277}
    var v21 uint32 = 647
    v1 = (v1 * 24)
    for fi22 := 0; fi22 < 1; fi22++ {
        delete(v2, (fi22 + fi22))
        if (!(1 <= 1)) {
            var v23 int32 = 42
            _ = len(v2)
            var v24 uint16 = 440
            _ = v23
            _ = v24
        }
        v25 := GenS1{F0: genFn1(2, 42), F1: 2147483647}
        _ = fi22
        _ = v25
    }
    if true {
        if ("foo bar" > (v19 + "\n\t")) {
            v26 := GenS1{F0: -9223372036854775808, F1: 0}
            if false {
                if false {
                    switch "" {
                    case "foo bar":
                        for fi27 := 0; fi27 < 3; fi27++ {
                            if true {
                                v28 := GenS3{F0: (v3 & (v3 * -1))}
                                switch v19 {
                                case "foo bar":
                                    for fi29 := 0; fi29 < 1; fi29++ {
                                        v30 := GenS1{F0: v3, F1: 1}
                                        _ = fi29
                                        _ = v30
                                    }
                                }
                                _ = v28
                            }
                            _ = fi27
                        }
                    case "a":
                        for fi31 := 0; fi31 < 14; fi31++ {
                            if true {
                                delete(v2, 4294967295)
                                var v32 uint32 = (^(v21))
                                _ = v32
                            }
                            if true {
                                switch fi31 {
                                case 127:
                                    var v33 int16 = -657
                                    _ = v33
                                case 4294967295:
                                    var v34 float32 = 0.0
                                    _ = v34
                                }
                                var v35 uint64 = 42
                                var v36 int8 = -1
                                _ = v35
                                _ = v36
                            }
                            _ = fi31
                        }
                    }
                    var v37 float64 = 0.0
                    v3 = 0
                    _ = v37
                }
                v2[42] = ((v1 | v1) < (v1 | v1))
                v38 := GenS1{F0: 42, F1: 625}
                _ = v38
            }
            _ = v26
        }
        var v39 int32 = 42
        _ = v39
    }
    var v40 float64 = 0.0
    v40 = v40
    var v41 uint = 1
    v19 = v19
    v21 = 42
    switch (v19 + "a") {
    case "a":
        v3 = int64(v1)
    case "":
        v42 := []int{128}
        _ = v42
    case "foo bar":
        var v43 float32 = 1.5
        _ = v43
    default:
        v19 = "a"
    }
    v44 := []int{177, 854, 954, 32767}
    v45 := GenS4{F0: (v41 + (v41 + v41))}
    v46 := GenS2{F0: 157, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v40
    _ = v41
    _ = v44
    _ = v45
    _ = v46
}
