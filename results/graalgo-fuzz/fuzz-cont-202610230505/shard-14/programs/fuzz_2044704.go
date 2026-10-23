package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (501 <= 255)
    v3 := []string{"", "hello", "foo bar", (("" + "\n\t") + "\n\t")}
    var v4 int8 = 1
    delete(v1, 0)
    v2 = (!false)
    if (2 < 122) {
        fmt.Println(v4, v2)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        v2 = (false || (v2 && (18446744073709551615 != 288)))
    }
    v2 = (true && false)
    v5 := GenS1{F0: (216 == 4294967295)}
    if ((!v2) || v2) {
        v6 := GenS1{F0: (1 != 1)}
        fmt.Println(v2)
        var v7 uint16 = 8
        _ = v6
        _ = v7
    } else {
        var v8 int32 = 0
        _ = v8
    }
    switch ("foo bar" + ("foo bar" + "a")) {
    case "a":
        for fi9 := 0; fi9 < 4; fi9++ {
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            _ = fi9
        }
    case "\n\t":
        var v10 int8 = -77
        _ = v10
    default:
        fmt.Println(v2, v2)
    }
    if (v4 > (v4 * v4)) {
        var v11 int8 = v4
        _ = v11
    }
    v2 = (4294967295 < 65535)
    v4 = -1
    for fi12 := 0; fi12 < 1; fi12++ {
        v4 = v4
        switch ((("\n\t" + "foo bar") + "foo bar") + "\n\t") {
        case "hello":
            if false {
                if (!false) {
                    fmt.Println(v2, fi12, v2)
                    v1[(-(fi12))] = 544
                }
            }
        default:
            v4 -= (v4 % (((v4 & v4)) | 1))
        }
        v13 := v1[fi12]
        _ = fi12
        _ = v13
    }
    fmt.Println(v4, v2)
    if (!v2) {
        switch -94 {
        case 2147483647:
            v1[875] = 0
        case -904:
            if (65535 > 2) {
                v14 := GenS1{F0: false}
                v15 := GenS1{F0: ((-2.25 / -2.25) != (0.0 - ((-2.25 / ((1.5 + ((float64(float32(3.14159)) * 3.14159) * -2.25)) / 0.1)) - 0.1)))}
                for fi16 := 0; fi16 < 17; fi16++ {
                    switch ("a" + "") {
                    case "":
                        v2 = ((!true) || v2)
                    }
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    for fi17 := 0; fi17 < 12; fi17++ {
                        fmt.Println(fi17, v4)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi17
                    }
                    _ = fi16
                }
                _ = v14
                _ = v15
            }
        case 32767:
            v3 = append(v3, ("\n\t" + ("" + "hello")))
        }
        v2 = v2
        v4 = v4
    } else {
        switch "" {
        case "\n\t":
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 uint8 = 127
                v19 -= v19
                fmt.Println(v4, fi18, v2)
                _ = fi18
                _ = v19
            }
        case "foo bar":
            v2 = v2
        case "":
            if false {
                switch -9223372036854775808 {
                case 9223372036854775807:
                    var v20 string = ("hello" + (("" + ((("foo bar" + "foo bar") + "foo bar") + "foo bar")) + ""))
                    _ = v20
                case 0:
                    fmt.Println(v2, v4)
                case -168:
                    var v21 int64 = 972
                    _ = v21
                }
                var v22 int64 = 285
                _ = v22
            }
        }
        fmt.Println(v2)
    }
    if false {
        if v2 {
            v23 := GenS1{F0: v2}
            var v24 bool = (("" != "") || true)
            fmt.Println(v2, v2, v2)
            _ = v23
            _ = v24
        } else {
            v4 = 46
            var v25 int64 = 1
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 uint8 = 127
        v28 := GenS1{F0: ((v4 <= v4) || (255 <= 65535))}
        _ = fi26
        _ = v27
        _ = v28
    }
    v2 = (((-2.25 + 0.0) <= ((float32((1.5 - 1e10)) - 0.1) + 1e10)) && v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
