package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint64; F1 uint32; F2 bool }
type GenS3 struct { F0 int64 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := []int{(fi4 % (((fi4 >> ((990) & 63))) | 1))}
        v6 := GenS1{F0: 381, F1: 1}
        var v7 uint32 = 0
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 float32 = 1e10
    for fi9 := 0; fi9 < 8; fi9++ {
        v8 -= (v8 - (1e10 - v8))
        var v10 uint8 = 128
        switch "" {
        case "\n\t":
            switch 342 {
            case 9223372036854775807:
                v11 := []int{fi9, (fi9 << ((4294967295) & 63)), 4294967295}
                _ = v11
            case -1:
                switch (fi9 >> ((65535) & 63)) {
                case -2147483648:
                    v10 = v10
                case 42:
                    v10 -= 128
                case 2147483647:
                    v12 := []int{(fi9 | fi9)}
                    _ = v12
                default:
                    var v13 bool = false
                    _ = v13
                }
            }
        case "":
            switch (("foo bar" + "") + "foo bar") {
            case "\n\t":
                v8 = 3.14159
            }
        case "a":
            v14 := GenS2{F0: 1, F1: 634, F2: false}
            _ = v14
        default:
            v8 += 0.0
        }
        _ = fi9
        _ = v10
    }
    var v15 uint = 2
    _ = v15
    _ = v8
    return 127
}
func genFn1() string {
    var v16 uint8 = 2
    var v17 string = ("\n\t" + "\n\t")
    var v18 uint16 = 65535
    _ = v16
    _ = v17
    _ = v18
    return "hello"
}

func main() {
    var v1 uint = 65535
    var v2 uint8 = 128
    var v3 int8 = -128
    var v19 float64 = (0.0 * float64(float32(float64(0.1))))
    v1 = (v1 + (v1 + 926))
    for fi20 := 0; fi20 < 10; fi20++ {
        v21 := []int{}
        _ = fi20
        _ = v21
    }
    v22 := []int{127, -2147483648}
    v23 := GenS3{F0: 42}
    for fi24 := 0; fi24 < 13; fi24++ {
        switch 9223372036854775807 {
        case 1:
            fmt.Println(v19, v2)
        case 9223372036854775807:
            var v25 string = "foo bar"
            _ = v25
        case 2147483647:
            var v26 bool = true
            _ = v26
        default:
            v3 += 42
        }
        switch ((("hello" + "hello") + "foo bar") + ("a" + "\n\t")) {
        case "":
            v27 := GenS3{F0: 1}
            _ = v27
        case "foo bar":
            v19 -= v19
        default:
            if true {
                var v28 int16 = 1000
                if false {
                    var v29 int32 = int32(v3)
                    _ = v29
                }
                _ = v28
            }
        }
        var v30 uint16 = 65535
        _ = fi24
        _ = v30
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        for fi32 := 0; fi32 < 1; fi32++ {
            var v33 uint = v1
            _ = fi32
            _ = v33
        }
        if (491 >= 42) {
            v2 -= (v2 ^ 2)
            if (v19 == v19) {
                v2 = (v2 * (v2 << (((^(v1))) & 63)))
            } else {
                for fi34 := 0; fi34 < 9; fi34++ {
                    if len(v22) > 0 {
                        v22[0] = (fi31 | -372)
                    }
                    v35 := GenS2{F0: 49, F1: 4294967295, F2: false}
                    _ = fi34
                    _ = v35
                }
            }
            v3 -= genFn0()
        }
        v1 -= (v1 >> (((^(v1))) & 63))
        _ = fi31
    }
    v36 := GenS2{F0: 826, F1: 2, F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v36
}
