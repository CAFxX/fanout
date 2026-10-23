package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int8; F2 int64 }

func main() {
    var v1 uint8 = 2
    v2 := []string{(("\n\t" + "") + ("hello" + "a")), ("hello" + "foo bar")}
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    v5 := make(map[string]int)
    var v6 uint16 = 0
    switch ("foo bar" + "foo bar") {
    case "":
        if (-128 == 127) {
            v7 := GenS1{F0: -1, F1: 128, F2: "hello"}
            for fi8 := 0; fi8 < 16; fi8++ {
                for fi9 := 0; fi9 < 11; fi9++ {
                    v10 := GenS1{F0: 34, F1: (v1 + 221), F2: "\n\t"}
                    _ = fi9
                    _ = v10
                }
                _ = len(v2)
                _ = fi8
            }
            _ = v7
        } else {
            if ((v1 + (^(v1))) >= v1) {
                var v11 int = -845
                if true {
                    var v12 uint64 = 42
                    _ = v12
                }
                var v13 string = "a"
                _ = v11
                _ = v13
            }
        }
    default:
        fmt.Println(v6, v1)
    }
    if false {
        var v14 uint8 = v1
        if true {
            if (127 > 362) {
                v6 -= (v6 | v6)
            }
            var v15 uint32 = 65535
            switch "" {
            case "":
                v6 = v6
            case "\n\t":
                switch "" {
                case "a":
                    v16 := GenS1{F0: 100, F1: v1, F2: "hello"}
                    _ = v16
                }
            case "foo bar":
                var v17 uint32 = (v15 >> ((958) & 63))
                _ = v17
            }
            _ = v15
        }
        var v18 float32 = (((((0.1 * (3.14159 / 3.14159)) * 0.0) * 0.0) / 1.0) + 3.14159)
        _ = v14
        _ = v18
    } else {
        v6 = v6
        for fi19 := 0; fi19 < 1; fi19++ {
            v3[370] = 2147483647
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 int32 = int32(fi19)
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    }
    v6 = 562
    fmt.Println(v6, v6)
    v22 := GenS2{F0: 127}
    switch int(v6) {
    case 128:
        switch "\n\t" {
        case "":
            var v23 bool = true
            _ = v23
        case "\n\t":
            for fi24 := 0; fi24 < 13; fi24++ {
                delete(v5, "foo bar")
                _ = fi24
            }
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = ""
            }
        default:
            var v25 int8 = 42
            _ = v25
        }
    case -771:
        if false {
            var v26 int = 9223372036854775807
            v27 := GenS3{F0: v6, F1: -128, F2: -832}
            _ = v26
            _ = v27
        }
    case 32768:
        for fi28 := 0; fi28 < 7; fi28++ {
            if (2 > 195) {
                if false {
                    var v29 int64 = -334
                    fmt.Println(fi28)
                    _ = v29
                }
                var v30 uint16 = (^(v6))
                delete(v3, -597)
                _ = v30
            }
            _ = fi28
        }
    default:
        var v31 uint64 = 0
        _ = v31
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 uint32 = 4294967295
        _ = fi32
        _ = v33
    }
    v34 := GenS1{F0: 0, F1: (v1 / (((v1 >> ((4294967295) & 63))) | 1)), F2: ("\n\t" + (((((((("\n\t" + ("hello" + (("" + "hello") + "\n\t"))) + "hello") + "a") + ("a" + "foo bar")) + (("" + ("" + ("\n\t" + ""))) + "\n\t")) + ("" + (("foo bar" + (("\n\t" + "foo bar") + "foo bar")) + ("" + "foo bar")))) + ("hello" + ("hello" + "a"))) + "\n\t"))}
    _ = len(v3)
    if (!false) {
        switch 32767 {
        case 1:
            var v35 string = "a"
            _ = v35
        default:
            var v36 uint64 = 1
            _ = v36
        }
    } else {
        v6 -= (^(v6))
        var v37 uint8 = (v1 >> ((4294967295) & 63))
        switch -525 {
        case 0:
            v38 := GenS3{F0: v6, F1: 100, F2: -9223372036854775808}
            _ = v38
        case -865:
            var v39 int = 9223372036854775807
            _ = v39
        }
        _ = v37
    }
    var v40 bool = ((true || ((v1 >> ((932) & 63)) == (v1 + (^(v1))))) || true)
    for fi41 := 0; fi41 < 12; fi41++ {
        v42 := GenS3{F0: (v6 >> ((0) & 63)), F1: 94, F2: -533}
        var v43 float64 = (1.5 / -2.25)
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
