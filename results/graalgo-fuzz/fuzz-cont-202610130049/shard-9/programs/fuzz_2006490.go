package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 bool; F2 bool }
type GenS4 struct { F0 uint16; F1 int16 }
func genFn0(p0 int8, p1 int32, p2 string) string {
    fmt.Println(p2, p1)
    var v4 int32 = (p1 + 0)
    var v5 int32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}
func genFn1(p0 int8) uint32 {
    v6 := GenS1{F0: 4294967295}
    var v7 int64 = -701
    var v8 int16 = 1000
    if (1e10 >= (3.14159 + ((-2.25 + 0.1) + -2.25))) {
        switch (("foo bar" + "foo bar") + ("foo bar" + (("\n\t" + ("foo bar" + "a")) + ""))) {
        case "foo bar":
            switch 9223372036854775807 {
            case 1:
                p0 -= (p0 % ((p0) | 1))
            case 128:
                if (!true) {
                    var v9 int16 = (v8 << ((930) & 63))
                    if ("" == "foo bar") {
                        v9 -= v8
                        v7 = v7
                        v10 := GenS2{F0: false}
                        _ = v10
                    } else {
                        var v11 int32 = 502
                        for fi12 := 0; fi12 < 19; fi12++ {
                            v13 := []int{fi12, (-(fi12)), fi12, (fi12 & 810)}
                            fmt.Println(p0, v9, v8)
                            p0 = 0
                            _ = fi12
                            _ = v13
                        }
                        fmt.Println(p0)
                        _ = v11
                    }
                    v14 := GenS3{F0: 734, F1: (!false), F2: true}
                    _ = v14
                    _ = v9
                } else {
                    fmt.Println(p0, v8)
                    switch ("a" + "") {
                    case "foo bar":
                        p0 = (p0 / (((p0 / (((-(p0))) | 1))) | 1))
                    }
                }
            default:
                v7 -= -338
            }
        case "\n\t":
            var v15 uint = 783
            _ = v15
        default:
            var v16 float64 = 0.0
            _ = v16
        }
        switch ("foo bar" + "foo bar") {
        case "":
            fmt.Println(v7, v7)
        default:
            var v17 int = -916
            _ = v17
        }
        v18 := GenS3{F0: 1, F1: true, F2: (403 > 1)}
        _ = v18
    }
    switch ("foo bar" + ("" + (("a" + "") + ""))) {
    case "hello":
        switch 32767 {
        case 700:
            switch "foo bar" {
            case "foo bar":
                var v19 int64 = (v7 * (^(v7)))
                _ = v19
            case "":
                var v20 string = (("hello" + "a") + "a")
                _ = v20
            }
        case -9223372036854775808:
            var v21 int16 = v8
            _ = v21
        case -907:
            for fi22 := 0; fi22 < 1; fi22++ {
                switch (fi22 / ((757) | 1)) {
                case 2147483647:
                    fmt.Println(p0, v7, fi22)
                }
                _ = fi22
            }
        default:
            fmt.Println(v7)
        }
    case "":
        var v23 uint64 = 2
        _ = v23
    default:
        switch "hello" {
        case "\n\t":
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 uint8 = 178
                v25 = v25
                fmt.Println(v7, v8)
                _ = fi24
                _ = v25
            }
        case "a":
            if (!(1.5 == 0.0)) {
                p0 = p0
                var v26 int8 = -12
                _ = v26
            }
        default:
            for fi27 := 0; fi27 < 1; fi27++ {
                var v28 int8 = p0
                var v29 string = "a"
                v30 := []int{(fi27 ^ 128), fi27, (fi27 ^ fi27), fi27}
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        }
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 814
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{1, 137}
    v3 := make(map[uint32]string)
    v31 := GenS1{F0: 101}
    if true {
        v32 := GenS2{F0: (!(4294967295 > 485))}
        _ = v32
    } else {
        if true {
            var v33 int8 = -1
            v33 = (v33 << ((255) & 63))
            _ = v33
        }
        var v34 string = "a"
        v34 = (v34 + "a")
        _ = v34
    }
    var v35 uint8 = 255
    v1[42] = 991
    if true {
        var v36 float64 = 1.5
        _ = v36
    }
    var v37 uint16 = 42
    if ((1e10 / 3.14159) <= (1.5 / 1e10)) {
        var v38 uint = 65535
        v37 += 1
        _ = v38
    } else {
        v37 += (v37 * v37)
        v39 := GenS4{F0: v37, F1: -60}
        _ = v39
    }
    if (-880 < -837) {
        for fi40 := 0; fi40 < 2; fi40++ {
            v41 := GenS4{F0: (v37 & 442), F1: 32767}
            _ = fi40
            _ = v41
        }
        v42 := v1[659]
        _ = v42
    }
    if len(v2) > 0 {
        v2[0] = -2147483648
    }
    var v43 float64 = 0.0
    fmt.Println(v43)
    v37 -= (v37 % (((v37 - v37)) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v37
    _ = v43
}
