package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int16; F1 int; F2 int64 }
type GenS3 struct { F0 string; F1 int32; F2 uint16 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 uint16; F1 string; F2 uint }
type GenS6 struct { F0 string; F1 int; F2 int16 }
func genFn0() int {
    var v7 int16 = 1000
    fmt.Println(v7)
    v8 := GenS1{F0: -2147483648, F1: 65535, F2: 812}
    if true {
        v9 := GenS2{F0: (v7 << ((4294967295) & 63)), F1: 9223372036854775807, F2: -1}
        _ = v9
    } else {
        v10 := GenS2{F0: v7, F1: 460, F2: 1}
        var v11 uint8 = 127
        _ = v10
        _ = v11
    }
    v7 -= 32767
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 float64 = float64(float32(0.1))
    v2 := []uint32{590, 995, 1, 369}
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    var v5 bool = false
    v6 := []int{}
    for fi12 := 0; fi12 < 2; fi12++ {
        v5 = (!false)
        if false {
            switch "foo bar" {
            case "foo bar":
                v13 := GenS3{F0: "hello", F1: -292, F2: 32767}
                _ = v13
            case "\n\t":
                var v14 uint32 = 2
                _ = v14
            default:
                var v15 string = "hello"
                _ = v15
            }
            _ = len(v6)
            switch "foo bar" {
            case "foo bar":
                v16 := GenS3{F0: ("" + ("foo bar" + (("hello" + "\n\t") + ""))), F1: -2147483648, F2: 262}
                _ = v16
            case "\n\t":
                v17 := GenS3{F0: "foo bar", F1: 996, F2: 440}
                _ = v17
            case "hello":
                v18 := GenS4{F0: 1000, F1: v5}
                _ = v18
            }
        }
        _ = fi12
    }
    if false {
        fmt.Println(v5, v5)
    } else {
        fmt.Println(v1, v5)
    }
    if (((2 != 181) && v5) && (v1 <= (v1 - (-2.25 / 1e10)))) {
        fmt.Println(v5, v1)
    }
    if len(v2) > 0 {
        v2[0] = 65535
    }
    switch "\n\t" {
    case "":
        var v19 uint32 = 4294967295
        _ = v19
    case "\n\t":
        switch 32767 {
        case 4294967295:
            fmt.Println(v1)
        default:
            switch "foo bar" {
            case "":
                for fi20 := 0; fi20 < 12; fi20++ {
                    var v21 int64 = -1
                    _ = fi20
                    _ = v21
                }
            }
        }
    }
    switch (("\n\t" + "\n\t") + "a") {
    case "hello":
        fmt.Println(v5, v5)
    default:
        for fi22 := 0; fi22 < 5; fi22++ {
            var v23 int32 = -746
            _ = fi22
            _ = v23
        }
    }
    v24 := GenS3{F0: ("" + "hello"), F1: 861, F2: 255}
    if (!false) {
        var v25 uint32 = 1
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 int8 = -98
            for fi28 := 0; fi28 < 3; fi28++ {
                v29 := GenS5{F0: 1, F1: "", F2: 2}
                v1 = (-2.25 - v1)
                switch ("hello" + ("a" + "")) {
                case "foo bar":
                    if true {
                        switch "a" {
                        case "":
                            switch fi26 {
                            case 128:
                                v27 += -1
                            }
                        default:
                            var v30 int64 = 462
                            _ = v30
                        }
                    } else {
                        v4[-9223372036854775808] = "hello"
                        fmt.Println(v25)
                        v1 = ((v1 + (v1 / 0.1)) + float64(1e10))
                    }
                case "a":
                    if (!v5) {
                        var v31 bool = v5
                        if ("hello" < "hello") {
                            if (0 > 32767) {
                                switch fi28 {
                                case -517:
                                    v25 += (^(v25))
                                default:
                                    v25 = v25
                                }
                                var v32 bool = v31
                                _ = len(v4)
                                v32 = (((!(v25 < 1)) || false) && v31)
                                _ = v32
                            }
                            var v33 uint16 = 1
                            _ = v33
                        } else {
                            v34 := GenS6{F0: "\n\t", F1: fi28, F2: 42}
                            v31 = v31
                            _ = v34
                        }
                        fmt.Println(v5, fi26, v25)
                        _ = v31
                    }
                    var v35 float32 = (float32((v1 + 1e10)) * float32(v1))
                    var v36 float64 = 1e10
                    _ = v35
                    _ = v36
                case "":
                    var v37 uint64 = 808
                    _ = v37
                default:
                    v25 -= (v25 + 42)
                }
                _ = fi28
                _ = v29
            }
            _ = fi26
            _ = v27
        }
        v5 = ("\n\t" > "\n\t")
        _ = v25
    } else {
        switch ("\n\t" + "hello") {
        case "a":
            for fi38 := 0; fi38 < 14; fi38++ {
                v39 := GenS6{F0: "\n\t", F1: fi38, F2: 0}
                _ = fi38
                _ = v39
            }
        case "foo bar":
            v1 = (0.0 * v1)
        }
    }
    var v40 int8 = 42
    switch 4294967295 {
    case -897:
        var v41 int64 = 42
        _ = v41
    case -207:
        v1 = float64((float32(v1) - 0.1))
    }
    var v42 string = "a"
    v42 = ("foo bar" + ("\n\t" + (v42 + "a")))
    if v5 {
        fmt.Println(v5, v42)
        v1 += float64(0.1)
    } else {
        if len(v2) > 0 {
            v2[0] = 643
        }
    }
    for fi43 := 0; fi43 < 6; fi43++ {
        var v44 float32 = (1e10 / (0.0 + 3.14159))
        _ = fi43
        _ = v44
    }
    fmt.Println(v42, v5, v1)
    v45 := GenS1{F0: 9223372036854775807, F1: 42, F2: 2}
    var v46 uint16 = 255
    if len(v6) > 0 {
        v6[0] = 127
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
