package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 int16 }
func genFn0(p0 uint64, p1 int16, p2 uint) uint32 {
    var v6 string = (((("" + "foo bar") + "") + "a") + "\n\t")
    v7 := []int{0}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 419
}
func genFn1(p0 int8, p1 uint32, p2 uint32) uint16 {
    for fi8 := 0; fi8 < 11; fi8++ {
        switch "foo bar" {
        case "hello":
            var v9 int16 = -1
            _ = v9
        }
        _ = fi8
    }
    p0 = p0
    switch ("hello" + "hello") {
    case "":
        v10 := GenS1{F0: "", F1: genFn0(4294967295, 42, 1)}
        _ = v10
    case "foo bar":
        var v11 float64 = 0.1
        _ = v11
    case "a":
        switch 2147483647 {
        case 0:
            var v12 int = 127
            _ = v12
        }
    default:
        v13 := GenS2{F0: 1, F1: 4294967295}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = -518
    var v3 float32 = float32(3.14159)
    v4 := make(map[int]uint32)
    v5 := []int{-2147483648, -847, 42, -528}
    fmt.Println(v2)
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := v1[("foo bar" + "a")]
        if true {
            if (998 == 520) {
                var v16 uint8 = 2
                var v17 uint16 = 2
                _ = v16
                _ = v17
            }
            var v18 uint64 = 4294967295
            var v19 string = "\n\t"
            _ = v18
            _ = v19
        }
        _ = fi14
        _ = v15
    }
    v20 := GenS2{F0: 100, F1: -1}
    _ = len(v1)
    if false {
        fmt.Println(v3, v3)
        v2 -= -69
        v3 = v3
    } else {
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 string = "\n\t"
            var v23 uint16 = 1
            _ = len(v4)
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    var v24 uint64 = 2
    for fi25 := 0; fi25 < 3; fi25++ {
        if false {
            v26 := GenS2{F0: 127, F1: fi25}
            switch "\n\t" {
            case "foo bar":
                v24 = (v24 ^ v24)
            case "":
                switch ("hello" + ("" + ("foo bar" + "\n\t"))) {
                case "hello":
                    if true {
                        var v27 int8 = -58
                        var v28 uint = 2
                        _ = v27
                        _ = v28
                    }
                }
            case "hello":
                var v29 bool = false
                _ = v29
            }
            _ = v26
        } else {
            fmt.Println(v2)
            if true {
                var v30 uint64 = 2
                for fi31 := 0; fi31 < 8; fi31++ {
                    var v32 string = ("" + ("foo bar" + ("" + "")))
                    v2 += 0
                    _ = fi31
                    _ = v32
                }
                v3 += v3
                _ = v30
            }
        }
        v3 += v3
        _ = fi25
    }
    fmt.Println(v2, v3)
    switch 0 {
    case -301:
        v33 := GenS3{F0: 141}
        _ = v33
    case 9223372036854775807:
        v34 := GenS3{F0: 248}
        _ = v34
    default:
        var v35 uint16 = 281
        _ = v35
    }
    v36 := GenS4{F0: v2, F1: 42}
    var v37 int = -184
    v2 = v2
    fmt.Println(v2)
    if false {
        var v38 uint16 = 42
        _ = len(v5)
        _ = v38
    } else {
        fmt.Println(v37, v37, v3)
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        var v40 bool = false
        var v41 int = v37
        fmt.Println(v40, v24)
        _ = fi39
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        if true {
            for fi43 := 0; fi43 < 13; fi43++ {
                if len(v5) > 0 {
                    v5[0] = fi43
                }
                delete(v1, "foo bar")
                switch (v37 + 127) {
                case 849:
                    v24 = v24
                }
                _ = fi43
            }
            v2 = (v2 | int64(fi42))
            if false {
                var v44 uint16 = 363
                delete(v4, 32767)
                _ = v44
            } else {
                for fi45 := 0; fi45 < 6; fi45++ {
                    v3 = 3.14159
                    var v46 bool = false
                    _ = fi45
                    _ = v46
                }
            }
        } else {
            v24 = 18446744073709551615
        }
        _ = len(v1)
        _ = fi42
    }
    v24 = (^(v24))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
