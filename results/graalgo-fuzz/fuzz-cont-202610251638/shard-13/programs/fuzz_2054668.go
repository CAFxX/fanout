package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int16 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int64 = 42
    var v2 string = (((("hello" + (("foo bar" + "\n\t") + "\n\t")) + "hello") + ((("a" + "") + "a") + "\n\t")) + ("foo bar" + "foo bar"))
    v3 := []int{32767, 0}
    var v4 int32 = -5
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v3 = append(v3, 42)
        var v5 uint16 = 65535
        _ = v5
    }
    v6 := GenS1{F0: 42, F1: false, F2: 818}
    switch v2 {
    case "hello":
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := GenS2{F0: 834, F1: -1, F2: 65535}
            _ = fi7
            _ = v8
        }
    case "":
        switch 9223372036854775807 {
        case 498:
            for fi9 := 0; fi9 < 17; fi9++ {
                fmt.Println(v4)
                for fi10 := 0; fi10 < 13; fi10++ {
                    fmt.Println(v2, v1, v2)
                    v2 = ((v2 + v2) + ("a" + v2))
                    v11 := GenS3{F0: (v1 * v1)}
                    v12 := GenS2{F0: 42, F1: (v1 >> ((255) & 63)), F2: 255}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = fi9
            }
        case 42:
            if true {
                var v13 uint32 = 904
                var v14 uint = 18446744073709551615
                switch (("" + ("foo bar" + v2)) + "foo bar") {
                case "foo bar":
                    for fi15 := 0; fi15 < 13; fi15++ {
                        var v16 float64 = (0.0 - -2.25)
                        v2 = "\n\t"
                        v17 := GenS3{F0: 1}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                default:
                    v14 = v14
                }
                _ = v13
                _ = v14
            } else {
                if ((904 < -580) && true) {
                    v18 := GenS4{F0: 0}
                    _ = v18
                }
            }
        default:
            v1 = 1
        }
    case "foo bar":
        var v19 int16 = 0
        _ = v19
    }
    var v20 uint64 = 564
    switch -9223372036854775808 {
    case 310:
        v1 -= (v1 * v1)
    default:
        v2 = v2
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        v20 += (^(v20))
        if true {
            for fi22 := 0; fi22 < 3; fi22++ {
                v23 := GenS2{F0: 2, F1: v1, F2: 18446744073709551615}
                fmt.Println(v4)
                _ = fi22
                _ = v23
            }
        }
        _ = fi21
    }
    switch 103 {
    case -616:
        var v24 uint64 = 2
        _ = v24
    default:
        var v25 int8 = 0
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
}
