package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint32) int64 {
    if false {
        var v7 uint8 = 2
        switch 9223372036854775807 {
        case 32767:
            p0 -= (p0 % (((p0 >> ((743) & 63))) | 1))
        case 128:
            if (384 == 1000) {
                if (0.0 != (1e10 * 0.1)) {
                    v8 := GenS1{F0: -128, F1: 0, F2: false}
                    var v9 string = "a"
                    v7 -= v7
                    _ = v8
                    _ = v9
                }
                p0 -= 4294967295
            }
        case -9223372036854775808:
            if false {
                fmt.Println(v7, p0)
                v10 := GenS1{F0: 0, F1: 2, F2: true}
                _ = v10
            }
        }
        switch ((((((("hello" + "\n\t") + ("" + "\n\t")) + "foo bar") + ((("" + ("\n\t" + "")) + "foo bar") + "a")) + (("a" + "\n\t") + "")) + "") + ((("a" + ("hello" + "foo bar")) + ((("foo bar" + ("hello" + "foo bar")) + "hello") + ("a" + "a"))) + ((("foo bar" + ("a" + "a")) + "hello") + "\n\t"))) {
        case "hello":
            fmt.Println(p0, v7)
        case "\n\t":
            var v11 int64 = -661
            _ = v11
        }
        _ = v7
    }
    fmt.Println(p0)
    p0 = 1
    _ = p0
    return -672
}
func genFn1(p0 bool, p1 string, p2 bool) int {
    if (!p0) {
        v12 := []int{0, 2147483647, -1}
        _ = v12
    }
    v13 := GenS2{F0: ((1e10 - (((float32(((0.0 / 1.0) + 0.1)) * 0.0) / ((-2.25 / 1.5) + (1e10 / 3.14159))) - float32(1.5))) == 0.0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    var v4 bool = (255 > 42)
    var v5 int64 = 0
    var v6 int8 = -128
    switch ("hello" + "a") {
    case "\n\t":
        for fi14 := 0; fi14 < 8; fi14++ {
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 int = 32767
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    default:
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int8 = v6
            _ = fi17
            _ = v18
        }
    }
    fmt.Println(v5, v6)
    for fi19 := 0; fi19 < 7; fi19++ {
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := GenS3{F0: v6, F1: 42}
            v22 := GenS2{F0: false}
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 int8 = v6
        v4 = false
        _ = fi19
        _ = v23
    }
    var v24 float32 = 0.0
    v4 = v4
    var v25 int16 = 0
    for fi26 := 0; fi26 < 6; fi26++ {
        for fi27 := 0; fi27 < 13; fi27++ {
            v24 = -2.25
            if (!v4) {
                var v28 int16 = (v25 % ((v25) | 1))
                var v29 bool = ("\n\t" == (("a" + "hello") + ("hello" + ("" + ""))))
                var v30 uint64 = 379
                if ((v5 / ((v5) | 1)) == (v5 - genFn0(852))) {
                    fmt.Println(v30, fi26)
                }
                _ = v28
                _ = v29
                _ = v30
            }
            fmt.Println(v6, v6, fi26)
            _ = fi27
        }
        _ = fi26
    }
    if true {
        v31 := GenS1{F0: v6, F1: 742, F2: v4}
        var v32 uint32 = 890
        _ = v31
        _ = v32
    }
    v24 -= v24
    v33 := []int{}
    v34 := v3[792]
    for fi35 := 0; fi35 < 9; fi35++ {
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        var v36 bool = false
        switch "\n\t" {
        case "hello":
            for fi37 := 0; fi37 < 6; fi37++ {
                var v38 int = -272
                fmt.Println(fi37, v38, fi35)
                v39 := GenS4{F0: 32767, F1: (^(v25))}
                if true {
                    v40 := GenS1{F0: -114, F1: 719, F2: v4}
                    var v41 string = "hello"
                    var v42 int32 = -687
                    _ = len(v33)
                    var v43 bool = v36
                    _ = v40
                    _ = v41
                    _ = v42
                    _ = v43
                }
                _ = fi37
                _ = v38
                _ = v39
            }
        default:
            var v44 int32 = -2147483648
            _ = v44
        }
        _ = fi35
        _ = v36
    }
    v2[2147483647] = (("hello" + "foo bar") + "")
    v45 := GenS5{F0: (v25 / ((1) | 1))}
    var v46 uint = 842
    if len(v33) > 0 {
        v33[0] = v34
    }
    v34 += v34
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
