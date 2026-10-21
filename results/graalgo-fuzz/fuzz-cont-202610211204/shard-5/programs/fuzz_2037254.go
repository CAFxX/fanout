package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 0
    var v3 int8 = -128
    v4 := []uint32{541, 450}
    v5 := GenS1{F0: 4294967295}
    switch (((("foo bar" + "") + "a") + ("\n\t" + ("foo bar" + "hello"))) + ("hello" + "foo bar")) {
    case "hello":
        var v6 uint = 936
        _ = v6
    }
    var v7 int8 = v3
    v3 = -1
    if (!true) {
        v7 = v3
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := GenS1{F0: 4294967295}
            switch "hello" {
            case "":
                v7 = (v3 % ((int8(v2)) | 1))
            default:
                v10 := GenS1{F0: fi8}
                _ = v10
            }
            _ = fi8
            _ = v9
        }
    } else {
        if (false || false) {
            switch ("foo bar" + ("foo bar" + "")) {
            case "a":
                fmt.Println(v7)
            case "":
                if (129 != 42) {
                    switch ("\n\t" + "hello") {
                    case "hello":
                        var v11 uint32 = 65535
                        _ = v11
                    case "foo bar":
                        v3 -= (v7 & (v3 << ((65535) & 63)))
                    default:
                        for fi12 := 0; fi12 < 12; fi12++ {
                            fmt.Println(v2, v2)
                            var v13 uint8 = 255
                            _ = len(v4)
                            _ = fi12
                            _ = v13
                        }
                    }
                } else {
                    fmt.Println(v2, v7)
                    var v14 string = "foo bar"
                    _ = v14
                }
            case "foo bar":
                var v15 int16 = 1
                _ = v15
            default:
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 float64 = float64(((1e10 + 1.5) / -2.25))
                    var v18 float32 = float32(((v17 - 3.14159) - (((0.1 - v17) / 1.5) / (v17 + v17))))
                    switch -758 {
                    case 824:
                        v1["foo bar"] = (-(fi16))
                    case 42:
                        var v19 uint64 = 697
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
        } else {
            fmt.Println(v3)
            v7 = v3
            var v20 int = -2147483648
            _ = v20
        }
    }
    _ = len(v4)
    v21 := GenS1{F0: 128}
    var v22 float64 = ((1.5 - (1.5 + 1.5)) + float64(0.0))
    v23 := GenS2{F0: (v2 & -2147483648), F1: 441, F2: 107}
    if (!(!(true || false))) {
        var v24 uint64 = 1
        var v25 uint32 = 2147483648
        _ = v24
        _ = v25
    } else {
        v22 += v22
        v7 -= (v7 & (^(v7)))
    }
    v7 = v7
    switch -982 {
    case 32768:
        switch 128 {
        case -1:
            if false {
                var v26 int64 = -1
                _ = v26
            }
        case 2147483647:
            v2 += (v2 % (((v2 & v2)) | 1))
        case -199:
            if ((0.1 - 1.5) != 1.5) {
                for fi27 := 0; fi27 < 10; fi27++ {
                    var v28 int32 = (v2 | v2)
                    _ = fi27
                    _ = v28
                }
                v29 := GenS1{F0: 126}
                v30 := GenS1{F0: 32768}
                _ = v29
                _ = v30
            }
        }
    case 127:
        for fi31 := 0; fi31 < 14; fi31++ {
            v32 := GenS3{F0: 655}
            v33 := GenS3{F0: 2}
            _ = fi31
            _ = v32
            _ = v33
        }
    case -9223372036854775808:
        var v34 int16 = 1000
        _ = v34
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
