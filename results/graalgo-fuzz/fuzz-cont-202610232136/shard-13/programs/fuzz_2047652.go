package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint; F2 uint16 }

func main() {
    var v1 int64 = -1
    var v2 int16 = -32768
    var v3 int = 4294967295
    var v4 uint16 = 65535
    var v5 int32 = 2147483647
    var v6 uint16 = (v4 + v4)
    for fi7 := 0; fi7 < 19; fi7++ {
        if true {
            v8 := []int{(v3 % (((v3 & fi7)) | 1)), fi7, fi7}
            _ = v8
        }
        switch 708 {
        case -1:
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := []int{fi7, 42, fi9, int(v4)}
                if len(v10) > 0 {
                    v10[0] = (fi7 ^ (fi9 / (((^(v3))) | 1)))
                }
                if false {
                    var v11 uint8 = 42
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
        case 2147483647:
            v6 = v4
        case -9223372036854775808:
            switch (v3 / ((189) | 1)) {
            case 9223372036854775807:
                v2 = -308
            case -250:
                if (((!false) && true) && false) {
                    v12 := []int{128, fi7, fi7, int(v2)}
                    _ = v12
                }
            case 32767:
                v3 += (v3 + (fi7 - v3))
            default:
                var v13 int8 = 28
                _ = v13
            }
        default:
            v5 = (v5 * v5)
        }
        _ = fi7
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 uint = 0
        fmt.Println(v6, v3)
        v3 -= 1
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 float64 = (0.1 / (0.1 * 3.14159))
        _ = fi16
        _ = v17
    }
    v5 -= 562
    v18 := GenS1{F0: 42}
    var v19 bool = true
    if ((4294967295 < 255) && (((v19 || ((v1 >> ((18446744073709551615) & 63)) != v1)) && ((v2 & 42) == (v2 - int16(v1)))) || v19)) {
        var v20 float32 = -2.25
        _ = v20
    }
    fmt.Println(v4)
    v21 := GenS1{F0: 65535}
    v2 = -278
    var v22 uint32 = 2147483648
    if (-2.25 >= (((-2.25 * float64(3.14159)) * -2.25) / 1e10)) {
        v23 := GenS2{F0: (-(v1)), F1: 2, F2: (v6 >> ((42) & 63))}
        v24 := []int{}
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 int = (fi25 % ((int(v5)) | 1))
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
            _ = fi25
            _ = v26
        }
        _ = v23
        _ = v24
    } else {
        var v27 uint8 = 1
        v22 += (v22 % (((v22 % ((2147483647) | 1))) | 1))
        var v28 uint32 = (^(v22))
        _ = v27
        _ = v28
    }
    switch ("\n\t" + ("hello" + "foo bar")) {
    case "\n\t":
        var v29 uint64 = 2
        _ = v29
    }
    switch (v3 & (v3 | (v3 % ((v3) | 1)))) {
    case 531:
        var v30 uint16 = (v4 + (v4 / ((v4) | 1)))
        _ = v30
    }
    var v31 int = v3
    var v32 bool = ((v5 >> ((255) & 63)) == (v5 << ((1) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
