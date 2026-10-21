package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint }
type GenS2 struct { F0 uint; F1 int64; F2 uint32 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint }
type GenS4 struct { F0 uint }
func genFn0(p0 int8) string {
    v7 := []int{176, 32768, 2147483647, -161}
    var v8 int32 = -815
    p0 = (p0 | (p0 >> ((320) & 63)))
    var v9 uint = 1
    fmt.Println(v8, v9, v9)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return ("\n\t" + ("" + ("a" + "foo bar")))
}
func genFn1(p0 int32, p1 int64) int8 {
    v10 := GenS1{F0: ("" + "foo bar"), F1: 2, F2: 0}
    p0 = 42
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    var v1 int = 1
    var v2 int64 = 0
    var v3 uint16 = 6
    var v4 float64 = 3.14159
    var v5 int16 = 832
    v6 := make(map[string]bool)
    v1 = v1
    v4 = 3.14159
    fmt.Println(v4, v5, v4)
    switch (v1 * v1) {
    case 128:
        var v11 int8 = 31
        _ = v11
    }
    v12 := GenS2{F0: 265, F1: v2, F2: 147}
    var v13 uint32 = 2147483647
    v14 := GenS1{F0: "hello", F1: 0, F2: 313}
    var v15 string = "a"
    for fi16 := 0; fi16 < 13; fi16++ {
        v1 = -1
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 int16 = v5
            var v19 int16 = v5
            v20 := GenS1{F0: ("a" + v15), F1: 834, F2: 4294967295}
            var v21 bool = true
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
            _ = v21
        }
        _ = fi16
    }
    if true {
        switch v15 {
        case "hello":
            if false {
                v22 := []int{4294967295, v1, (v1 - (v1 / ((int(v13)) | 1))), v1}
                v23 := GenS3{F0: int32(v2), F1: 18446744073709551615, F2: 4294967295}
                _ = v22
                _ = v23
            }
        case "a":
            var v24 uint8 = 65
            _ = v24
        }
        v4 += float64(0.0)
    }
    var v25 float64 = v4
    if false {
        fmt.Println(v2, v4, v15)
        var v26 uint64 = 197
        fmt.Println(v3, v26)
        _ = v26
    } else {
        for fi27 := 0; fi27 < 10; fi27++ {
            switch v15 {
            case "a":
                var v28 string = (("foo bar" + (("a" + v15) + v15)) + (((v15 + (("hello" + "") + "\n\t")) + "foo bar") + v15))
                _ = v28
            }
            _ = fi27
        }
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 uint64 = 518
            _ = fi29
            _ = v30
        }
    }
    switch (v15 + (v15 + ((("hello" + ("\n\t" + (v15 + (v15 + "hello")))) + v15) + "\n\t"))) {
    case "a":
        switch (v1 | 32768) {
        case 9223372036854775807:
            var v31 int8 = 1
            _ = v31
        }
    default:
        for fi32 := 0; fi32 < 1; fi32++ {
            for fi33 := 0; fi33 < 16; fi33++ {
                if ((430 <= 1) && ((!false) || (v15 != v15))) {
                    if true {
                        var v34 string = (v15 + "foo bar")
                        var v35 string = "foo bar"
                        _ = v34
                        _ = v35
                    } else {
                        v36 := []int{-1, v1, (fi33 - v1), (v1 + (^(fi32)))}
                        var v37 uint16 = (^(v3))
                        v38 := GenS4{F0: 255}
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    var v39 string = "\n\t"
                    _ = v39
                }
                _ = fi33
            }
            var v40 uint32 = v13
            _ = fi32
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
