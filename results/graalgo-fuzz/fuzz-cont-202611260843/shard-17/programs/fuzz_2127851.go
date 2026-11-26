package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint16) uint16 {
    var v4 float64 = 1e10
    v5 := []int{-1, 64}
    var v6 int16 = 886
    var v7 float32 = 0.0
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            v6 -= 32767
            _ = fi9
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return p1
}
func genFn1(p0 int16, p1 uint32, p2 int32) int64 {
    for fi10 := 0; fi10 < 17; fi10++ {
        if true {
            fmt.Println(fi10)
        } else {
            p2 += (p2 << ((255) & 63))
        }
        _ = fi10
    }
    var v11 string = "hello"
    if false {
        switch (v11 + "hello") {
        case "foo bar":
            var v12 int16 = p0
            _ = v12
        default:
            v13 := GenS1{F0: false, F1: 2, F2: 358}
            _ = v13
        }
        var v14 uint32 = 2147483647
        if true {
            v14 = 387
            switch 55 {
            case 2147483647:
                var v15 uint64 = 18446744073709551615
                _ = v15
            }
            fmt.Println(p0, p0, p2)
        } else {
            if ((p1 >> ((42) & 63)) < (v14 % ((p1) | 1))) {
                p1 = (v14 - p1)
                for fi16 := 0; fi16 < 4; fi16++ {
                    for fi17 := 0; fi17 < 8; fi17++ {
                        p0 = (-(p0))
                        v18 := []int{(fi16 << ((484) & 63)), 32768}
                        _ = fi17
                        _ = v18
                    }
                    fmt.Println(p1, fi16, p2)
                    _ = fi16
                }
                v19 := GenS1{F0: (42 <= -1), F1: 368, F2: 0}
                _ = v19
            }
            v20 := GenS2{F0: 934}
            _ = v20
        }
        _ = v14
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 1
}

func main() {
    v1 := []string{"", "a"}
    v2 := []string{(("a" + "\n\t") + "a"), "", ("hello" + "\n\t"), ("" + ((("" + "a") + "foo bar") + ("" + "foo bar")))}
    v3 := []string{}
    var v21 uint8 = 127
    v22 := GenS3{F0: 78, F1: genFn1(1, 2, 1)}
    switch -961 {
    case -9223372036854775808:
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := GenS4{F0: 18446744073709551615}
            var v25 uint64 = 196
            switch 127 {
            case 933:
                switch fi23 {
                case 32767:
                    var v26 float64 = 0.1
                    _ = v26
                case -9223372036854775808:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    fmt.Println(v21, fi23)
                }
            case 128:
                v27 := GenS3{F0: 42, F1: 465}
                _ = v27
            default:
                var v28 uint16 = 197
                _ = v28
            }
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    switch -340 {
    case -365:
        fmt.Println(v21)
    case 1:
        var v29 int32 = 1
        _ = v29
    case 127:
        var v30 uint = 306
        _ = v30
    }
    v21 = (v21 / ((v21) | 1))
    switch ("hello" + ("a" + "a")) {
    case "hello":
        if false {
            var v31 uint = 18446744073709551615
            var v32 uint64 = 4294967295
            var v33 int = 570
            _ = v31
            _ = v32
            _ = v33
        }
    }
    if true {
        var v34 uint64 = 143
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 uint8 = v21
            _ = fi35
            _ = v36
        }
        _ = v34
    } else {
        var v37 uint8 = v21
        _ = v37
    }
    _ = len(v3)
    v38 := GenS5{F0: 4294967295}
    var v39 int32 = int32(v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v38
    _ = v39
}
