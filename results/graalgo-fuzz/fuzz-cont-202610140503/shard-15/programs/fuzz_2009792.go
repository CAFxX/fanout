package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int16; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint64; F1 uint16; F2 bool }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int = -2147483648
    var v3 uint32 = 2147483648
    if false {
        if false {
            var v4 string = "hello"
            var v5 uint = 42
            v2 = v2
            _ = v4
            _ = v5
        }
    }
    v6 := GenS1{F0: ("\n\t" + "hello")}
    v7 := GenS2{F0: "foo bar", F1: -1}
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS3{F0: 32767, F1: (v3 >> ((18446744073709551615) & 63)), F2: 42}
        var v10 int8 = -84
        var v11 float64 = (0.0 / 1.5)
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        v2 = v2
        fmt.Println(v1, v1)
        _ = fi12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        if (!false) {
            for fi14 := 0; fi14 < 10; fi14++ {
                switch 9223372036854775807 {
                case 1:
                    if true {
                        var v15 uint64 = v1
                        v16 := GenS4{F0: v1, F1: -32768}
                        _ = v15
                        _ = v16
                    } else {
                        v2 -= (v2 / ((312) | 1))
                    }
                    fmt.Println(v3, fi13, fi13)
                    switch fi14 {
                    case 52:
                        fmt.Println(fi14)
                    case 0:
                        var v17 uint32 = (v3 << ((765) & 63))
                        _ = v17
                    case -325:
                        var v18 int32 = -2147483648
                        v2 = (^(fi13))
                        _ = v18
                    }
                    v19 := GenS2{F0: (("" + ("foo bar" + "")) + ("a" + "\n\t")), F1: -1}
                    _ = v19
                case 2147483647:
                    if false {
                        switch "\n\t" {
                        case "\n\t":
                            v1 = 0
                            var v20 uint32 = v3
                            var v21 int8 = 123
                            switch (v2 / (((v2 * int(v3))) | 1)) {
                            case -9223372036854775808:
                                var v22 uint64 = v1
                                _ = v22
                            case 128:
                                v23 := GenS2{F0: "\n\t", F1: 42}
                                _ = v23
                            }
                            var v24 uint64 = 817
                            _ = v20
                            _ = v21
                            _ = v24
                        default:
                            v25 := GenS5{F0: v1, F1: 65535, F2: (true && true)}
                            _ = v25
                        }
                        v26 := GenS3{F0: -620, F1: v3, F2: -127}
                        _ = v26
                    } else {
                        var v27 uint64 = 4294967295
                        _ = v27
                    }
                default:
                    v28 := GenS1{F0: "\n\t"}
                    _ = v28
                }
                fmt.Println(fi14)
                v1 -= (v1 >> ((45) & 63))
                _ = fi14
            }
        }
        _ = fi13
    }
    var v29 bool = (!false)
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 uint32 = v3
        switch v2 {
        case 173:
            var v32 uint8 = 42
            _ = v32
        case 2147483647:
            switch ("a" + "a") {
            case "a":
                var v33 string = ""
                _ = v33
            default:
                var v34 string = "foo bar"
                _ = v34
            }
        case 396:
            v1 -= v1
        }
        _ = fi30
        _ = v31
    }
    v35 := GenS1{F0: ""}
    if (!v29) {
        for fi36 := 0; fi36 < 14; fi36++ {
            var v37 float32 = ((1e10 + 0.0) * (1.5 + (0.0 * 3.14159)))
            _ = fi36
            _ = v37
        }
    } else {
        if true {
            var v38 int16 = -923
            if v29 {
                v39 := GenS1{F0: (("foo bar" + "hello") + (("hello" + ("a" + "a")) + "foo bar"))}
                _ = v39
            }
            _ = v38
        } else {
            v29 = ((42 <= 255) && ((v1 << ((255) & 63)) >= 42))
            fmt.Println(v3, v2)
        }
        for fi40 := 0; fi40 < 5; fi40++ {
            v3 -= (v3 << ((796) & 63))
            _ = fi40
        }
        for fi41 := 0; fi41 < 4; fi41++ {
            switch (("" + ("a" + ("foo bar" + ("\n\t" + "foo bar")))) + "hello") {
            case "":
                v2 += v2
            }
            _ = fi41
        }
    }
    fmt.Println(v2, v2, v3)
    v42 := GenS1{F0: ("a" + "hello")}
    v43 := GenS3{F0: -1, F1: 42, F2: -128}
    v44 := []int{(v2 - 32768)}
    if len(v44) > 0 {
        v44[0] = v2
    }
    fmt.Println(v1, v29, v1)
    v2 -= 0
    var v45 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v6
    _ = v7
}
