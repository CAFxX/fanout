package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint32 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int16 = 42
    v3 := []string{"hello", ""}
    v4 := make(map[uint32]int)
    v2 = (v2 * -1)
    if true {
        var v5 int8 = 1
        _ = v5
    }
    if true {
        var v6 int16 = (v2 % (((v2 * (v2 * -1))) | 1))
        switch ("foo bar" + "hello") {
        case "\n\t":
            if (v6 > v2) {
                if false {
                    v1 = v1
                    fmt.Println(v6, v6)
                } else {
                    v2 = v6
                }
                if true {
                    switch "a" {
                    case "foo bar":
                        var v7 uint = 4294967295
                        _ = v7
                    }
                    var v8 float64 = 0.1
                    var v9 float32 = 0.1
                    _ = v8
                    _ = v9
                }
                switch (("foo bar" + "\n\t") + (("hello" + "hello") + ("foo bar" + "a"))) {
                case "":
                    var v10 uint32 = 65535
                    _ = v10
                case "hello":
                    v3 = append(v3, "foo bar")
                case "\n\t":
                    fmt.Println(v1, v2)
                default:
                    var v11 int16 = 1000
                    _ = v11
                }
            }
        case "hello":
            if true {
                for fi12 := 0; fi12 < 7; fi12++ {
                    for fi13 := 0; fi13 < 17; fi13++ {
                        var v14 string = "foo bar"
                        _ = len(v3)
                        _ = fi13
                        _ = v14
                    }
                    _ = fi12
                }
                v6 = 32767
            } else {
                v6 += v2
            }
        case "a":
            v6 += int16(v1)
        }
        _ = v6
    }
    fmt.Println(v1, v1)
    for fi15 := 0; fi15 < 6; fi15++ {
        for fi16 := 0; fi16 < 16; fi16++ {
            v2 = (v2 + (v2 << ((1) & 63)))
            var v17 uint32 = 2
            var v18 int = fi15
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    var v19 uint32 = 567
    for fi20 := 0; fi20 < 10; fi20++ {
        _ = len(v3)
        for fi21 := 0; fi21 < 18; fi21++ {
            v19 += (v19 % (((v19 - 935)) | 1))
            _ = fi21
        }
        _ = fi20
    }
    var v22 uint16 = uint16(v19)
    if false {
        for fi23 := 0; fi23 < 10; fi23++ {
            fmt.Println(v2, v19)
            _ = fi23
        }
    }
    switch "foo bar" {
    case "":
        v24 := GenS1{F0: (v1 * v1), F1: -128, F2: 4294967295}
        _ = v24
    case "a":
        for fi25 := 0; fi25 < 13; fi25++ {
            for fi26 := 0; fi26 < 2; fi26++ {
                v27 := GenS1{F0: v1, F1: 100, F2: v19}
                switch 568 {
                case 32768:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    v19 = 65535
                }
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    default:
        var v28 uint64 = 0
        _ = v28
    }
    if false {
        v22 += v22
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        var v30 string = ("a" + "")
        if true {
            var v31 float64 = (3.14159 - 1e10)
            _ = v31
        }
        for fi32 := 0; fi32 < 2; fi32++ {
            v30 = v30
            fmt.Println(v30)
            v3 = append(v3, v30)
            _ = fi32
        }
        _ = fi29
        _ = v30
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        var v34 int32 = 861
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
