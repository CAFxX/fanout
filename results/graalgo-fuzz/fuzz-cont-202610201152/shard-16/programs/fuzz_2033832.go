package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 uint16; F1 uint8; F2 bool }
type GenS7 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint16, p1 int) int64 {
    switch (-(p1)) {
    case 128:
        for fi6 := 0; fi6 < 16; fi6++ {
            p0 += (p0 ^ 255)
            for fi7 := 0; fi7 < 4; fi7++ {
                v8 := GenS1{F0: 42, F1: -603}
                var v9 int16 = 0
                p0 = (p0 & (p0 % ((32767) | 1)))
                var v10 bool = ("foo bar" < "hello")
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            p0 -= p0
            _ = fi6
        }
    case 127:
        p0 += p0
    case 9223372036854775807:
        var v11 bool = (0 != 42)
        _ = v11
    default:
        fmt.Println(p1, p1)
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS1{F0: 740, F1: 32767}
        var v14 int16 = -1
        if false {
            switch (p1 << ((357) & 63)) {
            case 32768:
                p1 = p1
            case -1:
                v15 := GenS2{F0: 660, F1: -73}
                _ = v15
            }
        } else {
            if false {
                if false {
                    for fi16 := 0; fi16 < 19; fi16++ {
                        if (true && false) {
                            var v17 uint64 = 18446744073709551615
                            _ = v17
                        }
                        if false {
                            v18 := GenS3{F0: (p0 & p0), F1: 1}
                            v19 := GenS4{F0: false, F1: 2147483648}
                            var v20 string = "foo bar"
                            _ = v18
                            _ = v19
                            _ = v20
                        } else {
                            var v21 int64 = 0
                            var v22 bool = false
                            _ = v21
                            _ = v22
                        }
                        _ = fi16
                    }
                }
            }
            fmt.Println(p1, v14)
            v23 := []int{}
            _ = v23
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    p1 = (p1 % ((p1) | 1))
    v24 := []int{p1}
    _ = p0
    _ = p1
    _ = v24
    return 0
}

func main() {
    v1 := []int{4294967295, 2147483647, 32767}
    v2 := []string{}
    var v3 int = -9223372036854775808
    var v4 int32 = 1
    v5 := make(map[string]uint32)
    var v25 uint16 = 42
    var v26 string = (("hello" + (("foo bar" + "hello") + (("a" + "foo bar") + "a"))) + "\n\t")
    v27 := GenS5{F0: ((v26 + v26) + "a"), F1: v25}
    var v28 float32 = (0.0 * -2.25)
    switch v26 {
    case "":
        switch v3 {
        case 4294967295:
            var v29 uint16 = (v25 * v25)
            _ = v29
        case 128:
            v30 := GenS1{F0: 2147483648, F1: 32767}
            _ = v30
        case 1:
            if false {
                v31 := GenS6{F0: 0, F1: 127, F2: (v3 != (-(v3)))}
                switch ("hello" + "foo bar") {
                case "hello":
                    _ = len(v5)
                case "foo bar":
                    var v32 int = (v3 % ((v3) | 1))
                    _ = v32
                case "\n\t":
                    switch v3 {
                    case 42:
                        if (4294967295 == 0) {
                            v5["foo bar"] = 898
                        }
                    case 784:
                        v28 += 0.0
                    case 2147483647:
                        var v33 uint64 = 355
                        _ = v33
                    default:
                        var v34 float32 = (v28 / ((v28 * v28) - ((0.0 * v28) * (v28 - 0.0))))
                        _ = v34
                    }
                default:
                    var v35 bool = (255 != 18446744073709551615)
                    _ = v35
                }
                var v36 float64 = 1e10
                _ = v31
                _ = v36
            }
        }
    }
    if true {
        switch 0 {
        case -2147483648:
            for fi37 := 0; fi37 < 3; fi37++ {
                for fi38 := 0; fi38 < 4; fi38++ {
                    v39 := GenS1{F0: 304, F1: -587}
                    if (((v26 + (v26 + ("foo bar" + ""))) + (v26 + ("" + ("" + "a")))) == v26) {
                        v40 := GenS1{F0: 4294967295, F1: 469}
                        _ = v40
                    } else {
                        v3 = fi38
                    }
                    switch fi38 {
                    case -1:
                        if ((v28 - float32(3.14159)) <= v28) {
                            switch fi37 {
                            case 128:
                                v26 = (v26 + v26)
                            case -381:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            var v41 int32 = (v4 << ((2) & 63))
                            _ = v41
                        }
                        v28 += 1.5
                        var v42 uint = 322
                        _ = v42
                    case -310:
                        var v43 int8 = 127
                        v28 -= v28
                        switch fi37 {
                        case 4294967295:
                            var v44 int8 = (v43 - int8(fi37))
                            _ = v44
                        case 127:
                            v25 = (v25 & 2)
                        default:
                            _ = len(v1)
                        }
                        _ = v43
                    case -2147483648:
                        v3 = (fi38 * fi38)
                    }
                    if true {
                        v45 := GenS1{F0: 2147483647, F1: -1}
                        var v46 int = v3
                        _ = v45
                        _ = v46
                    }
                    _ = fi38
                    _ = v39
                }
                _ = fi37
            }
        case 32768:
            fmt.Println(v4, v25, v28)
        }
        switch v3 {
        case -551:
            v47 := GenS7{F0: 1, F1: 0}
            _ = v47
        case 0:
            v48 := GenS7{F0: 1, F1: 0}
            _ = v48
        case -930:
            v3 += 0
        }
        v26 = (v26 + v26)
    } else {
        var v49 int16 = 42
        _ = len(v5)
        _ = v49
    }
    if (1 < 1) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi50 := 0; fi50 < 5; fi50++ {
            fmt.Println(v28, v28, v3)
            v1 = append(v1, (^(fi50)))
            _ = fi50
        }
    }
    var v51 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
