package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint; F1 uint64; F2 int16 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 7; fi4++ {
        switch ((("hello" + "") + "a") + "a") {
        case "":
            v5 := []int{fi4, (fi4 % (((fi4 | 2147483647)) | 1)), fi4}
            _ = v5
        case "hello":
            switch "\n\t" {
            case "hello":
                var v6 string = "foo bar"
                _ = v6
            }
        case "a":
            switch -1 {
            case 4294967295:
                var v7 int16 = -1
                _ = v7
            case -164:
                for fi8 := 0; fi8 < 19; fi8++ {
                    var v9 int8 = -124
                    _ = fi8
                    _ = v9
                }
            case 32768:
                v10 := GenS1{F0: 42}
                _ = v10
            default:
                var v11 string = "\n\t"
                _ = v11
            }
        default:
            var v12 float32 = 3.14159
            _ = v12
        }
        v13 := []int{(fi4 << ((206) & 63)), (fi4 * -2147483648), fi4}
        _ = len(v13)
        _ = fi4
        _ = v13
    }
    switch 32768 {
    case 127:
        var v14 int64 = -989
        _ = v14
    case 365:
        switch "foo bar" {
        case "\n\t":
            var v15 uint8 = 36
            _ = v15
        default:
            var v16 float64 = 0.1
            _ = v16
        }
    case 32767:
        var v17 uint8 = 1
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        if true {
            if true {
                if false {
                    v19 := []int{(fi18 & 544), (fi18 / (((fi18 & (fi18 | fi18))) | 1)), fi18, fi18}
                    fmt.Println(fi18)
                    for fi20 := 0; fi20 < 7; fi20++ {
                        if true {
                            var v21 uint = 167
                            _ = v21
                        } else {
                            if (161 < 42) {
                                var v22 int64 = 0
                                _ = v22
                            } else {
                                var v23 uint32 = 0
                                _ = v23
                            }
                            var v24 uint32 = 2147483647
                            _ = v24
                        }
                        _ = fi20
                    }
                    _ = v19
                } else {
                    if true {
                        for fi25 := 0; fi25 < 18; fi25++ {
                            v26 := []int{(fi18 % ((fi18) | 1)), 1, 1}
                            v26 = append(v26, -1)
                            _ = fi25
                            _ = v26
                        }
                    } else {
                        var v27 string = "hello"
                        v27 = (v27 + v27)
                        _ = v27
                    }
                }
                v28 := GenS1{F0: 2}
                var v29 int32 = 2147483647
                _ = v28
                _ = v29
            }
        }
        for fi30 := 0; fi30 < 15; fi30++ {
            v31 := []int{-7, 140, (fi30 ^ (fi18 ^ fi18))}
            var v32 float32 = 0.1
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi18
    }
    return 255
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{127, 9223372036854775807, 1, -389}
    var v3 uint16 = 1
    v33 := GenS2{F0: 2}
    switch -150 {
    case 32768:
        var v34 uint32 = 2
        _ = v34
    case 357:
        for fi35 := 0; fi35 < 8; fi35++ {
            fmt.Println(v3, fi35)
            v36 := GenS1{F0: 0}
            _ = fi35
            _ = v36
        }
    case 32767:
        v3 -= v3
    default:
        fmt.Println(v3)
    }
    v3 -= v3
    var v37 int64 = 376
    v38 := GenS1{F0: 42}
    v37 = -1
    v39 := GenS2{F0: genFn0()}
    v37 += 9223372036854775807
    v40 := GenS2{F0: uint8(v3)}
    switch -1 {
    case 2147483647:
        v41 := GenS3{F0: 4294967295, F1: 1}
        _ = v41
    case 0:
        switch "" {
        case "\n\t":
            v3 = 0
        case "":
            switch "foo bar" {
            case "hello":
                var v42 int16 = -1
                _ = v42
            case "a":
                _ = len(v2)
            case "\n\t":
                fmt.Println(v37, v3)
            }
        case "a":
            var v43 int16 = -1
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 10; fi44++ {
        if true {
            switch (fi44 & (fi44 & -1)) {
            case -890:
                v3 -= (v3 >> ((18446744073709551615) & 63))
            case -1:
                v45 := GenS4{F0: 18446744073709551615, F1: 777, F2: -929}
                _ = v45
            }
            fmt.Println(v37)
        }
        v1["\n\t"] = fi44
        for fi46 := 0; fi46 < 17; fi46++ {
            switch "hello" {
            case "a":
                var v47 int = (fi46 << ((355) & 63))
                _ = len(v1)
                var v48 uint16 = v3
                var v49 int = (v47 ^ fi46)
                var v50 int = fi44
                _ = v47
                _ = v48
                _ = v49
                _ = v50
            case "":
                fmt.Println(v37, fi44)
            }
            var v51 int8 = 1
            _ = fi46
            _ = v51
        }
        _ = fi44
    }
    if (1.5 != -2.25) {
        v52 := GenS3{F0: 539, F1: 128}
        var v53 string = "\n\t"
        var v54 uint8 = 0
        _ = v52
        _ = v53
        _ = v54
    }
    v55 := GenS3{F0: 4294967295, F1: 3}
    for fi56 := 0; fi56 < 15; fi56++ {
        v57 := GenS4{F0: 655, F1: 18446744073709551615, F2: 32767}
        fmt.Println(v37, v3, fi56)
        if true {
            var v58 uint64 = 18446744073709551615
            delete(v1, "a")
            _ = v58
        } else {
            if (65535 == 4294967295) {
                var v59 uint32 = 1
                _ = v59
            }
        }
        _ = fi56
        _ = v57
    }
    var v60 int8 = 42
    var v61 uint = 65535
    v37 = v37
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v55
    _ = v60
    _ = v61
}
