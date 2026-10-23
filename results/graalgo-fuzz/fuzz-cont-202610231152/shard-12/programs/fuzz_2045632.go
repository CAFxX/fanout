package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint8; F2 bool }
type GenS4 struct { F0 uint64; F1 uint16; F2 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 122
    v3 := make(map[uint32]uint32)
    var v4 int8 = -128
    var v5 int16 = 42
    var v6 uint8 = (v2 >> ((0) & 63))
    v2 = v6
    fmt.Println(v5, v2, v4)
    if false {
        switch 2147483647 {
        case -364:
            for fi7 := 0; fi7 < 17; fi7++ {
                if (0 != v5) {
                    if (!true) {
                        v8 := GenS1{F0: (v5 / ((-509) | 1))}
                        _ = v8
                    }
                } else {
                    v9 := []int{(fi7 | (fi7 - 0)), -1}
                    var v10 string = "\n\t"
                    _ = v10
                    _ = v9
                }
                if false {
                    if false {
                        v2 += (^(v6))
                        v11 := GenS2{F0: fi7, F1: (-(fi7)), F2: 2}
                        _ = v11
                    }
                } else {
                    var v12 uint16 = 0
                    v13 := []int{fi7}
                    _ = v12
                    _ = v13
                }
                _ = fi7
            }
        case -9223372036854775808:
            v6 -= v6
        case 128:
            v6 = v6
        default:
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 uint64 = 18446744073709551615
                _ = fi14
                _ = v15
            }
        }
        var v16 uint = 159
        var v17 uint64 = 4294967295
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 11; fi18++ {
            if ((0.1 * 1e10) >= 0.0) {
                fmt.Println(v6)
                v19 := []int{(fi18 + (fi18 + (fi18 - (fi18 ^ (fi18 & fi18))))), fi18, (-(fi18)), 271}
                if true {
                    for fi20 := 0; fi20 < 18; fi20++ {
                        v21 := GenS3{F0: 42, F1: v6, F2: ("\n\t" != "hello")}
                        var v22 string = (("foo bar" + "hello") + "a")
                        switch 2147483647 {
                        case 4294967295:
                            v4 = -128
                        }
                        switch "" {
                        case "\n\t":
                            if false {
                                if (1.5 < 1.5) {
                                    var v23 uint16 = 255
                                    if len(v19) > 0 {
                                        v19[0] = (-(fi20))
                                    }
                                    _ = v23
                                } else {
                                    fmt.Println(v2, fi20, v5)
                                }
                            }
                        case "foo bar":
                            var v24 float64 = 3.14159
                            _ = v24
                        }
                        var v25 uint64 = 42
                        v5 = v5
                        v22 = v22
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v25
                    }
                    v2 += 0
                } else {
                    var v26 uint16 = 42
                    var v27 uint8 = v6
                    _ = v26
                    _ = v27
                }
                _ = v19
            }
            v28 := GenS2{F0: (fi18 | fi18), F1: fi18, F2: 4294967295}
            for fi29 := 0; fi29 < 13; fi29++ {
                fmt.Println(fi29, v4)
                var v30 uint16 = 1
                v31 := v3[132]
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi18
            _ = v28
        }
    }
    var v32 int32 = -1
    fmt.Println(v2, v5, v4)
    for fi33 := 0; fi33 < 3; fi33++ {
        switch "hello" {
        case "hello":
            fmt.Println(v32, fi33, v2)
        default:
            switch (("foo bar" + "foo bar") + "\n\t") {
            case "":
                var v34 uint8 = v2
                _ = v34
            case "foo bar":
                v5 -= v5
            }
        }
        switch fi33 {
        case -9223372036854775808:
            var v35 int64 = -1
            _ = v35
        }
        fmt.Println(v6, v4)
        _ = fi33
    }
    v2 -= (v2 - 46)
    v6 = v2
    fmt.Println(v32, v32, v32)
    for fi36 := 0; fi36 < 1; fi36++ {
        if (!(v4 >= int8(v5))) {
            v6 -= 1
        }
        v37 := []int{fi36}
        _ = fi36
        _ = v37
    }
    v38 := []int{431}
    for fi39 := 0; fi39 < 16; fi39++ {
        v40 := GenS4{F0: 0, F1: 0, F2: int8(v2)}
        for fi41 := 0; fi41 < 10; fi41++ {
            var v42 uint16 = 0
            _ = fi41
            _ = v42
        }
        _ = fi39
        _ = v40
    }
    switch 9223372036854775807 {
    case 2147483647:
        v4 = int8(v32)
    default:
        for fi43 := 0; fi43 < 1; fi43++ {
            for fi44 := 0; fi44 < 9; fi44++ {
                v45 := v3[0]
                var v46 uint64 = 0
                var v47 int64 = 42
                _ = fi44
                _ = v45
                _ = v46
                _ = v47
            }
            _ = fi43
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
