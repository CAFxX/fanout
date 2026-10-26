package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int; F1 string; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 int32; F2 int64 }
type GenS5 struct { F0 string; F1 int; F2 bool }
func genFn0() uint32 {
    var v5 int16 = -955
    v5 -= v5
    _ = v5
    return 65535
}

func main() {
    var v1 uint32 = 0
    var v2 uint16 = 65535
    var v3 bool = (1.5 < 1.5)
    var v4 uint8 = 128
    v6 := []int{127, 2147483647, 429}
    v7 := GenS1{F0: 2147483647, F1: (^(v4)), F2: 1}
    v8 := GenS1{F0: 2147483647, F1: v4, F2: 2}
    if (false && true) {
        v3 = (42 > 32768)
        var v9 float32 = 3.14159
        var v10 uint = 18446744073709551615
        _ = v10
        _ = v9
    } else {
        var v11 uint8 = 255
        switch ("foo bar" + "a") {
        case "":
            var v12 int64 = 42
            _ = v12
        default:
            v2 = (^(v2))
        }
        if v3 {
            v13 := GenS1{F0: -710, F1: v11, F2: 0}
            for fi14 := 0; fi14 < 2; fi14++ {
                fmt.Println(v11, fi14, v4)
                if (-888 != -731) {
                    if (false && (!(v3 || ((v1 << ((255) & 63)) == uint32(v4))))) {
                        fmt.Println(fi14)
                    }
                }
                var v15 uint = 58
                _ = fi14
                _ = v15
            }
            _ = v13
        }
        _ = v11
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        switch (fi16 & (fi16 ^ 42)) {
        case 299:
            v17 := GenS2{F0: fi16, F1: "hello", F2: v2}
            _ = v17
        case 127:
            v4 = (v4 & (v4 + (v4 * (v4 << ((18446744073709551615) & 63)))))
        default:
            var v18 int64 = 1
            _ = v18
        }
        v3 = (uint8(v1) <= 95)
        _ = fi16
    }
    if false {
        if (v4 != uint8(v2)) {
            if false {
                var v19 float64 = 0.0
                _ = v19
            }
        } else {
            fmt.Println(v4, v4, v3)
            v20 := GenS2{F0: -224, F1: (("" + ("" + "hello")) + "\n\t"), F2: (v2 % ((v2) | 1))}
            var v21 int8 = 127
            _ = v20
            _ = v21
        }
    } else {
        if v3 {
            for fi22 := 0; fi22 < 5; fi22++ {
                _ = len(v6)
                fmt.Println(v1, v4)
                _ = fi22
            }
            v23 := GenS1{F0: 42, F1: (v4 & 87), F2: uint64(v2)}
            var v24 int16 = 42
            _ = v23
            _ = v24
        } else {
            v1 = (v1 + genFn0())
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            switch fi25 {
            case -712:
                switch (("a" + "a") + "\n\t") {
                case "hello":
                    v26 := GenS1{F0: -2147483648, F1: 42, F2: 0}
                    _ = v26
                case "\n\t":
                    v27 := GenS3{F0: 42}
                    _ = v27
                case "":
                    if v3 {
                        var v28 int32 = 2147483647
                        fmt.Println(v3, v2)
                        _ = v28
                    }
                default:
                    var v29 float32 = 0.0
                    _ = v29
                }
            default:
                if ((fi25 % ((fi25) | 1)) < fi25) {
                    _ = len(v6)
                } else {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    for fi30 := 0; fi30 < 5; fi30++ {
                        var v31 uint16 = v2
                        v32 := GenS4{F0: (v1 >> ((18446744073709551615) & 63)), F1: -979, F2: 713}
                        v31 -= (v31 & 1)
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    v33 := GenS3{F0: -32768}
                    _ = v33
                }
            }
            var v34 bool = true
            _ = fi25
            _ = v34
        }
        if ("hello" >= "\n\t") {
            var v35 int = 4294967295
            fmt.Println(v35, v2, v1)
            v1 = v1
            _ = v35
        }
    }
    if false {
        v2 = (v2 >> ((1) & 63))
        if len(v6) > 0 {
            v6[0] = -336
        }
        v3 = v3
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v36 bool = true
    if false {
        if (0 < 2147483647) {
            var v37 bool = ((v1 ^ genFn0()) < (v1 ^ (v1 + 1)))
            for fi38 := 0; fi38 < 13; fi38++ {
                v6 = append(v6, fi38)
                _ = fi38
            }
            _ = v37
        } else {
            v4 += (v4 - 0)
            var v39 int32 = -1
            v1 += 0
            _ = v39
        }
    }
    if false {
        var v40 int8 = 127
        for fi41 := 0; fi41 < 14; fi41++ {
            if true {
                v42 := GenS2{F0: fi41, F1: ("foo bar" + ((("\n\t" + "foo bar") + "a") + ((("" + "\n\t") + "foo bar") + ""))), F2: v2}
                _ = v42
            }
            _ = fi41
        }
        v43 := GenS3{F0: -305}
        _ = v40
        _ = v43
    } else {
        if len(v6) > 0 {
            v6[0] = 0
        }
        v44 := GenS5{F0: ("foo bar" + "\n\t"), F1: 1, F2: true}
        _ = v44
    }
    if (false || (-2.25 == ((3.14159 * 1e10) + ((3.14159 + float32(float64(1e10))) + 1e10)))) {
        var v45 int32 = 1
        _ = v45
    } else {
        v3 = true
        switch 127 {
        case 1:
            if v3 {
                v1 += v1
            }
        default:
            v4 = 0
        }
        var v46 float64 = 0.0
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
