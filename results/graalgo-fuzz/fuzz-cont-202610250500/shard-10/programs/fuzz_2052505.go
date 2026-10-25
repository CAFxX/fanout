package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int, p1 uint8, p2 int32) uint64 {
    p1 += 0
    v5 := GenS1{F0: ("foo bar" + "\n\t")}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 uint16) int64 {
    v6 := GenS1{F0: ("" + "hello")}
    var v7 bool = ((1.5 + 0.0) < ((3.14159 * -2.25) + 1.5))
    var v8 uint16 = 977
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 uint32 = 1
    var v4 int16 = -632
    fmt.Println(v4, v4)
    var v9 uint64 = 1
    v3 += v3
    if (0 <= 271) {
        v10 := []int{42}
        if (!(v9 <= (^(v9)))) {
            var v11 uint16 = 97
            if (true && (128 > 149)) {
                var v12 int = -1
                v13 := GenS1{F0: "foo bar"}
                fmt.Println(v11, v4, v9)
                _ = v12
                _ = v13
            }
            switch ("foo bar" + ("a" + ("\n\t" + "foo bar"))) {
            case "\n\t":
                if ((("" + "\n\t") + ("" + "")) >= ((("\n\t" + "hello") + (("" + ("" + "hello")) + (("" + "hello") + "hello"))) + ((("hello" + "\n\t") + ("foo bar" + "hello")) + ("a" + ("a" + "foo bar"))))) {
                    switch 4294967295 {
                    case 624:
                        v4 = v4
                    case 127:
                        var v14 uint64 = (v9 % ((42) | 1))
                        _ = v14
                    }
                }
            case "foo bar":
                var v15 uint8 = 219
                _ = v15
            case "":
                v16 := v1[v3]
                _ = v16
            default:
                var v17 float32 = 1.5
                _ = v17
            }
            _ = v11
        }
        v4 = v4
        _ = v10
    }
    var v18 int32 = 513
    v2[("foo bar" + (("hello" + "\n\t") + "hello"))] = -831
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := []int{(fi19 << ((456) & 63)), (fi19 & fi19)}
        var v21 uint32 = (v3 % ((v3) | 1))
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    v9 = (v9 << ((2) & 63))
    var v22 int32 = v18
    switch "" {
    case "a":
        if (!((^(v4)) < v4)) {
            var v23 int = 138
            for fi24 := 0; fi24 < 9; fi24++ {
                for fi25 := 0; fi25 < 10; fi25++ {
                    v26 := GenS2{F0: ("" + ("" + ("foo bar" + ("hello" + "a")))), F1: 42}
                    if (0 < 0) {
                        _ = len(v2)
                        fmt.Println(v23, v4)
                    }
                    v27 := GenS3{F0: (v22 - -2147483648), F1: 0}
                    v4 -= v4
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = fi24
            }
            if false {
                v28 := []int{v23, 1, -821}
                _ = v28
            } else {
                v4 += int16(v18)
                switch v23 {
                case 128:
                    var v29 uint = 230
                    _ = v29
                }
                switch -753 {
                case -1:
                    v30 := GenS1{F0: "a"}
                    _ = v30
                case -9223372036854775808:
                    var v31 int32 = int32(v4)
                    _ = v31
                case -2147483648:
                    switch "a" {
                    case "":
                        switch "" {
                        case "\n\t":
                            v32 := GenS4{F0: 42, F1: (v9 - v9)}
                            _ = v32
                        case "a":
                            if true {
                                v3 = 297
                            }
                        case "":
                            var v33 bool = false
                            _ = v33
                        }
                    }
                }
            }
            _ = v23
        }
    case "hello":
        _ = len(v2)
    case "":
        v18 -= (v18 >> ((42) & 63))
    default:
        for fi34 := 0; fi34 < 20; fi34++ {
            for fi35 := 0; fi35 < 7; fi35++ {
                if (false || true) {
                    v3 = 2147483647
                    fmt.Println(v9)
                    v22 += (v22 >> ((42) & 63))
                    v36 := []int{(fi34 << ((2) & 63))}
                    _ = v36
                }
                var v37 float64 = (0.0 + (float64(0.1) / 1.0))
                _ = fi35
                _ = v37
            }
            v18 += 0
            _ = fi34
        }
    }
    var v38 int16 = (v4 | v4)
    v38 = (v38 >> ((934) & 63))
    switch (("a" + ("" + "")) + "") {
    case "foo bar":
        var v39 uint = 65535
        _ = v39
    case "":
        var v40 int8 = 42
        _ = v40
    case "a":
        v4 -= (v4 >> ((42) & 63))
    }
    v41 := GenS3{F0: (v22 + v18), F1: uint16(v3)}
    switch "hello" {
    case "\n\t":
        fmt.Println(v22)
    case "foo bar":
        v38 -= v38
    }
    if (false && false) {
        var v42 int32 = (v18 | v22)
        _ = v42
    } else {
        v3 = 939
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v38
    _ = v4
    _ = v41
    _ = v9
}
