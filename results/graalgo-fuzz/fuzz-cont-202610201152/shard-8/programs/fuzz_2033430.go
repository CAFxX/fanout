package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint64; F2 int16 }
func genFn0(p0 uint16) uint16 {
    v4 := GenS1{F0: -750, F1: 32767}
    v5 := GenS2{F0: ""}
    var v6 int8 = 69
    p0 += 6
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 255
}

func main() {
    var v1 bool = (44 >= 1)
    var v2 uint8 = 204
    v3 := []string{}
    if ((4294967295 >= 2147483648) && false) {
        var v7 string = "a"
        _ = v7
    } else {
        var v8 float32 = 1e10
        _ = v8
    }
    var v9 uint64 = 4294967295
    v9 = v9
    switch "\n\t" {
    case "hello":
        if (!false) {
            var v10 uint32 = 2147483647
            _ = v10
        }
    case "\n\t":
        switch "a" {
        case "foo bar":
            for fi11 := 0; fi11 < 15; fi11++ {
                if (42 >= 255) {
                    var v12 bool = v1
                    if ((float32(1.5) / 3.14159) == ((3.14159 * 1.5) - 0.0)) {
                        var v13 int = fi11
                        var v14 bool = (18446744073709551615 == 668)
                        _ = len(v3)
                        _ = v13
                        _ = v14
                    } else {
                        v1 = v12
                        for fi15 := 0; fi15 < 3; fi15++ {
                            if v12 {
                                v12 = (v2 != 127)
                                v16 := GenS1{F0: 42, F1: 835}
                                v1 = v12
                                _ = v16
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch "hello" {
                            case "hello":
                                v9 = uint64(v2)
                            case "":
                                v2 = v2
                            }
                            _ = fi15
                        }
                    }
                    v17 := GenS1{F0: -607, F1: 108}
                    _ = v12
                    _ = v17
                }
                var v18 uint8 = (^(v2))
                _ = fi11
                _ = v18
            }
        case "a":
            v1 = v1
        }
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 uint16 = 0
            _ = fi19
            _ = v20
        }
    }
    fmt.Println(v2, v2)
    v1 = (v1 && false)
    for fi21 := 0; fi21 < 1; fi21++ {
        fmt.Println(v2)
        v2 += v2
        _ = fi21
    }
    v22 := GenS3{F0: (v2 - (v2 % ((1) | 1))), F1: (v9 / (((v9 * v9)) | 1)), F2: 687}
    v1 = v1
    var v23 uint8 = v2
    var v24 uint64 = v9
    for fi25 := 0; fi25 < 18; fi25++ {
        fmt.Println(v1)
        var v26 int = (-(fi25))
        var v27 int64 = 628
        _ = fi25
        _ = v26
        _ = v27
    }
    switch (("a" + (("" + "foo bar") + "")) + "a") {
    case "hello":
        switch -321 {
        case 9223372036854775807:
            fmt.Println(v1)
        case -9223372036854775808:
            var v28 int16 = 1000
            _ = v28
        case 1:
            v29 := GenS2{F0: "a"}
            _ = v29
        }
    case "\n\t":
        fmt.Println(v2)
    default:
        for fi30 := 0; fi30 < 3; fi30++ {
            v9 += 403
            switch ("" + (("a" + "") + ("" + "hello"))) {
            case "a":
                if (("\n\t" + "\n\t") != "a") {
                    var v31 uint8 = (v23 / ((v2) | 1))
                    if v1 {
                        v32 := GenS3{F0: v23, F1: v24, F2: 32767}
                        v3 = append(v3, (("hello" + "a") + ("" + (((("\n\t" + "\n\t") + ("\n\t" + "foo bar")) + "") + ("foo bar" + "hello")))))
                        _ = v32
                    }
                    var v33 int64 = 1
                    _ = v31
                    _ = v33
                }
            case "hello":
                switch ("\n\t" + (("" + ("a" + ("a" + "\n\t"))) + "")) {
                case "hello":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    fmt.Println(v9)
                }
            }
            if (fi30 < fi30) {
                var v34 bool = v1
                var v35 uint64 = (v24 * v9)
                _ = v34
                _ = v35
            }
            _ = fi30
        }
    }
    var v36 int32 = 1
    switch "" {
    case "\n\t":
        v9 = v24
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "hello":
        fmt.Println(v1, v23)
    }
    fmt.Println(v2, v9, v1)
    v1 = (true && v1)
    v37 := GenS1{F0: (v36 & v36), F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v36
    _ = v37
    _ = v9
}
