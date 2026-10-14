package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint32) int16 {
    if (!false) {
        if false {
            if (!false) {
                p1 += (^(p1))
                v5 := GenS1{F0: 42, F1: 0, F2: (^(p1))}
                v6 := GenS1{F0: 665, F1: 4294967295, F2: 10}
                _ = v5
                _ = v6
            }
            var v7 uint64 = 922
            var v8 uint8 = 0
            _ = v7
            _ = v8
        } else {
            if false {
                for fi9 := 0; fi9 < 14; fi9++ {
                    fmt.Println(fi9, p0, p0)
                    v10 := GenS2{F0: ((p1 / ((2) | 1)) != (^(p1))), F1: (100 == -128)}
                    switch (fi9 / (((fi9 & fi9)) | 1)) {
                    case 2147483647:
                        for fi11 := 0; fi11 < 3; fi11++ {
                            v12 := GenS3{F0: 9223372036854775807, F1: 0, F2: 32767}
                            var v13 float32 = 1.5
                            if (-1 >= -1) {
                                p1 = (p1 - 4294967295)
                            }
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    case 128:
                        fmt.Println(p0, fi9, p0)
                    case -1:
                        v14 := GenS2{F0: true, F1: (true && ((fi9 << ((18446744073709551615) & 63)) != fi9))}
                        _ = v14
                    }
                    _ = fi9
                    _ = v10
                }
            }
        }
        var v15 int8 = -50
        v16 := GenS1{F0: 28, F1: 0, F2: p1}
        _ = v15
        _ = v16
    } else {
        v17 := GenS1{F0: 65535, F1: 0, F2: p1}
        _ = v17
    }
    switch ((("hello" + ("\n\t" + ("" + "a"))) + ("" + "a")) + "hello") {
    case "a":
        fmt.Println(p1, p0)
    case "":
        fmt.Println(p0)
    case "\n\t":
        for fi18 := 0; fi18 < 7; fi18++ {
            if true {
                if false {
                    fmt.Println(p1)
                } else {
                    var v19 int = fi18
                    _ = v19
                }
            }
            for fi20 := 0; fi20 < 15; fi20++ {
                v21 := GenS4{F0: 0}
                p1 = 0
                _ = fi20
                _ = v21
            }
            _ = fi18
        }
    default:
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := []int{127, fi22, (fi22 & (fi22 >> ((434) & 63))), (fi22 - fi22)}
            var v24 uint = 2
            for fi25 := 0; fi25 < 13; fi25++ {
                var v26 int = 0
                _ = fi25
                _ = v26
            }
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    p1 = 673
    var v27 uint8 = uint8(p1)
    _ = p0
    _ = p1
    _ = v27
    return -1
}

func main() {
    var v1 float64 = (-2.25 * float64((float32((((3.14159 + 0.1) - (0.1 * 0.0)) - -2.25)) * 0.1)))
    v2 := make(map[string]bool)
    var v3 uint64 = 1
    v4 := []string{}
    switch (("foo bar" + "foo bar") + ("hello" + "a")) {
    case "\n\t":
        if ((!(true && (101 > -1))) && ((426 == 42) && (v3 < v3))) {
            v28 := GenS4{F0: v3}
            var v29 int = 32767
            fmt.Println(v1, v1)
            _ = v28
            _ = v29
        } else {
            for fi30 := 0; fi30 < 10; fi30++ {
                v2["\n\t"] = (!((!(!(-1 >= 2147483647))) || false))
                _ = fi30
            }
            if len(v4) > 0 {
                v4[0] = "a"
            }
            var v31 uint8 = 1
            _ = v31
        }
    case "foo bar":
        fmt.Println(v3, v1)
    case "hello":
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 int64 = 46
            v34 := GenS2{F0: (false && false), F1: true}
            switch "" {
            case "":
                switch (("foo bar" + "foo bar") + "") {
                case "hello":
                    for fi35 := 0; fi35 < 12; fi35++ {
                        var v36 uint = 0
                        var v37 string = "a"
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                case "a":
                    fmt.Println(fi32, v3, fi32)
                }
            }
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    var v38 int = -812
    var v39 uint64 = v3
    v39 += (v3 - 2)
    fmt.Println(v3, v3)
    v39 -= 839
    v39 -= v39
    var v40 bool = (!false)
    for fi41 := 0; fi41 < 19; fi41++ {
        v4 = append(v4, ("" + "foo bar"))
        _ = fi41
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
