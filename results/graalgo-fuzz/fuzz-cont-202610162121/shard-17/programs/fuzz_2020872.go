package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 int64 }
func genFn0() int16 {
    var v6 uint32 = 355
    if (!true) {
        for fi7 := 0; fi7 < 17; fi7++ {
            v6 = (v6 & 655)
            for fi8 := 0; fi8 < 15; fi8++ {
                v6 = 2
                _ = fi8
            }
            v9 := GenS1{F0: "hello", F1: v6, F2: -838}
            _ = fi7
            _ = v9
        }
        var v10 bool = (true || false)
        switch 0 {
        case 426:
            var v11 uint16 = 200
            _ = v11
        case 873:
            v6 += v6
        case 4294967295:
            switch 235 {
            case 0:
                v10 = true
            }
        }
        _ = v10
    } else {
        var v12 uint = 858
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v6 += v6
        v6 = 489
        for fi14 := 0; fi14 < 1; fi14++ {
            v6 += 2
            _ = fi14
        }
        _ = fi13
    }
    v6 -= v6
    _ = v6
    return 1
}

func main() {
    var v1 float64 = 0.0
    var v2 uint32 = 42
    v3 := make(map[string]bool)
    var v4 int16 = -1
    v5 := make(map[uint32]string)
    v15 := GenS1{F0: ("" + "\n\t"), F1: 353, F2: 251}
    switch "\n\t" {
    case "":
        fmt.Println(v4, v1, v1)
    case "\n\t":
        if false {
            v16 := v3[""]
            v17 := []int{-2147483648}
            var v18 int64 = 42
            _ = v16
            _ = v17
            _ = v18
        } else {
            var v19 int = -1
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 bool = (((v4 - v4) <= genFn0()) || false)
                for fi22 := 0; fi22 < 14; fi22++ {
                    v23 := GenS1{F0: ("" + "hello"), F1: v2, F2: 940}
                    var v24 uint32 = (v2 | 2147483647)
                    if ((v19 | 32767) <= v19) {
                        v25 := GenS2{F0: "foo bar"}
                        var v26 bool = (!(!(1 == 560)))
                        v27 := GenS3{F0: v26}
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            var v28 int32 = 0
                            _ = v28
                        case "\n\t":
                            v24 = v2
                        }
                        if (154 > 2) {
                            var v29 bool = v21
                            _ = len(v3)
                            _ = v29
                        }
                    }
                    v21 = false
                    v30 := []int{fi20, (fi22 ^ fi20), fi22}
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v30
                }
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    }
    if (!false) {
        var v31 int = -1
        _ = v31
    } else {
        v1 = v1
        for fi32 := 0; fi32 < 16; fi32++ {
            fmt.Println(v1)
            fmt.Println(fi32, v2)
            v1 += v1
            _ = fi32
        }
    }
    fmt.Println(v1)
    v33 := GenS3{F0: false}
    if true {
        if ((v2 << ((2) & 63)) == v2) {
            v2 -= v2
        } else {
            _ = len(v5)
            _ = len(v3)
        }
    } else {
        if (false && (!true)) {
            v1 = 0.0
            var v34 int = 4294967295
            _ = v34
        }
        switch ((("foo bar" + "hello") + ("foo bar" + "\n\t")) + "") {
        case "a":
            switch ("\n\t" + "a") {
            case "\n\t":
                v2 = v2
            case "foo bar":
                if false {
                    v35 := []int{}
                    _ = v35
                }
            }
        case "foo bar":
            var v36 float32 = float32((1.5 / ((v1 - 3.14159) / v1)))
            _ = v36
        case "\n\t":
            var v37 float32 = (float32(v1) / 0.1)
            _ = v37
        default:
            v38 := GenS2{F0: ""}
            _ = v38
        }
        if false {
            var v39 uint32 = (^(v2))
            _ = v39
        }
    }
    var v40 uint16 = 514
    var v41 uint32 = (v2 / ((2) | 1))
    var v42 uint32 = (v41 | 56)
    for fi43 := 0; fi43 < 19; fi43++ {
        var v44 uint16 = (^(v40))
        var v45 int8 = 127
        _ = fi43
        _ = v44
        _ = v45
    }
    v46 := GenS4{F0: 2, F1: 593}
    v47 := v5[v2]
    if true {
        v2 = v2
        for fi48 := 0; fi48 < 8; fi48++ {
            var v49 int8 = -1
            if (!true) {
                v50 := []int{fi48, fi48, -385, (fi48 | (fi48 ^ fi48))}
                _ = v50
            }
            _ = fi48
            _ = v49
        }
    } else {
        if false {
            v51 := GenS2{F0: v47}
            _ = v51
        } else {
            v41 = (^(v42))
            v52 := GenS2{F0: "hello"}
            _ = v52
        }
    }
    v53 := []int{42, 1, 315}
    if (127 != -1) {
        v41 = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v53:", len(v53), v53)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v46
    _ = v47
    _ = v5
    _ = v53
}
