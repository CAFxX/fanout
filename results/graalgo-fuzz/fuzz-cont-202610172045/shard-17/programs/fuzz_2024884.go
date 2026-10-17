package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int16 }
func genFn0() int {
    var v4 string = "foo bar"
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 uint32 = 65535
        _ = fi5
        _ = v6
    }
    v4 = (("a" + (v4 + ((v4 + v4) + v4))) + "hello")
    switch 9223372036854775807 {
    case -511:
        switch "" {
        case "hello":
            v4 = v4
        }
    }
    _ = v4
    return 32768
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    var v7 string = "foo bar"
    var v8 float32 = 0.1
    v7 = (((((v7 + (("a" + "") + v7)) + "foo bar") + "a") + v7) + ("foo bar" + "foo bar"))
    v9 := v1[235]
    var v10 string = "a"
    v11 := v1[genFn0()]
    v12 := []int{genFn0(), 1}
    if (false && true) {
        var v13 int16 = 42
        v14 := v2[330]
        _ = v13
        _ = v14
    } else {
        v11 = ("a" + v7)
        var v15 uint = 525
        _ = v15
    }
    if false {
        v16 := GenS1{F0: 732}
        fmt.Println(v9, v10)
        _ = v16
    } else {
        var v17 uint64 = 18446744073709551615
        var v18 uint8 = 1
        _ = v17
        _ = v18
    }
    var v19 float64 = 1.5
    for fi20 := 0; fi20 < 10; fi20++ {
        if ("hello" < "hello") {
            v8 += float32(v19)
            if (!(!(1000 >= 52))) {
                if (!false) {
                    v10 = ("hello" + "")
                } else {
                    fmt.Println(fi20, v7, v11)
                    switch (((("hello" + "a") + "") + ((v9 + (v10 + v11)) + v11)) + "foo bar") {
                    case "\n\t":
                        if len(v12) > 0 {
                            v12[0] = fi20
                        }
                    default:
                        switch v7 {
                        case "a":
                            var v21 bool = true
                            _ = v21
                        default:
                            var v22 uint64 = 42
                            _ = v22
                        }
                    }
                }
            }
            var v23 uint8 = 176
            _ = v23
        } else {
            var v24 int16 = 42
            if (!(true || (!(!false)))) {
                switch (fi20 ^ genFn0()) {
                case -9223372036854775808:
                    v8 += v8
                case 32767:
                    for fi25 := 0; fi25 < 5; fi25++ {
                        if true {
                            if true {
                                var v26 int32 = 0
                                v27 := GenS1{F0: 4294967295}
                                v7 = v7
                                _ = v26
                                _ = v27
                            }
                            v28 := GenS2{F0: -382, F1: 56}
                            var v29 uint16 = 2
                            _ = v28
                            _ = v29
                        }
                        if ((v24 < -76) || (!false)) {
                            if len(v12) > 0 {
                                v12[0] = fi25
                            }
                            fmt.Println(v10, v8)
                            fmt.Println(v19, v24, v19)
                        }
                        v30 := GenS2{F0: 2147483647, F1: 0}
                        _ = fi25
                        _ = v30
                    }
                default:
                    var v31 float64 = float64(v8)
                    _ = v31
                }
            }
            _ = v24
        }
        var v32 uint32 = 4294967295
        _ = fi20
        _ = v32
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        v10 = v11
        _ = fi33
    }
    v10 = "foo bar"
    _ = len(v3)
    var v34 int16 = 42
    v35 := GenS3{F0: (v34 ^ v34)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v7
    _ = v8
    _ = v9
}
