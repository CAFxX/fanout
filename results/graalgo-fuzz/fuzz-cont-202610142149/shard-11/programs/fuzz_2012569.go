package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int32 }
type GenS2 struct { F0 bool; F1 uint; F2 bool }
type GenS3 struct { F0 int16; F1 uint16; F2 int8 }

func main() {
    var v1 int16 = 32767
    var v2 uint = 65535
    var v3 int = 128
    v4 := []string{("a" + ("" + ("\n\t" + "")))}
    var v5 uint = 65535
    for fi6 := 0; fi6 < 12; fi6++ {
        v3 += fi6
        switch "hello" {
        case "\n\t":
            v5 -= 2
        case "":
            var v7 float64 = (0.1 + (0.1 * (1e10 - (1e10 + (-2.25 + float64(-2.25))))))
            _ = v7
        default:
            v5 += (^(v5))
        }
        _ = fi6
    }
    if (1 != 1) {
        if (false || (("hello" + "\n\t") >= "hello")) {
            fmt.Println(v1)
            v8 := GenS1{F0: ("" + ("hello" + "")), F1: 271, F2: int32(v3)}
            _ = v8
        } else {
            var v9 uint32 = 995
            _ = v9
        }
        v10 := GenS1{F0: "a", F1: 1, F2: -2147483648}
        _ = v10
    } else {
        v11 := GenS1{F0: "", F1: 0, F2: 1}
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 float32 = 3.14159
        _ = fi12
        _ = v13
    }
    fmt.Println(v3)
    v14 := GenS2{F0: (!(42 > 128)), F1: 4294967295, F2: (false || (false && false))}
    if (!true) {
        v2 += 918
        fmt.Println(v2, v3)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        if false {
            var v15 uint64 = 58
            fmt.Println(v15)
            var v16 float32 = 3.14159
            _ = v15
            _ = v16
        } else {
            if (!(-9223372036854775808 < 42)) {
                for fi17 := 0; fi17 < 16; fi17++ {
                    v18 := GenS3{F0: v1, F1: 2, F2: 0}
                    var v19 uint = (v2 - (^(v5)))
                    switch v3 {
                    case 0:
                        v19 -= (v5 + 968)
                    case 2147483647:
                        v5 = 255
                    default:
                        v20 := GenS1{F0: "foo bar", F1: 898, F2: 42}
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
        }
        if (9223372036854775807 < 9223372036854775807) {
            switch -375 {
            case 128:
                fmt.Println(v5, v5)
            default:
                var v21 uint16 = 0
                _ = v21
            }
        }
        switch "hello" {
        case "foo bar":
            switch (v3 / ((v3) | 1)) {
            case 2147483647:
                v5 = (v2 & v5)
            case 42:
                v2 += v5
            default:
                var v22 float64 = 0.0
                _ = v22
            }
        case "":
            if false {
                v23 := GenS1{F0: "hello", F1: 357, F2: 1}
                if (!((-1 > 1) && (!(479 < -9223372036854775808)))) {
                    v24 := GenS3{F0: (v1 - v1), F1: 740, F2: -126}
                    v1 = v1
                    _ = v24
                }
                if true {
                    var v25 string = ("" + "foo bar")
                    v4 = append(v4, v25)
                    _ = v25
                }
                _ = v23
            }
        }
    }
    var v26 uint64 = 512
    v5 = (v2 << ((v5) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
