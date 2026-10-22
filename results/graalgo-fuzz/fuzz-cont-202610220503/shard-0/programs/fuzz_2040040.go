package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int32; F2 int16 }
type GenS3 struct { F0 uint8; F1 int16; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 uint; F2 uint }
func genFn0(p0 string, p1 bool, p2 int64) uint16 {
    var v6 float64 = 3.14159
    p0 = p0
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int16 = 1
        p2 = p2
        _ = fi7
        _ = v8
    }
    p1 = (v6 != ((3.14159 + float64((3.14159 * 1e10))) - 0.0))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}
func genFn1(p0 uint16, p1 int32) int8 {
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := []int{(fi9 * 0), fi9}
            p1 = p1
            _ = fi10
            _ = v11
        }
        switch (fi9 - fi9) {
        case 42:
            fmt.Println(fi9, p0)
        case 0:
            var v12 int64 = -9223372036854775808
            _ = v12
        case 128:
            switch -197 {
            case 2147483647:
                v13 := GenS1{F0: 854}
                _ = v13
            case -9223372036854775808:
                for fi14 := 0; fi14 < 7; fi14++ {
                    fmt.Println(p1, fi14)
                    _ = fi14
                }
            case -2147483648:
                p0 -= 32767
            }
        default:
            if ((p1 & 966) != p1) {
                fmt.Println(p0, fi9, fi9)
                var v15 uint16 = p0
                fmt.Println(p0, v15)
                _ = v15
            }
        }
        fmt.Println(p1)
        _ = fi9
    }
    var v16 float32 = 1.5
    if (!(0 != -302)) {
        if false {
            var v17 uint32 = 2147483648
            v18 := GenS2{F0: 42, F1: -1, F2: 42}
            _ = v17
            _ = v18
        } else {
            v19 := GenS1{F0: 1}
            fmt.Println(p0, p0)
            _ = v19
        }
        fmt.Println(p1, p0, v16)
    } else {
        if true {
            var v20 int32 = (p1 * -1)
            if true {
                v21 := []int{32767, 42, 127, 945}
                _ = v21
            }
            var v22 float32 = (-2.25 - -2.25)
            _ = v20
            _ = v22
        }
        switch -1 {
        case 42:
            fmt.Println(p0, p1)
        case 1:
            p1 += 0
        }
    }
    var v23 uint = 42
    _ = p0
    _ = p1
    _ = v16
    _ = v23
    return -11
}

func main() {
    var v1 int = -170
    var v2 uint8 = 73
    var v3 uint16 = 32767
    v4 := make(map[int]bool)
    var v5 int16 = -32768
    v24 := GenS1{F0: 0}
    var v25 string = (("a" + "") + "foo bar")
    var v26 int8 = 16
    var v27 uint64 = 106
    var v28 uint32 = 0
    v29 := []int{int(v3), (v1 * 127), (v1 ^ 4294967295)}
    if true {
        var v30 uint = 4294967295
        _ = v30
    } else {
        v3 -= (v3 >> ((4294967295) & 63))
        var v31 bool = ((v27 < uint64(v3)) || (1000 >= v5))
        if v31 {
            var v32 int = (^(v1))
            v1 += (v1 << ((809) & 63))
            _ = v32
        } else {
            _ = len(v4)
            var v33 uint32 = (v28 ^ v28)
            _ = v33
        }
        _ = v31
    }
    var v34 int32 = 2147483647
    fmt.Println(v1, v28)
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    fmt.Println(v26, v34, v2)
    var v35 float32 = 3.14159
    v27 -= (v27 | uint64(v2))
    var v36 uint16 = (v3 + 255)
    if false {
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
        var v37 uint32 = (v28 * (v28 - (v28 * v28)))
        _ = v37
    }
    v35 -= v35
    for fi38 := 0; fi38 < 15; fi38++ {
        fmt.Println(v27, v25, v25)
        if ((!false) || false) {
            for fi39 := 0; fi39 < 5; fi39++ {
                if true {
                    v2 -= (v2 ^ (v2 | uint8(v36)))
                } else {
                    var v40 uint64 = v27
                    var v41 string = v25
                    switch "hello" {
                    case "foo bar":
                        v34 += (v34 + (v34 & v34))
                    case "a":
                        v27 = v27
                    default:
                        var v42 uint64 = (v27 - v27)
                        _ = v42
                    }
                    v43 := GenS3{F0: (^(v2)), F1: v5, F2: v41}
                    _ = len(v4)
                    _ = v40
                    _ = v41
                    _ = v43
                }
                v44 := GenS3{F0: v2, F1: (^(v5)), F2: ((v25 + v25) + v25)}
                if ((v25 == ((v25 + "") + "a")) && false) {
                    switch fi38 {
                    case 4294967295:
                        fmt.Println(v5)
                    case 630:
                        v25 = "foo bar"
                        v25 = "a"
                    case 128:
                        if ((v3 - genFn0(v25, (!true), 357)) >= (^(v36))) {
                            if true {
                                var v45 uint16 = v3
                                _ = v45
                            } else {
                                v3 = (v36 | uint16(v28))
                            }
                            v5 = v5
                            fmt.Println(v35, v36)
                            fmt.Println(v28, fi39)
                        }
                    }
                    v46 := GenS4{F0: v25}
                    v2 -= v2
                    _ = v46
                } else {
                    v47 := GenS5{F0: (^(v27)), F1: 42, F2: 4294967295}
                    switch 128 {
                    case -1:
                        var v48 uint64 = (v27 ^ v27)
                        _ = v48
                    }
                    var v49 bool = (-2.25 >= (0.1 * (-2.25 * (0.1 + (1e10 + (1e10 * (-2.25 * 0.0)))))))
                    _ = v47
                    _ = v49
                }
                var v50 int16 = (-(v5))
                _ = fi39
                _ = v44
                _ = v50
            }
        }
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
