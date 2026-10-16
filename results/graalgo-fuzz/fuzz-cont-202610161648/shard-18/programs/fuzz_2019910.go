package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 string; F2 uint }
type GenS5 struct { F0 uint32; F1 uint16; F2 int16 }

func main() {
    var v1 string = "a"
    v2 := []int{249, -2147483648, -1}
    v3 := make(map[int]int)
    var v4 float32 = 0.0
    fmt.Println(v4, v1)
    v5 := GenS1{F0: (v1 + v1), F1: -677, F2: 32767}
    if true {
        var v6 bool = (789 < 718)
        _ = v6
    }
    switch "a" {
    case "a":
        switch "foo bar" {
        case "":
            if false {
                if (739 <= 42) {
                    var v7 bool = true
                    if v7 {
                        var v8 uint = 18446744073709551615
                        v9 := GenS2{F0: 32768, F1: 32768}
                        _ = v8
                        _ = v9
                    } else {
                        v7 = v7
                        if (!(-57 >= -12)) {
                            v1 = v1
                            if v7 {
                                v4 = 1.5
                                v1 = v1
                            } else {
                                delete(v3, -9223372036854775808)
                                v4 += v4
                                var v10 uint64 = 1
                                _ = v10
                            }
                        } else {
                            fmt.Println(v1, v4, v4)
                            for fi11 := 0; fi11 < 8; fi11++ {
                                if (v7 && (0 == -504)) {
                                    if (!(-1 <= -356)) {
                                        for fi12 := 0; fi12 < 10; fi12++ {
                                            var v13 bool = ((fi12 ^ -1) < (fi12 ^ (fi11 + -143)))
                                            v14 := GenS3{F0: 65535}
                                            v2 = append(v2, -1)
                                            if (v7 || v7) {
                                                v15 := GenS1{F0: ("" + v1), F1: 0, F2: -32768}
                                                _ = v15
                                            } else {
                                                _ = len(v2)
                                                if len(v2) > 0 {
                                                    v2[0] = (fi11 + fi12)
                                                }
                                            }
                                            _ = fi12
                                            _ = v13
                                            _ = v14
                                        }
                                    }
                                    fmt.Println(v4, v7)
                                } else {
                                    fmt.Println(v7, v7)
                                }
                                var v16 uint = 42
                                _ = fi11
                                _ = v16
                            }
                        }
                        v17 := GenS3{F0: 4294967295}
                        _ = v17
                    }
                    v1 = "\n\t"
                    _ = v7
                }
                if false {
                    v4 += -2.25
                    var v18 float32 = 1.5
                    _ = v18
                }
            } else {
                var v19 uint8 = 0
                var v20 string = (v1 + (((v1 + (("foo bar" + "a") + v1)) + (v1 + "\n\t")) + v1))
                v20 = v20
                _ = v19
                _ = v20
            }
        case "\n\t":
            _ = len(v3)
        case "a":
            for fi21 := 0; fi21 < 2; fi21++ {
                for fi22 := 0; fi22 < 20; fi22++ {
                    v4 += (v4 - v4)
                    v2 = append(v2, (fi22 ^ (fi21 - fi21)))
                    _ = fi22
                }
                var v23 uint8 = 2
                _ = fi21
                _ = v23
            }
        default:
            v1 = ("foo bar" + ("" + "\n\t"))
        }
    case "\n\t":
        switch v1 {
        case "hello":
            fmt.Println(v4, v4)
        case "":
            v24 := GenS4{F0: (1e10 != (3.14159 * 3.14159)), F1: v1, F2: 255}
            _ = v24
        }
    }
    if false {
        v1 = v1
    }
    v25 := GenS4{F0: false, F1: (v1 + "a"), F2: 65535}
    if len(v2) > 0 {
        v2[0] = -854
    }
    v1 = (v1 + ("hello" + v1))
    if ((-2.25 * 1e10) != -2.25) {
        if true {
            switch (((v1 + v1) + v1) + v1) {
            case "":
                for fi26 := 0; fi26 < 12; fi26++ {
                    if (1 == 4294967295) {
                        v2 = append(v2, fi26)
                        v27 := v3[375]
                        _ = v27
                    } else {
                        fmt.Println(v1)
                        switch -244 {
                        case -9223372036854775808:
                            if (true && (!(false || (0.1 == 0.0)))) {
                                v2 = append(v2, fi26)
                            } else {
                                switch (fi26 % ((fi26) | 1)) {
                                case 127:
                                    if false {
                                        fmt.Println(fi26, v4, fi26)
                                        for fi28 := 0; fi28 < 7; fi28++ {
                                            switch v1 {
                                            case "\n\t":
                                                fmt.Println(v1)
                                            case "":
                                                if true {
                                                    fmt.Println(v1, fi26)
                                                    fmt.Println(fi26, fi28)
                                                }
                                            }
                                            fmt.Println(fi28, fi28, fi26)
                                            fmt.Println(v1)
                                            _ = fi28
                                        }
                                    }
                                }
                                fmt.Println(fi26, v1, fi26)
                            }
                        }
                        var v29 uint32 = 4294967295
                        _ = v29
                    }
                    _ = fi26
                }
            default:
                var v30 string = (("\n\t" + v1) + ((v1 + "\n\t") + v1))
                _ = v30
            }
        }
    } else {
        v3[-2147483648] = -970
        if false {
            for fi31 := 0; fi31 < 14; fi31++ {
                fmt.Println(v1, v4)
                v32 := GenS4{F0: false, F1: "foo bar", F2: 907}
                _ = fi31
                _ = v32
            }
            fmt.Println(v1)
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v4 = (((((v4 / -2.25) / 3.14159) / ((1e10 / v4) / (v4 + v4))) - v4) + float32(float64((v4 / -2.25))))
        }
        if (2147483647 > 2147483647) {
            v33 := GenS4{F0: true, F1: v1, F2: 42}
            for fi34 := 0; fi34 < 7; fi34++ {
                fmt.Println(v4, fi34, v4)
                _ = fi34
            }
            for fi35 := 0; fi35 < 3; fi35++ {
                var v36 string = v1
                for fi37 := 0; fi37 < 6; fi37++ {
                    var v38 uint64 = 413
                    v36 = v36
                    v38 += v38
                    _ = fi37
                    _ = v38
                }
                for fi39 := 0; fi39 < 16; fi39++ {
                    var v40 uint32 = 55
                    v41 := GenS5{F0: (v40 % (((^(v40))) | 1)), F1: 65535, F2: 138}
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                _ = fi35
                _ = v36
            }
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
