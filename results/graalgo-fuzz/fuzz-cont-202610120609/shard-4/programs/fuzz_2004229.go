package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0() uint8 {
    if false {
        var v5 uint = 18446744073709551615
        v5 = (v5 - 2)
        v5 = (v5 % ((v5) | 1))
        _ = v5
    } else {
        v6 := GenS1{F0: 4294967295}
        var v7 float64 = (0.1 - 0.1)
        for fi8 := 0; fi8 < 7; fi8++ {
            fmt.Println(fi8)
            _ = fi8
        }
        _ = v6
        _ = v7
    }
    v9 := GenS2{F0: 1, F1: 32767}
    var v10 int64 = -9223372036854775808
    var v11 uint16 = 0
    _ = v10
    _ = v11
    _ = v9
    return 2
}

func main() {
    var v1 int16 = 645
    v2 := []int{1, 128, 4294967295, -734}
    v3 := []uint32{2147483648, 42, 721, 352}
    v4 := []string{}
    v1 += (v1 >> ((1) & 63))
    if false {
        for fi12 := 0; fi12 < 9; fi12++ {
            v1 -= v1
            for fi13 := 0; fi13 < 8; fi13++ {
                if len(v2) > 0 {
                    v2[0] = fi12
                }
                _ = fi13
            }
            _ = fi12
        }
        switch ("hello" + "foo bar") {
        case "foo bar":
            var v14 int16 = (v1 - 0)
            _ = v14
        default:
            fmt.Println(v1)
        }
    } else {
        var v15 uint16 = 975
        var v16 int64 = 1
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 int = fi17
        v18 = fi17
        var v19 float64 = (-2.25 / 1.5)
        _ = fi17
        _ = v18
        _ = v19
    }
    if (573 > -2147483648) {
        switch "foo bar" {
        case "\n\t":
            var v20 int = 128
            _ = v20
        case "hello":
            var v21 int8 = 37
            _ = v21
        case "a":
            _ = len(v4)
        default:
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 uint32 = 652
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(fi22, fi22, v1)
                _ = fi22
                _ = v23
            }
        }
        v1 += v1
    } else {
        for fi24 := 0; fi24 < 8; fi24++ {
            v25 := GenS2{F0: 0, F1: v1}
            if ((0.1 > 3.14159) || false) {
                if true {
                    var v26 float32 = float32(float64(0.0))
                    switch (fi24 << ((18446744073709551615) & 63)) {
                    case 4294967295:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case 0:
                        var v27 int = fi24
                        _ = v27
                    case -1:
                        v28 := GenS2{F0: 1, F1: (v1 | v1)}
                        _ = v28
                    default:
                        switch "\n\t" {
                        case "hello":
                            v1 = 32767
                        }
                    }
                    var v29 int64 = 1
                    _ = v26
                    _ = v29
                }
                v1 -= v1
                for fi30 := 0; fi30 < 6; fi30++ {
                    v1 = v1
                    _ = fi30
                }
            }
            switch -741 {
            case 32768:
                v1 += (v1 - v1)
            case 0:
                v31 := GenS2{F0: 2, F1: v1}
                _ = v31
            case 127:
                v1 -= (v1 << ((42) & 63))
            default:
                v32 := GenS3{F0: 216, F1: 1}
                _ = v32
            }
            _ = fi24
            _ = v25
        }
        switch -9223372036854775808 {
        case 1:
            if (688 >= 542) {
                var v33 uint = 65535
                switch ("a" + "") {
                case "":
                    var v34 float32 = 1.5
                    _ = v34
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "hello":
                    var v35 int8 = -21
                    _ = v35
                default:
                    v36 := GenS1{F0: 4294967295}
                    _ = v36
                }
                _ = v33
            } else {
                switch "\n\t" {
                case "a":
                    var v37 string = ("hello" + "a")
                    _ = v37
                case "":
                    switch ("\n\t" + ("hello" + "\n\t")) {
                    case "foo bar":
                        v1 -= v1
                    case "hello":
                        v2 = append(v2, -702)
                    default:
                        var v38 int64 = -1
                        _ = v38
                    }
                case "\n\t":
                    var v39 uint8 = 17
                    _ = v39
                }
                v1 -= (v1 & 32767)
            }
        }
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v40 uint16 = 42
    var v41 uint64 = 1
    var v42 int = 9223372036854775807
    v1 = -32768
    v41 = (v41 % (((v41 + v41)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
