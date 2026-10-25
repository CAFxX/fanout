package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint32, p1 string, p2 string) int8 {
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(p0, p0)
        _ = fi6
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 float32 = ((-2.25 * float32(float64(-2.25))) + float32((1e10 + 0.0)))
            v10 := []int{32767, (fi7 | fi8), 127, fi7}
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := []int{(fi7 - fi7), fi7, fi7, -1}
        _ = fi7
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 int64 = 435
    v2 := []uint32{0, 65535, 933}
    v3 := make(map[int]string)
    var v4 int16 = 932
    v5 := make(map[int]bool)
    fmt.Println(v4)
    switch 32768 {
    case 9223372036854775807:
        v1 = v1
    default:
        for fi12 := 0; fi12 < 2; fi12++ {
            fmt.Println(fi12)
            v1 = (v1 ^ (v1 - (^(v1))))
            _ = fi12
        }
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v3[fi13] = ""
        var v14 int32 = 1
        v14 += (v14 << ((329) & 63))
        _ = fi13
        _ = v14
    }
    if false {
        switch ("\n\t" + ("\n\t" + "hello")) {
        case "":
            for fi15 := 0; fi15 < 3; fi15++ {
                switch (("foo bar" + "foo bar") + "hello") {
                case "a":
                    if false {
                        delete(v3, (-(fi15)))
                    }
                }
                v16 := GenS1{F0: 32767, F1: 0}
                _ = fi15
                _ = v16
            }
        }
        var v17 int32 = 2147483647
        _ = v17
    } else {
        v18 := GenS2{F0: false, F1: (v1 + (v1 << ((255) & 63)))}
        _ = v18
    }
    fmt.Println(v4, v4)
    v19 := GenS2{F0: (32767 > 2), F1: 998}
    v2 = append(v2, 814)
    if (true && false) {
        var v20 uint16 = 2
        v21 := GenS3{F0: (v4 ^ v4)}
        var v22 uint16 = 65535
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 int = 1
    if (0 >= 42) {
        var v24 uint64 = 332
        _ = v24
    } else {
        var v25 int = v23
        if ("" < "\n\t") {
            v23 += v23
            fmt.Println(v4, v25, v25)
            v3[42] = ("foo bar" + "hello")
        } else {
            var v26 int32 = -2147483648
            if (v1 > int64(v4)) {
                var v27 uint16 = 74
                v28 := GenS4{F0: 2147483648}
                var v29 uint = 0
                _ = v27
                _ = v28
                _ = v29
            }
            var v30 uint8 = 42
            _ = v26
            _ = v30
        }
        var v31 uint32 = 2147483647
        _ = v25
        _ = v31
    }
    switch "" {
    case "hello":
        v32 := GenS4{F0: 4294967295}
        _ = v32
    case "":
        var v33 uint16 = 2
        _ = v33
    case "a":
        v34 := GenS1{F0: 0, F1: -735}
        _ = v34
    default:
        var v35 int16 = (^(v4))
        _ = v35
    }
    var v36 string = "a"
    fmt.Println(v36, v23, v36)
    if ((v1 + (v1 + -9223372036854775808)) >= (v1 << ((1000) & 63))) {
        v36 = v36
        v37 := v3[(^(v23))]
        _ = v37
    }
    var v38 uint = 1
    if true {
        v39 := GenS1{F0: 65535, F1: -178}
        fmt.Println(v23)
        v40 := GenS5{F0: 1, F1: 42}
        _ = v39
        _ = v40
    }
    fmt.Println(v1, v23, v4)
    var v41 uint32 = 827
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v38
    _ = v4
    _ = v41
    _ = v5
}
