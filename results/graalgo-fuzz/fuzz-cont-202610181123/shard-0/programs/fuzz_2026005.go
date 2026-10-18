package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint8; F1 string; F2 string }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int32; F1 bool; F2 uint8 }
func genFn0(p0 int, p1 int16) uint {
    fmt.Println(p0, p0)
    p1 += (p1 ^ 847)
    v5 := []int{p0, p0, -87}
    if len(v5) > 0 {
        v5[0] = (^(p0))
    }
    if true {
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 9; fi7++ {
                p0 += fi7
                var v8 uint64 = 0
                if false {
                    p0 += (fi6 / ((911) | 1))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(p1)
                } else {
                    fmt.Println(v8, fi7)
                }
                switch "foo bar" {
                case "a":
                    var v9 float64 = 3.14159
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        if false {
            v10 := GenS1{F0: -1, F1: 232}
            if (3 >= 255) {
                v11 := GenS2{F0: 127, F1: "foo bar", F2: ("hello" + "a")}
                v12 := GenS3{F0: 588, F1: p1}
                v5 = append(v5, p0)
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v14 := GenS4{F0: -120, F1: ((258 == 2) || (!false))}
                    v15 := GenS4{F0: fi13, F1: true}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v16 := GenS1{F0: 2147483647, F1: 1}
                p1 = p1
                _ = v16
            }
            for fi17 := 0; fi17 < 14; fi17++ {
                switch "a" {
                case "\n\t":
                    v18 := GenS5{F0: (p0 << ((255) & 63))}
                    _ = v18
                case "":
                    for fi19 := 0; fi19 < 12; fi19++ {
                        v20 := GenS6{F0: 18446744073709551615, F1: 65535}
                        v21 := GenS6{F0: 898, F1: 568}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                case "foo bar":
                    var v22 int8 = -1
                    _ = v22
                default:
                    p1 = p1
                }
                var v23 int = int(p1)
                switch (v23 + (fi17 - (p0 * 4294967295))) {
                case -1:
                    var v24 uint32 = 786
                    _ = v24
                case -2147483648:
                    var v25 int = v23
                    _ = v25
                case -9223372036854775808:
                    var v26 bool = ((!true) && (!(3.14159 >= 3.14159)))
                    _ = v26
                }
                _ = fi17
                _ = v23
            }
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 716
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]string)
    var v3 int32 = 1
    v4 := make(map[string]string)
    v1 -= (v1 ^ v1)
    for fi27 := 0; fi27 < 12; fi27++ {
        if (42 > 506) {
            var v28 string = ""
            _ = v28
        }
        _ = fi27
    }
    fmt.Println(v1)
    if (genFn0(2147483647, 42) > 4294967295) {
        v1 += v1
    } else {
        switch 127 {
        case -506:
            _ = len(v4)
        case 42:
            fmt.Println(v3, v1)
        default:
            if true {
                v29 := GenS3{F0: genFn0(-173, 0), F1: -1}
                _ = v29
            }
        }
    }
    switch -2147483648 {
    case 474:
        switch ((((("hello" + "") + "hello") + "foo bar") + "a") + "foo bar") {
        case "a":
            for fi30 := 0; fi30 < 18; fi30++ {
                for fi31 := 0; fi31 < 11; fi31++ {
                    v32 := GenS7{F0: v3, F1: true, F2: v1}
                    switch "" {
                    case "foo bar":
                        v1 = (v1 * v1)
                    }
                    _ = fi31
                    _ = v32
                }
                if true {
                    v33 := GenS7{F0: v3, F1: false, F2: v1}
                    v3 -= v3
                    _ = v33
                }
                _ = fi30
            }
        }
    case 42:
        var v34 int = 128
        _ = v34
    }
    var v35 uint64 = 0
    var v36 uint = 275
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
