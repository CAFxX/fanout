package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint16; F2 string }
type GenS3 struct { F0 bool; F1 string; F2 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = -2.25
    var v3 float64 = 1.5
    v4 := []int{-674, -51}
    fmt.Println(v2)
    var v5 string = "hello"
    v5 = v5
    if len(v4) > 0 {
        v4[0] = 127
    }
    var v6 uint8 = 127
    _ = len(v4)
    var v7 float64 = (v2 * (((((v3 / v3) - v3) - (v3 / ((3.14159 + (v2 / 3.14159)) * ((3.14159 + ((0.1 - 0.1) - v2)) * float64(1.5))))) + 0.0) / (v2 - v3)))
    if (168 > 4294967295) {
        for fi8 := 0; fi8 < 19; fi8++ {
            switch (fi8 + 9223372036854775807) {
            case 734:
                v9 := GenS1{F0: (-405 >= -245)}
                _ = v9
            case -553:
                v10 := GenS1{F0: (v3 < v2)}
                _ = v10
            default:
                for fi11 := 0; fi11 < 18; fi11++ {
                    v12 := GenS2{F0: 1, F1: 2, F2: (v5 + v5)}
                    var v13 uint = 836
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            _ = fi8
        }
        if len(v4) > 0 {
            v4[0] = -1
        }
        _ = len(v1)
    } else {
        var v14 uint8 = 42
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v15 uint64 = 1
        _ = v14
        _ = v15
    }
    fmt.Println(v5, v7)
    switch -1 {
    case -9223372036854775808:
        v5 = (v5 + v5)
    default:
        v4 = append(v4, 2147483647)
    }
    if false {
        v16 := GenS3{F0: false, F1: "hello", F2: 18446744073709551615}
        if true {
            for fi17 := 0; fi17 < 20; fi17++ {
                v18 := GenS3{F0: (false && (-54 >= -1)), F1: (("a" + (("\n\t" + "foo bar") + (v5 + "hello"))) + (("" + (v5 + "foo bar")) + ("" + "a"))), F2: 42}
                var v19 uint32 = 845
                _ = fi17
                _ = v18
                _ = v19
            }
            switch (v5 + v5) {
            case "a":
                v6 = v6
            case "hello":
                delete(v1, 0)
            case "":
                v5 = ("hello" + ("\n\t" + "\n\t"))
            default:
                v4 = append(v4, 0)
            }
        } else {
            if true {
                if true {
                    switch (("hello" + "") + v5) {
                    case "hello":
                        fmt.Println(v5, v2)
                    case "a":
                        v7 = (0.0 - (v2 * 1.5))
                    case "foo bar":
                        delete(v1, 1)
                    }
                }
                var v20 string = "\n\t"
                _ = v20
            } else {
                for fi21 := 0; fi21 < 18; fi21++ {
                    var v22 int32 = 1
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v2 += float64((1.5 / 3.14159))
                        var v24 float32 = 1e10
                        v6 = v6
                        _ = fi23
                        _ = v24
                    }
                    _ = fi21
                    _ = v22
                }
                switch "foo bar" {
                case "a":
                    var v25 uint = 616
                    _ = v25
                case "foo bar":
                    var v26 int16 = 42
                    _ = v26
                case "hello":
                    var v27 int = 1
                    _ = v27
                }
                v28 := GenS3{F0: false, F1: ("foo bar" + v5), F2: 1}
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 1; fi29++ {
            for fi30 := 0; fi30 < 4; fi30++ {
                v31 := GenS1{F0: false}
                v1[2147483648] = 2147483648
                v6 -= v6
                _ = fi30
                _ = v31
            }
            var v32 uint64 = 18446744073709551615
            for fi33 := 0; fi33 < 16; fi33++ {
                fmt.Println(fi33, fi29)
                _ = fi33
            }
            _ = fi29
            _ = v32
        }
        _ = v16
    }
    switch ("" + v5) {
    case "foo bar":
        var v34 int = 242
        _ = v34
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        var v36 uint16 = 42
        v6 = v6
        v5 = v5
        _ = fi35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
