package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS2 struct { F0 int; F1 uint64; F2 uint }
type GenS3 struct { F0 int16; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint }
type GenS5 struct { F0 int32; F1 uint16 }
func genFn0(p0 int16, p1 int32) uint16 {
    v6 := GenS1{F0: 184, F1: false, F2: 2}
    v7 := GenS1{F0: 0, F1: (1.5 > float64((0.1 - (((-2.25 / 1e10) / 1.0) * 3.14159)))), F2: 29}
    v8 := []int{int(p1), 2147483647, 4294967295}
    var v9 uint16 = 0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (v9 * 255)
}
func genFn1(p0 int16, p1 int64) int16 {
    p0 = p0
    if (!true) {
        p0 -= p0
    }
    var v10 float32 = (0.0 / 0.1)
    _ = p0
    _ = p1
    _ = v10
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    v5 := []string{("\n\t" + ""), ("hello" + "")}
    v11 := v4[1]
    if (!(18446744073709551615 < 18446744073709551615)) {
        var v12 float64 = 0.1
        _ = v12
    } else {
        v13 := GenS1{F0: 789, F1: false, F2: 747}
        if (!(false && true)) {
            var v14 int = v11
            var v15 uint64 = 42
            var v16 uint8 = 42
            _ = v14
            _ = v15
            _ = v16
        } else {
            if ((!false) && (-1 != 42)) {
                v11 += v11
                var v17 int = (v11 >> ((455) & 63))
                v18 := GenS1{F0: 1, F1: true, F2: 42}
                _ = v17
                _ = v18
            }
        }
        if false {
            for fi19 := 0; fi19 < 9; fi19++ {
                if false {
                    v20 := GenS1{F0: 42, F1: (!((!true) && true)), F2: 65535}
                    v21 := GenS2{F0: 1, F1: 4294967295, F2: 2}
                    v11 = fi19
                    _ = v20
                    _ = v21
                }
                if (fi19 < (v11 - fi19)) {
                    fmt.Println(fi19, v11)
                    fmt.Println(fi19, v11)
                    v11 = (fi19 & 4294967295)
                } else {
                    v5 = append(v5, "foo bar")
                    fmt.Println(v11, v11)
                    var v22 int16 = 651
                    _ = v22
                }
                _ = fi19
            }
            if true {
                v23 := GenS3{F0: -991, F1: 0, F2: 42}
                v24 := GenS4{F0: 1, F1: 1, F2: 18446744073709551615}
                fmt.Println(v11)
                _ = v23
                _ = v24
            } else {
                var v25 uint64 = 0
                _ = len(v2)
                _ = len(v4)
                _ = v25
            }
            fmt.Println(v11)
        }
        _ = v13
    }
    switch "a" {
    case "foo bar":
        var v26 int8 = -67
        _ = v26
    }
    var v27 uint8 = 128
    if (-1 > 0) {
        v11 += v11
        switch (v11 ^ (v11 >> ((864) & 63))) {
        case 4294967295:
            for fi28 := 0; fi28 < 9; fi28++ {
                v29 := GenS2{F0: (fi28 << ((42) & 63)), F1: 0, F2: 281}
                if true {
                    v27 = v27
                    fmt.Println(v11, v27, fi28)
                    v30 := GenS5{F0: 2147483647, F1: 844}
                    _ = v30
                } else {
                    var v31 uint32 = 841
                    _ = v31
                }
                var v32 uint16 = 255
                _ = fi28
                _ = v29
                _ = v32
            }
        case 2147483647:
            switch "a" {
            case "\n\t":
                if ((1 == 100) || true) {
                    for fi33 := 0; fi33 < 6; fi33++ {
                        v34 := GenS3{F0: 32767, F1: 766, F2: 27}
                        _ = fi33
                        _ = v34
                    }
                } else {
                    v11 -= v11
                    var v35 int16 = 408
                    var v36 uint64 = 2
                    _ = v35
                    _ = v36
                }
            case "":
                v37 := v3[0]
                _ = v37
            case "hello":
                switch v11 {
                case 32768:
                    v27 += v27
                }
            }
        }
    }
    _ = len(v4)
    var v38 int8 = -29
    var v39 bool = true
    var v40 int8 = 100
    switch ((("hello" + "a") + "") + "") {
    case "\n\t":
        _ = len(v3)
    case "hello":
        v38 = v40
    default:
        v40 = (v38 % ((1) | 1))
    }
    if v39 {
        v39 = (v27 < (v27 >> ((1) & 63)))
    } else {
        switch ("foo bar" + (("a" + "foo bar") + ("" + (((("a" + "") + "hello") + "hello") + ("a" + "foo bar"))))) {
        case "foo bar":
            if (!(65535 == 439)) {
                v40 = 1
            } else {
                v38 = 100
                v39 = v39
            }
        default:
            var v41 uint16 = 0
            _ = v41
        }
        var v42 float32 = -2.25
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
