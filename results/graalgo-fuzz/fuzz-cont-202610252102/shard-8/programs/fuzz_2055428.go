package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint; F2 int8 }
func genFn0(p0 uint, p1 uint64) uint {
    var v6 bool = ((!true) || false)
    for fi7 := 0; fi7 < 7; fi7++ {
        p0 += (p0 & p0)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return (p0 & p0)
}
func genFn1(p0 uint32, p1 uint8) uint8 {
    var v8 int8 = 1
    for fi9 := 0; fi9 < 1; fi9++ {
        switch 128 {
        case 32767:
            for fi10 := 0; fi10 < 3; fi10++ {
                if (true && true) {
                    if (-1 >= v8) {
                        p0 += (^(p0))
                    } else {
                        v11 := GenS1{F0: -1, F1: 2, F2: -1}
                        v12 := GenS1{F0: 183, F1: 0, F2: 1000}
                        switch ("hello" + "") {
                        case "foo bar":
                            if (!(2147483647 <= 0)) {
                                if false {
                                    v13 := GenS2{F0: 295}
                                    _ = v13
                                }
                                var v14 uint = 42
                                switch (fi9 + (^(fi10))) {
                                case -9223372036854775808:
                                    v15 := GenS2{F0: -805}
                                    _ = v15
                                case 545:
                                    v16 := []int{fi10, (fi10 + (fi10 | -545))}
                                    _ = v16
                                case 559:
                                    v17 := GenS1{F0: -2147483648, F1: 1, F2: -61}
                                    _ = v17
                                }
                                _ = v14
                            }
                        case "\n\t":
                            switch "a" {
                            case "hello":
                                v18 := []int{(fi9 & (fi10 / ((fi9) | 1))), fi9}
                                _ = v18
                            case "":
                                v19 := []int{(fi9 + fi9)}
                                _ = v19
                            case "foo bar":
                                v20 := GenS3{F0: -1, F1: genFn0(875, 42), F2: v8}
                                _ = v20
                            }
                        default:
                            v21 := []int{42, 9223372036854775807, fi9}
                            _ = v21
                        }
                        _ = v11
                        _ = v12
                    }
                } else {
                    p1 = p1
                }
                _ = fi10
            }
        }
        _ = fi9
    }
    if true {
        for fi22 := 0; fi22 < 10; fi22++ {
            v23 := []int{(fi22 & 9223372036854775807), (fi22 | (^(fi22))), (fi22 + (fi22 << ((255) & 63)))}
            if len(v23) > 0 {
                v23[0] = (fi22 / ((32767) | 1))
            }
            var v24 string = "hello"
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    p1 = (^(p1))
    v25 := []int{-805, -431}
    _ = p0
    _ = p1
    _ = v25
    _ = v8
    return (p1 & (p1 / ((255) | 1)))
}

func main() {
    v1 := []uint32{65535, 2147483647}
    v2 := make(map[int]bool)
    var v3 int8 = 42
    var v4 int64 = -1
    v5 := make(map[int]int)
    var v26 uint16 = 0
    if ((1e10 + 1.5) <= 3.14159) {
        if true {
            if false {
                if (!("hello" < "foo bar")) {
                    if true {
                        var v27 uint16 = 2
                        var v28 uint = 2
                        v27 += (v27 % (((v27 / ((v27) | 1))) | 1))
                        _ = v27
                        _ = v28
                    } else {
                        v29 := GenS3{F0: 0, F1: genFn0(0, 1), F2: v3}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v29
                    }
                }
            } else {
                var v30 int = 2
                _ = v30
            }
            var v31 int16 = 32767
            _ = v31
        }
        delete(v5, 9223372036854775807)
    }
    var v32 int32 = -2147483648
    var v33 float32 = float32((1e10 / 0.1))
    v34 := GenS1{F0: v32, F1: 128, F2: 42}
    if false {
        delete(v5, 0)
        switch ("" + "\n\t") {
        case "a":
            var v35 uint16 = 2
            _ = v35
        case "hello":
            if true {
                v33 = 0.0
            }
        case "\n\t":
            var v36 float64 = (1.5 * 0.0)
            _ = v36
        }
        _ = len(v1)
    } else {
        fmt.Println(v4)
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        var v38 bool = (!(float64(v33) >= 3.14159))
        _ = fi37
        _ = v38
    }
    _ = len(v5)
    for fi39 := 0; fi39 < 18; fi39++ {
        var v40 float64 = 3.14159
        _ = fi39
        _ = v40
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = (v3 % (((^(v3))) | 1))
    switch (("hello" + ("" + ("" + "\n\t"))) + "hello") {
    case "\n\t":
        for fi41 := 0; fi41 < 9; fi41++ {
            delete(v5, fi41)
            if ((("\n\t" + "a") + "\n\t") >= "foo bar") {
                _ = len(v1)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v3 += v3
                switch (fi41 & 0) {
                case 1:
                    v26 = v26
                case 42:
                    v4 -= v4
                default:
                    var v42 uint = 255
                    _ = v42
                }
            }
            if (!(v32 <= 515)) {
                v26 = v26
                v3 = 14
            } else {
                var v43 uint64 = 42
                fmt.Println(v33)
                var v44 uint16 = 446
                _ = v43
                _ = v44
            }
            _ = fi41
        }
    }
    for fi45 := 0; fi45 < 17; fi45++ {
        fmt.Println(fi45)
        _ = fi45
    }
    v26 += 460
    v3 += (-(v3))
    _ = len(v1)
    v46 := GenS3{F0: (^(v4)), F1: 1, F2: (v3 % ((v3) | 1))}
    if (true || true) {
        v32 = 0
    } else {
        v26 = (v26 | v26)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v46
    _ = v5
}
