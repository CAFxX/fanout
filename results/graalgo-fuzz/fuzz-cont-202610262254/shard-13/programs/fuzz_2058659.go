package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 bool; F1 int32; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int32; F1 bool; F2 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64; F1 string; F2 uint32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint8; F1 int64 }
type GenS10 struct { F0 string }
func genFn0(p0 uint8) string {
    p0 = p0
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := []int{(fi4 + 9223372036854775807), 32767}
        var v6 bool = false
        _ = fi4
        _ = v5
        _ = v6
    }
    switch -9223372036854775808 {
    case 42:
        var v7 uint = 0
        _ = v7
    case 32768:
        if (818 > 112) {
            for fi8 := 0; fi8 < 15; fi8++ {
                v9 := GenS1{F0: 42}
                if false {
                    var v10 uint8 = 128
                    fmt.Println(fi8)
                    _ = v10
                } else {
                    switch 1 {
                    case 42:
                        var v11 int16 = -1
                        _ = v11
                    }
                    fmt.Println(p0, p0)
                    v12 := GenS2{F0: 18446744073709551615, F1: p0}
                    _ = v12
                }
                _ = fi8
                _ = v9
            }
            for fi13 := 0; fi13 < 13; fi13++ {
                if true {
                    v14 := []int{(^(fi13)), (fi13 | fi13), 545, fi13}
                    _ = v14
                }
                for fi15 := 0; fi15 < 13; fi15++ {
                    if (2147483647 == -2147483648) {
                        p0 = (p0 * p0)
                        v16 := GenS3{F0: (42 != 1), F1: 701, F2: 949}
                        var v17 float64 = 1.5
                        _ = v16
                        _ = v17
                    } else {
                        var v18 int32 = 0
                        switch fi13 {
                        case -9223372036854775808:
                            var v19 int32 = 42
                            v20 := []int{fi15, fi15, fi15}
                            _ = v19
                            _ = v20
                        }
                        v18 = (v18 * (v18 + (v18 << ((561) & 63))))
                        var v21 uint16 = 65535
                        _ = v18
                        _ = v21
                    }
                    var v22 uint8 = 5
                    _ = fi15
                    _ = v22
                }
                _ = fi13
            }
        } else {
            switch "foo bar" {
            case "\n\t":
                p0 = 127
            }
            for fi23 := 0; fi23 < 19; fi23++ {
                v24 := GenS3{F0: false, F1: -1, F2: 2147483648}
                _ = fi23
                _ = v24
            }
            if true {
                p0 = 170
            }
        }
    }
    _ = p0
    return ("" + "foo bar")
}
func genFn1(p0 uint, p1 int32) uint32 {
    var v25 string = ("hello" + "a")
    for fi26 := 0; fi26 < 2; fi26++ {
        fmt.Println(fi26)
        for fi27 := 0; fi27 < 6; fi27++ {
            if true {
                if false {
                    if true {
                        switch fi26 {
                        case -366:
                            v28 := []int{fi26}
                            _ = v28
                        case -9223372036854775808:
                            switch (fi26 ^ (fi26 - fi26)) {
                            case 127:
                                var v29 int32 = 42
                                _ = v29
                            case 695:
                                v30 := GenS4{F0: (true || (-2.25 < 0.0)), F1: 1}
                                _ = v30
                            }
                        default:
                            fmt.Println(p0, p0)
                        }
                    }
                    var v31 uint8 = 255
                    if (32767 == fi26) {
                        var v32 uint32 = 577
                        var v33 int8 = 0
                        _ = v32
                        _ = v33
                    } else {
                        fmt.Println(fi27)
                        var v34 bool = (!(!true))
                        var v35 uint16 = 255
                        _ = v34
                        _ = v35
                    }
                    _ = v31
                }
            } else {
                switch (fi27 & (fi26 & fi26)) {
                case 9223372036854775807:
                    p0 += p0
                }
                v25 = v25
                p1 -= -2147483648
            }
            var v36 uint = (p0 << ((p0) & 63))
            switch (fi27 & (fi27 ^ (fi27 % ((fi27) | 1)))) {
            case 42:
                if true {
                    v37 := GenS5{F0: (p1 >> ((v36) & 63)), F1: false, F2: -2147483648}
                    v25 = "a"
                    var v38 int32 = 1
                    _ = v37
                    _ = v38
                } else {
                    switch ("\n\t" + ("\n\t" + v25)) {
                    case "a":
                        fmt.Println(v25, v36, fi27)
                        var v39 uint64 = 121
                        _ = v39
                    case "":
                        v40 := []int{(fi27 ^ (fi27 ^ -942)), fi27, 607}
                        _ = v40
                    default:
                        fmt.Println(v25, fi27, p1)
                    }
                    p0 = 4294967295
                    if (p1 == (p1 * (p1 * p1))) {
                        fmt.Println(p1, p1, p1)
                        var v41 int16 = -607
                        _ = v41
                    } else {
                        v42 := []int{}
                        fmt.Println(p1, fi27, p0)
                        v43 := GenS6{F0: -973}
                        _ = v42
                        _ = v43
                    }
                }
            }
            _ = fi27
            _ = v36
        }
        _ = fi26
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        switch (v25 + v25) {
        case "foo bar":
            p1 -= 42
        case "":
            p0 -= p0
        default:
            var v45 bool = true
            _ = v45
        }
        if ((p1 >= (p1 >> ((p0) & 63))) && (!true)) {
            for fi46 := 0; fi46 < 20; fi46++ {
                var v47 uint8 = 127
                v47 -= (v47 % ((v47) | 1))
                _ = fi46
                _ = v47
            }
            var v48 string = ("hello" + v25)
            _ = v48
        }
        if true {
            v25 = (v25 + (("a" + ("" + "")) + "\n\t"))
        }
        _ = fi44
    }
    var v49 int8 = -114
    _ = p0
    _ = p1
    _ = v25
    _ = v49
    return 0
}

