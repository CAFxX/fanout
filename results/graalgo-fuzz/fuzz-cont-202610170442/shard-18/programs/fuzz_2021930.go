package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 bool; F1 bool; F2 int64 }
type GenS4 struct { F0 int64; F1 uint; F2 int16 }

func main() {
    var v1 int = 128
    var v2 bool = false
    var v3 int64 = -234
    var v4 uint64 = 2
    switch v1 {
    case -2147483648:
        switch (("hello" + "\n\t") + "a") {
        case "":
            fmt.Println(v4, v2)
        default:
            v4 = v4
        }
    }
    var v5 int32 = 42
    v6 := []int{v1}
    v7 := GenS1{F0: (v5 * v5), F1: 4294967295, F2: 2}
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint64 = 238
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS1{F0: (v5 - (v5 ^ 42)), F1: 447, F2: 1}
            if (v2 || v2) {
                switch "a" {
                case "a":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            }
            if true {
                v3 += (v3 / ((v3) | 1))
            }
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    if len(v6) > 0 {
        v6[0] = 32768
    }
    var v12 int = -9223372036854775808
    v2 = (-131 <= (v5 & 42))
    fmt.Println(v1, v12, v1)
    if (v5 < -98) {
        fmt.Println(v1, v5)
        v13 := GenS1{F0: (v5 - 1), F1: 2, F2: 315}
        _ = v13
    }
    v5 = v5
    if true {
        if (59 > 2) {
            v14 := GenS2{F0: 610, F1: 1}
            var v15 bool = (!true)
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 14; fi16++ {
            v12 -= fi16
            _ = fi16
        }
    }
    if len(v6) > 0 {
        v6[0] = (v1 & v12)
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        switch "" {
        case "a":
            if len(v6) > 0 {
                v6[0] = (v12 % ((int(v3)) | 1))
            }
        default:
            fmt.Println(v2)
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v18 := GenS1{F0: v5, F1: 4294967295, F2: 2}
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        switch -9223372036854775808 {
        case 2147483647:
            v5 = (-(v5))
        case -1:
            var v20 int32 = v5
            _ = v20
        case 42:
            var v21 bool = ("hello" != "")
            _ = v21
        default:
            v12 = v12
        }
        if true {
            fmt.Println(v1, v4, v4)
            v12 -= fi19
        } else {
            v22 := GenS3{F0: true, F1: (((0.1 / float64((float32(-2.25) - (3.14159 - 1.5)))) / ((1.5 * float64((-2.25 + 1.5))) + 3.14159)) > (float64((0.0 + 3.14159)) - 1e10)), F2: 0}
            fmt.Println(v3)
            _ = v22
        }
        _ = fi19
    }
    switch ("\n\t" + ("a" + "\n\t")) {
    case "hello":
        if ((1.5 * 0.1) >= 1.5) {
            if (!v2) {
                v23 := GenS2{F0: v4, F1: 821}
                _ = v23
            }
            for fi24 := 0; fi24 < 17; fi24++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v25 uint64 = (v4 + (v4 >> ((826) & 63)))
                for fi26 := 0; fi26 < 12; fi26++ {
                    if len(v6) > 0 {
                        v6[0] = fi26
                    }
                    _ = fi26
                }
                _ = fi24
                _ = v25
            }
        } else {
            v27 := GenS4{F0: (^(v3)), F1: 65535, F2: 87}
            _ = v27
        }
    case "foo bar":
        var v28 uint32 = 65535
        _ = v28
    case "":
        switch ("foo bar" + "hello") {
        case "a":
            var v29 uint16 = 378
            _ = v29
        }
    default:
        if len(v6) > 0 {
            v6[0] = (v1 / (((v1 & (v12 >> ((0) & 63)))) | 1))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
