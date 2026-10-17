package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 int64; F2 int64 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 int8, p1 int8, p2 bool) int16 {
    fmt.Println(p1)
    fmt.Println(p0)
    v5 := []int{}
    for fi6 := 0; fi6 < 9; fi6++ {
        p0 = 0
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}
func genFn1(p0 int16) int16 {
    switch "" {
    case "hello":
        p0 += 897
    case "foo bar":
        if (4294967295 <= 4294967295) {
            v7 := GenS1{F0: 255}
            _ = v7
        }
    case "\n\t":
        fmt.Println(p0)
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        switch ("\n\t" + "\n\t") {
        case "a":
            p0 += 1000
        case "\n\t":
            v9 := []int{}
            _ = v9
        case "":
            if ((-2.25 / 0.1) <= float32(((0.0 / 1.5) - -2.25))) {
                v10 := GenS2{F0: 43}
                v11 := []int{fi8, 457, fi8, fi8}
                v11 = append(v11, (fi8 ^ (^(fi8))))
                _ = v10
                _ = v11
            }
        }
        p0 -= -354
        _ = fi8
    }
    _ = p0
    return (p0 ^ p0)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -2147483648
    v3 := make(map[int]int)
    var v4 int32 = v2
    v12 := GenS3{F0: 0}
    if false {
        v2 = (v4 << ((107) & 63))
        if false {
            if (true && ((0.1 * 3.14159) != (0.1 / 1e10))) {
                v2 = v2
                var v13 int32 = -2147483648
                _ = v13
            }
        }
    } else {
        switch -2147483648 {
        case 127:
            if false {
                delete(v3, 42)
                var v14 uint64 = 1
                _ = v14
            }
        case 907:
            var v15 string = "hello"
            _ = v15
        case 4294967295:
            if (0 >= 682) {
                v16 := v1[2147483648]
                fmt.Println(v2)
                _ = v16
            }
        }
        v17 := v3[-168]
        _ = v17
    }
    v18 := GenS2{F0: -39}
    switch "a" {
    case "hello":
        if ((563 != 400) && true) {
            if true {
                v19 := []int{}
                _ = v19
            } else {
                v4 = (v4 & v2)
                fmt.Println(v4, v4)
            }
            v20 := GenS4{F0: 723, F1: 1}
            _ = v20
        } else {
            switch 32767 {
            case 128:
                var v21 uint32 = 1
                _ = v21
            default:
                fmt.Println(v4, v2)
            }
            for fi22 := 0; fi22 < 16; fi22++ {
                for fi23 := 0; fi23 < 6; fi23++ {
                    v24 := GenS5{F0: "a"}
                    var v25 bool = false
                    var v26 int16 = 32767
                    v27 := GenS2{F0: -1}
                    v28 := GenS5{F0: "a"}
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                switch fi22 {
                case -2147483648:
                    v29 := GenS6{F0: 1, F1: 613, F2: 44}
                    _ = v29
                default:
                    switch -556 {
                    case 4294967295:
                        fmt.Println(fi22, v2)
                    default:
                        v4 -= v4
                    }
                }
                v4 -= v2
                _ = fi22
            }
            for fi30 := 0; fi30 < 4; fi30++ {
                v2 = 2147483647
                var v31 int = fi30
                _ = fi30
                _ = v31
            }
        }
    case "a":
        var v32 uint16 = 387
        _ = v32
    case "":
        var v33 int8 = -1
        _ = v33
    default:
        v4 = (v2 + v4)
    }
    v34 := GenS7{F0: 255}
    v4 -= v4
    var v35 uint = 65535
    for fi36 := 0; fi36 < 3; fi36++ {
        switch ((("" + ("" + ("a" + (("hello" + "a") + "foo bar")))) + "hello") + "hello") {
        case "\n\t":
            v4 = (^(v2))
        case "foo bar":
            var v37 float64 = (0.1 + 0.0)
            _ = v37
        case "hello":
            var v38 int64 = 0
            _ = v38
        default:
            for fi39 := 0; fi39 < 8; fi39++ {
                fmt.Println(v35, v35, v4)
                _ = fi39
            }
        }
        var v40 string = "hello"
        _ = fi36
        _ = v40
    }
    var v41 float32 = 3.14159
    v4 = v2
    for fi42 := 0; fi42 < 5; fi42++ {
        v2 = (^(v4))
        _ = fi42
    }
    v1[585] = "a"
    v41 -= 0.1
    v43 := []int{}
    v4 -= (v4 + 673)
    fmt.Println(v4, v35, v35)
    v4 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v41
    _ = v43
}
