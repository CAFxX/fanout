package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint }
type GenS3 struct { F0 uint8 }
func genFn0() uint {
    switch -333 {
    case 0:
        v5 := GenS1{F0: "a", F1: (257 <= -1)}
        _ = v5
    }
    var v6 uint32 = 2
    v6 -= v6
    _ = v6
    return uint(v6)
}

func main() {
    var v1 float32 = 1.5
    v2 := []string{}
    var v3 int = 42
    var v4 uint32 = 271
    v4 = (v4 % (((v4 / ((v4) | 1))) | 1))
    var v7 int32 = 42
    var v8 int16 = 32767
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            switch "" {
            case "":
                v3 += v3
            }
            var v11 bool = (128 <= 88)
            _ = fi10
            _ = v11
        }
        v12 := GenS2{F0: 112, F1: 128, F2: 887}
        fmt.Println(v1, v7)
        _ = fi9
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        _ = len(v2)
        v14 := GenS1{F0: "a", F1: true}
        if (((v7 + v7) != (-(v7))) && false) {
            _ = len(v2)
        } else {
            if false {
                for fi15 := 0; fi15 < 14; fi15++ {
                    v7 = (v7 | (v7 - 42))
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    var v16 uint16 = 2
                    _ = fi15
                    _ = v16
                }
                var v17 float32 = v1
                fmt.Println(v8, v3, v17)
                _ = v17
            } else {
                _ = len(v2)
                for fi18 := 0; fi18 < 7; fi18++ {
                    switch 128 {
                    case 42:
                        var v19 uint8 = 2
                        _ = v19
                    }
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    _ = fi18
                }
                var v20 uint16 = 1
                _ = v20
            }
            var v21 uint = 18446744073709551615
            v22 := GenS1{F0: ((("a" + "\n\t") + ("hello" + "\n\t")) + ("a" + (((("a" + "foo bar") + "hello") + (("a" + "") + "a")) + "foo bar"))), F1: true}
            _ = v21
            _ = v22
        }
        _ = fi13
        _ = v14
    }
    v1 += (v1 * ((v1 / 1.0) / (0.0 * v1)))
    v23 := GenS1{F0: ("a" + ""), F1: false}
    if (v8 > 339) {
        v4 -= (v4 + v4)
        fmt.Println(v4, v7)
        v7 = v7
    } else {
        v8 = (^(v8))
        v7 -= v7
        v24 := GenS2{F0: 212, F1: 1, F2: 18446744073709551615}
        _ = v24
    }
    v3 += (v3 >> ((546) & 63))
    switch (("foo bar" + "hello") + "a") {
    case "":
        if true {
            v25 := GenS3{F0: 127}
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 bool = true
                v8 += 1
                v2 = append(v2, (("" + "foo bar") + "foo bar"))
                _ = fi26
                _ = v27
            }
            switch "\n\t" {
            case "a":
                v4 = (v4 >> ((2) & 63))
            case "\n\t":
                if false {
                    v2 = append(v2, "")
                    var v28 uint = 0
                    _ = v28
                } else {
                    v29 := GenS2{F0: 1, F1: 93, F2: 2}
                    _ = len(v2)
                    v8 -= 1
                    _ = v29
                }
            case "hello":
                switch (^(v3)) {
                case 537:
                    var v30 int64 = 1
                    _ = v30
                case 128:
                    v2 = append(v2, "foo bar")
                default:
                    v7 -= -2147483648
                }
            default:
                v31 := GenS1{F0: "", F1: true}
                _ = v31
            }
            _ = v25
        } else {
            v2 = append(v2, ((("a" + ("a" + ("foo bar" + "foo bar"))) + "\n\t") + "foo bar"))
            var v32 uint32 = 349
            switch v3 {
            case 0:
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            }
            _ = v32
        }
    case "foo bar":
        v1 += float32(1e10)
    case "a":
        if false {
            v8 -= v8
            if false {
                switch v3 {
                case 2147483647:
                    var v33 float64 = (1e10 - 0.0)
                    _ = v33
                case 9223372036854775807:
                    var v34 float64 = ((float64(v1) + ((float64(v1) - -2.25) - 1e10)) + (-2.25 - float64((-2.25 / (v1 / 1e10)))))
                    _ = v34
                }
                v2 = append(v2, "\n\t")
            }
        } else {
            var v35 string = "a"
            var v36 uint32 = (v4 / ((42) | 1))
            _ = v35
            _ = v36
        }
    }
    var v37 uint8 = 62
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v37
    _ = v4
    _ = v7
    _ = v8
}
