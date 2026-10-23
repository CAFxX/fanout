package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 string) uint {
    v5 := []int{9223372036854775807}
    p0 = p0
    p0 = p0
    switch "a" {
    case "foo bar":
        if (1 <= 2) {
            for fi6 := 0; fi6 < 10; fi6++ {
                v7 := GenS1{F0: 0, F1: 2, F2: 18446744073709551615}
                for fi8 := 0; fi8 < 20; fi8++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi8
                }
                var v9 int8 = -128
                _ = fi6
                _ = v7
                _ = v9
            }
            switch (p0 + p0) {
            case "hello":
                for fi10 := 0; fi10 < 15; fi10++ {
                    p0 = ((p0 + (((p0 + "\n\t") + p0) + ((("" + p0) + (p0 + "a")) + p0))) + (p0 + ""))
                    if false {
                        switch (fi10 - (fi10 - fi10)) {
                        case 32768:
                            var v11 uint16 = 42
                            _ = v11
                        }
                        p0 = p0
                        var v12 int64 = 42
                        _ = v12
                    } else {
                        var v13 string = (p0 + "a")
                        for fi14 := 0; fi14 < 16; fi14++ {
                            fmt.Println(fi10)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            p0 = ((p0 + (((("hello" + v13) + p0) + "a") + "foo bar")) + v13)
                            _ = fi14
                        }
                        _ = v13
                    }
                    fmt.Println(fi10, p0)
                    _ = fi10
                }
            }
        }
    case "a":
        p0 = "hello"
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        p0 = p0
    }
    _ = len(v5)
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 bool, p1 int, p2 uint32) int {
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS1{F0: 65535, F1: 65535, F2: 822}
        v17 := []int{}
        var v18 float32 = ((3.14159 - (float32(0.1) * (1e10 / (1e10 / 1.5)))) + float32((((-2.25 + 0.1) * (1e10 + 0.0)) / (-2.25 + 1e10))))
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    fmt.Println(p2, p2)
    var v19 uint8 = 252
    v20 := GenS1{F0: 1, F1: 255, F2: 852}
    v21 := GenS1{F0: 65535, F1: 255, F2: uint64(p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    _ = v21
    return -83
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 42
    v3 := make(map[uint32]bool)
    var v4 bool = false
    v22 := GenS2{F0: 255}
    v4 = (-2147483648 <= -988)
    if v4 {
        var v23 uint8 = 127
        var v24 int = -9223372036854775808
        v2 += v2
        _ = v23
        _ = v24
    }
    v25 := []int{-9223372036854775808}
    var v26 uint16 = 65535
    for fi27 := 0; fi27 < 13; fi27++ {
        var v28 int64 = 830
        switch -9223372036854775808 {
        case 0:
            if (-128 > -1) {
                for fi29 := 0; fi29 < 9; fi29++ {
                    v30 := GenS3{F0: 32767}
                    _ = fi29
                    _ = v30
                }
                v26 = 308
            }
        case -551:
            if ((0.0 + 0.1) == ((0.1 / 3.14159) - 0.0)) {
                delete(v3, 2)
            }
        default:
            if (!true) {
                if true {
                    v28 += v28
                    if true {
                        var v31 uint8 = 207
                        var v32 int8 = -128
                        var v33 string = ("foo bar" + "hello")
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                } else {
                    if v4 {
                        _ = len(v25)
                        for fi34 := 0; fi34 < 6; fi34++ {
                            var v35 float32 = (float32(float64(3.14159)) + float32(-2.25))
                            fmt.Println(v26, v35)
                            _ = fi34
                            _ = v35
                        }
                        fmt.Println(v26)
                    }
                    v28 = v28
                    v26 = 96
                }
            }
        }
        v28 += 1
        _ = fi27
        _ = v28
    }
    var v36 uint16 = 2
    delete(v3, 65535)
    var v37 uint8 = 14
    for fi38 := 0; fi38 < 17; fi38++ {
        var v39 int16 = 126
        _ = fi38
        _ = v39
    }
    var v40 uint64 = 162
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v40
}
