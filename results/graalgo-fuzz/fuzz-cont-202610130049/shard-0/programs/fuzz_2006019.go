package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint16; F1 int64; F2 int16 }
type GenS5 struct { F0 int16; F1 uint16 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint32 = 42
    var v3 string = ("a" + "hello")
    var v4 bool = (true || (!(0 <= 100)))
    v5 := make(map[string]bool)
    var v6 int16 = 32767
    v3 = "\n\t"
    v3 = "\n\t"
    v7 := GenS1{F0: 699}
    v2 += v2
    v4 = (1 <= -561)
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(fi9, fi9)
            v10 := GenS2{F0: 255}
            if v4 {
                v5[v3] = v4
            } else {
                v11 := GenS3{F0: 255, F1: -273}
                fmt.Println(fi8, fi9, v6)
                v6 = v6
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        v12 := []int{}
        var v13 uint = 18446744073709551615
        _ = fi8
        _ = v12
        _ = v13
    }
    if true {
        var v14 string = v3
        fmt.Println(v14, v1)
        v15 := []int{0, 0, 860}
        _ = v14
        _ = v15
    }
    if v4 {
        for fi16 := 0; fi16 < 1; fi16++ {
            fmt.Println(v4, v2)
            if true {
                for fi17 := 0; fi17 < 6; fi17++ {
                    v18 := []int{-1, fi17, -9223372036854775808, fi16}
                    v19 := GenS4{F0: 255, F1: (v1 ^ 1), F2: (v6 + (v6 % ((v6) | 1)))}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            } else {
                if false {
                    var v20 int8 = 0
                    _ = v20
                } else {
                    v3 = v3
                    var v21 uint8 = 127
                    v6 += v6
                    _ = v21
                }
                var v22 uint32 = v2
                _ = v22
            }
            fmt.Println(v1, v6)
            _ = fi16
        }
        fmt.Println(v6, v4, v2)
        for fi23 := 0; fi23 < 11; fi23++ {
            v6 -= (^(v6))
            _ = fi23
        }
    } else {
        if (!(42 > 128)) {
            for fi24 := 0; fi24 < 20; fi24++ {
                _ = len(v5)
                v25 := GenS1{F0: 2}
                fmt.Println(v3)
                _ = fi24
                _ = v25
            }
            v4 = false
        }
    }
    if (-2.25 >= 0.1) {
        v26 := []int{}
        fmt.Println(v4)
        _ = v26
    } else {
        fmt.Println(v6, v1, v6)
        fmt.Println(v4, v2)
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        var v28 int16 = (v6 >> ((65535) & 63))
        switch "" {
        case "":
            v29 := GenS5{F0: -911, F1: 691}
            _ = v29
        }
        v30 := []int{}
        _ = fi27
        _ = v28
        _ = v30
    }
    v5[v3] = true
    v3 = (v3 + v3)
    for fi31 := 0; fi31 < 6; fi31++ {
        if (127 != 128) {
            for fi32 := 0; fi32 < 10; fi32++ {
                v2 += 42
                _ = fi32
            }
            fmt.Println(fi31, v6)
            for fi33 := 0; fi33 < 14; fi33++ {
                v3 = "\n\t"
                _ = fi33
            }
        }
        switch (fi31 ^ -451) {
        case 1:
            v34 := []int{fi31, 42, fi31, (fi31 / ((fi31) | 1))}
            _ = v34
        case 32767:
            delete(v5, (("hello" + "hello") + "foo bar"))
        case -812:
            v6 -= 42
        }
        _ = fi31
    }
    fmt.Println(v4, v6, v6)
    delete(v5, (("a" + (v3 + "")) + (("" + v3) + ("a" + (v3 + "")))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
