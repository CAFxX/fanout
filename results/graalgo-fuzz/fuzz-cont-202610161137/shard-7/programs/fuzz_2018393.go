package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint32; F2 int16 }
type GenS3 struct { F0 int; F1 uint; F2 uint }

func main() {
    var v1 bool = false
    var v2 float64 = 1.5
    var v3 bool = true
    var v4 bool = (!v1)
    var v5 uint = 367
    if (692 < 17) {
        if v4 {
            v6 := GenS1{F0: v5, F1: 18446744073709551615}
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := []int{}
                if (v2 < v2) {
                    v5 += v5
                    if (!false) {
                        v5 -= (v5 % (((^(v5))) | 1))
                        switch "hello" {
                        case "hello":
                            v3 = ((fi7 + fi7) >= fi7)
                        case "a":
                            fmt.Println(v1)
                        default:
                            _ = len(v8)
                        }
                    } else {
                        var v9 int8 = 89
                        _ = v9
                    }
                    if (true || ((-2.25 + 1.5) > (((v2 + 1e10) * float64(0.0)) - float64(0.0)))) {
                        v4 = false
                    }
                }
                _ = fi7
                _ = v8
            }
            _ = v6
        } else {
            v5 += (v5 >> ((898) & 63))
            var v10 float64 = v2
            v10 -= float64(float32(1.5))
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 float64 = (v2 * 0.0)
        v1 = true
        switch (fi11 >> (((v5 % ((v5) | 1))) & 63)) {
        case -340:
            v12 -= (v12 - v12)
        }
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        switch "" {
        case "a":
            var v14 int = fi13
            _ = v14
        case "foo bar":
            v15 := []int{42, 4294967295, (fi13 % ((fi13) | 1)), (fi13 % ((-1) | 1))}
            _ = v15
        }
        var v16 int32 = 0
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := GenS2{F0: (^(v5)), F1: 65535, F2: 42}
            if false {
                v19 := []int{(fi17 | (fi13 | fi13)), 2147483647}
                var v20 int32 = 0
                v16 = (v16 + v20)
                _ = v19
                _ = v20
            }
            _ = fi17
            _ = v18
        }
        _ = fi13
        _ = v16
    }
    var v21 int64 = 42
    if (v5 >= v5) {
        v22 := []int{}
        _ = v22
    }
    var v23 string = ("\n\t" + "a")
    v23 = "a"
    v21 -= (v21 >> (((v5 ^ 255)) & 63))
    switch v23 {
    case "hello":
        if true {
            switch (v23 + ("\n\t" + v23)) {
            case "":
                var v24 uint32 = 2147483648
                _ = v24
            case "\n\t":
                if true {
                    v23 = "\n\t"
                    if v4 {
                        v25 := GenS2{F0: 1, F1: 748, F2: 42}
                        _ = v25
                    } else {
                        fmt.Println(v21, v1)
                        v26 := GenS1{F0: (^(v5)), F1: 0}
                        var v27 int16 = 873
                        _ = v26
                        _ = v27
                    }
                } else {
                    v28 := []int{-620, -236, 786}
                    _ = v28
                }
            }
            v29 := []int{42, 325, 128}
            _ = v29
        }
    case "foo bar":
        v5 = v5
    case "":
        v3 = true
    }
    v30 := GenS2{F0: v5, F1: 1, F2: 1}
    for fi31 := 0; fi31 < 14; fi31++ {
        v32 := GenS3{F0: fi31, F1: v5, F2: (v5 + 0)}
        v33 := []int{(fi31 + (-(fi31))), -525}
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
