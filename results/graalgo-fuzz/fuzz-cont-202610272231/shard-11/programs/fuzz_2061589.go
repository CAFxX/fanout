package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint; F1 int; F2 string }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint64; F1 int32; F2 int64 }
type GenS5 struct { F0 int; F1 bool; F2 int32 }
func genFn0(p0 bool) bool {
    var v7 string = "\n\t"
    v8 := []int{652, 32768, -1}
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            fmt.Println(fi9, v7, fi9)
            p0 = true
            var v11 string = v7
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: p0, F1: 350}
        var v13 int16 = 1000
        _ = fi9
        _ = v12
        _ = v13
    }
    var v14 uint16 = 1
    _ = len(v8)
    _ = p0
    _ = v14
    _ = v7
    _ = v8
    return ((1 <= 32767) || (!(false || (158 <= 198))))
}

func main() {
    var v1 uint8 = 128
    var v2 bool = (!true)
    var v3 uint8 = (^(v1))
    var v4 uint64 = 4294967295
    v5 := make(map[uint32]int)
    var v6 uint = uint(v4)
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS1{F0: (v2 || (v2 || v2)), F1: 2147483647}
        if (42 < 0) {
            for fi17 := 0; fi17 < 13; fi17++ {
                v18 := GenS1{F0: (true && (9223372036854775807 != -934)), F1: 1}
                _ = fi17
                _ = v18
            }
        }
        var v19 bool = false
        _ = fi15
        _ = v16
        _ = v19
    }
    v20 := []int{0, -874, 9223372036854775807, 128}
    var v21 uint8 = v3
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 float32 = (3.14159 * 1.5)
        var v24 uint = (^(v6))
        if (164 <= 42) {
            switch fi22 {
            case 1:
                v4 = v4
            }
            switch (((("a" + ("foo bar" + "\n\t")) + "hello") + "foo bar") + ("hello" + "a")) {
            case "":
                v1 += 255
            case "foo bar":
                fmt.Println(fi22)
            default:
                if (true && v2) {
                    v25 := GenS2{F0: v24, F1: (fi22 + (fi22 | 924)), F2: (("" + ("\n\t" + "\n\t")) + "hello")}
                    for fi26 := 0; fi26 < 18; fi26++ {
                        if (v23 != 3.14159) {
                            v5[750] = (fi22 >> ((v6) & 63))
                            v27 := GenS3{F0: -512, F1: ""}
                            _ = v27
                        }
                        _ = fi26
                    }
                    v28 := v5[2147483647]
                    _ = v25
                    _ = v28
                }
            }
            fmt.Println(v23, v1, v21)
        } else {
            for fi29 := 0; fi29 < 5; fi29++ {
                fmt.Println(fi22, v2, v1)
                var v30 uint = v24
                _ = fi29
                _ = v30
            }
            fmt.Println(v3)
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    switch ("hello" + "a") {
    case "a":
        v3 = v21
    case "":
        fmt.Println(v3, v21)
    case "foo bar":
        var v31 float32 = -2.25
        _ = v31
    }
    v32 := GenS4{F0: 18446744073709551615, F1: 2147483647, F2: -859}
    for fi33 := 0; fi33 < 6; fi33++ {
        v34 := GenS5{F0: 1, F1: ((v2 || (float32((1.5 * (3.14159 + 1.5))) < 1e10)) || (((false || ((0 > 255) || (!(v2 || v2)))) || v2) || (127 > -126))), F2: -309}
        _ = fi33
        _ = v34
    }
    if true {
        if ((!v2) || (4294967295 <= 290)) {
            if (uint16(v21) != 1) {
                switch "hello" {
                case "a":
                    var v35 bool = v2
                    _ = v35
                case "\n\t":
                    v6 -= 0
                case "foo bar":
                    if true {
                        v3 += (^(v21))
                        switch -9223372036854775808 {
                        case 32768:
                            var v36 int = 127
                            _ = v36
                        default:
                            for fi37 := 0; fi37 < 15; fi37++ {
                                var v38 float64 = 0.1
                                fmt.Println(v21, fi37, v6)
                                v20 = append(v20, (fi37 << ((v6) & 63)))
                                _ = fi37
                                _ = v38
                            }
                        }
                    } else {
                        var v39 uint64 = uint64(v6)
                        v4 = 18446744073709551615
                        fmt.Println(v39, v3)
                        _ = v39
                    }
                }
                var v40 uint = (v6 - uint(v1))
                var v41 int64 = -839
                _ = v40
                _ = v41
            }
        }
        v42 := GenS2{F0: (^(v6)), F1: 170, F2: "foo bar"}
        switch "" {
        case "foo bar":
            v4 = 42
        case "":
            for fi43 := 0; fi43 < 16; fi43++ {
                for fi44 := 0; fi44 < 15; fi44++ {
                    var v45 int8 = -84
                    _ = fi44
                    _ = v45
                }
                fmt.Println(v6, v21, v3)
                v3 = (v3 / ((2) | 1))
                _ = fi43
            }
        case "a":
            fmt.Println(v2, v6, v4)
        }
        _ = v42
    }
    v6 -= uint(v3)
    fmt.Println(v2)
    var v46 bool = (uint64(v21) == uint64(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
