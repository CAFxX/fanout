package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint }
func genFn0(p0 int16, p1 uint32, p2 uint64) int16 {
    if (1.5 != 1.5) {
        var v7 uint16 = 65535
        var v8 uint32 = uint32(p2)
        var v9 uint64 = 2
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint32 = (p1 / ((p1) | 1))
    switch "\n\t" {
    case "":
        p2 -= p2
    case "hello":
        v11 := []int{}
        _ = v11
    case "a":
        if ((0.0 * 3.14159) >= -2.25) {
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 string = ("hello" + "hello")
                var v14 float32 = 0.0
                v15 := []int{(fi12 | fi12)}
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            fmt.Println(p2)
            fmt.Println(p2, p1, p1)
        }
    default:
        v16 := GenS1{F0: 0}
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v18 := []int{fi17, fi17, 4294967295}
        var v19 int64 = -9223372036854775808
        var v20 int = 325
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    if (!true) {
        var v21 uint64 = (^(p2))
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 335
    v3 := make(map[int]int)
    v4 := []string{}
    v5 := []string{("" + "")}
    v6 := make(map[int]string)
    v2 = (v2 + (^(v2)))
    if false {
        switch 570 {
        case 1:
            var v22 float32 = 0.1
            _ = v22
        case 127:
            v2 = (v2 * (v2 - v2))
        default:
            var v23 string = "hello"
            _ = v23
        }
    } else {
        for fi24 := 0; fi24 < 9; fi24++ {
            v2 = 18446744073709551615
            _ = fi24
        }
        if ((0 == -94) || false) {
            v25 := GenS1{F0: 0}
            var v26 uint64 = v2
            _ = v25
            _ = v26
        }
    }
    fmt.Println(v2)
    fmt.Println(v2)
    if true {
        switch 1 {
        case 626:
            var v27 string = "foo bar"
            _ = v27
        case 128:
            v28 := GenS1{F0: -1}
            _ = v28
        case 2147483647:
            v6[4294967295] = ""
        }
        if true {
            v29 := v3[32767]
            _ = v29
        }
    }
    switch ("hello" + "foo bar") {
    case "foo bar":
        if false {
            var v30 int8 = -1
            var v31 uint8 = 0
            _ = v30
            _ = v31
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case "":
        switch -295 {
        case 4294967295:
            var v32 uint16 = 255
            _ = v32
        case 9223372036854775807:
            v3[0] = 840
        }
    default:
        v33 := GenS2{F0: 2, F1: 1}
        _ = v33
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        if len(v5) > 0 {
            v5[0] = (("foo bar" + "") + "hello")
        }
        if ((!(!(220 < v2))) && false) {
            var v35 uint32 = 65535
            _ = v35
        } else {
            var v36 int16 = -457
            v2 = 724
            _ = v36
        }
        switch ((("a" + ("foo bar" + "hello")) + ("a" + (("" + ("" + "\n\t")) + "hello"))) + "hello") {
        case "":
            var v37 float64 = 1.5
            _ = v37
        }
        _ = fi34
    }
    v2 = (v2 ^ (v2 << ((255) & 63)))
    v38 := GenS1{F0: -862}
    v39 := GenS2{F0: 65535, F1: 991}
    v40 := GenS3{F0: 42, F1: ((-79 > 1) || (23 < 517)), F2: 42}
    if false {
        for fi41 := 0; fi41 < 3; fi41++ {
            v2 = v2
            if len(v5) > 0 {
                v5[0] = "foo bar"
            }
            _ = fi41
        }
        var v42 uint32 = 65535
        for fi43 := 0; fi43 < 14; fi43++ {
            var v44 uint8 = 2
            fmt.Println(v44, v2, v42)
            _ = fi43
            _ = v44
        }
        _ = v42
    } else {
        v45 := v3[32768]
        v45 = 42
        _ = v45
    }
    for fi46 := 0; fi46 < 15; fi46++ {
        fmt.Println(v2)
        switch -736 {
        case 128:
            var v47 int = (fi46 ^ (-(fi46)))
            _ = v47
        default:
            if true {
                for fi48 := 0; fi48 < 18; fi48++ {
                    var v49 uint16 = 42
                    fmt.Println(fi46, fi48, v49)
                    switch (fi48 * (fi48 & fi46)) {
                    case 4294967295:
                        var v50 int32 = -1
                        var v51 uint32 = 4294967295
                        v50 = (-(v50))
                        _ = v50
                        _ = v51
                    default:
                        v52 := GenS3{F0: (v49 - 794), F1: true, F2: 255}
                        _ = v52
                    }
                    _ = fi48
                    _ = v49
                }
            }
        }
        _ = fi46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
