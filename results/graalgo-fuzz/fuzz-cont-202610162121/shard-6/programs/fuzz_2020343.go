package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 bool, p1 uint16, p2 int32) int8 {
    if false {
        v6 := GenS1{F0: 2, F1: 2147483647}
        p1 = p1
        v7 := []int{2147483647, 9223372036854775807}
        _ = v6
        _ = v7
    } else {
        var v8 uint = 96
        _ = v8
    }
    var v9 float64 = 1.5
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 float32 = (1.5 + 3.14159)
        _ = fi10
        _ = v11
    }
    v12 := []int{383, 9223372036854775807, 128, -9223372036854775808}
    p0 = ((p1 << ((130) & 63)) != 255)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v9
    return 127
}

func main() {
    v1 := []int{843, 32768, 91}
    v2 := []string{"hello"}
    v3 := make(map[string]bool)
    v4 := make(map[int]bool)
    var v5 int8 = 47
    switch 32768 {
    case -1:
        v4[-383] = false
    default:
        v5 -= v5
    }
    switch 2147483647 {
    case 783:
        switch ("foo bar" + "a") {
        case "a":
            v13 := GenS1{F0: 0, F1: 2147483648}
            _ = v13
        case "hello":
            v5 += (v5 % ((42) | 1))
        case "foo bar":
            var v14 uint16 = 866
            _ = v14
        default:
            if true {
                if (!(255 != 257)) {
                    v5 = v5
                } else {
                    if true {
                        v15 := GenS1{F0: 106, F1: 244}
                        _ = v15
                    }
                }
                if true {
                    v16 := GenS2{F0: 32768}
                    var v17 string = ("foo bar" + "")
                    _ = v16
                    _ = v17
                } else {
                    v18 := GenS3{F0: false}
                    v3[("\n\t" + "\n\t")] = true
                    _ = v18
                }
                var v19 int64 = 0
                _ = v19
            } else {
                v20 := GenS2{F0: -2147483648}
                _ = v20
            }
        }
    default:
        if (!((((((float64((float32(1.5) + 3.14159)) * (0.0 * (-2.25 + float64(0.0)))) + 1e10) + float64((1.5 / 1.0))) + float64(0.1)) * 0.0) >= (0.0 - ((0.1 * (-2.25 / 1.0)) + 0.1)))) {
            var v21 int32 = -1
            switch 654 {
            case 42:
                v22 := GenS3{F0: (!false)}
                _ = v22
            }
            fmt.Println(v5)
            _ = v21
        }
    }
    fmt.Println(v5)
    if (2 <= 4294967295) {
        var v23 uint8 = 60
        var v24 uint = 255
        _ = v23
        _ = v24
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 += (v5 | (^(v5)))
    v5 = v5
    if false {
        if (!false) {
            v25 := GenS4{F0: false}
            for fi26 := 0; fi26 < 11; fi26++ {
                var v27 uint64 = 777
                _ = fi26
                _ = v27
            }
            switch 128 {
            case -88:
                switch (("foo bar" + "hello") + "foo bar") {
                case "hello":
                    var v28 uint64 = 47
                    _ = v28
                }
            case 127:
                v29 := v3["hello"]
                _ = v29
            case -658:
                v5 = (-(v5))
            }
            _ = v25
        }
    }
    if (false || false) {
        fmt.Println(v5)
        if true {
            v30 := GenS5{F0: (v5 + v5), F1: -976}
            var v31 uint = 255
            switch "foo bar" {
            case "":
                switch "a" {
                case "hello":
                    _ = len(v1)
                case "a":
                    v31 = v31
                case "\n\t":
                    v32 := GenS3{F0: (0 > -2147483648)}
                    _ = v32
                }
            }
            _ = v30
            _ = v31
        }
    }
    v1 = append(v1, 42)
    var v33 string = "hello"
    switch 924 {
    case -457:
        v34 := GenS6{F0: 131}
        _ = v34
    case 994:
        if true {
            v4[2147483647] = true
        }
    default:
        v33 = "\n\t"
    }
    v33 = (v33 + "\n\t")
    var v35 float64 = (3.14159 / (3.14159 / 1.0))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((!false) && (true && ((!(765 >= 32767)) && true))) {
        v35 = v35
    } else {
        v36 := GenS2{F0: 128}
        switch (v33 + ((((v33 + ((("foo bar" + v33) + "foo bar") + v33)) + "a") + (("a" + "hello") + v33)) + ((v33 + "a") + v33))) {
        case "":
            v5 += (^(v5))
        default:
            v35 = v35
        }
        _ = v36
    }
    switch 647 {
    case 4294967295:
        var v37 uint64 = 715
        _ = v37
    case -2147483648:
        var v38 uint32 = 160
        _ = v38
    case 32768:
        fmt.Println(v35, v5)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v4
    _ = v5
}
