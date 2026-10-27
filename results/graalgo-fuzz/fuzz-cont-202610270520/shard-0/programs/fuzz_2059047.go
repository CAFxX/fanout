package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int64; F2 int32 }
func genFn0(p0 bool, p1 int8, p2 int64) uint64 {
    for fi6 := 0; fi6 < 6; fi6++ {
        p0 = true
        _ = fi6
    }
    p0 = p0
    fmt.Println(p2)
    v7 := GenS1{F0: (p1 << ((1) & 63)), F1: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 72
}
func genFn1() bool {
    var v8 uint64 = 42
    if ((v8 | genFn0((100 < 42), 127, 9223372036854775807)) != (v8 << ((2) & 63))) {
        switch "foo bar" {
        case "":
            switch ("hello" + "foo bar") {
            case "hello":
                v8 = v8
            }
        case "a":
            switch 4294967295 {
            case 727:
                v9 := []int{4294967295, 42, 69, 262}
                _ = v9
            case 9223372036854775807:
                for fi10 := 0; fi10 < 4; fi10++ {
                    var v11 uint32 = 938
                    v11 += (v11 + 4294967295)
                    _ = fi10
                    _ = v11
                }
            case 128:
                for fi12 := 0; fi12 < 6; fi12++ {
                    v8 -= (v8 << ((852) & 63))
                    for fi13 := 0; fi13 < 6; fi13++ {
                        if false {
                            switch "foo bar" {
                            case "":
                                var v14 string = ("" + "\n\t")
                                _ = v14
                            case "hello":
                                if (0 == 0) {
                                    var v15 float32 = 1e10
                                    _ = v15
                                }
                            default:
                                v8 = v8
                            }
                            v16 := []int{}
                            var v17 uint32 = 42
                            _ = v16
                            _ = v17
                        }
                        _ = fi13
                    }
                    _ = fi12
                }
            }
        default:
            fmt.Println(v8)
        }
        v8 = 333
        v8 = v8
    }
    v8 += (v8 ^ v8)
    v8 = v8
    _ = v8
    return (true && ((65535 >= 973) || (1 < 42)))
}

func main() {
    v1 := []string{"a", "\n\t", ("foo bar" + ""), "\n\t"}
    var v2 uint8 = 75
    v3 := make(map[uint32]string)
    v4 := make(map[string]string)
    v5 := []string{"hello", "foo bar", "a"}
    v2 = v2
    switch 893 {
    case 1:
        if (100 >= 21) {
            switch "a" {
            case "foo bar":
                v2 -= 46
            }
            v2 = 42
        }
    case 32768:
        var v18 int32 = -2147483648
        _ = v18
    }
    if ("a" != ("foo bar" + "a")) {
        for fi19 := 0; fi19 < 2; fi19++ {
            if (0 <= 1) {
                v2 += 128
                v2 += 0
                var v20 uint64 = 449
                _ = v20
            } else {
                var v21 int = (fi19 / ((-465) | 1))
                _ = v21
            }
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 uint = 615
                v24 := GenS1{F0: -1, F1: 224}
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi19
        }
        _ = len(v4)
        v25 := v3[uint32(v2)]
        _ = v25
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v1 = append(v1, "a")
    var v26 uint16 = 0
    fmt.Println(v26)
    fmt.Println(v2)
    for fi27 := 0; fi27 < 9; fi27++ {
        for fi28 := 0; fi28 < 18; fi28++ {
            _ = len(v1)
            _ = fi28
        }
        var v29 int = fi27
        _ = fi27
        _ = v29
    }
    if (!(0 >= -1)) {
        fmt.Println(v2, v2)
        var v30 uint8 = (v2 / ((v2) | 1))
        _ = v30
    } else {
        v4["foo bar"] = ("a" + "hello")
    }
    v31 := GenS1{F0: 127, F1: 65535}
    switch "foo bar" {
    case "\n\t":
        for fi32 := 0; fi32 < 9; fi32++ {
            var v33 int16 = -1
            fmt.Println(v26)
            _ = fi32
            _ = v33
        }
    default:
        var v34 float32 = float32((0.0 * 0.0))
        _ = v34
    }
    delete(v3, 380)
    if (false && true) {
        if true {
            v26 = (v26 | 65535)
            v2 += (v2 & 2)
            fmt.Println(v2, v26)
        }
        for fi35 := 0; fi35 < 9; fi35++ {
            if true {
                v36 := GenS2{F0: 0}
                var v37 float64 = -2.25
                v38 := GenS3{F0: 1, F1: 1, F2: 1}
                _ = v36
                _ = v37
                _ = v38
            } else {
                for fi39 := 0; fi39 < 4; fi39++ {
                    v26 -= v26
                    _ = fi39
                }
            }
            fmt.Println(fi35, v26, v2)
            var v40 bool = false
            _ = fi35
            _ = v40
        }
        for fi41 := 0; fi41 < 10; fi41++ {
            var v42 bool = (255 <= 1)
            v42 = (4294967295 > 884)
            for fi43 := 0; fi43 < 9; fi43++ {
                v26 = (v26 + (v26 >> ((545) & 63)))
                _ = fi43
            }
            _ = fi41
            _ = v42
        }
    }
    v2 -= (v2 + (v2 * 57))
    var v44 int16 = 14
    v2 = uint8(v26)
    v44 = (v44 * v44)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v44
    _ = v5
}
