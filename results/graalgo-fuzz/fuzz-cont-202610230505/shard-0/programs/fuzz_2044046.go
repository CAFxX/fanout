package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint }
type GenS4 struct { F0 int32; F1 int; F2 bool }
type GenS5 struct { F0 int32 }
func genFn0(p0 int, p1 uint64) uint64 {
    p0 += (p0 % ((p0) | 1))
    for fi6 := 0; fi6 < 5; fi6++ {
        switch (fi6 % ((0) | 1)) {
        case -541:
            v7 := []int{fi6, 127, (p0 + (-(p0))), 32768}
            _ = v7
        }
        v8 := GenS1{F0: ("" + "a")}
        _ = fi6
        _ = v8
    }
    v9 := GenS2{F0: 42}
    var v10 uint16 = 744
    v11 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 239
}
func genFn1(p0 int) uint16 {
    v12 := GenS3{F0: 115, F1: 535, F2: 42}
    v13 := []int{(^(p0)), p0, (p0 + (p0 - (p0 >> ((56) & 63))))}
    for fi14 := 0; fi14 < 8; fi14++ {
        if (false || (690 < 42)) {
            p0 = (-(fi14))
        } else {
            if (!false) {
                for fi15 := 0; fi15 < 17; fi15++ {
                    var v16 uint64 = 1
                    if (fi15 < fi15) {
                        if false {
                            p0 = 642
                            fmt.Println(fi14, v16)
                        }
                        if len(v13) > 0 {
                            v13[0] = fi14
                        }
                        switch "hello" {
                        case "\n\t":
                            if true {
                                v16 = (v16 - 1)
                                var v17 uint32 = uint32(v16)
                                var v18 int32 = -797
                                _ = v17
                                _ = v18
                            } else {
                                var v19 uint = 494
                                _ = v19
                            }
                        case "foo bar":
                            v13 = append(v13, 4294967295)
                        default:
                            v13 = append(v13, (p0 >> ((732) & 63)))
                        }
                    } else {
                        v16 -= (v16 * (^(v16)))
                    }
                    if len(v13) > 0 {
                        v13[0] = 2147483647
                    }
                    _ = fi15
                    _ = v16
                }
                v20 := GenS4{F0: 42, F1: (p0 / ((p0) | 1)), F2: (true && false)}
                fmt.Println(fi14)
                _ = v20
            }
        }
        p0 += fi14
        for fi21 := 0; fi21 < 13; fi21++ {
            p0 = (fi21 | 0)
            fmt.Println(fi21, fi21, fi14)
            _ = fi21
        }
        _ = fi14
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        p0 = 32767
        if (32767 == 42) {
            p0 += 1
            v23 := GenS1{F0: "\n\t"}
            p0 += fi22
            _ = v23
        }
        fmt.Println(fi22, p0)
        _ = fi22
    }
    _ = p0
    _ = v12
    _ = v13
    return 255
}

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[int]bool)
    v3 := []string{"a", "a"}
    v4 := make(map[string]bool)
    var v5 int8 = 0
    _ = len(v1)
    var v24 bool = true
    var v25 uint16 = 482
    switch "a" {
    case "a":
        if (-2147483648 >= 126) {
            fmt.Println(v5, v25)
            if (!(v25 <= v25)) {
                var v26 int = 42
                var v27 int8 = -1
                _ = v26
                _ = v27
            } else {
                if v24 {
                    v5 = 42
                    if (!v24) {
                        for fi28 := 0; fi28 < 14; fi28++ {
                            var v29 int64 = 42
                            _ = fi28
                            _ = v29
                        }
                    } else {
                        fmt.Println(v25, v5, v5)
                    }
                    v30 := GenS5{F0: 40}
                    _ = v30
                } else {
                    var v31 uint = 65535
                    var v32 bool = false
                    for fi33 := 0; fi33 < 14; fi33++ {
                        v34 := GenS5{F0: -2147483648}
                        var v35 string = ("hello" + "foo bar")
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    _ = v31
                    _ = v32
                }
            }
        }
    case "":
        v25 = genFn1(822)
    case "\n\t":
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            var v36 uint16 = v25
            _ = v36
        }
    default:
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    }
    delete(v2, -305)
    var v37 int64 = 42
    var v38 uint8 = 127
    if len(v3) > 0 {
        v3[0] = ("foo bar" + "a")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
