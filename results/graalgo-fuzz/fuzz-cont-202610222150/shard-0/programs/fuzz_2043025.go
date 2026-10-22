package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint64) uint16 {
    switch 1 {
    case 32767:
        v5 := []int{}
        _ = v5
    case 854:
        fmt.Println(p0)
    case 2147483647:
        if false {
            p0 = p0
        }
    default:
        v6 := []int{9223372036854775807}
        _ = v6
    }
    v7 := GenS1{F0: p0, F1: -1, F2: 2}
    p0 = 1
    _ = p0
    _ = v7
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -32768
    var v3 uint16 = 2
    v4 := make(map[string]string)
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint32 = 1
        var v10 float64 = -2.25
        if true {
            var v11 string = "foo bar"
            if (true || ((!false) && true)) {
                v12 := GenS2{F0: 1, F1: (-1 < -128)}
                _ = v12
            }
            _ = v11
        } else {
            v9 -= (v9 << ((2) & 63))
            if true {
                v13 := GenS1{F0: 42, F1: 109, F2: 1}
                v3 += 762
                v14 := []int{}
                _ = v13
                _ = v14
            } else {
                v15 := v1["a"]
                var v16 float64 = v10
                _ = v15
                _ = v16
            }
            if true {
                v17 := []int{(fi8 + -1)}
                fmt.Println(v2)
                _ = v17
            }
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v18 uint = 468
    v3 += v3
    var v19 uint32 = 787
    var v20 int64 = 553
    v2 = v2
    for fi21 := 0; fi21 < 14; fi21++ {
        if (!false) {
            fmt.Println(v20)
            if (!((0.0 / -2.25) <= ((-2.25 / -2.25) + 0.1))) {
                switch "hello" {
                case "foo bar":
                    var v22 string = "hello"
                    _ = v22
                case "hello":
                    if (true && false) {
                        fmt.Println(v20, v20, v19)
                        switch ("\n\t" + "\n\t") {
                        case "hello":
                            v23 := GenS2{F0: 27, F1: ((2 >= 4294967295) || false)}
                            _ = v23
                        }
                        fmt.Println(v20, v2)
                    } else {
                        if true {
                            v19 += 2147483647
                            v19 = (v19 ^ uint32(v18))
                            var v24 bool = true
                            _ = v24
                        }
                        v18 += 42
                        v20 -= 600
                    }
                case "a":
                    v25 := GenS3{F0: 688, F1: (v19 - (v19 / (((v19 / (((v19 >> ((2) & 63))) | 1))) | 1))), F2: 1000}
                    _ = v25
                }
            }
            for fi26 := 0; fi26 < 20; fi26++ {
                if false {
                    var v27 uint32 = v19
                    _ = v27
                }
                _ = fi26
            }
        }
        for fi28 := 0; fi28 < 5; fi28++ {
            v29 := GenS2{F0: 245, F1: true}
            v30 := GenS4{F0: v3, F1: (v19 + 1)}
            _ = fi28
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 1; fi31++ {
            var v32 uint16 = v3
            _ = fi31
            _ = v32
        }
        _ = fi21
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        for fi34 := 0; fi34 < 20; fi34++ {
            if (true || (v18 > v18)) {
                var v35 bool = (v18 < 65535)
                _ = v35
            } else {
                fmt.Println(fi33, v18, fi34)
                v2 = 32767
                var v36 float32 = 3.14159
                v37 := GenS5{F0: -128}
                fmt.Println(v36, fi34, v20)
                _ = v36
                _ = v37
            }
            _ = fi34
        }
        for fi38 := 0; fi38 < 10; fi38++ {
            v39 := []int{-1, (-(fi33)), -362}
            v1["hello"] = ("hello" + "hello")
            v19 -= (v19 & (v19 - (v19 * 65535)))
            switch fi38 {
            case -2147483648:
                _ = len(v39)
            case -213:
                v3 += genFn0(552)
            case -875:
                v2 = 42
            }
            if len(v39) >= 2 {
                v39 = v39[1:len(v39)]
            }
            var v40 float64 = (-2.25 * 3.14159)
            v19 += v19
            _ = fi38
            _ = v39
            _ = v40
        }
        _ = fi33
    }
    v41 := GenS1{F0: 733, F1: 15, F2: 128}
    _ = len(v1)
    v42 := GenS5{F0: -128}
    var v43 int = 4294967295
    v44 := v4[("foo bar" + ("hello" + ""))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
