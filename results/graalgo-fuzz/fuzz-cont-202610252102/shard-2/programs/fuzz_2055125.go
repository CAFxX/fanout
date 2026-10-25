package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 float64 = float64(((1e10 * 1.5) + 3.14159))
    v2 := make(map[int]int)
    var v3 bool = (!(!(-32768 == 32767)))
    var v4 int8 = -128
    var v5 float32 = 3.14159
    _ = len(v2)
    delete(v2, 888)
    v6 := GenS1{F0: 628, F1: 189, F2: 2147483647}
    _ = len(v2)
    switch "hello" {
    case "foo bar":
        v5 += (((v5 / float32(float64((v5 + v5)))) / v5) - v5)
    default:
        v7 := GenS1{F0: -464, F1: 210, F2: -1}
        _ = v7
    }
    v8 := GenS1{F0: -32768, F1: 65535, F2: -1}
    switch 9223372036854775807 {
    case -1:
        var v9 float64 = (v1 * v1)
        _ = v9
    }
    fmt.Println(v5, v5, v5)
    if true {
        var v10 bool = v3
        _ = v10
    } else {
        v11 := []int{1}
        v12 := GenS1{F0: -1, F1: 2, F2: -1}
        v13 := GenS1{F0: -32768, F1: 0, F2: 1}
        _ = v11
        _ = v12
        _ = v13
    }
    v3 = ("hello" < ("" + "hello"))
    v5 += 0.1
    if (!(!(2 >= 210))) {
        fmt.Println(v1, v1)
    } else {
        var v14 string = ("a" + "foo bar")
        _ = v14
    }
    fmt.Println(v1, v3, v4)
    v15 := []int{66, 2147483647, -361}
    switch "foo bar" {
    case "foo bar":
        var v16 uint32 = 301
        _ = v16
    case "hello":
        var v17 string = "\n\t"
        _ = v17
    case "\n\t":
        if (v1 <= 0.1) {
            v4 += 100
            var v18 uint8 = 42
            if ((1e10 - 3.14159) <= v5) {
                var v19 uint = 42
                var v20 int16 = 1
                _ = v19
                _ = v20
            }
            _ = v18
        } else {
            v5 -= 1e10
            var v21 bool = ((!v3) && (((42 > -9223372036854775808) || true) || (134 > 42)))
            var v22 uint32 = 2147483648
            _ = v21
            _ = v22
        }
    default:
        if (v3 && (!false)) {
            switch "hello" {
            case "\n\t":
                if (v3 || (v4 > v4)) {
                    var v23 uint32 = 4294967295
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 5; fi24++ {
                        _ = len(v15)
                        _ = fi24
                    }
                    var v25 uint64 = 2
                    _ = v25
                }
            case "foo bar":
                if v3 {
                    for fi26 := 0; fi26 < 4; fi26++ {
                        if v3 {
                            v27 := GenS2{F0: 255}
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v4, v1)
                            case "":
                                if true {
                                    fmt.Println(v1, fi26)
                                    var v28 int8 = (v4 >> ((65535) & 63))
                                    _ = v28
                                }
                            default:
                                v3 = v3
                            }
                            _ = v27
                        }
                        fmt.Println(v1)
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = fi26
                    }
                    for fi29 := 0; fi29 < 2; fi29++ {
                        v3 = v3
                        _ = fi29
                    }
                } else {
                    var v30 uint16 = 2
                    if true {
                        _ = len(v2)
                        v31 := v2[-760]
                        fmt.Println(v30, v3, v30)
                        _ = v31
                    } else {
                        v32 := v2[2147483647]
                        var v33 int8 = 127
                        if v3 {
                            for fi34 := 0; fi34 < 18; fi34++ {
                                var v35 uint64 = 0
                                _ = fi34
                                _ = v35
                            }
                        } else {
                            v33 = (^(v4))
                            if (!(!(128 > 0))) {
                                fmt.Println(v30, v32)
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                for fi36 := 0; fi36 < 10; fi36++ {
                                    var v37 uint64 = 4294967295
                                    _ = fi36
                                    _ = v37
                                }
                            } else {
                                var v38 int16 = 42
                                _ = v38
                            }
                        }
                        _ = v32
                        _ = v33
                    }
                    var v39 string = ""
                    _ = v30
                    _ = v39
                }
            case "a":
                v5 += v5
            default:
                if (128 >= 14) {
                    v3 = v3
                    v40 := GenS1{F0: 42, F1: 1, F2: -2147483648}
                    _ = v40
                }
            }
            v41 := GenS3{F0: v4}
            _ = v41
        }
    }
    var v42 int64 = 1
    for fi43 := 0; fi43 < 6; fi43++ {
        var v44 uint32 = 1
        _ = fi43
        _ = v44
    }
    switch -1 {
    case -9223372036854775808:
        var v45 float64 = v1
        _ = v45
    case 42:
        v3 = v3
    case 128:
        v42 = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v8
}
