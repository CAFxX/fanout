package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 bool = (255 == 127)
    var v3 bool = ((((0.1 > -2.25) || true) || v2) || ((v2 || true) && ((867 > 105) || true)))
    var v4 uint = 2
    var v5 bool = true
    var v6 float32 = (0.0 - 0.0)
    var v7 string = ("foo bar" + ("a" + "hello"))
    v5 = (-1 < 45)
    v7 = v7
    switch ("\n\t" + "\n\t") {
    case "a":
        var v8 uint32 = 1
        _ = v8
    case "hello":
        switch 4294967295 {
        case -1:
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := GenS1{F0: -126}
                _ = fi9
                _ = v10
            }
        case -48:
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 uint = (v4 * 0)
                switch fi11 {
                case 9223372036854775807:
                    switch fi11 {
                    case 1:
                        var v13 int16 = 1000
                        _ = v13
                    case 0:
                        v14 := v1[v7]
                        _ = v14
                    default:
                        v7 = "hello"
                    }
                case -2147483648:
                    v12 += v12
                case 68:
                    switch fi11 {
                    case 128:
                        for fi15 := 0; fi15 < 15; fi15++ {
                            v3 = v5
                            var v16 int64 = -9223372036854775808
                            _ = fi15
                            _ = v16
                        }
                    }
                default:
                    switch (v7 + "foo bar") {
                    case "":
                        v7 = v7
                    default:
                        for fi17 := 0; fi17 < 7; fi17++ {
                            v18 := GenS1{F0: 42}
                            _ = fi17
                            _ = v18
                        }
                    }
                }
                var v19 bool = (!(!v3))
                _ = fi11
                _ = v12
                _ = v19
            }
        default:
            v7 = ("\n\t" + "hello")
        }
    case "\n\t":
        switch ("hello" + ((v7 + ((v7 + "a") + ("" + v7))) + (v7 + (v7 + "a")))) {
        case "foo bar":
            var v20 int8 = 42
            _ = v20
        case "":
            fmt.Println(v3)
        default:
            var v21 int8 = -1
            _ = v21
        }
    default:
        fmt.Println(v2, v4, v6)
    }
    v7 = v7
    v22 := GenS1{F0: -91}
    fmt.Println(v6)
    if (v2 && (!v2)) {
        fmt.Println(v7, v7)
        delete(v1, v7)
    }
    if true {
        v23 := []int{849, -603, -157}
        v2 = (-238 >= -665)
        _ = len(v23)
        _ = v23
    }
    switch (v7 + v7) {
    case "a":
        for fi24 := 0; fi24 < 4; fi24++ {
            v5 = (v5 && v3)
            switch v7 {
            case "":
                v7 = v7
            case "a":
                var v25 uint = v4
                _ = v25
            }
            _ = fi24
        }
    }
    v7 = v7
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 uint = v4
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 int16 = 0
            v30 := []int{int(v27), 0, (fi26 | (^(fi28)))}
            v5 = (v4 > 18446744073709551615)
            _ = fi28
            _ = v29
            _ = v30
        }
        v4 = (v4 / (((v4 << ((v4) & 63))) | 1))
        _ = fi26
        _ = v27
    }
    if v3 {
        v31 := []int{32767}
        for fi32 := 0; fi32 < 2; fi32++ {
            var v33 uint16 = 2
            var v34 uint = v4
            switch ((v7 + v7) + v7) {
            case "":
                fmt.Println(v33, v7, fi32)
            case "\n\t":
                switch "\n\t" {
                case "\n\t":
                    if (65535 != 2) {
                        fmt.Println(v3, v2)
                        fmt.Println(fi32, v6)
                        for fi35 := 0; fi35 < 12; fi35++ {
                            v36 := GenS1{F0: -128}
                            var v37 int64 = 9223372036854775807
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                    }
                case "foo bar":
                    v31 = append(v31, fi32)
                }
            case "foo bar":
                if v3 {
                    _ = len(v1)
                    if len(v31) > 0 {
                        v31[0] = fi32
                    }
                    v5 = v2
                } else {
                    for fi38 := 0; fi38 < 1; fi38++ {
                        v31 = append(v31, fi38)
                        _ = fi38
                    }
                }
            default:
                var v39 uint = v4
                _ = v39
            }
            _ = fi32
            _ = v33
            _ = v34
        }
        v40 := GenS1{F0: -128}
        _ = v31
        _ = v40
    } else {
        for fi41 := 0; fi41 < 14; fi41++ {
            fmt.Println(v5, v5, fi41)
            _ = fi41
        }
        var v42 uint = (v4 ^ (v4 >> ((v4) & 63)))
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
