package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint32; F1 int; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint8 }
func genFn0(p0 bool, p1 int64, p2 uint64) uint {
    if (42 == 0) {
        v5 := GenS1{F0: 0, F1: ("foo bar" + ("a" + "hello"))}
        var v6 int8 = -1
        _ = v5
        _ = v6
    }
    p0 = (!p0)
    if p0 {
        p2 = (^(p2))
        var v7 uint64 = p2
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := []int{(fi8 >> ((581) & 63)), fi8}
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 int16) uint64 {
    v10 := []int{254}
    fmt.Println(p0)
    switch "a" {
    case "hello":
        p0 += (p0 << ((124) & 63))
    case "foo bar":
        if false {
            var v11 int64 = -800
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            v11 -= 1
            _ = v11
        } else {
            var v12 string = "a"
            if false {
                if true {
                    if false {
                        if false {
                            var v13 int32 = -2147483648
                            v13 -= 1
                            _ = v13
                        } else {
                            var v14 int16 = p0
                            _ = v14
                        }
                        switch 32768 {
                        case 42:
                            p0 += p0
                        default:
                            _ = len(v10)
                        }
                    } else {
                        p0 -= p0
                    }
                    if len(v10) > 0 {
                        v10[0] = 32767
                    }
                    var v15 string = v12
                    _ = v15
                }
                if (!true) {
                    v12 = (v12 + (((v12 + "foo bar") + v12) + v12))
                } else {
                    switch 32768 {
                    case -333:
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    default:
                        switch ((v12 + "foo bar") + v12) {
                        case "hello":
                            if false {
                                v12 = v12
                                switch 1 {
                                case 0:
                                    v12 = (("a" + v12) + v12)
                                case -614:
                                    v12 = "a"
                                default:
                                    fmt.Println(v12, v12)
                                }
                            } else {
                                v16 := GenS2{F0: 65535, F1: 0, F2: 1}
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                _ = v16
                            }
                        case "foo bar":
                            v12 = (((v12 + (v12 + v12)) + "hello") + v12)
                        case "a":
                            var v17 uint16 = 0
                            _ = v17
                        }
                    }
                }
                var v18 int64 = 867
                _ = v18
            }
            fmt.Println(p0)
            _ = v12
        }
    case "a":
        v19 := GenS3{F0: p0, F1: 4294967295, F2: 351}
        _ = v19
    default:
        for fi20 := 0; fi20 < 5; fi20++ {
            fmt.Println(p0, p0)
            v21 := GenS4{F0: 2147483647, F1: 128}
            _ = fi20
            _ = v21
        }
    }
    switch 42 {
    case -1:
        p0 = (^(p0))
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        p0 = p0
        for fi23 := 0; fi23 < 1; fi23++ {
            switch ("a" + "a") {
            case "a":
                if true {
                    v24 := GenS4{F0: -646, F1: 2}
                    var v25 float32 = 0.1
                    _ = v24
                    _ = v25
                } else {
                    fmt.Println(fi23, fi22, fi22)
                }
            case "foo bar":
                fmt.Println(p0, fi22)
            case "hello":
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
            p0 = p0
            var v26 uint8 = 1
            _ = fi23
            _ = v26
        }
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 int64 = 0
            var v29 bool = (189 != 669)
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi22
    }
    _ = p0
    _ = v10
    return 848
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint8 = 128
    v3 := make(map[int]int)
    var v4 int8 = 1
    var v30 int = -1
    var v31 int16 = -17
    v1 = (v1 + 0)
    v32 := []int{v30}
    fmt.Println(v30, v1)
    v4 = v4
    var v33 uint32 = (v1 >> ((4294967295) & 63))
    switch (^(v30)) {
    case -877:
        v34 := GenS1{F0: (-(v30)), F1: ("foo bar" + ((("" + (((("" + "hello") + "a") + "hello") + "a")) + ("foo bar" + "a")) + ("a" + "\n\t")))}
        _ = v34
    }
    var v35 int16 = 32767
    for fi36 := 0; fi36 < 3; fi36++ {
        v37 := GenS1{F0: (fi36 * v30), F1: ("a" + "")}
        _ = fi36
        _ = v37
    }
    switch 4294967295 {
    case 42:
        v2 = 128
    }
    var v38 float64 = float64(0.1)
    fmt.Println(v38, v1, v30)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v38
    _ = v4
}
