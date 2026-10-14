package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 uint16; F2 bool }
func genFn0(p0 uint64, p1 int, p2 int16) int32 {
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := []int{fi6, 1, (^(fi5))}
            p2 -= (p2 << ((740) & 63))
            var v8 int16 = (p2 ^ p2)
            _ = fi6
            _ = v7
            _ = v8
        }
        switch (("a" + "foo bar") + "hello") {
        case "foo bar":
            var v9 int16 = p2
            _ = v9
        case "a":
            switch "foo bar" {
            case "foo bar":
                v10 := GenS1{F0: 0}
                _ = v10
            case "a":
                v11 := []int{(-(fi5)), p1, p1}
                _ = v11
            }
        }
        var v12 uint8 = 76
        _ = fi5
        _ = v12
    }
    p1 -= (-(p1))
    _ = p0
    _ = p1
    _ = p2
    return 231
}
func genFn1(p0 int32, p1 int16, p2 int32) bool {
    v13 := GenS1{F0: 4294967295}
    if (false && (false || false)) {
        v14 := []int{-9223372036854775808, 9223372036854775807}
        _ = v14
    } else {
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return true
}

func main() {
    var v1 int = 9223372036854775807
    var v2 bool = true
    var v3 string = "a"
    v4 := []int{(^(v1))}
    v1 += (v1 ^ v1)
    fmt.Println(v3, v1)
    v1 = v1
    switch ("hello" + "a") {
    case "foo bar":
        fmt.Println(v3)
    default:
        v3 = v3
    }
    switch (v1 + v1) {
    case -2147483648:
        v2 = (!(0.0 <= 3.14159))
    case -9223372036854775808:
        for fi15 := 0; fi15 < 6; fi15++ {
            if len(v4) > 0 {
                v4[0] = fi15
            }
            v16 := GenS1{F0: 803}
            switch v3 {
            case "\n\t":
                v3 = v3
            case "":
                switch fi15 {
                case 4294967295:
                    if true {
                        v17 := GenS1{F0: 444}
                        _ = v17
                    }
                case -886:
                    v3 = (v3 + v3)
                case 128:
                    for fi18 := 0; fi18 < 4; fi18++ {
                        v19 := GenS1{F0: 0}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v2)
                        _ = fi18
                        _ = v19
                    }
                }
            case "foo bar":
                if len(v4) > 0 {
                    v4[0] = v1
                }
            }
            _ = fi15
            _ = v16
        }
    }
    v3 = "hello"
    if (!v2) {
        var v20 int16 = 1
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 int8 = 1
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v20 = 32767
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 int = (^(v1))
            _ = fi23
            _ = v24
        }
        _ = v20
    } else {
        v25 := GenS2{F0: false}
        v2 = v2
        _ = v25
    }
    var v26 uint8 = 255
    fmt.Println(v1)
    v27 := GenS3{F0: 1}
    for fi28 := 0; fi28 < 2; fi28++ {
        for fi29 := 0; fi29 < 17; fi29++ {
            v30 := GenS4{F0: v2, F1: 131, F2: ((!v2) && v2)}
            switch v3 {
            case "hello":
                if v2 {
                    fmt.Println(v26)
                }
            case "a":
                if ((128 != 128) || v2) {
                    v3 = "foo bar"
                    var v31 int64 = 0
                    _ = v31
                }
            default:
                v32 := GenS4{F0: v2, F1: 255, F2: (116 != 32767)}
                _ = v32
            }
            if (2 != 18446744073709551615) {
                _ = len(v4)
            }
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
