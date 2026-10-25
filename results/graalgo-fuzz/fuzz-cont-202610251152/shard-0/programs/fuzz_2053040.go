package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int32; F2 int }
func genFn0(p0 uint64) uint {
    fmt.Println(p0)
    if (true && false) {
        var v4 uint8 = 255
        _ = v4
    } else {
        if (3.14159 != -2.25) {
            if (!true) {
                v5 := []int{-9223372036854775808, -1, 42}
                _ = v5
            }
            if (98 != 255) {
                p0 += p0
            } else {
                switch "foo bar" {
                case "a":
                    v6 := []int{448, 2147483647, 1, -2147483648}
                    _ = v6
                default:
                    p0 -= p0
                }
                fmt.Println(p0)
            }
            var v7 uint64 = p0
            _ = v7
        } else {
            var v8 bool = false
            var v9 uint16 = 1
            _ = v8
            _ = v9
        }
        switch 4294967295 {
        case 42:
            switch 0 {
            case 1:
                fmt.Println(p0)
            default:
                var v10 string = "\n\t"
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            fmt.Println(fi11, fi11)
            var v12 string = "hello"
            _ = fi11
            _ = v12
        }
    }
    p0 += (p0 << ((667) & 63))
    _ = p0
    return 18446744073709551615
}

func main() {
    var v1 bool = true
    var v2 int64 = -1
    var v3 int = 679
    var v13 uint32 = 4294967295
    switch -966 {
    case -9223372036854775808:
        if true {
            fmt.Println(v3, v2)
            v14 := GenS1{F0: 1000, F1: 125}
            switch v3 {
            case 366:
                v15 := GenS1{F0: 277, F1: 0}
                _ = v15
            default:
                v16 := GenS1{F0: 1, F1: -113}
                _ = v16
            }
            _ = v14
        }
    default:
        v2 = v2
    }
    if (3.14159 > (-2.25 / 3.14159)) {
        v17 := []int{(v3 << ((757) & 63)), (v3 / ((v3) | 1)), (-(v3)), v3}
        var v18 uint32 = v13
        if v1 {
            v2 = (v2 % (((v2 + v2)) | 1))
            v1 = v1
        } else {
            if (!(0.0 != 0.1)) {
                v17 = append(v17, (v3 ^ 127))
            } else {
                _ = len(v17)
                var v19 uint64 = 2
                if len(v17) > 0 {
                    v17[0] = 1
                }
                _ = v19
            }
            var v20 int64 = -1
            _ = v20
        }
        _ = v17
        _ = v18
    } else {
        v1 = false
    }
    switch (v3 << ((65535) & 63)) {
    case 1:
        var v21 uint8 = 128
        _ = v21
    case 42:
        v22 := []int{(v3 | v3), (v3 ^ (v3 / (((-(v3))) | 1)))}
        _ = v22
    }
    switch (("foo bar" + "a") + "a") {
    case "":
        v1 = false
    case "\n\t":
        v1 = v1
    }
    switch v3 {
    case 2147483647:
        var v23 int8 = 112
        _ = v23
    case 4294967295:
        var v24 uint32 = (^(v13))
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        var v26 int64 = v2
        _ = fi25
        _ = v26
    }
    v1 = v1
    for fi27 := 0; fi27 < 14; fi27++ {
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 int64 = v2
            fmt.Println(v3)
            v29 = (-(v2))
            _ = fi28
            _ = v29
        }
        for fi30 := 0; fi30 < 12; fi30++ {
            if (!(-1 <= int16(v13))) {
                v31 := []int{}
                switch ("a" + "foo bar") {
                case "\n\t":
                    v32 := GenS1{F0: 1, F1: int8(v13)}
                    _ = v32
                case "hello":
                    v33 := GenS1{F0: 296, F1: 127}
                    _ = v33
                case "foo bar":
                    v3 -= (fi27 / ((fi30) | 1))
                default:
                    if len(v31) > 0 {
                        v31[0] = fi27
                    }
                }
                if (!(!(v1 || false))) {
                    _ = len(v31)
                } else {
                    var v34 int = fi30
                    v34 = fi27
                    _ = v34
                }
                _ = v31
            } else {
                var v35 int8 = 100
                v13 += (v13 >> ((4294967295) & 63))
                var v36 int8 = (v35 | 0)
                _ = v35
                _ = v36
            }
            var v37 uint16 = 204
            var v38 int = (fi30 << ((42) & 63))
            v39 := GenS2{F0: 55}
            v40 := GenS1{F0: 0, F1: -1}
            _ = fi30
            _ = v37
            _ = v38
            _ = v39
            _ = v40
        }
        _ = fi27
    }
    v3 += -1
    v41 := GenS3{F0: 255, F1: -1, F2: (v3 ^ (v3 * (v3 & -457)))}
    if (v13 > v13) {
        v3 = -308
    }
    switch (-(v3)) {
    case 128:
        var v42 string = ("\n\t" + ("" + "foo bar"))
        _ = v42
    case 2147483647:
        v43 := []int{v3}
        _ = v43
    case 32767:
        switch "" {
        case "a":
            switch (("hello" + "foo bar") + "") {
            case "hello":
                v44 := GenS1{F0: -32768, F1: -1}
                _ = v44
            }
        case "hello":
            v45 := GenS3{F0: 37, F1: -2147483648, F2: 127}
            _ = v45
        case "\n\t":
            v46 := []int{(v3 >> ((114) & 63))}
            _ = v46
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v41
}
