package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int64 }
func genFn0(p0 int32) bool {
    var v7 uint = 2
    var v8 int64 = int64(v7)
    var v9 float64 = 1e10
    p0 -= (-(p0))
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (!(!(true && ((608 != p0) && false))))
}
func genFn1(p0 string, p1 int, p2 uint8) uint {
    var v10 uint32 = 2147483647
    v11 := []int{p1, (p1 | -759), int(v10)}
    p0 = ("a" + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 0
}

func main() {
    var v1 string = ("\n\t" + "foo bar")
    v2 := []int{-9223372036854775808}
    v3 := []uint32{562, 65535}
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 0
    v6 := make(map[uint32]int)
    fmt.Println(v5, v5)
    var v12 int64 = 157
    if false {
        v13 := GenS1{F0: 978, F1: 2, F2: 2}
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 int = (fi14 * -468)
            var v16 int8 = 100
            v16 -= v16
            _ = fi14
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            switch (-(fi17)) {
            case 32767:
                _ = len(v4)
            case -9223372036854775808:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 42:
                for fi18 := 0; fi18 < 2; fi18++ {
                    v5 += (v5 - v5)
                    fmt.Println(fi18)
                    _ = fi18
                }
            }
            _ = fi17
        }
        _ = v13
    } else {
        var v19 int = 418
        var v20 uint16 = 1
        _ = v19
        _ = v20
    }
    switch 4294967295 {
    case 128:
        var v21 bool = ((-(v12)) > v12)
        _ = v21
    case -69:
        v3 = append(v3, 65535)
    default:
        var v22 uint16 = 65535
        _ = v22
    }
    v23 := GenS2{F0: 524}
    if false {
        v12 = (v12 + (^(v12)))
        if ((3.14159 != 1e10) || (false || (true && true))) {
            var v24 int = 122
            for fi25 := 0; fi25 < 19; fi25++ {
                fmt.Println(v24, v24)
                _ = fi25
            }
            _ = v24
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v26 := GenS3{F0: "hello", F1: (v12 >> ((4294967295) & 63))}
    if (-1 > v12) {
        if true {
            var v27 int32 = 696
            v28 := GenS3{F0: v1, F1: v12}
            _ = v27
            _ = v28
        } else {
            fmt.Println(v1, v5)
        }
    }
    fmt.Println(v1, v1, v5)
    for fi29 := 0; fi29 < 3; fi29++ {
        var v30 uint = 0
        _ = fi29
        _ = v30
    }
    _ = len(v6)
    v3 = append(v3, 597)
    if (1000 <= 1000) {
        v1 = (v1 + v1)
    }
    v31 := GenS3{F0: v1, F1: v12}
    for fi32 := 0; fi32 < 18; fi32++ {
        fmt.Println(fi32, v1)
        switch v1 {
        case "":
            v6[65535] = 837
        }
        _ = fi32
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        for fi34 := 0; fi34 < 9; fi34++ {
            v6[4294967295] = fi33
            switch ((v1 + ("\n\t" + v1)) + ("" + ("\n\t" + ((v1 + "hello") + "foo bar")))) {
            case "foo bar":
                if ((0 != -2147483648) && ((!true) && true)) {
                    fmt.Println(v12, v5)
                    v12 = int64(fi33)
                } else {
                    var v35 uint16 = 32767
                    delete(v4, 42)
                    delete(v6, 1)
                    _ = v35
                }
            case "hello":
                v1 = ("" + "hello")
            case "\n\t":
                switch (-(fi33)) {
                case -197:
                    v2 = append(v2, (fi34 >> ((4294967295) & 63)))
                case -713:
                    var v36 float32 = 0.1
                    _ = v36
                case 1:
                    var v37 int = fi34
                    _ = v37
                default:
                    var v38 int32 = 2147483647
                    _ = v38
                }
            }
            _ = fi34
        }
        _ = fi33
    }
    v1 = ((("foo bar" + (((("a" + (v1 + "foo bar")) + ((v1 + v1) + ((("hello" + "foo bar") + v1) + (v1 + v1)))) + v1) + v1)) + (v1 + (((v1 + v1) + "\n\t") + v1))) + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
