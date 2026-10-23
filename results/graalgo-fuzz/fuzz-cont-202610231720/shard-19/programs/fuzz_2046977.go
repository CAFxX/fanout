package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint16; F2 string }
type GenS3 struct { F0 uint64; F1 uint64; F2 int64 }

func main() {
    var v1 uint16 = 32767
    v2 := []string{"\n\t", "hello"}
    var v3 float64 = (-2.25 - 1.5)
    v4 := []string{}
    var v5 uint64 = 0
    var v6 uint8 = 127
    v7 := GenS1{F0: -633}
    var v8 uint = 0
    v8 -= v8
    for fi9 := 0; fi9 < 17; fi9++ {
        v3 = (v3 + v3)
        _ = fi9
    }
    switch ("a" + "\n\t") {
    case "":
        var v10 uint32 = 2
        _ = v10
    case "foo bar":
        for fi11 := 0; fi11 < 8; fi11++ {
            switch "" {
            case "a":
                switch "foo bar" {
                case "":
                    if true {
                        fmt.Println(v1, v1)
                    } else {
                        var v12 bool = (-176 < 1)
                        _ = v12
                    }
                }
            case "hello":
                for fi13 := 0; fi13 < 4; fi13++ {
                    if len(v4) > 0 {
                        v4[0] = (("foo bar" + ("hello" + ("a" + ("foo bar" + ((("hello" + "foo bar") + (("a" + "hello") + "foo bar")) + ("" + ("hello" + ("a" + "a")))))))) + "hello")
                    }
                    _ = fi13
                }
            case "\n\t":
                var v14 float32 = 0.1
                _ = v14
            default:
                var v15 uint32 = 4294967295
                _ = v15
            }
            var v16 float64 = (v3 - v3)
            v17 := GenS1{F0: 4294967295}
            _ = fi11
            _ = v16
            _ = v17
        }
    case "\n\t":
        var v18 int16 = int16(v1)
        _ = v18
    }
    var v19 float32 = float32(v3)
    var v20 float32 = (v19 - v19)
    var v21 string = (("a" + (("a" + "") + "")) + ("\n\t" + "hello"))
    var v22 int32 = 2147483647
    var v23 int64 = 1
    v19 += float32((((float64((3.14159 * float32(((3.14159 * v3) - 0.0)))) * 0.1) * 0.0) / 0.1))
    var v24 uint16 = v1
    for fi25 := 0; fi25 < 15; fi25++ {
        _ = len(v2)
        _ = fi25
    }
    var v26 uint16 = 2
    switch v21 {
    case "hello":
        var v27 bool = true
        _ = v27
    }
    if (0 > 2147483648) {
        if true {
            for fi28 := 0; fi28 < 14; fi28++ {
                v29 := GenS1{F0: -2147483648}
                for fi30 := 0; fi30 < 8; fi30++ {
                    v6 += uint8(v1)
                    if false {
                        var v31 uint32 = 4294967295
                        _ = len(v4)
                        _ = v31
                    }
                    _ = fi30
                }
                _ = fi28
                _ = v29
            }
        } else {
            var v32 uint8 = v6
            var v33 int8 = 1
            switch v21 {
            case "a":
                for fi34 := 0; fi34 < 17; fi34++ {
                    switch (fi34 >> ((4294967295) & 63)) {
                    case 9223372036854775807:
                        v35 := GenS2{F0: (v33 * v33), F1: 162, F2: "foo bar"}
                        _ = v35
                    case 874:
                        v21 = "\n\t"
                    case 128:
                        v22 -= (v22 + (v22 ^ (v22 >> (((^(v8))) & 63))))
                    }
                    _ = fi34
                }
            case "":
                for fi36 := 0; fi36 < 17; fi36++ {
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    fmt.Println(v1, v22)
                    _ = fi36
                }
            }
            _ = v32
            _ = v33
        }
        for fi37 := 0; fi37 < 16; fi37++ {
            fmt.Println(v21)
            fmt.Println(v23, v5, v6)
            for fi38 := 0; fi38 < 10; fi38++ {
                var v39 int = -348
                _ = fi38
                _ = v39
            }
            _ = fi37
        }
    }
    v40 := GenS3{F0: v5, F1: (v5 * v5), F2: v23}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
