package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 string }
type GenS2 struct { F0 int32; F1 int64; F2 int16 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint = 255
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 65535
    var v5 uint8 = 42
    v6 := make(map[string]bool)
    var v7 uint = (v1 * v1)
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := []int{fi8, -1, (^(fi8)), 4294967295}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int = fi10
        switch (("\n\t" + "hello") + (("hello" + "a") + "a")) {
        case "foo bar":
            fmt.Println(v2, fi10)
        case "a":
            switch "" {
            case "":
                for fi12 := 0; fi12 < 9; fi12++ {
                    v13 := GenS1{F0: (v2 & 42), F1: -254, F2: ("a" + (("" + "foo bar") + "hello"))}
                    v4 += (^(v4))
                    _ = fi12
                    _ = v13
                }
            default:
                fmt.Println(fi10)
            }
        case "\n\t":
            switch "" {
            case "\n\t":
                _ = len(v3)
            case "":
                fmt.Println(v4)
            case "foo bar":
                var v14 float64 = (-2.25 * (3.14159 * (((3.14159 + 0.0) - 0.0) - (0.0 + -2.25))))
                _ = v14
            }
        default:
            v15 := v3[v4]
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    var v16 uint8 = 255
    if (140 < int(v2)) {
        fmt.Println(v7, v16)
    }
    switch "hello" {
    case "a":
        _ = len(v3)
    }
    var v17 uint16 = 1
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 uint64 = 1
        v20 := GenS2{F0: -2147483648, F1: v2, F2: 1}
        _ = fi18
        _ = v19
        _ = v20
    }
    v1 += 2
    for fi21 := 0; fi21 < 13; fi21++ {
        delete(v3, (^(v4)))
        var v22 float64 = ((0.1 - (0.0 - 3.14159)) - (-2.25 - 0.1))
        v23 := GenS3{F0: v16}
        _ = fi21
        _ = v22
        _ = v23
    }
    switch "a" {
    case "foo bar":
        v24 := GenS1{F0: (v2 + v2), F1: 1, F2: "hello"}
        _ = v24
    case "":
        var v25 string = (("hello" + "\n\t") + "foo bar")
        _ = v25
    case "\n\t":
        if (!false) {
            for fi26 := 0; fi26 < 12; fi26++ {
                var v27 int16 = int16(v1)
                v28 := GenS2{F0: 0, F1: (v2 + (v2 << ((v7) & 63))), F2: v27}
                _ = fi26
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 12; fi29++ {
                switch (fi29 | fi29) {
                case -2147483648:
                    delete(v6, "")
                case -255:
                    fmt.Println(v4, fi29)
                case 0:
                    if ((false && true) && false) {
                        v30 := GenS2{F0: 936, F1: v2, F2: 0}
                        for fi31 := 0; fi31 < 20; fi31++ {
                            switch fi31 {
                            case 9223372036854775807:
                                if true {
                                    v32 := []int{42}
                                    v32 = append(v32, fi29)
                                    if (!((false || (((1.5 + 3.14159) * (3.14159 - 0.0)) < (1e10 - 0.1))) && (true && (((1.5 + (3.14159 * 3.14159)) + 0.1) < 0.0)))) {
                                        v32 = append(v32, 9223372036854775807)
                                        v32 = append(v32, fi31)
                                    }
                                    _ = v32
                                } else {
                                    v33 := GenS2{F0: -1, F1: v2, F2: 0}
                                    _ = v33
                                }
                            case 128:
                                v16 -= v16
                            }
                            if true {
                                v34 := GenS2{F0: 0, F1: 755, F2: -564}
                                switch (^(fi29)) {
                                case 32767:
                                    if true {
                                        switch "hello" {
                                        case "a":
                                            var v35 string = ("\n\t" + ("\n\t" + "a"))
                                            _ = v35
                                        case "foo bar":
                                            var v36 int64 = 42
                                            _ = v36
                                        case "":
                                            fmt.Println(v17, v7, v5)
                                        }
                                        if false {
                                            if (!("foo bar" != ("a" + "hello"))) {
                                                fmt.Println(v16)
                                            }
                                            v7 -= (v7 / ((849) | 1))
                                        }
                                    }
                                default:
                                    fmt.Println(fi31, v7)
                                }
                                if false {
                                    var v37 uint8 = (v5 % ((v5) | 1))
                                    _ = v37
                                }
                                _ = v34
                            }
                            _ = fi31
                        }
                        for fi38 := 0; fi38 < 7; fi38++ {
                            if false {
                                switch "" {
                                case "":
                                    fmt.Println(fi29, v1, v16)
                                case "hello":
                                    var v39 uint32 = v4
                                    _ = v39
                                }
                                switch "hello" {
                                case "foo bar":
                                    if (0 <= 98) {
                                        v40 := v6[""]
                                        fmt.Println(v40, v5)
                                        _ = v40
                                    }
                                case "hello":
                                    v6["hello"] = ("foo bar" <= "a")
                                }
                                var v41 uint16 = v17
                                switch ("a" + "hello") {
                                case "foo bar":
                                    v42 := []int{int(v2), (fi29 << ((v7) & 63))}
                                    _ = v42
                                }
                                delete(v3, 2147483647)
                                _ = v41
                            }
                            var v43 uint32 = v4
                            _ = fi38
                            _ = v43
                        }
                        _ = v30
                    }
                }
                _ = fi29
            }
            v44 := []int{-253}
            _ = v44
        }
    default:
        for fi45 := 0; fi45 < 14; fi45++ {
            v46 := []int{(fi45 + (fi45 % (((fi45 * 2147483647)) | 1)))}
            v2 -= v2
            var v47 int64 = (v2 ^ (^(v2)))
            _ = fi45
            _ = v46
            _ = v47
        }
    }
    switch 2147483647 {
    case -123:
        v4 = uint32(v17)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
