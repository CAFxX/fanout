package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := []uint32{1}
    v2 := []string{"", "a", ((("\n\t" + "") + "hello") + ("foo bar" + "a")), "hello"}
    var v3 int8 = 127
    if ((4294967295 > 227) && (42 <= 9223372036854775807)) {
        switch "\n\t" {
        case "a":
            var v4 uint64 = 744
            _ = v4
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi5 := 0; fi5 < 17; fi5++ {
            fmt.Println(fi5, fi5)
            switch "hello" {
            case "hello":
                _ = len(v2)
            default:
                if false {
                    var v6 int32 = -65
                    switch (fi5 & fi5) {
                    case -734:
                        v7 := GenS1{F0: 42, F1: 472}
                        _ = v7
                    case -766:
                        v8 := GenS1{F0: 353, F1: 2}
                        _ = v8
                    case 1:
                        v6 = v6
                    }
                    switch (fi5 & (-(fi5))) {
                    case 0:
                        fmt.Println(v6, fi5)
                    }
                    _ = v6
                } else {
                    for fi9 := 0; fi9 < 8; fi9++ {
                        v3 -= 0
                        _ = fi9
                    }
                    switch 4294967295 {
                    case 2147483647:
                        var v10 int8 = (v3 << ((4294967295) & 63))
                        _ = v10
                    case 0:
                        if false {
                            var v11 uint32 = 2147483648
                            var v12 int32 = -1
                            _ = v11
                            _ = v12
                        }
                    }
                    var v13 uint = 2
                    _ = v13
                }
            }
            v3 += (v3 | (v3 * (^(v3))))
            _ = fi5
        }
        _ = len(v1)
    }
    fmt.Println(v3)
    fmt.Println(v3)
    if ((!true) || false) {
        var v14 float64 = 1e10
        _ = v14
    } else {
        if (!(!false)) {
            switch 0 {
            case 894:
                if false {
                    var v15 int64 = -425
                    var v16 uint64 = 2
                    _ = v15
                    _ = v16
                }
            case -281:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v17 := GenS1{F0: 42, F1: 4294967295}
            _ = v17
        } else {
            v18 := GenS2{F0: 0, F1: 1, F2: -1}
            _ = v18
        }
        v19 := GenS3{F0: 42}
        _ = v19
    }
    fmt.Println(v3)
    if true {
        var v20 uint64 = 606
        _ = v20
    }
    switch 1 {
    case -1:
        v3 = v3
    case -926:
        v3 -= v3
    case -9223372036854775808:
        var v21 int = -267
        _ = v21
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 int32 = 384
        _ = fi22
        _ = v23
    }
    v1 = append(v1, 4294967295)
    if (-1 >= -1) {
        var v24 string = ""
        _ = v24
    }
    v25 := GenS1{F0: 1000, F1: 720}
    v1 = append(v1, 946)
    if (!false) {
        var v26 int = 1
        fmt.Println(v26, v3)
        _ = v26
    }
    v1 = append(v1, 0)
    var v27 int16 = -549
    if true {
        v28 := GenS3{F0: 101}
        v3 -= -128
        _ = v28
    } else {
        v1 = append(v1, 693)
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        if len(v2) > 0 {
            v2[0] = ""
        }
        _ = fi29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
}
