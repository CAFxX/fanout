package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint16) string {
    v6 := []int{4294967295}
    var v7 int64 = 1
    p0 = p0
    if (!true) {
        for fi8 := 0; fi8 < 17; fi8++ {
            p0 = p0
            _ = fi8
        }
        v7 = (-(v7))
    } else {
        var v9 uint8 = 1
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ""
}
func genFn1() uint {
    var v10 int = 128
    if (true && (((0.0 + (float32(1e10) * 1.5)) > 1e10) || (!true))) {
        if (2 > 42) {
            switch "\n\t" {
            case "":
                v11 := GenS1{F0: 357}
                _ = v11
            case "hello":
                v12 := []int{v10, 9223372036854775807, (v10 << ((255) & 63)), 347}
                _ = v12
            }
        }
    } else {
        v13 := []int{v10, -59}
        _ = v13
    }
    v14 := GenS2{F0: ("foo bar" + ("a" + "\n\t")), F1: 0}
    _ = v10
    _ = v14
    return 42
}

func main() {
    v1 := []string{}
    v2 := []int{-9223372036854775808, 42, 127, 709}
    v3 := make(map[string]string)
    var v4 uint16 = 979
    v5 := make(map[uint32]string)
    var v15 uint = genFn1()
    fmt.Println(v15)
    var v16 string = "foo bar"
    v17 := GenS2{F0: (v16 + v16), F1: 42}
    v18 := GenS3{F0: -75, F1: 18446744073709551615, F2: -9223372036854775808}
    switch ((v16 + v16) + ("hello" + ("hello" + (v16 + v16)))) {
    case "hello":
        switch v16 {
        case "\n\t":
            switch v16 {
            case "a":
                _ = len(v5)
            default:
                switch v16 {
                case "hello":
                    var v19 int = 0
                    _ = v19
                }
            }
        case "":
            var v20 int16 = -1
            _ = v20
        case "a":
            v15 -= v15
        }
    case "foo bar":
        var v21 uint = (v15 % ((v15) | 1))
        _ = v21
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        if false {
            for fi23 := 0; fi23 < 4; fi23++ {
                fmt.Println(v16)
                switch (v16 + (("hello" + v16) + v16)) {
                case "":
                    v24 := GenS4{F0: uint16(v15)}
                    _ = v24
                default:
                    fmt.Println(v4, v4, fi22)
                }
                v16 = v16
                _ = fi23
            }
        } else {
            if false {
                var v25 int16 = -1
                if (-1 > 127) {
                    var v26 uint = (^(v15))
                    _ = v26
                }
                _ = v25
            } else {
                fmt.Println(fi22, v4, v15)
                v5[2] = ("\n\t" + "hello")
                v2 = append(v2, (fi22 << ((4294967295) & 63)))
            }
        }
        var v27 int16 = -32768
        if (2 > 405) {
            var v28 float32 = 0.1
            _ = v28
        } else {
            if ((0 != 42) && false) {
                delete(v3, (v16 + (v16 + (("foo bar" + v16) + (v16 + "a")))))
                v15 += v15
                for fi29 := 0; fi29 < 10; fi29++ {
                    fmt.Println(fi22, v16)
                    var v30 string = v16
                    _ = fi29
                    _ = v30
                }
            }
            if len(v2) > 0 {
                v2[0] = (fi22 - (fi22 >> (((v15 + (v15 * 255))) & 63)))
            }
        }
        _ = fi22
        _ = v27
    }
    v31 := GenS5{F0: 18446744073709551615}
    v16 = (((("" + v16) + v16) + ((("a" + v16) + v16) + v16)) + ("foo bar" + v16))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 = (v4 >> (((v15 % ((v15) | 1))) & 63))
    if ((0.1 + (((0.0 - (1.5 - 1e10)) + 1.5) - 1e10)) != -2.25) {
        delete(v5, 98)
    }
    if false {
        for fi32 := 0; fi32 < 11; fi32++ {
            v15 += (v15 - v15)
            v33 := GenS3{F0: 0, F1: 2, F2: 0}
            _ = fi32
            _ = v33
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
