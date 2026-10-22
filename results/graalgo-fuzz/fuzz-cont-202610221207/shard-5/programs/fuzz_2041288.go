package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint64 }
func genFn0(p0 bool) uint32 {
    var v6 int = 986
    var v7 int32 = 0
    _ = p0
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 float64 = (1e10 * 1e10)
    var v2 float64 = (1e10 * v1)
    var v3 int64 = 245
    v4 := make(map[int]int)
    v5 := make(map[uint32]bool)
    if false {
        v1 = 0.1
    } else {
        v8 := []int{208, 4294967295, -1, -389}
        _ = v8
    }
    fmt.Println(v1, v1, v3)
    var v9 string = ("" + ((("\n\t" + "foo bar") + "a") + "a"))
    v10 := GenS1{F0: 65535, F1: 42}
    v11 := GenS1{F0: 2147483648, F1: 0}
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint8 = 0
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 string = (v9 + v9)
            var v16 uint32 = 2147483647
            v15 = v15
            var v17 int16 = 1
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        var v18 uint8 = 2
        _ = fi12
        _ = v13
        _ = v18
    }
    var v19 float32 = ((-2.25 - ((float32(float64(-2.25)) * -2.25) / 0.1)) + ((3.14159 + (float32(float64(0.0)) * ((3.14159 * 3.14159) + 0.0))) / ((-2.25 * (1e10 + float32(1.5))) - 3.14159)))
    if true {
        for fi20 := 0; fi20 < 4; fi20++ {
            v19 = ((v19 + 3.14159) + (v19 - (0.1 - 0.0)))
            _ = fi20
        }
        var v21 int16 = 0
        var v22 int = -1
        _ = v21
        _ = v22
    }
    v23 := []int{-431, -388, 112, -9223372036854775808}
    if true {
        v2 += v2
        switch (("hello" + "") + ((v9 + "hello") + v9)) {
        case "":
            v24 := GenS1{F0: 4294967295, F1: 1}
            _ = v24
        case "a":
            switch 4294967295 {
            case 1:
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
            case 32767:
                var v25 uint16 = 255
                _ = v25
            default:
                var v26 uint8 = 1
                _ = v26
            }
        case "\n\t":
            v27 := GenS2{F0: 800}
            _ = v27
        default:
            switch ("" + (v9 + "hello")) {
            case "a":
                for fi28 := 0; fi28 < 19; fi28++ {
                    fmt.Println(v2, v3, v2)
                    _ = fi28
                }
            case "":
                switch 2147483647 {
                case -799:
                    for fi29 := 0; fi29 < 13; fi29++ {
                        fmt.Println(v19, v19, fi29)
                        _ = fi29
                    }
                case 0:
                    for fi30 := 0; fi30 < 10; fi30++ {
                        v31 := GenS2{F0: 0}
                        _ = fi30
                        _ = v31
                    }
                }
            case "\n\t":
                for fi32 := 0; fi32 < 7; fi32++ {
                    var v33 string = v9
                    _ = fi32
                    _ = v33
                }
            }
        }
        for fi34 := 0; fi34 < 6; fi34++ {
            v19 = 1e10
            v35 := GenS1{F0: 358, F1: -1}
            if true {
                v3 -= (^(v3))
            }
            _ = fi34
            _ = v35
        }
    }
    var v36 float64 = (v1 + v1)
    switch "\n\t" {
    case "a":
        v37 := GenS3{F0: v9}
        _ = v37
    case "hello":
        if (!(232 <= 2)) {
            delete(v4, 2147483647)
            v38 := GenS4{F0: -32768, F1: 582}
            _ = v38
        }
    }
    var v39 uint16 = 1
    switch 9223372036854775807 {
    case -426:
        v40 := GenS1{F0: 0, F1: -9223372036854775808}
        _ = v40
    }
    if (!true) {
        if true {
            for fi41 := 0; fi41 < 18; fi41++ {
                fmt.Println(v3)
                fmt.Println(v3, v19)
                v1 = 0.1
                _ = fi41
            }
        }
        for fi42 := 0; fi42 < 15; fi42++ {
            if ((779 == 2147483647) && true) {
                var v43 uint8 = 127
                for fi44 := 0; fi44 < 6; fi44++ {
                    if (-700 <= 934) {
                        v36 -= float64(v19)
                        var v45 float64 = float64((0.0 * v19))
                        _ = v45
                    } else {
                        var v46 int16 = int16(v39)
                        v1 = 1.5
                        _ = v46
                    }
                    var v47 string = "hello"
                    var v48 int = (fi44 & fi44)
                    var v49 bool = (!(!false))
                    _ = fi44
                    _ = v47
                    _ = v48
                    _ = v49
                }
                _ = len(v23)
                _ = v43
            } else {
                fmt.Println(v2, fi42, v2)
                if (!(-1 < 42)) {
                    if true {
                        v3 = (v3 / ((v3) | 1))
                    }
                }
                v1 += float64(1.5)
            }
            _ = len(v5)
            _ = fi42
        }
        fmt.Println(v36)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v9
}
