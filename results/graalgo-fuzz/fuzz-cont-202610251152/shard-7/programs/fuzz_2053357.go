package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
func genFn0(p0 int64) uint64 {
    v5 := GenS1{F0: 4294967295, F1: (p0 >> ((0) & 63))}
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: (fi6 << ((1) & 63)), F1: p0}
        _ = fi6
        _ = v7
    }
    switch "hello" {
    case "":
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 uint32 = 4294967295
            _ = fi8
            _ = v9
        }
    case "a":
        var v10 uint64 = 4294967295
        _ = v10
    case "hello":
        var v11 uint8 = 1
        _ = v11
    }
    var v12 string = ""
    _ = p0
    _ = v12
    _ = v5
    return 912
}
func genFn1() uint8 {
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := []int{fi13, 128}
        _ = fi13
        _ = v14
    }
    if true {
        v15 := GenS1{F0: 32768, F1: -1}
        var v16 bool = true
        _ = v15
        _ = v16
    }
    if ("hello" == ("a" + "a")) {
        if ("foo bar" >= "\n\t") {
            var v17 uint = 1
            _ = v17
        } else {
            var v18 uint = 255
            v19 := []int{2147483647, 1, 769, -1}
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 int32 = -747
            v22 := GenS1{F0: (^(fi20)), F1: -9223372036854775808}
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    return 255
}

func main() {
    var v1 int64 = -564
    var v2 int8 = 42
    v3 := make(map[uint32]bool)
    v4 := make(map[string]uint32)
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 float64 = 1e10
        var v25 int32 = 586
        fmt.Println(fi23, fi23, fi23)
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 int8 = -80
    v27 := GenS1{F0: 9223372036854775807, F1: -883}
    if (false || true) {
        var v28 float64 = 0.1
        _ = v28
    }
    if false {
        for fi29 := 0; fi29 < 20; fi29++ {
            _ = len(v3)
            var v30 int = fi29
            _ = fi29
            _ = v30
        }
        v3[520] = true
    } else {
        v2 += v2
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 uint = 2
            var v33 float64 = 1.5
            v26 += -128
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    v4["a"] = 1
    for fi34 := 0; fi34 < 5; fi34++ {
        fmt.Println(v1, fi34, v2)
        _ = fi34
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        switch (^(fi35)) {
        case 0:
            switch (fi35 ^ 32768) {
            case -1:
                var v36 float32 = 1e10
                _ = v36
            case 0:
                v1 = (v1 & v1)
            }
        case 924:
            for fi37 := 0; fi37 < 6; fi37++ {
                var v38 int = 42
                fmt.Println(v1)
                _ = fi37
                _ = v38
            }
        case 9223372036854775807:
            v39 := GenS1{F0: (fi35 << ((4294967295) & 63)), F1: (v1 & v1)}
            _ = v39
        default:
            _ = len(v4)
        }
        fmt.Println(fi35)
        v1 -= 1
        _ = fi35
    }
    var v40 uint = 42
    for fi41 := 0; fi41 < 3; fi41++ {
        var v42 int = fi41
        _ = fi41
        _ = v42
    }
    v43 := []int{42}
    if ("" > ("\n\t" + ("foo bar" + ""))) {
        v44 := GenS1{F0: -634, F1: (v1 & -9223372036854775808)}
        _ = v44
    } else {
        if (272 == -9223372036854775808) {
            if (0 == 1) {
                var v45 float64 = 1.5
                v46 := v4["a"]
                _ = v45
                _ = v46
            } else {
                v47 := GenS1{F0: 959, F1: v1}
                _ = v47
            }
        }
        fmt.Println(v26, v2, v40)
        for fi48 := 0; fi48 < 4; fi48++ {
            if (!true) {
                v1 += (v1 << ((v40) & 63))
            } else {
                var v49 uint8 = 0
                _ = v49
            }
            var v50 uint64 = 18446744073709551615
            _ = fi48
            _ = v50
        }
    }
    switch "foo bar" {
    case "foo bar":
        v2 = v26
    case "\n\t":
        fmt.Println(v2, v1, v40)
    case "hello":
        if true {
            for fi51 := 0; fi51 < 9; fi51++ {
                switch "\n\t" {
                case "a":
                    switch fi51 {
                    case 946:
                        var v52 uint8 = 1
                        _ = v52
                    }
                default:
                    switch ((("\n\t" + "a") + "") + (("foo bar" + "\n\t") + "\n\t")) {
                    case "foo bar":
                        v26 = -128
                    default:
                        for fi53 := 0; fi53 < 9; fi53++ {
                            fmt.Println(v1, v26, v1)
                            if true {
                                v40 += v40
                                switch (fi53 | fi51) {
                                case 2147483647:
                                    v26 += v26
                                case 32768:
                                    v1 = v1
                                }
                            }
                            var v54 uint8 = 255
                            _ = fi53
                            _ = v54
                        }
                    }
                }
                var v55 uint8 = genFn1()
                _ = fi51
                _ = v55
            }
        } else {
            var v56 uint32 = 2147483648
            v40 = 4294967295
            _ = v56
        }
    }
    var v57 float32 = 3.14159
    _ = len(v43)
    for fi58 := 0; fi58 < 19; fi58++ {
        switch (fi58 % (((^(fi58))) | 1)) {
        case 32767:
            v59 := GenS1{F0: fi58, F1: 42}
            _ = v59
        default:
            if true {
                for fi60 := 0; fi60 < 13; fi60++ {
                    fmt.Println(fi60)
                    v61 := GenS1{F0: 0, F1: v1}
                    var v62 int8 = v26
                    fmt.Println(v26, v26)
                    fmt.Println(v57, fi60)
                    _ = fi60
                    _ = v61
                    _ = v62
                }
            }
        }
        _ = fi58
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v57
}
