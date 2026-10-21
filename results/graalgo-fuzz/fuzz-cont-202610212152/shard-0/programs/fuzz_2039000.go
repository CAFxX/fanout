package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 1
    var v3 int16 = -190
    switch "hello" {
    case "foo bar":
        v2 -= v2
    case "a":
        if false {
            v4 := GenS1{F0: 551}
            _ = v4
        }
    case "hello":
        fmt.Println(v3, v3)
    }
    var v5 uint = 129
    if false {
        v6 := GenS2{F0: -723}
        v7 := []int{-2147483648, -373}
        _ = v6
        _ = v7
    }
    v8 := GenS3{F0: 0}
    var v9 string = "\n\t"
    var v10 uint64 = 438
    for fi11 := 0; fi11 < 17; fi11++ {
        switch (fi11 % ((2147483647) | 1)) {
        case 128:
            if (0 >= 557) {
                for fi12 := 0; fi12 < 5; fi12++ {
                    if true {
                        fmt.Println(v2, v5, v2)
                    }
                    _ = fi12
                }
            } else {
                v13 := []int{127, fi11, 32768, 128}
                _ = v13
            }
        }
        _ = fi11
    }
    var v14 uint32 = v2
    v3 += -32768
    var v15 bool = true
    v2 = v14
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := GenS1{F0: v14}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        fmt.Println(v14, v2)
        _ = fi18
    }
    switch -2147483648 {
    case 42:
        var v19 uint = v5
        _ = v19
    case -845:
        for fi20 := 0; fi20 < 20; fi20++ {
            v5 += 2
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 uint64 = 1
                _ = fi21
                _ = v22
            }
            v3 -= (v3 << ((65535) & 63))
            _ = fi20
        }
    case 32767:
        switch ((v9 + (v9 + "foo bar")) + v9) {
        case "":
            var v23 int64 = 695
            _ = v23
        case "a":
            v14 += (^(v14))
        default:
            for fi24 := 0; fi24 < 16; fi24++ {
                v25 := GenS4{F0: v9, F1: 483}
                _ = fi24
                _ = v25
            }
        }
    }
    switch 128 {
    case 128:
        fmt.Println(v5, v5, v14)
    case 9223372036854775807:
        if false {
            for fi26 := 0; fi26 < 15; fi26++ {
                v27 := GenS3{F0: 37}
                if v15 {
                    switch 1 {
                    case -1:
                        switch ("foo bar" + "foo bar") {
                        case "a":
                            v28 := []int{fi26, (fi26 & fi26), fi26}
                            _ = v28
                        case "hello":
                            var v29 float32 = (0.0 / (float32(-2.25) + float32((-2.25 * 1.5))))
                            _ = v29
                        case "\n\t":
                            var v30 uint8 = 132
                            _ = v30
                        }
                    }
                    fmt.Println(v10)
                    v14 = v14
                } else {
                    v31 := GenS5{F0: 255, F1: 0}
                    _ = v31
                }
                for fi32 := 0; fi32 < 5; fi32++ {
                    var v33 uint32 = (v2 + v14)
                    _ = fi32
                    _ = v33
                }
                _ = fi26
                _ = v27
            }
            v34 := GenS4{F0: "", F1: 9223372036854775807}
            _ = v34
        }
    case -305:
        fmt.Println(v9, v14)
    default:
        v9 = v9
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        fmt.Println(v3)
        var v36 string = v9
        v37 := GenS1{F0: 333}
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v8
    _ = v9
}
