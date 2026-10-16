package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 float64 = 1.5
    var v2 float64 = 0.0
    var v3 uint16 = 255
    v4 := make(map[string]string)
    v5 := make(map[uint32]uint32)
    v3 = v3
    var v6 float64 = -2.25
    var v7 float32 = 0.1
    fmt.Println(v1, v7, v6)
    switch "" {
    case "hello":
        switch "a" {
        case "hello":
            var v8 uint16 = (v3 & 0)
            _ = v8
        }
    case "a":
        v9 := v5[2]
        _ = v9
    case "foo bar":
        switch "hello" {
        case "foo bar":
            if (((940 <= 32767) || true) && (uint32(v3) > 0)) {
                delete(v4, ((("foo bar" + (("hello" + "\n\t") + "\n\t")) + (((("hello" + "\n\t") + "\n\t") + "\n\t") + (("\n\t" + "") + (("a" + "hello") + ("a" + "\n\t"))))) + ""))
                if true {
                    var v10 uint64 = 0
                    var v11 int8 = -1
                    _ = v10
                    _ = v11
                } else {
                    v7 += (-2.25 + ((((v7 * (v7 - v7)) - v7) + (-2.25 / 3.14159)) + 0.0))
                    v12 := GenS1{F0: 1, F1: 0, F2: 4294967295}
                    v3 += 1
                    _ = v12
                }
                for fi13 := 0; fi13 < 10; fi13++ {
                    var v14 int64 = 1
                    v6 = -2.25
                    _ = fi13
                    _ = v14
                }
            }
        default:
            for fi15 := 0; fi15 < 18; fi15++ {
                v16 := GenS1{F0: 2, F1: 1, F2: 232}
                if ("\n\t" >= ("a" + "\n\t")) {
                    var v17 string = "hello"
                    switch (fi15 % (((^(fi15))) | 1)) {
                    case 32767:
                        if (true && false) {
                            if (18446744073709551615 <= 1) {
                                _ = len(v5)
                                switch (v17 + v17) {
                                case "":
                                    v18 := GenS1{F0: 0, F1: 255, F2: uint(v3)}
                                    _ = v18
                                case "a":
                                    fmt.Println(v2, v2)
                                case "hello":
                                    for fi19 := 0; fi19 < 5; fi19++ {
                                        var v20 uint32 = 0
                                        v21 := GenS1{F0: 0, F1: 2, F2: 18446744073709551615}
                                        _ = fi19
                                        _ = v20
                                        _ = v21
                                    }
                                default:
                                    switch (v17 + v17) {
                                    case "hello":
                                        v17 = "hello"
                                    case "":
                                        v2 = (v6 - v2)
                                    case "foo bar":
                                        v22 := GenS2{F0: "hello", F1: v3}
                                        _ = v22
                                    }
                                }
                                v23 := GenS3{F0: true, F1: 42}
                                _ = v23
                            }
                            switch fi15 {
                            case 803:
                                v3 += v3
                            case 4294967295:
                                if false {
                                    switch (fi15 << ((913) & 63)) {
                                    case 829:
                                        if (-2.25 > v7) {
                                            v24 := []int{fi15, (fi15 >> ((42) & 63)), fi15}
                                            _ = len(v4)
                                            var v25 uint8 = 1
                                            _ = v24
                                            _ = v25
                                        }
                                    default:
                                        for fi26 := 0; fi26 < 14; fi26++ {
                                            var v27 uint = 4294967295
                                            v28 := GenS4{F0: -2147483648}
                                            _ = fi26
                                            _ = v27
                                            _ = v28
                                        }
                                    }
                                    fmt.Println(v1)
                                } else {
                                    var v29 uint8 = 255
                                    _ = v29
                                }
                            }
                            var v30 uint16 = v3
                            _ = v30
                        }
                    case 128:
                        fmt.Println(fi15, v1, v17)
                    }
                    var v31 int32 = -1
                    _ = v17
                    _ = v31
                }
                _ = fi15
                _ = v16
            }
        }
    }
    switch 32767 {
    case 711:
        if (!true) {
            v6 = v6
        } else {
            v6 = 0.0
        }
    case 1:
        v3 = v3
    }
    v32 := GenS1{F0: 239, F1: 0, F2: 656}
    var v33 int = 128
    var v34 uint32 = 295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
