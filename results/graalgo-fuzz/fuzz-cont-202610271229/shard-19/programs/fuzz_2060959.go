package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0() int {
    if true {
        var v7 float32 = float32(-2.25)
        switch 0 {
        case 391:
            fmt.Println(v7)
        }
        var v8 int16 = 0
        _ = v7
        _ = v8
    } else {
        var v9 float64 = 0.1
        _ = v9
    }
    switch 32767 {
    case 32767:
        for fi10 := 0; fi10 < 18; fi10++ {
            for fi11 := 0; fi11 < 8; fi11++ {
                if true {
                    var v12 uint64 = 18446744073709551615
                    if false {
                        var v13 int8 = 1
                        if false {
                            if false {
                                if true {
                                    v12 += (^(v12))
                                    fmt.Println(v12)
                                    v14 := []int{108}
                                    _ = v14
                                }
                                if false {
                                    fmt.Println(v12, fi11, v12)
                                }
                                v15 := []int{fi10, fi11, (fi10 & (fi10 ^ (fi10 & fi11))), (fi11 >> ((4294967295) & 63))}
                                _ = v15
                            } else {
                                v16 := []int{(fi10 >> ((255) & 63)), fi11, fi11, 127}
                                v13 = -91
                                v16 = append(v16, 1)
                                switch "\n\t" {
                                case "hello":
                                    v12 = (v12 % (((v12 << ((uint(v12)) & 63))) | 1))
                                }
                                v17 := GenS1{F0: 42}
                                fmt.Println(fi10, fi10, v12)
                                v18 := GenS1{F0: 18446744073709551615}
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            v12 -= (v12 & (^(v12)))
                        } else {
                            var v19 string = "hello"
                            _ = v19
                        }
                        _ = v13
                    }
                    _ = v12
                } else {
                    v20 := []int{fi10, fi10, -2147483648, (^(fi11))}
                    _ = v20
                }
                var v21 int64 = -9223372036854775808
                v21 += 42
                _ = fi11
                _ = v21
            }
            v22 := GenS1{F0: 0}
            fmt.Println(fi10)
            _ = fi10
            _ = v22
        }
    default:
        var v23 int32 = -2147483648
        _ = v23
    }
    return 4294967295
}
func genFn1(p0 bool) uint32 {
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 int16 = 42
        _ = fi24
        _ = v25
    }
    p0 = (!true)
    _ = p0
    return 4294967295
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]int)
    var v3 uint32 = 2147483648
    v4 := make(map[string]string)
    var v5 uint32 = (^(v3))
    var v6 int = 9223372036854775807
    if ((v1 ^ (v1 * (^(v1)))) < 32767) {
        v4["\n\t"] = (("\n\t" + "foo bar") + "hello")
        var v26 bool = false
        _ = v26
    }
    var v27 int8 = 59
    if true {
        fmt.Println(v1, v5)
        fmt.Println(v1)
        var v28 string = "a"
        _ = v28
    }
    v29 := GenS2{F0: (v3 >> ((65535) & 63)), F1: 73}
    fmt.Println(v5)
    fmt.Println(v1, v5, v27)
    var v30 uint = 65535
    var v31 uint = (v30 | uint(v1))
    v31 = 0
    switch "hello" {
    case "hello":
        v1 -= uint16(v30)
    case "":
        if ((v3 ^ 65535) > (^(v5))) {
            var v32 uint16 = (v1 << (((v31 * 4294967295)) & 63))
            var v33 uint = (v30 % (((v30 - v31)) | 1))
            v34 := []int{}
            _ = v32
            _ = v33
            _ = v34
        }
    case "\n\t":
        var v35 uint8 = 255
        _ = v35
    }
    v36 := []int{}
    var v37 float64 = 1e10
    var v38 string = "hello"
    var v39 string = ""
    fmt.Println(v38)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
