package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint8; F2 string }
type GenS3 struct { F0 int64; F1 int; F2 uint8 }
func genFn0(p0 int64, p1 uint, p2 uint16) int16 {
    p1 += 42
    v5 := []int{4294967295, 1, 32767}
    if (true || false) {
        var v6 float32 = 1e10
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint32 = 913
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    var v9 int = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v9
    return -32768
}
func genFn1(p0 int) int64 {
    switch (("" + ("hello" + "hello")) + "\n\t") {
    case "a":
        p0 += (p0 >> ((0) & 63))
    case "hello":
        p0 = (p0 / ((1) | 1))
    default:
        v10 := []int{32767, p0, -650}
        _ = v10
    }
    var v11 bool = false
    _ = p0
    _ = v11
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    var v3 int64 = 0
    v4 := []int{-1, -9223372036854775808}
    switch ((("hello" + "\n\t") + "") + (("\n\t" + ("\n\t" + "")) + "a")) {
    case "a":
        delete(v1, 374)
    case "foo bar":
        v2[325] = 108
    case "\n\t":
        var v12 uint = 0
        _ = v12
    }
    var v13 uint = 255
    if true {
        v14 := GenS1{F0: v13}
        for fi15 := 0; fi15 < 6; fi15++ {
            if false {
                var v16 uint8 = 255
                fmt.Println(fi15, v3)
                _ = v16
            } else {
                v17 := GenS1{F0: v13}
                v18 := GenS1{F0: v13}
                _ = v17
                _ = v18
            }
            _ = fi15
        }
        for fi19 := 0; fi19 < 1; fi19++ {
            for fi20 := 0; fi20 < 5; fi20++ {
                v21 := GenS2{F0: 962, F1: 2, F2: "\n\t"}
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        _ = v14
    }
    v13 += v13
    _ = len(v4)
    if (433 <= (v13 * 42)) {
        v13 = v13
        var v22 float64 = (1e10 / 1.5)
        switch 32768 {
        case 0:
            if false {
                var v23 bool = (1000 == 42)
                for fi24 := 0; fi24 < 8; fi24++ {
                    if (!v23) {
                        v25 := GenS3{F0: v3, F1: fi24, F2: 2}
                        var v26 float64 = (v22 / v22)
                        var v27 float64 = ((-2.25 / float64(-2.25)) / (v26 + v22))
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        var v28 uint64 = 1
                        for fi29 := 0; fi29 < 2; fi29++ {
                            v23 = v23
                            _ = fi29
                        }
                        var v30 int32 = -835
                        _ = v28
                        _ = v30
                    }
                    _ = fi24
                }
                if (v23 || false) {
                    v23 = v23
                    v22 += v22
                }
                _ = v23
            }
        case 2147483647:
            var v31 int16 = 0
            _ = v31
        case -380:
            fmt.Println(v3, v13)
        }
        _ = v22
    } else {
        switch ("hello" + ("a" + "foo bar")) {
        case "a":
            for fi32 := 0; fi32 < 1; fi32++ {
                var v33 uint64 = 207
                _ = fi32
                _ = v33
            }
        }
        var v34 float64 = float64(float32(1.5))
        fmt.Println(v34, v3, v3)
        _ = v34
    }
    fmt.Println(v13)
    v3 = genFn1(2147483647)
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    if (("foo bar" + ("" + "hello")) == ("" + "a")) {
        switch -1 {
        case -725:
            for fi35 := 0; fi35 < 6; fi35++ {
                v13 += (v13 ^ v13)
                _ = fi35
            }
        }
    }
    if (-36 != 1) {
        v1[316] = (731 != 523)
        switch "" {
        case "":
            for fi36 := 0; fi36 < 7; fi36++ {
                var v37 uint16 = 465
                v38 := GenS2{F0: 174, F1: 101, F2: ""}
                _ = fi36
                _ = v37
                _ = v38
            }
        }
        fmt.Println(v13)
    } else {
        fmt.Println(v13, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
