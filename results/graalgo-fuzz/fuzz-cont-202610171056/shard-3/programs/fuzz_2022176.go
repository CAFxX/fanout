package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int8; F2 int32 }
type GenS3 struct { F0 int; F1 int8; F2 uint }
func genFn0(p0 bool) int32 {
    for fi4 := 0; fi4 < 10; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            p0 = (762 >= 1)
            if (!((false || false) || (847 < -21))) {
                if p0 {
                    p0 = p0
                    p0 = (945 != 1)
                    if p0 {
                        var v6 uint32 = 2147483647
                        v7 := GenS1{F0: 42}
                        var v8 string = "foo bar"
                        if p0 {
                            fmt.Println(fi5)
                            p0 = ((fi5 << ((701) & 63)) >= fi4)
                        }
                        _ = v6
                        _ = v7
                        _ = v8
                    } else {
                        v9 := []int{(fi4 / (((fi4 | (fi5 ^ fi5))) | 1)), fi4, fi4, (fi5 / (((fi4 & 32767)) | 1))}
                        if len(v9) > 0 {
                            v9[0] = -263
                        }
                        _ = v9
                    }
                    var v10 uint32 = 2147483648
                    _ = v10
                }
                fmt.Println(p0, p0)
            }
            var v11 uint = 594
            _ = fi5
            _ = v11
        }
        v12 := []int{}
        _ = fi4
        _ = v12
    }
    p0 = p0
    p0 = p0
    for fi13 := 0; fi13 < 13; fi13++ {
        switch "hello" {
        case "hello":
            v14 := []int{(fi13 + fi13), (fi13 * (-(fi13))), (fi13 ^ fi13)}
            _ = v14
        case "foo bar":
            p0 = true
        }
        var v15 int16 = -954
        if true {
            var v16 int64 = 42
            p0 = true
            _ = v16
        } else {
            var v17 bool = false
            for fi18 := 0; fi18 < 15; fi18++ {
                v19 := GenS1{F0: 0}
                p0 = v17
                v20 := []int{(fi13 * (fi18 & 42)), fi13, (fi18 % (((^(fi18))) | 1))}
                _ = fi18
                _ = v19
                _ = v20
            }
            if v17 {
                switch "foo bar" {
                case "foo bar":
                    v21 := []int{fi13, fi13}
                    _ = v21
                }
                for fi22 := 0; fi22 < 14; fi22++ {
                    switch "foo bar" {
                    case "":
                        fmt.Println(fi22, fi13, fi13)
                    case "foo bar":
                        p0 = p0
                    default:
                        v23 := []int{}
                        _ = v23
                    }
                    v24 := GenS2{F0: 379, F1: 100, F2: 0}
                    _ = fi22
                    _ = v24
                }
            } else {
                var v25 int8 = 0
                _ = v25
            }
            _ = v17
        }
        _ = fi13
        _ = v15
    }
    _ = p0
    return 2147483647
}

func main() {
    var v1 float32 = float32(0.1)
    v2 := []string{"hello", "foo bar"}
    v3 := make(map[int]uint32)
    v1 -= (v1 / 0.1)
    var v26 float64 = (3.14159 - 0.0)
    v27 := GenS3{F0: 32768, F1: 111, F2: 65535}
    for fi28 := 0; fi28 < 9; fi28++ {
        switch (("a" + "\n\t") + ("foo bar" + "\n\t")) {
        case "\n\t":
            var v29 float32 = (-2.25 + 0.1)
            _ = v29
        }
        switch "foo bar" {
        case "a":
            if (2 == 2) {
                if (271 == 4294967295) {
                    v26 = (((0.0 / v26) - v26) / (v26 + -2.25))
                    for fi30 := 0; fi30 < 20; fi30++ {
                        var v31 string = "a"
                        delete(v3, fi28)
                        fmt.Println(v26)
                        _ = fi30
                        _ = v31
                    }
                    fmt.Println(fi28, v1, v1)
                }
                for fi32 := 0; fi32 < 1; fi32++ {
                    var v33 uint = 276
                    delete(v3, (fi32 + (-(fi28))))
                    _ = fi32
                    _ = v33
                }
                v26 = v26
            }
        case "":
            var v34 uint = 42
            _ = v34
        }
        _ = fi28
    }
    var v35 uint8 = 2
    for fi36 := 0; fi36 < 19; fi36++ {
        for fi37 := 0; fi37 < 15; fi37++ {
            delete(v3, fi36)
            fmt.Println(v1, fi36)
            _ = fi37
        }
        if false {
            var v38 uint16 = 2
            var v39 uint8 = v35
            v35 -= (v35 / ((1) | 1))
            _ = v38
            _ = v39
        }
        for fi40 := 0; fi40 < 3; fi40++ {
            delete(v3, (fi36 + (fi36 << ((0) & 63))))
            _ = fi40
        }
        _ = fi36
    }
    var v41 uint8 = 2
    switch -9223372036854775808 {
    case 4294967295:
        switch -2147483648 {
        case 1:
            v41 += v35
        case 32767:
            v1 -= 0.0
        default:
            for fi42 := 0; fi42 < 13; fi42++ {
                var v43 int32 = 0
                v41 = 193
                _ = fi42
                _ = v43
            }
        }
    case 32767:
        var v44 bool = (0 < 2)
        _ = v44
    default:
        v35 -= (v35 >> ((65535) & 63))
    }
    var v45 float32 = v1
    var v46 uint32 = 388
    var v47 int = 127
    v3[v47] = v46
    v46 = v46
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v41
    _ = v45
    _ = v46
    _ = v47
}
