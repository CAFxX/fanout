package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 uint32 }

func main() {
    var v1 int = 81
    v2 := make(map[uint32]string)
    var v3 uint8 = 42
    fmt.Println(v1)
    v2[2147483647] = "a"
    switch ("\n\t" + "hello") {
    case "foo bar":
        var v4 int64 = 0
        _ = v4
    case "\n\t":
        var v5 string = "a"
        _ = v5
    case "":
        switch "" {
        case "hello":
            v6 := GenS1{F0: 0, F1: 2147483647}
            _ = v6
        case "\n\t":
            fmt.Println(v3, v1)
        case "":
            switch (v1 - (^(v1))) {
            case 4294967295:
                delete(v2, 1)
            case -289:
                for fi7 := 0; fi7 < 7; fi7++ {
                    for fi8 := 0; fi8 < 3; fi8++ {
                        if (2 > 65535) {
                            var v9 bool = ("\n\t" == "hello")
                            if true {
                                v9 = true
                                switch (-(v1)) {
                                case 9223372036854775807:
                                    fmt.Println(fi7, v1)
                                case 301:
                                    v1 += (^(fi7))
                                case 0:
                                    var v10 string = ((("" + "\n\t") + ("hello" + "\n\t")) + "a")
                                    _ = v10
                                }
                                v11 := []int{}
                                _ = v11
                            } else {
                                fmt.Println(v3)
                                v12 := GenS1{F0: 600, F1: -710}
                                v9 = true
                                _ = v12
                            }
                            _ = v9
                        } else {
                            if true {
                                fmt.Println(fi7, v1, v3)
                                v13 := GenS2{F0: 4294967295, F1: 42}
                                if ((0.1 >= ((1e10 + 1e10) + float64((-2.25 / 1e10)))) && ((!true) || (false || false))) {
                                    var v14 uint = 42
                                    if (1.5 != -2.25) {
                                        var v15 int8 = 127
                                        switch (((("hello" + "") + "a") + "hello") + "\n\t") {
                                        case "":
                                            v1 += (fi7 >> ((1) & 63))
                                        }
                                        _ = v15
                                    }
                                    _ = v14
                                } else {
                                    if false {
                                        switch fi8 {
                                        case -308:
                                            if ((((!true) && false) || true) && (true && false)) {
                                                fmt.Println(v1, fi7, v1)
                                                v2[479] = ("foo bar" + (("a" + "hello") + "\n\t"))
                                                switch "foo bar" {
                                                case "foo bar":
                                                    if (-1 == 42) {
                                                        if true {
                                                            var v16 uint = 47
                                                            _ = v16
                                                        }
                                                        fmt.Println(fi7, fi8)
                                                    }
                                                }
                                            } else {
                                                var v17 uint32 = 65535
                                                delete(v2, v17)
                                                _ = v17
                                            }
                                        case 1:
                                            if ((255 >= 32767) || false) {
                                                switch 4294967295 {
                                                case -506:
                                                    v3 -= v3
                                                case 946:
                                                    v18 := []int{32768}
                                                    _ = v18
                                                case 9223372036854775807:
                                                    var v19 uint32 = 0
                                                    _ = v19
                                                }
                                                fmt.Println(fi7, v3, fi7)
                                                if ((2 == 1) && true) {
                                                    v20 := []int{(v1 * 1)}
                                                    _ = v20
                                                } else {
                                                    fmt.Println(fi8)
                                                }
                                            }
                                        case 32768:
                                            v21 := GenS3{F0: 0, F1: "a"}
                                            _ = v21
                                        }
                                        var v22 bool = true
                                        _ = v22
                                    }
                                }
                                _ = v13
                            } else {
                                var v23 uint = 42
                                v1 -= fi7
                                _ = v23
                            }
                        }
                        switch fi7 {
                        case -727:
                            if false {
                                var v24 uint32 = 0
                                var v25 int16 = -1
                                v2[v24] = (("hello" + ("" + ("" + "foo bar"))) + "hello")
                                _ = v24
                                _ = v25
                            } else {
                                v26 := GenS3{F0: 814, F1: "\n\t"}
                                _ = v26
                            }
                        case 9223372036854775807:
                            fmt.Println(v1, v1, v3)
                        case 1:
                            fmt.Println(fi8, v3)
                        }
                        _ = fi8
                    }
                    if (0.1 > 0.1) {
                        for fi27 := 0; fi27 < 17; fi27++ {
                            var v28 int64 = 42
                            fmt.Println(fi27)
                            v2[2] = ("" + "")
                            if (!false) {
                                var v29 int8 = -125
                                var v30 uint = 1
                                _ = v29
                                _ = v30
                            }
                            _ = fi27
                            _ = v28
                        }
                        v1 = (fi7 - fi7)
                    } else {
                        var v31 int8 = 127
                        v3 += v3
                        _ = v31
                    }
                    _ = fi7
                }
            case 190:
                for fi32 := 0; fi32 < 3; fi32++ {
                    v1 = (^(fi32))
                    _ = fi32
                }
            default:
                if false {
                    fmt.Println(v3, v1)
                } else {
                    if false {
                        var v33 uint64 = 707
                        _ = v33
                    }
                    fmt.Println(v3, v1)
                    v34 := GenS4{F0: 42}
                    _ = v34
                }
            }
        default:
            v3 = v3
        }
    }
    var v35 string = "a"
    var v36 int8 = 1
    var v37 float32 = ((0.0 + (3.14159 * 0.1)) + 3.14159)
    fmt.Println(v3, v1, v3)
    if false {
        v38 := GenS5{F0: -2147483648, F1: 312}
        var v39 int = -1
        var v40 uint64 = 4294967295
        _ = v38
        _ = v39
        _ = v40
    }
    v37 = 1e10
    v37 += -2.25
    _ = len(v2)
    v3 = v3
    var v41 int8 = v36
    switch (v1 ^ (v1 + v1)) {
    case 1:
        var v42 uint64 = 521
        _ = v42
    default:
        for fi43 := 0; fi43 < 10; fi43++ {
            _ = len(v2)
            _ = fi43
        }
    }
    var v44 uint8 = v3
    for fi45 := 0; fi45 < 18; fi45++ {
        switch v35 {
        case "\n\t":
            v35 = (v35 + v35)
        default:
            v46 := GenS2{F0: 0, F1: -1}
            _ = v46
        }
        _ = fi45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v41
    _ = v44
}