func main() {
    var v1 int32 = 395
    v2 := make(map[int]int)
    var v3 uint16 = 2
    var v50 float32 = float32(0.0)
    v50 = v50
    if (957 >= 1) {
        v51 := GenS7{F0: 759, F1: (("" + "a") + "hello"), F2: genFn1(1, 42)}
        for fi52 := 0; fi52 < 2; fi52++ {
            for fi53 := 0; fi53 < 20; fi53++ {
                fmt.Println(v1, fi52, fi53)
                var v54 uint32 = 2147483648
                _ = fi53
                _ = v54
            }
            switch (fi52 ^ fi52) {
            case 596:
                for fi55 := 0; fi55 < 17; fi55++ {
                    if true {
                        fmt.Println(v50, v50, fi55)
                    }
                    v56 := GenS8{F0: -554}
                    v1 -= (^(v1))
                    _ = fi55
                    _ = v56
                }
            }
            delete(v2, (fi52 * 0))
            _ = fi52
        }
        _ = v51
    }
    if (!((false || true) && true)) {
        fmt.Println(v1, v50)
        switch "\n\t" {
        case "\n\t":
            var v57 int64 = 578
            _ = v57
        default:
            v58 := GenS9{F0: 119, F1: -1}
            _ = v58
        }
    }
    v1 -= v1
    v59 := GenS3{F0: ((0.0 + 0.1) >= 1.5), F1: (^(v1)), F2: 708}
    v3 += v3
    v60 := v2[42]
    v61 := []int{}
    switch "hello" {
    case "hello":
        var v62 uint32 = 2147483647
        _ = v62
    }
    var v63 float32 = float32(1.5)
    if (!true) {
        v64 := GenS1{F0: 0}
        v60 += v60
        v65 := v2[v60]
        _ = v64
        _ = v65
    } else {
        v50 = v63
        var v66 int64 = 0
        var v67 int64 = (v66 - (v66 & 195))
        _ = v66
        _ = v67
    }
    delete(v2, (v60 % (((v60 + (v60 & v60))) | 1)))
    v63 = v63
    v68 := GenS10{F0: "hello"}
    switch ("" + "foo bar") {
    case "\n\t":
        v69 := GenS2{F0: 42, F1: 1}
        _ = v69
    default:
        var v70 uint8 = 101
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v59
    _ = v60
    _ = v61
    _ = v63
    _ = v68
}
