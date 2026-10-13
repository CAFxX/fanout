package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint32 }
func genFn0() string {
    var v6 float64 = (1.5 + ((0.0 - -2.25) + (3.14159 / 0.1)))
    var v7 uint32 = 133
    switch "hello" {
    case "":
        var v8 int32 = 0
        _ = v8
    case "foo bar":
        switch ("hello" + "\n\t") {
        case "\n\t":
            for fi9 := 0; fi9 < 15; fi9++ {
                if (true || false) {
                    var v10 float32 = (3.14159 / float32(v6))
                    _ = v10
                } else {
                    switch fi9 {
                    case -329:
                        for fi11 := 0; fi11 < 15; fi11++ {
                            v12 := GenS1{F0: -2147483648, F1: 1}
                            v6 += v6
                            _ = fi11
                            _ = v12
                        }
                    }
                    fmt.Println(v6)
                    if false {
                        var v13 uint8 = 197
                        _ = v13
                    }
                }
                v7 += 1
                v7 = (v7 >> ((18446744073709551615) & 63))
                _ = fi9
            }
        case "hello":
            fmt.Println(v6)
        case "":
            var v14 uint32 = 211
            _ = v14
        }
    }
    _ = v6
    _ = v7
    return ""
}
func genFn1(p0 int32, p1 uint64, p2 string) int64 {
    var v15 string = ("foo bar" + p2)
    switch v15 {
    case "foo bar":
        var v16 int64 = -1
        _ = v16
    case "a":
        var v17 int32 = p0
        _ = v17
    case "\n\t":
        var v18 bool = ((^(p1)) != (p1 + 0))
        _ = v18
    }
    var v19 uint = 465
    if false {
        v19 = (v19 + (v19 + v19))
        p0 += p0
    } else {
        v20 := []int{}
        fmt.Println(v19, v15, p2)
        for fi21 := 0; fi21 < 10; fi21++ {
            if (!((v19 ^ v19) == (^(v19)))) {
                if true {
                    v19 = 255
                    var v22 uint16 = 822
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    _ = v22
                }
                p2 = ""
            } else {
                v20 = append(v20, (^(fi21)))
                p1 = (p1 + p1)
            }
            fmt.Println(p1, v15)
            _ = fi21
        }
        _ = v20
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v19
    return 42
}

func main() {
    var v1 int8 = -79
    v2 := []string{("foo bar" + ""), "foo bar"}
    var v3 int16 = -828
    var v4 float32 = (0.1 + (1e10 * (((1.5 / (-2.25 + (((-2.25 / 1.5) + 1e10) / 1.0))) * 1.5) * ((-2.25 / 1e10) * ((float32(3.14159) * -2.25) / (float32((float64(1e10) - 0.1)) + 0.1))))))
    var v5 string = "a"
    fmt.Println(v4, v1)
    switch -1 {
    case 556:
        v23 := GenS1{F0: -2147483648, F1: 574}
        _ = v23
    case -120:
        v4 += v4
    case 42:
        v1 = (-(v1))
    }
    var v24 int8 = (v1 & v1)
    if len(v2) > 0 {
        v2[0] = ("a" + v5)
    }
    fmt.Println(v5, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v25 float32 = (v4 + 0.0)
    fmt.Println(v1, v1, v25)
    v3 = 32767
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 uint16 = 1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v28 uint32 = 2147483647
        _ = fi26
        _ = v27
        _ = v28
    }
    switch "foo bar" {
    case "\n\t":
        if ((2147483647 <= -537) && (false && (-1 > 2147483647))) {
            v4 = v4
            switch "" {
            case "\n\t":
                v25 = (-2.25 / v25)
            }
        } else {
            if (!false) {
                switch ((v5 + v5) + v5) {
                case "hello":
                    switch v5 {
                    case "a":
                        switch 2147483647 {
                        case -2147483648:
                            switch (v5 + v5) {
                            case "hello":
                                if true {
                                    fmt.Println(v3, v5)
                                }
                            case "a":
                                v29 := GenS2{F0: 631, F1: 32767}
                                _ = v29
                            case "foo bar":
                                v30 := GenS1{F0: 1, F1: 255}
                                _ = v30
                            }
                        case 0:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    case "foo bar":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "hello":
                        fmt.Println(v4, v25, v1)
                    default:
                        _ = len(v2)
                    }
                }
                var v31 int8 = 100
                v32 := GenS2{F0: 9223372036854775807, F1: 2}
                _ = v31
                _ = v32
            }
        }
    case "":
        v3 = v3
    default:
        fmt.Println(v1)
    }
    if len(v2) > 0 {
        v2[0] = (v5 + v5)
    }
    if true {
        for fi33 := 0; fi33 < 7; fi33++ {
            var v34 uint = 65535
            v35 := GenS3{F0: 65535}
            _ = fi33
            _ = v34
            _ = v35
        }
        v24 = v1
        for fi36 := 0; fi36 < 6; fi36++ {
            v2 = append(v2, ("hello" + (v5 + (v5 + "foo bar"))))
            v2 = append(v2, v5)
            for fi37 := 0; fi37 < 12; fi37++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v38 := GenS1{F0: 42, F1: 259}
                _ = fi37
                _ = v38
            }
            _ = fi36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
