package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint; F2 int }
type GenS4 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint32) int16 {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 float64 = ((-2.25 - -2.25) / ((-2.25 / -2.25) - (3.14159 - 1.5)))
        fmt.Println(p0)
        var v8 int = (fi6 % (((fi6 & (fi6 / ((fi6) | 1)))) | 1))
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 1 {
    case 127:
        v9 := []int{128}
        _ = v9
    default:
        v10 := []int{128, 966, 42}
        _ = v10
    }
    switch "hello" {
    case "a":
        if (!true) {
            var v11 uint32 = p0
            fmt.Println(v11)
            v12 := GenS1{F0: 142, F1: -127}
            _ = v11
            _ = v12
        }
    case "foo bar":
        if (!true) {
            v13 := GenS2{F0: 255}
            _ = v13
        } else {
            switch -700 {
            case 4294967295:
                switch 32768 {
                case 42:
                    var v14 int64 = -523
                    _ = v14
                case 0:
                    if true {
                        var v15 uint16 = 255
                        v16 := GenS1{F0: 419, F1: 405}
                        _ = v15
                        _ = v16
                    }
                default:
                    switch -2147483648 {
                    case 927:
                        var v17 int = 42
                        _ = v17
                    }
                }
            case 0:
                var v18 float32 = (0.0 * (1e10 + (1e10 * (float32((-2.25 / 1e10)) * (1.5 * -2.25)))))
                _ = v18
            }
            var v19 int16 = -32768
            v20 := []int{769, -504, -152, 2147483647}
            _ = v19
            _ = v20
        }
    case "\n\t":
        v21 := GenS2{F0: 0}
        _ = v21
    default:
        var v22 int16 = 32767
        _ = v22
    }
    v23 := []int{9223372036854775807}
    fmt.Println(p0)
    _ = p0
    _ = v23
    return 1
}

func main() {
    var v1 int16 = -32768
    v2 := []int{970, 1, 782, 1}
    var v3 string = "hello"
    v4 := []uint32{42, 42}
    v5 := []uint32{435}
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    switch v3 {
    case "foo bar":
        var v24 uint16 = 1
        _ = v24
    default:
        for fi25 := 0; fi25 < 10; fi25++ {
            switch ((("" + "a") + v3) + v3) {
            case "hello":
                var v26 int = fi25
                _ = v26
            case "":
                var v27 int64 = -348
                _ = v27
            default:
                for fi28 := 0; fi28 < 14; fi28++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v29 := GenS3{F0: v3, F1: 65535, F2: 42}
                    v3 = (((("hello" + "a") + (v3 + v3)) + "a") + "hello")
                    if false {
                        v3 = "hello"
                    } else {
                        var v30 float64 = -2.25
                        v30 -= 1.5
                        var v31 uint64 = 132
                        _ = v30
                        _ = v31
                    }
                    var v32 uint64 = 0
                    var v33 uint64 = v32
                    _ = fi28
                    _ = v29
                    _ = v32
                    _ = v33
                }
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi25
        }
    }
    fmt.Println(v1, v3)
    for fi34 := 0; fi34 < 4; fi34++ {
        for fi35 := 0; fi35 < 12; fi35++ {
            v3 = ("" + (("a" + v3) + v3))
            _ = fi35
        }
        _ = fi34
    }
    var v36 int = 127
    _ = len(v4)
    fmt.Println(v36)
    var v37 string = v3
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        var v39 string = (v37 + "")
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi38
        _ = v39
    }
    if true {
        for fi40 := 0; fi40 < 2; fi40++ {
            for fi41 := 0; fi41 < 5; fi41++ {
                if len(v2) > 0 {
                    v2[0] = (fi41 >> ((1) & 63))
                }
                switch (v37 + "hello") {
                case "foo bar":
                    v42 := GenS3{F0: v37, F1: 844, F2: (-(fi40))}
                    _ = v42
                }
                v4 = append(v4, 0)
                _ = fi41
            }
            _ = fi40
        }
        v43 := GenS2{F0: 4294967295}
        switch (-(v36)) {
        case -9223372036854775808:
            fmt.Println(v1, v1, v37)
        default:
            v37 = ("hello" + "hello")
        }
        _ = v43
    } else {
        fmt.Println(v37, v36)
        if len(v2) > 0 {
            v2[0] = (v36 + (v36 * (v36 - 175)))
        }
    }
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    v4 = append(v4, 42)
    if false {
        v44 := GenS4{F0: 305, F1: 127}
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
