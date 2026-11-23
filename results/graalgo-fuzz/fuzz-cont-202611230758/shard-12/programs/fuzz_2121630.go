package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 uint64; F2 int64 }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 int64; F1 uint; F2 int }
func genFn0(p0 int32) int {
    switch (("hello" + "a") + "foo bar") {
    case "foo bar":
        var v5 uint16 = 1
        _ = v5
    case "\n\t":
        var v6 string = (("" + "\n\t") + "a")
        _ = v6
    default:
        fmt.Println(p0)
    }
    var v7 float32 = float32(-2.25)
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := GenS1{F0: 0, F1: 1}
            if true {
                if true {
                    v11 := GenS1{F0: 2, F1: 1000}
                    _ = v11
                }
                if (true && false) {
                    v12 := GenS2{F0: -1, F1: 0, F2: 2}
                    var v13 uint32 = 65535
                    v14 := []int{fi8, fi9}
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    switch -1 {
                    case -925:
                        fmt.Println(fi9, p0, fi9)
                    case 128:
                        var v15 uint64 = 350
                        _ = v15
                    default:
                        var v16 int8 = -57
                        if (false && (!(4294967295 > 4294967295))) {
                            switch (fi8 & fi8) {
                            case -707:
                                switch "\n\t" {
                                case "a":
                                    v17 := GenS2{F0: int64(v16), F1: 42, F2: 403}
                                    _ = v17
                                default:
                                    v18 := []int{}
                                    _ = v18
                                }
                            case -533:
                                if false {
                                    v19 := GenS1{F0: 129, F1: -32768}
                                    switch fi8 {
                                    case 303:
                                        v20 := GenS1{F0: 127, F1: 1000}
                                        _ = v20
                                    case 1:
                                        p0 = 296
                                    case 4294967295:
                                        var v21 bool = (231 > 1)
                                        _ = v21
                                    }
                                    v22 := GenS3{F0: 0}
                                    if true {
                                        v23 := GenS4{F0: -500, F1: 42, F2: 9223372036854775807}
                                        v16 = (v16 & v16)
                                        if false {
                                            fmt.Println(p0, fi9, v16)
                                            v24 := GenS1{F0: 127, F1: -344}
                                            fmt.Println(fi8, v7)
                                            v25 := GenS5{F0: ("" != "a"), F1: (p0 / (((p0 / ((-302) | 1))) | 1))}
                                            _ = v24
                                            _ = v25
                                        } else {
                                            if true {
                                                fmt.Println(fi8, fi8)
                                                switch 0 {
                                                case 127:
                                                    v7 += v7
                                                case 128:
                                                    v16 -= v16
                                                }
                                            }
                                            v26 := []int{9223372036854775807}
                                            v16 = v16
                                            _ = v26
                                        }
                                        fmt.Println(fi8, p0, fi9)
                                        v27 := []int{(fi9 & 4294967295), -9223372036854775808}
                                        _ = v23
                                        _ = v27
                                    } else {
                                        v28 := GenS6{F0: p0, F1: 62}
                                        _ = v28
                                    }
                                    _ = v19
                                    _ = v22
                                } else {
                                    fmt.Println(fi9, fi9)
                                    var v29 float32 = 1.5
                                    var v30 float32 = 1e10
                                    _ = v29
                                    _ = v30
                                }
                            case 110:
                                v31 := []int{fi9}
                                _ = v31
                            default:
                                v16 -= 42
                            }
                        }
                        var v32 uint16 = 255
                        _ = v16
                        _ = v32
                    }
                    switch fi9 {
                    case 4294967295:
                        v33 := GenS7{F0: 255, F1: 4294967295}
                        _ = v33
                    }
                }
                v7 -= v7
            } else {
                switch ((("a" + ("" + "hello")) + "hello") + "\n\t") {
                case "a":
                    var v34 int = (fi8 + fi9)
                    _ = v34
                }
                var v35 float64 = 1.5
                _ = v35
            }
            _ = fi9
            _ = v10
        }
        p0 = p0
        v36 := GenS1{F0: 128, F1: 426}
        _ = fi8
        _ = v36
    }
    _ = p0
    _ = v7
    return -90
}

func main() {
    var v1 float64 = 1e10
    var v2 uint8 = 132
    v3 := []string{}
    var v4 int32 = 1
    var v37 int16 = 0
    var v38 uint = 42
    v37 = (v37 << (((^(v38))) & 63))
    if true {
        for fi39 := 0; fi39 < 17; fi39++ {
            v2 = v2
            v1 -= v1
            var v40 int64 = -1
            _ = fi39
            _ = v40
        }
        v4 += v4
        v41 := GenS8{F0: -906, F1: v38, F2: genFn0(0)}
        _ = v41
    } else {
        fmt.Println(v38)
        v38 = v38
        v4 = (v4 % ((42) | 1))
    }
    switch genFn0(v4) {
    case 586:
        v38 -= 1
    default:
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v37)
        case "a":
            fmt.Println(v2, v2)
        }
    }
    switch 42 {
    case 9223372036854775807:
        var v42 uint = (^(v38))
        _ = v42
    }
    if false {
        v3 = append(v3, "a")
        var v43 int32 = (-(v4))
        _ = v43
    }
    var v44 int32 = v4
    v4 = (^(v4))
    v38 = (v38 + v38)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v45 int16 = -689
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v44
    _ = v45
}
