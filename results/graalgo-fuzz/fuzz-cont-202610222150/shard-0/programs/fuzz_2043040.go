package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 string; F1 bool }
func genFn0(p0 uint64, p1 int32) int32 {
    p1 = p1
    switch (("foo bar" + "\n\t") + ("hello" + "hello")) {
    case "hello":
        v4 := []int{272, -282}
        _ = v4
    }
    p1 -= p1
    p1 = p1
    var v5 uint16 = 42
    _ = p0
    _ = p1
    _ = v5
    return p1
}
func genFn1(p0 uint8, p1 string, p2 int32) bool {
    for fi6 := 0; fi6 < 1; fi6++ {
        switch (("foo bar" + (p1 + "foo bar")) + "\n\t") {
        case "hello":
            v7 := GenS1{F0: false}
            _ = v7
        case "":
            switch "hello" {
            case "a":
                if (-1 > 0) {
                    if true {
                        v8 := GenS1{F0: false}
                        switch p1 {
                        case "\n\t":
                            p0 = (p0 / ((p0) | 1))
                        }
                        var v9 int16 = 32767
                        _ = v8
                        _ = v9
                    } else {
                        var v10 int = fi6
                        var v11 int = (v10 ^ (v10 % ((v10) | 1)))
                        _ = v10
                        _ = v11
                    }
                    var v12 float64 = 0.1
                    switch p1 {
                    case "hello":
                        var v13 uint16 = 255
                        _ = v13
                    }
                    _ = v12
                }
            case "":
                v14 := []int{fi6, 362}
                _ = v14
            case "hello":
                v15 := GenS1{F0: true}
                _ = v15
            }
        case "a":
            p1 = "\n\t"
        }
        _ = fi6
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        fmt.Println(p1, fi16, p1)
        _ = fi16
    }
    p0 = p0
    p2 += (^(p2))
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    v1 := []string{"a", ("\n\t" + ""), ("hello" + ""), ("a" + ("a" + "foo bar"))}
    var v2 uint32 = 0
    var v3 int32 = 0
    var v17 bool = true
    var v18 uint64 = 42
    if false {
        v1 = append(v1, ("foo bar" + "a"))
        v19 := GenS1{F0: v17}
        var v20 uint = 2
        _ = v19
        _ = v20
    }
    var v21 int16 = 1
    for fi22 := 0; fi22 < 5; fi22++ {
        if true {
            v23 := GenS2{F0: -7}
            v1 = append(v1, "hello")
            switch "a" {
            case "a":
                for fi24 := 0; fi24 < 8; fi24++ {
                    switch "" {
                    case "\n\t":
                        if (210 >= 241) {
                            v17 = v17
                        }
                    case "foo bar":
                        v25 := GenS3{F0: int16(v18), F1: (^(v21))}
                        _ = v25
                    default:
                        v26 := GenS4{F0: 0}
                        _ = v26
                    }
                    v27 := GenS1{F0: v17}
                    _ = fi24
                    _ = v27
                }
            case "hello":
                var v28 bool = true
                _ = v28
            }
            _ = v23
        }
        var v29 int8 = 1
        _ = fi22
        _ = v29
    }
    var v30 uint = 65535
    fmt.Println(v3, v18, v2)
    var v31 int64 = 1
    if v17 {
        if true {
            for fi32 := 0; fi32 < 3; fi32++ {
                _ = len(v1)
                var v33 uint64 = (v18 % (((v18 & 42)) | 1))
                _ = fi32
                _ = v33
            }
            if len(v1) > 0 {
                v1[0] = "a"
            }
            v34 := GenS4{F0: 556}
            _ = v34
        } else {
            var v35 bool = v17
            v3 = (v3 / (((v3 | genFn0((v18 & (^(v18))), int32(v18)))) | 1))
            var v36 string = "foo bar"
            _ = v35
            _ = v36
        }
        var v37 uint16 = 65535
        v1 = append(v1, "")
        _ = v37
    } else {
        _ = len(v1)
    }
    var v38 string = "foo bar"
    v39 := GenS5{F0: (v38 + v38), F1: false}
    if v17 {
        v31 -= 1
    } else {
        v40 := GenS6{F0: "hello", F1: true}
        var v41 uint = v30
        _ = v40
        _ = v41
    }
    if (!(1.5 != 1e10)) {
        var v42 int8 = 100
        fmt.Println(v17, v2, v31)
        _ = v42
    }
    v1 = append(v1, "a")
    var v43 int = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v39
    _ = v43
}
