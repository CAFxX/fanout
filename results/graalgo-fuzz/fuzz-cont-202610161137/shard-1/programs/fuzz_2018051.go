package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 int; F2 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int64; F1 int64 }
func genFn0(p0 bool) int64 {
    if (p0 || p0) {
        v5 := GenS1{F0: 127}
        v6 := GenS2{F0: 0}
        _ = v5
        _ = v6
    }
    p0 = (127 >= -128)
    p0 = (!(true && (255 > 255)))
    if p0 {
        var v7 uint32 = 940
        if (!(709 != 524)) {
            v8 := GenS1{F0: 128}
            _ = v8
        }
        _ = v7
    } else {
        var v9 bool = (!true)
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS1{F0: 127}
            var v12 uint = 665
            var v13 uint64 = 1
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        v14 := []int{-174, -9223372036854775808}
        _ = v14
        _ = v9
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 int = 9223372036854775807
    var v3 float32 = (1.5 / (-2.25 - 1e10))
    var v4 float32 = v3
    switch (v2 / (((v2 & (v2 * (-(v2))))) | 1)) {
    case 42:
        if (0 > 2147483647) {
            v15 := GenS2{F0: 1}
            v2 -= v2
            if ((v2 % ((v2) | 1)) < (v2 % (((-(v2))) | 1))) {
                v4 -= ((float32(0.0) - (v4 / v3)) + 0.0)
                v16 := GenS3{F0: -286}
                _ = v16
            }
            _ = v15
        }
    case 128:
        v17 := GenS4{F0: -9223372036854775808, F1: -953}
        _ = v17
    case -1:
        if (false || (true || (int32(v2) > 1))) {
            fmt.Println(v3)
            switch (-(v2)) {
            case -602:
                var v18 int32 = 0
                _ = v18
            case 157:
                v19 := []int{-1}
                _ = v19
            }
        } else {
            v20 := GenS1{F0: 2}
            _ = v20
        }
    }
    v21 := GenS5{F0: -261}
    for fi22 := 0; fi22 < 8; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            if false {
                var v24 uint64 = 4294967295
                v24 += 42
                _ = v24
            } else {
                var v25 uint64 = 0
                v26 := []int{}
                switch "\n\t" {
                case "\n\t":
                    switch ("a" + "\n\t") {
                    case "":
                        fmt.Println(v2, v25, fi23)
                    case "hello":
                        var v27 int16 = -32768
                        _ = v27
                    }
                case "foo bar":
                    var v28 uint = 123
                    _ = v28
                }
                if false {
                    v29 := GenS2{F0: 0}
                    fmt.Println(fi22)
                    _ = v29
                }
                _ = v25
                _ = v26
            }
            _ = fi23
        }
        v2 += 777
        _ = fi22
    }
    var v30 float64 = float64(v4)
    _ = len(v1)
    v31 := GenS1{F0: 2}
    fmt.Println(v4, v3, v30)
    v32 := GenS6{F0: (true && true), F1: -9223372036854775808, F2: 0}
    switch -1 {
    case 128:
        if true {
            for fi33 := 0; fi33 < 2; fi33++ {
                var v34 int8 = 42
                v1["foo bar"] = -880
                if true {
                    fmt.Println(v30, v2, fi33)
                    for fi35 := 0; fi35 < 7; fi35++ {
                        v2 -= 32768
                        var v36 int32 = -1
                        v37 := []int{}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    var v38 uint16 = 32767
                    _ = v38
                }
                _ = fi33
                _ = v34
            }
            fmt.Println(v30, v3, v30)
            if ((((("\n\t" + "\n\t") + ("a" + (("\n\t" + (((("foo bar" + "hello") + ("hello" + "")) + "a") + "hello")) + "hello"))) + ("foo bar" + "hello")) + "") == ((("hello" + "hello") + (("\n\t" + "") + "")) + "\n\t")) {
                v2 -= v2
                fmt.Println(v30, v2)
                v39 := GenS4{F0: genFn0(true), F1: genFn0((556 < 0))}
                _ = v39
            } else {
                v2 = (v2 / ((v2) | 1))
                v4 = float32(1.5)
            }
        }
    case 32767:
        var v40 uint64 = 2
        _ = v40
    }
    var v41 uint16 = 954
    switch (v2 | (v2 & 0)) {
    case 1:
        for fi42 := 0; fi42 < 14; fi42++ {
            fmt.Println(v41, fi42, v3)
            var v43 bool = true
            _ = fi42
            _ = v43
        }
    case 711:
        var v44 float64 = (float64(1e10) - 0.1)
        _ = v44
    }
    fmt.Println(v41)
    var v45 uint = 4294967295
    if (v45 < (v45 / (((v45 * (v45 * 65535))) | 1))) {
        if (1 < -1) {
            var v46 uint8 = 124
            if false {
                switch v2 {
                case 273:
                    v47 := []int{}
                    _ = v47
                case 85:
                    var v48 int = (v2 % (((v2 & (v2 ^ v2))) | 1))
                    _ = v48
                }
                v30 = (v30 / (1.5 - -2.25))
            } else {
                if false {
                    var v49 uint32 = uint32(v46)
                    v41 = v41
                    v49 = v49
                    _ = v49
                } else {
                    v1["a"] = v2
                }
            }
            _ = v46
        } else {
            _ = len(v1)
        }
        if ((v41 + v41) < (v41 % ((v41) | 1))) {
            fmt.Println(v3, v3)
            v50 := []int{(-(v2)), 42}
            v51 := GenS7{F0: 533}
            _ = v50
            _ = v51
        } else {
            var v52 int = int(v41)
            v45 = (v45 ^ 1)
            v3 = v4
            _ = v52
        }
    } else {
        for fi53 := 0; fi53 < 12; fi53++ {
            fmt.Println(v3)
            v30 += (v30 / v30)
            _ = fi53
        }
        v54 := GenS8{F0: 888, F1: genFn0(true)}
        _ = v54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v41
    _ = v45
}
