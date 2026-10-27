package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint64 }
func genFn0(p0 uint64, p1 uint64, p2 uint8) int32 {
    var v5 uint64 = (^(p0))
    for fi6 := 0; fi6 < 2; fi6++ {
        switch (fi6 | (fi6 | fi6)) {
        case 32767:
            switch -148 {
            case 32768:
                fmt.Println(v5, v5)
            case 42:
                v7 := GenS1{F0: 1, F1: ("a" + ((("a" + "a") + "foo bar") + ("foo bar" + "hello"))), F2: 638}
                _ = v7
            }
        default:
            switch "\n\t" {
            case "foo bar":
                var v8 int8 = 1
                _ = v8
            case "\n\t":
                for fi9 := 0; fi9 < 19; fi9++ {
                    p1 = 2
                    _ = fi9
                }
            }
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 uint8 = 1
            v12 := GenS1{F0: 0, F1: ("" + "\n\t"), F2: 9223372036854775807}
            var v13 float32 = float32(1e10)
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 391
}
func genFn1(p0 uint32) int {
    var v14 int = 0
    v15 := []int{(v14 ^ (v14 % ((32767) | 1)))}
    v14 += (v14 ^ v14)
    for fi16 := 0; fi16 < 9; fi16++ {
        p0 = (p0 << ((187) & 63))
        switch (fi16 + (fi16 << ((255) & 63))) {
        case -1:
            if (930 > 817) {
                v17 := GenS1{F0: 877, F1: "\n\t", F2: 511}
                v14 -= v14
                _ = v17
            } else {
                v14 = (v14 - (v14 - fi16))
                var v18 int8 = 42
                v19 := GenS2{F0: 42}
                _ = v18
                _ = v19
            }
        case -546:
            if (-1 == 1000) {
                var v20 int64 = -9223372036854775808
                var v21 bool = true
                fmt.Println(v21, p0, p0)
                _ = v20
                _ = v21
            } else {
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                p0 += 42
            }
        default:
            p0 -= p0
        }
        _ = fi16
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        for fi23 := 0; fi23 < 15; fi23++ {
            fmt.Println(fi22, v14, fi23)
            _ = fi23
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := GenS1{F0: 32767, F1: (("hello" + "hello") + ""), F2: -179}
            _ = fi24
            _ = v25
        }
        _ = fi22
    }
    _ = p0
    _ = v14
    _ = v15
    return (v14 >> ((1) & 63))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    v3 := []int{4294967295, 384, 2147483647}
    v4 := []int{-325, -1, -9223372036854775808}
    v26 := GenS3{F0: ("" + "a"), F1: 770}
    if true {
        var v27 uint8 = 0
        _ = v27
    }
    var v28 int32 = -2147483648
    if true {
        v28 = (v28 << ((18446744073709551615) & 63))
    }
    if ((2147483648 < 42) && (true && false)) {
        if false {
            v28 = (v28 ^ genFn0(1, 4294967295, 2))
            for fi29 := 0; fi29 < 5; fi29++ {
                fmt.Println(v28, fi29)
                for fi30 := 0; fi30 < 11; fi30++ {
                    var v31 uint32 = 787
                    var v32 int32 = v28
                    var v33 int32 = (v32 & v32)
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                if len(v4) > 0 {
                    v4[0] = fi29
                }
                _ = fi29
            }
        } else {
            v34 := GenS3{F0: "foo bar", F1: 2}
            v4 = append(v4, 127)
            _ = v34
        }
    }
    v28 = v28
    v28 -= (v28 ^ -2147483648)
    for fi35 := 0; fi35 < 17; fi35++ {
        switch fi35 {
        case -1:
            v28 += 746
        default:
            v28 -= (v28 * (v28 / (((-(v28))) | 1)))
        }
        switch "hello" {
        case "\n\t":
            v28 -= v28
        }
        _ = fi35
    }
    fmt.Println(v28)
    if false {
        if false {
            var v36 uint32 = 436
            var v37 string = ((("\n\t" + "hello") + ("hello" + "\n\t")) + "foo bar")
            if (!(((v36 - v36) < (v36 ^ (v36 ^ 441))) && (19 < 1))) {
                var v38 int32 = (v28 | v28)
                if true {
                    v39 := GenS1{F0: 1, F1: v37, F2: 1}
                    _ = len(v2)
                    _ = v39
                } else {
                    var v40 uint64 = 883
                    var v41 int = 2147483647
                    var v42 uint64 = v40
                    _ = v40
                    _ = v41
                    _ = v42
                }
                v43 := GenS1{F0: 42, F1: v37, F2: 92}
                _ = v38
                _ = v43
            }
            _ = v36
            _ = v37
        }
        if (1000 != 42) {
            if (false || (v28 == (v28 * v28))) {
                fmt.Println(v28)
                v28 += (v28 % ((v28) | 1))
                var v44 uint64 = 42
                _ = v44
            } else {
                for fi45 := 0; fi45 < 5; fi45++ {
                    var v46 uint32 = 2147483648
                    _ = fi45
                    _ = v46
                }
                delete(v1, 0)
                switch ("a" + "hello") {
                case "foo bar":
                    v28 = genFn0(433, 2, 127)
                case "a":
                    v28 = v28
                }
            }
            if (!(255 > 2)) {
                fmt.Println(v28)
                fmt.Println(v28)
            } else {
                v47 := v1[-9223372036854775808]
                v28 = 2147483647
                _ = v47
            }
        }
    } else {
        if true {
            if false {
                if len(v3) > 0 {
                    v3[0] = 1
                }
                v28 -= v28
            }
        } else {
            v28 = v28
            fmt.Println(v28)
            v28 = (-(v28))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
}
