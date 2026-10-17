package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    v5 := make(map[string]int)
    var v6 int32 = 228
    var v7 float32 = 1e10
    fmt.Println(v7)
    v6 = 808
    v8 := GenS1{F0: 128, F1: ("" + "hello")}
    switch -2147483648 {
    case -1:
        switch "hello" {
        case "foo bar":
            v7 -= v7
        }
    default:
        v6 += (v6 + (v6 >> ((42) & 63)))
    }
    var v9 int8 = -1
    v10 := GenS2{F0: -512}
    if (1e10 == 0.0) {
        v7 -= v7
    } else {
        for fi11 := 0; fi11 < 14; fi11++ {
            v7 = v7
            v7 += v7
            for fi12 := 0; fi12 < 8; fi12++ {
                v6 -= 0
                v6 += v6
                v13 := GenS3{F0: 0}
                if true {
                    var v14 uint32 = 2147483647
                    _ = v14
                }
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        v15 := GenS1{F0: 32767, F1: "foo bar"}
        _ = v15
    }
    _ = len(v5)
    var v16 int32 = v6
    for fi17 := 0; fi17 < 5; fi17++ {
        if false {
            var v18 uint = 42
            fmt.Println(v6, fi17)
            _ = v18
        } else {
            v9 += 5
            v9 = v9
            if ("\n\t" != "") {
                fmt.Println(v16)
                v19 := []int{fi17, 789}
                v9 -= 0
                _ = v19
            } else {
                v20 := []int{}
                for fi21 := 0; fi21 < 4; fi21++ {
                    v22 := GenS4{F0: 4294967295}
                    v7 = 0.1
                    if (2 >= 4294967295) {
                        var v23 int = (fi17 / ((-9223372036854775808) | 1))
                        _ = v23
                    } else {
                        var v24 int64 = -789
                        var v25 bool = (true || true)
                        v26 := GenS5{F0: 54}
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    _ = fi21
                    _ = v22
                }
                _ = len(v2)
                _ = v20
            }
        }
        v27 := v5[""]
        switch ("\n\t" + "a") {
        case "foo bar":
            fmt.Println(v7, v9, v9)
        case "hello":
            v28 := []int{-9223372036854775808, v27}
            _ = v28
        case "a":
            switch (v27 + -1) {
            case 32768:
                v6 = (^(v6))
            case -526:
                v16 += (^(v6))
            case -926:
                fmt.Println(v9, fi17)
            }
        }
        _ = fi17
        _ = v27
    }
    v29 := []int{}
    v7 = v7
    v30 := GenS2{F0: 830}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
