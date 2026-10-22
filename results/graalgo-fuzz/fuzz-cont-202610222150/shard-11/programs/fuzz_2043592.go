package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8; F1 int64 }
func genFn0(p0 bool, p1 int8) int8 {
    if true {
        v6 := []int{-836}
        _ = v6
    }
    var v7 int16 = 32767
    if true {
        v7 = v7
        p1 = -71
    } else {
        p0 = p0
        for fi8 := 0; fi8 < 5; fi8++ {
            if false {
                v9 := GenS1{F0: ((v7 ^ -900) > (v7 & -63)), F1: int32(p1), F2: 42}
                p0 = p0
                v10 := GenS2{F0: (v7 * -359)}
                _ = v10
                _ = v9
            }
            fmt.Println(v7, fi8)
            _ = fi8
        }
    }
    var v11 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return (p1 & p1)
}

func main() {
    v1 := []string{("a" + "\n\t")}
    v2 := []int{}
    v3 := make(map[string]string)
    var v4 int32 = 2147483647
    var v5 float32 = 0.0
    var v12 uint = 1
    if (-345 >= 32767) {
        fmt.Println(v12, v5)
    }
    v1 = append(v1, "foo bar")
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v4 -= 133
        if (42 >= 843) {
            v4 += 0
            var v14 uint32 = 2
            _ = v14
        } else {
            var v15 uint8 = 42
            var v16 int8 = 101
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    fmt.Println(v5)
    var v17 int16 = 42
    if true {
        v17 = 0
        v18 := GenS2{F0: v17}
        v12 += (v12 & (v12 | 42))
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        switch ("" + "a") {
        case "hello":
            v20 := GenS3{F0: 38}
            _ = v20
        case "":
            switch 32767 {
            case 331:
                var v21 string = ("foo bar" + "a")
                _ = v21
            default:
                if (!true) {
                    v4 -= v4
                } else {
                    for fi22 := 0; fi22 < 1; fi22++ {
                        var v23 uint16 = 1
                        switch (fi19 + (fi19 >> (((v12 % ((v12) | 1))) & 63))) {
                        case 0:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        v24 := GenS3{F0: 455}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 3; fi25++ {
                        v12 -= v12
                        if true {
                            v12 += 1
                            if true {
                                fmt.Println(fi25)
                                v26 := GenS2{F0: v17}
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                                _ = v26
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v27 string = "hello"
                                _ = v27
                            }
                            var v28 int8 = -82
                            _ = v28
                        }
                        v4 -= (v4 * v4)
                        if (float64(((v5 / 1.5) * v5)) > 0.0) {
                            v12 += (v12 >> (((^(v12))) & 63))
                            switch "" {
                            case "a":
                                v29 := GenS4{F0: 2147483647, F1: 300}
                                _ = v29
                            default:
                                var v30 float64 = -2.25
                                _ = v30
                            }
                            switch "foo bar" {
                            case "\n\t":
                                v31 := GenS5{F0: false, F1: 1}
                                _ = v31
                            case "a":
                                var v32 int64 = -1
                                if (true || ((false || (!(0 <= 65535))) || true)) {
                                    v33 := GenS5{F0: (1 != 639), F1: (-(v4))}
                                    fmt.Println(fi25, v12, v4)
                                    _ = v33
                                } else {
                                    v32 -= v32
                                    fmt.Println(fi19, v32)
                                    var v34 int16 = (v17 + v17)
                                    _ = v34
                                }
                                _ = len(v3)
                                _ = v32
                            case "":
                                v35 := GenS1{F0: (true || (((!(1 > 42)) && true) || true)), F1: -793, F2: v12}
                                _ = v35
                            }
                            var v36 float64 = 1e10
                            _ = v36
                        }
                        _ = fi25
                    }
                }
            }
        }
        _ = fi19
    }
    if true {
        if false {
            for fi37 := 0; fi37 < 8; fi37++ {
                for fi38 := 0; fi38 < 16; fi38++ {
                    if (128 != 2) {
                        v5 -= 0.0
                        fmt.Println(v4, v12)
                        v17 = (v17 << ((2) & 63))
                    }
                    _ = fi38
                }
                for fi39 := 0; fi39 < 10; fi39++ {
                    v2 = append(v2, (fi39 * 127))
                    v40 := GenS3{F0: v12}
                    _ = fi39
                    _ = v40
                }
                for fi41 := 0; fi41 < 9; fi41++ {
                    v12 -= v12
                    v12 = (^(v12))
                    _ = fi41
                }
                _ = fi37
            }
            fmt.Println(v5)
        } else {
            for fi42 := 0; fi42 < 3; fi42++ {
                v43 := GenS6{F0: 1}
                v12 -= (v12 << (((^(v12))) & 63))
                _ = fi42
                _ = v43
            }
            switch ("\n\t" + ("foo bar" + "")) {
            case "":
                v17 += v17
            default:
                v12 += v12
            }
        }
        if len(v2) > 0 {
            v2[0] = 42
        }
    }
    v5 = -2.25
    if (1000 >= 1000) {
        if (127 >= 1) {
            fmt.Println(v12)
        }
        v44 := GenS7{F0: 128, F1: -1}
        _ = v44
    } else {
        v5 -= (0.0 * v5)
        for fi45 := 0; fi45 < 13; fi45++ {
            for fi46 := 0; fi46 < 17; fi46++ {
                v12 = (v12 + (^(v12)))
                _ = fi46
            }
            if false {
                v17 = v17
            }
            v4 = v4
            _ = fi45
        }
        switch ("" + "a") {
        case "foo bar":
            var v47 float32 = (3.14159 / 3.14159)
            _ = v47
        case "\n\t":
            var v48 string = "a"
            _ = v48
        }
    }
    fmt.Println(v4, v4, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
