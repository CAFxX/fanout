package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int; F2 int8 }
func genFn0(p0 int64, p1 uint8) int32 {
    switch "hello" {
    case "a":
        if (38 < -128) {
            if true {
                if (!true) {
                    var v6 int64 = p0
                    _ = v6
                } else {
                    for fi7 := 0; fi7 < 10; fi7++ {
                        fmt.Println(fi7, fi7)
                        p1 += p1
                        v8 := []int{}
                        _ = fi7
                        _ = v8
                    }
                }
                p0 -= p0
            }
        } else {
            for fi9 := 0; fi9 < 1; fi9++ {
                v10 := GenS1{F0: "a", F1: -89, F2: fi9}
                _ = fi9
                _ = v10
            }
            p0 -= p0
            var v11 int64 = -234
            _ = v11
        }
    default:
        fmt.Println(p1, p1)
    }
    var v12 uint = 913
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 int64 = p0
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        for fi17 := 0; fi17 < 9; fi17++ {
            if false {
                p0 = -735
            } else {
                if (!true) {
                    v18 := []int{fi16}
                    var v19 uint = (v12 % ((v12) | 1))
                    v18 = append(v18, fi16)
                    _ = v18
                    _ = v19
                } else {
                    p0 = (p0 * p0)
                }
            }
            _ = fi17
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = v12
    return -1
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 0
    var v3 int64 = 1
    v4 := []uint32{42, 1, 2147483647}
    var v5 uint = 1
    var v20 uint8 = 201
    v4 = append(v4, 2)
    for fi21 := 0; fi21 < 10; fi21++ {
        if true {
            v22 := GenS1{F0: ((("" + "a") + "a") + "a"), F1: -128, F2: (-(fi21))}
            if true {
                v23 := GenS1{F0: "a", F1: 127, F2: fi21}
                _ = v23
            }
            fmt.Println(v20, v3, v5)
            _ = v22
        }
        v24 := GenS1{F0: "", F1: -46, F2: -913}
        v25 := GenS2{F0: fi21}
        _ = fi21
        _ = v24
        _ = v25
    }
    switch "" {
    case "\n\t":
        v20 = (v20 >> ((1) & 63))
    case "hello":
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 int32 = 0
            _ = fi26
            _ = v27
        }
    case "foo bar":
        for fi28 := 0; fi28 < 6; fi28++ {
            for fi29 := 0; fi29 < 13; fi29++ {
                _ = len(v1)
                var v30 uint64 = 2
                v20 = (^(v20))
                var v31 int16 = v2
                _ = fi29
                _ = v30
                _ = v31
            }
            v3 -= v3
            var v32 int8 = 127
            _ = fi28
            _ = v32
        }
    }
    v5 = uint(v20)
    v3 -= v3
    fmt.Println(v3, v3, v2)
    fmt.Println(v3)
    var v33 uint16 = uint16(v20)
    v20 -= (v20 | v20)
    var v34 int8 = 19
    var v35 string = ("a" + "")
    var v36 uint = (^(v5))
    var v37 uint8 = (v20 >> ((42) & 63))
    if true {
        var v38 uint64 = 2
        v39 := GenS2{F0: 42}
        var v40 float32 = (-2.25 * 1.5)
        _ = v38
        _ = v39
        _ = v40
    } else {
        switch -1 {
        case 830:
            for fi41 := 0; fi41 < 6; fi41++ {
                v3 -= v3
                _ = fi41
            }
        case 683:
            switch "\n\t" {
            case "\n\t":
                switch -727 {
                case -645:
                    v42 := GenS3{F0: -1, F1: 1, F2: v34}
                    _ = v42
                case 4294967295:
                    fmt.Println(v3, v5, v34)
                case -1:
                    var v43 bool = (v35 == "hello")
                    _ = v43
                }
            case "":
                switch ("a" + v35) {
                case "\n\t":
                    v20 += uint8(v33)
                }
            case "a":
                fmt.Println(v20, v20, v5)
            }
        case 0:
            v44 := GenS2{F0: 32768}
            _ = v44
        }
    }
    var v45 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v45
    _ = v5
}
