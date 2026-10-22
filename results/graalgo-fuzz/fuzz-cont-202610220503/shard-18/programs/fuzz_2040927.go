package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 32767
    v1 -= (v1 & v1)
    v3 -= (^(v3))
    if (96 > 0) {
        switch 672 {
        case 495:
            switch "a" {
            case "hello":
                v4 := GenS1{F0: -435, F1: 42}
                _ = v4
            case "a":
                v5 := GenS1{F0: -1, F1: -847}
                _ = v5
            }
        case 127:
            switch "foo bar" {
            case "\n\t":
                switch ("\n\t" + (("foo bar" + "a") + (("a" + "foo bar") + ("a" + ("hello" + ((("hello" + "a") + "foo bar") + "foo bar")))))) {
                case "":
                    for fi6 := 0; fi6 < 4; fi6++ {
                        v2[uint32(v3)] = 965
                        v7 := GenS1{F0: 0, F1: (fi6 << ((42) & 63))}
                        v3 = 32767
                        _ = fi6
                        _ = v7
                    }
                case "\n\t":
                    v8 := GenS1{F0: 0, F1: 175}
                    _ = v8
                }
            }
        case -120:
            if (!(!(-32768 > 661))) {
                for fi9 := 0; fi9 < 17; fi9++ {
                    for fi10 := 0; fi10 < 5; fi10++ {
                        if true {
                            v11 := GenS1{F0: 1, F1: fi9}
                            v1 += uint64(v3)
                            _ = v11
                        }
                        var v12 string = "a"
                        if false {
                            switch fi9 {
                            case 42:
                                v1 += v1
                            }
                            var v13 int = (fi10 * fi10)
                            _ = v13
                        }
                        _ = fi10
                        _ = v12
                    }
                    _ = fi9
                }
            }
        default:
            for fi14 := 0; fi14 < 20; fi14++ {
                for fi15 := 0; fi15 < 10; fi15++ {
                    if (4294967295 == 42) {
                        v2[42] = 4294967295
                    }
                    _ = fi15
                }
                v1 = 465
                _ = fi14
            }
        }
        var v16 uint8 = 255
        var v17 uint8 = (^(v16))
        _ = v16
        _ = v17
    } else {
        var v18 uint16 = v3
        var v19 int32 = 2147483647
        _ = v18
        _ = v19
    }
    v1 = (v1 | v1)
    v3 = v3
    var v20 uint64 = 319
    var v21 uint16 = (v3 - v3)
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 float32 = float32(float64(0.0))
        if false {
            var v24 int32 = -731
            _ = v24
        } else {
            var v25 float64 = 3.14159
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    switch 621 {
    case 443:
        v26 := GenS1{F0: -1, F1: -395}
        _ = v26
    case 42:
        for fi27 := 0; fi27 < 6; fi27++ {
            v21 = v21
            _ = fi27
        }
    default:
        if false {
            var v28 string = (("foo bar" + "a") + "hello")
            var v29 uint64 = (v1 * (v1 >> ((341) & 63)))
            fmt.Println(v3, v21, v1)
            _ = v28
            _ = v29
        }
    }
    var v30 int32 = 0
    for fi31 := 0; fi31 < 3; fi31++ {
        v32 := []int{}
        _ = fi31
        _ = v32
    }
    var v33 int32 = -566
    v30 = (v33 - v33)
    var v34 int = -875
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v33
    _ = v34
}
