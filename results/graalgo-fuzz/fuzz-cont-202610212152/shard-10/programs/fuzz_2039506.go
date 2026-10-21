package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 bool; F2 int64 }
type GenS5 struct { F0 uint8; F1 int16 }
func genFn0(p0 string, p1 uint8, p2 int) int16 {
    p2 += p2
    v6 := GenS1{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 32767
}
func genFn1() int {
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS2{F0: 781, F1: "a"}
        for fi9 := 0; fi9 < 2; fi9++ {
            if (42 == 128) {
                var v10 int = fi9
                _ = v10
            } else {
                fmt.Println(fi7, fi7)
                var v11 string = (("a" + "\n\t") + "\n\t")
                v12 := GenS1{F0: 4294967295}
                _ = v11
                _ = v12
            }
            _ = fi9
        }
        v13 := []int{(fi7 & fi7), fi7, 2147483647}
        _ = fi7
        _ = v13
        _ = v8
    }
    var v14 uint8 = 192
    _ = v14
    return -248
}

func main() {
    var v1 float64 = 0.1
    v2 := []int{-775}
    var v3 int8 = 33
    v4 := make(map[uint32]int)
    var v5 int8 = v3
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 string = "hello"
        _ = fi15
        _ = v16
    }
    v17 := GenS3{F0: 127}
    if true {
        v18 := GenS1{F0: -1}
        _ = v18
    }
    var v19 string = (((("hello" + "a") + (("hello" + "foo bar") + "")) + "") + "")
    v1 = v1
    if false {
        for fi20 := 0; fi20 < 1; fi20++ {
            if (!true) {
                fmt.Println(v19, v5)
            }
            var v21 float32 = 3.14159
            v22 := GenS2{F0: 4294967295, F1: (v19 + v19)}
            _ = fi20
            _ = v21
            _ = v22
        }
    } else {
        switch "hello" {
        case "\n\t":
            var v23 uint = 4294967295
            _ = v23
        case "":
            for fi24 := 0; fi24 < 3; fi24++ {
                v25 := v4[177]
                _ = fi24
                _ = v25
            }
        default:
            if (true || true) {
                if (!(0.0 != 1.5)) {
                    fmt.Println(v1, v3)
                } else {
                    if true {
                        v5 = v5
                    } else {
                        v19 = (("foo bar" + v19) + "foo bar")
                        for fi26 := 0; fi26 < 14; fi26++ {
                            var v27 int32 = -696
                            _ = fi26
                            _ = v27
                        }
                    }
                    v2 = append(v2, 4294967295)
                    if false {
                        if len(v2) > 0 {
                            v2[0] = 2147483647
                        }
                        v4[2147483647] = 32768
                    } else {
                        v4[4294967295] = -2147483648
                        v28 := v4[0]
                        _ = v28
                    }
                }
            }
        }
        if true {
            switch (v19 + "\n\t") {
            case "a":
                switch -923 {
                case -625:
                    v29 := GenS4{F0: 1, F1: (1000 < -32768), F2: 914}
                    _ = v29
                }
            case "foo bar":
                if false {
                    fmt.Println(v19, v1)
                }
            case "\n\t":
                v1 = 1.5
            default:
                v30 := v4[2]
                _ = v30
            }
        }
        switch -2147483648 {
        case -562:
            var v31 string = v19
            _ = v31
        default:
            delete(v4, 65535)
        }
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        for fi33 := 0; fi33 < 19; fi33++ {
            v5 = v5
            fmt.Println(v3, v3, v1)
            v34 := GenS5{F0: 1, F1: 32767}
            switch "" {
            case "hello":
                if true {
                    v1 = -2.25
                    if (!(42 != 0)) {
                        fmt.Println(fi33, v19, v19)
                        var v35 int8 = v5
                        _ = len(v2)
                        _ = v35
                    }
                }
                var v36 float32 = -2.25
                _ = v36
            default:
                var v37 float32 = 1e10
                _ = v37
            }
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    v38 := GenS1{F0: 128}
    if false {
        fmt.Println(v5, v5, v1)
    }
    if (3 < 101) {
        delete(v4, 780)
    }
    var v39 uint16 = 352
    var v40 int = -923
    v41 := GenS5{F0: 71, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
