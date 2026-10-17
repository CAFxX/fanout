package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 string; F2 int }
type GenS4 struct { F0 uint8; F1 int; F2 int16 }
func genFn0(p0 string, p1 uint64) uint {
    var v5 int16 = 0
    switch (p0 + (("foo bar" + "foo bar") + (("\n\t" + "hello") + "a"))) {
    case "":
        if true {
            var v6 int64 = 42
            v6 = v6
            _ = v6
        } else {
            switch p0 {
            case "\n\t":
                for fi7 := 0; fi7 < 14; fi7++ {
                    var v8 bool = true
                    if v8 {
                        v8 = (123 < -1)
                        var v9 int64 = -1
                        _ = v9
                    } else {
                        v10 := []int{(fi7 & fi7), (fi7 / ((9223372036854775807) | 1))}
                        v11 := GenS1{F0: 2}
                        _ = v10
                        _ = v11
                    }
                    _ = fi7
                    _ = v8
                }
            default:
                fmt.Println(v5, p1)
            }
            p0 = p0
        }
    case "foo bar":
        v12 := []int{0, 809, -305, -891}
        _ = v12
    case "\n\t":
        var v13 float64 = 3.14159
        _ = v13
    }
    var v14 uint32 = 1
    _ = p0
    _ = p1
    _ = v14
    _ = v5
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    var v4 uint32 = 390
    _ = len(v2)
    var v15 uint8 = 31
    v16 := GenS1{F0: 2}
    v15 = 255
    for fi17 := 0; fi17 < 1; fi17++ {
        fmt.Println(fi17)
        fmt.Println(v15)
        v4 = (v4 >> ((746) & 63))
        _ = fi17
    }
    v4 += (v4 >> ((42) & 63))
    v18 := []int{-798}
    v4 = v4
    if (!(genFn0("\n\t", 1) >= 2)) {
        v15 = 127
    }
    delete(v3, 2147483647)
    if false {
        if (false || true) {
            v15 = v15
            switch 328 {
            case 42:
                if (false || false) {
                    v19 := GenS1{F0: 42}
                    v20 := GenS1{F0: 416}
                    _ = v19
                    _ = v20
                } else {
                    if true {
                        for fi21 := 0; fi21 < 6; fi21++ {
                            v22 := GenS1{F0: 255}
                            for fi23 := 0; fi23 < 6; fi23++ {
                                v15 += v15
                                fmt.Println(v4)
                                v4 -= 2147483647
                                _ = fi23
                            }
                            var v24 uint16 = 445
                            _ = fi21
                            _ = v22
                            _ = v24
                        }
                    }
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                }
            }
        }
        switch ("foo bar" + ("hello" + "hello")) {
        case "a":
            if false {
                for fi25 := 0; fi25 < 20; fi25++ {
                    if (false && (32767 > 42)) {
                        for fi26 := 0; fi26 < 11; fi26++ {
                            v15 += v15
                            _ = fi26
                        }
                    }
                    var v27 int = (fi25 ^ fi25)
                    _ = fi25
                    _ = v27
                }
                var v28 int = -40
                _ = v28
            } else {
                v29 := GenS1{F0: 301}
                if false {
                    if true {
                        var v30 float64 = 1.5
                        _ = v30
                    }
                    if true {
                        for fi31 := 0; fi31 < 13; fi31++ {
                            _ = len(v2)
                            _ = fi31
                        }
                    } else {
                        v4 = v4
                        switch 264 {
                        case 1:
                            v32 := GenS1{F0: 255}
                            _ = v32
                        }
                    }
                    for fi33 := 0; fi33 < 11; fi33++ {
                        var v34 uint = 4294967295
                        fmt.Println(v4, v15, v4)
                        v15 = v15
                        _ = fi33
                        _ = v34
                    }
                }
                _ = v29
            }
        case "\n\t":
            if false {
                v4 += v4
                var v35 bool = true
                _ = v35
            } else {
                var v36 float32 = 0.1
                _ = v36
            }
        case "foo bar":
            v4 = (v4 | v4)
        default:
            v37 := GenS2{F0: (v15 & (v15 * (v15 & v15)))}
            _ = v37
        }
        switch "hello" {
        case "\n\t":
            for fi38 := 0; fi38 < 12; fi38++ {
                var v39 string = "foo bar"
                for fi40 := 0; fi40 < 14; fi40++ {
                    switch (^(fi38)) {
                    case 32768:
                        v41 := GenS1{F0: 255}
                        _ = v41
                    case 128:
                        if (uint(v15) > 2) {
                            v42 := GenS3{F0: 981, F1: v39, F2: fi40}
                            if (true || (!((!true) && (true && true)))) {
                                v43 := GenS2{F0: (v15 * v15)}
                                var v44 uint8 = (^(v15))
                                var v45 int = fi38
                                _ = v43
                                _ = v44
                                _ = v45
                            } else {
                                var v46 float64 = 1.5
                                v39 = v39
                                if (42 <= 518) {
                                    fmt.Println(v39)
                                    _ = len(v18)
                                }
                                fmt.Println(v39)
                                _ = v46
                            }
                            _ = len(v3)
                            _ = v42
                        }
                        v2[v39] = 826
                        v39 = ("" + v39)
                    case 1:
                        v4 -= (v4 | 0)
                    }
                    var v47 uint = 424
                    v47 = 957
                    _ = fi40
                    _ = v47
                }
                _ = fi38
                _ = v39
            }
        }
    }
    v48 := GenS4{F0: v15, F1: 32768, F2: 32767}
    v15 = v15
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v48
}
