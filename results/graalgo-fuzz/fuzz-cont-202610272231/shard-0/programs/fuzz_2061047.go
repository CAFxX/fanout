package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int }

func main() {
    var v1 bool = (1 > 845)
    v2 := make(map[uint32]string)
    v3 := []uint32{42}
    var v4 bool = ((v1 && v1) || (v1 || false))
    v5 := GenS1{F0: 42, F1: -758, F2: 4294967295}
    var v6 int32 = 42
    var v7 int = -9223372036854775808
    var v8 string = ("\n\t" + ("" + "foo bar"))
    if (0.0 >= 3.14159) {
        for fi9 := 0; fi9 < 2; fi9++ {
            v4 = false
            delete(v2, 4294967295)
            _ = fi9
        }
        var v10 uint64 = 52
        _ = v10
    }
    var v11 int = (v7 ^ (v7 ^ v7))
    if (0 <= 127) {
        var v12 int = v7
        v13 := v2[2]
        _ = v12
        _ = v13
    }
    var v14 float32 = -2.25
    v15 := GenS1{F0: (v7 % ((37) | 1)), F1: 42, F2: (v7 * -9223372036854775808)}
    for fi16 := 0; fi16 < 8; fi16++ {
        if (v1 && v1) {
            var v17 uint32 = 297
            var v18 int64 = -1
            _ = v17
            _ = v18
        } else {
            v19 := GenS1{F0: v11, F1: -544, F2: fi16}
            var v20 int32 = v6
            _ = v19
            _ = v20
        }
        _ = fi16
    }
    v21 := GenS1{F0: (v7 / ((v7) | 1)), F1: 42, F2: (v11 % ((4294967295) | 1))}
    v22 := GenS1{F0: 128, F1: 993, F2: -2147483648}
    switch ((v8 + v8) + ((v8 + v8) + (v8 + "a"))) {
    case "foo bar":
        switch v8 {
        case "":
            v6 = v6
        case "\n\t":
            v11 -= (v11 % (((v11 << ((65535) & 63))) | 1))
        case "a":
            switch (v7 * v7) {
            case 32768:
                v1 = ((!true) && ((1 >= 141) || (-1 <= -29)))
            default:
                switch v8 {
                case "foo bar":
                    var v23 uint32 = 4294967295
                    _ = v23
                case "":
                    var v24 int64 = 106
                    _ = v24
                }
            }
        default:
            var v25 uint64 = 42
            _ = v25
        }
    case "":
        if (702 > 409) {
            v11 -= (v11 << ((255) & 63))
        }
    case "\n\t":
        for fi26 := 0; fi26 < 7; fi26++ {
            v27 := GenS1{F0: v7, F1: 42, F2: -521}
            if (!(!v1)) {
                var v28 bool = v4
                var v29 int = 454
                if false {
                    _ = len(v2)
                    if (((!v28) || true) || (!(2 == 1))) {
                        if (-948 >= int16(v6)) {
                            v11 = v11
                        }
                        for fi30 := 0; fi30 < 8; fi30++ {
                            var v31 int16 = 1000
                            _ = fi30
                            _ = v31
                        }
                    } else {
                        if ((v8 == v8) && (!(20 >= 18446744073709551615))) {
                            for fi32 := 0; fi32 < 2; fi32++ {
                                v4 = (!(1 > 255))
                                _ = fi32
                            }
                        } else {
                            var v33 bool = (v1 || (65535 <= 42))
                            _ = v33
                        }
                    }
                    v34 := GenS1{F0: 0, F1: 639, F2: -2147483648}
                    _ = v34
                }
                _ = v28
                _ = v29
            }
            _ = fi26
            _ = v27
        }
    }
    if false {
        var v35 int = (v7 | -9223372036854775808)
        if ((v8 + "hello") == (((v8 + (v8 + ((v8 + v8) + "hello"))) + ((("" + v8) + "") + ((v8 + "a") + (v8 + "\n\t")))) + (("a" + "") + ((v8 + ("" + (v8 + ("hello" + v8)))) + "a")))) {
            v6 -= (-(v6))
            v36 := v2[2]
            _ = v36
        }
        _ = v35
    } else {
        v4 = (18446744073709551615 <= 18446744073709551615)
        fmt.Println(v11, v8)
    }
    var v37 uint8 = 42
    var v38 uint8 = (v37 % ((v37) | 1))
    var v39 bool = (18446744073709551615 <= 42)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
