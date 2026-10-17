package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 uint; F2 uint16 }

func main() {
    var v1 uint32 = 129
    var v2 int16 = 825
    v3 := make(map[string]uint32)
    var v4 bool = false
    switch 9223372036854775807 {
    case -2147483648:
        v4 = (!v4)
    default:
        var v5 int64 = 42
        _ = v5
    }
    v6 := GenS1{F0: 42, F1: 18446744073709551615, F2: 4294967295}
    if (-2147483648 < 32768) {
        v1 = (v1 << ((100) & 63))
    }
    switch 628 {
    case -288:
        v2 -= v2
    }
    if v4 {
        v2 = -901
        if (0 == -1) {
            fmt.Println(v1, v2)
        }
    } else {
        switch "foo bar" {
        case "a":
            var v7 uint16 = 2
            _ = v7
        case "hello":
            if v4 {
                v4 = true
            }
        case "\n\t":
            v8 := GenS2{F0: int64(v2), F1: -352}
            _ = v8
        default:
            v1 -= (v1 / ((0) | 1))
        }
        for fi9 := 0; fi9 < 15; fi9++ {
            v4 = v4
            delete(v3, "hello")
            v10 := GenS2{F0: 1, F1: 42}
            _ = fi9
            _ = v10
        }
        if (-9223372036854775808 <= int64(v1)) {
            v11 := GenS3{F0: "\n\t"}
            _ = v11
        }
    }
    if v4 {
        v12 := GenS4{F0: ("foo bar" + ("" + "\n\t"))}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 uint8 = 42
            _ = fi13
            _ = v14
        }
    }
    v15 := []int{0, 42, 128}
    if (!v4) {
        var v16 uint8 = 0
        _ = v16
    } else {
        fmt.Println(v4)
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v15 = append(v15, fi17)
        v18 := v3["hello"]
        switch "foo bar" {
        case "hello":
            if true {
                if v4 {
                    var v19 int = fi17
                    var v20 uint = 129
                    var v21 uint64 = 2
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    v22 := GenS5{F0: 42, F1: 2, F2: 65535}
                    if (((true && (true || (4294967295 > 4294967295))) && v4) || (v4 && (809 != 2))) {
                        if ((v2 ^ (v2 >> ((255) & 63))) >= v2) {
                            v23 := GenS2{F0: 1, F1: (fi17 % ((42) | 1))}
                            v2 = (v2 & (v2 & (v2 * (v2 | (v2 % (((v2 ^ -960)) | 1))))))
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 12; fi24++ {
                            if len(v15) > 0 {
                                v15[0] = fi17
                            }
                            _ = fi24
                        }
                        var v25 uint64 = 2
                        _ = v25
                    }
                    _ = v22
                }
                fmt.Println(v2, v2)
            }
        case "a":
            var v26 float32 = float32(0.1)
            _ = v26
        case "foo bar":
            if v4 {
                var v27 uint32 = (v18 | v1)
                for fi28 := 0; fi28 < 15; fi28++ {
                    v18 = v18
                    v29 := GenS3{F0: ""}
                    _ = fi28
                    _ = v29
                }
                _ = v27
            } else {
                for fi30 := 0; fi30 < 7; fi30++ {
                    fmt.Println(v2)
                    v31 := v3["foo bar"]
                    v15 = append(v15, fi17)
                    _ = fi30
                    _ = v31
                }
                if (int32(v1) != 0) {
                    fmt.Println(v4, v2)
                    switch (fi17 * fi17) {
                    case 9223372036854775807:
                        var v32 bool = v4
                        _ = v32
                    case -9223372036854775808:
                        fmt.Println(fi17, fi17)
                    case 4294967295:
                        v33 := GenS1{F0: 278, F1: 0, F2: 42}
                        _ = v33
                    default:
                        for fi34 := 0; fi34 < 15; fi34++ {
                            var v35 string = "\n\t"
                            _ = len(v15)
                            _ = fi34
                            _ = v35
                        }
                    }
                } else {
                    var v36 uint = 743
                    v37 := GenS5{F0: (v1 & v1), F1: (v36 / ((18446744073709551615) | 1)), F2: 550}
                    _ = v36
                    _ = v37
                }
            }
        default:
            var v38 int = fi17
            _ = v38
        }
        _ = fi17
        _ = v18
    }
    switch ("" + "hello") {
    case "foo bar":
        v15 = append(v15, 4294967295)
    case "":
        for fi39 := 0; fi39 < 14; fi39++ {
            v40 := v3[("\n\t" + ("\n\t" + ""))]
            v41 := GenS5{F0: 1, F1: 42, F2: 364}
            var v42 uint8 = 1
            _ = fi39
            _ = v40
            _ = v41
            _ = v42
        }
    }
    var v43 int64 = 0
    var v44 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v6
}
