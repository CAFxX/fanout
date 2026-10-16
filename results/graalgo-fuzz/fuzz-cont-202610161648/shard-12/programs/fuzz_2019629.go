package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() uint {
    var v4 int = 42
    var v5 float32 = 0.0
    if true {
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 uint64 = 4294967295
            var v8 float32 = (v5 * (v5 / ((1.5 + v5) - (v5 / ((v5 - 0.0) - float32(0.0))))))
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 float32 = 1.5
        v5 = -2.25
        _ = v9
    } else {
        var v10 string = ""
        var v11 bool = true
        v12 := []int{(^(v4)), -477}
        _ = v10
        _ = v11
        _ = v12
    }
    _ = v4
    _ = v5
    return 148
}
func genFn1(p0 int16, p1 int64, p2 uint64) uint8 {
    var v13 int8 = 42
    switch 127 {
    case -400:
        v14 := []int{278, -354, 0}
        _ = v14
    case -1:
        if ((65535 > uint32(p2)) || false) {
            v15 := GenS1{F0: 42}
            for fi16 := 0; fi16 < 3; fi16++ {
                if ((p2 / ((42) | 1)) > 4294967295) {
                    p1 = p1
                }
                switch "" {
                case "a":
                    for fi17 := 0; fi17 < 16; fi17++ {
                        var v18 uint32 = 649
                        switch "hello" {
                        case "a":
                            fmt.Println(fi16, p0)
                            if ((fi17 << ((18446744073709551615) & 63)) == (fi16 | fi16)) {
                                fmt.Println(fi17)
                            } else {
                                v19 := []int{(fi16 | 127)}
                                _ = v19
                            }
                            switch (fi17 & fi16) {
                            case 64:
                                var v20 uint = 525
                                _ = v20
                            case 537:
                                var v21 float64 = 0.0
                                _ = v21
                            default:
                                var v22 uint32 = (v18 % (((v18 + v18)) | 1))
                                _ = v22
                            }
                        }
                        _ = fi17
                        _ = v18
                    }
                }
                _ = fi16
            }
            _ = v15
        }
    case 9223372036854775807:
        switch "foo bar" {
        case "foo bar":
            v23 := []int{int(p0), 127, -2147483648}
            _ = v23
        case "\n\t":
            v24 := []int{-1, -9223372036854775808, -403, -625}
            _ = v24
        }
    }
    var v25 float32 = float32(0.0)
    p0 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v25
    return 164
}

func main() {
    var v1 int16 = 1
    var v2 float64 = -2.25
    var v3 float32 = 0.0
    var v26 int32 = 1
    v26 += 1
    v1 = v1
    var v27 int16 = 1
    var v28 uint8 = 42
    v29 := []int{int(v27)}
    var v30 string = (((("\n\t" + "a") + "foo bar") + "a") + "a")
    var v31 float32 = (((v3 / float32(1e10)) + v3) * 0.0)
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
