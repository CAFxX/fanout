package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint; F1 int; F2 string }
func genFn0(p0 uint16, p1 uint64, p2 int8) bool {
    v6 := GenS1{F0: -237}
    var v7 uint32 = 189
    switch "hello" {
    case "\n\t":
        v8 := GenS1{F0: -9223372036854775808}
        _ = v8
    default:
        switch 558 {
        case 9223372036854775807:
            switch "foo bar" {
            case "\n\t":
                if (false && true) {
                    v9 := GenS1{F0: 1}
                    v10 := GenS1{F0: -28}
                    fmt.Println(p1)
                    _ = v10
                    _ = v9
                } else {
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v12 := []int{}
                        _ = fi11
                        _ = v12
                    }
                }
            default:
                if (-1 != -9223372036854775808) {
                    p2 = 127
                    var v13 bool = (-9223372036854775808 > int(p2))
                    switch -9223372036854775808 {
                    case 42:
                        var v14 bool = (!(true || false))
                        _ = v14
                    case 11:
                        v15 := GenS1{F0: 857}
                        _ = v15
                    }
                    _ = v13
                } else {
                    v7 = v7
                }
            }
        case -9223372036854775808:
            var v16 bool = true
            _ = v16
        default:
            p1 = (^(p1))
        }
    }
    v17 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v6
    _ = v7
    return (false && ("" == "foo bar"))
}
func genFn1(p0 int64) uint64 {
    for fi18 := 0; fi18 < 7; fi18++ {
        p0 = p0
        var v19 uint16 = 0
        _ = fi18
        _ = v19
    }
    var v20 uint32 = 4294967295
    if false {
        var v21 int16 = 125
        if true {
            var v22 string = ("foo bar" + "foo bar")
            _ = v22
        } else {
            v23 := GenS2{F0: false, F1: (v20 >> ((569) & 63))}
            v24 := []int{67, 484, 128}
            var v25 int8 = -1
            _ = v23
            _ = v24
            _ = v25
        }
        _ = v21
    } else {
        var v26 uint32 = (^(v20))
        var v27 bool = ((true || false) && (true || (!((!false) && true))))
        var v28 uint32 = (v20 | (v20 + (^(v26))))
        _ = v26
        _ = v27
        _ = v28
    }
    var v29 float32 = -2.25
    v30 := GenS2{F0: (false || false), F1: (^(v20))}
    _ = p0
    _ = v20
    _ = v29
    _ = v30
    return 973
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 42
    var v3 uint32 = 1
    var v4 int16 = 42
    v5 := []uint32{1}
    if (128 >= 127) {
        var v31 uint64 = 267
        _ = v31
    }
    var v32 uint16 = 255
    v4 += int16(v3)
    var v33 uint16 = 232
    for fi34 := 0; fi34 < 11; fi34++ {
        v32 = 328
        if false {
            v32 -= v32
            v32 = (v33 | (v33 + v32))
        } else {
            v4 = (-(v4))
            fmt.Println(v2, fi34)
        }
        _ = fi34
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        v36 := GenS2{F0: (!(v33 <= v33)), F1: (v3 ^ (v3 + v3))}
        v37 := GenS3{F0: 127, F1: 42}
        switch 32767 {
        case 957:
            switch "foo bar" {
            case "foo bar":
                var v38 float32 = float32(3.14159)
                _ = v38
            case "\n\t":
                var v39 uint8 = 255
                _ = v39
            }
        default:
            for fi40 := 0; fi40 < 5; fi40++ {
                var v41 uint8 = 243
                v42 := GenS3{F0: -65, F1: (v41 / ((v41) | 1))}
                fmt.Println(fi35, fi40, v33)
                _ = fi40
                _ = v41
                _ = v42
            }
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    switch "foo bar" {
    case "\n\t":
        v43 := GenS3{F0: int8(v32), F1: 127}
        _ = v43
    case "foo bar":
        for fi44 := 0; fi44 < 18; fi44++ {
            fmt.Println(v33, v33, v3)
            for fi45 := 0; fi45 < 3; fi45++ {
                v46 := GenS3{F0: 75, F1: 3}
                fmt.Println(v33, fi44)
                v47 := GenS4{F0: 255, F1: fi45, F2: "a"}
                _ = fi45
                _ = v46
                _ = v47
            }
            _ = fi44
        }
    case "":
        var v48 int8 = 100
        _ = v48
    }
    v32 -= (v33 << ((907) & 63))
    v1[uint32(v32)] = v3
    var v49 uint32 = (v3 - (v3 ^ v3))
    if (100 == 72) {
        for fi50 := 0; fi50 < 13; fi50++ {
            var v51 int32 = 1
            delete(v1, (v49 - (v49 - v3)))
            v32 += 32767
            _ = fi50
            _ = v51
        }
        v4 = (v4 >> ((2) & 63))
        v52 := GenS1{F0: -1}
        _ = v52
    }
    switch "" {
    case "a":
        _ = len(v1)
    }
    if ((v33 + 1) != (^(v32))) {
        var v53 bool = (2 < 1)
        for fi54 := 0; fi54 < 11; fi54++ {
            for fi55 := 0; fi55 < 13; fi55++ {
                var v56 string = "a"
                switch fi55 {
                case -1:
                    fmt.Println(v3)
                case 853:
                    v57 := GenS1{F0: 505}
                    _ = v57
                case -2147483648:
                    v49 = (v49 << ((42) & 63))
                }
                _ = fi55
                _ = v56
            }
            _ = fi54
        }
        var v58 uint = uint(v49)
        _ = v53
        _ = v58
    } else {
        v32 = (v32 << ((4294967295) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v49
    _ = v5
}
