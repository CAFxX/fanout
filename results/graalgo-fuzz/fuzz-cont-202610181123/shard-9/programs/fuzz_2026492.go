package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 bool = (-256 >= 718)
    var v2 int64 = -1
    var v3 int = -9223372036854775808
    var v4 float32 = 1e10
    v5 := []uint32{2, 4294967295}
    var v6 int64 = (v2 / ((663) | 1))
    v3 += (-(v3))
    for fi7 := 0; fi7 < 7; fi7++ {
        if (v1 || v1) {
            var v8 int16 = 42
            if v1 {
                var v9 uint = 315
                var v10 int8 = -128
                if v1 {
                    v6 = v2
                    if false {
                        if len(v5) > 0 {
                            v5[0] = 42
                        }
                    }
                    var v11 int64 = (v2 >> (((v9 + v9)) & 63))
                    _ = v11
                }
                _ = v10
                _ = v9
            }
            _ = v8
        }
        if v1 {
            for fi12 := 0; fi12 < 1; fi12++ {
                v3 = (v3 * v3)
                switch ((("a" + "a") + ((("hello" + "hello") + ("foo bar" + "")) + ("" + "hello"))) + (("hello" + ("" + (("hello" + "hello") + "a"))) + "foo bar")) {
                case "hello":
                    var v13 bool = (663 < 255)
                    _ = v13
                }
                _ = fi12
            }
            var v14 uint32 = 522
            _ = v14
        } else {
            switch "\n\t" {
            case "hello":
                var v15 uint = 60
                _ = v15
            }
            switch (fi7 << ((4294967295) & 63)) {
            case 32768:
                var v16 uint32 = 42
                _ = v16
            case -9223372036854775808:
                var v17 int16 = 1000
                _ = v17
            default:
                var v18 uint64 = 4294967295
                _ = v18
            }
            var v19 float32 = v4
            _ = v19
        }
        _ = fi7
    }
    switch "hello" {
    case "\n\t":
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := GenS1{F0: 0}
            _ = fi20
            _ = v21
        }
    default:
        if (749 == 1) {
            v22 := GenS2{F0: 1}
            v6 -= 0
            _ = v22
        } else {
            for fi23 := 0; fi23 < 1; fi23++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                fmt.Println(fi23, v4, v2)
                _ = fi23
            }
            v2 -= 9223372036854775807
        }
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 uint8 = 42
        _ = len(v5)
        _ = fi24
        _ = v25
    }
    v26 := GenS2{F0: 0}
    switch ((("foo bar" + ("foo bar" + "hello")) + "a") + ("" + "\n\t")) {
    case "hello":
        v5 = append(v5, 0)
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = 545
        }
    }
    switch "\n\t" {
    case "a":
        v3 = (v3 & 32768)
    }
    fmt.Println(v2)
    v6 -= 0
    fmt.Println(v6, v2)
    fmt.Println(v6, v1, v4)
    fmt.Println(v3)
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
