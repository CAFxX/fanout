package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 string }
func genFn0(p0 int32, p1 uint8, p2 uint) uint32 {
    var v6 uint = p2
    p0 = (p0 / (((p0 | 1)) | 1))
    v7 := []int{1}
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 662
}
func genFn1() int64 {
    switch 128 {
    case 2147483647:
        v8 := []int{}
        _ = v8
    }
    var v9 float64 = (0.0 / -2.25)
    v10 := GenS1{F0: 2, F1: -513}
    var v11 float32 = (1.5 + 1e10)
    switch -706 {
    case -1:
        for fi12 := 0; fi12 < 9; fi12++ {
            if false {
                var v13 uint16 = 2
                _ = v13
            }
            v14 := []int{fi12, (^(fi12)), (fi12 & (fi12 & 9223372036854775807)), fi12}
            _ = fi12
            _ = v14
        }
    }
    _ = v10
    _ = v11
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{("hello" + ""), "\n\t", ("hello" + ""), ("hello" + "foo bar")}
    v3 := make(map[string]string)
    var v4 bool = true
    v5 := []string{}
    switch 1 {
    case 1:
        v4 = true
    case 4294967295:
        _ = len(v3)
    default:
        if (0 == 195) {
            if true {
                v4 = (2147483648 <= 2)
            } else {
                var v15 int = 42
                if (447 <= 4294967295) {
                    v4 = true
                    var v16 uint = 918
                    fmt.Println(v15)
                    _ = v16
                } else {
                    v15 -= v15
                    v15 = v15
                }
                if (113 > -1) {
                    var v17 int8 = 127
                    v4 = (!(!v4))
                    _ = v17
                }
                _ = v15
            }
            if v4 {
                v18 := GenS1{F0: 519, F1: 2147483647}
                _ = v18
            }
        }
    }
    v5 = append(v5, ("" + "foo bar"))
    v19 := GenS2{F0: 2, F1: 979}
    var v20 int16 = -32768
    _ = len(v1)
    if (("hello" == "") && v4) {
        var v21 int8 = -128
        var v22 int64 = -567
        var v23 int16 = v20
        _ = v21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 string = "hello"
        v4 = (!v4)
        var v26 uint = 943
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        if false {
            for fi28 := 0; fi28 < 6; fi28++ {
                fmt.Println(v4, v4)
                v4 = v4
                _ = fi28
            }
            var v29 uint16 = 774
            _ = v29
        }
        switch "a" {
        case "hello":
            delete(v1, (fi27 >> ((556) & 63)))
        }
        _ = fi27
    }
    var v30 bool = v4
    switch "foo bar" {
    case "a":
        v31 := GenS1{F0: 65535, F1: 1}
        _ = v31
    case "\n\t":
        for fi32 := 0; fi32 < 14; fi32++ {
            v30 = v30
            v30 = v30
            switch ("" + "foo bar") {
            case "foo bar":
                fmt.Println(v30, v4)
            }
            _ = fi32
        }
    case "":
        switch 32767 {
        case 2147483647:
            if v4 {
                for fi33 := 0; fi33 < 7; fi33++ {
                    v34 := GenS3{F0: 1}
                    for fi35 := 0; fi35 < 11; fi35++ {
                        var v36 int16 = (v20 % ((v20) | 1))
                        v37 := GenS1{F0: 42, F1: 42}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    _ = fi33
                    _ = v34
                }
            }
        case 127:
            v20 += 700
        case 942:
            var v38 string = ("" + "foo bar")
            _ = v38
        }
    }
    var v39 int32 = -1
    var v40 int8 = 1
    v41 := GenS4{F0: 929, F1: "\n\t"}
    var v42 bool = v30
    v42 = v4
    var v43 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
