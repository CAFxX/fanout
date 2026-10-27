package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 bool }
type GenS2 struct { F0 int32; F1 uint64 }

func main() {
    var v1 int8 = -128
    var v2 uint8 = 128
    var v3 uint32 = 2
    var v4 int32 = 1
    switch "a" {
    case "foo bar":
        var v5 uint32 = 2147483648
        _ = v5
    case "":
        if (false && (!false)) {
            if false {
                var v6 int8 = v1
                if (!(true || false)) {
                    fmt.Println(v2, v1, v4)
                }
                for fi7 := 0; fi7 < 2; fi7++ {
                    var v8 float32 = -2.25
                    fmt.Println(v6, v3)
                    _ = fi7
                    _ = v8
                }
                _ = v6
            }
        } else {
            var v9 int64 = 42
            fmt.Println(v1, v3, v3)
            _ = v9
        }
    case "hello":
        v10 := GenS1{F0: 18446744073709551615, F1: v2, F2: true}
        _ = v10
    }
    fmt.Println(v3, v3)
    v11 := []int{889, -948}
    var v12 uint = 255
    var v13 bool = false
    switch 233 {
    case -809:
        v14 := GenS1{F0: v12, F1: (v2 ^ (^(v2))), F2: ("\n\t" < "foo bar")}
        _ = v14
    case -2147483648:
        v13 = (((v2 == v2) || v13) || true)
    case 9223372036854775807:
        for fi15 := 0; fi15 < 1; fi15++ {
            for fi16 := 0; fi16 < 6; fi16++ {
                if v13 {
                    if (42 <= 420) {
                        switch "a" {
                        case "hello":
                            v17 := GenS1{F0: (v12 - v12), F1: v2, F2: v13}
                            _ = v17
                        }
                        v18 := GenS1{F0: uint(v3), F1: (v2 * 42), F2: ((255 <= 2) && false)}
                        v3 -= v3
                        _ = v18
                    } else {
                        v19 := GenS1{F0: v12, F1: 127, F2: v13}
                        v20 := GenS1{F0: (v12 + v12), F1: 79, F2: true}
                        _ = v19
                        _ = v20
                    }
                }
                fmt.Println(fi15, v3)
                var v21 bool = ((float32((0.1 + 1e10)) / (0.1 - 3.14159)) <= 1e10)
                var v22 int32 = 0
                _ = fi16
                _ = v21
                _ = v22
            }
            _ = fi15
        }
    default:
        if len(v11) > 0 {
            v11[0] = 127
        }
    }
    var v23 uint = 640
    v3 += uint32(v23)
    if true {
        if len(v11) > 0 {
            v11[0] = 1
        }
        for fi24 := 0; fi24 < 18; fi24++ {
            if v13 {
                v25 := GenS1{F0: 1, F1: 118, F2: v13}
                _ = v25
            } else {
                v26 := GenS1{F0: v23, F1: (v2 + 153), F2: v13}
                fmt.Println(v12, v3, v2)
                for fi27 := 0; fi27 < 1; fi27++ {
                    var v28 int64 = int64(v4)
                    var v29 int8 = -1
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = v26
            }
            _ = fi24
        }
    } else {
        switch -9223372036854775808 {
        case -445:
            if false {
                v1 -= v1
            } else {
                if len(v11) > 0 {
                    v11[0] = -1
                }
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            }
        }
        v11 = append(v11, 1)
    }
    v12 -= v23
    for fi30 := 0; fi30 < 3; fi30++ {
        fmt.Println(v23)
        var v31 uint16 = 52
        _ = fi30
        _ = v31
    }
    if ((v23 << (((v23 - v12)) & 63)) > v23) {
        switch ("a" + "\n\t") {
        case "hello":
            var v32 int64 = -278
            _ = v32
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        if len(v11) > 0 {
            v11[0] = 127
        }
    } else {
        var v33 uint16 = 775
        v34 := GenS2{F0: 119, F1: 1}
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
