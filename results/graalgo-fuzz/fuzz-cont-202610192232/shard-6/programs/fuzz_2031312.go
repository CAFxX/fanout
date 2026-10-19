package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint; F1 int64; F2 int32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 bool }
func genFn0(p0 string) uint64 {
    var v4 int64 = 322
    v5 := GenS1{F0: (v4 % ((v4) | 1)), F1: 0, F2: 1}
    var v6 float32 = 0.0
    v6 += float32(1e10)
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 string = (("\n\t" + ("hello" + "foo bar")) + "foo bar")
    var v2 uint = 0
    v3 := []uint32{65535, 65535, 508, 631}
    v2 = v2
    v7 := GenS1{F0: -863, F1: 262, F2: 636}
    var v8 uint32 = 2147483648
    v2 = v2
    for fi9 := 0; fi9 < 3; fi9++ {
        if false {
            var v10 uint16 = 32767
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 uint64 = 18446744073709551615
                if len(v3) > 0 {
                    v3[0] = v8
                }
                if false {
                    v12 += genFn0("\n\t")
                    if false {
                        if true {
                            v13 := GenS1{F0: 9223372036854775807, F1: 42, F2: -1}
                            var v14 uint8 = 42
                            _ = v13
                            _ = v14
                        }
                    }
                    switch (v1 + v1) {
                    case "\n\t":
                        if (!false) {
                            fmt.Println(v8)
                            fmt.Println(v2, v12)
                        } else {
                            if (!true) {
                                var v15 string = v1
                                v15 = v15
                                if (false || (!(fi9 != fi9))) {
                                    var v16 bool = false
                                    _ = v16
                                } else {
                                    v8 -= v8
                                    var v17 int8 = -1
                                    _ = v17
                                }
                                _ = v15
                            }
                            switch (fi11 - (fi9 & fi9)) {
                            case 0:
                                v8 = 2
                            case 32768:
                                switch v1 {
                                case "a":
                                    var v18 float32 = float32(-2.25)
                                    _ = v18
                                }
                            }
                            fmt.Println(v10)
                        }
                    case "hello":
                        switch v1 {
                        case "\n\t":
                            v1 = "a"
                        case "":
                            v12 = 74
                        }
                    }
                } else {
                    var v19 int64 = -919
                    _ = v19
                }
                _ = fi11
                _ = v12
            }
            var v20 int = fi9
            _ = v10
            _ = v20
        }
        _ = fi9
    }
    if true {
        var v21 int8 = -103
        _ = v21
    } else {
        var v22 int = -1
        v23 := GenS1{F0: -9223372036854775808, F1: 164, F2: -1}
        _ = v22
        _ = v23
    }
    var v24 float32 = (((1e10 + (3.14159 * 1.5)) / 0.1) * ((((3.14159 - 3.14159) + 1.5) / ((3.14159 - float32((3.14159 - 1e10))) / 1.5)) - 3.14159))
    if (-1 >= 32767) {
        switch ((v1 + "\n\t") + v1) {
        case "\n\t":
            v25 := GenS2{F0: 1, F1: 620, F2: -258}
            _ = v25
        case "a":
            var v26 int64 = 0
            _ = v26
        case "":
            v27 := GenS3{F0: 1000, F1: 42}
            _ = v27
        default:
            v28 := GenS4{F0: false}
            _ = v28
        }
        fmt.Println(v8, v2)
        v1 = (v1 + ("foo bar" + v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v7
    _ = v8
}
