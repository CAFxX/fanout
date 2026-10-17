package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 string; F2 string }
func genFn0(p0 bool, p1 int16, p2 uint) bool {
    for fi5 := 0; fi5 < 20; fi5++ {
        p0 = (42 >= 516)
        _ = fi5
    }
    p2 -= 957
    p1 -= p1
    v6 := []int{42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}
func genFn1(p0 int8, p1 uint32, p2 int) uint64 {
    switch "\n\t" {
    case "hello":
        var v7 uint32 = (p1 + p1)
        _ = v7
    case "a":
        if (!false) {
            p1 = p1
            var v8 float32 = (0.1 / 1.5)
            p0 = p0
            _ = v8
        } else {
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 string = ("foo bar" + "")
                v11 := GenS1{F0: -44, F1: 126}
                _ = fi9
                _ = v10
                _ = v11
            }
        }
    default:
        var v12 uint = 65535
        _ = v12
    }
    p0 += p0
    var v13 int8 = p0
    v14 := []int{p2, p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 2
}

func main() {
    var v1 uint8 = 122
    var v2 bool = true
    v3 := []string{"foo bar"}
    v4 := make(map[uint32]string)
    var v15 uint64 = 2
    for fi16 := 0; fi16 < 2; fi16++ {
        switch "" {
        case "hello":
            var v17 int = fi16
            _ = v17
        case "foo bar":
            v2 = (!true)
        default:
            v18 := GenS1{F0: 32767, F1: 509}
            _ = v18
        }
        _ = fi16
    }
    v2 = (v1 < 128)
    v19 := GenS2{F0: 100, F1: 934}
    v15 = (v15 & (v15 / ((v15) | 1)))
    if false {
        if len(v3) > 0 {
            v3[0] = ""
        }
        _ = len(v3)
        switch -476 {
        case 32767:
            _ = len(v3)
        }
    }
    fmt.Println(v15, v2)
    var v20 float64 = 1.5
    v15 -= v15
    for fi21 := 0; fi21 < 6; fi21++ {
        fmt.Println(v2, v20, v1)
        _ = fi21
    }
    fmt.Println(v2)
    for fi22 := 0; fi22 < 17; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            switch (^(fi22)) {
            case -397:
                if v2 {
                    var v24 uint32 = 0
                    fmt.Println(v15, fi23)
                    _ = v24
                } else {
                    v25 := GenS3{F0: 200}
                    _ = v25
                }
            default:
                v26 := GenS4{F0: 733}
                _ = v26
            }
            var v27 uint32 = 2
            _ = fi23
            _ = v27
        }
        fmt.Println(v2, v1)
        v28 := GenS2{F0: -1, F1: 32767}
        _ = fi22
        _ = v28
    }
    if false {
        if (!true) {
            for fi29 := 0; fi29 < 12; fi29++ {
                v30 := GenS5{F0: 186, F1: "\n\t", F2: "foo bar"}
                var v31 uint64 = (v15 << ((255) & 63))
                for fi32 := 0; fi32 < 13; fi32++ {
                    v15 = v15
                    _ = len(v3)
                    _ = fi32
                }
                _ = fi29
                _ = v30
                _ = v31
            }
        } else {
            var v33 float32 = 1e10
            var v34 bool = false
            _ = v33
            _ = v34
        }
        fmt.Println(v20)
        v35 := GenS2{F0: 1, F1: -32768}
        _ = v35
    } else {
        if ((float32(v20) - 0.0) == (0.1 - 0.1)) {
            var v36 int64 = 1
            var v37 string = (("foo bar" + ("foo bar" + (("\n\t" + ("a" + "foo bar")) + ""))) + ("a" + "foo bar"))
            if v2 {
                switch (v37 + v37) {
                case "\n\t":
                    v36 = -1
                case "hello":
                    for fi38 := 0; fi38 < 4; fi38++ {
                        for fi39 := 0; fi39 < 3; fi39++ {
                            _ = len(v4)
                            _ = fi39
                        }
                        var v40 uint64 = v15
                        v41 := GenS4{F0: v15}
                        _ = fi38
                        _ = v40
                        _ = v41
                    }
                case "foo bar":
                    v15 -= v15
                default:
                    var v42 int = 0
                    _ = v42
                }
                v2 = (((v37 + (v37 + ((v37 + v37) + v37))) + v37) != "hello")
                for fi43 := 0; fi43 < 14; fi43++ {
                    var v44 uint16 = 65535
                    _ = fi43
                    _ = v44
                }
            } else {
                for fi45 := 0; fi45 < 17; fi45++ {
                    if v2 {
                        delete(v4, 4294967295)
                    }
                    v15 = (v15 % ((genFn1(127, 2147483648, fi45)) | 1))
                    var v46 int16 = -414
                    _ = fi45
                    _ = v46
                }
                var v47 float64 = float64((0.1 - 1e10))
                _ = v47
            }
            _ = v36
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
