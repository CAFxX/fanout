package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint16; F2 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 string }
func genFn0(p0 int) uint32 {
    if false {
        if (!(35 > 127)) {
            fmt.Println(p0)
            v6 := GenS1{F0: 308}
            p0 = p0
            _ = v6
        }
        p0 = (p0 | p0)
    }
    p0 -= p0
    fmt.Println(p0)
    if true {
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS2{F0: 919}
            fmt.Println(fi7, fi7)
            _ = fi7
            _ = v8
        }
        p0 += p0
    } else {
        v9 := []int{}
        _ = v9
    }
    _ = p0
    return 0
}
func genFn1(p0 bool) int {
    var v10 uint16 = 42
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 int32 = 0
        switch "a" {
        case "":
            if p0 {
                var v13 bool = (!("" < ("hello" + ("a" + "a"))))
                _ = v13
            } else {
                fmt.Println(v12)
                p0 = false
            }
        case "hello":
            switch ("hello" + "a") {
            case "hello":
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS2{F0: 1}
                    fmt.Println(fi14, v12)
                    _ = fi14
                    _ = v15
                }
            }
        default:
            if (p0 && p0) {
                var v16 string = ("a" + "")
                _ = v16
            } else {
                for fi17 := 0; fi17 < 18; fi17++ {
                    fmt.Println(fi11, p0)
                    fmt.Println(p0, v12)
                    _ = fi17
                }
                for fi18 := 0; fi18 < 10; fi18++ {
                    var v19 int = fi18
                    var v20 int = fi18
                    var v21 uint = 255
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                fmt.Println(p0, p0, fi11)
            }
        }
        _ = fi11
        _ = v12
    }
    if (!(54 >= 42)) {
        var v22 int32 = -1
        var v23 uint8 = 0
        _ = v22
        _ = v23
    } else {
        p0 = false
        v10 = (v10 & (v10 % (((v10 / (((v10 ^ v10)) | 1))) | 1)))
        for fi24 := 0; fi24 < 17; fi24++ {
            fmt.Println(p0, p0, p0)
            v25 := []int{}
            v26 := GenS2{F0: 690}
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    var v27 uint64 = 18446744073709551615
    if p0 {
        var v28 string = ""
        _ = v28
    } else {
        switch ((("" + "\n\t") + "\n\t") + "\n\t") {
        case "\n\t":
            v27 += (^(v27))
        }
        v29 := GenS3{F0: -73, F1: v10, F2: 546}
        _ = v29
    }
    _ = p0
    _ = v10
    _ = v27
    return 664
}

func main() {
    var v1 int32 = 2147483647
    var v2 float64 = 0.1
    v3 := make(map[string]int)
    var v4 uint64 = 175
    var v5 int8 = -1
    v30 := GenS4{F0: "hello"}
    v31 := []int{-419, 480}
    v3["a"] = 359
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 int16 = -1
        v1 = (v1 & v1)
        _ = fi32
        _ = v33
    }
    v4 += (v4 + v4)
    var v34 int8 = (v5 & 67)
    switch 127 {
    case -2147483648:
        v1 += (-(v1))
    case 9223372036854775807:
        var v35 uint32 = 42
        _ = v35
    }
    v34 = (-(v5))
    _ = len(v3)
    if ((!true) && ((^(v1)) < 42)) {
        v2 += v2
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        v37 := GenS3{F0: 32767, F1: 42, F2: -997}
        switch 32767 {
        case -403:
            var v38 bool = (100 < (v34 * v34))
            _ = v38
        case -9223372036854775808:
            v3["foo bar"] = fi36
        }
        var v39 int32 = v1
        _ = fi36
        _ = v37
        _ = v39
    }
    v40 := GenS5{F0: true, F1: ""}
    v41 := GenS6{F0: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
