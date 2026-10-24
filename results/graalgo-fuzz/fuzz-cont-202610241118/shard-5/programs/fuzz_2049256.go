package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 bool; F1 int64; F2 int16 }
func genFn0(p0 int64) uint64 {
    fmt.Println(p0)
    var v4 float32 = 3.14159
    v5 := []int{9223372036854775807, 0}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 int16, p1 uint) uint {
    v6 := GenS1{F0: 0, F1: (-69 != -1)}
    p1 = (^(p1))
    if (2147483647 < 2147483648) {
        v7 := []int{1, 32768, 986, 0}
        var v8 string = ("" + ("\n\t" + "a"))
        for fi9 := 0; fi9 < 11; fi9++ {
            if true {
                switch 32767 {
                case 967:
                    v7 = append(v7, fi9)
                default:
                    if (2 < 42) {
                        v10 := GenS1{F0: fi9, F1: false}
                        _ = v10
                    }
                }
                if ((2 > 255) || true) {
                    if true {
                        fmt.Println(fi9)
                    } else {
                        if (true && ((p1 == 255) && ((^(p0)) >= p0))) {
                            v11 := GenS1{F0: 127, F1: true}
                            _ = v11
                        }
                    }
                    _ = len(v7)
                    for fi12 := 0; fi12 < 10; fi12++ {
                        if false {
                            if (!false) {
                                var v13 uint64 = genFn0(-9223372036854775808)
                                _ = v13
                            }
                            switch v8 {
                            case "":
                                fmt.Println(p0, p1, v8)
                                fmt.Println(p1, p0)
                                var v14 uint32 = 65535
                                _ = v14
                            }
                            p1 = p1
                            if (-32768 < (p0 << ((p1) & 63))) {
                                fmt.Println(fi12, v8)
                                if len(v7) > 0 {
                                    v7[0] = 9223372036854775807
                                }
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                            v15 := GenS2{F0: 27, F1: -1}
                            _ = v15
                        }
                        _ = fi12
                    }
                }
                p1 += 65535
            }
            var v16 string = "foo bar"
            for fi17 := 0; fi17 < 4; fi17++ {
                fmt.Println(p1)
                p1 += p1
                _ = fi17
            }
            _ = fi9
            _ = v16
        }
        _ = v7
        _ = v8
    }
    v18 := []int{-395}
    var v19 float32 = -2.25
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v6
    return (p1 & 65535)
}

func main() {
    var v1 int64 = 920
    var v2 string = "a"
    v3 := make(map[string]string)
    switch "" {
    case "\n\t":
        v2 = "foo bar"
    }
    if false {
        v1 = 405
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS1{F0: fi20, F1: true}
        var v22 uint8 = 42
        _ = fi20
        _ = v21
        _ = v22
    }
    if true {
        var v23 uint32 = 478
        _ = v23
    } else {
        v2 = v2
        v2 = (v2 + v2)
        var v24 uint8 = 74
        _ = v24
    }
    v25 := GenS3{F0: true, F1: -9223372036854775808, F2: -32768}
    switch (v2 + ("hello" + v2)) {
    case "hello":
        v26 := []int{162}
        _ = v26
    }
    fmt.Println(v2, v2)
    var v27 bool = false
    if true {
        var v28 int = 127
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 int64 = v1
            for fi31 := 0; fi31 < 4; fi31++ {
                v1 = 0
                switch v2 {
                case "a":
                    var v32 int32 = 636
                    _ = v32
                case "\n\t":
                    v33 := GenS2{F0: 255, F1: v1}
                    _ = v33
                }
                _ = fi31
            }
            var v34 float32 = 0.1
            _ = fi29
            _ = v30
            _ = v34
        }
        _ = v28
    }
    var v35 float32 = ((0.0 * 1.5) * 1e10)
    v1 = 9223372036854775807
    var v36 string = v2
    delete(v3, "foo bar")
    var v37 uint32 = 256
    fmt.Println(v35, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v37
}
