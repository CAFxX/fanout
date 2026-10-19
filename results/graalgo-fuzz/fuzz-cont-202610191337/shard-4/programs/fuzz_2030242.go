package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int8; F1 uint64; F2 bool }

func main() {
    v1 := []string{"\n\t"}
    v2 := []uint32{2, 65535}
    var v3 uint8 = 218
    var v4 uint32 = 2
    var v5 float64 = 0.1
    if false {
        v6 := GenS1{F0: 703, F1: 0}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v4, v5)
        _ = v6
    }
    if (false || (false || (false || (v5 < (3.14159 * ((v5 + -2.25) - v5)))))) {
        switch ("" + "hello") {
        case "":
            v4 = 0
        case "\n\t":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    } else {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 float32 = float32(v5)
            fmt.Println(v3)
            var v9 uint = 65535
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    switch "" {
    case "hello":
        var v10 bool = true
        _ = v10
    case "\n\t":
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 bool = true
            var v13 int32 = -530
            v14 := GenS1{F0: (v13 * (v13 << ((506) & 63))), F1: 1}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    case "":
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS1{F0: 42, F1: int32(v3)}
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 int = 2147483647
                if len(v2) > 0 {
                    v2[0] = 981
                }
                v19 := GenS2{F0: 127, F1: 18446744073709551615, F2: false}
                _ = fi17
                _ = v18
                _ = v19
            }
            if true {
                var v20 float32 = ((1.5 + -2.25) * 0.0)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if false {
                    for fi21 := 0; fi21 < 13; fi21++ {
                        fmt.Println(v5, v3)
                        v22 := GenS2{F0: 100, F1: uint64(v4), F2: (false || true)}
                        _ = fi21
                        _ = v22
                    }
                }
                _ = v20
            } else {
                v5 = 1.5
            }
            _ = fi15
            _ = v16
        }
    }
    v4 = v4
    var v23 uint32 = (v4 >> ((42) & 63))
    fmt.Println(v23, v5)
    var v24 uint16 = 255
    for fi25 := 0; fi25 < 5; fi25++ {
        var v26 int64 = 0
        v5 += 0.0
        _ = fi25
        _ = v26
    }
    var v27 int = 169
    var v28 bool = (int8(v27) != 1)
    v1 = append(v1, (("foo bar" + ("" + "")) + "a"))
    v28 = v28
    v27 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
