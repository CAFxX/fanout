package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint = 46
    var v2 uint8 = 40
    var v3 int16 = 1
    v4 := []int{}
    var v5 int = 42
    v6 := []string{"hello", "\n\t", ("a" + "hello"), "a"}
    v5 += 127
    var v7 uint16 = 255
    var v8 uint32 = 718
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (!true) {
        v2 -= 130
    } else {
        if len(v4) > 0 {
            v4[0] = v5
        }
        if false {
            switch "\n\t" {
            case "a":
                for fi9 := 0; fi9 < 8; fi9++ {
                    var v10 float32 = (0.1 / (-2.25 * 1e10))
                    var v11 int64 = -9223372036854775808
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case "":
                for fi12 := 0; fi12 < 10; fi12++ {
                    v13 := GenS1{F0: 0}
                    v7 -= 42
                    _ = fi12
                    _ = v13
                }
            case "hello":
                if len(v4) > 0 {
                    v4[0] = v5
                }
            }
            v3 = (v3 | v3)
        } else {
            var v14 uint = v1
            v4 = append(v4, 787)
            for fi15 := 0; fi15 < 18; fi15++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v16 := GenS2{F0: 2147483647}
                var v17 int64 = -1
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        }
    }
    v2 = v2
    if len(v4) > 0 {
        v4[0] = (v5 >> (((^(v1))) & 63))
    }
    switch (^(v5)) {
    case 127:
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS3{F0: v1}
            var v20 uint8 = 59
            _ = fi18
            _ = v19
            _ = v20
        }
    case -2147483648:
        v8 -= (v8 % ((2147483647) | 1))
    case 128:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v3 = (v3 | (-(v3)))
    var v21 int = 879
    switch "hello" {
    case "a":
        for fi22 := 0; fi22 < 19; fi22++ {
            for fi23 := 0; fi23 < 5; fi23++ {
                fmt.Println(fi22, fi23)
                _ = fi23
            }
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 float64 = (0.0 + 1e10)
                switch v21 {
                case 70:
                    fmt.Println(v8, fi24)
                default:
                    fmt.Println(v1, fi22)
                }
                _ = fi24
                _ = v25
            }
            _ = fi22
        }
    case "\n\t":
        v26 := GenS4{F0: 1}
        _ = v26
    case "":
        v2 = (v2 % ((v2) | 1))
    }
    if (false || (42 < int32(v7))) {
        switch ("foo bar" + "") {
        case "hello":
            fmt.Println(v8, v21)
        case "":
            var v27 uint = (v1 >> ((v1) & 63))
            _ = v27
        default:
            for fi28 := 0; fi28 < 12; fi28++ {
                fmt.Println(v2, v1, v5)
                switch "" {
                case "foo bar":
                    v5 = int(v2)
                case "hello":
                    v29 := GenS2{F0: uint32(v2)}
                    _ = v29
                }
                _ = fi28
            }
        }
        v30 := GenS4{F0: -1}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
