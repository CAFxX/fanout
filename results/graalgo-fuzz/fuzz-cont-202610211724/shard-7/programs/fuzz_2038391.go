package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int }
func genFn0(p0 uint32, p1 int) uint {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 bool = (!false)
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 uint64) uint8 {
    v8 := []int{}
    if true {
        for fi9 := 0; fi9 < 7; fi9++ {
            if (1 <= 2147483647) {
                p0 -= (p0 >> ((2) & 63))
            } else {
                var v10 int64 = 0
                if true {
                    var v11 int32 = 2147483647
                    p0 = (p0 >> ((2) & 63))
                    _ = v11
                } else {
                    var v12 uint = 42
                    if false {
                        v10 += 0
                        var v13 int16 = 405
                        _ = v13
                    }
                    var v14 bool = (433 > -1)
                    _ = v12
                    _ = v14
                }
                _ = v10
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi9
        }
        p0 = (p0 & (p0 << ((255) & 63)))
    }
    _ = p0
    _ = v8
    return 128
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 42
    var v3 uint32 = 559
    var v4 int8 = -128
    v5 := make(map[string]bool)
    v15 := GenS1{F0: v4, F1: 32767, F2: 32767}
    v16 := GenS2{F0: 2}
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 uint8 = genFn1(v2)
        if false {
            var v19 float32 = (0.0 - (1e10 - 1.5))
            v20 := []int{}
            _ = v19
            _ = v20
        }
        var v21 int = 127
        _ = fi17
        _ = v18
        _ = v21
    }
    switch 0 {
    case 127:
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 bool = true
            _ = fi22
            _ = v23
        }
    case 128:
        for fi24 := 0; fi24 < 20; fi24++ {
            v25 := GenS1{F0: v4, F1: 32767, F2: 346}
            v4 -= (v4 >> ((843) & 63))
            _ = fi24
            _ = v25
        }
    }
    v3 -= v3
    v26 := []int{-914, 42}
    fmt.Println(v3)
    v27 := v5["hello"]
    switch ("a" + "") {
    case "hello":
        switch ("foo bar" + ("hello" + "hello")) {
        case "hello":
            fmt.Println(v4, v3, v27)
        case "a":
            _ = len(v1)
        case "":
            if true {
                switch "" {
                case "":
                    for fi28 := 0; fi28 < 8; fi28++ {
                        if (!(!(true || v27))) {
                            v27 = (true || v27)
                        }
                        var v29 bool = (v27 || v27)
                        _ = fi28
                        _ = v29
                    }
                }
                var v30 uint32 = (v3 - (v3 & 2))
                v31 := GenS2{F0: 376}
                _ = v30
                _ = v31
            }
        }
    }
    _ = len(v26)
    for fi32 := 0; fi32 < 14; fi32++ {
        if (((1e10 - 1.5) + float32(1.5)) < 0.0) {
            fmt.Println(v27, v3, v27)
            if false {
                var v33 float32 = (((((0.0 * 1.5) * (0.0 * 0.0)) - 1e10) / 1e10) * (0.1 * 3.14159))
                fmt.Println(fi32, v33, v2)
                _ = v33
            } else {
                if (v27 && v27) {
                    fmt.Println(v27, v3)
                    for fi34 := 0; fi34 < 8; fi34++ {
                        v35 := GenS1{F0: v4, F1: -751, F2: 399}
                        switch ("foo bar" + "foo bar") {
                        case "":
                            if (fi34 >= fi34) {
                                v36 := GenS2{F0: 255}
                                if v27 {
                                    v37 := GenS3{F0: 127, F1: (fi32 | fi34)}
                                    _ = v37
                                }
                                var v38 uint8 = 1
                                _ = v36
                                _ = v38
                            } else {
                                var v39 string = ("hello" + "\n\t")
                                fmt.Println(fi32, v4, fi32)
                                v26 = append(v26, 0)
                                _ = v39
                            }
                        case "\n\t":
                            var v40 uint8 = 255
                            _ = v40
                        }
                        if (652 > 655) {
                            v41 := GenS3{F0: 156, F1: fi34}
                            delete(v5, "")
                            _ = len(v26)
                            fmt.Println(v3, v27)
                            delete(v5, "\n\t")
                            _ = v41
                        } else {
                            var v42 uint64 = 2
                            fmt.Println(v27, v4, v27)
                            if v27 {
                                v43 := GenS1{F0: -99, F1: 616, F2: 255}
                                var v44 int32 = -1
                                var v45 string = ("foo bar" + "foo bar")
                                _ = v43
                                _ = v44
                                _ = v45
                            } else {
                                if v27 {
                                    if len(v26) >= 2 {
                                        v26 = v26[1:len(v26)]
                                    }
                                }
                            }
                            v27 = v27
                            _ = v42
                        }
                        _ = fi34
                        _ = v35
                    }
                    v3 = (v3 >> ((genFn0(v3, fi32)) & 63))
                }
            }
            for fi46 := 0; fi46 < 9; fi46++ {
                v47 := GenS1{F0: v4, F1: 247, F2: 42}
                if len(v26) > 0 {
                    v26[0] = int(v4)
                }
                if (4294967295 == (^(v2))) {
                    v2 = v2
                    if v27 {
                        v2 = v2
                    }
                } else {
                    fmt.Println(v4, v27, v2)
                }
                _ = fi46
                _ = v47
            }
        }
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
