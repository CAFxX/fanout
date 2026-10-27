package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int64 }

func main() {
    var v1 uint8 = 2
    var v2 int16 = 1000
    var v3 int32 = -884
    v4 := make(map[string]uint32)
    var v5 uint16 = 274
    var v6 int8 = 42
    v6 += 53
    v1 = (v1 ^ v1)
    switch "hello" {
    case "a":
        switch "\n\t" {
        case "\n\t":
            delete(v4, (("foo bar" + (((("hello" + ("foo bar" + "")) + "a") + ("a" + "a")) + (("" + "") + ("" + "\n\t")))) + "hello"))
        }
    case "hello":
        v7 := GenS1{F0: v5, F1: 42}
        _ = v7
    case "\n\t":
        var v8 int32 = v3
        _ = v8
    default:
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 bool = ((810 >= 167) || true)
            if false {
                var v11 float64 = -2.25
                var v12 int32 = 294
                _ = v11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
    }
    fmt.Println(v5, v1, v6)
    var v13 string = ((("foo bar" + (("foo bar" + "foo bar") + "hello")) + "hello") + "foo bar")
    v14 := v4[((v13 + v13) + v13)]
    var v15 float32 = -2.25
    if (18446744073709551615 == 18446744073709551615) {
        var v16 uint32 = (v14 - v14)
        var v17 bool = (2147483647 != 1)
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        v1 -= 128
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS2{F0: int(v3), F1: 0, F2: 975}
            if ((("" + "foo bar") < (((v13 + v13) + v13) + v13)) && (true && (int64(v3) < -332))) {
                if (((v15 + 1.5) - (v15 * (((float32(0.1) + 0.1) / ((((0.0 + (3.14159 * (v15 * float32(3.14159)))) - v15) + 3.14159) + ((v15 - 3.14159) * 0.0))) + float32(1.5)))) >= float32(1e10)) {
                    if (true || true) {
                        v21 := GenS1{F0: 32767, F1: 65535}
                        v22 := GenS2{F0: (-(fi19)), F1: (v1 << ((uint(v5)) & 63)), F2: 4294967295}
                        var v23 int32 = (v3 % (((v3 - v3)) | 1))
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        fmt.Println(v3, v13, v6)
                        var v24 uint = 4294967295
                        v25 := GenS3{F0: (fi18 % ((fi19) | 1))}
                        _ = v24
                        _ = v25
                    }
                }
                switch (fi19 >> ((0) & 63)) {
                case 1:
                    v26 := GenS4{F0: -737, F1: v13}
                    _ = v26
                default:
                    v1 = v1
                }
                v2 = (v2 % ((v2) | 1))
            }
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                v6 = v6
            case "a":
                fmt.Println(fi18, fi19, v15)
            case "hello":
                var v27 bool = false
                _ = v27
            default:
                var v28 int = -9223372036854775808
                _ = v28
            }
            switch (v13 + v13) {
            case "foo bar":
                if false {
                    switch fi18 {
                    case 9223372036854775807:
                        var v29 int8 = (v6 ^ (v6 / ((100) | 1)))
                        _ = v29
                    case 42:
                        v30 := GenS2{F0: (fi19 / ((fi19) | 1)), F1: 0, F2: 798}
                        var v31 float32 = v15
                        _ = v30
                        _ = v31
                    }
                    var v32 float64 = 1.5
                    if (((v14 << ((1) & 63)) != 1) || true) {
                        if (-2147483648 > fi18) {
                            fmt.Println(v5, v3, fi18)
                        } else {
                            v15 = v15
                        }
                        v33 := GenS1{F0: 0, F1: v14}
                        _ = v33
                    }
                    _ = v32
                } else {
                    v34 := GenS4{F0: v3, F1: ((("a" + (v13 + (v13 + ((v13 + v13) + "a")))) + "hello") + "a")}
                    _ = v34
                }
            }
            _ = fi19
            _ = v20
        }
        v2 = int16(v6)
        _ = fi18
    }
    v35 := GenS5{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
