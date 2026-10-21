package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int32 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 float64 = -2.25
    var v2 uint8 = 42
    var v3 int = -503
    if (-300 > v3) {
        fmt.Println(v1, v3)
        if true {
            v4 := GenS1{F0: 797, F1: "hello", F2: -2147483648}
            v2 = v2
            _ = v4
        } else {
            v1 -= 3.14159
        }
        for fi5 := 0; fi5 < 13; fi5++ {
            v6 := GenS1{F0: 691, F1: "foo bar", F2: -1}
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 int32 = -2147483648
                _ = fi7
                _ = v8
            }
            var v9 int64 = 356
            _ = fi5
            _ = v6
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            v1 = v1
            for fi11 := 0; fi11 < 5; fi11++ {
                v3 = v3
                switch 32768 {
                case -1:
                    switch "foo bar" {
                    case "a":
                        v1 = (((v1 - (-2.25 + v1)) * 1e10) / -2.25)
                    case "hello":
                        if ((!(!false)) || (false || false)) {
                            v12 := GenS1{F0: 4294967295, F1: ((("hello" + ((((("hello" + ("a" + ((("\n\t" + ("" + ("" + ""))) + "") + "\n\t"))) + ((("hello" + "a") + "\n\t") + ((("foo bar" + "hello") + ("a" + "\n\t")) + "\n\t"))) + "") + ("foo bar" + "foo bar")) + ("a" + "hello"))) + ("a" + "\n\t")) + "foo bar"), F2: 1}
                            _ = v12
                        }
                        v2 = v2
                        var v13 uint16 = 2
                        _ = v13
                    case "\n\t":
                        v2 += v2
                    default:
                        fmt.Println(v2)
                    }
                case 32768:
                    v14 := []int{}
                    _ = v14
                case -2147483648:
                    var v15 float64 = 0.1
                    _ = v15
                default:
                    fmt.Println(v3, fi11)
                }
                v1 = (((v1 / v1) - (1.5 + v1)) - (v1 + 0.0))
                _ = fi11
            }
            _ = fi10
        }
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 uint8 = 1
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := []int{}
            v17 -= (v2 + (v2 * (v17 << ((239) & 63))))
            _ = fi18
            _ = v19
        }
        switch "hello" {
        case "":
            v2 -= v2
        case "a":
            v20 := GenS1{F0: 0, F1: "\n\t", F2: 1}
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    var v21 float32 = -2.25
    v3 -= (v3 >> ((1) & 63))
    v2 += v2
    switch (v3 * v3) {
    case -638:
        v22 := GenS1{F0: 4294967295, F1: "hello", F2: 42}
        _ = v22
    case 32768:
        if false {
            var v23 float64 = (v1 * v1)
            fmt.Println(v23, v2)
            _ = v23
        }
    }
    switch ("a" + ("foo bar" + ("foo bar" + ("hello" + "")))) {
    case "\n\t":
        v24 := GenS2{F0: (105 > 84), F1: v2, F2: 899}
        _ = v24
    case "foo bar":
        for fi25 := 0; fi25 < 8; fi25++ {
            v3 = (-(v3))
            var v26 uint8 = 233
            var v27 uint32 = 463
            _ = fi25
            _ = v26
            _ = v27
        }
    default:
        fmt.Println(v2, v21, v21)
    }
    var v28 string = "\n\t"
    v29 := []int{(v3 | (v3 + (v3 | (v3 | (v3 / (((v3 + (v3 / (((v3 + int(v2))) | 1)))) | 1)))))), v3}
    v30 := GenS3{F0: (v3 / ((v3) | 1))}
    v21 = 3.14159
    v31 := GenS4{F0: 769}
    switch v28 {
    case "hello":
        switch (v28 + (("foo bar" + v28) + v28)) {
        case "":
            for fi32 := 0; fi32 < 19; fi32++ {
                for fi33 := 0; fi33 < 19; fi33++ {
                    v34 := GenS5{F0: (v2 * (^(v2))), F1: 460}
                    _ = fi33
                    _ = v34
                }
                v35 := GenS1{F0: 18446744073709551615, F1: (v28 + v28), F2: 42}
                _ = fi32
                _ = v35
            }
        case "hello":
            v36 := GenS4{F0: 42}
            _ = v36
        }
    case "":
        if len(v29) > 0 {
            v29[0] = v3
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
