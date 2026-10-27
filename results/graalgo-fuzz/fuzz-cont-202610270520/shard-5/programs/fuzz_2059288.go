package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int8 }

func main() {
    var v1 uint64 = 4294967295
    var v2 int64 = 863
    v3 := make(map[uint32]uint32)
    var v4 float32 = 0.1
    var v5 int8 = 100
    var v6 uint8 = 255
    v7 := GenS1{F0: 657}
    _ = len(v3)
    if true {
        fmt.Println(v4, v6, v2)
        if (!(493 > 0)) {
            var v8 int8 = (v5 * v5)
            v9 := GenS2{F0: 32767}
            if (v6 <= v6) {
                v10 := GenS2{F0: 1}
                var v11 int32 = 576
                _ = v10
                _ = v11
            }
            _ = v8
            _ = v9
        } else {
            v4 = ((-2.25 / ((float32((0.1 * -2.25)) * v4) / (v4 - v4))) * 0.1)
        }
    }
    fmt.Println(v4, v6, v6)
    for fi12 := 0; fi12 < 12; fi12++ {
        if true {
            fmt.Println(v4, fi12)
            v5 -= v5
        }
        _ = fi12
    }
    v13 := []int{625}
    switch "a" {
    case "hello":
        for fi14 := 0; fi14 < 16; fi14++ {
            switch (fi14 - (fi14 % (((fi14 - fi14)) | 1))) {
            case 2147483647:
                v15 := GenS3{F0: (^(v6)), F1: 255, F2: (fi14 + fi14)}
                _ = v15
            }
            for fi16 := 0; fi16 < 8; fi16++ {
                v17 := GenS1{F0: 4294967295}
                _ = fi16
                _ = v17
            }
            _ = fi14
        }
    case "":
        if ((("a" + ("foo bar" + "foo bar")) + ("hello" + "")) != "") {
            fmt.Println(v2, v6, v4)
        }
    case "foo bar":
        v2 = v2
    }
    v18 := GenS4{F0: 42}
    for fi19 := 0; fi19 < 18; fi19++ {
        delete(v3, 826)
        for fi20 := 0; fi20 < 6; fi20++ {
            v6 = 128
            fmt.Println(fi19, v5, v2)
            switch ("\n\t" + "\n\t") {
            case "a":
                v21 := GenS5{F0: v2, F1: 32767}
                _ = v21
            default:
                switch (fi20 % (((fi19 - 32768)) | 1)) {
                case 62:
                    v13 = append(v13, -1)
                }
            }
            _ = fi20
        }
        if false {
            switch int(v1) {
            case -136:
                if true {
                    switch fi19 {
                    case 32767:
                        v4 += (v4 - v4)
                    }
                    v1 -= (v1 + v1)
                    var v22 int = fi19
                    _ = v22
                } else {
                    v23 := GenS5{F0: (v2 ^ v2), F1: 0}
                    v1 = (v1 + (^(v1)))
                    _ = v23
                }
            default:
                if ((v6 ^ v6) <= v6) {
                    var v24 float32 = float32(-2.25)
                    v25 := GenS1{F0: 2147483648}
                    _ = v24
                    _ = v25
                }
            }
        }
        _ = fi19
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        switch (fi26 - (^(fi26))) {
        case -439:
            switch fi26 {
            case 1:
                v27 := GenS6{F0: v5}
                _ = v27
            case -9223372036854775808:
                v5 += (v5 % ((67) | 1))
            }
        default:
            if (true && true) {
                if false {
                    v2 -= v2
                } else {
                    if len(v13) > 0 {
                        v13[0] = fi26
                    }
                }
                v1 -= 18446744073709551615
            } else {
                var v28 int = 32768
                for fi29 := 0; fi29 < 10; fi29++ {
                    v28 += (fi26 * (fi29 << ((892) & 63)))
                    v30 := GenS4{F0: 0}
                    switch fi29 {
                    case 123:
                        if (true && true) {
                            if (32767 != 42) {
                                v2 -= v2
                                switch fi29 {
                                case 0:
                                    v4 = 0.1
                                }
                                _ = len(v13)
                                fmt.Println(v5, v1, v5)
                                delete(v3, 42)
                            } else {
                                v2 -= v2
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                var v31 uint16 = 10
                                switch (-(fi26)) {
                                case -9223372036854775808:
                                    _ = len(v13)
                                case 9223372036854775807:
                                    fmt.Println(v28)
                                }
                                _ = v31
                            }
                        } else {
                            v3[2147483647] = 1
                            v5 = (v5 * 0)
                            if true {
                                if true {
                                    v4 += v4
                                    var v32 int = (^(fi29))
                                    _ = v32
                                }
                            }
                        }
                    }
                    _ = fi29
                    _ = v30
                }
                _ = v28
            }
        }
        _ = fi26
    }
    v1 = (v1 ^ (v1 + (v1 << ((4294967295) & 63))))
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 int16 = 863
        _ = fi33
        _ = v34
    }
    v5 = v5
    fmt.Println(v6, v2)
    var v35 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
