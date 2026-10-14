package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int32 }

func main() {
    var v1 uint8 = 1
    var v2 int = 357
    var v3 string = "foo bar"
    v4 := []string{"a", (("a" + "hello") + v3), v3}
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            if true {
                var v7 uint32 = 0
                _ = v7
            } else {
                var v8 float64 = (0.1 + 1e10)
                var v9 uint64 = 2
                switch v3 {
                case "hello":
                    v9 += 313
                case "foo bar":
                    v10 := GenS1{F0: v1, F1: (v9 + v9)}
                    _ = v10
                case "a":
                    v11 := GenS1{F0: v1, F1: 309}
                    _ = v11
                }
                _ = v8
                _ = v9
            }
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                var v12 uint = 18446744073709551615
                _ = v12
            }
            v13 := GenS1{F0: (v1 ^ 255), F1: 18446744073709551615}
            _ = fi6
            _ = v13
        }
        v14 := GenS1{F0: (v1 % (((v1 >> ((255) & 63))) | 1)), F1: 807}
        _ = fi5
        _ = v14
    }
    var v15 uint32 = 1
    switch v3 {
    case "":
        var v16 uint = 4294967295
        _ = v16
    default:
        if false {
            switch ((v3 + "\n\t") + (v3 + v3)) {
            case "\n\t":
                var v17 uint16 = 42
                _ = v17
            case "foo bar":
                v4 = append(v4, v3)
            }
        }
    }
    v2 = v2
    if len(v4) > 0 {
        v4[0] = v3
    }
    v3 = ((((((v3 + v3) + ("\n\t" + v3)) + v3) + "\n\t") + "a") + (v3 + ("\n\t" + "")))
    var v18 int8 = -128
    if ("a" > (v3 + v3)) {
        if (((!false) && (("a" + (v3 + ((v3 + v3) + v3))) <= (("" + ("" + (v3 + (v3 + v3)))) + ("foo bar" + ("\n\t" + v3))))) || true) {
            fmt.Println(v1)
            _ = len(v4)
        } else {
            var v19 bool = false
            v1 -= 42
            _ = v19
        }
        v1 = (v1 >> ((42) & 63))
    } else {
        var v20 uint = 517
        v1 = (^(v1))
        _ = v20
    }
    if ((860 > 9223372036854775807) && false) {
        if len(v4) > 0 {
            v4[0] = "a"
        }
        switch 9223372036854775807 {
        case 128:
            var v21 float64 = (0.0 * 3.14159)
            _ = v21
        case 1:
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        default:
            v1 += (^(v1))
        }
    }
    if (0 <= 0) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    switch v2 {
    case 34:
        _ = len(v4)
    case -2147483648:
        fmt.Println(v18)
    }
    if true {
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
        if true {
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            var v22 string = ""
            var v23 uint = 255
            _ = v22
            _ = v23
        }
        switch ("\n\t" + v3) {
        case "\n\t":
            if true {
                v18 += (v18 - 42)
                v2 = v2
            } else {
                var v24 uint64 = 596
                for fi25 := 0; fi25 < 6; fi25++ {
                    var v26 uint = 978
                    _ = fi25
                    _ = v26
                }
                var v27 int32 = 2147483647
                _ = v24
                _ = v27
            }
        case "foo bar":
            for fi28 := 0; fi28 < 3; fi28++ {
                var v29 float64 = float64(1e10)
                if (!(!false)) {
                    var v30 bool = (!(!(false && true)))
                    v31 := GenS1{F0: (v1 / ((69) | 1)), F1: 1}
                    _ = v30
                    _ = v31
                }
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
                _ = fi28
                _ = v29
            }
        case "":
            if (!false) {
                v32 := GenS1{F0: (v1 ^ v1), F1: 18446744073709551615}
                _ = v32
            }
        default:
            switch (^(v2)) {
            case 0:
                var v33 int8 = (v18 / ((int8(v2)) | 1))
                _ = v33
            case -844:
                v3 = v3
            case 2147483647:
                v15 -= v15
            }
        }
    } else {
        for fi34 := 0; fi34 < 14; fi34++ {
            if len(v4) > 0 {
                v4[0] = v3
            }
            _ = fi34
        }
        v18 = -9
        v18 = v18
    }
    var v35 int = v2
    v36 := GenS2{F0: 1}
    v37 := GenS3{F0: 42, F1: 540}
    v38 := GenS3{F0: -991, F1: 2147483647}
    fmt.Println(v3)
    for fi39 := 0; fi39 < 7; fi39++ {
        fmt.Println(v1, v2)
        v1 = 127
        v40 := GenS1{F0: v1, F1: uint64(v1)}
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
