package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint64; F2 bool }
type GenS5 struct { F0 uint8; F1 int16; F2 int }
type GenS6 struct { F0 int; F1 string; F2 string }
func genFn0(p0 bool, p1 int8, p2 bool) int8 {
    for fi5 := 0; fi5 < 8; fi5++ {
        p1 = p1
        v6 := GenS1{F0: "foo bar"}
        _ = fi5
        _ = v6
    }
    p0 = ((42 == 2) || (127 > 255))
    var v7 int64 = 9223372036854775807
    switch "foo bar" {
    case "hello":
        for fi8 := 0; fi8 < 6; fi8++ {
            fmt.Println(p0)
            v9 := []int{}
            _ = fi8
            _ = v9
        }
    default:
        var v10 int8 = p1
        _ = v10
    }
    var v11 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return p1
}
func genFn1(p0 uint8, p1 uint16) bool {
    fmt.Println(p0, p1)
    switch "\n\t" {
    case "a":
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 int8 = 0
            _ = fi12
            _ = v13
        }
    case "":
        var v14 uint16 = 0
        _ = v14
    case "foo bar":
        switch (("" + ("\n\t" + ("" + ""))) + "a") {
        case "hello":
            var v15 float64 = float64(0.1)
            _ = v15
        case "a":
            v16 := GenS2{F0: (p0 >> ((4294967295) & 63)), F1: 166}
            _ = v16
        }
    default:
        var v17 int = 42
        _ = v17
    }
    _ = p0
    _ = p1
    return false
}

func main() {
    var v1 string = "\n\t"
    var v2 int32 = 0
    v3 := make(map[int]uint32)
    var v4 int64 = 9223372036854775807
    var v18 uint16 = 65535
    var v19 float64 = (0.0 / ((-2.25 + 1.5) + -2.25))
    v4 -= v4
    var v20 int8 = 34
    switch (v1 + v1) {
    case "a":
        switch (v1 + v1) {
        case "hello":
            v21 := GenS1{F0: "\n\t"}
            _ = v21
        case "\n\t":
            var v22 uint = 65535
            _ = v22
        case "a":
            if ((!false) || false) {
                for fi23 := 0; fi23 < 7; fi23++ {
                    if ((false || (!(true && (!true)))) && (float32(v19) != 1.5)) {
                        v4 += v4
                        switch fi23 {
                        case 414:
                            v24 := GenS3{F0: 18446744073709551615, F1: v4, F2: 255}
                            _ = v24
                        default:
                            var v25 int64 = 579
                            _ = v25
                        }
                        v26 := GenS4{F0: 65535, F1: 150, F2: true}
                        _ = v26
                    } else {
                        var v27 uint32 = uint32(v18)
                        _ = v27
                    }
                    _ = fi23
                }
                v1 = ((("" + "a") + "hello") + "a")
                for fi28 := 0; fi28 < 18; fi28++ {
                    var v29 int32 = -2147483648
                    _ = fi28
                    _ = v29
                }
            } else {
                v30 := []int{1, 558, 0}
                if (42 <= -474) {
                    fmt.Println(v1, v18)
                    var v31 string = "\n\t"
                    v4 = 1
                    _ = v31
                } else {
                    v2 -= 42
                }
                var v32 int = -2147483648
                _ = v30
                _ = v32
            }
        default:
            if true {
                v20 = int8(v2)
            } else {
                var v33 int = 2147483647
                var v34 uint = 0
                v18 += (v18 >> (((v34 >> (((v34 & v34)) & 63))) & 63))
                _ = v33
                _ = v34
            }
        }
    }
    if true {
        v35 := []int{}
        v36 := GenS1{F0: v1}
        v37 := GenS3{F0: 348, F1: -260, F2: 128}
        _ = v35
        _ = v36
        _ = v37
    }
    var v38 uint32 = 408
    fmt.Println(v2)
    if (!false) {
        var v39 int = -1
        _ = v39
    } else {
        var v40 uint32 = 2147483648
        var v41 uint16 = v18
        var v42 int = 1
        _ = v40
        _ = v41
        _ = v42
    }
    v43 := []int{4294967295}
    fmt.Println(v38, v19, v2)
    v44 := GenS5{F0: 42, F1: -942, F2: 42}
    var v45 uint64 = 18446744073709551615
    if (1e10 == 1.5) {
        v46 := GenS6{F0: -403, F1: "foo bar", F2: "a"}
        fmt.Println(v2, v20)
        var v47 bool = true
        _ = v46
        _ = v47
    }
    fmt.Println(v38, v45)
    var v48 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v48
}
