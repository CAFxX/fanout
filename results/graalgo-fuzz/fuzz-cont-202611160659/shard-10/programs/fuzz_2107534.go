package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int64; F2 uint8 }
type GenS4 struct { F0 int64; F1 int8 }
func genFn0(p0 uint8) uint64 {
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: p0}
        for fi9 := 0; fi9 < 8; fi9++ {
            fmt.Println(fi7)
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    p0 -= (p0 * (p0 % ((p0) | 1)))
    p0 = (p0 * (p0 & 128))
    switch "hello" {
    case "a":
        var v10 int16 = -1
        _ = v10
    case "":
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := []int{fi11}
            var v13 uint8 = 42
            var v14 uint16 = 2
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    case "foo bar":
        var v15 int8 = -110
        _ = v15
    default:
        var v16 uint8 = 207
        _ = v16
    }
    p0 = (p0 * (^(p0)))
    _ = p0
    return 18446744073709551615
}
func genFn1(p0 uint16, p1 uint64) uint64 {
    v17 := []int{128, -9223372036854775808, -1, 4294967295}
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    if (-1 <= 2147483647) {
        p1 = (^(p1))
    } else {
        switch 42 {
        case 32768:
            for fi18 := 0; fi18 < 2; fi18++ {
                for fi19 := 0; fi19 < 13; fi19++ {
                    if len(v17) > 0 {
                        v17[0] = (fi18 & 813)
                    }
                    if (float32(-2.25) <= 1e10) {
                        if (false && true) {
                            p0 = p0
                            v20 := GenS1{F0: 2}
                            _ = v20
                        } else {
                            p0 -= p0
                            v21 := GenS2{F0: ""}
                            switch (fi18 / ((fi18) | 1)) {
                            case -2147483648:
                                p1 = (p1 << ((255) & 63))
                            case 991:
                                v17 = append(v17, (fi19 + fi18))
                            default:
                                fmt.Println(fi18, p0, fi18)
                            }
                            v22 := GenS2{F0: "foo bar"}
                            _ = v21
                            _ = v22
                        }
                        if true {
                            switch (fi18 | (-(fi18))) {
                            case -34:
                                _ = len(v17)
                            case 1:
                                var v23 bool = false
                                _ = v23
                            }
                        } else {
                            var v24 int16 = -474
                            fmt.Println(fi19)
                            v25 := GenS3{F0: -1, F1: -1, F2: 206}
                            v24 -= (v24 | 779)
                            v26 := GenS1{F0: 2}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    } else {
                        var v27 uint8 = uint8(p1)
                        if len(v17) > 0 {
                            v17[0] = (fi19 / (((fi18 & fi19)) | 1))
                        }
                        if true {
                            var v28 int32 = 1
                            _ = v28
                        }
                        fmt.Println(p1, p1, fi19)
                        _ = v27
                    }
                    _ = fi19
                }
                fmt.Println(fi18)
                if true {
                    var v29 int8 = int8(fi18)
                    _ = v29
                }
                _ = fi18
            }
        case 128:
            if true {
                var v30 uint8 = 255
                var v31 int32 = 1
                fmt.Println(v30, v31)
                _ = v30
                _ = v31
            } else {
                p1 += genFn0(42)
            }
        }
        p1 += 59
    }
    var v32 uint16 = 42
    _ = p0
    _ = p1
    _ = v17
    _ = v32
    return p1
}

func main() {
    var v1 uint64 = 238
    var v2 int32 = 114
    v3 := make(map[uint32]string)
    var v4 int = 9223372036854775807
    var v5 uint32 = 937
    var v6 int8 = 42
    var v33 uint32 = 279
    var v34 uint16 = 255
    fmt.Println(v1)
    v35 := GenS2{F0: ((("hello" + "\n\t") + ("hello" + "a")) + "hello")}
    var v36 uint = 1
    var v37 int8 = 127
    for fi38 := 0; fi38 < 15; fi38++ {
        switch 1 {
        case -9223372036854775808:
            if false {
                fmt.Println(v6)
            } else {
                v5 += v33
                v36 -= (v36 | 42)
                var v39 int = fi38
                _ = v39
            }
        case 4294967295:
            switch 4294967295 {
            case 2147483647:
                for fi40 := 0; fi40 < 4; fi40++ {
                    _ = len(v3)
                    switch (v4 / (((^(fi40))) | 1)) {
                    case -514:
                        var v41 uint = v36
                        if ((^(v1)) != 18446744073709551615) {
                            v42 := v3[v5]
                            v36 = 4294967295
                            _ = v42
                        }
                        _ = v41
                    case 2147483647:
                        v34 = (^(v34))
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            v43 := []int{921}
                            _ = v43
                        }
                    }
                    fmt.Println(fi38, v1)
                    _ = fi40
                }
            }
        case -373:
            v44 := GenS1{F0: 1}
            _ = v44
        default:
            fmt.Println(v36)
        }
        for fi45 := 0; fi45 < 12; fi45++ {
            v5 = (v5 / ((v5) | 1))
            v46 := GenS4{F0: 9223372036854775807, F1: v6}
            v5 = 986
            _ = fi45
            _ = v46
        }
        _ = fi38
    }
    var v47 int32 = -312
    v48 := GenS2{F0: "hello"}
    if (42 == 491) {
        v34 = (^(v34))
        v49 := []int{-2147483648}
        fmt.Println(v6)
        _ = v49
    }
    if true {
        v33 = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
