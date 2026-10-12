package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0(p0 bool, p1 bool) int {
    v4 := GenS1{F0: -528}
    switch -2147483648 {
    case 1:
        var v5 int64 = 1
        _ = v5
    case -958:
        p0 = (false && (!((p1 && ((p1 && (!(p0 && false))) || true)) || p0)))
    case 0:
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := []int{128, fi6, 396, fi6}
            _ = fi6
            _ = v7
        }
    }
    var v8 uint64 = 2
    if (0 != 894) {
        if (1 > -743) {
            var v9 uint64 = (v8 & v8)
            p1 = (!p0)
            _ = v9
        }
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS1{F0: 1}
            _ = fi10
            _ = v11
        }
        v12 := []int{42}
        _ = v12
    } else {
        if ((v8 & 2) > v8) {
            v13 := []int{-2147483648}
            var v14 int = 128
            v8 -= v8
            _ = v13
            _ = v14
        } else {
            fmt.Println(p0, v8)
            v15 := GenS1{F0: 1}
            _ = v15
        }
    }
    var v16 int16 = 0
    _ = p0
    _ = p1
    _ = v16
    _ = v4
    _ = v8
    return 0
}
func genFn1(p0 bool, p1 uint) string {
    p1 = p1
    p1 += p1
    var v17 int16 = 496
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := []int{(^(fi18)), fi18, fi18}
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        _ = fi18
        _ = v19
    }
    var v20 int32 = -539
    _ = p0
    _ = p1
    _ = v17
    _ = v20
    return "hello"
}

func main() {
    var v1 uint8 = 246
    var v2 float32 = float32(1.5)
    var v3 float32 = 1e10
    var v21 int16 = 32767
    fmt.Println(v2, v2)
    var v22 int8 = 16
    v23 := []int{}
    if len(v23) > 0 {
        v23[0] = 42
    }
    v21 = v21
    v24 := GenS2{F0: 285}
    var v25 int64 = 43
    for fi26 := 0; fi26 < 2; fi26++ {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        var v27 string = ("hello" + ("a" + ("\n\t" + "a")))
        switch (fi26 / ((genFn0(false, true)) | 1)) {
        case 32768:
            for fi28 := 0; fi28 < 4; fi28++ {
                switch (fi26 / ((int(v1)) | 1)) {
                case 32767:
                    if (9223372036854775807 <= (^(v25))) {
                        v22 += v22
                        fmt.Println(v25)
                        var v29 uint32 = 1
                        v27 = v27
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                        _ = v29
                    } else {
                        var v30 float32 = v2
                        _ = v30
                    }
                case -2147483648:
                    switch genFn0((-225 >= 0), (!true)) {
                    case -2147483648:
                        fmt.Println(v21, v25, v25)
                    case 4294967295:
                        var v31 float64 = 0.0
                        _ = v31
                    }
                case -709:
                    v23 = append(v23, fi26)
                default:
                    var v32 uint8 = (^(v1))
                    _ = v32
                }
                switch (fi28 ^ fi28) {
                case 32767:
                    fmt.Println(v3, v27)
                case 9223372036854775807:
                    fmt.Println(v22, v22)
                case 32768:
                    var v33 uint = 65535
                    _ = v33
                }
                _ = fi28
            }
        case 81:
            if (!(-925 <= -1)) {
                if (v1 >= v1) {
                    v34 := GenS2{F0: 2147483647}
                    _ = v34
                }
                _ = len(v23)
                if false {
                    v35 := GenS2{F0: 42}
                    v23 = append(v23, 950)
                    v25 -= (v25 - 42)
                    _ = v35
                } else {
                    v27 = "hello"
                }
            } else {
                for fi36 := 0; fi36 < 2; fi36++ {
                    var v37 uint32 = 420
                    v25 = v25
                    switch v27 {
                    case "foo bar":
                        v25 = (^(v25))
                    }
                    _ = fi36
                    _ = v37
                }
                if false {
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                    v38 := GenS1{F0: v21}
                    _ = v38
                } else {
                    switch ((v27 + "") + (v27 + "\n\t")) {
                    case "":
                        var v39 int = 32767
                        _ = v39
                    case "hello":
                        v40 := GenS3{F0: 0, F1: 2147483647}
                        _ = v40
                    case "foo bar":
                        fmt.Println(v25, fi26)
                    }
                }
            }
        default:
            v41 := GenS3{F0: v1, F1: 4294967295}
            _ = v41
        }
        _ = fi26
        _ = v27
    }
    if (false || true) {
        var v42 int8 = -128
        v1 = (v1 - (v1 / ((v1) | 1)))
        _ = v42
    }
    v22 += v22
    if (false && true) {
        v1 = v1
    }
    v3 = 1.5
    var v43 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v43
}
