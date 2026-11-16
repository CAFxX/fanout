package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }

func main() {
    var v1 uint32 = 291
    var v2 int64 = 1
    var v3 uint32 = (v1 - (^(v1)))
    var v4 int8 = 100
    var v5 uint8 = 0
    var v6 float32 = 1e10
    var v7 string = "\n\t"
    var v8 uint32 = v1
    if (!false) {
        var v9 bool = false
        _ = v9
    } else {
        var v10 uint8 = v5
        _ = v10
    }
    if (!(!true)) {
        v11 := GenS1{F0: false, F1: v2}
        var v12 uint32 = v3
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := []int{(fi13 << ((1) & 63)), (-(fi13)), (fi13 * fi13), 32767}
            fmt.Println(v3, v4)
            switch -2147483648 {
            case 9223372036854775807:
                var v15 uint8 = 128
                _ = v15
            case 128:
                v3 += (^(v1))
            case 641:
                v14 = append(v14, fi13)
            }
            _ = fi13
            _ = v14
        }
        var v16 uint64 = 801
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        v18 := GenS1{F0: true, F1: (v2 >> ((2) & 63))}
        switch (fi17 - (-(fi17))) {
        case 4294967295:
            var v19 int8 = v4
            _ = v19
        case 127:
            v5 += 2
        }
        _ = fi17
        _ = v18
    }
    if false {
        if ((true || false) || (2 >= 18446744073709551615)) {
            v6 = ((v6 * (0.0 - float32(1.5))) + ((0.0 + ((v6 + 1.5) / v6)) * 1e10))
            switch "foo bar" {
            case "a":
                var v20 int8 = (v4 - 42)
                _ = v20
            case "\n\t":
                for fi21 := 0; fi21 < 12; fi21++ {
                    fmt.Println(v2, v1, v6)
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v23 := GenS2{F0: 460}
                        _ = fi22
                        _ = v23
                    }
                    _ = fi21
                }
            }
        }
        for fi24 := 0; fi24 < 19; fi24++ {
            switch v7 {
            case "\n\t":
                fmt.Println(v4)
            default:
                var v25 uint16 = 480
                _ = v25
            }
            _ = fi24
        }
    } else {
        if (false && false) {
            v26 := GenS1{F0: (v7 == "\n\t"), F1: 0}
            v27 := GenS1{F0: (255 > 255), F1: -9223372036854775808}
            v28 := []int{9223372036854775807, 127}
            _ = v26
            _ = v27
            _ = v28
        }
    }
    if false {
        var v29 int32 = 1
        _ = v29
    } else {
        var v30 bool = (18446744073709551615 != 382)
        v31 := GenS2{F0: 0}
        _ = v30
        _ = v31
    }
    if (!(true && true)) {
        if (!(32767 >= 42)) {
            v3 += 399
        }
    }
    v32 := []int{342, 127, -880}
    for fi33 := 0; fi33 < 13; fi33++ {
        var v34 float32 = ((-2.25 + (((1e10 + v6) + ((3.14159 / ((-2.25 - v6) - v6)) * 0.1)) * 1.5)) + v6)
        var v35 int = fi33
        v36 := GenS1{F0: ("a" >= ""), F1: 1}
        _ = fi33
        _ = v34
        _ = v35
        _ = v36
    }
    v3 -= 2147483647
    v37 := GenS3{F0: 1000}
    if (int16(v2) == 32767) {
        if (127 <= 127) {
            v8 = (v1 << ((65535) & 63))
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
        }
        v38 := GenS3{F0: 809}
        _ = v38
    } else {
        var v39 bool = (-1 > 32767)
        _ = v39
    }
    v40 := GenS3{F0: int16(v8)}
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        v42 := GenS4{F0: v7}
        v43 := GenS3{F0: 328}
        _ = fi41
        _ = v42
        _ = v43
    }
    if (true || (((-363 >= 2147483647) && (v3 >= v1)) || false)) {
        var v44 int16 = -189
        _ = v44
    } else {
        v5 += v5
        var v45 int32 = -150
        if len(v32) > 0 {
            v32[0] = -2147483648
        }
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
