package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 bool }
type GenS2 struct { F0 int; F1 bool; F2 string }
type GenS3 struct { F0 uint; F1 int64 }
func genFn0(p0 int, p1 string) uint8 {
    v4 := []int{}
    var v5 uint64 = 4294967295
    if true {
        v5 = 409
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 37
}
func genFn1(p0 string, p1 uint, p2 uint64) int {
    switch p0 {
    case "a":
        fmt.Println(p1, p2, p0)
    case "\n\t":
        if false {
            switch 172 {
            case 488:
                for fi6 := 0; fi6 < 15; fi6++ {
                    var v7 int64 = 258
                    if true {
                        v7 = -9223372036854775808
                        fmt.Println(p1, v7)
                        if (!false) {
                            v8 := []int{(fi6 << (((p1 / ((1) | 1))) & 63)), fi6, (^(fi6))}
                            switch p0 {
                            case "a":
                                var v9 uint32 = 0
                                _ = v9
                            case "foo bar":
                                p0 = p0
                            case "\n\t":
                                var v10 bool = false
                                _ = v10
                            }
                            _ = v8
                        }
                    } else {
                        p1 = p1
                        switch (p0 + p0) {
                        case "\n\t":
                            if (!(!((101 != 0) || true))) {
                                if (776 < 2147483647) {
                                    var v11 uint = 255
                                    fmt.Println(fi6, p1)
                                    _ = v11
                                }
                                if ((fi6 & (fi6 + 0)) == fi6) {
                                    var v12 string = ""
                                    var v13 int64 = 9223372036854775807
                                    _ = v12
                                    _ = v13
                                }
                                v14 := GenS1{F0: 206, F1: "\n\t", F2: (255 <= 1)}
                                _ = v14
                            }
                        default:
                            var v15 string = p0
                            _ = v15
                        }
                    }
                    v16 := GenS1{F0: (p2 * (p2 & 243)), F1: "foo bar", F2: true}
                    _ = fi6
                    _ = v16
                    _ = v7
                }
            case 42:
                v17 := GenS1{F0: p2, F1: ((p0 + (p0 + p0)) + p0), F2: true}
                _ = v17
            default:
                var v18 uint8 = 103
                _ = v18
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                p2 = p2
                _ = fi19
            }
        }
    case "hello":
        switch ((("a" + p0) + p0) + p0) {
        case "a":
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 uint = (p1 & p1)
                var v22 int64 = -1
                _ = fi20
                _ = v21
                _ = v22
            }
        }
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            fmt.Println(fi23, fi24, fi23)
            if false {
                var v25 uint64 = (p2 / ((p2) | 1))
                if (((fi23 % ((-284) | 1)) >= 32767) && true) {
                    var v26 uint8 = 42
                    switch fi24 {
                    case -368:
                        v27 := GenS2{F0: 4294967295, F1: (((fi23 + fi23) > (fi23 % ((fi24) | 1))) && (fi23 >= fi24)), F2: p0}
                        var v28 string = ((p0 + p0) + (p0 + ("foo bar" + p0)))
                        var v29 int8 = -1
                        _ = v27
                        _ = v28
                        _ = v29
                    case 0:
                        if ((v26 + v26) < v26) {
                            p1 += p1
                            var v30 bool = ((1 != 34) || (false && true))
                            var v31 uint32 = 65535
                            _ = v30
                            _ = v31
                        }
                    default:
                        v32 := []int{(fi24 ^ 1), 2147483647}
                        _ = v32
                    }
                    v33 := GenS2{F0: (fi24 % (((fi23 * -2147483648)) | 1)), F1: true, F2: ((p0 + ((p0 + (p0 + ("foo bar" + (p0 + (p0 + ((p0 + p0) + "a")))))) + p0)) + p0)}
                    _ = v26
                    _ = v33
                } else {
                    if (!true) {
                        var v34 float32 = 0.0
                        _ = v34
                    }
                    fmt.Println(p2, p1)
                    var v35 float32 = (1.5 * (float32(-2.25) - 3.14159))
                    v25 = (v25 + p2)
                    _ = v35
                }
                _ = v25
            } else {
                var v36 int16 = -72
                v36 -= (v36 % (((v36 + v36)) | 1))
                v37 := []int{fi24}
                fmt.Println(v36)
                _ = v36
                _ = v37
            }
            p0 = (p0 + p0)
            v38 := GenS3{F0: p1, F1: 42}
            _ = fi24
            _ = v38
        }
        _ = fi23
    }
    var v39 string = "\n\t"
    v39 = v39
    for fi40 := 0; fi40 < 17; fi40++ {
        for fi41 := 0; fi41 < 8; fi41++ {
            if true {
                v42 := []int{}
                _ = v42
            }
            _ = fi41
        }
        p0 = p0
        _ = fi40
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    return 0
}

func main() {
    v1 := []uint32{1}
    var v2 int16 = -32768
    v3 := []uint32{94, 0, 4294967295, 42}
    for fi43 := 0; fi43 < 16; fi43++ {
        fmt.Println(fi43)
        fmt.Println(v2)
        _ = fi43
    }
    v2 += v2
    var v44 int64 = 1
    for fi45 := 0; fi45 < 17; fi45++ {
        if true {
            var v46 uint32 = 1
            v2 = 0
            if (!((3.14159 * 0.0) != 1e10)) {
                for fi47 := 0; fi47 < 9; fi47++ {
                    fmt.Println(v46, fi47, v2)
                    _ = fi47
                }
                var v48 uint32 = v46
                if ((float32((-2.25 * (((1e10 - 0.0) / 1e10) / 1e10))) == 3.14159) && false) {
                    fmt.Println(v2)
                }
                _ = v48
            } else {
                v2 = v2
                var v49 uint32 = v46
                _ = v49
            }
            _ = v46
        }
        _ = fi45
    }
    if (65535 <= 65535) {
        var v50 uint8 = 127
        var v51 uint64 = 4294967295
        v2 += (-(v2))
        _ = v50
        _ = v51
    }
    v3 = append(v3, 932)
    v2 -= v2
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
}
