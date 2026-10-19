package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint64, p1 int64, p2 uint64) int8 {
    var v7 float64 = 0.0
    var v8 float64 = (((v7 + v7) - v7) / ((v7 / v7) - v7))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 bool) uint {
    v9 := GenS1{F0: "hello"}
    var v10 uint32 = 0
    v10 -= v10
    fmt.Println(v10, p0)
    _ = p0
    _ = v10
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 0
    var v2 uint = 127
    var v3 uint32 = 2147483648
    var v4 float32 = 0.1
    var v5 bool = false
    var v6 uint32 = (^(v3))
    if v5 {
        for fi11 := 0; fi11 < 9; fi11++ {
            v1 = 956
            _ = fi11
        }
    } else {
        switch 42 {
        case 42:
            v6 = (v6 / ((135) | 1))
        default:
            v2 -= (v2 ^ 1)
        }
        if ("\n\t" != "hello") {
            v4 = 0.0
        }
    }
    if true {
        var v12 int = -9223372036854775808
        _ = v12
    }
    switch "foo bar" {
    case "foo bar":
        var v13 bool = ((v6 << ((v2) & 63)) <= (v3 % ((v6) | 1)))
        _ = v13
    case "\n\t":
        v14 := []int{0, -1}
        _ = v14
    }
    var v15 int8 = 1
    for fi16 := 0; fi16 < 16; fi16++ {
        if (746 < v6) {
            var v17 uint16 = 65535
            var v18 uint64 = (v1 - v1)
            _ = v17
            _ = v18
        }
        _ = fi16
    }
    var v19 string = "\n\t"
    var v20 int32 = 0
    var v21 int16 = 42
    var v22 int16 = v21
    v23 := GenS1{F0: "foo bar"}
    var v24 int32 = 0
    var v25 uint64 = (v1 - (v1 >> (((v2 | (v2 + v2))) & 63)))
    fmt.Println(v22, v25)
    for fi26 := 0; fi26 < 20; fi26++ {
        switch fi26 {
        case 1:
            v27 := []int{(fi26 % (((^(fi26))) | 1)), 0, fi26, (fi26 / (((fi26 << (((v2 % (((v2 * genFn1(((v2 >> ((v2) & 63)) == genFn1((4294967295 < 42)))))) | 1))) & 63))) | 1))}
            _ = v27
        case 711:
            if false {
                if v5 {
                    switch (^(fi26)) {
                    case 1:
                        v20 = v20
                    }
                    var v28 uint = (v2 ^ (v2 ^ v2))
                    for fi29 := 0; fi29 < 17; fi29++ {
                        v30 := []int{fi29, fi26, fi29}
                        var v31 uint16 = 495
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = v28
                }
                v32 := GenS1{F0: v19}
                v25 = (v1 / (((v1 - v1)) | 1))
                _ = v32
            }
        case 75:
            for fi33 := 0; fi33 < 1; fi33++ {
                fmt.Println(v22, v15, fi26)
                var v34 int32 = v20
                _ = fi33
                _ = v34
            }
        }
        for fi35 := 0; fi35 < 3; fi35++ {
            var v36 float64 = ((((3.14159 * 0.0) + 0.1) * -2.25) * (-2.25 - ((-2.25 * (0.0 + 3.14159)) * (1.5 - 1.5))))
            _ = fi35
            _ = v36
        }
        _ = fi26
    }
    switch -833 {
    case 2147483647:
        var v37 uint64 = v25
        _ = v37
    default:
        fmt.Println(v24)
    }
    fmt.Println(v20)
    v38 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
