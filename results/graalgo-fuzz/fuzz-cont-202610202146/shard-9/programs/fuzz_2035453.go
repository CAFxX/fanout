package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int }
type GenS2 struct { F0 uint8; F1 bool }
func genFn0(p0 int32, p1 int16, p2 uint8) int32 {
    p1 = p1
    if false {
        v4 := []int{}
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 uint16 = 1
            _ = fi5
            _ = v6
        }
        var v7 int16 = (^(p1))
        _ = v4
        _ = v7
    }
    p0 -= p0
    v8 := []int{-9223372036854775808, -9223372036854775808, int(p2), 32767}
    if false {
        var v9 string = ("foo bar" + ("a" + "a"))
        switch 109 {
        case 0:
            if (v9 == "hello") {
                var v10 float64 = 0.1
                _ = v10
            }
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p0
}
func genFn1(p0 uint32, p1 int32) uint64 {
    switch 0 {
    case 32767:
        var v11 float64 = 1e10
        _ = v11
    case 42:
        var v12 uint16 = 2
        _ = v12
    default:
        p1 = -1
    }
    v13 := []int{}
    _ = p0
    _ = p1
    _ = v13
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -2147483648
    var v3 int64 = -9223372036854775808
    v14 := GenS1{F0: 1, F1: ("" + "\n\t"), F2: v2}
    fmt.Println(v3, v3)
    if (!false) {
        switch -678 {
        case -2147483648:
            switch ("foo bar" + "hello") {
            case "a":
                fmt.Println(v3, v3)
            case "foo bar":
                switch 2147483647 {
                case 127:
                    var v15 int64 = -9223372036854775808
                    _ = v15
                case -1:
                    switch -2147483648 {
                    case 1:
                        switch "a" {
                        case "hello":
                            _ = len(v1)
                        default:
                            var v16 int32 = -680
                            _ = v16
                        }
                    case 32768:
                        switch "foo bar" {
                        case "a":
                            var v17 uint16 = 32767
                            _ = v17
                        case "hello":
                            v18 := GenS1{F0: 72, F1: (("" + "foo bar") + ("hello" + "foo bar")), F2: v2}
                            _ = v18
                        case "":
                            v19 := []int{-900, v2, (v2 ^ (v2 - (v2 + (^(v2))))), (v2 + v2)}
                            _ = v19
                        default:
                            var v20 int32 = -601
                            _ = v20
                        }
                    default:
                        switch (-(v2)) {
                        case 32768:
                            var v21 bool = (2 >= 18446744073709551615)
                            _ = v21
                        case 9223372036854775807:
                            v22 := []int{(v2 * v2), -1, v2}
                            _ = v22
                        case 315:
                            if (("a" + "foo bar") >= ("foo bar" + (("\n\t" + "a") + "foo bar"))) {
                                switch (-(v2)) {
                                case -9223372036854775808:
                                    fmt.Println(v3, v3)
                                case 128:
                                    var v23 uint8 = 42
                                    _ = v23
                                }
                                v3 += 9223372036854775807
                            }
                        }
                    }
                case 128:
                    v3 -= v3
                default:
                    var v24 uint8 = 127
                    _ = v24
                }
            }
        case 9223372036854775807:
            var v25 uint16 = 400
            _ = v25
        case 2147483647:
            fmt.Println(v2)
        default:
            if false {
                v3 = 503
                v3 -= v3
                fmt.Println(v3, v3)
            }
        }
        for fi26 := 0; fi26 < 17; fi26++ {
            switch "a" {
            case "foo bar":
                v27 := GenS2{F0: 255, F1: true}
                _ = v27
            case "":
                if true {
                    v2 = -1
                    var v28 uint8 = 192
                    _ = v28
                } else {
                    fmt.Println(fi26, v2, fi26)
                    var v29 uint8 = 128
                    v29 = v29
                    _ = v29
                }
            default:
                var v30 uint8 = 0
                _ = v30
            }
            fmt.Println(v3)
            v3 += int64(fi26)
            _ = fi26
        }
        v2 = v2
    }
    v31 := []int{v2, (v2 << ((532) & 63)), 0, 9223372036854775807}
    if false {
        v1[("hello" + "a")] = 199
        if (1.5 < -2.25) {
            var v32 int8 = -83
            _ = v32
        }
    }
    v3 = (v3 >> ((65535) & 63))
    v2 -= (v2 & (v2 - 0))
    delete(v1, "")
    v3 = (v3 + v3)
    switch ((((("hello" + "") + "a") + "") + ("a" + "")) + (("" + ("a" + "foo bar")) + (("\n\t" + "") + "a"))) {
    case "a":
        switch (v2 % (((v2 & 2147483647)) | 1)) {
        case 32767:
            if true {
                v3 += v3
                v31 = append(v31, (v2 | 2147483647))
            } else {
                for fi33 := 0; fi33 < 2; fi33++ {
                    delete(v1, "foo bar")
                    var v34 uint32 = 1
                    v2 = v2
                    _ = fi33
                    _ = v34
                }
            }
        case 1:
            fmt.Println(v3, v2)
        default:
            var v35 float64 = (3.14159 - (0.0 - 0.1))
            _ = v35
        }
    case "foo bar":
        var v36 string = (("a" + ("a" + (("\n\t" + ("a" + "foo bar")) + ("foo bar" + "hello")))) + ("a" + "foo bar"))
        _ = v36
    }
    v3 = 144
    for fi37 := 0; fi37 < 2; fi37++ {
        v3 += (v3 & 9223372036854775807)
        _ = fi37
    }
    if len(v31) > 0 {
        v31[0] = (v2 ^ (-(v2)))
    }
    v2 += 622
    v2 = v2
    for fi38 := 0; fi38 < 18; fi38++ {
        var v39 string = "hello"
        fmt.Println(v3, v2)
        fmt.Println(fi38, fi38, v2)
        _ = fi38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
}
