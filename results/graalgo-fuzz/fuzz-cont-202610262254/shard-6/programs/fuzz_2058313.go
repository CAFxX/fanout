package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
func genFn0(p0 string) int16 {
    p0 = ("foo bar" + (p0 + ""))
    v5 := GenS1{F0: true}
    var v6 uint32 = 2
    if true {
        if false {
            if (!(-2.25 <= 1e10)) {
                if (-128 >= -1) {
                    if (128 == 128) {
                        for fi7 := 0; fi7 < 14; fi7++ {
                            fmt.Println(fi7)
                            _ = fi7
                        }
                        v6 -= (v6 & v6)
                    } else {
                        v8 := GenS1{F0: (41 <= 128)}
                        _ = v8
                    }
                }
                switch -1 {
                case 2147483647:
                    for fi9 := 0; fi9 < 4; fi9++ {
                        p0 = ((p0 + (p0 + "hello")) + p0)
                        v6 -= 2147483648
                        _ = fi9
                    }
                }
                switch "" {
                case "a":
                    var v10 uint16 = 32767
                    _ = v10
                case "":
                    v11 := []int{}
                    _ = v11
                default:
                    switch ((p0 + p0) + "hello") {
                    case "a":
                        if (!true) {
                            for fi12 := 0; fi12 < 9; fi12++ {
                                var v13 uint64 = 2
                                _ = fi12
                                _ = v13
                            }
                            var v14 int16 = 1
                            _ = v14
                        } else {
                            v15 := GenS2{F0: p0}
                            for fi16 := 0; fi16 < 9; fi16++ {
                                p0 = p0
                                if (128 != 42) {
                                    p0 = p0
                                    if false {
                                        v17 := []int{(fi16 % ((fi16) | 1)), 0, (fi16 | (-(fi16)))}
                                        v6 -= 2147483647
                                        _ = v17
                                    }
                                } else {
                                    for fi18 := 0; fi18 < 11; fi18++ {
                                        var v19 uint64 = 4294967295
                                        _ = fi18
                                        _ = v19
                                    }
                                    v20 := GenS3{F0: (^(fi16))}
                                    _ = v20
                                }
                                _ = fi16
                            }
                            _ = v15
                        }
                    case "hello":
                        for fi21 := 0; fi21 < 11; fi21++ {
                            for fi22 := 0; fi22 < 6; fi22++ {
                                fmt.Println(fi22, fi22, p0)
                                _ = fi22
                            }
                            p0 = ""
                            _ = fi21
                        }
                    }
                }
            }
            v23 := []int{}
            var v24 uint32 = v6
            _ = v23
            _ = v24
        }
        var v25 bool = (!(false || (9223372036854775807 < 4294967295)))
        _ = v25
    }
    var v26 int32 = 1
    _ = p0
    _ = v26
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 string) uint16 {
    p0 = p0
    v27 := []int{}
    var v28 uint = 2
    _ = p0
    _ = v27
    _ = v28
    return 255
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 42
    v3 := []int{2147483647}
    var v4 float32 = 1e10
    var v29 uint16 = 768
    var v30 uint8 = 2
    for fi31 := 0; fi31 < 3; fi31++ {
        v4 += (v4 - v4)
        _ = fi31
    }
    if (!(3.14159 >= (1.5 - 3.14159))) {
        if ((false && true) || (!((-2147483648 > -9223372036854775808) && true))) {
            v29 = (^(v29))
            v30 = 128
            v32 := GenS2{F0: "\n\t"}
            _ = v32
        }
    }
    v33 := GenS1{F0: true}
    var v34 uint32 = 42
    for fi35 := 0; fi35 < 19; fi35++ {
        var v36 uint64 = v2
        _ = fi35
        _ = v36
    }
    var v37 float64 = ((1.5 * (0.1 / (float64(v4) * -2.25))) + float64(3.14159))
    v37 = v37
    for fi38 := 0; fi38 < 18; fi38++ {
        v4 = -2.25
        for fi39 := 0; fi39 < 11; fi39++ {
            var v40 int16 = 1
            _ = fi39
            _ = v40
        }
        fmt.Println(v34)
        _ = fi38
    }
    v30 = (v30 - v30)
    var v41 int16 = genFn0(((("hello" + "") + "") + ("foo bar" + "hello")))
    var v42 uint32 = (v34 & (^(v34)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v37
    _ = v4
    _ = v41
    _ = v42
}
