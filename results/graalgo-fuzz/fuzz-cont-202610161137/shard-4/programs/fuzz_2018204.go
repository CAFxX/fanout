package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint16 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 bool; F1 int32; F2 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 uint64 = 4294967295
    v4 := make(map[string]string)
    if ((v3 & (v3 >> ((244) & 63))) >= 4294967295) {
        if (0 <= uint8(v3)) {
            fmt.Println(v3)
            var v5 uint = 42
            _ = v5
        } else {
            fmt.Println(v3)
        }
        var v6 int16 = -32768
        v7 := GenS1{F0: 32767, F1: (v6 & 42), F2: 2}
        _ = v6
        _ = v7
    }
    if (!(32767 < -32768)) {
        switch (("foo bar" + "a") + ("foo bar" + "\n\t")) {
        case "\n\t":
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 int32 = 42
                v10 := GenS1{F0: 313, F1: 42, F2: 255}
                _ = fi8
                _ = v10
                _ = v9
            }
        case "a":
            fmt.Println(v3)
        case "hello":
            v11 := GenS1{F0: 32767, F1: -1, F2: 2}
            _ = v11
        }
        var v12 uint16 = 410
        v12 = (^(v12))
        _ = v12
    }
    var v13 int = 128
    for fi14 := 0; fi14 < 1; fi14++ {
        if ((!((42 == 1) && (152 < 255))) || true) {
            v3 = 2
        } else {
            v15 := GenS1{F0: 0, F1: 42, F2: 567}
            _ = v15
        }
        _ = fi14
    }
    v3 = v3
    var v16 string = (("\n\t" + ("hello" + ((("\n\t" + "") + (("foo bar" + ("hello" + "a")) + "hello")) + ""))) + "foo bar")
    v13 = (-(v13))
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            v13 = -9223372036854775808
            v19 := GenS1{F0: 0, F1: -57, F2: 229}
            v16 = v16
            _ = fi18
            _ = v19
        }
        var v20 uint32 = 0
        switch v16 {
        case "hello":
            if true {
                if false {
                    v21 := GenS1{F0: 665, F1: 0, F2: 1}
                    v22 := v1[(v20 & v20)]
                    var v23 int32 = -617
                    _ = v21
                    _ = v22
                    _ = v23
                }
                fmt.Println(v13, v13, v13)
                v24 := GenS2{F0: "a", F1: (0.1 != 3.14159)}
                _ = v24
            } else {
                var v25 bool = ((!(2 < 42)) && true)
                _ = v25
            }
        case "":
            v26 := GenS1{F0: 995, F1: -153, F2: 0}
            _ = v26
        case "a":
            if false {
                v27 := v2[(((v16 + "a") + "a") + ("" + "hello"))]
                if ((0 < 0) && false) {
                    v16 = (v27 + v16)
                    v3 -= v3
                    switch v16 {
                    case "a":
                        for fi28 := 0; fi28 < 15; fi28++ {
                            v29 := GenS1{F0: -1, F1: -32768, F2: 424}
                            _ = fi28
                            _ = v29
                        }
                    case "hello":
                        if (!true) {
                            v30 := GenS2{F0: v27, F1: false}
                            if true {
                                var v31 int64 = 9223372036854775807
                                _ = v31
                            } else {
                                v32 := v2[((v16 + v27) + "foo bar")]
                                _ = v32
                            }
                            _ = v30
                        }
                    }
                } else {
                    switch "foo bar" {
                    case "":
                        v13 = (^(fi17))
                    case "foo bar":
                        v33 := GenS1{F0: -1, F1: 0, F2: 1}
                        _ = v33
                    case "\n\t":
                        var v34 bool = (18446744073709551615 <= 0)
                        _ = v34
                    default:
                        _ = len(v2)
                    }
                }
                _ = v27
            } else {
                v3 -= (v3 * v3)
            }
        }
        _ = fi17
        _ = v20
    }
    v35 := []int{(v13 ^ (v13 >> ((897) & 63))), (-(v13)), (v13 - v13)}
    v3 += (v3 / (((v3 & v3)) | 1))
    v16 = ("foo bar" + v16)
    if ((2 >= 42) || (42 != -79)) {
        var v36 uint32 = 65535
        v16 = v16
        _ = v36
    }
    v37 := GenS3{F0: (!(-2147483648 > 0)), F1: 2147483647, F2: -853}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
}
