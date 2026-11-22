package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int8 = 1
    var v2 string = ""
    v3 := make(map[string]bool)
    var v4 uint32 = 4294967295
    var v5 int32 = 383
    v6 := make(map[string]int)
    v5 += v5
    v4 -= v4
    v7 := GenS1{F0: v2}
    v8 := GenS1{F0: "\n\t"}
    if (-2.25 <= 0.0) {
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS1{F0: ""}
            _ = fi9
            _ = v10
        }
        fmt.Println(v4, v2)
    } else {
        var v11 uint16 = 0
        _ = v11
    }
    v1 += v1
    if false {
        v6[v2] = 2147483647
        v12 := GenS2{F0: 1, F1: v2}
        v13 := GenS3{F0: v2, F1: 0}
        _ = v12
        _ = v13
    } else {
        var v14 uint32 = v4
        var v15 uint16 = 42
        _ = v14
        _ = v15
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v2, v5, v4)
    default:
        var v16 int8 = v1
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        if false {
            v18 := v6["a"]
            if false {
                if (v1 < v1) {
                    fmt.Println(v18, v1)
                    v4 += (^(v4))
                } else {
                    if (127 >= 1) {
                        v5 += 94
                        v19 := []int{}
                        switch (-(v18)) {
                        case 32768:
                            var v20 float32 = 0.0
                            _ = v20
                        default:
                            v6["foo bar"] = (^(v18))
                        }
                        _ = v19
                    } else {
                        fmt.Println(v1)
                        if false {
                            var v21 bool = (32767 < 42)
                            _ = v21
                        } else {
                            v22 := GenS2{F0: -1, F1: v2}
                            v23 := v6[""]
                            for fi24 := 0; fi24 < 19; fi24++ {
                                if (65535 > 42) {
                                    var v25 uint8 = 42
                                    if (false || true) {
                                        var v26 float64 = 0.1
                                        var v27 uint64 = 4294967295
                                        _ = v26
                                        _ = v27
                                    } else {
                                        fmt.Println(v18, v23, v18)
                                        var v28 int = v23
                                        var v29 int64 = 435
                                        _ = v28
                                        _ = v29
                                    }
                                    _ = v25
                                }
                                if false {
                                    v2 = v2
                                    fmt.Println(v4, v5, v5)
                                }
                                v23 = fi17
                                _ = fi24
                            }
                            _ = v22
                            _ = v23
                        }
                    }
                }
            } else {
                var v30 uint8 = 19
                _ = v30
            }
            v1 = 111
            _ = v18
        }
        var v31 int64 = -254
        var v32 int8 = -123
        _ = fi17
        _ = v31
        _ = v32
    }
    var v33 int16 = 1000
    var v34 uint32 = (v4 << ((255) & 63))
    var v35 int = 32768
    v36 := GenS2{F0: (v33 | (v33 ^ -1)), F1: (v2 + ((v2 + (v2 + "")) + (v2 + "a")))}
    if (127 >= 190) {
        for fi37 := 0; fi37 < 8; fi37++ {
            var v38 bool = (-32768 > v33)
            _ = fi37
            _ = v38
        }
        var v39 int8 = v1
        _ = v39
    } else {
        _ = len(v6)
        v40 := GenS4{F0: 30}
        v41 := GenS1{F0: v2}
        _ = v40
        _ = v41
    }
    v4 = v34
    for fi42 := 0; fi42 < 7; fi42++ {
        v43 := []int{int(v1), fi42, v35, (fi42 << ((379) & 63))}
        for fi44 := 0; fi44 < 3; fi44++ {
            switch (("" + (v2 + v2)) + v2) {
            case "\n\t":
                delete(v3, ("" + "a"))
            case "foo bar":
                fmt.Println(v33, v5)
            case "hello":
                v5 = v5
            default:
                if (true || (true && true)) {
                    v33 -= v33
                    v5 += v5
                    if true {
                        switch (v2 + (v2 + v2)) {
                        case "foo bar":
                            _ = len(v6)
                        case "\n\t":
                            v1 = -49
                            fmt.Println(fi44, v33)
                            v45 := GenS4{F0: 1}
                            _ = v45
                        case "hello":
                            if true {
                                v43 = append(v43, 476)
                                v2 = ("" + ("hello" + ("a" + "foo bar")))
                            }
                        default:
                            fmt.Println(v1, v2)
                        }
                        var v46 string = ((v2 + v2) + "a")
                        _ = v46
                    } else {
                        v35 += fi44
                        if false {
                            delete(v6, "hello")
                        }
                        fmt.Println(v33, v35, v1)
                    }
                    v47 := GenS3{F0: ((v2 + v2) + (v2 + ("hello" + "hello"))), F1: 18446744073709551615}
                    _ = len(v3)
                    _ = v47
                }
            }
            _ = fi44
        }
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
