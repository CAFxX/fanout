package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int; F1 uint64; F2 string }
func genFn0(p0 uint32, p1 uint16) uint8 {
    var v7 string = "\n\t"
    v7 = "a"
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 2147483647
    var v2 float32 = 1.5
    var v3 uint16 = 1
    var v4 int64 = -1
    var v5 uint32 = 2147483647
    var v6 string = "hello"
    v6 = ("hello" + "hello")
    v8 := GenS1{F0: 65535, F1: 993, F2: -128}
    if (100 >= 100) {
        var v9 string = (v6 + "\n\t")
        v10 := []int{}
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := GenS2{F0: false}
            v13 := GenS3{F0: 1, F1: 2147483647}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
        _ = v9
    }
    var v14 bool = false
    for fi15 := 0; fi15 < 17; fi15++ {
        if false {
            switch v6 {
            case "foo bar":
                var v16 uint64 = 1
                _ = v16
            case "hello":
                switch fi15 {
                case 4294967295:
                    v3 -= v3
                case 127:
                    fmt.Println(v3, v2, v5)
                }
            case "\n\t":
                v6 = (v6 + v6)
            }
            fmt.Println(v5, v1, v5)
        }
        var v17 float32 = v2
        _ = fi15
        _ = v17
    }
    switch 4294967295 {
    case -725:
        v14 = true
    case -9223372036854775808:
        v14 = false
    default:
        if (("" + (v6 + v6)) >= v6) {
            v18 := GenS3{F0: 2147483647, F1: -1}
            _ = v18
        } else {
            v19 := []int{}
            _ = v19
        }
    }
    if (!v14) {
        var v20 float32 = v2
        _ = v20
    } else {
        for fi21 := 0; fi21 < 13; fi21++ {
            switch (fi21 - fi21) {
            case 1:
                for fi22 := 0; fi22 < 15; fi22++ {
                    var v23 float64 = -2.25
                    var v24 uint64 = 4294967295
                    v6 = ((v6 + (((v6 + v6) + "a") + v6)) + "foo bar")
                    v14 = (!false)
                    v6 = ("a" + v6)
                    if v14 {
                        fmt.Println(v6, v6, v14)
                    }
                    var v25 int16 = -1
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
            }
            var v26 uint8 = 2
            _ = fi21
            _ = v26
        }
        v3 = 739
    }
    if ((229 <= 0) || (!(true && (v14 || v14)))) {
        v3 = (v3 % ((v3) | 1))
        v27 := []int{}
        _ = v27
    } else {
        for fi28 := 0; fi28 < 20; fi28++ {
            switch (v6 + ("a" + ((("a" + "hello") + v6) + "hello"))) {
            case "hello":
                if (v6 == "") {
                    v4 -= 393
                    v4 = v4
                }
            case "a":
                if (410 <= 0) {
                    var v29 uint16 = 32767
                    fmt.Println(v5, v14, fi28)
                    _ = v29
                } else {
                    v30 := []int{255}
                    _ = v30
                }
            default:
                for fi31 := 0; fi31 < 14; fi31++ {
                    v32 := GenS2{F0: v14}
                    v1 -= (v1 / ((v5) | 1))
                    _ = fi31
                    _ = v32
                }
            }
            v6 = (("hello" + v6) + "hello")
            fmt.Println(v6, v6)
            _ = fi28
        }
        v33 := GenS4{F0: uint8(v1), F1: (^(v1)), F2: 10}
        _ = v33
    }
    v34 := []int{373}
    v1 += v1
    var v35 int16 = 0
    fmt.Println(v6)
    var v36 bool = v14
    if v36 {
        if (42 > -2147483648) {
            v37 := GenS5{F0: 304, F1: 4294967295, F2: v6}
            var v38 uint = 1
            v39 := GenS5{F0: 2147483647, F1: 1, F2: (v6 + ("" + v6))}
            _ = v37
            _ = v38
            _ = v39
        } else {
            if len(v34) > 0 {
                v34[0] = -261
            }
            var v40 int = -356
            _ = v40
        }
    }
    fmt.Println(v5, v36, v35)
    fmt.Println(v5, v36, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
