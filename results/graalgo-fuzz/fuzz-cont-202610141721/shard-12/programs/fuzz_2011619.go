package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32, p1 uint) int {
    var v7 int32 = -1
    if true {
        var v8 int = -2147483648
        switch "a" {
        case "hello":
            v8 -= (^(v8))
        case "\n\t":
            if true {
                p0 += v7
                v9 := GenS1{F0: 255}
                for fi10 := 0; fi10 < 7; fi10++ {
                    var v11 uint16 = 2
                    if true {
                        var v12 int16 = -32768
                        p0 -= 399
                        var v13 string = ""
                        _ = v12
                        _ = v13
                    } else {
                        fmt.Println(p1, v7)
                        v14 := []int{int(p0), -9223372036854775808, 32767, 42}
                        var v15 int = (v8 ^ 1)
                        _ = v14
                        _ = v15
                    }
                    _ = fi10
                    _ = v11
                }
                _ = v9
            } else {
                p1 = (p1 & p1)
                v16 := GenS1{F0: 1}
                _ = v16
            }
        }
        fmt.Println(p1, v8, p0)
        _ = v8
    } else {
        v17 := []int{32768, 4294967295, 32768, 127}
        var v18 uint = (p1 | p1)
        var v19 uint64 = 0
        _ = v17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := []int{fi20, (fi20 & (-(fi20))), (fi20 << ((p1) & 63)), (^(fi20))}
        var v22 int16 = -1
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println(p1, p0)
    p1 -= (p1 << ((p1) & 63))
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 string, p1 uint16, p2 bool) int16 {
    if (182 >= -9223372036854775808) {
        fmt.Println(p1)
    }
    if true {
        p1 = p1
        for fi23 := 0; fi23 < 16; fi23++ {
            v24 := []int{-807, (fi23 / ((genFn0(2147483647, 18446744073709551615)) | 1))}
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 int16 = 288
                _ = fi25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
    } else {
        p0 = "\n\t"
    }
    p2 = p2
    v27 := []int{genFn0(42, 255), 128}
    switch 32768 {
    case -2147483648:
        var v28 uint16 = (p1 << ((255) & 63))
        _ = v28
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 273
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 2
    var v3 bool = true
    var v4 int = 2147483647
    v5 := []int{v4, (v4 + v4), v4, (-(v4))}
    v6 := make(map[string]uint32)
    for fi29 := 0; fi29 < 20; fi29++ {
        var v30 int32 = -850
        v30 += (v30 / (((-(v30))) | 1))
        switch "foo bar" {
        case "a":
            switch "" {
            case "\n\t":
                var v31 int8 = 8
                _ = v31
            case "a":
                var v32 uint32 = 2
                _ = v32
            case "hello":
                if v3 {
                    v3 = false
                    v33 := GenS1{F0: 190}
                    fmt.Println(v3, v2)
                    _ = v33
                } else {
                    if (4294967295 < 1) {
                        var v34 int16 = 67
                        for fi35 := 0; fi35 < 17; fi35++ {
                            v6["hello"] = uint32(v2)
                            if v3 {
                                var v36 int64 = 1
                                var v37 string = ((("hello" + "a") + "") + "")
                                var v38 uint16 = 2
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                            _ = fi35
                        }
                        _ = v34
                    }
                }
            }
        case "":
            for fi39 := 0; fi39 < 16; fi39++ {
                v3 = (false || v3)
                var v40 uint = 18446744073709551615
                var v41 float32 = (0.1 - 1e10)
                _ = fi39
                _ = v40
                _ = v41
            }
        case "foo bar":
            switch "a" {
            case "a":
                fmt.Println(v3, v2)
            case "":
                _ = len(v5)
            default:
                _ = len(v5)
            }
        }
        _ = fi29
        _ = v30
    }
    fmt.Println(v4, v2, v2)
    v42 := v6["foo bar"]
    fmt.Println(v2, v4)
    v43 := GenS2{F0: 32767}
    v44 := v1[(("" + "\n\t") + "\n\t")]
    var v45 int16 = genFn1("foo bar", (v2 & v2), false)
    v46 := GenS1{F0: 93}
    var v47 uint = uint(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
