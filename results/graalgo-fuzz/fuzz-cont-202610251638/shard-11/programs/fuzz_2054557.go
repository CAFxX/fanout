package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 string = ""
    var v2 uint16 = 42
    v3 := make(map[string]int)
    v4 := []string{}
    v5 := make(map[string]string)
    v6 := make(map[int]string)
    if false {
        v2 -= 65535
    }
    var v7 float64 = 0.1
    if (-1 > -655) {
        var v8 uint = 2
        v9 := GenS1{F0: 282}
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        _ = v8
        _ = v9
    }
    if true {
        if (255 == 128) {
            v7 = float64(float32(v7))
            v1 = "\n\t"
        }
    }
    v10 := GenS1{F0: 42}
    if (v7 > (v7 / v7)) {
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := GenS1{F0: 66}
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: 546}
        delete(v3, v1)
        _ = v13
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            delete(v6, 32768)
            if (!(1 != 143)) {
                switch "foo bar" {
                case "hello":
                    if (0 > 9223372036854775807) {
                        v2 = v2
                        fmt.Println(fi14)
                    } else {
                        if (0.0 <= 1.5) {
                            v1 = v1
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                        } else {
                            fmt.Println(v2)
                            var v15 uint8 = 0
                            if true {
                                v16 := GenS2{F0: -1}
                                fmt.Println(v1, v7, v7)
                                _ = v16
                            } else {
                                for fi17 := 0; fi17 < 15; fi17++ {
                                    fmt.Println(v2)
                                    var v18 float64 = -2.25
                                    v7 -= float64(0.0)
                                    _ = fi17
                                    _ = v18
                                }
                                v2 -= (v2 * v2)
                            }
                            _ = v15
                        }
                    }
                }
                var v19 bool = true
                var v20 int16 = 1
                _ = v19
                _ = v20
            } else {
                var v21 int16 = 1
                fmt.Println(v7, v21, v21)
                _ = v21
            }
            v22 := GenS3{F0: ("hello" + v1), F1: 0}
            _ = fi14
            _ = v22
        }
    }
    v7 = 3.14159
    v1 = v1
    if ((490 >= 204) && false) {
        v7 -= v7
        _ = len(v3)
    } else {
        if (4294967295 == 1) {
            v2 = v2
            _ = len(v4)
        }
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 float64 = (float64(3.14159) * (float64((1.5 / 0.1)) * float64(0.0)))
        var v25 string = v1
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 bool = false
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = v1
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        for fi28 := 0; fi28 < 8; fi28++ {
            v29 := GenS4{F0: 335}
            v4 = append(v4, ("foo bar" + v1))
            fmt.Println(v7)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (!true) {
                delete(v6, (fi27 % (((fi28 * 2147483647)) | 1)))
                v2 += 255
            }
            switch (fi27 * (fi28 - 0)) {
            case 32768:
                var v30 int16 = -271
                _ = v30
            default:
                v7 -= (-2.25 / v7)
            }
            _ = fi28
            _ = v29
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v2)
        _ = fi27
    }
    switch "\n\t" {
    case "\n\t":
        var v31 bool = (42 >= 2147483647)
        _ = v31
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        v7 = (v7 / ((1.5 * (v7 * (((float64(float32(v7)) - -2.25) / 1e10) / 0.1))) + v7))
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
