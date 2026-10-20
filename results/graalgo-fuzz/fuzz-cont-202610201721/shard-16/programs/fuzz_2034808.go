package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int8; F1 int64 }
func genFn0(p0 uint32, p1 bool, p2 int64) int32 {
    p0 = p0
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int64 = (p2 << ((16) & 63))
        _ = fi5
        _ = v6
    }
    var v7 uint = 4294967295
    v8 := []int{127, 164, 127, 0}
    if false {
        var v9 string = ""
        for fi10 := 0; fi10 < 3; fi10++ {
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 uint64 = 4294967295
                v13 := GenS1{F0: 42}
                v14 := GenS1{F0: -128}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            if p1 {
                var v15 uint16 = 65535
                for fi16 := 0; fi16 < 4; fi16++ {
                    if (v7 > v7) {
                        if len(v8) > 0 {
                            v8[0] = fi16
                        }
                    }
                    _ = len(v8)
                    v8 = append(v8, fi10)
                    _ = fi16
                }
                _ = len(v8)
                _ = v15
            }
            _ = fi10
        }
        _ = v9
    } else {
        var v17 int32 = 0
        if (!p1) {
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 int16 = -1
                var v20 uint64 = 0
                var v21 float64 = (float64(0.0) * -2.25)
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            var v22 bool = (((3.14159 + 1e10) * -2.25) <= 0.1)
            _ = v22
        }
        v17 -= v17
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int16, p1 uint32, p2 uint) int16 {
    switch 9223372036854775807 {
    case 32768:
        if ("a" > "hello") {
            v23 := GenS1{F0: 1}
            _ = v23
        }
    case -901:
        fmt.Println(p0, p1, p1)
    default:
        var v24 uint8 = 2
        _ = v24
    }
    switch ("hello" + "") {
    case "":
        var v25 uint16 = 488
        _ = v25
    case "\n\t":
        p0 = p0
    case "hello":
        p0 = p0
    }
    var v26 uint16 = 207
    v27 := GenS2{F0: v26, F1: "a"}
    p1 = 508
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return p0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{1}
    var v3 float64 = (float64(((0.0 + 1.5) * 0.0)) / (1.5 / 0.1))
    v4 := make(map[int]string)
    v28 := GenS2{F0: 968, F1: "a"}
    v3 = float64(0.0)
    v3 = v3
    var v29 uint16 = 32767
    var v30 int32 = 0
    fmt.Println(v30, v3)
    var v31 int8 = 100
    for fi32 := 0; fi32 < 19; fi32++ {
        var v33 uint32 = 440
        var v34 uint8 = 127
        v29 = v29
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 int32 = v30
    fmt.Println(v35, v3, v35)
    switch 2147483647 {
    case 2147483647:
        var v36 uint64 = 0
        _ = v36
    case 892:
        var v37 uint16 = v29
        _ = v37
    default:
        switch "foo bar" {
        case "":
            v38 := GenS1{F0: (v31 ^ (v31 * v31))}
            _ = v38
        case "a":
            v39 := GenS2{F0: v29, F1: "\n\t"}
            _ = v39
        case "\n\t":
            switch 788 {
            case -898:
                var v40 uint32 = 65535
                _ = v40
            }
        }
    }
    v31 -= (v31 | (v31 >> ((424) & 63)))
    for fi41 := 0; fi41 < 9; fi41++ {
        if false {
            v4[(^(fi41))] = ""
        }
        for fi42 := 0; fi42 < 14; fi42++ {
            var v43 bool = true
            _ = fi42
            _ = v43
        }
        _ = fi41
    }
    v3 += v3
    v31 += 1
    v30 = (^(v30))
    for fi44 := 0; fi44 < 9; fi44++ {
        var v45 float32 = ((1e10 * -2.25) - float32((((v3 * float64(3.14159)) / ((v3 + 3.14159) - v3)) + (float64(3.14159) * v3))))
        _ = fi44
        _ = v45
    }
    switch "foo bar" {
    case "hello":
        v46 := GenS3{F0: v31, F1: -9223372036854775808}
        _ = v46
    case "a":
        if (42 == 4294967295) {
            var v47 uint64 = 432
            _ = v47
        } else {
            var v48 float32 = 1e10
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v48
        }
    case "":
        switch 128 {
        case 32768:
            var v49 bool = true
            _ = v49
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
}
