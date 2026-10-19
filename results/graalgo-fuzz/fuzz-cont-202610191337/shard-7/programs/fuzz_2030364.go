package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint; F2 string }
func genFn0(p0 int32, p1 uint64, p2 int) int64 {
    var v7 int = 638
    switch (v7 & (v7 / ((p2) | 1))) {
    case 128:
        switch ("hello" + ("" + ("hello" + "a"))) {
        case "foo bar":
            if false {
                for fi8 := 0; fi8 < 16; fi8++ {
                    v9 := []int{(-(v7)), (p2 << ((42) & 63)), p2}
                    _ = fi8
                    _ = v9
                }
            }
        }
    case -224:
        for fi10 := 0; fi10 < 9; fi10++ {
            for fi11 := 0; fi11 < 2; fi11++ {
                switch p2 {
                case -1:
                    switch p2 {
                    case 1:
                        switch ("\n\t" + "a") {
                        case "hello":
                            v12 := []int{}
                            _ = v12
                        case "foo bar":
                            v13 := GenS1{F0: 1}
                            _ = v13
                        }
                    case -784:
                        var v14 float32 = ((((1e10 - float32(0.1)) * 1e10) * 3.14159) + -2.25)
                        _ = v14
                    case 4294967295:
                        var v15 int8 = -1
                        _ = v15
                    }
                case 402:
                    var v16 int = v7
                    _ = v16
                case 127:
                    p1 -= 0
                }
                if (!(true || ((^(p1)) < 1))) {
                    v17 := GenS2{F0: 1, F1: (p0 ^ 23)}
                    v18 := []int{-9223372036854775808, (v7 ^ -1), p2}
                    _ = v17
                    _ = v18
                }
                _ = fi11
            }
            var v19 int16 = 0
            var v20 uint = 2
            _ = fi10
            _ = v19
            _ = v20
        }
    case 0:
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := []int{0}
            _ = fi21
            _ = v22
        }
    default:
        var v23 int16 = 1
        _ = v23
    }
    fmt.Println(p1, v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -654
}

func main() {
    var v1 uint = 255
    v2 := make(map[int]uint32)
    var v3 int = 32768
    v4 := make(map[string]int)
    var v5 uint64 = 859
    v6 := []int{42, v3, v3, (v3 - v3)}
    v24 := GenS3{F0: 0, F1: (^(v1)), F2: "hello"}
    v25 := GenS2{F0: -657, F1: -2147483648}
    delete(v4, ("hello" + "foo bar"))
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 int = int(v1)
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        delete(v4, "\n\t")
        switch (v3 / (((fi28 - (^(fi28)))) | 1)) {
        case -1:
            var v29 float64 = float64((-2.25 - (3.14159 - 0.0)))
            _ = v29
        }
        _ = fi28
    }
    var v30 uint8 = 42
    var v31 string = ("" + ((("hello" + "") + "a") + ("hello" + "\n\t")))
    for fi32 := 0; fi32 < 19; fi32++ {
        v33 := GenS2{F0: 42, F1: 42}
        fmt.Println(v30, v3, v1)
        _ = fi32
        _ = v33
    }
    if len(v6) > 0 {
        v6[0] = 9223372036854775807
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
