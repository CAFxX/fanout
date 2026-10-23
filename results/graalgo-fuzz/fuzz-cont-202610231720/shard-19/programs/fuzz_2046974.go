package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int16; F2 int64 }
type GenS3 struct { F0 uint; F1 int; F2 int32 }
func genFn0() uint64 {
    switch ("" + (("foo bar" + ("hello" + "")) + "a")) {
    case "foo bar":
        if true {
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := GenS1{F0: 42}
                var v9 int = (^(fi7))
                v10 := []int{v9}
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        } else {
            if (!true) {
                v11 := GenS1{F0: 0}
                _ = v11
            }
            var v12 float32 = (1e10 + (3.14159 / -2.25))
            _ = v12
        }
    case "a":
        v13 := GenS1{F0: 4294967295}
        _ = v13
    case "":
        var v14 bool = false
        _ = v14
    default:
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 int16 = -594
            _ = fi15
            _ = v16
        }
    }
    var v17 int32 = 99
    fmt.Println(v17)
    v18 := []int{-726}
    _ = v17
    _ = v18
    return 2
}
func genFn1() int64 {
    for fi19 := 0; fi19 < 10; fi19++ {
        v20 := []int{(fi19 / (((-(fi19))) | 1)), fi19, (fi19 & fi19), (fi19 + fi19)}
        _ = fi19
        _ = v20
    }
    switch (("foo bar" + "\n\t") + "a") {
    case "foo bar":
        var v21 uint16 = 32767
        _ = v21
    }
    switch 127 {
    case 32768:
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 uint8 = 80
            var v24 bool = false
            _ = fi22
            _ = v23
            _ = v24
        }
    default:
        v25 := GenS2{F0: -1, F1: -371, F2: 1}
        _ = v25
    }
    if true {
        var v26 uint = 18446744073709551615
        v27 := []int{}
        fmt.Println(v26)
        _ = v26
        _ = v27
    }
    v28 := GenS1{F0: 42}
    _ = v28
    return -53
}

func main() {
    v1 := []string{(("" + "a") + "a"), ((("foo bar" + "\n\t") + "foo bar") + "\n\t")}
    var v2 uint64 = 18446744073709551615
    var v3 string = ((("foo bar" + "hello") + "a") + "a")
    var v4 string = "hello"
    v5 := []string{}
    var v6 uint64 = v2
    var v29 float32 = 3.14159
    var v30 bool = (!true)
    fmt.Println(v4, v30, v2)
    if len(v1) > 0 {
        v1[0] = "a"
    }
    var v31 int64 = 59
    v32 := GenS3{F0: 65535, F1: 32767, F2: -2147483648}
    for fi33 := 0; fi33 < 6; fi33++ {
        var v34 int8 = -116
        var v35 int = (fi33 | -776)
        var v36 int32 = 0
        _ = fi33
        _ = v34
        _ = v35
        _ = v36
    }
    if ((v6 | 8) <= (^(v2))) {
        v29 += 0.0
    } else {
        v6 += (v2 ^ (v6 / ((42) | 1)))
    }
    if len(v5) > 0 {
        v5[0] = ""
    }
    v37 := GenS3{F0: 4294967295, F1: 399, F2: -305}
    var v38 uint64 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
