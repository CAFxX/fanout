package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16; F1 uint64; F2 uint16 }
func genFn0(p0 bool) uint32 {
    var v5 uint8 = 0
    switch "\n\t" {
    case "hello":
        for fi6 := 0; fi6 < 3; fi6++ {
            for fi7 := 0; fi7 < 5; fi7++ {
                if false {
                    if p0 {
                        v8 := []int{}
                        _ = v8
                    } else {
                        var v9 uint16 = 437
                        v10 := GenS1{F0: (!((float32((float64(((0.1 - -2.25) * 0.0)) * 1.5)) >= 0.1) && p0)), F1: 42}
                        _ = v10
                        _ = v9
                    }
                    fmt.Println(v5, fi7, p0)
                    v11 := GenS1{F0: true, F1: 353}
                    var v12 uint16 = 2
                    var v13 float32 = 0.0
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    var v14 bool = true
                    _ = v14
                }
                _ = fi7
            }
            if true {
                var v15 float32 = float32(0.0)
                var v16 int8 = 30
                if ((0.1 - v15) == ((v15 + (v15 * 3.14159)) + 1e10)) {
                    v17 := GenS2{F0: 1, F1: 42}
                    if p0 {
                        fmt.Println(v16, v15, v5)
                    }
                    _ = v17
                }
                _ = v15
                _ = v16
            }
            _ = fi6
        }
    case "\n\t":
        for fi18 := 0; fi18 < 10; fi18++ {
            if ((!false) || (!p0)) {
                v19 := GenS3{F0: "foo bar", F1: 0}
                _ = v19
            }
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 uint8 = 1
                _ = fi20
                _ = v21
            }
            _ = fi18
        }
    default:
        v22 := GenS4{F0: 42}
        _ = v22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        v24 := GenS2{F0: 761, F1: 255}
        _ = fi23
        _ = v24
    }
    _ = p0
    _ = v5
    return 234
}
func genFn1(p0 string, p1 int32, p2 uint) uint64 {
    p1 = 2147483647
    for fi25 := 0; fi25 < 13; fi25++ {
        p1 -= int32(fi25)
        if true {
            p2 += (p2 % ((516) | 1))
        }
        var v26 int = (fi25 << ((65535) & 63))
        _ = fi25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 770
    v2 := make(map[int]string)
    var v3 bool = true
    var v4 uint32 = (v1 * (v1 << ((4294967295) & 63)))
    for fi27 := 0; fi27 < 20; fi27++ {
        v28 := GenS5{F0: 128, F1: 72, F2: 65535}
        _ = fi27
        _ = v28
    }
    var v29 float32 = (((1.5 * 1e10) * float32(0.1)) + 1e10)
    v4 = (v4 << ((0) & 63))
    if (true || (!((1 > 9223372036854775807) && v3))) {
        for fi30 := 0; fi30 < 6; fi30++ {
            switch (("foo bar" + ("a" + (("" + "foo bar") + ("" + "")))) + "a") {
            case "a":
                switch (fi30 % ((fi30) | 1)) {
                case 32768:
                    if false {
                        v31 := GenS2{F0: genFn1("a", -1, 382), F1: 908}
                        for fi32 := 0; fi32 < 2; fi32++ {
                            fmt.Println(fi32)
                            _ = fi32
                        }
                        if v3 {
                            var v33 float64 = float64(v29)
                            var v34 string = "\n\t"
                            v35 := []int{fi30}
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        _ = v31
                    }
                case -419:
                    v36 := GenS6{F0: -1}
                    _ = v36
                case 32767:
                    v3 = (-1 < 2147483647)
                default:
                    v37 := []int{(fi30 << ((65535) & 63)), fi30, fi30}
                    _ = v37
                }
            case "":
                var v38 uint16 = 0
                _ = v38
            }
            var v39 float64 = 0.0
            var v40 string = ""
            _ = fi30
            _ = v39
            _ = v40
        }
    } else {
        var v41 int16 = 32767
        fmt.Println(v29, v1, v3)
        _ = v41
    }
    var v42 int8 = 42
    if v3 {
        var v43 uint = uint(v1)
        var v44 int64 = -9223372036854775808
        _ = v43
        _ = v44
    }
    v3 = ("a" != "")
    var v45 int16 = 32767
    if v3 {
        v46 := GenS7{F0: (v45 & (^(v45))), F1: genFn1((("" + "\n\t") + "hello"), 0, 526), F2: 935}
        _ = v46
    }
    var v47 uint64 = 1
    v48 := []int{57, 9223372036854775807}
    v42 = (v42 / ((-37) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v42
    _ = v45
    _ = v47
    _ = v48
}
