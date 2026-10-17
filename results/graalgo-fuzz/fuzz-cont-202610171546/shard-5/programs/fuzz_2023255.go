package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0(p0 int8) int {
    fmt.Println(p0)
    v7 := GenS1{F0: (("" + "foo bar") + ("hello" + "a"))}
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            switch (fi9 | fi9) {
            case -1:
                p0 += 100
            }
            p0 += 20
            _ = fi9
        }
        v10 := []int{fi8, (fi8 << ((255) & 63)), (fi8 / ((fi8) | 1))}
        _ = fi8
        _ = v10
    }
    if false {
        fmt.Println(p0)
        var v11 uint64 = 90
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 uint32 = 580
            v14 := []int{}
            var v15 string = "\n\t"
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = v11
    } else {
        p0 -= p0
        v16 := []int{346}
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 bool = false
            var v19 uint64 = 4294967295
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v16
    }
    _ = p0
    _ = v7
    return 32768
}
func genFn1(p0 int64, p1 uint32) string {
    v20 := []int{238, 128, -9223372036854775808, 2147483647}
    if (p1 == (^(p1))) {
        var v21 int = 1
        var v22 uint16 = 255
        var v23 uint8 = 188
        _ = v21
        _ = v22
        _ = v23
    } else {
        var v24 int8 = 42
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        v25 := GenS1{F0: "\n\t"}
        _ = v24
        _ = v25
    }
    v26 := GenS1{F0: ("" + "\n\t")}
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    _ = p0
    _ = p1
    _ = v20
    _ = v26
    return "hello"
}

func main() {
    var v1 string = ""
    var v2 string = (v1 + ((v1 + v1) + "a"))
    var v3 string = "a"
    var v4 bool = (true || true)
    var v5 uint16 = 32767
    var v6 float32 = (1e10 - -2.25)
    v1 = v3
    v4 = true
    var v27 int64 = 0
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 int16 = 0
        fmt.Println(v27)
        var v30 float64 = (0.1 - 3.14159)
        _ = fi28
        _ = v29
        _ = v30
    }
    switch v3 {
    case "":
        v5 = (v5 & (^(v5)))
    default:
        var v31 int16 = 42
        _ = v31
    }
    v2 = ((v1 + v3) + (v3 + ("" + (((v1 + "") + "hello") + "foo bar"))))
    for fi32 := 0; fi32 < 14; fi32++ {
        if (!v4) {
            if v4 {
                v5 -= (v5 / ((1) | 1))
                switch fi32 {
                case -126:
                    var v33 float32 = float32(1e10)
                    _ = v33
                }
                v34 := []int{}
                _ = v34
            }
            v35 := []int{}
            var v36 int32 = 0
            _ = v35
            _ = v36
        } else {
            v37 := []int{(fi32 | genFn0(127)), (fi32 & (fi32 ^ 229))}
            _ = v37
        }
        if v4 {
            v4 = v4
            for fi38 := 0; fi38 < 16; fi38++ {
                v1 = v2
                v27 += 9223372036854775807
                _ = fi38
            }
        }
        if (-1 > 42) {
            for fi39 := 0; fi39 < 10; fi39++ {
                v40 := []int{fi39, fi32, fi32, 42}
                var v41 string = "hello"
                _ = fi39
                _ = v40
                _ = v41
            }
        }
        _ = fi32
    }
    var v42 int32 = 377
    v43 := GenS2{F0: "", F1: (v27 / (((v27 + v27)) | 1))}
    if v4 {
        var v44 int16 = 135
        _ = v44
    }
    if (genFn0(100) != 280) {
        var v45 string = v1
        v46 := GenS3{F0: 694, F1: v4}
        fmt.Println(v2)
        _ = v45
        _ = v46
    } else {
        v47 := GenS2{F0: (v3 + v1), F1: -770}
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
