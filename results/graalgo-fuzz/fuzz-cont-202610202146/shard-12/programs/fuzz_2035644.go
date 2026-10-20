package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 string; F2 uint32 }
func genFn0() uint16 {
    var v5 int16 = 269
    v6 := GenS1{F0: (false || (!("foo bar" <= "a"))), F1: 238}
    _ = v5
    _ = v6
    return 17
}
func genFn1() int32 {
    var v7 int16 = 1
    var v8 int16 = v7
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 int32 = 45
    v2 := []string{"hello", "a"}
    v3 := []int{42, -9223372036854775808, 32768, 32767}
    var v4 int16 = 31
    if false {
        if (false && (!true)) {
            if (1 != 4294967295) {
                v1 = (v1 % ((-732) | 1))
                v4 -= v4
            } else {
                fmt.Println(v4)
            }
        } else {
            if (!(119 < 128)) {
                v1 = v1
                for fi9 := 0; fi9 < 14; fi9++ {
                    var v10 float32 = ((-2.25 + ((3.14159 * 1.5) - ((3.14159 / (1e10 * float32(0.1))) - 0.0))) / 1e10)
                    _ = fi9
                    _ = v10
                }
                switch "a" {
                case "hello":
                    var v11 int16 = -991
                    _ = v11
                case "":
                    _ = len(v2)
                }
            } else {
                _ = len(v3)
                fmt.Println(v1)
                v1 += (v1 - 2147483647)
            }
        }
        for fi12 := 0; fi12 < 16; fi12++ {
            fmt.Println(v1)
            _ = fi12
        }
    }
    var v13 int16 = v4
    v13 += (v4 - (v4 - v13))
    fmt.Println(v13)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v4, v4, v4)
    v14 := GenS2{F0: 0, F1: -1, F2: 133}
    if true {
        for fi15 := 0; fi15 < 10; fi15++ {
            for fi16 := 0; fi16 < 7; fi16++ {
                v4 -= int16(fi15)
                _ = fi16
            }
            _ = len(v2)
            var v17 int16 = (v4 / ((-272) | 1))
            _ = fi15
            _ = v17
        }
        if (255 < 92) {
            switch "" {
            case "\n\t":
                v1 -= genFn1()
            case "":
                if ((v13 > (-(v4))) || (140 > 128)) {
                    var v18 uint32 = 65535
                    _ = v18
                }
            case "foo bar":
                var v19 int = -2147483648
                _ = v19
            default:
                v20 := GenS3{F0: 32768}
                _ = v20
            }
        } else {
            v1 -= genFn1()
        }
        switch 472 {
        case 32767:
            if (!((true && false) && (true && (279 <= 32768)))) {
                v21 := GenS2{F0: 1, F1: 127, F2: 42}
                v13 = v4
                _ = v21
            }
        case 1:
            var v22 uint = 42
            _ = v22
        default:
            var v23 uint32 = 0
            _ = v23
        }
    } else {
        fmt.Println(v4, v4)
        v13 -= v4
        var v24 int8 = 127
        _ = v24
    }
    if true {
        for fi25 := 0; fi25 < 3; fi25++ {
            _ = len(v2)
            for fi26 := 0; fi26 < 20; fi26++ {
                v13 = 1000
                v27 := GenS3{F0: fi25}
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
        if false {
            for fi28 := 0; fi28 < 16; fi28++ {
                v1 -= (v1 >> ((18446744073709551615) & 63))
                _ = fi28
            }
            if (4294967295 >= 42) {
                fmt.Println(v1, v4, v1)
            } else {
                v13 += 1000
            }
            v2 = append(v2, "foo bar")
        } else {
            var v29 int16 = v4
            _ = v29
        }
    }
    var v30 float32 = 0.0
    v31 := GenS4{F0: 1, F1: "foo bar", F2: 1}
    var v32 float32 = v30
    v1 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
