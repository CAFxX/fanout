package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
func genFn0(p0 uint16, p1 uint8) int {
    v5 := GenS1{F0: 42}
    fmt.Println(p0)
    for fi6 := 0; fi6 < 17; fi6++ {
        p0 -= p0
        p1 = (^(p1))
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := GenS2{F0: 18446744073709551615, F1: -32768}
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    if ((34 < 0) || true) {
        v9 := []int{}
        _ = v9
    } else {
        switch 32767 {
        case 826:
            var v10 string = ("" + "hello")
            _ = v10
        case 95:
            v11 := GenS2{F0: 4294967295, F1: 32767}
            _ = v11
        case 768:
            for fi12 := 0; fi12 < 1; fi12++ {
                fmt.Println(p1, p1, fi12)
                var v13 uint64 = 4294967295
                _ = fi12
                _ = v13
            }
        }
        p1 = (^(p1))
        if (!false) {
            var v14 string = "\n\t"
            var v15 int8 = 2
            if ((!(true && (255 > 670))) && false) {
                v16 := []int{}
                for fi17 := 0; fi17 < 19; fi17++ {
                    for fi18 := 0; fi18 < 13; fi18++ {
                        _ = len(v16)
                        v19 := GenS2{F0: uint(p1), F1: -32768}
                        _ = fi18
                        _ = v19
                    }
                    if (p1 > (p1 ^ p1)) {
                        v20 := GenS3{F0: uint(p1)}
                        _ = v20
                    }
                    var v21 int32 = 226
                    _ = fi17
                    _ = v21
                }
                _ = v16
            }
            _ = v14
            _ = v15
        } else {
            var v22 uint8 = (p1 >> ((4294967295) & 63))
            v23 := GenS4{F0: 4294967295}
            _ = v22
            _ = v23
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 9223372036854775807
}
func genFn1(p0 uint16, p1 int16, p2 uint8) string {
    if (!false) {
        var v24 string = "a"
        var v25 float32 = 1e10
        _ = v24
        _ = v25
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 int32 = 232
            var v28 float32 = 1e10
            fmt.Println(v28)
            _ = fi26
            _ = v27
            _ = v28
        }
    case 42:
        if false {
            p1 = p1
            if false {
                v29 := GenS2{F0: 255, F1: (^(p1))}
                var v30 float32 = 1.5
                _ = v29
                _ = v30
            } else {
                var v31 bool = false
                for fi32 := 0; fi32 < 10; fi32++ {
                    switch (fi32 >> ((474) & 63)) {
                    case 32768:
                        var v33 int = (fi32 * fi32)
                        _ = v33
                    case 282:
                        fmt.Println(p0, p1)
                    case 127:
                        if v31 {
                            v34 := []int{}
                            var v35 int = fi32
                            if len(v34) > 0 {
                                v34[0] = genFn0((p0 & (^(p0))), 255)
                            }
                            _ = v34
                            _ = v35
                        }
                    default:
                        if true {
                            v36 := GenS3{F0: 810}
                            var v37 bool = true
                            if (((^(p0)) >= p0) || v37) {
                                if (!(!((0 < 9223372036854775807) && false))) {
                                    var v38 uint8 = p2
                                    _ = v38
                                }
                                v37 = true
                            }
                            _ = v36
                            _ = v37
                        }
                    }
                    _ = fi32
                }
                _ = v31
            }
            var v39 uint64 = 4294967295
            _ = v39
        }
    case 1:
        var v40 int32 = 0
        _ = v40
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        v42 := []int{genFn0((p0 - 255), 127), fi41}
        for fi43 := 0; fi43 < 2; fi43++ {
            if false {
                var v44 uint16 = 1
                if len(v42) > 0 {
                    v42[0] = (fi41 % (((-(fi43))) | 1))
                }
                _ = v44
            }
            _ = fi43
        }
        fmt.Println(p0, fi41, p1)
        _ = fi41
        _ = v42
    }
    _ = p0
    _ = p1
    _ = p2
    return ("foo bar" + "foo bar")
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 1.5
    v3 := make(map[int]uint32)
    var v4 float32 = 0.1
    switch -364 {
    case 4294967295:
        fmt.Println(v4)
    }
    for fi45 := 0; fi45 < 16; fi45++ {
        if false {
            v2 += (1.5 * v4)
        }
        v2 -= (v4 / (v2 + ((1.5 / float32(3.14159)) - float32(1e10))))
        for fi46 := 0; fi46 < 19; fi46++ {
            var v47 uint8 = 10
            fmt.Println(v2)
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    for fi48 := 0; fi48 < 11; fi48++ {
        var v49 uint8 = 255
        switch 127 {
        case 32767:
            var v50 uint64 = 42
            _ = v50
        }
        _ = fi48
        _ = v49
    }
    for fi51 := 0; fi51 < 16; fi51++ {
        v2 = ((v2 + (3.14159 / ((v2 * v4) + (float32(((3.14159 - 0.1) / 3.14159)) / float32(((-2.25 - (1e10 * 3.14159)) - (1.5 - 0.1))))))) - v2)
        switch "foo bar" {
        case "foo bar":
            var v52 string = "a"
            _ = v52
        }
        _ = fi51
    }
    var v53 float32 = -2.25
    var v54 int8 = 21
    for fi55 := 0; fi55 < 8; fi55++ {
        fmt.Println(fi55, fi55)
        var v56 string = "a"
        _ = fi55
        _ = v56
    }
    var v57 string = "foo bar"
    var v58 bool = (!(4294967295 != 7))
    delete(v3, -159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v57
    _ = v58
}
