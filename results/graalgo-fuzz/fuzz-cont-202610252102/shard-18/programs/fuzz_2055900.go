package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }

func main() {
    var v1 string = ("a" + "")
    var v2 int32 = -2147483648
    var v3 int64 = -1
    var v4 float64 = (3.14159 + 1.5)
    v4 = (1e10 - (1e10 - (v4 - (v4 - v4))))
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 uint64 = 0
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{(fi7 / ((272) | 1)), 0, (fi7 ^ -1), fi7}
        switch fi7 {
        case 128:
            v9 := GenS1{F0: (v1 + (("a" + v1) + "")), F1: 65535, F2: 933}
            _ = v9
        case -450:
            if false {
                var v10 string = (v1 + "foo bar")
                v8 = append(v8, (fi7 >> ((0) & 63)))
                var v11 int = -9223372036854775808
                _ = v10
                _ = v11
            }
        case -922:
            switch (v1 + (v1 + v1)) {
            case "":
                var v12 float32 = 0.1
                _ = v12
            case "\n\t":
                for fi13 := 0; fi13 < 4; fi13++ {
                    v3 = (v3 >> ((385) & 63))
                    _ = fi13
                }
            }
        default:
            v3 += (v3 & v3)
        }
        var v14 string = v1
        _ = fi7
        _ = v14
        _ = v8
    }
    if true {
        for fi15 := 0; fi15 < 3; fi15++ {
            if (true || false) {
                v2 = (v2 >> ((18446744073709551615) & 63))
                switch v1 {
                case "hello":
                    switch v1 {
                    case "":
                        v16 := []int{(fi15 * fi15), (fi15 >> ((65535) & 63)), (fi15 + (fi15 << ((417) & 63))), -9223372036854775808}
                        _ = v16
                    case "a":
                        v17 := GenS1{F0: v1, F1: 2, F2: 65535}
                        _ = v17
                    case "hello":
                        var v18 uint32 = 2
                        _ = v18
                    default:
                        v1 = v1
                    }
                case "a":
                    var v19 int64 = (^(v3))
                    _ = v19
                }
                v1 = ((("a" + v1) + (v1 + v1)) + "a")
            } else {
                v20 := []int{(fi15 % ((32768) | 1)), 4294967295, (fi15 * fi15)}
                _ = v20
            }
            _ = fi15
        }
        v1 = "foo bar"
        if (true || (!(!false))) {
            for fi21 := 0; fi21 < 17; fi21++ {
                if false {
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v23 := []int{}
                        v1 = v1
                        var v24 int = fi22
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 6; fi25++ {
                        fmt.Println(fi25)
                        v26 := []int{}
                        var v27 uint16 = 737
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    fmt.Println(v3, v3, fi21)
                } else {
                    var v28 float32 = float32(v4)
                    _ = v28
                }
                _ = fi21
            }
            var v29 uint8 = 2
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 string = "a"
                var v32 int32 = (v2 % (((v2 | v2)) | 1))
                v33 := GenS2{F0: 4294967295, F1: 4294967295, F2: v1}
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
            _ = v29
        } else {
            fmt.Println(v1, v1, v4)
            for fi34 := 0; fi34 < 19; fi34++ {
                v35 := []int{}
                v36 := GenS2{F0: 65535, F1: 42, F2: v1}
                _ = fi34
                _ = v35
                _ = v36
            }
        }
    }
    switch -477 {
    case -9223372036854775808:
        fmt.Println(v3, v4, v4)
    case 32768:
        v3 += v3
    }
    v37 := []int{127}
    v4 = ((float64((float32(0.0) + 0.1)) / v4) / v4)
    fmt.Println(v3)
    var v38 uint = 1
    if len(v37) > 0 {
        v37[0] = -1
    }
    if false {
        if false {
            switch v1 {
            case "":
                var v39 int32 = v2
                _ = v39
            default:
                v1 = v1
            }
            var v40 int64 = v3
            v1 = v1
            _ = v40
        }
        v1 = ""
    }
    v37 = append(v37, 9223372036854775807)
    v41 := GenS1{F0: "\n\t", F1: 42, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v41
}
