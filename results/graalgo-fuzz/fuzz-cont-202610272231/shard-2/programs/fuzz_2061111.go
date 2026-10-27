package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint64; F1 int; F2 uint32 }
type GenS6 struct { F0 int8 }
func genFn0() bool {
    if false {
        v5 := GenS1{F0: 2147483647}
        var v6 float64 = (1e10 + 1e10)
        if false {
            var v7 float32 = float32(v6)
            var v8 int16 = 594
            _ = v7
            _ = v8
        }
        _ = v5
        _ = v6
    }
    v9 := GenS2{F0: 2}
    var v10 bool = false
    _ = v10
    _ = v9
    return v10
}

func main() {
    var v1 int64 = 42
    v2 := []uint32{954, 0, 42, 65535}
    v3 := []uint32{4294967295, 28, 65535}
    var v4 int32 = 1
    if true {
        v11 := GenS3{F0: 2, F1: 117}
        var v12 string = ((("a" + ("a" + "foo bar")) + "hello") + "")
        v12 = ("\n\t" + v12)
        _ = v11
        _ = v12
    } else {
        v4 = (-(v4))
        var v13 int64 = v1
        var v14 uint16 = 42
        _ = v13
        _ = v14
    }
    v1 += (v1 ^ (v1 + v1))
    var v15 int32 = v4
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 16; fi17++ {
            v15 = (v15 & 2147483647)
            v18 := GenS4{F0: 2, F1: ((v4 <= 2147483647) || false), F2: 558}
            if len(v3) > 0 {
                v3[0] = 814
            }
            if true {
                switch "" {
                case "foo bar":
                    v1 -= -9223372036854775808
                case "\n\t":
                    fmt.Println(v15, fi17, v1)
                }
                v19 := GenS5{F0: 940, F1: (-(fi16)), F2: 2}
                var v20 uint32 = 0
                _ = v19
                _ = v20
            }
            v2 = append(v2, 42)
            v21 := GenS5{F0: 18446744073709551615, F1: 704, F2: 1}
            v15 -= (v4 / ((v15) | 1))
            _ = fi17
            _ = v18
            _ = v21
        }
        switch -2147483648 {
        case 0:
            v15 = 2147483647
        case 128:
            for fi22 := 0; fi22 < 5; fi22++ {
                v3 = append(v3, 189)
                v1 += v1
                _ = len(v3)
                _ = fi22
            }
        case 32768:
            var v23 uint = 0
            _ = v23
        }
        var v24 uint8 = 253
        _ = fi16
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        _ = len(v2)
        switch (fi25 % ((-1) | 1)) {
        case 42:
            var v26 int8 = 0
            _ = v26
        case 0:
            switch ("\n\t" + "hello") {
            case "":
                fmt.Println(v4, fi25)
            default:
                var v27 uint8 = 147
                _ = v27
            }
        case -684:
            if true {
                if len(v3) > 0 {
                    v3[0] = 0
                }
                var v28 uint32 = 42
                _ = v28
            } else {
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
            }
        }
        _ = fi25
    }
    v15 += (v4 << ((1) & 63))
    var v29 string = ("" + "")
    if (960 <= 255) {
        if (350 != 0) {
            v29 = v29
            if (-1 <= -2147483648) {
                for fi30 := 0; fi30 < 11; fi30++ {
                    v31 := GenS5{F0: 2, F1: (fi30 + (fi30 % ((-2147483648) | 1))), F2: 42}
                    fmt.Println(v1, v1)
                    _ = fi30
                    _ = v31
                }
            } else {
                var v32 float64 = 3.14159
                _ = v32
            }
            v4 = v4
        } else {
            var v33 uint64 = 18446744073709551615
            if (59 >= 1) {
                v29 = (v29 + "hello")
            }
            v1 = v1
            _ = v33
        }
        fmt.Println(v15, v1, v4)
        var v34 uint64 = 488
        _ = v34
    }
    v35 := GenS6{F0: -1}
    for fi36 := 0; fi36 < 14; fi36++ {
        if false {
            var v37 uint8 = 2
            var v38 int32 = 1
            _ = v37
            _ = v38
        }
        if (32767 <= -435) {
            fmt.Println(v1)
            if len(v3) > 0 {
                v3[0] = 811
            }
            var v39 uint16 = 0
            _ = v39
        }
        _ = fi36
    }
    switch 0 {
    case 0:
        switch v29 {
        case "\n\t":
            v1 += (v1 / (((v1 ^ -9223372036854775808)) | 1))
        }
    case 2147483647:
        v3 = append(v3, 193)
    default:
        v2 = append(v2, 42)
    }
    for fi40 := 0; fi40 < 20; fi40++ {
        v41 := GenS2{F0: 1}
        fmt.Println(fi40, v1, v15)
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v4
}
