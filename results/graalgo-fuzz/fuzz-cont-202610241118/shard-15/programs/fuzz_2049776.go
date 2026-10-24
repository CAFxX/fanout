package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{904, 65535, 4294967295}
    var v3 uint32 = 0
    var v4 uint32 = v3
    if ((float64((float32(0.1) * 1.5)) / 1.5) < 0.1) {
        v5 := GenS1{F0: (v3 ^ v3), F1: -1}
        if (1e10 != ((3.14159 * 1.5) + 0.0)) {
            v3 += (^(v3))
        }
        var v6 int16 = -32768
        _ = v5
        _ = v6
    }
    v7 := v1[v4]
    if true {
        if (962 == 42) {
            v4 -= 2
            switch ("hello" + "foo bar") {
            case "foo bar":
                v7 -= 2
            }
            v7 = v4
        }
        fmt.Println(v4)
    }
    var v8 int64 = -9223372036854775808
    for fi9 := 0; fi9 < 15; fi9++ {
        if ((true && (-1 != fi9)) || true) {
            var v10 uint = 801
            _ = v10
        } else {
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS1{F0: (v7 ^ (v7 >> ((42) & 63))), F1: v8}
                if len(v2) > 0 {
                    v2[0] = 204
                }
                v4 += (v3 - v4)
                _ = fi11
                _ = v12
            }
            var v13 uint64 = 1
            var v14 bool = (!(!(2 > 255)))
            _ = v13
            _ = v14
        }
        _ = fi9
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        if false {
            var v16 bool = (!("a" <= "foo bar"))
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 bool = true
                if len(v2) > 0 {
                    v2[0] = v7
                }
                _ = fi17
                _ = v18
            }
            if (42 == -516) {
                var v19 bool = true
                _ = v19
            } else {
                v16 = (v16 && (!v16))
                var v20 float64 = (0.1 / 1.0)
                v21 := v1[v3]
                _ = v20
                _ = v21
            }
            _ = v16
        } else {
            var v22 uint = 616
            if true {
                v23 := GenS2{F0: 741, F1: 255}
                v7 = v4
                _ = v23
            }
            _ = v22
        }
        _ = fi15
    }
    v24 := GenS1{F0: (^(v4)), F1: (v8 & (v8 - v8))}
    v25 := v1[v7]
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((true || (-2.25 == (3.14159 - ((0.0 + -2.25) - 0.1)))) || (0.1 <= 1.5)) {
        v4 -= v7
    } else {
        if (884 < 2) {
            v2 = append(v2, (v25 % (((v4 << ((841) & 63))) | 1)))
        } else {
            for fi26 := 0; fi26 < 17; fi26++ {
                v3 = 148
                v25 -= (^(v25))
                v27 := GenS1{F0: v4, F1: (v8 & (v8 / (((v8 << ((935) & 63))) | 1)))}
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 12; fi28++ {
                fmt.Println(v4)
                for fi29 := 0; fi29 < 5; fi29++ {
                    var v30 int64 = v8
                    v30 -= (v8 % ((v8) | 1))
                    _ = fi29
                    _ = v30
                }
                _ = fi28
            }
        }
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        for fi32 := 0; fi32 < 19; fi32++ {
            if ((42 < -2147483648) && false) {
                var v33 uint8 = 128
                var v34 int16 = 1000
                v1[(v4 % ((v4) | 1))] = 65535
                _ = v33
                _ = v34
            }
            var v35 bool = (0 > 2)
            _ = fi32
            _ = v35
        }
        switch 42 {
        case -739:
            if (!(true || false)) {
                v8 = 845
                v25 = v4
            } else {
                var v36 int16 = 32767
                _ = v36
            }
        case 127:
            v8 = (-(v8))
        case 42:
            var v37 uint16 = 54
            _ = v37
        }
        _ = fi31
    }
    var v38 uint16 = 0
    switch "hello" {
    case "\n\t":
        var v39 int8 = 1
        _ = v39
    }
    v40 := GenS1{F0: (v7 + (v25 | v7)), F1: (v8 + 1)}
    switch -1 {
    case 375:
        for fi41 := 0; fi41 < 4; fi41++ {
            var v42 uint64 = 18446744073709551615
            v42 = (^(v42))
            _ = fi41
            _ = v42
        }
    case 886:
        if len(v2) > 0 {
            v2[0] = v25
        }
    }
    switch "foo bar" {
    case "a":
        if (!(!(4294967295 >= int(v25)))) {
            fmt.Println(v25, v3, v7)
            var v43 int8 = 63
            var v44 float32 = 0.0
            _ = v43
            _ = v44
        }
    }
    v8 -= v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v7
    _ = v8
}
