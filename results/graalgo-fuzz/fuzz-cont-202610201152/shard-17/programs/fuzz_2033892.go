package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int64) uint {
    fmt.Println(p0)
    p0 = 878
    var v6 string = ("hello" + ("\n\t" + ""))
    _ = p0
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 733
    var v3 float64 = float64(0.1)
    v4 := []int{2147483647, 878}
    v5 := make(map[int]int)
    v7 := GenS1{F0: 110, F1: 2147483647, F2: 32767}
    v2 -= v2
    v2 = (v2 ^ v2)
    if (!false) {
        v8 := GenS2{F0: 1}
        _ = v8
    } else {
        v9 := GenS2{F0: -9223372036854775808}
        v3 -= 0.0
        v2 += v2
        _ = v9
    }
    var v10 int = 32768
    for fi11 := 0; fi11 < 5; fi11++ {
        switch ("hello" + ("foo bar" + "hello")) {
        case "hello":
            v3 -= ((v3 * (v3 - 0.0)) / 1.0)
        }
        _ = fi11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        v2 -= (^(v2))
        if true {
            for fi13 := 0; fi13 < 2; fi13++ {
                v2 += v2
                _ = fi13
            }
            var v14 int16 = 0
            _ = v14
        } else {
            v10 -= fi12
            if ((v10 / ((-2147483648) | 1)) == fi12) {
                v2 = int16(v10)
                fmt.Println(v3)
                v10 += (v10 % ((42) | 1))
            } else {
                if true {
                    switch "\n\t" {
                    case "a":
                        if (64 > -47) {
                            for fi15 := 0; fi15 < 2; fi15++ {
                                v2 += v2
                                v10 += v10
                                var v16 float32 = 0.1
                                _ = fi15
                                _ = v16
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v17 := GenS2{F0: 1}
                            _ = v17
                        }
                    case "foo bar":
                        var v18 string = "foo bar"
                        _ = v18
                    }
                }
            }
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 uint8 = 42
                var v21 uint8 = 2
                fmt.Println(fi12, v20)
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        _ = fi12
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        if (v2 < v2) {
            for fi23 := 0; fi23 < 10; fi23++ {
                fmt.Println(v2, fi23, fi23)
                v3 = 0.0
                fmt.Println(fi23, fi22)
                _ = fi23
            }
            v10 -= 32767
        }
        if (237 != 221) {
            var v24 string = "a"
            _ = v24
        }
        _ = fi22
    }
    v10 = (v10 << ((42) & 63))
    v3 = v3
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        var v25 int16 = (v2 + (v2 + (v2 / (((v2 >> ((4294967295) & 63))) | 1))))
        switch "foo bar" {
        case "\n\t":
            for fi26 := 0; fi26 < 1; fi26++ {
                v25 -= (v2 & v2)
                var v27 string = (("a" + ("\n\t" + "hello")) + (("hello" + "") + "hello"))
                var v28 uint16 = 32767
                _ = fi26
                _ = v27
                _ = v28
            }
        case "a":
            v10 = (v10 - v10)
        case "hello":
            _ = len(v5)
        default:
            switch ((("hello" + "hello") + "hello") + "\n\t") {
            case "a":
                v2 = v2
            case "foo bar":
                v29 := GenS3{F0: 252}
                _ = v29
            case "\n\t":
                fmt.Println(v2, v2)
            default:
                var v30 int32 = 0
                _ = v30
            }
        }
        _ = v25
    } else {
        switch ("" + "foo bar") {
        case "foo bar":
            delete(v5, v10)
        }
        for fi31 := 0; fi31 < 4; fi31++ {
            fmt.Println(v3)
            _ = fi31
        }
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        for fi33 := 0; fi33 < 9; fi33++ {
            var v34 int16 = v2
            _ = fi33
            _ = v34
        }
        switch "hello" {
        case "\n\t":
            fmt.Println(v10)
        default:
            fmt.Println(v3, fi32, v2)
        }
        _ = fi32
    }
    var v35 string = (("hello" + "a") + ("\n\t" + ""))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v7
}
