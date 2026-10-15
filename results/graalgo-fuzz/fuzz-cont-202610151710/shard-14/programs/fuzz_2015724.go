package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -2147483648
    v3 := []int{2147483647}
    v4 := []string{"hello"}
    var v5 uint64 = 18446744073709551615
    v5 += (v5 - v5)
    for fi6 := 0; fi6 < 14; fi6++ {
        switch -9223372036854775808 {
        case 9223372036854775807:
            var v7 int32 = v2
            _ = v7
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi6
    }
    switch (("a" + "hello") + "foo bar") {
    case "":
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 int = -220
                v3 = append(v3, fi8)
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    case "hello":
        var v11 float64 = float64(-2.25)
        _ = v11
    default:
        if (!(!false)) {
            switch "foo bar" {
            case "\n\t":
                v2 -= v2
            case "foo bar":
                v12 := GenS1{F0: 0, F1: false}
                _ = v12
            case "hello":
                switch "foo bar" {
                case "\n\t":
                    v5 = v5
                case "hello":
                    for fi13 := 0; fi13 < 12; fi13++ {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            v4 = append(v4, "\n\t")
                            _ = fi14
                        }
                        if (false && (!true)) {
                            v15 := GenS2{F0: "a"}
                            _ = v15
                        } else {
                            if false {
                                var v16 string = ("a" + ("" + ""))
                                fmt.Println(v2)
                                _ = v16
                            }
                        }
                        _ = fi13
                    }
                case "a":
                    var v17 int32 = v2
                    _ = v17
                }
            default:
                switch 354 {
                case -2147483648:
                    if false {
                        var v18 uint64 = v5
                        if ((false && false) && ((v2 >> ((65535) & 63)) < -879)) {
                            var v19 int8 = -102
                            _ = v19
                        } else {
                            v3 = append(v3, -451)
                            for fi20 := 0; fi20 < 11; fi20++ {
                                var v21 string = "a"
                                if (false || true) {
                                    v21 = ((v21 + v21) + "foo bar")
                                }
                                fmt.Println(v18, v2, v18)
                                _ = fi20
                                _ = v21
                            }
                            v2 -= v2
                        }
                        _ = v18
                    }
                case 128:
                    var v22 int8 = int8(v2)
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 15; fi23++ {
                        v4 = append(v4, "")
                        v5 = (v5 % (((v5 << ((4294967295) & 63))) | 1))
                        switch (("" + "\n\t") + ("a" + "foo bar")) {
                        case "":
                            if false {
                                fmt.Println(fi23, fi23)
                                for fi24 := 0; fi24 < 3; fi24++ {
                                    v25 := GenS2{F0: (("foo bar" + "hello") + ("hello" + ((("a" + ("\n\t" + "\n\t")) + "") + ("hello" + ("foo bar" + ("hello" + "hello"))))))}
                                    switch ("\n\t" + ("a" + ("a" + "a"))) {
                                    case "":
                                        var v26 int32 = 0
                                        _ = v26
                                    default:
                                        if false {
                                            v4 = append(v4, "hello")
                                            v2 = -1
                                        }
                                    }
                                    if false {
                                        var v27 float32 = ((0.0 * 1e10) * 1.5)
                                        v28 := GenS2{F0: "hello"}
                                        _ = v27
                                        _ = v28
                                    }
                                    _ = fi24
                                    _ = v25
                                }
                            }
                        case "a":
                            v29 := GenS3{F0: false}
                            _ = v29
                        }
                        _ = fi23
                    }
                }
            }
        }
    }
    v4 = append(v4, "a")
    var v30 uint64 = (v5 % ((0) | 1))
    var v31 float32 = (-2.25 / 3.14159)
    v2 -= v2
    switch "foo bar" {
    case "\n\t":
        v2 = (v2 + (v2 / ((v2) | 1)))
    default:
        for fi32 := 0; fi32 < 11; fi32++ {
            for fi33 := 0; fi33 < 4; fi33++ {
                fmt.Println(v5, fi33, v5)
                var v34 bool = (!(false && (float32(1e10) <= 0.1)))
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
    }
    var v35 float32 = float32((1e10 + float64((3.14159 * 1e10))))
    var v36 uint64 = (v30 >> ((839) & 63))
    var v37 int8 = 127
    switch "foo bar" {
    case "hello":
        v38 := GenS1{F0: 42, F1: (v2 <= (v2 ^ v2))}
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
