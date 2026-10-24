package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 bool }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint16 = 65535
    v2 := []uint32{4294967295, 0, 827}
    var v3 uint = 18446744073709551615
    v4 := []uint32{}
    var v5 uint32 = 1
    switch ("hello" + "") {
    case "\n\t":
        if true {
            v5 -= (v5 | v5)
            if ("\n\t" > "foo bar") {
                for fi6 := 0; fi6 < 19; fi6++ {
                    for fi7 := 0; fi7 < 16; fi7++ {
                        var v8 int32 = 1
                        var v9 int32 = (v8 % (((v8 + (v8 - int32(v1)))) | 1))
                        _ = len(v2)
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    if false {
                        v10 := GenS1{F0: (v5 / (((v5 & (v5 & v5))) | 1)), F1: 42, F2: false}
                        _ = v10
                    }
                    _ = fi6
                }
            } else {
                if true {
                    fmt.Println(v5, v5)
                    for fi11 := 0; fi11 < 12; fi11++ {
                        v3 = v3
                        fmt.Println(v1, v1)
                        _ = fi11
                    }
                }
                var v12 uint16 = v1
                _ = v12
            }
        } else {
            if true {
                if (70 <= 0) {
                    var v13 int32 = 2147483647
                    _ = v13
                }
                v3 -= 0
                v1 = v1
            } else {
                switch -284 {
                case 32767:
                    v14 := GenS1{F0: (v5 << (((v3 << ((v3) & 63))) & 63)), F1: 42, F2: false}
                    _ = v14
                }
                for fi15 := 0; fi15 < 11; fi15++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi16 := 0; fi16 < 10; fi16++ {
                        _ = len(v4)
                        v1 -= (v1 - v1)
                        v2 = append(v2, v5)
                        _ = fi16
                    }
                    _ = fi15
                }
                v17 := GenS1{F0: 546, F1: 128, F2: true}
                _ = v17
            }
        }
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := GenS1{F0: (v5 & v5), F1: uint8(v1), F2: ("hello" == "\n\t")}
        v1 += 65535
        _ = fi18
        _ = v19
    }
    switch 2147483647 {
    case 127:
        var v20 uint16 = v1
        _ = v20
    case -1:
        var v21 uint32 = v5
        _ = v21
    case -2147483648:
        _ = len(v4)
    default:
        v1 += (v1 * v1)
    }
    v2 = append(v2, (v5 * v5))
    if (32767 > 0) {
        switch 32767 {
        case -9223372036854775808:
            var v22 uint16 = v1
            _ = v22
        case -225:
            if (!false) {
                var v23 float32 = 1.5
                _ = v23
            }
        case 1:
            fmt.Println(v3)
        default:
            if false {
                v1 = v1
                fmt.Println(v3)
                v5 = v5
            }
        }
        var v24 uint16 = v1
        v25 := GenS1{F0: 0, F1: 42, F2: (!(-2147483648 == 0))}
        _ = v24
        _ = v25
    } else {
        var v26 int = 4294967295
        var v27 uint8 = 1
        var v28 uint = (v3 - v3)
        _ = v26
        _ = v27
        _ = v28
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        for fi29 := 0; fi29 < 12; fi29++ {
            if (false && true) {
                for fi30 := 0; fi30 < 3; fi30++ {
                    var v31 float32 = 3.14159
                    var v32 string = "foo bar"
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
            v1 = (^(v1))
            var v33 int16 = -1
            _ = fi29
            _ = v33
        }
        fmt.Println(v5)
        v1 += (v1 << ((18446744073709551615) & 63))
    }
    v34 := GenS1{F0: v5, F1: 1, F2: false}
    v1 += 42
    var v35 bool = false
    var v36 uint = 603
    for fi37 := 0; fi37 < 19; fi37++ {
        if false {
            for fi38 := 0; fi38 < 1; fi38++ {
                fmt.Println(fi38, v36, fi37)
                v35 = (((("" + "a") + "a") <= ("foo bar" + ("" + "hello"))) || (!true))
                _ = fi38
            }
        } else {
            switch "" {
            case "":
                fmt.Println(v35, v35, v1)
            case "foo bar":
                _ = len(v4)
            case "a":
                for fi39 := 0; fi39 < 14; fi39++ {
                    var v40 int16 = 1
                    v35 = true
                    _ = fi39
                    _ = v40
                }
            }
        }
        v36 -= uint(v5)
        var v41 string = "hello"
        _ = fi37
        _ = v41
    }
    var v42 uint8 = 1
    fmt.Println(v35)
    if ((v5 | 0) != v5) {
        v35 = v35
    }
    v43 := GenS2{F0: 42}
    var v44 float64 = 3.14159
    v44 += 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
