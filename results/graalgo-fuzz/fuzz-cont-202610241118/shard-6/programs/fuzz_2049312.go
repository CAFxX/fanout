package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 bool; F2 uint16 }
type GenS3 struct { F0 int }
func genFn0(p0 uint8, p1 int16) int {
    switch "\n\t" {
    case "hello":
        p0 += p0
    case "foo bar":
        for fi5 := 0; fi5 < 13; fi5++ {
            if false {
                for fi6 := 0; fi6 < 3; fi6++ {
                    v7 := GenS1{F0: 42}
                    if true {
                        switch (-(fi5)) {
                        case -1:
                            fmt.Println(p0, p1, fi6)
                        case 32768:
                            switch ("foo bar" + ("\n\t" + (("\n\t" + "foo bar") + "\n\t"))) {
                            case "\n\t":
                                v8 := []int{(fi5 ^ (^(fi5))), fi5, fi5, int(p0)}
                                _ = v8
                            default:
                                var v9 uint32 = 2147483648
                                _ = v9
                            }
                        case 4294967295:
                            var v10 float32 = (0.0 - -2.25)
                            _ = v10
                        default:
                            var v11 uint = 4294967295
                            _ = v11
                        }
                        var v12 uint = 2
                        v12 += v12
                        p1 = (p1 >> (((v12 + v12)) & 63))
                        p0 += p0
                        _ = v12
                    } else {
                        switch ("a" + "") {
                        case "":
                            v13 := []int{}
                            _ = v13
                        case "hello":
                            var v14 int8 = 1
                            _ = v14
                        default:
                            p0 = 49
                        }
                        switch (fi6 >> ((18446744073709551615) & 63)) {
                        case -785:
                            var v15 uint = 2
                            _ = v15
                        }
                        fmt.Println(p0, fi6)
                    }
                    v16 := GenS2{F0: 91, F1: (false || (-2.25 != 1e10)), F2: 32767}
                    _ = fi6
                    _ = v16
                    _ = v7
                }
            } else {
                var v17 uint64 = 4294967295
                _ = v17
            }
            v18 := []int{fi5, fi5}
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            _ = fi5
            _ = v18
        }
    }
    v19 := []int{-583}
    for fi20 := 0; fi20 < 2; fi20++ {
        v21 := GenS1{F0: -662}
        _ = fi20
        _ = v21
    }
    fmt.Println(p0, p1)
    v22 := GenS1{F0: 2147483647}
    _ = p0
    _ = p1
    _ = v19
    _ = v22
    return 0
}

func main() {
    v1 := []int{594, -283}
    v2 := []int{32768}
    var v3 float64 = 0.1
    var v4 float32 = 0.0
    switch "foo bar" {
    case "hello":
        fmt.Println(v3)
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        if (1.5 == 0.1) {
            if true {
                fmt.Println(v3, v3)
                switch "a" {
                case "\n\t":
                    v24 := GenS3{F0: 128}
                    _ = v24
                case "":
                    switch (("a" + "") + ("a" + "foo bar")) {
                    case "hello":
                        v3 += ((float64(0.0) - v3) / (((1.5 - v3) / (v3 / v3)) + 1e10))
                    case "":
                        for fi25 := 0; fi25 < 5; fi25++ {
                            fmt.Println(fi23)
                            _ = fi25
                        }
                    }
                default:
                    v4 = ((float32(((v3 / 1.5) - v3)) * v4) * ((0.1 / v4) - ((((v4 + (float32(1e10) + (float32((float64(float32(0.0)) * 1.5)) - v4))) - 3.14159) + (((v4 - v4) / v4) + v4)) / 0.1)))
                }
                switch (fi23 >> ((644) & 63)) {
                case 9223372036854775807:
                    v26 := GenS1{F0: 42}
                    _ = v26
                default:
                    fmt.Println(v3, fi23)
                }
            }
            v3 -= 1.5
            v4 += 0.0
        }
        fmt.Println(fi23)
        _ = fi23
    }
    fmt.Println(v3, v3)
    _ = len(v1)
    var v27 uint64 = 18446744073709551615
    for fi28 := 0; fi28 < 3; fi28++ {
        if len(v2) > 0 {
            v2[0] = (fi28 - 42)
        }
        _ = fi28
    }
    switch int(v27) {
    case -260:
        var v29 int64 = -176
        _ = v29
    case 32768:
        var v30 uint = 0
        _ = v30
    }
    var v31 bool = (-70 != 100)
    v32 := GenS2{F0: -1, F1: true, F2: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
