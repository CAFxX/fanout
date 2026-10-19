package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"hello", "\n\t", "\n\t"}
    v3 := []string{}
    v4 := make(map[uint32]bool)
    v1[("a" + "")] = 2147483647
    if true {
        var v5 int64 = -1
        _ = v5
    } else {
        var v6 float64 = 3.14159
        var v7 uint = 0
        _ = v6
        _ = v7
    }
    var v8 int = 127
    if false {
        v8 += 97
        v8 = v8
    } else {
        switch (v8 % ((4294967295) | 1)) {
        case 0:
            v9 := v1["a"]
            _ = v9
        case 4294967295:
            switch "\n\t" {
            case "\n\t":
                v10 := GenS1{F0: 42}
                _ = v10
            case "hello":
                var v11 uint = 2
                _ = v11
            case "foo bar":
                switch "\n\t" {
                case "":
                    if true {
                        var v12 bool = (v8 <= 32767)
                        v12 = v12
                        _ = v12
                    } else {
                        v8 += v8
                        v8 -= -9223372036854775808
                        var v13 int = (v8 >> ((255) & 63))
                        _ = v13
                    }
                }
            default:
                switch "foo bar" {
                case "foo bar":
                    v14 := GenS1{F0: -1}
                    _ = v14
                default:
                    v8 = v8
                }
            }
        case 1:
            var v15 int32 = -2147483648
            _ = v15
        }
        switch v8 {
        case 128:
            v2 = append(v2, "hello")
        case 553:
            switch 128 {
            case 2147483647:
                delete(v1, "")
            case 9223372036854775807:
                switch ("a" + "hello") {
                case "foo bar":
                    v1["\n\t"] = 4294967295
                default:
                    var v16 int = -215
                    _ = v16
                }
            default:
                for fi17 := 0; fi17 < 1; fi17++ {
                    for fi18 := 0; fi18 < 17; fi18++ {
                        switch ((("foo bar" + "hello") + ("a" + "foo bar")) + "a") {
                        case "hello":
                            if (((!true) && (float32(3.14159) > (((-2.25 / 1e10) + (0.1 / (float32((0.0 - (float64((0.1 / -2.25)) * (-2.25 - float64(1e10))))) * 0.1))) - (((1.5 - ((3.14159 / (float32((-2.25 * 0.1)) / 1.0)) - (3.14159 - (0.0 + 0.0)))) * (((1e10 * (float32(-2.25) * float32(0.0))) + 3.14159) * 3.14159)) + 0.1)))) || true) {
                                var v19 int8 = 61
                                fmt.Println(fi17, fi17)
                                var v20 uint32 = 109
                                _ = v19
                                _ = v20
                            }
                        }
                        _ = fi18
                    }
                    var v21 bool = true
                    _ = fi17
                    _ = v21
                }
            }
        default:
            var v22 bool = true
            _ = v22
        }
    }
    if len(v3) > 0 {
        v3[0] = ("foo bar" + (("hello" + ((("a" + "") + ("a" + ("a" + "hello"))) + "")) + "hello"))
    }
    if false {
        var v23 int64 = 42
        if true {
            for fi24 := 0; fi24 < 4; fi24++ {
                switch 335 {
                case 32767:
                    fmt.Println(v23, fi24, fi24)
                }
                _ = fi24
            }
            for fi25 := 0; fi25 < 7; fi25++ {
                switch (("" + ("hello" + "\n\t")) + ("a" + "a")) {
                case "\n\t":
                    v26 := GenS1{F0: v23}
                    _ = v26
                case "foo bar":
                    fmt.Println(v23)
                case "a":
                    for fi27 := 0; fi27 < 5; fi27++ {
                        var v28 uint8 = 1
                        v29 := GenS1{F0: v23}
                        var v30 int16 = -32768
                        v30 = v30
                        if true {
                            v3 = append(v3, ("" + "foo bar"))
                            v28 += (v28 / ((v28) | 1))
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                }
                v23 -= -1
                if (65535 != 167) {
                    v8 += 128
                    v4[796] = (!false)
                } else {
                    var v31 float64 = 1.5
                    _ = v31
                }
                _ = fi25
            }
            switch v8 {
            case -897:
                var v32 float32 = 0.0
                _ = v32
            case -2147483648:
                var v33 float64 = 1e10
                _ = v33
            }
        }
        _ = v23
    }
    var v34 uint8 = 113
    v35 := GenS2{F0: (^(v8)), F1: 892}
    fmt.Println(v8, v34)
    v8 += (v8 - v8)
    v36 := GenS2{F0: (v8 & v8), F1: -267}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v8
}
