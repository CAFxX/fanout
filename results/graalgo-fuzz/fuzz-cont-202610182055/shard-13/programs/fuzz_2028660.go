package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint8; F2 bool }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint32 }

func main() {
    v1 := []string{"\n\t", "\n\t", "", "hello"}
    v2 := []uint32{42, 65535, 2, 89}
    var v3 uint32 = 861
    var v4 float32 = 1.5
    var v5 uint8 = 128
    var v6 uint16 = 52
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(v5, v3, v3)
        for fi8 := 0; fi8 < 6; fi8++ {
            v4 = ((v4 * float32(-2.25)) / 1.0)
            _ = fi8
        }
        _ = fi7
    }
    v5 = (v5 << ((18446744073709551615) & 63))
    v1 = append(v1, ("a" + ("foo bar" + "\n\t")))
    v3 += (v3 % ((uint32(v6)) | 1))
    if true {
        if (-9223372036854775808 <= 435) {
            if ((("hello" + ("foo bar" + "")) + "") < "hello") {
                v6 = (v6 ^ 32767)
                v2 = append(v2, (v3 | 2147483648))
                v9 := GenS1{F0: int32(v5)}
                _ = v9
            }
            fmt.Println(v5)
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v10 int8 = -91
        var v11 int16 = -1
        _ = v10
        _ = v11
    } else {
        v2 = append(v2, (^(v3)))
    }
    var v12 int64 = 0
    v5 -= (v5 ^ 17)
    if (true || (1 != 32767)) {
        v5 = v5
        var v13 uint16 = (^(v6))
        fmt.Println(v12)
        _ = v13
    } else {
        v5 = (v5 >> ((65535) & 63))
        v14 := GenS2{F0: ((uint64(v3) <= 2) || (1e10 > 1e10)), F1: (v5 / (((v5 & v5)) | 1)), F2: (false || true)}
        v3 -= (v3 << ((1) & 63))
        _ = v14
    }
    var v15 uint16 = 32767
    v16 := GenS3{F0: v3, F1: 42}
    if false {
        v6 = (v15 >> ((255) & 63))
        fmt.Println(v6, v3)
    } else {
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 string = ("a" + "foo bar")
            var v19 uint16 = v15
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        v3 = (v3 - 827)
    }
    var v20 bool = (!(!false))
    fmt.Println(v20, v15, v12)
    if (v20 || true) {
        switch 4294967295 {
        case 4294967295:
            v4 += v4
        default:
            var v21 int = 32767
            _ = v21
        }
        v12 -= (v12 >> ((1) & 63))
        switch "a" {
        case "foo bar":
            if (127 < -459) {
                v22 := GenS4{F0: (v3 ^ v3)}
                if ((v12 | v12) > v12) {
                    v20 = v20
                    var v23 int64 = 0
                    if (!v20) {
                        if false {
                            v24 := GenS4{F0: 2147483648}
                            fmt.Println(v12)
                            _ = v24
                        }
                        v20 = false
                        for fi25 := 0; fi25 < 11; fi25++ {
                            var v26 bool = false
                            var v27 uint16 = (v6 % ((v15) | 1))
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                    _ = v23
                }
                if false {
                    fmt.Println(v5, v20)
                    v15 += v15
                    if len(v2) > 0 {
                        v2[0] = (^(v3))
                    }
                }
                _ = v22
            }
        case "hello":
            v6 = 420
        }
    }
    v2 = append(v2, v3)
    if true {
        _ = len(v2)
        v15 += v6
        fmt.Println(v6, v15)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
