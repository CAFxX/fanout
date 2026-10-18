package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 string }

func main() {
    v1 := make(map[string]string)
    var v2 string = (("a" + "a") + "\n\t")
    v3 := []uint32{2147483648, 201}
    if false {
        v2 = v2
    }
    switch (v2 + ("hello" + "a")) {
    case "\n\t":
        v2 = v2
    case "":
        if false {
            var v4 bool = true
            v4 = v4
            v5 := GenS1{F0: 42, F1: 65535}
            _ = v4
            _ = v5
        }
    default:
        switch 4294967295 {
        case 42:
            var v6 bool = false
            _ = v6
        case 32767:
            if true {
                _ = len(v1)
                v2 = v2
                v2 = "a"
            } else {
                for fi7 := 0; fi7 < 18; fi7++ {
                    for fi8 := 0; fi8 < 9; fi8++ {
                        var v9 string = ("foo bar" + (((v2 + v2) + v2) + v2))
                        fmt.Println(fi7, fi7, fi7)
                        _ = fi8
                        _ = v9
                    }
                    var v10 float64 = float64(1e10)
                    _ = fi7
                    _ = v10
                }
                var v11 float64 = -2.25
                _ = v11
            }
        case -2147483648:
            for fi12 := 0; fi12 < 9; fi12++ {
                v2 = "foo bar"
                v2 = "foo bar"
                fmt.Println(fi12)
                _ = fi12
            }
        }
    }
    switch v2 {
    case "foo bar":
        var v13 int16 = 42
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS1{F0: 255, F1: 2147483647}
        _ = fi14
        _ = v15
    }
    var v16 string = "hello"
    v17 := GenS2{F0: 1}
    var v18 int64 = 42
    fmt.Println(v2)
    if (4294967295 >= 324) {
        for fi19 := 0; fi19 < 8; fi19++ {
            for fi20 := 0; fi20 < 3; fi20++ {
                v21 := GenS3{F0: (570 >= 321), F1: 1}
                var v22 uint16 = 213
                v23 := GenS2{F0: -1}
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            _ = fi19
        }
    } else {
        for fi24 := 0; fi24 < 6; fi24++ {
            switch int(v18) {
            case -1:
                v25 := GenS1{F0: 39, F1: 2}
                _ = v25
            case 456:
                v26 := GenS2{F0: 4294967295}
                _ = v26
            default:
                v2 = (v16 + v2)
            }
            _ = fi24
        }
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        var v28 float32 = 0.1
        fmt.Println(v18)
        if false {
            if ((-2.25 == 0.0) || false) {
                var v29 uint16 = 2
                var v30 uint = 42
                fmt.Println(v29, v28, fi27)
                _ = v29
                _ = v30
            }
            var v31 uint64 = 855
            _ = v31
        } else {
            fmt.Println(v16)
            var v32 int16 = -794
            _ = v32
        }
        _ = fi27
        _ = v28
    }
    var v33 uint = 42
    v34 := GenS4{F0: ("\n\t" + (v16 + "a"))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
}
