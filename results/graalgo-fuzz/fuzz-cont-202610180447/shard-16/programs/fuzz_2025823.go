package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0(p0 int64, p1 int64, p2 uint) uint16 {
    v5 := GenS1{F0: -1}
    if (0 <= -1) {
        v6 := GenS1{F0: (p0 + (-(p0)))}
        p0 -= p1
        var v7 float32 = 0.0
        _ = v6
        _ = v7
    } else {
        fmt.Println(p2, p2)
    }
    var v8 uint64 = 2
    v9 := []int{-553, 127, -1, -407}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    _ = v9
    return 42
}
func genFn1() int {
    v10 := []int{9223372036854775807, -291}
    var v11 string = ("hello" + "foo bar")
    _ = v10
    _ = v11
    return 128
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[string]string)
    v3 := []int{(v1 * v1), (v1 << ((255) & 63))}
    var v4 float64 = (((3.14159 * float64(3.14159)) + 1.5) - 0.1)
    fmt.Println(v4)
    v4 += (1.5 / 0.1)
    if false {
        fmt.Println(v4, v4)
        fmt.Println(v4, v4)
        var v12 float64 = v4
        _ = v12
    } else {
        fmt.Println(v1)
        var v13 uint32 = 42
        if true {
            v14 := GenS1{F0: -191}
            if (1.5 <= v4) {
                v15 := GenS1{F0: 1}
                v16 := GenS1{F0: 1}
                var v17 int64 = 1
                _ = v15
                _ = v16
                _ = v17
            } else {
                switch "hello" {
                case "foo bar":
                    switch "\n\t" {
                    case "":
                        v13 = v13
                    case "a":
                        v2["hello"] = ("a" + "foo bar")
                    }
                case "hello":
                    fmt.Println(v13)
                }
            }
            _ = v14
        } else {
            fmt.Println(v1, v13)
        }
        _ = v13
    }
    if len(v3) > 0 {
        v3[0] = v1
    }
    v1 += (v1 * 32767)
    var v18 uint8 = 127
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 float64 = (v4 * v4)
        if true {
            if len(v3) > 0 {
                v3[0] = fi19
            }
            if (124 > 0) {
                fmt.Println(v20, v18, v1)
            } else {
                switch fi19 {
                case -226:
                    delete(v2, "foo bar")
                }
                var v21 int64 = -514
                v4 += 0.1
                _ = v21
            }
        } else {
            var v22 uint8 = v18
            for fi23 := 0; fi23 < 10; fi23++ {
                switch (fi23 | 2147483647) {
                case 9223372036854775807:
                    var v24 int32 = int32(v1)
                    _ = v24
                case 127:
                    v2[("foo bar" + (("\n\t" + ("hello" + ("foo bar" + ""))) + ("foo bar" + "a")))] = ("a" + (("\n\t" + "foo bar") + "\n\t"))
                case 559:
                    v25 := GenS1{F0: -245}
                    _ = v25
                }
                v26 := GenS1{F0: 1}
                v3 = append(v3, fi19)
                _ = fi23
                _ = v26
            }
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        switch ("\n\t" + "foo bar") {
        case "a":
            for fi28 := 0; fi28 < 4; fi28++ {
                fmt.Println(v1, v1)
                _ = fi28
            }
        case "":
            _ = len(v2)
        case "foo bar":
            for fi29 := 0; fi29 < 17; fi29++ {
                v3 = append(v3, v1)
                _ = fi29
            }
        default:
            if (1e10 >= ((v4 + v4) + -2.25)) {
                var v30 int64 = 1
                var v31 bool = (!(!true))
                _ = v30
                _ = v31
            } else {
                v1 = fi27
                if (0 >= 42) {
                    var v32 uint = 451
                    _ = v32
                } else {
                    if false {
                        var v33 float32 = 1e10
                        v34 := GenS2{F0: "a", F1: 1}
                        v35 := GenS1{F0: int64(v18)}
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    switch "foo bar" {
                    case "a":
                        var v36 bool = (42 != 65535)
                        _ = v36
                    case "":
                        if false {
                            if true {
                                switch -658 {
                                case 32767:
                                    var v37 uint = 296
                                    _ = v37
                                case 128:
                                    var v38 uint = 4294967295
                                    _ = v38
                                case 32768:
                                    v1 = v1
                                }
                                var v39 bool = (uint16(v18) == 699)
                                _ = v39
                            } else {
                                for fi40 := 0; fi40 < 11; fi40++ {
                                    switch (-(v1)) {
                                    case 2147483647:
                                        v41 := v2[("" + ("hello" + "hello"))]
                                        _ = v41
                                    default:
                                        delete(v2, "")
                                    }
                                    v1 += fi27
                                    _ = len(v2)
                                    _ = fi40
                                }
                            }
                            fmt.Println(fi27)
                        }
                    case "hello":
                        v1 = (fi27 & 127)
                    }
                }
            }
        }
        for fi42 := 0; fi42 < 11; fi42++ {
            v43 := GenS3{F0: 2, F1: 1}
            v44 := GenS3{F0: 776, F1: 0}
            v4 += (v4 + (1.5 / 0.1))
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi27
    }
    delete(v2, ("foo bar" + ""))
    v45 := GenS1{F0: 290}
    var v46 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
}
