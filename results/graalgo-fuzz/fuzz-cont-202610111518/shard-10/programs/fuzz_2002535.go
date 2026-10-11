package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint16; F2 int64 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint32; F1 int16; F2 uint64 }
func genFn0(p0 string, p1 int32) int {
    switch "hello" {
    case "\n\t":
        p0 = (p0 + p0)
    case "foo bar":
        v5 := GenS1{F0: 4294967295}
        _ = v5
    default:
        v6 := []int{42, 42, 128}
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int64 = 366
        v9 := GenS1{F0: 0}
        fmt.Println(p1)
        _ = fi7
        _ = v8
        _ = v9
    }
    if (0.1 >= 3.14159) {
        for fi10 := 0; fi10 < 7; fi10++ {
            if (fi10 != -827) {
                var v11 int32 = (p1 & -912)
                for fi12 := 0; fi12 < 3; fi12++ {
                    var v13 int16 = 42
                    if true {
                        var v14 float32 = 1e10
                        if true {
                            if (true && true) {
                                switch p0 {
                                case "\n\t":
                                    if false {
                                        var v15 uint16 = 197
                                        switch (-(fi12)) {
                                        case 0:
                                            v14 += v14
                                        }
                                        v11 = 42
                                        _ = v15
                                    }
                                }
                                if (!true) {
                                    var v16 float64 = 1.5
                                    v16 += (v16 * v16)
                                    _ = v16
                                }
                            }
                            v17 := []int{fi12, (^(fi12)), 32767, fi10}
                            v18 := GenS2{F0: 42, F1: v11, F2: 162}
                            _ = v17
                            _ = v18
                        }
                        var v19 float64 = -2.25
                        _ = v14
                        _ = v19
                    }
                    v13 -= v13
                    _ = fi12
                    _ = v13
                }
                _ = v11
            } else {
                if true {
                    fmt.Println(p1)
                    v20 := GenS1{F0: 751}
                    p1 = 276
                    _ = v20
                } else {
                    v21 := []int{434}
                    v22 := GenS3{F0: 18446744073709551615}
                    switch ((p0 + p0) + ((p0 + p0) + p0)) {
                    case "":
                        var v23 uint64 = 4294967295
                        _ = v23
                    case "hello":
                        v21 = append(v21, (fi10 ^ (^(fi10))))
                    }
                    _ = v21
                    _ = v22
                }
                var v24 uint16 = 255
                _ = v24
            }
            _ = fi10
        }
        if false {
            for fi25 := 0; fi25 < 18; fi25++ {
                p0 = "foo bar"
                _ = fi25
            }
            for fi26 := 0; fi26 < 14; fi26++ {
                for fi27 := 0; fi27 < 10; fi27++ {
                    v28 := GenS4{F0: (p1 << ((18446744073709551615) & 63)), F1: 42, F2: 743}
                    fmt.Println(fi26, p0, p0)
                    var v29 string = "hello"
                    v30 := []int{(fi27 * (^(fi27))), (fi27 + 792), (fi26 << ((197) & 63)), (fi26 & 46)}
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                v31 := []int{(fi26 | fi26), (fi26 % ((fi26) | 1)), (-(fi26)), fi26}
                fmt.Println(p1, p0)
                _ = fi26
                _ = v31
            }
            if (!false) {
                v32 := GenS1{F0: 2}
                _ = v32
            }
        } else {
            fmt.Println(p0)
            p1 = -2147483648
        }
    }
    p1 -= p1
    v33 := []int{-527, 32767, 4294967295}
    _ = p0
    _ = p1
    _ = v33
    return -375
}

func main() {
    v1 := []int{32768, -918}
    v2 := make(map[string]int)
    var v3 int64 = -659
    var v4 string = ""
    if false {
        if len(v1) > 0 {
            v1[0] = 83
        }
        if true {
            fmt.Println(v3, v3)
            v34 := GenS2{F0: 2, F1: -1, F2: 128}
            var v35 float64 = ((3.14159 / (1e10 * float64(-2.25))) + ((0.0 - 1.5) * (0.1 - 1e10)))
            _ = v34
            _ = v35
        }
    }
    _ = len(v2)
    var v36 uint = 42
    v1 = append(v1, genFn0(v4, 0))
    if true {
        if (float64(1.5) < (float64(float32((-2.25 - 0.0))) - (0.0 + 3.14159))) {
            var v37 int = genFn0(("hello" + v4), -1)
            var v38 int8 = 100
            _ = v37
            _ = v38
        }
        v3 = v3
        var v39 bool = true
        _ = v39
    } else {
        _ = len(v1)
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        switch v4 {
        case "a":
            v4 = v4
        default:
            _ = len(v2)
        }
        _ = fi40
    }
    switch (("hello" + v4) + "hello") {
    case "a":
        v41 := GenS3{F0: 1}
        _ = v41
    case "hello":
        v42 := GenS5{F0: 364, F1: 811}
        _ = v42
    case "\n\t":
        v43 := GenS6{F0: 2, F1: 0, F2: 4294967295}
        _ = v43
    }
    var v44 string = v4
    fmt.Println(v36)
    v36 += 174
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v44
}
