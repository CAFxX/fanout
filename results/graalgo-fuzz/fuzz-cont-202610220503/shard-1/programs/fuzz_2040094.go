package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 uint8 }
func genFn0(p0 uint, p1 uint16, p2 int) uint16 {
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: "\n\t", F1: 948, F2: 42}
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS1{F0: "", F1: p1, F2: 922}
            _ = fi9
            _ = v10
        }
        var v11 int8 = 42
        _ = fi7
        _ = v11
        _ = v8
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        p2 = p2
        if false {
            switch p2 {
            case 32767:
                if (186 != uint8(p1)) {
                    v13 := []int{}
                    _ = v13
                } else {
                    v14 := []int{}
                    p2 = fi12
                    _ = v14
                }
            case 42:
                p2 += 0
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                if ((127 >= 237) && (-2.25 == 1e10)) {
                    switch ("\n\t" + ("hello" + "a")) {
                    case "":
                        var v16 int64 = 1
                        v17 := []int{fi15, -213, fi12}
                        if false {
                            fmt.Println(p2, p0, p2)
                            p1 = p1
                            p2 += p2
                        }
                        _ = v16
                        _ = v17
                    }
                    switch ("\n\t" + ("\n\t" + "hello")) {
                    case "":
                        var v18 uint32 = 0
                        _ = v18
                    case "\n\t":
                        fmt.Println(p0)
                    case "foo bar":
                        v19 := GenS2{F0: 4294967295}
                        _ = v19
                    default:
                        switch p2 {
                        case 128:
                            var v20 int16 = -62
                            _ = v20
                        default:
                            if (true || ((0 >= -9223372036854775808) || ((!((!false) || true)) && true))) {
                                v21 := GenS1{F0: "\n\t", F1: p1, F2: -2147483648}
                                _ = v21
                            }
                        }
                    }
                    p2 -= 498
                } else {
                    p1 = p1
                    if (4294967295 == 42) {
                        var v22 int8 = 42
                        v23 := GenS2{F0: -1}
                        _ = v22
                        _ = v23
                    }
                    fmt.Println(p0, fi15, fi15)
                    v24 := []int{(-(fi15))}
                    var v25 int32 = 1
                    _ = v24
                    _ = v25
                }
                _ = fi15
            }
        } else {
            if (true || false) {
                switch (p2 % ((p2) | 1)) {
                case 2147483647:
                    switch (fi12 ^ 2147483647) {
                    case 340:
                        var v26 uint8 = 1
                        _ = v26
                    default:
                        v27 := GenS3{F0: 698}
                        _ = v27
                    }
                default:
                    var v28 uint = p0
                    _ = v28
                }
                p1 = p1
            } else {
                p2 -= (p2 << ((0) & 63))
                for fi29 := 0; fi29 < 12; fi29++ {
                    var v30 int8 = 42
                    v31 := []int{(^(fi29)), fi29}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 % (((p1 % ((p1) | 1))) | 1))
}

func main() {
    var v1 uint64 = 2
    var v2 bool = (1 == 37)
    v3 := []int{32767, 410, 128}
    var v4 uint64 = (v1 | (v1 / ((42) | 1)))
    var v5 bool = v2
    var v6 uint = 255
    v4 += v4
    var v32 int64 = 1
    fmt.Println(v5, v2)
    var v33 float64 = 3.14159
    v34 := GenS1{F0: "a", F1: 255, F2: 42}
    v3 = append(v3, 486)
    fmt.Println(v32, v5)
    v1 -= 1
    v2 = v2
    var v35 bool = (2 < 616)
    var v36 int16 = 1
    if (v4 >= v1) {
        fmt.Println(v1, v2, v1)
    }
    if v5 {
        var v37 float64 = v33
        switch -927 {
        case 42:
            var v38 int = 128
            _ = v38
        default:
            _ = len(v3)
        }
        v36 += -1
        _ = v37
    }
    v39 := GenS4{F0: 160, F1: 74}
    if (1e10 < 0.1) {
        if v5 {
            var v40 int64 = (v32 / ((-9223372036854775808) | 1))
            _ = v40
        } else {
            fmt.Println(v1, v6, v4)
        }
        v41 := GenS1{F0: ("\n\t" + "a"), F1: 0, F2: 2147483647}
        _ = v41
    } else {
        var v42 uint16 = genFn0((v6 / ((42) | 1)), 65535, 2147483647)
        var v43 uint8 = 200
        _ = v42
        _ = v43
    }
    v33 += 1.5
    for fi44 := 0; fi44 < 17; fi44++ {
        for fi45 := 0; fi45 < 6; fi45++ {
            var v46 int32 = -2147483648
            _ = fi45
            _ = v46
        }
        _ = fi44
    }
    v47 := GenS2{F0: -754}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
