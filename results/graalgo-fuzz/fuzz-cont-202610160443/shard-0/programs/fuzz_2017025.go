package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 bool }
func genFn0(p0 string, p1 int8, p2 uint8) bool {
    p1 = (p1 % (((p1 / (((p1 + 23)) | 1))) | 1))
    p1 = 100
    var v6 uint8 = p2
    var v7 uint8 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return ((true || (0.0 == ((1.5 * 1e10) / float64(1.5)))) || false)
}

func main() {
    var v1 int = 42
    var v2 float32 = (1.5 - (3.14159 * float32(1.5)))
    var v3 uint64 = 2
    v4 := []string{("\n\t" + "a")}
    var v5 int64 = 42
    var v8 int8 = 42
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch ("a" + ("foo bar" + (("a" + (("" + ("hello" + ("\n\t" + ("" + "foo bar")))) + "foo bar")) + "\n\t"))) {
    case "\n\t":
        v5 = -1
    case "":
        var v9 float64 = ((1e10 - 3.14159) - 1e10)
        _ = v9
    default:
        fmt.Println(v3, v1, v8)
    }
    v1 -= v1
    switch (^(v1)) {
    case 127:
        if len(v4) > 0 {
            v4[0] = "a"
        }
    default:
        if (true || true) {
            var v10 int64 = 9223372036854775807
            _ = v10
        } else {
            if (!(962 >= 65535)) {
                if false {
                    var v11 uint64 = (v3 >> ((141) & 63))
                    _ = v11
                }
                if len(v4) > 0 {
                    v4[0] = "a"
                }
                v12 := GenS1{F0: 65535}
                _ = v12
            }
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if ((("foo bar" + ("\n\t" + ("a" + "hello"))) + ((((("a" + "foo bar") + ((("\n\t" + "\n\t") + ("\n\t" + "hello")) + "")) + "a") + "hello") + "hello")) >= "") {
        var v13 bool = (!((685 > 1) || false))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v13
    }
    switch "" {
    case "":
        switch "foo bar" {
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
        default:
            for fi14 := 0; fi14 < 20; fi14++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi14
            }
        }
    case "foo bar":
        switch (v1 & v1) {
        case 4294967295:
            v5 += v5
        }
    case "hello":
        var v15 uint64 = v3
        _ = v15
    default:
        if (527 == 2) {
            for fi16 := 0; fi16 < 18; fi16++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v17 := GenS2{F0: 1000, F1: true}
                var v18 float64 = (0.0 + -2.25)
                _ = fi16
                _ = v17
                _ = v18
            }
            if false {
                var v19 uint64 = (v3 / ((798) | 1))
                v5 += (^(v5))
                var v20 int = (v1 & (v1 ^ 32767))
                _ = v19
                _ = v20
            }
            if false {
                v8 = 1
                switch 79 {
                case -801:
                    v5 += 42
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        } else {
            v21 := GenS2{F0: 1, F1: ((!(-2.25 <= 1.5)) && (675 >= 2))}
            _ = v21
        }
    }
    v1 += (v1 >> ((2) & 63))
    fmt.Println(v1, v1)
    var v22 float32 = (1e10 + v2)
    var v23 uint8 = 255
    if true {
        var v24 int32 = 1
        var v25 int8 = (v8 ^ (v8 % (((v8 - (v8 / ((v8) | 1)))) | 1)))
        v24 += v24
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
