package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int32; F2 int64 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 uint64; F1 int64 }
func genFn0(p0 bool, p1 uint, p2 uint32) int8 {
    switch -1 {
    case 321:
        p1 -= 90
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        if (!((fi6 ^ fi6) > fi6)) {
            v7 := GenS1{F0: 42}
            var v8 int8 = -128
            _ = v7
            _ = v8
        } else {
            p1 -= uint(p2)
            switch (fi6 | fi6) {
            case 0:
                v9 := GenS1{F0: 0}
                _ = v9
            }
        }
        switch (fi6 - (fi6 - fi6)) {
        case 4294967295:
            var v10 uint32 = p2
            _ = v10
        case -1:
            if false {
                if (true || p0) {
                    p0 = (p0 && (!p0))
                    var v11 uint = (^(p1))
                    var v12 string = "\n\t"
                    _ = v11
                    _ = v12
                } else {
                    p0 = (!p0)
                }
                for fi13 := 0; fi13 < 4; fi13++ {
                    v14 := []int{fi6}
                    p2 = p2
                    _ = fi13
                    _ = v14
                }
                var v15 int = fi6
                _ = v15
            } else {
                v16 := []int{fi6, fi6, fi6, 42}
                var v17 bool = true
                var v18 int = fi6
                _ = v16
                _ = v17
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 float64 = ((0.0 + ((0.1 - (1e10 + (0.0 * 0.1))) + 0.0)) / 0.1)
            _ = fi19
            _ = v20
        }
        _ = fi6
    }
    v21 := []int{-762, 127, 4294967295, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return -92
}

func main() {
    v1 := []uint32{1, 4294967295, 0, 812}
    v2 := []string{"", ("foo bar" + ((("hello" + "") + "\n\t") + "hello"))}
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    var v5 uint32 = 2147483648
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 uint16 = 42
        _ = fi22
        _ = v23
    }
    v5 += v5
    if (false && (-1 <= -32768)) {
        if (false || (-313 > 42)) {
            for fi24 := 0; fi24 < 17; fi24++ {
                if (((0.0 * 0.1) / 1.5) != ((((((0.1 * 3.14159) / 1.0) + ((0.1 / float32(3.14159)) / 1.0)) - 1.5) - (3.14159 / float32(1e10))) / 1.0)) {
                    fmt.Println(fi24, v5)
                    var v25 uint32 = 935
                    _ = len(v4)
                    _ = v25
                } else {
                    if false {
                        if (65535 < 32767) {
                            fmt.Println(v5, fi24)
                        }
                        v26 := GenS2{F0: 1000, F1: -280, F2: -498}
                        _ = v26
                    }
                    if false {
                        delete(v4, "foo bar")
                    } else {
                        for fi27 := 0; fi27 < 5; fi27++ {
                            v28 := v4["a"]
                            v5 = v5
                            var v29 float64 = -2.25
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        var v30 float32 = 3.14159
                        var v31 uint16 = 32767
                        _ = v30
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 17; fi32++ {
                        v5 = (v5 | (v5 << ((1) & 63)))
                        _ = fi32
                    }
                }
                _ = fi24
            }
            var v33 bool = (!(true && (!(1 >= -870))))
            _ = v33
        } else {
            fmt.Println(v5)
            for fi34 := 0; fi34 < 11; fi34++ {
                var v35 uint16 = 382
                for fi36 := 0; fi36 < 1; fi36++ {
                    var v37 int = 4294967295
                    v35 = (v35 + v35)
                    v35 += v35
                    _ = fi36
                    _ = v37
                }
                _ = fi34
                _ = v35
            }
            v1 = append(v1, v5)
        }
        for fi38 := 0; fi38 < 7; fi38++ {
            var v39 string = "\n\t"
            _ = fi38
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        if (380 <= -9223372036854775808) {
            _ = len(v3)
        }
        for fi41 := 0; fi41 < 19; fi41++ {
            switch "foo bar" {
            case "":
                var v42 uint64 = 42
                _ = v42
            default:
                v43 := GenS3{F0: -1, F1: 4294967295}
                _ = v43
            }
            _ = fi41
        }
        _ = fi40
    }
    for fi44 := 0; fi44 < 1; fi44++ {
        v45 := GenS2{F0: 42, F1: 1, F2: 711}
        v46 := GenS4{F0: 42, F1: -248}
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 uint32 = v5
    v47 = 2147483648
    var v48 int8 = 42
    var v49 float64 = (1.5 / (((0.1 * 3.14159) / 1e10) * (1e10 + float64(1e10))))
    v49 = v49
    fmt.Println(v47, v5)
    v50 := GenS4{F0: 0, F1: 0}
    switch "hello" {
    case "foo bar":
        var v51 uint64 = 2
        _ = v51
    default:
        var v52 bool = (!false)
        _ = v52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
