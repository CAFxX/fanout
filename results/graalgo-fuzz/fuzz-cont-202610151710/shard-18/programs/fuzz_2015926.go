package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 string; F2 int64 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint8; F1 int64; F2 int64 }
func genFn0(p0 int16, p1 int, p2 int8) uint32 {
    p0 -= 920
    for fi6 := 0; fi6 < 20; fi6++ {
        switch (^(p1)) {
        case 903:
            v7 := []int{}
            _ = v7
        case 128:
            for fi8 := 0; fi8 < 18; fi8++ {
                p2 = p2
                var v9 string = ("foo bar" + "hello")
                v10 := []int{fi8, p1, (p1 | 127), 0}
                _ = fi8
                _ = v10
                _ = v9
            }
        case -2147483648:
            var v11 int8 = p2
            _ = v11
        default:
            v12 := GenS1{F0: 0, F1: false}
            _ = v12
        }
        _ = fi6
    }
    var v13 int16 = (p0 + p0)
    fmt.Println(p2, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    v3 := []int{}
    var v4 int64 = -592
    var v5 uint64 = 687
    _ = len(v3)
    v1[uint32(v5)] = ""
    v4 -= v4
    v14 := GenS1{F0: v4, F1: false}
    if (-46 == 100) {
        var v15 int8 = 42
        _ = len(v1)
        for fi16 := 0; fi16 < 2; fi16++ {
            if (1 <= 1) {
                if len(v3) > 0 {
                    v3[0] = (^(fi16))
                }
                switch (("hello" + "a") + "") {
                case "\n\t":
                    if (112 == 0) {
                        var v17 string = (("" + "foo bar") + "foo bar")
                        _ = v17
                    }
                case "":
                    var v18 uint32 = 2147483648
                    _ = v18
                case "foo bar":
                    switch "\n\t" {
                    case "":
                        v3 = append(v3, fi16)
                    case "\n\t":
                        v19 := GenS1{F0: v4, F1: (!(32767 > 255))}
                        _ = v19
                    }
                default:
                    v15 = v15
                }
                if true {
                    var v20 int32 = 2147483647
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 bool = false
                        switch "hello" {
                        case "\n\t":
                            v20 -= 370
                        case "a":
                            var v23 uint16 = 65535
                            _ = v23
                        }
                        v24 := GenS2{F0: 0}
                        _ = fi21
                        _ = v22
                        _ = v24
                    }
                    fmt.Println(v5)
                    _ = v20
                }
            }
            if true {
                for fi25 := 0; fi25 < 9; fi25++ {
                    v4 += (v4 + (v4 / ((v4) | 1)))
                    v15 += 1
                    var v26 string = "\n\t"
                    if ((v15 << ((0) & 63)) < (-(v15))) {
                        v3 = append(v3, (fi16 + 32768))
                        var v27 int = (-(fi25))
                        v5 = 4294967295
                        var v28 int32 = -2147483648
                        _ = v27
                        _ = v28
                    } else {
                        v29 := GenS3{F0: fi25, F1: v26, F2: v4}
                        v30 := GenS4{F0: 0, F1: 0, F2: 0}
                        _ = v29
                        _ = v30
                    }
                    _ = fi25
                    _ = v26
                }
                v31 := GenS1{F0: (-(v4)), F1: false}
                v4 = (v4 << ((2) & 63))
                _ = v31
            } else {
                fmt.Println(v15, v4)
            }
            v4 += -1
            _ = fi16
        }
        _ = v15
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        var v33 float32 = float32(0.0)
        fmt.Println(v33, fi32, v4)
        var v34 uint = 1
        _ = fi32
        _ = v33
        _ = v34
    }
    if (!(2147483647 <= 2)) {
        var v35 uint = 4294967295
        _ = v35
    }
    var v36 int32 = 0
    switch "" {
    case "":
        var v37 string = ""
        _ = v37
    case "foo bar":
        if false {
            var v38 float32 = float32(3.14159)
            var v39 float64 = 0.0
            _ = v38
            _ = v39
        } else {
            delete(v1, 2147483647)
            var v40 int16 = 1000
            if false {
                for fi41 := 0; fi41 < 6; fi41++ {
                    if len(v3) > 0 {
                        v3[0] = (fi41 >> ((943) & 63))
                    }
                    v36 -= 42
                    _ = fi41
                }
            } else {
                var v42 string = "foo bar"
                _ = v42
            }
            _ = v40
        }
    case "hello":
        var v43 uint16 = 255
        _ = v43
    }
    v44 := GenS5{F0: 0, F1: (v4 / ((v4) | 1)), F2: (v4 << ((uint(v5)) & 63))}
    v45 := GenS2{F0: -32768}
    var v46 int16 = 32767
    var v47 float32 = 1e10
    delete(v2, "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
