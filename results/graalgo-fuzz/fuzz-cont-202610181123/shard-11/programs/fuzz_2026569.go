package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }

func main() {
    v1 := []uint32{4294967295, 1, 1, 1}
    var v2 uint16 = 32767
    var v3 uint32 = 65535
    v4 := make(map[uint32]bool)
    var v5 int = -650
    var v6 uint32 = v3
    switch "" {
    case "foo bar":
        if false {
            if (!true) {
                delete(v4, (v6 & (v3 % (((v6 >> ((162) & 63))) | 1))))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v5, v5, v6)
            } else {
                switch (v5 * (-(v5))) {
                case -1:
                    v7 := v4[2]
                    _ = v7
                case 127:
                    var v8 float64 = 3.14159
                    _ = v8
                case 32768:
                    var v9 uint32 = 2147483647
                    _ = v9
                default:
                    switch ("\n\t" + ("a" + ("a" + ("foo bar" + (("" + ("" + ((((((("foo bar" + ((("\n\t" + ("a" + "\n\t")) + (("foo bar" + "") + ("foo bar" + (("hello" + ("foo bar" + "a")) + "a")))) + "\n\t")) + ("foo bar" + ("hello" + "hello"))) + (((((("" + ("a" + ("" + ((((("" + ("hello" + "")) + "foo bar") + "hello") + "a") + "")))) + "\n\t") + (("" + "\n\t") + "\n\t")) + ("\n\t" + "foo bar")) + "a") + ((("a" + (("\n\t" + "a") + "hello")) + (("foo bar" + "") + "a")) + ""))) + "hello") + "\n\t") + "foo bar") + "foo bar"))) + "hello"))))) {
                    case "\n\t":
                        v10 := GenS1{F0: 1}
                        _ = v10
                    default:
                        _ = len(v4)
                    }
                }
                v5 -= (v5 / ((-515) | 1))
                switch v5 {
                case 127:
                    var v11 int8 = -22
                    _ = v11
                }
            }
        } else {
            v12 := GenS1{F0: 4294967295}
            v5 = (v5 / (((v5 / ((v5) | 1))) | 1))
            if true {
                var v13 int64 = 9223372036854775807
                fmt.Println(v13, v13)
                var v14 int64 = 1
                _ = v13
                _ = v14
            }
            _ = v12
        }
    }
    v6 = uint32(v2)
    for fi15 := 0; fi15 < 4; fi15++ {
        if false {
            var v16 int = fi15
            fmt.Println(v5)
            _ = v16
        } else {
            var v17 int64 = 1
            _ = v17
        }
        var v18 uint32 = (v3 * v3)
        _ = fi15
        _ = v18
    }
    v2 = 143
    switch ("foo bar" + ("\n\t" + (((("foo bar" + "") + "foo bar") + "foo bar") + "hello"))) {
    case "a":
        fmt.Println(v6, v5)
    case "hello":
        switch ("hello" + "hello") {
        case "hello":
            v19 := GenS2{F0: (false || (!true))}
            _ = v19
        default:
            var v20 uint32 = v3
            _ = v20
        }
    }
    fmt.Println(v5, v6, v3)
    fmt.Println(v5)
    if false {
        v21 := GenS1{F0: 0}
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
