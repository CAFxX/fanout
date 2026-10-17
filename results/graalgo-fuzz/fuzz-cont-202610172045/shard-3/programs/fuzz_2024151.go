package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int16; F2 int32 }
func genFn0() uint {
    var v6 uint64 = 42
    if true {
        fmt.Println(v6)
    }
    if ((0 == 45) && (false || (2 >= 18446744073709551615))) {
        var v7 int64 = -752
        var v8 float64 = (0.0 * 0.0)
        switch "" {
        case "a":
            switch ("" + ("\n\t" + "\n\t")) {
            case "":
                v9 := GenS1{F0: 61}
                _ = v9
            case "a":
                var v10 uint = 0
                _ = v10
            }
        case "":
            if (false || ((v7 % (((v7 - -490)) | 1)) > v7)) {
                fmt.Println(v7, v7, v6)
                var v11 uint32 = 2
                fmt.Println(v7, v6, v7)
                _ = v11
            } else {
                var v12 int8 = 1
                _ = v12
            }
        default:
            v8 = (3.14159 + -2.25)
        }
        _ = v7
        _ = v8
    }
    v13 := []int{0, 1, -9223372036854775808}
    _ = v13
    _ = v6
    return 257
}
func genFn1(p0 bool, p1 int) int {
    if true {
        var v14 float32 = 0.1
        p0 = p0
        var v15 uint16 = 1
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 uint = 18446744073709551615
            p1 -= (p1 * (fi16 / (((fi16 << (((v17 / ((42) | 1))) & 63))) | 1)))
            p0 = p0
            _ = fi16
            _ = v17
        }
        p0 = ((1 != 930) || true)
        p0 = p0
    }
    var v18 int16 = 32767
    var v19 uint32 = 265
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return (p1 ^ (p1 | p1))
}

func main() {
    v1 := []int{-181, -325, -901}
    v2 := []uint32{2147483648, 721, 242}
    var v3 float64 = (float64(-2.25) + 3.14159)
    var v4 string = ""
    v5 := []uint32{42, 1, 190, 297}
    var v20 int32 = 2147483647
    v3 = float64(0.0)
    v4 = ("" + ((v4 + "a") + ("a" + "a")))
    if len(v1) > 0 {
        v1[0] = genFn1((127 >= 100), genFn1((v20 < (v20 << ((1) & 63))), 0))
    }
    switch -108 {
    case -1:
        v3 -= v3
    case -9223372036854775808:
        v2 = append(v2, 1)
    case 9223372036854775807:
        v20 = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        v20 = 862
        if true {
            var v22 uint16 = 32767
            var v23 string = ("hello" + "")
            _ = v22
            _ = v23
        }
        _ = len(v5)
        _ = fi21
    }
    if (!(-2.25 == 1e10)) {
        if true {
            v3 = (float64(3.14159) * v3)
            v20 += v20
        } else {
            switch (v4 + "foo bar") {
            case "a":
                v24 := GenS1{F0: 18446744073709551615}
                _ = v24
            case "":
                var v25 float32 = (float32(v3) + (float32(-2.25) / float32((v3 - v3))))
                _ = v25
            }
            v26 := GenS2{F0: 368, F1: -1, F2: -60}
            _ = v26
        }
        var v27 bool = (0 != 957)
        var v28 int64 = 1
        _ = v27
        _ = v28
    } else {
        var v29 int32 = (v20 << ((255) & 63))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v3 = v3
        _ = v29
    }
    v4 = v4
    for fi30 := 0; fi30 < 15; fi30++ {
        for fi31 := 0; fi31 < 20; fi31++ {
            v20 = v20
            v32 := GenS1{F0: 2}
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch v4 {
                case "foo bar":
                    fmt.Println(fi31, v4, v3)
                    var v33 uint32 = 901
                    var v34 string = (v4 + (v4 + "hello"))
                    _ = v33
                    _ = v34
                case "hello":
                    v3 = (float64(float32(float64(0.1))) / 0.1)
                case "":
                    switch "a" {
                    case "hello":
                        v3 -= v3
                    }
                default:
                    v35 := GenS1{F0: 0}
                    _ = v35
                }
                v3 = (v3 + v3)
            } else {
                if (false || (((3.14159 < 3.14159) || (fi30 >= (fi31 & fi31))) || (482 <= 18446744073709551615))) {
                    fmt.Println(v20)
                }
                fmt.Println(v4, fi30, v3)
                v4 = v4
            }
            var v36 uint = 2
            _ = fi31
            _ = v32
            _ = v36
        }
        var v37 int16 = 1000
        _ = fi30
        _ = v37
    }
    v4 = v4
    for fi38 := 0; fi38 < 12; fi38++ {
        var v39 uint8 = 42
        var v40 int8 = 42
        _ = fi38
        _ = v39
        _ = v40
    }
    v4 = v4
    v2 = append(v2, 910)
    v3 += v3
    for fi41 := 0; fi41 < 20; fi41++ {
        if len(v2) > 0 {
            v2[0] = 313
        }
        var v42 uint8 = 42
        _ = fi41
        _ = v42
    }
    for fi43 := 0; fi43 < 13; fi43++ {
        if ((0.0 < v3) || (false || (-9223372036854775808 == -9223372036854775808))) {
            for fi44 := 0; fi44 < 13; fi44++ {
                v3 = (-2.25 * ((0.0 * float64(-2.25)) / 0.1))
                _ = len(v5)
                _ = fi44
            }
            var v45 uint64 = 1
            _ = v45
        } else {
            v20 = v20
        }
        _ = len(v5)
        var v46 int32 = v20
        _ = fi43
        _ = v46
    }
    switch ("hello" + v4) {
    case "a":
        var v47 int16 = -690
        _ = v47
    case "hello":
        var v48 uint8 = 128
        _ = v48
    case "foo bar":
        fmt.Println(v20, v3, v4)
    }
    var v49 uint16 = 385
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v49
    _ = v5
}
