package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 uint32; F1 string }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    var v3 uint8 = 183
    if false {
        v4 := GenS1{F0: -938}
        v5 := []int{2147483647, 2147483647}
        _ = v4
        _ = v5
    }
    v3 += 1
    if (false || (("\n\t" + ("\n\t" + "\n\t")) >= "")) {
        fmt.Println(v3)
    } else {
        var v6 uint8 = 42
        for fi7 := 0; fi7 < 5; fi7++ {
            if (true || (false && false)) {
                v8 := []int{127, (fi7 - (fi7 | (fi7 ^ (fi7 & (fi7 - -280))))), fi7, 42}
                _ = v8
            }
            _ = fi7
        }
        fmt.Println(v3, v6)
        _ = v6
    }
    switch ("hello" + ("\n\t" + ((("a" + (("foo bar" + ("a" + "foo bar")) + ("a" + ("a" + "")))) + (("" + "a") + (("" + (("a" + "a") + "\n\t")) + "a"))) + ""))) {
    case "hello":
        if true {
            switch 0 {
            case 850:
                v9 := GenS2{F0: 9223372036854775807}
                _ = v9
            case -774:
                var v10 uint32 = 206
                _ = v10
            }
            v11 := GenS2{F0: -1}
            if true {
                for fi12 := 0; fi12 < 13; fi12++ {
                    v3 = 241
                    fmt.Println(v3, v3)
                    var v13 float64 = 1.5
                    _ = fi12
                    _ = v13
                }
                switch 32768 {
                case 0:
                    v14 := []int{4294967295, 937, -1, 127}
                    _ = v14
                case 127:
                    v2[42] = -2147483648
                case 434:
                    fmt.Println(v3)
                }
                v3 = 96
            } else {
                v15 := []int{-1, 42, 2147483647}
                _ = len(v15)
                _ = v15
            }
            _ = v11
        }
    default:
        if (!(!false)) {
            var v16 uint = 18446744073709551615
            v16 = (v16 | 255)
            _ = v16
        }
    }
    v17 := GenS3{F0: 42}
    var v18 uint8 = (^(v3))
    var v19 uint32 = 69
    v18 += v3
    if true {
        var v20 uint16 = 904
        for fi21 := 0; fi21 < 2; fi21++ {
            fmt.Println(v20, v3, v19)
            v3 += 1
            v22 := GenS2{F0: 437}
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            v18 += v3
            v18 = (v18 << ((1) & 63))
            v20 -= 0
            _ = fi23
        }
        _ = v20
    }
    v24 := []int{32767, 4294967295}
    if len(v24) > 0 {
        v24[0] = 127
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 uint = 2
        for fi27 := 0; fi27 < 11; fi27++ {
            if false {
                var v28 int = (fi25 * (fi27 / ((fi27) | 1)))
                _ = v28
            } else {
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
                v29 := GenS4{F0: v19, F1: int64(v26), F2: (v19 % ((v19) | 1))}
                v30 := GenS5{F0: v19, F1: (v3 * (v18 >> ((uint(v18)) & 63)))}
                _ = v29
                _ = v30
            }
            v31 := GenS6{F0: v19, F1: (("a" + ("foo bar" + "\n\t")) + "")}
            v18 = (v3 & 127)
            _ = fi27
            _ = v31
        }
        v18 -= (v18 * (v3 << ((4294967295) & 63)))
        _ = fi25
        _ = v26
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        fmt.Println(fi32)
        v33 := GenS6{F0: (^(v19)), F1: "\n\t"}
        v19 += 2147483648
        _ = fi32
        _ = v33
    }
    var v34 int32 = 0
    v19 = v19
    if len(v24) > 0 {
        v24[0] = 127
    }
    _ = len(v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v34
}
