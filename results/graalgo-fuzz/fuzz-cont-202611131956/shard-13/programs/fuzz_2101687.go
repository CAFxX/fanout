package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint16; F2 int16 }
func genFn0(p0 uint, p1 int32) uint16 {
    v4 := GenS1{F0: 2147483647, F1: 127, F2: -128}
    p1 = (^(p1))
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 int8 = -128
            var v8 int16 = -840
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 int8 = 0
        var v10 bool = (127 <= 234)
        _ = fi5
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    return 115
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[uint32]string)
    var v3 float64 = -2.25
    switch ("hello" + "hello") {
    case "":
        v11 := GenS1{F0: 912, F1: (v1 | 0), F2: -30}
        _ = v11
    default:
        if (true && (!(42 < 1000))) {
            var v12 int64 = -81
            var v13 uint16 = 0
            _ = v12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        if true {
            v1 += v1
            if true {
                var v15 int32 = -362
                _ = v15
            }
        } else {
            v16 := []int{42, fi14}
            for fi17 := 0; fi17 < 7; fi17++ {
                v3 = (3.14159 * float64(((3.14159 - 0.1) - 1.5)))
                fmt.Println(fi14, v3)
                v3 = 1e10
                _ = fi17
            }
            _ = v16
        }
        v18 := GenS1{F0: 190, F1: (v1 & (^(v1))), F2: -1}
        _ = fi14
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        switch (((("a" + "hello") + "") + ("foo bar" + "a")) + (((("foo bar" + "a") + "\n\t") + ("\n\t" + "a")) + "hello")) {
        case "foo bar":
            if ("a" >= "a") {
                v1 -= v1
                switch fi19 {
                case 42:
                    _ = len(v2)
                }
            }
        case "hello":
            var v20 uint64 = 467
            _ = v20
        case "a":
            v21 := GenS1{F0: 0, F1: (^(v1)), F2: -57}
            _ = v21
        default:
            for fi22 := 0; fi22 < 7; fi22++ {
                fmt.Println(v1, v3, fi22)
                var v23 bool = true
                _ = fi22
                _ = v23
            }
        }
        _ = fi19
    }
    if (42 > 0) {
        v1 -= 146
        var v24 int = 4294967295
        _ = v24
    } else {
        v3 = ((v3 * 1e10) + (-2.25 - v3))
        switch -1 {
        case -2147483648:
            for fi25 := 0; fi25 < 18; fi25++ {
                fmt.Println(v1)
                _ = fi25
            }
        default:
            fmt.Println(v3)
        }
        v1 -= 255
    }
    v26 := GenS2{F0: 65535, F1: 216}
    v27 := GenS3{F0: 2}
    v28 := GenS4{F0: 42}
    v29 := []int{-716, -2147483648, 128, -9223372036854775808}
    fmt.Println(v3, v1)
    v30 := GenS5{F0: -1, F1: 175, F2: 1000}
    v31 := GenS4{F0: 446}
    switch ((("foo bar" + ("" + "a")) + ((("foo bar" + ("" + ("hello" + "\n\t"))) + "\n\t") + "foo bar")) + "\n\t") {
    case "a":
        for fi32 := 0; fi32 < 19; fi32++ {
            v1 = v1
            _ = fi32
        }
    case "":
        v1 += 2
    case "hello":
        switch "a" {
        case "\n\t":
            v1 += 60
        case "":
            _ = len(v29)
        default:
            fmt.Println(v3, v1)
        }
    default:
        _ = len(v29)
    }
    if false {
        if true {
            if true {
                switch 307 {
                case -2147483648:
                    var v33 float64 = v3
                    _ = v33
                case 128:
                    v29 = append(v29, 32767)
                }
                fmt.Println(v1, v3)
            }
            if true {
                _ = len(v2)
            } else {
                var v34 uint = 0
                fmt.Println(v3)
                _ = v34
            }
        } else {
            var v35 uint64 = uint64(v1)
            var v36 int = -791
            v3 -= ((v3 / (v3 - v3)) + 1.5)
            _ = v35
            _ = v36
        }
        for fi37 := 0; fi37 < 15; fi37++ {
            for fi38 := 0; fi38 < 4; fi38++ {
                var v39 uint8 = 2
                var v40 float32 = float32((float64(0.1) - float64(1.5)))
                v40 -= (3.14159 * v40)
                switch (("a" + "hello") + "hello") {
                case "foo bar":
                    fmt.Println(v40)
                }
                _ = fi38
                _ = v39
                _ = v40
            }
            var v41 uint8 = 127
            _ = fi37
            _ = v41
        }
        fmt.Println(v3, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
