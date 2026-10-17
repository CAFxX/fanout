package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int64 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int; F1 uint8; F2 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 string; F2 uint8 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint32) string {
    var v7 bool = (2 > 255)
    fmt.Println(v7, p1, v7)
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int16 = -32768
        _ = fi8
        _ = v9
    }
    switch ((("" + "hello") + "hello") + "\n\t") {
    case "":
        var v10 uint = 42
        _ = v10
    case "hello":
        fmt.Println(p1)
    case "a":
        v11 := GenS1{F0: 86, F1: 103, F2: -74}
        _ = v11
    }
    var v12 int64 = 42
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    return ("foo bar" + "foo bar")
}
func genFn1(p0 uint64, p1 uint16, p2 int16) string {
    if (false && (!true)) {
        v13 := []int{-9223372036854775808, -253, 2147483647, -931}
        _ = v13
    }
    var v14 uint64 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (("\n\t" + "foo bar") + (("" + "hello") + ("" + "a")))
}

func main() {
    v1 := []int{-583, 127}
    var v2 uint8 = 170
    var v3 int = 0
    var v4 string = ""
    var v5 string = v4
    var v6 int32 = -2147483648
    for fi15 := 0; fi15 < 8; fi15++ {
        _ = len(v1)
        _ = fi15
    }
    var v16 bool = (1 <= 65535)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v17 := GenS2{F0: 32767, F1: 65535}
    v5 = "foo bar"
    v4 = ("\n\t" + v5)
    v16 = (v16 || (!v16))
    switch 2147483647 {
    case 0:
        var v18 int32 = 660
        _ = v18
    }
    switch (v3 / (((-(v3))) | 1)) {
    case 369:
        switch ((v5 + v4) + v5) {
        case "\n\t":
            var v19 uint16 = 255
            _ = v19
        case "":
            _ = len(v1)
        }
    case 517:
        for fi20 := 0; fi20 < 3; fi20++ {
            if v16 {
                for fi21 := 0; fi21 < 12; fi21++ {
                    if v16 {
                        switch (v3 + fi21) {
                        case 347:
                            v2 += v2
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (!(fi21 == fi20)) {
                            if true {
                                switch (v3 >> ((42) & 63)) {
                                case 1:
                                    v6 += 0
                                case 42:
                                    fmt.Println(v6, v5)
                                }
                                v22 := GenS1{F0: 9223372036854775807, F1: 836, F2: int64(v6)}
                                _ = v22
                            } else {
                                v23 := GenS3{F0: fi21, F1: v2, F2: 4294967295}
                                if len(v1) > 0 {
                                    v1[0] = fi21
                                }
                                fmt.Println(v3, fi20, v6)
                                v1 = append(v1, (fi20 | 1))
                                v1 = append(v1, (fi21 & fi21))
                                fmt.Println(fi20)
                                v24 := GenS3{F0: v3, F1: 128, F2: 255}
                                _ = v23
                                _ = v24
                            }
                            var v25 int64 = 1
                            switch (v5 + "hello") {
                            case "a":
                                v26 := GenS2{F0: 0, F1: 1}
                                _ = v26
                            }
                            _ = v25
                        } else {
                            var v27 int16 = -8
                            _ = v27
                        }
                    } else {
                        v1 = append(v1, -693)
                        v1 = append(v1, 0)
                        v5 = "\n\t"
                    }
                    _ = fi21
                }
            } else {
                if (v5 > "a") {
                    switch fi20 {
                    case -674:
                        for fi28 := 0; fi28 < 18; fi28++ {
                            if ((v6 - 1) != (v6 % ((1) | 1))) {
                                var v29 float64 = -2.25
                                switch (fi28 * 127) {
                                case 9223372036854775807:
                                    v3 += -278
                                case -832:
                                    var v30 int = (v3 >> ((42) & 63))
                                    _ = v30
                                default:
                                    var v31 int16 = 0
                                    _ = v31
                                }
                                _ = v29
                            } else {
                                var v32 float32 = 0.1
                                _ = v32
                            }
                            _ = fi28
                        }
                    }
                }
                var v33 string = (v5 + v4)
                _ = v33
            }
            switch ((v5 + "\n\t") + "") {
            case "":
                if v16 {
                    if v16 {
                        fmt.Println(v5, v16)
                        fmt.Println(v6)
                        v1 = append(v1, (v3 % (((v3 / (((fi20 + (fi20 / ((fi20) | 1)))) | 1))) | 1)))
                    }
                    if true {
                        var v34 bool = ((((0.1 + (float64(1.5) + 3.14159)) == 0.0) && true) && v16)
                        _ = v34
                    } else {
                        var v35 bool = v16
                        _ = v35
                    }
                    v36 := GenS4{F0: (v6 - (v6 + v6))}
                    _ = v36
                }
            case "hello":
                v37 := GenS5{F0: v6, F1: (v5 + (v5 + v4)), F2: 157}
                _ = v37
            case "foo bar":
                v1 = append(v1, (^(fi20)))
            default:
                v5 = "a"
            }
            var v38 float64 = 0.1
            _ = fi20
            _ = v38
        }
    }
    if ((2 >= 65535) || ((v3 < (v3 << ((17) & 63))) || (42 <= 317))) {
        if v16 {
            v16 = ((true && (!(v16 || (v2 <= 0)))) || v16)
        }
    } else {
        var v39 uint64 = 18446744073709551615
        if len(v1) > 0 {
            v1[0] = v3
        }
        _ = v39
    }
    if (2147483648 <= 42) {
        fmt.Println(v5, v6)
    } else {
        fmt.Println(v6, v16)
    }
    var v40 uint8 = 130
    v41 := GenS6{F0: 234}
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = v3
    }
    var v42 int16 = 1
    v6 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
