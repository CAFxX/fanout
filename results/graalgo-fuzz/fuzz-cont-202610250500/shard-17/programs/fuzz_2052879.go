package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int; F2 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint; F2 uint }
func genFn0(p0 int8, p1 uint32) int8 {
    p0 = (p0 << ((547) & 63))
    switch ((("hello" + "foo bar") + "hello") + "foo bar") {
    case "":
        switch "a" {
        case "hello":
            p1 += (^(p1))
        case "foo bar":
            var v6 bool = false
            _ = v6
        }
    case "a":
        p0 += (-(p0))
    case "foo bar":
        var v7 int32 = -1
        _ = v7
    default:
        v8 := GenS1{F0: 42}
        _ = v8
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 string = "a"
    v2 := []uint32{612}
    v3 := []string{v1, v1, v1}
    v4 := make(map[uint32]uint32)
    v5 := []string{}
    v1 = (v1 + "hello")
    switch (v1 + ("\n\t" + v1)) {
    case "hello":
        v9 := GenS1{F0: 2}
        _ = v9
    case "":
        v1 = "\n\t"
    case "foo bar":
        var v10 float32 = float32(0.1)
        _ = v10
    default:
        var v11 float64 = 0.1
        _ = v11
    }
    v1 = "foo bar"
    v12 := GenS2{F0: -1, F1: 127, F2: 376}
    if false {
        for fi13 := 0; fi13 < 4; fi13++ {
            if (0.0 <= 1e10) {
                v4[42] = 586
            }
            var v14 uint32 = 2
            v1 = ""
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 6; fi15++ {
            switch (fi15 * (fi15 >> ((4294967295) & 63))) {
            case -217:
                v16 := GenS1{F0: 594}
                _ = v16
            case 23:
                for fi17 := 0; fi17 < 18; fi17++ {
                    v1 = v1
                    switch (fi15 >> ((1) & 63)) {
                    case 127:
                        fmt.Println(fi17, v1)
                    case 428:
                        var v18 int8 = 0
                        _ = v18
                    default:
                        _ = len(v4)
                    }
                    v4[829] = 42
                    _ = fi17
                }
            case 42:
                v19 := GenS3{F0: ((v1 + "foo bar") + v1)}
                _ = v19
            default:
                for fi20 := 0; fi20 < 17; fi20++ {
                    v3 = append(v3, (("" + "foo bar") + ("foo bar" + (v1 + ((v1 + ((v1 + (v1 + v1)) + (v1 + v1))) + (v1 + ""))))))
                    if (18446744073709551615 == 4294967295) {
                        fmt.Println(v1, fi20, v1)
                        var v21 uint = 1
                        _ = v21
                    }
                    var v22 float32 = 3.14159
                    _ = fi20
                    _ = v22
                }
            }
            v1 = v1
            fmt.Println(v1, fi15)
            _ = fi15
        }
    } else {
        for fi23 := 0; fi23 < 7; fi23++ {
            fmt.Println(fi23, fi23)
            _ = fi23
        }
    }
    v24 := GenS4{F0: 598, F1: 65535, F2: 2}
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 2; fi26++ {
            fmt.Println(v1, v1)
            v1 = ((v1 + v1) + (((((((v1 + "foo bar") + "foo bar") + v1) + v1) + ((v1 + (v1 + (v1 + "hello"))) + v1)) + v1) + v1))
            fmt.Println(fi25)
            _ = fi26
        }
        switch (v1 + (v1 + v1)) {
        case "":
            v1 = v1
        case "\n\t":
            if (true || true) {
                switch (fi25 & fi25) {
                case 9223372036854775807:
                    var v27 float64 = float64(0.0)
                    _ = v27
                case 289:
                    for fi28 := 0; fi28 < 11; fi28++ {
                        var v29 float32 = 3.14159
                        var v30 float64 = (((0.1 - -2.25) * 0.1) - float64(v29))
                        if len(v2) > 0 {
                            v2[0] = 798
                        }
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case -611:
                    v1 = ""
                default:
                    v31 := GenS2{F0: 2147483647, F1: (fi25 << ((921) & 63)), F2: 32767}
                    _ = v31
                }
                if ((0 >= -32768) && (405 >= 44)) {
                    var v32 uint8 = 42
                    _ = v32
                } else {
                    v1 = (v1 + "")
                }
                v1 = "\n\t"
            } else {
                var v33 int8 = 42
                var v34 string = (((("foo bar" + v1) + v1) + "a") + "a")
                _ = v33
                _ = v34
            }
        }
        if true {
            fmt.Println(fi25)
        }
        _ = fi25
    }
    var v35 string = (v1 + v1)
    fmt.Println(v35, v35)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
