package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int8 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 string; F1 int32; F2 uint }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 bool) int64 {
    if (-1 == 0) {
        switch "a" {
        case "":
            switch -2147483648 {
            case 144:
                for fi5 := 0; fi5 < 10; fi5++ {
                    switch "\n\t" {
                    case "foo bar":
                        p0 = p0
                    }
                    for fi6 := 0; fi6 < 8; fi6++ {
                        switch (^(fi5)) {
                        case -1:
                            var v7 bool = (474 != 10)
                            _ = v7
                        case 0:
                            var v8 float32 = 0.1
                            _ = v8
                        case 4294967295:
                            v9 := []int{fi6, (fi5 ^ 127), (fi6 << ((1) & 63)), (fi6 ^ 32768)}
                            _ = v9
                        default:
                            var v10 float32 = (0.1 - (((1e10 / (1.5 * 0.1)) + (0.0 + 3.14159)) * ((3.14159 * 1e10) - -2.25)))
                            _ = v10
                        }
                        switch "a" {
                        case "":
                            v11 := GenS1{F0: -1, F1: (p0 || (2147483647 < 165)), F2: 100}
                            _ = v11
                        default:
                            var v12 float64 = 0.1
                            _ = v12
                        }
                        switch fi6 {
                        case 2147483647:
                            if (p0 && p0) {
                                p0 = p0
                                var v13 float32 = 3.14159
                                _ = v13
                            }
                        }
                        var v14 int8 = -128
                        var v15 int8 = 42
                        v16 := []int{-528, fi6, (fi6 / ((fi6) | 1)), 32768}
                        var v17 float32 = 3.14159
                        if true {
                            var v18 int8 = v15
                            var v19 float32 = (v17 - v17)
                            _ = v18
                            _ = v19
                        }
                        _ = fi6
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    var v20 uint64 = 511
                    _ = fi5
                    _ = v20
                }
            case 42:
                p0 = p0
            case 128:
                v21 := GenS1{F0: 11, F1: (302 == 32767), F2: -27}
                _ = v21
            default:
                p0 = (-2147483648 <= -1)
            }
        case "\n\t":
            v22 := GenS2{F0: -1, F1: -128}
            _ = v22
        default:
            if (-655 >= 42) {
                p0 = (1 >= -9223372036854775808)
                switch ("\n\t" + "") {
                case "foo bar":
                    var v23 int = 2147483647
                    _ = v23
                }
                p0 = (0 != 0)
            }
        }
    } else {
        var v24 bool = (0 < 1)
        p0 = p0
        _ = v24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        if p0 {
            for fi26 := 0; fi26 < 7; fi26++ {
                fmt.Println(fi25, fi26, fi25)
                _ = fi26
            }
        } else {
            var v27 uint8 = 2
            _ = v27
        }
        var v28 uint = 1
        _ = fi25
        _ = v28
    }
    switch 128 {
    case -1:
        var v29 uint16 = 818
        _ = v29
    }
    var v30 uint32 = 65535
    v31 := []int{4294967295, 9223372036854775807, -579, 285}
    _ = p0
    _ = v30
    _ = v31
    return -9223372036854775808
}
func genFn1(p0 string, p1 uint64) bool {
    v32 := GenS3{F0: p0, F1: 524, F2: 0}
    fmt.Println(p1, p1)
    v33 := GenS2{F0: -1, F1: -53}
    v34 := GenS4{F0: p0, F1: 42}
    _ = p0
    _ = p1
    _ = v32
    _ = v33
    _ = v34
    return true
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 1
    var v3 int8 = -9
    var v4 uint32 = 127
    var v35 string = ("\n\t" + ("a" + ""))
    switch ((v35 + (v35 + (((v35 + "a") + v35) + v35))) + v35) {
    case "\n\t":
        v3 += (v3 - 127)
    case "a":
        v2 -= (v2 & genFn0(false))
    case "hello":
        var v36 float64 = (3.14159 * (-2.25 * ((0.0 / float64((1e10 - ((0.0 - 1e10) * 0.1)))) + float64(float32(3.14159)))))
        _ = v36
    default:
        switch (v35 + "foo bar") {
        case "foo bar":
            var v37 string = (v35 + v35)
            _ = v37
        case "":
            if true {
                if (!true) {
                    v3 -= v3
                    for fi38 := 0; fi38 < 9; fi38++ {
                        var v39 float64 = ((1.5 * (0.0 - 1e10)) / 3.14159)
                        var v40 int32 = -1
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    _ = len(v1)
                }
            } else {
                v2 -= (v2 / ((-9223372036854775808) | 1))
                var v41 uint = 255
                var v42 float64 = 3.14159
                _ = v41
                _ = v42
            }
        }
    }
    v43 := GenS4{F0: "\n\t", F1: 623}
    var v44 string = v35
    var v45 uint64 = 1
    v1[v35] = false
    if true {
        v44 = v35
        v45 += 1
        var v46 string = ((v44 + v35) + v44)
        _ = v46
    }
    fmt.Println(v44, v45, v4)
    for fi47 := 0; fi47 < 13; fi47++ {
        var v48 int16 = 1000
        var v49 string = ((((v44 + "") + ((v44 + ("\n\t" + ((v35 + ((v35 + v35) + ("a" + ("\n\t" + v35)))) + ("foo bar" + v35)))) + v35)) + "foo bar") + ("foo bar" + v35))
        fmt.Println(v4)
        _ = fi47
        _ = v48
        _ = v49
    }
    for fi50 := 0; fi50 < 7; fi50++ {
        if false {
            fmt.Println(v2)
        } else {
            v51 := []int{-758, fi50, -492}
            for fi52 := 0; fi52 < 10; fi52++ {
                var v53 uint16 = 42
                _ = fi52
                _ = v53
            }
            fmt.Println(v45)
            _ = v51
        }
        v54 := []int{int(v4), fi50}
        _ = fi50
        _ = v54
    }
    v45 = (v45 / (((^(v45))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
