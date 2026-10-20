package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int64 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint64; F1 uint8; F2 int8 }

func main() {
    var v1 uint = 1
    var v2 uint64 = 42
    v3 := []string{("foo bar" + "foo bar"), ("" + "\n\t"), "hello", ""}
    var v4 int16 = -1
    fmt.Println(v1, v4)
    if (-561 != 9223372036854775807) {
        v2 += (v2 << (((v1 + (v1 | (v1 >> ((1) & 63))))) & 63))
    }
    v4 = -1
    fmt.Println(v2, v2, v4)
    v2 = (v2 / ((v2) | 1))
    switch 4294967295 {
    case 42:
        v2 = (v2 << ((v1) & 63))
    case 32768:
        fmt.Println(v4)
    }
    for fi5 := 0; fi5 < 2; fi5++ {
        v2 = 4294967295
        v2 += 4294967295
        _ = fi5
    }
    if (!(v2 <= 42)) {
        var v6 uint8 = 176
        v7 := GenS1{F0: "\n\t", F1: 1, F2: 1}
        _ = v6
        _ = v7
    } else {
        var v8 uint16 = 0
        if (v1 > (v1 + 384)) {
            var v9 float32 = float32(((1.5 - -2.25) * -2.25))
            _ = v9
        }
        _ = v8
    }
    var v10 float64 = float64(float32((-2.25 - 1e10)))
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS2{F0: -1, F1: ("a" <= "a")}
        v13 := GenS2{F0: (v4 % ((v4) | 1)), F1: (true || ((533 >= 65535) || (v4 != (v4 >> ((v1) & 63)))))}
        _ = fi11
        _ = v12
        _ = v13
    }
    switch (("" + "a") + "foo bar") {
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "":
        v2 -= 2
    default:
        if true {
            v14 := GenS1{F0: ((("hello" + ("" + "")) + "foo bar") + "\n\t"), F1: 0, F2: 302}
            var v15 float32 = 1.5
            switch "foo bar" {
            case "":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "a":
                v2 -= v2
            }
            _ = v14
            _ = v15
        } else {
            if false {
                if true {
                    v2 -= (v2 | 208)
                }
            }
        }
    }
    v16 := GenS2{F0: v4, F1: (("" + "") == (("a" + ("" + "foo bar")) + "a"))}
    fmt.Println(v4)
    if (-1 > 42) {
        v1 = v1
    }
    if true {
        for fi17 := 0; fi17 < 11; fi17++ {
            v1 = v1
            var v18 int8 = 52
            _ = fi17
            _ = v18
        }
    } else {
        switch -723 {
        case -858:
            for fi19 := 0; fi19 < 13; fi19++ {
                v1 -= 2
                v20 := GenS3{F0: (v2 << ((427) & 63)), F1: 196, F2: 100}
                _ = fi19
                _ = v20
            }
        case -265:
            v4 = v4
        case 127:
            _ = len(v3)
        default:
            v2 = 126
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
