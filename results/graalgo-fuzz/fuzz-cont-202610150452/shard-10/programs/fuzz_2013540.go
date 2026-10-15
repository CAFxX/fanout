package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 int64 }
func genFn0(p0 uint64, p1 uint8) int64 {
    for fi4 := 0; fi4 < 11; fi4++ {
        if true {
            var v5 string = ("hello" + "a")
            _ = v5
        } else {
            p1 = uint8(p0)
            for fi6 := 0; fi6 < 17; fi6++ {
                var v7 bool = (false || (255 >= 1))
                _ = fi6
                _ = v7
            }
            v8 := GenS1{F0: ("a" + "a")}
            _ = v8
        }
        _ = fi4
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := []int{int(p1), fi9}
        _ = fi9
        _ = v10
    }
    var v11 int16 = 32767
    _ = p0
    _ = p1
    _ = v11
    return -166
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]string)
    v3 := make(map[string]bool)
    fmt.Println(v1)
    if ((v1 >> ((0) & 63)) == (v1 % (((-(v1))) | 1))) {
        v12 := []int{127, -1}
        _ = v12
    } else {
        if true {
            switch -516 {
            case 0:
                var v13 int32 = 718
                _ = v13
            case -922:
                v14 := GenS1{F0: ("" + "foo bar")}
                _ = v14
            case -2147483648:
                v1 -= (v1 / (((v1 & 0)) | 1))
            }
            fmt.Println(v1)
            fmt.Println(v1)
        }
        switch "\n\t" {
        case "hello":
            for fi15 := 0; fi15 < 19; fi15++ {
                v16 := GenS2{F0: 300}
                fmt.Println(v1)
                _ = fi15
                _ = v16
            }
        }
        v17 := GenS3{F0: true}
        _ = v17
    }
    if false {
        var v18 uint32 = 501
        v18 = v18
        if (128 > 128) {
            if false {
                var v19 uint16 = 297
                _ = v19
            }
            var v20 uint64 = 129
            _ = v20
        } else {
            v21 := []int{32767, -1}
            var v22 uint = 571
            _ = v21
            _ = v22
        }
        _ = v18
    }
    switch -594 {
    case -844:
        fmt.Println(v1)
    }
    var v23 uint8 = 117
    if (2 <= 2) {
        fmt.Println(v1, v23)
    }
    var v24 bool = false
    v25 := []int{4294967295, 134}
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 13; fi27++ {
            var v28 bool = (v24 && (!v24))
            if (0.1 == (float32(0.0) * -2.25)) {
                v29 := GenS4{F0: 18446744073709551615, F1: v1}
                switch (^(fi27)) {
                case -2147483648:
                    if (!v24) {
                        _ = len(v2)
                        switch 32767 {
                        case 2147483647:
                            var v30 bool = v28
                            _ = v30
                        }
                    }
                default:
                    var v31 float64 = (-2.25 * 1.5)
                    _ = v31
                }
                _ = v29
            } else {
                var v32 bool = (v28 && ((65535 > 4294967295) || v24))
                var v33 uint16 = 974
                v2[2147483648] = "hello"
                fmt.Println(v32)
                _ = v32
                _ = v33
            }
            _ = fi27
            _ = v28
        }
        v34 := GenS2{F0: -1}
        _ = fi26
        _ = v34
    }
    v35 := GenS3{F0: v24}
    v36 := GenS3{F0: true}
    switch 305 {
    case -942:
        for fi37 := 0; fi37 < 7; fi37++ {
            v23 = v23
            fmt.Println(v23, fi37)
            v24 = (-1 != v1)
            _ = fi37
        }
    case -2147483648:
        var v38 uint16 = 0
        _ = v38
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        fmt.Println(v24, fi39)
        _ = fi39
    }
    var v40 float32 = 1.5
    v24 = false
    var v41 bool = (!v24)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v40
    _ = v41
}
