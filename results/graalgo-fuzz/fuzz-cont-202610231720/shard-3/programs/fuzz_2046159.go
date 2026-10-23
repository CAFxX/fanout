package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 bool; F1 bool }
func genFn0() uint64 {
    var v5 uint8 = 42
    for fi6 := 0; fi6 < 10; fi6++ {
        if true {
            v7 := GenS1{F0: 1, F1: false}
            var v8 uint16 = 647
            _ = v7
            _ = v8
        } else {
            var v9 bool = true
            var v10 uint = 0
            var v11 uint32 = 511
            _ = v10
            _ = v11
            _ = v9
        }
        if (42 < uint32(v5)) {
            switch 0 {
            case -613:
                switch "hello" {
                case "\n\t":
                    v12 := GenS2{F0: "hello", F1: 1}
                    _ = v12
                }
            case 32768:
                switch fi6 {
                case 32767:
                    fmt.Println(fi6, v5)
                }
            case 1:
                v13 := []int{}
                _ = v13
            }
            v5 = (v5 | (^(v5)))
        }
        _ = fi6
    }
    if false {
        v5 += (v5 + v5)
        var v14 uint = 18446744073709551615
        _ = v14
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := []int{int(v5), fi15, (-(fi15)), (fi15 + 127)}
        v17 := GenS1{F0: -32768, F1: false}
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = v5
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    var v3 bool = (!true)
    v4 := []string{"foo bar", "\n\t", ""}
    for fi18 := 0; fi18 < 18; fi18++ {
        v4 = append(v4, "hello")
        fmt.Println(fi18, fi18)
        if v3 {
            v2[fi18] = (("\n\t" + "hello") <= "foo bar")
        }
        _ = fi18
    }
    v3 = (-1 > -102)
    v3 = v3
    if v3 {
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 string = ("\n\t" + "foo bar")
            var v21 int16 = 32767
            delete(v2, (fi19 >> ((534) & 63)))
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 uint = 4294967295
        v3 = false
        _ = v22
    }
    var v23 float64 = ((1.5 - 1.5) + 0.1)
    v4 = append(v4, "hello")
    v3 = v3
    v24 := GenS2{F0: (("a" + "a") + "\n\t"), F1: 0}
    var v25 uint32 = 2
    _ = len(v1)
    var v26 int16 = 1000
    v3 = (456 != 799)
    if false {
        var v27 uint64 = 767
        _ = v27
    } else {
        for fi28 := 0; fi28 < 8; fi28++ {
            switch fi28 {
            case -939:
                _ = len(v2)
            case -9223372036854775808:
                v26 = (v26 * v26)
            case -122:
                fmt.Println(v3, v26, v26)
            }
            _ = fi28
        }
    }
    _ = len(v4)
    switch 16 {
    case 4294967295:
        switch 32767 {
        case 128:
            for fi29 := 0; fi29 < 17; fi29++ {
                var v30 int8 = 42
                v31 := v1[(^(v25))]
                _ = fi29
                _ = v30
                _ = v31
            }
        default:
            var v32 bool = (v3 || (1 >= 2147483648))
            _ = v32
        }
    case -9223372036854775808:
        v3 = v3
    case -94:
        for fi33 := 0; fi33 < 14; fi33++ {
            v4 = append(v4, ("\n\t" + "hello"))
            if (((1e10 - v23) <= v23) || (("" == "a") || v3)) {
                switch (fi33 & fi33) {
                case -1:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            } else {
                delete(v1, v25)
            }
            _ = fi33
        }
    default:
        v34 := GenS3{F0: v3, F1: false}
        _ = v34
    }
    var v35 float64 = v23
    v25 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v4
}
