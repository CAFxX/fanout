package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -60
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    if (true && (!true)) {
        if true {
            v5 := GenS1{F0: 255, F1: 42, F2: 65535}
            v2 = v2
            fmt.Println(v2)
            _ = v5
        }
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 string = ("a" + (("a" + ("" + "\n\t")) + "hello"))
            _ = fi6
            _ = v7
        }
        _ = len(v3)
    } else {
        v2 = (v2 << ((65535) & 63))
        var v8 uint32 = 65535
        _ = v8
    }
    switch "\n\t" {
    case "":
        var v9 uint32 = 1
        _ = v9
    default:
        switch "" {
        case "hello":
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS2{F0: -1}
                if (!((1 < 255) && true)) {
                    var v12 uint32 = 674
                    var v13 int64 = v2
                    for fi14 := 0; fi14 < 10; fi14++ {
                        var v15 uint32 = 202
                        _ = fi14
                        _ = v15
                    }
                    _ = v12
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
        case "foo bar":
            if false {
                v2 = v2
                for fi16 := 0; fi16 < 15; fi16++ {
                    v17 := []int{-985, (fi16 & 32768), fi16}
                    _ = fi16
                    _ = v17
                }
            }
        }
    }
    var v18 string = "hello"
    v18 = v18
    v19 := []int{32768, 32767, 9223372036854775807}
    v3[("" + (v18 + v18))] = true
    for fi20 := 0; fi20 < 10; fi20++ {
        fmt.Println(fi20, v18, v2)
        _ = fi20
    }
    if (false && true) {
        var v21 uint16 = 575
        for fi22 := 0; fi22 < 1; fi22++ {
            if true {
                if false {
                    switch int(v21) {
                    case 254:
                        if len(v19) > 0 {
                            v19[0] = (fi22 / ((fi22) | 1))
                        }
                    case 475:
                        for fi23 := 0; fi23 < 11; fi23++ {
                            v2 = v2
                            _ = fi23
                        }
                    default:
                        switch (^(fi22)) {
                        case 127:
                            v2 = v2
                        default:
                            if (0.0 == 0.1) {
                                if true {
                                    v2 += v2
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        if (fi22 == (fi22 | -2147483648)) {
                                            v24 := GenS1{F0: 255, F1: 888, F2: 1}
                                            fmt.Println(v21, v2)
                                            _ = v24
                                        } else {
                                            for fi25 := 0; fi25 < 6; fi25++ {
                                                fmt.Println(fi22, fi25, v21)
                                                _ = fi25
                                            }
                                        }
                                    case "hello":
                                        v18 = v18
                                    case "":
                                        v26 := v4[876]
                                        _ = v26
                                    }
                                    switch v18 {
                                    case "\n\t":
                                        if true {
                                            fmt.Println(v21, fi22)
                                            var v27 uint64 = 4294967295
                                            _ = v27
                                        } else {
                                            var v28 int64 = (v2 | (v2 % ((-1) | 1)))
                                            v19 = append(v19, (fi22 & fi22))
                                            var v29 string = v18
                                            _ = v28
                                            _ = v29
                                        }
                                    case "":
                                        var v30 int8 = 42
                                        _ = v30
                                    case "a":
                                        if (!(985 == 4294967295)) {
                                            for fi31 := 0; fi31 < 6; fi31++ {
                                                _ = len(v3)
                                                var v32 uint = 0
                                                _ = fi31
                                                _ = v32
                                            }
                                        } else {
                                            fmt.Println(v2, v18, v2)
                                            var v33 uint = 4294967295
                                            v18 = v18
                                            _ = v33
                                        }
                                    }
                                }
                                var v34 uint = 2
                                var v35 int32 = 1
                                _ = v34
                                _ = v35
                            } else {
                                if (0.0 >= (1e10 + 1.5)) {
                                    var v36 uint8 = 94
                                    _ = v36
                                } else {
                                    fmt.Println(v2)
                                }
                                var v37 bool = true
                                _ = v37
                            }
                        }
                    }
                }
                fmt.Println(v2, v21, v18)
                v38 := GenS2{F0: 1}
                _ = v38
            }
            if true {
                fmt.Println(fi22, v18, v2)
            } else {
                for fi39 := 0; fi39 < 16; fi39++ {
                    v21 = v21
                    v40 := v3[v18]
                    _ = fi39
                    _ = v40
                }
                var v41 uint = 4294967295
                _ = v41
            }
            _ = fi22
        }
        _ = v21
    }
    var v42 int8 = 1
    switch ((v18 + v18) + "hello") {
    case "foo bar":
        switch v18 {
        case "":
            v43 := GenS3{F0: 0, F1: 4294967295}
            _ = v43
        default:
            for fi44 := 0; fi44 < 20; fi44++ {
                var v45 uint = 4294967295
                _ = fi44
                _ = v45
            }
        }
    }
    var v46 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v46
}
