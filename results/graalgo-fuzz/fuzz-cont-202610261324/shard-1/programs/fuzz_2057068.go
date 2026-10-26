package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() uint32 {
    if true {
        var v5 float32 = 1e10
        _ = v5
    }
    if true {
        var v6 int16 = -334
        _ = v6
    } else {
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 uint = 18446744073709551615
            v8 -= (^(v8))
            for fi9 := 0; fi9 < 13; fi9++ {
                if true {
                    var v10 uint = v8
                    switch "foo bar" {
                    case "":
                        v11 := []int{(fi9 >> ((v10) & 63)), (^(fi9))}
                        _ = v11
                    case "hello":
                        var v12 int64 = -700
                        _ = v12
                    }
                    _ = v10
                }
                var v13 uint8 = 128
                switch fi7 {
                case -9223372036854775808:
                    if ((fi7 | fi7) < (fi7 | (fi7 & (fi9 + fi9)))) {
                        switch (fi9 | 128) {
                        case 895:
                            if true {
                                v14 := []int{-2147483648}
                                v8 = v8
                                _ = v14
                            }
                        case 127:
                            v8 = (v8 / ((42) | 1))
                            v15 := []int{32767, fi9}
                            _ = v15
                        case 2147483647:
                            v13 = 120
                        }
                    }
                }
                _ = fi9
                _ = v13
            }
            _ = fi7
            _ = v8
        }
    }
    if (1 >= -1) {
        switch 128 {
        case -495:
            var v16 int64 = -9223372036854775808
            _ = v16
        case 658:
            v17 := GenS1{F0: 42}
            _ = v17
        case 4294967295:
            v18 := GenS1{F0: 0}
            _ = v18
        }
        var v19 int64 = 42
        _ = v19
    }
    var v20 uint = 18446744073709551615
    _ = v20
    return 1
}
func genFn1() uint32 {
    v21 := []int{-9223372036854775808, 32768}
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := GenS1{F0: 42}
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
        _ = fi22
        _ = v23
    }
    var v24 int64 = 0
    var v25 int64 = -1
    _ = v21
    _ = v24
    _ = v25
    return 2
}

func main() {
    var v1 uint8 = 206
    var v2 int = 32767
    var v3 bool = (0 < 0)
    v4 := make(map[int]int)
    v3 = ((!((!v3) && v3)) || false)
    fmt.Println(v3, v3)
    var v26 bool = false
    if v26 {
        v26 = true
    }
    var v27 uint = 924
    v27 = v27
    v2 += v2
    if (42 > 89) {
        if (true || ((0.0 != -2.25) && (true && (42 < 546)))) {
            var v28 uint = (v27 - (^(v27)))
            v29 := GenS1{F0: 1}
            _ = v28
            _ = v29
        }
        v30 := GenS1{F0: 757}
        var v31 int16 = -32768
        _ = v30
        _ = v31
    }
    var v32 int8 = 1
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v4
}
