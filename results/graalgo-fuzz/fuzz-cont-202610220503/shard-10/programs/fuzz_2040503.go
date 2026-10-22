package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 int16 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int32 = 2147483647
        v7 += 42
        for fi8 := 0; fi8 < 13; fi8++ {
            fmt.Println(fi8, v7)
            v9 := GenS1{F0: 128}
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 int64 = 42
    v11 := GenS2{F0: v10, F1: 42, F2: 2147483647}
    fmt.Println(v10)
    var v12 int8 = -128
    _ = v10
    _ = v11
    _ = v12
    return 42
}
func genFn1() uint16 {
    v13 := []int{1, -2147483648}
    _ = len(v13)
    var v14 uint16 = 2
    v14 += (^(v14))
    fmt.Println(v14)
    _ = v13
    _ = v14
    return 1
}

func main() {
    var v1 string = ""
    var v2 int32 = 1
    v3 := []int{128, 32768}
    v4 := []string{"foo bar", "\n\t", v1, v1}
    v5 := make(map[int]uint32)
    v5[-943] = 42
    switch (v1 + "") {
    case "hello":
        v15 := GenS1{F0: 1}
        _ = v15
    default:
        v4 = append(v4, "hello")
    }
    var v16 int8 = 100
    for fi17 := 0; fi17 < 4; fi17++ {
        v2 = v2
        var v18 int64 = -1
        _ = fi17
        _ = v18
    }
    delete(v5, 2147483647)
    if ((("a" + (v1 + v1)) + ("a" + "foo bar")) != v1) {
        if false {
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 int = -2147483648
                v21 := GenS3{F0: 42}
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 3; fi22++ {
            v23 := GenS2{F0: -9223372036854775808, F1: 1, F2: (fi22 | fi22)}
            _ = fi22
            _ = v23
        }
    } else {
        v16 = -88
        v24 := GenS2{F0: 674, F1: 255, F2: -2147483648}
        v1 = (v1 + (("a" + (((v1 + ((v1 + v1) + v1)) + ("hello" + v1)) + v1)) + (((v1 + "foo bar") + "\n\t") + v1)))
        _ = v24
    }
    if true {
        v2 = v2
        for fi25 := 0; fi25 < 9; fi25++ {
            if false {
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 int16 = 42
                    fmt.Println(fi25)
                    v16 = (v16 % ((v16) | 1))
                    _ = fi26
                    _ = v27
                }
            } else {
                var v28 bool = false
                _ = v28
            }
            v1 = v1
            _ = len(v3)
            _ = fi25
        }
    }
    delete(v5, 128)
    switch v1 {
    case "":
        fmt.Println(v2)
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "\n\t":
        for fi29 := 0; fi29 < 18; fi29++ {
            var v30 uint8 = 255
            v31 := GenS4{F0: 2}
            v32 := GenS3{F0: 4294967295}
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
    }
    _ = len(v5)
    v16 -= v16
    var v33 bool = false
    v1 = v1
    if (((v2 << ((67) & 63)) == 892) || (-2.25 < 0.1)) {
        var v34 int64 = -535
        var v35 int64 = v34
        if (v35 < (v35 << ((2) & 63))) {
            fmt.Println(v33, v1, v34)
            var v36 uint = 0
            for fi37 := 0; fi37 < 17; fi37++ {
                v38 := GenS5{F0: 65535, F1: 1000}
                v35 -= v35
                var v39 string = v1
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = v36
        }
        _ = v34
        _ = v35
    } else {
        var v40 int32 = v2
        _ = v40
    }
    v33 = v33
    if (0.1 == -2.25) {
        var v41 int64 = 42
        fmt.Println(v16)
        _ = v41
    } else {
        var v42 uint = 2
        var v43 uint64 = 4294967295
        var v44 uint32 = 910
        _ = v42
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
