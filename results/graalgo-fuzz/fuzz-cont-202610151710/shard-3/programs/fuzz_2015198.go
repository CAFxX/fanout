package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int; F1 int16; F2 int }
func genFn0() uint16 {
    if true {
        switch (("hello" + "") + "a") {
        case "foo bar":
            var v7 int32 = -2147483648
            _ = v7
        case "":
            var v8 uint = 255
            _ = v8
        case "\n\t":
            v9 := GenS1{F0: 2, F1: 63}
            _ = v9
        default:
            var v10 uint32 = 42
            _ = v10
        }
        for fi11 := 0; fi11 < 16; fi11++ {
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 uint64 = 2
                v14 := GenS2{F0: 404, F1: 127, F2: -128}
                if false {
                    var v15 float32 = (1e10 / float32((1e10 - (0.1 / (float64((0.1 / 1e10)) + (0.1 + (float64(0.0) - ((3.14159 - 1e10) / 1.5))))))))
                    _ = v15
                } else {
                    fmt.Println(fi11, fi11, fi11)
                    v16 := GenS3{F0: 65535, F1: (v13 / (((v13 | v13)) | 1))}
                    switch fi12 {
                    case 351:
                        var v17 string = "\n\t"
                        _ = v17
                    case 1:
                        if (!((237 <= 0) && true)) {
                            v18 := []int{127}
                            _ = v18
                        }
                    }
                    _ = v16
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            var v19 string = "a"
            v19 = ("a" + v19)
            _ = fi11
            _ = v19
        }
    }
    var v20 uint64 = 864
    _ = v20
    return 255
}

func main() {
    var v1 int16 = -32768
    v2 := []int{}
    var v3 bool = true
    var v4 uint = 444
    var v5 float32 = (3.14159 * (float32(1.5) / (1e10 - (3.14159 * 3.14159))))
    var v6 uint64 = 4294967295
    for fi21 := 0; fi21 < 18; fi21++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v22 int64 = 825
        v5 += v5
        _ = fi21
        _ = v22
    }
    if (v3 && v3) {
        fmt.Println(v3, v3)
    } else {
        v2 = append(v2, -9223372036854775808)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v23 uint = (^(v4))
        _ = v23
    }
    v1 += v1
    if (1 != -2147483648) {
        v24 := GenS1{F0: 68, F1: -1}
        var v25 uint32 = 2147483647
        _ = v24
        _ = v25
    }
    switch "hello" {
    case "a":
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 float64 = 1.5
            for fi28 := 0; fi28 < 14; fi28++ {
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    var v29 int = fi26
                    _ = v29
                default:
                    v6 = 256
                }
                _ = fi28
            }
            var v30 bool = true
            _ = fi26
            _ = v27
            _ = v30
        }
    case "foo bar":
        var v31 uint8 = 255
        _ = v31
    case "\n\t":
        var v32 int = -9223372036854775808
        _ = v32
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            if len(v2) > 0 {
                v2[0] = (^(fi33))
            }
            switch "a" {
            case "":
                var v34 uint64 = (^(v6))
                _ = v34
            case "\n\t":
                for fi35 := 0; fi35 < 3; fi35++ {
                    var v36 string = "a"
                    var v37 int32 = 771
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            case "hello":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v38 := GenS3{F0: 2, F1: 2}
            _ = v38
        } else {
            _ = len(v2)
        }
        _ = fi33
    }
    v3 = false
    v39 := GenS4{F0: -2147483648, F1: v1, F2: 545}
    v40 := GenS3{F0: 42, F1: (v6 ^ (v6 % ((uint64(v4)) | 1)))}
    var v41 int32 = 0
    v42 := GenS3{F0: 65535, F1: (v6 & (v6 - v6))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
