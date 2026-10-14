package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint }
type GenS2 struct { F0 int64; F1 int32; F2 int16 }
type GenS3 struct { F0 bool; F1 uint; F2 bool }
type GenS4 struct { F0 uint16; F1 int8; F2 string }

func main() {
    var v1 float64 = 0.1
    var v2 string = ("hello" + ("hello" + "hello"))
    var v3 string = ("hello" + (v2 + "foo bar"))
    v4 := make(map[string]string)
    var v5 uint8 = 0
    if false {
        v1 -= float64((3.14159 * (((1e10 + -2.25) + 1.5) / float32(((0.1 / (-2.25 / (float64(-2.25) + v1))) / v1)))))
        v1 -= v1
        for fi6 := 0; fi6 < 17; fi6++ {
            if false {
                var v7 uint16 = 402
                v2 = (v3 + (("" + (v2 + v2)) + ""))
                v1 += (((((0.0 - 0.1) * 3.14159) / v1) + ((v1 / 1.5) + v1)) - (v1 + (((1.5 / v1) + 1.5) / v1)))
                _ = v7
            }
            _ = fi6
        }
    } else {
        var v8 string = v3
        if (0 >= -128) {
            if (1 >= 0) {
                var v9 string = v8
                v5 = v5
                _ = v9
            } else {
                v8 = "a"
            }
            var v10 int8 = -81
            v11 := []int{42}
            _ = v10
            _ = v11
        }
        _ = v8
    }
    v1 += (float64(3.14159) / (v1 / 1.0))
    v1 += v1
    for fi12 := 0; fi12 < 12; fi12++ {
        switch fi12 {
        case -1:
            var v13 float32 = 0.0
            _ = v13
        case 4294967295:
            if (((v3 + "hello") + (("a" + v3) + v2)) <= v3) {
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 uint32 = 0
                    if (1e10 >= 3.14159) {
                        if true {
                            v16 := []int{fi12, 32767, (fi14 | -2147483648)}
                            if false {
                                if false {
                                    _ = len(v4)
                                    v5 += (v5 >> ((42) & 63))
                                    fmt.Println(v5, v1)
                                }
                                v17 := GenS1{F0: 4294967295, F1: (fi14 - (fi14 + fi14)), F2: 65535}
                                _ = v17
                            }
                            v18 := GenS1{F0: 2, F1: (fi14 ^ (-(fi12))), F2: 791}
                            _ = v16
                            _ = v18
                        }
                        switch (fi12 % ((fi12) | 1)) {
                        case -21:
                            var v19 uint8 = 220
                            _ = v19
                        }
                        fmt.Println(v1, v15)
                    } else {
                        v15 += v15
                    }
                    _ = fi14
                    _ = v15
                }
                if (!true) {
                    delete(v4, v2)
                    var v20 bool = (!true)
                    var v21 int8 = 72
                    _ = v20
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 15; fi22++ {
                        var v23 uint64 = 18446744073709551615
                        fmt.Println(v23)
                        var v24 uint16 = 880
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    v1 = 0.1
                    var v25 int = (fi12 & 9223372036854775807)
                    _ = v25
                }
            }
        }
        delete(v4, (v3 + (v3 + v2)))
        var v26 uint64 = 42
        _ = fi12
        _ = v26
    }
    v27 := GenS2{F0: 1, F1: 1, F2: 202}
    for fi28 := 0; fi28 < 2; fi28++ {
        switch (v3 + (v2 + "")) {
        case "a":
            var v29 bool = false
            _ = v29
        case "foo bar":
            var v30 float64 = (float64(-2.25) + ((1.5 - 0.0) / 1.5))
            _ = v30
        case "\n\t":
            var v31 int16 = 32767
            _ = v31
        }
        v32 := GenS3{F0: true, F1: 663, F2: false}
        _ = fi28
        _ = v32
    }
    switch ("a" + ("" + v2)) {
    case "a":
        v5 -= (v5 | 0)
    case "foo bar":
        for fi33 := 0; fi33 < 15; fi33++ {
            for fi34 := 0; fi34 < 2; fi34++ {
                fmt.Println(fi34)
                var v35 uint32 = 2
                var v36 bool = true
                _ = fi34
                _ = v35
                _ = v36
            }
            v5 = v5
            switch "\n\t" {
            case "a":
                for fi37 := 0; fi37 < 4; fi37++ {
                    fmt.Println(v1, fi33)
                    _ = fi37
                }
            case "foo bar":
                v1 = float64((-2.25 / 3.14159))
            case "":
                for fi38 := 0; fi38 < 8; fi38++ {
                    var v39 int64 = 568
                    switch fi38 {
                    case 452:
                        v5 = v5
                    case 0:
                        v3 = v3
                    case 953:
                        if true {
                            v40 := GenS4{F0: 1, F1: -108, F2: "foo bar"}
                            _ = v40
                        }
                    default:
                        var v41 uint = 736
                        _ = v41
                    }
                    _ = fi38
                    _ = v39
                }
            }
            _ = fi33
        }
    default:
        v1 -= v1
    }
    switch "a" {
    case "":
        _ = len(v4)
    case "hello":
        switch -9223372036854775808 {
        case 42:
            var v42 uint16 = 1
            _ = v42
        }
    default:
        v1 = (v1 * v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
