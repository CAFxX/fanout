package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int16; F1 int32; F2 int32 }
func genFn0(p0 int, p1 uint64, p2 int8) uint16 {
    var v5 uint = 2
    var v6 uint16 = 255
    var v7 string = "foo bar"
    v7 = v7
    v8 := GenS1{F0: (v7 + "a"), F1: int8(p0), F2: 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v6
}
func genFn1(p0 uint16, p1 uint64, p2 int8) int64 {
    var v9 uint8 = 191
    fmt.Println(p2)
    fmt.Println(v9, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 929
}

func main() {
    var v1 uint16 = 730
    var v2 int64 = 941
    var v3 uint32 = 559
    var v4 float64 = (float64(3.14159) * float64((3.14159 * 0.1)))
    v10 := GenS1{F0: "foo bar", F1: -91, F2: 4294967295}
    if (2 == genFn0(4294967295, uint64(v1), 100)) {
        v11 := []int{}
        v12 := GenS2{F0: 65535}
        _ = v11
        _ = v12
    } else {
        switch "" {
        case "hello":
            fmt.Println(v1, v4)
        case "a":
            var v13 uint64 = 18446744073709551615
            _ = v13
        case "\n\t":
            v3 -= 0
        }
        v14 := []int{127, 128, 836, 1}
        _ = v14
    }
    v4 = v4
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int32 = 2147483647
        v17 := []int{fi15, -861, fi15, -793}
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS2{F0: genFn0(127, 286, 33)}
        _ = fi18
        _ = v19
    }
    fmt.Println(v4, v4)
    v20 := GenS3{F0: "foo bar", F1: 127}
    if false {
        var v21 int8 = -128
        for fi22 := 0; fi22 < 6; fi22++ {
            fmt.Println(v1, v3)
            v4 -= v4
            _ = fi22
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 uint8 = 2
            if (v21 < (v21 << ((1) & 63))) {
                v25 := []int{fi23, fi23, fi23}
                v1 -= v1
                _ = v25
            } else {
                v26 := []int{fi23}
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        _ = v21
    }
    if true {
        for fi27 := 0; fi27 < 3; fi27++ {
            v28 := GenS2{F0: (^(v1))}
            v4 -= (0.1 * (v4 / 1e10))
            v29 := []int{0, 1, (fi27 - 128), fi27}
            _ = fi27
            _ = v28
            _ = v29
        }
        switch 1 {
        case 200:
            var v30 float32 = 1.5
            _ = v30
        }
        switch 4294967295 {
        case 2147483647:
            if ((^(v1)) == (^(v1))) {
                v31 := []int{}
                v2 += v2
                var v32 int8 = 100
                _ = v31
                _ = v32
            } else {
                v33 := GenS2{F0: v1}
                _ = v33
            }
        case 0:
            v34 := GenS1{F0: "", F1: 127, F2: 4294967295}
            _ = v34
        default:
            v35 := []int{}
            _ = v35
        }
    }
    v36 := []int{42, 32768, -990, 128}
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    var v37 int16 = int16(v1)
    for fi38 := 0; fi38 < 18; fi38++ {
        var v39 uint32 = v3
        for fi40 := 0; fi40 < 1; fi40++ {
            v2 = v2
            if false {
                fmt.Println(fi40, v3, v2)
                if (((("\n\t" + "a") + ("foo bar" + "foo bar")) <= ("" + ("foo bar" + "a"))) && (!true)) {
                    fmt.Println(v1)
                    var v41 uint8 = 78
                    var v42 int64 = (v2 * genFn1(genFn0(fi40, 2, -117), 18446744073709551615, 87))
                    _ = v41
                    _ = v42
                }
            } else {
                if true {
                    v43 := GenS3{F0: "foo bar", F1: 42}
                    _ = v43
                }
                v44 := GenS2{F0: 65535}
                var v45 uint16 = v1
                _ = v44
                _ = v45
            }
            var v46 float64 = v4
            if len(v36) > 0 {
                v36[0] = 9223372036854775807
            }
            _ = fi40
            _ = v46
        }
        if true {
            fmt.Println(fi38, v37, v2)
            var v47 uint = 42
            fmt.Println(v3, v2, v37)
            _ = v47
        } else {
            v48 := GenS4{F0: v37, F1: 42, F2: -1}
            v37 += v37
            _ = v48
        }
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
