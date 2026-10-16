package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
func genFn0() int {
    if true {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint32 = 2147483648
            for fi9 := 0; fi9 < 9; fi9++ {
                switch (("" + "foo bar") + "") {
                case "":
                    v10 := []int{fi7, (^(fi7))}
                    _ = v10
                default:
                    if ((fi9 / (((fi9 ^ -2147483648)) | 1)) < fi9) {
                        var v11 float64 = (0.1 + (0.1 * 1e10))
                        var v12 float64 = 0.1
                        _ = v11
                        _ = v12
                    } else {
                        var v13 string = "hello"
                        _ = v13
                    }
                }
                var v14 uint = 500
                v15 := GenS1{F0: 350, F1: (v8 / (((v8 * 359)) | 1)), F2: false}
                v16 := []int{}
                _ = fi9
                _ = v14
                _ = v15
                _ = v16
            }
            v17 := []int{fi7}
            _ = fi7
            _ = v17
            _ = v8
        }
        var v18 string = "foo bar"
        _ = v18
    }
    v19 := GenS2{F0: 4294967295}
    var v20 int8 = -1
    var v21 int64 = -833
    v22 := []int{}
    _ = v19
    _ = v20
    _ = v21
    _ = v22
    return 27
}
func genFn1(p0 string) string {
    var v23 int32 = -236
    if (4294967295 <= 18) {
        p0 = (p0 + (p0 + "a"))
    }
    v24 := []int{9223372036854775807}
    var v25 int16 = -684
    for fi26 := 0; fi26 < 12; fi26++ {
        var v27 int32 = -2147483648
        if ((fi26 <= fi26) || (genFn0() != fi26)) {
            if len(v24) > 0 {
                v24[0] = fi26
            }
            v28 := GenS3{F0: v27, F1: 65535}
            _ = v28
        } else {
            v24 = append(v24, fi26)
            v23 -= v23
        }
        if false {
            if (42 >= 0) {
                var v29 int16 = 0
                for fi30 := 0; fi30 < 17; fi30++ {
                    var v31 uint16 = 448
                    v32 := GenS4{F0: 675}
                    fmt.Println(p0, fi30)
                    var v33 int32 = int32(v31)
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = v29
            }
        }
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = v23
    _ = v24
    _ = v25
    return "a"
}

func main() {
    var v1 int = -248
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    var v4 int8 = -104
    v5 := make(map[int]bool)
    v6 := make(map[string]bool)
    fmt.Println(v1, v1)
    var v34 int64 = -9223372036854775808
    v35 := []int{128, genFn0(), v1, v1}
    fmt.Println(v4)
    v34 -= v34
    delete(v3, 467)
    if true {
        _ = len(v35)
        v1 -= v1
        switch (v1 & int(v34)) {
        case 1:
            v1 += -2147483648
        case 2147483647:
            for fi36 := 0; fi36 < 20; fi36++ {
                var v37 uint = 18446744073709551615
                var v38 uint8 = 42
                if true {
                    var v39 int8 = -1
                    _ = v39
                }
                _ = fi36
                _ = v37
                _ = v38
            }
        default:
            v1 -= v1
        }
    } else {
        v40 := GenS5{F0: genFn0()}
        v4 -= v4
        _ = v40
    }
    var v41 int = -2147483648
    v42 := GenS2{F0: 1}
    v4 = (v4 * v4)
    var v43 int8 = v4
    if false {
        v44 := GenS1{F0: 1000, F1: 422, F2: false}
        if true {
            if (!false) {
                v45 := GenS3{F0: -1, F1: 911}
                switch "" {
                case "hello":
                    v35 = append(v35, v41)
                case "a":
                    v35 = append(v35, genFn0())
                }
                v46 := GenS2{F0: 2147483648}
                _ = v45
                _ = v46
            }
        } else {
            v41 += (v41 | -2147483648)
            _ = len(v6)
        }
        _ = v44
    }
    v34 = (v34 | v34)
    var v47 string = "foo bar"
    var v48 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
