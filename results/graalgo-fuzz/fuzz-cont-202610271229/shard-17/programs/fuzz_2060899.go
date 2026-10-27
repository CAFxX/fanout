package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 bool }
func genFn0(p0 int32, p1 int16) string {
    v4 := []int{}
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: 105, F1: 1}
        v7 := GenS1{F0: 0, F1: 1}
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 bool = true
            _ = fi8
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    _ = p0
    _ = p1
    _ = v4
    return ("a" + ("a" + ""))
}
func genFn1(p0 int8) string {
    if false {
        v10 := []int{667, int(p0)}
        if (false && false) {
            if (((^(p0)) < p0) || (3.14159 > 0.1)) {
                _ = len(v10)
            }
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            case "a":
                if (!(false || (!(-424 < 9223372036854775807)))) {
                    if (42 == 128) {
                        if (!(2 > 197)) {
                            if (!false) {
                                var v11 int32 = 883
                                _ = v11
                            } else {
                                var v12 int32 = -2147483648
                                _ = v12
                            }
                        }
                        for fi13 := 0; fi13 < 1; fi13++ {
                            fmt.Println(p0, fi13)
                            for fi14 := 0; fi14 < 15; fi14++ {
                                if len(v10) > 0 {
                                    v10[0] = 127
                                }
                                fmt.Println(fi13, p0, fi13)
                                if ((p0 * (p0 << ((767) & 63))) != p0) {
                                    p0 = p0
                                    var v15 uint = 65535
                                    _ = v15
                                }
                                _ = fi14
                            }
                            v16 := GenS2{F0: fi13}
                            _ = fi13
                            _ = v16
                        }
                    }
                    p0 = p0
                }
            case "":
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
            var v17 int64 = 9223372036854775807
            _ = v17
        } else {
            p0 += (p0 * -19)
            fmt.Println(p0)
            for fi18 := 0; fi18 < 8; fi18++ {
                switch 9223372036854775807 {
                case -9223372036854775808:
                    var v19 uint8 = 1
                    _ = v19
                case 4294967295:
                    var v20 uint16 = 42
                    _ = v20
                }
                if false {
                    fmt.Println(fi18, fi18)
                    p0 -= (p0 - p0)
                } else {
                    v21 := GenS2{F0: fi18}
                    _ = v21
                }
                _ = fi18
            }
        }
        switch 128 {
        case -792:
            v22 := GenS2{F0: -751}
            _ = v22
        case 0:
            var v23 string = "\n\t"
            _ = v23
        case 278:
            for fi24 := 0; fi24 < 11; fi24++ {
                _ = len(v10)
                if ((("\n\t" + (("\n\t" + "") + "foo bar")) <= (("" + "") + "")) && (1e10 <= (1.5 / 1.0))) {
                    if len(v10) > 0 {
                        v10[0] = 32768
                    }
                    v10 = append(v10, fi24)
                    if (!false) {
                        p0 = p0
                    }
                }
                _ = fi24
            }
        }
        _ = v10
    }
    var v25 int8 = (p0 * p0)
    fmt.Println(v25)
    _ = p0
    _ = v25
    return (("\n\t" + "hello") + "\n\t")
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{}
    var v3 int16 = 457
    v1["\n\t"] = 51
    for fi26 := 0; fi26 < 9; fi26++ {
        v3 = v3
        var v27 uint32 = 188
        _ = fi26
        _ = v27
    }
    v3 += (v3 << ((232) & 63))
    fmt.Println(v3)
    switch ((((("\n\t" + ("foo bar" + "\n\t")) + "foo bar") + "\n\t") + "foo bar") + "hello") {
    case "":
        _ = len(v1)
    case "hello":
        var v28 float32 = 0.0
        _ = v28
    case "a":
        delete(v1, ("a" + "hello"))
    }
    v3 = 1
    for fi29 := 0; fi29 < 15; fi29++ {
        fmt.Println(fi29, fi29)
        var v30 int = fi29
        for fi31 := 0; fi31 < 6; fi31++ {
            if true {
                v32 := GenS3{F0: "", F1: false}
                v33 := v1[("a" + "")]
                _ = v32
                _ = v33
            }
            var v34 uint32 = 2147483648
            switch "\n\t" {
            case "hello":
                switch v30 {
                case 42:
                    v34 -= v34
                case 4294967295:
                    var v35 bool = (!(!((0.0 <= 3.14159) && true)))
                    _ = v35
                case 737:
                    v3 = (v3 << ((42) & 63))
                default:
                    var v36 uint16 = 67
                    _ = v36
                }
            case "a":
                v3 += (v3 + v3)
                v34 = 0
                v1[(((("foo bar" + ("foo bar" + (("\n\t" + ((("foo bar" + "foo bar") + (("hello" + "\n\t") + "hello")) + ("a" + ("foo bar" + (("a" + "a") + "foo bar"))))) + (("foo bar" + "\n\t") + "hello")))) + ("" + "foo bar")) + "\n\t") + "foo bar")] = v34
                v3 = int16(v34)
            case "":
                _ = len(v1)
            }
            _ = fi31
            _ = v34
        }
        _ = fi29
        _ = v30
    }
    switch "" {
    case "foo bar":
        v37 := GenS2{F0: -1}
        _ = v37
    }
    if (630 <= 42) {
        if (true && true) {
            v3 -= 1
            fmt.Println(v3)
            v38 := GenS1{F0: -128, F1: 255}
            _ = v38
        }
        for fi39 := 0; fi39 < 8; fi39++ {
            fmt.Println(fi39, fi39)
            var v40 uint64 = 1
            v3 += v3
            _ = fi39
            _ = v40
        }
    } else {
        v41 := GenS2{F0: 42}
        _ = v41
    }
    var v42 int64 = 9223372036854775807
    var v43 uint64 = 42
    v1[((("hello" + "foo bar") + "hello") + "foo bar")] = 2147483647
    v42 -= v42
    var v44 int64 = 42
    v43 += 709
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v44
}
