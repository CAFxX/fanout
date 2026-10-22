package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint64) uint64 {
    if true {
        for fi5 := 0; fi5 < 6; fi5++ {
            switch (fi5 | (fi5 / (((fi5 << ((255) & 63))) | 1))) {
            case 328:
                p0 = (p0 ^ (p0 >> ((65535) & 63)))
            case -957:
                var v6 uint = 611
                _ = v6
            case 0:
                var v7 float64 = ((1.5 * 3.14159) * float64(-2.25))
                _ = v7
            default:
                fmt.Println(fi5)
            }
            _ = fi5
        }
        v8 := []int{793, 42}
        p0 += (^(p0))
        _ = v8
    } else {
        switch "\n\t" {
        case "hello":
            switch 42 {
            case -777:
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 int8 = -56
                    var v11 int32 = 53
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case 9223372036854775807:
                if false {
                    v12 := []int{26, 4294967295}
                    var v13 uint8 = 128
                    _ = v12
                    _ = v13
                } else {
                    p0 = p0
                    p0 -= (p0 / (((^(p0))) | 1))
                }
            }
        case "foo bar":
            for fi14 := 0; fi14 < 1; fi14++ {
                fmt.Println(p0, fi14)
                v15 := []int{fi14, (fi14 >> ((2) & 63)), 128}
                _ = fi14
                _ = v15
            }
        case "a":
            p0 = p0
        }
        v16 := []int{32767}
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 int8 = -1
            var v19 int64 = -9223372036854775808
            var v20 uint32 = 975
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        _ = v16
    }
    if true {
        v21 := GenS1{F0: 42, F1: 950}
        var v22 uint64 = (p0 & (p0 ^ p0))
        _ = v21
        _ = v22
    }
    var v23 bool = false
    _ = p0
    _ = v23
    return 18446744073709551615
}

func main() {
    var v1 int8 = 1
    v2 := []int{}
    v3 := []int{0, 128}
    v4 := []int{}
    if false {
        v24 := GenS1{F0: 32767, F1: 0}
        fmt.Println(v1)
        _ = v24
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        v1 = (v1 & v1)
        _ = fi25
    }
    var v26 uint = 65535
    for fi27 := 0; fi27 < 9; fi27++ {
        if ((-2.25 / ((-2.25 - 1.5) / -2.25)) < 0.1) {
            var v28 uint16 = 255
            v29 := GenS1{F0: v28, F1: 42}
            _ = v28
            _ = v29
        }
        if (true || (!true)) {
            for fi30 := 0; fi30 < 16; fi30++ {
                v26 += (v26 % (((v26 | (v26 - 447))) | 1))
                var v31 int64 = -117
                var v32 float64 = 0.0
                _ = fi30
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 2; fi33++ {
                if false {
                    switch "" {
                    case "\n\t":
                        switch fi27 {
                        case 4294967295:
                            var v34 uint64 = 4294967295
                            _ = v34
                        }
                    case "hello":
                        if (!true) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v1 += (v1 * (v1 >> ((v26) & 63)))
                        } else {
                            var v35 int8 = v1
                            v4 = append(v4, fi27)
                            if len(v4) > 0 {
                                v4[0] = (fi33 * (fi33 & fi27))
                            }
                            _ = v35
                        }
                    default:
                        var v36 int32 = -1
                        _ = v36
                    }
                }
                var v37 int64 = -1
                _ = fi33
                _ = v37
            }
            var v38 string = "foo bar"
            _ = v38
        }
        v1 = (v1 % (((v1 + (v1 & (v1 << ((v26) & 63))))) | 1))
        _ = fi27
    }
    v26 = (v26 % (((v26 * v26)) | 1))
    v4 = append(v4, 127)
    v1 -= (v1 >> ((v26) & 63))
    for fi39 := 0; fi39 < 10; fi39++ {
        var v40 string = ("a" + ("" + "a"))
        _ = fi39
        _ = v40
    }
    var v41 int64 = -1
    if false {
        _ = len(v2)
        var v42 uint64 = 2
        _ = v42
    } else {
        var v43 float64 = 0.0
        var v44 int = 32767
        for fi45 := 0; fi45 < 3; fi45++ {
            var v46 float32 = 0.0
            _ = fi45
            _ = v46
        }
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v41
}
