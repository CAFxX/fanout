package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint; F2 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 bool; F2 uint }

func main() {
    v1 := []int{4294967295, 32768, 32767, 32768}
    var v2 uint64 = 161
    var v3 uint16 = 255
    fmt.Println(v2, v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2, v3)
    v2 = v2
    v3 += (v3 << ((18446744073709551615) & 63))
    switch 0 {
    case 32767:
        if true {
            for fi4 := 0; fi4 < 9; fi4++ {
                var v5 int = int(v2)
                if (!(true && true)) {
                    switch "a" {
                    case "hello":
                        if (0 >= 9223372036854775807) {
                            var v6 int16 = 745
                            _ = v6
                        }
                    case "\n\t":
                        v7 := GenS1{F0: -860}
                        _ = v7
                    }
                } else {
                    v8 := GenS1{F0: -1}
                    switch "a" {
                    case "":
                        if (1000 == -1) {
                            v9 := GenS1{F0: -958}
                            _ = v9
                        }
                    case "hello":
                        fmt.Println(v2, v5)
                    case "\n\t":
                        var v10 int16 = int16(fi4)
                        _ = v10
                    }
                    _ = v8
                }
                _ = fi4
                _ = v5
            }
        }
    case 4294967295:
        for fi11 := 0; fi11 < 17; fi11++ {
            if true {
                v1 = append(v1, fi11)
                v3 = v3
            } else {
                v12 := GenS2{F0: "foo bar", F1: 179, F2: 127}
                _ = v12
            }
            for fi13 := 0; fi13 < 16; fi13++ {
                if true {
                    v2 -= (v2 ^ v2)
                    _ = len(v1)
                    _ = len(v1)
                } else {
                    v1 = append(v1, (fi13 * fi11))
                    var v14 uint16 = (v3 + v3)
                    var v15 uint64 = (v2 & (v2 % (((v2 << ((403) & 63))) | 1)))
                    _ = v14
                    _ = v15
                }
                if false {
                    v3 -= v3
                    var v16 int16 = -32768
                    _ = v16
                } else {
                    v17 := GenS2{F0: "a", F1: 42, F2: 100}
                    var v18 int8 = 20
                    switch "a" {
                    case "hello":
                        var v19 int32 = -884
                        _ = v19
                    default:
                        v2 = (v2 - v2)
                    }
                    _ = v17
                    _ = v18
                }
                var v20 uint = 42
                _ = fi13
                _ = v20
            }
            _ = fi11
        }
    }
    v1 = append(v1, 4294967295)
    var v21 int16 = -139
    fmt.Println(v3, v2, v3)
    var v22 uint32 = 65535
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 int64 = 9223372036854775807
        for fi25 := 0; fi25 < 1; fi25++ {
            switch (fi25 | fi25) {
            case -253:
                v26 := GenS3{F0: 711}
                _ = v26
            case -280:
                var v27 int16 = v21
                _ = v27
            case 32768:
                var v28 int64 = v24
                _ = v28
            }
            var v29 uint32 = (v22 / ((v22) | 1))
            var v30 int8 = -1
            var v31 float64 = -2.25
            fmt.Println(v30)
            _ = fi25
            _ = v29
            _ = v30
            _ = v31
        }
        if true {
            switch ("hello" + "hello") {
            case "a":
                for fi32 := 0; fi32 < 12; fi32++ {
                    if len(v1) > 0 {
                        v1[0] = fi23
                    }
                    switch (-(fi23)) {
                    case -9223372036854775808:
                        v24 = (^(v24))
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    fmt.Println(fi32)
                    _ = fi32
                }
            case "foo bar":
                for fi33 := 0; fi33 < 9; fi33++ {
                    v24 += (v24 >> ((1) & 63))
                    v34 := GenS1{F0: 496}
                    _ = fi33
                    _ = v34
                }
            default:
                fmt.Println(v2, v2, v2)
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v35 int8 = 69
            v35 -= v35
            _ = v35
        }
        _ = fi23
        _ = v24
    }
    switch ("hello" + ("a" + "\n\t")) {
    case "foo bar":
        for fi36 := 0; fi36 < 16; fi36++ {
            var v37 uint = 1
            v2 += 1
            _ = fi36
            _ = v37
        }
    }
    if false {
        fmt.Println(v21, v3, v21)
    }
    v22 = v22
    for fi38 := 0; fi38 < 8; fi38++ {
        var v39 string = ("hello" + "\n\t")
        if (false && (717 != -116)) {
            switch 140 {
            case 9223372036854775807:
                var v40 int = 4294967295
                _ = v40
            case 32767:
                for fi41 := 0; fi41 < 1; fi41++ {
                    v3 = (v3 ^ (v3 * (v3 | uint16(v22))))
                    _ = fi41
                }
            default:
                v42 := GenS2{F0: "hello", F1: 529, F2: -128}
                _ = v42
            }
            var v43 uint64 = v2
            if len(v1) > 0 {
                v1[0] = (fi38 & (fi38 & 9223372036854775807))
            }
            _ = v43
        } else {
            switch "a" {
            case "foo bar":
                v1 = append(v1, fi38)
            case "a":
                for fi44 := 0; fi44 < 11; fi44++ {
                    var v45 int32 = -467
                    if true {
                        if true {
                            fmt.Println(v22)
                        }
                        var v46 string = v39
                        var v47 int8 = 103
                        var v48 uint16 = 640
                        _ = v46
                        _ = v47
                        _ = v48
                    } else {
                        var v49 uint32 = 42
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v49
                    }
                    v21 = v21
                    if ((-(v21)) <= (v21 << ((1) & 63))) {
                        if (!true) {
                            v3 -= v3
                        } else {
                            var v50 float64 = (1.5 * 1e10)
                            _ = v50
                        }
                        v22 += v22
                    }
                    _ = fi44
                    _ = v45
                }
            }
            v39 = v39
            v21 = (-(v21))
        }
        if (!(v21 >= 42)) {
            var v51 int16 = v21
            _ = v51
        } else {
            if false {
                if (!true) {
                    v52 := GenS4{F0: true}
                    fmt.Println(v22, v39)
                    _ = v52
                } else {
                    switch fi38 {
                    case 128:
                        var v53 int8 = -73
                        _ = v53
                    case -1:
                        var v54 uint16 = 224
                        _ = v54
                    }
                    if (4294967295 < 42) {
                        for fi55 := 0; fi55 < 3; fi55++ {
                            v3 = (v3 * (v3 & (v3 - (^(v3)))))
                            v2 = (^(v2))
                            _ = fi55
                        }
                        v21 += v21
                        v3 = 255
                    } else {
                        v22 = 2147483647
                        switch (fi38 >> ((125) & 63)) {
                        case 4294967295:
                            v56 := GenS5{F0: v2, F1: true, F2: 529}
                            _ = v56
                        }
                        var v57 uint64 = uint64(v22)
                        _ = v57
                    }
                    var v58 float32 = 1e10
                    _ = v58
                }
            } else {
                var v59 uint = 42
                if (!true) {
                    var v60 int = fi38
                    _ = v60
                }
                _ = v59
            }
            var v61 int = 32768
            _ = v61
        }
        _ = fi38
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
