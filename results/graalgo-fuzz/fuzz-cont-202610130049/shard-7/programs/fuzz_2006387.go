package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int32; F1 int32; F2 bool }
type GenS4 struct { F0 uint; F1 int8 }

func main() {
    v1 := []string{"hello", "a", "a"}
    v2 := []string{("" + ((("\n\t" + "foo bar") + "") + (("\n\t" + "foo bar") + ("foo bar" + "hello")))), "a", "hello", ""}
    var v3 int8 = 0
    var v4 int8 = 1
    var v5 uint64 = 983
    var v6 int16 = 42
    for fi7 := 0; fi7 < 1; fi7++ {
        _ = len(v2)
        v5 -= (v5 & v5)
        _ = fi7
    }
    v3 -= v4
    switch -9223372036854775808 {
    case 32767:
        var v8 int16 = v6
        _ = v8
    case 2147483647:
        v2 = append(v2, "foo bar")
    default:
        v9 := GenS1{F0: 42, F1: 9223372036854775807}
        _ = v9
    }
    fmt.Println(v3)
    v6 += 1
    v10 := GenS2{F0: (v3 >> ((18446744073709551615) & 63)), F1: v6}
    for fi11 := 0; fi11 < 15; fi11++ {
        if (0.0 != 0.0) {
            v12 := GenS3{F0: 345, F1: 0, F2: (true && (!(true || true)))}
            fmt.Println(v4, fi11)
            var v13 float64 = -2.25
            _ = v12
            _ = v13
        } else {
            v3 += v4
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if len(v2) > 0 {
            v2[0] = "a"
        }
        _ = fi11
    }
    if (!(!true)) {
        fmt.Println(v5)
        v14 := GenS3{F0: int32(v4), F1: 2147483647, F2: ("\n\t" == "hello")}
        v15 := GenS1{F0: 743, F1: 9223372036854775807}
        _ = v14
        _ = v15
    }
    v16 := GenS4{F0: 2, F1: 42}
    for fi17 := 0; fi17 < 1; fi17++ {
        switch fi17 {
        case -311:
            v3 += v4
        case -83:
            var v18 float64 = (0.1 - -2.25)
            _ = v18
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            if false {
                switch (fi19 << ((0) & 63)) {
                case -2147483648:
                    var v20 uint64 = v5
                    _ = v20
                case 1:
                    v6 -= 1000
                case 127:
                    v5 = (v5 << ((2) & 63))
                default:
                    if false {
                        switch fi17 {
                        case 128:
                            var v21 float64 = 0.1
                            _ = v21
                        case 32768:
                            var v22 int16 = (^(v6))
                            _ = v22
                        }
                        v4 = (v4 / ((v3) | 1))
                        v23 := GenS2{F0: (v4 + v3), F1: v6}
                        _ = v23
                    }
                }
                var v24 uint = 2
                _ = v24
            } else {
                fmt.Println(v4)
            }
            _ = fi19
        }
        _ = fi17
    }
    var v25 int16 = (^(v6))
    v6 -= -1
    v6 += v6
    var v26 float64 = 1.5
    var v27 float32 = 1.5
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 uint64 = (v5 % (((v5 + v5)) | 1))
        _ = fi28
        _ = v29
    }
    var v30 uint16 = 15
    var v31 int32 = -642
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
