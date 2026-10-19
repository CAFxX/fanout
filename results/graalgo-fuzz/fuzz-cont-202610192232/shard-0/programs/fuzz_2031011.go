package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int32 }

func main() {
    var v1 int64 = -1
    v2 := []string{"hello"}
    v3 := []string{"", "a", ("" + (("\n\t" + ("\n\t" + ("" + "a"))) + (("\n\t" + "foo bar") + "foo bar")))}
    var v4 int8 = 1
    var v5 int16 = 1000
    var v6 string = ("foo bar" + ("foo bar" + ""))
    var v7 int64 = v1
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: 65535}
        var v10 int = fi8
        _ = fi8
        _ = v10
        _ = v9
    }
    if (!(127 <= 32767)) {
        if false {
            v11 := GenS2{F0: 42, F1: 0, F2: 2147483647}
            v5 = 295
            _ = len(v2)
            _ = v11
        }
    } else {
        var v12 uint64 = 42
        _ = v12
    }
    switch 1 {
    case -844:
        if (!(false || true)) {
            v13 := GenS1{F0: 318}
            if len(v2) > 0 {
                v2[0] = v6
            }
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 int32 = -1
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 float32 = 3.14159
                    var v18 uint = 1
                    v19 := GenS1{F0: 1}
                    _ = len(v3)
                    var v20 uint16 = 1
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                v5 = (v5 & (v5 << ((65535) & 63)))
                _ = fi14
                _ = v15
            }
            _ = v13
        } else {
            _ = len(v3)
            switch v6 {
            case "a":
                switch 1 {
                case 2147483647:
                    for fi21 := 0; fi21 < 20; fi21++ {
                        if false {
                            v22 := GenS1{F0: 1}
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 9; fi23++ {
                            if ((fi21 >> ((531) & 63)) == -193) {
                                v1 = v1
                            } else {
                                var v24 int = int(v7)
                                var v25 uint = 0
                                _ = v24
                                _ = v25
                            }
                            var v26 int64 = v1
                            _ = fi23
                            _ = v26
                        }
                        var v27 uint16 = 488
                        _ = fi21
                        _ = v27
                    }
                default:
                    v28 := GenS2{F0: 65535, F1: (v5 + 42), F2: 42}
                    _ = v28
                }
            default:
                v6 = v6
            }
        }
    default:
        var v29 int64 = 449
        _ = v29
    }
    fmt.Println(v7, v7, v7)
    switch -2147483648 {
    case -9223372036854775808:
        fmt.Println(v6, v1, v1)
    case 4294967295:
        if (!true) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v30 int = -9223372036854775808
            v31 := GenS3{F0: 127, F1: -124}
            _ = v30
            _ = v31
        }
    }
    switch "\n\t" {
    case "\n\t":
        v6 = "\n\t"
    case "foo bar":
        var v32 uint64 = 4294967295
        _ = v32
    default:
        if true {
            var v33 uint = 1
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            _ = v33
        } else {
            v34 := GenS3{F0: 2, F1: -429}
            if true {
                v4 = 42
            } else {
                v3 = append(v3, v6)
                var v35 int64 = (v7 >> ((42) & 63))
                v7 = (v35 ^ 549)
                _ = v35
            }
            _ = v34
        }
    }
    _ = len(v3)
    fmt.Println(v5, v1)
    v36 := GenS1{F0: 136}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
