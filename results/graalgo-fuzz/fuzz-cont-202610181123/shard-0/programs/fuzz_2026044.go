package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint8; F1 uint16; F2 int }

func main() {
    var v1 bool = false
    var v2 bool = v1
    var v3 int8 = 39
    v4 := []string{"a", "a"}
    var v5 uint32 = 65535
    v6 := make(map[string]bool)
    if false {
        var v7 bool = (v1 || true)
        _ = v7
    }
    switch ((("hello" + "a") + "\n\t") + "hello") {
    case "foo bar":
        delete(v6, "hello")
    }
    v2 = (312 <= 2)
    v5 = v5
    fmt.Println(v5, v3, v1)
    delete(v6, "foo bar")
    v1 = (18446744073709551615 > 42)
    switch 1 {
    case -281:
        _ = len(v4)
    }
    if (!true) {
        switch "" {
        case "foo bar":
            v8 := GenS1{F0: v3, F1: (("\n\t" + ("foo bar" + ("foo bar" + "a"))) + (("" + "") + "\n\t"))}
            _ = v8
        case "a":
            var v9 bool = (-9223372036854775808 >= -1)
            _ = v9
        case "hello":
            v10 := GenS2{F0: 0}
            _ = v10
        default:
            if ((0 < -1) || v2) {
                v5 -= v5
                var v11 uint = 255
                _ = v11
            } else {
                v12 := GenS2{F0: 18446744073709551615}
                v13 := GenS3{F0: -9223372036854775808, F1: 42}
                switch 116 {
                case 2147483647:
                    v14 := GenS3{F0: -353, F1: 0}
                    _ = v14
                case -9223372036854775808:
                    for fi15 := 0; fi15 < 16; fi15++ {
                        var v16 uint = 65535
                        v2 = false
                        var v17 uint64 = 42
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                case -883:
                    v18 := GenS2{F0: 42}
                    _ = v18
                default:
                    v5 += (v5 + (v5 / (((v5 % (((^(v5))) | 1))) | 1)))
                }
                _ = v12
                _ = v13
            }
        }
        var v19 uint64 = 2
        _ = v19
    } else {
        v5 = (v5 >> ((255) & 63))
        var v20 uint64 = 0
        fmt.Println(v2)
        _ = v20
    }
    v21 := GenS4{F0: 1, F1: 174, F2: 595}
    for fi22 := 0; fi22 < 15; fi22++ {
        for fi23 := 0; fi23 < 9; fi23++ {
            v5 += (^(v5))
            _ = fi23
        }
        v5 += (v5 % (((v5 * v5)) | 1))
        fmt.Println(fi22)
        _ = fi22
    }
    if v2 {
        v2 = v2
    }
    v1 = v2
    if ((!true) || ((!(-171 > -911)) || v2)) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v3, v3, v2)
        for fi24 := 0; fi24 < 2; fi24++ {
            for fi25 := 0; fi25 < 10; fi25++ {
                fmt.Println(fi25, fi24, v1)
                if (-1 <= 661) {
                    var v26 uint8 = 1
                    v27 := v6[""]
                    v27 = v2
                    _ = v26
                    _ = v27
                } else {
                    if (1 != 1) {
                        var v28 uint16 = 32767
                        _ = v28
                    }
                    var v29 uint32 = 0
                    _ = v29
                }
                var v30 bool = v1
                _ = fi25
                _ = v30
            }
            _ = fi24
        }
    }
    var v31 uint64 = 849
    v31 = 359
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
