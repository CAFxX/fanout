package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int8 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
func genFn0(p0 int32, p1 int16, p2 string) int32 {
    p2 = p2
    for fi4 := 0; fi4 < 3; fi4++ {
        if false {
            if (1.5 < 3.14159) {
                fmt.Println(fi4, p0, p2)
                switch (^(fi4)) {
                case 9223372036854775807:
                    var v5 float64 = float64(0.1)
                    _ = v5
                case 4294967295:
                    p1 -= (^(p1))
                case 924:
                    for fi6 := 0; fi6 < 2; fi6++ {
                        p2 = ("" + p2)
                        fmt.Println(p1, fi6, fi6)
                        _ = fi6
                    }
                }
            }
            if (!false) {
                p1 = -989
                v7 := GenS1{F0: false, F1: p2, F2: -1}
                var v8 uint32 = 2147483647
                _ = v7
                _ = v8
            } else {
                for fi9 := 0; fi9 < 7; fi9++ {
                    if false {
                        p1 += -874
                        fmt.Println(p0, fi9)
                    } else {
                        if true {
                            if (!false) {
                                if ("hello" == ((p2 + ((p2 + "a") + (((p2 + p2) + "") + "a"))) + (p2 + (("\n\t" + "") + p2)))) {
                                    var v10 uint32 = 2147483647
                                    var v11 int8 = -1
                                    _ = v10
                                    _ = v11
                                }
                                if (false && (0 < p0)) {
                                    p0 += -1
                                }
                            } else {
                                p2 = ("\n\t" + "")
                                var v12 string = "a"
                                v13 := []int{}
                                _ = v12
                                _ = v13
                            }
                            v14 := []int{}
                            _ = len(v14)
                            _ = v14
                        }
                    }
                    _ = fi9
                }
                p1 += p1
            }
        } else {
            if true {
                if true {
                    v15 := []int{fi4, (fi4 - fi4), fi4, -460}
                    p1 += p1
                    _ = v15
                } else {
                    v16 := []int{}
                    for fi17 := 0; fi17 < 15; fi17++ {
                        var v18 uint64 = 4294967295
                        _ = fi17
                        _ = v18
                    }
                    _ = len(v16)
                    _ = v16
                }
                for fi19 := 0; fi19 < 7; fi19++ {
                    var v20 string = p2
                    if true {
                        p1 -= p1
                        if (42 <= int8(p1)) {
                            p0 -= (-(p0))
                            fmt.Println(fi19, v20, p0)
                        } else {
                            p0 -= p0
                            p2 = ((((("foo bar" + "\n\t") + v20) + v20) + (v20 + (p2 + (v20 + p2)))) + p2)
                        }
                    }
                    v21 := GenS1{F0: true, F1: v20, F2: int8(fi19)}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            fmt.Println(p0)
            p2 = "foo bar"
        }
        var v22 uint16 = 1
        fmt.Println(fi4, p2, p0)
        _ = fi4
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 * p0)
}

func main() {
    var v1 uint = 519
    var v2 int16 = -168
    v3 := make(map[string]int)
    for fi23 := 0; fi23 < 5; fi23++ {
        v24 := GenS1{F0: false, F1: (("foo bar" + "hello") + ("\n\t" + "hello")), F2: -128}
        var v25 int64 = -550
        _ = fi23
        _ = v24
        _ = v25
    }
    v26 := GenS1{F0: ((!(!false)) && false), F1: "foo bar", F2: -28}
    switch "a" {
    case "a":
        fmt.Println(v2)
    case "":
        var v27 uint32 = 93
        _ = v27
    default:
        v1 = 255
    }
    v1 = v1
    fmt.Println(v1, v2)
    v28 := GenS1{F0: false, F1: "\n\t", F2: 127}
    var v29 uint16 = 42
    v2 = (v2 | v2)
    fmt.Println(v1, v2, v2)
    v30 := GenS2{F0: (!true), F1: "\n\t"}
    var v31 float64 = (1e10 + 3.14159)
    if (true && (127 != 32767)) {
        for fi32 := 0; fi32 < 10; fi32++ {
            v31 -= v31
            v33 := GenS3{F0: (v29 | v29)}
            _ = fi32
            _ = v33
        }
        v34 := GenS2{F0: true, F1: ""}
        _ = v34
    }
    v35 := GenS4{F0: 437}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
}
