package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int64 }
func genFn0() uint {
    var v5 float32 = (3.14159 + 1e10)
    if (!true) {
        switch "a" {
        case "hello":
            v6 := []int{454, 32767}
            _ = v6
        case "\n\t":
            for fi7 := 0; fi7 < 9; fi7++ {
                switch fi7 {
                case 32768:
                    v8 := []int{(fi7 / (((^(fi7))) | 1)), 32768, 9223372036854775807, fi7}
                    _ = v8
                }
                _ = fi7
            }
        case "":
            var v9 uint8 = 122
            _ = v9
        default:
            var v10 int16 = 1
            _ = v10
        }
        var v11 int8 = 1
        switch (("" + "a") + "foo bar") {
        case "":
            v11 += (v11 << ((4294967295) & 63))
        default:
            v12 := GenS1{F0: 2147483647, F1: 124, F2: -1}
            _ = v12
        }
        _ = v11
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        fmt.Println(fi13, v5)
        v5 = (float32((3.14159 / 0.1)) * float32((3.14159 / (3.14159 * (3.14159 + 1e10)))))
        _ = fi13
    }
    _ = v5
    return 255
}
func genFn1() bool {
    v14 := GenS1{F0: 96, F1: 255, F2: 0}
    if false {
        v15 := []int{0}
        var v16 int = 681
        _ = v15
        _ = v16
    }
    _ = v14
    return (1.5 <= (float32((1e10 / 1.0)) / (1e10 / 1.0)))
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := []uint32{2147483648, 792, 42}
    var v4 int16 = -316
    v4 -= v4
    v17 := GenS1{F0: 1, F1: 42, F2: -9223372036854775808}
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            v4 = 317
            if (-284 != 1) {
                var v20 int64 = -439
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    switch -9223372036854775808 {
                    case -579:
                        v20 += (-(v20))
                    }
                    v20 = (v20 & v20)
                }
                _ = v20
            } else {
                v4 -= 42
            }
            _ = fi19
        }
        var v21 uint64 = 42
        var v22 float64 = 0.1
        _ = fi18
        _ = v21
        _ = v22
    }
    delete(v2, ("hello" + "hello"))
    v4 = (v4 * -32768)
    if true {
        v4 = (-(v4))
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        switch "\n\t" {
        case "a":
            if false {
                if len(v3) > 0 {
                    v3[0] = 353
                }
            }
        case "foo bar":
            switch (fi23 << ((2) & 63)) {
            case 1:
                v4 = (v4 + int16(fi23))
            case -2147483648:
                if len(v3) > 0 {
                    v3[0] = 220
                }
            case 879:
                v4 = 0
            }
        }
        switch 4294967295 {
        case 1:
            if (-1 >= 81) {
                v4 = 1000
                var v24 float64 = 0.0
                _ = v24
            }
        case 9223372036854775807:
            switch (fi23 << ((1) & 63)) {
            case 1:
                fmt.Println(v4)
            }
        case -255:
            v3 = append(v3, 971)
        default:
            if (2 <= 4294967295) {
                v25 := v2["hello"]
                _ = v25
            }
        }
        var v26 string = ("" + "foo bar")
        _ = fi23
        _ = v26
    }
    var v27 int = 127
    switch ("" + "hello") {
    case "\n\t":
        if true {
            if (float32(0.0) < (1e10 + float32(0.1))) {
                switch "a" {
                case "\n\t":
                    v27 = -2147483648
                case "a":
                    v28 := v1["\n\t"]
                    _ = v28
                }
            } else {
                var v29 uint64 = 468
                switch "hello" {
                case "":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "a":
                    var v30 int64 = -880
                    _ = v30
                }
                var v31 uint = genFn0()
                _ = v29
                _ = v31
            }
            _ = len(v2)
        }
    case "":
        fmt.Println(v27, v27)
    case "foo bar":
        fmt.Println(v4, v27)
    default:
        var v32 uint8 = 255
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
