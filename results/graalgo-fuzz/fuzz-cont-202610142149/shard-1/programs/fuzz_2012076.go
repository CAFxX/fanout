package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int16; F2 bool }
type GenS5 struct { F0 uint16; F1 int64; F2 int }
func genFn0(p0 uint64, p1 string) int {
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            if false {
                p1 = ((p1 + (("\n\t" + "a") + ((p1 + "\n\t") + ""))) + (p1 + p1))
            } else {
                v8 := []int{(-(fi7)), 0, fi6}
                p1 = p1
                var v9 uint8 = 0
                var v10 bool = (!true)
                switch p1 {
                case "a":
                    v11 := GenS1{F0: 128, F1: 24, F2: v9}
                    _ = v11
                case "":
                    if v10 {
                        p1 = (p1 + p1)
                        switch fi7 {
                        case -200:
                            _ = len(v8)
                            switch fi7 {
                            case 128:
                                p0 -= uint64(v9)
                            default:
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            }
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        case 144:
                            _ = len(v8)
                        default:
                            v12 := GenS1{F0: v9, F1: 32767, F2: v9}
                            _ = v12
                        }
                    } else {
                        v13 := GenS2{F0: 0}
                        _ = v13
                    }
                case "\n\t":
                    var v14 int16 = 1
                    _ = v14
                default:
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
                _ = v10
                _ = v8
                _ = v9
            }
            switch (fi7 & 824) {
            case 4294967295:
                v15 := []int{}
                _ = v15
            case 151:
                p1 = (p1 + (p1 + ("foo bar" + p1)))
            default:
                var v16 int64 = 1
                _ = v16
            }
            _ = fi7
        }
        v17 := GenS2{F0: 100}
        _ = fi6
        _ = v17
    }
    var v18 uint8 = 42
    var v19 string = "foo bar"
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 982
}

func main() {
    var v1 uint8 = 42
    var v2 string = "a"
    var v3 uint = 65535
    v4 := make(map[string]int)
    var v5 int64 = -9223372036854775808
    if (!(2 == 42)) {
        var v20 int8 = 1
        v20 -= (v20 | (v20 * v20))
        v1 = 42
        _ = v20
    }
    switch 42 {
    case 19:
        var v21 float64 = ((0.1 * -2.25) * 1e10)
        _ = v21
    default:
        v22 := []int{718}
        _ = v22
    }
    v23 := GenS1{F0: (v1 - (v1 % ((v1) | 1))), F1: 1, F2: (v1 ^ (v1 << (((v3 | (^(v3)))) & 63)))}
    v24 := GenS3{F0: 782}
    v3 -= v3
    for fi25 := 0; fi25 < 3; fi25++ {
        v5 = (v5 / ((v5) | 1))
        switch "foo bar" {
        case "\n\t":
            if true {
                switch "foo bar" {
                case "a":
                    v2 = v2
                case "":
                    v26 := GenS4{F0: 255, F1: 1, F2: false}
                    _ = v26
                case "\n\t":
                    var v27 int8 = -1
                    _ = v27
                default:
                    for fi28 := 0; fi28 < 18; fi28++ {
                        if true {
                            v29 := GenS4{F0: 255, F1: 0, F2: true}
                            var v30 uint8 = (v1 / (((v1 * (v1 ^ (v1 - v1)))) | 1))
                            fmt.Println(fi28, fi28)
                            switch fi28 {
                            case 325:
                                v31 := []int{-2147483648, fi25, (fi28 + 1), (fi25 + fi28)}
                                _ = v31
                            case -9223372036854775808:
                                var v32 int64 = int64(v30)
                                _ = v32
                            case 128:
                                var v33 uint = 0
                                _ = v33
                            default:
                                var v34 uint = (v3 - 42)
                                _ = v34
                            }
                            _ = v29
                            _ = v30
                        } else {
                            v35 := GenS3{F0: 2147483648}
                            _ = v35
                        }
                        _ = fi28
                    }
                }
                v36 := []int{genFn0(0, ("" + "hello")), genFn0(787, "\n\t"), (^(fi25)), genFn0(0, v2)}
                _ = v36
            }
        }
        var v37 int8 = -1
        _ = fi25
        _ = v37
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        switch v2 {
        case "foo bar":
            for fi39 := 0; fi39 < 5; fi39++ {
                var v40 int = 9223372036854775807
                v41 := []int{42, (-(v40)), fi38}
                v1 = 1
                _ = fi39
                _ = v40
                _ = v41
            }
        case "hello":
            var v42 int = (fi38 ^ (fi38 & fi38))
            _ = v42
        case "\n\t":
            v2 = v2
        default:
            for fi43 := 0; fi43 < 10; fi43++ {
                v44 := GenS1{F0: 254, F1: 65535, F2: v1}
                var v45 uint64 = 870
                _ = fi43
                _ = v44
                _ = v45
            }
        }
        _ = fi38
    }
    v46 := GenS5{F0: 2, F1: (v5 >> ((1) & 63)), F2: 9223372036854775807}
    if (42 > 0) {
        switch ("foo bar" + "a") {
        case "hello":
            delete(v4, v2)
        case "\n\t":
            if false {
                var v47 int16 = 121
                v2 = v2
                _ = v47
            }
        case "foo bar":
            var v48 int8 = -128
            _ = v48
        }
    } else {
        v5 = v5
        if (false && true) {
            var v49 int = -9223372036854775808
            _ = v49
        }
        v1 += (^(v1))
    }
    var v50 float64 = 0.0
    var v51 float32 = 0.1
    v52 := []int{-9223372036854775808, -994, 32768, 9223372036854775807}
    var v53 int16 = 1000
    var v54 uint16 = 32767
    var v55 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
}
