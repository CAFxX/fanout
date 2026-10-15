package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint64) int64 {
    var v6 uint16 = 2
    v6 -= 65535
    _ = p0
    _ = v6
    return -9223372036854775808
}
func genFn1() int {
    if false {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 uint16 = 0
            switch ("hello" + "foo bar") {
            case "a":
                var v9 int = (-(fi7))
                _ = v9
            case "foo bar":
                var v10 uint16 = (v8 - (v8 ^ (^(v8))))
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        switch -303 {
        case -9223372036854775808:
            if false {
                var v11 uint8 = 94
                _ = v11
            } else {
                v12 := GenS1{F0: 0, F1: 0}
                _ = v12
            }
        case -155:
            var v13 string = "hello"
            _ = v13
        }
    } else {
        v14 := []int{32768, -842}
        _ = v14
    }
    var v15 float32 = 0.0
    for fi16 := 0; fi16 < 20; fi16++ {
        switch fi16 {
        case 4294967295:
            v17 := []int{(fi16 | (fi16 & fi16))}
            _ = v17
        default:
            var v18 uint32 = 2147483648
            _ = v18
        }
        _ = fi16
    }
    _ = v15
    return -504
}

func main() {
    var v1 string = (("" + ((("\n\t" + ("foo bar" + (("\n\t" + "\n\t") + "\n\t"))) + ("hello" + "foo bar")) + "\n\t")) + "hello")
    var v2 uint = 1
    var v3 int16 = 1
    v4 := make(map[int]int)
    v5 := []int{2147483647, 32768, -301}
    var v19 int32 = 42
    fmt.Println(v19)
    var v20 bool = true
    for fi21 := 0; fi21 < 10; fi21++ {
        v3 = (v3 % (((v3 << (((v2 ^ v2)) & 63))) | 1))
        fmt.Println(v3)
        _ = fi21
    }
    v22 := v4[genFn1()]
    if false {
        v23 := GenS1{F0: -987, F1: 42}
        var v24 float32 = float32((0.0 * float64(0.0)))
        v25 := GenS1{F0: 103, F1: 127}
        _ = v23
        _ = v24
        _ = v25
    } else {
        v26 := GenS2{F0: -1}
        _ = v26
    }
    v5 = append(v5, 1)
    for fi27 := 0; fi27 < 4; fi27++ {
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 int8 = 0
            fmt.Println(v22, v2)
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    var v30 uint16 = 967
    v31 := GenS1{F0: (v19 * 0), F1: 1}
    fmt.Println(v19, v30)
    v32 := GenS3{F0: (v22 * (v22 * v22))}
    v30 = (^(v30))
    fmt.Println(v20)
    v4[genFn1()] = v22
    if true {
        switch v22 {
        case 645:
            v5 = append(v5, v22)
        case 4294967295:
            var v33 uint32 = 445
            _ = v33
        }
    } else {
        switch (-(v22)) {
        case 4294967295:
            v34 := v4[genFn1()]
            _ = v34
        case -624:
            v1 = (("foo bar" + v1) + "\n\t")
        }
        var v35 uint8 = 97
        _ = v35
    }
    v36 := GenS3{F0: v22}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
    _ = v4
    _ = v5
}
