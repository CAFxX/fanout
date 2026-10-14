package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 uint32; F2 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    v3 := []string{"foo bar", "hello"}
    var v4 int16 = -635
    var v5 int64 = 1
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            v5 += (v5 % ((42) | 1))
            var v8 int8 = -1
            v1[1] = 4294967295
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 bool = false
    fmt.Println(v5, v4, v5)
    switch 705 {
    case -9223372036854775808:
        v2[2147483647] = "a"
    case 1:
        v5 = v5
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS1{F0: "hello", F1: (fi10 & 188), F2: 4294967295}
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        _ = fi10
        _ = v11
    }
    if v9 {
        v12 := GenS2{F0: 2147483648}
        _ = v12
    } else {
        v2[127] = "\n\t"
        var v13 uint8 = 214
        switch 394 {
        case 2147483647:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 954:
            v14 := GenS2{F0: 2}
            _ = v14
        default:
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 int32 = -1
                v17 := GenS3{F0: 140, F1: -126, F2: 65535}
                v18 := GenS4{F0: 42}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        }
        _ = v13
    }
    if v9 {
        var v19 int = 32768
        v20 := GenS1{F0: "foo bar", F1: (v19 + 128), F2: 4294967295}
        _ = v19
        _ = v20
    }
    var v21 uint = 4294967295
    for fi22 := 0; fi22 < 13; fi22++ {
        v23 := v2[652]
        var v24 int = 32767
        if (true || (!(9223372036854775807 >= int64(v21)))) {
            fmt.Println(fi22)
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 float32 = -2.25
    if (!false) {
        var v26 int = -276
        v21 += v21
        _ = v26
    }
    delete(v2, -818)
    var v27 uint64 = 2
    switch ("a" + "\n\t") {
    case "hello":
        var v28 int8 = 1
        _ = v28
    case "a":
        v29 := GenS5{F0: -360, F1: 4294967295, F2: v27}
        _ = v29
    case "":
        if (!v9) {
            switch -9223372036854775808 {
            case -229:
                var v30 int = -296
                _ = v30
            case -9223372036854775808:
                v31 := GenS3{F0: v27, F1: v4, F2: 2147483648}
                _ = v31
            case 9223372036854775807:
                for fi32 := 0; fi32 < 15; fi32++ {
                    v2[(fi32 >> (((^(v21))) & 63))] = "hello"
                    var v33 uint8 = 223
                    v34 := GenS4{F0: 0}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
            var v35 uint16 = 242
            v3 = append(v3, "a")
            _ = v35
        }
    default:
        for fi36 := 0; fi36 < 20; fi36++ {
            switch fi36 {
            case 32767:
                fmt.Println(v21, v9, v4)
            }
            fmt.Println(v9, v9, v5)
            _ = fi36
        }
    }
    var v37 int16 = v4
    v21 += v21
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v9
}
