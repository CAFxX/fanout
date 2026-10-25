package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 bool; F1 uint16; F2 int64 }
func genFn0(p0 uint64, p1 int64) string {
    if false {
        for fi6 := 0; fi6 < 15; fi6++ {
            if true {
                var v7 uint16 = 919
                _ = v7
            }
            for fi8 := 0; fi8 < 17; fi8++ {
                p1 = (p1 >> ((18446744073709551615) & 63))
                v9 := []int{}
                _ = fi8
                _ = v9
            }
            switch (fi6 ^ (fi6 / ((32768) | 1))) {
            case -9223372036854775808:
                v10 := []int{128, (fi6 | 32767), 2147483647}
                _ = v10
            }
            _ = fi6
        }
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if true {
            v12 := GenS1{F0: 411}
            p0 = (p0 + (p0 / ((p0) | 1)))
            v13 := GenS2{F0: 2}
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    if (42 <= 41) {
        v14 := GenS3{F0: 18446744073709551615, F1: p0}
        if (3.14159 < float64((3.14159 + 0.0))) {
            v15 := []int{-46, -1}
            _ = v15
        } else {
            p0 = p0
            var v16 bool = false
            _ = v16
        }
        switch "foo bar" {
        case "hello":
            var v17 string = ("foo bar" + ("hello" + ""))
            _ = v17
        case "a":
            v18 := []int{}
            _ = v18
        default:
            var v19 uint16 = 0
            _ = v19
        }
        _ = v14
    } else {
        switch -352 {
        case 127:
            v20 := []int{42, 2147483647}
            _ = v20
        default:
            var v21 int64 = 9223372036854775807
            _ = v21
        }
        switch "" {
        case "\n\t":
            p0 = (p0 % (((^(p0))) | 1))
        case "foo bar":
            v22 := []int{-238, -1, 298, -9223372036854775808}
            _ = v22
        default:
            v23 := GenS2{F0: 1}
            _ = v23
        }
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return "foo bar"
}
func genFn1() bool {
    var v24 bool = false
    v24 = true
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := GenS4{F0: fi25, F1: "a"}
        _ = fi25
        _ = v26
    }
    var v27 uint = 42
    v28 := GenS3{F0: (v27 ^ 18446744073709551615), F1: 18446744073709551615}
    _ = v24
    _ = v27
    _ = v28
    return true
}

func main() {
    var v1 float32 = (0.1 + float32((float64((float32(1.5) + float32(float64(((float32(0.1) / 1.5) / 1.5))))) + -2.25)))
    var v2 int = -631
    v3 := make(map[string]int)
    v4 := []uint32{2147483647, 436, 2147483647, 392}
    v5 := make(map[string]uint32)
    var v29 string = "a"
    v30 := GenS4{F0: (v2 >> ((671) & 63)), F1: (v29 + v29)}
    for fi31 := 0; fi31 < 17; fi31++ {
        fmt.Println(v1)
        var v32 int64 = 9223372036854775807
        delete(v3, "hello")
        _ = fi31
        _ = v32
    }
    v33 := v5[("hello" + "\n\t")]
    var v34 float64 = (0.1 / 3.14159)
    for fi35 := 0; fi35 < 5; fi35++ {
        v36 := GenS3{F0: 4294967295, F1: 49}
        _ = fi35
        _ = v36
    }
    switch (("\n\t" + v29) + v29) {
    case "a":
        var v37 uint16 = 365
        _ = v37
    case "":
        if (!(v1 == (1e10 + v1))) {
            if false {
                if len(v4) > 0 {
                    v4[0] = (v33 & (v33 ^ (v33 % ((v33) | 1))))
                }
                v3["a"] = v2
            }
            switch "" {
            case "foo bar":
                var v38 int64 = 170
                _ = v38
            case "hello":
                v39 := GenS1{F0: 0}
                _ = v39
            case "a":
                v29 = "foo bar"
            default:
                v40 := GenS2{F0: 285}
                _ = v40
            }
            v41 := GenS4{F0: (v2 - (v2 << ((42) & 63))), F1: v29}
            _ = v41
        }
    }
    v29 = ("a" + "\n\t")
    if true {
        v42 := GenS3{F0: 1, F1: 4294967295}
        v43 := GenS3{F0: 2, F1: 1}
        if (654 <= uint64(v33)) {
            v1 = v1
            v44 := GenS1{F0: 255}
            v2 = 431
            _ = v44
        } else {
            v45 := GenS5{F0: (false || (true || true)), F1: 65535, F2: 9223372036854775807}
            _ = v45
        }
        _ = v42
        _ = v43
    } else {
        v33 = v33
        v46 := v5["foo bar"]
        _ = v46
    }
    var v47 int = -2147483648
    v29 = (((v29 + (v29 + "a")) + (v29 + "hello")) + ("hello" + v29))
    v47 -= v47
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v47 = v47
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
    _ = v47
    _ = v5
}
