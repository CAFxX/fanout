package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 int8, p1 int32) int16 {
    v5 := GenS1{F0: (!(0 > 673))}
    for fi6 := 0; fi6 < 9; fi6++ {
        if true {
            if (fi6 <= fi6) {
                v7 := GenS2{F0: p0, F1: 2147483647, F2: -1}
                var v8 bool = ((^(fi6)) >= (^(fi6)))
                if true {
                    var v9 bool = (235 < 146)
                    if v9 {
                        for fi10 := 0; fi10 < 9; fi10++ {
                            var v11 string = "a"
                            var v12 int = (fi6 + -571)
                            if (-2.25 != 0.0) {
                                v13 := []int{}
                                _ = v13
                            }
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                    }
                    for fi14 := 0; fi14 < 7; fi14++ {
                        var v15 int32 = -2147483648
                        v16 := []int{(^(fi6)), int(p1), -9223372036854775808}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    _ = v9
                } else {
                    v17 := GenS2{F0: p0, F1: 65535, F2: 0}
                    for fi18 := 0; fi18 < 5; fi18++ {
                        v19 := GenS1{F0: (168 <= 228)}
                        _ = fi18
                        _ = v19
                    }
                    _ = v17
                }
                _ = v7
                _ = v8
            } else {
                if (32767 != 255) {
                    var v20 uint64 = 42
                    var v21 int16 = 32767
                    for fi22 := 0; fi22 < 15; fi22++ {
                        v23 := GenS1{F0: (!false)}
                        v24 := GenS1{F0: (!(p0 > -1))}
                        var v25 uint = 0
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    _ = v20
                    _ = v21
                } else {
                    var v26 float32 = 0.0
                    p1 += (p1 ^ (p1 - -2147483648))
                    _ = v26
                }
                var v27 uint16 = 0
                _ = v27
            }
        } else {
            var v28 bool = true
            _ = v28
        }
        fmt.Println(p0, p1)
        fmt.Println(fi6, p1)
        _ = fi6
    }
    v29 := []int{}
    _ = p0
    _ = p1
    _ = v29
    _ = v5
    return -1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    var v3 float32 = 3.14159
    var v4 string = ""
    var v30 int64 = -967
    var v31 int64 = v30
    v4 = "hello"
    for fi32 := 0; fi32 < 10; fi32++ {
        fmt.Println(v3)
        v33 := v1[128]
        v4 = v4
        _ = fi32
        _ = v33
    }
    v34 := []int{260, 128, 568, 42}
    v35 := GenS2{F0: -114, F1: 2147483647, F2: (-(v31))}
    var v36 uint16 = 65535
    v36 = v36
    if (!true) {
        v34 = append(v34, -2147483648)
        _ = len(v2)
        var v37 uint16 = (^(v36))
        _ = v37
    } else {
        if (94 <= 125) {
            v30 = v31
            v36 = 32767
            v4 = (("foo bar" + v4) + v4)
        }
    }
    v34 = append(v34, 9223372036854775807)
    for fi38 := 0; fi38 < 18; fi38++ {
        v3 = (v3 / v3)
        _ = fi38
    }
    switch v4 {
    case "foo bar":
        fmt.Println(v31)
    case "\n\t":
        for fi39 := 0; fi39 < 18; fi39++ {
            v30 = (v31 & v31)
            _ = fi39
        }
    case "hello":
        fmt.Println(v4)
    }
    fmt.Println(v36, v36)
    if len(v34) >= 2 {
        v34 = v34[1:len(v34)]
    }
    for fi40 := 0; fi40 < 4; fi40++ {
        for fi41 := 0; fi41 < 5; fi41++ {
            if true {
                v42 := GenS3{F0: 128, F1: (!((fi41 / ((fi41) | 1)) == (fi40 << ((42) & 63))))}
                var v43 int64 = v31
                _ = v42
                _ = v43
            } else {
                var v44 uint8 = 178
                var v45 int64 = (v30 % ((v31) | 1))
                _ = v44
                _ = v45
            }
            v31 = 9223372036854775807
            var v46 int8 = 0
            _ = fi41
            _ = v46
        }
        var v47 int = (^(fi40))
        switch fi40 {
        case 42:
            v3 += (v3 / (v3 + v3))
        case 1:
            if len(v34) > 0 {
                v34[0] = -508
            }
        case -1:
            for fi48 := 0; fi48 < 12; fi48++ {
                switch fi48 {
                case 42:
                    if true {
                        var v49 uint8 = 60
                        _ = v49
                    }
                }
                _ = fi48
            }
        }
        _ = fi40
        _ = v47
    }
    fmt.Println(v30)
    switch v4 {
    case "foo bar":
        v31 -= (v31 >> ((857) & 63))
    case "\n\t":
        var v50 int16 = 42
        _ = v50
    case "hello":
        var v51 int8 = 100
        _ = v51
    default:
        v52 := GenS4{F0: ((v4 + ("hello" + v4)) + "hello"), F1: 11}
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
