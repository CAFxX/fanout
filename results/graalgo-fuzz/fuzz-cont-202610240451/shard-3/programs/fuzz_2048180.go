package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int64; F1 int32 }
func genFn0(p0 uint, p1 uint64) int8 {
    p0 += (p0 & (^(p0)))
    switch ("hello" + "a") {
    case "":
        switch "foo bar" {
        case "foo bar":
            var v5 int32 = -1
            _ = v5
        }
    case "foo bar":
        p1 += (p1 * 2)
    case "hello":
        p0 = (^(p0))
    default:
        var v6 int16 = 1000
        _ = v6
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 uint, p1 uint32, p2 bool) int8 {
    var v7 float64 = 1.5
    var v8 string = "hello"
    v9 := []int{61, 128, 127, -550}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 31
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{42}
    v3 := []string{}
    var v4 uint = 255
    v1[979] = (true && (803 == 32768))
    var v10 float64 = 0.1
    switch (((("" + ("hello" + ("\n\t" + "a"))) + ("" + ("a" + ""))) + "foo bar") + "") {
    case "foo bar":
        if (-1 != 32768) {
            var v11 uint32 = 2147483648
            _ = v11
        } else {
            for fi12 := 0; fi12 < 19; fi12++ {
                switch (fi12 - fi12) {
                case -955:
                    if true {
                        fmt.Println(fi12, v4)
                        if false {
                            v13 := v1[42]
                            fmt.Println(v4, v10, fi12)
                            _ = v13
                        }
                        v4 = (v4 | v4)
                    } else {
                        var v14 uint32 = 847
                        _ = v14
                    }
                default:
                    for fi15 := 0; fi15 < 2; fi15++ {
                        var v16 int16 = -828
                        var v17 uint = (v4 % ((v4) | 1))
                        if (!(!(100 <= genFn0((^(v17)), 0)))) {
                            v4 = v4
                        }
                        if true {
                            var v18 uint32 = 2
                            fmt.Println(v10, fi12, v17)
                            v17 -= 1
                            _ = v18
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
                var v19 uint = (v4 + v4)
                v3 = append(v3, "\n\t")
                _ = fi12
                _ = v19
            }
            var v20 float32 = (3.14159 - -2.25)
            _ = v20
        }
    case "a":
        v4 = 423
    case "hello":
        fmt.Println(v10, v4)
    }
    v10 -= (0.0 - v10)
    v4 = (v4 & (v4 | (v4 & v4)))
    if true {
        var v21 uint8 = 42
        if false {
            switch -947 {
            case -2147483648:
                v1[4294967295] = false
            case 0:
                _ = len(v3)
            case 1:
                _ = len(v2)
            }
        }
        v22 := v1[2147483648]
        _ = v21
        _ = v22
    }
    v4 = (v4 + v4)
    delete(v1, 951)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "\n\t" {
    case "hello":
        if ((!(578 > 2147483647)) || (-1 < -345)) {
            switch 42 {
            case 1:
                v10 -= v10
            default:
                fmt.Println(v4, v4)
            }
            v23 := GenS1{F0: 2, F1: -2147483648}
            _ = v23
        }
    case "\n\t":
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := GenS2{F0: -822, F1: 1}
            var v26 bool = true
            _ = fi24
            _ = v25
            _ = v26
        }
    default:
        var v27 uint64 = 42
        _ = v27
    }
    var v28 int32 = 777
    var v29 int64 = 42
    var v30 int64 = (v29 & (^(v29)))
    v31 := v1[2]
    if (v28 < v28) {
        var v32 string = "foo bar"
        _ = v32
    } else {
        if true {
            v29 += -1
            v29 -= (^(v29))
        } else {
            var v33 int16 = 952
            v34 := GenS1{F0: 43, F1: (v28 / ((v28) | 1))}
            _ = v33
            _ = v34
        }
        _ = len(v1)
    }
    var v35 float32 = (1.5 + 1e10)
    var v36 uint64 = 0
    var v37 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
