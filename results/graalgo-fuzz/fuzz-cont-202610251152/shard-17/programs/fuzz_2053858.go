package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 string }
func genFn0(p0 bool, p1 string, p2 uint) string {
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 42}
        fmt.Println(p2, p0, p0)
        _ = fi7
        _ = v8
    }
    v9 := []int{0, 4294967295, 597, 128}
    if p0 {
        p1 = "foo bar"
        var v10 uint8 = 4
        var v11 uint32 = 2147483647
        _ = v10
        _ = v11
    } else {
        if (true && p0) {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p1
}
func genFn1(p0 uint8, p1 string, p2 uint32) uint {
    p2 += p2
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 uint16 = 482
        v14 := []int{fi12, (^(fi12)), fi12}
        v15 := GenS2{F0: 42, F1: (-(fi12))}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    if (0 > -9223372036854775808) {
        p0 = p0
    }
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return 399
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 2
    var v3 int32 = -827
    v4 := make(map[uint32]string)
    var v5 uint16 = 0
    var v6 int64 = -612
    v5 = v5
    v16 := GenS1{F0: v2}
    _ = len(v4)
    v6 = (-(v6))
    fmt.Println(v5, v5)
    if (!false) {
        fmt.Println(v3)
        v3 = v3
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            for fi18 := 0; fi18 < 11; fi18++ {
                v2 += 255
                v19 := GenS2{F0: 18446744073709551615, F1: -303}
                if false {
                    switch (fi17 / (((-(fi18))) | 1)) {
                    case -37:
                        v20 := GenS3{F0: ""}
                        _ = v20
                    }
                    fmt.Println(fi18)
                }
                _ = fi18
                _ = v19
            }
            for fi21 := 0; fi21 < 15; fi21++ {
                v22 := GenS2{F0: 2, F1: 0}
                _ = fi21
                _ = v22
            }
            _ = fi17
        }
    }
    v3 -= v3
    v5 = (v5 + (^(v5)))
    v2 = (^(v2))
    var v23 bool = (!(1 > 288))
    var v24 bool = true
    v6 -= (-(v6))
    v5 += v5
    v4[4294967295] = "hello"
    if v24 {
        var v25 int16 = 42
        var v26 string = "\n\t"
        v5 -= 0
        _ = v25
        _ = v26
    } else {
        fmt.Println(v3, v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
