package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int32; F2 bool }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := []string{(("hello" + ("\n\t" + "a")) + ("a" + (("a" + "hello") + ("" + "\n\t")))), "", "a", "a"}
    var v2 int64 = 248
    var v3 uint = 18446744073709551615
    var v4 uint8 = 128
    v5 := make(map[int]uint32)
    v4 -= (v4 | v4)
    var v6 bool = true
    switch -9223372036854775808 {
    case 23:
        switch (("foo bar" + "hello") + "a") {
        case "a":
            var v7 bool = ((1.5 - 0.1) > -2.25)
            _ = v7
        case "foo bar":
            if true {
                if len(v1) > 0 {
                    v1[0] = ((("hello" + "a") + "hello") + "")
                }
                for fi8 := 0; fi8 < 6; fi8++ {
                    v6 = true
                    if false {
                        for fi9 := 0; fi9 < 19; fi9++ {
                            if (uint64(v3) <= 4294967295) {
                                v1 = append(v1, "hello")
                            } else {
                                switch "" {
                                case "foo bar":
                                    _ = len(v5)
                                case "":
                                    var v10 string = "a"
                                    _ = v10
                                }
                            }
                            _ = fi9
                        }
                        v6 = ((v3 >> ((v3) & 63)) < v3)
                        var v11 int64 = v2
                        _ = v11
                    }
                    _ = fi8
                }
                if (0.0 != 1.5) {
                    var v12 bool = (127 <= -1)
                    _ = v12
                }
            } else {
                v4 -= uint8(v3)
                switch (((("" + ("\n\t" + ((("hello" + "hello") + "") + "a"))) + (("\n\t" + ((("\n\t" + "foo bar") + "") + "\n\t")) + ("hello" + ("\n\t" + "a")))) + "foo bar") + "a") {
                case "hello":
                    fmt.Println(v4)
                case "\n\t":
                    var v13 string = ("\n\t" + "\n\t")
                    _ = v13
                }
            }
        }
    }
    var v14 uint8 = v4
    var v15 int64 = v2
    for fi16 := 0; fi16 < 11; fi16++ {
        if false {
            v17 := GenS1{F0: 411}
            v14 -= 2
            _ = v17
        } else {
            switch ("a" + "") {
            case "":
                var v18 uint8 = 128
                _ = v18
            case "a":
                var v19 int16 = 32767
                _ = v19
            }
        }
        v5[(^(fi16))] = 269
        _ = fi16
    }
    switch (((("foo bar" + (("\n\t" + "") + ((("foo bar" + "a") + "\n\t") + ("\n\t" + ("foo bar" + "hello"))))) + "hello") + "") + ("foo bar" + "\n\t")) {
    case "\n\t":
        if (!v6) {
            v6 = (true && (v14 != v4))
            _ = len(v5)
            v20 := GenS1{F0: 2}
            _ = v20
        }
    default:
        v21 := GenS2{F0: "hello"}
        _ = v21
    }
    if v6 {
        v22 := GenS3{F0: 65535, F1: 42, F2: ("foo bar" > "foo bar")}
        _ = v22
    } else {
        var v23 bool = ((((0.1 / 1.0) + 1.5) + -2.25) >= (0.1 * (3.14159 + -2.25)))
        switch 104 {
        case 32768:
            v24 := v5[-9223372036854775808]
            _ = v24
        case 32767:
            if (int64(v3) <= (v15 / (((v15 * (-(v15)))) | 1))) {
                fmt.Println(v3, v3, v4)
                for fi25 := 0; fi25 < 6; fi25++ {
                    for fi26 := 0; fi26 < 4; fi26++ {
                        if true {
                            if true {
                                switch "\n\t" {
                                case "\n\t":
                                    v5[fi26] = 42
                                case "a":
                                    v2 += (v2 * v15)
                                }
                                var v27 int16 = 0
                                var v28 uint64 = 42
                                _ = v27
                                _ = v28
                            }
                            v29 := GenS2{F0: ("\n\t" + ("\n\t" + "foo bar"))}
                            var v30 bool = (v14 > v14)
                            _ = v29
                            _ = v30
                        }
                        var v31 int16 = 42
                        v6 = (-1 >= 457)
                        _ = fi26
                        _ = v31
                    }
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(fi25, v6)
                    }
                    _ = fi25
                }
                for fi32 := 0; fi32 < 6; fi32++ {
                    v2 = -9223372036854775808
                    switch int(v15) {
                    case 1:
                        v3 = v3
                    case 9223372036854775807:
                        switch "" {
                        case "":
                            for fi33 := 0; fi33 < 20; fi33++ {
                                v5[51] = uint32(v14)
                                v6 = (1 < 42)
                                if ((false || v6) || true) {
                                    v14 = 42
                                    if (v6 && v23) {
                                        delete(v5, 776)
                                    } else {
                                        v14 += v4
                                    }
                                } else {
                                    delete(v5, -1)
                                    var v34 string = "\n\t"
                                    switch (^(fi33)) {
                                    case 262:
                                        v35 := GenS2{F0: v34}
                                        if (123 >= 0) {
                                            var v36 int32 = -1
                                            _ = v36
                                        }
                                        _ = v35
                                    case -472:
                                        v37 := GenS4{F0: v15, F1: 165}
                                        _ = v37
                                    case 610:
                                        if (32767 > 32767) {
                                            if len(v1) > 0 {
                                                v1[0] = (v34 + v34)
                                            }
                                            v3 = v3
                                            if v23 {
                                                v38 := GenS5{F0: (v4 << (((v3 & (v3 ^ v3))) & 63)), F1: 65535}
                                                var v39 int8 = 127
                                                if v23 {
                                                    v34 = v34
                                                    var v40 int16 = 32767
                                                    if (v15 >= 42) {
                                                        v23 = (v39 <= v39)
                                                    }
                                                    _ = v40
                                                }
                                                _ = v38
                                                _ = v39
                                            }
                                        } else {
                                            v41 := GenS3{F0: 0, F1: -1, F2: true}
                                            v4 -= v4
                                            var v42 int16 = 9
                                            fmt.Println(fi32, v2, fi33)
                                            _ = v41
                                            _ = v42
                                        }
                                    }
                                    _ = v34
                                }
                                fmt.Println(v14, fi33)
                                fmt.Println(v14, v14)
                                _ = fi33
                            }
                        case "hello":
                            v15 = v2
                        case "foo bar":
                            var v43 int64 = -346
                            _ = v43
                        }
                    case -2147483648:
                        var v44 uint = v3
                        _ = v44
                    }
                    _ = fi32
                }
            } else {
                var v45 bool = v6
                _ = v45
            }
        }
        v2 += 0
        _ = v23
    }
    var v46 uint32 = 318
    var v47 int8 = 1
    var v48 int = -874
    v2 -= (v2 * 932)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
