package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int }
type GenS4 struct { F0 int16; F1 int16; F2 uint8 }

func main() {
    var v1 uint16 = 40
    var v2 uint8 = 255
    v3 := make(map[string]uint32)
    var v4 int32 = -704
    var v5 float32 = (3.14159 - 0.0)
    var v6 string = "foo bar"
    for fi7 := 0; fi7 < 2; fi7++ {
        switch "foo bar" {
        case "a":
            for fi8 := 0; fi8 < 17; fi8++ {
                v4 = 2147483647
                v4 = (v4 & (^(v4)))
                _ = fi8
            }
        }
        fmt.Println(v6)
        _ = fi7
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int = 580
        _ = fi9
        _ = v10
    }
    fmt.Println(v5, v4)
    var v11 int = -9223372036854775808
    fmt.Println(v1, v4)
    var v12 int8 = 42
    var v13 float64 = (-2.25 - (-2.25 - (0.1 / float64(v5))))
    var v14 int16 = 1
    delete(v3, "")
    v5 += v5
    if true {
        var v15 uint32 = 2
        v16 := GenS1{F0: v6}
        _ = v15
        _ = v16
    }
    var v17 bool = true
    switch v11 {
    case 4294967295:
        fmt.Println(v4)
    default:
        switch (v11 % ((4294967295) | 1)) {
        case -9223372036854775808:
            v18 := GenS2{F0: (^(v11)), F1: 0}
            _ = v18
        case -2147483648:
            _ = len(v3)
        case 2147483647:
            v19 := []int{(v11 << ((1) & 63))}
            _ = v19
        }
    }
    v20 := v3["\n\t"]
    for fi21 := 0; fi21 < 11; fi21++ {
        for fi22 := 0; fi22 < 20; fi22++ {
            if v17 {
                v23 := GenS1{F0: v6}
                _ = v23
            }
            var v24 float64 = v13
            if true {
                v12 = (v12 - (v12 % ((100) | 1)))
                var v25 int32 = v4
                v26 := []int{}
                v27 := GenS1{F0: "hello"}
                _ = v25
                _ = v26
                _ = v27
            }
            _ = fi22
            _ = v24
        }
        var v28 uint16 = v1
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 int = fi21
            switch ((v6 + "\n\t") + ((((((("\n\t" + v6) + "foo bar") + ("\n\t" + ((v6 + v6) + ("\n\t" + v6)))) + "hello") + v6) + v6) + v6)) {
            case "hello":
                switch v6 {
                case "":
                    fmt.Println(v6, v17, v4)
                case "a":
                    switch fi29 {
                    case -2147483648:
                        v31 := []int{(fi29 - (fi21 / (((-(fi29))) | 1))), 128}
                        _ = v31
                    case -1:
                        var v32 uint32 = v20
                        v33 := GenS3{F0: 288, F1: 727, F2: v30}
                        v34 := GenS3{F0: v1, F1: 858, F2: (v11 / (((v30 - (fi29 << ((18446744073709551615) & 63)))) | 1))}
                        _ = v32
                        _ = v33
                        _ = v34
                    case 42:
                        v20 = (v20 / ((uint32(v1)) | 1))
                    default:
                        var v35 int16 = v14
                        _ = v35
                    }
                }
            case "a":
                v36 := GenS3{F0: v1, F1: 4294967295, F2: int(v2)}
                _ = v36
            case "foo bar":
                fmt.Println(v12)
            default:
                if (v17 || v17) {
                    v37 := GenS2{F0: (v11 << ((760) & 63)), F1: 0}
                    v2 += v2
                    v4 -= (v4 >> ((121) & 63))
                    _ = v37
                } else {
                    v38 := GenS3{F0: 2, F1: 4294967295, F2: fi29}
                    var v39 int = (^(v30))
                    _ = v38
                    _ = v39
                }
            }
            switch fi29 {
            case 1:
                v28 += (v28 << ((1) & 63))
            case -1:
                v40 := GenS4{F0: v14, F1: (v14 ^ (v14 * v14)), F2: (v2 + (v2 % ((127) | 1)))}
                _ = v40
            }
            _ = fi29
            _ = v30
        }
        _ = fi21
        _ = v28
    }
    fmt.Println(v13, v20, v20)
    v14 -= 0
    var v41 bool = (!v17)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
