package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint }
func genFn0(p0 bool, p1 uint16) uint {
    if (1 == 0) {
        var v4 int = 4294967295
        v4 -= 128
        var v5 uint64 = 1
        _ = v4
        _ = v5
    } else {
        p1 = (p1 / (((p1 | p1)) | 1))
    }
    for fi6 := 0; fi6 < 12; fi6++ {
        switch "foo bar" {
        case "a":
            var v7 uint16 = 2
            _ = v7
        case "":
            var v8 uint16 = 255
            _ = v8
        case "foo bar":
            if (p0 || (!(true || (p0 && ((p0 && false) || p0))))) {
                if (p1 <= (^(p1))) {
                    switch "" {
                    case "":
                        for fi9 := 0; fi9 < 10; fi9++ {
                            var v10 bool = (p0 && false)
                            _ = fi9
                            _ = v10
                        }
                    case "foo bar":
                        var v11 uint16 = 255
                        _ = v11
                    case "\n\t":
                        var v12 int16 = 1000
                        _ = v12
                    default:
                        for fi13 := 0; fi13 < 18; fi13++ {
                            v14 := []int{fi13}
                            _ = fi13
                            _ = v14
                        }
                    }
                    for fi15 := 0; fi15 < 19; fi15++ {
                        p0 = false
                        v16 := GenS1{F0: fi6}
                        var v17 string = ("\n\t" + "foo bar")
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
                var v18 int64 = 0
                _ = v18
            }
        default:
            v19 := GenS1{F0: (fi6 - (fi6 << ((632) & 63)))}
            _ = v19
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (42 < 128)
    var v3 uint64 = 4294967295
    var v20 int16 = 1
    switch -715 {
    case 32767:
        if (65535 != 332) {
            var v21 uint64 = v3
            v22 := v1[127]
            _ = v21
            _ = v22
        }
    case 287:
        if v2 {
            v3 += 1
        }
    default:
        var v23 int8 = -1
        _ = v23
    }
    var v24 uint64 = (^(v3))
    if v2 {
        v20 -= v20
        for fi25 := 0; fi25 < 7; fi25++ {
            switch "foo bar" {
            case "foo bar":
                var v26 int8 = 0
                _ = v26
            case "\n\t":
                fmt.Println(v3, v2)
            case "hello":
                v27 := []int{(fi25 / ((fi25) | 1))}
                _ = v27
            default:
                if v2 {
                    v2 = v2
                    fmt.Println(v20, v20, v3)
                    v28 := GenS1{F0: fi25}
                    _ = v28
                } else {
                    v24 -= (^(v24))
                }
            }
            _ = fi25
        }
    }
    v29 := GenS2{F0: 0, F1: 963}
    v3 = v3
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 float32 = 3.14159
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        v2 = true
        for fi33 := 0; fi33 < 2; fi33++ {
            fmt.Println(v20, v20, fi33)
            v2 = false
            _ = fi33
        }
        switch fi32 {
        case 9223372036854775807:
            var v34 int32 = -2147483648
            _ = v34
        case -2147483648:
            v35 := []int{fi32}
            _ = v35
        case 4294967295:
            var v36 int16 = v20
            _ = v36
        }
        _ = fi32
    }
    var v37 int16 = 42
    if v2 {
        v38 := GenS2{F0: 0, F1: 4294967295}
        v39 := GenS1{F0: 9223372036854775807}
        _ = v38
        _ = v39
    }
    for fi40 := 0; fi40 < 19; fi40++ {
        v24 = v24
        for fi41 := 0; fi41 < 9; fi41++ {
            var v42 uint = 18446744073709551615
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    delete(v1, 42)
    fmt.Println(v37, v37, v37)
    v43 := GenS2{F0: 1, F1: 65535}
    _ = len(v1)
    v44 := []int{32767, -636, 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v29
    _ = v3
    _ = v37
    _ = v43
    _ = v44
}
