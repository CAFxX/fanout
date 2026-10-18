package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint16) int32 {
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 += 42
        var v8 float64 = (3.14159 - 0.0)
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 2, F1: 65535, F2: 0}
    _ = p0
    _ = v9
    return -2147483648
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 0
    var v3 uint64 = 1
    var v4 int64 = 737
    v5 := []uint32{917}
    v6 := []string{"hello", ("foo bar" + ("a" + "")), "", ("" + ("hello" + "foo bar"))}
    for fi10 := 0; fi10 < 2; fi10++ {
        switch ("foo bar" + "\n\t") {
        case "hello":
            fmt.Println(fi10)
        }
        if (true || (0 == 0)) {
            delete(v1, fi10)
        } else {
            v1[(fi10 % ((fi10) | 1))] = ("" + ("foo bar" + (("" + (("foo bar" + (("foo bar" + "a") + "\n\t")) + "\n\t")) + "\n\t")))
        }
        _ = fi10
    }
    v11 := GenS1{F0: 255, F1: 2, F2: 4294967295}
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 uint64 = v3
        _ = fi12
        _ = v13
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS1{F0: 212, F1: 2147483647, F2: 316}
        switch (fi14 << ((1) & 63)) {
        case -1:
            switch 32768 {
            case -9223372036854775808:
                fmt.Println(v4)
            }
        case -201:
            if true {
                fmt.Println(v2)
                var v16 int64 = v4
                v17 := GenS2{F0: 1}
                _ = v16
                _ = v17
            }
        default:
            fmt.Println(v3, fi14, v2)
        }
        _ = fi14
        _ = v15
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        v6 = append(v6, "hello")
        var v19 int16 = 549
        _ = fi18
        _ = v19
    }
    v3 -= v3
    var v20 int16 = 1
    var v21 float32 = 0.1
    if (-9223372036854775808 > -2147483648) {
        v2 -= v2
        if ((!true) || false) {
            v22 := GenS2{F0: 92}
            _ = v22
        } else {
            fmt.Println(v4, v4, v20)
            v23 := GenS3{F0: v3, F1: 2147483647}
            var v24 bool = false
            _ = v23
            _ = v24
        }
    } else {
        v25 := GenS4{F0: ((!false) && ((("" < ("\n\t" + "\n\t")) || (!false)) && ("\n\t" <= (("" + (("foo bar" + "") + (("hello" + "foo bar") + ("\n\t" + "\n\t")))) + "\n\t"))))}
        if (v2 != (^(v2))) {
            v4 = 660
        }
        v20 += v20
        _ = v25
    }
    switch "" {
    case "":
        var v26 int16 = (v20 | (v20 << ((18446744073709551615) & 63)))
        _ = v26
    case "hello":
        switch -870 {
        case -650:
            switch 310 {
            case 42:
                v27 := GenS3{F0: (v3 - (v3 / (((v3 + v3)) | 1))), F1: genFn0((v2 << ((4294967295) & 63)))}
                _ = v27
            case 2147483647:
                fmt.Println(v3, v20, v20)
            case -359:
                fmt.Println(v2)
            default:
                if (-2147483648 == -742) {
                    v20 += v20
                }
            }
        case 32768:
            switch int(v20) {
            case -9223372036854775808:
                switch -1 {
                case 1:
                    _ = len(v5)
                case 148:
                    v2 += (v2 >> ((751) & 63))
                }
            case 1:
                v2 -= v2
            default:
                var v28 uint32 = 1
                _ = v28
            }
        default:
            if len(v5) > 0 {
                v5[0] = 65535
            }
        }
    case "\n\t":
        var v29 uint32 = 2147483648
        _ = v29
    }
    v30 := GenS5{F0: 2147483647}
    _ = len(v1)
    var v31 int8 = 88
    v21 -= -2.25
    v21 = v21
    var v32 int32 = 585
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
