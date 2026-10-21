package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int64; F1 bool; F2 string }
type GenS5 struct { F0 int8; F1 uint8; F2 int }
func genFn0() int8 {
    var v4 int32 = -1
    var v5 float64 = 0.1
    var v6 string = (((("\n\t" + "hello") + "hello") + "foo bar") + "\n\t")
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 string = v6
        if (false || ((1 != 119) && false)) {
            fmt.Println(v5, v6, v5)
            v9 := GenS1{F0: 128}
            _ = v9
        } else {
            v6 = v8
            if false {
                if false {
                    if false {
                        for fi10 := 0; fi10 < 19; fi10++ {
                            v5 = (v5 / (-2.25 / 1e10))
                            v11 := GenS2{F0: 65535, F1: 1}
                            var v12 uint16 = 0
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                    } else {
                        fmt.Println(v4, v6)
                    }
                    if false {
                        v13 := GenS2{F0: 686, F1: -1}
                        _ = v13
                    } else {
                        v14 := []int{}
                        for fi15 := 0; fi15 < 2; fi15++ {
                            var v16 bool = true
                            _ = fi15
                            _ = v16
                        }
                        _ = v14
                    }
                    var v17 uint32 = 2147483647
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 17; fi18++ {
                        switch fi7 {
                        case 4294967295:
                            if (false || true) {
                                fmt.Println(v8, fi7)
                            } else {
                                var v19 int8 = 1
                                _ = v19
                            }
                        case 32768:
                            v20 := []int{4294967295, fi7}
                            _ = v20
                        default:
                            var v21 uint32 = 65535
                            _ = v21
                        }
                        switch "foo bar" {
                        case "":
                            if true {
                                var v22 int64 = -9223372036854775808
                                _ = v22
                            }
                        default:
                            switch fi7 {
                            case 4294967295:
                                v23 := []int{fi18, fi18, (fi7 | 429), 4294967295}
                                _ = v23
                            case 32768:
                                v5 += v5
                            case -2147483648:
                                v24 := GenS1{F0: 121}
                                _ = v24
                            }
                        }
                        _ = fi18
                    }
                    fmt.Println(v8, fi7, v5)
                    fmt.Println(v6, v5)
                }
            } else {
                v25 := GenS3{F0: 2, F1: 2147483647}
                v5 = 0.0
                _ = v25
            }
        }
        _ = fi7
        _ = v8
    }
    var v26 float64 = v5
    _ = v26
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint16 = 42
    var v2 uint = 4294967295
    v3 := make(map[string]bool)
    if false {
        if false {
            if false {
                var v27 string = ("" + "hello")
                for fi28 := 0; fi28 < 3; fi28++ {
                    var v29 uint16 = v1
                    v30 := v3[("a" + v27)]
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v1 += v1
                _ = v27
            }
            if false {
                fmt.Println(v2, v1)
            } else {
                v31 := []int{754, -19, 0, -2147483648}
                fmt.Println(v1, v2)
                _ = len(v3)
                _ = v31
            }
        } else {
            delete(v3, "foo bar")
        }
        for fi32 := 0; fi32 < 14; fi32++ {
            if true {
                var v33 uint = (^(v2))
                _ = v33
            } else {
                fmt.Println(v1, fi32)
            }
            var v34 int32 = 0
            v35 := GenS4{F0: 0, F1: (2 >= 18446744073709551615), F2: ("\n\t" + "a")}
            _ = fi32
            _ = v34
            _ = v35
        }
        v36 := GenS3{F0: 2, F1: 384}
        _ = v36
    }
    v37 := []int{0, -77, 2147483647}
    _ = len(v3)
    fmt.Println(v1, v1)
    var v38 float64 = 1e10
    fmt.Println(v1)
    var v39 bool = false
    for fi40 := 0; fi40 < 1; fi40++ {
        v2 -= 2
        _ = fi40
    }
    v1 -= 32767
    v39 = v39
    for fi41 := 0; fi41 < 6; fi41++ {
        fmt.Println(v39, fi41, fi41)
        v37 = append(v37, -9223372036854775808)
        v3["a"] = (65535 > 1)
        _ = fi41
    }
    switch "hello" {
    case "hello":
        if true {
            v42 := GenS5{F0: genFn0(), F1: 0, F2: 127}
            _ = v42
        } else {
            var v43 uint32 = 65535
            var v44 uint16 = v1
            _ = v43
            _ = v44
        }
    case "foo bar":
        var v45 uint64 = 4294967295
        _ = v45
    }
    var v46 float32 = (float32(v38) - 1.5)
    for fi47 := 0; fi47 < 4; fi47++ {
        var v48 uint16 = (^(v1))
        v2 = (v2 >> ((v2) & 63))
        v46 = v46
        _ = fi47
        _ = v48
    }
    if len(v37) >= 2 {
        v37 = v37[1:len(v37)]
    }
    for fi49 := 0; fi49 < 9; fi49++ {
        var v50 int32 = 1
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v46, v39)
        }
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v46
}
