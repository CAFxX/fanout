package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int64; F1 int16; F2 uint16 }
type GenS5 struct { F0 bool; F1 uint16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 uint64 = 264
    v5 := []uint32{2147483648, 2}
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int32 = 1
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        if true {
            var v10 float32 = (3.14159 + 0.1)
            fmt.Println(v10, v4, v4)
            fmt.Println(v4, fi9)
            _ = v10
        } else {
            fmt.Println(fi9, fi9)
            if ((3.14159 + -2.25) < float32(1e10)) {
                var v11 float32 = float32(3.14159)
                _ = v11
            }
        }
        var v12 bool = true
        if (true && ((fi9 | fi9) == (fi9 | (^(fi9))))) {
            v12 = (!("" < "foo bar"))
        }
        _ = fi9
        _ = v12
    }
    v13 := GenS1{F0: "\n\t", F1: 0, F2: 964}
    fmt.Println(v4)
    v14 := GenS2{F0: ""}
    var v15 uint = 0
    fmt.Println(v15)
    v15 += v15
    switch 42 {
    case 9223372036854775807:
        v4 = v4
    case 1:
        v16 := GenS2{F0: (((("\n\t" + "a") + "foo bar") + "") + ((("hello" + (("\n\t" + "a") + ("" + ("" + (("a" + (("\n\t" + "foo bar") + "a")) + ""))))) + ("a" + "hello")) + "\n\t"))}
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        switch ("" + "foo bar") {
        case "\n\t":
            v15 = (v15 / (((v15 + uint(v4))) | 1))
        case "a":
            if true {
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 int8 = 0
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 14; fi20++ {
                    if false {
                        if ((819 != 921) && (1 > 0)) {
                            delete(v1, ((("a" + "foo bar") + "foo bar") + ("hello" + "hello")))
                        } else {
                            v21 := v2[(fi17 << (((v15 << ((v15) & 63))) & 63))]
                            var v22 uint16 = 931
                            _ = v21
                            _ = v22
                        }
                    }
                    if (-1 <= 1) {
                        var v23 uint = v15
                        if len(v5) > 0 {
                            v5[0] = 2147483647
                        }
                        v24 := GenS3{F0: 100, F1: (fi20 % ((fi20) | 1))}
                        if len(v5) > 0 {
                            v5[0] = 65535
                        }
                        _ = v23
                        _ = v24
                    } else {
                        if (0.1 != 3.14159) {
                            var v25 uint16 = 255
                            _ = len(v5)
                            v4 += 4294967295
                            _ = v25
                        }
                        v15 = 18446744073709551615
                        v26 := GenS3{F0: -43, F1: fi17}
                        var v27 int32 = -2147483648
                        var v28 string = (("" + "foo bar") + "foo bar")
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    _ = fi20
                }
            }
        }
        _ = fi17
    }
    if len(v5) > 0 {
        v5[0] = uint32(v4)
    }
    var v29 int = 916
    switch v29 {
    case 127:
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 int = (v29 / (((^(fi30))) | 1))
            var v32 int16 = 137
            v15 = v15
            _ = fi30
            _ = v31
            _ = v32
        }
    case 32768:
        v33 := GenS4{F0: 777, F1: 99, F2: 333}
        _ = v33
    case 4294967295:
        v34 := v2[(v29 | -67)]
        _ = v34
    }
    fmt.Println(v29, v15, v15)
    for fi35 := 0; fi35 < 5; fi35++ {
        var v36 string = "hello"
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 int = v29
            var v39 string = v36
            v40 := GenS5{F0: false, F1: 399}
            var v41 uint32 = 53
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
            _ = v41
        }
        var v42 bool = true
        _ = fi35
        _ = v36
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
