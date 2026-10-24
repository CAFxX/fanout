package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int16; F2 string }

func main() {
    var v1 uint32 = 2147483648
    var v2 int8 = -1
    v3 := []string{"\n\t", "\n\t", ""}
    var v4 uint32 = (v1 - v1)
    if (true || true) {
        if true {
            var v5 float32 = (1.5 / (1e10 + (-2.25 * 0.1)))
            var v6 uint16 = 401
            _ = v5
            _ = v6
        }
    } else {
        switch 1 {
        case -1:
            for fi7 := 0; fi7 < 11; fi7++ {
                fmt.Println(v1, v4)
                _ = fi7
            }
        case 2147483647:
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        case -291:
            v4 -= v1
        default:
            for fi8 := 0; fi8 < 10; fi8++ {
                v1 += 4294967295
                v1 = (v4 * (v1 >> ((154) & 63)))
                _ = fi8
            }
        }
    }
    var v9 uint = 65535
    if (false || true) {
        switch 127 {
        case 2147483647:
            var v10 float32 = (1e10 - 0.0)
            _ = v10
        case 9223372036854775807:
            if ((v1 >> ((v9) & 63)) != (v4 * (v1 & 2147483647))) {
                var v11 int16 = 42
                _ = v11
            }
        case 32768:
            var v12 uint8 = 0
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS1{F0: 0}
        _ = fi13
        _ = v14
    }
    if (v9 <= v9) {
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 uint64 = 552
            var v17 int64 = 9223372036854775807
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 string = "hello"
        _ = v18
    } else {
        v19 := GenS1{F0: 487}
        var v20 int8 = 42
        v4 = v1
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: 0}
    var v22 int = 497
    var v23 string = ("foo bar" + ("\n\t" + ""))
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 int16 = 32767
        if (18446744073709551615 < 4294967295) {
            v1 -= uint32(v9)
            fmt.Println(v4, v22, v4)
        }
        if (false && (!(true && true))) {
            v26 := GenS1{F0: 18446744073709551615}
            v27 := GenS2{F0: -2147483648}
            _ = v26
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v2)
    v28 := GenS1{F0: 18446744073709551615}
    v29 := GenS1{F0: uint64(v1)}
    for fi30 := 0; fi30 < 19; fi30++ {
        fmt.Println(fi30, v22)
        _ = fi30
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 int8 = v2
        fmt.Println(v9, v4, v22)
        _ = fi31
        _ = v32
    }
    var v33 int64 = -1
    for fi34 := 0; fi34 < 4; fi34++ {
        v35 := GenS1{F0: 800}
        v4 = 4294967295
        if true {
            if true {
                var v36 uint64 = 0
                for fi37 := 0; fi37 < 13; fi37++ {
                    var v38 uint64 = v36
                    _ = fi37
                    _ = v38
                }
                for fi39 := 0; fi39 < 10; fi39++ {
                    _ = len(v3)
                    _ = fi39
                }
                _ = v36
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v22 += (fi34 + (fi34 >> ((924) & 63)))
        } else {
            v40 := GenS3{F0: 1, F1: 32767, F2: (((v23 + ((v23 + "hello") + "a")) + (((v23 + ("a" + v23)) + v23) + v23)) + "foo bar")}
            for fi41 := 0; fi41 < 14; fi41++ {
                _ = len(v3)
                switch v23 {
                case "a":
                    var v42 int64 = -547
                    _ = v42
                default:
                    fmt.Println(v2)
                }
                _ = fi41
            }
            var v43 float64 = (0.1 - -2.25)
            _ = v40
            _ = v43
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v9
}
