package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    v4 := []string{("" + ("" + ("\n\t" + "foo bar"))), ("foo bar" + ("\n\t" + ("" + "foo bar"))), "foo bar"}
    v5 := make(map[string]bool)
    for fi6 := 0; fi6 < 18; fi6++ {
        if (true && false) {
            var v7 int = fi6
            if (("\n\t" == ("" + (("hello" + (("\n\t" + ("foo bar" + "hello")) + "\n\t")) + ("foo bar" + (("\n\t" + "a") + "a"))))) && false) {
                if true {
                    var v8 int64 = -401
                    _ = v8
                }
                v9 := GenS1{F0: 0, F1: 1}
                var v10 int = v7
                _ = v10
                _ = v9
            } else {
                var v11 bool = (42 != 42)
                _ = v11
            }
            _ = v7
        }
        _ = fi6
    }
    var v12 int32 = -841
    v13 := GenS1{F0: 42, F1: 42}
    switch "\n\t" {
    case "a":
        v12 = v12
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        v12 = (v12 >> ((0) & 63))
        var v15 float64 = ((float64(3.14159) / float64(3.14159)) - 3.14159)
        if false {
            var v16 int = fi14
            v15 += (float64(-2.25) * v15)
            v15 = v15
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        if (!("" >= "foo bar")) {
            fmt.Println(fi17, fi17)
        } else {
            v12 -= v12
        }
        if false {
            var v18 uint32 = 2147483647
            _ = v18
        }
        var v19 float64 = 0.1
        _ = fi17
        _ = v19
    }
    if len(v4) > 0 {
        v4[0] = ""
    }
    if true {
        for fi20 := 0; fi20 < 19; fi20++ {
            v12 -= v12
            var v21 uint8 = 2
            v22 := GenS1{F0: -611, F1: 4294967295}
            _ = fi20
            _ = v21
            _ = v22
        }
        delete(v2, "")
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 bool = (0.0 == 1.5)
            v24 = (float32(float64((0.0 + float32(((-2.25 + -2.25) / float64(float32(-2.25))))))) <= 3.14159)
            var v25 uint8 = 2
            _ = fi23
            _ = v24
            _ = v25
        }
    } else {
        v26 := GenS1{F0: -1, F1: 65535}
        fmt.Println(v12)
        _ = v26
    }
    var v27 uint16 = 1
    for fi28 := 0; fi28 < 10; fi28++ {
        v27 = v27
        var v29 float32 = (1.5 * 1e10)
        v30 := GenS1{F0: -484, F1: 4294967295}
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println(v27)
    for fi31 := 0; fi31 < 20; fi31++ {
        v27 += v27
        _ = fi31
    }
    v12 = v12
    v4 = append(v4, "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
