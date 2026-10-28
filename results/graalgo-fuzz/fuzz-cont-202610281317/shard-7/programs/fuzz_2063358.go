package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 uint16; F2 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 int64; F2 int8 }
func genFn0(p0 int32, p1 uint32) int64 {
    switch "a" {
    case "foo bar":
        v6 := GenS1{F0: 4294967295, F1: ("" + "foo bar")}
        _ = v6
    }
    p1 = (^(p1))
    if true {
        switch "" {
        case "foo bar":
            p0 = 2147483647
        case "":
            v7 := GenS1{F0: 664, F1: ("a" + "hello")}
            _ = v7
        default:
            fmt.Println(p0)
        }
        var v8 int16 = -32768
        var v9 int64 = 0
        _ = v8
        _ = v9
    } else {
        if true {
            var v10 string = "hello"
            switch -1 {
            case 1:
                v11 := []int{127, -3, 809}
                _ = v11
            }
            _ = v10
        } else {
            v12 := GenS2{F0: 625}
            switch "foo bar" {
            case "":
                v13 := []int{-953, 271, 127}
                _ = v13
            case "foo bar":
                switch 32768 {
                case 0:
                    p0 += (^(p0))
                }
            case "hello":
                var v14 uint64 = 762
                _ = v14
            default:
                p1 += (p1 ^ (p1 * (p1 + p1)))
            }
            _ = v12
        }
        v15 := []int{0, -2147483648, 108}
        fmt.Println(p0, p0)
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := []int{}
        p1 = 0
        switch "\n\t" {
        case "hello":
            p1 += (^(p1))
        }
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 255}
    _ = p0
    _ = p1
    _ = v18
    return -1
}
func genFn1(p0 uint64) int16 {
    fmt.Println(p0)
    var v19 int16 = 0
    switch "\n\t" {
    case "a":
        p0 = (p0 | p0)
    case "\n\t":
        v19 += (v19 + (v19 & v19))
    }
    _ = p0
    _ = v19
    return (v19 * -32768)
}

func main() {
    var v1 bool = (!(true || true))
    var v2 uint8 = 117
    var v3 string = "\n\t"
    var v4 uint32 = 0
    v5 := make(map[int]string)
    var v20 uint32 = v4
    var v21 int16 = genFn1(18446744073709551615)
    fmt.Println(v20, v20, v4)
    delete(v5, 128)
    var v22 uint64 = 2
    var v23 uint8 = v2
    v24 := v5[762]
    switch (("a" + ((v3 + (v3 + "a")) + v24)) + (v24 + v24)) {
    case "hello":
        var v25 uint32 = 1
        _ = v25
    case "foo bar":
        if v1 {
            var v26 uint8 = v2
            _ = v26
        }
    default:
        if (v1 || v1) {
            for fi27 := 0; fi27 < 5; fi27++ {
                v2 += (^(v2))
                _ = fi27
            }
        }
    }
    v3 = "a"
    if (!(!v1)) {
        switch (v3 + ("" + v24)) {
        case "":
            for fi28 := 0; fi28 < 16; fi28++ {
                for fi29 := 0; fi29 < 17; fi29++ {
                    fmt.Println(v21, v1)
                    if (false || v1) {
                        v30 := GenS3{F0: (v1 || (v24 <= "hello"))}
                        _ = v30
                    }
                    _ = fi29
                }
                _ = fi28
            }
        case "foo bar":
            v31 := GenS4{F0: 65535}
            _ = v31
        case "hello":
            v32 := []int{-238, 128, 793, 9223372036854775807}
            _ = v32
        }
        var v33 uint32 = 797
        if v1 {
            v2 += v2
            if ((v22 * 537) == 42) {
                v33 = v4
                var v34 int8 = -13
                _ = v34
            } else {
                v35 := GenS3{F0: v1}
                _ = v35
            }
            switch 127 {
            case 9223372036854775807:
                var v36 int = 1
                _ = v36
            case 0:
                var v37 uint32 = (^(v20))
                _ = v37
            case 127:
                v4 = v33
            }
        } else {
            fmt.Println(v1)
        }
        _ = v33
    }
    for fi38 := 0; fi38 < 5; fi38++ {
        if false {
            switch ("" + v3) {
            case "":
                switch (v3 + v3) {
                case "hello":
                    v21 = (v21 - genFn1((^(v22))))
                case "foo bar":
                    _ = len(v5)
                }
            case "a":
                v39 := GenS5{F0: v3, F1: 867, F2: ""}
                _ = v39
            }
            v40 := GenS6{F0: "\n\t"}
            _ = v40
        } else {
            v5[fi38] = v3
        }
        switch 4294967295 {
        case 4294967295:
            if (!(!(!(778 != 2)))) {
                if (v21 >= (v21 / ((v21) | 1))) {
                    var v41 int64 = 1
                    _ = v41
                } else {
                    var v42 float32 = -2.25
                    var v43 float64 = 1e10
                    _ = v42
                    _ = v43
                }
                switch "hello" {
                case "foo bar":
                    v44 := []int{(^(fi38))}
                    _ = v44
                case "hello":
                    v45 := GenS4{F0: 2}
                    _ = v45
                default:
                    v46 := GenS7{F0: -657, F1: -439, F2: 17}
                    _ = v46
                }
            } else {
                var v47 uint16 = 65535
                var v48 int16 = genFn1(v22)
                for fi49 := 0; fi49 < 9; fi49++ {
                    var v50 int = (^(fi49))
                    _ = fi49
                    _ = v50
                }
                _ = v47
                _ = v48
            }
        }
        _ = fi38
    }
    if (!v1) {
        fmt.Println(v21, v23)
    }
    var v51 float32 = 1e10
    v52 := []int{-37, 42}
    fmt.Println(v22, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v52:", len(v52), v52)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
}
