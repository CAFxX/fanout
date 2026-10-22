package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 42
    var v3 bool = (!false)
    v4 := make(map[string]string)
    v2 -= v2
    switch 127 {
    case 0:
        v5 := GenS1{F0: -34, F1: 2147483647}
        _ = v5
    case 2147483647:
        var v6 int16 = 1
        _ = v6
    case -2147483648:
        if (1.5 == -2.25) {
            if false {
                v2 = 1
                if v3 {
                    var v7 int = 392
                    _ = v7
                } else {
                    if (!(!(v2 != 42))) {
                        for fi8 := 0; fi8 < 9; fi8++ {
                            fmt.Println(v2, v3, fi8)
                            if v3 {
                                if v3 {
                                    fmt.Println(v3, fi8)
                                    var v9 float32 = 0.0
                                    _ = v9
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        v10 := v1["a"]
                                        _ = v10
                                    default:
                                        for fi11 := 0; fi11 < 12; fi11++ {
                                            fmt.Println(v3, fi11, fi8)
                                            var v12 int32 = 2147483647
                                            if (!(42 < -128)) {
                                                v13 := []int{}
                                                fmt.Println(v2, v3, v3)
                                                v12 -= v12
                                                _ = v13
                                            } else {
                                                switch (fi8 ^ 42) {
                                                case -375:
                                                    switch (("hello" + "a") + ("\n\t" + "a")) {
                                                    case "hello":
                                                        v14 := GenS1{F0: -65, F1: 2147483648}
                                                        _ = v14
                                                    default:
                                                        if v3 {
                                                            v15 := GenS1{F0: 1, F1: 2}
                                                            v16 := GenS2{F0: 1, F1: -5, F2: 42}
                                                            var v17 uint16 = 0
                                                            _ = v15
                                                            _ = v16
                                                            _ = v17
                                                        }
                                                    }
                                                }
                                                _ = len(v4)
                                            }
                                            _ = fi11
                                            _ = v12
                                        }
                                    }
                                }
                            } else {
                                delete(v4, (("a" + "hello") + "hello"))
                                v18 := []int{-2147483648, 128, fi8, fi8}
                                v19 := GenS1{F0: 59, F1: 2147483648}
                                _ = v18
                                _ = v19
                            }
                            _ = fi8
                        }
                        v20 := GenS1{F0: 54, F1: 2}
                        _ = v20
                    }
                    if (42 <= 847) {
                        for fi21 := 0; fi21 < 11; fi21++ {
                            delete(v1, "\n\t")
                            v2 += 42
                            _ = fi21
                        }
                        for fi22 := 0; fi22 < 18; fi22++ {
                            switch fi22 {
                            case 42:
                                var v23 int64 = 9223372036854775807
                                _ = v23
                            case 127:
                                var v24 uint16 = 548
                                _ = v24
                            case -9223372036854775808:
                                v3 = true
                            }
                            v25 := GenS2{F0: 0, F1: -62, F2: 65535}
                            _ = fi22
                            _ = v25
                        }
                    } else {
                        v26 := GenS1{F0: -1, F1: 1}
                        _ = v26
                    }
                }
                v2 = 143
            }
            if false {
                fmt.Println(v3, v3)
            }
            v3 = (42 == 204)
        } else {
            fmt.Println(v2)
        }
    default:
        v27 := GenS2{F0: 1, F1: -128, F2: 2}
        _ = v27
    }
    var v28 int16 = 1
    v3 = (v3 || true)
    v29 := []int{32768, 1, -444}
    v30 := v1["foo bar"]
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 uint = 907
        for fi33 := 0; fi33 < 18; fi33++ {
            v28 += v28
            v34 := GenS1{F0: -1, F1: (v30 * 42)}
            var v35 string = "foo bar"
            _ = fi33
            _ = v34
            _ = v35
        }
        if len(v29) > 0 {
            v29[0] = fi31
        }
        _ = fi31
        _ = v32
    }
    var v36 int64 = -1
    var v37 string = ""
    switch v37 {
    case "\n\t":
        fmt.Println(v3, v30, v30)
    }
    var v38 bool = v3
    var v39 float64 = 1e10
    if v3 {
        for fi40 := 0; fi40 < 11; fi40++ {
            var v41 uint64 = (^(v2))
            if (v2 > v41) {
                if v38 {
                    v2 += v41
                    v1["hello"] = 0
                } else {
                    var v42 uint32 = v30
                    v38 = v3
                    _ = v42
                }
            } else {
                v28 += 0
            }
            _ = fi40
            _ = v41
        }
    } else {
        var v43 uint64 = 2
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
